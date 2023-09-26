# VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gid** | Pointer to **NullableString** | 16-digit Västtrafik service journey gid that the trip leg is a part of. | [optional] 
**Direction** | Pointer to **NullableString** | A description of the direction. | [optional] 
**Line** | Pointer to [**VTApiPlaneraResaWebV4ModelsJourneyDetailsLineDetailsApiModel**](VTApiPlaneraResaWebV4ModelsJourneyDetailsLineDetailsApiModel.md) |  | [optional] 
**ServiceJourneyCoordinates** | Pointer to [**[]VTApiPlaneraResaWebV4ModelsCoordinateApiModel**](VTApiPlaneraResaWebV4ModelsCoordinateApiModel.md) | The coordinates on the service journey. | [optional] 
**CallsOnServiceJourney** | Pointer to [**[]VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel**](VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel.md) | All calls on the service journey. | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel

`func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel() *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGid

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) GetGid() string`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) GetGidOk() (*string, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) SetGid(v string)`

SetGid sets Gid field to given value.

### HasGid

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) HasGid() bool`

HasGid returns a boolean if a field has been set.

### SetGidNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) SetGidNil(b bool)`

 SetGidNil sets the value for Gid to be an explicit nil

### UnsetGid
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) UnsetGid()`

UnsetGid ensures that no value is present for Gid, not even an explicit nil
### GetDirection

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetLine

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) GetLine() VTApiPlaneraResaWebV4ModelsJourneyDetailsLineDetailsApiModel`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) GetLineOk() (*VTApiPlaneraResaWebV4ModelsJourneyDetailsLineDetailsApiModel, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) SetLine(v VTApiPlaneraResaWebV4ModelsJourneyDetailsLineDetailsApiModel)`

SetLine sets Line field to given value.

### HasLine

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetServiceJourneyCoordinates

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) GetServiceJourneyCoordinates() []VTApiPlaneraResaWebV4ModelsCoordinateApiModel`

GetServiceJourneyCoordinates returns the ServiceJourneyCoordinates field if non-nil, zero value otherwise.

### GetServiceJourneyCoordinatesOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) GetServiceJourneyCoordinatesOk() (*[]VTApiPlaneraResaWebV4ModelsCoordinateApiModel, bool)`

GetServiceJourneyCoordinatesOk returns a tuple with the ServiceJourneyCoordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceJourneyCoordinates

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) SetServiceJourneyCoordinates(v []VTApiPlaneraResaWebV4ModelsCoordinateApiModel)`

SetServiceJourneyCoordinates sets ServiceJourneyCoordinates field to given value.

### HasServiceJourneyCoordinates

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) HasServiceJourneyCoordinates() bool`

HasServiceJourneyCoordinates returns a boolean if a field has been set.

### SetServiceJourneyCoordinatesNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) SetServiceJourneyCoordinatesNil(b bool)`

 SetServiceJourneyCoordinatesNil sets the value for ServiceJourneyCoordinates to be an explicit nil

### UnsetServiceJourneyCoordinates
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) UnsetServiceJourneyCoordinates()`

UnsetServiceJourneyCoordinates ensures that no value is present for ServiceJourneyCoordinates, not even an explicit nil
### GetCallsOnServiceJourney

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) GetCallsOnServiceJourney() []VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel`

GetCallsOnServiceJourney returns the CallsOnServiceJourney field if non-nil, zero value otherwise.

### GetCallsOnServiceJourneyOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) GetCallsOnServiceJourneyOk() (*[]VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel, bool)`

GetCallsOnServiceJourneyOk returns a tuple with the CallsOnServiceJourney field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallsOnServiceJourney

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) SetCallsOnServiceJourney(v []VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel)`

SetCallsOnServiceJourney sets CallsOnServiceJourney field to given value.

### HasCallsOnServiceJourney

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) HasCallsOnServiceJourney() bool`

HasCallsOnServiceJourney returns a boolean if a field has been set.

### SetCallsOnServiceJourneyNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) SetCallsOnServiceJourneyNil(b bool)`

 SetCallsOnServiceJourneyNil sets the value for CallsOnServiceJourney to be an explicit nil

### UnsetCallsOnServiceJourney
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) UnsetCallsOnServiceJourney()`

UnsetCallsOnServiceJourney ensures that no value is present for CallsOnServiceJourney, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


