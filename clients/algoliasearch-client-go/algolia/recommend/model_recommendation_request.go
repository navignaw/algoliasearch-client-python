// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package recommend

import (
	"encoding/json"
	"fmt"
)

// RecommendationRequest struct for RecommendationRequest
type RecommendationRequest struct {
	Model RecommendationModels `json:"model"`
	// Unique identifier of the object.
	ObjectID string `json:"objectID"`
	// The Algolia index name.
	IndexName string `json:"indexName"`
	// The threshold to use when filtering recommendations by their score.
	Threshold int32 `json:"threshold"`
	// The max number of recommendations to retrieve. If it's set to 0, all the recommendations of the objectID may be returned.
	MaxRecommendations *int32              `json:"maxRecommendations,omitempty"`
	QueryParameters    *SearchParamsObject `json:"queryParameters,omitempty"`
	FallbackParameters *SearchParamsObject `json:"fallbackParameters,omitempty"`
}

type RecommendationRequestOption func(f *RecommendationRequest)

func WithRecommendationRequestMaxRecommendations(val int32) RecommendationRequestOption {
	return func(f *RecommendationRequest) {
		f.MaxRecommendations = &val
	}
}

func WithRecommendationRequestQueryParameters(val SearchParamsObject) RecommendationRequestOption {
	return func(f *RecommendationRequest) {
		f.QueryParameters = &val
	}
}

func WithRecommendationRequestFallbackParameters(val SearchParamsObject) RecommendationRequestOption {
	return func(f *RecommendationRequest) {
		f.FallbackParameters = &val
	}
}

// NewRecommendationRequest instantiates a new RecommendationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationRequest(model RecommendationModels, objectID string, indexName string, threshold int32, opts ...RecommendationRequestOption) *RecommendationRequest {
	this := &RecommendationRequest{}
	this.Model = model
	this.ObjectID = objectID
	this.IndexName = indexName
	this.Threshold = threshold
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewRecommendationRequestWithDefaults instantiates a new RecommendationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationRequestWithDefaults() *RecommendationRequest {
	this := &RecommendationRequest{}
	var maxRecommendations int32 = 0
	this.MaxRecommendations = &maxRecommendations
	return this
}

// GetModel returns the Model field value
func (o *RecommendationRequest) GetModel() RecommendationModels {
	if o == nil {
		var ret RecommendationModels
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *RecommendationRequest) GetModelOk() (*RecommendationModels, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *RecommendationRequest) SetModel(v RecommendationModels) {
	o.Model = v
}

// GetObjectID returns the ObjectID field value
func (o *RecommendationRequest) GetObjectID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectID
}

// GetObjectIDOk returns a tuple with the ObjectID field value
// and a boolean to check if the value has been set.
func (o *RecommendationRequest) GetObjectIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectID, true
}

// SetObjectID sets field value
func (o *RecommendationRequest) SetObjectID(v string) {
	o.ObjectID = v
}

// GetIndexName returns the IndexName field value
func (o *RecommendationRequest) GetIndexName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IndexName
}

// GetIndexNameOk returns a tuple with the IndexName field value
// and a boolean to check if the value has been set.
func (o *RecommendationRequest) GetIndexNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IndexName, true
}

// SetIndexName sets field value
func (o *RecommendationRequest) SetIndexName(v string) {
	o.IndexName = v
}

// GetThreshold returns the Threshold field value
func (o *RecommendationRequest) GetThreshold() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value
// and a boolean to check if the value has been set.
func (o *RecommendationRequest) GetThresholdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threshold, true
}

// SetThreshold sets field value
func (o *RecommendationRequest) SetThreshold(v int32) {
	o.Threshold = v
}

// GetMaxRecommendations returns the MaxRecommendations field value if set, zero value otherwise.
func (o *RecommendationRequest) GetMaxRecommendations() int32 {
	if o == nil || o.MaxRecommendations == nil {
		var ret int32
		return ret
	}
	return *o.MaxRecommendations
}

// GetMaxRecommendationsOk returns a tuple with the MaxRecommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationRequest) GetMaxRecommendationsOk() (*int32, bool) {
	if o == nil || o.MaxRecommendations == nil {
		return nil, false
	}
	return o.MaxRecommendations, true
}

// HasMaxRecommendations returns a boolean if a field has been set.
func (o *RecommendationRequest) HasMaxRecommendations() bool {
	if o != nil && o.MaxRecommendations != nil {
		return true
	}

	return false
}

// SetMaxRecommendations gets a reference to the given int32 and assigns it to the MaxRecommendations field.
func (o *RecommendationRequest) SetMaxRecommendations(v int32) {
	o.MaxRecommendations = &v
}

// GetQueryParameters returns the QueryParameters field value if set, zero value otherwise.
func (o *RecommendationRequest) GetQueryParameters() SearchParamsObject {
	if o == nil || o.QueryParameters == nil {
		var ret SearchParamsObject
		return ret
	}
	return *o.QueryParameters
}

// GetQueryParametersOk returns a tuple with the QueryParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationRequest) GetQueryParametersOk() (*SearchParamsObject, bool) {
	if o == nil || o.QueryParameters == nil {
		return nil, false
	}
	return o.QueryParameters, true
}

// HasQueryParameters returns a boolean if a field has been set.
func (o *RecommendationRequest) HasQueryParameters() bool {
	if o != nil && o.QueryParameters != nil {
		return true
	}

	return false
}

// SetQueryParameters gets a reference to the given SearchParamsObject and assigns it to the QueryParameters field.
func (o *RecommendationRequest) SetQueryParameters(v SearchParamsObject) {
	o.QueryParameters = &v
}

// GetFallbackParameters returns the FallbackParameters field value if set, zero value otherwise.
func (o *RecommendationRequest) GetFallbackParameters() SearchParamsObject {
	if o == nil || o.FallbackParameters == nil {
		var ret SearchParamsObject
		return ret
	}
	return *o.FallbackParameters
}

// GetFallbackParametersOk returns a tuple with the FallbackParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationRequest) GetFallbackParametersOk() (*SearchParamsObject, bool) {
	if o == nil || o.FallbackParameters == nil {
		return nil, false
	}
	return o.FallbackParameters, true
}

// HasFallbackParameters returns a boolean if a field has been set.
func (o *RecommendationRequest) HasFallbackParameters() bool {
	if o != nil && o.FallbackParameters != nil {
		return true
	}

	return false
}

// SetFallbackParameters gets a reference to the given SearchParamsObject and assigns it to the FallbackParameters field.
func (o *RecommendationRequest) SetFallbackParameters(v SearchParamsObject) {
	o.FallbackParameters = &v
}

func (o RecommendationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["model"] = o.Model
	}
	if true {
		toSerialize["objectID"] = o.ObjectID
	}
	if true {
		toSerialize["indexName"] = o.IndexName
	}
	if true {
		toSerialize["threshold"] = o.Threshold
	}
	if o.MaxRecommendations != nil {
		toSerialize["maxRecommendations"] = o.MaxRecommendations
	}
	if o.QueryParameters != nil {
		toSerialize["queryParameters"] = o.QueryParameters
	}
	if o.FallbackParameters != nil {
		toSerialize["fallbackParameters"] = o.FallbackParameters
	}
	return json.Marshal(toSerialize)
}

func (o RecommendationRequest) String() string {
	out := ""
	out += fmt.Sprintf("  model=%v\n", o.Model)
	out += fmt.Sprintf("  objectID=%v\n", o.ObjectID)
	out += fmt.Sprintf("  indexName=%v\n", o.IndexName)
	out += fmt.Sprintf("  threshold=%v\n", o.Threshold)
	out += fmt.Sprintf("  maxRecommendations=%v\n", o.MaxRecommendations)
	out += fmt.Sprintf("  queryParameters=%v\n", o.QueryParameters)
	out += fmt.Sprintf("  fallbackParameters=%v\n", o.FallbackParameters)
	return fmt.Sprintf("RecommendationRequest {\n%s}", out)
}

type NullableRecommendationRequest struct {
	value *RecommendationRequest
	isSet bool
}

func (v NullableRecommendationRequest) Get() *RecommendationRequest {
	return v.value
}

func (v *NullableRecommendationRequest) Set(val *RecommendationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationRequest(val *RecommendationRequest) *NullableRecommendationRequest {
	return &NullableRecommendationRequest{value: val, isSet: true}
}

func (v NullableRecommendationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
