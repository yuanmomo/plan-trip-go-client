/*
Planera Resa

Sök och planera resor med Västtrafik

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// StopPointsApiService StopPointsApi service
type StopPointsApiService service

type ApiStopPointsStopPointGidArrivalsDetailsReferenceDetailsGetRequest struct {
	ctx context.Context
	ApiService *StopPointsApiService
	detailsReference string
	stopPointGid string
	includes *[]VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType
}

// The additional information to include in the response.
func (r ApiStopPointsStopPointGidArrivalsDetailsReferenceDetailsGetRequest) Includes(includes []VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType) ApiStopPointsStopPointGidArrivalsDetailsReferenceDetailsGetRequest {
	r.includes = &includes
	return r
}

func (r ApiStopPointsStopPointGidArrivalsDetailsReferenceDetailsGetRequest) Execute() (*VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalDetailsApiModel, *http.Response, error) {
	return r.ApiService.StopPointsStopPointGidArrivalsDetailsReferenceDetailsGetExecute(r)
}

/*
StopPointsStopPointGidArrivalsDetailsReferenceDetailsGet Returns details about an arrival.

Sample request:

    GET /stop-points/9022014001760003/arrivals/{detailsReference}/details?includes=servicejourneycalls

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param detailsReference The reference to the service journey, received from the arrivals call. A detailsReference is only valid during the same day as it was generated.
 @param stopPointGid
 @return ApiStopPointsStopPointGidArrivalsDetailsReferenceDetailsGetRequest
*/
func (a *StopPointsApiService) StopPointsStopPointGidArrivalsDetailsReferenceDetailsGet(ctx context.Context, detailsReference string, stopPointGid string) ApiStopPointsStopPointGidArrivalsDetailsReferenceDetailsGetRequest {
	return ApiStopPointsStopPointGidArrivalsDetailsReferenceDetailsGetRequest{
		ApiService: a,
		ctx: ctx,
		detailsReference: detailsReference,
		stopPointGid: stopPointGid,
	}
}

// Execute executes the request
//  @return VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalDetailsApiModel
func (a *StopPointsApiService) StopPointsStopPointGidArrivalsDetailsReferenceDetailsGetExecute(r ApiStopPointsStopPointGidArrivalsDetailsReferenceDetailsGetRequest) (*VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalDetailsApiModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalDetailsApiModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StopPointsApiService.StopPointsStopPointGidArrivalsDetailsReferenceDetailsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stop-points/{stopPointGid}/arrivals/{detailsReference}/details"
	localVarPath = strings.Replace(localVarPath, "{"+"detailsReference"+"}", url.PathEscape(parameterValueToString(r.detailsReference, "detailsReference")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"stopPointGid"+"}", url.PathEscape(parameterValueToString(r.stopPointGid, "stopPointGid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includes != nil {
		t := *r.includes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "includes", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "includes", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v VTApiPlaneraResaWebV4ModelsApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v MicrosoftAspNetCoreMvcProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiStopPointsStopPointGidArrivalsGetRequest struct {
	ctx context.Context
	ApiService *StopPointsApiService
	stopPointGid string
	startDateTime *string
	timeSpanInMinutes *int32
	maxArrivalsPerLineAndDirection *int32
	limit *int32
	offset *int32
	directionGid *string
}

// The start of the time interval for which to get upcoming arrivals. Must be specified in RFC 3339 format or be null which means that the current time on the server is used.
func (r ApiStopPointsStopPointGidArrivalsGetRequest) StartDateTime(startDateTime string) ApiStopPointsStopPointGidArrivalsGetRequest {
	r.startDateTime = &startDateTime
	return r
}

// The number of minutes from the start time for which to get upcoming arrivals. Allowed values are between 0 and 1440.
func (r ApiStopPointsStopPointGidArrivalsGetRequest) TimeSpanInMinutes(timeSpanInMinutes int32) ApiStopPointsStopPointGidArrivalsGetRequest {
	r.timeSpanInMinutes = &timeSpanInMinutes
	return r
}

// The maximum number of arrivals for a single line in a specific direction.
func (r ApiStopPointsStopPointGidArrivalsGetRequest) MaxArrivalsPerLineAndDirection(maxArrivalsPerLineAndDirection int32) ApiStopPointsStopPointGidArrivalsGetRequest {
	r.maxArrivalsPerLineAndDirection = &maxArrivalsPerLineAndDirection
	return r
}

// The number of results to return.
func (r ApiStopPointsStopPointGidArrivalsGetRequest) Limit(limit int32) ApiStopPointsStopPointGidArrivalsGetRequest {
	r.limit = &limit
	return r
}

// The zero-based start offset of the pagination.
func (r ApiStopPointsStopPointGidArrivalsGetRequest) Offset(offset int32) ApiStopPointsStopPointGidArrivalsGetRequest {
	r.offset = &offset
	return r
}

// Filter result by last stop on journey. Must be a 16-digit Västtrafik stop area
func (r ApiStopPointsStopPointGidArrivalsGetRequest) DirectionGid(directionGid string) ApiStopPointsStopPointGidArrivalsGetRequest {
	r.directionGid = &directionGid
	return r
}

func (r ApiStopPointsStopPointGidArrivalsGetRequest) Execute() (*VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsGetArrivalsResponse, *http.Response, error) {
	return r.ApiService.StopPointsStopPointGidArrivalsGetExecute(r)
}

/*
StopPointsStopPointGidArrivalsGet Returns arrivals to the specified stop point at the specified time.

Sample request:

    GET /stop-points/9022014001760003/arrivals

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stopPointGid The 16-digit Västtrafik gid of the stop point.
 @return ApiStopPointsStopPointGidArrivalsGetRequest
*/
func (a *StopPointsApiService) StopPointsStopPointGidArrivalsGet(ctx context.Context, stopPointGid string) ApiStopPointsStopPointGidArrivalsGetRequest {
	return ApiStopPointsStopPointGidArrivalsGetRequest{
		ApiService: a,
		ctx: ctx,
		stopPointGid: stopPointGid,
	}
}

// Execute executes the request
//  @return VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsGetArrivalsResponse
func (a *StopPointsApiService) StopPointsStopPointGidArrivalsGetExecute(r ApiStopPointsStopPointGidArrivalsGetRequest) (*VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsGetArrivalsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsGetArrivalsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StopPointsApiService.StopPointsStopPointGidArrivalsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stop-points/{stopPointGid}/arrivals"
	localVarPath = strings.Replace(localVarPath, "{"+"stopPointGid"+"}", url.PathEscape(parameterValueToString(r.stopPointGid, "stopPointGid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startDateTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startDateTime", r.startDateTime, "")
	}
	if r.timeSpanInMinutes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timeSpanInMinutes", r.timeSpanInMinutes, "")
	}
	if r.maxArrivalsPerLineAndDirection != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxArrivalsPerLineAndDirection", r.maxArrivalsPerLineAndDirection, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.directionGid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "directionGid", r.directionGid, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v VTApiPlaneraResaWebV4ModelsApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v MicrosoftAspNetCoreMvcProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiStopPointsStopPointGidDeparturesDetailsReferenceDetailsGetRequest struct {
	ctx context.Context
	ApiService *StopPointsApiService
	detailsReference string
	stopPointGid string
	includes *[]VTApiPlaneraResaWebV4ModelsDepartureDetailsIncludeType
}

// The additional information to include in the response.
func (r ApiStopPointsStopPointGidDeparturesDetailsReferenceDetailsGetRequest) Includes(includes []VTApiPlaneraResaWebV4ModelsDepartureDetailsIncludeType) ApiStopPointsStopPointGidDeparturesDetailsReferenceDetailsGetRequest {
	r.includes = &includes
	return r
}

func (r ApiStopPointsStopPointGidDeparturesDetailsReferenceDetailsGetRequest) Execute() (*VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureDetailsApiModel, *http.Response, error) {
	return r.ApiService.StopPointsStopPointGidDeparturesDetailsReferenceDetailsGetExecute(r)
}

/*
StopPointsStopPointGidDeparturesDetailsReferenceDetailsGet Returns details about a departure.

Sample request:

    GET /stop-points/9022014001760003/departures/{detailsReference}/details?includes=servicejourneycalls

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param detailsReference The reference to the service journey, received from the departures call. A detailsReference is only valid during the same day as it was generated.
 @param stopPointGid
 @return ApiStopPointsStopPointGidDeparturesDetailsReferenceDetailsGetRequest
*/
func (a *StopPointsApiService) StopPointsStopPointGidDeparturesDetailsReferenceDetailsGet(ctx context.Context, detailsReference string, stopPointGid string) ApiStopPointsStopPointGidDeparturesDetailsReferenceDetailsGetRequest {
	return ApiStopPointsStopPointGidDeparturesDetailsReferenceDetailsGetRequest{
		ApiService: a,
		ctx: ctx,
		detailsReference: detailsReference,
		stopPointGid: stopPointGid,
	}
}

// Execute executes the request
//  @return VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureDetailsApiModel
func (a *StopPointsApiService) StopPointsStopPointGidDeparturesDetailsReferenceDetailsGetExecute(r ApiStopPointsStopPointGidDeparturesDetailsReferenceDetailsGetRequest) (*VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureDetailsApiModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureDetailsApiModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StopPointsApiService.StopPointsStopPointGidDeparturesDetailsReferenceDetailsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stop-points/{stopPointGid}/departures/{detailsReference}/details"
	localVarPath = strings.Replace(localVarPath, "{"+"detailsReference"+"}", url.PathEscape(parameterValueToString(r.detailsReference, "detailsReference")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"stopPointGid"+"}", url.PathEscape(parameterValueToString(r.stopPointGid, "stopPointGid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includes != nil {
		t := *r.includes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "includes", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "includes", t, "multi")
		}
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v VTApiPlaneraResaWebV4ModelsApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v MicrosoftAspNetCoreMvcProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiStopPointsStopPointGidDeparturesGetRequest struct {
	ctx context.Context
	ApiService *StopPointsApiService
	stopPointGid string
	startDateTime *string
	timeSpanInMinutes *int32
	maxDeparturesPerLineAndDirection *int32
	limit *int32
	offset *int32
	includeOccupancy *bool
	directionGid *string
}

// The start of the time interval for which to get upcoming departures. Must be specified in RFC 3339 format or be null which means that the current time on the server is used.
func (r ApiStopPointsStopPointGidDeparturesGetRequest) StartDateTime(startDateTime string) ApiStopPointsStopPointGidDeparturesGetRequest {
	r.startDateTime = &startDateTime
	return r
}

// The number of minutes from the start time for which to get upcoming departures. Allowed values are between 0 and 1440.
func (r ApiStopPointsStopPointGidDeparturesGetRequest) TimeSpanInMinutes(timeSpanInMinutes int32) ApiStopPointsStopPointGidDeparturesGetRequest {
	r.timeSpanInMinutes = &timeSpanInMinutes
	return r
}

// The maximum number of departures for a single line in a specific direction.
func (r ApiStopPointsStopPointGidDeparturesGetRequest) MaxDeparturesPerLineAndDirection(maxDeparturesPerLineAndDirection int32) ApiStopPointsStopPointGidDeparturesGetRequest {
	r.maxDeparturesPerLineAndDirection = &maxDeparturesPerLineAndDirection
	return r
}

// The number of results to return.
func (r ApiStopPointsStopPointGidDeparturesGetRequest) Limit(limit int32) ApiStopPointsStopPointGidDeparturesGetRequest {
	r.limit = &limit
	return r
}

// The zero-based start offset of the pagination.
func (r ApiStopPointsStopPointGidDeparturesGetRequest) Offset(offset int32) ApiStopPointsStopPointGidDeparturesGetRequest {
	r.offset = &offset
	return r
}

// Includes occupancy in departure.
func (r ApiStopPointsStopPointGidDeparturesGetRequest) IncludeOccupancy(includeOccupancy bool) ApiStopPointsStopPointGidDeparturesGetRequest {
	r.includeOccupancy = &includeOccupancy
	return r
}

// Filter result by last stop on journey. Must be a 16-digit Västtrafik stop area
func (r ApiStopPointsStopPointGidDeparturesGetRequest) DirectionGid(directionGid string) ApiStopPointsStopPointGidDeparturesGetRequest {
	r.directionGid = &directionGid
	return r
}

func (r ApiStopPointsStopPointGidDeparturesGetRequest) Execute() (*VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsGetDeparturesResponse, *http.Response, error) {
	return r.ApiService.StopPointsStopPointGidDeparturesGetExecute(r)
}

/*
StopPointsStopPointGidDeparturesGet Returns departures from the specified stop point at the specified time.

Sample request:

    GET /stop-points/9022014001760003/departures

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param stopPointGid The 16-digit Västtrafik gid of the stop point.
 @return ApiStopPointsStopPointGidDeparturesGetRequest
*/
func (a *StopPointsApiService) StopPointsStopPointGidDeparturesGet(ctx context.Context, stopPointGid string) ApiStopPointsStopPointGidDeparturesGetRequest {
	return ApiStopPointsStopPointGidDeparturesGetRequest{
		ApiService: a,
		ctx: ctx,
		stopPointGid: stopPointGid,
	}
}

// Execute executes the request
//  @return VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsGetDeparturesResponse
func (a *StopPointsApiService) StopPointsStopPointGidDeparturesGetExecute(r ApiStopPointsStopPointGidDeparturesGetRequest) (*VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsGetDeparturesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsGetDeparturesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StopPointsApiService.StopPointsStopPointGidDeparturesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/stop-points/{stopPointGid}/departures"
	localVarPath = strings.Replace(localVarPath, "{"+"stopPointGid"+"}", url.PathEscape(parameterValueToString(r.stopPointGid, "stopPointGid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startDateTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startDateTime", r.startDateTime, "")
	}
	if r.timeSpanInMinutes != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timeSpanInMinutes", r.timeSpanInMinutes, "")
	}
	if r.maxDeparturesPerLineAndDirection != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "maxDeparturesPerLineAndDirection", r.maxDeparturesPerLineAndDirection, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.includeOccupancy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeOccupancy", r.includeOccupancy, "")
	}
	if r.directionGid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "directionGid", r.directionGid, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/plain", "application/json", "text/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v VTApiPlaneraResaWebV4ModelsApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v MicrosoftAspNetCoreMvcProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}