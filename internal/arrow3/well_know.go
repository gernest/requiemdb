package arrow3

import (
	commonv1 "go.opentelemetry.io/proto/otlp/common/v1"
)

var (
	otelAnyDescriptor = (&commonv1.AnyValue{}).ProtoReflect().Descriptor()
)
