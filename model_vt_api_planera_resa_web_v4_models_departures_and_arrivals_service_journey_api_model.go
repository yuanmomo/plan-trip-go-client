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

// checks if the VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel{}

// VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel Information about a service journey of a departure or arrival.
type VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel struct {
	// 16-digit Västtrafik service journey gid.
	Gid string `json:"gid"`
	// A description of the origin.
	Origin NullableString `json:"origin,omitempty"`
	// A description of the direction.
	Direction NullableString `json:"direction,omitempty"`
	Line *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel `json:"line,omitempty"`
}

// NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel instantiates a new VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel(gid string) *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel {
	this := VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel{}
	this.Gid = gid
	return &this
}

// NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel {
	this := VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel{}
	return &this
}

// GetGid returns the Gid field value
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) GetGid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gid
}

// GetGidOk returns a tuple with the Gid field value
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) GetGidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gid, true
}

// SetGid sets field value
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) SetGid(v string) {
	o.Gid = v
}

// GetOrigin returns the Origin field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) GetOrigin() string {
	if o == nil || IsNil(o.Origin.Get()) {
		var ret string
		return ret
	}
	return *o.Origin.Get()
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) GetOriginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Origin.Get(), o.Origin.IsSet()
}

// HasOrigin returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) HasOrigin() bool {
	if o != nil && o.Origin.IsSet() {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given NullableString and assigns it to the Origin field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) SetOrigin(v string) {
	o.Origin.Set(&v)
}
// SetOriginNil sets the value for Origin to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) SetOriginNil() {
	o.Origin.Set(nil)
}

// UnsetOrigin ensures that no value is present for Origin, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) UnsetOrigin() {
	o.Origin.Unset()
}

// GetDirection returns the Direction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) GetDirection() string {
	if o == nil || IsNil(o.Direction.Get()) {
		var ret string
		return ret
	}
	return *o.Direction.Get()
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) GetDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Direction.Get(), o.Direction.IsSet()
}

// HasDirection returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) HasDirection() bool {
	if o != nil && o.Direction.IsSet() {
		return true
	}

	return false
}

// SetDirection gets a reference to the given NullableString and assigns it to the Direction field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) SetDirection(v string) {
	o.Direction.Set(&v)
}
// SetDirectionNil sets the value for Direction to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) SetDirectionNil() {
	o.Direction.Set(nil)
}

// UnsetDirection ensures that no value is present for Direction, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) UnsetDirection() {
	o.Direction.Unset()
}

// GetLine returns the Line field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) GetLine() VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel {
	if o == nil || IsNil(o.Line) {
		var ret VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) GetLineOk() (*VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel, bool) {
	if o == nil || IsNil(o.Line) {
		return nil, false
	}
	return o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) HasLine() bool {
	if o != nil && !IsNil(o.Line) {
		return true
	}

	return false
}

// SetLine gets a reference to the given VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel and assigns it to the Line field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) SetLine(v VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsLineApiModel) {
	o.Line = &v
}

func (o VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gid"] = o.Gid
	if o.Origin.IsSet() {
		toSerialize["origin"] = o.Origin.Get()
	}
	if o.Direction.IsSet() {
		toSerialize["direction"] = o.Direction.Get()
	}
	if !IsNil(o.Line) {
		toSerialize["line"] = o.Line
	}
	return toSerialize, nil
}

type NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel struct {
	value *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel
	isSet bool
}

func (v NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) Get() *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel {
	return v.value
}

func (v *NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) Set(val *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel(val *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) *NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel {
	return &NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel{value: val, isSet: true}
}

func (v NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


