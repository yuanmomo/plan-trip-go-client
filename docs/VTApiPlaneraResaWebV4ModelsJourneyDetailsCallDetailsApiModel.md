# VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StopPoint** | [**VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel**](VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel.md) |  | 
**PlannedArrivalTime** | Pointer to **NullableString** | The planned arrival time for the call in RFC 3339 format. | [optional] 
**PlannedDepartureTime** | Pointer to **NullableString** | The planned departure time for the call in RFC 3339 format. | [optional] 
**EstimatedArrivalTime** | Pointer to **NullableString** | The estimated arrival time for the call in RFC 3339 format. | [optional] 
**EstimatedDepartureTime** | Pointer to **NullableString** | The estimated departure time for the call in RFC 3339 format. | [optional] 
**EstimatedOtherwisePlannedArrivalTime** | Pointer to **NullableString** | The best known time of the call in RFC 3339 format. Is EstimatedArrivalTime if exists, otherwise PlannedArrivalTime. | [optional] [readonly] 
**EstimatedOtherwisePlannedDepartureTime** | Pointer to **NullableString** | The best known time of the call in RFC 3339 format. Is EstimatedDepartureTime if exists, otherwise PlannedDepartureTime. | [optional] [readonly] 
**PlannedPlatform** | Pointer to **NullableString** | The planned platform of the call. | [optional] 
**EstimatedPlatform** | Pointer to **NullableString** | The estimated platform of the call. | [optional] 
**Latitude** | Pointer to **NullableFloat64** | The latitude of the stop point of the call. | [optional] 
**Longitude** | Pointer to **NullableFloat64** | The longitude of the stop point of the call. | [optional] 
**Index** | Pointer to **NullableString** | The index of the stop point of the call. | [optional] 
**IsOnTripLeg** | Pointer to **NullableBool** | The call is on the trip leg. | [optional] 
**IsTripLegStart** | Pointer to **NullableBool** | The call is the first on the trip leg. | [optional] 
**IsTripLegStop** | Pointer to **NullableBool** | The call is the last on the trip leg. | [optional] 
**TariffZones** | Pointer to [**[]VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel**](VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel.md) | The primary tariff zone of the call. A call can be related to a stop area with multiple tariff zones  and this is the zone that for example should be displayed in overviews etc. | [optional] 
**Occupancy** | Pointer to [**VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel**](VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel.md) |  | [optional] 
**IsCancelled** | Pointer to **bool** | Flag indicating if the call is cancelled. | [optional] 
**IsDepartureCancelled** | Pointer to **NullableBool** | Flag indicating if the departure is cancelled. | [optional] 
**IsArrivalCancelled** | Pointer to **NullableBool** | Flag indicating if the arrival is cancelled. | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel

`func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel(stopPoint VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel, ) *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStopPoint

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetStopPoint() VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel`

GetStopPoint returns the StopPoint field if non-nil, zero value otherwise.

### GetStopPointOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetStopPointOk() (*VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel, bool)`

GetStopPointOk returns a tuple with the StopPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPoint

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetStopPoint(v VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel)`

SetStopPoint sets StopPoint field to given value.


### GetPlannedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetPlannedArrivalTime() string`

GetPlannedArrivalTime returns the PlannedArrivalTime field if non-nil, zero value otherwise.

### GetPlannedArrivalTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetPlannedArrivalTimeOk() (*string, bool)`

GetPlannedArrivalTimeOk returns a tuple with the PlannedArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetPlannedArrivalTime(v string)`

SetPlannedArrivalTime sets PlannedArrivalTime field to given value.

### HasPlannedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) HasPlannedArrivalTime() bool`

HasPlannedArrivalTime returns a boolean if a field has been set.

### SetPlannedArrivalTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetPlannedArrivalTimeNil(b bool)`

 SetPlannedArrivalTimeNil sets the value for PlannedArrivalTime to be an explicit nil

### UnsetPlannedArrivalTime
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) UnsetPlannedArrivalTime()`

UnsetPlannedArrivalTime ensures that no value is present for PlannedArrivalTime, not even an explicit nil
### GetPlannedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetPlannedDepartureTime() string`

GetPlannedDepartureTime returns the PlannedDepartureTime field if non-nil, zero value otherwise.

### GetPlannedDepartureTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetPlannedDepartureTimeOk() (*string, bool)`

GetPlannedDepartureTimeOk returns a tuple with the PlannedDepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetPlannedDepartureTime(v string)`

SetPlannedDepartureTime sets PlannedDepartureTime field to given value.

### HasPlannedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) HasPlannedDepartureTime() bool`

HasPlannedDepartureTime returns a boolean if a field has been set.

### SetPlannedDepartureTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetPlannedDepartureTimeNil(b bool)`

 SetPlannedDepartureTimeNil sets the value for PlannedDepartureTime to be an explicit nil

### UnsetPlannedDepartureTime
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) UnsetPlannedDepartureTime()`

UnsetPlannedDepartureTime ensures that no value is present for PlannedDepartureTime, not even an explicit nil
### GetEstimatedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetEstimatedArrivalTime() string`

GetEstimatedArrivalTime returns the EstimatedArrivalTime field if non-nil, zero value otherwise.

### GetEstimatedArrivalTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetEstimatedArrivalTimeOk() (*string, bool)`

GetEstimatedArrivalTimeOk returns a tuple with the EstimatedArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetEstimatedArrivalTime(v string)`

SetEstimatedArrivalTime sets EstimatedArrivalTime field to given value.

### HasEstimatedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) HasEstimatedArrivalTime() bool`

HasEstimatedArrivalTime returns a boolean if a field has been set.

### SetEstimatedArrivalTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetEstimatedArrivalTimeNil(b bool)`

 SetEstimatedArrivalTimeNil sets the value for EstimatedArrivalTime to be an explicit nil

### UnsetEstimatedArrivalTime
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) UnsetEstimatedArrivalTime()`

UnsetEstimatedArrivalTime ensures that no value is present for EstimatedArrivalTime, not even an explicit nil
### GetEstimatedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetEstimatedDepartureTime() string`

GetEstimatedDepartureTime returns the EstimatedDepartureTime field if non-nil, zero value otherwise.

### GetEstimatedDepartureTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetEstimatedDepartureTimeOk() (*string, bool)`

GetEstimatedDepartureTimeOk returns a tuple with the EstimatedDepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetEstimatedDepartureTime(v string)`

SetEstimatedDepartureTime sets EstimatedDepartureTime field to given value.

### HasEstimatedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) HasEstimatedDepartureTime() bool`

HasEstimatedDepartureTime returns a boolean if a field has been set.

### SetEstimatedDepartureTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetEstimatedDepartureTimeNil(b bool)`

 SetEstimatedDepartureTimeNil sets the value for EstimatedDepartureTime to be an explicit nil

### UnsetEstimatedDepartureTime
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) UnsetEstimatedDepartureTime()`

UnsetEstimatedDepartureTime ensures that no value is present for EstimatedDepartureTime, not even an explicit nil
### GetEstimatedOtherwisePlannedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetEstimatedOtherwisePlannedArrivalTime() string`

GetEstimatedOtherwisePlannedArrivalTime returns the EstimatedOtherwisePlannedArrivalTime field if non-nil, zero value otherwise.

### GetEstimatedOtherwisePlannedArrivalTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetEstimatedOtherwisePlannedArrivalTimeOk() (*string, bool)`

GetEstimatedOtherwisePlannedArrivalTimeOk returns a tuple with the EstimatedOtherwisePlannedArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedOtherwisePlannedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetEstimatedOtherwisePlannedArrivalTime(v string)`

SetEstimatedOtherwisePlannedArrivalTime sets EstimatedOtherwisePlannedArrivalTime field to given value.

### HasEstimatedOtherwisePlannedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) HasEstimatedOtherwisePlannedArrivalTime() bool`

HasEstimatedOtherwisePlannedArrivalTime returns a boolean if a field has been set.

### SetEstimatedOtherwisePlannedArrivalTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetEstimatedOtherwisePlannedArrivalTimeNil(b bool)`

 SetEstimatedOtherwisePlannedArrivalTimeNil sets the value for EstimatedOtherwisePlannedArrivalTime to be an explicit nil

### UnsetEstimatedOtherwisePlannedArrivalTime
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) UnsetEstimatedOtherwisePlannedArrivalTime()`

UnsetEstimatedOtherwisePlannedArrivalTime ensures that no value is present for EstimatedOtherwisePlannedArrivalTime, not even an explicit nil
### GetEstimatedOtherwisePlannedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetEstimatedOtherwisePlannedDepartureTime() string`

GetEstimatedOtherwisePlannedDepartureTime returns the EstimatedOtherwisePlannedDepartureTime field if non-nil, zero value otherwise.

### GetEstimatedOtherwisePlannedDepartureTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetEstimatedOtherwisePlannedDepartureTimeOk() (*string, bool)`

GetEstimatedOtherwisePlannedDepartureTimeOk returns a tuple with the EstimatedOtherwisePlannedDepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedOtherwisePlannedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetEstimatedOtherwisePlannedDepartureTime(v string)`

SetEstimatedOtherwisePlannedDepartureTime sets EstimatedOtherwisePlannedDepartureTime field to given value.

### HasEstimatedOtherwisePlannedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) HasEstimatedOtherwisePlannedDepartureTime() bool`

HasEstimatedOtherwisePlannedDepartureTime returns a boolean if a field has been set.

### SetEstimatedOtherwisePlannedDepartureTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetEstimatedOtherwisePlannedDepartureTimeNil(b bool)`

 SetEstimatedOtherwisePlannedDepartureTimeNil sets the value for EstimatedOtherwisePlannedDepartureTime to be an explicit nil

### UnsetEstimatedOtherwisePlannedDepartureTime
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) UnsetEstimatedOtherwisePlannedDepartureTime()`

UnsetEstimatedOtherwisePlannedDepartureTime ensures that no value is present for EstimatedOtherwisePlannedDepartureTime, not even an explicit nil
### GetPlannedPlatform

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetPlannedPlatform() string`

GetPlannedPlatform returns the PlannedPlatform field if non-nil, zero value otherwise.

### GetPlannedPlatformOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetPlannedPlatformOk() (*string, bool)`

GetPlannedPlatformOk returns a tuple with the PlannedPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedPlatform

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetPlannedPlatform(v string)`

SetPlannedPlatform sets PlannedPlatform field to given value.

### HasPlannedPlatform

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) HasPlannedPlatform() bool`

HasPlannedPlatform returns a boolean if a field has been set.

### SetPlannedPlatformNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetPlannedPlatformNil(b bool)`

 SetPlannedPlatformNil sets the value for PlannedPlatform to be an explicit nil

### UnsetPlannedPlatform
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) UnsetPlannedPlatform()`

UnsetPlannedPlatform ensures that no value is present for PlannedPlatform, not even an explicit nil
### GetEstimatedPlatform

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetEstimatedPlatform() string`

GetEstimatedPlatform returns the EstimatedPlatform field if non-nil, zero value otherwise.

### GetEstimatedPlatformOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetEstimatedPlatformOk() (*string, bool)`

GetEstimatedPlatformOk returns a tuple with the EstimatedPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedPlatform

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetEstimatedPlatform(v string)`

SetEstimatedPlatform sets EstimatedPlatform field to given value.

### HasEstimatedPlatform

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) HasEstimatedPlatform() bool`

HasEstimatedPlatform returns a boolean if a field has been set.

### SetEstimatedPlatformNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetEstimatedPlatformNil(b bool)`

 SetEstimatedPlatformNil sets the value for EstimatedPlatform to be an explicit nil

### UnsetEstimatedPlatform
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) UnsetEstimatedPlatform()`

UnsetEstimatedPlatform ensures that no value is present for EstimatedPlatform, not even an explicit nil
### GetLatitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetIndex

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### SetIndexNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetIndexNil(b bool)`

 SetIndexNil sets the value for Index to be an explicit nil

### UnsetIndex
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) UnsetIndex()`

UnsetIndex ensures that no value is present for Index, not even an explicit nil
### GetIsOnTripLeg

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetIsOnTripLeg() bool`

GetIsOnTripLeg returns the IsOnTripLeg field if non-nil, zero value otherwise.

### GetIsOnTripLegOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetIsOnTripLegOk() (*bool, bool)`

GetIsOnTripLegOk returns a tuple with the IsOnTripLeg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnTripLeg

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetIsOnTripLeg(v bool)`

SetIsOnTripLeg sets IsOnTripLeg field to given value.

### HasIsOnTripLeg

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) HasIsOnTripLeg() bool`

HasIsOnTripLeg returns a boolean if a field has been set.

### SetIsOnTripLegNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetIsOnTripLegNil(b bool)`

 SetIsOnTripLegNil sets the value for IsOnTripLeg to be an explicit nil

### UnsetIsOnTripLeg
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) UnsetIsOnTripLeg()`

UnsetIsOnTripLeg ensures that no value is present for IsOnTripLeg, not even an explicit nil
### GetIsTripLegStart

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetIsTripLegStart() bool`

GetIsTripLegStart returns the IsTripLegStart field if non-nil, zero value otherwise.

### GetIsTripLegStartOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetIsTripLegStartOk() (*bool, bool)`

GetIsTripLegStartOk returns a tuple with the IsTripLegStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTripLegStart

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetIsTripLegStart(v bool)`

SetIsTripLegStart sets IsTripLegStart field to given value.

### HasIsTripLegStart

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) HasIsTripLegStart() bool`

HasIsTripLegStart returns a boolean if a field has been set.

### SetIsTripLegStartNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetIsTripLegStartNil(b bool)`

 SetIsTripLegStartNil sets the value for IsTripLegStart to be an explicit nil

### UnsetIsTripLegStart
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) UnsetIsTripLegStart()`

UnsetIsTripLegStart ensures that no value is present for IsTripLegStart, not even an explicit nil
### GetIsTripLegStop

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetIsTripLegStop() bool`

GetIsTripLegStop returns the IsTripLegStop field if non-nil, zero value otherwise.

### GetIsTripLegStopOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetIsTripLegStopOk() (*bool, bool)`

GetIsTripLegStopOk returns a tuple with the IsTripLegStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTripLegStop

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetIsTripLegStop(v bool)`

SetIsTripLegStop sets IsTripLegStop field to given value.

### HasIsTripLegStop

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) HasIsTripLegStop() bool`

HasIsTripLegStop returns a boolean if a field has been set.

### SetIsTripLegStopNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetIsTripLegStopNil(b bool)`

 SetIsTripLegStopNil sets the value for IsTripLegStop to be an explicit nil

### UnsetIsTripLegStop
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) UnsetIsTripLegStop()`

UnsetIsTripLegStop ensures that no value is present for IsTripLegStop, not even an explicit nil
### GetTariffZones

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetTariffZones() []VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel`

GetTariffZones returns the TariffZones field if non-nil, zero value otherwise.

### GetTariffZonesOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetTariffZonesOk() (*[]VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel, bool)`

GetTariffZonesOk returns a tuple with the TariffZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffZones

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetTariffZones(v []VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel)`

SetTariffZones sets TariffZones field to given value.

### HasTariffZones

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) HasTariffZones() bool`

HasTariffZones returns a boolean if a field has been set.

### SetTariffZonesNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetTariffZonesNil(b bool)`

 SetTariffZonesNil sets the value for TariffZones to be an explicit nil

### UnsetTariffZones
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) UnsetTariffZones()`

UnsetTariffZones ensures that no value is present for TariffZones, not even an explicit nil
### GetOccupancy

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetOccupancy() VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel`

GetOccupancy returns the Occupancy field if non-nil, zero value otherwise.

### GetOccupancyOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetOccupancyOk() (*VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel, bool)`

GetOccupancyOk returns a tuple with the Occupancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupancy

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetOccupancy(v VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel)`

SetOccupancy sets Occupancy field to given value.

### HasOccupancy

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) HasOccupancy() bool`

HasOccupancy returns a boolean if a field has been set.

### GetIsCancelled

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetIsCancelled() bool`

GetIsCancelled returns the IsCancelled field if non-nil, zero value otherwise.

### GetIsCancelledOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetIsCancelledOk() (*bool, bool)`

GetIsCancelledOk returns a tuple with the IsCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCancelled

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetIsCancelled(v bool)`

SetIsCancelled sets IsCancelled field to given value.

### HasIsCancelled

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) HasIsCancelled() bool`

HasIsCancelled returns a boolean if a field has been set.

### GetIsDepartureCancelled

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetIsDepartureCancelled() bool`

GetIsDepartureCancelled returns the IsDepartureCancelled field if non-nil, zero value otherwise.

### GetIsDepartureCancelledOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetIsDepartureCancelledOk() (*bool, bool)`

GetIsDepartureCancelledOk returns a tuple with the IsDepartureCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDepartureCancelled

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetIsDepartureCancelled(v bool)`

SetIsDepartureCancelled sets IsDepartureCancelled field to given value.

### HasIsDepartureCancelled

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) HasIsDepartureCancelled() bool`

HasIsDepartureCancelled returns a boolean if a field has been set.

### SetIsDepartureCancelledNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetIsDepartureCancelledNil(b bool)`

 SetIsDepartureCancelledNil sets the value for IsDepartureCancelled to be an explicit nil

### UnsetIsDepartureCancelled
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) UnsetIsDepartureCancelled()`

UnsetIsDepartureCancelled ensures that no value is present for IsDepartureCancelled, not even an explicit nil
### GetIsArrivalCancelled

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetIsArrivalCancelled() bool`

GetIsArrivalCancelled returns the IsArrivalCancelled field if non-nil, zero value otherwise.

### GetIsArrivalCancelledOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) GetIsArrivalCancelledOk() (*bool, bool)`

GetIsArrivalCancelledOk returns a tuple with the IsArrivalCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArrivalCancelled

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetIsArrivalCancelled(v bool)`

SetIsArrivalCancelled sets IsArrivalCancelled field to given value.

### HasIsArrivalCancelled

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) HasIsArrivalCancelled() bool`

HasIsArrivalCancelled returns a boolean if a field has been set.

### SetIsArrivalCancelledNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) SetIsArrivalCancelledNil(b bool)`

 SetIsArrivalCancelledNil sets the value for IsArrivalCancelled to be an explicit nil

### UnsetIsArrivalCancelled
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) UnsetIsArrivalCancelled()`

UnsetIsArrivalCancelled ensures that no value is present for IsArrivalCancelled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


