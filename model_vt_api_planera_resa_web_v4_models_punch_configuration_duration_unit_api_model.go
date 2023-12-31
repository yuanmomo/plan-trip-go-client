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

// VTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel Unit of duration of validity of a single punch.
type VTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel string

// List of VT.ApiPlaneraResa.Web.V4.Models.PunchConfigurationDurationUnitApiModel
const (
	VTAPIPLANERARESAWEBV4MODELSPUNCHCONFIGURATIONDURATIONUNITAPIMODEL_HOURS VTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel = "hours"
)

// All allowed values of VTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel enum
var AllowedVTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModelEnumValues = []VTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel{
	"hours",
}

func (v *VTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel(value)
	for _, existing := range AllowedVTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel", value)
}

// NewVTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModelFromValue returns a pointer to a valid VTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModelFromValue(v string) (*VTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel, error) {
	ev := VTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel: valid values are %v", v, AllowedVTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel) IsValid() bool {
	for _, existing := range AllowedVTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VT.ApiPlaneraResa.Web.V4.Models.PunchConfigurationDurationUnitApiModel value
func (v VTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel) Ptr() *VTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel {
	return &v
}

type NullableVTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel struct {
	value *VTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel
	isSet bool
}

func (v NullableVTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel) Get() *VTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel {
	return v.value
}

func (v *NullableVTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel) Set(val *VTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableVTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableVTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel(val *VTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel) *NullableVTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel {
	return &NullableVTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel{value: val, isSet: true}
}

func (v NullableVTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVTApiPlaneraResaWebV4ModelsPunchConfigurationDurationUnitApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

