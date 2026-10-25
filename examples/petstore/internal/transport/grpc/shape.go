package grpc

import (
	"reflect"

	"buf.build/go/protovalidate"
	"google.golang.org/protobuf/proto"

	"github.com/stephen-mcelhose/testable-spec-pattern/examples/petstore/internal/pet"
)

// validateShape is the stereotypical proto shape check at the transport boundary:
// nil guard, then protovalidate.Validate against buf.validate on the request.
// Shape violations MUST map to InvalidArgument; the application service MUST NOT
// be called on shape failure.
func validateShape[T proto.Message](msg T) error {
	var zero T
	if any(msg) == any(zero) || isNilProto(msg) {
		return &pet.InvalidArgumentError{Message: "request is required"}
	}
	if err := protovalidate.Validate(msg); err != nil {
		return &pet.InvalidArgumentError{Message: err.Error()}
	}
	return nil
}

func isNilProto(msg proto.Message) bool {
	v := reflect.ValueOf(msg)
	switch v.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Map, reflect.Slice, reflect.Chan, reflect.Func:
		return v.IsNil()
	default:
		return false
	}
}
