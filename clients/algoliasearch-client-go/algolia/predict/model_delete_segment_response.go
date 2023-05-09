// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package predict

import (
	"encoding/json"
	"fmt"
)

// DeleteSegmentResponse struct for DeleteSegmentResponse
type DeleteSegmentResponse struct {
	// The ID of the segment.
	SegmentID string `json:"segmentID" validate:"required"`
	// The date and time at which the segment will be re-ingested.
	DeletedUntil string `json:"deletedUntil" validate:"required"`
}

// NewDeleteSegmentResponse instantiates a new DeleteSegmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteSegmentResponse(segmentID string, deletedUntil string) *DeleteSegmentResponse {
	this := &DeleteSegmentResponse{}
	this.SegmentID = segmentID
	this.DeletedUntil = deletedUntil
	return this
}

// NewDeleteSegmentResponseWithDefaults instantiates a new DeleteSegmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteSegmentResponseWithDefaults() *DeleteSegmentResponse {
	this := &DeleteSegmentResponse{}
	return this
}

// GetSegmentID returns the SegmentID field value
func (o *DeleteSegmentResponse) GetSegmentID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SegmentID
}

// GetSegmentIDOk returns a tuple with the SegmentID field value
// and a boolean to check if the value has been set.
func (o *DeleteSegmentResponse) GetSegmentIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SegmentID, true
}

// SetSegmentID sets field value
func (o *DeleteSegmentResponse) SetSegmentID(v string) {
	o.SegmentID = v
}

// GetDeletedUntil returns the DeletedUntil field value
func (o *DeleteSegmentResponse) GetDeletedUntil() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeletedUntil
}

// GetDeletedUntilOk returns a tuple with the DeletedUntil field value
// and a boolean to check if the value has been set.
func (o *DeleteSegmentResponse) GetDeletedUntilOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedUntil, true
}

// SetDeletedUntil sets field value
func (o *DeleteSegmentResponse) SetDeletedUntil(v string) {
	o.DeletedUntil = v
}

func (o DeleteSegmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["segmentID"] = o.SegmentID
	}
	if true {
		toSerialize["deletedUntil"] = o.DeletedUntil
	}
	return json.Marshal(toSerialize)
}

func (o DeleteSegmentResponse) String() string {
	out := ""
	out += fmt.Sprintf("  segmentID=%v\n", o.SegmentID)
	out += fmt.Sprintf("  deletedUntil=%v\n", o.DeletedUntil)
	return fmt.Sprintf("DeleteSegmentResponse {\n%s}", out)
}

type NullableDeleteSegmentResponse struct {
	value *DeleteSegmentResponse
	isSet bool
}

func (v NullableDeleteSegmentResponse) Get() *DeleteSegmentResponse {
	return v.value
}

func (v *NullableDeleteSegmentResponse) Set(val *DeleteSegmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteSegmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteSegmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteSegmentResponse(val *DeleteSegmentResponse) *NullableDeleteSegmentResponse {
	return &NullableDeleteSegmentResponse{value: val, isSet: true}
}

func (v NullableDeleteSegmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteSegmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
