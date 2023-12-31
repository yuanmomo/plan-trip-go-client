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

// checks if the VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel{}

// VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel Information about a line of a departure or arrival service journey.
type VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel struct {
	// 16-digit Västtrafik line gid.
	Gid NullableString `json:"gid,omitempty"`
	// The line name.
	Name NullableString `json:"name,omitempty"`
	// The short name of the line, usually 5 characters or less.
	ShortName NullableString `json:"shortName,omitempty"`
	// The designation of the line.
	Designation NullableString `json:"designation,omitempty"`
	// The background color of the line symbol.
	BackgroundColor NullableString `json:"backgroundColor,omitempty"`
	// The foreground color of the line symbol.
	ForegroundColor NullableString `json:"foregroundColor,omitempty"`
	// The border color of the line symbol.
	BorderColor NullableString `json:"borderColor,omitempty"`
	TransportMode *VTApiPlaneraResaCoreModelsTransportMode `json:"transportMode,omitempty"`
	TransportSubMode *VTApiPlaneraResaCoreModelsTransportSubMode `json:"transportSubMode,omitempty"`
	// Flag indicating if the line is wheelchair accessible.
	IsWheelchairAccessible *bool `json:"isWheelchairAccessible,omitempty"`
}

// NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel instantiates a new VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel() *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel {
	this := VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel{}
	return &this
}

// NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel {
	this := VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel{}
	return &this
}

// GetGid returns the Gid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) GetGid() string {
	if o == nil || IsNil(o.Gid.Get()) {
		var ret string
		return ret
	}
	return *o.Gid.Get()
}

// GetGidOk returns a tuple with the Gid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) GetGidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gid.Get(), o.Gid.IsSet()
}

// HasGid returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) HasGid() bool {
	if o != nil && o.Gid.IsSet() {
		return true
	}

	return false
}

// SetGid gets a reference to the given NullableString and assigns it to the Gid field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) SetGid(v string) {
	o.Gid.Set(&v)
}
// SetGidNil sets the value for Gid to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) SetGidNil() {
	o.Gid.Set(nil)
}

// UnsetGid ensures that no value is present for Gid, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) UnsetGid() {
	o.Gid.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) UnsetName() {
	o.Name.Unset()
}

// GetShortName returns the ShortName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) GetShortName() string {
	if o == nil || IsNil(o.ShortName.Get()) {
		var ret string
		return ret
	}
	return *o.ShortName.Get()
}

// GetShortNameOk returns a tuple with the ShortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) GetShortNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ShortName.Get(), o.ShortName.IsSet()
}

// HasShortName returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) HasShortName() bool {
	if o != nil && o.ShortName.IsSet() {
		return true
	}

	return false
}

// SetShortName gets a reference to the given NullableString and assigns it to the ShortName field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) SetShortName(v string) {
	o.ShortName.Set(&v)
}
// SetShortNameNil sets the value for ShortName to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) SetShortNameNil() {
	o.ShortName.Set(nil)
}

// UnsetShortName ensures that no value is present for ShortName, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) UnsetShortName() {
	o.ShortName.Unset()
}

// GetDesignation returns the Designation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) GetDesignation() string {
	if o == nil || IsNil(o.Designation.Get()) {
		var ret string
		return ret
	}
	return *o.Designation.Get()
}

// GetDesignationOk returns a tuple with the Designation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) GetDesignationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Designation.Get(), o.Designation.IsSet()
}

// HasDesignation returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) HasDesignation() bool {
	if o != nil && o.Designation.IsSet() {
		return true
	}

	return false
}

// SetDesignation gets a reference to the given NullableString and assigns it to the Designation field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) SetDesignation(v string) {
	o.Designation.Set(&v)
}
// SetDesignationNil sets the value for Designation to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) SetDesignationNil() {
	o.Designation.Set(nil)
}

// UnsetDesignation ensures that no value is present for Designation, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) UnsetDesignation() {
	o.Designation.Unset()
}

// GetBackgroundColor returns the BackgroundColor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) GetBackgroundColor() string {
	if o == nil || IsNil(o.BackgroundColor.Get()) {
		var ret string
		return ret
	}
	return *o.BackgroundColor.Get()
}

// GetBackgroundColorOk returns a tuple with the BackgroundColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) GetBackgroundColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BackgroundColor.Get(), o.BackgroundColor.IsSet()
}

// HasBackgroundColor returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) HasBackgroundColor() bool {
	if o != nil && o.BackgroundColor.IsSet() {
		return true
	}

	return false
}

// SetBackgroundColor gets a reference to the given NullableString and assigns it to the BackgroundColor field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) SetBackgroundColor(v string) {
	o.BackgroundColor.Set(&v)
}
// SetBackgroundColorNil sets the value for BackgroundColor to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) SetBackgroundColorNil() {
	o.BackgroundColor.Set(nil)
}

// UnsetBackgroundColor ensures that no value is present for BackgroundColor, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) UnsetBackgroundColor() {
	o.BackgroundColor.Unset()
}

// GetForegroundColor returns the ForegroundColor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) GetForegroundColor() string {
	if o == nil || IsNil(o.ForegroundColor.Get()) {
		var ret string
		return ret
	}
	return *o.ForegroundColor.Get()
}

// GetForegroundColorOk returns a tuple with the ForegroundColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) GetForegroundColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ForegroundColor.Get(), o.ForegroundColor.IsSet()
}

// HasForegroundColor returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) HasForegroundColor() bool {
	if o != nil && o.ForegroundColor.IsSet() {
		return true
	}

	return false
}

// SetForegroundColor gets a reference to the given NullableString and assigns it to the ForegroundColor field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) SetForegroundColor(v string) {
	o.ForegroundColor.Set(&v)
}
// SetForegroundColorNil sets the value for ForegroundColor to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) SetForegroundColorNil() {
	o.ForegroundColor.Set(nil)
}

// UnsetForegroundColor ensures that no value is present for ForegroundColor, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) UnsetForegroundColor() {
	o.ForegroundColor.Unset()
}

// GetBorderColor returns the BorderColor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) GetBorderColor() string {
	if o == nil || IsNil(o.BorderColor.Get()) {
		var ret string
		return ret
	}
	return *o.BorderColor.Get()
}

// GetBorderColorOk returns a tuple with the BorderColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) GetBorderColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.BorderColor.Get(), o.BorderColor.IsSet()
}

// HasBorderColor returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) HasBorderColor() bool {
	if o != nil && o.BorderColor.IsSet() {
		return true
	}

	return false
}

// SetBorderColor gets a reference to the given NullableString and assigns it to the BorderColor field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) SetBorderColor(v string) {
	o.BorderColor.Set(&v)
}
// SetBorderColorNil sets the value for BorderColor to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) SetBorderColorNil() {
	o.BorderColor.Set(nil)
}

// UnsetBorderColor ensures that no value is present for BorderColor, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) UnsetBorderColor() {
	o.BorderColor.Unset()
}

// GetTransportMode returns the TransportMode field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) GetTransportMode() VTApiPlaneraResaCoreModelsTransportMode {
	if o == nil || IsNil(o.TransportMode) {
		var ret VTApiPlaneraResaCoreModelsTransportMode
		return ret
	}
	return *o.TransportMode
}

// GetTransportModeOk returns a tuple with the TransportMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) GetTransportModeOk() (*VTApiPlaneraResaCoreModelsTransportMode, bool) {
	if o == nil || IsNil(o.TransportMode) {
		return nil, false
	}
	return o.TransportMode, true
}

// HasTransportMode returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) HasTransportMode() bool {
	if o != nil && !IsNil(o.TransportMode) {
		return true
	}

	return false
}

// SetTransportMode gets a reference to the given VTApiPlaneraResaCoreModelsTransportMode and assigns it to the TransportMode field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) SetTransportMode(v VTApiPlaneraResaCoreModelsTransportMode) {
	o.TransportMode = &v
}

// GetTransportSubMode returns the TransportSubMode field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) GetTransportSubMode() VTApiPlaneraResaCoreModelsTransportSubMode {
	if o == nil || IsNil(o.TransportSubMode) {
		var ret VTApiPlaneraResaCoreModelsTransportSubMode
		return ret
	}
	return *o.TransportSubMode
}

// GetTransportSubModeOk returns a tuple with the TransportSubMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) GetTransportSubModeOk() (*VTApiPlaneraResaCoreModelsTransportSubMode, bool) {
	if o == nil || IsNil(o.TransportSubMode) {
		return nil, false
	}
	return o.TransportSubMode, true
}

// HasTransportSubMode returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) HasTransportSubMode() bool {
	if o != nil && !IsNil(o.TransportSubMode) {
		return true
	}

	return false
}

// SetTransportSubMode gets a reference to the given VTApiPlaneraResaCoreModelsTransportSubMode and assigns it to the TransportSubMode field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) SetTransportSubMode(v VTApiPlaneraResaCoreModelsTransportSubMode) {
	o.TransportSubMode = &v
}

// GetIsWheelchairAccessible returns the IsWheelchairAccessible field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) GetIsWheelchairAccessible() bool {
	if o == nil || IsNil(o.IsWheelchairAccessible) {
		var ret bool
		return ret
	}
	return *o.IsWheelchairAccessible
}

// GetIsWheelchairAccessibleOk returns a tuple with the IsWheelchairAccessible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) GetIsWheelchairAccessibleOk() (*bool, bool) {
	if o == nil || IsNil(o.IsWheelchairAccessible) {
		return nil, false
	}
	return o.IsWheelchairAccessible, true
}

// HasIsWheelchairAccessible returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) HasIsWheelchairAccessible() bool {
	if o != nil && !IsNil(o.IsWheelchairAccessible) {
		return true
	}

	return false
}

// SetIsWheelchairAccessible gets a reference to the given bool and assigns it to the IsWheelchairAccessible field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) SetIsWheelchairAccessible(v bool) {
	o.IsWheelchairAccessible = &v
}

func (o VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Gid.IsSet() {
		toSerialize["gid"] = o.Gid.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ShortName.IsSet() {
		toSerialize["shortName"] = o.ShortName.Get()
	}
	if o.Designation.IsSet() {
		toSerialize["designation"] = o.Designation.Get()
	}
	if o.BackgroundColor.IsSet() {
		toSerialize["backgroundColor"] = o.BackgroundColor.Get()
	}
	if o.ForegroundColor.IsSet() {
		toSerialize["foregroundColor"] = o.ForegroundColor.Get()
	}
	if o.BorderColor.IsSet() {
		toSerialize["borderColor"] = o.BorderColor.Get()
	}
	if !IsNil(o.TransportMode) {
		toSerialize["transportMode"] = o.TransportMode
	}
	if !IsNil(o.TransportSubMode) {
		toSerialize["transportSubMode"] = o.TransportSubMode
	}
	if !IsNil(o.IsWheelchairAccessible) {
		toSerialize["isWheelchairAccessible"] = o.IsWheelchairAccessible
	}
	return toSerialize, nil
}

type NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel struct {
	value *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel
	isSet bool
}

func (v NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) Get() *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel {
	return v.value
}

func (v *NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) Set(val *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel(val *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) *NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel {
	return &NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel{value: val, isSet: true}
}

func (v NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


