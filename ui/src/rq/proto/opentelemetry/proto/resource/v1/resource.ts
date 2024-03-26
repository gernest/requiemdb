// @generated by protobuf-ts 2.9.1 with parameter generate_dependencies,long_type_number
// @generated from protobuf file "opentelemetry/proto/resource/v1/resource.proto" (package "opentelemetry.proto.resource.v1", syntax proto3)
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
import { KeyValue } from "../../common/v1/common";
/**
 * Resource information.
 *
 * @generated from protobuf message opentelemetry.proto.resource.v1.Resource
 */
export interface Resource {
    /**
     * Set of attributes that describe the resource.
     * Attribute keys MUST be unique (it is not allowed to have more than one
     * attribute with the same key).
     *
     * @generated from protobuf field: repeated opentelemetry.proto.common.v1.KeyValue attributes = 1;
     */
    attributes: KeyValue[];
    /**
     * dropped_attributes_count is the number of dropped attributes. If the value is 0, then
     * no attributes were dropped.
     *
     * @generated from protobuf field: uint32 dropped_attributes_count = 2;
     */
    droppedAttributesCount: number;
}
// @generated message type with reflection information, may provide speed optimized methods
class Resource$Type extends MessageType<Resource> {
    constructor() {
        super("opentelemetry.proto.resource.v1.Resource", [
            { no: 1, name: "attributes", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => KeyValue },
            { no: 2, name: "dropped_attributes_count", kind: "scalar", T: 13 /*ScalarType.UINT32*/ }
        ]);
    }
    create(value?: PartialMessage<Resource>): Resource {
        const message = { attributes: [], droppedAttributesCount: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Resource>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Resource): Resource {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated opentelemetry.proto.common.v1.KeyValue attributes */ 1:
                    message.attributes.push(KeyValue.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* uint32 dropped_attributes_count */ 2:
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
    internalBinaryWrite(message: Resource, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* repeated opentelemetry.proto.common.v1.KeyValue attributes = 1; */
        for (let i = 0; i < message.attributes.length; i++)
            KeyValue.internalBinaryWrite(message.attributes[i], writer.tag(1, WireType.LengthDelimited).fork(), options).join();
        /* uint32 dropped_attributes_count = 2; */
        if (message.droppedAttributesCount !== 0)
            writer.tag(2, WireType.Varint).uint32(message.droppedAttributesCount);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message opentelemetry.proto.resource.v1.Resource
 */
export const Resource = new Resource$Type();
