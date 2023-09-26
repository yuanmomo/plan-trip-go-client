# VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetailsReference** | Pointer to **NullableString** | Journey reference | [optional] 
**Line** | Pointer to [**VTApiPlaneraResaWebV4ModelsPositionsLineDetailsApiModel**](VTApiPlaneraResaWebV4ModelsPositionsLineDetailsApiModel.md) |  | [optional] 
**Notes** | Pointer to [**[]VTApiPlaneraResaCoreModelsNote**](VTApiPlaneraResaCoreModelsNote.md) | Journey notes | [optional] 
**Name** | Pointer to **NullableString** | Journey name | [optional] 
**Direction** | Pointer to **NullableString** | Journey direction | [optional] 
**Latitude** | Pointer to **NullableFloat64** | Current latitude of journey | [optional] 
**Longitude** | Pointer to **NullableFloat64** | Current longitude of journey | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel

`func NewVTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel() *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel`

NewVTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel instantiates a new VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel`

NewVTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetailsReference

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) GetDetailsReference() string`

GetDetailsReference returns the DetailsReference field if non-nil, zero value otherwise.

### GetDetailsReferenceOk

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) GetDetailsReferenceOk() (*string, bool)`

GetDetailsReferenceOk returns a tuple with the DetailsReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailsReference

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) SetDetailsReference(v string)`

SetDetailsReference sets DetailsReference field to given value.

### HasDetailsReference

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) HasDetailsReference() bool`

HasDetailsReference returns a boolean if a field has been set.

### SetDetailsReferenceNil

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) SetDetailsReferenceNil(b bool)`

 SetDetailsReferenceNil sets the value for DetailsReference to be an explicit nil

### UnsetDetailsReference
`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) UnsetDetailsReference()`

UnsetDetailsReference ensures that no value is present for DetailsReference, not even an explicit nil
### GetLine

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) GetLine() VTApiPlaneraResaWebV4ModelsPositionsLineDetailsApiModel`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) GetLineOk() (*VTApiPlaneraResaWebV4ModelsPositionsLineDetailsApiModel, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) SetLine(v VTApiPlaneraResaWebV4ModelsPositionsLineDetailsApiModel)`

SetLine sets Line field to given value.

### HasLine

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetNotes

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) GetNotes() []VTApiPlaneraResaCoreModelsNote`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) GetNotesOk() (*[]VTApiPlaneraResaCoreModelsNote, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) SetNotes(v []VTApiPlaneraResaCoreModelsNote)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetName

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDirection

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetLatitude

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *VTApiPlaneraResaWebV4ModelsPositionsJourneyPositionApiModel) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


