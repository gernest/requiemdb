// @generated by protobuf-ts 2.9.1 with parameter generate_dependencies
// @generated from protobuf file "opentelemetry/proto/common/v1/common.proto" (package "opentelemetry.proto.common.v1", syntax proto3)
// tslint:disable
//
// Copyright 2019, OpenTelemetry Authors
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
/**
 * AnyValue is used to represent any type of attribute value. AnyValue may contain a
 * primitive value such as a string or integer or it may contain an arbitrary nested
 * object containing arrays, key-value lists and primitives.
 *
 * @generated from protobuf message opentelemetry.proto.common.v1.AnyValue
 */
export interface AnyValue {
    /**
     * @generated from protobuf oneof: value
     */
    value: {
        oneofKind: "stringValue";
        /**
         * @generated from protobuf field: string string_value = 1;
         */
        stringValue: string;
    } | {
        oneofKind: "boolValue";
        /**
         * @generated from protobuf field: bool bool_value = 2;
         */
        boolValue: boolean;
    } | {
        oneofKind: "intValue";
        /**
         * @generated from protobuf field: int64 int_value = 3;
         */
        intValue: bigint;
    } | {
        oneofKind: "doubleValue";
        /**
         * @generated from protobuf field: double double_value = 4;
         */
        doubleValue: number;
    } | {
        oneofKind: "arrayValue";
        /**
         * @generated from protobuf field: opentelemetry.proto.common.v1.ArrayValue array_value = 5;
         */
        arrayValue: ArrayValue;
    } | {
        oneofKind: "kvlistValue";
        /**
         * @generated from protobuf field: opentelemetry.proto.common.v1.KeyValueList kvlist_value = 6;
         */
        kvlistValue: KeyValueList;
    } | {
        oneofKind: "bytesValue";
        /**
         * @generated from protobuf field: bytes bytes_value = 7;
         */
        bytesValue: Uint8Array;
    } | {
        oneofKind: undefined;
    };
}
/**
 * ArrayValue is a list of AnyValue messages. We need ArrayValue as a message
 * since oneof in AnyValue does not allow repeated fields.
 *
 * @generated from protobuf message opentelemetry.proto.common.v1.ArrayValue
 */
export interface ArrayValue {
    /**
     * Array of values. The array may be empty (contain 0 elements).
     *
     * @generated from protobuf field: repeated opentelemetry.proto.common.v1.AnyValue values = 1;
     */
    values: AnyValue[];
}
/**
 * KeyValueList is a list of KeyValue messages. We need KeyValueList as a message
 * since `oneof` in AnyValue does not allow repeated fields. Everywhere else where we need
 * a list of KeyValue messages (e.g. in Span) we use `repeated KeyValue` directly to
 * avoid unnecessary extra wrapping (which slows down the protocol). The 2 approaches
 * are semantically equivalent.
 *
 * @generated from protobuf message opentelemetry.proto.common.v1.KeyValueList
 */
export interface KeyValueList {
    /**
     * A collection of key/value pairs of key-value pairs. The list may be empty (may
     * contain 0 elements).
     * The keys MUST be unique (it is not allowed to have more than one
     * value with the same key).
     *
     * @generated from protobuf field: repeated opentelemetry.proto.common.v1.KeyValue values = 1;
     */
    values: KeyValue[];
}
/**
 * KeyValue is a key-value pair that is used to store Span attributes, Link
 * attributes, etc.
 *
 * @generated from protobuf message opentelemetry.proto.common.v1.KeyValue
 */
export interface KeyValue {
    /**
     * @generated from protobuf field: string key = 1;
     */
    key: string;
    /**
     * @generated from protobuf field: opentelemetry.proto.common.v1.AnyValue value = 2;
     */
    value?: AnyValue;
}
/**
 * InstrumentationScope is a message representing the instrumentation scope information
 * such as the fully qualified name and version.
 *
 * @generated from protobuf message opentelemetry.proto.common.v1.InstrumentationScope
 */
export interface InstrumentationScope {
    /**
     * An empty instrumentation scope name means the name is unknown.
     *
     * @generated from protobuf field: string name = 1;
     */
    name: string;
    /**
     * @generated from protobuf field: string version = 2;
     */
    version: string;
    /**
     * Additional attributes that describe the scope. [Optional].
     * Attribute keys MUST be unique (it is not allowed to have more than one
     * attribute with the same key).
     *
     * @generated from protobuf field: repeated opentelemetry.proto.common.v1.KeyValue attributes = 3;
     */
    attributes: KeyValue[];
    /**
     * @generated from protobuf field: uint32 dropped_attributes_count = 4;
     */
    droppedAttributesCount: number;
}
// @generated message type with reflection information, may provide speed optimized methods
class AnyValue$Type extends MessageType<AnyValue> {
    constructor() {
        super("opentelemetry.proto.common.v1.AnyValue", [
            { no: 1, name: "string_value", kind: "scalar", oneof: "value", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "bool_value", kind: "scalar", oneof: "value", T: 8 /*ScalarType.BOOL*/ },
            { no: 3, name: "int_value", kind: "scalar", oneof: "value", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 4, name: "double_value", kind: "scalar", oneof: "value", T: 1 /*ScalarType.DOUBLE*/ },
            { no: 5, name: "array_value", kind: "message", oneof: "value", T: () => ArrayValue },
            { no: 6, name: "kvlist_value", kind: "message", oneof: "value", T: () => KeyValueList },
            { no: 7, name: "bytes_value", kind: "scalar", oneof: "value", T: 12 /*ScalarType.BYTES*/ }
        ]);
    }
    create(value?: PartialMessage<AnyValue>): AnyValue {
        const message = { value: { oneofKind: undefined } };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<AnyValue>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: AnyValue): AnyValue {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string string_value */ 1:
                    message.value = {
                        oneofKind: "stringValue",
                        stringValue: reader.string()
                    };
                    break;
                case /* bool bool_value */ 2:
                    message.value = {
                        oneofKind: "boolValue",
                        boolValue: reader.bool()
                    };
                    break;
                case /* int64 int_value */ 3:
                    message.value = {
                        oneofKind: "intValue",
                        intValue: reader.int64().toBigInt()
                    };
                    break;
                case /* double double_value */ 4:
                    message.value = {
                        oneofKind: "doubleValue",
                        doubleValue: reader.double()
                    };
                    break;
                case /* opentelemetry.proto.common.v1.ArrayValue array_value */ 5:
                    message.value = {
                        oneofKind: "arrayValue",
                        arrayValue: ArrayValue.internalBinaryRead(reader, reader.uint32(), options, (message.value as any).arrayValue)
                    };
                    break;
                case /* opentelemetry.proto.common.v1.KeyValueList kvlist_value */ 6:
                    message.value = {
                        oneofKind: "kvlistValue",
                        kvlistValue: KeyValueList.internalBinaryRead(reader, reader.uint32(), options, (message.value as any).kvlistValue)
                    };
                    break;
                case /* bytes bytes_value */ 7:
                    message.value = {
                        oneofKind: "bytesValue",
                        bytesValue: reader.bytes()
                    };
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
    internalBinaryWrite(message: AnyValue, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string string_value = 1; */
        if (message.value.oneofKind === "stringValue")
            writer.tag(1, WireType.LengthDelimited).string(message.value.stringValue);
        /* bool bool_value = 2; */
        if (message.value.oneofKind === "boolValue")
            writer.tag(2, WireType.Varint).bool(message.value.boolValue);
        /* int64 int_value = 3; */
        if (message.value.oneofKind === "intValue")
            writer.tag(3, WireType.Varint).int64(message.value.intValue);
        /* double double_value = 4; */
        if (message.value.oneofKind === "doubleValue")
            writer.tag(4, WireType.Bit64).double(message.value.doubleValue);
        /* opentelemetry.proto.common.v1.ArrayValue array_value = 5; */
        if (message.value.oneofKind === "arrayValue")
            ArrayValue.internalBinaryWrite(message.value.arrayValue, writer.tag(5, WireType.LengthDelimited).fork(), options).join();
        /* opentelemetry.proto.common.v1.KeyValueList kvlist_value = 6; */
        if (message.value.oneofKind === "kvlistValue")
            KeyValueList.internalBinaryWrite(message.value.kvlistValue, writer.tag(6, WireType.LengthDelimited).fork(), options).join();
        /* bytes bytes_value = 7; */
        if (message.value.oneofKind === "bytesValue")
            writer.tag(7, WireType.LengthDelimited).bytes(message.value.bytesValue);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message opentelemetry.proto.common.v1.AnyValue
 */
export const AnyValue = new AnyValue$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ArrayValue$Type extends MessageType<ArrayValue> {
    constructor() {
        super("opentelemetry.proto.common.v1.ArrayValue", [
            { no: 1, name: "values", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => AnyValue }
        ]);
    }
    create(value?: PartialMessage<ArrayValue>): ArrayValue {
        const message = { values: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<ArrayValue>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ArrayValue): ArrayValue {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated opentelemetry.proto.common.v1.AnyValue values */ 1:
                    message.values.push(AnyValue.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: ArrayValue, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated opentelemetry.proto.common.v1.AnyValue values = 1; */
        for (let i = 0; i < message.values.length; i++)
            AnyValue.internalBinaryWrite(message.values[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message opentelemetry.proto.common.v1.ArrayValue
 */
export const ArrayValue = new ArrayValue$Type();
// @generated message type with reflection information, may provide speed optimized methods
class KeyValueList$Type extends MessageType<KeyValueList> {
    constructor() {
        super("opentelemetry.proto.common.v1.KeyValueList", [
            { no: 1, name: "values", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => KeyValue }
        ]);
    }
    create(value?: PartialMessage<KeyValueList>): KeyValueList {
        const message = { values: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<KeyValueList>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: KeyValueList): KeyValueList {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated opentelemetry.proto.common.v1.KeyValue values */ 1:
                    message.values.push(KeyValue.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: KeyValueList, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated opentelemetry.proto.common.v1.KeyValue values = 1; */
        for (let i = 0; i < message.values.length; i++)
            KeyValue.internalBinaryWrite(message.values[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message opentelemetry.proto.common.v1.KeyValueList
 */
export const KeyValueList = new KeyValueList$Type();
// @generated message type with reflection information, may provide speed optimized methods
class KeyValue$Type extends MessageType<KeyValue> {
    constructor() {
        super("opentelemetry.proto.common.v1.KeyValue", [
            { no: 1, name: "key", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "value", kind: "message", T: () => AnyValue }
        ]);
    }
    create(value?: PartialMessage<KeyValue>): KeyValue {
        const message = { key: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<KeyValue>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: KeyValue): KeyValue {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string key */ 1:
                    message.key = reader.string();
                    break;
                case /* opentelemetry.proto.common.v1.AnyValue value */ 2:
                    message.value = AnyValue.internalBinaryRead(reader, reader.uint32(), options, message.value);
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
    internalBinaryWrite(message: KeyValue, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string key = 1; */
        if (message.key !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.key);
        /* opentelemetry.proto.common.v1.AnyValue value = 2; */
        if (message.value)
            AnyValue.internalBinaryWrite(message.value, writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message opentelemetry.proto.common.v1.KeyValue
 */
export const KeyValue = new KeyValue$Type();
// @generated message type with reflection information, may provide speed optimized methods
class InstrumentationScope$Type extends MessageType<InstrumentationScope> {
    constructor() {
        super("opentelemetry.proto.common.v1.InstrumentationScope", [
            { no: 1, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "version", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "attributes", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => KeyValue },
            { no: 4, name: "dropped_attributes_count", kind: "scalar", T: 13 /*ScalarType.UINT32*/ }
        ]);
    }
    create(value?: PartialMessage<InstrumentationScope>): InstrumentationScope {
        const message = { name: "", version: "", attributes: [], droppedAttributesCount: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<InstrumentationScope>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: InstrumentationScope): InstrumentationScope {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string name */ 1:
                    message.name = reader.string();
                    break;
                case /* string version */ 2:
                    message.version = reader.string();
                    break;
                case /* repeated opentelemetry.proto.common.v1.KeyValue attributes */ 3:
                    message.attributes.push(KeyValue.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* uint32 dropped_attributes_count */ 4:
                    message.droppedAttributesCount = reader.uint32();
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
    internalBinaryWrite(message: InstrumentationScope, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string name = 1; */
        if (message.name !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.name);
        /* string version = 2; */
        if (message.version !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.version);
        /* repeated opentelemetry.proto.common.v1.KeyValue attributes = 3; */
        for (let i = 0; i < message.attributes.length; i++)
            KeyValue.internalBinaryWrite(message.attributes[i], writer.tag(3, WireType.LengthDelimited).fork(), options).join();
        /* uint32 dropped_attributes_count = 4; */
        if (message.droppedAttributesCount !== 0)
            writer.tag(4, WireType.Varint).uint32(message.droppedAttributesCount);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message opentelemetry.proto.common.v1.InstrumentationScope
 */
export const InstrumentationScope = new InstrumentationScope$Type();
