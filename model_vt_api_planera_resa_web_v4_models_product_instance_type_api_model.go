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

// VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel Specifies whether or not the product is dynamically based on the journey route.
type VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel string

// List of VT.ApiPlaneraResa.Web.V4.Models.ProductInstanceTypeApiModel
const (
	VTAPIPLANERARESAWEBV4MODELSPRODUCTINSTANCETYPEAPIMODEL_STATIC VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel = "static"
	VTAPIPLANERARESAWEBV4MODELSPRODUCTINSTANCETYPEAPIMODEL_DYNAMIC VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel = "dynamic"
)

// All allowed values of VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel enum
var AllowedVTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModelEnumValues = []VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel{
	"static",
	"dynamic",
}

func (v *VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel(value)
	for _, existing := range AllowedVTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel", value)
}

// NewVTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModelFromValue returns a pointer to a valid VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModelFromValue(v string) (*VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel, error) {
	ev := VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel: valid values are %v", v, AllowedVTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel) IsValid() bool {
	for _, existing := range AllowedVTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VT.ApiPlaneraResa.Web.V4.Models.ProductInstanceTypeApiModel value
func (v VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel) Ptr() *VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel {
	return &v
}

type NullableVTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel struct {
	value *VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel
	isSet bool
}

func (v NullableVTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel) Get() *VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel {
	return v.value
}

func (v *NullableVTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel) Set(val *VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableVTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableVTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel(val *VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel) *NullableVTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel {
	return &NullableVTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel{value: val, isSet: true}
}

func (v NullableVTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

