# VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransportMode** | Pointer to [**VTApiPlaneraResaCoreModelsTransportMode**](VTApiPlaneraResaCoreModelsTransportMode.md) |  | [optional] 
**TransportSubMode** | Pointer to [**VTApiPlaneraResaCoreModelsTransportSubMode**](VTApiPlaneraResaCoreModelsTransportSubMode.md) |  | [optional] 
**Origin** | Pointer to [**VTApiPlaneraResaWebV4ModelsJourneysLinkEndpointApiModel**](VTApiPlaneraResaWebV4ModelsJourneysLinkEndpointApiModel.md) |  | [optional] 
**Destination** | Pointer to [**VTApiPlaneraResaWebV4ModelsJourneysLinkEndpointApiModel**](VTApiPlaneraResaWebV4ModelsJourneysLinkEndpointApiModel.md) |  | [optional] 
**Notes** | Pointer to [**[]VTApiPlaneraResaCoreModelsNote**](VTApiPlaneraResaCoreModelsNote.md) | An ordered list (most important first) of notes related to the access link. | [optional] 
**DistanceInMeters** | Pointer to **NullableInt32** | Distance in meters. | [optional] 
**PlannedDepartureTime** | Pointer to **NullableString** | The planned departure time in RFC 3339 format. | [optional] 
**PlannedArrivalTime** | Pointer to **NullableString** | The planned arrival time in RFC 3339 format. | [optional] 
**PlannedDurationInMinutes** | Pointer to **NullableInt32** | The planned duration in minutes. | [optional] 
**EstimatedDepartureTime** | Pointer to **NullableString** | The estimated departure time in RFC 3339 format, if available. | [optional] 
**EstimatedArrivalTime** | Pointer to **NullableString** | The estimated arrival time in RFC 3339 format, if available. | [optional] 
**EstimatedDurationInMinutes** | Pointer to **NullableInt32** | The estimated duration in minutes, if available. | [optional] 
**EstimatedNumberOfSteps** | Pointer to **NullableInt32** | Number of steps based on the distance and an estimated step length of 0.65 meters. | [optional] 
**LinkCoordinates** | Pointer to [**[]VTApiPlaneraResaWebV4ModelsCoordinateApiModel**](VTApiPlaneraResaWebV4ModelsCoordinateApiModel.md) | The coordinates for the link. | [optional] 
**Segments** | Pointer to [**[]VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel**](VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel.md) | The segments that make up this link. | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel

`func NewVTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel() *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel instantiates a new VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransportMode

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetTransportMode() VTApiPlaneraResaCoreModelsTransportMode`

GetTransportMode returns the TransportMode field if non-nil, zero value otherwise.

### GetTransportModeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetTransportModeOk() (*VTApiPlaneraResaCoreModelsTransportMode, bool)`

GetTransportModeOk returns a tuple with the TransportMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportMode

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetTransportMode(v VTApiPlaneraResaCoreModelsTransportMode)`

SetTransportMode sets TransportMode field to given value.

### HasTransportMode

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) HasTransportMode() bool`

HasTransportMode returns a boolean if a field has been set.

### GetTransportSubMode

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetTransportSubMode() VTApiPlaneraResaCoreModelsTransportSubMode`

GetTransportSubMode returns the TransportSubMode field if non-nil, zero value otherwise.

### GetTransportSubModeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetTransportSubModeOk() (*VTApiPlaneraResaCoreModelsTransportSubMode, bool)`

GetTransportSubModeOk returns a tuple with the TransportSubMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportSubMode

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetTransportSubMode(v VTApiPlaneraResaCoreModelsTransportSubMode)`

SetTransportSubMode sets TransportSubMode field to given value.

### HasTransportSubMode

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) HasTransportSubMode() bool`

HasTransportSubMode returns a boolean if a field has been set.

### GetOrigin

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetOrigin() VTApiPlaneraResaWebV4ModelsJourneysLinkEndpointApiModel`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetOriginOk() (*VTApiPlaneraResaWebV4ModelsJourneysLinkEndpointApiModel, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetOrigin(v VTApiPlaneraResaWebV4ModelsJourneysLinkEndpointApiModel)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetDestination

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetDestination() VTApiPlaneraResaWebV4ModelsJourneysLinkEndpointApiModel`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetDestinationOk() (*VTApiPlaneraResaWebV4ModelsJourneysLinkEndpointApiModel, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetDestination(v VTApiPlaneraResaWebV4ModelsJourneysLinkEndpointApiModel)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetNotes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetNotes() []VTApiPlaneraResaCoreModelsNote`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetNotesOk() (*[]VTApiPlaneraResaCoreModelsNote, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetNotes(v []VTApiPlaneraResaCoreModelsNote)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetDistanceInMeters

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetDistanceInMeters() int32`

GetDistanceInMeters returns the DistanceInMeters field if non-nil, zero value otherwise.

### GetDistanceInMetersOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetDistanceInMetersOk() (*int32, bool)`

GetDistanceInMetersOk returns a tuple with the DistanceInMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceInMeters

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetDistanceInMeters(v int32)`

SetDistanceInMeters sets DistanceInMeters field to given value.

### HasDistanceInMeters

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) HasDistanceInMeters() bool`

HasDistanceInMeters returns a boolean if a field has been set.

### SetDistanceInMetersNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetDistanceInMetersNil(b bool)`

 SetDistanceInMetersNil sets the value for DistanceInMeters to be an explicit nil

### UnsetDistanceInMeters
`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) UnsetDistanceInMeters()`

UnsetDistanceInMeters ensures that no value is present for DistanceInMeters, not even an explicit nil
### GetPlannedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetPlannedDepartureTime() string`

GetPlannedDepartureTime returns the PlannedDepartureTime field if non-nil, zero value otherwise.

### GetPlannedDepartureTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetPlannedDepartureTimeOk() (*string, bool)`

GetPlannedDepartureTimeOk returns a tuple with the PlannedDepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetPlannedDepartureTime(v string)`

SetPlannedDepartureTime sets PlannedDepartureTime field to given value.

### HasPlannedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) HasPlannedDepartureTime() bool`

HasPlannedDepartureTime returns a boolean if a field has been set.

### SetPlannedDepartureTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetPlannedDepartureTimeNil(b bool)`

 SetPlannedDepartureTimeNil sets the value for PlannedDepartureTime to be an explicit nil

### UnsetPlannedDepartureTime
`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) UnsetPlannedDepartureTime()`

UnsetPlannedDepartureTime ensures that no value is present for PlannedDepartureTime, not even an explicit nil
### GetPlannedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetPlannedArrivalTime() string`

GetPlannedArrivalTime returns the PlannedArrivalTime field if non-nil, zero value otherwise.

### GetPlannedArrivalTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetPlannedArrivalTimeOk() (*string, bool)`

GetPlannedArrivalTimeOk returns a tuple with the PlannedArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetPlannedArrivalTime(v string)`

SetPlannedArrivalTime sets PlannedArrivalTime field to given value.

### HasPlannedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) HasPlannedArrivalTime() bool`

HasPlannedArrivalTime returns a boolean if a field has been set.

### SetPlannedArrivalTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetPlannedArrivalTimeNil(b bool)`

 SetPlannedArrivalTimeNil sets the value for PlannedArrivalTime to be an explicit nil

### UnsetPlannedArrivalTime
`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) UnsetPlannedArrivalTime()`

UnsetPlannedArrivalTime ensures that no value is present for PlannedArrivalTime, not even an explicit nil
### GetPlannedDurationInMinutes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetPlannedDurationInMinutes() int32`

GetPlannedDurationInMinutes returns the PlannedDurationInMinutes field if non-nil, zero value otherwise.

### GetPlannedDurationInMinutesOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetPlannedDurationInMinutesOk() (*int32, bool)`

GetPlannedDurationInMinutesOk returns a tuple with the PlannedDurationInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDurationInMinutes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetPlannedDurationInMinutes(v int32)`

SetPlannedDurationInMinutes sets PlannedDurationInMinutes field to given value.

### HasPlannedDurationInMinutes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) HasPlannedDurationInMinutes() bool`

HasPlannedDurationInMinutes returns a boolean if a field has been set.

### SetPlannedDurationInMinutesNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetPlannedDurationInMinutesNil(b bool)`

 SetPlannedDurationInMinutesNil sets the value for PlannedDurationInMinutes to be an explicit nil

### UnsetPlannedDurationInMinutes
`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) UnsetPlannedDurationInMinutes()`

UnsetPlannedDurationInMinutes ensures that no value is present for PlannedDurationInMinutes, not even an explicit nil
### GetEstimatedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetEstimatedDepartureTime() string`

GetEstimatedDepartureTime returns the EstimatedDepartureTime field if non-nil, zero value otherwise.

### GetEstimatedDepartureTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetEstimatedDepartureTimeOk() (*string, bool)`

GetEstimatedDepartureTimeOk returns a tuple with the EstimatedDepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetEstimatedDepartureTime(v string)`

SetEstimatedDepartureTime sets EstimatedDepartureTime field to given value.

### HasEstimatedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) HasEstimatedDepartureTime() bool`

HasEstimatedDepartureTime returns a boolean if a field has been set.

### SetEstimatedDepartureTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetEstimatedDepartureTimeNil(b bool)`

 SetEstimatedDepartureTimeNil sets the value for EstimatedDepartureTime to be an explicit nil

### UnsetEstimatedDepartureTime
`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) UnsetEstimatedDepartureTime()`

UnsetEstimatedDepartureTime ensures that no value is present for EstimatedDepartureTime, not even an explicit nil
### GetEstimatedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetEstimatedArrivalTime() string`

GetEstimatedArrivalTime returns the EstimatedArrivalTime field if non-nil, zero value otherwise.

### GetEstimatedArrivalTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetEstimatedArrivalTimeOk() (*string, bool)`

GetEstimatedArrivalTimeOk returns a tuple with the EstimatedArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetEstimatedArrivalTime(v string)`

SetEstimatedArrivalTime sets EstimatedArrivalTime field to given value.

### HasEstimatedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) HasEstimatedArrivalTime() bool`

HasEstimatedArrivalTime returns a boolean if a field has been set.

### SetEstimatedArrivalTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetEstimatedArrivalTimeNil(b bool)`

 SetEstimatedArrivalTimeNil sets the value for EstimatedArrivalTime to be an explicit nil

### UnsetEstimatedArrivalTime
`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) UnsetEstimatedArrivalTime()`

UnsetEstimatedArrivalTime ensures that no value is present for EstimatedArrivalTime, not even an explicit nil
### GetEstimatedDurationInMinutes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetEstimatedDurationInMinutes() int32`

GetEstimatedDurationInMinutes returns the EstimatedDurationInMinutes field if non-nil, zero value otherwise.

### GetEstimatedDurationInMinutesOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetEstimatedDurationInMinutesOk() (*int32, bool)`

GetEstimatedDurationInMinutesOk returns a tuple with the EstimatedDurationInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDurationInMinutes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetEstimatedDurationInMinutes(v int32)`

SetEstimatedDurationInMinutes sets EstimatedDurationInMinutes field to given value.

### HasEstimatedDurationInMinutes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) HasEstimatedDurationInMinutes() bool`

HasEstimatedDurationInMinutes returns a boolean if a field has been set.

### SetEstimatedDurationInMinutesNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetEstimatedDurationInMinutesNil(b bool)`

 SetEstimatedDurationInMinutesNil sets the value for EstimatedDurationInMinutes to be an explicit nil

### UnsetEstimatedDurationInMinutes
`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) UnsetEstimatedDurationInMinutes()`

UnsetEstimatedDurationInMinutes ensures that no value is present for EstimatedDurationInMinutes, not even an explicit nil
### GetEstimatedNumberOfSteps

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetEstimatedNumberOfSteps() int32`

GetEstimatedNumberOfSteps returns the EstimatedNumberOfSteps field if non-nil, zero value otherwise.

### GetEstimatedNumberOfStepsOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetEstimatedNumberOfStepsOk() (*int32, bool)`

GetEstimatedNumberOfStepsOk returns a tuple with the EstimatedNumberOfSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedNumberOfSteps

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetEstimatedNumberOfSteps(v int32)`

SetEstimatedNumberOfSteps sets EstimatedNumberOfSteps field to given value.

### HasEstimatedNumberOfSteps

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) HasEstimatedNumberOfSteps() bool`

HasEstimatedNumberOfSteps returns a boolean if a field has been set.

### SetEstimatedNumberOfStepsNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetEstimatedNumberOfStepsNil(b bool)`

 SetEstimatedNumberOfStepsNil sets the value for EstimatedNumberOfSteps to be an explicit nil

### UnsetEstimatedNumberOfSteps
`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) UnsetEstimatedNumberOfSteps()`

UnsetEstimatedNumberOfSteps ensures that no value is present for EstimatedNumberOfSteps, not even an explicit nil
### GetLinkCoordinates

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetLinkCoordinates() []VTApiPlaneraResaWebV4ModelsCoordinateApiModel`

GetLinkCoordinates returns the LinkCoordinates field if non-nil, zero value otherwise.

### GetLinkCoordinatesOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetLinkCoordinatesOk() (*[]VTApiPlaneraResaWebV4ModelsCoordinateApiModel, bool)`

GetLinkCoordinatesOk returns a tuple with the LinkCoordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkCoordinates

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetLinkCoordinates(v []VTApiPlaneraResaWebV4ModelsCoordinateApiModel)`

SetLinkCoordinates sets LinkCoordinates field to given value.

### HasLinkCoordinates

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) HasLinkCoordinates() bool`

HasLinkCoordinates returns a boolean if a field has been set.

### SetLinkCoordinatesNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetLinkCoordinatesNil(b bool)`

 SetLinkCoordinatesNil sets the value for LinkCoordinates to be an explicit nil

### UnsetLinkCoordinates
`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) UnsetLinkCoordinates()`

UnsetLinkCoordinates ensures that no value is present for LinkCoordinates, not even an explicit nil
### GetSegments

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetSegments() []VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel`

GetSegments returns the Segments field if non-nil, zero value otherwise.

### GetSegmentsOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) GetSegmentsOk() (*[]VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel, bool)`

GetSegmentsOk returns a tuple with the Segments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegments

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetSegments(v []VTApiPlaneraResaWebV4ModelsJourneysLinkSegmentApiModel)`

SetSegments sets Segments field to given value.

### HasSegments

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) HasSegments() bool`

HasSegments returns a boolean if a field has been set.

### SetSegmentsNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) SetSegmentsNil(b bool)`

 SetSegmentsNil sets the value for Segments to be an explicit nil

### UnsetSegments
`func (o *VTApiPlaneraResaWebV4ModelsJourneysDestinationLinkApiModel) UnsetSegments()`

UnsetSegments ensures that no value is present for Segments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


