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

// VTApiPlaneraResaCoreModelsDateTimeRelatesToType the model 'VTApiPlaneraResaCoreModelsDateTimeRelatesToType'
type VTApiPlaneraResaCoreModelsDateTimeRelatesToType string

// List of VT.ApiPlaneraResa.Core.Models.DateTimeRelatesToType
const (
	VTAPIPLANERARESACOREMODELSDATETIMERELATESTOTYPE_DEPARTURE VTApiPlaneraResaCoreModelsDateTimeRelatesToType = "departure"
	VTAPIPLANERARESACOREMODELSDATETIMERELATESTOTYPE_ARRIVAL VTApiPlaneraResaCoreModelsDateTimeRelatesToType = "arrival"
)

// All allowed values of VTApiPlaneraResaCoreModelsDateTimeRelatesToType enum
var AllowedVTApiPlaneraResaCoreModelsDateTimeRelatesToTypeEnumValues = []VTApiPlaneraResaCoreModelsDateTimeRelatesToType{
	"departure",
	"arrival",
}

func (v *VTApiPlaneraResaCoreModelsDateTimeRelatesToType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VTApiPlaneraResaCoreModelsDateTimeRelatesToType(value)
	for _, existing := range AllowedVTApiPlaneraResaCoreModelsDateTimeRelatesToTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VTApiPlaneraResaCoreModelsDateTimeRelatesToType", value)
}

// NewVTApiPlaneraResaCoreModelsDateTimeRelatesToTypeFromValue returns a pointer to a valid VTApiPlaneraResaCoreModelsDateTimeRelatesToType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVTApiPlaneraResaCoreModelsDateTimeRelatesToTypeFromValue(v string) (*VTApiPlaneraResaCoreModelsDateTimeRelatesToType, error) {
	ev := VTApiPlaneraResaCoreModelsDateTimeRelatesToType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VTApiPlaneraResaCoreModelsDateTimeRelatesToType: valid values are %v", v, AllowedVTApiPlaneraResaCoreModelsDateTimeRelatesToTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VTApiPlaneraResaCoreModelsDateTimeRelatesToType) IsValid() bool {
	for _, existing := range AllowedVTApiPlaneraResaCoreModelsDateTimeRelatesToTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VT.ApiPlaneraResa.Core.Models.DateTimeRelatesToType value
func (v VTApiPlaneraResaCoreModelsDateTimeRelatesToType) Ptr() *VTApiPlaneraResaCoreModelsDateTimeRelatesToType {
	return &v
}

type NullableVTApiPlaneraResaCoreModelsDateTimeRelatesToType struct {
	value *VTApiPlaneraResaCoreModelsDateTimeRelatesToType
	isSet bool
}

func (v NullableVTApiPlaneraResaCoreModelsDateTimeRelatesToType) Get() *VTApiPlaneraResaCoreModelsDateTimeRelatesToType {
	return v.value
}

func (v *NullableVTApiPlaneraResaCoreModelsDateTimeRelatesToType) Set(val *VTApiPlaneraResaCoreModelsDateTimeRelatesToType) {
	v.value = val
	v.isSet = true
}

func (v NullableVTApiPlaneraResaCoreModelsDateTimeRelatesToType) IsSet() bool {
	return v.isSet
}

func (v *NullableVTApiPlaneraResaCoreModelsDateTimeRelatesToType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVTApiPlaneraResaCoreModelsDateTimeRelatesToType(val *VTApiPlaneraResaCoreModelsDateTimeRelatesToType) *NullableVTApiPlaneraResaCoreModelsDateTimeRelatesToType {
	return &NullableVTApiPlaneraResaCoreModelsDateTimeRelatesToType{value: val, isSet: true}
}

func (v NullableVTApiPlaneraResaCoreModelsDateTimeRelatesToType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVTApiPlaneraResaCoreModelsDateTimeRelatesToType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
