# VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetailsReference** | Pointer to **NullableString** | A reference that should be used when getting detailed information about the journey. | [optional] 
**ServiceJourney** | Pointer to [**VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel**](VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel.md) |  | [optional] 
**StopPoint** | [**VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel**](VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel.md) |  | 
**PlannedTime** | **string** | The planned time of the call in RFC 3339 format. | 
**EstimatedTime** | Pointer to **NullableString** | The estimated time of the call in RFC 3339 format. | [optional] 
**EstimatedOtherwisePlannedTime** | Pointer to **NullableString** | The best known time of the call in RFC 3339 format. Is EstimatedTime if exists, otherwise PlannedTime. | [optional] [readonly] 
**IsCancelled** | Pointer to **bool** | Flag indicating if the departure or arrival is cancelled. | [optional] 
**IsPartCancelled** | Pointer to **bool** | Flag indicating if the departure or arrival is partially cancelled. | [optional] 
**Occupancy** | Pointer to [**VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel**](VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel.md) |  | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel

`func NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel(stopPoint VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel, plannedTime string, ) *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel`

NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel instantiates a new VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel`

NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetailsReference

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetDetailsReference() string`

GetDetailsReference returns the DetailsReference field if non-nil, zero value otherwise.

### GetDetailsReferenceOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetDetailsReferenceOk() (*string, bool)`

GetDetailsReferenceOk returns a tuple with the DetailsReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailsReference

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) SetDetailsReference(v string)`

SetDetailsReference sets DetailsReference field to given value.

### HasDetailsReference

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) HasDetailsReference() bool`

HasDetailsReference returns a boolean if a field has been set.

### SetDetailsReferenceNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) SetDetailsReferenceNil(b bool)`

 SetDetailsReferenceNil sets the value for DetailsReference to be an explicit nil

### UnsetDetailsReference
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) UnsetDetailsReference()`

UnsetDetailsReference ensures that no value is present for DetailsReference, not even an explicit nil
### GetServiceJourney

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetServiceJourney() VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel`

GetServiceJourney returns the ServiceJourney field if non-nil, zero value otherwise.

### GetServiceJourneyOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetServiceJourneyOk() (*VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel, bool)`

GetServiceJourneyOk returns a tuple with the ServiceJourney field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceJourney

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) SetServiceJourney(v VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel)`

SetServiceJourney sets ServiceJourney field to given value.

### HasServiceJourney

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) HasServiceJourney() bool`

HasServiceJourney returns a boolean if a field has been set.

### GetStopPoint

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetStopPoint() VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel`

GetStopPoint returns the StopPoint field if non-nil, zero value otherwise.

### GetStopPointOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetStopPointOk() (*VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel, bool)`

GetStopPointOk returns a tuple with the StopPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPoint

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) SetStopPoint(v VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel)`

SetStopPoint sets StopPoint field to given value.


### GetPlannedTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetPlannedTime() string`

GetPlannedTime returns the PlannedTime field if non-nil, zero value otherwise.

### GetPlannedTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetPlannedTimeOk() (*string, bool)`

GetPlannedTimeOk returns a tuple with the PlannedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) SetPlannedTime(v string)`

SetPlannedTime sets PlannedTime field to given value.


### GetEstimatedTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetEstimatedTime() string`

GetEstimatedTime returns the EstimatedTime field if non-nil, zero value otherwise.

### GetEstimatedTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetEstimatedTimeOk() (*string, bool)`

GetEstimatedTimeOk returns a tuple with the EstimatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) SetEstimatedTime(v string)`

SetEstimatedTime sets EstimatedTime field to given value.

### HasEstimatedTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) HasEstimatedTime() bool`

HasEstimatedTime returns a boolean if a field has been set.

### SetEstimatedTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) SetEstimatedTimeNil(b bool)`

 SetEstimatedTimeNil sets the value for EstimatedTime to be an explicit nil

### UnsetEstimatedTime
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) UnsetEstimatedTime()`

UnsetEstimatedTime ensures that no value is present for EstimatedTime, not even an explicit nil
### GetEstimatedOtherwisePlannedTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetEstimatedOtherwisePlannedTime() string`

GetEstimatedOtherwisePlannedTime returns the EstimatedOtherwisePlannedTime field if non-nil, zero value otherwise.

### GetEstimatedOtherwisePlannedTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetEstimatedOtherwisePlannedTimeOk() (*string, bool)`

GetEstimatedOtherwisePlannedTimeOk returns a tuple with the EstimatedOtherwisePlannedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedOtherwisePlannedTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) SetEstimatedOtherwisePlannedTime(v string)`

SetEstimatedOtherwisePlannedTime sets EstimatedOtherwisePlannedTime field to given value.

### HasEstimatedOtherwisePlannedTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) HasEstimatedOtherwisePlannedTime() bool`

HasEstimatedOtherwisePlannedTime returns a boolean if a field has been set.

### SetEstimatedOtherwisePlannedTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) SetEstimatedOtherwisePlannedTimeNil(b bool)`

 SetEstimatedOtherwisePlannedTimeNil sets the value for EstimatedOtherwisePlannedTime to be an explicit nil

### UnsetEstimatedOtherwisePlannedTime
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) UnsetEstimatedOtherwisePlannedTime()`

UnsetEstimatedOtherwisePlannedTime ensures that no value is present for EstimatedOtherwisePlannedTime, not even an explicit nil
### GetIsCancelled

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetIsCancelled() bool`

GetIsCancelled returns the IsCancelled field if non-nil, zero value otherwise.

### GetIsCancelledOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetIsCancelledOk() (*bool, bool)`

GetIsCancelledOk returns a tuple with the IsCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCancelled

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) SetIsCancelled(v bool)`

SetIsCancelled sets IsCancelled field to given value.

### HasIsCancelled

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) HasIsCancelled() bool`

HasIsCancelled returns a boolean if a field has been set.

### GetIsPartCancelled

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetIsPartCancelled() bool`

GetIsPartCancelled returns the IsPartCancelled field if non-nil, zero value otherwise.

### GetIsPartCancelledOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetIsPartCancelledOk() (*bool, bool)`

GetIsPartCancelledOk returns a tuple with the IsPartCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartCancelled

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) SetIsPartCancelled(v bool)`

SetIsPartCancelled sets IsPartCancelled field to given value.

### HasIsPartCancelled

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) HasIsPartCancelled() bool`

HasIsPartCancelled returns a boolean if a field has been set.

### GetOccupancy

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetOccupancy() VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel`

GetOccupancy returns the Occupancy field if non-nil, zero value otherwise.

### GetOccupancyOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetOccupancyOk() (*VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel, bool)`

GetOccupancyOk returns a tuple with the Occupancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupancy

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) SetOccupancy(v VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel)`

SetOccupancy sets Occupancy field to given value.

### HasOccupancy

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) HasOccupancy() bool`

HasOccupancy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


