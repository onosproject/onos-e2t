// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/xnap/v1/xnap_pdu_descriptions.proto

package xnappdudescriptionsv1

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

// Validate checks the field values on XnApPDu with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *XnApPDu) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on XnApPDu with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in XnApPDuMultiError, or nil if none found.
func (m *XnApPDu) ValidateAll() error {
	return m.validate(true)
}

func (m *XnApPDu) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.XnApPdu.(type) {

	case *XnApPDu_InitiatingMessage:

		if all {
			switch v := interface{}(m.GetInitiatingMessage()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, XnApPDuValidationError{
						field:  "InitiatingMessage",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, XnApPDuValidationError{
						field:  "InitiatingMessage",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetInitiatingMessage()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return XnApPDuValidationError{
					field:  "InitiatingMessage",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *XnApPDu_SuccessfulOutcome:

		if all {
			switch v := interface{}(m.GetSuccessfulOutcome()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, XnApPDuValidationError{
						field:  "SuccessfulOutcome",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, XnApPDuValidationError{
						field:  "SuccessfulOutcome",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetSuccessfulOutcome()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return XnApPDuValidationError{
					field:  "SuccessfulOutcome",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *XnApPDu_UnsuccessfulOutcome:

		if all {
			switch v := interface{}(m.GetUnsuccessfulOutcome()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, XnApPDuValidationError{
						field:  "UnsuccessfulOutcome",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, XnApPDuValidationError{
						field:  "UnsuccessfulOutcome",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUnsuccessfulOutcome()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return XnApPDuValidationError{
					field:  "UnsuccessfulOutcome",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return XnApPDuMultiError(errors)
	}

	return nil
}

// XnApPDuMultiError is an error wrapping multiple validation errors returned
// by XnApPDu.ValidateAll() if the designated constraints aren't met.
type XnApPDuMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m XnApPDuMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m XnApPDuMultiError) AllErrors() []error { return m }

// XnApPDuValidationError is the validation error returned by XnApPDu.Validate
// if the designated constraints aren't met.
type XnApPDuValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e XnApPDuValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e XnApPDuValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e XnApPDuValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e XnApPDuValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e XnApPDuValidationError) ErrorName() string { return "XnApPDuValidationError" }

// Error satisfies the builtin error interface
func (e XnApPDuValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sXnApPDu.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = XnApPDuValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = XnApPDuValidationError{}

// Validate checks the field values on InitiatingMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *InitiatingMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InitiatingMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// InitiatingMessageMultiError, or nil if none found.
func (m *InitiatingMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *InitiatingMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ProcedureCode

	// no validation rules for Criticality

	if all {
		switch v := interface{}(m.GetValue()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, InitiatingMessageValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, InitiatingMessageValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InitiatingMessageValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return InitiatingMessageMultiError(errors)
	}

	return nil
}

// InitiatingMessageMultiError is an error wrapping multiple validation errors
// returned by InitiatingMessage.ValidateAll() if the designated constraints
// aren't met.
type InitiatingMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InitiatingMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InitiatingMessageMultiError) AllErrors() []error { return m }

// InitiatingMessageValidationError is the validation error returned by
// InitiatingMessage.Validate if the designated constraints aren't met.
type InitiatingMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InitiatingMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InitiatingMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InitiatingMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InitiatingMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InitiatingMessageValidationError) ErrorName() string {
	return "InitiatingMessageValidationError"
}

// Error satisfies the builtin error interface
func (e InitiatingMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInitiatingMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InitiatingMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InitiatingMessageValidationError{}

// Validate checks the field values on
// InitiatingMessageXnApElementaryProcedures with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *InitiatingMessageXnApElementaryProcedures) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// InitiatingMessageXnApElementaryProcedures with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// InitiatingMessageXnApElementaryProceduresMultiError, or nil if none found.
func (m *InitiatingMessageXnApElementaryProcedures) ValidateAll() error {
	return m.validate(true)
}

func (m *InitiatingMessageXnApElementaryProcedures) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.ImValues.(type) {

	case *InitiatingMessageXnApElementaryProcedures_XnSetupRequest:

		if all {
			switch v := interface{}(m.GetXnSetupRequest()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, InitiatingMessageXnApElementaryProceduresValidationError{
						field:  "XnSetupRequest",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, InitiatingMessageXnApElementaryProceduresValidationError{
						field:  "XnSetupRequest",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetXnSetupRequest()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return InitiatingMessageXnApElementaryProceduresValidationError{
					field:  "XnSetupRequest",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return InitiatingMessageXnApElementaryProceduresMultiError(errors)
	}

	return nil
}

// InitiatingMessageXnApElementaryProceduresMultiError is an error wrapping
// multiple validation errors returned by
// InitiatingMessageXnApElementaryProcedures.ValidateAll() if the designated
// constraints aren't met.
type InitiatingMessageXnApElementaryProceduresMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InitiatingMessageXnApElementaryProceduresMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InitiatingMessageXnApElementaryProceduresMultiError) AllErrors() []error { return m }

// InitiatingMessageXnApElementaryProceduresValidationError is the validation
// error returned by InitiatingMessageXnApElementaryProcedures.Validate if the
// designated constraints aren't met.
type InitiatingMessageXnApElementaryProceduresValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InitiatingMessageXnApElementaryProceduresValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InitiatingMessageXnApElementaryProceduresValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InitiatingMessageXnApElementaryProceduresValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InitiatingMessageXnApElementaryProceduresValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InitiatingMessageXnApElementaryProceduresValidationError) ErrorName() string {
	return "InitiatingMessageXnApElementaryProceduresValidationError"
}

// Error satisfies the builtin error interface
func (e InitiatingMessageXnApElementaryProceduresValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInitiatingMessageXnApElementaryProcedures.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InitiatingMessageXnApElementaryProceduresValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InitiatingMessageXnApElementaryProceduresValidationError{}

// Validate checks the field values on SuccessfulOutcome with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SuccessfulOutcome) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SuccessfulOutcome with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SuccessfulOutcomeMultiError, or nil if none found.
func (m *SuccessfulOutcome) ValidateAll() error {
	return m.validate(true)
}

func (m *SuccessfulOutcome) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ProcedureCode

	// no validation rules for Criticality

	if all {
		switch v := interface{}(m.GetValue()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SuccessfulOutcomeValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SuccessfulOutcomeValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SuccessfulOutcomeValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SuccessfulOutcomeMultiError(errors)
	}

	return nil
}

// SuccessfulOutcomeMultiError is an error wrapping multiple validation errors
// returned by SuccessfulOutcome.ValidateAll() if the designated constraints
// aren't met.
type SuccessfulOutcomeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SuccessfulOutcomeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SuccessfulOutcomeMultiError) AllErrors() []error { return m }

// SuccessfulOutcomeValidationError is the validation error returned by
// SuccessfulOutcome.Validate if the designated constraints aren't met.
type SuccessfulOutcomeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SuccessfulOutcomeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SuccessfulOutcomeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SuccessfulOutcomeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SuccessfulOutcomeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SuccessfulOutcomeValidationError) ErrorName() string {
	return "SuccessfulOutcomeValidationError"
}

// Error satisfies the builtin error interface
func (e SuccessfulOutcomeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSuccessfulOutcome.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SuccessfulOutcomeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SuccessfulOutcomeValidationError{}

// Validate checks the field values on
// SuccessfulOutcomeXnApElementaryProcedures with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SuccessfulOutcomeXnApElementaryProcedures) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// SuccessfulOutcomeXnApElementaryProcedures with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// SuccessfulOutcomeXnApElementaryProceduresMultiError, or nil if none found.
func (m *SuccessfulOutcomeXnApElementaryProcedures) ValidateAll() error {
	return m.validate(true)
}

func (m *SuccessfulOutcomeXnApElementaryProcedures) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.SoValues.(type) {

	case *SuccessfulOutcomeXnApElementaryProcedures_XnSetupResponse:

		if all {
			switch v := interface{}(m.GetXnSetupResponse()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SuccessfulOutcomeXnApElementaryProceduresValidationError{
						field:  "XnSetupResponse",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SuccessfulOutcomeXnApElementaryProceduresValidationError{
						field:  "XnSetupResponse",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetXnSetupResponse()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SuccessfulOutcomeXnApElementaryProceduresValidationError{
					field:  "XnSetupResponse",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SuccessfulOutcomeXnApElementaryProceduresMultiError(errors)
	}

	return nil
}

// SuccessfulOutcomeXnApElementaryProceduresMultiError is an error wrapping
// multiple validation errors returned by
// SuccessfulOutcomeXnApElementaryProcedures.ValidateAll() if the designated
// constraints aren't met.
type SuccessfulOutcomeXnApElementaryProceduresMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SuccessfulOutcomeXnApElementaryProceduresMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SuccessfulOutcomeXnApElementaryProceduresMultiError) AllErrors() []error { return m }

// SuccessfulOutcomeXnApElementaryProceduresValidationError is the validation
// error returned by SuccessfulOutcomeXnApElementaryProcedures.Validate if the
// designated constraints aren't met.
type SuccessfulOutcomeXnApElementaryProceduresValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SuccessfulOutcomeXnApElementaryProceduresValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SuccessfulOutcomeXnApElementaryProceduresValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SuccessfulOutcomeXnApElementaryProceduresValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SuccessfulOutcomeXnApElementaryProceduresValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SuccessfulOutcomeXnApElementaryProceduresValidationError) ErrorName() string {
	return "SuccessfulOutcomeXnApElementaryProceduresValidationError"
}

// Error satisfies the builtin error interface
func (e SuccessfulOutcomeXnApElementaryProceduresValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSuccessfulOutcomeXnApElementaryProcedures.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SuccessfulOutcomeXnApElementaryProceduresValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SuccessfulOutcomeXnApElementaryProceduresValidationError{}

// Validate checks the field values on UnsuccessfulOutcome with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UnsuccessfulOutcome) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UnsuccessfulOutcome with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UnsuccessfulOutcomeMultiError, or nil if none found.
func (m *UnsuccessfulOutcome) ValidateAll() error {
	return m.validate(true)
}

func (m *UnsuccessfulOutcome) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ProcedureCode

	// no validation rules for Criticality

	if all {
		switch v := interface{}(m.GetValue()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UnsuccessfulOutcomeValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UnsuccessfulOutcomeValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UnsuccessfulOutcomeValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UnsuccessfulOutcomeMultiError(errors)
	}

	return nil
}

// UnsuccessfulOutcomeMultiError is an error wrapping multiple validation
// errors returned by UnsuccessfulOutcome.ValidateAll() if the designated
// constraints aren't met.
type UnsuccessfulOutcomeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UnsuccessfulOutcomeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UnsuccessfulOutcomeMultiError) AllErrors() []error { return m }

// UnsuccessfulOutcomeValidationError is the validation error returned by
// UnsuccessfulOutcome.Validate if the designated constraints aren't met.
type UnsuccessfulOutcomeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnsuccessfulOutcomeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnsuccessfulOutcomeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnsuccessfulOutcomeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnsuccessfulOutcomeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnsuccessfulOutcomeValidationError) ErrorName() string {
	return "UnsuccessfulOutcomeValidationError"
}

// Error satisfies the builtin error interface
func (e UnsuccessfulOutcomeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnsuccessfulOutcome.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnsuccessfulOutcomeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnsuccessfulOutcomeValidationError{}

// Validate checks the field values on
// UnsuccessfulOutcomeXnApElementaryProcedures with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UnsuccessfulOutcomeXnApElementaryProcedures) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// UnsuccessfulOutcomeXnApElementaryProcedures with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// UnsuccessfulOutcomeXnApElementaryProceduresMultiError, or nil if none found.
func (m *UnsuccessfulOutcomeXnApElementaryProcedures) ValidateAll() error {
	return m.validate(true)
}

func (m *UnsuccessfulOutcomeXnApElementaryProcedures) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.UoValues.(type) {

	case *UnsuccessfulOutcomeXnApElementaryProcedures_XnSetupFailure:

		if all {
			switch v := interface{}(m.GetXnSetupFailure()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, UnsuccessfulOutcomeXnApElementaryProceduresValidationError{
						field:  "XnSetupFailure",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, UnsuccessfulOutcomeXnApElementaryProceduresValidationError{
						field:  "XnSetupFailure",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetXnSetupFailure()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UnsuccessfulOutcomeXnApElementaryProceduresValidationError{
					field:  "XnSetupFailure",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return UnsuccessfulOutcomeXnApElementaryProceduresMultiError(errors)
	}

	return nil
}

// UnsuccessfulOutcomeXnApElementaryProceduresMultiError is an error wrapping
// multiple validation errors returned by
// UnsuccessfulOutcomeXnApElementaryProcedures.ValidateAll() if the designated
// constraints aren't met.
type UnsuccessfulOutcomeXnApElementaryProceduresMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UnsuccessfulOutcomeXnApElementaryProceduresMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UnsuccessfulOutcomeXnApElementaryProceduresMultiError) AllErrors() []error { return m }

// UnsuccessfulOutcomeXnApElementaryProceduresValidationError is the validation
// error returned by UnsuccessfulOutcomeXnApElementaryProcedures.Validate if
// the designated constraints aren't met.
type UnsuccessfulOutcomeXnApElementaryProceduresValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UnsuccessfulOutcomeXnApElementaryProceduresValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UnsuccessfulOutcomeXnApElementaryProceduresValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UnsuccessfulOutcomeXnApElementaryProceduresValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UnsuccessfulOutcomeXnApElementaryProceduresValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UnsuccessfulOutcomeXnApElementaryProceduresValidationError) ErrorName() string {
	return "UnsuccessfulOutcomeXnApElementaryProceduresValidationError"
}

// Error satisfies the builtin error interface
func (e UnsuccessfulOutcomeXnApElementaryProceduresValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUnsuccessfulOutcomeXnApElementaryProcedures.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UnsuccessfulOutcomeXnApElementaryProceduresValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UnsuccessfulOutcomeXnApElementaryProceduresValidationError{}
