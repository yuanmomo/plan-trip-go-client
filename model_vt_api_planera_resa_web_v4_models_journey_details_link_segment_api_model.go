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

// checks if the VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel{}

// VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel Represents a segment of a departure access link, arrival access link or destination link.
type VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel struct {
	// Segment name.
	Name NullableString `json:"name,omitempty"`
	Maneuver *VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver `json:"maneuver,omitempty"`
	Orientation *VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation `json:"orientation,omitempty"`
	// Description for the maneuver.
	ManeuverDescription NullableString `json:"maneuverDescription,omitempty"`
	// Distance for this segment in meter.
	DistanceInMeters NullableInt32 `json:"distanceInMeters,omitempty"`
}

// NewVTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel() *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel {
	this := VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel{}
	return &this
}

// NewVTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel {
	this := VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) UnsetName() {
	o.Name.Unset()
}

// GetManeuver returns the Maneuver field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) GetManeuver() VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver {
	if o == nil || IsNil(o.Maneuver) {
		var ret VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver
		return ret
	}
	return *o.Maneuver
}

// GetManeuverOk returns a tuple with the Maneuver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) GetManeuverOk() (*VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver, bool) {
	if o == nil || IsNil(o.Maneuver) {
		return nil, false
	}
	return o.Maneuver, true
}

// HasManeuver returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) HasManeuver() bool {
	if o != nil && !IsNil(o.Maneuver) {
		return true
	}

	return false
}

// SetManeuver gets a reference to the given VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver and assigns it to the Maneuver field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) SetManeuver(v VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver) {
	o.Maneuver = &v
}

// GetOrientation returns the Orientation field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) GetOrientation() VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation {
	if o == nil || IsNil(o.Orientation) {
		var ret VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation
		return ret
	}
	return *o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) GetOrientationOk() (*VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation, bool) {
	if o == nil || IsNil(o.Orientation) {
		return nil, false
	}
	return o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) HasOrientation() bool {
	if o != nil && !IsNil(o.Orientation) {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation and assigns it to the Orientation field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) SetOrientation(v VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation) {
	o.Orientation = &v
}

// GetManeuverDescription returns the ManeuverDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) GetManeuverDescription() string {
	if o == nil || IsNil(o.ManeuverDescription.Get()) {
		var ret string
		return ret
	}
	return *o.ManeuverDescription.Get()
}

// GetManeuverDescriptionOk returns a tuple with the ManeuverDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) GetManeuverDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ManeuverDescription.Get(), o.ManeuverDescription.IsSet()
}

// HasManeuverDescription returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) HasManeuverDescription() bool {
	if o != nil && o.ManeuverDescription.IsSet() {
		return true
	}

	return false
}

// SetManeuverDescription gets a reference to the given NullableString and assigns it to the ManeuverDescription field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) SetManeuverDescription(v string) {
	o.ManeuverDescription.Set(&v)
}
// SetManeuverDescriptionNil sets the value for ManeuverDescription to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) SetManeuverDescriptionNil() {
	o.ManeuverDescription.Set(nil)
}

// UnsetManeuverDescription ensures that no value is present for ManeuverDescription, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) UnsetManeuverDescription() {
	o.ManeuverDescription.Unset()
}

// GetDistanceInMeters returns the DistanceInMeters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) GetDistanceInMeters() int32 {
	if o == nil || IsNil(o.DistanceInMeters.Get()) {
		var ret int32
		return ret
	}
	return *o.DistanceInMeters.Get()
}

// GetDistanceInMetersOk returns a tuple with the DistanceInMeters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) GetDistanceInMetersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DistanceInMeters.Get(), o.DistanceInMeters.IsSet()
}

// HasDistanceInMeters returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) HasDistanceInMeters() bool {
	if o != nil && o.DistanceInMeters.IsSet() {
		return true
	}

	return false
}

// SetDistanceInMeters gets a reference to the given NullableInt32 and assigns it to the DistanceInMeters field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) SetDistanceInMeters(v int32) {
	o.DistanceInMeters.Set(&v)
}
// SetDistanceInMetersNil sets the value for DistanceInMeters to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) SetDistanceInMetersNil() {
	o.DistanceInMeters.Set(nil)
}

// UnsetDistanceInMeters ensures that no value is present for DistanceInMeters, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) UnsetDistanceInMeters() {
	o.DistanceInMeters.Unset()
}

func (o VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Maneuver) {
		toSerialize["maneuver"] = o.Maneuver
	}
	if !IsNil(o.Orientation) {
		toSerialize["orientation"] = o.Orientation
	}
	if o.ManeuverDescription.IsSet() {
		toSerialize["maneuverDescription"] = o.ManeuverDescription.Get()
	}
	if o.DistanceInMeters.IsSet() {
		toSerialize["distanceInMeters"] = o.DistanceInMeters.Get()
	}
	return toSerialize, nil
}

type NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel struct {
	value *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel
	isSet bool
}

func (v NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) Get() *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel {
	return v.value
}

func (v *NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) Set(val *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel(val *VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) *NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel {
	return &NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel{value: val, isSet: true}
}

func (v NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


