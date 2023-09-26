/*
Planera Resa

Sök och planera resor med Västtrafik

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation Specifies an orientation of a link segment.
type VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation string

// List of VT.ApiPlaneraResa.Web.V4.Models.LinkSegmentOrientation
const (
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTORIENTATION_UNKNOWN VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation = "unknown"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTORIENTATION_NORTH VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation = "north"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTORIENTATION_SOUTH VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation = "south"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTORIENTATION_EAST VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation = "east"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTORIENTATION_WEST VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation = "west"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTORIENTATION_NORTHEAST VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation = "northeast"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTORIENTATION_SOUTHEAST VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation = "southeast"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTORIENTATION_NORTHWEST VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation = "northwest"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTORIENTATION_SOUTHWEST VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation = "southwest"
)

// All allowed values of VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation enum
var AllowedVTApiPlaneraResaWebV4ModelsLinkSegmentOrientationEnumValues = []VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation{
	"unknown",
	"north",
	"south",
	"east",
	"west",
	"northeast",
	"southeast",
	"northwest",
	"southwest",
}

func (v *VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation(value)
	for _, existing := range AllowedVTApiPlaneraResaWebV4ModelsLinkSegmentOrientationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation", value)
}

// NewVTApiPlaneraResaWebV4ModelsLinkSegmentOrientationFromValue returns a pointer to a valid VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVTApiPlaneraResaWebV4ModelsLinkSegmentOrientationFromValue(v string) (*VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation, error) {
	ev := VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation: valid values are %v", v, AllowedVTApiPlaneraResaWebV4ModelsLinkSegmentOrientationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation) IsValid() bool {
	for _, existing := range AllowedVTApiPlaneraResaWebV4ModelsLinkSegmentOrientationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VT.ApiPlaneraResa.Web.V4.Models.LinkSegmentOrientation value
func (v VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation) Ptr() *VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation {
	return &v
}

type NullableVTApiPlaneraResaWebV4ModelsLinkSegmentOrientation struct {
	value *VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation
	isSet bool
}

func (v NullableVTApiPlaneraResaWebV4ModelsLinkSegmentOrientation) Get() *VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation {
	return v.value
}

func (v *NullableVTApiPlaneraResaWebV4ModelsLinkSegmentOrientation) Set(val *VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation) {
	v.value = val
	v.isSet = true
}

func (v NullableVTApiPlaneraResaWebV4ModelsLinkSegmentOrientation) IsSet() bool {
	return v.isSet
}

func (v *NullableVTApiPlaneraResaWebV4ModelsLinkSegmentOrientation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVTApiPlaneraResaWebV4ModelsLinkSegmentOrientation(val *VTApiPlaneraResaWebV4ModelsLinkSegmentOrientation) *NullableVTApiPlaneraResaWebV4ModelsLinkSegmentOrientation {
	return &NullableVTApiPlaneraResaWebV4ModelsLinkSegmentOrientation{value: val, isSet: true}
}

func (v NullableVTApiPlaneraResaWebV4ModelsLinkSegmentOrientation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVTApiPlaneraResaWebV4ModelsLinkSegmentOrientation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
