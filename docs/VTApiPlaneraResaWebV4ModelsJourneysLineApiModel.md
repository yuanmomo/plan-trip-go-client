# VTApiPlaneraResaWebV4ModelsJourneysLineApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The line name. | [optional] 
**BackgroundColor** | Pointer to **NullableString** | The background color of the line symbol. | [optional] 
**ForegroundColor** | Pointer to **NullableString** | The foreground color of the line symbol. | [optional] 
**BorderColor** | Pointer to **NullableString** | The border color of the line symbol. | [optional] 
**TransportMode** | Pointer to [**VTApiPlaneraResaCoreModelsTransportMode**](VTApiPlaneraResaCoreModelsTransportMode.md) |  | [optional] 
**TransportSubMode** | Pointer to [**VTApiPlaneraResaCoreModelsTransportSubMode**](VTApiPlaneraResaCoreModelsTransportSubMode.md) |  | [optional] 
**ShortName** | Pointer to **NullableString** | The short name of the line, usually 5 characters or less. | [optional] 
**Designation** | Pointer to **NullableString** | The designation of the line. | [optional] 
**IsWheelchairAccessible** | Pointer to **bool** | Flag indicating if the line is wheelchair accessible. | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsJourneysLineApiModel

`func NewVTApiPlaneraResaWebV4ModelsJourneysLineApiModel() *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneysLineApiModel instantiates a new VTApiPlaneraResaWebV4ModelsJourneysLineApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsJourneysLineApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsJourneysLineApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneysLineApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneysLineApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetBackgroundColor

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### SetBackgroundColorNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) SetBackgroundColorNil(b bool)`

 SetBackgroundColorNil sets the value for BackgroundColor to be an explicit nil

### UnsetBackgroundColor
`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) UnsetBackgroundColor()`

UnsetBackgroundColor ensures that no value is present for BackgroundColor, not even an explicit nil
### GetForegroundColor

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) GetForegroundColor() string`

GetForegroundColor returns the ForegroundColor field if non-nil, zero value otherwise.

### GetForegroundColorOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) GetForegroundColorOk() (*string, bool)`

GetForegroundColorOk returns a tuple with the ForegroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForegroundColor

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) SetForegroundColor(v string)`

SetForegroundColor sets ForegroundColor field to given value.

### HasForegroundColor

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) HasForegroundColor() bool`

HasForegroundColor returns a boolean if a field has been set.

### SetForegroundColorNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) SetForegroundColorNil(b bool)`

 SetForegroundColorNil sets the value for ForegroundColor to be an explicit nil

### UnsetForegroundColor
`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) UnsetForegroundColor()`

UnsetForegroundColor ensures that no value is present for ForegroundColor, not even an explicit nil
### GetBorderColor

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) GetBorderColor() string`

GetBorderColor returns the BorderColor field if non-nil, zero value otherwise.

### GetBorderColorOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) GetBorderColorOk() (*string, bool)`

GetBorderColorOk returns a tuple with the BorderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorderColor

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) SetBorderColor(v string)`

SetBorderColor sets BorderColor field to given value.

### HasBorderColor

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) HasBorderColor() bool`

HasBorderColor returns a boolean if a field has been set.

### SetBorderColorNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) SetBorderColorNil(b bool)`

 SetBorderColorNil sets the value for BorderColor to be an explicit nil

### UnsetBorderColor
`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) UnsetBorderColor()`

UnsetBorderColor ensures that no value is present for BorderColor, not even an explicit nil
### GetTransportMode

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) GetTransportMode() VTApiPlaneraResaCoreModelsTransportMode`

GetTransportMode returns the TransportMode field if non-nil, zero value otherwise.

### GetTransportModeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) GetTransportModeOk() (*VTApiPlaneraResaCoreModelsTransportMode, bool)`

GetTransportModeOk returns a tuple with the TransportMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportMode

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) SetTransportMode(v VTApiPlaneraResaCoreModelsTransportMode)`

SetTransportMode sets TransportMode field to given value.

### HasTransportMode

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) HasTransportMode() bool`

HasTransportMode returns a boolean if a field has been set.

### GetTransportSubMode

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) GetTransportSubMode() VTApiPlaneraResaCoreModelsTransportSubMode`

GetTransportSubMode returns the TransportSubMode field if non-nil, zero value otherwise.

### GetTransportSubModeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) GetTransportSubModeOk() (*VTApiPlaneraResaCoreModelsTransportSubMode, bool)`

GetTransportSubModeOk returns a tuple with the TransportSubMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportSubMode

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) SetTransportSubMode(v VTApiPlaneraResaCoreModelsTransportSubMode)`

SetTransportSubMode sets TransportSubMode field to given value.

### HasTransportSubMode

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) HasTransportSubMode() bool`

HasTransportSubMode returns a boolean if a field has been set.

### GetShortName

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### SetShortNameNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) SetShortNameNil(b bool)`

 SetShortNameNil sets the value for ShortName to be an explicit nil

### UnsetShortName
`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) UnsetShortName()`

UnsetShortName ensures that no value is present for ShortName, not even an explicit nil
### GetDesignation

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) GetDesignation() string`

GetDesignation returns the Designation field if non-nil, zero value otherwise.

### GetDesignationOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) GetDesignationOk() (*string, bool)`

GetDesignationOk returns a tuple with the Designation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesignation

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) SetDesignation(v string)`

SetDesignation sets Designation field to given value.

### HasDesignation

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) HasDesignation() bool`

HasDesignation returns a boolean if a field has been set.

### SetDesignationNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) SetDesignationNil(b bool)`

 SetDesignationNil sets the value for Designation to be an explicit nil

### UnsetDesignation
`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) UnsetDesignation()`

UnsetDesignation ensures that no value is present for Designation, not even an explicit nil
### GetIsWheelchairAccessible

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) GetIsWheelchairAccessible() bool`

GetIsWheelchairAccessible returns the IsWheelchairAccessible field if non-nil, zero value otherwise.

### GetIsWheelchairAccessibleOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) GetIsWheelchairAccessibleOk() (*bool, bool)`

GetIsWheelchairAccessibleOk returns a tuple with the IsWheelchairAccessible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWheelchairAccessible

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) SetIsWheelchairAccessible(v bool)`

SetIsWheelchairAccessible sets IsWheelchairAccessible field to given value.

### HasIsWheelchairAccessible

`func (o *VTApiPlaneraResaWebV4ModelsJourneysLineApiModel) HasIsWheelchairAccessible() bool`

HasIsWheelchairAccessible returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


