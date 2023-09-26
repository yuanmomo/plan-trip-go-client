# VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gid** | Pointer to **NullableString** | The 16-digit Västtrafik gid of the stop area. | [optional] 
**Name** | Pointer to **NullableString** | The stop area name. | [optional] 
**Latitude** | Pointer to **float64** | The latitude of the stop point. | [optional] 
**Longitude** | Pointer to **float64** | The longitude of the stop point. | [optional] 
**TariffZone1** | Pointer to [**VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel**](VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel.md) |  | [optional] 
**TariffZone2** | Pointer to [**VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel**](VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel.md) |  | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel

`func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel() *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGid

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) GetGid() string`

GetGid returns the Gid field if non-nil, zero value otherwise.

### GetGidOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) GetGidOk() (*string, bool)`

GetGidOk returns a tuple with the Gid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGid

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) SetGid(v string)`

SetGid sets Gid field to given value.

### HasGid

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) HasGid() bool`

HasGid returns a boolean if a field has been set.

### SetGidNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) SetGidNil(b bool)`

 SetGidNil sets the value for Gid to be an explicit nil

### UnsetGid
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) UnsetGid()`

UnsetGid ensures that no value is present for Gid, not even an explicit nil
### GetName

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetLatitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetTariffZone1

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) GetTariffZone1() VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel`

GetTariffZone1 returns the TariffZone1 field if non-nil, zero value otherwise.

### GetTariffZone1Ok

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) GetTariffZone1Ok() (*VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel, bool)`

GetTariffZone1Ok returns a tuple with the TariffZone1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffZone1

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) SetTariffZone1(v VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel)`

SetTariffZone1 sets TariffZone1 field to given value.

### HasTariffZone1

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) HasTariffZone1() bool`

HasTariffZone1 returns a boolean if a field has been set.

### GetTariffZone2

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) GetTariffZone2() VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel`

GetTariffZone2 returns the TariffZone2 field if non-nil, zero value otherwise.

### GetTariffZone2Ok

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) GetTariffZone2Ok() (*VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel, bool)`

GetTariffZone2Ok returns a tuple with the TariffZone2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTariffZone2

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) SetTariffZone2(v VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel)`

SetTariffZone2 sets TariffZone2 field to given value.

### HasTariffZone2

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsStopAreaApiModel) HasTariffZone2() bool`

HasTariffZone2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


