// @generated by protobuf-ts 2.9.1 with parameter generate_dependencies
// @generated from protobuf file "rq/v1/sample.proto" (package "v1", syntax proto3)
// tslint:disable
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
/**
 * @generated from protobuf message v1.Sample
 */
export interface Sample {
    /**
     * Serialized Data object, compressed with zstd. We use bytes here because we
     * automatically sore Sample as a arrow.Record.
     *
     * @generated from protobuf field: bytes data = 1;
     */
    data: Uint8Array;
    /**
     * Minimum timetamp observed in this sample in milliseconds
     *
     * @generated from protobuf field: uint64 min_ts = 2;
     */
    minTs: bigint;
    /**
     * Maximum timestamp observed in this sample in milliseconds
     *
     * @generated from protobuf field: uint64 max_ts = 3;
     */
    maxTs: bigint;
    /**
     * Date in nillisecond in which the sample was taken
     *
     * @generated from protobuf field: uint64 date = 4;
     */
    date: bigint;
}
/**
 * @generated from protobuf enum v1.PREFIX
 */
export enum PREFIX {
    /**
     * @generated from protobuf enum value: RESOURCE_SCHEMA = 0;
     */
    RESOURCE_SCHEMA = 0,
    /**
     * @generated from protobuf enum value: RESOURCE_ATTRIBUTES = 1;
     */
    RESOURCE_ATTRIBUTES = 1,
    /**
     * @generated from protobuf enum value: SCOPE_SCHEMA = 2;
     */
    SCOPE_SCHEMA = 2,
    /**
     * @generated from protobuf enum value: SCOPE_NAME = 3;
     */
    SCOPE_NAME = 3,
    /**
     * @generated from protobuf enum value: SCOPE_VERSION = 4;
     */
    SCOPE_VERSION = 4,
    /**
     * @generated from protobuf enum value: SCOPE_ATTRIBUTES = 5;
     */
    SCOPE_ATTRIBUTES = 5,
    /**
     * @generated from protobuf enum value: NAME = 6;
     */
    NAME = 6,
    /**
     * @generated from protobuf enum value: ATTRIBUTES = 7;
     */
    ATTRIBUTES = 7,
    /**
     * @generated from protobuf enum value: TRACE_ID = 8;
     */
    TRACE_ID = 8,
    /**
     * @generated from protobuf enum value: SPAN_ID = 9;
     */
    SPAN_ID = 9,
    /**
     * @generated from protobuf enum value: PARENT_SPAN_ID = 10;
     */
    PARENT_SPAN_ID = 10,
    /**
     * @generated from protobuf enum value: LOGS_LEVEL = 11;
     */
    LOGS_LEVEL = 11
}
/**
 * @generated from protobuf enum v1.SampleKind
 */
export enum SampleKind {
    /**
     * @generated from protobuf enum value: METRICS = 0;
     */
    METRICS = 0,
    /**
     * @generated from protobuf enum value: TRACES = 2;
     */
    TRACES = 2,
    /**
     * @generated from protobuf enum value: LOGS = 3;
     */
    LOGS = 3,
    /**
     * @generated from protobuf enum value: SNIPPETS = 4;
     */
    SNIPPETS = 4
}
// @generated message type with reflection information, may provide speed optimized methods
class Sample$Type extends MessageType<Sample> {
    constructor() {
        super("v1.Sample", [
            { no: 1, name: "data", kind: "scalar", T: 12 /*ScalarType.BYTES*/ },
            { no: 2, name: "min_ts", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 3, name: "max_ts", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 4, name: "date", kind: "scalar", T: 4 /*ScalarType.UINT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
    create(value?: PartialMessage<Sample>): Sample {
        const message = { data: new Uint8Array(0), minTs: 0n, maxTs: 0n, date: 0n };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Sample>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Sample): Sample {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bytes data */ 1:
                    message.data = reader.bytes();
                    break;
                case /* uint64 min_ts */ 2:
                    message.minTs = reader.uint64().toBigInt();
                    break;
                case /* uint64 max_ts */ 3:
                    message.maxTs = reader.uint64().toBigInt();
                    break;
                case /* uint64 date */ 4:
                    message.date = reader.uint64().toBigInt();
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
    internalBinaryWrite(message: Sample, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bytes data = 1; */
        if (message.data.length)
            writer.tag(1, WireType.LengthDelimited).bytes(message.data);
        /* uint64 min_ts = 2; */
        if (message.minTs !== 0n)
            writer.tag(2, WireType.Varint).uint64(message.minTs);
        /* uint64 max_ts = 3; */
        if (message.maxTs !== 0n)
            writer.tag(3, WireType.Varint).uint64(message.maxTs);
        /* uint64 date = 4; */
        if (message.date !== 0n)
            writer.tag(4, WireType.Varint).uint64(message.date);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message v1.Sample
 */
export const Sample = new Sample$Type();