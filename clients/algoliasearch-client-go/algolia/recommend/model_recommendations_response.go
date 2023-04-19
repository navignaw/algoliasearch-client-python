// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package recommend

import (
	"encoding/json"
	"fmt"
)

// RecommendationsResponse struct for RecommendationsResponse
type RecommendationsResponse struct {
	// If a search encounters an index that is being A/B tested, abTestID reports the ongoing A/B test ID.
	AbTestID *int32 `json:"abTestID,omitempty"`
	// If a search encounters an index that is being A/B tested, abTestVariantID reports the variant ID of the index used (starting at 1).
	AbTestVariantID *int32 `json:"abTestVariantID,omitempty"`
	// The computed geo location.
	AroundLatLng *string `json:"aroundLatLng,omitempty"`
	// The automatically computed radius. For legacy reasons, this parameter is a string and not an integer.
	AutomaticRadius *string `json:"automaticRadius,omitempty"`
	// Whether the facet count is exhaustive or approximate.
	ExhaustiveFacetsCount *bool `json:"exhaustiveFacetsCount,omitempty"`
	// Indicate if the nbHits count was exhaustive or approximate.
	ExhaustiveNbHits bool `json:"exhaustiveNbHits"`
	// Indicate if the typo-tolerance search was exhaustive or approximate (only included when typo-tolerance is enabled).
	ExhaustiveTypo *bool `json:"exhaustiveTypo,omitempty"`
	// A mapping of each facet name to the corresponding facet counts.
	Facets *map[string]map[string]int32 `json:"facets,omitempty"`
	// Statistics for numerical facets.
	FacetsStats *map[string]FacetsStats `json:"facets_stats,omitempty"`
	// Set the number of hits per page.
	HitsPerPage int32 `json:"hitsPerPage"`
	// Index name used for the query.
	Index *string `json:"index,omitempty"`
	// Index name used for the query. In the case of an A/B test, the targeted index isn't always the index used by the query.
	IndexUsed *string `json:"indexUsed,omitempty"`
	// Used to return warnings about the query.
	Message *string `json:"message,omitempty"`
	// Number of hits that the search query matched.
	NbHits int32 `json:"nbHits"`
	// Number of pages available for the current query.
	NbPages int32 `json:"nbPages"`
	// The number of hits selected and sorted by the relevant sort algorithm.
	NbSortedHits *int32 `json:"nbSortedHits,omitempty"`
	// Specify the page to retrieve.
	Page int32 `json:"page"`
	// A url-encoded string of all search parameters.
	Params   string                      `json:"params"`
	Redirect *BaseSearchResponseRedirect `json:"redirect,omitempty"`
	// The query string that will be searched, after normalization.
	ParsedQuery *string `json:"parsedQuery,omitempty"`
	// Time the server took to process the request, in milliseconds.
	ProcessingTimeMS int32 `json:"processingTimeMS"`
	// The text to search in the index.
	Query string `json:"query"`
	// A markup text indicating which parts of the original query have been removed in order to retrieve a non-empty result set.
	QueryAfterRemoval *string `json:"queryAfterRemoval,omitempty"`
	// Actual host name of the server that processed the request.
	ServerUsed *string `json:"serverUsed,omitempty"`
	// Lets you store custom data in your indices.
	UserData         map[string]interface{} `json:"userData,omitempty"`
	RenderingContent *RenderingContent      `json:"renderingContent,omitempty"`
	Hits             []RecommendHit         `json:"hits"`
}

type RecommendationsResponseOption func(f *RecommendationsResponse)

func WithRecommendationsResponseAbTestID(val int32) RecommendationsResponseOption {
	return func(f *RecommendationsResponse) {
		f.AbTestID = &val
	}
}

func WithRecommendationsResponseAbTestVariantID(val int32) RecommendationsResponseOption {
	return func(f *RecommendationsResponse) {
		f.AbTestVariantID = &val
	}
}

func WithRecommendationsResponseAroundLatLng(val string) RecommendationsResponseOption {
	return func(f *RecommendationsResponse) {
		f.AroundLatLng = &val
	}
}

func WithRecommendationsResponseAutomaticRadius(val string) RecommendationsResponseOption {
	return func(f *RecommendationsResponse) {
		f.AutomaticRadius = &val
	}
}

func WithRecommendationsResponseExhaustiveFacetsCount(val bool) RecommendationsResponseOption {
	return func(f *RecommendationsResponse) {
		f.ExhaustiveFacetsCount = &val
	}
}

func WithRecommendationsResponseExhaustiveTypo(val bool) RecommendationsResponseOption {
	return func(f *RecommendationsResponse) {
		f.ExhaustiveTypo = &val
	}
}

func WithRecommendationsResponseFacets(val map[string]map[string]int32) RecommendationsResponseOption {
	return func(f *RecommendationsResponse) {
		f.Facets = &val
	}
}

func WithRecommendationsResponseFacetsStats(val map[string]FacetsStats) RecommendationsResponseOption {
	return func(f *RecommendationsResponse) {
		f.FacetsStats = &val
	}
}

func WithRecommendationsResponseIndex(val string) RecommendationsResponseOption {
	return func(f *RecommendationsResponse) {
		f.Index = &val
	}
}

func WithRecommendationsResponseIndexUsed(val string) RecommendationsResponseOption {
	return func(f *RecommendationsResponse) {
		f.IndexUsed = &val
	}
}

func WithRecommendationsResponseMessage(val string) RecommendationsResponseOption {
	return func(f *RecommendationsResponse) {
		f.Message = &val
	}
}

func WithRecommendationsResponseNbSortedHits(val int32) RecommendationsResponseOption {
	return func(f *RecommendationsResponse) {
		f.NbSortedHits = &val
	}
}

func WithRecommendationsResponseRedirect(val BaseSearchResponseRedirect) RecommendationsResponseOption {
	return func(f *RecommendationsResponse) {
		f.Redirect = &val
	}
}

func WithRecommendationsResponseParsedQuery(val string) RecommendationsResponseOption {
	return func(f *RecommendationsResponse) {
		f.ParsedQuery = &val
	}
}

func WithRecommendationsResponseQueryAfterRemoval(val string) RecommendationsResponseOption {
	return func(f *RecommendationsResponse) {
		f.QueryAfterRemoval = &val
	}
}

func WithRecommendationsResponseServerUsed(val string) RecommendationsResponseOption {
	return func(f *RecommendationsResponse) {
		f.ServerUsed = &val
	}
}

func WithRecommendationsResponseUserData(val map[string]interface{}) RecommendationsResponseOption {
	return func(f *RecommendationsResponse) {
		f.UserData = val
	}
}

func WithRecommendationsResponseRenderingContent(val RenderingContent) RecommendationsResponseOption {
	return func(f *RecommendationsResponse) {
		f.RenderingContent = &val
	}
}

// NewRecommendationsResponse instantiates a new RecommendationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecommendationsResponse(exhaustiveNbHits bool, hitsPerPage int32, nbHits int32, nbPages int32, page int32, params string, processingTimeMS int32, query string, hits []RecommendHit, opts ...RecommendationsResponseOption) *RecommendationsResponse {
	this := &RecommendationsResponse{}
	this.ExhaustiveNbHits = exhaustiveNbHits
	this.HitsPerPage = hitsPerPage
	this.NbHits = nbHits
	this.NbPages = nbPages
	this.Page = page
	this.Params = params
	this.ProcessingTimeMS = processingTimeMS
	this.Query = query
	this.Hits = hits
	for _, opt := range opts {
		opt(this)
	}
	return this
}

// NewRecommendationsResponseWithDefaults instantiates a new RecommendationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecommendationsResponseWithDefaults() *RecommendationsResponse {
	this := &RecommendationsResponse{}
	var hitsPerPage int32 = 20
	this.HitsPerPage = hitsPerPage
	var page int32 = 0
	this.Page = page
	var query string = ""
	this.Query = query
	return this
}

// GetAbTestID returns the AbTestID field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetAbTestID() int32 {
	if o == nil || o.AbTestID == nil {
		var ret int32
		return ret
	}
	return *o.AbTestID
}

// GetAbTestIDOk returns a tuple with the AbTestID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetAbTestIDOk() (*int32, bool) {
	if o == nil || o.AbTestID == nil {
		return nil, false
	}
	return o.AbTestID, true
}

// HasAbTestID returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasAbTestID() bool {
	if o != nil && o.AbTestID != nil {
		return true
	}

	return false
}

// SetAbTestID gets a reference to the given int32 and assigns it to the AbTestID field.
func (o *RecommendationsResponse) SetAbTestID(v int32) {
	o.AbTestID = &v
}

// GetAbTestVariantID returns the AbTestVariantID field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetAbTestVariantID() int32 {
	if o == nil || o.AbTestVariantID == nil {
		var ret int32
		return ret
	}
	return *o.AbTestVariantID
}

// GetAbTestVariantIDOk returns a tuple with the AbTestVariantID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetAbTestVariantIDOk() (*int32, bool) {
	if o == nil || o.AbTestVariantID == nil {
		return nil, false
	}
	return o.AbTestVariantID, true
}

// HasAbTestVariantID returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasAbTestVariantID() bool {
	if o != nil && o.AbTestVariantID != nil {
		return true
	}

	return false
}

// SetAbTestVariantID gets a reference to the given int32 and assigns it to the AbTestVariantID field.
func (o *RecommendationsResponse) SetAbTestVariantID(v int32) {
	o.AbTestVariantID = &v
}

// GetAroundLatLng returns the AroundLatLng field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetAroundLatLng() string {
	if o == nil || o.AroundLatLng == nil {
		var ret string
		return ret
	}
	return *o.AroundLatLng
}

// GetAroundLatLngOk returns a tuple with the AroundLatLng field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetAroundLatLngOk() (*string, bool) {
	if o == nil || o.AroundLatLng == nil {
		return nil, false
	}
	return o.AroundLatLng, true
}

// HasAroundLatLng returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasAroundLatLng() bool {
	if o != nil && o.AroundLatLng != nil {
		return true
	}

	return false
}

// SetAroundLatLng gets a reference to the given string and assigns it to the AroundLatLng field.
func (o *RecommendationsResponse) SetAroundLatLng(v string) {
	o.AroundLatLng = &v
}

// GetAutomaticRadius returns the AutomaticRadius field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetAutomaticRadius() string {
	if o == nil || o.AutomaticRadius == nil {
		var ret string
		return ret
	}
	return *o.AutomaticRadius
}

// GetAutomaticRadiusOk returns a tuple with the AutomaticRadius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetAutomaticRadiusOk() (*string, bool) {
	if o == nil || o.AutomaticRadius == nil {
		return nil, false
	}
	return o.AutomaticRadius, true
}

// HasAutomaticRadius returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasAutomaticRadius() bool {
	if o != nil && o.AutomaticRadius != nil {
		return true
	}

	return false
}

// SetAutomaticRadius gets a reference to the given string and assigns it to the AutomaticRadius field.
func (o *RecommendationsResponse) SetAutomaticRadius(v string) {
	o.AutomaticRadius = &v
}

// GetExhaustiveFacetsCount returns the ExhaustiveFacetsCount field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetExhaustiveFacetsCount() bool {
	if o == nil || o.ExhaustiveFacetsCount == nil {
		var ret bool
		return ret
	}
	return *o.ExhaustiveFacetsCount
}

// GetExhaustiveFacetsCountOk returns a tuple with the ExhaustiveFacetsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetExhaustiveFacetsCountOk() (*bool, bool) {
	if o == nil || o.ExhaustiveFacetsCount == nil {
		return nil, false
	}
	return o.ExhaustiveFacetsCount, true
}

// HasExhaustiveFacetsCount returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasExhaustiveFacetsCount() bool {
	if o != nil && o.ExhaustiveFacetsCount != nil {
		return true
	}

	return false
}

// SetExhaustiveFacetsCount gets a reference to the given bool and assigns it to the ExhaustiveFacetsCount field.
func (o *RecommendationsResponse) SetExhaustiveFacetsCount(v bool) {
	o.ExhaustiveFacetsCount = &v
}

// GetExhaustiveNbHits returns the ExhaustiveNbHits field value
func (o *RecommendationsResponse) GetExhaustiveNbHits() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ExhaustiveNbHits
}

// GetExhaustiveNbHitsOk returns a tuple with the ExhaustiveNbHits field value
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetExhaustiveNbHitsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExhaustiveNbHits, true
}

// SetExhaustiveNbHits sets field value
func (o *RecommendationsResponse) SetExhaustiveNbHits(v bool) {
	o.ExhaustiveNbHits = v
}

// GetExhaustiveTypo returns the ExhaustiveTypo field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetExhaustiveTypo() bool {
	if o == nil || o.ExhaustiveTypo == nil {
		var ret bool
		return ret
	}
	return *o.ExhaustiveTypo
}

// GetExhaustiveTypoOk returns a tuple with the ExhaustiveTypo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetExhaustiveTypoOk() (*bool, bool) {
	if o == nil || o.ExhaustiveTypo == nil {
		return nil, false
	}
	return o.ExhaustiveTypo, true
}

// HasExhaustiveTypo returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasExhaustiveTypo() bool {
	if o != nil && o.ExhaustiveTypo != nil {
		return true
	}

	return false
}

// SetExhaustiveTypo gets a reference to the given bool and assigns it to the ExhaustiveTypo field.
func (o *RecommendationsResponse) SetExhaustiveTypo(v bool) {
	o.ExhaustiveTypo = &v
}

// GetFacets returns the Facets field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetFacets() map[string]map[string]int32 {
	if o == nil || o.Facets == nil {
		var ret map[string]map[string]int32
		return ret
	}
	return *o.Facets
}

// GetFacetsOk returns a tuple with the Facets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetFacetsOk() (*map[string]map[string]int32, bool) {
	if o == nil || o.Facets == nil {
		return nil, false
	}
	return o.Facets, true
}

// HasFacets returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasFacets() bool {
	if o != nil && o.Facets != nil {
		return true
	}

	return false
}

// SetFacets gets a reference to the given map[string]map[string]int32 and assigns it to the Facets field.
func (o *RecommendationsResponse) SetFacets(v map[string]map[string]int32) {
	o.Facets = &v
}

// GetFacetsStats returns the FacetsStats field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetFacetsStats() map[string]FacetsStats {
	if o == nil || o.FacetsStats == nil {
		var ret map[string]FacetsStats
		return ret
	}
	return *o.FacetsStats
}

// GetFacetsStatsOk returns a tuple with the FacetsStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetFacetsStatsOk() (*map[string]FacetsStats, bool) {
	if o == nil || o.FacetsStats == nil {
		return nil, false
	}
	return o.FacetsStats, true
}

// HasFacetsStats returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasFacetsStats() bool {
	if o != nil && o.FacetsStats != nil {
		return true
	}

	return false
}

// SetFacetsStats gets a reference to the given map[string]FacetsStats and assigns it to the FacetsStats field.
func (o *RecommendationsResponse) SetFacetsStats(v map[string]FacetsStats) {
	o.FacetsStats = &v
}

// GetHitsPerPage returns the HitsPerPage field value
func (o *RecommendationsResponse) GetHitsPerPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.HitsPerPage
}

// GetHitsPerPageOk returns a tuple with the HitsPerPage field value
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetHitsPerPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HitsPerPage, true
}

// SetHitsPerPage sets field value
func (o *RecommendationsResponse) SetHitsPerPage(v int32) {
	o.HitsPerPage = v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetIndex() string {
	if o == nil || o.Index == nil {
		var ret string
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetIndexOk() (*string, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given string and assigns it to the Index field.
func (o *RecommendationsResponse) SetIndex(v string) {
	o.Index = &v
}

// GetIndexUsed returns the IndexUsed field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetIndexUsed() string {
	if o == nil || o.IndexUsed == nil {
		var ret string
		return ret
	}
	return *o.IndexUsed
}

// GetIndexUsedOk returns a tuple with the IndexUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetIndexUsedOk() (*string, bool) {
	if o == nil || o.IndexUsed == nil {
		return nil, false
	}
	return o.IndexUsed, true
}

// HasIndexUsed returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasIndexUsed() bool {
	if o != nil && o.IndexUsed != nil {
		return true
	}

	return false
}

// SetIndexUsed gets a reference to the given string and assigns it to the IndexUsed field.
func (o *RecommendationsResponse) SetIndexUsed(v string) {
	o.IndexUsed = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *RecommendationsResponse) SetMessage(v string) {
	o.Message = &v
}

// GetNbHits returns the NbHits field value
func (o *RecommendationsResponse) GetNbHits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NbHits
}

// GetNbHitsOk returns a tuple with the NbHits field value
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetNbHitsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NbHits, true
}

// SetNbHits sets field value
func (o *RecommendationsResponse) SetNbHits(v int32) {
	o.NbHits = v
}

// GetNbPages returns the NbPages field value
func (o *RecommendationsResponse) GetNbPages() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NbPages
}

// GetNbPagesOk returns a tuple with the NbPages field value
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetNbPagesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NbPages, true
}

// SetNbPages sets field value
func (o *RecommendationsResponse) SetNbPages(v int32) {
	o.NbPages = v
}

// GetNbSortedHits returns the NbSortedHits field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetNbSortedHits() int32 {
	if o == nil || o.NbSortedHits == nil {
		var ret int32
		return ret
	}
	return *o.NbSortedHits
}

// GetNbSortedHitsOk returns a tuple with the NbSortedHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetNbSortedHitsOk() (*int32, bool) {
	if o == nil || o.NbSortedHits == nil {
		return nil, false
	}
	return o.NbSortedHits, true
}

// HasNbSortedHits returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasNbSortedHits() bool {
	if o != nil && o.NbSortedHits != nil {
		return true
	}

	return false
}

// SetNbSortedHits gets a reference to the given int32 and assigns it to the NbSortedHits field.
func (o *RecommendationsResponse) SetNbSortedHits(v int32) {
	o.NbSortedHits = &v
}

// GetPage returns the Page field value
func (o *RecommendationsResponse) GetPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *RecommendationsResponse) SetPage(v int32) {
	o.Page = v
}

// GetParams returns the Params field value
func (o *RecommendationsResponse) GetParams() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetParamsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *RecommendationsResponse) SetParams(v string) {
	o.Params = v
}

// GetRedirect returns the Redirect field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetRedirect() BaseSearchResponseRedirect {
	if o == nil || o.Redirect == nil {
		var ret BaseSearchResponseRedirect
		return ret
	}
	return *o.Redirect
}

// GetRedirectOk returns a tuple with the Redirect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetRedirectOk() (*BaseSearchResponseRedirect, bool) {
	if o == nil || o.Redirect == nil {
		return nil, false
	}
	return o.Redirect, true
}

// HasRedirect returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasRedirect() bool {
	if o != nil && o.Redirect != nil {
		return true
	}

	return false
}

// SetRedirect gets a reference to the given BaseSearchResponseRedirect and assigns it to the Redirect field.
func (o *RecommendationsResponse) SetRedirect(v BaseSearchResponseRedirect) {
	o.Redirect = &v
}

// GetParsedQuery returns the ParsedQuery field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetParsedQuery() string {
	if o == nil || o.ParsedQuery == nil {
		var ret string
		return ret
	}
	return *o.ParsedQuery
}

// GetParsedQueryOk returns a tuple with the ParsedQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetParsedQueryOk() (*string, bool) {
	if o == nil || o.ParsedQuery == nil {
		return nil, false
	}
	return o.ParsedQuery, true
}

// HasParsedQuery returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasParsedQuery() bool {
	if o != nil && o.ParsedQuery != nil {
		return true
	}

	return false
}

// SetParsedQuery gets a reference to the given string and assigns it to the ParsedQuery field.
func (o *RecommendationsResponse) SetParsedQuery(v string) {
	o.ParsedQuery = &v
}

// GetProcessingTimeMS returns the ProcessingTimeMS field value
func (o *RecommendationsResponse) GetProcessingTimeMS() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProcessingTimeMS
}

// GetProcessingTimeMSOk returns a tuple with the ProcessingTimeMS field value
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetProcessingTimeMSOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProcessingTimeMS, true
}

// SetProcessingTimeMS sets field value
func (o *RecommendationsResponse) SetProcessingTimeMS(v int32) {
	o.ProcessingTimeMS = v
}

// GetQuery returns the Query field value
func (o *RecommendationsResponse) GetQuery() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Query
}

// GetQueryOk returns a tuple with the Query field value
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetQueryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Query, true
}

// SetQuery sets field value
func (o *RecommendationsResponse) SetQuery(v string) {
	o.Query = v
}

// GetQueryAfterRemoval returns the QueryAfterRemoval field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetQueryAfterRemoval() string {
	if o == nil || o.QueryAfterRemoval == nil {
		var ret string
		return ret
	}
	return *o.QueryAfterRemoval
}

// GetQueryAfterRemovalOk returns a tuple with the QueryAfterRemoval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetQueryAfterRemovalOk() (*string, bool) {
	if o == nil || o.QueryAfterRemoval == nil {
		return nil, false
	}
	return o.QueryAfterRemoval, true
}

// HasQueryAfterRemoval returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasQueryAfterRemoval() bool {
	if o != nil && o.QueryAfterRemoval != nil {
		return true
	}

	return false
}

// SetQueryAfterRemoval gets a reference to the given string and assigns it to the QueryAfterRemoval field.
func (o *RecommendationsResponse) SetQueryAfterRemoval(v string) {
	o.QueryAfterRemoval = &v
}

// GetServerUsed returns the ServerUsed field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetServerUsed() string {
	if o == nil || o.ServerUsed == nil {
		var ret string
		return ret
	}
	return *o.ServerUsed
}

// GetServerUsedOk returns a tuple with the ServerUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetServerUsedOk() (*string, bool) {
	if o == nil || o.ServerUsed == nil {
		return nil, false
	}
	return o.ServerUsed, true
}

// HasServerUsed returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasServerUsed() bool {
	if o != nil && o.ServerUsed != nil {
		return true
	}

	return false
}

// SetServerUsed gets a reference to the given string and assigns it to the ServerUsed field.
func (o *RecommendationsResponse) SetServerUsed(v string) {
	o.ServerUsed = &v
}

// GetUserData returns the UserData field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetUserData() map[string]interface{} {
	if o == nil || o.UserData == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.UserData
}

// GetUserDataOk returns a tuple with the UserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetUserDataOk() (map[string]interface{}, bool) {
	if o == nil || o.UserData == nil {
		return nil, false
	}
	return o.UserData, true
}

// HasUserData returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasUserData() bool {
	if o != nil && o.UserData != nil {
		return true
	}

	return false
}

// SetUserData gets a reference to the given map[string]interface{} and assigns it to the UserData field.
func (o *RecommendationsResponse) SetUserData(v map[string]interface{}) {
	o.UserData = v
}

// GetRenderingContent returns the RenderingContent field value if set, zero value otherwise.
func (o *RecommendationsResponse) GetRenderingContent() RenderingContent {
	if o == nil || o.RenderingContent == nil {
		var ret RenderingContent
		return ret
	}
	return *o.RenderingContent
}

// GetRenderingContentOk returns a tuple with the RenderingContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetRenderingContentOk() (*RenderingContent, bool) {
	if o == nil || o.RenderingContent == nil {
		return nil, false
	}
	return o.RenderingContent, true
}

// HasRenderingContent returns a boolean if a field has been set.
func (o *RecommendationsResponse) HasRenderingContent() bool {
	if o != nil && o.RenderingContent != nil {
		return true
	}

	return false
}

// SetRenderingContent gets a reference to the given RenderingContent and assigns it to the RenderingContent field.
func (o *RecommendationsResponse) SetRenderingContent(v RenderingContent) {
	o.RenderingContent = &v
}

// GetHits returns the Hits field value
func (o *RecommendationsResponse) GetHits() []RecommendHit {
	if o == nil {
		var ret []RecommendHit
		return ret
	}

	return o.Hits
}

// GetHitsOk returns a tuple with the Hits field value
// and a boolean to check if the value has been set.
func (o *RecommendationsResponse) GetHitsOk() ([]RecommendHit, bool) {
	if o == nil {
		return nil, false
	}
	return o.Hits, true
}

// SetHits sets field value
func (o *RecommendationsResponse) SetHits(v []RecommendHit) {
	o.Hits = v
}

func (o RecommendationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]any{}
	if o.AbTestID != nil {
		toSerialize["abTestID"] = o.AbTestID
	}
	if o.AbTestVariantID != nil {
		toSerialize["abTestVariantID"] = o.AbTestVariantID
	}
	if o.AroundLatLng != nil {
		toSerialize["aroundLatLng"] = o.AroundLatLng
	}
	if o.AutomaticRadius != nil {
		toSerialize["automaticRadius"] = o.AutomaticRadius
	}
	if o.ExhaustiveFacetsCount != nil {
		toSerialize["exhaustiveFacetsCount"] = o.ExhaustiveFacetsCount
	}
	if true {
		toSerialize["exhaustiveNbHits"] = o.ExhaustiveNbHits
	}
	if o.ExhaustiveTypo != nil {
		toSerialize["exhaustiveTypo"] = o.ExhaustiveTypo
	}
	if o.Facets != nil {
		toSerialize["facets"] = o.Facets
	}
	if o.FacetsStats != nil {
		toSerialize["facets_stats"] = o.FacetsStats
	}
	if true {
		toSerialize["hitsPerPage"] = o.HitsPerPage
	}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.IndexUsed != nil {
		toSerialize["indexUsed"] = o.IndexUsed
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["nbHits"] = o.NbHits
	}
	if true {
		toSerialize["nbPages"] = o.NbPages
	}
	if o.NbSortedHits != nil {
		toSerialize["nbSortedHits"] = o.NbSortedHits
	}
	if true {
		toSerialize["page"] = o.Page
	}
	if true {
		toSerialize["params"] = o.Params
	}
	if o.Redirect != nil {
		toSerialize["redirect"] = o.Redirect
	}
	if o.ParsedQuery != nil {
		toSerialize["parsedQuery"] = o.ParsedQuery
	}
	if true {
		toSerialize["processingTimeMS"] = o.ProcessingTimeMS
	}
	if true {
		toSerialize["query"] = o.Query
	}
	if o.QueryAfterRemoval != nil {
		toSerialize["queryAfterRemoval"] = o.QueryAfterRemoval
	}
	if o.ServerUsed != nil {
		toSerialize["serverUsed"] = o.ServerUsed
	}
	if o.UserData != nil {
		toSerialize["userData"] = o.UserData
	}
	if o.RenderingContent != nil {
		toSerialize["renderingContent"] = o.RenderingContent
	}
	if true {
		toSerialize["hits"] = o.Hits
	}
	return json.Marshal(toSerialize)
}

func (o RecommendationsResponse) String() string {
	out := ""
	out += fmt.Sprintf("  abTestID=%v\n", o.AbTestID)
	out += fmt.Sprintf("  abTestVariantID=%v\n", o.AbTestVariantID)
	out += fmt.Sprintf("  aroundLatLng=%v\n", o.AroundLatLng)
	out += fmt.Sprintf("  automaticRadius=%v\n", o.AutomaticRadius)
	out += fmt.Sprintf("  exhaustiveFacetsCount=%v\n", o.ExhaustiveFacetsCount)
	out += fmt.Sprintf("  exhaustiveNbHits=%v\n", o.ExhaustiveNbHits)
	out += fmt.Sprintf("  exhaustiveTypo=%v\n", o.ExhaustiveTypo)
	out += fmt.Sprintf("  facets=%v\n", o.Facets)
	out += fmt.Sprintf("  facets_stats=%v\n", o.FacetsStats)
	out += fmt.Sprintf("  hitsPerPage=%v\n", o.HitsPerPage)
	out += fmt.Sprintf("  index=%v\n", o.Index)
	out += fmt.Sprintf("  indexUsed=%v\n", o.IndexUsed)
	out += fmt.Sprintf("  message=%v\n", o.Message)
	out += fmt.Sprintf("  nbHits=%v\n", o.NbHits)
	out += fmt.Sprintf("  nbPages=%v\n", o.NbPages)
	out += fmt.Sprintf("  nbSortedHits=%v\n", o.NbSortedHits)
	out += fmt.Sprintf("  page=%v\n", o.Page)
	out += fmt.Sprintf("  params=%v\n", o.Params)
	out += fmt.Sprintf("  redirect=%v\n", o.Redirect)
	out += fmt.Sprintf("  parsedQuery=%v\n", o.ParsedQuery)
	out += fmt.Sprintf("  processingTimeMS=%v\n", o.ProcessingTimeMS)
	out += fmt.Sprintf("  query=%v\n", o.Query)
	out += fmt.Sprintf("  queryAfterRemoval=%v\n", o.QueryAfterRemoval)
	out += fmt.Sprintf("  serverUsed=%v\n", o.ServerUsed)
	out += fmt.Sprintf("  userData=%v\n", o.UserData)
	out += fmt.Sprintf("  renderingContent=%v\n", o.RenderingContent)
	out += fmt.Sprintf("  hits=%v\n", o.Hits)
	return fmt.Sprintf("RecommendationsResponse {\n%s}", out)
}

type NullableRecommendationsResponse struct {
	value *RecommendationsResponse
	isSet bool
}

func (v NullableRecommendationsResponse) Get() *RecommendationsResponse {
	return v.value
}

func (v *NullableRecommendationsResponse) Set(val *RecommendationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRecommendationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRecommendationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecommendationsResponse(val *RecommendationsResponse) *NullableRecommendationsResponse {
	return &NullableRecommendationsResponse{value: val, isSet: true}
}

func (v NullableRecommendationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecommendationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
