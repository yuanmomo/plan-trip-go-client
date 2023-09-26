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

// checks if the VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse{}

// VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse struct for VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse
type VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse struct {
	// The results.
	Results []VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel `json:"results,omitempty"`
	Pagination *VTApiPlaneraResaWebV4ModelsPaginationProperties `json:"pagination,omitempty"`
	Links *VTApiPlaneraResaWebV4ModelsPaginationLinks `json:"links,omitempty"`
}

// NewVTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse instantiates a new VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse() *VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse {
	this := VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse{}
	return &this
}

// NewVTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponseWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponseWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse {
	this := VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse) GetResults() []VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel {
	if o == nil {
		var ret []VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse) GetResultsOk() ([]VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse) HasResults() bool {
	if o != nil && IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel and assigns it to the Results field.
func (o *VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse) SetResults(v []VTApiPlaneraResaWebV4ModelsJourneysJourneyApiModel) {
	o.Results = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse) GetPagination() VTApiPlaneraResaWebV4ModelsPaginationProperties {
	if o == nil || IsNil(o.Pagination) {
		var ret VTApiPlaneraResaWebV4ModelsPaginationProperties
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse) GetPaginationOk() (*VTApiPlaneraResaWebV4ModelsPaginationProperties, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given VTApiPlaneraResaWebV4ModelsPaginationProperties and assigns it to the Pagination field.
func (o *VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse) SetPagination(v VTApiPlaneraResaWebV4ModelsPaginationProperties) {
	o.Pagination = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse) GetLinks() VTApiPlaneraResaWebV4ModelsPaginationLinks {
	if o == nil || IsNil(o.Links) {
		var ret VTApiPlaneraResaWebV4ModelsPaginationLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse) GetLinksOk() (*VTApiPlaneraResaWebV4ModelsPaginationLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given VTApiPlaneraResaWebV4ModelsPaginationLinks and assigns it to the Links field.
func (o *VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse) SetLinks(v VTApiPlaneraResaWebV4ModelsPaginationLinks) {
	o.Links = &v
}

func (o VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableVTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse struct {
	value *VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse
	isSet bool
}

func (v NullableVTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse) Get() *VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse {
	return v.value
}

func (v *NullableVTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse) Set(val *VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse(val *VTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse) *NullableVTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse {
	return &NullableVTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse{value: val, isSet: true}
}

func (v NullableVTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVTApiPlaneraResaWebV4ModelsJourneysGetJourneysResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


