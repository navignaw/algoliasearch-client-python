// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package suggestions

import (
	"encoding/json"
	"fmt"
)

// ErrorBase Error.
type ErrorBase struct {
	Message              *string `json:"message,omitempty"`
	AdditionalProperties map[string]any
}

type _ErrorBase ErrorBase

type ErrorBaseOption func(f *ErrorBase)

func WithErrorBaseMessage(val string) ErrorBaseOption {
	return func(f *ErrorBase) {
		f.Message = &val
	}
}

// NewErrorBase instantiates a new ErrorBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorBase(opts ...ErrorBaseOption) *ErrorBase {
	this := &ErrorBase{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewErrorBaseWithDefaults instantiates a new ErrorBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorBaseWithDefaults() *ErrorBase {
	this := &ErrorBase{}
	return this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorBase) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorBase) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorBase) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorBase) SetMessage(v string) {
	o.Message = &v
}

func (o ErrorBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ErrorBase) UnmarshalJSON(bytes []byte) (err error) {
	varErrorBase := _ErrorBase{}

	if err = json.Unmarshal(bytes, &varErrorBase); err == nil {
		*o = ErrorBase(varErrorBase)
	}

	additionalProperties := make(map[string]any)

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

func (o ErrorBase) String() string {
	out := ""
	out += fmt.Sprintf("  message=%v\n", o.Message)
	for key, value := range o.AdditionalProperties {
		out += fmt.Sprintf("  %s=%v\n", key, value)
	}
	return fmt.Sprintf("ErrorBase {\n%s}", out)
}

type NullableErrorBase struct {
	value *ErrorBase
	isSet bool
}

func (v NullableErrorBase) Get() *ErrorBase {
	return v.value
}

func (v *NullableErrorBase) Set(val *ErrorBase) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorBase) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorBase(val *ErrorBase) *NullableErrorBase {
	return &NullableErrorBase{value: val, isSet: true}
}

func (v NullableErrorBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
