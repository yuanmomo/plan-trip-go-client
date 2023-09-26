# VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gid** | **string** | 16-digit Västtrafik service journey gid. | 
**Direction** | Pointer to **NullableString** | A description of the direction. | [optional] 
**Number** | Pointer to **NullableString** | Västtrafik service journey number that the trip leg is a part of. | [optional] 
**Line** | Pointer to [**VTApiPlaneraResaWebV4ModelsJourneysLineApiModel**](VTApiPlaneraResaWebV4ModelsJourneysLineApiModel.md) |  | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel

`func NewVTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel(gid string, ) *VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel instantiates a new VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGid

`func (o *VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel) GetGid() string`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel) GetGidOk() (*string, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel) SetGid(v string)`

SetGid sets Gid field to given value.


### GetDirection

`func (o *VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetNumber

`func (o *VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetLine

`func (o *VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel) GetLine() VTApiPlaneraResaWebV4ModelsJourneysLineApiModel`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel) GetLineOk() (*VTApiPlaneraResaWebV4ModelsJourneysLineApiModel, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel) SetLine(v VTApiPlaneraResaWebV4ModelsJourneysLineApiModel)`

SetLine sets Line field to given value.

### HasLine

`func (o *VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel) HasLine() bool`

HasLine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


