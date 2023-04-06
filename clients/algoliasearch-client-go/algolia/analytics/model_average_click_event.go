// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package analytics

import (
	"encoding/json"
	"fmt"
)

// AverageClickEvent struct for AverageClickEvent
type AverageClickEvent struct {
	// The average of all the click count event.
	Average float64 `json:"average"`
	// The number of click event.
	ClickCount int32 `json:"clickCount"`
	// Date of the event.
	Date string `json:"date"`
}

// NewAverageClickEvent instantiates a new AverageClickEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAverageClickEvent(average float64, clickCount int32, date string) *AverageClickEvent {
	this := &AverageClickEvent{}
	this.Average = average
	this.ClickCount = clickCount
	this.Date = date
	return this
}

// NewAverageClickEventWithDefaults instantiates a new AverageClickEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAverageClickEventWithDefaults() *AverageClickEvent {
	this := &AverageClickEvent{}
	return this
}

// GetAverage returns the Average field value
func (o *AverageClickEvent) GetAverage() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Average
}

// GetAverageOk returns a tuple with the Average field value
// and a boolean to check if the value has been set.
func (o *AverageClickEvent) GetAverageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Average, true
}

// SetAverage sets field value
func (o *AverageClickEvent) SetAverage(v float64) {
	o.Average = v
}

// GetClickCount returns the ClickCount field value
func (o *AverageClickEvent) GetClickCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClickCount
}

// GetClickCountOk returns a tuple with the ClickCount field value
// and a boolean to check if the value has been set.
func (o *AverageClickEvent) GetClickCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClickCount, true
}

// SetClickCount sets field value
func (o *AverageClickEvent) SetClickCount(v int32) {
	o.ClickCount = v
}

// GetDate returns the Date field value
func (o *AverageClickEvent) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *AverageClickEvent) GetDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *AverageClickEvent) SetDate(v string) {
	o.Date = v
}

func (o AverageClickEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["average"] = o.Average
	}
	if true {
		toSerialize["clickCount"] = o.ClickCount
	}
	if true {
		toSerialize["date"] = o.Date
	}
	return json.Marshal(toSerialize)
}

func (o AverageClickEvent) String() string {
	out := ""
	out += fmt.Sprintf("  average=%v\n", o.Average)
	out += fmt.Sprintf("  clickCount=%v\n", o.ClickCount)
	out += fmt.Sprintf("  date=%v\n", o.Date)
	return fmt.Sprintf("AverageClickEvent {\n%s}", out)
}

type NullableAverageClickEvent struct {
	value *AverageClickEvent
	isSet bool
}

func (v NullableAverageClickEvent) Get() *AverageClickEvent {
	return v.value
}

func (v *NullableAverageClickEvent) Set(val *AverageClickEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAverageClickEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAverageClickEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAverageClickEvent(val *AverageClickEvent) *NullableAverageClickEvent {
	return &NullableAverageClickEvent{value: val, isSet: true}
}

func (v NullableAverageClickEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAverageClickEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
