# VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gid** | **string** | 16-digit Västtrafik service journey gid. | 
**Origin** | Pointer to **NullableString** | A description of the origin. | [optional] 
**Direction** | Pointer to **NullableString** | A description of the direction. | [optional] 
**Line** | Pointer to [**VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel**](VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel.md) |  | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel

`func NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel(gid string, ) *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel`

NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel instantiates a new VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel`

NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGid

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) GetGid() string`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) GetGidOk() (*string, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) SetGid(v string)`

SetGid sets Gid field to given value.


### GetOrigin

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### SetOriginNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) SetOriginNil(b bool)`

 SetOriginNil sets the value for Origin to be an explicit nil

### UnsetOrigin
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) UnsetOrigin()`

UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
### GetDirection

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### SetDirectionNil

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) SetDirectionNil(b bool)`

 SetDirectionNil sets the value for Direction to be an explicit nil

### UnsetDirection
`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) UnsetDirection()`

UnsetDirection ensures that no value is present for Direction, not even an explicit nil
### GetLine

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) GetLine() VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) GetLineOk() (*VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) SetLine(v VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel)`

SetLine sets Line field to given value.

### HasLine

`func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) HasLine() bool`

HasLine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


