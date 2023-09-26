# VTApiPlaneraResaWebV4ModelsJourneysCallApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StopPoint** | [**VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel**](VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel.md) |  | 
**PlannedTime** | **string** | The planned time of the call in RFC 3339 format. | 
**EstimatedTime** | Pointer to **NullableString** | The estimated time of the call in RFC 3339 format. | [optional] 
**EstimatedOtherwisePlannedTime** | Pointer to **NullableString** | The best known time of the call in RFC 3339 format. Is EstimatedTime if exists, otherwise PlannedTime. | [optional] [readonly] 
**Notes** | Pointer to [**[]VTApiPlaneraResaCoreModelsNote**](VTApiPlaneraResaCoreModelsNote.md) | An ordered list (most important first) of notes related to the call. | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsJourneysCallApiModel

`func NewVTApiPlaneraResaWebV4ModelsJourneysCallApiModel(stopPoint VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel, plannedTime string, ) *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneysCallApiModel instantiates a new VTApiPlaneraResaWebV4ModelsJourneysCallApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsJourneysCallApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsJourneysCallApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneysCallApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneysCallApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStopPoint

`func (o *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel) GetStopPoint() VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel`

GetStopPoint returns the StopPoint field if non-nil, zero value otherwise.

### GetStopPointOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel) GetStopPointOk() (*VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel, bool)`

GetStopPointOk returns a tuple with the StopPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPoint

`func (o *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel) SetStopPoint(v VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel)`

SetStopPoint sets StopPoint field to given value.


### GetPlannedTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel) GetPlannedTime() string`

GetPlannedTime returns the PlannedTime field if non-nil, zero value otherwise.

### GetPlannedTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel) GetPlannedTimeOk() (*string, bool)`

GetPlannedTimeOk returns a tuple with the PlannedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel) SetPlannedTime(v string)`

SetPlannedTime sets PlannedTime field to given value.


### GetEstimatedTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel) GetEstimatedTime() string`

GetEstimatedTime returns the EstimatedTime field if non-nil, zero value otherwise.

### GetEstimatedTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel) GetEstimatedTimeOk() (*string, bool)`

GetEstimatedTimeOk returns a tuple with the EstimatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel) SetEstimatedTime(v string)`

SetEstimatedTime sets EstimatedTime field to given value.

### HasEstimatedTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel) HasEstimatedTime() bool`

HasEstimatedTime returns a boolean if a field has been set.

### SetEstimatedTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel) SetEstimatedTimeNil(b bool)`

 SetEstimatedTimeNil sets the value for EstimatedTime to be an explicit nil

### UnsetEstimatedTime
`func (o *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel) UnsetEstimatedTime()`

UnsetEstimatedTime ensures that no value is present for EstimatedTime, not even an explicit nil
### GetEstimatedOtherwisePlannedTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel) GetEstimatedOtherwisePlannedTime() string`

GetEstimatedOtherwisePlannedTime returns the EstimatedOtherwisePlannedTime field if non-nil, zero value otherwise.

### GetEstimatedOtherwisePlannedTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel) GetEstimatedOtherwisePlannedTimeOk() (*string, bool)`

GetEstimatedOtherwisePlannedTimeOk returns a tuple with the EstimatedOtherwisePlannedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedOtherwisePlannedTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel) SetEstimatedOtherwisePlannedTime(v string)`

SetEstimatedOtherwisePlannedTime sets EstimatedOtherwisePlannedTime field to given value.

### HasEstimatedOtherwisePlannedTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel) HasEstimatedOtherwisePlannedTime() bool`

HasEstimatedOtherwisePlannedTime returns a boolean if a field has been set.

### SetEstimatedOtherwisePlannedTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel) SetEstimatedOtherwisePlannedTimeNil(b bool)`

 SetEstimatedOtherwisePlannedTimeNil sets the value for EstimatedOtherwisePlannedTime to be an explicit nil

### UnsetEstimatedOtherwisePlannedTime
`func (o *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel) UnsetEstimatedOtherwisePlannedTime()`

UnsetEstimatedOtherwisePlannedTime ensures that no value is present for EstimatedOtherwisePlannedTime, not even an explicit nil
### GetNotes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel) GetNotes() []VTApiPlaneraResaCoreModelsNote`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel) GetNotesOk() (*[]VTApiPlaneraResaCoreModelsNote, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel) SetNotes(v []VTApiPlaneraResaCoreModelsNote)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *VTApiPlaneraResaWebV4ModelsJourneysCallApiModel) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


