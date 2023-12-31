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

// checks if the VTApiPlaneraResaCoreModelsNote type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VTApiPlaneraResaCoreModelsNote{}

// VTApiPlaneraResaCoreModelsNote struct for VTApiPlaneraResaCoreModelsNote
type VTApiPlaneraResaCoreModelsNote struct {
	Type NullableString `json:"type,omitempty"`
	Severity *VTApiPlaneraResaCoreModelsSeverity `json:"severity,omitempty"`
	Text NullableString `json:"text,omitempty"`
}

// NewVTApiPlaneraResaCoreModelsNote instantiates a new VTApiPlaneraResaCoreModelsNote object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVTApiPlaneraResaCoreModelsNote() *VTApiPlaneraResaCoreModelsNote {
	this := VTApiPlaneraResaCoreModelsNote{}
	return &this
}

// NewVTApiPlaneraResaCoreModelsNoteWithDefaults instantiates a new VTApiPlaneraResaCoreModelsNote object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVTApiPlaneraResaCoreModelsNoteWithDefaults() *VTApiPlaneraResaCoreModelsNote {
	this := VTApiPlaneraResaCoreModelsNote{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaCoreModelsNote) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaCoreModelsNote) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *VTApiPlaneraResaCoreModelsNote) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *VTApiPlaneraResaCoreModelsNote) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *VTApiPlaneraResaCoreModelsNote) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *VTApiPlaneraResaCoreModelsNote) UnsetType() {
	o.Type.Unset()
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *VTApiPlaneraResaCoreModelsNote) GetSeverity() VTApiPlaneraResaCoreModelsSeverity {
	if o == nil || IsNil(o.Severity) {
		var ret VTApiPlaneraResaCoreModelsSeverity
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaCoreModelsNote) GetSeverityOk() (*VTApiPlaneraResaCoreModelsSeverity, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *VTApiPlaneraResaCoreModelsNote) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given VTApiPlaneraResaCoreModelsSeverity and assigns it to the Severity field.
func (o *VTApiPlaneraResaCoreModelsNote) SetSeverity(v VTApiPlaneraResaCoreModelsSeverity) {
	o.Severity = &v
}

// GetText returns the Text field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaCoreModelsNote) GetText() string {
	if o == nil || IsNil(o.Text.Get()) {
		var ret string
		return ret
	}
	return *o.Text.Get()
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaCoreModelsNote) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text.Get(), o.Text.IsSet()
}

// HasText returns a boolean if a field has been set.
func (o *VTApiPlaneraResaCoreModelsNote) HasText() bool {
	if o != nil && o.Text.IsSet() {
		return true
	}

	return false
}

// SetText gets a reference to the given NullableString and assigns it to the Text field.
func (o *VTApiPlaneraResaCoreModelsNote) SetText(v string) {
	o.Text.Set(&v)
}
// SetTextNil sets the value for Text to be an explicit nil
func (o *VTApiPlaneraResaCoreModelsNote) SetTextNil() {
	o.Text.Set(nil)
}

// UnsetText ensures that no value is present for Text, not even an explicit nil
func (o *VTApiPlaneraResaCoreModelsNote) UnsetText() {
	o.Text.Unset()
}

func (o VTApiPlaneraResaCoreModelsNote) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VTApiPlaneraResaCoreModelsNote) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if o.Text.IsSet() {
		toSerialize["text"] = o.Text.Get()
	}
	return toSerialize, nil
}

type NullableVTApiPlaneraResaCoreModelsNote struct {
	value *VTApiPlaneraResaCoreModelsNote
	isSet bool
}

func (v NullableVTApiPlaneraResaCoreModelsNote) Get() *VTApiPlaneraResaCoreModelsNote {
	return v.value
}

func (v *NullableVTApiPlaneraResaCoreModelsNote) Set(val *VTApiPlaneraResaCoreModelsNote) {
	v.value = val
	v.isSet = true
}

func (v NullableVTApiPlaneraResaCoreModelsNote) IsSet() bool {
	return v.isSet
}

func (v *NullableVTApiPlaneraResaCoreModelsNote) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVTApiPlaneraResaCoreModelsNote(val *VTApiPlaneraResaCoreModelsNote) *NullableVTApiPlaneraResaCoreModelsNote {
	return &NullableVTApiPlaneraResaCoreModelsNote{value: val, isSet: true}
}

func (v NullableVTApiPlaneraResaCoreModelsNote) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVTApiPlaneraResaCoreModelsNote) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


