# VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gid** | Pointer to **NullableString** | The 16-digit Västtrafik gid. | [optional] 
**Name** | **string** | The location name. | 
**LocationType** | [**VTApiPlaneraResaCoreModelsLocationType**](VTApiPlaneraResaCoreModelsLocationType.md) |  | 
**Latitude** | Pointer to **NullableFloat64** | The WGS84 latitude of the location. | [optional] 
**Longitude** | Pointer to **NullableFloat64** | The WGS84 longitude of the location. | [optional] 
**PlannedTime** | **string** | The planned time in RFC 3339 format. | 
**EstimatedTime** | Pointer to **NullableString** | The estimated time in RFC 3339 format. | [optional] 
**EstimatedOtherwisePlannedTime** | Pointer to **NullableString** | The best known time of the link in RFC 3339 format. Is EstimatedTime if exists, otherwise PlannedTime. | [optional] [readonly] 
**Notes** | Pointer to [**[]VTApiPlaneraResaCoreModelsNote**](VTApiPlaneraResaCoreModelsNote.md) | An ordered list (most important first) of notes related to the end point. | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel

`func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel(name string, locationType VTApiPlaneraResaCoreModelsLocationType, plannedTime string, ) *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGid

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) GetGid() string`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) GetGidOk() (*string, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) SetGid(v string)`

SetGid sets Gid field to given value.

### HasGid

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) HasGid() bool`

HasGid returns a boolean if a field has been set.

### SetGidNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) SetGidNil(b bool)`

 SetGidNil sets the value for Gid to be an explicit nil

### UnsetGid
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) UnsetGid()`

UnsetGid ensures that no value is present for Gid, not even an explicit nil
### GetName

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetLocationType

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) GetLocationType() VTApiPlaneraResaCoreModelsLocationType`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) GetLocationTypeOk() (*VTApiPlaneraResaCoreModelsLocationType, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) SetLocationType(v VTApiPlaneraResaCoreModelsLocationType)`

SetLocationType sets LocationType field to given value.


### GetLatitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetPlannedTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) GetPlannedTime() string`

GetPlannedTime returns the PlannedTime field if non-nil, zero value otherwise.

### GetPlannedTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) GetPlannedTimeOk() (*string, bool)`

GetPlannedTimeOk returns a tuple with the PlannedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) SetPlannedTime(v string)`

SetPlannedTime sets PlannedTime field to given value.


### GetEstimatedTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) GetEstimatedTime() string`

GetEstimatedTime returns the EstimatedTime field if non-nil, zero value otherwise.

### GetEstimatedTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) GetEstimatedTimeOk() (*string, bool)`

GetEstimatedTimeOk returns a tuple with the EstimatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) SetEstimatedTime(v string)`

SetEstimatedTime sets EstimatedTime field to given value.

### HasEstimatedTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) HasEstimatedTime() bool`

HasEstimatedTime returns a boolean if a field has been set.

### SetEstimatedTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) SetEstimatedTimeNil(b bool)`

 SetEstimatedTimeNil sets the value for EstimatedTime to be an explicit nil

### UnsetEstimatedTime
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) UnsetEstimatedTime()`

UnsetEstimatedTime ensures that no value is present for EstimatedTime, not even an explicit nil
### GetEstimatedOtherwisePlannedTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) GetEstimatedOtherwisePlannedTime() string`

GetEstimatedOtherwisePlannedTime returns the EstimatedOtherwisePlannedTime field if non-nil, zero value otherwise.

### GetEstimatedOtherwisePlannedTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) GetEstimatedOtherwisePlannedTimeOk() (*string, bool)`

GetEstimatedOtherwisePlannedTimeOk returns a tuple with the EstimatedOtherwisePlannedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedOtherwisePlannedTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) SetEstimatedOtherwisePlannedTime(v string)`

SetEstimatedOtherwisePlannedTime sets EstimatedOtherwisePlannedTime field to given value.

### HasEstimatedOtherwisePlannedTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) HasEstimatedOtherwisePlannedTime() bool`

HasEstimatedOtherwisePlannedTime returns a boolean if a field has been set.

### SetEstimatedOtherwisePlannedTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) SetEstimatedOtherwisePlannedTimeNil(b bool)`

 SetEstimatedOtherwisePlannedTimeNil sets the value for EstimatedOtherwisePlannedTime to be an explicit nil

### UnsetEstimatedOtherwisePlannedTime
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) UnsetEstimatedOtherwisePlannedTime()`

UnsetEstimatedOtherwisePlannedTime ensures that no value is present for EstimatedOtherwisePlannedTime, not even an explicit nil
### GetNotes

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) GetNotes() []VTApiPlaneraResaCoreModelsNote`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) GetNotesOk() (*[]VTApiPlaneraResaCoreModelsNote, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) SetNotes(v []VTApiPlaneraResaCoreModelsNote)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkEndpointApiModel) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


