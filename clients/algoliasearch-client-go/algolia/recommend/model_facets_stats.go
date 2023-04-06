// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package recommend

import (
	"encoding/json"
	"fmt"
)

// FacetsStats struct for FacetsStats
type FacetsStats struct {
	// The minimum value in the result set.
	Min *int32 `json:"min,omitempty"`
	// The maximum value in the result set.
	Max *int32 `json:"max,omitempty"`
	// The average facet value in the result set.
	Avg *int32 `json:"avg,omitempty"`
	// The sum of all values in the result set.
	Sum *int32 `json:"sum,omitempty"`
}

type FacetsStatsOption func(f *FacetsStats)

func WithFacetsStatsMin(val int32) FacetsStatsOption {
	return func(f *FacetsStats) {
		f.Min = &val
	}
}

func WithFacetsStatsMax(val int32) FacetsStatsOption {
	return func(f *FacetsStats) {
		f.Max = &val
	}
}

func WithFacetsStatsAvg(val int32) FacetsStatsOption {
	return func(f *FacetsStats) {
		f.Avg = &val
	}
}

func WithFacetsStatsSum(val int32) FacetsStatsOption {
	return func(f *FacetsStats) {
		f.Sum = &val
	}
}

// NewFacetsStats instantiates a new FacetsStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFacetsStats(opts ...FacetsStatsOption) *FacetsStats {
	this := &FacetsStats{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewFacetsStatsWithDefaults instantiates a new FacetsStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFacetsStatsWithDefaults() *FacetsStats {
	this := &FacetsStats{}
	return this
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *FacetsStats) GetMin() int32 {
	if o == nil || o.Min == nil {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetsStats) GetMinOk() (*int32, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *FacetsStats) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *FacetsStats) SetMin(v int32) {
	o.Min = &v
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *FacetsStats) GetMax() int32 {
	if o == nil || o.Max == nil {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetsStats) GetMaxOk() (*int32, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *FacetsStats) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *FacetsStats) SetMax(v int32) {
	o.Max = &v
}

// GetAvg returns the Avg field value if set, zero value otherwise.
func (o *FacetsStats) GetAvg() int32 {
	if o == nil || o.Avg == nil {
		var ret int32
		return ret
	}
	return *o.Avg
}

// GetAvgOk returns a tuple with the Avg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetsStats) GetAvgOk() (*int32, bool) {
	if o == nil || o.Avg == nil {
		return nil, false
	}
	return o.Avg, true
}

// HasAvg returns a boolean if a field has been set.
func (o *FacetsStats) HasAvg() bool {
	if o != nil && o.Avg != nil {
		return true
	}

	return false
}

// SetAvg gets a reference to the given int32 and assigns it to the Avg field.
func (o *FacetsStats) SetAvg(v int32) {
	o.Avg = &v
}

// GetSum returns the Sum field value if set, zero value otherwise.
func (o *FacetsStats) GetSum() int32 {
	if o == nil || o.Sum == nil {
		var ret int32
		return ret
	}
	return *o.Sum
}

// GetSumOk returns a tuple with the Sum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FacetsStats) GetSumOk() (*int32, bool) {
	if o == nil || o.Sum == nil {
		return nil, false
	}
	return o.Sum, true
}

// HasSum returns a boolean if a field has been set.
func (o *FacetsStats) HasSum() bool {
	if o != nil && o.Sum != nil {
		return true
	}

	return false
}

// SetSum gets a reference to the given int32 and assigns it to the Sum field.
func (o *FacetsStats) SetSum(v int32) {
	o.Sum = &v
}

func (o FacetsStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.Avg != nil {
		toSerialize["avg"] = o.Avg
	}
	if o.Sum != nil {
		toSerialize["sum"] = o.Sum
	}
	return json.Marshal(toSerialize)
}

func (o FacetsStats) String() string {
	out := ""
	out += fmt.Sprintf("  min=%v\n", o.Min)
	out += fmt.Sprintf("  max=%v\n", o.Max)
	out += fmt.Sprintf("  avg=%v\n", o.Avg)
	out += fmt.Sprintf("  sum=%v\n", o.Sum)
	return fmt.Sprintf("FacetsStats {\n%s}", out)
}

type NullableFacetsStats struct {
	value *FacetsStats
	isSet bool
}

func (v NullableFacetsStats) Get() *FacetsStats {
	return v.value
}

func (v *NullableFacetsStats) Set(val *FacetsStats) {
	v.value = val
	v.isSet = true
}

func (v NullableFacetsStats) IsSet() bool {
	return v.isSet
}

func (v *NullableFacetsStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFacetsStats(val *FacetsStats) *NullableFacetsStats {
	return &NullableFacetsStats{value: val, isSet: true}
}

func (v NullableFacetsStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFacetsStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
