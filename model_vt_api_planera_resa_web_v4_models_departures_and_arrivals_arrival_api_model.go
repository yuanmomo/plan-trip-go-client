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

// checks if the VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel{}

// VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel struct for VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel
type VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel struct {
	// A reference that should be used when getting detailed information about the journey.
	DetailsReference NullableString `json:"detailsReference,omitempty"`
	ServiceJourney *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel `json:"serviceJourney,omitempty"`
	StopPoint VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel `json:"stopPoint"`
	// The planned time of the call in RFC 3339 format.
	PlannedTime string `json:"plannedTime"`
	// The estimated time of the call in RFC 3339 format.
	EstimatedTime NullableString `json:"estimatedTime,omitempty"`
	// The best known time of the call in RFC 3339 format. Is EstimatedTime if exists, otherwise PlannedTime.
	EstimatedOtherwisePlannedTime NullableString `json:"estimatedOtherwisePlannedTime,omitempty"`
	// Flag indicating if the departure or arrival is cancelled.
	IsCancelled *bool `json:"isCancelled,omitempty"`
	// Flag indicating if the departure or arrival is partially cancelled.
	IsPartCancelled *bool `json:"isPartCancelled,omitempty"`
}

// NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel instantiates a new VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel(stopPoint VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel, plannedTime string) *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel {
	this := VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel{}
	this.StopPoint = stopPoint
	this.PlannedTime = plannedTime
	return &this
}

// NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel {
	this := VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel{}
	return &this
}

// GetDetailsReference returns the DetailsReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetDetailsReference() string {
	if o == nil || IsNil(o.DetailsReference.Get()) {
		var ret string
		return ret
	}
	return *o.DetailsReference.Get()
}

// GetDetailsReferenceOk returns a tuple with the DetailsReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetDetailsReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DetailsReference.Get(), o.DetailsReference.IsSet()
}

// HasDetailsReference returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) HasDetailsReference() bool {
	if o != nil && o.DetailsReference.IsSet() {
		return true
	}

	return false
}

// SetDetailsReference gets a reference to the given NullableString and assigns it to the DetailsReference field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) SetDetailsReference(v string) {
	o.DetailsReference.Set(&v)
}
// SetDetailsReferenceNil sets the value for DetailsReference to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) SetDetailsReferenceNil() {
	o.DetailsReference.Set(nil)
}

// UnsetDetailsReference ensures that no value is present for DetailsReference, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) UnsetDetailsReference() {
	o.DetailsReference.Unset()
}

// GetServiceJourney returns the ServiceJourney field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetServiceJourney() VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel {
	if o == nil || IsNil(o.ServiceJourney) {
		var ret VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel
		return ret
	}
	return *o.ServiceJourney
}

// GetServiceJourneyOk returns a tuple with the ServiceJourney field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetServiceJourneyOk() (*VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel, bool) {
	if o == nil || IsNil(o.ServiceJourney) {
		return nil, false
	}
	return o.ServiceJourney, true
}

// HasServiceJourney returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) HasServiceJourney() bool {
	if o != nil && !IsNil(o.ServiceJourney) {
		return true
	}

	return false
}

// SetServiceJourney gets a reference to the given VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel and assigns it to the ServiceJourney field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) SetServiceJourney(v VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) {
	o.ServiceJourney = &v
}

// GetStopPoint returns the StopPoint field value
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetStopPoint() VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel {
	if o == nil {
		var ret VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel
		return ret
	}

	return o.StopPoint
}

// GetStopPointOk returns a tuple with the StopPoint field value
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetStopPointOk() (*VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StopPoint, true
}

// SetStopPoint sets field value
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) SetStopPoint(v VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) {
	o.StopPoint = v
}

// GetPlannedTime returns the PlannedTime field value
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetPlannedTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlannedTime
}

// GetPlannedTimeOk returns a tuple with the PlannedTime field value
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetPlannedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlannedTime, true
}

// SetPlannedTime sets field value
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) SetPlannedTime(v string) {
	o.PlannedTime = v
}

// GetEstimatedTime returns the EstimatedTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetEstimatedTime() string {
	if o == nil || IsNil(o.EstimatedTime.Get()) {
		var ret string
		return ret
	}
	return *o.EstimatedTime.Get()
}

// GetEstimatedTimeOk returns a tuple with the EstimatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetEstimatedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedTime.Get(), o.EstimatedTime.IsSet()
}

// HasEstimatedTime returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) HasEstimatedTime() bool {
	if o != nil && o.EstimatedTime.IsSet() {
		return true
	}

	return false
}

// SetEstimatedTime gets a reference to the given NullableString and assigns it to the EstimatedTime field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) SetEstimatedTime(v string) {
	o.EstimatedTime.Set(&v)
}
// SetEstimatedTimeNil sets the value for EstimatedTime to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) SetEstimatedTimeNil() {
	o.EstimatedTime.Set(nil)
}

// UnsetEstimatedTime ensures that no value is present for EstimatedTime, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) UnsetEstimatedTime() {
	o.EstimatedTime.Unset()
}

// GetEstimatedOtherwisePlannedTime returns the EstimatedOtherwisePlannedTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetEstimatedOtherwisePlannedTime() string {
	if o == nil || IsNil(o.EstimatedOtherwisePlannedTime.Get()) {
		var ret string
		return ret
	}
	return *o.EstimatedOtherwisePlannedTime.Get()
}

// GetEstimatedOtherwisePlannedTimeOk returns a tuple with the EstimatedOtherwisePlannedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetEstimatedOtherwisePlannedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedOtherwisePlannedTime.Get(), o.EstimatedOtherwisePlannedTime.IsSet()
}

// HasEstimatedOtherwisePlannedTime returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) HasEstimatedOtherwisePlannedTime() bool {
	if o != nil && o.EstimatedOtherwisePlannedTime.IsSet() {
		return true
	}

	return false
}

// SetEstimatedOtherwisePlannedTime gets a reference to the given NullableString and assigns it to the EstimatedOtherwisePlannedTime field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) SetEstimatedOtherwisePlannedTime(v string) {
	o.EstimatedOtherwisePlannedTime.Set(&v)
}
// SetEstimatedOtherwisePlannedTimeNil sets the value for EstimatedOtherwisePlannedTime to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) SetEstimatedOtherwisePlannedTimeNil() {
	o.EstimatedOtherwisePlannedTime.Set(nil)
}

// UnsetEstimatedOtherwisePlannedTime ensures that no value is present for EstimatedOtherwisePlannedTime, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) UnsetEstimatedOtherwisePlannedTime() {
	o.EstimatedOtherwisePlannedTime.Unset()
}

// GetIsCancelled returns the IsCancelled field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetIsCancelled() bool {
	if o == nil || IsNil(o.IsCancelled) {
		var ret bool
		return ret
	}
	return *o.IsCancelled
}

// GetIsCancelledOk returns a tuple with the IsCancelled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetIsCancelledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCancelled) {
		return nil, false
	}
	return o.IsCancelled, true
}

// HasIsCancelled returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) HasIsCancelled() bool {
	if o != nil && !IsNil(o.IsCancelled) {
		return true
	}

	return false
}

// SetIsCancelled gets a reference to the given bool and assigns it to the IsCancelled field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) SetIsCancelled(v bool) {
	o.IsCancelled = &v
}

// GetIsPartCancelled returns the IsPartCancelled field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetIsPartCancelled() bool {
	if o == nil || IsNil(o.IsPartCancelled) {
		var ret bool
		return ret
	}
	return *o.IsPartCancelled
}

// GetIsPartCancelledOk returns a tuple with the IsPartCancelled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) GetIsPartCancelledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPartCancelled) {
		return nil, false
	}
	return o.IsPartCancelled, true
}

// HasIsPartCancelled returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) HasIsPartCancelled() bool {
	if o != nil && !IsNil(o.IsPartCancelled) {
		return true
	}

	return false
}

// SetIsPartCancelled gets a reference to the given bool and assigns it to the IsPartCancelled field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) SetIsPartCancelled(v bool) {
	o.IsPartCancelled = &v
}

func (o VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DetailsReference.IsSet() {
		toSerialize["detailsReference"] = o.DetailsReference.Get()
	}
	if !IsNil(o.ServiceJourney) {
		toSerialize["serviceJourney"] = o.ServiceJourney
	}
	toSerialize["stopPoint"] = o.StopPoint
	toSerialize["plannedTime"] = o.PlannedTime
	if o.EstimatedTime.IsSet() {
		toSerialize["estimatedTime"] = o.EstimatedTime.Get()
	}
	if o.EstimatedOtherwisePlannedTime.IsSet() {
		toSerialize["estimatedOtherwisePlannedTime"] = o.EstimatedOtherwisePlannedTime.Get()
	}
	if !IsNil(o.IsCancelled) {
		toSerialize["isCancelled"] = o.IsCancelled
	}
	if !IsNil(o.IsPartCancelled) {
		toSerialize["isPartCancelled"] = o.IsPartCancelled
	}
	return toSerialize, nil
}

type NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel struct {
	value *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel
	isSet bool
}

func (v NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) Get() *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel {
	return v.value
}

func (v *NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) Set(val *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel(val *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) *NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel {
	return &NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel{value: val, isSet: true}
}

func (v NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsArrivalApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


