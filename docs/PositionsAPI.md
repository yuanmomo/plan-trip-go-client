# \PositionsApi

All URIs are relative to *https://ext-api.vasttrafik.se/pr/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PositionsGet**](PositionsApi.md#PositionsGet) | **Get** /positions | Returns journey positions within a bounding box



## PositionsGet

> []VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel PositionsGet(ctx).LowerLeftLat(lowerLeftLat).LowerLeftLong(lowerLeftLong).UpperRightLat(upperRightLat).UpperRightLong(upperRightLong).TransportModes(transportModes).DetailsReferences(detailsReferences).LineDesignations(lineDesignations).Limit(limit).Execute()

Returns journey positions within a bounding box



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
    lowerLeftLat := float64(1.2) // float64 | Lower left latitude of bounding box.
    lowerLeftLong := float64(1.2) // float64 | Lower left longitude of bounding box.
    upperRightLat := float64(1.2) // float64 | Upper right latitude of bounding box.
    upperRightLong := float64(1.2) // float64 | Upper right longitude of bounding box.
    transportModes := []openapiclient.VTApiPlaneraResaCoreModelsPositionTransportMode{openapiclient.VT.ApiPlaneraResa.Core.Models.PositionTransportMode("tram")} // []VTApiPlaneraResaCoreModelsPositionTransportMode | The transport modes to include when searching for journeys, if none specified all transport modes are included. (optional)
    detailsReferences := []string{"Inner_example"} // []string | Filter journeys by one or more journey details reference. (optional)
    lineDesignations := []string{"Inner_example"} // []string | Only journeys running the given lineDesignations (case sensitive) are part of the result. (optional)
    limit := int32(56) // int32 | Maximum number of journeys in response. Range from 1 to 200. Defaults to 100 (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PositionsApi.PositionsGet(context.Background()).LowerLeftLat(lowerLeftLat).LowerLeftLong(lowerLeftLong).UpperRightLat(upperRightLat).UpperRightLong(upperRightLong).TransportModes(transportModes).DetailsReferences(detailsReferences).LineDesignations(lineDesignations).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PositionsApi.PositionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PositionsGet`: []VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel
    fmt.Fprintf(os.Stdout, "Response from `PositionsApi.PositionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPositionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lowerLeftLat** | **float64** | Lower left latitude of bounding box. | 
 **lowerLeftLong** | **float64** | Lower left longitude of bounding box. | 
 **upperRightLat** | **float64** | Upper right latitude of bounding box. | 
 **upperRightLong** | **float64** | Upper right longitude of bounding box. | 
 **transportModes** | [**[]VTApiPlaneraResaCoreModelsPositionTransportMode**](VTApiPlaneraResaCoreModelsPositionTransportMode.md) | The transport modes to include when searching for journeys, if none specified all transport modes are included. | 
 **detailsReferences** | **[]string** | Filter journeys by one or more journey details reference. | 
 **lineDesignations** | **[]string** | Only journeys running the given lineDesignations (case sensitive) are part of the result. | 
 **limit** | **int32** | Maximum number of journeys in response. Range from 1 to 200. Defaults to 100 | [default to 100]

### Return type

[**[]VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel**](VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

