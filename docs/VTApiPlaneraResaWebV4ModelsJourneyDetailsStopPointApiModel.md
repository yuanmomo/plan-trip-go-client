# VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gid** | **string** | The 16-digit Västtrafik gid of the stop point. | 
**Name** | **string** | The stop point name. | 
**Platform** | Pointer to **NullableString** | The platform of the stop point. | [optional] 
**Latitude** | Pointer to **NullableFloat64** | The latitude coordinate of the stop point. | [optional] 
**Longitude** | Pointer to **NullableFloat64** | The longitude coordinate of the stop point. | [optional] 
**StopArea** | Pointer to [**VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel**](VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel.md) |  | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel

`func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel(gid string, name string, ) *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGid

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) GetGid() string`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) GetGidOk() (*string, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) SetGid(v string)`

SetGid sets Gid field to given value.


### GetName

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetPlatform

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetLatitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetStopArea

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) GetStopArea() VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel`

GetStopArea returns the StopArea field if non-nil, zero value otherwise.

### GetStopAreaOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) GetStopAreaOk() (*VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel, bool)`

GetStopAreaOk returns a tuple with the StopArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopArea

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) SetStopArea(v VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel)`

SetStopArea sets StopArea field to given value.

### HasStopArea

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopPointApiModel) HasStopArea() bool`

HasStopArea returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


