// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// BaseSearchResponseRedirect This parameter is for internal use only.
type BaseSearchResponseRedirect struct {
	Index []RedirectRuleIndexMetadata `json:"index,omitempty"`
}

type BaseSearchResponseRedirectOption func(f *BaseSearchResponseRedirect)

func WithBaseSearchResponseRedirectIndex(val []RedirectRuleIndexMetadata) BaseSearchResponseRedirectOption {
	return func(f *BaseSearchResponseRedirect) {
		f.Index = val
	}
}

// NewBaseSearchResponseRedirect instantiates a new BaseSearchResponseRedirect object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseSearchResponseRedirect(opts ...BaseSearchResponseRedirectOption) *BaseSearchResponseRedirect {
	this := &BaseSearchResponseRedirect{}
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewBaseSearchResponseRedirectWithDefaults instantiates a new BaseSearchResponseRedirect object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseSearchResponseRedirectWithDefaults() *BaseSearchResponseRedirect {
	this := &BaseSearchResponseRedirect{}
	return this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *BaseSearchResponseRedirect) GetIndex() []RedirectRuleIndexMetadata {
	if o == nil || o.Index == nil {
		var ret []RedirectRuleIndexMetadata
		return ret
	}
	return o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSearchResponseRedirect) GetIndexOk() ([]RedirectRuleIndexMetadata, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *BaseSearchResponseRedirect) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given []RedirectRuleIndexMetadata and assigns it to the Index field.
func (o *BaseSearchResponseRedirect) SetIndex(v []RedirectRuleIndexMetadata) {
	o.Index = v
}

func (o BaseSearchResponseRedirect) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	return json.Marshal(toSerialize)
}

func (o BaseSearchResponseRedirect) String() string {
	out := ""
	out += fmt.Sprintf("  index=%v\n", o.Index)
	return fmt.Sprintf("BaseSearchResponseRedirect {\n%s}", out)
}

type NullableBaseSearchResponseRedirect struct {
	value *BaseSearchResponseRedirect
	isSet bool
}

func (v NullableBaseSearchResponseRedirect) Get() *BaseSearchResponseRedirect {
	return v.value
}

func (v *NullableBaseSearchResponseRedirect) Set(val *BaseSearchResponseRedirect) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseSearchResponseRedirect) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseSearchResponseRedirect) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseSearchResponseRedirect(val *BaseSearchResponseRedirect) *NullableBaseSearchResponseRedirect {
	return &NullableBaseSearchResponseRedirect{value: val, isSet: true}
}

func (v NullableBaseSearchResponseRedirect) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseSearchResponseRedirect) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
