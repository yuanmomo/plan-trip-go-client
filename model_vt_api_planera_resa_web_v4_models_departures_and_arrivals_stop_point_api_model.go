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

// checks if the VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel{}

// VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel struct for VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel
type VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel struct {
	// The 16-digit Västtrafik gid of the stop point.
	Gid string `json:"gid"`
	// The stop point name.
	Name string `json:"name"`
	// The platform of the stop point.
	Platform NullableString `json:"platform,omitempty"`
	// The latitude coordinate of the stop point
	Latitude NullableFloat64 `json:"latitude,omitempty"`
	// The logitude coordinate of the stop point
	Longitude NullableFloat64 `json:"longitude,omitempty"`
}

// NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel instantiates a new VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel(gid string, name string) *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel {
	this := VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel{}
	this.Gid = gid
	this.Name = name
	return &this
}

// NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel {
	this := VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel{}
	return &this
}

// GetGid returns the Gid field value
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) GetGid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gid
}

// GetGidOk returns a tuple with the Gid field value
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) GetGidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gid, true
}

// SetGid sets field value
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) SetGid(v string) {
	o.Gid = v
}

// GetName returns the Name field value
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) SetName(v string) {
	o.Name = v
}

// GetPlatform returns the Platform field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) GetPlatform() string {
	if o == nil || IsNil(o.Platform.Get()) {
		var ret string
		return ret
	}
	return *o.Platform.Get()
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) GetPlatformOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Platform.Get(), o.Platform.IsSet()
}

// HasPlatform returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) HasPlatform() bool {
	if o != nil && o.Platform.IsSet() {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given NullableString and assigns it to the Platform field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) SetPlatform(v string) {
	o.Platform.Set(&v)
}
// SetPlatformNil sets the value for Platform to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) SetPlatformNil() {
	o.Platform.Set(nil)
}

// UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) UnsetPlatform() {
	o.Platform.Unset()
}

// GetLatitude returns the Latitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) GetLatitude() float64 {
	if o == nil || IsNil(o.Latitude.Get()) {
		var ret float64
		return ret
	}
	return *o.Latitude.Get()
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) GetLatitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Latitude.Get(), o.Latitude.IsSet()
}

// HasLatitude returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) HasLatitude() bool {
	if o != nil && o.Latitude.IsSet() {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given NullableFloat64 and assigns it to the Latitude field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) SetLatitude(v float64) {
	o.Latitude.Set(&v)
}
// SetLatitudeNil sets the value for Latitude to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) SetLatitudeNil() {
	o.Latitude.Set(nil)
}

// UnsetLatitude ensures that no value is present for Latitude, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) UnsetLatitude() {
	o.Latitude.Unset()
}

// GetLongitude returns the Longitude field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) GetLongitude() float64 {
	if o == nil || IsNil(o.Longitude.Get()) {
		var ret float64
		return ret
	}
	return *o.Longitude.Get()
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) GetLongitudeOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Longitude.Get(), o.Longitude.IsSet()
}

// HasLongitude returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) HasLongitude() bool {
	if o != nil && o.Longitude.IsSet() {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given NullableFloat64 and assigns it to the Longitude field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) SetLongitude(v float64) {
	o.Longitude.Set(&v)
}
// SetLongitudeNil sets the value for Longitude to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) SetLongitudeNil() {
	o.Longitude.Set(nil)
}

// UnsetLongitude ensures that no value is present for Longitude, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) UnsetLongitude() {
	o.Longitude.Unset()
}

func (o VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["gid"] = o.Gid
	toSerialize["name"] = o.Name
	if o.Platform.IsSet() {
		toSerialize["platform"] = o.Platform.Get()
	}
	if o.Latitude.IsSet() {
		toSerialize["latitude"] = o.Latitude.Get()
	}
	if o.Longitude.IsSet() {
		toSerialize["longitude"] = o.Longitude.Get()
	}
	return toSerialize, nil
}

type NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel struct {
	value *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel
	isSet bool
}

func (v NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) Get() *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel {
	return v.value
}

func (v *NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) Set(val *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel(val *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) *NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel {
	return &NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel{value: val, isSet: true}
}

func (v NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

