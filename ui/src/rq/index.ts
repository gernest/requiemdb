import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';

export * from "./proto"

import { RQClient } from "./proto";

const createTransport = () => {
    let transport = new GrpcWebFetchTransport({
        baseUrl: window.location.origin,
    });

    if (process.env.NODE_ENV === 'development') {
        transport = new GrpcWebFetchTransport({
            baseUrl: process.env.REACT_APP_API_URL!,
        })
    }
    return transport
}

export const createOTSClient = () => {
    return new RQClient(createTransport())
}