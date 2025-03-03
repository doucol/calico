import React from 'react';
import { useOutletContext } from 'react-router-dom';
import { ApiError, FlowLog } from '@/types/api';
import FlowLogsList from '../FlowLogsList';

export type FlowLogsContext = {
    view: 'all' | 'denied';
    flowLogs: FlowLog[];
    error: ApiError | null;
    onRowClicked: () => void;
};

const FlowLogsContainer: React.FC = () => {
    const { flowLogs, error, onRowClicked } =
        useOutletContext<FlowLogsContext>();
    // const { data, isLoading, error } = useFlowLogs(
    //     view === 'denied' ? { action: 'deny' } : undefined,
    // );

    return (
        <FlowLogsList
            flowLogs={flowLogs}
            isLoading={false}
            error={error}
            onRowClicked={onRowClicked}
        />
    );
};

export default FlowLogsContainer;
