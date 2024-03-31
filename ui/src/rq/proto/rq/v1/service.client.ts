// @generated by protobuf-ts 2.9.1 with parameter generate_dependencies
// @generated from protobuf file "rq/v1/service.proto" (package "v1", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { RQ } from "./service";
import type { Version } from "./service";
import type { GetVersionRequest } from "./service";
import type { GetSnippetResponse } from "./service";
import type { GetSnippetRequest } from "./service";
import type { SnippetInfo_List } from "./service";
import type { ListStippetsRequest } from "./service";
import type { RenameSnippetResponse } from "./service";
import type { RenameSnippetRequest } from "./service";
import type { UploadSnippetResponse } from "./service";
import type { UploadSnippetRequest } from "./service";
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
     * @generated from protobuf rpc: UploadSnippet(v1.UploadSnippetRequest) returns (v1.UploadSnippetResponse);
     */
    uploadSnippet(input: UploadSnippetRequest, options?: RpcOptions): UnaryCall<UploadSnippetRequest, UploadSnippetResponse>;
    /**
     * @generated from protobuf rpc: RenameSnippet(v1.RenameSnippetRequest) returns (v1.RenameSnippetResponse);
     */
    renameSnippet(input: RenameSnippetRequest, options?: RpcOptions): UnaryCall<RenameSnippetRequest, RenameSnippetResponse>;
    /**
     * @generated from protobuf rpc: ListSnippets(v1.ListStippetsRequest) returns (v1.SnippetInfo.List);
     */
    listSnippets(input: ListStippetsRequest, options?: RpcOptions): UnaryCall<ListStippetsRequest, SnippetInfo_List>;
    /**
     * @generated from protobuf rpc: GetSnippet(v1.GetSnippetRequest) returns (v1.GetSnippetResponse);
     */
    getSnippet(input: GetSnippetRequest, options?: RpcOptions): UnaryCall<GetSnippetRequest, GetSnippetResponse>;
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
     * @generated from protobuf rpc: UploadSnippet(v1.UploadSnippetRequest) returns (v1.UploadSnippetResponse);
     */
    uploadSnippet(input: UploadSnippetRequest, options?: RpcOptions): UnaryCall<UploadSnippetRequest, UploadSnippetResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<UploadSnippetRequest, UploadSnippetResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: RenameSnippet(v1.RenameSnippetRequest) returns (v1.RenameSnippetResponse);
     */
    renameSnippet(input: RenameSnippetRequest, options?: RpcOptions): UnaryCall<RenameSnippetRequest, RenameSnippetResponse> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<RenameSnippetRequest, RenameSnippetResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: ListSnippets(v1.ListStippetsRequest) returns (v1.SnippetInfo.List);
     */
    listSnippets(input: ListStippetsRequest, options?: RpcOptions): UnaryCall<ListStippetsRequest, SnippetInfo_List> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<ListStippetsRequest, SnippetInfo_List>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: GetSnippet(v1.GetSnippetRequest) returns (v1.GetSnippetResponse);
     */
    getSnippet(input: GetSnippetRequest, options?: RpcOptions): UnaryCall<GetSnippetRequest, GetSnippetResponse> {
        const method = this.methods[4], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetSnippetRequest, GetSnippetResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: GetVersion(v1.GetVersionRequest) returns (v1.Version);
     */
    getVersion(input: GetVersionRequest, options?: RpcOptions): UnaryCall<GetVersionRequest, Version> {
        const method = this.methods[5], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetVersionRequest, Version>("unary", this._transport, method, opt, input);
    }
}
