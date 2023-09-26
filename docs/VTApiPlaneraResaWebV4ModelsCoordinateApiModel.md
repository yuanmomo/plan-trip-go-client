# VTApiPlaneraResaWebV4ModelsCoordinateApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Latitude** | Pointer to **float64** | The latitude of this position (WGS84). | [optional] 
**Longitude** | Pointer to **float64** | The longitude of this position (WGS84). | [optional] 
**Elevation** | Pointer to **NullableFloat64** | The elevation of this position (WGS84). | [optional] 
**IsOnTripLeg** | Pointer to **NullableBool** | The coordinate is on the tripleg. | [optional] 
**IsTripLegStart** | Pointer to **NullableBool** | The coordinate is on the first call of the leg. | [optional] 
**IsTripLegStop** | Pointer to **NullableBool** | The coordinate is on the last call of the leg. | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsCoordinateApiModel

`func NewVTApiPlaneraResaWebV4ModelsCoordinateApiModel() *VTApiPlaneraResaWebV4ModelsCoordinateApiModel`

NewVTApiPlaneraResaWebV4ModelsCoordinateApiModel instantiates a new VTApiPlaneraResaWebV4ModelsCoordinateApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsCoordinateApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsCoordinateApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsCoordinateApiModel`

NewVTApiPlaneraResaWebV4ModelsCoordinateApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsCoordinateApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatitude

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) GetLatitude() float64`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) GetLatitudeOk() (*float64, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) SetLatitude(v float64)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) GetLongitude() float64`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) GetLongitudeOk() (*float64, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) SetLongitude(v float64)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetElevation

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) GetElevation() float64`

GetElevation returns the Elevation field if non-nil, zero value otherwise.

### GetElevationOk

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) GetElevationOk() (*float64, bool)`

GetElevationOk returns a tuple with the Elevation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElevation

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) SetElevation(v float64)`

SetElevation sets Elevation field to given value.

### HasElevation

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) HasElevation() bool`

HasElevation returns a boolean if a field has been set.

### SetElevationNil

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) SetElevationNil(b bool)`

 SetElevationNil sets the value for Elevation to be an explicit nil

### UnsetElevation
`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) UnsetElevation()`

UnsetElevation ensures that no value is present for Elevation, not even an explicit nil
### GetIsOnTripLeg

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) GetIsOnTripLeg() bool`

GetIsOnTripLeg returns the IsOnTripLeg field if non-nil, zero value otherwise.

### GetIsOnTripLegOk

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) GetIsOnTripLegOk() (*bool, bool)`

GetIsOnTripLegOk returns a tuple with the IsOnTripLeg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOnTripLeg

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) SetIsOnTripLeg(v bool)`

SetIsOnTripLeg sets IsOnTripLeg field to given value.

### HasIsOnTripLeg

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) HasIsOnTripLeg() bool`

HasIsOnTripLeg returns a boolean if a field has been set.

### SetIsOnTripLegNil

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) SetIsOnTripLegNil(b bool)`

 SetIsOnTripLegNil sets the value for IsOnTripLeg to be an explicit nil

### UnsetIsOnTripLeg
`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) UnsetIsOnTripLeg()`

UnsetIsOnTripLeg ensures that no value is present for IsOnTripLeg, not even an explicit nil
### GetIsTripLegStart

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) GetIsTripLegStart() bool`

GetIsTripLegStart returns the IsTripLegStart field if non-nil, zero value otherwise.

### GetIsTripLegStartOk

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) GetIsTripLegStartOk() (*bool, bool)`

GetIsTripLegStartOk returns a tuple with the IsTripLegStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTripLegStart

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) SetIsTripLegStart(v bool)`

SetIsTripLegStart sets IsTripLegStart field to given value.

### HasIsTripLegStart

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) HasIsTripLegStart() bool`

HasIsTripLegStart returns a boolean if a field has been set.

### SetIsTripLegStartNil

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) SetIsTripLegStartNil(b bool)`

 SetIsTripLegStartNil sets the value for IsTripLegStart to be an explicit nil

### UnsetIsTripLegStart
`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) UnsetIsTripLegStart()`

UnsetIsTripLegStart ensures that no value is present for IsTripLegStart, not even an explicit nil
### GetIsTripLegStop

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) GetIsTripLegStop() bool`

GetIsTripLegStop returns the IsTripLegStop field if non-nil, zero value otherwise.

### GetIsTripLegStopOk

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) GetIsTripLegStopOk() (*bool, bool)`

GetIsTripLegStopOk returns a tuple with the IsTripLegStop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTripLegStop

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) SetIsTripLegStop(v bool)`

SetIsTripLegStop sets IsTripLegStop field to given value.

### HasIsTripLegStop

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) HasIsTripLegStop() bool`

HasIsTripLegStop returns a boolean if a field has been set.

### SetIsTripLegStopNil

`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) SetIsTripLegStopNil(b bool)`

 SetIsTripLegStopNil sets the value for IsTripLegStop to be an explicit nil

### UnsetIsTripLegStop
`func (o *VTApiPlaneraResaWebV4ModelsCoordinateApiModel) UnsetIsTripLegStop()`

UnsetIsTripLegStop ensures that no value is present for IsTripLegStop, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


