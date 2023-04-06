// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package search

import (
	"encoding/json"
	"fmt"
)

// AddApiKeyResponse struct for AddApiKeyResponse
type AddApiKeyResponse struct {
	// The API key.
	Key string `json:"key"`
	// Date of creation (ISO-8601 format).
	CreatedAt string `json:"createdAt"`
}

// NewAddApiKeyResponse instantiates a new AddApiKeyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddApiKeyResponse(key string, createdAt string) *AddApiKeyResponse {
	this := &AddApiKeyResponse{}
	this.Key = key
	this.CreatedAt = createdAt
	return this
}

// NewAddApiKeyResponseWithDefaults instantiates a new AddApiKeyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddApiKeyResponseWithDefaults() *AddApiKeyResponse {
	this := &AddApiKeyResponse{}
	return this
}

// GetKey returns the Key field value
func (o *AddApiKeyResponse) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *AddApiKeyResponse) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *AddApiKeyResponse) SetKey(v string) {
	o.Key = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AddApiKeyResponse) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AddApiKeyResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AddApiKeyResponse) SetCreatedAt(v string) {
	o.CreatedAt = v
}

func (o AddApiKeyResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["key"] = o.Key
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

func (o AddApiKeyResponse) String() string {
	out := ""
	out += fmt.Sprintf("  key=%v\n", o.Key)
	out += fmt.Sprintf("  createdAt=%v\n", o.CreatedAt)
	return fmt.Sprintf("AddApiKeyResponse {\n%s}", out)
}

type NullableAddApiKeyResponse struct {
	value *AddApiKeyResponse
	isSet bool
}

func (v NullableAddApiKeyResponse) Get() *AddApiKeyResponse {
	return v.value
}

func (v *NullableAddApiKeyResponse) Set(val *AddApiKeyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAddApiKeyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAddApiKeyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddApiKeyResponse(val *AddApiKeyResponse) *NullableAddApiKeyResponse {
	return &NullableAddApiKeyResponse{value: val, isSet: true}
}

func (v NullableAddApiKeyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddApiKeyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
