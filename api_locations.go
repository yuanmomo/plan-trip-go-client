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
	"reflect"
)


// LocationsApiService LocationsApi service
type LocationsApiService service

type ApiLocationsByCoordinatesGetRequest struct {
	ctx context.Context
	ApiService *LocationsApiService
	latitude *float64
	longitude *float64
	radiusInMeters *int32
	types *[]VTApiPlaneraResaWebV4ModelsLocationByCoordinatesType
	limit *int32
	offset *int32
}

// The latitude.
func (r ApiLocationsByCoordinatesGetRequest) Latitude(latitude float64) ApiLocationsByCoordinatesGetRequest {
	r.latitude = &latitude
	return r
}

// The longitude.
func (r ApiLocationsByCoordinatesGetRequest) Longitude(longitude float64) ApiLocationsByCoordinatesGetRequest {
	r.longitude = &longitude
	return r
}

// The search radius from the coordinates specified in meters. Must be a positive integer &gt; 0.
func (r ApiLocationsByCoordinatesGetRequest) RadiusInMeters(radiusInMeters int32) ApiLocationsByCoordinatesGetRequest {
	r.radiusInMeters = &radiusInMeters
	return r
}

// The location types to include in the response, if none specified all locations types are included.
func (r ApiLocationsByCoordinatesGetRequest) Types(types []VTApiPlaneraResaWebV4ModelsLocationByCoordinatesType) ApiLocationsByCoordinatesGetRequest {
	r.types = &types
	return r
}

// The number of results to return.
func (r ApiLocationsByCoordinatesGetRequest) Limit(limit int32) ApiLocationsByCoordinatesGetRequest {
	r.limit = &limit
	return r
}

// The zero-based start offset of the pagination.
func (r ApiLocationsByCoordinatesGetRequest) Offset(offset int32) ApiLocationsByCoordinatesGetRequest {
	r.offset = &offset
	return r
}

func (r ApiLocationsByCoordinatesGetRequest) Execute() (*VTApiPlaneraResaWebV4ModelsLocationsGetLocationsResponse, *http.Response, error) {
	return r.ApiService.LocationsByCoordinatesGetExecute(r)
}

/*
LocationsByCoordinatesGet Returns the locations nearest the specified coordinates. Currently only stop areas, stop points and meta-stations are supported.

Sample request:

    GET /locations/by-coordinates?latitude=57.708734&longitude=11.974764&radiusInMeters=500&limit=10&offset=0

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLocationsByCoordinatesGetRequest
*/
func (a *LocationsApiService) LocationsByCoordinatesGet(ctx context.Context) ApiLocationsByCoordinatesGetRequest {
	return ApiLocationsByCoordinatesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return VTApiPlaneraResaWebV4ModelsLocationsGetLocationsResponse
func (a *LocationsApiService) LocationsByCoordinatesGetExecute(r ApiLocationsByCoordinatesGetRequest) (*VTApiPlaneraResaWebV4ModelsLocationsGetLocationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VTApiPlaneraResaWebV4ModelsLocationsGetLocationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LocationsApiService.LocationsByCoordinatesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/locations/by-coordinates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.latitude == nil {
		return localVarReturnValue, nil, reportError("latitude is required and must be specified")
	}
	if r.longitude == nil {
		return localVarReturnValue, nil, reportError("longitude is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "latitude", r.latitude, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "longitude", r.longitude, "")
	if r.radiusInMeters != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "radiusInMeters", r.radiusInMeters, "")
	}
	if r.types != nil {
		t := *r.types
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "types", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "types", t, "multi")
		}
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
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

type ApiLocationsByTextGetRequest struct {
	ctx context.Context
	ApiService *LocationsApiService
	q *string
	types *[]VTApiPlaneraResaWebV4ModelsLocationByTextType
	limit *int32
	offset *int32
}

// The search text (e.g. &#39;brunn&#39;, &#39;cent&#39; or &#39;Kungsgatan&#39;). The maximum length allowed is 256 characters.
func (r ApiLocationsByTextGetRequest) Q(q string) ApiLocationsByTextGetRequest {
	r.q = &q
	return r
}

// The location types to include in the response, if none specified all locations types are included.
func (r ApiLocationsByTextGetRequest) Types(types []VTApiPlaneraResaWebV4ModelsLocationByTextType) ApiLocationsByTextGetRequest {
	r.types = &types
	return r
}

// The number of results to return.
func (r ApiLocationsByTextGetRequest) Limit(limit int32) ApiLocationsByTextGetRequest {
	r.limit = &limit
	return r
}

// The zero-based start offset of the pagination.
func (r ApiLocationsByTextGetRequest) Offset(offset int32) ApiLocationsByTextGetRequest {
	r.offset = &offset
	return r
}

func (r ApiLocationsByTextGetRequest) Execute() (*VTApiPlaneraResaWebV4ModelsLocationsGetLocationsResponse, *http.Response, error) {
	return r.ApiService.LocationsByTextGetExecute(r)
}

/*
LocationsByTextGet Returns locations matching the specified text. Currently only stop areas, addresses, points of interest and meta-stations are supported.

Sample request:

    GET /locations/by-text?q=brunnsparken&limit=10&offset=0

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLocationsByTextGetRequest
*/
func (a *LocationsApiService) LocationsByTextGet(ctx context.Context) ApiLocationsByTextGetRequest {
	return ApiLocationsByTextGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return VTApiPlaneraResaWebV4ModelsLocationsGetLocationsResponse
func (a *LocationsApiService) LocationsByTextGetExecute(r ApiLocationsByTextGetRequest) (*VTApiPlaneraResaWebV4ModelsLocationsGetLocationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *VTApiPlaneraResaWebV4ModelsLocationsGetLocationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LocationsApiService.LocationsByTextGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/locations/by-text"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.q == nil {
		return localVarReturnValue, nil, reportError("q is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "")
	if r.types != nil {
		t := *r.types
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "types", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "types", t, "multi")
		}
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
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
