// @generated by protobuf-ts 2.9.1 with parameter generate_dependencies,long_type_number
// @generated from protobuf file "rq/v1/service.proto" (package "v1", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { Duration } from "../../google/protobuf/duration";
import { Timestamp } from "../../google/protobuf/timestamp";
/**
 * @generated from protobuf message v1.GetSnippetRequest
 */
export interface GetSnippetRequest {
    /**
     * @generated from protobuf field: string name = 1;
     */
    name: string;
}
/**
 * @generated from protobuf message v1.GetSnippetResponse
 */
export interface GetSnippetResponse {
    /**
     * @generated from protobuf field: string name = 1;
     */
    name: string;
    /**
     * @generated from protobuf field: string description = 2;
     */
    description: string;
    /**
     * @generated from protobuf field: bytes raw = 3;
     */
    raw: Uint8Array;
    /**
     * @generated from protobuf field: google.protobuf.Timestamp created_at = 4;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: google.protobuf.Timestamp updated_at = 5;
     */
    updatedAt?: Timestamp;
}
/**
 * @generated from protobuf message v1.ListStippetsRequest
 */
export interface ListStippetsRequest {
}
/**
 * @generated from protobuf message v1.ListSnippetsResponse
 */
export interface ListSnippetsResponse {
    /**
     * @generated from protobuf field: repeated v1.SnippetInfo snippets = 1;
     */
    snippets: SnippetInfo[];
}
/**
 * @generated from protobuf message v1.SnippetInfo
 */
export interface SnippetInfo {
    /**
     * @generated from protobuf field: string name = 1;
     */
    name: string;
    /**
     * @generated from protobuf field: string description = 2;
     */
    description: string;
    /**
     * @generated from protobuf field: google.protobuf.Timestamp created_at = 3;
     */
    createdAt?: Timestamp;
    /**
     * @generated from protobuf field: google.protobuf.Timestamp updated_at = 4;
     */
    updatedAt?: Timestamp;
}
/**
 * @generated from protobuf message v1.UploadSnippetRequest
 */
export interface UploadSnippetRequest {
    /**
     * @generated from protobuf field: string name = 1;
     */
    name: string;
    /**
     * @generated from protobuf field: bytes data = 2;
     */
    data: Uint8Array;
}
/**
 * @generated from protobuf message v1.UploadSnippetResponse
 */
export interface UploadSnippetResponse {
}
/**
 * @generated from protobuf message v1.QueryRequest
 */
export interface QueryRequest {
    /**
     * @generated from protobuf oneof: request
     */
    request: {
        oneofKind: "scriptName";
        /**
         * Name of existing snippet.
         *
         * @generated from protobuf field: string script_name = 1;
         */
        scriptName: string;
    } | {
        oneofKind: "scriptData";
        /**
         * This will be compiled and executed, It is not recommended. We cache for
         * faster execution .
         *
         * Useful when experimenting.
         *
         * @generated from protobuf field: bytes script_data = 2;
         */
        scriptData: Uint8Array;
    } | {
        oneofKind: undefined;
    };
    /**
     * when true any logs associated with script execution will be included in the
     * response.
     *
     * @generated from protobuf field: bool include_logs = 3;
     */
    includeLogs: boolean;
    /**
     * @generated from protobuf field: google.protobuf.Timestamp start_date = 4;
     */
    startDate?: Timestamp;
    /**
     * @generated from protobuf field: google.protobuf.Timestamp end_date = 5;
     */
    endDate?: Timestamp;
}
/**
 * @generated from protobuf message v1.QueryResponse
 */
export interface QueryResponse {
    /**
     * @generated from protobuf field: v1.Result value = 1;
     */
    value?: Result;
    /**
     * @generated from protobuf field: v1.Timings timings = 2;
     */
    timings?: Timings; // Number of samples processed by the query snippet.
}
/**
 * @generated from protobuf message v1.Timings
 */
export interface Timings {
    /**
     * Time taken to compile the query snippet before evealuation.
     *
     * @generated from protobuf field: google.protobuf.Duration compiling = 1;
     */
    compiling?: Duration;
    /**
     * Time taken to evaluate compiled query.
     *
     * @generated from protobuf field: google.protobuf.Duration evaluating = 2;
     */
    evaluating?: Duration;
}
/**
 * @generated from protobuf message v1.Result
 */
export interface Result {
}
// @generated message type with reflection information, may provide speed optimized methods
class GetSnippetRequest$Type extends MessageType<GetSnippetRequest> {
    constructor() {
        super("v1.GetSnippetRequest", [
            { no: 1, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<GetSnippetRequest>): GetSnippetRequest {
        const message = { name: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<GetSnippetRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetSnippetRequest): GetSnippetRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string name */ 1:
                    message.name = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: GetSnippetRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string name = 1; */
        if (message.name !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.name);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.GetSnippetRequest
 */
export const GetSnippetRequest = new GetSnippetRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetSnippetResponse$Type extends MessageType<GetSnippetResponse> {
    constructor() {
        super("v1.GetSnippetResponse", [
            { no: 1, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "description", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "raw", kind: "scalar", T: 12 /*ScalarType.BYTES*/ },
            { no: 4, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 5, name: "updated_at", kind: "message", T: () => Timestamp }
        ]);
    }
    create(value?: PartialMessage<GetSnippetResponse>): GetSnippetResponse {
        const message = { name: "", description: "", raw: new Uint8Array(0) };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<GetSnippetResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: GetSnippetResponse): GetSnippetResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string name */ 1:
                    message.name = reader.string();
                    break;
                case /* string description */ 2:
                    message.description = reader.string();
                    break;
                case /* bytes raw */ 3:
                    message.raw = reader.bytes();
                    break;
                case /* google.protobuf.Timestamp created_at */ 4:
                    message.createdAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.createdAt);
                    break;
                case /* google.protobuf.Timestamp updated_at */ 5:
                    message.updatedAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.updatedAt);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: GetSnippetResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string name = 1; */
        if (message.name !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.name);
        /* string description = 2; */
        if (message.description !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.description);
        /* bytes raw = 3; */
        if (message.raw.length)
            writer.tag(3, WireType.LengthDelimited).bytes(message.raw);
        /* google.protobuf.Timestamp created_at = 4; */
        if (message.createdAt)
            Timestamp.internalBinaryWrite(message.createdAt, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* google.protobuf.Timestamp updated_at = 5; */
        if (message.updatedAt)
            Timestamp.internalBinaryWrite(message.updatedAt, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.GetSnippetResponse
 */
export const GetSnippetResponse = new GetSnippetResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListStippetsRequest$Type extends MessageType<ListStippetsRequest> {
    constructor() {
        super("v1.ListStippetsRequest", []);
    }
    create(value?: PartialMessage<ListStippetsRequest>): ListStippetsRequest {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<ListStippetsRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ListStippetsRequest): ListStippetsRequest {
        return target ?? this.create();
    }
    internalBinaryWrite(message: ListStippetsRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.ListStippetsRequest
 */
export const ListStippetsRequest = new ListStippetsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ListSnippetsResponse$Type extends MessageType<ListSnippetsResponse> {
    constructor() {
        super("v1.ListSnippetsResponse", [
            { no: 1, name: "snippets", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => SnippetInfo }
        ]);
    }
    create(value?: PartialMessage<ListSnippetsResponse>): ListSnippetsResponse {
        const message = { snippets: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<ListSnippetsResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ListSnippetsResponse): ListSnippetsResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated v1.SnippetInfo snippets */ 1:
                    message.snippets.push(SnippetInfo.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: ListSnippetsResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated v1.SnippetInfo snippets = 1; */
        for (let i = 0; i < message.snippets.length; i++)
            SnippetInfo.internalBinaryWrite(message.snippets[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.ListSnippetsResponse
 */
export const ListSnippetsResponse = new ListSnippetsResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SnippetInfo$Type extends MessageType<SnippetInfo> {
    constructor() {
        super("v1.SnippetInfo", [
            { no: 1, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "description", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "created_at", kind: "message", T: () => Timestamp },
            { no: 4, name: "updated_at", kind: "message", T: () => Timestamp }
        ]);
    }
    create(value?: PartialMessage<SnippetInfo>): SnippetInfo {
        const message = { name: "", description: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<SnippetInfo>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SnippetInfo): SnippetInfo {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string name */ 1:
                    message.name = reader.string();
                    break;
                case /* string description */ 2:
                    message.description = reader.string();
                    break;
                case /* google.protobuf.Timestamp created_at */ 3:
                    message.createdAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.createdAt);
                    break;
                case /* google.protobuf.Timestamp updated_at */ 4:
                    message.updatedAt = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.updatedAt);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: SnippetInfo, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string name = 1; */
        if (message.name !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.name);
        /* string description = 2; */
        if (message.description !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.description);
        /* google.protobuf.Timestamp created_at = 3; */
        if (message.createdAt)
            Timestamp.internalBinaryWrite(message.createdAt, writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* google.protobuf.Timestamp updated_at = 4; */
        if (message.updatedAt)
            Timestamp.internalBinaryWrite(message.updatedAt, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.SnippetInfo
 */
export const SnippetInfo = new SnippetInfo$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UploadSnippetRequest$Type extends MessageType<UploadSnippetRequest> {
    constructor() {
        super("v1.UploadSnippetRequest", [
            { no: 1, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "data", kind: "scalar", T: 12 /*ScalarType.BYTES*/ }
        ]);
    }
    create(value?: PartialMessage<UploadSnippetRequest>): UploadSnippetRequest {
        const message = { name: "", data: new Uint8Array(0) };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<UploadSnippetRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UploadSnippetRequest): UploadSnippetRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string name */ 1:
                    message.name = reader.string();
                    break;
                case /* bytes data */ 2:
                    message.data = reader.bytes();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: UploadSnippetRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string name = 1; */
        if (message.name !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.name);
        /* bytes data = 2; */
        if (message.data.length)
            writer.tag(2, WireType.LengthDelimited).bytes(message.data);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.UploadSnippetRequest
 */
export const UploadSnippetRequest = new UploadSnippetRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class UploadSnippetResponse$Type extends MessageType<UploadSnippetResponse> {
    constructor() {
        super("v1.UploadSnippetResponse", []);
    }
    create(value?: PartialMessage<UploadSnippetResponse>): UploadSnippetResponse {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<UploadSnippetResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: UploadSnippetResponse): UploadSnippetResponse {
        return target ?? this.create();
    }
    internalBinaryWrite(message: UploadSnippetResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.UploadSnippetResponse
 */
export const UploadSnippetResponse = new UploadSnippetResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class QueryRequest$Type extends MessageType<QueryRequest> {
    constructor() {
        super("v1.QueryRequest", [
            { no: 1, name: "script_name", kind: "scalar", oneof: "request", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "script_data", kind: "scalar", oneof: "request", T: 12 /*ScalarType.BYTES*/ },
            { no: 3, name: "include_logs", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 4, name: "start_date", kind: "message", T: () => Timestamp },
            { no: 5, name: "end_date", kind: "message", T: () => Timestamp }
        ]);
    }
    create(value?: PartialMessage<QueryRequest>): QueryRequest {
        const message = { request: { oneofKind: undefined }, includeLogs: false };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<QueryRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: QueryRequest): QueryRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string script_name */ 1:
                    message.request = {
                        oneofKind: "scriptName",
                        scriptName: reader.string()
                    };
                    break;
                case /* bytes script_data */ 2:
                    message.request = {
                        oneofKind: "scriptData",
                        scriptData: reader.bytes()
                    };
                    break;
                case /* bool include_logs */ 3:
                    message.includeLogs = reader.bool();
                    break;
                case /* google.protobuf.Timestamp start_date */ 4:
                    message.startDate = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.startDate);
                    break;
                case /* google.protobuf.Timestamp end_date */ 5:
                    message.endDate = Timestamp.internalBinaryRead(reader, reader.uint32(), options, message.endDate);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: QueryRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string script_name = 1; */
        if (message.request.oneofKind === "scriptName")
            writer.tag(1, WireType.LengthDelimited).string(message.request.scriptName);
        /* bytes script_data = 2; */
        if (message.request.oneofKind === "scriptData")
            writer.tag(2, WireType.LengthDelimited).bytes(message.request.scriptData);
        /* bool include_logs = 3; */
        if (message.includeLogs !== false)
            writer.tag(3, WireType.Varint).bool(message.includeLogs);
        /* google.protobuf.Timestamp start_date = 4; */
        if (message.startDate)
            Timestamp.internalBinaryWrite(message.startDate, writer.tag(4, WireType.LengthDelimited).fork(), options).join();
        /* google.protobuf.Timestamp end_date = 5; */
        if (message.endDate)
            Timestamp.internalBinaryWrite(message.endDate, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.QueryRequest
 */
export const QueryRequest = new QueryRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class QueryResponse$Type extends MessageType<QueryResponse> {
    constructor() {
        super("v1.QueryResponse", [
            { no: 1, name: "value", kind: "message", T: () => Result },
            { no: 2, name: "timings", kind: "message", T: () => Timings }
        ]);
    }
    create(value?: PartialMessage<QueryResponse>): QueryResponse {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<QueryResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: QueryResponse): QueryResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* v1.Result value */ 1:
                    message.value = Result.internalBinaryRead(reader, reader.uint32(), options, message.value);
                    break;
                case /* v1.Timings timings */ 2:
                    message.timings = Timings.internalBinaryRead(reader, reader.uint32(), options, message.timings);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: QueryResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* v1.Result value = 1; */
        if (message.value)
            Result.internalBinaryWrite(message.value, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* v1.Timings timings = 2; */
        if (message.timings)
            Timings.internalBinaryWrite(message.timings, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.QueryResponse
 */
export const QueryResponse = new QueryResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Timings$Type extends MessageType<Timings> {
    constructor() {
        super("v1.Timings", [
            { no: 1, name: "compiling", kind: "message", T: () => Duration },
            { no: 2, name: "evaluating", kind: "message", T: () => Duration }
        ]);
    }
    create(value?: PartialMessage<Timings>): Timings {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Timings>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Timings): Timings {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* google.protobuf.Duration compiling */ 1:
                    message.compiling = Duration.internalBinaryRead(reader, reader.uint32(), options, message.compiling);
                    break;
                case /* google.protobuf.Duration evaluating */ 2:
                    message.evaluating = Duration.internalBinaryRead(reader, reader.uint32(), options, message.evaluating);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: Timings, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* google.protobuf.Duration compiling = 1; */
        if (message.compiling)
            Duration.internalBinaryWrite(message.compiling, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* google.protobuf.Duration evaluating = 2; */
        if (message.evaluating)
            Duration.internalBinaryWrite(message.evaluating, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.Timings
 */
export const Timings = new Timings$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Result$Type extends MessageType<Result> {
    constructor() {
        super("v1.Result", []);
    }
    create(value?: PartialMessage<Result>): Result {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Result>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Result): Result {
        return target ?? this.create();
    }
    internalBinaryWrite(message: Result, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.Result
 */
export const Result = new Result$Type();
/**
 * @generated ServiceType for protobuf service v1.RQ
 */
export const RQ = new ServiceType("v1.RQ", [
    { name: "Query", options: { "google.api.http": { post: "/api/v1/query", body: "*" } }, I: QueryRequest, O: QueryResponse },
    { name: "UploadSnippet", options: { "google.api.http": { post: "/api/v1/upload", body: "*" } }, I: UploadSnippetRequest, O: UploadSnippetResponse },
    { name: "ListSnippets", options: { "google.api.http": { get: "/api/v1/list" } }, I: ListStippetsRequest, O: ListSnippetsResponse },
    { name: "GetSnippet", options: { "google.api.http": { get: "/api/v1/snippet/{name}" } }, I: GetSnippetRequest, O: GetSnippetResponse }
]);
