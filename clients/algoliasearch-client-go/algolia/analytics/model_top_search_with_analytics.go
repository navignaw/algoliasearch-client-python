// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package analytics

import (
	"encoding/json"
	"fmt"
)

// TopSearchWithAnalytics struct for TopSearchWithAnalytics
type TopSearchWithAnalytics struct {
	// The search query.
	Search string `json:"search" validate:"required"`
	// The number of occurrences.
	Count int32 `json:"count" validate:"required"`
	// The click-through rate.
	ClickThroughRate float64 `json:"clickThroughRate" validate:"required"`
	// The average position of all the click count event.
	AverageClickPosition int32 `json:"averageClickPosition" validate:"required"`
	// The conversion rate.
	ConversionRate float64 `json:"conversionRate" validate:"required"`
	// The number of tracked search click.
	TrackedSearchCount int32 `json:"trackedSearchCount" validate:"required"`
	// The number of click event.
	ClickCount int32 `json:"clickCount" validate:"required"`
	// The number of converted clicks.
	ConversionCount int32 `json:"conversionCount" validate:"required"`
	// Number of hits that the search query matched.
	NbHits int32 `json:"nbHits" validate:"required"`
}

// NewTopSearchWithAnalytics instantiates a new TopSearchWithAnalytics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTopSearchWithAnalytics(search string, count int32, clickThroughRate float64, averageClickPosition int32, conversionRate float64, trackedSearchCount int32, clickCount int32, conversionCount int32, nbHits int32) *TopSearchWithAnalytics {
	this := &TopSearchWithAnalytics{}
	this.Search = search
	this.Count = count
	this.ClickThroughRate = clickThroughRate
	this.AverageClickPosition = averageClickPosition
	this.ConversionRate = conversionRate
	this.TrackedSearchCount = trackedSearchCount
	this.ClickCount = clickCount
	this.ConversionCount = conversionCount
	this.NbHits = nbHits
	return this
}

// NewTopSearchWithAnalyticsWithDefaults instantiates a new TopSearchWithAnalytics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopSearchWithAnalyticsWithDefaults() *TopSearchWithAnalytics {
	this := &TopSearchWithAnalytics{}
	return this
}

// GetSearch returns the Search field value
func (o *TopSearchWithAnalytics) GetSearch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Search
}

// GetSearchOk returns a tuple with the Search field value
// and a boolean to check if the value has been set.
func (o *TopSearchWithAnalytics) GetSearchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Search, true
}

// SetSearch sets field value
func (o *TopSearchWithAnalytics) SetSearch(v string) {
	o.Search = v
}

// GetCount returns the Count field value
func (o *TopSearchWithAnalytics) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *TopSearchWithAnalytics) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *TopSearchWithAnalytics) SetCount(v int32) {
	o.Count = v
}

// GetClickThroughRate returns the ClickThroughRate field value
func (o *TopSearchWithAnalytics) GetClickThroughRate() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.ClickThroughRate
}

// GetClickThroughRateOk returns a tuple with the ClickThroughRate field value
// and a boolean to check if the value has been set.
func (o *TopSearchWithAnalytics) GetClickThroughRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClickThroughRate, true
}

// SetClickThroughRate sets field value
func (o *TopSearchWithAnalytics) SetClickThroughRate(v float64) {
	o.ClickThroughRate = v
}

// GetAverageClickPosition returns the AverageClickPosition field value
func (o *TopSearchWithAnalytics) GetAverageClickPosition() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AverageClickPosition
}

// GetAverageClickPositionOk returns a tuple with the AverageClickPosition field value
// and a boolean to check if the value has been set.
func (o *TopSearchWithAnalytics) GetAverageClickPositionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AverageClickPosition, true
}

// SetAverageClickPosition sets field value
func (o *TopSearchWithAnalytics) SetAverageClickPosition(v int32) {
	o.AverageClickPosition = v
}

// GetConversionRate returns the ConversionRate field value
func (o *TopSearchWithAnalytics) GetConversionRate() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.ConversionRate
}

// GetConversionRateOk returns a tuple with the ConversionRate field value
// and a boolean to check if the value has been set.
func (o *TopSearchWithAnalytics) GetConversionRateOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversionRate, true
}

// SetConversionRate sets field value
func (o *TopSearchWithAnalytics) SetConversionRate(v float64) {
	o.ConversionRate = v
}

// GetTrackedSearchCount returns the TrackedSearchCount field value
func (o *TopSearchWithAnalytics) GetTrackedSearchCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TrackedSearchCount
}

// GetTrackedSearchCountOk returns a tuple with the TrackedSearchCount field value
// and a boolean to check if the value has been set.
func (o *TopSearchWithAnalytics) GetTrackedSearchCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrackedSearchCount, true
}

// SetTrackedSearchCount sets field value
func (o *TopSearchWithAnalytics) SetTrackedSearchCount(v int32) {
	o.TrackedSearchCount = v
}

// GetClickCount returns the ClickCount field value
func (o *TopSearchWithAnalytics) GetClickCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClickCount
}

// GetClickCountOk returns a tuple with the ClickCount field value
// and a boolean to check if the value has been set.
func (o *TopSearchWithAnalytics) GetClickCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClickCount, true
}

// SetClickCount sets field value
func (o *TopSearchWithAnalytics) SetClickCount(v int32) {
	o.ClickCount = v
}

// GetConversionCount returns the ConversionCount field value
func (o *TopSearchWithAnalytics) GetConversionCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ConversionCount
}

// GetConversionCountOk returns a tuple with the ConversionCount field value
// and a boolean to check if the value has been set.
func (o *TopSearchWithAnalytics) GetConversionCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConversionCount, true
}

// SetConversionCount sets field value
func (o *TopSearchWithAnalytics) SetConversionCount(v int32) {
	o.ConversionCount = v
}

// GetNbHits returns the NbHits field value
func (o *TopSearchWithAnalytics) GetNbHits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NbHits
}

// GetNbHitsOk returns a tuple with the NbHits field value
// and a boolean to check if the value has been set.
func (o *TopSearchWithAnalytics) GetNbHitsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NbHits, true
}

// SetNbHits sets field value
func (o *TopSearchWithAnalytics) SetNbHits(v int32) {
	o.NbHits = v
}

func (o TopSearchWithAnalytics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if true {
		toSerialize["search"] = o.Search
	}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["clickThroughRate"] = o.ClickThroughRate
	}
	if true {
		toSerialize["averageClickPosition"] = o.AverageClickPosition
	}
	if true {
		toSerialize["conversionRate"] = o.ConversionRate
	}
	if true {
		toSerialize["trackedSearchCount"] = o.TrackedSearchCount
	}
	if true {
		toSerialize["clickCount"] = o.ClickCount
	}
	if true {
		toSerialize["conversionCount"] = o.ConversionCount
	}
	if true {
		toSerialize["nbHits"] = o.NbHits
	}
	return json.Marshal(toSerialize)
}

func (o TopSearchWithAnalytics) String() string {
	out := ""
	out += fmt.Sprintf("  search=%v\n", o.Search)
	out += fmt.Sprintf("  count=%v\n", o.Count)
	out += fmt.Sprintf("  clickThroughRate=%v\n", o.ClickThroughRate)
	out += fmt.Sprintf("  averageClickPosition=%v\n", o.AverageClickPosition)
	out += fmt.Sprintf("  conversionRate=%v\n", o.ConversionRate)
	out += fmt.Sprintf("  trackedSearchCount=%v\n", o.TrackedSearchCount)
	out += fmt.Sprintf("  clickCount=%v\n", o.ClickCount)
	out += fmt.Sprintf("  conversionCount=%v\n", o.ConversionCount)
	out += fmt.Sprintf("  nbHits=%v\n", o.NbHits)
	return fmt.Sprintf("TopSearchWithAnalytics {\n%s}", out)
}

type NullableTopSearchWithAnalytics struct {
	value *TopSearchWithAnalytics
	isSet bool
}

func (v NullableTopSearchWithAnalytics) Get() *TopSearchWithAnalytics {
	return v.value
}

func (v *NullableTopSearchWithAnalytics) Set(val *TopSearchWithAnalytics) {
	v.value = val
	v.isSet = true
}

func (v NullableTopSearchWithAnalytics) IsSet() bool {
	return v.isSet
}

func (v *NullableTopSearchWithAnalytics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopSearchWithAnalytics(val *TopSearchWithAnalytics) *NullableTopSearchWithAnalytics {
	return &NullableTopSearchWithAnalytics{value: val, isSet: true}
}

func (v NullableTopSearchWithAnalytics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopSearchWithAnalytics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
