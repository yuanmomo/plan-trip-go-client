# \JourneysApi

All URIs are relative to *https://ext-api.vasttrafik.se/pr/v4*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JourneysDetailsReferenceDetailsGet**](JourneysApi.md#JourneysDetailsReferenceDetailsGet) | **Get** /journeys/{detailsReference}/details | Returns details about a journey.
[**JourneysGet**](JourneysApi.md#JourneysGet) | **Get** /journeys | Returns journeys matching the specified search parameters.
[**JourneysReconstructGet**](JourneysApi.md#JourneysReconstructGet) | **Get** /journeys/reconstruct | Reconstructs a journey based on the given reconstruction reference, received from the search journeys query.
[**JourneysValidTimeIntervalGet**](JourneysApi.md#JourneysValidTimeIntervalGet) | **Get** /journeys/valid-time-interval | Returns a time interval for when journey data is available.



## JourneysDetailsReferenceDetailsGet

> VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel JourneysDetailsReferenceDetailsGet(ctx, detailsReference).Includes(includes).ChannelIds(channelIds).ProductTypes(productTypes).TravellerCategories(travellerCategories).Execute()

Returns details about a journey.



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
    detailsReference := "detailsReference_example" // string | The reference to the journey, received from the search journeys query. A detailsReference is only valid during the same day as it was generated.
    includes := []openapiclient.VTApiPlaneraResaWebV4ModelsJourneyDetailsIncludeType{openapiclient.VT.ApiPlaneraResa.Web.V4.Models.JourneyDetailsIncludeType("ticketsuggestions")} // []VTApiPlaneraResaWebV4ModelsJourneyDetailsIncludeType | The additional information to include in the response. (optional)
    channelIds := []int32{int32(123)} // []int32 | List of channel ids to include if 'ticketsuggestions' is set in the 'includes' parameter. Optional parameter. (optional)
    productTypes := []int32{int32(123)} // []int32 | List of product type ids to include if 'ticketsuggestions' is set in the 'includes' parameter. Optional parameter. (optional)
    travellerCategories := []openapiclient.VTApiPlaneraResaCoreModelsTravellerCategory{openapiclient.VT.ApiPlaneraResa.Core.Models.TravellerCategory("unknown")} // []VTApiPlaneraResaCoreModelsTravellerCategory | List of traveller category ids to include if 'ticketsuggestions' is set in the 'includes' parameter. Optional parameter. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JourneysApi.JourneysDetailsReferenceDetailsGet(context.Background(), detailsReference).Includes(includes).ChannelIds(channelIds).ProductTypes(productTypes).TravellerCategories(travellerCategories).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JourneysApi.JourneysDetailsReferenceDetailsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JourneysDetailsReferenceDetailsGet`: VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel
    fmt.Fprintf(os.Stdout, "Response from `JourneysApi.JourneysDetailsReferenceDetailsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**detailsReference** | **string** | The reference to the journey, received from the search journeys query. A detailsReference is only valid during the same day as it was generated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJourneysDetailsReferenceDetailsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includes** | [**[]VTApiPlaneraResaWebV4ModelsJourneyDetailsIncludeType**](VTApiPlaneraResaWebV4ModelsJourneyDetailsIncludeType.md) | The additional information to include in the response. | 
 **channelIds** | **[]int32** | List of channel ids to include if &#39;ticketsuggestions&#39; is set in the &#39;includes&#39; parameter. Optional parameter. | 
 **productTypes** | **[]int32** | List of product type ids to include if &#39;ticketsuggestions&#39; is set in the &#39;includes&#39; parameter. Optional parameter. | 
 **travellerCategories** | [**[]VTApiPlaneraResaCoreModelsTravellerCategory**](VTApiPlaneraResaCoreModelsTravellerCategory.md) | List of traveller category ids to include if &#39;ticketsuggestions&#39; is set in the &#39;includes&#39; parameter. Optional parameter. | 

### Return type

[**VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel**](VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JourneysGet

> VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse JourneysGet(ctx).OriginGid(originGid).OriginName(originName).OriginLatitude(originLatitude).OriginLongitude(originLongitude).DestinationGid(destinationGid).DestinationName(destinationName).DestinationLatitude(destinationLatitude).DestinationLongitude(destinationLongitude).DateTime(dateTime).DateTimeRelatesTo(dateTimeRelatesTo).PaginationReference(paginationReference).Limit(limit).TransportModes(transportModes).TransportSubModes(transportSubModes).OnlyDirectConnections(onlyDirectConnections).IncludeNearbyStopAreas(includeNearbyStopAreas).ViaGid(viaGid).OriginWalk(originWalk).DestWalk(destWalk).OriginBike(originBike).DestBike(destBike).TotalBike(totalBike).OriginCar(originCar).DestCar(destCar).OriginPark(originPark).DestPark(destPark).InterchangeDurationInMinutes(interchangeDurationInMinutes).IncludeOccupancy(includeOccupancy).Execute()

Returns journeys matching the specified search parameters.



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
    originGid := "originGid_example" // string | The 16-digit Västtrafik gid of the origin location (which could be either a stop area (e.g. '9021014001760000'), a stop point (e.g. '9022014001760004') or meta-station (e.g. '0000000800000022')). (optional)
    originName := "originName_example" // string | The name of the origin location. The maximum length allowed is 256 characters. (optional)
    originLatitude := float64(1.2) // float64 | The latitude of the origin location. (optional)
    originLongitude := float64(1.2) // float64 | The longitude of the origin location. (optional)
    destinationGid := "destinationGid_example" // string | The 16-digit Västtrafik gid of the destination location (which could be either a stop area, stop point or meta-station). (optional)
    destinationName := "destinationName_example" // string | The name of the destination location. The maximum length allowed is 256 characters. (optional)
    destinationLatitude := float64(1.2) // float64 | The latitude of the destination location. (optional)
    destinationLongitude := float64(1.2) // float64 | The longitude of the destination location. (optional)
    dateTime := time.Now() // time.Time | The datetime for which to search journeys. Must be specified in RFC 3339 format or be null which means that the current time on the server is used. The related dateTimeRelatesTo parameter specifies if the time is related to the arrival or departure. (optional)
    dateTimeRelatesTo := openapiclient.VT.ApiPlaneraResa.Core.Models.DateTimeRelatesToType("departure") // VTApiPlaneraResaCoreModelsDateTimeRelatesToType | Specifies if the datetime is related to the departure or arrival of the journey. (optional)
    paginationReference := "paginationReference_example" // string | Pagination reference from a previous search. (optional)
    limit := int32(56) // int32 | The number of results to return. Not guaranteed to return the specified number of results and usually not more than 7 results. (optional) (default to 10)
    transportModes := []openapiclient.VTApiPlaneraResaWebV4ModelsJourneyTransportMode{openapiclient.VT.ApiPlaneraResa.Web.V4.Models.JourneyTransportMode("tram")} // []VTApiPlaneraResaWebV4ModelsJourneyTransportMode | The transport modes to include when searching for journeys, if none specified all transport modes are included. (optional)
    transportSubModes := []openapiclient.VTApiPlaneraResaWebV4ModelsJourneyTransportSubMode{openapiclient.VT.ApiPlaneraResa.Web.V4.Models.JourneyTransportSubMode("vasttagen")} // []VTApiPlaneraResaWebV4ModelsJourneyTransportSubMode | The transport sub modes to include when searching for journeys, if none specified all transport sub modes are included. Only supported in combination with transportMode 'train'. (optional)
    onlyDirectConnections := true // bool | Only include direct connections, e.g. journeys with one trip leg. (optional) (default to false)
    includeNearbyStopAreas := true // bool | Includes nearby stop areas when searching for a journey to or from a stop area or stop point. This means that the search algorithm will take additional stop points of other stop areas nearby the given start and destination stop area into account. These additional stop points are reachable by walk. E.g when true a journey suggestion may include a departure access link (initial walking leg) to a stop point of a stop area close by the specified origin stop area. (optional) (default to false)
    viaGid := int64(789) // int64 | The 16-digit Västtrafik gid of the via location (which must be a stop area). (optional)
    originWalk := "originWalk_example" // string | Enables/disables using footpaths in the beginning of a trip when searching from an address. To fine-tune the minimum and/or maximum distance to the next public transport station, provide these values separated by comma. The values are expressed in meters. To enable walk with a minimum distance of 0 meters and a maximum distance of 3 kilometers, set the parameter originWalk=1,0,3000. If default distances should be used, skip the values, e.g 1,,. This will enable walk with the default minimum (0 meters) and the default maximum (2000 meters). (optional)
    destWalk := "destWalk_example" // string | Enables/disables using footpaths at the end of a trip when searching to an address. To fine-tune the minimum and/or maximum distance to the next public transport station, provide these values separated by comma. The values are expressed in meters. To enable walk with a minimum distance of 0 meters and a maximum distance of 3 kilometers, set the parameter destWalk=1,0,3000. If default distances should be used, skip the values, e.g 1,,. This will enable walk with the default minimum (0 meters) and the default maximum (2000 meters). (optional)
    originBike := "originBike_example" // string | Enables/disables using bike paths in the beginning of a trip when searching from an address. To fine-tune the minimum and/or maximum distance to the next public transport station, provide these values separated by comma. The values are expressed in meters. To enable bike with a minimum distance of 1000 meters and a maximum distance of 5 kilometers, set the parameter originBike=1,1000,5000. If default distances should be used, skip the values, e.g 1,,. This will enable bike with the default minimum (0 meters) and the default maximum (3000 meters). (optional)
    destBike := "destBike_example" // string | Enables/disables using bike paths at the end of a trip when searching to an address. To fine-tune the minimum and/or maximum distance to the next public transport station, provide these values separated by comma. The values are expressed in meters. To enable bike with a minimum distance of 1000 meters and a maximum distance of 5 kilometers, set the parameter destBike=1,1000,5000. If default distances should be used, skip the values, e.g 1,,. This will enable bike with the default minimum (0 meters) and the default maximum (3000 meters). (optional)
    totalBike := "totalBike_example" // string | Enables/disables using bike routes for the whole trip. To fine-tune the minimum and/or maximum distance to the next public transport station, provide these values separated by comma. The values are expressed in meters. To enable bike with a minimum distance of 0 meters and a maximum distance of 20 kilometers, set the parameter totalBike=1,0,20000. If default distances should be used, skip the values, e.g 1,,. This will enable bike with the default minimum (0 meters) and the default maximum (25000 meters). (optional)
    originCar := "originCar_example" // string | Enables/disables using car in the beginning of a trip when searching from an address. To fine-tune the minimum and/or maximum distance to the next public transport station, provide these values separated by comma. The values are expressed in meters. To enable car with a minimum distance of 2000 meters and a maximum distance of 7 kilometers, set the parameter origincar=1,2000,7000. If default distances should be used, skip the values, e.g 1,,. This will enable car with the default minimum (0 meters) and the default maximum (5000 meters). (optional)
    destCar := "destCar_example" // string | Enables/disables using car at the end of a trip when searching to an address. To fine-tune the minimum and/or maximum distance to the next public transport station, provide these values separated by comma. The values are expressed in meters. To enable car with a minimum distance of 2000 meters and a maximum distance of 7 kilometers, set the parameter destCar=1,2000,7000. If default distances should be used, skip the values, e.g 1,,. This will enable car with the default minimum (0 meters) and the default maximum (5000 meters). (optional)
    originPark := "originPark_example" // string | Enables/disables using Park and Ride in the beginning of a trip when searching from an address. To fine-tune the minimum and/or maximum distance to the next public transport station, provide these values separated by comma. The values are expressed in meters. To enable Park and Ride with a minimum distance of 3000 meters and a maximum distance of 70 kilometers, set the parameter originPark=1,3000,70000. If default distances should be used, skip the values, e.g 1,,. This will enable Park and Ride with the default minimum (2000 meters) and the default maximum (50000 meters). (optional)
    destPark := "destPark_example" // string | Enables/disables using Park and Ride at the end of a trip when searching to an address. To fine-tune the minimum and/or maximum distance to the next public transport station, provide these values separated by comma. The values are expressed in meters. To enable Park and Ride with a minimum distance of 3000 meters and a maximum distance of 70 kilometers, set the parameter destPark=1,3000,70000. If default distances should be used, skip the values, e.g 1,,. This will enable Park and Ride with the default minimum (2000 meters) and the default maximum (50000 meters). (optional)
    interchangeDurationInMinutes := int32(56) // int32 | The minimum number of minutes between arrival and departure for a connection to be valid and the trip included in the search results, ignoring the default value. (optional)
    includeOccupancy := true // bool | Includes occupancy in journey. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JourneysApi.JourneysGet(context.Background()).OriginGid(originGid).OriginName(originName).OriginLatitude(originLatitude).OriginLongitude(originLongitude).DestinationGid(destinationGid).DestinationName(destinationName).DestinationLatitude(destinationLatitude).DestinationLongitude(destinationLongitude).DateTime(dateTime).DateTimeRelatesTo(dateTimeRelatesTo).PaginationReference(paginationReference).Limit(limit).TransportModes(transportModes).TransportSubModes(transportSubModes).OnlyDirectConnections(onlyDirectConnections).IncludeNearbyStopAreas(includeNearbyStopAreas).ViaGid(viaGid).OriginWalk(originWalk).DestWalk(destWalk).OriginBike(originBike).DestBike(destBike).TotalBike(totalBike).OriginCar(originCar).DestCar(destCar).OriginPark(originPark).DestPark(destPark).InterchangeDurationInMinutes(interchangeDurationInMinutes).IncludeOccupancy(includeOccupancy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JourneysApi.JourneysGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JourneysGet`: VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse
    fmt.Fprintf(os.Stdout, "Response from `JourneysApi.JourneysGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJourneysGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **originGid** | **string** | The 16-digit Västtrafik gid of the origin location (which could be either a stop area (e.g. &#39;9021014001760000&#39;), a stop point (e.g. &#39;9022014001760004&#39;) or meta-station (e.g. &#39;0000000800000022&#39;)). | 
 **originName** | **string** | The name of the origin location. The maximum length allowed is 256 characters. | 
 **originLatitude** | **float64** | The latitude of the origin location. | 
 **originLongitude** | **float64** | The longitude of the origin location. | 
 **destinationGid** | **string** | The 16-digit Västtrafik gid of the destination location (which could be either a stop area, stop point or meta-station). | 
 **destinationName** | **string** | The name of the destination location. The maximum length allowed is 256 characters. | 
 **destinationLatitude** | **float64** | The latitude of the destination location. | 
 **destinationLongitude** | **float64** | The longitude of the destination location. | 
 **dateTime** | **time.Time** | The datetime for which to search journeys. Must be specified in RFC 3339 format or be null which means that the current time on the server is used. The related dateTimeRelatesTo parameter specifies if the time is related to the arrival or departure. | 
 **dateTimeRelatesTo** | [**VTApiPlaneraResaCoreModelsDateTimeRelatesToType**](VTApiPlaneraResaCoreModelsDateTimeRelatesToType.md) | Specifies if the datetime is related to the departure or arrival of the journey. | 
 **paginationReference** | **string** | Pagination reference from a previous search. | 
 **limit** | **int32** | The number of results to return. Not guaranteed to return the specified number of results and usually not more than 7 results. | [default to 10]
 **transportModes** | [**[]VTApiPlaneraResaWebV4ModelsJourneyTransportMode**](VTApiPlaneraResaWebV4ModelsJourneyTransportMode.md) | The transport modes to include when searching for journeys, if none specified all transport modes are included. | 
 **transportSubModes** | [**[]VTApiPlaneraResaWebV4ModelsJourneyTransportSubMode**](VTApiPlaneraResaWebV4ModelsJourneyTransportSubMode.md) | The transport sub modes to include when searching for journeys, if none specified all transport sub modes are included. Only supported in combination with transportMode &#39;train&#39;. | 
 **onlyDirectConnections** | **bool** | Only include direct connections, e.g. journeys with one trip leg. | [default to false]
 **includeNearbyStopAreas** | **bool** | Includes nearby stop areas when searching for a journey to or from a stop area or stop point. This means that the search algorithm will take additional stop points of other stop areas nearby the given start and destination stop area into account. These additional stop points are reachable by walk. E.g when true a journey suggestion may include a departure access link (initial walking leg) to a stop point of a stop area close by the specified origin stop area. | [default to false]
 **viaGid** | **int64** | The 16-digit Västtrafik gid of the via location (which must be a stop area). | 
 **originWalk** | **string** | Enables/disables using footpaths in the beginning of a trip when searching from an address. To fine-tune the minimum and/or maximum distance to the next public transport station, provide these values separated by comma. The values are expressed in meters. To enable walk with a minimum distance of 0 meters and a maximum distance of 3 kilometers, set the parameter originWalk&#x3D;1,0,3000. If default distances should be used, skip the values, e.g 1,,. This will enable walk with the default minimum (0 meters) and the default maximum (2000 meters). | 
 **destWalk** | **string** | Enables/disables using footpaths at the end of a trip when searching to an address. To fine-tune the minimum and/or maximum distance to the next public transport station, provide these values separated by comma. The values are expressed in meters. To enable walk with a minimum distance of 0 meters and a maximum distance of 3 kilometers, set the parameter destWalk&#x3D;1,0,3000. If default distances should be used, skip the values, e.g 1,,. This will enable walk with the default minimum (0 meters) and the default maximum (2000 meters). | 
 **originBike** | **string** | Enables/disables using bike paths in the beginning of a trip when searching from an address. To fine-tune the minimum and/or maximum distance to the next public transport station, provide these values separated by comma. The values are expressed in meters. To enable bike with a minimum distance of 1000 meters and a maximum distance of 5 kilometers, set the parameter originBike&#x3D;1,1000,5000. If default distances should be used, skip the values, e.g 1,,. This will enable bike with the default minimum (0 meters) and the default maximum (3000 meters). | 
 **destBike** | **string** | Enables/disables using bike paths at the end of a trip when searching to an address. To fine-tune the minimum and/or maximum distance to the next public transport station, provide these values separated by comma. The values are expressed in meters. To enable bike with a minimum distance of 1000 meters and a maximum distance of 5 kilometers, set the parameter destBike&#x3D;1,1000,5000. If default distances should be used, skip the values, e.g 1,,. This will enable bike with the default minimum (0 meters) and the default maximum (3000 meters). | 
 **totalBike** | **string** | Enables/disables using bike routes for the whole trip. To fine-tune the minimum and/or maximum distance to the next public transport station, provide these values separated by comma. The values are expressed in meters. To enable bike with a minimum distance of 0 meters and a maximum distance of 20 kilometers, set the parameter totalBike&#x3D;1,0,20000. If default distances should be used, skip the values, e.g 1,,. This will enable bike with the default minimum (0 meters) and the default maximum (25000 meters). | 
 **originCar** | **string** | Enables/disables using car in the beginning of a trip when searching from an address. To fine-tune the minimum and/or maximum distance to the next public transport station, provide these values separated by comma. The values are expressed in meters. To enable car with a minimum distance of 2000 meters and a maximum distance of 7 kilometers, set the parameter origincar&#x3D;1,2000,7000. If default distances should be used, skip the values, e.g 1,,. This will enable car with the default minimum (0 meters) and the default maximum (5000 meters). | 
 **destCar** | **string** | Enables/disables using car at the end of a trip when searching to an address. To fine-tune the minimum and/or maximum distance to the next public transport station, provide these values separated by comma. The values are expressed in meters. To enable car with a minimum distance of 2000 meters and a maximum distance of 7 kilometers, set the parameter destCar&#x3D;1,2000,7000. If default distances should be used, skip the values, e.g 1,,. This will enable car with the default minimum (0 meters) and the default maximum (5000 meters). | 
 **originPark** | **string** | Enables/disables using Park and Ride in the beginning of a trip when searching from an address. To fine-tune the minimum and/or maximum distance to the next public transport station, provide these values separated by comma. The values are expressed in meters. To enable Park and Ride with a minimum distance of 3000 meters and a maximum distance of 70 kilometers, set the parameter originPark&#x3D;1,3000,70000. If default distances should be used, skip the values, e.g 1,,. This will enable Park and Ride with the default minimum (2000 meters) and the default maximum (50000 meters). | 
 **destPark** | **string** | Enables/disables using Park and Ride at the end of a trip when searching to an address. To fine-tune the minimum and/or maximum distance to the next public transport station, provide these values separated by comma. The values are expressed in meters. To enable Park and Ride with a minimum distance of 3000 meters and a maximum distance of 70 kilometers, set the parameter destPark&#x3D;1,3000,70000. If default distances should be used, skip the values, e.g 1,,. This will enable Park and Ride with the default minimum (2000 meters) and the default maximum (50000 meters). | 
 **interchangeDurationInMinutes** | **int32** | The minimum number of minutes between arrival and departure for a connection to be valid and the trip included in the search results, ignoring the default value. | 
 **includeOccupancy** | **bool** | Includes occupancy in journey. | [default to false]

### Return type

[**VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse**](VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JourneysReconstructGet

> VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel JourneysReconstructGet(ctx).Ref(ref).IncludeOccupancy(includeOccupancy).Execute()

Reconstructs a journey based on the given reconstruction reference, received from the search journeys query.



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
    ref := "ref_example" // string | The reconstruction reference. A reconstructionReference is valid as long as the original journey search is valid.
    includeOccupancy := true // bool | Includes occupancy in journey. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JourneysApi.JourneysReconstructGet(context.Background()).Ref(ref).IncludeOccupancy(includeOccupancy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JourneysApi.JourneysReconstructGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JourneysReconstructGet`: VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel
    fmt.Fprintf(os.Stdout, "Response from `JourneysApi.JourneysReconstructGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiJourneysReconstructGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ref** | **string** | The reconstruction reference. A reconstructionReference is valid as long as the original journey search is valid. | 
 **includeOccupancy** | **bool** | Includes occupancy in journey. | [default to false]

### Return type

[**VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel**](VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JourneysValidTimeIntervalGet

> VTApiPlaneraResaWebV4ModelsValidTimeIntervalApiModel JourneysValidTimeIntervalGet(ctx).Execute()

Returns a time interval for when journey data is available.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JourneysApi.JourneysValidTimeIntervalGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JourneysApi.JourneysValidTimeIntervalGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JourneysValidTimeIntervalGet`: VTApiPlaneraResaWebV4ModelsValidTimeIntervalApiModel
    fmt.Fprintf(os.Stdout, "Response from `JourneysApi.JourneysValidTimeIntervalGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiJourneysValidTimeIntervalGetRequest struct via the builder pattern


### Return type

[**VTApiPlaneraResaWebV4ModelsValidTimeIntervalApiModel**](VTApiPlaneraResaWebV4ModelsValidTimeIntervalApiModel.md)

### Authorization

[auth](../README.md#auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, text/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

