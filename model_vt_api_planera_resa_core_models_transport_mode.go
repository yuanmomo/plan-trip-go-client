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

// VTApiPlaneraResaCoreModelsTransportMode the model 'VTApiPlaneraResaCoreModelsTransportMode'
type VTApiPlaneraResaCoreModelsTransportMode string

// List of VT.ApiPlaneraResa.Core.Models.TransportMode
const (
	VTAPIPLANERARESACOREMODELSTRANSPORTMODE_UNKNOWN VTApiPlaneraResaCoreModelsTransportMode = "unknown"
	VTAPIPLANERARESACOREMODELSTRANSPORTMODE_NONE VTApiPlaneraResaCoreModelsTransportMode = "none"
	VTAPIPLANERARESACOREMODELSTRANSPORTMODE_TRAM VTApiPlaneraResaCoreModelsTransportMode = "tram"
	VTAPIPLANERARESACOREMODELSTRANSPORTMODE_BUS VTApiPlaneraResaCoreModelsTransportMode = "bus"
	VTAPIPLANERARESACOREMODELSTRANSPORTMODE_FERRY VTApiPlaneraResaCoreModelsTransportMode = "ferry"
	VTAPIPLANERARESACOREMODELSTRANSPORTMODE_TRAIN VTApiPlaneraResaCoreModelsTransportMode = "train"
	VTAPIPLANERARESACOREMODELSTRANSPORTMODE_TAXI VTApiPlaneraResaCoreModelsTransportMode = "taxi"
	VTAPIPLANERARESACOREMODELSTRANSPORTMODE_WALK VTApiPlaneraResaCoreModelsTransportMode = "walk"
	VTAPIPLANERARESACOREMODELSTRANSPORTMODE_BIKE VTApiPlaneraResaCoreModelsTransportMode = "bike"
	VTAPIPLANERARESACOREMODELSTRANSPORTMODE_CAR VTApiPlaneraResaCoreModelsTransportMode = "car"
)

// All allowed values of VTApiPlaneraResaCoreModelsTransportMode enum
var AllowedVTApiPlaneraResaCoreModelsTransportModeEnumValues = []VTApiPlaneraResaCoreModelsTransportMode{
	"unknown",
	"none",
	"tram",
	"bus",
	"ferry",
	"train",
	"taxi",
	"walk",
	"bike",
	"car",
}

func (v *VTApiPlaneraResaCoreModelsTransportMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VTApiPlaneraResaCoreModelsTransportMode(value)
	for _, existing := range AllowedVTApiPlaneraResaCoreModelsTransportModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VTApiPlaneraResaCoreModelsTransportMode", value)
}

// NewVTApiPlaneraResaCoreModelsTransportModeFromValue returns a pointer to a valid VTApiPlaneraResaCoreModelsTransportMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVTApiPlaneraResaCoreModelsTransportModeFromValue(v string) (*VTApiPlaneraResaCoreModelsTransportMode, error) {
	ev := VTApiPlaneraResaCoreModelsTransportMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VTApiPlaneraResaCoreModelsTransportMode: valid values are %v", v, AllowedVTApiPlaneraResaCoreModelsTransportModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VTApiPlaneraResaCoreModelsTransportMode) IsValid() bool {
	for _, existing := range AllowedVTApiPlaneraResaCoreModelsTransportModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VT.ApiPlaneraResa.Core.Models.TransportMode value
func (v VTApiPlaneraResaCoreModelsTransportMode) Ptr() *VTApiPlaneraResaCoreModelsTransportMode {
	return &v
}

type NullableVTApiPlaneraResaCoreModelsTransportMode struct {
	value *VTApiPlaneraResaCoreModelsTransportMode
	isSet bool
}

func (v NullableVTApiPlaneraResaCoreModelsTransportMode) Get() *VTApiPlaneraResaCoreModelsTransportMode {
	return v.value
}

func (v *NullableVTApiPlaneraResaCoreModelsTransportMode) Set(val *VTApiPlaneraResaCoreModelsTransportMode) {
	v.value = val
	v.isSet = true
}

func (v NullableVTApiPlaneraResaCoreModelsTransportMode) IsSet() bool {
	return v.isSet
}

func (v *NullableVTApiPlaneraResaCoreModelsTransportMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVTApiPlaneraResaCoreModelsTransportMode(val *VTApiPlaneraResaCoreModelsTransportMode) *NullableVTApiPlaneraResaCoreModelsTransportMode {
	return &NullableVTApiPlaneraResaCoreModelsTransportMode{value: val, isSet: true}
}

func (v NullableVTApiPlaneraResaCoreModelsTransportMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVTApiPlaneraResaCoreModelsTransportMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
