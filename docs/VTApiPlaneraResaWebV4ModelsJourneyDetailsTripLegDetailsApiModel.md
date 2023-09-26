# VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceJourneys** | Pointer to [**[]VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel**](VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel.md) | The service journey for the trip leg. | [optional] 
**CallsOnTripLeg** | Pointer to [**[]VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel**](VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel.md) | The calls on the trip leg. | [optional] 
**TripLegCoordinates** | Pointer to [**[]VTApiPlaneraResaWebV4ModelsCoordinateApiModel**](VTApiPlaneraResaWebV4ModelsCoordinateApiModel.md) | The coordinates for the trip leg. | [optional] 
**TariffZones** | Pointer to [**[]VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel**](VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel.md) | The tariff zones that the trip leg traverses. | [optional] 
**IsCancelled** | Pointer to **bool** | Flag indicating if the trip leg is cancelled. | [optional] 
**IsPartCancelled** | Pointer to **bool** | Flag indicating if the trip leg is partially cancelled. | [optional] 
**Occupancy** | Pointer to [**VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel**](VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel.md) |  | [optional] 
**JourneyLegIndex** | Pointer to **int32** | Index of Leg in Journey | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel

`func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel() *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceJourneys

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) GetServiceJourneys() []VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel`

GetServiceJourneys returns the ServiceJourneys field if non-nil, zero value otherwise.

### GetServiceJourneysOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) GetServiceJourneysOk() (*[]VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel, bool)`

GetServiceJourneysOk returns a tuple with the ServiceJourneys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceJourneys

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) SetServiceJourneys(v []VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel)`

SetServiceJourneys sets ServiceJourneys field to given value.

### HasServiceJourneys

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) HasServiceJourneys() bool`

HasServiceJourneys returns a boolean if a field has been set.

### SetServiceJourneysNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) SetServiceJourneysNil(b bool)`

 SetServiceJourneysNil sets the value for ServiceJourneys to be an explicit nil

### UnsetServiceJourneys
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) UnsetServiceJourneys()`

UnsetServiceJourneys ensures that no value is present for ServiceJourneys, not even an explicit nil
### GetCallsOnTripLeg

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) GetCallsOnTripLeg() []VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel`

GetCallsOnTripLeg returns the CallsOnTripLeg field if non-nil, zero value otherwise.

### GetCallsOnTripLegOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) GetCallsOnTripLegOk() (*[]VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel, bool)`

GetCallsOnTripLegOk returns a tuple with the CallsOnTripLeg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallsOnTripLeg

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) SetCallsOnTripLeg(v []VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel)`

SetCallsOnTripLeg sets CallsOnTripLeg field to given value.

### HasCallsOnTripLeg

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) HasCallsOnTripLeg() bool`

HasCallsOnTripLeg returns a boolean if a field has been set.

### SetCallsOnTripLegNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) SetCallsOnTripLegNil(b bool)`

 SetCallsOnTripLegNil sets the value for CallsOnTripLeg to be an explicit nil

### UnsetCallsOnTripLeg
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) UnsetCallsOnTripLeg()`

UnsetCallsOnTripLeg ensures that no value is present for CallsOnTripLeg, not even an explicit nil
### GetTripLegCoordinates

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) GetTripLegCoordinates() []VTApiPlaneraResaWebV4ModelsCoordinateApiModel`

GetTripLegCoordinates returns the TripLegCoordinates field if non-nil, zero value otherwise.

### GetTripLegCoordinatesOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) GetTripLegCoordinatesOk() (*[]VTApiPlaneraResaWebV4ModelsCoordinateApiModel, bool)`

GetTripLegCoordinatesOk returns a tuple with the TripLegCoordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTripLegCoordinates

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) SetTripLegCoordinates(v []VTApiPlaneraResaWebV4ModelsCoordinateApiModel)`

SetTripLegCoordinates sets TripLegCoordinates field to given value.

### HasTripLegCoordinates

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) HasTripLegCoordinates() bool`

HasTripLegCoordinates returns a boolean if a field has been set.

### SetTripLegCoordinatesNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) SetTripLegCoordinatesNil(b bool)`

 SetTripLegCoordinatesNil sets the value for TripLegCoordinates to be an explicit nil

### UnsetTripLegCoordinates
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) UnsetTripLegCoordinates()`

UnsetTripLegCoordinates ensures that no value is present for TripLegCoordinates, not even an explicit nil
### GetTariffZones

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) GetTariffZones() []VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel`

GetTariffZones returns the TariffZones field if non-nil, zero value otherwise.

### GetTariffZonesOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) GetTariffZonesOk() (*[]VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel, bool)`

GetTariffZonesOk returns a tuple with the TariffZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffZones

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) SetTariffZones(v []VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel)`

SetTariffZones sets TariffZones field to given value.

### HasTariffZones

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) HasTariffZones() bool`

HasTariffZones returns a boolean if a field has been set.

### SetTariffZonesNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) SetTariffZonesNil(b bool)`

 SetTariffZonesNil sets the value for TariffZones to be an explicit nil

### UnsetTariffZones
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) UnsetTariffZones()`

UnsetTariffZones ensures that no value is present for TariffZones, not even an explicit nil
### GetIsCancelled

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) GetIsCancelled() bool`

GetIsCancelled returns the IsCancelled field if non-nil, zero value otherwise.

### GetIsCancelledOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) GetIsCancelledOk() (*bool, bool)`

GetIsCancelledOk returns a tuple with the IsCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCancelled

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) SetIsCancelled(v bool)`

SetIsCancelled sets IsCancelled field to given value.

### HasIsCancelled

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) HasIsCancelled() bool`

HasIsCancelled returns a boolean if a field has been set.

### GetIsPartCancelled

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) GetIsPartCancelled() bool`

GetIsPartCancelled returns the IsPartCancelled field if non-nil, zero value otherwise.

### GetIsPartCancelledOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) GetIsPartCancelledOk() (*bool, bool)`

GetIsPartCancelledOk returns a tuple with the IsPartCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartCancelled

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) SetIsPartCancelled(v bool)`

SetIsPartCancelled sets IsPartCancelled field to given value.

### HasIsPartCancelled

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) HasIsPartCancelled() bool`

HasIsPartCancelled returns a boolean if a field has been set.

### GetOccupancy

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) GetOccupancy() VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel`

GetOccupancy returns the Occupancy field if non-nil, zero value otherwise.

### GetOccupancyOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) GetOccupancyOk() (*VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel, bool)`

GetOccupancyOk returns a tuple with the Occupancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupancy

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) SetOccupancy(v VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel)`

SetOccupancy sets Occupancy field to given value.

### HasOccupancy

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) HasOccupancy() bool`

HasOccupancy returns a boolean if a field has been set.

### GetJourneyLegIndex

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) GetJourneyLegIndex() int32`

GetJourneyLegIndex returns the JourneyLegIndex field if non-nil, zero value otherwise.

### GetJourneyLegIndexOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) GetJourneyLegIndexOk() (*int32, bool)`

GetJourneyLegIndexOk returns a tuple with the JourneyLegIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJourneyLegIndex

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) SetJourneyLegIndex(v int32)`

SetJourneyLegIndex sets JourneyLegIndex field to given value.

### HasJourneyLegIndex

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel) HasJourneyLegIndex() bool`

HasJourneyLegIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


