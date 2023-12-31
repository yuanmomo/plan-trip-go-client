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

// VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver Specifies a maneuver to be executed for a link segment.
type VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver string

// List of VT.ApiPlaneraResa.Web.V4.Models.LinkSegmentManeuver
const (
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_NONE VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "none"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_FROM VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "from"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_TO VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "to"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_ON VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "on"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_LEFT VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "left"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_RIGHT VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "right"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_KEEPLEFT VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "keepleft"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_KEEPRIGHT VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "keepright"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_HALFLEFT VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "halfleft"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_HALFRIGHT VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "halfright"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_KEEPHALFLEFT VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "keephalfleft"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_KEEPHALFRIGHT VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "keephalfright"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_SHARPLEFT VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "sharpleft"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_SHARPRIGHT VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "sharpright"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_KEEPSHARPLEFT VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "keepsharpleft"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_KEEPSHARPRIGHT VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "keepsharpright"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_STRAIGHT VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "straight"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_UTURN VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "uturn"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_ENTER VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "enter"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_LEAVE VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "leave"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_ENTERROUNDABOUT VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "enterroundabout"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_STAYINROUNDABOUT VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "stayinroundabout"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_LEAVEROUNDABOUT VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "leaveroundabout"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_ENTERFERRY VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "enterferry"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_LEAVEFERRY VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "leaveferry"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_CHANGEHIGHWAY VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "changehighway"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_CHECKINFERRY VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "checkinferry"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_CHECKOUTFERRY VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "checkoutferry"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_ELEVATOR VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "elevator"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_ELEVATORDOWN VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "elevatordown"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_ELEVATORUP VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "elevatorup"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_ESCALATOR VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "escalator"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_ESCALATORDOWN VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "escalatordown"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_ESCALATORUP VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "escalatorup"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_STAIRS VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "stairs"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_STAIRSDOWN VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "stairsdown"
	VTAPIPLANERARESAWEBV4MODELSLINKSEGMENTMANEUVER_STAIRSUP VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver = "stairsup"
)

// All allowed values of VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver enum
var AllowedVTApiPlaneraResaWebV4ModelsLinkSegmentManeuverEnumValues = []VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver{
	"none",
	"from",
	"to",
	"on",
	"left",
	"right",
	"keepleft",
	"keepright",
	"halfleft",
	"halfright",
	"keephalfleft",
	"keephalfright",
	"sharpleft",
	"sharpright",
	"keepsharpleft",
	"keepsharpright",
	"straight",
	"uturn",
	"enter",
	"leave",
	"enterroundabout",
	"stayinroundabout",
	"leaveroundabout",
	"enterferry",
	"leaveferry",
	"changehighway",
	"checkinferry",
	"checkoutferry",
	"elevator",
	"elevatordown",
	"elevatorup",
	"escalator",
	"escalatordown",
	"escalatorup",
	"stairs",
	"stairsdown",
	"stairsup",
}

func (v *VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver(value)
	for _, existing := range AllowedVTApiPlaneraResaWebV4ModelsLinkSegmentManeuverEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver", value)
}

// NewVTApiPlaneraResaWebV4ModelsLinkSegmentManeuverFromValue returns a pointer to a valid VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVTApiPlaneraResaWebV4ModelsLinkSegmentManeuverFromValue(v string) (*VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver, error) {
	ev := VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver: valid values are %v", v, AllowedVTApiPlaneraResaWebV4ModelsLinkSegmentManeuverEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver) IsValid() bool {
	for _, existing := range AllowedVTApiPlaneraResaWebV4ModelsLinkSegmentManeuverEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VT.ApiPlaneraResa.Web.V4.Models.LinkSegmentManeuver value
func (v VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver) Ptr() *VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver {
	return &v
}

type NullableVTApiPlaneraResaWebV4ModelsLinkSegmentManeuver struct {
	value *VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver
	isSet bool
}

func (v NullableVTApiPlaneraResaWebV4ModelsLinkSegmentManeuver) Get() *VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver {
	return v.value
}

func (v *NullableVTApiPlaneraResaWebV4ModelsLinkSegmentManeuver) Set(val *VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver) {
	v.value = val
	v.isSet = true
}

func (v NullableVTApiPlaneraResaWebV4ModelsLinkSegmentManeuver) IsSet() bool {
	return v.isSet
}

func (v *NullableVTApiPlaneraResaWebV4ModelsLinkSegmentManeuver) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVTApiPlaneraResaWebV4ModelsLinkSegmentManeuver(val *VTApiPlaneraResaWebV4ModelsLinkSegmentManeuver) *NullableVTApiPlaneraResaWebV4ModelsLinkSegmentManeuver {
	return &NullableVTApiPlaneraResaWebV4ModelsLinkSegmentManeuver{value: val, isSet: true}
}

func (v NullableVTApiPlaneraResaWebV4ModelsLinkSegmentManeuver) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVTApiPlaneraResaWebV4ModelsLinkSegmentManeuver) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

