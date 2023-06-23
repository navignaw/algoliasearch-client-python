// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.
package insights

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/algolia/algoliasearch-client-go/v4/algolia/call"
)

type Option struct {
	optionType string
	name       string
	value      string
}

func QueryParamOption(name string, val any) Option {
	return Option{
		optionType: "query",
		name:       name,
		value:      parameterToString(val),
	}
}

func HeaderParamOption(name string, val any) Option {
	return Option{
		optionType: "header",
		name:       name,
		value:      parameterToString(val),
	}
}

func (r *ApiDelRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return err
	}
	if v, ok := req["path"]; ok {
		err = json.Unmarshal(v, &r.path)
		if err != nil {
			err = json.Unmarshal(b, &r.path)
			if err != nil {
				return err
			}
		}
	}
	if v, ok := req["parameters"]; ok {
		err = json.Unmarshal(v, &r.parameters)
		if err != nil {
			err = json.Unmarshal(b, &r.parameters)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// ApiDelRequest represents the request with all the parameters for the API call.
type ApiDelRequest struct {
	path       string
	parameters map[string]interface{}
}

// NewApiDelRequest creates an instance of the ApiDelRequest to be used for the API call.
func (c *APIClient) NewApiDelRequest(path string) ApiDelRequest {
	return ApiDelRequest{
		path: path,
	}
}

// WithParameters adds the parameters to the ApiDelRequest and returns the request for chaining.
func (r ApiDelRequest) WithParameters(parameters map[string]interface{}) ApiDelRequest {
	r.parameters = parameters
	return r
}

/*
Del Send requests to the Algolia REST API. Wraps DelWithContext using context.Background.

This method allow you to send requests to the Algolia REST API.

Request can be constructed by NewApiDelRequest with parameters below.

	@param path string - The path of the API endpoint to target, anything after the /1 needs to be specified.
	@param parameters map[string]interface{} - Query parameters to be applied to the current query.
	@return map[string]interface{}
*/
func (c *APIClient) Del(r ApiDelRequest, opts ...Option) (map[string]interface{}, error) {
	return c.DelWithContext(context.Background(), r, opts...)
}

/*
Del Send requests to the Algolia REST API.

This method allow you to send requests to the Algolia REST API.

Request can be constructed by NewApiDelRequest with parameters below.

	@param path string - The path of the API endpoint to target, anything after the /1 needs to be specified.
	@param parameters map[string]interface{} - Query parameters to be applied to the current query.
	@return map[string]interface{}
*/
func (c *APIClient) DelWithContext(ctx context.Context, r ApiDelRequest, opts ...Option) (map[string]interface{}, error) {
	var (
		postBody    any
		returnValue map[string]interface{}
	)

	requestPath := "/1{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(parameterToString(r.path)), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

	if !isNilorEmpty(r.parameters) {
		for k, v := range r.parameters {
			queryParams.Set(k, parameterToString(v))
		}
	}

	// optional params if any
	for _, opt := range opts {
		switch opt.optionType {
		case "query":
			queryParams.Set(opt.name, opt.value)
		case "header":
			headers[opt.name] = opt.value
		}
	}

	req, err := c.prepareRequest(ctx, requestPath, http.MethodDelete, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, err := c.callAPI(req, call.Write)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))
	if err != nil {
		return returnValue, err
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}
		if res.StatusCode == 400 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 402 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 403 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 404 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
		}
		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody, res.Header.Get("Content-Type"))
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

func (r *ApiGetRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return err
	}
	if v, ok := req["path"]; ok {
		err = json.Unmarshal(v, &r.path)
		if err != nil {
			err = json.Unmarshal(b, &r.path)
			if err != nil {
				return err
			}
		}
	}
	if v, ok := req["parameters"]; ok {
		err = json.Unmarshal(v, &r.parameters)
		if err != nil {
			err = json.Unmarshal(b, &r.parameters)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// ApiGetRequest represents the request with all the parameters for the API call.
type ApiGetRequest struct {
	path       string
	parameters map[string]interface{}
}

// NewApiGetRequest creates an instance of the ApiGetRequest to be used for the API call.
func (c *APIClient) NewApiGetRequest(path string) ApiGetRequest {
	return ApiGetRequest{
		path: path,
	}
}

// WithParameters adds the parameters to the ApiGetRequest and returns the request for chaining.
func (r ApiGetRequest) WithParameters(parameters map[string]interface{}) ApiGetRequest {
	r.parameters = parameters
	return r
}

/*
Get Send requests to the Algolia REST API. Wraps GetWithContext using context.Background.

This method allow you to send requests to the Algolia REST API.

Request can be constructed by NewApiGetRequest with parameters below.

	@param path string - The path of the API endpoint to target, anything after the /1 needs to be specified.
	@param parameters map[string]interface{} - Query parameters to be applied to the current query.
	@return map[string]interface{}
*/
func (c *APIClient) Get(r ApiGetRequest, opts ...Option) (map[string]interface{}, error) {
	return c.GetWithContext(context.Background(), r, opts...)
}

/*
Get Send requests to the Algolia REST API.

This method allow you to send requests to the Algolia REST API.

Request can be constructed by NewApiGetRequest with parameters below.

	@param path string - The path of the API endpoint to target, anything after the /1 needs to be specified.
	@param parameters map[string]interface{} - Query parameters to be applied to the current query.
	@return map[string]interface{}
*/
func (c *APIClient) GetWithContext(ctx context.Context, r ApiGetRequest, opts ...Option) (map[string]interface{}, error) {
	var (
		postBody    any
		returnValue map[string]interface{}
	)

	requestPath := "/1{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(parameterToString(r.path)), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

	if !isNilorEmpty(r.parameters) {
		for k, v := range r.parameters {
			queryParams.Set(k, parameterToString(v))
		}
	}

	// optional params if any
	for _, opt := range opts {
		switch opt.optionType {
		case "query":
			queryParams.Set(opt.name, opt.value)
		case "header":
			headers[opt.name] = opt.value
		}
	}

	req, err := c.prepareRequest(ctx, requestPath, http.MethodGet, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, err := c.callAPI(req, call.Write)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))
	if err != nil {
		return returnValue, err
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}
		if res.StatusCode == 400 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 402 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 403 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 404 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
		}
		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody, res.Header.Get("Content-Type"))
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

func (r *ApiPostRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return err
	}
	if v, ok := req["path"]; ok {
		err = json.Unmarshal(v, &r.path)
		if err != nil {
			err = json.Unmarshal(b, &r.path)
			if err != nil {
				return err
			}
		}
	}
	if v, ok := req["parameters"]; ok {
		err = json.Unmarshal(v, &r.parameters)
		if err != nil {
			err = json.Unmarshal(b, &r.parameters)
			if err != nil {
				return err
			}
		}
	}
	if v, ok := req["body"]; ok {
		err = json.Unmarshal(v, &r.body)
		if err != nil {
			err = json.Unmarshal(b, &r.body)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// ApiPostRequest represents the request with all the parameters for the API call.
type ApiPostRequest struct {
	path       string
	parameters map[string]interface{}
	body       map[string]interface{}
}

// NewApiPostRequest creates an instance of the ApiPostRequest to be used for the API call.
func (c *APIClient) NewApiPostRequest(path string) ApiPostRequest {
	return ApiPostRequest{
		path: path,
	}
}

// WithParameters adds the parameters to the ApiPostRequest and returns the request for chaining.
func (r ApiPostRequest) WithParameters(parameters map[string]interface{}) ApiPostRequest {
	r.parameters = parameters
	return r
}

// WithBody adds the body to the ApiPostRequest and returns the request for chaining.
func (r ApiPostRequest) WithBody(body map[string]interface{}) ApiPostRequest {
	r.body = body
	return r
}

/*
Post Send requests to the Algolia REST API. Wraps PostWithContext using context.Background.

This method allow you to send requests to the Algolia REST API.

Request can be constructed by NewApiPostRequest with parameters below.

	@param path string - The path of the API endpoint to target, anything after the /1 needs to be specified.
	@param parameters map[string]interface{} - Query parameters to be applied to the current query.
	@param body map[string]interface{} - The parameters to send with the custom request.
	@return map[string]interface{}
*/
func (c *APIClient) Post(r ApiPostRequest, opts ...Option) (map[string]interface{}, error) {
	return c.PostWithContext(context.Background(), r, opts...)
}

/*
Post Send requests to the Algolia REST API.

This method allow you to send requests to the Algolia REST API.

Request can be constructed by NewApiPostRequest with parameters below.

	@param path string - The path of the API endpoint to target, anything after the /1 needs to be specified.
	@param parameters map[string]interface{} - Query parameters to be applied to the current query.
	@param body map[string]interface{} - The parameters to send with the custom request.
	@return map[string]interface{}
*/
func (c *APIClient) PostWithContext(ctx context.Context, r ApiPostRequest, opts ...Option) (map[string]interface{}, error) {
	var (
		postBody    any
		returnValue map[string]interface{}
	)

	requestPath := "/1{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(parameterToString(r.path)), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

	if !isNilorEmpty(r.parameters) {
		for k, v := range r.parameters {
			queryParams.Set(k, parameterToString(v))
		}
	}

	// optional params if any
	for _, opt := range opts {
		switch opt.optionType {
		case "query":
			queryParams.Set(opt.name, opt.value)
		case "header":
			headers[opt.name] = opt.value
		}
	}

	// body params
	if isNilorEmpty(r.body) {
		postBody = "{}"
	} else {
		postBody = r.body
	}
	req, err := c.prepareRequest(ctx, requestPath, http.MethodPost, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, err := c.callAPI(req, call.Write)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))
	if err != nil {
		return returnValue, err
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}
		if res.StatusCode == 400 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 402 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 403 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 404 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
		}
		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody, res.Header.Get("Content-Type"))
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

func (r *ApiPushEventsRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return err
	}
	if v, ok := req["insightsEvents"]; ok {
		err = json.Unmarshal(v, &r.insightsEvents)
		if err != nil {
			err = json.Unmarshal(b, &r.insightsEvents)
			if err != nil {
				return err
			}
		}
	} else {
		err = json.Unmarshal(b, &r.insightsEvents)
		if err != nil {
			return err
		}
	}

	return nil
}

// ApiPushEventsRequest represents the request with all the parameters for the API call.
type ApiPushEventsRequest struct {
	insightsEvents *InsightsEvents
}

// NewApiPushEventsRequest creates an instance of the ApiPushEventsRequest to be used for the API call.
func (c *APIClient) NewApiPushEventsRequest(insightsEvents *InsightsEvents) ApiPushEventsRequest {
	return ApiPushEventsRequest{
		insightsEvents: insightsEvents,
	}
}

/*
PushEvents Send events. Wraps PushEventsWithContext using context.Background.

Send a list of events to the Insights API.

You can include up to 1,000 events in a single request,
but the request body must be smaller than 2&nbsp;MB.

Request can be constructed by NewApiPushEventsRequest with parameters below.

	@param insightsEvents InsightsEvents
	@return EventsResponse
*/
func (c *APIClient) PushEvents(r ApiPushEventsRequest, opts ...Option) (*EventsResponse, error) {
	return c.PushEventsWithContext(context.Background(), r, opts...)
}

/*
PushEvents Send events.

Send a list of events to the Insights API.

You can include up to 1,000 events in a single request,
but the request body must be smaller than 2&nbsp;MB.

Request can be constructed by NewApiPushEventsRequest with parameters below.

	@param insightsEvents InsightsEvents
	@return EventsResponse
*/
func (c *APIClient) PushEventsWithContext(ctx context.Context, r ApiPushEventsRequest, opts ...Option) (*EventsResponse, error) {
	var (
		postBody    any
		returnValue *EventsResponse
	)

	requestPath := "/1/events"

	headers := make(map[string]string)
	queryParams := url.Values{}
	if r.insightsEvents == nil {
		return returnValue, reportError("insightsEvents is required and must be specified")
	}

	// optional params if any
	for _, opt := range opts {
		switch opt.optionType {
		case "query":
			queryParams.Set(opt.name, opt.value)
		case "header":
			headers[opt.name] = opt.value
		}
	}

	// body params
	postBody = r.insightsEvents
	req, err := c.prepareRequest(ctx, requestPath, http.MethodPost, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, err := c.callAPI(req, call.Write)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))
	if err != nil {
		return returnValue, err
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}
		if res.StatusCode == 400 {
			var v string
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 401 {
			var v EventsResponse
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 404 {
			var v EventsResponse
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 405 {
			var v EventsResponse
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 413 {
			var v EventsResponse
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 422 {
			var v EventsResponse
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
		}
		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody, res.Header.Get("Content-Type"))
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}

func (r *ApiPutRequest) UnmarshalJSON(b []byte) error {
	req := map[string]json.RawMessage{}
	err := json.Unmarshal(b, &req)
	if err != nil {
		return err
	}
	if v, ok := req["path"]; ok {
		err = json.Unmarshal(v, &r.path)
		if err != nil {
			err = json.Unmarshal(b, &r.path)
			if err != nil {
				return err
			}
		}
	}
	if v, ok := req["parameters"]; ok {
		err = json.Unmarshal(v, &r.parameters)
		if err != nil {
			err = json.Unmarshal(b, &r.parameters)
			if err != nil {
				return err
			}
		}
	}
	if v, ok := req["body"]; ok {
		err = json.Unmarshal(v, &r.body)
		if err != nil {
			err = json.Unmarshal(b, &r.body)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// ApiPutRequest represents the request with all the parameters for the API call.
type ApiPutRequest struct {
	path       string
	parameters map[string]interface{}
	body       map[string]interface{}
}

// NewApiPutRequest creates an instance of the ApiPutRequest to be used for the API call.
func (c *APIClient) NewApiPutRequest(path string) ApiPutRequest {
	return ApiPutRequest{
		path: path,
	}
}

// WithParameters adds the parameters to the ApiPutRequest and returns the request for chaining.
func (r ApiPutRequest) WithParameters(parameters map[string]interface{}) ApiPutRequest {
	r.parameters = parameters
	return r
}

// WithBody adds the body to the ApiPutRequest and returns the request for chaining.
func (r ApiPutRequest) WithBody(body map[string]interface{}) ApiPutRequest {
	r.body = body
	return r
}

/*
Put Send requests to the Algolia REST API. Wraps PutWithContext using context.Background.

This method allow you to send requests to the Algolia REST API.

Request can be constructed by NewApiPutRequest with parameters below.

	@param path string - The path of the API endpoint to target, anything after the /1 needs to be specified.
	@param parameters map[string]interface{} - Query parameters to be applied to the current query.
	@param body map[string]interface{} - The parameters to send with the custom request.
	@return map[string]interface{}
*/
func (c *APIClient) Put(r ApiPutRequest, opts ...Option) (map[string]interface{}, error) {
	return c.PutWithContext(context.Background(), r, opts...)
}

/*
Put Send requests to the Algolia REST API.

This method allow you to send requests to the Algolia REST API.

Request can be constructed by NewApiPutRequest with parameters below.

	@param path string - The path of the API endpoint to target, anything after the /1 needs to be specified.
	@param parameters map[string]interface{} - Query parameters to be applied to the current query.
	@param body map[string]interface{} - The parameters to send with the custom request.
	@return map[string]interface{}
*/
func (c *APIClient) PutWithContext(ctx context.Context, r ApiPutRequest, opts ...Option) (map[string]interface{}, error) {
	var (
		postBody    any
		returnValue map[string]interface{}
	)

	requestPath := "/1{path}"
	requestPath = strings.Replace(requestPath, "{"+"path"+"}", url.PathEscape(parameterToString(r.path)), -1)

	headers := make(map[string]string)
	queryParams := url.Values{}

	if !isNilorEmpty(r.parameters) {
		for k, v := range r.parameters {
			queryParams.Set(k, parameterToString(v))
		}
	}

	// optional params if any
	for _, opt := range opts {
		switch opt.optionType {
		case "query":
			queryParams.Set(opt.name, opt.value)
		case "header":
			headers[opt.name] = opt.value
		}
	}

	// body params
	if isNilorEmpty(r.body) {
		postBody = "{}"
	} else {
		postBody = r.body
	}
	req, err := c.prepareRequest(ctx, requestPath, http.MethodPut, postBody, headers, queryParams)
	if err != nil {
		return returnValue, err
	}

	res, err := c.callAPI(req, call.Write)
	if err != nil {
		return returnValue, err
	}
	if res == nil {
		return returnValue, reportError("res is nil")
	}

	resBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	res.Body = io.NopCloser(bytes.NewBuffer(resBody))
	if err != nil {
		return returnValue, err
	}

	if res.StatusCode >= 300 {
		newErr := &APIError{
			Message: string(resBody),
			Status:  res.StatusCode,
		}
		if res.StatusCode == 400 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 402 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 403 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
			return returnValue, newErr
		}
		if res.StatusCode == 404 {
			var v ErrorBase
			err = c.decode(&v, resBody, res.Header.Get("Content-Type"))
			if err != nil {
				newErr.Message = err.Error()
				return returnValue, newErr
			}
		}
		return returnValue, newErr
	}

	err = c.decode(&returnValue, resBody, res.Header.Get("Content-Type"))
	if err != nil {
		return returnValue, reportError("cannot decode result: %w", err)
	}

	return returnValue, nil
}
