// Copyright (c) 2020 Tigera, Inc. All rights reserved.
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

package idalloc

import (
	"errors"
	"sort"

	"github.com/projectcalico/calico/libcalico-go/lib/set"
)

type IndexRange struct {
	Min, Max int
}

func (r IndexRange) contains(a int) bool {
	return r.Min <= a && r.Max >= a
}

// ByMaxIndex sorts collections of IndexRange structs in order of their starting/lower index
type ByMaxIndex []IndexRange

// Len is the number of indexranges in the collection
func (i ByMaxIndex) Len() int { return len(i) }

// Less reports whether the element with index a
// must sort before the element with index b.
func (i ByMaxIndex) Less(a, b int) bool { return i[a].Max < i[b].Max }

// Swap swaps the elements with indexes a and b.
func (i ByMaxIndex) Swap(a, b int) { i[a], i[b] = i[b], i[a] }

type IndexAllocator struct {
	indexStack *stack
	exclusions []IndexRange
}

// NewIndexAllocator returns an index allocator from the provided indexRanges
// any indices falling within the specified exclusions will not be returned, even if designated by indexRanges
func NewIndexAllocator(indexRanges []IndexRange, exclusions []IndexRange) *IndexAllocator {
	// sort index ranges in descending order of their Max bound
	if len(indexRanges) > 1 {
		sort.Sort(sort.Reverse(ByMaxIndex(indexRanges)))
	}

	r := &IndexAllocator{
		indexStack: &stack{},
		exclusions: exclusions,
	}

	var lowestIndex int
	for j, indexRange := range indexRanges {
		if j == 0 {
			// keep track of the most recent Min to prevent an overlapping range
			// from creating duplicate indices in the final index stack
			lowestIndex = indexRange.Max + 1
		}

		// Push in reverse order so that the lowest index will come out first.
	populating:
		for i := indexRange.Max; i >= indexRange.Min; i-- {
			// skip overlapping range indices
			if i >= lowestIndex {
				continue
			}
			// skip exclusions
			for _, excl := range exclusions {
				if excl.contains(i) {
					continue populating
				}
			}
			r.indexStack.Push(i)
			lowestIndex = i
		}
	}
	return r
}

func (r *IndexAllocator) GrabIndex() (int, error) {
	if r.indexStack.IsEmpty() {
		return 0, errors.New("no more indices available")
	}
	return r.indexStack.Pop(), nil
}

func (r *IndexAllocator) ReleaseIndex(index int) {
	r.indexStack.Push(index)
}

// GrabBlock tries to grab a contiguous block of indices from the stack
func (r *IndexAllocator) GrabBlock(len int) (set.Set[int], error) {
	indices := set.New[int]()
	for i := 0; i < len; i++ {
		idx, err := r.GrabIndex()
		if err != nil {
			return indices, err
		}
		indices.Add(idx)
	}
	return indices, nil
}

type (
	// stack is an implementation of the stack data structure.
	// Adapted from https://github.com/golang-collections/collections/blob/604e922904d35e97f98a774db7881f049cd8d970/stack/stack.go.
	stack struct {
		top *node
	}
	node struct {
		value int
		prev  *node
	}
)

// IsEmpty returns true if the stack is empty.
func (s *stack) IsEmpty() bool {
	return s.top == nil
}

// Pop the top item of the stack and return it
func (s *stack) Pop() int {
	if s.IsEmpty() {
		return 0
	}

	n := s.top
	s.top = n.prev
	return n.value
}

// Push a value onto the top of the stack
func (s *stack) Push(value int) {
	s.top = &node{value, s.top}
}
