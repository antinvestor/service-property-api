// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: property.proto

package property_v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on Locality with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Locality) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetID() != "" {

		if l := utf8.RuneCountInString(m.GetID()); l < 3 || l > 40 {
			return LocalityValidationError{
				field:  "ID",
				reason: "value length must be between 3 and 40 runes, inclusive",
			}
		}

		if !_Locality_ID_Pattern.MatchString(m.GetID()) {
			return LocalityValidationError{
				field:  "ID",
				reason: "value does not match regex pattern \"[0-9a-z_-]{3,20}\"",
			}
		}

	}

	if m.GetParentID() != "" {

		if l := utf8.RuneCountInString(m.GetParentID()); l < 3 || l > 40 {
			return LocalityValidationError{
				field:  "ParentID",
				reason: "value length must be between 3 and 40 runes, inclusive",
			}
		}

		if !_Locality_ParentID_Pattern.MatchString(m.GetParentID()) {
			return LocalityValidationError{
				field:  "ParentID",
				reason: "value does not match regex pattern \"[0-9a-z_-]{3,20}\"",
			}
		}

	}

	if l := utf8.RuneCountInString(m.GetBoundary()); l < 10 || l > 2048 {
		return LocalityValidationError{
			field:  "Boundary",
			reason: "value length must be between 10 and 2048 runes, inclusive",
		}
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 3 || l > 60 {
		return LocalityValidationError{
			field:  "Name",
			reason: "value length must be between 3 and 60 runes, inclusive",
		}
	}

	if m.GetDescription() != "" {

		if utf8.RuneCountInString(m.GetDescription()) < 50 {
			return LocalityValidationError{
				field:  "Description",
				reason: "value length must be at least 50 runes",
			}
		}

	}

	// no validation rules for Extras

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocalityValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// LocalityValidationError is the validation error returned by
// Locality.Validate if the designated constraints aren't met.
type LocalityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocalityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocalityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocalityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocalityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocalityValidationError) ErrorName() string { return "LocalityValidationError" }

// Error satisfies the builtin error interface
func (e LocalityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocality.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocalityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocalityValidationError{}

var _Locality_ID_Pattern = regexp.MustCompile("[0-9a-z_-]{3,20}")

var _Locality_ParentID_Pattern = regexp.MustCompile("[0-9a-z_-]{3,20}")

// Validate checks the field values on PropertyState with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *PropertyState) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetID()); l < 3 || l > 40 {
		return PropertyStateValidationError{
			field:  "ID",
			reason: "value length must be between 3 and 40 runes, inclusive",
		}
	}

	if !_PropertyState_ID_Pattern.MatchString(m.GetID()) {
		return PropertyStateValidationError{
			field:  "ID",
			reason: "value does not match regex pattern \"[0-9a-z_-]{3,20}\"",
		}
	}

	if l := utf8.RuneCountInString(m.GetPropertyID()); l < 3 || l > 40 {
		return PropertyStateValidationError{
			field:  "PropertyID",
			reason: "value length must be between 3 and 40 runes, inclusive",
		}
	}

	if !_PropertyState_PropertyID_Pattern.MatchString(m.GetPropertyID()) {
		return PropertyStateValidationError{
			field:  "PropertyID",
			reason: "value does not match regex pattern \"[0-9a-z_-]{3,20}\"",
		}
	}

	// no validation rules for State

	// no validation rules for Status

	if l := utf8.RuneCountInString(m.GetName()); l < 3 || l > 60 {
		return PropertyStateValidationError{
			field:  "Name",
			reason: "value length must be between 3 and 60 runes, inclusive",
		}
	}

	if m.GetDescription() != "" {

		if utf8.RuneCountInString(m.GetDescription()) < 50 {
			return PropertyStateValidationError{
				field:  "Description",
				reason: "value length must be at least 50 runes",
			}
		}

	}

	// no validation rules for Extras

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PropertyStateValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// PropertyStateValidationError is the validation error returned by
// PropertyState.Validate if the designated constraints aren't met.
type PropertyStateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PropertyStateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PropertyStateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PropertyStateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PropertyStateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PropertyStateValidationError) ErrorName() string { return "PropertyStateValidationError" }

// Error satisfies the builtin error interface
func (e PropertyStateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPropertyState.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PropertyStateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PropertyStateValidationError{}

var _PropertyState_ID_Pattern = regexp.MustCompile("[0-9a-z_-]{3,20}")

var _PropertyState_PropertyID_Pattern = regexp.MustCompile("[0-9a-z_-]{3,20}")

// Validate checks the field values on PropertyType with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *PropertyType) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetID() != "" {

		if l := utf8.RuneCountInString(m.GetID()); l < 3 || l > 40 {
			return PropertyTypeValidationError{
				field:  "ID",
				reason: "value length must be between 3 and 40 runes, inclusive",
			}
		}

		if !_PropertyType_ID_Pattern.MatchString(m.GetID()) {
			return PropertyTypeValidationError{
				field:  "ID",
				reason: "value does not match regex pattern \"[0-9a-z_-]{3,20}\"",
			}
		}

	}

	if l := utf8.RuneCountInString(m.GetName()); l < 3 || l > 60 {
		return PropertyTypeValidationError{
			field:  "Name",
			reason: "value length must be between 3 and 60 runes, inclusive",
		}
	}

	if m.GetDescription() != "" {

		if utf8.RuneCountInString(m.GetDescription()) < 50 {
			return PropertyTypeValidationError{
				field:  "Description",
				reason: "value length must be at least 50 runes",
			}
		}

	}

	// no validation rules for Extra

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PropertyTypeValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// PropertyTypeValidationError is the validation error returned by
// PropertyType.Validate if the designated constraints aren't met.
type PropertyTypeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PropertyTypeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PropertyTypeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PropertyTypeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PropertyTypeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PropertyTypeValidationError) ErrorName() string { return "PropertyTypeValidationError" }

// Error satisfies the builtin error interface
func (e PropertyTypeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPropertyType.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PropertyTypeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PropertyTypeValidationError{}

var _PropertyType_ID_Pattern = regexp.MustCompile("[0-9a-z_-]{3,20}")

// Validate checks the field values on Subscription with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Subscription) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetID() != "" {

		if l := utf8.RuneCountInString(m.GetID()); l < 3 || l > 40 {
			return SubscriptionValidationError{
				field:  "ID",
				reason: "value length must be between 3 and 40 runes, inclusive",
			}
		}

		if !_Subscription_ID_Pattern.MatchString(m.GetID()) {
			return SubscriptionValidationError{
				field:  "ID",
				reason: "value does not match regex pattern \"[0-9a-z_-]{3,20}\"",
			}
		}

	}

	if l := utf8.RuneCountInString(m.GetPropertyID()); l < 3 || l > 40 {
		return SubscriptionValidationError{
			field:  "PropertyID",
			reason: "value length must be between 3 and 40 runes, inclusive",
		}
	}

	if m.GetProfileID() != "" {

		if l := utf8.RuneCountInString(m.GetProfileID()); l < 3 || l > 50 {
			return SubscriptionValidationError{
				field:  "ProfileID",
				reason: "value length must be between 3 and 50 runes, inclusive",
			}
		}

	}

	if m.GetRole() != "" {

		if l := utf8.RuneCountInString(m.GetRole()); l < 3 || l > 50 {
			return SubscriptionValidationError{
				field:  "Role",
				reason: "value length must be between 3 and 50 runes, inclusive",
			}
		}

	}

	// no validation rules for Extra

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SubscriptionValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetExpiresAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SubscriptionValidationError{
				field:  "ExpiresAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SubscriptionValidationError is the validation error returned by
// Subscription.Validate if the designated constraints aren't met.
type SubscriptionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubscriptionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubscriptionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubscriptionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubscriptionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubscriptionValidationError) ErrorName() string { return "SubscriptionValidationError" }

// Error satisfies the builtin error interface
func (e SubscriptionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubscription.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubscriptionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubscriptionValidationError{}

var _Subscription_ID_Pattern = regexp.MustCompile("[0-9a-z_-]{3,20}")

// Validate checks the field values on Property with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Property) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetID() != "" {

		if l := utf8.RuneCountInString(m.GetID()); l < 3 || l > 40 {
			return PropertyValidationError{
				field:  "ID",
				reason: "value length must be between 3 and 40 runes, inclusive",
			}
		}

		if !_Property_ID_Pattern.MatchString(m.GetID()) {
			return PropertyValidationError{
				field:  "ID",
				reason: "value does not match regex pattern \"[0-9a-z_-]{3,20}\"",
			}
		}

	}

	if m.GetParentID() != "" {

		if l := utf8.RuneCountInString(m.GetParentID()); l < 3 || l > 40 {
			return PropertyValidationError{
				field:  "ParentID",
				reason: "value length must be between 3 and 40 runes, inclusive",
			}
		}

		if !_Property_ParentID_Pattern.MatchString(m.GetParentID()) {
			return PropertyValidationError{
				field:  "ParentID",
				reason: "value does not match regex pattern \"[0-9a-z_-]{3,20}\"",
			}
		}

	}

	if l := utf8.RuneCountInString(m.GetName()); l < 3 || l > 60 {
		return PropertyValidationError{
			field:  "Name",
			reason: "value length must be between 3 and 60 runes, inclusive",
		}
	}

	if m.GetDescription() != "" {

		if utf8.RuneCountInString(m.GetDescription()) < 50 {
			return PropertyValidationError{
				field:  "Description",
				reason: "value length must be at least 50 runes",
			}
		}

	}

	if m.GetPropertyType() == nil {
		return PropertyValidationError{
			field:  "PropertyType",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetPropertyType()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PropertyValidationError{
				field:  "PropertyType",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetLocality()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PropertyValidationError{
				field:  "Locality",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if t := m.GetStartedAt(); t != nil {
		ts, err := t.AsTime(), t.CheckValid()
		if err != nil {
			return PropertyValidationError{
				field:  "StartedAt",
				reason: "value is not a valid timestamp",
				cause:  err,
			}
		}

		now := time.Now()

		if ts.Sub(now) >= 0 {
			return PropertyValidationError{
				field:  "StartedAt",
				reason: "value must be less than now",
			}
		}

	}

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PropertyValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetExtra()) > 0 {

		for key, val := range m.GetExtra() {
			_ = val

			// no validation rules for Extra[key]

			if utf8.RuneCountInString(val) < 3 {
				return PropertyValidationError{
					field:  fmt.Sprintf("Extra[%v]", key),
					reason: "value length must be at least 3 runes",
				}
			}

		}

	}

	return nil
}

// PropertyValidationError is the validation error returned by
// Property.Validate if the designated constraints aren't met.
type PropertyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PropertyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PropertyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PropertyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PropertyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PropertyValidationError) ErrorName() string { return "PropertyValidationError" }

// Error satisfies the builtin error interface
func (e PropertyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProperty.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PropertyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PropertyValidationError{}

var _Property_ID_Pattern = regexp.MustCompile("[0-9a-z_-]{3,20}")

var _Property_ParentID_Pattern = regexp.MustCompile("[0-9a-z_-]{3,20}")

// Validate checks the field values on RequestID with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *RequestID) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetID()); l < 3 || l > 40 {
		return RequestIDValidationError{
			field:  "ID",
			reason: "value length must be between 3 and 40 runes, inclusive",
		}
	}

	if !_RequestID_ID_Pattern.MatchString(m.GetID()) {
		return RequestIDValidationError{
			field:  "ID",
			reason: "value does not match regex pattern \"[0-9a-z_-]{3,20}\"",
		}
	}

	return nil
}

// RequestIDValidationError is the validation error returned by
// RequestID.Validate if the designated constraints aren't met.
type RequestIDValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestIDValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestIDValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestIDValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestIDValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestIDValidationError) ErrorName() string { return "RequestIDValidationError" }

// Error satisfies the builtin error interface
func (e RequestIDValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequestID.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestIDValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestIDValidationError{}

var _RequestID_ID_Pattern = regexp.MustCompile("[0-9a-z_-]{3,20}")

// Validate checks the field values on UpdateRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UpdateRequest) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetID()); l < 3 || l > 40 {
		return UpdateRequestValidationError{
			field:  "ID",
			reason: "value length must be between 3 and 40 runes, inclusive",
		}
	}

	if !_UpdateRequest_ID_Pattern.MatchString(m.GetID()) {
		return UpdateRequestValidationError{
			field:  "ID",
			reason: "value does not match regex pattern \"[0-9a-z_-]{3,20}\"",
		}
	}

	// no validation rules for State

	// no validation rules for Status

	if m.GetName() != "" {

		if l := utf8.RuneCountInString(m.GetName()); l < 3 || l > 60 {
			return UpdateRequestValidationError{
				field:  "Name",
				reason: "value length must be between 3 and 60 runes, inclusive",
			}
		}

	}

	if m.GetDescription() != "" {

		if utf8.RuneCountInString(m.GetDescription()) < 60 {
			return UpdateRequestValidationError{
				field:  "Description",
				reason: "value length must be at least 60 runes",
			}
		}

	}

	if m.GetGuardianID() != "" {

		if l := utf8.RuneCountInString(m.GetGuardianID()); l < 3 || l > 40 {
			return UpdateRequestValidationError{
				field:  "GuardianID",
				reason: "value length must be between 3 and 40 runes, inclusive",
			}
		}

		if !_UpdateRequest_GuardianID_Pattern.MatchString(m.GetGuardianID()) {
			return UpdateRequestValidationError{
				field:  "GuardianID",
				reason: "value does not match regex pattern \"[0-9a-z_-]{3,20}\"",
			}
		}

	}

	if m.GetLocalityID() != "" {

		if l := utf8.RuneCountInString(m.GetLocalityID()); l < 3 || l > 40 {
			return UpdateRequestValidationError{
				field:  "LocalityID",
				reason: "value length must be between 3 and 40 runes, inclusive",
			}
		}

		if !_UpdateRequest_LocalityID_Pattern.MatchString(m.GetLocalityID()) {
			return UpdateRequestValidationError{
				field:  "LocalityID",
				reason: "value does not match regex pattern \"[0-9a-z_-]{3,20}\"",
			}
		}

	}

	if len(m.GetExtras()) > 0 {

		for key, val := range m.GetExtras() {
			_ = val

			// no validation rules for Extras[key]

			if utf8.RuneCountInString(val) < 3 {
				return UpdateRequestValidationError{
					field:  fmt.Sprintf("Extras[%v]", key),
					reason: "value length must be at least 3 runes",
				}
			}

		}

	}

	return nil
}

// UpdateRequestValidationError is the validation error returned by
// UpdateRequest.Validate if the designated constraints aren't met.
type UpdateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRequestValidationError) ErrorName() string { return "UpdateRequestValidationError" }

// Error satisfies the builtin error interface
func (e UpdateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRequestValidationError{}

var _UpdateRequest_ID_Pattern = regexp.MustCompile("[0-9a-z_-]{3,20}")

var _UpdateRequest_GuardianID_Pattern = regexp.MustCompile("[0-9a-z_-]{3,20}")

var _UpdateRequest_LocalityID_Pattern = regexp.MustCompile("[0-9a-z_-]{3,20}")

// Validate checks the field values on SearchRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SearchRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Query

	return nil
}

// SearchRequestValidationError is the validation error returned by
// SearchRequest.Validate if the designated constraints aren't met.
type SearchRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchRequestValidationError) ErrorName() string { return "SearchRequestValidationError" }

// Error satisfies the builtin error interface
func (e SearchRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchRequestValidationError{}

// Validate checks the field values on SubscriptionListRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SubscriptionListRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for PropertyID

	// no validation rules for Query

	return nil
}

// SubscriptionListRequestValidationError is the validation error returned by
// SubscriptionListRequest.Validate if the designated constraints aren't met.
type SubscriptionListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SubscriptionListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SubscriptionListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SubscriptionListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SubscriptionListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SubscriptionListRequestValidationError) ErrorName() string {
	return "SubscriptionListRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SubscriptionListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSubscriptionListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SubscriptionListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SubscriptionListRequestValidationError{}
