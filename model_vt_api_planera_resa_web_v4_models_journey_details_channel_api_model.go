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

// checks if the VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel{}

// VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel Information about a channel related to the ticket suggestion.
type VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel struct {
	// The channel id.
	Id *int32 `json:"id,omitempty"`
	// The channel-specific ticket name, set if the channel is configured to override the default  product name.
	TicketName NullableString `json:"ticketName,omitempty"`
}

// NewVTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel() *VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel {
	this := VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel{}
	return &this
}

// NewVTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel {
	this := VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel) SetId(v int32) {
	o.Id = &v
}

// GetTicketName returns the TicketName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel) GetTicketName() string {
	if o == nil || IsNil(o.TicketName.Get()) {
		var ret string
		return ret
	}
	return *o.TicketName.Get()
}

// GetTicketNameOk returns a tuple with the TicketName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel) GetTicketNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TicketName.Get(), o.TicketName.IsSet()
}

// HasTicketName returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel) HasTicketName() bool {
	if o != nil && o.TicketName.IsSet() {
		return true
	}

	return false
}

// SetTicketName gets a reference to the given NullableString and assigns it to the TicketName field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel) SetTicketName(v string) {
	o.TicketName.Set(&v)
}
// SetTicketNameNil sets the value for TicketName to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel) SetTicketNameNil() {
	o.TicketName.Set(nil)
}

// UnsetTicketName ensures that no value is present for TicketName, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel) UnsetTicketName() {
	o.TicketName.Unset()
}

func (o VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.TicketName.IsSet() {
		toSerialize["ticketName"] = o.TicketName.Get()
	}
	return toSerialize, nil
}

type NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel struct {
	value *VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel
	isSet bool
}

func (v NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel) Get() *VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel {
	return v.value
}

func (v *NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel) Set(val *VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel(val *VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel) *NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel {
	return &NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel{value: val, isSet: true}
}

func (v NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

