// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/f1ap_4g/v1/f1ap_commondatatypes.proto

package f1apcommondatatypesv1

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

// Validate checks the field values on PrivateIeID with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PrivateIeID) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PrivateIeID with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PrivateIeIDMultiError, or
// nil if none found.
func (m *PrivateIeID) ValidateAll() error {
	return m.validate(true)
}

func (m *PrivateIeID) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.PrivateIeId.(type) {

	case *PrivateIeID_Local:
		// no validation rules for Local

	case *PrivateIeID_Global:
		// no validation rules for Global

	}

	if len(errors) > 0 {
		return PrivateIeIDMultiError(errors)
	}

	return nil
}

// PrivateIeIDMultiError is an error wrapping multiple validation errors
// returned by PrivateIeID.ValidateAll() if the designated constraints aren't met.
type PrivateIeIDMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PrivateIeIDMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PrivateIeIDMultiError) AllErrors() []error { return m }

// PrivateIeIDValidationError is the validation error returned by
// PrivateIeID.Validate if the designated constraints aren't met.
type PrivateIeIDValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PrivateIeIDValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PrivateIeIDValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PrivateIeIDValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PrivateIeIDValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PrivateIeIDValidationError) ErrorName() string { return "PrivateIeIDValidationError" }

// Error satisfies the builtin error interface
func (e PrivateIeIDValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPrivateIeID.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PrivateIeIDValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PrivateIeIDValidationError{}

// Validate checks the field values on ProcedureCode with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ProcedureCode) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProcedureCode with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ProcedureCodeMultiError, or
// nil if none found.
func (m *ProcedureCode) ValidateAll() error {
	return m.validate(true)
}

func (m *ProcedureCode) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if val := m.GetValue(); val < 0 || val > 255 {
		err := ProcedureCodeValidationError{
			field:  "Value",
			reason: "value must be inside range [0, 255]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ProcedureCodeMultiError(errors)
	}

	return nil
}

// ProcedureCodeMultiError is an error wrapping multiple validation errors
// returned by ProcedureCode.ValidateAll() if the designated constraints
// aren't met.
type ProcedureCodeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProcedureCodeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProcedureCodeMultiError) AllErrors() []error { return m }

// ProcedureCodeValidationError is the validation error returned by
// ProcedureCode.Validate if the designated constraints aren't met.
type ProcedureCodeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProcedureCodeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProcedureCodeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProcedureCodeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProcedureCodeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProcedureCodeValidationError) ErrorName() string { return "ProcedureCodeValidationError" }

// Error satisfies the builtin error interface
func (e ProcedureCodeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProcedureCode.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProcedureCodeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProcedureCodeValidationError{}

// Validate checks the field values on ProtocolExtensionId with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ProtocolExtensionId) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProtocolExtensionId with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProtocolExtensionIdMultiError, or nil if none found.
func (m *ProtocolExtensionId) ValidateAll() error {
	return m.validate(true)
}

func (m *ProtocolExtensionId) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if val := m.GetValue(); val < 0 || val > 65535 {
		err := ProtocolExtensionIdValidationError{
			field:  "Value",
			reason: "value must be inside range [0, 65535]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ProtocolExtensionIdMultiError(errors)
	}

	return nil
}

// ProtocolExtensionIdMultiError is an error wrapping multiple validation
// errors returned by ProtocolExtensionId.ValidateAll() if the designated
// constraints aren't met.
type ProtocolExtensionIdMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProtocolExtensionIdMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProtocolExtensionIdMultiError) AllErrors() []error { return m }

// ProtocolExtensionIdValidationError is the validation error returned by
// ProtocolExtensionId.Validate if the designated constraints aren't met.
type ProtocolExtensionIdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProtocolExtensionIdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProtocolExtensionIdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProtocolExtensionIdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProtocolExtensionIdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProtocolExtensionIdValidationError) ErrorName() string {
	return "ProtocolExtensionIdValidationError"
}

// Error satisfies the builtin error interface
func (e ProtocolExtensionIdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProtocolExtensionId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProtocolExtensionIdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProtocolExtensionIdValidationError{}

// Validate checks the field values on ProtocolIeID with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ProtocolIeID) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProtocolIeID with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ProtocolIeIDMultiError, or
// nil if none found.
func (m *ProtocolIeID) ValidateAll() error {
	return m.validate(true)
}

func (m *ProtocolIeID) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if val := m.GetValue(); val < 0 || val > 65535 {
		err := ProtocolIeIDValidationError{
			field:  "Value",
			reason: "value must be inside range [0, 65535]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ProtocolIeIDMultiError(errors)
	}

	return nil
}

// ProtocolIeIDMultiError is an error wrapping multiple validation errors
// returned by ProtocolIeID.ValidateAll() if the designated constraints aren't met.
type ProtocolIeIDMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProtocolIeIDMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProtocolIeIDMultiError) AllErrors() []error { return m }

// ProtocolIeIDValidationError is the validation error returned by
// ProtocolIeID.Validate if the designated constraints aren't met.
type ProtocolIeIDValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProtocolIeIDValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProtocolIeIDValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProtocolIeIDValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProtocolIeIDValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProtocolIeIDValidationError) ErrorName() string { return "ProtocolIeIDValidationError" }

// Error satisfies the builtin error interface
func (e ProtocolIeIDValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProtocolIeID.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProtocolIeIDValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProtocolIeIDValidationError{}
