// Copyright (c) 2020-2025 Tigera, Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package intdataplane

import (
	"errors"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"

	"github.com/projectcalico/calico/felix/bpf/events"
)

var (
	bpfEventsCounters = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "felix_bpf_events",
		Help: "Number of events generated by BPF dataplane split by type/category",
	}, []string{"type"})
)

func init() {
	prometheus.MustRegister(bpfEventsCounters)
}

type bpfEventHandler func(e events.Event)

type bpfEventSink struct {
	handlers []bpfEventHandler
	counter  prometheus.Counter
}

type bpfEventPoller struct {
	events events.Events
	sinks  map[events.Type]bpfEventSink

	prometheusLostCounter   prometheus.Counter
	prometheusNoSinkCounter prometheus.Counter

	numLostEventsSinceLastLog int
	numGoodEventsSinceLastLog int
	lastLostEventsLogTime     time.Time
}

func newBpfEventPoller(e events.Events) *bpfEventPoller {
	p := &bpfEventPoller{
		events:                e,
		sinks:                 make(map[events.Type]bpfEventSink),
		lastLostEventsLogTime: time.Now(),
	}

	var err error

	p.prometheusLostCounter, err = bpfEventsCounters.GetMetricWithLabelValues("lost")
	if err != nil {
		log.WithError(err).Panic("Failed to create prometheus bpf events lost counter")
	}
	p.prometheusNoSinkCounter, err = bpfEventsCounters.GetMetricWithLabelValues("no_sink")
	if err != nil {
		log.WithError(err).Panic("Failed to create prometheus bpf events no sink counter")
	}

	return p
}

func (p *bpfEventPoller) Register(t events.Type, handler bpfEventHandler) {
	sink, ok := p.sinks[t]
	if !ok {
		var err error
		sink.counter, err = bpfEventsCounters.GetMetricWithLabelValues(t.String())
		if err != nil {
			log.WithError(err).Panicf("Failed to create prometheus bpf events %q counter", t.String())
		}
	}
	sink.handlers = append(sink.handlers, handler)
	p.sinks[t] = sink
}

func (p *bpfEventPoller) Start() error {
	if len(p.sinks) == 0 {
		return errors.New("no event sinks registered")
	}

	go p.run()
	return nil
}

func (p *bpfEventPoller) run() {
	for {
		event, err := p.events.Next()
		if err != nil {
			if lost, ok := err.(events.ErrLostEvents); ok {
				p.numLostEventsSinceLastLog += lost.Num()
				p.prometheusLostCounter.Add(float64(lost))
			}
			if time.Since(p.lastLostEventsLogTime) > time.Second*60 {
				log.WithError(err).WithFields(log.Fields{
					"numLostSinceLastLog":       p.numLostEventsSinceLastLog,
					"numGoodEventsSinceLastLog": p.numGoodEventsSinceLastLog,
				}).Warn("Failed to get next event; eBPF-based statistics and/or flow logs collection may be " +
					"impacted (rate limited to 1 log per 60s)")
				p.numLostEventsSinceLastLog = 0
				p.numGoodEventsSinceLastLog = 0
				p.lastLostEventsLogTime = time.Now()
			}
			continue
		}

		sink, ok := p.sinks[event.Type()]
		if !ok {
			log.Warnf("Event type %d without a sink", event.Type())
			p.prometheusNoSinkCounter.Inc()
			continue
		}

		for _, handler := range sink.handlers {
			handler(event)
		}
		p.numGoodEventsSinceLastLog++
		sink.counter.Inc()
	}
}