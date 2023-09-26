# VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gid** | Pointer to **NullableString** | 16-digit Västtrafik service journey gid. | [optional] 
**Direction** | Pointer to **NullableString** | A description of the direction. | [optional] 
**Line** | Pointer to [**VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineDetailsApiModel**](VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineDetailsApiModel.md) |  | [optional] 
**ServiceJourneyCoordinates** | Pointer to [**[]VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCoordinateApiModel**](VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCoordinateApiModel.md) | The coordinates of the service journey. | [optional] 
**CallsOnServiceJourney** | Pointer to [**[]VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel**](VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel.md) | All calls on the service journey. | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel

`func NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel() *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel`

NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel instantiates a new VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel`

NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGid

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) GetGid() string`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) GetGidOk() (*string, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) SetGid(v string)`

SetGid sets Gid field to given value.

### HasGid

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) HasGid() bool`

HasGid returns a boolean if a field has been set.

### SetGidNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) SetGidNil(b bool)`

 SetGidNil sets the value for Gid to be an explicit nil

### UnsetGid
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) UnsetGid()`

UnsetGid ensures that no value is present for Gid, not even an explicit nil
### GetDirection

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetLine

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) GetLine() VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineDetailsApiModel`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) GetLineOk() (*VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineDetailsApiModel, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) SetLine(v VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineDetailsApiModel)`

SetLine sets Line field to given value.

### HasLine

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetServiceJourneyCoordinates

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) GetServiceJourneyCoordinates() []VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCoordinateApiModel`

GetServiceJourneyCoordinates returns the ServiceJourneyCoordinates field if non-nil, zero value otherwise.

### GetServiceJourneyCoordinatesOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) GetServiceJourneyCoordinatesOk() (*[]VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCoordinateApiModel, bool)`

GetServiceJourneyCoordinatesOk returns a tuple with the ServiceJourneyCoordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceJourneyCoordinates

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) SetServiceJourneyCoordinates(v []VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCoordinateApiModel)`

SetServiceJourneyCoordinates sets ServiceJourneyCoordinates field to given value.

### HasServiceJourneyCoordinates

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) HasServiceJourneyCoordinates() bool`

HasServiceJourneyCoordinates returns a boolean if a field has been set.

### SetServiceJourneyCoordinatesNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) SetServiceJourneyCoordinatesNil(b bool)`

 SetServiceJourneyCoordinatesNil sets the value for ServiceJourneyCoordinates to be an explicit nil

### UnsetServiceJourneyCoordinates
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) UnsetServiceJourneyCoordinates()`

UnsetServiceJourneyCoordinates ensures that no value is present for ServiceJourneyCoordinates, not even an explicit nil
### GetCallsOnServiceJourney

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) GetCallsOnServiceJourney() []VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel`

GetCallsOnServiceJourney returns the CallsOnServiceJourney field if non-nil, zero value otherwise.

### GetCallsOnServiceJourneyOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) GetCallsOnServiceJourneyOk() (*[]VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel, bool)`

GetCallsOnServiceJourneyOk returns a tuple with the CallsOnServiceJourney field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallsOnServiceJourney

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) SetCallsOnServiceJourney(v []VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsCallDetailsApiModel)`

SetCallsOnServiceJourney sets CallsOnServiceJourney field to given value.

### HasCallsOnServiceJourney

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) HasCallsOnServiceJourney() bool`

HasCallsOnServiceJourney returns a boolean if a field has been set.

### SetCallsOnServiceJourneyNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) SetCallsOnServiceJourneyNil(b bool)`

 SetCallsOnServiceJourneyNil sets the value for CallsOnServiceJourney to be an explicit nil

### UnsetCallsOnServiceJourney
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyDetailsApiModel) UnsetCallsOnServiceJourney()`

UnsetCallsOnServiceJourney ensures that no value is present for CallsOnServiceJourney, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


