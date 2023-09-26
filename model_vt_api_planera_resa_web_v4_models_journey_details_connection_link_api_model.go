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

// checks if the VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel{}

// VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel Information about a walk, bike or car link between two public transport trip legs.
type VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel struct {
	TransportMode *VTApiPlaneraResaCoreModelsTransportMode `json:"transportMode,omitempty"`
	TransportSubMode *VTApiPlaneraResaCoreModelsTransportSubMode `json:"transportSubMode,omitempty"`
	Origin *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallApiModel `json:"origin,omitempty"`
	Destination *VTApiPlaneraResaWebV4ModelsJourneyDetailsCallApiModel `json:"destination,omitempty"`
	// An ordered list (most important first) of notes related to the access link.
	Notes []VTApiPlaneraResaCoreModelsNote `json:"notes,omitempty"`
	// Distance in meters.
	DistanceInMeters NullableInt32 `json:"distanceInMeters,omitempty"`
	// The planned departure time in RFC 3339 format.
	PlannedDepartureTime NullableString `json:"plannedDepartureTime,omitempty"`
	// The planned arrival time in RFC 3339 format.
	PlannedArrivalTime NullableString `json:"plannedArrivalTime,omitempty"`
	// The planned duration in minutes.
	PlannedDurationInMinutes NullableInt32 `json:"plannedDurationInMinutes,omitempty"`
	// The estimated departure time in RFC 3339 format, if available.
	EstimatedDepartureTime NullableString `json:"estimatedDepartureTime,omitempty"`
	// The estimated arrival time in RFC 3339 format, if available.
	EstimatedArrivalTime NullableString `json:"estimatedArrivalTime,omitempty"`
	// The estimated duration in minutes, if available.
	EstimatedDurationInMinutes NullableInt32 `json:"estimatedDurationInMinutes,omitempty"`
	// Number of steps based on the distance and an estimated step length of 0.65 meters.
	EstimatedNumberOfSteps NullableInt32 `json:"estimatedNumberOfSteps,omitempty"`
	// The coordinates for the link.
	LinkCoordinates []VTApiPlaneraResaWebV4ModelsCoordinateApiModel `json:"linkCoordinates,omitempty"`
	// The segments that make up this link.
	Segments []VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel `json:"segments,omitempty"`
	// Index of Leg in Journey
	JourneyLegIndex *int32 `json:"journeyLegIndex,omitempty"`
}

// NewVTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel() *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel {
	this := VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel{}
	return &this
}

// NewVTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel {
	this := VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel{}
	return &this
}

// GetTransportMode returns the TransportMode field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetTransportMode() VTApiPlaneraResaCoreModelsTransportMode {
	if o == nil || IsNil(o.TransportMode) {
		var ret VTApiPlaneraResaCoreModelsTransportMode
		return ret
	}
	return *o.TransportMode
}

// GetTransportModeOk returns a tuple with the TransportMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetTransportModeOk() (*VTApiPlaneraResaCoreModelsTransportMode, bool) {
	if o == nil || IsNil(o.TransportMode) {
		return nil, false
	}
	return o.TransportMode, true
}

// HasTransportMode returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) HasTransportMode() bool {
	if o != nil && !IsNil(o.TransportMode) {
		return true
	}

	return false
}

// SetTransportMode gets a reference to the given VTApiPlaneraResaCoreModelsTransportMode and assigns it to the TransportMode field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) SetTransportMode(v VTApiPlaneraResaCoreModelsTransportMode) {
	o.TransportMode = &v
}

// GetTransportSubMode returns the TransportSubMode field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetTransportSubMode() VTApiPlaneraResaCoreModelsTransportSubMode {
	if o == nil || IsNil(o.TransportSubMode) {
		var ret VTApiPlaneraResaCoreModelsTransportSubMode
		return ret
	}
	return *o.TransportSubMode
}

// GetTransportSubModeOk returns a tuple with the TransportSubMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetTransportSubModeOk() (*VTApiPlaneraResaCoreModelsTransportSubMode, bool) {
	if o == nil || IsNil(o.TransportSubMode) {
		return nil, false
	}
	return o.TransportSubMode, true
}

// HasTransportSubMode returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) HasTransportSubMode() bool {
	if o != nil && !IsNil(o.TransportSubMode) {
		return true
	}

	return false
}

// SetTransportSubMode gets a reference to the given VTApiPlaneraResaCoreModelsTransportSubMode and assigns it to the TransportSubMode field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) SetTransportSubMode(v VTApiPlaneraResaCoreModelsTransportSubMode) {
	o.TransportSubMode = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetOrigin() VTApiPlaneraResaWebV4ModelsJourneyDetailsCallApiModel {
	if o == nil || IsNil(o.Origin) {
		var ret VTApiPlaneraResaWebV4ModelsJourneyDetailsCallApiModel
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetOriginOk() (*VTApiPlaneraResaWebV4ModelsJourneyDetailsCallApiModel, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given VTApiPlaneraResaWebV4ModelsJourneyDetailsCallApiModel and assigns it to the Origin field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) SetOrigin(v VTApiPlaneraResaWebV4ModelsJourneyDetailsCallApiModel) {
	o.Origin = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetDestination() VTApiPlaneraResaWebV4ModelsJourneyDetailsCallApiModel {
	if o == nil || IsNil(o.Destination) {
		var ret VTApiPlaneraResaWebV4ModelsJourneyDetailsCallApiModel
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetDestinationOk() (*VTApiPlaneraResaWebV4ModelsJourneyDetailsCallApiModel, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given VTApiPlaneraResaWebV4ModelsJourneyDetailsCallApiModel and assigns it to the Destination field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) SetDestination(v VTApiPlaneraResaWebV4ModelsJourneyDetailsCallApiModel) {
	o.Destination = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetNotes() []VTApiPlaneraResaCoreModelsNote {
	if o == nil {
		var ret []VTApiPlaneraResaCoreModelsNote
		return ret
	}
	return o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetNotesOk() ([]VTApiPlaneraResaCoreModelsNote, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) HasNotes() bool {
	if o != nil && IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given []VTApiPlaneraResaCoreModelsNote and assigns it to the Notes field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) SetNotes(v []VTApiPlaneraResaCoreModelsNote) {
	o.Notes = v
}

// GetDistanceInMeters returns the DistanceInMeters field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetDistanceInMeters() int32 {
	if o == nil || IsNil(o.DistanceInMeters.Get()) {
		var ret int32
		return ret
	}
	return *o.DistanceInMeters.Get()
}

// GetDistanceInMetersOk returns a tuple with the DistanceInMeters field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetDistanceInMetersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.DistanceInMeters.Get(), o.DistanceInMeters.IsSet()
}

// HasDistanceInMeters returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) HasDistanceInMeters() bool {
	if o != nil && o.DistanceInMeters.IsSet() {
		return true
	}

	return false
}

// SetDistanceInMeters gets a reference to the given NullableInt32 and assigns it to the DistanceInMeters field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) SetDistanceInMeters(v int32) {
	o.DistanceInMeters.Set(&v)
}
// SetDistanceInMetersNil sets the value for DistanceInMeters to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) SetDistanceInMetersNil() {
	o.DistanceInMeters.Set(nil)
}

// UnsetDistanceInMeters ensures that no value is present for DistanceInMeters, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) UnsetDistanceInMeters() {
	o.DistanceInMeters.Unset()
}

// GetPlannedDepartureTime returns the PlannedDepartureTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetPlannedDepartureTime() string {
	if o == nil || IsNil(o.PlannedDepartureTime.Get()) {
		var ret string
		return ret
	}
	return *o.PlannedDepartureTime.Get()
}

// GetPlannedDepartureTimeOk returns a tuple with the PlannedDepartureTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetPlannedDepartureTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlannedDepartureTime.Get(), o.PlannedDepartureTime.IsSet()
}

// HasPlannedDepartureTime returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) HasPlannedDepartureTime() bool {
	if o != nil && o.PlannedDepartureTime.IsSet() {
		return true
	}

	return false
}

// SetPlannedDepartureTime gets a reference to the given NullableString and assigns it to the PlannedDepartureTime field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) SetPlannedDepartureTime(v string) {
	o.PlannedDepartureTime.Set(&v)
}
// SetPlannedDepartureTimeNil sets the value for PlannedDepartureTime to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) SetPlannedDepartureTimeNil() {
	o.PlannedDepartureTime.Set(nil)
}

// UnsetPlannedDepartureTime ensures that no value is present for PlannedDepartureTime, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) UnsetPlannedDepartureTime() {
	o.PlannedDepartureTime.Unset()
}

// GetPlannedArrivalTime returns the PlannedArrivalTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetPlannedArrivalTime() string {
	if o == nil || IsNil(o.PlannedArrivalTime.Get()) {
		var ret string
		return ret
	}
	return *o.PlannedArrivalTime.Get()
}

// GetPlannedArrivalTimeOk returns a tuple with the PlannedArrivalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetPlannedArrivalTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlannedArrivalTime.Get(), o.PlannedArrivalTime.IsSet()
}

// HasPlannedArrivalTime returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) HasPlannedArrivalTime() bool {
	if o != nil && o.PlannedArrivalTime.IsSet() {
		return true
	}

	return false
}

// SetPlannedArrivalTime gets a reference to the given NullableString and assigns it to the PlannedArrivalTime field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) SetPlannedArrivalTime(v string) {
	o.PlannedArrivalTime.Set(&v)
}
// SetPlannedArrivalTimeNil sets the value for PlannedArrivalTime to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) SetPlannedArrivalTimeNil() {
	o.PlannedArrivalTime.Set(nil)
}

// UnsetPlannedArrivalTime ensures that no value is present for PlannedArrivalTime, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) UnsetPlannedArrivalTime() {
	o.PlannedArrivalTime.Unset()
}

// GetPlannedDurationInMinutes returns the PlannedDurationInMinutes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetPlannedDurationInMinutes() int32 {
	if o == nil || IsNil(o.PlannedDurationInMinutes.Get()) {
		var ret int32
		return ret
	}
	return *o.PlannedDurationInMinutes.Get()
}

// GetPlannedDurationInMinutesOk returns a tuple with the PlannedDurationInMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetPlannedDurationInMinutesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlannedDurationInMinutes.Get(), o.PlannedDurationInMinutes.IsSet()
}

// HasPlannedDurationInMinutes returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) HasPlannedDurationInMinutes() bool {
	if o != nil && o.PlannedDurationInMinutes.IsSet() {
		return true
	}

	return false
}

// SetPlannedDurationInMinutes gets a reference to the given NullableInt32 and assigns it to the PlannedDurationInMinutes field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) SetPlannedDurationInMinutes(v int32) {
	o.PlannedDurationInMinutes.Set(&v)
}
// SetPlannedDurationInMinutesNil sets the value for PlannedDurationInMinutes to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) SetPlannedDurationInMinutesNil() {
	o.PlannedDurationInMinutes.Set(nil)
}

// UnsetPlannedDurationInMinutes ensures that no value is present for PlannedDurationInMinutes, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) UnsetPlannedDurationInMinutes() {
	o.PlannedDurationInMinutes.Unset()
}

// GetEstimatedDepartureTime returns the EstimatedDepartureTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetEstimatedDepartureTime() string {
	if o == nil || IsNil(o.EstimatedDepartureTime.Get()) {
		var ret string
		return ret
	}
	return *o.EstimatedDepartureTime.Get()
}

// GetEstimatedDepartureTimeOk returns a tuple with the EstimatedDepartureTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetEstimatedDepartureTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedDepartureTime.Get(), o.EstimatedDepartureTime.IsSet()
}

// HasEstimatedDepartureTime returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) HasEstimatedDepartureTime() bool {
	if o != nil && o.EstimatedDepartureTime.IsSet() {
		return true
	}

	return false
}

// SetEstimatedDepartureTime gets a reference to the given NullableString and assigns it to the EstimatedDepartureTime field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) SetEstimatedDepartureTime(v string) {
	o.EstimatedDepartureTime.Set(&v)
}
// SetEstimatedDepartureTimeNil sets the value for EstimatedDepartureTime to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) SetEstimatedDepartureTimeNil() {
	o.EstimatedDepartureTime.Set(nil)
}

// UnsetEstimatedDepartureTime ensures that no value is present for EstimatedDepartureTime, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) UnsetEstimatedDepartureTime() {
	o.EstimatedDepartureTime.Unset()
}

// GetEstimatedArrivalTime returns the EstimatedArrivalTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetEstimatedArrivalTime() string {
	if o == nil || IsNil(o.EstimatedArrivalTime.Get()) {
		var ret string
		return ret
	}
	return *o.EstimatedArrivalTime.Get()
}

// GetEstimatedArrivalTimeOk returns a tuple with the EstimatedArrivalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetEstimatedArrivalTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedArrivalTime.Get(), o.EstimatedArrivalTime.IsSet()
}

// HasEstimatedArrivalTime returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) HasEstimatedArrivalTime() bool {
	if o != nil && o.EstimatedArrivalTime.IsSet() {
		return true
	}

	return false
}

// SetEstimatedArrivalTime gets a reference to the given NullableString and assigns it to the EstimatedArrivalTime field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) SetEstimatedArrivalTime(v string) {
	o.EstimatedArrivalTime.Set(&v)
}
// SetEstimatedArrivalTimeNil sets the value for EstimatedArrivalTime to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) SetEstimatedArrivalTimeNil() {
	o.EstimatedArrivalTime.Set(nil)
}

// UnsetEstimatedArrivalTime ensures that no value is present for EstimatedArrivalTime, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) UnsetEstimatedArrivalTime() {
	o.EstimatedArrivalTime.Unset()
}

// GetEstimatedDurationInMinutes returns the EstimatedDurationInMinutes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetEstimatedDurationInMinutes() int32 {
	if o == nil || IsNil(o.EstimatedDurationInMinutes.Get()) {
		var ret int32
		return ret
	}
	return *o.EstimatedDurationInMinutes.Get()
}

// GetEstimatedDurationInMinutesOk returns a tuple with the EstimatedDurationInMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetEstimatedDurationInMinutesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedDurationInMinutes.Get(), o.EstimatedDurationInMinutes.IsSet()
}

// HasEstimatedDurationInMinutes returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) HasEstimatedDurationInMinutes() bool {
	if o != nil && o.EstimatedDurationInMinutes.IsSet() {
		return true
	}

	return false
}

// SetEstimatedDurationInMinutes gets a reference to the given NullableInt32 and assigns it to the EstimatedDurationInMinutes field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) SetEstimatedDurationInMinutes(v int32) {
	o.EstimatedDurationInMinutes.Set(&v)
}
// SetEstimatedDurationInMinutesNil sets the value for EstimatedDurationInMinutes to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) SetEstimatedDurationInMinutesNil() {
	o.EstimatedDurationInMinutes.Set(nil)
}

// UnsetEstimatedDurationInMinutes ensures that no value is present for EstimatedDurationInMinutes, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) UnsetEstimatedDurationInMinutes() {
	o.EstimatedDurationInMinutes.Unset()
}

// GetEstimatedNumberOfSteps returns the EstimatedNumberOfSteps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetEstimatedNumberOfSteps() int32 {
	if o == nil || IsNil(o.EstimatedNumberOfSteps.Get()) {
		var ret int32
		return ret
	}
	return *o.EstimatedNumberOfSteps.Get()
}

// GetEstimatedNumberOfStepsOk returns a tuple with the EstimatedNumberOfSteps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetEstimatedNumberOfStepsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedNumberOfSteps.Get(), o.EstimatedNumberOfSteps.IsSet()
}

// HasEstimatedNumberOfSteps returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) HasEstimatedNumberOfSteps() bool {
	if o != nil && o.EstimatedNumberOfSteps.IsSet() {
		return true
	}

	return false
}

// SetEstimatedNumberOfSteps gets a reference to the given NullableInt32 and assigns it to the EstimatedNumberOfSteps field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) SetEstimatedNumberOfSteps(v int32) {
	o.EstimatedNumberOfSteps.Set(&v)
}
// SetEstimatedNumberOfStepsNil sets the value for EstimatedNumberOfSteps to be an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) SetEstimatedNumberOfStepsNil() {
	o.EstimatedNumberOfSteps.Set(nil)
}

// UnsetEstimatedNumberOfSteps ensures that no value is present for EstimatedNumberOfSteps, not even an explicit nil
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) UnsetEstimatedNumberOfSteps() {
	o.EstimatedNumberOfSteps.Unset()
}

// GetLinkCoordinates returns the LinkCoordinates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetLinkCoordinates() []VTApiPlaneraResaWebV4ModelsCoordinateApiModel {
	if o == nil {
		var ret []VTApiPlaneraResaWebV4ModelsCoordinateApiModel
		return ret
	}
	return o.LinkCoordinates
}

// GetLinkCoordinatesOk returns a tuple with the LinkCoordinates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetLinkCoordinatesOk() ([]VTApiPlaneraResaWebV4ModelsCoordinateApiModel, bool) {
	if o == nil || IsNil(o.LinkCoordinates) {
		return nil, false
	}
	return o.LinkCoordinates, true
}

// HasLinkCoordinates returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) HasLinkCoordinates() bool {
	if o != nil && IsNil(o.LinkCoordinates) {
		return true
	}

	return false
}

// SetLinkCoordinates gets a reference to the given []VTApiPlaneraResaWebV4ModelsCoordinateApiModel and assigns it to the LinkCoordinates field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) SetLinkCoordinates(v []VTApiPlaneraResaWebV4ModelsCoordinateApiModel) {
	o.LinkCoordinates = v
}

// GetSegments returns the Segments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetSegments() []VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel {
	if o == nil {
		var ret []VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel
		return ret
	}
	return o.Segments
}

// GetSegmentsOk returns a tuple with the Segments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetSegmentsOk() ([]VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel, bool) {
	if o == nil || IsNil(o.Segments) {
		return nil, false
	}
	return o.Segments, true
}

// HasSegments returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) HasSegments() bool {
	if o != nil && IsNil(o.Segments) {
		return true
	}

	return false
}

// SetSegments gets a reference to the given []VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel and assigns it to the Segments field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) SetSegments(v []VTApiPlaneraResaWebV4ModelsJourneyDetailsLinkSegmentApiModel) {
	o.Segments = v
}

// GetJourneyLegIndex returns the JourneyLegIndex field value if set, zero value otherwise.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetJourneyLegIndex() int32 {
	if o == nil || IsNil(o.JourneyLegIndex) {
		var ret int32
		return ret
	}
	return *o.JourneyLegIndex
}

// GetJourneyLegIndexOk returns a tuple with the JourneyLegIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) GetJourneyLegIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.JourneyLegIndex) {
		return nil, false
	}
	return o.JourneyLegIndex, true
}

// HasJourneyLegIndex returns a boolean if a field has been set.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) HasJourneyLegIndex() bool {
	if o != nil && !IsNil(o.JourneyLegIndex) {
		return true
	}

	return false
}

// SetJourneyLegIndex gets a reference to the given int32 and assigns it to the JourneyLegIndex field.
func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) SetJourneyLegIndex(v int32) {
	o.JourneyLegIndex = &v
}

func (o VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TransportMode) {
		toSerialize["transportMode"] = o.TransportMode
	}
	if !IsNil(o.TransportSubMode) {
		toSerialize["transportSubMode"] = o.TransportSubMode
	}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if o.Notes != nil {
		toSerialize["notes"] = o.Notes
	}
	if o.DistanceInMeters.IsSet() {
		toSerialize["distanceInMeters"] = o.DistanceInMeters.Get()
	}
	if o.PlannedDepartureTime.IsSet() {
		toSerialize["plannedDepartureTime"] = o.PlannedDepartureTime.Get()
	}
	if o.PlannedArrivalTime.IsSet() {
		toSerialize["plannedArrivalTime"] = o.PlannedArrivalTime.Get()
	}
	if o.PlannedDurationInMinutes.IsSet() {
		toSerialize["plannedDurationInMinutes"] = o.PlannedDurationInMinutes.Get()
	}
	if o.EstimatedDepartureTime.IsSet() {
		toSerialize["estimatedDepartureTime"] = o.EstimatedDepartureTime.Get()
	}
	if o.EstimatedArrivalTime.IsSet() {
		toSerialize["estimatedArrivalTime"] = o.EstimatedArrivalTime.Get()
	}
	if o.EstimatedDurationInMinutes.IsSet() {
		toSerialize["estimatedDurationInMinutes"] = o.EstimatedDurationInMinutes.Get()
	}
	if o.EstimatedNumberOfSteps.IsSet() {
		toSerialize["estimatedNumberOfSteps"] = o.EstimatedNumberOfSteps.Get()
	}
	if o.LinkCoordinates != nil {
		toSerialize["linkCoordinates"] = o.LinkCoordinates
	}
	if o.Segments != nil {
		toSerialize["segments"] = o.Segments
	}
	if !IsNil(o.JourneyLegIndex) {
		toSerialize["journeyLegIndex"] = o.JourneyLegIndex
	}
	return toSerialize, nil
}

type NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel struct {
	value *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel
	isSet bool
}

func (v NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) Get() *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel {
	return v.value
}

func (v *NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) Set(val *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) {
	v.value = val
	v.isSet = true
}

func (v NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) IsSet() bool {
	return v.isSet
}

func (v *NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel(val *VTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) *NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel {
	return &NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel{value: val, isSet: true}
}

func (v NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVTApiPlaneraResaWebV4ModelsJourneyDetailsConnectionLinkApiModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

