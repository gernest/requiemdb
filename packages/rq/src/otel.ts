
export interface LogsData {
    resource_logs: ResourceLogs[]
}

export interface ResourceLogs {
    resource: Resource
    scope_logs: ScopeLogs[]
    schema_url: string
}

export interface ScopeLogs {
    scope: InstrumentationScope
    log_records: LogRecord
    schema_url: string
}

export interface LogRecord {
    time_unix_nano: number
    observed_time_unix_nano: number
    severity_number: SeverityNumber
    severity_text: string
    body: AnyValue
    attributes: KeyValue[]
    trace_id: Uint8Array
    span_id: Uint8Array
}

enum SeverityNumber {
    SEVERITY_NUMBER_UNSPECIFIED = 0,
    SEVERITY_NUMBER_TRACE = 1,
    SEVERITY_NUMBER_TRACE2 = 2,
    SEVERITY_NUMBER_TRACE3 = 3,
    SEVERITY_NUMBER_TRACE4 = 4,
    SEVERITY_NUMBER_DEBUG = 5,
    SEVERITY_NUMBER_DEBUG2 = 6,
    SEVERITY_NUMBER_DEBUG3 = 7,
    SEVERITY_NUMBER_DEBUG4 = 8,
    SEVERITY_NUMBER_INFO = 9,
    SEVERITY_NUMBER_INFO2 = 10,
    SEVERITY_NUMBER_INFO3 = 11,
    SEVERITY_NUMBER_INFO4 = 12,
    SEVERITY_NUMBER_WARN = 13,
    SEVERITY_NUMBER_WARN2 = 14,
    SEVERITY_NUMBER_WARN3 = 15,
    SEVERITY_NUMBER_WARN4 = 16,
    SEVERITY_NUMBER_ERROR = 17,
    SEVERITY_NUMBER_ERROR2 = 18,
    SEVERITY_NUMBER_ERROR3 = 19,
    SEVERITY_NUMBER_ERROR4 = 20,
    SEVERITY_NUMBER_FATAL = 21,
    SEVERITY_NUMBER_FATAL2 = 22,
    SEVERITY_NUMBER_FATAL3 = 23,
    SEVERITY_NUMBER_FATAL4 = 24,
}

export interface MetricsData {
    resource_metrics: ResourceMetrics[]
}

export interface ResourceMetrics {
    resource?: Resource
    scope_metrics: ScopeMetrics[]
    schema_url: string
}

export interface ScopeMetrics {
    scope?: InstrumentationScope
    metrics: Metric[]
    schema_url: string
}

export interface Metric {
    name: string
    description: string
    unit: string
    GetGauge(): Gauge
    GetSum(): Sum
    GetHistogram(): Histogram
    GetExponentialHistogram(): ExponentialHistogram
    GetSummary(): Summary
}

export interface Gauge {
    data_points: NumberDataPoint[]
}

export interface Sum {
    data_points: NumberDataPoint[]
    aggregation_temporality: AggregationTemporality
    is_monotonic: boolean
}

export interface Histogram {
    data_points: HistogramDataPoint[]
    aggregation_temporality: AggregationTemporality
}

export interface ExponentialHistogram {
    data_points: ExponentialHistogramDataPoint[]
    aggregation_temporality: AggregationTemporality
}

export interface Summary {
    data_points: SummaryDataPoint[]
}

export enum AggregationTemporality {
    AGGREGATION_TEMPORALITY_UNSPECIFIED = 0,
    AGGREGATION_TEMPORALITY_DELTA = 1,
    AGGREGATION_TEMPORALITY_CUMULATIVE = 2,
}

export interface NumberDataPoint {
    attributes: KeyValue[]
    start_time_unix_nano: number
    time_unix_nano: number
    GetAsDouble(): number
    GetAsInt(): number
    exemplars: Exemplar[]
}

export interface HistogramDataPoint {
    attributes: KeyValue[]
    start_time_unix_nano: number
    time_unix_nano: number
    count: number
    sum?: number
    bucket_counts: number[]
    explicit_bounds: number[]
    exemplars: Exemplar[]
    min: number[]
    max: number[]
}

export interface ExponentialHistogramDataPoint {
    attributes: KeyValue[]
    start_time_unix_nano: number
    time_unix_nano: number
    count: number
    sum: number[]
    scale: number
    zero_count: number
    positive: Buckets
    negative: Buckets
    exemplars: Exemplar[]
    min: number[]
    max: number[]
    zero_threshold: number[]
}

export interface Buckets {
    offset: number
    bucket_counts: number[]
}

export interface SummaryDataPoint {
    attributes: KeyValue[]
    start_time_unix_nano: number
    time_unix_nano: number
    count: number
    sum: number
    quantile_values: ValueAtQuantile[]
}

export interface ValueAtQuantile {
    quantile: number
    value: number
}

export interface Exemplar {
    filtered_attributes: KeyValue[]
    time_unix_nano: number
    span_id?: Uint8Array
    trace_id?: Uint8Array
    GetAsDouble(): number
    GetAsInt(): number
}

export interface Resource {
    attributes?: KeyValue[]
}

export interface KeyValue {
    key: string
    value: AnyValue
}

export interface AnyValue {
    GetStringValue(): string
    GetBoolValue(): boolean
    GetIntValue(): boolean
}

export interface InstrumentationScope {
    name: string
    version: string
    attributes?: KeyValue[]
}
