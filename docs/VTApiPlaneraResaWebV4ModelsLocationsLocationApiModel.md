# VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gid** | Pointer to **NullableString** | The 16-digit Västtrafik gid. | [optional] 
**Name** | **string** | The location name. | 
**LocationType** | [**VTApiPlaneraResaCoreModelsLocationType**](VTApiPlaneraResaCoreModelsLocationType.md) |  | 
**Latitude** | Pointer to **NullableFloat64** | The WGS84 latitude of the location. | [optional] 
**Longitude** | Pointer to **NullableFloat64** | The WGS84 longitude of the location. | [optional] 
**Platform** | Pointer to **NullableString** | The location platform, only available for stop points. | [optional] 
**StraightLineDistanceInMeters** | Pointer to **NullableInt32** | The location straight line distance in meters. | [optional] 
**HasLocalService** | Pointer to **NullableBool** | Is \&quot;Närtrafik\&quot; (Local Service) available for the location?  Values are only available for LocationType: StopArea, PointOfInterest and Address.  Values are only available for endpoint: locations/by-text. | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsLocationsLocationApiModel

`func NewVTApiPlaneraResaWebV4ModelsLocationsLocationApiModel(name string, locationType VTApiPlaneraResaCoreModelsLocationType, ) *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel`

NewVTApiPlaneraResaWebV4ModelsLocationsLocationApiModel instantiates a new VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsLocationsLocationApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsLocationsLocationApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel`

NewVTApiPlaneraResaWebV4ModelsLocationsLocationApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGid

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) GetGid() string`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) GetGidOk() (*string, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) SetGid(v string)`

SetGid sets Gid field to given value.

### HasGid

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) HasGid() bool`

HasGid returns a boolean if a field has been set.

### SetGidNil

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) SetGidNil(b bool)`

 SetGidNil sets the value for Gid to be an explicit nil

### UnsetGid
`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) UnsetGid()`

UnsetGid ensures that no value is present for Gid, not even an explicit nil
### GetName

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) SetName(v string)`

SetName sets Name field to given value.


### GetLocationType

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) GetLocationType() VTApiPlaneraResaCoreModelsLocationType`

GetLocationType returns the LocationType field if non-nil, zero value otherwise.

### GetLocationTypeOk

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) GetLocationTypeOk() (*VTApiPlaneraResaCoreModelsLocationType, bool)`

GetLocationTypeOk returns a tuple with the LocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationType

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) SetLocationType(v VTApiPlaneraResaCoreModelsLocationType)`

SetLocationType sets LocationType field to given value.


### GetLatitude

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### SetLatitudeNil

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) SetLatitudeNil(b bool)`

 SetLatitudeNil sets the value for Latitude to be an explicit nil

### UnsetLatitude
`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) UnsetLatitude()`

UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
### GetLongitude

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### SetLongitudeNil

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) SetLongitudeNil(b bool)`

 SetLongitudeNil sets the value for Longitude to be an explicit nil

### UnsetLongitude
`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) UnsetLongitude()`

UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
### GetPlatform

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetStraightLineDistanceInMeters

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) GetStraightLineDistanceInMeters() int32`

GetStraightLineDistanceInMeters returns the StraightLineDistanceInMeters field if non-nil, zero value otherwise.

### GetStraightLineDistanceInMetersOk

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) GetStraightLineDistanceInMetersOk() (*int32, bool)`

GetStraightLineDistanceInMetersOk returns a tuple with the StraightLineDistanceInMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStraightLineDistanceInMeters

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) SetStraightLineDistanceInMeters(v int32)`

SetStraightLineDistanceInMeters sets StraightLineDistanceInMeters field to given value.

### HasStraightLineDistanceInMeters

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) HasStraightLineDistanceInMeters() bool`

HasStraightLineDistanceInMeters returns a boolean if a field has been set.

### SetStraightLineDistanceInMetersNil

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) SetStraightLineDistanceInMetersNil(b bool)`

 SetStraightLineDistanceInMetersNil sets the value for StraightLineDistanceInMeters to be an explicit nil

### UnsetStraightLineDistanceInMeters
`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) UnsetStraightLineDistanceInMeters()`

UnsetStraightLineDistanceInMeters ensures that no value is present for StraightLineDistanceInMeters, not even an explicit nil
### GetHasLocalService

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) GetHasLocalService() bool`

GetHasLocalService returns the HasLocalService field if non-nil, zero value otherwise.

### GetHasLocalServiceOk

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) GetHasLocalServiceOk() (*bool, bool)`

GetHasLocalServiceOk returns a tuple with the HasLocalService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLocalService

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) SetHasLocalService(v bool)`

SetHasLocalService sets HasLocalService field to given value.

### HasHasLocalService

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) HasHasLocalService() bool`

HasHasLocalService returns a boolean if a field has been set.

### SetHasLocalServiceNil

`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) SetHasLocalServiceNil(b bool)`

 SetHasLocalServiceNil sets the value for HasLocalService to be an explicit nil

### UnsetHasLocalService
`func (o *VTApiPlaneraResaWebV4ModelsLocationsLocationApiModel) UnsetHasLocalService()`

UnsetHasLocalService ensures that no value is present for HasLocalService, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


