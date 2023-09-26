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

// checks if the VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel{}

// VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel Information about a service journey.
type VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel struct {
	// 16-digit Västtrafik service journey gid that the trip leg is a part of.
	Gid NullableString `json:"gid,omitempty"`
	// A description of the direction.
	Direction NullableString `json:"direction,omitempty"`
	Line *VTApiPlaneraResaWebV4ModelsJourneyDetailsLineDetailsApiModel `json:"line,omitempty"`
	// The coordinates on the service journey.
	ServiceJourneyCoordinates []VTApiPlaneraResaWebV4ModelsCoordinateApiModel `json:"serviceJourneyCoordinates,omitempty"`
	// All calls on the service journey.
	CallsOnServiceJourney []VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel `json:"callsOnServiceJourney,omitempty"`
}

// NewVTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel() *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel {
	this := VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel{}
	return &this
}

// NewVTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel {
	this := VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel{}
	return &this
}

// GetGid returns the Gid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) GetGid() string {
	if o == nil || IsNil(o.Gid.Get()) {
		var ret string
		return ret
	}
	return *o.Gid.Get()
}

// GetGidOk returns a tuple with the Gid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) GetGidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Gid.Get(), o.Gid.IsSet()
}

// HasGid returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) HasGid() bool {
	if o != nil && o.Gid.IsSet() {
		return true
	}

	return false
}

// SetGid gets a reference to the given NullableString and assigns it to the Gid field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) SetGid(v string) {
	o.Gid.Set(&v)
}
// SetGidNil sets the value for Gid to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) SetGidNil() {
	o.Gid.Set(nil)
}

// UnsetGid ensures that no value is present for Gid, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) UnsetGid() {
	o.Gid.Unset()
}

// GetDirection returns the Direction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) GetDirection() string {
	if o == nil || IsNil(o.Direction.Get()) {
		var ret string
		return ret
	}
	return *o.Direction.Get()
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) GetDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Direction.Get(), o.Direction.IsSet()
}

// HasDirection returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) HasDirection() bool {
	if o != nil && o.Direction.IsSet() {
		return true
	}

	return false
}

// SetDirection gets a reference to the given NullableString and assigns it to the Direction field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) SetDirection(v string) {
	o.Direction.Set(&v)
}
// SetDirectionNil sets the value for Direction to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) SetDirectionNil() {
	o.Direction.Set(nil)
}

// UnsetDirection ensures that no value is present for Direction, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) UnsetDirection() {
	o.Direction.Unset()
}

// GetLine returns the Line field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) GetLine() VTApiPlaneraResaWebV4ModelsJourneyDetailsLineDetailsApiModel {
	if o == nil || IsNil(o.Line) {
		var ret VTApiPlaneraResaWebV4ModelsJourneyDetailsLineDetailsApiModel
		return ret
	}
	return *o.Line
}

// GetLineOk returns a tuple with the Line field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) GetLineOk() (*VTApiPlaneraResaWebV4ModelsJourneyDetailsLineDetailsApiModel, bool) {
	if o == nil || IsNil(o.Line) {
		return nil, false
	}
	return o.Line, true
}

// HasLine returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) HasLine() bool {
	if o != nil && !IsNil(o.Line) {
		return true
	}

	return false
}

// SetLine gets a reference to the given VTApiPlaneraResaWebV4ModelsJourneyDetailsLineDetailsApiModel and assigns it to the Line field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) SetLine(v VTApiPlaneraResaWebV4ModelsJourneyDetailsLineDetailsApiModel) {
	o.Line = &v
}

// GetServiceJourneyCoordinates returns the ServiceJourneyCoordinates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) GetServiceJourneyCoordinates() []VTApiPlaneraResaWebV4ModelsCoordinateApiModel {
	if o == nil {
		var ret []VTApiPlaneraResaWebV4ModelsCoordinateApiModel
		return ret
	}
	return o.ServiceJourneyCoordinates
}

// GetServiceJourneyCoordinatesOk returns a tuple with the ServiceJourneyCoordinates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) GetServiceJourneyCoordinatesOk() ([]VTApiPlaneraResaWebV4ModelsCoordinateApiModel, bool) {
	if o == nil || IsNil(o.ServiceJourneyCoordinates) {
		return nil, false
	}
	return o.ServiceJourneyCoordinates, true
}

// HasServiceJourneyCoordinates returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) HasServiceJourneyCoordinates() bool {
	if o != nil && IsNil(o.ServiceJourneyCoordinates) {
		return true
	}

	return false
}

// SetServiceJourneyCoordinates gets a reference to the given []VTApiPlaneraResaWebV4ModelsCoordinateApiModel and assigns it to the ServiceJourneyCoordinates field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) SetServiceJourneyCoordinates(v []VTApiPlaneraResaWebV4ModelsCoordinateApiModel) {
	o.ServiceJourneyCoordinates = v
}

// GetCallsOnServiceJourney returns the CallsOnServiceJourney field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) GetCallsOnServiceJourney() []VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel {
	if o == nil {
		var ret []VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel
		return ret
	}
	return o.CallsOnServiceJourney
}

// GetCallsOnServiceJourneyOk returns a tuple with the CallsOnServiceJourney field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) GetCallsOnServiceJourneyOk() ([]VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel, bool) {
	if o == nil || IsNil(o.CallsOnServiceJourney) {
		return nil, false
	}
	return o.CallsOnServiceJourney, true
}

// HasCallsOnServiceJourney returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) HasCallsOnServiceJourney() bool {
	if o != nil && IsNil(o.CallsOnServiceJourney) {
		return true
	}

	return false
}

// SetCallsOnServiceJourney gets a reference to the given []VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel and assigns it to the CallsOnServiceJourney field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) SetCallsOnServiceJourney(v []VTApiPlaneraResaWebV4ModelsJourneyDetailsCallDetailsApiModel) {
	o.CallsOnServiceJourney = v
}

func (o VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Gid.IsSet() {
		toSerialize["gid"] = o.Gid.Get()
	}
	if o.Direction.IsSet() {
		toSerialize["direction"] = o.Direction.Get()
	}
	if !IsNil(o.Line) {
		toSerialize["line"] = o.Line
	}
	if o.ServiceJourneyCoordinates != nil {
		toSerialize["serviceJourneyCoordinates"] = o.ServiceJourneyCoordinates
	}
	if o.CallsOnServiceJourney != nil {
		toSerialize["callsOnServiceJourney"] = o.CallsOnServiceJourney
	}
	return toSerialize, nil
}

type NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel struct {
	value *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel
	isSet bool
}

func (v NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) Get() *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel {
	return v.value
}

func (v *NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) Set(val *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel(val *VTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) *NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel {
	return &NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel{value: val, isSet: true}
}

func (v NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsServiceJourneyApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

