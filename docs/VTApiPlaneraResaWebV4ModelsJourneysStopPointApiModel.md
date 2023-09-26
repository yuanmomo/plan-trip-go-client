# VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gid** | **string** | The 16-digit Västtrafik gid of the stop point. | 
**Name** | **string** | The stop point name. | 
**Platform** | Pointer to **NullableString** | The platform of the stop point. | [optional] 
**Latitude** | Pointer to **NullableFloat64** | The latitude coordinate of the stop point. | [optional] 
**Longitude** | Pointer to **NullableFloat64** | The logitude coordinate of the stop point. | [optional] 
**StopArea** | Pointer to [**VTApiPlaneraResaWebV4ModelsJourneysStopAreaApiModel**](VTApiPlaneraResaWebV4ModelsJourneysStopAreaApiModel.md) |  | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel

`func NewVTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel(gid string, name string, ) *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel instantiates a new VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsJourneysStopPointApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsJourneysStopPointApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneysStopPointApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGid

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) GetGid() string`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) GetGidOk() (*string, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) SetGid(v string)`

SetGid sets Gid field to given value.


### GetName

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetPlatform

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetLatitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetStopArea

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) GetStopArea() VTApiPlaneraResaWebV4ModelsJourneysStopAreaApiModel`

GetStopArea returns the StopArea field if non-nil, zero value otherwise.

### GetStopAreaOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) GetStopAreaOk() (*VTApiPlaneraResaWebV4ModelsJourneysStopAreaApiModel, bool)`

GetStopAreaOk returns a tuple with the StopArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopArea

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) SetStopArea(v VTApiPlaneraResaWebV4ModelsJourneysStopAreaApiModel)`

SetStopArea sets StopArea field to given value.

### HasStopArea

`func (o *VTApiPlaneraResaWebV4ModelsJourneysStopPointApiModel) HasStopArea() bool`

HasStopArea returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


