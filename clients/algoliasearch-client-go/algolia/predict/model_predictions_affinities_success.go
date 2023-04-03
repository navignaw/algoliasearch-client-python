// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package predict

import (
	"encoding/json"
	"fmt"
)

// PredictionsAffinitiesSuccess Prediction for the **affinities** model.
type PredictionsAffinitiesSuccess struct {
	Value         []Affinity `json:"value"`
	LastUpdatedAt string     `json:"lastUpdatedAt"`
}

// NewPredictionsAffinitiesSuccess instantiates a new PredictionsAffinitiesSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPredictionsAffinitiesSuccess(value []Affinity, lastUpdatedAt string) *PredictionsAffinitiesSuccess {
	this := &PredictionsAffinitiesSuccess{}
	this.Value = value
	this.LastUpdatedAt = lastUpdatedAt
	return this
}

// NewPredictionsAffinitiesSuccessWithDefaults instantiates a new PredictionsAffinitiesSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPredictionsAffinitiesSuccessWithDefaults() *PredictionsAffinitiesSuccess {
	this := &PredictionsAffinitiesSuccess{}
	return this
}

// GetValue returns the Value field value
func (o *PredictionsAffinitiesSuccess) GetValue() []Affinity {
	if o == nil {
		var ret []Affinity
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PredictionsAffinitiesSuccess) GetValueOk() ([]Affinity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *PredictionsAffinitiesSuccess) SetValue(v []Affinity) {
	o.Value = v
}

// GetLastUpdatedAt returns the LastUpdatedAt field value
func (o *PredictionsAffinitiesSuccess) GetLastUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdatedAt
}

// GetLastUpdatedAtOk returns a tuple with the LastUpdatedAt field value
// and a boolean to check if the value has been set.
func (o *PredictionsAffinitiesSuccess) GetLastUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdatedAt, true
}

// SetLastUpdatedAt sets field value
func (o *PredictionsAffinitiesSuccess) SetLastUpdatedAt(v string) {
	o.LastUpdatedAt = v
}

func (o PredictionsAffinitiesSuccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["lastUpdatedAt"] = o.LastUpdatedAt
	}
	return json.Marshal(toSerialize)
}

func (o PredictionsAffinitiesSuccess) String() string {
	out := "PredictionsAffinitiesSuccess {\n"
	out += fmt.Sprintf("  value=%v\n", o.Value)
	out += fmt.Sprintf("  lastUpdatedAt=%v\n", o.LastUpdatedAt)
	out += "}"
	return out
}

type NullablePredictionsAffinitiesSuccess struct {
	value *PredictionsAffinitiesSuccess
	isSet bool
}

func (v NullablePredictionsAffinitiesSuccess) Get() *PredictionsAffinitiesSuccess {
	return v.value
}

func (v *NullablePredictionsAffinitiesSuccess) Set(val *PredictionsAffinitiesSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullablePredictionsAffinitiesSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullablePredictionsAffinitiesSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePredictionsAffinitiesSuccess(val *PredictionsAffinitiesSuccess) *NullablePredictionsAffinitiesSuccess {
	return &NullablePredictionsAffinitiesSuccess{value: val, isSet: true}
}

func (v NullablePredictionsAffinitiesSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePredictionsAffinitiesSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}