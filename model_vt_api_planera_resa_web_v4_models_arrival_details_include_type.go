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

// VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType The different kinds of detailed information that could be included in a get arrival details request.
type VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType string

// List of VT.ApiPlaneraResa.Web.V4.Models.ArrivalDetailsIncludeType
const (
	VTAPIPLANERARESAWEBV4MODELSARRIVALDETAILSINCLUDETYPE_SERVICEJOURNEYCALLS VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType = "servicejourneycalls"
	VTAPIPLANERARESAWEBV4MODELSARRIVALDETAILSINCLUDETYPE_SERVICEJOURNEYCOORDINATES VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType = "servicejourneycoordinates"
)

// All allowed values of VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType enum
var AllowedVTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeTypeEnumValues = []VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType{
	"servicejourneycalls",
	"servicejourneycoordinates",
}

func (v *VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType(value)
	for _, existing := range AllowedVTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType", value)
}

// NewVTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeTypeFromValue returns a pointer to a valid VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeTypeFromValue(v string) (*VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType, error) {
	ev := VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType: valid values are %v", v, AllowedVTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType) IsValid() bool {
	for _, existing := range AllowedVTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VT.ApiPlaneraResa.Web.V4.Models.ArrivalDetailsIncludeType value
func (v VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType) Ptr() *VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType {
	return &v
}

type NullableVTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType struct {
	value *VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType
	isSet bool
}

func (v NullableVTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType) Get() *VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType {
	return v.value
}

func (v *NullableVTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType) Set(val *VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType) {
	v.value = val
	v.isSet = true
}

func (v NullableVTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType) IsSet() bool {
	return v.isSet
}

func (v *NullableVTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType(val *VTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType) *NullableVTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType {
	return &NullableVTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType{value: val, isSet: true}
}

func (v NullableVTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVTApiPlaneraResaWebV4ModelsArrivalDetailsIncludeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

