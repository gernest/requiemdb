// @generated by protobuf-ts 2.9.1 with parameter generate_dependencies
// @generated from protobuf file "rq/v1/service.proto" (package "v1", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { RQ } from "./service";
import type { Version } from "./service";
import type { GetVersionRequest } from "./service";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { QueryResponse } from "./service";
import type { QueryRequest } from "./service";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service v1.RQ
 */
export interface IRQClient {
    /**
     * @generated from protobuf rpc: Query(v1.QueryRequest) returns (v1.QueryResponse);
     */
    query(input: QueryRequest, options?: RpcOptions): UnaryCall<QueryRequest, QueryResponse>;
    /**
     * @generated from protobuf rpc: GetVersion(v1.GetVersionRequest) returns (v1.Version);
     */
    getVersion(input: GetVersionRequest, options?: RpcOptions): UnaryCall<GetVersionRequest, Version>;
}
/**
 * @generated from protobuf service v1.RQ
 */
export class RQClient implements IRQClient, ServiceInfo {
    typeName = RQ.typeName;
    methods = RQ.methods;
    options = RQ.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @generated from protobuf rpc: Query(v1.QueryRequest) returns (v1.QueryResponse);
     */
    query(input: QueryRequest, options?: RpcOptions): UnaryCall<QueryRequest, QueryResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<QueryRequest, QueryResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: GetVersion(v1.GetVersionRequest) returns (v1.Version);
     */
    getVersion(input: GetVersionRequest, options?: RpcOptions): UnaryCall<GetVersionRequest, Version> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetVersionRequest, Version>("unary", this._transport, method, opt, input);
    }
}
