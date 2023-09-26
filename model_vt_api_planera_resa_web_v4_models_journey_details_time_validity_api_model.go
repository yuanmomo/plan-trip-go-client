/*
Planera Resa

Sök och planera resor med Västtrafik

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel{}

// VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel Information about the time validity of a ticket suggestion.
type VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel struct {
	Type *VTApiPlaneraResaCoreModelsTimeValidityType `json:"type,omitempty"`
	// The amount of the unit specified by the Unit property. Always used together with the Unit property.
	Amount NullableInt32 `json:"amount,omitempty"`
	Unit *VTApiPlaneraResaCoreModelsTimeValidityUnit `json:"unit,omitempty"`
	// The from date of a date interval specified in RFC 3339 format. Always used together with the  ToDate property.
	FromDate NullableString `json:"fromDate,omitempty"`
	// The to date of a date interval specified in RFC 3339 format. Always used together with the  FromDate property.
	ToDate NullableString `json:"toDate,omitempty"`
	// The from time of a datetime interval specified in RFC 3339 format. Always used together with  the ToDateTime property.
	FromDateTime NullableString `json:"fromDateTime,omitempty"`
	// The to time of a datetime interval specified in RFC 3339 format. Always used together with  the FromDateTime property.
	ToDateTime NullableString `json:"toDateTime,omitempty"`
}

// NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel() *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel {
	this := VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel{}
	return &this
}

// NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel {
	this := VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetType() VTApiPlaneraResaCoreModelsTimeValidityType {
	if o == nil || IsNil(o.Type) {
		var ret VTApiPlaneraResaCoreModelsTimeValidityType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetTypeOk() (*VTApiPlaneraResaCoreModelsTimeValidityType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given VTApiPlaneraResaCoreModelsTimeValidityType and assigns it to the Type field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) SetType(v VTApiPlaneraResaCoreModelsTimeValidityType) {
	o.Type = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetAmount() int32 {
	if o == nil || IsNil(o.Amount.Get()) {
		var ret int32
		return ret
	}
	return *o.Amount.Get()
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Amount.Get(), o.Amount.IsSet()
}

// HasAmount returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) HasAmount() bool {
	if o != nil && o.Amount.IsSet() {
		return true
	}

	return false
}

// SetAmount gets a reference to the given NullableInt32 and assigns it to the Amount field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) SetAmount(v int32) {
	o.Amount.Set(&v)
}
// SetAmountNil sets the value for Amount to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) SetAmountNil() {
	o.Amount.Set(nil)
}

// UnsetAmount ensures that no value is present for Amount, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) UnsetAmount() {
	o.Amount.Unset()
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetUnit() VTApiPlaneraResaCoreModelsTimeValidityUnit {
	if o == nil || IsNil(o.Unit) {
		var ret VTApiPlaneraResaCoreModelsTimeValidityUnit
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetUnitOk() (*VTApiPlaneraResaCoreModelsTimeValidityUnit, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given VTApiPlaneraResaCoreModelsTimeValidityUnit and assigns it to the Unit field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) SetUnit(v VTApiPlaneraResaCoreModelsTimeValidityUnit) {
	o.Unit = &v
}

// GetFromDate returns the FromDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetFromDate() string {
	if o == nil || IsNil(o.FromDate.Get()) {
		var ret string
		return ret
	}
	return *o.FromDate.Get()
}

// GetFromDateOk returns a tuple with the FromDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetFromDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FromDate.Get(), o.FromDate.IsSet()
}

// HasFromDate returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) HasFromDate() bool {
	if o != nil && o.FromDate.IsSet() {
		return true
	}

	return false
}

// SetFromDate gets a reference to the given NullableString and assigns it to the FromDate field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) SetFromDate(v string) {
	o.FromDate.Set(&v)
}
// SetFromDateNil sets the value for FromDate to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) SetFromDateNil() {
	o.FromDate.Set(nil)
}

// UnsetFromDate ensures that no value is present for FromDate, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) UnsetFromDate() {
	o.FromDate.Unset()
}

// GetToDate returns the ToDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetToDate() string {
	if o == nil || IsNil(o.ToDate.Get()) {
		var ret string
		return ret
	}
	return *o.ToDate.Get()
}

// GetToDateOk returns a tuple with the ToDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetToDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ToDate.Get(), o.ToDate.IsSet()
}

// HasToDate returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) HasToDate() bool {
	if o != nil && o.ToDate.IsSet() {
		return true
	}

	return false
}

// SetToDate gets a reference to the given NullableString and assigns it to the ToDate field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) SetToDate(v string) {
	o.ToDate.Set(&v)
}
// SetToDateNil sets the value for ToDate to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) SetToDateNil() {
	o.ToDate.Set(nil)
}

// UnsetToDate ensures that no value is present for ToDate, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) UnsetToDate() {
	o.ToDate.Unset()
}

// GetFromDateTime returns the FromDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetFromDateTime() string {
	if o == nil || IsNil(o.FromDateTime.Get()) {
		var ret string
		return ret
	}
	return *o.FromDateTime.Get()
}

// GetFromDateTimeOk returns a tuple with the FromDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetFromDateTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FromDateTime.Get(), o.FromDateTime.IsSet()
}

// HasFromDateTime returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) HasFromDateTime() bool {
	if o != nil && o.FromDateTime.IsSet() {
		return true
	}

	return false
}

// SetFromDateTime gets a reference to the given NullableString and assigns it to the FromDateTime field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) SetFromDateTime(v string) {
	o.FromDateTime.Set(&v)
}
// SetFromDateTimeNil sets the value for FromDateTime to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) SetFromDateTimeNil() {
	o.FromDateTime.Set(nil)
}

// UnsetFromDateTime ensures that no value is present for FromDateTime, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) UnsetFromDateTime() {
	o.FromDateTime.Unset()
}

// GetToDateTime returns the ToDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetToDateTime() string {
	if o == nil || IsNil(o.ToDateTime.Get()) {
		var ret string
		return ret
	}
	return *o.ToDateTime.Get()
}

// GetToDateTimeOk returns a tuple with the ToDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) GetToDateTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ToDateTime.Get(), o.ToDateTime.IsSet()
}

// HasToDateTime returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) HasToDateTime() bool {
	if o != nil && o.ToDateTime.IsSet() {
		return true
	}

	return false
}

// SetToDateTime gets a reference to the given NullableString and assigns it to the ToDateTime field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) SetToDateTime(v string) {
	o.ToDateTime.Set(&v)
}
// SetToDateTimeNil sets the value for ToDateTime to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) SetToDateTimeNil() {
	o.ToDateTime.Set(nil)
}

// UnsetToDateTime ensures that no value is present for ToDateTime, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) UnsetToDateTime() {
	o.ToDateTime.Unset()
}

func (o VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.Amount.IsSet() {
		toSerialize["amount"] = o.Amount.Get()
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	if o.FromDate.IsSet() {
		toSerialize["fromDate"] = o.FromDate.Get()
	}
	if o.ToDate.IsSet() {
		toSerialize["toDate"] = o.ToDate.Get()
	}
	if o.FromDateTime.IsSet() {
		toSerialize["fromDateTime"] = o.FromDateTime.Get()
	}
	if o.ToDateTime.IsSet() {
		toSerialize["toDateTime"] = o.ToDateTime.Get()
	}
	return toSerialize, nil
}

type NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel struct {
	value *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel
	isSet bool
}

func (v NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) Get() *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel {
	return v.value
}

func (v *NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) Set(val *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel(val *VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) *NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel {
	return &NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel{value: val, isSet: true}
}

func (v NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

