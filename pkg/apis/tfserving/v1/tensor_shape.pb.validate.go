// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/apis/tfserving/v1/tensor_shape.proto

package tfserving

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on TensorShapeProto with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *TensorShapeProto) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TensorShapeProto with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TensorShapeProtoMultiError, or nil if none found.
func (m *TensorShapeProto) ValidateAll() error {
	return m.validate(true)
}

func (m *TensorShapeProto) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetDim() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TensorShapeProtoValidationError{
						field:  fmt.Sprintf("Dim[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TensorShapeProtoValidationError{
						field:  fmt.Sprintf("Dim[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TensorShapeProtoValidationError{
					field:  fmt.Sprintf("Dim[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for UnknownRank

	if len(errors) > 0 {
		return TensorShapeProtoMultiError(errors)
	}

	return nil
}

// TensorShapeProtoMultiError is an error wrapping multiple validation errors
// returned by TensorShapeProto.ValidateAll() if the designated constraints
// aren't met.
type TensorShapeProtoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TensorShapeProtoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TensorShapeProtoMultiError) AllErrors() []error { return m }

// TensorShapeProtoValidationError is the validation error returned by
// TensorShapeProto.Validate if the designated constraints aren't met.
type TensorShapeProtoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TensorShapeProtoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TensorShapeProtoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TensorShapeProtoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TensorShapeProtoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TensorShapeProtoValidationError) ErrorName() string { return "TensorShapeProtoValidationError" }

// Error satisfies the builtin error interface
func (e TensorShapeProtoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTensorShapeProto.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TensorShapeProtoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TensorShapeProtoValidationError{}

// Validate checks the field values on TensorShapeProto_Dim with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *TensorShapeProto_Dim) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TensorShapeProto_Dim with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// TensorShapeProto_DimMultiError, or nil if none found.
func (m *TensorShapeProto_Dim) ValidateAll() error {
	return m.validate(true)
}

func (m *TensorShapeProto_Dim) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Size

	// no validation rules for Name

	if len(errors) > 0 {
		return TensorShapeProto_DimMultiError(errors)
	}

	return nil
}

// TensorShapeProto_DimMultiError is an error wrapping multiple validation
// errors returned by TensorShapeProto_Dim.ValidateAll() if the designated
// constraints aren't met.
type TensorShapeProto_DimMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TensorShapeProto_DimMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TensorShapeProto_DimMultiError) AllErrors() []error { return m }

// TensorShapeProto_DimValidationError is the validation error returned by
// TensorShapeProto_Dim.Validate if the designated constraints aren't met.
type TensorShapeProto_DimValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TensorShapeProto_DimValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TensorShapeProto_DimValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TensorShapeProto_DimValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TensorShapeProto_DimValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TensorShapeProto_DimValidationError) ErrorName() string {
	return "TensorShapeProto_DimValidationError"
}

// Error satisfies the builtin error interface
func (e TensorShapeProto_DimValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTensorShapeProto_Dim.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TensorShapeProto_DimValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TensorShapeProto_DimValidationError{}
