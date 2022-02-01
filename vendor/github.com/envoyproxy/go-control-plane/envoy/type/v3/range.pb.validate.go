// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/type/v3/range.proto

package envoy_type_v3

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

// Validate checks the field values on Int64Range with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Int64Range) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Int64Range with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Int64RangeMultiError, or
// nil if none found.
func (m *Int64Range) ValidateAll() error {
	return m.validate(true)
}

func (m *Int64Range) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Start

	// no validation rules for End

	if len(errors) > 0 {
		return Int64RangeMultiError(errors)
	}
	return nil
}

// Int64RangeMultiError is an error wrapping multiple validation errors
// returned by Int64Range.ValidateAll() if the designated constraints aren't met.
type Int64RangeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Int64RangeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Int64RangeMultiError) AllErrors() []error { return m }

// Int64RangeValidationError is the validation error returned by
// Int64Range.Validate if the designated constraints aren't met.
type Int64RangeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Int64RangeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Int64RangeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Int64RangeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Int64RangeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Int64RangeValidationError) ErrorName() string { return "Int64RangeValidationError" }

// Error satisfies the builtin error interface
func (e Int64RangeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInt64Range.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Int64RangeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Int64RangeValidationError{}

// Validate checks the field values on Int32Range with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Int32Range) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Int32Range with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Int32RangeMultiError, or
// nil if none found.
func (m *Int32Range) ValidateAll() error {
	return m.validate(true)
}

func (m *Int32Range) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Start

	// no validation rules for End

	if len(errors) > 0 {
		return Int32RangeMultiError(errors)
	}
	return nil
}

// Int32RangeMultiError is an error wrapping multiple validation errors
// returned by Int32Range.ValidateAll() if the designated constraints aren't met.
type Int32RangeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Int32RangeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Int32RangeMultiError) AllErrors() []error { return m }

// Int32RangeValidationError is the validation error returned by
// Int32Range.Validate if the designated constraints aren't met.
type Int32RangeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Int32RangeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Int32RangeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Int32RangeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Int32RangeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Int32RangeValidationError) ErrorName() string { return "Int32RangeValidationError" }

// Error satisfies the builtin error interface
func (e Int32RangeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInt32Range.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Int32RangeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Int32RangeValidationError{}

// Validate checks the field values on DoubleRange with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DoubleRange) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DoubleRange with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DoubleRangeMultiError, or
// nil if none found.
func (m *DoubleRange) ValidateAll() error {
	return m.validate(true)
}

func (m *DoubleRange) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Start

	// no validation rules for End

	if len(errors) > 0 {
		return DoubleRangeMultiError(errors)
	}
	return nil
}

// DoubleRangeMultiError is an error wrapping multiple validation errors
// returned by DoubleRange.ValidateAll() if the designated constraints aren't met.
type DoubleRangeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DoubleRangeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DoubleRangeMultiError) AllErrors() []error { return m }

// DoubleRangeValidationError is the validation error returned by
// DoubleRange.Validate if the designated constraints aren't met.
type DoubleRangeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DoubleRangeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DoubleRangeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DoubleRangeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DoubleRangeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DoubleRangeValidationError) ErrorName() string { return "DoubleRangeValidationError" }

// Error satisfies the builtin error interface
func (e DoubleRangeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDoubleRange.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DoubleRangeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DoubleRangeValidationError{}
