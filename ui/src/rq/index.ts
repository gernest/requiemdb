import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';

export * from "./proto"

import { RQClient } from "./proto";

const createTransport = () => {
    return new GrpcWebFetchTransport({
        baseUrl: window.location.origin,
    });
}

export const createRQClient = () => {
    return new RQClient(createTransport())
}