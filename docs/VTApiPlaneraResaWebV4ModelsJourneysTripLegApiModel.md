# VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Origin** | [**VTApiPlaneraResaWebV4ModelsJourneysCallApiModel**](VTApiPlaneraResaWebV4ModelsJourneysCallApiModel.md) |  | 
**Destination** | [**VTApiPlaneraResaWebV4ModelsJourneysCallApiModel**](VTApiPlaneraResaWebV4ModelsJourneysCallApiModel.md) |  | 
**IsCancelled** | **bool** | Flag indicating if the trip leg is cancelled. | 
**IsPartCancelled** | Pointer to **bool** | Flag indicating if the trip leg is partially cancelled. | [optional] 
**ServiceJourney** | Pointer to [**VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel**](VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel.md) |  | [optional] 
**Notes** | Pointer to [**[]VTApiPlaneraResaCoreModelsNote**](VTApiPlaneraResaCoreModelsNote.md) | An ordered list (most important first) of notes related to the trip leg. | [optional] 
**EstimatedDistanceInMeters** | Pointer to **NullableInt32** | Estimated distance in meters. Only for transport mode Walk. | [optional] 
**PlannedConnectingTimeInMinutes** | Pointer to **NullableInt32** | The planned (according to timetable) connecting time in minutes relative to  the previous public transport trip leg (if any). | [optional] 
**EstimatedConnectingTimeInMinutes** | Pointer to **NullableInt32** | The estimated (according to real-time data) connecting time in minutes relative to  the previous public transport trip leg (if any). | [optional] 
**IsRiskOfMissingConnection** | Pointer to **NullableBool** | Flag indicating that there is less than 5 minutes margin between arriving at the  origin stop point and the departure from that stop point. | [optional] 
**PlannedDepartureTime** | Pointer to **NullableString** | The planned departure time in RFC 3339 format. | [optional] 
**PlannedArrivalTime** | Pointer to **NullableString** | The planned arrival time in RFC 3339 format. | [optional] 
**PlannedDurationInMinutes** | Pointer to **NullableInt32** | The planned duration in minutes. | [optional] 
**EstimatedDepartureTime** | Pointer to **NullableString** | The estimated departure time in RFC 3339 format, if available. | [optional] 
**EstimatedArrivalTime** | Pointer to **NullableString** | The estimated arrival time in RFC 3339 format, if available. | [optional] 
**EstimatedDurationInMinutes** | Pointer to **NullableInt32** | The estimated duration in minutes, if available. | [optional] 
**EstimatedOtherwisePlannedArrivalTime** | Pointer to **NullableString** | The best known time of the arrival in RFC 3339 format. Is EstimatedArrivalTime if exists, otherwise PlannedArrivalTime. | [optional] [readonly] 
**EstimatedOtherwisePlannedDepartureTime** | Pointer to **NullableString** | The best known time of the departure in RFC 3339 format. Is EstimatedDepartureTime if exists, otherwise PlannedDepartureTime. | [optional] [readonly] 
**Occupancy** | Pointer to [**VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel**](VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel.md) |  | [optional] 
**JourneyLegIndex** | Pointer to **int32** | Index of Leg in Journey | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel

`func NewVTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel(origin VTApiPlaneraResaWebV4ModelsJourneysCallApiModel, destination VTApiPlaneraResaWebV4ModelsJourneysCallApiModel, isCancelled bool, ) *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel instantiates a new VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsJourneysTripLegApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsJourneysTripLegApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneysTripLegApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrigin

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetOrigin() VTApiPlaneraResaWebV4ModelsJourneysCallApiModel`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetOriginOk() (*VTApiPlaneraResaWebV4ModelsJourneysCallApiModel, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetOrigin(v VTApiPlaneraResaWebV4ModelsJourneysCallApiModel)`

SetOrigin sets Origin field to given value.


### GetDestination

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetDestination() VTApiPlaneraResaWebV4ModelsJourneysCallApiModel`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetDestinationOk() (*VTApiPlaneraResaWebV4ModelsJourneysCallApiModel, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetDestination(v VTApiPlaneraResaWebV4ModelsJourneysCallApiModel)`

SetDestination sets Destination field to given value.


### GetIsCancelled

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetIsCancelled() bool`

GetIsCancelled returns the IsCancelled field if non-nil, zero value otherwise.

### GetIsCancelledOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetIsCancelledOk() (*bool, bool)`

GetIsCancelledOk returns a tuple with the IsCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCancelled

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetIsCancelled(v bool)`

SetIsCancelled sets IsCancelled field to given value.


### GetIsPartCancelled

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetIsPartCancelled() bool`

GetIsPartCancelled returns the IsPartCancelled field if non-nil, zero value otherwise.

### GetIsPartCancelledOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetIsPartCancelledOk() (*bool, bool)`

GetIsPartCancelledOk returns a tuple with the IsPartCancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPartCancelled

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetIsPartCancelled(v bool)`

SetIsPartCancelled sets IsPartCancelled field to given value.

### HasIsPartCancelled

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) HasIsPartCancelled() bool`

HasIsPartCancelled returns a boolean if a field has been set.

### GetServiceJourney

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetServiceJourney() VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel`

GetServiceJourney returns the ServiceJourney field if non-nil, zero value otherwise.

### GetServiceJourneyOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetServiceJourneyOk() (*VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel, bool)`

GetServiceJourneyOk returns a tuple with the ServiceJourney field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceJourney

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetServiceJourney(v VTApiPlaneraResaWebV4ModelsJourneysServiceJourneyApiModel)`

SetServiceJourney sets ServiceJourney field to given value.

### HasServiceJourney

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) HasServiceJourney() bool`

HasServiceJourney returns a boolean if a field has been set.

### GetNotes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetNotes() []VTApiPlaneraResaCoreModelsNote`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetNotesOk() (*[]VTApiPlaneraResaCoreModelsNote, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetNotes(v []VTApiPlaneraResaCoreModelsNote)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetEstimatedDistanceInMeters

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetEstimatedDistanceInMeters() int32`

GetEstimatedDistanceInMeters returns the EstimatedDistanceInMeters field if non-nil, zero value otherwise.

### GetEstimatedDistanceInMetersOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetEstimatedDistanceInMetersOk() (*int32, bool)`

GetEstimatedDistanceInMetersOk returns a tuple with the EstimatedDistanceInMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDistanceInMeters

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetEstimatedDistanceInMeters(v int32)`

SetEstimatedDistanceInMeters sets EstimatedDistanceInMeters field to given value.

### HasEstimatedDistanceInMeters

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) HasEstimatedDistanceInMeters() bool`

HasEstimatedDistanceInMeters returns a boolean if a field has been set.

### SetEstimatedDistanceInMetersNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetEstimatedDistanceInMetersNil(b bool)`

 SetEstimatedDistanceInMetersNil sets the value for EstimatedDistanceInMeters to be an explicit nil

### UnsetEstimatedDistanceInMeters
`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) UnsetEstimatedDistanceInMeters()`

UnsetEstimatedDistanceInMeters ensures that no value is present for EstimatedDistanceInMeters, not even an explicit nil
### GetPlannedConnectingTimeInMinutes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetPlannedConnectingTimeInMinutes() int32`

GetPlannedConnectingTimeInMinutes returns the PlannedConnectingTimeInMinutes field if non-nil, zero value otherwise.

### GetPlannedConnectingTimeInMinutesOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetPlannedConnectingTimeInMinutesOk() (*int32, bool)`

GetPlannedConnectingTimeInMinutesOk returns a tuple with the PlannedConnectingTimeInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedConnectingTimeInMinutes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetPlannedConnectingTimeInMinutes(v int32)`

SetPlannedConnectingTimeInMinutes sets PlannedConnectingTimeInMinutes field to given value.

### HasPlannedConnectingTimeInMinutes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) HasPlannedConnectingTimeInMinutes() bool`

HasPlannedConnectingTimeInMinutes returns a boolean if a field has been set.

### SetPlannedConnectingTimeInMinutesNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetPlannedConnectingTimeInMinutesNil(b bool)`

 SetPlannedConnectingTimeInMinutesNil sets the value for PlannedConnectingTimeInMinutes to be an explicit nil

### UnsetPlannedConnectingTimeInMinutes
`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) UnsetPlannedConnectingTimeInMinutes()`

UnsetPlannedConnectingTimeInMinutes ensures that no value is present for PlannedConnectingTimeInMinutes, not even an explicit nil
### GetEstimatedConnectingTimeInMinutes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetEstimatedConnectingTimeInMinutes() int32`

GetEstimatedConnectingTimeInMinutes returns the EstimatedConnectingTimeInMinutes field if non-nil, zero value otherwise.

### GetEstimatedConnectingTimeInMinutesOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetEstimatedConnectingTimeInMinutesOk() (*int32, bool)`

GetEstimatedConnectingTimeInMinutesOk returns a tuple with the EstimatedConnectingTimeInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedConnectingTimeInMinutes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetEstimatedConnectingTimeInMinutes(v int32)`

SetEstimatedConnectingTimeInMinutes sets EstimatedConnectingTimeInMinutes field to given value.

### HasEstimatedConnectingTimeInMinutes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) HasEstimatedConnectingTimeInMinutes() bool`

HasEstimatedConnectingTimeInMinutes returns a boolean if a field has been set.

### SetEstimatedConnectingTimeInMinutesNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetEstimatedConnectingTimeInMinutesNil(b bool)`

 SetEstimatedConnectingTimeInMinutesNil sets the value for EstimatedConnectingTimeInMinutes to be an explicit nil

### UnsetEstimatedConnectingTimeInMinutes
`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) UnsetEstimatedConnectingTimeInMinutes()`

UnsetEstimatedConnectingTimeInMinutes ensures that no value is present for EstimatedConnectingTimeInMinutes, not even an explicit nil
### GetIsRiskOfMissingConnection

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetIsRiskOfMissingConnection() bool`

GetIsRiskOfMissingConnection returns the IsRiskOfMissingConnection field if non-nil, zero value otherwise.

### GetIsRiskOfMissingConnectionOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetIsRiskOfMissingConnectionOk() (*bool, bool)`

GetIsRiskOfMissingConnectionOk returns a tuple with the IsRiskOfMissingConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRiskOfMissingConnection

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetIsRiskOfMissingConnection(v bool)`

SetIsRiskOfMissingConnection sets IsRiskOfMissingConnection field to given value.

### HasIsRiskOfMissingConnection

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) HasIsRiskOfMissingConnection() bool`

HasIsRiskOfMissingConnection returns a boolean if a field has been set.

### SetIsRiskOfMissingConnectionNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetIsRiskOfMissingConnectionNil(b bool)`

 SetIsRiskOfMissingConnectionNil sets the value for IsRiskOfMissingConnection to be an explicit nil

### UnsetIsRiskOfMissingConnection
`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) UnsetIsRiskOfMissingConnection()`

UnsetIsRiskOfMissingConnection ensures that no value is present for IsRiskOfMissingConnection, not even an explicit nil
### GetPlannedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetPlannedDepartureTime() string`

GetPlannedDepartureTime returns the PlannedDepartureTime field if non-nil, zero value otherwise.

### GetPlannedDepartureTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetPlannedDepartureTimeOk() (*string, bool)`

GetPlannedDepartureTimeOk returns a tuple with the PlannedDepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetPlannedDepartureTime(v string)`

SetPlannedDepartureTime sets PlannedDepartureTime field to given value.

### HasPlannedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) HasPlannedDepartureTime() bool`

HasPlannedDepartureTime returns a boolean if a field has been set.

### SetPlannedDepartureTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetPlannedDepartureTimeNil(b bool)`

 SetPlannedDepartureTimeNil sets the value for PlannedDepartureTime to be an explicit nil

### UnsetPlannedDepartureTime
`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) UnsetPlannedDepartureTime()`

UnsetPlannedDepartureTime ensures that no value is present for PlannedDepartureTime, not even an explicit nil
### GetPlannedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetPlannedArrivalTime() string`

GetPlannedArrivalTime returns the PlannedArrivalTime field if non-nil, zero value otherwise.

### GetPlannedArrivalTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetPlannedArrivalTimeOk() (*string, bool)`

GetPlannedArrivalTimeOk returns a tuple with the PlannedArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetPlannedArrivalTime(v string)`

SetPlannedArrivalTime sets PlannedArrivalTime field to given value.

### HasPlannedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) HasPlannedArrivalTime() bool`

HasPlannedArrivalTime returns a boolean if a field has been set.

### SetPlannedArrivalTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetPlannedArrivalTimeNil(b bool)`

 SetPlannedArrivalTimeNil sets the value for PlannedArrivalTime to be an explicit nil

### UnsetPlannedArrivalTime
`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) UnsetPlannedArrivalTime()`

UnsetPlannedArrivalTime ensures that no value is present for PlannedArrivalTime, not even an explicit nil
### GetPlannedDurationInMinutes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetPlannedDurationInMinutes() int32`

GetPlannedDurationInMinutes returns the PlannedDurationInMinutes field if non-nil, zero value otherwise.

### GetPlannedDurationInMinutesOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetPlannedDurationInMinutesOk() (*int32, bool)`

GetPlannedDurationInMinutesOk returns a tuple with the PlannedDurationInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannedDurationInMinutes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetPlannedDurationInMinutes(v int32)`

SetPlannedDurationInMinutes sets PlannedDurationInMinutes field to given value.

### HasPlannedDurationInMinutes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) HasPlannedDurationInMinutes() bool`

HasPlannedDurationInMinutes returns a boolean if a field has been set.

### SetPlannedDurationInMinutesNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetPlannedDurationInMinutesNil(b bool)`

 SetPlannedDurationInMinutesNil sets the value for PlannedDurationInMinutes to be an explicit nil

### UnsetPlannedDurationInMinutes
`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) UnsetPlannedDurationInMinutes()`

UnsetPlannedDurationInMinutes ensures that no value is present for PlannedDurationInMinutes, not even an explicit nil
### GetEstimatedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetEstimatedDepartureTime() string`

GetEstimatedDepartureTime returns the EstimatedDepartureTime field if non-nil, zero value otherwise.

### GetEstimatedDepartureTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetEstimatedDepartureTimeOk() (*string, bool)`

GetEstimatedDepartureTimeOk returns a tuple with the EstimatedDepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetEstimatedDepartureTime(v string)`

SetEstimatedDepartureTime sets EstimatedDepartureTime field to given value.

### HasEstimatedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) HasEstimatedDepartureTime() bool`

HasEstimatedDepartureTime returns a boolean if a field has been set.

### SetEstimatedDepartureTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetEstimatedDepartureTimeNil(b bool)`

 SetEstimatedDepartureTimeNil sets the value for EstimatedDepartureTime to be an explicit nil

### UnsetEstimatedDepartureTime
`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) UnsetEstimatedDepartureTime()`

UnsetEstimatedDepartureTime ensures that no value is present for EstimatedDepartureTime, not even an explicit nil
### GetEstimatedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetEstimatedArrivalTime() string`

GetEstimatedArrivalTime returns the EstimatedArrivalTime field if non-nil, zero value otherwise.

### GetEstimatedArrivalTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetEstimatedArrivalTimeOk() (*string, bool)`

GetEstimatedArrivalTimeOk returns a tuple with the EstimatedArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetEstimatedArrivalTime(v string)`

SetEstimatedArrivalTime sets EstimatedArrivalTime field to given value.

### HasEstimatedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) HasEstimatedArrivalTime() bool`

HasEstimatedArrivalTime returns a boolean if a field has been set.

### SetEstimatedArrivalTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetEstimatedArrivalTimeNil(b bool)`

 SetEstimatedArrivalTimeNil sets the value for EstimatedArrivalTime to be an explicit nil

### UnsetEstimatedArrivalTime
`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) UnsetEstimatedArrivalTime()`

UnsetEstimatedArrivalTime ensures that no value is present for EstimatedArrivalTime, not even an explicit nil
### GetEstimatedDurationInMinutes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetEstimatedDurationInMinutes() int32`

GetEstimatedDurationInMinutes returns the EstimatedDurationInMinutes field if non-nil, zero value otherwise.

### GetEstimatedDurationInMinutesOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetEstimatedDurationInMinutesOk() (*int32, bool)`

GetEstimatedDurationInMinutesOk returns a tuple with the EstimatedDurationInMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDurationInMinutes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetEstimatedDurationInMinutes(v int32)`

SetEstimatedDurationInMinutes sets EstimatedDurationInMinutes field to given value.

### HasEstimatedDurationInMinutes

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) HasEstimatedDurationInMinutes() bool`

HasEstimatedDurationInMinutes returns a boolean if a field has been set.

### SetEstimatedDurationInMinutesNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetEstimatedDurationInMinutesNil(b bool)`

 SetEstimatedDurationInMinutesNil sets the value for EstimatedDurationInMinutes to be an explicit nil

### UnsetEstimatedDurationInMinutes
`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) UnsetEstimatedDurationInMinutes()`

UnsetEstimatedDurationInMinutes ensures that no value is present for EstimatedDurationInMinutes, not even an explicit nil
### GetEstimatedOtherwisePlannedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetEstimatedOtherwisePlannedArrivalTime() string`

GetEstimatedOtherwisePlannedArrivalTime returns the EstimatedOtherwisePlannedArrivalTime field if non-nil, zero value otherwise.

### GetEstimatedOtherwisePlannedArrivalTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetEstimatedOtherwisePlannedArrivalTimeOk() (*string, bool)`

GetEstimatedOtherwisePlannedArrivalTimeOk returns a tuple with the EstimatedOtherwisePlannedArrivalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedOtherwisePlannedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetEstimatedOtherwisePlannedArrivalTime(v string)`

SetEstimatedOtherwisePlannedArrivalTime sets EstimatedOtherwisePlannedArrivalTime field to given value.

### HasEstimatedOtherwisePlannedArrivalTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) HasEstimatedOtherwisePlannedArrivalTime() bool`

HasEstimatedOtherwisePlannedArrivalTime returns a boolean if a field has been set.

### SetEstimatedOtherwisePlannedArrivalTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetEstimatedOtherwisePlannedArrivalTimeNil(b bool)`

 SetEstimatedOtherwisePlannedArrivalTimeNil sets the value for EstimatedOtherwisePlannedArrivalTime to be an explicit nil

### UnsetEstimatedOtherwisePlannedArrivalTime
`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) UnsetEstimatedOtherwisePlannedArrivalTime()`

UnsetEstimatedOtherwisePlannedArrivalTime ensures that no value is present for EstimatedOtherwisePlannedArrivalTime, not even an explicit nil
### GetEstimatedOtherwisePlannedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetEstimatedOtherwisePlannedDepartureTime() string`

GetEstimatedOtherwisePlannedDepartureTime returns the EstimatedOtherwisePlannedDepartureTime field if non-nil, zero value otherwise.

### GetEstimatedOtherwisePlannedDepartureTimeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetEstimatedOtherwisePlannedDepartureTimeOk() (*string, bool)`

GetEstimatedOtherwisePlannedDepartureTimeOk returns a tuple with the EstimatedOtherwisePlannedDepartureTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedOtherwisePlannedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetEstimatedOtherwisePlannedDepartureTime(v string)`

SetEstimatedOtherwisePlannedDepartureTime sets EstimatedOtherwisePlannedDepartureTime field to given value.

### HasEstimatedOtherwisePlannedDepartureTime

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) HasEstimatedOtherwisePlannedDepartureTime() bool`

HasEstimatedOtherwisePlannedDepartureTime returns a boolean if a field has been set.

### SetEstimatedOtherwisePlannedDepartureTimeNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetEstimatedOtherwisePlannedDepartureTimeNil(b bool)`

 SetEstimatedOtherwisePlannedDepartureTimeNil sets the value for EstimatedOtherwisePlannedDepartureTime to be an explicit nil

### UnsetEstimatedOtherwisePlannedDepartureTime
`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) UnsetEstimatedOtherwisePlannedDepartureTime()`

UnsetEstimatedOtherwisePlannedDepartureTime ensures that no value is present for EstimatedOtherwisePlannedDepartureTime, not even an explicit nil
### GetOccupancy

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetOccupancy() VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel`

GetOccupancy returns the Occupancy field if non-nil, zero value otherwise.

### GetOccupancyOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetOccupancyOk() (*VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel, bool)`

GetOccupancyOk returns a tuple with the Occupancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccupancy

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetOccupancy(v VTApiPlaneraResaWebV4ModelsOccupancyInformationApiModel)`

SetOccupancy sets Occupancy field to given value.

### HasOccupancy

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) HasOccupancy() bool`

HasOccupancy returns a boolean if a field has been set.

### GetJourneyLegIndex

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetJourneyLegIndex() int32`

GetJourneyLegIndex returns the JourneyLegIndex field if non-nil, zero value otherwise.

### GetJourneyLegIndexOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) GetJourneyLegIndexOk() (*int32, bool)`

GetJourneyLegIndexOk returns a tuple with the JourneyLegIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJourneyLegIndex

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) SetJourneyLegIndex(v int32)`

SetJourneyLegIndex sets JourneyLegIndex field to given value.

### HasJourneyLegIndex

`func (o *VTApiPlaneraResaWebV4ModelsJourneysTripLegApiModel) HasJourneyLegIndex() bool`

HasJourneyLegIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


