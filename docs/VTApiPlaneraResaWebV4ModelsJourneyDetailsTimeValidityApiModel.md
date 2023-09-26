# VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**VTApiPlaneraResaCoreModelsTimeValidityType**](VTApiPlaneraResaCoreModelsTimeValidityType.md) |  | [optional] 
**Amount** | Pointer to **NullableInt32** | The amount of the unit specified by the Unit property. Always used together with the Unit property. | [optional] 
**Unit** | Pointer to [**VTApiPlaneraResaCoreModelsTimeValidityUnit**](VTApiPlaneraResaCoreModelsTimeValidityUnit.md) |  | [optional] 
**FromDate** | Pointer to **NullableString** | The from date of a date interval specified in RFC 3339 format. Always used together with the  ToDate property. | [optional] 
**ToDate** | Pointer to **NullableString** | The to date of a date interval specified in RFC 3339 format. Always used together with the  FromDate property. | [optional] 
**FromDateTime** | Pointer to **NullableString** | The from time of a datetime interval specified in RFC 3339 format. Always used together with  the ToDateTime property. | [optional] 
**ToDateTime** | Pointer to **NullableString** | The to time of a datetime interval specified in RFC 3339 format. Always used together with  the FromDateTime property. | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel

`func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel() *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetType() VTApiPlaneraResaCoreModelsTimeValidityType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetTypeOk() (*VTApiPlaneraResaCoreModelsTimeValidityType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) SetType(v VTApiPlaneraResaCoreModelsTimeValidityType)`

SetType sets Type field to given value.

### HasType

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAmount

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### SetAmountNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) SetAmountNil(b bool)`

 SetAmountNil sets the value for Amount to be an explicit nil

### UnsetAmount
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) UnsetAmount()`

UnsetAmount ensures that no value is present for Amount, not even an explicit nil
### GetUnit

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetUnit() VTApiPlaneraResaCoreModelsTimeValidityUnit`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetUnitOk() (*VTApiPlaneraResaCoreModelsTimeValidityUnit, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) SetUnit(v VTApiPlaneraResaCoreModelsTimeValidityUnit)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetFromDate

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetFromDate() string`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetFromDateOk() (*string, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) SetFromDate(v string)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### SetFromDateNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) SetFromDateNil(b bool)`

 SetFromDateNil sets the value for FromDate to be an explicit nil

### UnsetFromDate
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) UnsetFromDate()`

UnsetFromDate ensures that no value is present for FromDate, not even an explicit nil
### GetToDate

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetToDate() string`

GetToDate returns the ToDate field if non-nil, zero value otherwise.

### GetToDateOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetToDateOk() (*string, bool)`

GetToDateOk returns a tuple with the ToDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDate

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) SetToDate(v string)`

SetToDate sets ToDate field to given value.

### HasToDate

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) HasToDate() bool`

HasToDate returns a boolean if a field has been set.

### SetToDateNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) SetToDateNil(b bool)`

 SetToDateNil sets the value for ToDate to be an explicit nil

### UnsetToDate
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) UnsetToDate()`

UnsetToDate ensures that no value is present for ToDate, not even an explicit nil
### GetFromDateTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetFromDateTime() string`

GetFromDateTime returns the FromDateTime field if non-nil, zero value otherwise.

### GetFromDateTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetFromDateTimeOk() (*string, bool)`

GetFromDateTimeOk returns a tuple with the FromDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDateTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) SetFromDateTime(v string)`

SetFromDateTime sets FromDateTime field to given value.

### HasFromDateTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) HasFromDateTime() bool`

HasFromDateTime returns a boolean if a field has been set.

### SetFromDateTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) SetFromDateTimeNil(b bool)`

 SetFromDateTimeNil sets the value for FromDateTime to be an explicit nil

### UnsetFromDateTime
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) UnsetFromDateTime()`

UnsetFromDateTime ensures that no value is present for FromDateTime, not even an explicit nil
### GetToDateTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetToDateTime() string`

GetToDateTime returns the ToDateTime field if non-nil, zero value otherwise.

### GetToDateTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetToDateTimeOk() (*string, bool)`

GetToDateTimeOk returns a tuple with the ToDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToDateTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) SetToDateTime(v string)`

SetToDateTime sets ToDateTime field to given value.

### HasToDateTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) HasToDateTime() bool`

HasToDateTime returns a boolean if a field has been set.

### SetToDateTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) SetToDateTimeNil(b bool)`

 SetToDateTimeNil sets the value for ToDateTime to be an explicit nil

### UnsetToDateTime
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) UnsetToDateTime()`

UnsetToDateTime ensures that no value is present for ToDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


