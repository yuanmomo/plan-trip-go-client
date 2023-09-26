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

// VTApiPlaneraResaWebV4ModelsOccupancyInformationSource Represents the source from which the occupancy information originate.
type VTApiPlaneraResaWebV4ModelsOccupancyInformationSource string

// List of VT.ApiPlaneraResa.Web.V4.Models.OccupancyInformationSource
const (
	VTAPIPLANERARESAWEBV4MODELSOCCUPANCYINFORMATIONSOURCE_PREDICTION VTApiPlaneraResaWebV4ModelsOccupancyInformationSource = "prediction"
	VTAPIPLANERARESAWEBV4MODELSOCCUPANCYINFORMATIONSOURCE_REALTIME VTApiPlaneraResaWebV4ModelsOccupancyInformationSource = "realtime"
)

// All allowed values of VTApiPlaneraResaWebV4ModelsOccupancyInformationSource enum
var AllowedVTApiPlaneraResaWebV4ModelsOccupancyInformationSourceEnumValues = []VTApiPlaneraResaWebV4ModelsOccupancyInformationSource{
	"prediction",
	"realtime",
}

func (v *VTApiPlaneraResaWebV4ModelsOccupancyInformationSource) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VTApiPlaneraResaWebV4ModelsOccupancyInformationSource(value)
	for _, existing := range AllowedVTApiPlaneraResaWebV4ModelsOccupancyInformationSourceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VTApiPlaneraResaWebV4ModelsOccupancyInformationSource", value)
}

// NewVTApiPlaneraResaWebV4ModelsOccupancyInformationSourceFromValue returns a pointer to a valid VTApiPlaneraResaWebV4ModelsOccupancyInformationSource
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVTApiPlaneraResaWebV4ModelsOccupancyInformationSourceFromValue(v string) (*VTApiPlaneraResaWebV4ModelsOccupancyInformationSource, error) {
	ev := VTApiPlaneraResaWebV4ModelsOccupancyInformationSource(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VTApiPlaneraResaWebV4ModelsOccupancyInformationSource: valid values are %v", v, AllowedVTApiPlaneraResaWebV4ModelsOccupancyInformationSourceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VTApiPlaneraResaWebV4ModelsOccupancyInformationSource) IsValid() bool {
	for _, existing := range AllowedVTApiPlaneraResaWebV4ModelsOccupancyInformationSourceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VT.ApiPlaneraResa.Web.V4.Models.OccupancyInformationSource value
func (v VTApiPlaneraResaWebV4ModelsOccupancyInformationSource) Ptr() *VTApiPlaneraResaWebV4ModelsOccupancyInformationSource {
	return &v
}

type NullableVTApiPlaneraResaWebV4ModelsOccupancyInformationSource struct {
	value *VTApiPlaneraResaWebV4ModelsOccupancyInformationSource
	isSet bool
}

func (v NullableVTApiPlaneraResaWebV4ModelsOccupancyInformationSource) Get() *VTApiPlaneraResaWebV4ModelsOccupancyInformationSource {
	return v.value
}

func (v *NullableVTApiPlaneraResaWebV4ModelsOccupancyInformationSource) Set(val *VTApiPlaneraResaWebV4ModelsOccupancyInformationSource) {
	v.value = val
	v.isSet = true
}

func (v NullableVTApiPlaneraResaWebV4ModelsOccupancyInformationSource) IsSet() bool {
	return v.isSet
}

func (v *NullableVTApiPlaneraResaWebV4ModelsOccupancyInformationSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVTApiPlaneraResaWebV4ModelsOccupancyInformationSource(val *VTApiPlaneraResaWebV4ModelsOccupancyInformationSource) *NullableVTApiPlaneraResaWebV4ModelsOccupancyInformationSource {
	return &NullableVTApiPlaneraResaWebV4ModelsOccupancyInformationSource{value: val, isSet: true}
}

func (v NullableVTApiPlaneraResaWebV4ModelsOccupancyInformationSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVTApiPlaneraResaWebV4ModelsOccupancyInformationSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

