# \LocationsApi

All URIs are relative to *https://ext-api.vasttrafik.se/pr/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LocationsByCoordinatesGet**](LocationsApi.md#LocationsByCoordinatesGet) | **Get** /locations/by-coordinates | Returns the locations nearest the specified coordinates. Currently only stop areas, stop points and meta-stations are supported.
[**LocationsByTextGet**](LocationsApi.md#LocationsByTextGet) | **Get** /locations/by-text | Returns locations matching the specified text. Currently only stop areas, addresses, points of interest and meta-stations are supported.



## LocationsByCoordinatesGet

> VTApiPlaneraResaWebV4ModelsLocationsGetLocationsResponse LocationsByCoordinatesGet(ctx).Latitude(latitude).Longitude(longitude).RadiusInMeters(radiusInMeters).Types(types).Limit(limit).Offset(offset).Execute()

Returns the locations nearest the specified coordinates. Currently only stop areas, stop points and meta-stations are supported.



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
    latitude := float64(1.2) // float64 | The latitude.
    longitude := float64(1.2) // float64 | The longitude.
    radiusInMeters := int32(56) // int32 | The search radius from the coordinates specified in meters. Must be a positive integer > 0. (optional) (default to 500)
    types := []openapiclient.VTApiPlaneraResaWebV4ModelsLocationByCoordinatesType{openapiclient.VT.ApiPlaneraResa.Web.V4.Models.LocationByCoordinatesType("stoparea")} // []VTApiPlaneraResaWebV4ModelsLocationByCoordinatesType | The location types to include in the response, if none specified all locations types are included. (optional)
    limit := int32(56) // int32 | The number of results to return. (optional) (default to 10)
    offset := int32(56) // int32 | The zero-based start offset of the pagination. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationsApi.LocationsByCoordinatesGet(context.Background()).Latitude(latitude).Longitude(longitude).RadiusInMeters(radiusInMeters).Types(types).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.LocationsByCoordinatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LocationsByCoordinatesGet`: VTApiPlaneraResaWebV4ModelsLocationsGetLocationsResponse
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.LocationsByCoordinatesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLocationsByCoordinatesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **latitude** | **float64** | The latitude. | 
 **longitude** | **float64** | The longitude. | 
 **radiusInMeters** | **int32** | The search radius from the coordinates specified in meters. Must be a positive integer &gt; 0. | [default to 500]
 **types** | [**[]VTApiPlaneraResaWebV4ModelsLocationByCoordinatesType**](VTApiPlaneraResaWebV4ModelsLocationByCoordinatesType.md) | The location types to include in the response, if none specified all locations types are included. | 
 **limit** | **int32** | The number of results to return. | [default to 10]
 **offset** | **int32** | The zero-based start offset of the pagination. | [default to 0]

### Return type

[**VTApiPlaneraResaWebV4ModelsLocationsGetLocationsResponse**](VTApiPlaneraResaWebV4ModelsLocationsGetLocationsResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LocationsByTextGet

> VTApiPlaneraResaWebV4ModelsLocationsGetLocationsResponse LocationsByTextGet(ctx).Q(q).Types(types).Limit(limit).Offset(offset).Execute()

Returns locations matching the specified text. Currently only stop areas, addresses, points of interest and meta-stations are supported.



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
    q := "q_example" // string | The search text (e.g. 'brunn', 'cent' or 'Kungsgatan'). The maximum length allowed is 256 characters.
    types := []openapiclient.VTApiPlaneraResaWebV4ModelsLocationByTextType{openapiclient.VT.ApiPlaneraResa.Web.V4.Models.LocationByTextType("stoparea")} // []VTApiPlaneraResaWebV4ModelsLocationByTextType | The location types to include in the response, if none specified all locations types are included. (optional)
    limit := int32(56) // int32 | The number of results to return. (optional) (default to 10)
    offset := int32(56) // int32 | The zero-based start offset of the pagination. (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LocationsApi.LocationsByTextGet(context.Background()).Q(q).Types(types).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LocationsApi.LocationsByTextGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LocationsByTextGet`: VTApiPlaneraResaWebV4ModelsLocationsGetLocationsResponse
    fmt.Fprintf(os.Stdout, "Response from `LocationsApi.LocationsByTextGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLocationsByTextGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | The search text (e.g. &#39;brunn&#39;, &#39;cent&#39; or &#39;Kungsgatan&#39;). The maximum length allowed is 256 characters. | 
 **types** | [**[]VTApiPlaneraResaWebV4ModelsLocationByTextType**](VTApiPlaneraResaWebV4ModelsLocationByTextType.md) | The location types to include in the response, if none specified all locations types are included. | 
 **limit** | **int32** | The number of results to return. | [default to 10]
 **offset** | **int32** | The zero-based start offset of the pagination. | [default to 0]

### Return type

[**VTApiPlaneraResaWebV4ModelsLocationsGetLocationsResponse**](VTApiPlaneraResaWebV4ModelsLocationsGetLocationsResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

