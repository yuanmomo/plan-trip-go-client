# VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReconstructionReference** | Pointer to **NullableString** | A reference that can be used to reconstruct this journey at a later time. | [optional] 
**DetailsReference** | Pointer to **NullableString** | A reference that should be used when getting detailed information about the journey. | [optional] 
**DepartureAccessLink** | Pointer to [**VTApiPlaneraResaWebV4ModelsJourneysDepartureAccessLinkApiModel**](VTApiPlaneraResaWebV4ModelsJourneysDepartureAccessLinkApiModel.md) |  | [optional] 
**TripLegs** | Pointer to [**[]VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel**](VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel.md) | A list of trip legs on a journey, when applicable. A journey has either one or more trip legs or one  destination link. | [optional] 
**ConnectionLinks** | Pointer to [**[]VTApiPlaneraResaWebV4ModelsJourneysConnectionLinkApiModel**](VTApiPlaneraResaWebV4ModelsJourneysConnectionLinkApiModel.md) | A list of ConnectionLinks between TripLegs, when applicable. The internal order of TripLegs and ConnectionLinks is defined by Index-property on the objects. | [optional] 
**ArrivalAccessLink** | Pointer to [**VTApiPlaneraResaWebV4ModelsJourneysArrivalAccessLinkApiModel**](VTApiPlaneraResaWebV4ModelsJourneysArrivalAccessLinkApiModel.md) |  | [optional] 
**DestinationLink** | Pointer to [**VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel**](VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel.md) |  | [optional] 
**IsDeparted** | Pointer to **NullableBool** | Flag indicating if the first trip leg of the journey is departed. | [optional] 
**Occupancy** | Pointer to [**VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel**](VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel.md) |  | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel

`func NewVTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel() *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel instantiates a new VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsJourneysJourneyApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsJourneysJourneyApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneysJourneyApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReconstructionReference

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) GetReconstructionReference() string`

GetReconstructionReference returns the ReconstructionReference field if non-nil, zero value otherwise.

### GetReconstructionReferenceOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) GetReconstructionReferenceOk() (*string, bool)`

GetReconstructionReferenceOk returns a tuple with the ReconstructionReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReconstructionReference

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) SetReconstructionReference(v string)`

SetReconstructionReference sets ReconstructionReference field to given value.

### HasReconstructionReference

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) HasReconstructionReference() bool`

HasReconstructionReference returns a boolean if a field has been set.

### SetReconstructionReferenceNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) SetReconstructionReferenceNil(b bool)`

 SetReconstructionReferenceNil sets the value for ReconstructionReference to be an explicit nil

### UnsetReconstructionReference
`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) UnsetReconstructionReference()`

UnsetReconstructionReference ensures that no value is present for ReconstructionReference, not even an explicit nil
### GetDetailsReference

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) GetDetailsReference() string`

GetDetailsReference returns the DetailsReference field if non-nil, zero value otherwise.

### GetDetailsReferenceOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) GetDetailsReferenceOk() (*string, bool)`

GetDetailsReferenceOk returns a tuple with the DetailsReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailsReference

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) SetDetailsReference(v string)`

SetDetailsReference sets DetailsReference field to given value.

### HasDetailsReference

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) HasDetailsReference() bool`

HasDetailsReference returns a boolean if a field has been set.

### SetDetailsReferenceNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) SetDetailsReferenceNil(b bool)`

 SetDetailsReferenceNil sets the value for DetailsReference to be an explicit nil

### UnsetDetailsReference
`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) UnsetDetailsReference()`

UnsetDetailsReference ensures that no value is present for DetailsReference, not even an explicit nil
### GetDepartureAccessLink

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) GetDepartureAccessLink() VTApiPlaneraResaWebV4ModelsJourneysDepartureAccessLinkApiModel`

GetDepartureAccessLink returns the DepartureAccessLink field if non-nil, zero value otherwise.

### GetDepartureAccessLinkOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) GetDepartureAccessLinkOk() (*VTApiPlaneraResaWebV4ModelsJourneysDepartureAccessLinkApiModel, bool)`

GetDepartureAccessLinkOk returns a tuple with the DepartureAccessLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartureAccessLink

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) SetDepartureAccessLink(v VTApiPlaneraResaWebV4ModelsJourneysDepartureAccessLinkApiModel)`

SetDepartureAccessLink sets DepartureAccessLink field to given value.

### HasDepartureAccessLink

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) HasDepartureAccessLink() bool`

HasDepartureAccessLink returns a boolean if a field has been set.

### GetTripLegs

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) GetTripLegs() []VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel`

GetTripLegs returns the TripLegs field if non-nil, zero value otherwise.

### GetTripLegsOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) GetTripLegsOk() (*[]VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel, bool)`

GetTripLegsOk returns a tuple with the TripLegs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTripLegs

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) SetTripLegs(v []VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel)`

SetTripLegs sets TripLegs field to given value.

### HasTripLegs

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) HasTripLegs() bool`

HasTripLegs returns a boolean if a field has been set.

### SetTripLegsNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) SetTripLegsNil(b bool)`

 SetTripLegsNil sets the value for TripLegs to be an explicit nil

### UnsetTripLegs
`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) UnsetTripLegs()`

UnsetTripLegs ensures that no value is present for TripLegs, not even an explicit nil
### GetConnectionLinks

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) GetConnectionLinks() []VTApiPlaneraResaWebV4ModelsJourneysConnectionLinkApiModel`

GetConnectionLinks returns the ConnectionLinks field if non-nil, zero value otherwise.

### GetConnectionLinksOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) GetConnectionLinksOk() (*[]VTApiPlaneraResaWebV4ModelsJourneysConnectionLinkApiModel, bool)`

GetConnectionLinksOk returns a tuple with the ConnectionLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionLinks

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) SetConnectionLinks(v []VTApiPlaneraResaWebV4ModelsJourneysConnectionLinkApiModel)`

SetConnectionLinks sets ConnectionLinks field to given value.

### HasConnectionLinks

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) HasConnectionLinks() bool`

HasConnectionLinks returns a boolean if a field has been set.

### SetConnectionLinksNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) SetConnectionLinksNil(b bool)`

 SetConnectionLinksNil sets the value for ConnectionLinks to be an explicit nil

### UnsetConnectionLinks
`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) UnsetConnectionLinks()`

UnsetConnectionLinks ensures that no value is present for ConnectionLinks, not even an explicit nil
### GetArrivalAccessLink

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) GetArrivalAccessLink() VTApiPlaneraResaWebV4ModelsJourneysArrivalAccessLinkApiModel`

GetArrivalAccessLink returns the ArrivalAccessLink field if non-nil, zero value otherwise.

### GetArrivalAccessLinkOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) GetArrivalAccessLinkOk() (*VTApiPlaneraResaWebV4ModelsJourneysArrivalAccessLinkApiModel, bool)`

GetArrivalAccessLinkOk returns a tuple with the ArrivalAccessLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalAccessLink

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) SetArrivalAccessLink(v VTApiPlaneraResaWebV4ModelsJourneysArrivalAccessLinkApiModel)`

SetArrivalAccessLink sets ArrivalAccessLink field to given value.

### HasArrivalAccessLink

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) HasArrivalAccessLink() bool`

HasArrivalAccessLink returns a boolean if a field has been set.

### GetDestinationLink

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) GetDestinationLink() VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel`

GetDestinationLink returns the DestinationLink field if non-nil, zero value otherwise.

### GetDestinationLinkOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) GetDestinationLinkOk() (*VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel, bool)`

GetDestinationLinkOk returns a tuple with the DestinationLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLink

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) SetDestinationLink(v VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel)`

SetDestinationLink sets DestinationLink field to given value.

### HasDestinationLink

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) HasDestinationLink() bool`

HasDestinationLink returns a boolean if a field has been set.

### GetIsDeparted

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) GetIsDeparted() bool`

GetIsDeparted returns the IsDeparted field if non-nil, zero value otherwise.

### GetIsDepartedOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) GetIsDepartedOk() (*bool, bool)`

GetIsDepartedOk returns a tuple with the IsDeparted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeparted

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) SetIsDeparted(v bool)`

SetIsDeparted sets IsDeparted field to given value.

### HasIsDeparted

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) HasIsDeparted() bool`

HasIsDeparted returns a boolean if a field has been set.

### SetIsDepartedNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) SetIsDepartedNil(b bool)`

 SetIsDepartedNil sets the value for IsDeparted to be an explicit nil

### UnsetIsDeparted
`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) UnsetIsDeparted()`

UnsetIsDeparted ensures that no value is present for IsDeparted, not even an explicit nil
### GetOccupancy

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) GetOccupancy() VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel`

GetOccupancy returns the Occupancy field if non-nil, zero value otherwise.

### GetOccupancyOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) GetOccupancyOk() (*VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel, bool)`

GetOccupancyOk returns a tuple with the Occupancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupancy

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) SetOccupancy(v VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel)`

SetOccupancy sets Occupancy field to given value.

### HasOccupancy

`func (o *VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) HasOccupancy() bool`

HasOccupancy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


