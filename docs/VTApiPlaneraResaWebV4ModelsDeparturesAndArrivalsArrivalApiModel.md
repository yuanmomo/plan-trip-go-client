# VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel

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

## Methods

### NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel

`func NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel(stopPoint VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel, plannedTime string, ) *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel`

NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel instantiates a new VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel`

NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetailsReference

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetDetailsReference() string`

GetDetailsReference returns the DetailsReference field if non-nil, zero value otherwise.

### GetDetailsReferenceOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetDetailsReferenceOk() (*string, bool)`

GetDetailsReferenceOk returns a tuple with the DetailsReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailsReference

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) SetDetailsReference(v string)`

SetDetailsReference sets DetailsReference field to given value.

### HasDetailsReference

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) HasDetailsReference() bool`

HasDetailsReference returns a boolean if a field has been set.

### SetDetailsReferenceNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) SetDetailsReferenceNil(b bool)`

 SetDetailsReferenceNil sets the value for DetailsReference to be an explicit nil

### UnsetDetailsReference
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) UnsetDetailsReference()`

UnsetDetailsReference ensures that no value is present for DetailsReference, not even an explicit nil
### GetServiceJourney

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetServiceJourney() VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel`

GetServiceJourney returns the ServiceJourney field if non-nil, zero value otherwise.

### GetServiceJourneyOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetServiceJourneyOk() (*VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel, bool)`

GetServiceJourneyOk returns a tuple with the ServiceJourney field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceJourney

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) SetServiceJourney(v VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel)`

SetServiceJourney sets ServiceJourney field to given value.

### HasServiceJourney

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) HasServiceJourney() bool`

HasServiceJourney returns a boolean if a field has been set.

### GetStopPoint

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetStopPoint() VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel`

GetStopPoint returns the StopPoint field if non-nil, zero value otherwise.

### GetStopPointOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetStopPointOk() (*VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel, bool)`

GetStopPointOk returns a tuple with the StopPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPoint

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) SetStopPoint(v VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel)`

SetStopPoint sets StopPoint field to given value.


### GetPlannedTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetPlannedTime() string`

GetPlannedTime returns the PlannedTime field if non-nil, zero value otherwise.

### GetPlannedTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetPlannedTimeOk() (*string, bool)`

GetPlannedTimeOk returns a tuple with the PlannedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) SetPlannedTime(v string)`

SetPlannedTime sets PlannedTime field to given value.


### GetEstimatedTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetEstimatedTime() string`

GetEstimatedTime returns the EstimatedTime field if non-nil, zero value otherwise.

### GetEstimatedTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetEstimatedTimeOk() (*string, bool)`

GetEstimatedTimeOk returns a tuple with the EstimatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) SetEstimatedTime(v string)`

SetEstimatedTime sets EstimatedTime field to given value.

### HasEstimatedTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) HasEstimatedTime() bool`

HasEstimatedTime returns a boolean if a field has been set.

### SetEstimatedTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) SetEstimatedTimeNil(b bool)`

 SetEstimatedTimeNil sets the value for EstimatedTime to be an explicit nil

### UnsetEstimatedTime
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) UnsetEstimatedTime()`

UnsetEstimatedTime ensures that no value is present for EstimatedTime, not even an explicit nil
### GetEstimatedOtherwisePlannedTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetEstimatedOtherwisePlannedTime() string`

GetEstimatedOtherwisePlannedTime returns the EstimatedOtherwisePlannedTime field if non-nil, zero value otherwise.

### GetEstimatedOtherwisePlannedTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetEstimatedOtherwisePlannedTimeOk() (*string, bool)`

GetEstimatedOtherwisePlannedTimeOk returns a tuple with the EstimatedOtherwisePlannedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedOtherwisePlannedTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) SetEstimatedOtherwisePlannedTime(v string)`

SetEstimatedOtherwisePlannedTime sets EstimatedOtherwisePlannedTime field to given value.

### HasEstimatedOtherwisePlannedTime

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) HasEstimatedOtherwisePlannedTime() bool`

HasEstimatedOtherwisePlannedTime returns a boolean if a field has been set.

### SetEstimatedOtherwisePlannedTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) SetEstimatedOtherwisePlannedTimeNil(b bool)`

 SetEstimatedOtherwisePlannedTimeNil sets the value for EstimatedOtherwisePlannedTime to be an explicit nil

### UnsetEstimatedOtherwisePlannedTime
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) UnsetEstimatedOtherwisePlannedTime()`

UnsetEstimatedOtherwisePlannedTime ensures that no value is present for EstimatedOtherwisePlannedTime, not even an explicit nil
### GetIsCancelled

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetIsCancelled() bool`

GetIsCancelled returns the IsCancelled field if non-nil, zero value otherwise.

### GetIsCancelledOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetIsCancelledOk() (*bool, bool)`

GetIsCancelledOk returns a tuple with the IsCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCancelled

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) SetIsCancelled(v bool)`

SetIsCancelled sets IsCancelled field to given value.

### HasIsCancelled

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) HasIsCancelled() bool`

HasIsCancelled returns a boolean if a field has been set.

### GetIsPartCancelled

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetIsPartCancelled() bool`

GetIsPartCancelled returns the IsPartCancelled field if non-nil, zero value otherwise.

### GetIsPartCancelledOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetIsPartCancelledOk() (*bool, bool)`

GetIsPartCancelledOk returns a tuple with the IsPartCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartCancelled

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) SetIsPartCancelled(v bool)`

SetIsPartCancelled sets IsPartCancelled field to given value.

### HasIsPartCancelled

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) HasIsPartCancelled() bool`

HasIsPartCancelled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


