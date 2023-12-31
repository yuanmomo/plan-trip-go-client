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

// checks if the VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel{}

// VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel struct for VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel
type VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel struct {
	// The 16-digit Västtrafik gid of the tariff zone.
	Gid NullableString `json:"gid,omitempty"`
	// The name of the tariff zone.
	Name NullableString `json:"name,omitempty"`
	// The tariff zone number.
	Number *int32 `json:"number,omitempty"`
	// The short name of the tariff zone.
	ShortName NullableString `json:"shortName,omitempty"`
}

// NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel() *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel {
	this := VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel{}
	return &this
}

// NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel {
	this := VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel{}
	return &this
}

// GetGid returns the Gid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) GetGid() string {
	if o == nil || IsNil(o.Gid.Get()) {
		var ret string
		return ret
	}
	return *o.Gid.Get()
}

// GetGidOk returns a tuple with the Gid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) GetGidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gid.Get(), o.Gid.IsSet()
}

// HasGid returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) HasGid() bool {
	if o != nil && o.Gid.IsSet() {
		return true
	}

	return false
}

// SetGid gets a reference to the given NullableString and assigns it to the Gid field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) SetGid(v string) {
	o.Gid.Set(&v)
}
// SetGidNil sets the value for Gid to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) SetGidNil() {
	o.Gid.Set(nil)
}

// UnsetGid ensures that no value is present for Gid, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) UnsetGid() {
	o.Gid.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) UnsetName() {
	o.Name.Unset()
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) GetNumber() int32 {
	if o == nil || IsNil(o.Number) {
		var ret int32
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) GetNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given int32 and assigns it to the Number field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) SetNumber(v int32) {
	o.Number = &v
}

// GetShortName returns the ShortName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) GetShortName() string {
	if o == nil || IsNil(o.ShortName.Get()) {
		var ret string
		return ret
	}
	return *o.ShortName.Get()
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) GetShortNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShortName.Get(), o.ShortName.IsSet()
}

// HasShortName returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) HasShortName() bool {
	if o != nil && o.ShortName.IsSet() {
		return true
	}

	return false
}

// SetShortName gets a reference to the given NullableString and assigns it to the ShortName field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) SetShortName(v string) {
	o.ShortName.Set(&v)
}
// SetShortNameNil sets the value for ShortName to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) SetShortNameNil() {
	o.ShortName.Set(nil)
}

// UnsetShortName ensures that no value is present for ShortName, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) UnsetShortName() {
	o.ShortName.Unset()
}

func (o VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Gid.IsSet() {
		toSerialize["gid"] = o.Gid.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	if o.ShortName.IsSet() {
		toSerialize["shortName"] = o.ShortName.Get()
	}
	return toSerialize, nil
}

type NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel struct {
	value *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel
	isSet bool
}

func (v NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) Get() *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel {
	return v.value
}

func (v *NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) Set(val *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel(val *VTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) *NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel {
	return &NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel{value: val, isSet: true}
}

func (v NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsTariffZoneApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


