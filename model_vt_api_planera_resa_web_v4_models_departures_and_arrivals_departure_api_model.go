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

// checks if the VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel{}

// VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel struct for VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel
type VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel struct {
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
	Occupancy *VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel `json:"occupancy,omitempty"`
}

// NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel instantiates a new VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel(stopPoint VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel, plannedTime string) *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel {
	this := VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel{}
	this.StopPoint = stopPoint
	this.PlannedTime = plannedTime
	return &this
}

// NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel {
	this := VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel{}
	return &this
}

// GetDetailsReference returns the DetailsReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetDetailsReference() string {
	if o == nil || IsNil(o.DetailsReference.Get()) {
		var ret string
		return ret
	}
	return *o.DetailsReference.Get()
}

// GetDetailsReferenceOk returns a tuple with the DetailsReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetDetailsReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DetailsReference.Get(), o.DetailsReference.IsSet()
}

// HasDetailsReference returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) HasDetailsReference() bool {
	if o != nil && o.DetailsReference.IsSet() {
		return true
	}

	return false
}

// SetDetailsReference gets a reference to the given NullableString and assigns it to the DetailsReference field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) SetDetailsReference(v string) {
	o.DetailsReference.Set(&v)
}
// SetDetailsReferenceNil sets the value for DetailsReference to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) SetDetailsReferenceNil() {
	o.DetailsReference.Set(nil)
}

// UnsetDetailsReference ensures that no value is present for DetailsReference, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) UnsetDetailsReference() {
	o.DetailsReference.Unset()
}

// GetServiceJourney returns the ServiceJourney field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetServiceJourney() VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel {
	if o == nil || IsNil(o.ServiceJourney) {
		var ret VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel
		return ret
	}
	return *o.ServiceJourney
}

// GetServiceJourneyOk returns a tuple with the ServiceJourney field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetServiceJourneyOk() (*VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel, bool) {
	if o == nil || IsNil(o.ServiceJourney) {
		return nil, false
	}
	return o.ServiceJourney, true
}

// HasServiceJourney returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) HasServiceJourney() bool {
	if o != nil && !IsNil(o.ServiceJourney) {
		return true
	}

	return false
}

// SetServiceJourney gets a reference to the given VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel and assigns it to the ServiceJourney field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) SetServiceJourney(v VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsServiceJourneyApiModel) {
	o.ServiceJourney = &v
}

// GetStopPoint returns the StopPoint field value
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetStopPoint() VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel {
	if o == nil {
		var ret VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel
		return ret
	}

	return o.StopPoint
}

// GetStopPointOk returns a tuple with the StopPoint field value
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetStopPointOk() (*VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StopPoint, true
}

// SetStopPoint sets field value
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) SetStopPoint(v VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsStopPointApiModel) {
	o.StopPoint = v
}

// GetPlannedTime returns the PlannedTime field value
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetPlannedTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlannedTime
}

// GetPlannedTimeOk returns a tuple with the PlannedTime field value
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetPlannedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlannedTime, true
}

// SetPlannedTime sets field value
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) SetPlannedTime(v string) {
	o.PlannedTime = v
}

// GetEstimatedTime returns the EstimatedTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetEstimatedTime() string {
	if o == nil || IsNil(o.EstimatedTime.Get()) {
		var ret string
		return ret
	}
	return *o.EstimatedTime.Get()
}

// GetEstimatedTimeOk returns a tuple with the EstimatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetEstimatedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedTime.Get(), o.EstimatedTime.IsSet()
}

// HasEstimatedTime returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) HasEstimatedTime() bool {
	if o != nil && o.EstimatedTime.IsSet() {
		return true
	}

	return false
}

// SetEstimatedTime gets a reference to the given NullableString and assigns it to the EstimatedTime field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) SetEstimatedTime(v string) {
	o.EstimatedTime.Set(&v)
}
// SetEstimatedTimeNil sets the value for EstimatedTime to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) SetEstimatedTimeNil() {
	o.EstimatedTime.Set(nil)
}

// UnsetEstimatedTime ensures that no value is present for EstimatedTime, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) UnsetEstimatedTime() {
	o.EstimatedTime.Unset()
}

// GetEstimatedOtherwisePlannedTime returns the EstimatedOtherwisePlannedTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetEstimatedOtherwisePlannedTime() string {
	if o == nil || IsNil(o.EstimatedOtherwisePlannedTime.Get()) {
		var ret string
		return ret
	}
	return *o.EstimatedOtherwisePlannedTime.Get()
}

// GetEstimatedOtherwisePlannedTimeOk returns a tuple with the EstimatedOtherwisePlannedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetEstimatedOtherwisePlannedTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedOtherwisePlannedTime.Get(), o.EstimatedOtherwisePlannedTime.IsSet()
}

// HasEstimatedOtherwisePlannedTime returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) HasEstimatedOtherwisePlannedTime() bool {
	if o != nil && o.EstimatedOtherwisePlannedTime.IsSet() {
		return true
	}

	return false
}

// SetEstimatedOtherwisePlannedTime gets a reference to the given NullableString and assigns it to the EstimatedOtherwisePlannedTime field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) SetEstimatedOtherwisePlannedTime(v string) {
	o.EstimatedOtherwisePlannedTime.Set(&v)
}
// SetEstimatedOtherwisePlannedTimeNil sets the value for EstimatedOtherwisePlannedTime to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) SetEstimatedOtherwisePlannedTimeNil() {
	o.EstimatedOtherwisePlannedTime.Set(nil)
}

// UnsetEstimatedOtherwisePlannedTime ensures that no value is present for EstimatedOtherwisePlannedTime, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) UnsetEstimatedOtherwisePlannedTime() {
	o.EstimatedOtherwisePlannedTime.Unset()
}

// GetIsCancelled returns the IsCancelled field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetIsCancelled() bool {
	if o == nil || IsNil(o.IsCancelled) {
		var ret bool
		return ret
	}
	return *o.IsCancelled
}

// GetIsCancelledOk returns a tuple with the IsCancelled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetIsCancelledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCancelled) {
		return nil, false
	}
	return o.IsCancelled, true
}

// HasIsCancelled returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) HasIsCancelled() bool {
	if o != nil && !IsNil(o.IsCancelled) {
		return true
	}

	return false
}

// SetIsCancelled gets a reference to the given bool and assigns it to the IsCancelled field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) SetIsCancelled(v bool) {
	o.IsCancelled = &v
}

// GetIsPartCancelled returns the IsPartCancelled field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetIsPartCancelled() bool {
	if o == nil || IsNil(o.IsPartCancelled) {
		var ret bool
		return ret
	}
	return *o.IsPartCancelled
}

// GetIsPartCancelledOk returns a tuple with the IsPartCancelled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetIsPartCancelledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPartCancelled) {
		return nil, false
	}
	return o.IsPartCancelled, true
}

// HasIsPartCancelled returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) HasIsPartCancelled() bool {
	if o != nil && !IsNil(o.IsPartCancelled) {
		return true
	}

	return false
}

// SetIsPartCancelled gets a reference to the given bool and assigns it to the IsPartCancelled field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) SetIsPartCancelled(v bool) {
	o.IsPartCancelled = &v
}

// GetOccupancy returns the Occupancy field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetOccupancy() VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel {
	if o == nil || IsNil(o.Occupancy) {
		var ret VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel
		return ret
	}
	return *o.Occupancy
}

// GetOccupancyOk returns a tuple with the Occupancy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) GetOccupancyOk() (*VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel, bool) {
	if o == nil || IsNil(o.Occupancy) {
		return nil, false
	}
	return o.Occupancy, true
}

// HasOccupancy returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) HasOccupancy() bool {
	if o != nil && !IsNil(o.Occupancy) {
		return true
	}

	return false
}

// SetOccupancy gets a reference to the given VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel and assigns it to the Occupancy field.
func (o *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) SetOccupancy(v VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel) {
	o.Occupancy = &v
}

func (o VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Occupancy) {
		toSerialize["occupancy"] = o.Occupancy
	}
	return toSerialize, nil
}

type NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel struct {
	value *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel
	isSet bool
}

func (v NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) Get() *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel {
	return v.value
}

func (v *NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) Set(val *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel(val *VTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) *NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel {
	return &NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel{value: val, isSet: true}
}

func (v NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVTApiPlaneraResaWebV4ModelsDeparturesAndArrivalsDepartureApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

