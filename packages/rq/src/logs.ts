// @generated by protobuf-ts 2.9.1 with parameter generate_dependencies,long_type_number
// @generated from protobuf file "opentelemetry/proto/logs/v1/logs.proto" (package "opentelemetry.proto.logs.v1", syntax proto3)
// tslint:disable
//
// Copyright 2020, OpenTelemetry Authors
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
//
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
import { KeyValue } from "./common";
import { AnyValue } from "./common";
import { InstrumentationScope } from "./common";
import { Resource } from "./resource";
/**
 * LogsData represents the logs data that can be stored in a persistent storage,
 * OR can be embedded by other protocols that transfer OTLP logs data but do not
 * implement the OTLP protocol.
 *
 * The main difference between this message and collector protocol is that
 * in this message there will not be any "control" or "metadata" specific to
 * OTLP protocol.
 *
 * When new fields are added into this message, the OTLP request MUST be updated
 * as well.
 *
 * @generated from protobuf message opentelemetry.proto.logs.v1.LogsData
 */
export interface LogsData {
    /**
     * An array of ResourceLogs.
     * For data coming from a single resource this array will typically contain
     * one element. Intermediary nodes that receive data from multiple origins
     * typically batch the data before forwarding further and in that case this
     * array will contain multiple elements.
     *
     * @generated from protobuf field: repeated opentelemetry.proto.logs.v1.ResourceLogs resource_logs = 1;
     */
    resourceLogs: ResourceLogs[];
}
/**
 * A collection of ScopeLogs from a Resource.
 *
 * @generated from protobuf message opentelemetry.proto.logs.v1.ResourceLogs
 */
export interface ResourceLogs {
    /**
     * The resource for the logs in this message.
     * If this field is not set then resource info is unknown.
     *
     * @generated from protobuf field: opentelemetry.proto.resource.v1.Resource resource = 1;
     */
    resource?: Resource;
    /**
     * A list of ScopeLogs that originate from a resource.
     *
     * @generated from protobuf field: repeated opentelemetry.proto.logs.v1.ScopeLogs scope_logs = 2;
     */
    scopeLogs: ScopeLogs[];
    /**
     * The Schema URL, if known. This is the identifier of the Schema that the resource data
     * is recorded in. To learn more about Schema URL see
     * https://opentelemetry.io/docs/specs/otel/schemas/#schema-url
     * This schema_url applies to the data in the "resource" field. It does not apply
     * to the data in the "scope_logs" field which have their own schema_url field.
     *
     * @generated from protobuf field: string schema_url = 3;
     */
    schemaUrl: string;
}
/**
 * A collection of Logs produced by a Scope.
 *
 * @generated from protobuf message opentelemetry.proto.logs.v1.ScopeLogs
 */
export interface ScopeLogs {
    /**
     * The instrumentation scope information for the logs in this message.
     * Semantically when InstrumentationScope isn't set, it is equivalent with
     * an empty instrumentation scope name (unknown).
     *
     * @generated from protobuf field: opentelemetry.proto.common.v1.InstrumentationScope scope = 1;
     */
    scope?: InstrumentationScope;
    /**
     * A list of log records.
     *
     * @generated from protobuf field: repeated opentelemetry.proto.logs.v1.LogRecord log_records = 2;
     */
    logRecords: LogRecord[];
    /**
     * The Schema URL, if known. This is the identifier of the Schema that the log data
     * is recorded in. To learn more about Schema URL see
     * https://opentelemetry.io/docs/specs/otel/schemas/#schema-url
     * This schema_url applies to all logs in the "logs" field.
     *
     * @generated from protobuf field: string schema_url = 3;
     */
    schemaUrl: string;
}
/**
 * A log record according to OpenTelemetry Log Data Model:
 * https://github.com/open-telemetry/oteps/blob/main/text/logs/0097-log-data-model.md
 *
 * @generated from protobuf message opentelemetry.proto.logs.v1.LogRecord
 */
export interface LogRecord {
    /**
     * time_unix_nano is the time when the event occurred.
     * Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
     * Value of 0 indicates unknown or missing timestamp.
     *
     * @generated from protobuf field: fixed64 time_unix_nano = 1;
     */
    timeUnixNano: number;
    /**
     * Time when the event was observed by the collection system.
     * For events that originate in OpenTelemetry (e.g. using OpenTelemetry Logging SDK)
     * this timestamp is typically set at the generation time and is equal to Timestamp.
     * For events originating externally and collected by OpenTelemetry (e.g. using
     * Collector) this is the time when OpenTelemetry's code observed the event measured
     * by the clock of the OpenTelemetry code. This field MUST be set once the event is
     * observed by OpenTelemetry.
     *
     * For converting OpenTelemetry log data to formats that support only one timestamp or
     * when receiving OpenTelemetry log data by recipients that support only one timestamp
     * internally the following logic is recommended:
     *   - Use time_unix_nano if it is present, otherwise use observed_time_unix_nano.
     *
     * Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
     * Value of 0 indicates unknown or missing timestamp.
     *
     * @generated from protobuf field: fixed64 observed_time_unix_nano = 11;
     */
    observedTimeUnixNano: number;
    /**
     * Numerical value of the severity, normalized to values described in Log Data Model.
     * [Optional].
     *
     * @generated from protobuf field: opentelemetry.proto.logs.v1.SeverityNumber severity_number = 2;
     */
    severityNumber: SeverityNumber;
    /**
     * The severity text (also known as log level). The original string representation as
     * it is known at the source. [Optional].
     *
     * @generated from protobuf field: string severity_text = 3;
     */
    severityText: string;
    /**
     * A value containing the body of the log record. Can be for example a human-readable
     * string message (including multi-line) describing the event in a free form or it can
     * be a structured data composed of arrays and maps of other values. [Optional].
     *
     * @generated from protobuf field: opentelemetry.proto.common.v1.AnyValue body = 5;
     */
    body?: AnyValue;
    /**
     * Additional attributes that describe the specific event occurrence. [Optional].
     * Attribute keys MUST be unique (it is not allowed to have more than one
     * attribute with the same key).
     *
     * @generated from protobuf field: repeated opentelemetry.proto.common.v1.KeyValue attributes = 6;
     */
    attributes: KeyValue[];
    /**
     * @generated from protobuf field: uint32 dropped_attributes_count = 7;
     */
    droppedAttributesCount: number;
    /**
     * Flags, a bit field. 8 least significant bits are the trace flags as
     * defined in W3C Trace Context specification. 24 most significant bits are reserved
     * and must be set to 0. Readers must not assume that 24 most significant bits
     * will be zero and must correctly mask the bits when reading 8-bit trace flag (use
     * flags & LOG_RECORD_FLAGS_TRACE_FLAGS_MASK). [Optional].
     *
     * @generated from protobuf field: fixed32 flags = 8;
     */
    flags: number;
    /**
     * A unique identifier for a trace. All logs from the same trace share
     * the same `trace_id`. The ID is a 16-byte array. An ID with all zeroes OR
     * of length other than 16 bytes is considered invalid (empty string in OTLP/JSON
     * is zero-length and thus is also invalid).
     *
     * This field is optional.
     *
     * The receivers SHOULD assume that the log record is not associated with a
     * trace if any of the following is true:
     *   - the field is not present,
     *   - the field contains an invalid value.
     *
     * @generated from protobuf field: bytes trace_id = 9;
     */
    traceId: Uint8Array;
    /**
     * A unique identifier for a span within a trace, assigned when the span
     * is created. The ID is an 8-byte array. An ID with all zeroes OR of length
     * other than 8 bytes is considered invalid (empty string in OTLP/JSON
     * is zero-length and thus is also invalid).
     *
     * This field is optional. If the sender specifies a valid span_id then it SHOULD also
     * specify a valid trace_id.
     *
     * The receivers SHOULD assume that the log record is not associated with a
     * span if any of the following is true:
     *   - the field is not present,
     *   - the field contains an invalid value.
     *
     * @generated from protobuf field: bytes span_id = 10;
     */
    spanId: Uint8Array;
}
/**
 * Possible values for LogRecord.SeverityNumber.
 *
 * @generated from protobuf enum opentelemetry.proto.logs.v1.SeverityNumber
 */
export enum SeverityNumber {
    /**
     * UNSPECIFIED is the default SeverityNumber, it MUST NOT be used.
     *
     * @generated from protobuf enum value: SEVERITY_NUMBER_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * @generated from protobuf enum value: SEVERITY_NUMBER_TRACE = 1;
     */
    TRACE = 1,
    /**
     * @generated from protobuf enum value: SEVERITY_NUMBER_TRACE2 = 2;
     */
    TRACE2 = 2,
    /**
     * @generated from protobuf enum value: SEVERITY_NUMBER_TRACE3 = 3;
     */
    TRACE3 = 3,
    /**
     * @generated from protobuf enum value: SEVERITY_NUMBER_TRACE4 = 4;
     */
    TRACE4 = 4,
    /**
     * @generated from protobuf enum value: SEVERITY_NUMBER_DEBUG = 5;
     */
    DEBUG = 5,
    /**
     * @generated from protobuf enum value: SEVERITY_NUMBER_DEBUG2 = 6;
     */
    DEBUG2 = 6,
    /**
     * @generated from protobuf enum value: SEVERITY_NUMBER_DEBUG3 = 7;
     */
    DEBUG3 = 7,
    /**
     * @generated from protobuf enum value: SEVERITY_NUMBER_DEBUG4 = 8;
     */
    DEBUG4 = 8,
    /**
     * @generated from protobuf enum value: SEVERITY_NUMBER_INFO = 9;
     */
    INFO = 9,
    /**
     * @generated from protobuf enum value: SEVERITY_NUMBER_INFO2 = 10;
     */
    INFO2 = 10,
    /**
     * @generated from protobuf enum value: SEVERITY_NUMBER_INFO3 = 11;
     */
    INFO3 = 11,
    /**
     * @generated from protobuf enum value: SEVERITY_NUMBER_INFO4 = 12;
     */
    INFO4 = 12,
    /**
     * @generated from protobuf enum value: SEVERITY_NUMBER_WARN = 13;
     */
    WARN = 13,
    /**
     * @generated from protobuf enum value: SEVERITY_NUMBER_WARN2 = 14;
     */
    WARN2 = 14,
    /**
     * @generated from protobuf enum value: SEVERITY_NUMBER_WARN3 = 15;
     */
    WARN3 = 15,
    /**
     * @generated from protobuf enum value: SEVERITY_NUMBER_WARN4 = 16;
     */
    WARN4 = 16,
    /**
     * @generated from protobuf enum value: SEVERITY_NUMBER_ERROR = 17;
     */
    ERROR = 17,
    /**
     * @generated from protobuf enum value: SEVERITY_NUMBER_ERROR2 = 18;
     */
    ERROR2 = 18,
    /**
     * @generated from protobuf enum value: SEVERITY_NUMBER_ERROR3 = 19;
     */
    ERROR3 = 19,
    /**
     * @generated from protobuf enum value: SEVERITY_NUMBER_ERROR4 = 20;
     */
    ERROR4 = 20,
    /**
     * @generated from protobuf enum value: SEVERITY_NUMBER_FATAL = 21;
     */
    FATAL = 21,
    /**
     * @generated from protobuf enum value: SEVERITY_NUMBER_FATAL2 = 22;
     */
    FATAL2 = 22,
    /**
     * @generated from protobuf enum value: SEVERITY_NUMBER_FATAL3 = 23;
     */
    FATAL3 = 23,
    /**
     * @generated from protobuf enum value: SEVERITY_NUMBER_FATAL4 = 24;
     */
    FATAL4 = 24
}
/**
 * LogRecordFlags represents constants used to interpret the
 * LogRecord.flags field, which is protobuf 'fixed32' type and is to
 * be used as bit-fields. Each non-zero value defined in this enum is
 * a bit-mask.  To extract the bit-field, for example, use an
 * expression like:
 *
 *   (logRecord.flags & LOG_RECORD_FLAGS_TRACE_FLAGS_MASK)
 *
 *
 * @generated from protobuf enum opentelemetry.proto.logs.v1.LogRecordFlags
 */
export enum LogRecordFlags {
    /**
     * The zero value for the enum. Should not be used for comparisons.
     * Instead use bitwise "and" with the appropriate mask as shown above.
     *
     * @generated from protobuf enum value: LOG_RECORD_FLAGS_DO_NOT_USE = 0;
     */
    DO_NOT_USE = 0,
    /**
     * Bits 0-7 are used for trace flags.
     *
     * @generated from protobuf enum value: LOG_RECORD_FLAGS_TRACE_FLAGS_MASK = 255;
     */
    TRACE_FLAGS_MASK = 255
}
// @generated message type with reflection information, may provide speed optimized methods
class LogsData$Type extends MessageType<LogsData> {
    constructor() {
        super("opentelemetry.proto.logs.v1.LogsData", [
            { no: 1, name: "resource_logs", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => ResourceLogs }
        ]);
    }
    create(value?: PartialMessage<LogsData>): LogsData {
        const message = { resourceLogs: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<LogsData>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: LogsData): LogsData {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated opentelemetry.proto.logs.v1.ResourceLogs resource_logs */ 1:
                    message.resourceLogs.push(ResourceLogs.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: LogsData, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated opentelemetry.proto.logs.v1.ResourceLogs resource_logs = 1; */
        for (let i = 0; i < message.resourceLogs.length; i++)
            ResourceLogs.internalBinaryWrite(message.resourceLogs[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message opentelemetry.proto.logs.v1.LogsData
 */
export const LogsData = new LogsData$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ResourceLogs$Type extends MessageType<ResourceLogs> {
    constructor() {
        super("opentelemetry.proto.logs.v1.ResourceLogs", [
            { no: 1, name: "resource", kind: "message", T: () => Resource },
            { no: 2, name: "scope_logs", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => ScopeLogs },
            { no: 3, name: "schema_url", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<ResourceLogs>): ResourceLogs {
        const message = { scopeLogs: [], schemaUrl: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<ResourceLogs>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ResourceLogs): ResourceLogs {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* opentelemetry.proto.resource.v1.Resource resource */ 1:
                    message.resource = Resource.internalBinaryRead(reader, reader.uint32(), options, message.resource);
                    break;
                case /* repeated opentelemetry.proto.logs.v1.ScopeLogs scope_logs */ 2:
                    message.scopeLogs.push(ScopeLogs.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* string schema_url */ 3:
                    message.schemaUrl = reader.string();
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
    internalBinaryWrite(message: ResourceLogs, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* opentelemetry.proto.resource.v1.Resource resource = 1; */
        if (message.resource)
            Resource.internalBinaryWrite(message.resource, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* repeated opentelemetry.proto.logs.v1.ScopeLogs scope_logs = 2; */
        for (let i = 0; i < message.scopeLogs.length; i++)
            ScopeLogs.internalBinaryWrite(message.scopeLogs[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* string schema_url = 3; */
        if (message.schemaUrl !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.schemaUrl);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message opentelemetry.proto.logs.v1.ResourceLogs
 */
export const ResourceLogs = new ResourceLogs$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ScopeLogs$Type extends MessageType<ScopeLogs> {
    constructor() {
        super("opentelemetry.proto.logs.v1.ScopeLogs", [
            { no: 1, name: "scope", kind: "message", T: () => InstrumentationScope },
            { no: 2, name: "log_records", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => LogRecord },
            { no: 3, name: "schema_url", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<ScopeLogs>): ScopeLogs {
        const message = { logRecords: [], schemaUrl: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<ScopeLogs>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ScopeLogs): ScopeLogs {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* opentelemetry.proto.common.v1.InstrumentationScope scope */ 1:
                    message.scope = InstrumentationScope.internalBinaryRead(reader, reader.uint32(), options, message.scope);
                    break;
                case /* repeated opentelemetry.proto.logs.v1.LogRecord log_records */ 2:
                    message.logRecords.push(LogRecord.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* string schema_url */ 3:
                    message.schemaUrl = reader.string();
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
    internalBinaryWrite(message: ScopeLogs, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* opentelemetry.proto.common.v1.InstrumentationScope scope = 1; */
        if (message.scope)
            InstrumentationScope.internalBinaryWrite(message.scope, writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* repeated opentelemetry.proto.logs.v1.LogRecord log_records = 2; */
        for (let i = 0; i < message.logRecords.length; i++)
            LogRecord.internalBinaryWrite(message.logRecords[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* string schema_url = 3; */
        if (message.schemaUrl !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.schemaUrl);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message opentelemetry.proto.logs.v1.ScopeLogs
 */
export const ScopeLogs = new ScopeLogs$Type();
// @generated message type with reflection information, may provide speed optimized methods
class LogRecord$Type extends MessageType<LogRecord> {
    constructor() {
        super("opentelemetry.proto.logs.v1.LogRecord", [
            { no: 1, name: "time_unix_nano", kind: "scalar", T: 6 /*ScalarType.FIXED64*/, L: 2 /*LongType.NUMBER*/ },
            { no: 11, name: "observed_time_unix_nano", kind: "scalar", T: 6 /*ScalarType.FIXED64*/, L: 2 /*LongType.NUMBER*/ },
            { no: 2, name: "severity_number", kind: "enum", T: () => ["opentelemetry.proto.logs.v1.SeverityNumber", SeverityNumber, "SEVERITY_NUMBER_"] },
            { no: 3, name: "severity_text", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 5, name: "body", kind: "message", T: () => AnyValue },
            { no: 6, name: "attributes", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => KeyValue },
            { no: 7, name: "dropped_attributes_count", kind: "scalar", T: 13 /*ScalarType.UINT32*/ },
            { no: 8, name: "flags", kind: "scalar", T: 7 /*ScalarType.FIXED32*/ },
            { no: 9, name: "trace_id", kind: "scalar", T: 12 /*ScalarType.BYTES*/ },
            { no: 10, name: "span_id", kind: "scalar", T: 12 /*ScalarType.BYTES*/ }
        ]);
    }
    create(value?: PartialMessage<LogRecord>): LogRecord {
        const message = { timeUnixNano: 0, observedTimeUnixNano: 0, severityNumber: 0, severityText: "", attributes: [], droppedAttributesCount: 0, flags: 0, traceId: new Uint8Array(0), spanId: new Uint8Array(0) };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<LogRecord>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: LogRecord): LogRecord {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* fixed64 time_unix_nano */ 1:
                    message.timeUnixNano = reader.fixed64().toNumber();
                    break;
                case /* fixed64 observed_time_unix_nano */ 11:
                    message.observedTimeUnixNano = reader.fixed64().toNumber();
                    break;
                case /* opentelemetry.proto.logs.v1.SeverityNumber severity_number */ 2:
                    message.severityNumber = reader.int32();
                    break;
                case /* string severity_text */ 3:
                    message.severityText = reader.string();
                    break;
                case /* opentelemetry.proto.common.v1.AnyValue body */ 5:
                    message.body = AnyValue.internalBinaryRead(reader, reader.uint32(), options, message.body);
                    break;
                case /* repeated opentelemetry.proto.common.v1.KeyValue attributes */ 6:
                    message.attributes.push(KeyValue.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* uint32 dropped_attributes_count */ 7:
                    message.droppedAttributesCount = reader.uint32();
                    break;
                case /* fixed32 flags */ 8:
                    message.flags = reader.fixed32();
                    break;
                case /* bytes trace_id */ 9:
                    message.traceId = reader.bytes();
                    break;
                case /* bytes span_id */ 10:
                    message.spanId = reader.bytes();
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
    internalBinaryWrite(message: LogRecord, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* fixed64 time_unix_nano = 1; */
        if (message.timeUnixNano !== 0)
            writer.tag(1, WireType.Bit64).fixed64(message.timeUnixNano);
        /* fixed64 observed_time_unix_nano = 11; */
        if (message.observedTimeUnixNano !== 0)
            writer.tag(11, WireType.Bit64).fixed64(message.observedTimeUnixNano);
        /* opentelemetry.proto.logs.v1.SeverityNumber severity_number = 2; */
        if (message.severityNumber !== 0)
            writer.tag(2, WireType.Varint).int32(message.severityNumber);
        /* string severity_text = 3; */
        if (message.severityText !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.severityText);
        /* opentelemetry.proto.common.v1.AnyValue body = 5; */
        if (message.body)
            AnyValue.internalBinaryWrite(message.body, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        /* repeated opentelemetry.proto.common.v1.KeyValue attributes = 6; */
        for (let i = 0; i < message.attributes.length; i++)
            KeyValue.internalBinaryWrite(message.attributes[i], writer.tag(6, WireType.LengthDelimited).fork(), options).join();
        /* uint32 dropped_attributes_count = 7; */
        if (message.droppedAttributesCount !== 0)
            writer.tag(7, WireType.Varint).uint32(message.droppedAttributesCount);
        /* fixed32 flags = 8; */
        if (message.flags !== 0)
            writer.tag(8, WireType.Bit32).fixed32(message.flags);
        /* bytes trace_id = 9; */
        if (message.traceId.length)
            writer.tag(9, WireType.LengthDelimited).bytes(message.traceId);
        /* bytes span_id = 10; */
        if (message.spanId.length)
            writer.tag(10, WireType.LengthDelimited).bytes(message.spanId);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message opentelemetry.proto.logs.v1.LogRecord
 */
export const LogRecord = new LogRecord$Type();
