// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package ingestion

import (
	"encoding/json"
	"fmt"
)

// ScheduleDateUtilsInput The input for a `schedule` task whose source is of type `bigquery` and for which extracted data spans a fixed number of days.
type ScheduleDateUtilsInput struct {
	// The timeframe of the extraction, in number of days from today.
	Timeframe int32 `json:"timeframe" validate:"required"`
}

// NewScheduleDateUtilsInput instantiates a new ScheduleDateUtilsInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleDateUtilsInput(timeframe int32) *ScheduleDateUtilsInput {
	this := &ScheduleDateUtilsInput{}
	this.Timeframe = timeframe
	return this
}

// NewScheduleDateUtilsInputWithDefaults instantiates a new ScheduleDateUtilsInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleDateUtilsInputWithDefaults() *ScheduleDateUtilsInput {
	this := &ScheduleDateUtilsInput{}
	return this
}

// GetTimeframe returns the Timeframe field value
func (o *ScheduleDateUtilsInput) GetTimeframe() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timeframe
}

// GetTimeframeOk returns a tuple with the Timeframe field value
// and a boolean to check if the value has been set.
func (o *ScheduleDateUtilsInput) GetTimeframeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeframe, true
}

// SetTimeframe sets field value
func (o *ScheduleDateUtilsInput) SetTimeframe(v int32) {
	o.Timeframe = v
}

func (o ScheduleDateUtilsInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["timeframe"] = o.Timeframe
	}
	return json.Marshal(toSerialize)
}

func (o ScheduleDateUtilsInput) String() string {
	out := ""
	out += fmt.Sprintf("  timeframe=%v\n", o.Timeframe)
	return fmt.Sprintf("ScheduleDateUtilsInput {\n%s}", out)
}

type NullableScheduleDateUtilsInput struct {
	value *ScheduleDateUtilsInput
	isSet bool
}

func (v NullableScheduleDateUtilsInput) Get() *ScheduleDateUtilsInput {
	return v.value
}

func (v *NullableScheduleDateUtilsInput) Set(val *ScheduleDateUtilsInput) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleDateUtilsInput) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleDateUtilsInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleDateUtilsInput(val *ScheduleDateUtilsInput) *NullableScheduleDateUtilsInput {
	return &NullableScheduleDateUtilsInput{value: val, isSet: true}
}

func (v NullableScheduleDateUtilsInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleDateUtilsInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
