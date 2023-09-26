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

// VTApiPlaneraResaCoreModelsTimeValidityType the model 'VTApiPlaneraResaCoreModelsTimeValidityType'
type VTApiPlaneraResaCoreModelsTimeValidityType string

// List of VT.ApiPlaneraResa.Core.Models.TimeValidityType
const (
	VTAPIPLANERARESACOREMODELSTIMEVALIDITYTYPE_UNKNOWN VTApiPlaneraResaCoreModelsTimeValidityType = "unknown"
	VTAPIPLANERARESACOREMODELSTIMEVALIDITYTYPE_AMOUNTUNIT VTApiPlaneraResaCoreModelsTimeValidityType = "amountunit"
	VTAPIPLANERARESACOREMODELSTIMEVALIDITYTYPE_FROMTODATE VTApiPlaneraResaCoreModelsTimeValidityType = "fromtodate"
	VTAPIPLANERARESACOREMODELSTIMEVALIDITYTYPE_FROMTODATETIME VTApiPlaneraResaCoreModelsTimeValidityType = "fromtodatetime"
)

// All allowed values of VTApiPlaneraResaCoreModelsTimeValidityType enum
var AllowedVTApiPlaneraResaCoreModelsTimeValidityTypeEnumValues = []VTApiPlaneraResaCoreModelsTimeValidityType{
	"unknown",
	"amountunit",
	"fromtodate",
	"fromtodatetime",
}

func (v *VTApiPlaneraResaCoreModelsTimeValidityType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VTApiPlaneraResaCoreModelsTimeValidityType(value)
	for _, existing := range AllowedVTApiPlaneraResaCoreModelsTimeValidityTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VTApiPlaneraResaCoreModelsTimeValidityType", value)
}

// NewVTApiPlaneraResaCoreModelsTimeValidityTypeFromValue returns a pointer to a valid VTApiPlaneraResaCoreModelsTimeValidityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVTApiPlaneraResaCoreModelsTimeValidityTypeFromValue(v string) (*VTApiPlaneraResaCoreModelsTimeValidityType, error) {
	ev := VTApiPlaneraResaCoreModelsTimeValidityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VTApiPlaneraResaCoreModelsTimeValidityType: valid values are %v", v, AllowedVTApiPlaneraResaCoreModelsTimeValidityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VTApiPlaneraResaCoreModelsTimeValidityType) IsValid() bool {
	for _, existing := range AllowedVTApiPlaneraResaCoreModelsTimeValidityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VT.ApiPlaneraResa.Core.Models.TimeValidityType value
func (v VTApiPlaneraResaCoreModelsTimeValidityType) Ptr() *VTApiPlaneraResaCoreModelsTimeValidityType {
	return &v
}

type NullableVTApiPlaneraResaCoreModelsTimeValidityType struct {
	value *VTApiPlaneraResaCoreModelsTimeValidityType
	isSet bool
}

func (v NullableVTApiPlaneraResaCoreModelsTimeValidityType) Get() *VTApiPlaneraResaCoreModelsTimeValidityType {
	return v.value
}

func (v *NullableVTApiPlaneraResaCoreModelsTimeValidityType) Set(val *VTApiPlaneraResaCoreModelsTimeValidityType) {
	v.value = val
	v.isSet = true
}

func (v NullableVTApiPlaneraResaCoreModelsTimeValidityType) IsSet() bool {
	return v.isSet
}

func (v *NullableVTApiPlaneraResaCoreModelsTimeValidityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVTApiPlaneraResaCoreModelsTimeValidityType(val *VTApiPlaneraResaCoreModelsTimeValidityType) *NullableVTApiPlaneraResaCoreModelsTimeValidityType {
	return &NullableVTApiPlaneraResaCoreModelsTimeValidityType{value: val, isSet: true}
}

func (v NullableVTApiPlaneraResaCoreModelsTimeValidityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVTApiPlaneraResaCoreModelsTimeValidityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
