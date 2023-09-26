# VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasError** | Pointer to **bool** | Flag indicating that an error occurred while getting ticket suggestions. | [optional] 
**TicketSuggestions** | Pointer to [**[]VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel**](VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel.md) | Ticket suggestions for a journey. | [optional] 
**TicketValidities** | Pointer to [**[]VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketValidityApiModel**](VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketValidityApiModel.md) | An array with the tickets from the existingTickets-array in the post-body. Validity for the journey is indicated for each ticket in the array. Included if &#39;ticketsuggestions&#39; is passed in the includes array in the request, otherwise null. | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel

`func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel() *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasError

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel) GetHasError() bool`

GetHasError returns the HasError field if non-nil, zero value otherwise.

### GetHasErrorOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel) GetHasErrorOk() (*bool, bool)`

GetHasErrorOk returns a tuple with the HasError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasError

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel) SetHasError(v bool)`

SetHasError sets HasError field to given value.

### HasHasError

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel) HasHasError() bool`

HasHasError returns a boolean if a field has been set.

### GetTicketSuggestions

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel) GetTicketSuggestions() []VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel`

GetTicketSuggestions returns the TicketSuggestions field if non-nil, zero value otherwise.

### GetTicketSuggestionsOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel) GetTicketSuggestionsOk() (*[]VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel, bool)`

GetTicketSuggestionsOk returns a tuple with the TicketSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketSuggestions

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel) SetTicketSuggestions(v []VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel)`

SetTicketSuggestions sets TicketSuggestions field to given value.

### HasTicketSuggestions

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel) HasTicketSuggestions() bool`

HasTicketSuggestions returns a boolean if a field has been set.

### SetTicketSuggestionsNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel) SetTicketSuggestionsNil(b bool)`

 SetTicketSuggestionsNil sets the value for TicketSuggestions to be an explicit nil

### UnsetTicketSuggestions
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel) UnsetTicketSuggestions()`

UnsetTicketSuggestions ensures that no value is present for TicketSuggestions, not even an explicit nil
### GetTicketValidities

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel) GetTicketValidities() []VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketValidityApiModel`

GetTicketValidities returns the TicketValidities field if non-nil, zero value otherwise.

### GetTicketValiditiesOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel) GetTicketValiditiesOk() (*[]VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketValidityApiModel, bool)`

GetTicketValiditiesOk returns a tuple with the TicketValidities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketValidities

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel) SetTicketValidities(v []VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketValidityApiModel)`

SetTicketValidities sets TicketValidities field to given value.

### HasTicketValidities

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel) HasTicketValidities() bool`

HasTicketValidities returns a boolean if a field has been set.

### SetTicketValiditiesNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel) SetTicketValiditiesNil(b bool)`

 SetTicketValiditiesNil sets the value for TicketValidities to be an explicit nil

### UnsetTicketValidities
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionsResultApiModel) UnsetTicketValidities()`

UnsetTicketValidities ensures that no value is present for TicketValidities, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


