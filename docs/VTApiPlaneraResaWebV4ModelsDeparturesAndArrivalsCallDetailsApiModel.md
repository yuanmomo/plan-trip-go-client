# VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel

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
**Occupancy** | Pointer to [**VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel**](VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel.md) |  | [optional] 
**IsCancelled** | Pointer to **bool** | Flag indicating if the call is cancelled. | [optional] 
**IsDepartureCancelled** | Pointer to **NullableBool** | Flag indicating if the departure is cancelled. | [optional] 
**IsArrivalCancelled** | Pointer to **NullableBool** | Flag indicating if the arrival is cancelled. | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel

`func NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel(stopPoint VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel, ) *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel`

NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel instantiates a new VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel`

NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStopPoint

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetStopPoint() VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel`

GetStopPoint returns the StopPoint field if non-nil, zero value otherwise.

### GetStopPointOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetStopPointOk() (*VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel, bool)`

GetStopPointOk returns a tuple with the StopPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPoint

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetStopPoint(v VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel)`

SetStopPoint sets StopPoint field to given value.


### GetPlannedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetPlannedArrivalTime() string`

GetPlannedArrivalTime returns the PlannedArrivalTime field if non-nil, zero value otherwise.

### GetPlannedArrivalTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetPlannedArrivalTimeOk() (*string, bool)`

GetPlannedArrivalTimeOk returns a tuple with the PlannedArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetPlannedArrivalTime(v string)`

SetPlannedArrivalTime sets PlannedArrivalTime field to given value.

### HasPlannedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) HasPlannedArrivalTime() bool`

HasPlannedArrivalTime returns a boolean if a field has been set.

### SetPlannedArrivalTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetPlannedArrivalTimeNil(b bool)`

 SetPlannedArrivalTimeNil sets the value for PlannedArrivalTime to be an explicit nil

### UnsetPlannedArrivalTime
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) UnsetPlannedArrivalTime()`

UnsetPlannedArrivalTime ensures that no value is present for PlannedArrivalTime, not even an explicit nil
### GetPlannedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetPlannedDepartureTime() string`

GetPlannedDepartureTime returns the PlannedDepartureTime field if non-nil, zero value otherwise.

### GetPlannedDepartureTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetPlannedDepartureTimeOk() (*string, bool)`

GetPlannedDepartureTimeOk returns a tuple with the PlannedDepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetPlannedDepartureTime(v string)`

SetPlannedDepartureTime sets PlannedDepartureTime field to given value.

### HasPlannedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) HasPlannedDepartureTime() bool`

HasPlannedDepartureTime returns a boolean if a field has been set.

### SetPlannedDepartureTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetPlannedDepartureTimeNil(b bool)`

 SetPlannedDepartureTimeNil sets the value for PlannedDepartureTime to be an explicit nil

### UnsetPlannedDepartureTime
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) UnsetPlannedDepartureTime()`

UnsetPlannedDepartureTime ensures that no value is present for PlannedDepartureTime, not even an explicit nil
### GetEstimatedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetEstimatedArrivalTime() string`

GetEstimatedArrivalTime returns the EstimatedArrivalTime field if non-nil, zero value otherwise.

### GetEstimatedArrivalTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetEstimatedArrivalTimeOk() (*string, bool)`

GetEstimatedArrivalTimeOk returns a tuple with the EstimatedArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetEstimatedArrivalTime(v string)`

SetEstimatedArrivalTime sets EstimatedArrivalTime field to given value.

### HasEstimatedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) HasEstimatedArrivalTime() bool`

HasEstimatedArrivalTime returns a boolean if a field has been set.

### SetEstimatedArrivalTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetEstimatedArrivalTimeNil(b bool)`

 SetEstimatedArrivalTimeNil sets the value for EstimatedArrivalTime to be an explicit nil

### UnsetEstimatedArrivalTime
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) UnsetEstimatedArrivalTime()`

UnsetEstimatedArrivalTime ensures that no value is present for EstimatedArrivalTime, not even an explicit nil
### GetEstimatedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetEstimatedDepartureTime() string`

GetEstimatedDepartureTime returns the EstimatedDepartureTime field if non-nil, zero value otherwise.

### GetEstimatedDepartureTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetEstimatedDepartureTimeOk() (*string, bool)`

GetEstimatedDepartureTimeOk returns a tuple with the EstimatedDepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetEstimatedDepartureTime(v string)`

SetEstimatedDepartureTime sets EstimatedDepartureTime field to given value.

### HasEstimatedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) HasEstimatedDepartureTime() bool`

HasEstimatedDepartureTime returns a boolean if a field has been set.

### SetEstimatedDepartureTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetEstimatedDepartureTimeNil(b bool)`

 SetEstimatedDepartureTimeNil sets the value for EstimatedDepartureTime to be an explicit nil

### UnsetEstimatedDepartureTime
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) UnsetEstimatedDepartureTime()`

UnsetEstimatedDepartureTime ensures that no value is present for EstimatedDepartureTime, not even an explicit nil
### GetEstimatedOtherwisePlannedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetEstimatedOtherwisePlannedArrivalTime() string`

GetEstimatedOtherwisePlannedArrivalTime returns the EstimatedOtherwisePlannedArrivalTime field if non-nil, zero value otherwise.

### GetEstimatedOtherwisePlannedArrivalTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetEstimatedOtherwisePlannedArrivalTimeOk() (*string, bool)`

GetEstimatedOtherwisePlannedArrivalTimeOk returns a tuple with the EstimatedOtherwisePlannedArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedOtherwisePlannedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetEstimatedOtherwisePlannedArrivalTime(v string)`

SetEstimatedOtherwisePlannedArrivalTime sets EstimatedOtherwisePlannedArrivalTime field to given value.

### HasEstimatedOtherwisePlannedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) HasEstimatedOtherwisePlannedArrivalTime() bool`

HasEstimatedOtherwisePlannedArrivalTime returns a boolean if a field has been set.

### SetEstimatedOtherwisePlannedArrivalTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetEstimatedOtherwisePlannedArrivalTimeNil(b bool)`

 SetEstimatedOtherwisePlannedArrivalTimeNil sets the value for EstimatedOtherwisePlannedArrivalTime to be an explicit nil

### UnsetEstimatedOtherwisePlannedArrivalTime
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) UnsetEstimatedOtherwisePlannedArrivalTime()`

UnsetEstimatedOtherwisePlannedArrivalTime ensures that no value is present for EstimatedOtherwisePlannedArrivalTime, not even an explicit nil
### GetEstimatedOtherwisePlannedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetEstimatedOtherwisePlannedDepartureTime() string`

GetEstimatedOtherwisePlannedDepartureTime returns the EstimatedOtherwisePlannedDepartureTime field if non-nil, zero value otherwise.

### GetEstimatedOtherwisePlannedDepartureTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetEstimatedOtherwisePlannedDepartureTimeOk() (*string, bool)`

GetEstimatedOtherwisePlannedDepartureTimeOk returns a tuple with the EstimatedOtherwisePlannedDepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedOtherwisePlannedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetEstimatedOtherwisePlannedDepartureTime(v string)`

SetEstimatedOtherwisePlannedDepartureTime sets EstimatedOtherwisePlannedDepartureTime field to given value.

### HasEstimatedOtherwisePlannedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) HasEstimatedOtherwisePlannedDepartureTime() bool`

HasEstimatedOtherwisePlannedDepartureTime returns a boolean if a field has been set.

### SetEstimatedOtherwisePlannedDepartureTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetEstimatedOtherwisePlannedDepartureTimeNil(b bool)`

 SetEstimatedOtherwisePlannedDepartureTimeNil sets the value for EstimatedOtherwisePlannedDepartureTime to be an explicit nil

### UnsetEstimatedOtherwisePlannedDepartureTime
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) UnsetEstimatedOtherwisePlannedDepartureTime()`

UnsetEstimatedOtherwisePlannedDepartureTime ensures that no value is present for EstimatedOtherwisePlannedDepartureTime, not even an explicit nil
### GetPlannedPlatform

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetPlannedPlatform() string`

GetPlannedPlatform returns the PlannedPlatform field if non-nil, zero value otherwise.

### GetPlannedPlatformOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetPlannedPlatformOk() (*string, bool)`

GetPlannedPlatformOk returns a tuple with the PlannedPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedPlatform

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetPlannedPlatform(v string)`

SetPlannedPlatform sets PlannedPlatform field to given value.

### HasPlannedPlatform

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) HasPlannedPlatform() bool`

HasPlannedPlatform returns a boolean if a field has been set.

### SetPlannedPlatformNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetPlannedPlatformNil(b bool)`

 SetPlannedPlatformNil sets the value for PlannedPlatform to be an explicit nil

### UnsetPlannedPlatform
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) UnsetPlannedPlatform()`

UnsetPlannedPlatform ensures that no value is present for PlannedPlatform, not even an explicit nil
### GetEstimatedPlatform

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetEstimatedPlatform() string`

GetEstimatedPlatform returns the EstimatedPlatform field if non-nil, zero value otherwise.

### GetEstimatedPlatformOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetEstimatedPlatformOk() (*string, bool)`

GetEstimatedPlatformOk returns a tuple with the EstimatedPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedPlatform

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetEstimatedPlatform(v string)`

SetEstimatedPlatform sets EstimatedPlatform field to given value.

### HasEstimatedPlatform

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) HasEstimatedPlatform() bool`

HasEstimatedPlatform returns a boolean if a field has been set.

### SetEstimatedPlatformNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetEstimatedPlatformNil(b bool)`

 SetEstimatedPlatformNil sets the value for EstimatedPlatform to be an explicit nil

### UnsetEstimatedPlatform
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) UnsetEstimatedPlatform()`

UnsetEstimatedPlatform ensures that no value is present for EstimatedPlatform, not even an explicit nil
### GetLatitude

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetIndex

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetIndex() string`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetIndexOk() (*string, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetIndex(v string)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### SetIndexNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetIndexNil(b bool)`

 SetIndexNil sets the value for Index to be an explicit nil

### UnsetIndex
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) UnsetIndex()`

UnsetIndex ensures that no value is present for Index, not even an explicit nil
### GetOccupancy

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetOccupancy() VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel`

GetOccupancy returns the Occupancy field if non-nil, zero value otherwise.

### GetOccupancyOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetOccupancyOk() (*VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel, bool)`

GetOccupancyOk returns a tuple with the Occupancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupancy

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetOccupancy(v VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel)`

SetOccupancy sets Occupancy field to given value.

### HasOccupancy

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) HasOccupancy() bool`

HasOccupancy returns a boolean if a field has been set.

### GetIsCancelled

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetIsCancelled() bool`

GetIsCancelled returns the IsCancelled field if non-nil, zero value otherwise.

### GetIsCancelledOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetIsCancelledOk() (*bool, bool)`

GetIsCancelledOk returns a tuple with the IsCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCancelled

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetIsCancelled(v bool)`

SetIsCancelled sets IsCancelled field to given value.

### HasIsCancelled

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) HasIsCancelled() bool`

HasIsCancelled returns a boolean if a field has been set.

### GetIsDepartureCancelled

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetIsDepartureCancelled() bool`

GetIsDepartureCancelled returns the IsDepartureCancelled field if non-nil, zero value otherwise.

### GetIsDepartureCancelledOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetIsDepartureCancelledOk() (*bool, bool)`

GetIsDepartureCancelledOk returns a tuple with the IsDepartureCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDepartureCancelled

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetIsDepartureCancelled(v bool)`

SetIsDepartureCancelled sets IsDepartureCancelled field to given value.

### HasIsDepartureCancelled

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) HasIsDepartureCancelled() bool`

HasIsDepartureCancelled returns a boolean if a field has been set.

### SetIsDepartureCancelledNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetIsDepartureCancelledNil(b bool)`

 SetIsDepartureCancelledNil sets the value for IsDepartureCancelled to be an explicit nil

### UnsetIsDepartureCancelled
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) UnsetIsDepartureCancelled()`

UnsetIsDepartureCancelled ensures that no value is present for IsDepartureCancelled, not even an explicit nil
### GetIsArrivalCancelled

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetIsArrivalCancelled() bool`

GetIsArrivalCancelled returns the IsArrivalCancelled field if non-nil, zero value otherwise.

### GetIsArrivalCancelledOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) GetIsArrivalCancelledOk() (*bool, bool)`

GetIsArrivalCancelledOk returns a tuple with the IsArrivalCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsArrivalCancelled

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetIsArrivalCancelled(v bool)`

SetIsArrivalCancelled sets IsArrivalCancelled field to given value.

### HasIsArrivalCancelled

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) HasIsArrivalCancelled() bool`

HasIsArrivalCancelled returns a boolean if a field has been set.

### SetIsArrivalCancelledNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) SetIsArrivalCancelledNil(b bool)`

 SetIsArrivalCancelledNil sets the value for IsArrivalCancelled to be an explicit nil

### UnsetIsArrivalCancelled
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel) UnsetIsArrivalCancelled()`

UnsetIsArrivalCancelled ensures that no value is present for IsArrivalCancelled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


