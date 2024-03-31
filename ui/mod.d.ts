declare module  '@requiemdb/rq'{
 /**
 * AggregationTemporality defines how a metric aggregator reports aggregated
 * values. It describes how those values relate to the time interval over
 * which they are aggregated.
 *
 * @generated from protobuf enum opentelemetry.proto.metrics.v1.AggregationTemporality
 */
export  enum AggregationTemporality {
    /**
     * UNSPECIFIED is the default AggregationTemporality, it MUST not be used.
     *
     * @generated from protobuf enum value: AGGREGATION_TEMPORALITY_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * DELTA is an AggregationTemporality for a metric aggregator which reports
     * changes since last report time. Successive metrics contain aggregation of
     * values from continuous and non-overlapping intervals.
     *
     * The values for a DELTA metric are based only on the time interval
     * associated with one measurement cycle. There is no dependency on
     * previous measurements like is the case for CUMULATIVE metrics.
     *
     * For example, consider a system measuring the number of requests that
     * it receives and reports the sum of these requests every second as a
     * DELTA metric:
     *
     *   1. The system starts receiving at time=t_0.
     *   2. A request is received, the system measures 1 request.
     *   3. A request is received, the system measures 1 request.
     *   4. A request is received, the system measures 1 request.
     *   5. The 1 second collection cycle ends. A metric is exported for the
     *      number of requests received over the interval of time t_0 to
     *      t_0+1 with a value of 3.
     *   6. A request is received, the system measures 1 request.
     *   7. A request is received, the system measures 1 request.
     *   8. The 1 second collection cycle ends. A metric is exported for the
     *      number of requests received over the interval of time t_0+1 to
     *      t_0+2 with a value of 2.
     *
     * @generated from protobuf enum value: AGGREGATION_TEMPORALITY_DELTA = 1;
     */
    DELTA = 1,
    /**
     * CUMULATIVE is an AggregationTemporality for a metric aggregator which
     * reports changes since a fixed start time. This means that current values
     * of a CUMULATIVE metric depend on all previous measurements since the
     * start time. Because of this, the sender is required to retain this state
     * in some form. If this state is lost or invalidated, the CUMULATIVE metric
     * values MUST be reset and a new fixed start time following the last
     * reported measurement time sent MUST be used.
     *
     * For example, consider a system measuring the number of requests that
     * it receives and reports the sum of these requests every second as a
     * CUMULATIVE metric:
     *
     *   1. The system starts receiving at time=t_0.
     *   2. A request is received, the system measures 1 request.
     *   3. A request is received, the system measures 1 request.
     *   4. A request is received, the system measures 1 request.
     *   5. The 1 second collection cycle ends. A metric is exported for the
     *      number of requests received over the interval of time t_0 to
     *      t_0+1 with a value of 3.
     *   6. A request is received, the system measures 1 request.
     *   7. A request is received, the system measures 1 request.
     *   8. The 1 second collection cycle ends. A metric is exported for the
     *      number of requests received over the interval of time t_0 to
     *      t_0+2 with a value of 5.
     *   9. The system experiences a fault and loses state.
     *   10. The system recovers and resumes receiving at time=t_1.
     *   11. A request is received, the system measures 1 request.
     *   12. The 1 second collection cycle ends. A metric is exported for the
     *      number of requests received over the interval of time t_1 to
     *      t_0+1 with a value of 1.
     *
     * Note: Even though, when reporting changes since last report time, using
     * CUMULATIVE is valid, it is not recommended. This may cause problems for
     * systems that do not use start_time to determine when the aggregation
     * value was reset (e.g. Prometheus).
     *
     * @generated from protobuf enum value: AGGREGATION_TEMPORALITY_CUMULATIVE = 2;
     */
    CUMULATIVE = 2
}

export class AnyValue$Type extends MessageType<AnyValue> {
    constructor();
    create(value?: PartialMessage<AnyValue>): AnyValue;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: AnyValue): AnyValue;
    internalBinaryWrite(message: AnyValue, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * AnyValue is used to represent any type of attribute value. AnyValue may contain a
 * primitive value such as a string or integer or it may contain an arbitrary nested
 * object containing arrays, key-value lists and primitives.
 *
 * @generated from protobuf message opentelemetry.proto.common.v1.AnyValue
 */
export  interface AnyValue {
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
        intValue: number;
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
 * @generated MessageType for protobuf message opentelemetry.proto.common.v1.AnyValue
 */
export  const AnyValue: AnyValue$Type;

export class ArrayValue$Type extends MessageType<ArrayValue> {
    constructor();
    create(value?: PartialMessage<ArrayValue>): ArrayValue;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ArrayValue): ArrayValue;
    internalBinaryWrite(message: ArrayValue, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * ArrayValue is a list of AnyValue messages. We need ArrayValue as a message
 * since oneof in AnyValue does not allow repeated fields.
 *
 * @generated from protobuf message opentelemetry.proto.common.v1.ArrayValue
 */
export  interface ArrayValue {
    /**
     * Array of values. The array may be empty (contain 0 elements).
     *
     * @generated from protobuf field: repeated opentelemetry.proto.common.v1.AnyValue values = 1;
     */
    values: AnyValue[];
}

/**
 * @generated MessageType for protobuf message opentelemetry.proto.common.v1.ArrayValue
 */
export  const ArrayValue: ArrayValue$Type;

/**
 * Options for reading binary data.
 */
export interface BinaryReadOptions {
    /**
     * Shall unknown fields be read, ignored or raise an error?
     *
     * `true`: stores the unknown field on a symbol property of the
     * message. This is the default behaviour.
     *
     * `false`: ignores the unknown field.
     *
     * `"throw"`: throws an error.
     *
     * `UnknownFieldReader`: Your own behaviour for unknown fields.
     */
    readUnknownField: boolean | 'throw' | UnknownFieldReader;
    /**
     * Allows to use a custom implementation to parse binary data.
     */
    readerFactory: (bytes: Uint8Array) => IBinaryReader;
}

/**
 * Options for writing binary data.
 */
export interface BinaryWriteOptions {
    /**
     * Shall unknown fields be written back on wire?
     *
     * `true`: unknown fields stored in a symbol property of the message
     * are written back. This is the default behaviour.
     *
     * `false`: unknown fields are not written.
     *
     * `UnknownFieldWriter`: Your own behaviour for unknown fields.
     */
    writeUnknownFields: boolean | UnknownFieldWriter;
    /**
     * Allows to use a custom implementation to encode binary data.
     */
    writerFactory: () => IBinaryWriter;
}

export  class Config {
    base: Scan;
    constructor(scope: Scan_SCOPE);
    scan(): ScanData;
    limit(num_samples: number): this;
    reverse(): this;
    resourceSchema(schema: string): this;
    scopeSchema(schema: string): this;
    scopeName(name: string): this;
    scopeVersion(version: string): this;
    name(value: string): this;
    traceId(value: string): this;
    spanId(value: string): this;
    parentSpanId(value: string): this;
    logLevel(value: string): this;
    resourceAttr(key: string, value: string): this;
    scopeAttr(key: string, value: string): this;
    attr(key: string, value: string): this;
    private baseFilter;
    private attrFilter;
    filter(f: Scan_Filter): this;
    /**
     *
     * @returns samples for the last 15 minutes
     */
    latest(): this;
    today(): this;
    thisWeek(): this;
    thisYear(): this;
    /**
     *
     * @param duration is ISO 8601 duration string
     * @returns
     */
    ago(duration: string): this;
    thisMonth(): this;
    protected setRange(ts: any): this;
    private createTimeRange;
}

export class Data$Type extends MessageType<Data> {
    constructor();
}

/**
 * @generated from protobuf message v1.Data
 */
export  interface Data {
    /**
     * @generated from protobuf oneof: data
     */
    data: {
        oneofKind: "metrics";
        /**
         * @generated from protobuf field: opentelemetry.proto.metrics.v1.MetricsData metrics = 1;
         */
        metrics: MetricsData;
    } | {
        oneofKind: "logs";
        /**
         * @generated from protobuf field: opentelemetry.proto.logs.v1.LogsData logs = 2;
         */
        logs: LogsData;
    } | {
        oneofKind: "trace";
        /**
         * @generated from protobuf field: opentelemetry.proto.trace.v1.TracesData trace = 3;
         */
        trace: TracesData;
    } | {
        oneofKind: undefined;
    };
}

/**
 * @generated MessageType for protobuf message v1.Data
 */
export  const Data: Data$Type;

/**
 * DataPointFlags is defined as a protobuf 'uint32' type and is to be used as a
 * bit-field representing 32 distinct boolean flags.  Each flag defined in this
 * enum is a bit-mask.  To test the presence of a single flag in the flags of
 * a data point, for example, use an expression like:
 *
 *   (point.flags & DATA_POINT_FLAGS_NO_RECORDED_VALUE_MASK) == DATA_POINT_FLAGS_NO_RECORDED_VALUE_MASK
 *
 *
 * @generated from protobuf enum opentelemetry.proto.metrics.v1.DataPointFlags
 */
export  enum DataPointFlags {
    /**
     * The zero value for the enum. Should not be used for comparisons.
     * Instead use bitwise "and" with the appropriate mask as shown above.
     *
     * @generated from protobuf enum value: DATA_POINT_FLAGS_DO_NOT_USE = 0;
     */
    DO_NOT_USE = 0,
    /**
     * This DataPoint is valid but has no recorded value.  This value
     * SHOULD be used to reflect explicitly missing data in a series, as
     * for an equivalent to the Prometheus "staleness marker".
     *
     * @generated from protobuf enum value: DATA_POINT_FLAGS_NO_RECORDED_VALUE_MASK = 1;
     */
    NO_RECORDED_VALUE_MASK = 1
}

/**
 * Describes a protobuf enum for runtime reflection.
 *
 * The tuple consists of:
 *
 *
 * [0] the protobuf type name
 *
 * The type name follows the same rules as message type names.
 * See `MessageInfo` for details.
 *
 *
 * [1] the enum object generated by Typescript
 *
 * We generate standard Typescript enums for protobuf enums. They are compiled
 * to lookup objects that map from numerical value to name strings and vice
 * versa and can also contain alias names.
 *
 * See https://www.typescriptlang.org/docs/handbook/enums.html#reverse-mappings
 *
 * We use this lookup feature to when encoding / decoding JSON format. The
 * enum is guaranteed to have a value for 0. We generate an entry for 0 if
 * none was declared in .proto because we would need to support custom default
 * values if we didn't.
 *
 *
 * [2] the prefix shared by all original enum values (optional)
 *
 * If all values of a protobuf enum share a prefix, it is dropped in the
 * generated enum. For example, the protobuf enum `enum My { MY_FOO, MY_BAR }`
 * becomes the typescript enum `enum My { FOO, BAR }`.
 *
 * Because the JSON format requires the original value name, we store the
 * dropped prefix here, so that the JSON format implementation can restore
 * the original value names.
 */
export type EnumInfo = readonly [
/**
* The protobuf type name of the enum
*/
string, 
/**
* The enum object generated by Typescript
*/
    {
    [key: number]: string;
    [k: string]: number | string;
}, 
/**
* The prefix shared by all original enum values
*/
string?];

export class Exemplar$Type extends MessageType<Exemplar> {
    constructor();
    create(value?: PartialMessage<Exemplar>): Exemplar;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Exemplar): Exemplar;
    internalBinaryWrite(message: Exemplar, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * A representation of an exemplar, which is a sample input measurement.
 * Exemplars also hold information about the environment when the measurement
 * was recorded, for example the span and trace ID of the active span when the
 * exemplar was recorded.
 *
 * @generated from protobuf message opentelemetry.proto.metrics.v1.Exemplar
 */
export  interface Exemplar {
    /**
     * The set of key/value pairs that were filtered out by the aggregator, but
     * recorded alongside the original measurement. Only key/value pairs that were
     * filtered out by the aggregator should be included
     *
     * @generated from protobuf field: repeated opentelemetry.proto.common.v1.KeyValue filtered_attributes = 7;
     */
    filteredAttributes: KeyValue[];
    /**
     * time_unix_nano is the exact time when this exemplar was recorded
     *
     * Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January
     * 1970.
     *
     * @generated from protobuf field: fixed64 time_unix_nano = 2;
     */
    timeUnixNano: number;
    /**
     * @generated from protobuf oneof: value
     */
    value: {
        oneofKind: "asDouble";
        /**
         * @generated from protobuf field: double as_double = 3;
         */
        asDouble: number;
    } | {
        oneofKind: "asInt";
        /**
         * @generated from protobuf field: sfixed64 as_int = 6;
         */
        asInt: number;
    } | {
        oneofKind: undefined;
    };
    /**
     * (Optional) Span ID of the exemplar trace.
     * span_id may be missing if the measurement is not recorded inside a trace
     * or if the trace is not sampled.
     *
     * @generated from protobuf field: bytes span_id = 4;
     */
    spanId: Uint8Array;
    /**
     * (Optional) Trace ID of the exemplar trace.
     * trace_id may be missing if the measurement is not recorded inside a trace
     * or if the trace is not sampled.
     *
     * @generated from protobuf field: bytes trace_id = 5;
     */
    traceId: Uint8Array;
}

/**
 * @generated MessageType for protobuf message opentelemetry.proto.metrics.v1.Exemplar
 */
export  const Exemplar: Exemplar$Type;

export class ExponentialHistogram$Type extends MessageType<ExponentialHistogram> {
    constructor();
    create(value?: PartialMessage<ExponentialHistogram>): ExponentialHistogram;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ExponentialHistogram): ExponentialHistogram;
    internalBinaryWrite(message: ExponentialHistogram, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * ExponentialHistogram represents the type of a metric that is calculated by aggregating
 * as a ExponentialHistogram of all reported double measurements over a time interval.
 *
 * @generated from protobuf message opentelemetry.proto.metrics.v1.ExponentialHistogram
 */
export  interface ExponentialHistogram {
    /**
     * @generated from protobuf field: repeated opentelemetry.proto.metrics.v1.ExponentialHistogramDataPoint data_points = 1;
     */
    dataPoints: ExponentialHistogramDataPoint[];
    /**
     * aggregation_temporality describes if the aggregator reports delta changes
     * since last report time, or cumulative changes since a fixed start time.
     *
     * @generated from protobuf field: opentelemetry.proto.metrics.v1.AggregationTemporality aggregation_temporality = 2;
     */
    aggregationTemporality: AggregationTemporality;
}

/**
 * @generated MessageType for protobuf message opentelemetry.proto.metrics.v1.ExponentialHistogram
 */
export  const ExponentialHistogram: ExponentialHistogram$Type;

export class ExponentialHistogramDataPoint$Type extends MessageType<ExponentialHistogramDataPoint> {
    constructor();
    create(value?: PartialMessage<ExponentialHistogramDataPoint>): ExponentialHistogramDataPoint;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ExponentialHistogramDataPoint): ExponentialHistogramDataPoint;
    internalBinaryWrite(message: ExponentialHistogramDataPoint, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * ExponentialHistogramDataPoint is a single data point in a timeseries that describes the
 * time-varying values of a ExponentialHistogram of double values. A ExponentialHistogram contains
 * summary statistics for a population of values, it may optionally contain the
 * distribution of those values across a set of buckets.
 *
 *
 * @generated from protobuf message opentelemetry.proto.metrics.v1.ExponentialHistogramDataPoint
 */
export  interface ExponentialHistogramDataPoint {
    /**
     * The set of key/value pairs that uniquely identify the timeseries from
     * where this point belongs. The list may be empty (may contain 0 elements).
     * Attribute keys MUST be unique (it is not allowed to have more than one
     * attribute with the same key).
     *
     * @generated from protobuf field: repeated opentelemetry.proto.common.v1.KeyValue attributes = 1;
     */
    attributes: KeyValue[];
    /**
     * StartTimeUnixNano is optional but strongly encouraged, see the
     * the detailed comments above Metric.
     *
     * Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January
     * 1970.
     *
     * @generated from protobuf field: fixed64 start_time_unix_nano = 2;
     */
    startTimeUnixNano: number;
    /**
     * TimeUnixNano is required, see the detailed comments above Metric.
     *
     * Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January
     * 1970.
     *
     * @generated from protobuf field: fixed64 time_unix_nano = 3;
     */
    timeUnixNano: number;
    /**
     * count is the number of values in the population. Must be
     * non-negative. This value must be equal to the sum of the "bucket_counts"
     * values in the positive and negative Buckets plus the "zero_count" field.
     *
     * @generated from protobuf field: fixed64 count = 4;
     */
    count: number;
    /**
     * sum of the values in the population. If count is zero then this field
     * must be zero.
     *
     * Note: Sum should only be filled out when measuring non-negative discrete
     * events, and is assumed to be monotonic over the values of these events.
     * Negative events *can* be recorded, but sum should not be filled out when
     * doing so.  This is specifically to enforce compatibility w/ OpenMetrics,
     * see: https://github.com/OpenObservability/OpenMetrics/blob/main/specification/OpenMetrics.md#histogram
     *
     * @generated from protobuf field: optional double sum = 5;
     */
    sum?: number;
    /**
     * scale describes the resolution of the histogram.  Boundaries are
     * located at powers of the base, where:
     *
     *   base = (2^(2^-scale))
     *
     * The histogram bucket identified by `index`, a signed integer,
     * contains values that are greater than (base^index) and
     * less than or equal to (base^(index+1)).
     *
     * The positive and negative ranges of the histogram are expressed
     * separately.  Negative values are mapped by their absolute value
     * into the negative range using the same scale as the positive range.
     *
     * scale is not restricted by the protocol, as the permissible
     * values depend on the range of the data.
     *
     * @generated from protobuf field: sint32 scale = 6;
     */
    scale: number;
    /**
     * zero_count is the count of values that are either exactly zero or
     * within the region considered zero by the instrumentation at the
     * tolerated degree of precision.  This bucket stores values that
     * cannot be expressed using the standard exponential formula as
     * well as values that have been rounded to zero.
     *
     * Implementations MAY consider the zero bucket to have probability
     * mass equal to (zero_count / count).
     *
     * @generated from protobuf field: fixed64 zero_count = 7;
     */
    zeroCount: number;
    /**
     * positive carries the positive range of exponential bucket counts.
     *
     * @generated from protobuf field: opentelemetry.proto.metrics.v1.ExponentialHistogramDataPoint.Buckets positive = 8;
     */
    positive?: ExponentialHistogramDataPoint_Buckets;
    /**
     * negative carries the negative range of exponential bucket counts.
     *
     * @generated from protobuf field: opentelemetry.proto.metrics.v1.ExponentialHistogramDataPoint.Buckets negative = 9;
     */
    negative?: ExponentialHistogramDataPoint_Buckets;
    /**
     * Flags that apply to this specific data point.  See DataPointFlags
     * for the available flags and their meaning.
     *
     * @generated from protobuf field: uint32 flags = 10;
     */
    flags: number;
    /**
     * (Optional) List of exemplars collected from
     * measurements that were used to form the data point
     *
     * @generated from protobuf field: repeated opentelemetry.proto.metrics.v1.Exemplar exemplars = 11;
     */
    exemplars: Exemplar[];
    /**
     * min is the minimum value over (start_time, end_time].
     *
     * @generated from protobuf field: optional double min = 12;
     */
    min?: number;
    /**
     * max is the maximum value over (start_time, end_time].
     *
     * @generated from protobuf field: optional double max = 13;
     */
    max?: number;
    /**
     * ZeroThreshold may be optionally set to convey the width of the zero
     * region. Where the zero region is defined as the closed interval
     * [-ZeroThreshold, ZeroThreshold].
     * When ZeroThreshold is 0, zero count bucket stores values that cannot be
     * expressed using the standard exponential formula as well as values that
     * have been rounded to zero.
     *
     * @generated from protobuf field: double zero_threshold = 14;
     */
    zeroThreshold: number;
}

/**
 * @generated MessageType for protobuf message opentelemetry.proto.metrics.v1.ExponentialHistogramDataPoint
 */
export  const ExponentialHistogramDataPoint: ExponentialHistogramDataPoint$Type;

export class ExponentialHistogramDataPoint_Buckets$Type extends MessageType<ExponentialHistogramDataPoint_Buckets> {
    constructor();
    create(value?: PartialMessage<ExponentialHistogramDataPoint_Buckets>): ExponentialHistogramDataPoint_Buckets;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ExponentialHistogramDataPoint_Buckets): ExponentialHistogramDataPoint_Buckets;
    internalBinaryWrite(message: ExponentialHistogramDataPoint_Buckets, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * Buckets are a set of bucket counts, encoded in a contiguous array
 * of counts.
 *
 * @generated from protobuf message opentelemetry.proto.metrics.v1.ExponentialHistogramDataPoint.Buckets
 */
export  interface ExponentialHistogramDataPoint_Buckets {
    /**
     * Offset is the bucket index of the first entry in the bucket_counts array.
     *
     * Note: This uses a varint encoding as a simple form of compression.
     *
     * @generated from protobuf field: sint32 offset = 1;
     */
    offset: number;
    /**
     * bucket_counts is an array of count values, where bucket_counts[i] carries
     * the count of the bucket at index (offset+i). bucket_counts[i] is the count
     * of values greater than base^(offset+i) and less than or equal to
     * base^(offset+i+1).
     *
     * Note: By contrast, the explicit HistogramDataPoint uses
     * fixed64.  This field is expected to have many buckets,
     * especially zeros, so uint64 has been selected to ensure
     * varint encoding.
     *
     * @generated from protobuf field: repeated uint64 bucket_counts = 2;
     */
    bucketCounts: number[];
}

/**
 * @generated MessageType for protobuf message opentelemetry.proto.metrics.v1.ExponentialHistogramDataPoint.Buckets
 */
export  const ExponentialHistogramDataPoint_Buckets: ExponentialHistogramDataPoint_Buckets$Type;

/**
 * Describes a field of a protobuf message for runtime
 * reflection. We distinguish between the following
 * kinds of fields:
 *
 * "scalar": string, bool, float, int32, etc.
 * See https://developers.google.com/protocol-buffers/docs/proto3#scalar
 *
 * "enum": field was declared with an enum type.
 *
 * "message": field was declared with a message type.
 *
 * "map": field was declared with map<K,V>.
 *
 *
 * Every field, regardless of it's kind, always has the following properties:
 *
 * "no": The field number of the .proto field.
 * "name": The original name of the .proto field.
 * "localName": The name of the field as used in generated code.
 * "jsonName": The name for JSON serialization / deserialization.
 * "options": Custom field options from the .proto source in JSON format.
 *
 *
 * Other properties:
 *
 * - Fields of kind "scalar", "enum" and "message" can have a "repeat" type.
 * - Fields of kind "scalar" and "enum" can have a "repeat" type.
 * - Fields of kind "scalar", "enum" and "message" can be member of a "oneof".
 *
 * A field can be only have one of the above properties set.
 *
 * Options for "scalar" fields:
 *
 * - 64 bit integral types can provide "L" - the JavaScript representation
 *   type.
 *
 */
export type FieldInfo = fiRules<fiScalar> | fiRules<fiEnum> | fiRules<fiMessage> | fiRules<fiMap>;

export interface fiEnum extends fiShared {
    kind: 'enum';
    /**
     * Enum type information for the field.
     */
    T: () => EnumInfo;
    /**
     * Is the field repeated?
     */
    repeat: RepeatType;
    /**
     * Is the field optional?
     */
    opt: boolean;
}

export interface fiMap extends fiShared {
    kind: 'map';
    /**
     * Map key type.
     *
     * The key_type can be any integral or string type
     * (so, any scalar type except for floating point
     * types and bytes)
     */
    K: Exclude<ScalarType, ScalarType.FLOAT | ScalarType.DOUBLE | ScalarType.BYTES>;
    /**
     * Map value type. Can be a `ScalarType`, enum type information,
     * or type handler for a message.
     */
    V: {
        kind: 'scalar';
        T: ScalarType;
        L?: LongType;
    } | {
        kind: 'enum';
        T: () => EnumInfo;
    } | {
        kind: 'message';
        T: () => IMessageType<any>;
    };
}

export interface fiMessage extends fiShared {
    kind: 'message';
    /**
     * Message handler for the field.
     */
    T: () => IMessageType<any>;
    /**
     * Is the field repeated?
     */
    repeat: RepeatType;
}

export type fiPartialRules<T> = Omit<T, 'jsonName' | 'localName' | 'oneof' | 'repeat' | 'opt'> & ({
    localName?: string;
    jsonName?: string;
    repeat?: RepeatType.NO;
    opt?: false;
    oneof?: undefined;
} | {
    localName?: string;
    jsonName?: string;
    repeat?: RepeatType.NO;
    opt: true;
    oneof?: undefined;
} | {
    localName?: string;
    jsonName?: string;
    repeat: RepeatType.PACKED | RepeatType.UNPACKED;
    opt?: false;
    oneof?: undefined;
} | {
    localName?: string;
    jsonName?: string;
    repeat?: RepeatType.NO;
    opt?: false;
    oneof: string;
});

export type fiRules<T> = Omit<T, 'oneof' | 'repeat' | 'opt'> & ({
    repeat: RepeatType.NO;
    opt: false;
    oneof: undefined;
} | {
    repeat: RepeatType.NO;
    opt: true;
    oneof: undefined;
} | {
    repeat: RepeatType.PACKED | RepeatType.UNPACKED;
    opt: false;
    oneof: undefined;
} | {
    repeat: RepeatType.NO;
    opt: false;
    oneof: string;
});

export interface fiScalar extends fiShared {
    kind: 'scalar';
    /**
     * Scalar type of the field.
     */
    T: ScalarType;
    /**
     * Representation of 64 bit integral types (int64, uint64, sint64,
     * fixed64, sfixed64).
     *
     * If this option is set for other scalar types, it is ignored.
     * Omitting this option is equivalent to `STRING`.
     */
    L?: LongType;
    /**
     * Is the field repeated?
     */
    repeat: RepeatType;
    /**
     * Is the field optional?
     */
    opt: boolean;
}

export interface fiShared {
    /**
     * The field number of the .proto field.
     */
    no: number;
    /**
     * The original name of the .proto field.
     */
    name: string;
    /**
     * The name of the field as used in generated code.
     */
    localName: string;
    /**
     * The name for JSON serialization / deserialization.
     */
    jsonName: string;
    /**
     * The name of the `oneof` group, if this field belongs to one.
     */
    oneof: string | undefined;
    /**
     * Contains custom field options from the .proto source in JSON format.
     */
    options?: {
        [extensionName: string]: JsonValue;
    };
}

export class Gauge$Type extends MessageType<Gauge> {
    constructor();
    create(value?: PartialMessage<Gauge>): Gauge;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Gauge): Gauge;
    internalBinaryWrite(message: Gauge, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * Gauge represents the type of a scalar metric that always exports the
 * "current value" for every data point. It should be used for an "unknown"
 * aggregation.
 *
 * A Gauge does not support different aggregation temporalities. Given the
 * aggregation is unknown, points cannot be combined using the same
 * aggregation, regardless of aggregation temporalities. Therefore,
 * AggregationTemporality is not included. Consequently, this also means
 * "StartTimeUnixNano" is ignored for all data points.
 *
 * @generated from protobuf message opentelemetry.proto.metrics.v1.Gauge
 */
export  interface Gauge {
    /**
     * @generated from protobuf field: repeated opentelemetry.proto.metrics.v1.NumberDataPoint data_points = 1;
     */
    dataPoints: NumberDataPoint[];
}

/**
 * @generated MessageType for protobuf message opentelemetry.proto.metrics.v1.Gauge
 */
export  const Gauge: Gauge$Type;

export class Histogram$Type extends MessageType<Histogram> {
    constructor();
    create(value?: PartialMessage<Histogram>): Histogram;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Histogram): Histogram;
    internalBinaryWrite(message: Histogram, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * Histogram represents the type of a metric that is calculated by aggregating
 * as a Histogram of all reported measurements over a time interval.
 *
 * @generated from protobuf message opentelemetry.proto.metrics.v1.Histogram
 */
export  interface Histogram {
    /**
     * @generated from protobuf field: repeated opentelemetry.proto.metrics.v1.HistogramDataPoint data_points = 1;
     */
    dataPoints: HistogramDataPoint[];
    /**
     * aggregation_temporality describes if the aggregator reports delta changes
     * since last report time, or cumulative changes since a fixed start time.
     *
     * @generated from protobuf field: opentelemetry.proto.metrics.v1.AggregationTemporality aggregation_temporality = 2;
     */
    aggregationTemporality: AggregationTemporality;
}

/**
 * @generated MessageType for protobuf message opentelemetry.proto.metrics.v1.Histogram
 */
export  const Histogram: Histogram$Type;

export class HistogramDataPoint$Type extends MessageType<HistogramDataPoint> {
    constructor();
    create(value?: PartialMessage<HistogramDataPoint>): HistogramDataPoint;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: HistogramDataPoint): HistogramDataPoint;
    internalBinaryWrite(message: HistogramDataPoint, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * HistogramDataPoint is a single data point in a timeseries that describes the
 * time-varying values of a Histogram. A Histogram contains summary statistics
 * for a population of values, it may optionally contain the distribution of
 * those values across a set of buckets.
 *
 * If the histogram contains the distribution of values, then both
 * "explicit_bounds" and "bucket counts" fields must be defined.
 * If the histogram does not contain the distribution of values, then both
 * "explicit_bounds" and "bucket_counts" must be omitted and only "count" and
 * "sum" are known.
 *
 * @generated from protobuf message opentelemetry.proto.metrics.v1.HistogramDataPoint
 */
export  interface HistogramDataPoint {
    /**
     * The set of key/value pairs that uniquely identify the timeseries from
     * where this point belongs. The list may be empty (may contain 0 elements).
     * Attribute keys MUST be unique (it is not allowed to have more than one
     * attribute with the same key).
     *
     * @generated from protobuf field: repeated opentelemetry.proto.common.v1.KeyValue attributes = 9;
     */
    attributes: KeyValue[];
    /**
     * StartTimeUnixNano is optional but strongly encouraged, see the
     * the detailed comments above Metric.
     *
     * Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January
     * 1970.
     *
     * @generated from protobuf field: fixed64 start_time_unix_nano = 2;
     */
    startTimeUnixNano: number;
    /**
     * TimeUnixNano is required, see the detailed comments above Metric.
     *
     * Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January
     * 1970.
     *
     * @generated from protobuf field: fixed64 time_unix_nano = 3;
     */
    timeUnixNano: number;
    /**
     * count is the number of values in the population. Must be non-negative. This
     * value must be equal to the sum of the "count" fields in buckets if a
     * histogram is provided.
     *
     * @generated from protobuf field: fixed64 count = 4;
     */
    count: number;
    /**
     * sum of the values in the population. If count is zero then this field
     * must be zero.
     *
     * Note: Sum should only be filled out when measuring non-negative discrete
     * events, and is assumed to be monotonic over the values of these events.
     * Negative events *can* be recorded, but sum should not be filled out when
     * doing so.  This is specifically to enforce compatibility w/ OpenMetrics,
     * see: https://github.com/OpenObservability/OpenMetrics/blob/main/specification/OpenMetrics.md#histogram
     *
     * @generated from protobuf field: optional double sum = 5;
     */
    sum?: number;
    /**
     * bucket_counts is an optional field contains the count values of histogram
     * for each bucket.
     *
     * The sum of the bucket_counts must equal the value in the count field.
     *
     * The number of elements in bucket_counts array must be by one greater than
     * the number of elements in explicit_bounds array.
     *
     * @generated from protobuf field: repeated fixed64 bucket_counts = 6;
     */
    bucketCounts: number[];
    /**
     * explicit_bounds specifies buckets with explicitly defined bounds for values.
     *
     * The boundaries for bucket at index i are:
     *
     * (-infinity, explicit_bounds[i]] for i == 0
     * (explicit_bounds[i-1], explicit_bounds[i]] for 0 < i < size(explicit_bounds)
     * (explicit_bounds[i-1], +infinity) for i == size(explicit_bounds)
     *
     * The values in the explicit_bounds array must be strictly increasing.
     *
     * Histogram buckets are inclusive of their upper boundary, except the last
     * bucket where the boundary is at infinity. This format is intentionally
     * compatible with the OpenMetrics histogram definition.
     *
     * @generated from protobuf field: repeated double explicit_bounds = 7;
     */
    explicitBounds: number[];
    /**
     * (Optional) List of exemplars collected from
     * measurements that were used to form the data point
     *
     * @generated from protobuf field: repeated opentelemetry.proto.metrics.v1.Exemplar exemplars = 8;
     */
    exemplars: Exemplar[];
    /**
     * Flags that apply to this specific data point.  See DataPointFlags
     * for the available flags and their meaning.
     *
     * @generated from protobuf field: uint32 flags = 10;
     */
    flags: number;
    /**
     * min is the minimum value over (start_time, end_time].
     *
     * @generated from protobuf field: optional double min = 11;
     */
    min?: number;
    /**
     * max is the maximum value over (start_time, end_time].
     *
     * @generated from protobuf field: optional double max = 12;
     */
    max?: number;
}

/**
 * @generated MessageType for protobuf message opentelemetry.proto.metrics.v1.HistogramDataPoint
 */
export  const HistogramDataPoint: HistogramDataPoint$Type;

/**
 * This interface is used throughout @protobuf-ts to read
 * protobuf binary format.
 *
 * While not completely compatible, this interface is closely aligned
 * with the `Reader` class of `protobufjs` to make it easier to swap
 * the implementation.
 */
export interface IBinaryReader {
    /**
     * Current position.
     */
    readonly pos: number;
    /**
     * Number of bytes available in this reader.
     */
    readonly len: number;
    /**
     * Reads a tag - field number and wire type.
     */
    tag(): [number, WireType];
    /**
     * Skip one element on the wire and return the skipped data.
     */
    skip(wireType: WireType): Uint8Array;
    /**
     * Read a `int32` field, a signed 32 bit varint.
     */
    uint32(): number;
    /**
     * Read a `sint32` field, a signed, zigzag-encoded 32-bit varint.
     */
    int32(): number;
    /**
     * Read a `sint32` field, a signed, zigzag-encoded 32-bit varint.
     */
    sint32(): number;
    /**
     * Read a `int64` field, a signed 64-bit varint.
     */
    int64(): PbLong;
    /**
     * Read a `sint64` field, a signed, zig-zag-encoded 64-bit varint.
     */
    sint64(): PbLong;
    /**
     * Read a `fixed64` field, a signed, fixed-length 64-bit integer.
     */
    sfixed64(): PbLong;
    /**
     * Read a `uint64` field, an unsigned 64-bit varint.
     */
    uint64(): PbULong;
    /**
     * Read a `fixed64` field, an unsigned, fixed-length 64 bit integer.
     */
    fixed64(): PbULong;
    /**
     * Read a `bool` field, a variant.
     */
    bool(): boolean;
    /**
     * Read a `fixed32` field, an unsigned, fixed-length 32-bit integer.
     */
    fixed32(): number;
    /**
     * Read a `sfixed32` field, a signed, fixed-length 32-bit integer.
     */
    sfixed32(): number;
    /**
     * Read a `float` field, 32-bit floating point number.
     */
    float(): number;
    /**
     * Read a `double` field, a 64-bit floating point number.
     */
    double(): number;
    /**
     * Read a `bytes` field, length-delimited arbitrary data.
     */
    bytes(): Uint8Array;
    /**
     * Read a `string` field, length-delimited data converted to UTF-8 text.
     */
    string(): string;
}

/**
 * This interface is used throughout @protobuf-ts to write
 * protobuf binary format.
 *
 * While not completely compatible, this interface is closely aligned
 * with the `Writer` class of `protobufjs` to make it easier to swap
 * the implementation.
 */
export interface IBinaryWriter {
    /**
     * Return all bytes written and reset this writer.
     */
    finish(): Uint8Array;
    /**
     * Start a new fork for length-delimited data like a message
     * or a packed repeated field.
     *
     * Must be joined later with `join()`.
     */
    fork(): IBinaryWriter;
    /**
     * Join the last fork. Write its length and bytes, then
     * return to the previous state.
     */
    join(): IBinaryWriter;
    /**
     * Writes a tag (field number and wire type).
     *
     * Equivalent to `uint32( (fieldNo << 3 | type) >>> 0 )`
     *
     * Generated code should compute the tag ahead of time and call `uint32()`.
     */
    tag(fieldNo: number, type: WireType): IBinaryWriter;
    /**
     * Write a chunk of raw bytes.
     */
    raw(chunk: Uint8Array): IBinaryWriter;
    /**
     * Write a `uint32` value, an unsigned 32 bit varint.
     */
    uint32(value: number): IBinaryWriter;
    /**
     * Write a `int32` value, a signed 32 bit varint.
     */
    int32(value: number): IBinaryWriter;
    /**
     * Write a `sint32` value, a signed, zigzag-encoded 32-bit varint.
     */
    sint32(value: number): IBinaryWriter;
    /**
     * Write a `int64` value, a signed 64-bit varint.
     */
    int64(value: string | number | bigint): IBinaryWriter;
    /**
     * Write a `uint64` value, an unsigned 64-bit varint.
     */
    uint64(value: string | number | bigint): IBinaryWriter;
    /**
     * Write a `sint64` value, a signed, zig-zag-encoded 64-bit varint.
     */
    sint64(value: string | number | bigint): IBinaryWriter;
    /**
     * Write a `fixed64` value, an unsigned, fixed-length 64 bit integer.
     */
    fixed64(value: string | number | bigint): IBinaryWriter;
    /**
     * Write a `fixed64` value, a signed, fixed-length 64-bit integer.
     */
    sfixed64(value: string | number | bigint): IBinaryWriter;
    /**
     * Write a `bool` value, a variant.
     */
    bool(value: boolean): IBinaryWriter;
    /**
     * Write a `fixed32` value, an unsigned, fixed-length 32-bit integer.
     */
    fixed32(value: number): IBinaryWriter;
    /**
     * Write a `sfixed32` value, a signed, fixed-length 32-bit integer.
     */
    sfixed32(value: number): IBinaryWriter;
    /**
     * Write a `float` value, 32-bit floating point number.
     */
    float(value: number): IBinaryWriter;
    /**
     * Write a `double` value, a 64-bit floating point number.
     */
    double(value: number): IBinaryWriter;
    /**
     * Write a `bytes` value, length-delimited arbitrary data.
     */
    bytes(value: Uint8Array): IBinaryWriter;
    /**
     * Write a `string` value, length-delimited data converted to UTF-8 text.
     */
    string(value: string): IBinaryWriter;
}

/**
 * A message type provides an API to work with messages of a specific type.
 * It also exposes reflection information that can be used to work with a
 * message of unknown type.
 */
export interface IMessageType<T extends object> extends MessageInfo {
    /**
     * The protobuf type name of the message, including package and
     * parent types if present.
     *
     * Examples:
     * 'MyNamespaceLessMessage'
     * 'my_package.MyMessage'
     * 'my_package.ParentMessage.ChildMessage'
     */
    readonly typeName: string;
    /**
     * Simple information for each message field, in the order
     * of declaration in the .proto.
     */
    readonly fields: readonly FieldInfo[];
    /**
     * Contains custom message options from the .proto source in JSON format.
     */
    readonly options: {
        [extensionName: string]: JsonValue;
    };
    /**
     * Contains the prototype for messages returned by create() which
     * includes the `MESSAGE_TYPE` symbol pointing back to `this`.
     */
    readonly messagePrototype?: Readonly<{}> | undefined;
    /**
     * Create a new message with default values.
     *
     * For example, a protobuf `string name = 1;` has the default value `""`.
     */
    create(): T;
    /**
     * Create a new message from partial data.
     *
     * Unknown fields are discarded.
     *
     * `PartialMessage<T>` is similar to `Partial<T>`,
     * but it is recursive, and it keeps `oneof` groups
     * intact.
     */
    create(value: PartialMessage<T>): T;
    /**
     * Create a new message from binary format.
     */
    fromBinary(data: Uint8Array, options?: Partial<BinaryReadOptions>): T;
    /**
     * Write the message to binary format.
     */
    toBinary(message: T, options?: Partial<BinaryWriteOptions>): Uint8Array;
    /**
     * Read a new message from a JSON value.
     */
    fromJson(json: JsonValue, options?: Partial<JsonReadOptions>): T;
    /**
     * Read a new message from a JSON string.
     * This is equivalent to `T.fromJson(JSON.parse(json))`.
     */
    fromJsonString(json: string, options?: Partial<JsonReadOptions>): T;
    /**
     * Convert the message to canonical JSON value.
     */
    toJson(message: T, options?: Partial<JsonWriteOptions>): JsonValue;
    /**
     * Convert the message to canonical JSON string.
     * This is equivalent to `JSON.stringify(T.toJson(t))`
     */
    toJsonString(message: T, options?: Partial<JsonWriteStringOptions>): string;
    /**
     * Clone the message.
     *
     * Unknown fields are discarded.
     */
    clone(message: T): T;
    /**
     * Copy partial data into the target message.
     *
     * If a singular scalar or enum field is present in the source, it
     * replaces the field in the target.
     *
     * If a singular message field is present in the source, it is merged
     * with the target field by calling mergePartial() of the responsible
     * message type.
     *
     * If a repeated field is present in the source, its values replace
     * all values in the target array, removing extraneous values.
     * Repeated message fields are copied, not merged.
     *
     * If a map field is present in the source, entries are added to the
     * target map, replacing entries with the same key. Entries that only
     * exist in the target remain. Entries with message values are copied,
     * not merged.
     *
     * Note that this function differs from protobuf merge semantics,
     * which appends repeated fields.
     */
    mergePartial(target: T, source: PartialMessage<T>): void;
    /**
     * Determines whether two message of the same type have the same field values.
     * Checks for deep equality, traversing repeated fields, oneof groups, maps
     * and messages recursively.
     * Will also return true if both messages are `undefined`.
     */
    equals(a: T | undefined, b: T | undefined): boolean;
    /**
     * Is the given value assignable to our message type
     * and contains no [excess properties](https://www.typescriptlang.org/docs/handbook/interfaces.html#excess-property-checks)?
     */
    is(arg: any, depth?: number): arg is T;
    /**
     * Is the given value assignable to our message type,
     * regardless of [excess properties](https://www.typescriptlang.org/docs/handbook/interfaces.html#excess-property-checks)?
     */
    isAssignable(arg: any, depth?: number): arg is T;
    /**
     * This is an internal method. If you just want to read a message from
     * JSON, use `fromJson()` or `fromJsonString()`.
     *
     * Reads JSON value and merges the fields into the target
     * according to protobuf rules. If the target is omitted,
     * a new instance is created first.
     */
    internalJsonRead(json: JsonValue, options: JsonReadOptions, target?: T): T;
    /**
     * This is an internal method. If you just want to write a message
     * to JSON, use `toJson()` or `toJsonString().
     *
     * Writes JSON value and returns it.
     */
    internalJsonWrite(message: T, options: JsonWriteOptions): JsonValue;
    /**
     * This is an internal method. If you just want to write a message
     * in binary format, use `toBinary()`.
     *
     * Serializes the message in binary format and appends it to the given
     * writer. Returns passed writer.
     */
    internalBinaryWrite(message: T, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
    /**
     * This is an internal method. If you just want to read a message from
     * binary data, use `fromBinary()`.
     *
     * Reads data from binary format and merges the fields into
     * the target according to protobuf rules. If the target is
     * omitted, a new instance is created first.
     */
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: T): T;
}

export class InstrumentationScope$Type extends MessageType<InstrumentationScope> {
    constructor();
    create(value?: PartialMessage<InstrumentationScope>): InstrumentationScope;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: InstrumentationScope): InstrumentationScope;
    internalBinaryWrite(message: InstrumentationScope, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * InstrumentationScope is a message representing the instrumentation scope information
 * such as the fully qualified name and version.
 *
 * @generated from protobuf message opentelemetry.proto.common.v1.InstrumentationScope
 */
export  interface InstrumentationScope {
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

/**
 * @generated MessageType for protobuf message opentelemetry.proto.common.v1.InstrumentationScope
 */
export  const InstrumentationScope: InstrumentationScope$Type;

export interface JsonArray extends Array<JsonValue> {
}

/**
 * Represents a JSON object.
 */
export type JsonObject = {
    [k: string]: JsonValue;
};

export type JsonOptionsMap = {
    [extensionName: string]: JsonValue;
};

/**
 * Options for parsing JSON data.
 * All boolean options default to `false`.
 */
export interface JsonReadOptions {
    /**
     * Ignore unknown fields: Proto3 JSON parser should reject unknown fields
     * by default. This option ignores unknown fields in parsing, as well as
     * unrecognized enum string representations.
     */
    ignoreUnknownFields: boolean;
    /**
     * This option is required to read `google.protobuf.Any`
     * from JSON format.
     */
    typeRegistry?: readonly IMessageType<any>[];
}

/**
 * Represents any possible JSON value:
 * - number
 * - string
 * - boolean
 * - null
 * - object (with any JSON value as property)
 * - array (with any JSON value as element)
 */
export type JsonValue = number | string | boolean | null | JsonObject | JsonArray;

/**
 * Options for serializing to JSON object.
 * All boolean options default to `false`.
 */
export interface JsonWriteOptions {
    /**
     * Emit fields with default values: Fields with default values are omitted
     * by default in proto3 JSON output. This option overrides this behavior
     * and outputs fields with their default values.
     */
    emitDefaultValues: boolean;
    /**
     * Emit enum values as integers instead of strings: The name of an enum
     * value is used by default in JSON output. An option may be provided to
     * use the numeric value of the enum value instead.
     */
    enumAsInteger: boolean;
    /**
     * Use proto field name instead of lowerCamelCase name: By default proto3
     * JSON printer should convert the field name to lowerCamelCase and use
     * that as the JSON name. An implementation may provide an option to use
     * proto field name as the JSON name instead. Proto3 JSON parsers are
     * required to accept both the converted lowerCamelCase name and the proto
     * field name.
     */
    useProtoFieldName: boolean;
    /**
     * This option is required to write `google.protobuf.Any`
     * to JSON format.
     */
    typeRegistry?: readonly IMessageType<any>[];
}

/**
 * Options for serializing to JSON string.
 * All options default to `false` or `0`.
 */
export interface JsonWriteStringOptions extends JsonWriteOptions {
    prettySpaces: number;
}

export class KeyValue$Type extends MessageType<KeyValue> {
    constructor();
    create(value?: PartialMessage<KeyValue>): KeyValue;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: KeyValue): KeyValue;
    internalBinaryWrite(message: KeyValue, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * KeyValue is a key-value pair that is used to store Span attributes, Link
 * attributes, etc.
 *
 * @generated from protobuf message opentelemetry.proto.common.v1.KeyValue
 */
export  interface KeyValue {
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
 * @generated MessageType for protobuf message opentelemetry.proto.common.v1.KeyValue
 */
export  const KeyValue: KeyValue$Type;

export class KeyValueList$Type extends MessageType<KeyValueList> {
    constructor();
    create(value?: PartialMessage<KeyValueList>): KeyValueList;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: KeyValueList): KeyValueList;
    internalBinaryWrite(message: KeyValueList, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
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
export  interface KeyValueList {
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
 * @generated MessageType for protobuf message opentelemetry.proto.common.v1.KeyValueList
 */
export  const KeyValueList: KeyValueList$Type;

export class ListValue$Type extends MessageType<ListValue> {
    constructor();
    /**
     * Encode `ListValue` to JSON array.
     */
    internalJsonWrite(message: ListValue, options: JsonWriteOptions): JsonValue;
    /**
     * Decode `ListValue` from JSON array.
     */
    internalJsonRead(json: JsonValue, options: JsonReadOptions, target?: ListValue): ListValue;
}

/**
 * `ListValue` is a wrapper around a repeated field of values.
 *
 * The JSON representation for `ListValue` is JSON array.
 *
 * @generated from protobuf message google.protobuf.ListValue
 */
export interface ListValue {
    /**
     * Repeated field of dynamically typed values.
     *
     * @generated from protobuf field: repeated google.protobuf.Value values = 1;
     */
    values: Value[];
}

/**
 * @generated MessageType for protobuf message google.protobuf.ListValue
 */
export const ListValue: ListValue$Type;

export class LogRecord$Type extends MessageType<LogRecord> {
    constructor();
    create(value?: PartialMessage<LogRecord>): LogRecord;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: LogRecord): LogRecord;
    internalBinaryWrite(message: LogRecord, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * A log record according to OpenTelemetry Log Data Model:
 * https://github.com/open-telemetry/oteps/blob/main/text/logs/0097-log-data-model.md
 *
 * @generated from protobuf message opentelemetry.proto.logs.v1.LogRecord
 */
export  interface LogRecord {
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
 * @generated MessageType for protobuf message opentelemetry.proto.logs.v1.LogRecord
 */
export  const LogRecord: LogRecord$Type;

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
export  enum LogRecordFlags {
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

export  class Logs extends Config {
    constructor();
}

export class LogsData$Type extends MessageType<LogsData> {
    constructor();
    create(value?: PartialMessage<LogsData>): LogsData;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: LogsData): LogsData;
    internalBinaryWrite(message: LogsData, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

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
export  interface LogsData {
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
 * @generated MessageType for protobuf message opentelemetry.proto.logs.v1.LogsData
 */
export  const LogsData: LogsData$Type;

/**
 * JavaScript representation of 64 bit integral types. Equivalent to the
 * field option "jstype".
 *
 * By default, protobuf-ts represents 64 bit types as `bigint`.
 *
 * You can change the default behaviour by enabling the plugin parameter
 * `long_type_string`, which will represent 64 bit types as `string`.
 *
 * Alternatively, you can change the behaviour for individual fields
 * with the field option "jstype":
 *
 * ```protobuf
 * uint64 my_field = 1 [jstype = JS_STRING];
 * uint64 other_field = 2 [jstype = JS_NUMBER];
 * ```
 */
export enum LongType {
    /**
     * Use JavaScript `bigint`.
     *
     * Field option `[jstype = JS_NORMAL]`.
     */
    BIGINT = 0,
    /**
     * Use JavaScript `string`.
     *
     * Field option `[jstype = JS_STRING]`.
     */
    STRING = 1,
    /**
     * Use JavaScript `number`.
     *
     * Large values will loose precision.
     *
     * Field option `[jstype = JS_NUMBER]`.
     */
    NUMBER = 2
}

/**
 * Describes a protobuf message for runtime reflection.
 */
export interface MessageInfo {
    /**
     * The protobuf type name of the message, including package and
     * parent types if present.
     *
     * If the .proto file included a `package` statement, the type name
     * starts with '.'.
     *
     * Examples:
     * 'MyNamespaceLessMessage'
     * '.my_package.MyMessage'
     * '.my_package.ParentMessage.ChildMessage'
     */
    readonly typeName: string;
    /**
     * Simple information for each message field, in the order
     * of declaration in the source .proto.
     */
    readonly fields: readonly FieldInfo[];
    /**
     * Contains custom message options from the .proto source in JSON format.
     */
    readonly options: {
        [extensionName: string]: JsonValue;
    };
}

/**
 * This standard message type provides reflection-based
 * operations to work with a message.
 */
export class MessageType<T extends object> implements IMessageType<T> {
    /**
     * The protobuf type name of the message, including package and
     * parent types if present.
     *
     * If the .proto file included a `package` statement,
     * the type name will always start with a '.'.
     *
     * Examples:
     * 'MyNamespaceLessMessage'
     * '.my_package.MyMessage'
     * '.my_package.ParentMessage.ChildMessage'
     */
    readonly typeName: string;
    /**
     * Simple information for each message field, in the order
     * of declaration in the .proto.
     */
    readonly fields: readonly FieldInfo[];
    /**
     * Contains custom service options from the .proto source in JSON format.
     */
    readonly options: JsonOptionsMap;
    /**
     * Contains the prototype for messages returned by create() which
     * includes the `MESSAGE_TYPE` symbol pointing back to `this`.
     */
    readonly messagePrototype?: Readonly<{}> | undefined;
    protected readonly defaultCheckDepth = 16;
    protected readonly refTypeCheck: ReflectionTypeCheck;
    protected readonly refJsonReader: ReflectionJsonReader;
    protected readonly refJsonWriter: ReflectionJsonWriter;
    protected readonly refBinReader: ReflectionBinaryReader;
    protected readonly refBinWriter: ReflectionBinaryWriter;
    constructor(name: string, fields: readonly PartialFieldInfo[], options?: JsonOptionsMap);
    /**
     * Create a new message with default values.
     *
     * For example, a protobuf `string name = 1;` has the default value `""`.
     */
    create(): T;
    /**
     * Create a new message from partial data.
     * Where a field is omitted, the default value is used.
     *
     * Unknown fields are discarded.
     *
     * `PartialMessage<T>` is similar to `Partial<T>`,
     * but it is recursive, and it keeps `oneof` groups
     * intact.
     */
    create(value: PartialMessage<T>): T;
    /**
     * Clone the message.
     *
     * Unknown fields are discarded.
     */
    clone(message: T): T;
    /**
     * Determines whether two message of the same type have the same field values.
     * Checks for deep equality, traversing repeated fields, oneof groups, maps
     * and messages recursively.
     * Will also return true if both messages are `undefined`.
     */
    equals(a: T | undefined, b: T | undefined): boolean;
    /**
     * Is the given value assignable to our message type
     * and contains no [excess properties](https://www.typescriptlang.org/docs/handbook/interfaces.html#excess-property-checks)?
     */
    is(arg: any, depth?: number): arg is T;
    /**
     * Is the given value assignable to our message type,
     * regardless of [excess properties](https://www.typescriptlang.org/docs/handbook/interfaces.html#excess-property-checks)?
     */
    isAssignable(arg: any, depth?: number): arg is T;
    /**
     * Copy partial data into the target message.
     */
    mergePartial(target: T, source: PartialMessage<T>): void;
    /**
     * Create a new message from binary format.
     */
    fromBinary(data: Uint8Array, options?: Partial<BinaryReadOptions>): T;
    /**
     * Read a new message from a JSON value.
     */
    fromJson(json: JsonValue, options?: Partial<JsonReadOptions>): T;
    /**
     * Read a new message from a JSON string.
     * This is equivalent to `T.fromJson(JSON.parse(json))`.
     */
    fromJsonString(json: string, options?: Partial<JsonReadOptions>): T;
    /**
     * Write the message to canonical JSON value.
     */
    toJson(message: T, options?: Partial<JsonWriteOptions>): JsonValue;
    /**
     * Convert the message to canonical JSON string.
     * This is equivalent to `JSON.stringify(T.toJson(t))`
     */
    toJsonString(message: T, options?: Partial<JsonWriteStringOptions>): string;
    /**
     * Write the message to binary format.
     */
    toBinary(message: T, options?: Partial<BinaryWriteOptions>): Uint8Array;
    /**
     * This is an internal method. If you just want to read a message from
     * JSON, use `fromJson()` or `fromJsonString()`.
     *
     * Reads JSON value and merges the fields into the target
     * according to protobuf rules. If the target is omitted,
     * a new instance is created first.
     */
    internalJsonRead(json: JsonValue, options: JsonReadOptions, target?: T): T;
    /**
     * This is an internal method. If you just want to write a message
     * to JSON, use `toJson()` or `toJsonString().
     *
     * Writes JSON value and returns it.
     */
    internalJsonWrite(message: T, options: JsonWriteOptions): JsonValue;
    /**
     * This is an internal method. If you just want to write a message
     * in binary format, use `toBinary()`.
     *
     * Serializes the message in binary format and appends it to the given
     * writer. Returns passed writer.
     */
    internalBinaryWrite(message: T, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
    /**
     * This is an internal method. If you just want to read a message from
     * binary data, use `fromBinary()`.
     *
     * Reads data from binary format and merges the fields into
     * the target according to protobuf rules. If the target is
     * omitted, a new instance is created first.
     */
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: T): T;
}

export class Metric$Type extends MessageType<Metric> {
    constructor();
    create(value?: PartialMessage<Metric>): Metric;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Metric): Metric;
    internalBinaryWrite(message: Metric, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * Defines a Metric which has one or more timeseries.  The following is a
 * brief summary of the Metric data model.  For more details, see:
 *
 *   https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/metrics/data-model.md
 *
 *
 * The data model and relation between entities is shown in the
 * diagram below. Here, "DataPoint" is the term used to refer to any
 * one of the specific data point value types, and "points" is the term used
 * to refer to any one of the lists of points contained in the Metric.
 *
 * - Metric is composed of a metadata and data.
 * - Metadata part contains a name, description, unit.
 * - Data is one of the possible types (Sum, Gauge, Histogram, Summary).
 * - DataPoint contains timestamps, attributes, and one of the possible value type
 *   fields.
 *
 *     Metric
 *  +------------+
 *  |name        |
 *  |description |
 *  |unit        |     +------------------------------------+
 *  |data        |---> |Gauge, Sum, Histogram, Summary, ... |
 *  +------------+     +------------------------------------+
 *
 *    Data [One of Gauge, Sum, Histogram, Summary, ...]
 *  +-----------+
 *  |...        |  // Metadata about the Data.
 *  |points     |--+
 *  +-----------+  |
 *                 |      +---------------------------+
 *                 |      |DataPoint 1                |
 *                 v      |+------+------+   +------+ |
 *              +-----+   ||label |label |...|label | |
 *              |  1  |-->||value1|value2|...|valueN| |
 *              +-----+   |+------+------+   +------+ |
 *              |  .  |   |+-----+                    |
 *              |  .  |   ||value|                    |
 *              |  .  |   |+-----+                    |
 *              |  .  |   +---------------------------+
 *              |  .  |                   .
 *              |  .  |                   .
 *              |  .  |                   .
 *              |  .  |   +---------------------------+
 *              |  .  |   |DataPoint M                |
 *              +-----+   |+------+------+   +------+ |
 *              |  M  |-->||label |label |...|label | |
 *              +-----+   ||value1|value2|...|valueN| |
 *                        |+------+------+   +------+ |
 *                        |+-----+                    |
 *                        ||value|                    |
 *                        |+-----+                    |
 *                        +---------------------------+
 *
 * Each distinct type of DataPoint represents the output of a specific
 * aggregation function, the result of applying the DataPoint's
 * associated function of to one or more measurements.
 *
 * All DataPoint types have three common fields:
 * - Attributes includes key-value pairs associated with the data point
 * - TimeUnixNano is required, set to the end time of the aggregation
 * - StartTimeUnixNano is optional, but strongly encouraged for DataPoints
 *   having an AggregationTemporality field, as discussed below.
 *
 * Both TimeUnixNano and StartTimeUnixNano values are expressed as
 * UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
 *
 * # TimeUnixNano
 *
 * This field is required, having consistent interpretation across
 * DataPoint types.  TimeUnixNano is the moment corresponding to when
 * the data point's aggregate value was captured.
 *
 * Data points with the 0 value for TimeUnixNano SHOULD be rejected
 * by consumers.
 *
 * # StartTimeUnixNano
 *
 * StartTimeUnixNano in general allows detecting when a sequence of
 * observations is unbroken.  This field indicates to consumers the
 * start time for points with cumulative and delta
 * AggregationTemporality, and it should be included whenever possible
 * to support correct rate calculation.  Although it may be omitted
 * when the start time is truly unknown, setting StartTimeUnixNano is
 * strongly encouraged.
 *
 * @generated from protobuf message opentelemetry.proto.metrics.v1.Metric
 */
export  interface Metric {
    /**
     * name of the metric.
     *
     * @generated from protobuf field: string name = 1;
     */
    name: string;
    /**
     * description of the metric, which can be used in documentation.
     *
     * @generated from protobuf field: string description = 2;
     */
    description: string;
    /**
     * unit in which the metric value is reported. Follows the format
     * described by http://unitsofmeasure.org/ucum.html.
     *
     * @generated from protobuf field: string unit = 3;
     */
    unit: string;
    /**
     * @generated from protobuf oneof: data
     */
    data: {
        oneofKind: "gauge";
        /**
         * @generated from protobuf field: opentelemetry.proto.metrics.v1.Gauge gauge = 5;
         */
        gauge: Gauge;
    } | {
        oneofKind: "sum";
        /**
         * @generated from protobuf field: opentelemetry.proto.metrics.v1.Sum sum = 7;
         */
        sum: Sum;
    } | {
        oneofKind: "histogram";
        /**
         * @generated from protobuf field: opentelemetry.proto.metrics.v1.Histogram histogram = 9;
         */
        histogram: Histogram;
    } | {
        oneofKind: "exponentialHistogram";
        /**
         * @generated from protobuf field: opentelemetry.proto.metrics.v1.ExponentialHistogram exponential_histogram = 10;
         */
        exponentialHistogram: ExponentialHistogram;
    } | {
        oneofKind: "summary";
        /**
         * @generated from protobuf field: opentelemetry.proto.metrics.v1.Summary summary = 11;
         */
        summary: Summary;
    } | {
        oneofKind: undefined;
    };
    /**
     * Additional metadata attributes that describe the metric. [Optional].
     * Attributes are non-identifying.
     * Consumers SHOULD NOT need to be aware of these attributes.
     * These attributes MAY be used to encode information allowing
     * for lossless roundtrip translation to / from another data model.
     * Attribute keys MUST be unique (it is not allowed to have more than one
     * attribute with the same key).
     *
     * @generated from protobuf field: repeated opentelemetry.proto.common.v1.KeyValue metadata = 12;
     */
    metadata: KeyValue[];
}

/**
 * @generated MessageType for protobuf message opentelemetry.proto.metrics.v1.Metric
 */
export  const Metric: Metric$Type;

export  class Metrics extends Config {
    constructor(name?: string);
    query(): ScanData;
}

export class MetricsData$Type extends MessageType<MetricsData> {
    constructor();
    create(value?: PartialMessage<MetricsData>): MetricsData;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: MetricsData): MetricsData;
    internalBinaryWrite(message: MetricsData, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * MetricsData represents the metrics data that can be stored in a persistent
 * storage, OR can be embedded by other protocols that transfer OTLP metrics
 * data but do not implement the OTLP protocol.
 *
 * The main difference between this message and collector protocol is that
 * in this message there will not be any "control" or "metadata" specific to
 * OTLP protocol.
 *
 * When new fields are added into this message, the OTLP request MUST be updated
 * as well.
 *
 * @generated from protobuf message opentelemetry.proto.metrics.v1.MetricsData
 */
export  interface MetricsData {
    /**
     * An array of ResourceMetrics.
     * For data coming from a single resource this array will typically contain
     * one element. Intermediary nodes that receive data from multiple origins
     * typically batch the data before forwarding further and in that case this
     * array will contain multiple elements.
     *
     * @generated from protobuf field: repeated opentelemetry.proto.metrics.v1.ResourceMetrics resource_metrics = 1;
     */
    resourceMetrics: ResourceMetrics[];
}

/**
 * @generated MessageType for protobuf message opentelemetry.proto.metrics.v1.MetricsData
 */
export  const MetricsData: MetricsData$Type;

/**
 * `NullValue` is a singleton enumeration to represent the null value for the
 * `Value` type union.
 *
 *  The JSON representation for `NullValue` is JSON `null`.
 *
 * @generated from protobuf enum google.protobuf.NullValue
 */
export enum NullValue {
    /**
     * Null value.
     *
     * @generated from protobuf enum value: NULL_VALUE = 0;
     */
    NULL_VALUE = 0
}

export class NumberDataPoint$Type extends MessageType<NumberDataPoint> {
    constructor();
    create(value?: PartialMessage<NumberDataPoint>): NumberDataPoint;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: NumberDataPoint): NumberDataPoint;
    internalBinaryWrite(message: NumberDataPoint, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * NumberDataPoint is a single data point in a timeseries that describes the
 * time-varying scalar value of a metric.
 *
 * @generated from protobuf message opentelemetry.proto.metrics.v1.NumberDataPoint
 */
export  interface NumberDataPoint {
    /**
     * The set of key/value pairs that uniquely identify the timeseries from
     * where this point belongs. The list may be empty (may contain 0 elements).
     * Attribute keys MUST be unique (it is not allowed to have more than one
     * attribute with the same key).
     *
     * @generated from protobuf field: repeated opentelemetry.proto.common.v1.KeyValue attributes = 7;
     */
    attributes: KeyValue[];
    /**
     * StartTimeUnixNano is optional but strongly encouraged, see the
     * the detailed comments above Metric.
     *
     * Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January
     * 1970.
     *
     * @generated from protobuf field: fixed64 start_time_unix_nano = 2;
     */
    startTimeUnixNano: number;
    /**
     * TimeUnixNano is required, see the detailed comments above Metric.
     *
     * Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January
     * 1970.
     *
     * @generated from protobuf field: fixed64 time_unix_nano = 3;
     */
    timeUnixNano: number;
    /**
     * @generated from protobuf oneof: value
     */
    value: {
        oneofKind: "asDouble";
        /**
         * @generated from protobuf field: double as_double = 4;
         */
        asDouble: number;
    } | {
        oneofKind: "asInt";
        /**
         * @generated from protobuf field: sfixed64 as_int = 6;
         */
        asInt: number;
    } | {
        oneofKind: undefined;
    };
    /**
     * (Optional) List of exemplars collected from
     * measurements that were used to form the data point
     *
     * @generated from protobuf field: repeated opentelemetry.proto.metrics.v1.Exemplar exemplars = 5;
     */
    exemplars: Exemplar[];
    /**
     * Flags that apply to this specific data point.  See DataPointFlags
     * for the available flags and their meaning.
     *
     * @generated from protobuf field: uint32 flags = 8;
     */
    flags: number;
}

/**
 * @generated MessageType for protobuf message opentelemetry.proto.metrics.v1.NumberDataPoint
 */
export  const NumberDataPoint: NumberDataPoint$Type;

export type PartialField<T> = T extends (Date | Uint8Array | bigint | boolean | string | number) ? T : T extends Array<infer U> ? Array<PartialField<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<PartialField<U>> : T extends {
    oneofKind: string;
} ? T : T extends {
    oneofKind: undefined;
} ? T : T extends object ? PartialMessage<T> : T;

/**
 * Version of `FieldInfo` that allows the following properties
 * to be omitted:
 * - "localName", "jsonName": can be omitted if equal to lowerCamelCase(name)
 * - "opt": can be omitted if false
 * - "repeat", can be omitted if RepeatType.NO
 *
 * Use `normalizeFieldInfo()` to fill the omitted fields with
 * their standard values.
 */
export type PartialFieldInfo = fiPartialRules<fiScalar> | fiPartialRules<fiEnum> | fiPartialRules<fiMessage> | fiPartialRules<fiMap>;

/**
 * Similar to `Partial<T>`, but recursive, and keeps `oneof` groups
 * intact.
 */
export type PartialMessage<T extends object> = {
    [K in keyof T]?: PartialField<T[K]>;
};

/**
 * Version of `MessageInfo` that allows the following properties
 * to be omitted:
 * - "fields": omitting means the message has no fields
 * - "options": omitting means the message has no options
 */
export type PartialMessageInfo = PartialPartial<MessageInfo, "fields" | "options">;

export type PartialPartial<T, K extends keyof T> = Partial<Pick<T, K>> & Omit<T, K>;

/**
 * 64-bit signed integer as two 32-bit values.
 * Converts between `string`, `number` and `bigint` representations.
 */
export class PbLong extends SharedPbLong {
    /**
     * long 0 singleton.
     */
    static ZERO: PbLong;
    /**
     * Create instance from a `string`, `number` or `bigint`.
     */
    static from(value: string | number | bigint): PbLong;
    /**
     * Do we have a minus sign?
     */
    isNegative(): boolean;
    /**
     * Negate two's complement.
     * Invert all the bits and add one to the result.
     */
    negate(): PbLong;
    /**
     * Convert to decimal string.
     */
    toString(): string;
    /**
     * Convert to native bigint.
     */
    toBigInt(): bigint;
}

/**
 * 64-bit unsigned integer as two 32-bit values.
 * Converts between `string`, `number` and `bigint` representations.
 */
export class PbULong extends SharedPbLong {
    /**
     * ulong 0 singleton.
     */
    static ZERO: PbULong;
    /**
     * Create instance from a `string`, `number` or `bigint`.
     */
    static from(value: string | number | bigint): PbULong;
    /**
     * Convert to decimal string.
     */
    toString(): string;
    /**
     * Convert to native bigint.
     */
    toBigInt(): bigint;
}

/**
 * Reads proto3 messages in binary format using reflection information.
 *
 * https://developers.google.com/protocol-buffers/docs/encoding
 */
export class ReflectionBinaryReader {
    private readonly info;
    protected fieldNoToField?: ReadonlyMap<number, FieldInfo>;
    constructor(info: PartialMessageInfo);
    protected prepare(): void;
    /**
     * Reads a message from binary format into the target message.
     *
     * Repeated fields are appended. Map entries are added, overwriting
     * existing keys.
     *
     * If a message field is already present, it will be merged with the
     * new data.
     */
    read<T extends object>(reader: IBinaryReader, message: T, options: BinaryReadOptions, length?: number): void;
    /**
     * Read a map field, expecting key field = 1, value field = 2
     */
    protected mapEntry(field: FieldInfo & {
        kind: "map";
    }, reader: IBinaryReader, options: BinaryReadOptions): [string | number, UnknownMap[string]];
    protected scalar(reader: IBinaryReader, type: ScalarType, longType: LongType | undefined): UnknownScalar;
}

/**
 * Writes proto3 messages in binary format using reflection information.
 *
 * https://developers.google.com/protocol-buffers/docs/encoding
 */
export class ReflectionBinaryWriter {
    private readonly info;
    protected fields?: readonly FieldInfo[];
    constructor(info: PartialMessageInfo);
    protected prepare(): void;
    /**
     * Writes the message to binary format.
     */
    write<T extends object>(message: T, writer: IBinaryWriter, options: BinaryWriteOptions): void;
    protected mapEntry(writer: IBinaryWriter, options: BinaryWriteOptions, field: FieldInfo & {
        kind: 'map';
    }, key: any, value: any): void;
    protected message(writer: IBinaryWriter, options: BinaryWriteOptions, handler: IMessageType<any>, fieldNo: number, value: any): void;
    /**
     * Write a single scalar value.
     */
    protected scalar(writer: IBinaryWriter, type: ScalarType, fieldNo: number, value: any, emitDefault: boolean): void;
    /**
     * Write an array of scalar values in packed format.
     */
    protected packed(writer: IBinaryWriter, type: ScalarType, fieldNo: number, value: any[]): void;
    /**
     * Get information for writing a scalar value.
     *
     * Returns tuple:
     * [0]: appropriate WireType
     * [1]: name of the appropriate method of IBinaryWriter
     * [2]: whether the given value is a default value
     *
     * If argument `value` is omitted, [2] is always false.
     */
    protected scalarInfo(type: ScalarType, value?: any): [WireType, "int32" | "string" | "bool" | "uint32" | "double" | "float" | "int64" | "uint64" | "fixed64" | "bytes" | "fixed32" | "sfixed32" | "sfixed64" | "sint32" | "sint64", boolean];
}

/**
 * Reads proto3 messages in canonical JSON format using reflection information.
 *
 * https://developers.google.com/protocol-buffers/docs/proto3#json
 */
export class ReflectionJsonReader {
    private readonly info;
    /**
     * JSON key to field.
     * Accepts the original proto field name in the .proto, the
     * lowerCamelCase name, or the name specified by the json_name option.
     */
    private fMap?;
    constructor(info: PartialMessageInfo);
    protected prepare(): void;
    assert(condition: any, fieldName: string, jsonValue: JsonValue): asserts condition;
    /**
     * Reads a message from canonical JSON format into the target message.
     *
     * Repeated fields are appended. Map entries are added, overwriting
     * existing keys.
     *
     * If a message field is already present, it will be merged with the
     * new data.
     */
    read<T extends object>(input: JsonObject, message: T, options: JsonReadOptions): void;
    /**
     * Returns `false` for unrecognized string representations.
     *
     * google.protobuf.NullValue accepts only JSON `null` (or the old `"NULL_VALUE"`).
     */
    enum(type: EnumInfo, json: unknown, fieldName: string, ignoreUnknownFields: boolean): UnknownEnum | false;
    scalar(json: JsonValue, type: ScalarType, longType: LongType | undefined, fieldName: string): UnknownScalar;
}

/**
 * Writes proto3 messages in canonical JSON format using reflection
 * information.
 *
 * https://developers.google.com/protocol-buffers/docs/proto3#json
 */
export class ReflectionJsonWriter {
    private readonly fields;
    constructor(info: PartialMessageInfo);
    /**
     * Converts the message to a JSON object, based on the field descriptors.
     */
    write<T extends object>(message: T, options: JsonWriteOptions): JsonValue;
    field(field: FieldInfo, value: unknown, options: JsonWriteOptions): JsonValue | undefined;
    /**
     * Returns `null` as the default for google.protobuf.NullValue.
     */
    enum(type: EnumInfo, value: unknown, fieldName: string, optional: boolean, emitDefaultValues: boolean, enumAsInteger: boolean): JsonValue | undefined;
    message(type: IMessageType<any>, value: unknown, fieldName: string, options: JsonWriteOptions): JsonValue | undefined;
    scalar(type: ScalarType, value: unknown, fieldName: string, optional: false, emitDefaultValues: boolean): JsonValue;
    scalar(type: ScalarType, value: unknown, fieldName: string, optional: boolean, emitDefaultValues: boolean): JsonValue | undefined;
}

export class ReflectionTypeCheck {
    private readonly fields;
    private data;
    constructor(info: PartialMessageInfo);
    private prepare;
    /**
     * Is the argument a valid message as specified by the
     * reflection information?
     *
     * Checks all field types recursively. The `depth`
     * specifies how deep into the structure the check will be.
     *
     * With a depth of 0, only the presence of fields
     * is checked.
     *
     * With a depth of 1 or more, the field types are checked.
     *
     * With a depth of 2 or more, the members of map, repeated
     * and message fields are checked.
     *
     * Message fields will be checked recursively with depth - 1.
     *
     * The number of map entries / repeated values being checked
     * is < depth.
     */
    is(message: any, depth: number, allowExcessProperties?: boolean): boolean;
    private field;
    private message;
    private messages;
    private scalar;
    private scalars;
    private mapKeys;
}

/**
 * Serialize value and exit the script. This must only be called once, subsequent calls have
 * no effect.
 * @param value
 */
export  const render: (value: Struct | Data | ScanData) => void;

/**
 * Protobuf 2.1.0 introduced packed repeated fields.
 * Setting the field option `[packed = true]` enables packing.
 *
 * In proto3, all repeated fields are packed by default.
 * Setting the field option `[packed = false]` disables packing.
 *
 * Packed repeated fields are encoded with a single tag,
 * then a length-delimiter, then the element values.
 *
 * Unpacked repeated fields are encoded with a tag and
 * value for each element.
 *
 * `bytes` and `string` cannot be packed.
 */
export enum RepeatType {
    /**
     * The field is not repeated.
     */
    NO = 0,
    /**
     * The field is repeated and should be packed.
     * Invalid for `bytes` and `string`, they cannot be packed.
     */
    PACKED = 1,
    /**
     * The field is repeated but should not be packed.
     * The only valid repeat type for repeated `bytes` and `string`.
     */
    UNPACKED = 2
}

export class Resource$Type extends MessageType<Resource> {
    constructor();
    create(value?: PartialMessage<Resource>): Resource;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Resource): Resource;
    internalBinaryWrite(message: Resource, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * Resource information.
 *
 * @generated from protobuf message opentelemetry.proto.resource.v1.Resource
 */
export  interface Resource {
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

/**
 * @generated MessageType for protobuf message opentelemetry.proto.resource.v1.Resource
 */
export  const Resource: Resource$Type;

export class ResourceLogs$Type extends MessageType<ResourceLogs> {
    constructor();
    create(value?: PartialMessage<ResourceLogs>): ResourceLogs;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ResourceLogs): ResourceLogs;
    internalBinaryWrite(message: ResourceLogs, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * A collection of ScopeLogs from a Resource.
 *
 * @generated from protobuf message opentelemetry.proto.logs.v1.ResourceLogs
 */
export  interface ResourceLogs {
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
 * @generated MessageType for protobuf message opentelemetry.proto.logs.v1.ResourceLogs
 */
export  const ResourceLogs: ResourceLogs$Type;

export class ResourceMetrics$Type extends MessageType<ResourceMetrics> {
    constructor();
    create(value?: PartialMessage<ResourceMetrics>): ResourceMetrics;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ResourceMetrics): ResourceMetrics;
    internalBinaryWrite(message: ResourceMetrics, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * A collection of ScopeMetrics from a Resource.
 *
 * @generated from protobuf message opentelemetry.proto.metrics.v1.ResourceMetrics
 */
export  interface ResourceMetrics {
    /**
     * The resource for the metrics in this message.
     * If this field is not set then no resource info is known.
     *
     * @generated from protobuf field: opentelemetry.proto.resource.v1.Resource resource = 1;
     */
    resource?: Resource;
    /**
     * A list of metrics that originate from a resource.
     *
     * @generated from protobuf field: repeated opentelemetry.proto.metrics.v1.ScopeMetrics scope_metrics = 2;
     */
    scopeMetrics: ScopeMetrics[];
    /**
     * The Schema URL, if known. This is the identifier of the Schema that the resource data
     * is recorded in. To learn more about Schema URL see
     * https://opentelemetry.io/docs/specs/otel/schemas/#schema-url
     * This schema_url applies to the data in the "resource" field. It does not apply
     * to the data in the "scope_metrics" field which have their own schema_url field.
     *
     * @generated from protobuf field: string schema_url = 3;
     */
    schemaUrl: string;
}

/**
 * @generated MessageType for protobuf message opentelemetry.proto.metrics.v1.ResourceMetrics
 */
export  const ResourceMetrics: ResourceMetrics$Type;

export class ResourceSpans$Type extends MessageType<ResourceSpans> {
    constructor();
    create(value?: PartialMessage<ResourceSpans>): ResourceSpans;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ResourceSpans): ResourceSpans;
    internalBinaryWrite(message: ResourceSpans, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * A collection of ScopeSpans from a Resource.
 *
 * @generated from protobuf message opentelemetry.proto.trace.v1.ResourceSpans
 */
export  interface ResourceSpans {
    /**
     * The resource for the spans in this message.
     * If this field is not set then no resource info is known.
     *
     * @generated from protobuf field: opentelemetry.proto.resource.v1.Resource resource = 1;
     */
    resource?: Resource;
    /**
     * A list of ScopeSpans that originate from a resource.
     *
     * @generated from protobuf field: repeated opentelemetry.proto.trace.v1.ScopeSpans scope_spans = 2;
     */
    scopeSpans: ScopeSpans[];
    /**
     * The Schema URL, if known. This is the identifier of the Schema that the resource data
     * is recorded in. To learn more about Schema URL see
     * https://opentelemetry.io/docs/specs/otel/schemas/#schema-url
     * This schema_url applies to the data in the "resource" field. It does not apply
     * to the data in the "scope_spans" field which have their own schema_url field.
     *
     * @generated from protobuf field: string schema_url = 3;
     */
    schemaUrl: string;
}

/**
 * @generated MessageType for protobuf message opentelemetry.proto.trace.v1.ResourceSpans
 */
export  const ResourceSpans: ResourceSpans$Type;

/**
 * Scalar value types. This is a subset of field types declared by protobuf
 * enum google.protobuf.FieldDescriptorProto.Type The types GROUP and MESSAGE
 * are omitted, but the numerical values are identical.
 */
export enum ScalarType {
    DOUBLE = 1,
    FLOAT = 2,
    INT64 = 3,
    UINT64 = 4,
    INT32 = 5,
    FIXED64 = 6,
    FIXED32 = 7,
    BOOL = 8,
    STRING = 9,
    BYTES = 12,
    UINT32 = 13,
    SFIXED32 = 15,
    SFIXED64 = 16,
    SINT32 = 17,
    SINT64 = 18
}

export class Scan$Type extends MessageType<Scan> {
    constructor();
}

/**
 * @generated from protobuf message v1.Scan
 */
export  interface Scan {
    /**
     * @generated from protobuf field: v1.Scan.SCOPE scope = 1;
     */
    scope: Scan_SCOPE;
    /**
     * @generated from protobuf field: v1.Scan.TimeRange time_range = 2;
     */
    timeRange?: Scan_TimeRange;
    /**
     * @generated from protobuf field: repeated v1.Scan.Filter filters = 3;
     */
    filters: Scan_Filter[];
    /**
     * Number of samples to process. Defauluts to no limit.
     *
     * @generated from protobuf field: uint64 limit = 4;
     */
    limit: number;
    /**
     * Scans in reverse order, with latest samples comming first.  To get the
     * latest sample you can set reverse to true and limit 1.
     *
     * @generated from protobuf field: bool reverse = 5;
     */
    reverse: boolean;
}

/**
 * @generated MessageType for protobuf message v1.Scan
 */
export  const Scan: Scan$Type;

export class Scan_AttrFilter$Type extends MessageType<Scan_AttrFilter> {
    constructor();
}

/**
 * @generated from protobuf message v1.Scan.AttrFilter
 */
export  interface Scan_AttrFilter {
    /**
     * @generated from protobuf field: v1.Scan.AttributeProp prop = 1;
     */
    prop: Scan_AttributeProp;
    /**
     * @generated from protobuf field: string key = 2;
     */
    key: string;
    /**
     * @generated from protobuf field: string value = 3;
     */
    value: string;
}

/**
 * @generated MessageType for protobuf message v1.Scan.AttrFilter
 */
export  const Scan_AttrFilter: Scan_AttrFilter$Type;

/**
 * @generated from protobuf enum v1.Scan.AttributeProp
 */
export  enum Scan_AttributeProp {
    /**
     * @generated from protobuf enum value: UNKOWN_ATTR = 0;
     */
    UNKOWN_ATTR = 0,
    /**
     * @generated from protobuf enum value: RESOURCE_ATTRIBUTES = 1;
     */
    RESOURCE_ATTRIBUTES = 1,
    /**
     * @generated from protobuf enum value: SCOPE_ATTRIBUTES = 5;
     */
    SCOPE_ATTRIBUTES = 5,
    /**
     * @generated from protobuf enum value: ATTRIBUTES = 7;
     */
    ATTRIBUTES = 7
}

export class Scan_BaseFilter$Type extends MessageType<Scan_BaseFilter> {
    constructor();
}

/**
 * @generated from protobuf message v1.Scan.BaseFilter
 */
export  interface Scan_BaseFilter {
    /**
     * @generated from protobuf field: v1.Scan.BaseProp prop = 1;
     */
    prop: Scan_BaseProp;
    /**
     * @generated from protobuf field: string value = 2;
     */
    value: string;
}

/**
 * @generated MessageType for protobuf message v1.Scan.BaseFilter
 */
export  const Scan_BaseFilter: Scan_BaseFilter$Type;

/**
 * @generated from protobuf enum v1.Scan.BaseProp
 */
export  enum Scan_BaseProp {
    /**
     * @generated from protobuf enum value: RESOURCE_SCHEMA = 0;
     */
    RESOURCE_SCHEMA = 0,
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
     * @generated from protobuf enum value: NAME = 6;
     */
    NAME = 6,
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

export class Scan_Filter$Type extends MessageType<Scan_Filter> {
    constructor();
}

/**
 * @generated from protobuf message v1.Scan.Filter
 */
export  interface Scan_Filter {
    /**
     * @generated from protobuf oneof: value
     */
    value: {
        oneofKind: "base";
        /**
         * @generated from protobuf field: v1.Scan.BaseFilter base = 1;
         */
        base: Scan_BaseFilter;
    } | {
        oneofKind: "attr";
        /**
         * @generated from protobuf field: v1.Scan.AttrFilter attr = 2;
         */
        attr: Scan_AttrFilter;
    } | {
        oneofKind: undefined;
    };
}

/**
 * @generated MessageType for protobuf message v1.Scan.Filter
 */
export  const Scan_Filter: Scan_Filter$Type;

/**
 * @generated from protobuf enum v1.Scan.SCOPE
 */
export  enum Scan_SCOPE {
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

export class Scan_TimeRange$Type extends MessageType<Scan_TimeRange> {
    constructor();
}

/**
 * @generated from protobuf message v1.Scan.TimeRange
 */
export  interface Scan_TimeRange {
    /**
     * @generated from protobuf field: google.protobuf.Timestamp start = 1;
     */
    start?: Timestamp;
    /**
     * @generated from protobuf field: google.protobuf.Timestamp end = 2;
     */
    end?: Timestamp;
}

/**
 * @generated MessageType for protobuf message v1.Scan.TimeRange
 */
export  const Scan_TimeRange: Scan_TimeRange$Type;

export  class ScanData {
    private ptr;
    constructor(ptr: any);
    toData(): Data;
    formData(data: Data): ScanData;
    static is(value: any): boolean;
}

export class ScopeLogs$Type extends MessageType<ScopeLogs> {
    constructor();
    create(value?: PartialMessage<ScopeLogs>): ScopeLogs;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ScopeLogs): ScopeLogs;
    internalBinaryWrite(message: ScopeLogs, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * A collection of Logs produced by a Scope.
 *
 * @generated from protobuf message opentelemetry.proto.logs.v1.ScopeLogs
 */
export  interface ScopeLogs {
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
 * @generated MessageType for protobuf message opentelemetry.proto.logs.v1.ScopeLogs
 */
export  const ScopeLogs: ScopeLogs$Type;

export class ScopeMetrics$Type extends MessageType<ScopeMetrics> {
    constructor();
    create(value?: PartialMessage<ScopeMetrics>): ScopeMetrics;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ScopeMetrics): ScopeMetrics;
    internalBinaryWrite(message: ScopeMetrics, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * A collection of Metrics produced by an Scope.
 *
 * @generated from protobuf message opentelemetry.proto.metrics.v1.ScopeMetrics
 */
export  interface ScopeMetrics {
    /**
     * The instrumentation scope information for the metrics in this message.
     * Semantically when InstrumentationScope isn't set, it is equivalent with
     * an empty instrumentation scope name (unknown).
     *
     * @generated from protobuf field: opentelemetry.proto.common.v1.InstrumentationScope scope = 1;
     */
    scope?: InstrumentationScope;
    /**
     * A list of metrics that originate from an instrumentation library.
     *
     * @generated from protobuf field: repeated opentelemetry.proto.metrics.v1.Metric metrics = 2;
     */
    metrics: Metric[];
    /**
     * The Schema URL, if known. This is the identifier of the Schema that the metric data
     * is recorded in. To learn more about Schema URL see
     * https://opentelemetry.io/docs/specs/otel/schemas/#schema-url
     * This schema_url applies to all metrics in the "metrics" field.
     *
     * @generated from protobuf field: string schema_url = 3;
     */
    schemaUrl: string;
}

/**
 * @generated MessageType for protobuf message opentelemetry.proto.metrics.v1.ScopeMetrics
 */
export  const ScopeMetrics: ScopeMetrics$Type;

export class ScopeSpans$Type extends MessageType<ScopeSpans> {
    constructor();
    create(value?: PartialMessage<ScopeSpans>): ScopeSpans;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ScopeSpans): ScopeSpans;
    internalBinaryWrite(message: ScopeSpans, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * A collection of Spans produced by an InstrumentationScope.
 *
 * @generated from protobuf message opentelemetry.proto.trace.v1.ScopeSpans
 */
export  interface ScopeSpans {
    /**
     * The instrumentation scope information for the spans in this message.
     * Semantically when InstrumentationScope isn't set, it is equivalent with
     * an empty instrumentation scope name (unknown).
     *
     * @generated from protobuf field: opentelemetry.proto.common.v1.InstrumentationScope scope = 1;
     */
    scope?: InstrumentationScope;
    /**
     * A list of Spans that originate from an instrumentation scope.
     *
     * @generated from protobuf field: repeated opentelemetry.proto.trace.v1.Span spans = 2;
     */
    spans: Span[];
    /**
     * The Schema URL, if known. This is the identifier of the Schema that the span data
     * is recorded in. To learn more about Schema URL see
     * https://opentelemetry.io/docs/specs/otel/schemas/#schema-url
     * This schema_url applies to all spans and span events in the "spans" field.
     *
     * @generated from protobuf field: string schema_url = 3;
     */
    schemaUrl: string;
}

/**
 * @generated MessageType for protobuf message opentelemetry.proto.trace.v1.ScopeSpans
 */
export  const ScopeSpans: ScopeSpans$Type;

/**
 * Possible values for LogRecord.SeverityNumber.
 *
 * @generated from protobuf enum opentelemetry.proto.logs.v1.SeverityNumber
 */
export  enum SeverityNumber {
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

export abstract class SharedPbLong {
    /**
     * Low 32 bits.
     */
    readonly lo: number;
    /**
     * High 32 bits.
     */
    readonly hi: number;
    /**
     * Create a new instance with the given bits.
     */
    constructor(lo: number, hi: number);
    /**
     * Is this instance equal to 0?
     */
    isZero(): boolean;
    /**
     * Convert to a native number.
     */
    toNumber(): number;
    /**
     * Convert to decimal string.
     */
    abstract toString(): string;
    /**
     * Convert to native bigint.
     */
    abstract toBigInt(): bigint;
}

export class Span$Type extends MessageType<Span> {
    constructor();
    create(value?: PartialMessage<Span>): Span;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Span): Span;
    internalBinaryWrite(message: Span, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * A Span represents a single operation performed by a single component of the system.
 *
 * The next available field id is 17.
 *
 * @generated from protobuf message opentelemetry.proto.trace.v1.Span
 */
export  interface Span {
    /**
     * A unique identifier for a trace. All spans from the same trace share
     * the same `trace_id`. The ID is a 16-byte array. An ID with all zeroes OR
     * of length other than 16 bytes is considered invalid (empty string in OTLP/JSON
     * is zero-length and thus is also invalid).
     *
     * This field is required.
     *
     * @generated from protobuf field: bytes trace_id = 1;
     */
    traceId: Uint8Array;
    /**
     * A unique identifier for a span within a trace, assigned when the span
     * is created. The ID is an 8-byte array. An ID with all zeroes OR of length
     * other than 8 bytes is considered invalid (empty string in OTLP/JSON
     * is zero-length and thus is also invalid).
     *
     * This field is required.
     *
     * @generated from protobuf field: bytes span_id = 2;
     */
    spanId: Uint8Array;
    /**
     * trace_state conveys information about request position in multiple distributed tracing graphs.
     * It is a trace_state in w3c-trace-context format: https://www.w3.org/TR/trace-context/#tracestate-header
     * See also https://github.com/w3c/distributed-tracing for more details about this field.
     *
     * @generated from protobuf field: string trace_state = 3;
     */
    traceState: string;
    /**
     * The `span_id` of this span's parent span. If this is a root span, then this
     * field must be empty. The ID is an 8-byte array.
     *
     * @generated from protobuf field: bytes parent_span_id = 4;
     */
    parentSpanId: Uint8Array;
    /**
     * Flags, a bit field.
     *
     * Bits 0-7 (8 least significant bits) are the trace flags as defined in W3C Trace
     * Context specification. To read the 8-bit W3C trace flag, use
     * `flags & SPAN_FLAGS_TRACE_FLAGS_MASK`.
     *
     * See https://www.w3.org/TR/trace-context-2/#trace-flags for the flag definitions.
     *
     * Bits 8 and 9 represent the 3 states of whether a span's parent
     * is remote. The states are (unknown, is not remote, is remote).
     * To read whether the value is known, use `(flags & SPAN_FLAGS_CONTEXT_HAS_IS_REMOTE_MASK) != 0`.
     * To read whether the span is remote, use `(flags & SPAN_FLAGS_CONTEXT_IS_REMOTE_MASK) != 0`.
     *
     * When creating span messages, if the message is logically forwarded from another source
     * with an equivalent flags fields (i.e., usually another OTLP span message), the field SHOULD
     * be copied as-is. If creating from a source that does not have an equivalent flags field
     * (such as a runtime representation of an OpenTelemetry span), the high 22 bits MUST
     * be set to zero.
     * Readers MUST NOT assume that bits 10-31 (22 most significant bits) will be zero.
     *
     * [Optional].
     *
     * @generated from protobuf field: fixed32 flags = 16;
     */
    flags: number;
    /**
     * A description of the span's operation.
     *
     * For example, the name can be a qualified method name or a file name
     * and a line number where the operation is called. A best practice is to use
     * the same display name at the same call point in an application.
     * This makes it easier to correlate spans in different traces.
     *
     * This field is semantically required to be set to non-empty string.
     * Empty value is equivalent to an unknown span name.
     *
     * This field is required.
     *
     * @generated from protobuf field: string name = 5;
     */
    name: string;
    /**
     * Distinguishes between spans generated in a particular context. For example,
     * two spans with the same name may be distinguished using `CLIENT` (caller)
     * and `SERVER` (callee) to identify queueing latency associated with the span.
     *
     * @generated from protobuf field: opentelemetry.proto.trace.v1.Span.SpanKind kind = 6;
     */
    kind: Span_SpanKind;
    /**
     * start_time_unix_nano is the start time of the span. On the client side, this is the time
     * kept by the local machine where the span execution starts. On the server side, this
     * is the time when the server's application handler starts running.
     * Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
     *
     * This field is semantically required and it is expected that end_time >= start_time.
     *
     * @generated from protobuf field: fixed64 start_time_unix_nano = 7;
     */
    startTimeUnixNano: number;
    /**
     * end_time_unix_nano is the end time of the span. On the client side, this is the time
     * kept by the local machine where the span execution ends. On the server side, this
     * is the time when the server application handler stops running.
     * Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
     *
     * This field is semantically required and it is expected that end_time >= start_time.
     *
     * @generated from protobuf field: fixed64 end_time_unix_nano = 8;
     */
    endTimeUnixNano: number;
    /**
     * attributes is a collection of key/value pairs. Note, global attributes
     * like server name can be set using the resource API. Examples of attributes:
     *
     *     "/http/user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36"
     *     "/http/server_latency": 300
     *     "example.com/myattribute": true
     *     "example.com/score": 10.239
     *
     * The OpenTelemetry API specification further restricts the allowed value types:
     * https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/common/README.md#attribute
     * Attribute keys MUST be unique (it is not allowed to have more than one
     * attribute with the same key).
     *
     * @generated from protobuf field: repeated opentelemetry.proto.common.v1.KeyValue attributes = 9;
     */
    attributes: KeyValue[];
    /**
     * dropped_attributes_count is the number of attributes that were discarded. Attributes
     * can be discarded because their keys are too long or because there are too many
     * attributes. If this value is 0, then no attributes were dropped.
     *
     * @generated from protobuf field: uint32 dropped_attributes_count = 10;
     */
    droppedAttributesCount: number;
    /**
     * events is a collection of Event items.
     *
     * @generated from protobuf field: repeated opentelemetry.proto.trace.v1.Span.Event events = 11;
     */
    events: Span_Event[];
    /**
     * dropped_events_count is the number of dropped events. If the value is 0, then no
     * events were dropped.
     *
     * @generated from protobuf field: uint32 dropped_events_count = 12;
     */
    droppedEventsCount: number;
    /**
     * links is a collection of Links, which are references from this span to a span
     * in the same or different trace.
     *
     * @generated from protobuf field: repeated opentelemetry.proto.trace.v1.Span.Link links = 13;
     */
    links: Span_Link[];
    /**
     * dropped_links_count is the number of dropped links after the maximum size was
     * enforced. If this value is 0, then no links were dropped.
     *
     * @generated from protobuf field: uint32 dropped_links_count = 14;
     */
    droppedLinksCount: number;
    /**
     * An optional final status for this span. Semantically when Status isn't set, it means
     * span's status code is unset, i.e. assume STATUS_CODE_UNSET (code = 0).
     *
     * @generated from protobuf field: opentelemetry.proto.trace.v1.Status status = 15;
     */
    status?: Status;
}

/**
 * @generated MessageType for protobuf message opentelemetry.proto.trace.v1.Span
 */
export  const Span: Span$Type;

export class Span_Event$Type extends MessageType<Span_Event> {
    constructor();
    create(value?: PartialMessage<Span_Event>): Span_Event;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Span_Event): Span_Event;
    internalBinaryWrite(message: Span_Event, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * Event is a time-stamped annotation of the span, consisting of user-supplied
 * text description and key-value pairs.
 *
 * @generated from protobuf message opentelemetry.proto.trace.v1.Span.Event
 */
export  interface Span_Event {
    /**
     * time_unix_nano is the time the event occurred.
     *
     * @generated from protobuf field: fixed64 time_unix_nano = 1;
     */
    timeUnixNano: number;
    /**
     * name of the event.
     * This field is semantically required to be set to non-empty string.
     *
     * @generated from protobuf field: string name = 2;
     */
    name: string;
    /**
     * attributes is a collection of attribute key/value pairs on the event.
     * Attribute keys MUST be unique (it is not allowed to have more than one
     * attribute with the same key).
     *
     * @generated from protobuf field: repeated opentelemetry.proto.common.v1.KeyValue attributes = 3;
     */
    attributes: KeyValue[];
    /**
     * dropped_attributes_count is the number of dropped attributes. If the value is 0,
     * then no attributes were dropped.
     *
     * @generated from protobuf field: uint32 dropped_attributes_count = 4;
     */
    droppedAttributesCount: number;
}

/**
 * @generated MessageType for protobuf message opentelemetry.proto.trace.v1.Span.Event
 */
export  const Span_Event: Span_Event$Type;

export class Span_Link$Type extends MessageType<Span_Link> {
    constructor();
    create(value?: PartialMessage<Span_Link>): Span_Link;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Span_Link): Span_Link;
    internalBinaryWrite(message: Span_Link, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * A pointer from the current span to another span in the same trace or in a
 * different trace. For example, this can be used in batching operations,
 * where a single batch handler processes multiple requests from different
 * traces or when the handler receives a request from a different project.
 *
 * @generated from protobuf message opentelemetry.proto.trace.v1.Span.Link
 */
export  interface Span_Link {
    /**
     * A unique identifier of a trace that this linked span is part of. The ID is a
     * 16-byte array.
     *
     * @generated from protobuf field: bytes trace_id = 1;
     */
    traceId: Uint8Array;
    /**
     * A unique identifier for the linked span. The ID is an 8-byte array.
     *
     * @generated from protobuf field: bytes span_id = 2;
     */
    spanId: Uint8Array;
    /**
     * The trace_state associated with the link.
     *
     * @generated from protobuf field: string trace_state = 3;
     */
    traceState: string;
    /**
     * attributes is a collection of attribute key/value pairs on the link.
     * Attribute keys MUST be unique (it is not allowed to have more than one
     * attribute with the same key).
     *
     * @generated from protobuf field: repeated opentelemetry.proto.common.v1.KeyValue attributes = 4;
     */
    attributes: KeyValue[];
    /**
     * dropped_attributes_count is the number of dropped attributes. If the value is 0,
     * then no attributes were dropped.
     *
     * @generated from protobuf field: uint32 dropped_attributes_count = 5;
     */
    droppedAttributesCount: number;
    /**
     * Flags, a bit field.
     *
     * Bits 0-7 (8 least significant bits) are the trace flags as defined in W3C Trace
     * Context specification. To read the 8-bit W3C trace flag, use
     * `flags & SPAN_FLAGS_TRACE_FLAGS_MASK`.
     *
     * See https://www.w3.org/TR/trace-context-2/#trace-flags for the flag definitions.
     *
     * Bits 8 and 9 represent the 3 states of whether the link is remote.
     * The states are (unknown, is not remote, is remote).
     * To read whether the value is known, use `(flags & SPAN_FLAGS_CONTEXT_HAS_IS_REMOTE_MASK) != 0`.
     * To read whether the link is remote, use `(flags & SPAN_FLAGS_CONTEXT_IS_REMOTE_MASK) != 0`.
     *
     * Readers MUST NOT assume that bits 10-31 (22 most significant bits) will be zero.
     * When creating new spans, bits 10-31 (most-significant 22-bits) MUST be zero.
     *
     * [Optional].
     *
     * @generated from protobuf field: fixed32 flags = 6;
     */
    flags: number;
}

/**
 * @generated MessageType for protobuf message opentelemetry.proto.trace.v1.Span.Link
 */
export  const Span_Link: Span_Link$Type;

/**
 * SpanKind is the type of span. Can be used to specify additional relationships between spans
 * in addition to a parent/child relationship.
 *
 * @generated from protobuf enum opentelemetry.proto.trace.v1.Span.SpanKind
 */
export  enum Span_SpanKind {
    /**
     * Unspecified. Do NOT use as default.
     * Implementations MAY assume SpanKind to be INTERNAL when receiving UNSPECIFIED.
     *
     * @generated from protobuf enum value: SPAN_KIND_UNSPECIFIED = 0;
     */
    UNSPECIFIED = 0,
    /**
     * Indicates that the span represents an internal operation within an application,
     * as opposed to an operation happening at the boundaries. Default value.
     *
     * @generated from protobuf enum value: SPAN_KIND_INTERNAL = 1;
     */
    INTERNAL = 1,
    /**
     * Indicates that the span covers server-side handling of an RPC or other
     * remote network request.
     *
     * @generated from protobuf enum value: SPAN_KIND_SERVER = 2;
     */
    SERVER = 2,
    /**
     * Indicates that the span describes a request to some remote service.
     *
     * @generated from protobuf enum value: SPAN_KIND_CLIENT = 3;
     */
    CLIENT = 3,
    /**
     * Indicates that the span describes a producer sending a message to a broker.
     * Unlike CLIENT and SERVER, there is often no direct critical path latency relationship
     * between producer and consumer spans. A PRODUCER span ends when the message was accepted
     * by the broker while the logical processing of the message might span a much longer time.
     *
     * @generated from protobuf enum value: SPAN_KIND_PRODUCER = 4;
     */
    PRODUCER = 4,
    /**
     * Indicates that the span describes consumer receiving a message from a broker.
     * Like the PRODUCER kind, there is often no direct critical path latency relationship
     * between producer and consumer spans.
     *
     * @generated from protobuf enum value: SPAN_KIND_CONSUMER = 5;
     */
    CONSUMER = 5
}

/**
 * SpanFlags represents constants used to interpret the
 * Span.flags field, which is protobuf 'fixed32' type and is to
 * be used as bit-fields. Each non-zero value defined in this enum is
 * a bit-mask.  To extract the bit-field, for example, use an
 * expression like:
 *
 *   (span.flags & SPAN_FLAGS_TRACE_FLAGS_MASK)
 *
 * See https://www.w3.org/TR/trace-context-2/#trace-flags for the flag definitions.
 *
 * Note that Span flags were introduced in version 1.1 of the
 * OpenTelemetry protocol.  Older Span producers do not set this
 * field, consequently consumers should not rely on the absence of a
 * particular flag bit to indicate the presence of a particular feature.
 *
 * @generated from protobuf enum opentelemetry.proto.trace.v1.SpanFlags
 */
export  enum SpanFlags {
    /**
     * The zero value for the enum. Should not be used for comparisons.
     * Instead use bitwise "and" with the appropriate mask as shown above.
     *
     * @generated from protobuf enum value: SPAN_FLAGS_DO_NOT_USE = 0;
     */
    DO_NOT_USE = 0,
    /**
     * Bits 0-7 are used for trace flags.
     *
     * @generated from protobuf enum value: SPAN_FLAGS_TRACE_FLAGS_MASK = 255;
     */
    TRACE_FLAGS_MASK = 255,
    /**
     * Bits 8 and 9 are used to indicate that the parent span or link span is remote.
     * Bit 8 (`HAS_IS_REMOTE`) indicates whether the value is known.
     * Bit 9 (`IS_REMOTE`) indicates whether the span or link is remote.
     *
     * @generated from protobuf enum value: SPAN_FLAGS_CONTEXT_HAS_IS_REMOTE_MASK = 256;
     */
    CONTEXT_HAS_IS_REMOTE_MASK = 256,
    /**
     * @generated from protobuf enum value: SPAN_FLAGS_CONTEXT_IS_REMOTE_MASK = 512;
     */
    CONTEXT_IS_REMOTE_MASK = 512
}

export class Status$Type extends MessageType<Status> {
    constructor();
    create(value?: PartialMessage<Status>): Status;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Status): Status;
    internalBinaryWrite(message: Status, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * The Status type defines a logical error model that is suitable for different
 * programming environments, including REST APIs and RPC APIs.
 *
 * @generated from protobuf message opentelemetry.proto.trace.v1.Status
 */
export  interface Status {
    /**
     * A developer-facing human readable error message.
     *
     * @generated from protobuf field: string message = 2;
     */
    message: string;
    /**
     * The status code.
     *
     * @generated from protobuf field: opentelemetry.proto.trace.v1.Status.StatusCode code = 3;
     */
    code: Status_StatusCode;
}

/**
 * @generated MessageType for protobuf message opentelemetry.proto.trace.v1.Status
 */
export  const Status: Status$Type;

/**
 * For the semantics of status codes see
 * https://github.com/open-telemetry/opentelemetry-specification/blob/main/specification/trace/api.md#set-status
 *
 * @generated from protobuf enum opentelemetry.proto.trace.v1.Status.StatusCode
 */
export  enum Status_StatusCode {
    /**
     * The default status.
     *
     * @generated from protobuf enum value: STATUS_CODE_UNSET = 0;
     */
    UNSET = 0,
    /**
     * The Span has been validated by an Application developer or Operator to
     * have completed successfully.
     *
     * @generated from protobuf enum value: STATUS_CODE_OK = 1;
     */
    OK = 1,
    /**
     * The Span contains an error.
     *
     * @generated from protobuf enum value: STATUS_CODE_ERROR = 2;
     */
    ERROR = 2
}

export class Struct$Type extends MessageType<Struct> {
    constructor();
    /**
     * Encode `Struct` to JSON object.
     */
    internalJsonWrite(message: Struct, options: JsonWriteOptions): JsonValue;
    /**
     * Decode `Struct` from JSON object.
     */
    internalJsonRead(json: JsonValue, options: JsonReadOptions, target?: Struct): Struct;
}

/**
 * `Struct` represents a structured data value, consisting of fields
 * which map to dynamically typed values. In some languages, `Struct`
 * might be supported by a native representation. For example, in
 * scripting languages like JS a struct is represented as an
 * object. The details of that representation are described together
 * with the proto support for the language.
 *
 * The JSON representation for `Struct` is JSON object.
 *
 * @generated from protobuf message google.protobuf.Struct
 */
export interface Struct {
    /**
     * Unordered map of dynamically typed values.
     *
     * @generated from protobuf field: map<string, google.protobuf.Value> fields = 1;
     */
    fields: {
        [key: string]: Value;
    };
}

/**
 * @generated MessageType for protobuf message google.protobuf.Struct
 */
export const Struct: Struct$Type;

export class Sum$Type extends MessageType<Sum> {
    constructor();
    create(value?: PartialMessage<Sum>): Sum;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Sum): Sum;
    internalBinaryWrite(message: Sum, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * Sum represents the type of a scalar metric that is calculated as a sum of all
 * reported measurements over a time interval.
 *
 * @generated from protobuf message opentelemetry.proto.metrics.v1.Sum
 */
export  interface Sum {
    /**
     * @generated from protobuf field: repeated opentelemetry.proto.metrics.v1.NumberDataPoint data_points = 1;
     */
    dataPoints: NumberDataPoint[];
    /**
     * aggregation_temporality describes if the aggregator reports delta changes
     * since last report time, or cumulative changes since a fixed start time.
     *
     * @generated from protobuf field: opentelemetry.proto.metrics.v1.AggregationTemporality aggregation_temporality = 2;
     */
    aggregationTemporality: AggregationTemporality;
    /**
     * If "true" means that the sum is monotonic.
     *
     * @generated from protobuf field: bool is_monotonic = 3;
     */
    isMonotonic: boolean;
}

/**
 * @generated MessageType for protobuf message opentelemetry.proto.metrics.v1.Sum
 */
export  const Sum: Sum$Type;

export class Summary$Type extends MessageType<Summary> {
    constructor();
    create(value?: PartialMessage<Summary>): Summary;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Summary): Summary;
    internalBinaryWrite(message: Summary, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * Summary metric data are used to convey quantile summaries,
 * a Prometheus (see: https://prometheus.io/docs/concepts/metric_types/#summary)
 * and OpenMetrics (see: https://github.com/OpenObservability/OpenMetrics/blob/4dbf6075567ab43296eed941037c12951faafb92/protos/prometheus.proto#L45)
 * data type. These data points cannot always be merged in a meaningful way.
 * While they can be useful in some applications, histogram data points are
 * recommended for new applications.
 *
 * @generated from protobuf message opentelemetry.proto.metrics.v1.Summary
 */
export  interface Summary {
    /**
     * @generated from protobuf field: repeated opentelemetry.proto.metrics.v1.SummaryDataPoint data_points = 1;
     */
    dataPoints: SummaryDataPoint[];
}

/**
 * @generated MessageType for protobuf message opentelemetry.proto.metrics.v1.Summary
 */
export  const Summary: Summary$Type;

export class SummaryDataPoint$Type extends MessageType<SummaryDataPoint> {
    constructor();
    create(value?: PartialMessage<SummaryDataPoint>): SummaryDataPoint;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SummaryDataPoint): SummaryDataPoint;
    internalBinaryWrite(message: SummaryDataPoint, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * SummaryDataPoint is a single data point in a timeseries that describes the
 * time-varying values of a Summary metric.
 *
 * @generated from protobuf message opentelemetry.proto.metrics.v1.SummaryDataPoint
 */
export  interface SummaryDataPoint {
    /**
     * The set of key/value pairs that uniquely identify the timeseries from
     * where this point belongs. The list may be empty (may contain 0 elements).
     * Attribute keys MUST be unique (it is not allowed to have more than one
     * attribute with the same key).
     *
     * @generated from protobuf field: repeated opentelemetry.proto.common.v1.KeyValue attributes = 7;
     */
    attributes: KeyValue[];
    /**
     * StartTimeUnixNano is optional but strongly encouraged, see the
     * the detailed comments above Metric.
     *
     * Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January
     * 1970.
     *
     * @generated from protobuf field: fixed64 start_time_unix_nano = 2;
     */
    startTimeUnixNano: number;
    /**
     * TimeUnixNano is required, see the detailed comments above Metric.
     *
     * Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January
     * 1970.
     *
     * @generated from protobuf field: fixed64 time_unix_nano = 3;
     */
    timeUnixNano: number;
    /**
     * count is the number of values in the population. Must be non-negative.
     *
     * @generated from protobuf field: fixed64 count = 4;
     */
    count: number;
    /**
     * sum of the values in the population. If count is zero then this field
     * must be zero.
     *
     * Note: Sum should only be filled out when measuring non-negative discrete
     * events, and is assumed to be monotonic over the values of these events.
     * Negative events *can* be recorded, but sum should not be filled out when
     * doing so.  This is specifically to enforce compatibility w/ OpenMetrics,
     * see: https://github.com/OpenObservability/OpenMetrics/blob/main/specification/OpenMetrics.md#summary
     *
     * @generated from protobuf field: double sum = 5;
     */
    sum: number;
    /**
     * (Optional) list of values at different quantiles of the distribution calculated
     * from the current snapshot. The quantiles must be strictly increasing.
     *
     * @generated from protobuf field: repeated opentelemetry.proto.metrics.v1.SummaryDataPoint.ValueAtQuantile quantile_values = 6;
     */
    quantileValues: SummaryDataPoint_ValueAtQuantile[];
    /**
     * Flags that apply to this specific data point.  See DataPointFlags
     * for the available flags and their meaning.
     *
     * @generated from protobuf field: uint32 flags = 8;
     */
    flags: number;
}

/**
 * @generated MessageType for protobuf message opentelemetry.proto.metrics.v1.SummaryDataPoint
 */
export  const SummaryDataPoint: SummaryDataPoint$Type;

export class SummaryDataPoint_ValueAtQuantile$Type extends MessageType<SummaryDataPoint_ValueAtQuantile> {
    constructor();
    create(value?: PartialMessage<SummaryDataPoint_ValueAtQuantile>): SummaryDataPoint_ValueAtQuantile;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SummaryDataPoint_ValueAtQuantile): SummaryDataPoint_ValueAtQuantile;
    internalBinaryWrite(message: SummaryDataPoint_ValueAtQuantile, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * Represents the value at a given quantile of a distribution.
 *
 * To record Min and Max values following conventions are used:
 * - The 1.0 quantile is equivalent to the maximum value observed.
 * - The 0.0 quantile is equivalent to the minimum value observed.
 *
 * See the following issue for more context:
 * https://github.com/open-telemetry/opentelemetry-proto/issues/125
 *
 * @generated from protobuf message opentelemetry.proto.metrics.v1.SummaryDataPoint.ValueAtQuantile
 */
export  interface SummaryDataPoint_ValueAtQuantile {
    /**
     * The quantile of a distribution. Must be in the interval
     * [0.0, 1.0].
     *
     * @generated from protobuf field: double quantile = 1;
     */
    quantile: number;
    /**
     * The value at the given quantile of a distribution.
     *
     * Quantile values must NOT be negative.
     *
     * @generated from protobuf field: double value = 2;
     */
    value: number;
}

/**
 * @generated MessageType for protobuf message opentelemetry.proto.metrics.v1.SummaryDataPoint.ValueAtQuantile
 */
export  const SummaryDataPoint_ValueAtQuantile: SummaryDataPoint_ValueAtQuantile$Type;

export class Timestamp$Type extends MessageType<Timestamp> {
    constructor();
    /**
     * Creates a new `Timestamp` for the current time.
     */
    now(): Timestamp;
    /**
     * Converts a `Timestamp` to a JavaScript Date.
     */
    toDate(message: Timestamp): Date;
    /**
     * Converts a JavaScript Date to a `Timestamp`.
     */
    fromDate(date: Date): Timestamp;
    /**
     * In JSON format, the `Timestamp` type is encoded as a string
     * in the RFC 3339 format.
     */
    internalJsonWrite(message: Timestamp, options: JsonWriteOptions): JsonValue;
    /**
     * In JSON format, the `Timestamp` type is encoded as a string
     * in the RFC 3339 format.
     */
    internalJsonRead(json: JsonValue, options: JsonReadOptions, target?: Timestamp): Timestamp;
}

/**
 * A Timestamp represents a point in time independent of any time zone or local
 * calendar, encoded as a count of seconds and fractions of seconds at
 * nanosecond resolution. The count is relative to an epoch at UTC midnight on
 * January 1, 1970, in the proleptic Gregorian calendar which extends the
 * Gregorian calendar backwards to year one.
 *
 * All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
 * second table is needed for interpretation, using a [24-hour linear
 * smear](https://developers.google.com/time/smear).
 *
 * The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
 * restricting to that range, we ensure that we can convert to and from [RFC
 * 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
 *
 * # Examples
 *
 * Example 1: Compute Timestamp from POSIX `time()`.
 *
 *     Timestamp timestamp;
 *     timestamp.set_seconds(time(NULL));
 *     timestamp.set_nanos(0);
 *
 * Example 2: Compute Timestamp from POSIX `gettimeofday()`.
 *
 *     struct timeval tv;
 *     gettimeofday(&tv, NULL);
 *
 *     Timestamp timestamp;
 *     timestamp.set_seconds(tv.tv_sec);
 *     timestamp.set_nanos(tv.tv_usec * 1000);
 *
 * Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
 *
 *     FILETIME ft;
 *     GetSystemTimeAsFileTime(&ft);
 *     UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
 *
 *     // A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
 *     // is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
 *     Timestamp timestamp;
 *     timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
 *     timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
 *
 * Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
 *
 *     long millis = System.currentTimeMillis();
 *
 *     Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
 *         .setNanos((int) ((millis % 1000) * 1000000)).build();
 *
 * Example 5: Compute Timestamp from Java `Instant.now()`.
 *
 *     Instant now = Instant.now();
 *
 *     Timestamp timestamp =
 *         Timestamp.newBuilder().setSeconds(now.getEpochSecond())
 *             .setNanos(now.getNano()).build();
 *
 * Example 6: Compute Timestamp from current time in Python.
 *
 *     timestamp = Timestamp()
 *     timestamp.GetCurrentTime()
 *
 * # JSON Mapping
 *
 * In JSON format, the Timestamp type is encoded as a string in the
 * [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the
 * format is "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z"
 * where {year} is always expressed using four digits while {month}, {day},
 * {hour}, {min}, and {sec} are zero-padded to two digits each. The fractional
 * seconds, which can go up to 9 digits (i.e. up to 1 nanosecond resolution),
 * are optional. The "Z" suffix indicates the timezone ("UTC"); the timezone
 * is required. A proto3 JSON serializer should always use UTC (as indicated by
 * "Z") when printing the Timestamp type and a proto3 JSON parser should be
 * able to accept both UTC and other timezones (as indicated by an offset).
 *
 * For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past
 * 01:30 UTC on January 15, 2017.
 *
 * In JavaScript, one can convert a Date object to this format using the
 * standard
 * [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
 * method. In Python, a standard `datetime.datetime` object can be converted
 * to this format using
 * [`strftime`](https://docs.python.org/2/library/time.html#time.strftime) with
 * the time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one can use
 * the Joda Time's [`ISODateTimeFormat.dateTime()`](
 * http://www.joda.org/joda-time/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime%2D%2D
 * ) to obtain a formatter capable of generating timestamps in this format.
 *
 *
 * @generated from protobuf message google.protobuf.Timestamp
 */
export  interface Timestamp {
    /**
     * Represents seconds of UTC time since Unix epoch
     * 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
     * 9999-12-31T23:59:59Z inclusive.
     *
     * @generated from protobuf field: int64 seconds = 1;
     */
    seconds: number;
    /**
     * Non-negative fractions of a second at nanosecond resolution. Negative
     * second values with fractions must still have non-negative nanos values
     * that count forward in time. Must be from 0 to 999,999,999
     * inclusive.
     *
     * @generated from protobuf field: int32 nanos = 2;
     */
    nanos: number;
}

/**
 * @generated MessageType for protobuf message google.protobuf.Timestamp
 */
export  const Timestamp: Timestamp$Type;

export  class Trace extends Config {
    constructor();
}

export class TracesData$Type extends MessageType<TracesData> {
    constructor();
    create(value?: PartialMessage<TracesData>): TracesData;
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: TracesData): TracesData;
    internalBinaryWrite(message: TracesData, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter;
}

/**
 * TracesData represents the traces data that can be stored in a persistent storage,
 * OR can be embedded by other protocols that transfer OTLP traces data but do
 * not implement the OTLP protocol.
 *
 * The main difference between this message and collector protocol is that
 * in this message there will not be any "control" or "metadata" specific to
 * OTLP protocol.
 *
 * When new fields are added into this message, the OTLP request MUST be updated
 * as well.
 *
 * @generated from protobuf message opentelemetry.proto.trace.v1.TracesData
 */
export  interface TracesData {
    /**
     * An array of ResourceSpans.
     * For data coming from a single resource this array will typically contain
     * one element. Intermediary nodes that receive data from multiple origins
     * typically batch the data before forwarding further and in that case this
     * array will contain multiple elements.
     *
     * @generated from protobuf field: repeated opentelemetry.proto.trace.v1.ResourceSpans resource_spans = 1;
     */
    resourceSpans: ResourceSpans[];
}

/**
 * @generated MessageType for protobuf message opentelemetry.proto.trace.v1.TracesData
 */
export  const TracesData: TracesData$Type;

/**
 * A enum field of unknown type.
 */
export type UnknownEnum = number;

/**
 * Store an unknown field for a message somewhere.
 */
export type UnknownFieldReader = (typeName: string, message: any, fieldNo: number, wireType: WireType, data: Uint8Array) => void;

/**
 * Write unknown fields stored for the message to the writer.
 */
export type UnknownFieldWriter = (typeName: string, message: any, writer: IBinaryWriter) => void;

/**
 * A map field of unknown type.
 */
export type UnknownMap<T = UnknownMessage | UnknownScalar | UnknownEnum> = {
    [key: string]: T;
};

/**
 * A message of unknown type.
 */
export interface UnknownMessage {
    [k: string]: UnknownScalar | UnknownEnum | UnknownMessage | UnknownOneofGroup | UnknownMap | UnknownScalar[] | UnknownMessage[] | UnknownEnum[] | undefined;
}

/**
 * A unknown oneof group. See `isOneofGroup()` for details.
 */
export type UnknownOneofGroup = {
    oneofKind: undefined | string;
    [k: string]: UnknownScalar | UnknownEnum | UnknownMessage | undefined;
};

/**
 * A scalar field of unknown type.
 */
export type UnknownScalar = string | number | bigint | boolean | Uint8Array;

export class Value$Type extends MessageType<Value> {
    constructor();
    /**
     * Encode `Value` to JSON value.
     */
    internalJsonWrite(message: Value, options: JsonWriteOptions): JsonValue;
    /**
     * Decode `Value` from JSON value.
     */
    internalJsonRead(json: JsonValue, options: JsonReadOptions, target?: Value): Value;
}

/**
 * `Value` represents a dynamically typed value which can be either
 * null, a number, a string, a boolean, a recursive struct value, or a
 * list of values. A producer of value is expected to set one of these
 * variants. Absence of any variant indicates an error.
 *
 * The JSON representation for `Value` is JSON value.
 *
 * @generated from protobuf message google.protobuf.Value
 */
export interface Value {
    /**
     * @generated from protobuf oneof: kind
     */
    kind: {
        oneofKind: "nullValue";
        /**
         * Represents a null value.
         *
         * @generated from protobuf field: google.protobuf.NullValue null_value = 1;
         */
        nullValue: NullValue;
    } | {
        oneofKind: "numberValue";
        /**
         * Represents a double value.
         *
         * @generated from protobuf field: double number_value = 2;
         */
        numberValue: number;
    } | {
        oneofKind: "stringValue";
        /**
         * Represents a string value.
         *
         * @generated from protobuf field: string string_value = 3;
         */
        stringValue: string;
    } | {
        oneofKind: "boolValue";
        /**
         * Represents a boolean value.
         *
         * @generated from protobuf field: bool bool_value = 4;
         */
        boolValue: boolean;
    } | {
        oneofKind: "structValue";
        /**
         * Represents a structured value.
         *
         * @generated from protobuf field: google.protobuf.Struct struct_value = 5;
         */
        structValue: Struct;
    } | {
        oneofKind: "listValue";
        /**
         * Represents a repeated `Value`.
         *
         * @generated from protobuf field: google.protobuf.ListValue list_value = 6;
         */
        listValue: ListValue;
    } | {
        oneofKind: undefined;
    };
}

/**
 * @generated MessageType for protobuf message google.protobuf.Value
 */
export const Value: Value$Type;

/**
 * Protobuf binary format wire types.
 *
 * A wire type provides just enough information to find the length of the
 * following value.
 *
 * See https://developers.google.com/protocol-buffers/docs/encoding#structure
 */
export enum WireType {
    /**
     * Used for int32, int64, uint32, uint64, sint32, sint64, bool, enum
     */
    Varint = 0,
    /**
     * Used for fixed64, sfixed64, double.
     * Always 8 bytes with little-endian byte order.
     */
    Bit64 = 1,
    /**
     * Used for string, bytes, embedded messages, packed repeated fields
     *
     * Only repeated numeric types (types which use the varint, 32-bit,
     * or 64-bit wire types) can be packed. In proto3, such fields are
     * packed by default.
     */
    LengthDelimited = 2,
    /**
     * Used for groups
     * @deprecated
     */
    StartGroup = 3,
    /**
     * Used for groups
     * @deprecated
     */
    EndGroup = 4,
    /**
     * Used for fixed32, sfixed32, float.
     * Always 4 bytes with little-endian byte order.
     */
    Bit32 = 5
}

export { }

}