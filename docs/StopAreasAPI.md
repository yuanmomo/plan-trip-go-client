# \StopAreasApi

All URIs are relative to *https://ext-api.vasttrafik.se/pr/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StopAreasStopAreaGidArrivalsDetailsReferenceDetailsGet**](StopAreasApi.md#StopAreasStopAreaGidArrivalsDetailsReferenceDetailsGet) | **Get** /stop-areas/{stopAreaGid}/arrivals/{detailsReference}/details | Returns details about an arrival.
[**StopAreasStopAreaGidArrivalsGet**](StopAreasApi.md#StopAreasStopAreaGidArrivalsGet) | **Get** /stop-areas/{stopAreaGid}/arrivals | Returns arrivals to the specified stop area at the specified time.
[**StopAreasStopAreaGidDeparturesDetailsReferenceDetailsGet**](StopAreasApi.md#StopAreasStopAreaGidDeparturesDetailsReferenceDetailsGet) | **Get** /stop-areas/{stopAreaGid}/departures/{detailsReference}/details | Returns details about a departure.
[**StopAreasStopAreaGidDeparturesGet**](StopAreasApi.md#StopAreasStopAreaGidDeparturesGet) | **Get** /stop-areas/{stopAreaGid}/departures | Returns departures from the specified stop area at the specified time.



## StopAreasStopAreaGidArrivalsDetailsReferenceDetailsGet

> VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalDetailsApiModel StopAreasStopAreaGidArrivalsDetailsReferenceDetailsGet(ctx, detailsReference, stopAreaGid).Includes(includes).Execute()

Returns details about an arrival.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    detailsReference := "detailsReference_example" // string | The reference to the service journey, received from the arrivals call. A detailsReference is only valid during the same day as it was generated.
    stopAreaGid := "stopAreaGid_example" // string | 
    includes := []openapiclient.VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType{openapiclient.VT.ApiPlaneraResa.Web.V4.Models.ArrivalDetailsIncludeType("servicejourneycalls")} // []VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType | The additional information to include in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StopAreasApi.StopAreasStopAreaGidArrivalsDetailsReferenceDetailsGet(context.Background(), detailsReference, stopAreaGid).Includes(includes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StopAreasApi.StopAreasStopAreaGidArrivalsDetailsReferenceDetailsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopAreasStopAreaGidArrivalsDetailsReferenceDetailsGet`: VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalDetailsApiModel
    fmt.Fprintf(os.Stdout, "Response from `StopAreasApi.StopAreasStopAreaGidArrivalsDetailsReferenceDetailsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**detailsReference** | **string** | The reference to the service journey, received from the arrivals call. A detailsReference is only valid during the same day as it was generated. | 
**stopAreaGid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopAreasStopAreaGidArrivalsDetailsReferenceDetailsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includes** | [**[]VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType**](VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType.md) | The additional information to include in the response. | 

### Return type

[**VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalDetailsApiModel**](VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalDetailsApiModel.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopAreasStopAreaGidArrivalsGet

> VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsGetArrivalsResponse StopAreasStopAreaGidArrivalsGet(ctx, stopAreaGid).StartDateTime(startDateTime).Platforms(platforms).TimeSpanInMinutes(timeSpanInMinutes).MaxArrivalsPerLineAndDirection(maxArrivalsPerLineAndDirection).Limit(limit).Offset(offset).DirectionGid(directionGid).Execute()

Returns arrivals to the specified stop area at the specified time.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    stopAreaGid := "stopAreaGid_example" // string | The 16-digit Västtrafik gid of the stop area.
    startDateTime := time.Now() // time.Time | The start of the time interval for which to get upcoming arrivals. Must be specified in RFC 3339 format or be null which means that the current time on the server is used. (optional)
    platforms := "platforms_example" // string | Filter results by platform. Multiple platforms are separated by comma. Case sensitive. (optional)
    timeSpanInMinutes := int32(56) // int32 | The number of minutes from the start time for which to get upcoming arrivals. Allowed values are between 0 and 1440. (optional) (default to 60)
    maxArrivalsPerLineAndDirection := int32(56) // int32 | The maximum number of arrivals for a single line in a specific direction. (optional) (default to 2)
    limit := int32(56) // int32 | The number of results to return. (optional) (default to 10)
    offset := int32(56) // int32 | The zero-based start offset of the pagination. (optional) (default to 0)
    directionGid := "directionGid_example" // string | Filter result by last stop on journey. Must be a 16-digit Västtrafik stop area (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StopAreasApi.StopAreasStopAreaGidArrivalsGet(context.Background(), stopAreaGid).StartDateTime(startDateTime).Platforms(platforms).TimeSpanInMinutes(timeSpanInMinutes).MaxArrivalsPerLineAndDirection(maxArrivalsPerLineAndDirection).Limit(limit).Offset(offset).DirectionGid(directionGid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StopAreasApi.StopAreasStopAreaGidArrivalsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopAreasStopAreaGidArrivalsGet`: VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsGetArrivalsResponse
    fmt.Fprintf(os.Stdout, "Response from `StopAreasApi.StopAreasStopAreaGidArrivalsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stopAreaGid** | **string** | The 16-digit Västtrafik gid of the stop area. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopAreasStopAreaGidArrivalsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDateTime** | **time.Time** | The start of the time interval for which to get upcoming arrivals. Must be specified in RFC 3339 format or be null which means that the current time on the server is used. | 
 **platforms** | **string** | Filter results by platform. Multiple platforms are separated by comma. Case sensitive. | 
 **timeSpanInMinutes** | **int32** | The number of minutes from the start time for which to get upcoming arrivals. Allowed values are between 0 and 1440. | [default to 60]
 **maxArrivalsPerLineAndDirection** | **int32** | The maximum number of arrivals for a single line in a specific direction. | [default to 2]
 **limit** | **int32** | The number of results to return. | [default to 10]
 **offset** | **int32** | The zero-based start offset of the pagination. | [default to 0]
 **directionGid** | **string** | Filter result by last stop on journey. Must be a 16-digit Västtrafik stop area | 

### Return type

[**VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsGetArrivalsResponse**](VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsGetArrivalsResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopAreasStopAreaGidDeparturesDetailsReferenceDetailsGet

> VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureDetailsApiModel StopAreasStopAreaGidDeparturesDetailsReferenceDetailsGet(ctx, detailsReference, stopAreaGid).Includes(includes).Execute()

Returns details about a departure.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    detailsReference := "detailsReference_example" // string | The reference to the service journey, received from the departures call. A detailsReference is only valid during the same day as it was generated.
    stopAreaGid := "stopAreaGid_example" // string | 
    includes := []openapiclient.VTApiPlaneraResaWebV4ModelsDepartureDetailsIncludeType{openapiclient.VT.ApiPlaneraResa.Web.V4.Models.DepartureDetailsIncludeType("servicejourneycalls")} // []VTApiPlaneraResaWebV4ModelsDepartureDetailsIncludeType | The additional information to include in the response. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StopAreasApi.StopAreasStopAreaGidDeparturesDetailsReferenceDetailsGet(context.Background(), detailsReference, stopAreaGid).Includes(includes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StopAreasApi.StopAreasStopAreaGidDeparturesDetailsReferenceDetailsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopAreasStopAreaGidDeparturesDetailsReferenceDetailsGet`: VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureDetailsApiModel
    fmt.Fprintf(os.Stdout, "Response from `StopAreasApi.StopAreasStopAreaGidDeparturesDetailsReferenceDetailsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**detailsReference** | **string** | The reference to the service journey, received from the departures call. A detailsReference is only valid during the same day as it was generated. | 
**stopAreaGid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopAreasStopAreaGidDeparturesDetailsReferenceDetailsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includes** | [**[]VTApiPlaneraResaWebV4ModelsDepartureDetailsIncludeType**](VTApiPlaneraResaWebV4ModelsDepartureDetailsIncludeType.md) | The additional information to include in the response. | 

### Return type

[**VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureDetailsApiModel**](VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureDetailsApiModel.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopAreasStopAreaGidDeparturesGet

> VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsGetDeparturesResponse StopAreasStopAreaGidDeparturesGet(ctx, stopAreaGid).StartDateTime(startDateTime).Platforms(platforms).TimeSpanInMinutes(timeSpanInMinutes).MaxDeparturesPerLineAndDirection(maxDeparturesPerLineAndDirection).Limit(limit).Offset(offset).IncludeOccupancy(includeOccupancy).DirectionGid(directionGid).Execute()

Returns departures from the specified stop area at the specified time.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    stopAreaGid := "stopAreaGid_example" // string | The 16-digit Västtrafik gid of the stop area.
    startDateTime := time.Now() // time.Time | The start of the time interval for which to get upcoming departures. Must be specified in RFC 3339 format or be null which means that the current time on the server is used. (optional)
    platforms := "platforms_example" // string | Filter results by platform. Multiple platforms are separated by comma. Case sensitive. (optional)
    timeSpanInMinutes := int32(56) // int32 | The number of minutes from the start time for which to get upcoming departures. Allowed values are between 0 and 1440. (optional) (default to 60)
    maxDeparturesPerLineAndDirection := int32(56) // int32 | The maximum number of departures for a single line in a specific direction. (optional) (default to 2)
    limit := int32(56) // int32 | The number of results to return. (optional) (default to 10)
    offset := int32(56) // int32 | The zero-based start offset of the pagination. (optional) (default to 0)
    includeOccupancy := true // bool | Includes occupancy in departure. (optional) (default to false)
    directionGid := "directionGid_example" // string | Filter result by last stop on journey. Must be a 16-digit Västtrafik stop area (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StopAreasApi.StopAreasStopAreaGidDeparturesGet(context.Background(), stopAreaGid).StartDateTime(startDateTime).Platforms(platforms).TimeSpanInMinutes(timeSpanInMinutes).MaxDeparturesPerLineAndDirection(maxDeparturesPerLineAndDirection).Limit(limit).Offset(offset).IncludeOccupancy(includeOccupancy).DirectionGid(directionGid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StopAreasApi.StopAreasStopAreaGidDeparturesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopAreasStopAreaGidDeparturesGet`: VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsGetDeparturesResponse
    fmt.Fprintf(os.Stdout, "Response from `StopAreasApi.StopAreasStopAreaGidDeparturesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stopAreaGid** | **string** | The 16-digit Västtrafik gid of the stop area. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopAreasStopAreaGidDeparturesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDateTime** | **time.Time** | The start of the time interval for which to get upcoming departures. Must be specified in RFC 3339 format or be null which means that the current time on the server is used. | 
 **platforms** | **string** | Filter results by platform. Multiple platforms are separated by comma. Case sensitive. | 
 **timeSpanInMinutes** | **int32** | The number of minutes from the start time for which to get upcoming departures. Allowed values are between 0 and 1440. | [default to 60]
 **maxDeparturesPerLineAndDirection** | **int32** | The maximum number of departures for a single line in a specific direction. | [default to 2]
 **limit** | **int32** | The number of results to return. | [default to 10]
 **offset** | **int32** | The zero-based start offset of the pagination. | [default to 0]
 **includeOccupancy** | **bool** | Includes occupancy in departure. | [default to false]
 **directionGid** | **string** | Filter result by last stop on journey. Must be a 16-digit Västtrafik stop area | 

### Return type

[**VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsGetDeparturesResponse**](VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsGetDeparturesResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

