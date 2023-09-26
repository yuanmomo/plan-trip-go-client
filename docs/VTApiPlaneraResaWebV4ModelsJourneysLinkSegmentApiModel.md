# VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | Segment name. | [optional] 
**Maneuver** | Pointer to [**VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver**](VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver.md) |  | [optional] 
**Orientation** | Pointer to [**VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation**](VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation.md) |  | [optional] 
**ManeuverDescription** | Pointer to **NullableString** | Description for the maneuver. | [optional] 
**DistanceInMeters** | Pointer to **NullableInt32** | Distance for this segment in meter. | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel

`func NewVTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel() *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel instantiates a new VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetManeuver

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) GetManeuver() VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver`

GetManeuver returns the Maneuver field if non-nil, zero value otherwise.

### GetManeuverOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) GetManeuverOk() (*VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver, bool)`

GetManeuverOk returns a tuple with the Maneuver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManeuver

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) SetManeuver(v VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver)`

SetManeuver sets Maneuver field to given value.

### HasManeuver

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) HasManeuver() bool`

HasManeuver returns a boolean if a field has been set.

### GetOrientation

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) GetOrientation() VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) GetOrientationOk() (*VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) SetOrientation(v VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation)`

SetOrientation sets Orientation field to given value.

### HasOrientation

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) HasOrientation() bool`

HasOrientation returns a boolean if a field has been set.

### GetManeuverDescription

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) GetManeuverDescription() string`

GetManeuverDescription returns the ManeuverDescription field if non-nil, zero value otherwise.

### GetManeuverDescriptionOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) GetManeuverDescriptionOk() (*string, bool)`

GetManeuverDescriptionOk returns a tuple with the ManeuverDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManeuverDescription

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) SetManeuverDescription(v string)`

SetManeuverDescription sets ManeuverDescription field to given value.

### HasManeuverDescription

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) HasManeuverDescription() bool`

HasManeuverDescription returns a boolean if a field has been set.

### SetManeuverDescriptionNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) SetManeuverDescriptionNil(b bool)`

 SetManeuverDescriptionNil sets the value for ManeuverDescription to be an explicit nil

### UnsetManeuverDescription
`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) UnsetManeuverDescription()`

UnsetManeuverDescription ensures that no value is present for ManeuverDescription, not even an explicit nil
### GetDistanceInMeters

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) GetDistanceInMeters() int32`

GetDistanceInMeters returns the DistanceInMeters field if non-nil, zero value otherwise.

### GetDistanceInMetersOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) GetDistanceInMetersOk() (*int32, bool)`

GetDistanceInMetersOk returns a tuple with the DistanceInMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceInMeters

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) SetDistanceInMeters(v int32)`

SetDistanceInMeters sets DistanceInMeters field to given value.

### HasDistanceInMeters

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) HasDistanceInMeters() bool`

HasDistanceInMeters returns a boolean if a field has been set.

### SetDistanceInMetersNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) SetDistanceInMetersNil(b bool)`

 SetDistanceInMetersNil sets the value for DistanceInMeters to be an explicit nil

### UnsetDistanceInMeters
`func (o *VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel) UnsetDistanceInMeters()`

UnsetDistanceInMeters ensures that no value is present for DistanceInMeters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


