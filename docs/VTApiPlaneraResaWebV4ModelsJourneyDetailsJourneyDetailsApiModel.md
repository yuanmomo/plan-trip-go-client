# VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DepartureAccessLink** | Pointer to [**VTApiPlaneraResaWebV4ModelsJourneyDetailsDepartureAccessLinkApiModel**](VTApiPlaneraResaWebV4ModelsJourneyDetailsDepartureAccessLinkApiModel.md) |  | [optional] 
**TripLegs** | Pointer to [**[]VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel**](VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel.md) | Detailed information, including stops, about the trip legs. | [optional] 
**ConnectionLinks** | Pointer to [**[]VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel**](VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel.md) | A list of ConnectionLinks between TripLegs, when applicable. The internal order of TripLegs and ConnectionLinks is defined by Index-property on the objects. | [optional] 
**ArrivalAccessLink** | Pointer to [**VTApiPlaneraResaWebV4ModelsJourneyDetailsArrivalAccessLinkApiModel**](VTApiPlaneraResaWebV4ModelsJourneyDetailsArrivalAccessLinkApiModel.md) |  | [optional] 
**DestinationLink** | Pointer to [**VTApiPlaneraResaWebV4ModelsJourneyDetailsDestinationLinkApiModel**](VTApiPlaneraResaWebV4ModelsJourneyDetailsDestinationLinkApiModel.md) |  | [optional] 
**TicketSuggestionsResult** | Pointer to [**VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel**](VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel.md) |  | [optional] 
**TariffZones** | Pointer to [**[]VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel**](VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel.md) | The tariff zones that the journey traverses. | [optional] 
**Occupancy** | Pointer to [**VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel**](VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel.md) |  | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel

`func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel() *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepartureAccessLink

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) GetDepartureAccessLink() VTApiPlaneraResaWebV4ModelsJourneyDetailsDepartureAccessLinkApiModel`

GetDepartureAccessLink returns the DepartureAccessLink field if non-nil, zero value otherwise.

### GetDepartureAccessLinkOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) GetDepartureAccessLinkOk() (*VTApiPlaneraResaWebV4ModelsJourneyDetailsDepartureAccessLinkApiModel, bool)`

GetDepartureAccessLinkOk returns a tuple with the DepartureAccessLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureAccessLink

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) SetDepartureAccessLink(v VTApiPlaneraResaWebV4ModelsJourneyDetailsDepartureAccessLinkApiModel)`

SetDepartureAccessLink sets DepartureAccessLink field to given value.

### HasDepartureAccessLink

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) HasDepartureAccessLink() bool`

HasDepartureAccessLink returns a boolean if a field has been set.

### GetTripLegs

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) GetTripLegs() []VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel`

GetTripLegs returns the TripLegs field if non-nil, zero value otherwise.

### GetTripLegsOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) GetTripLegsOk() (*[]VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel, bool)`

GetTripLegsOk returns a tuple with the TripLegs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTripLegs

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) SetTripLegs(v []VTApiPlaneraResaWebV4ModelsJourneyDetailsTripLegDetailsApiModel)`

SetTripLegs sets TripLegs field to given value.

### HasTripLegs

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) HasTripLegs() bool`

HasTripLegs returns a boolean if a field has been set.

### SetTripLegsNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) SetTripLegsNil(b bool)`

 SetTripLegsNil sets the value for TripLegs to be an explicit nil

### UnsetTripLegs
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) UnsetTripLegs()`

UnsetTripLegs ensures that no value is present for TripLegs, not even an explicit nil
### GetConnectionLinks

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) GetConnectionLinks() []VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel`

GetConnectionLinks returns the ConnectionLinks field if non-nil, zero value otherwise.

### GetConnectionLinksOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) GetConnectionLinksOk() (*[]VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel, bool)`

GetConnectionLinksOk returns a tuple with the ConnectionLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionLinks

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) SetConnectionLinks(v []VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel)`

SetConnectionLinks sets ConnectionLinks field to given value.

### HasConnectionLinks

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) HasConnectionLinks() bool`

HasConnectionLinks returns a boolean if a field has been set.

### SetConnectionLinksNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) SetConnectionLinksNil(b bool)`

 SetConnectionLinksNil sets the value for ConnectionLinks to be an explicit nil

### UnsetConnectionLinks
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) UnsetConnectionLinks()`

UnsetConnectionLinks ensures that no value is present for ConnectionLinks, not even an explicit nil
### GetArrivalAccessLink

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) GetArrivalAccessLink() VTApiPlaneraResaWebV4ModelsJourneyDetailsArrivalAccessLinkApiModel`

GetArrivalAccessLink returns the ArrivalAccessLink field if non-nil, zero value otherwise.

### GetArrivalAccessLinkOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) GetArrivalAccessLinkOk() (*VTApiPlaneraResaWebV4ModelsJourneyDetailsArrivalAccessLinkApiModel, bool)`

GetArrivalAccessLinkOk returns a tuple with the ArrivalAccessLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalAccessLink

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) SetArrivalAccessLink(v VTApiPlaneraResaWebV4ModelsJourneyDetailsArrivalAccessLinkApiModel)`

SetArrivalAccessLink sets ArrivalAccessLink field to given value.

### HasArrivalAccessLink

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) HasArrivalAccessLink() bool`

HasArrivalAccessLink returns a boolean if a field has been set.

### GetDestinationLink

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) GetDestinationLink() VTApiPlaneraResaWebV4ModelsJourneyDetailsDestinationLinkApiModel`

GetDestinationLink returns the DestinationLink field if non-nil, zero value otherwise.

### GetDestinationLinkOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) GetDestinationLinkOk() (*VTApiPlaneraResaWebV4ModelsJourneyDetailsDestinationLinkApiModel, bool)`

GetDestinationLinkOk returns a tuple with the DestinationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLink

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) SetDestinationLink(v VTApiPlaneraResaWebV4ModelsJourneyDetailsDestinationLinkApiModel)`

SetDestinationLink sets DestinationLink field to given value.

### HasDestinationLink

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) HasDestinationLink() bool`

HasDestinationLink returns a boolean if a field has been set.

### GetTicketSuggestionsResult

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) GetTicketSuggestionsResult() VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel`

GetTicketSuggestionsResult returns the TicketSuggestionsResult field if non-nil, zero value otherwise.

### GetTicketSuggestionsResultOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) GetTicketSuggestionsResultOk() (*VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel, bool)`

GetTicketSuggestionsResultOk returns a tuple with the TicketSuggestionsResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketSuggestionsResult

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) SetTicketSuggestionsResult(v VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel)`

SetTicketSuggestionsResult sets TicketSuggestionsResult field to given value.

### HasTicketSuggestionsResult

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) HasTicketSuggestionsResult() bool`

HasTicketSuggestionsResult returns a boolean if a field has been set.

### GetTariffZones

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) GetTariffZones() []VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel`

GetTariffZones returns the TariffZones field if non-nil, zero value otherwise.

### GetTariffZonesOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) GetTariffZonesOk() (*[]VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel, bool)`

GetTariffZonesOk returns a tuple with the TariffZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffZones

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) SetTariffZones(v []VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel)`

SetTariffZones sets TariffZones field to given value.

### HasTariffZones

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) HasTariffZones() bool`

HasTariffZones returns a boolean if a field has been set.

### SetTariffZonesNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) SetTariffZonesNil(b bool)`

 SetTariffZonesNil sets the value for TariffZones to be an explicit nil

### UnsetTariffZones
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) UnsetTariffZones()`

UnsetTariffZones ensures that no value is present for TariffZones, not even an explicit nil
### GetOccupancy

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) GetOccupancy() VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel`

GetOccupancy returns the Occupancy field if non-nil, zero value otherwise.

### GetOccupancyOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) GetOccupancyOk() (*VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel, bool)`

GetOccupancyOk returns a tuple with the Occupancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupancy

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) SetOccupancy(v VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel)`

SetOccupancy sets Occupancy field to given value.

### HasOccupancy

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsJourneyDetailsApiModel) HasOccupancy() bool`

HasOccupancy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


