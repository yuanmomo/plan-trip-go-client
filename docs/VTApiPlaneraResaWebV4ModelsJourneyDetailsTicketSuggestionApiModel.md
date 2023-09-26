# VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **int32** | The product id. | [optional] 
**ProductName** | Pointer to **NullableString** | The product name. | [optional] 
**ProductType** | Pointer to **int32** | The product type. | [optional] 
**TravellerCategory** | Pointer to [**VTApiPlaneraResaCoreModelsTravellerCategory**](VTApiPlaneraResaCoreModelsTravellerCategory.md) |  | [optional] 
**PriceInSek** | Pointer to **float64** | The product price in SEK. | [optional] 
**TimeValidity** | Pointer to [**VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel**](VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel.md) |  | [optional] 
**TimeLimitation** | Pointer to [**VTApiPlaneraResaCoreModelsTimeLimitation**](VTApiPlaneraResaCoreModelsTimeLimitation.md) |  | [optional] 
**SaleChannels** | Pointer to [**[]VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel**](VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel.md) | A list of the channels that sell the product. | [optional] 
**ValidZones** | Pointer to [**[]VTApiPlaneraResaWebV4ModelsJourneyDetailsZoneApiModel**](VTApiPlaneraResaWebV4ModelsJourneyDetailsZoneApiModel.md) | A list of the valid zones for the ticket. | [optional] 
**ProductInstanceType** | Pointer to [**VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel**](VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel.md) |  | [optional] 
**PunchConfiguration** | Pointer to [**VTApiPlaneraResaWebV4ModelsJourneyDetailsPunchConfigurationApiModel**](VTApiPlaneraResaWebV4ModelsJourneyDetailsPunchConfigurationApiModel.md) |  | [optional] 
**OfferSpecification** | Pointer to **NullableString** | Used to get ticket offer. | [optional] 

## Methods

### NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel

`func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel() *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModelWithDefaults

`func NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModelWithDefaults() *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel`

NewVTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModelWithDefaults instantiates a new VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) GetProductId() int32`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) GetProductIdOk() (*int32, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) SetProductId(v int32)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### GetProductName

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) SetProductName(v string)`

SetProductName sets ProductName field to given value.

### HasProductName

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) HasProductName() bool`

HasProductName returns a boolean if a field has been set.

### SetProductNameNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) SetProductNameNil(b bool)`

 SetProductNameNil sets the value for ProductName to be an explicit nil

### UnsetProductName
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) UnsetProductName()`

UnsetProductName ensures that no value is present for ProductName, not even an explicit nil
### GetProductType

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) GetProductType() int32`

GetProductType returns the ProductType field if non-nil, zero value otherwise.

### GetProductTypeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) GetProductTypeOk() (*int32, bool)`

GetProductTypeOk returns a tuple with the ProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductType

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) SetProductType(v int32)`

SetProductType sets ProductType field to given value.

### HasProductType

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) HasProductType() bool`

HasProductType returns a boolean if a field has been set.

### GetTravellerCategory

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) GetTravellerCategory() VTApiPlaneraResaCoreModelsTravellerCategory`

GetTravellerCategory returns the TravellerCategory field if non-nil, zero value otherwise.

### GetTravellerCategoryOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) GetTravellerCategoryOk() (*VTApiPlaneraResaCoreModelsTravellerCategory, bool)`

GetTravellerCategoryOk returns a tuple with the TravellerCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravellerCategory

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) SetTravellerCategory(v VTApiPlaneraResaCoreModelsTravellerCategory)`

SetTravellerCategory sets TravellerCategory field to given value.

### HasTravellerCategory

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) HasTravellerCategory() bool`

HasTravellerCategory returns a boolean if a field has been set.

### GetPriceInSek

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) GetPriceInSek() float64`

GetPriceInSek returns the PriceInSek field if non-nil, zero value otherwise.

### GetPriceInSekOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) GetPriceInSekOk() (*float64, bool)`

GetPriceInSekOk returns a tuple with the PriceInSek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceInSek

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) SetPriceInSek(v float64)`

SetPriceInSek sets PriceInSek field to given value.

### HasPriceInSek

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) HasPriceInSek() bool`

HasPriceInSek returns a boolean if a field has been set.

### GetTimeValidity

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) GetTimeValidity() VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel`

GetTimeValidity returns the TimeValidity field if non-nil, zero value otherwise.

### GetTimeValidityOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) GetTimeValidityOk() (*VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel, bool)`

GetTimeValidityOk returns a tuple with the TimeValidity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeValidity

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) SetTimeValidity(v VTApiPlaneraResaWebV4ModelsJourneyDetailsTimeValidityApiModel)`

SetTimeValidity sets TimeValidity field to given value.

### HasTimeValidity

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) HasTimeValidity() bool`

HasTimeValidity returns a boolean if a field has been set.

### GetTimeLimitation

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) GetTimeLimitation() VTApiPlaneraResaCoreModelsTimeLimitation`

GetTimeLimitation returns the TimeLimitation field if non-nil, zero value otherwise.

### GetTimeLimitationOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) GetTimeLimitationOk() (*VTApiPlaneraResaCoreModelsTimeLimitation, bool)`

GetTimeLimitationOk returns a tuple with the TimeLimitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimitation

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) SetTimeLimitation(v VTApiPlaneraResaCoreModelsTimeLimitation)`

SetTimeLimitation sets TimeLimitation field to given value.

### HasTimeLimitation

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) HasTimeLimitation() bool`

HasTimeLimitation returns a boolean if a field has been set.

### GetSaleChannels

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) GetSaleChannels() []VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel`

GetSaleChannels returns the SaleChannels field if non-nil, zero value otherwise.

### GetSaleChannelsOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) GetSaleChannelsOk() (*[]VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel, bool)`

GetSaleChannelsOk returns a tuple with the SaleChannels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleChannels

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) SetSaleChannels(v []VTApiPlaneraResaWebV4ModelsJourneyDetailsChannelApiModel)`

SetSaleChannels sets SaleChannels field to given value.

### HasSaleChannels

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) HasSaleChannels() bool`

HasSaleChannels returns a boolean if a field has been set.

### SetSaleChannelsNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) SetSaleChannelsNil(b bool)`

 SetSaleChannelsNil sets the value for SaleChannels to be an explicit nil

### UnsetSaleChannels
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) UnsetSaleChannels()`

UnsetSaleChannels ensures that no value is present for SaleChannels, not even an explicit nil
### GetValidZones

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) GetValidZones() []VTApiPlaneraResaWebV4ModelsJourneyDetailsZoneApiModel`

GetValidZones returns the ValidZones field if non-nil, zero value otherwise.

### GetValidZonesOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) GetValidZonesOk() (*[]VTApiPlaneraResaWebV4ModelsJourneyDetailsZoneApiModel, bool)`

GetValidZonesOk returns a tuple with the ValidZones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidZones

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) SetValidZones(v []VTApiPlaneraResaWebV4ModelsJourneyDetailsZoneApiModel)`

SetValidZones sets ValidZones field to given value.

### HasValidZones

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) HasValidZones() bool`

HasValidZones returns a boolean if a field has been set.

### SetValidZonesNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) SetValidZonesNil(b bool)`

 SetValidZonesNil sets the value for ValidZones to be an explicit nil

### UnsetValidZones
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) UnsetValidZones()`

UnsetValidZones ensures that no value is present for ValidZones, not even an explicit nil
### GetProductInstanceType

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) GetProductInstanceType() VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel`

GetProductInstanceType returns the ProductInstanceType field if non-nil, zero value otherwise.

### GetProductInstanceTypeOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) GetProductInstanceTypeOk() (*VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel, bool)`

GetProductInstanceTypeOk returns a tuple with the ProductInstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductInstanceType

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) SetProductInstanceType(v VTApiPlaneraResaWebV4ModelsProductInstanceTypeApiModel)`

SetProductInstanceType sets ProductInstanceType field to given value.

### HasProductInstanceType

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) HasProductInstanceType() bool`

HasProductInstanceType returns a boolean if a field has been set.

### GetPunchConfiguration

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) GetPunchConfiguration() VTApiPlaneraResaWebV4ModelsJourneyDetailsPunchConfigurationApiModel`

GetPunchConfiguration returns the PunchConfiguration field if non-nil, zero value otherwise.

### GetPunchConfigurationOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) GetPunchConfigurationOk() (*VTApiPlaneraResaWebV4ModelsJourneyDetailsPunchConfigurationApiModel, bool)`

GetPunchConfigurationOk returns a tuple with the PunchConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPunchConfiguration

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) SetPunchConfiguration(v VTApiPlaneraResaWebV4ModelsJourneyDetailsPunchConfigurationApiModel)`

SetPunchConfiguration sets PunchConfiguration field to given value.

### HasPunchConfiguration

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) HasPunchConfiguration() bool`

HasPunchConfiguration returns a boolean if a field has been set.

### GetOfferSpecification

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) GetOfferSpecification() string`

GetOfferSpecification returns the OfferSpecification field if non-nil, zero value otherwise.

### GetOfferSpecificationOk

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) GetOfferSpecificationOk() (*string, bool)`

GetOfferSpecificationOk returns a tuple with the OfferSpecification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferSpecification

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) SetOfferSpecification(v string)`

SetOfferSpecification sets OfferSpecification field to given value.

### HasOfferSpecification

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) HasOfferSpecification() bool`

HasOfferSpecification returns a boolean if a field has been set.

### SetOfferSpecificationNil

`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) SetOfferSpecificationNil(b bool)`

 SetOfferSpecificationNil sets the value for OfferSpecification to be an explicit nil

### UnsetOfferSpecification
`func (o *VTApiPlaneraResaWebV4ModelsJourneyDetailsTicketSuggestionApiModel) UnsetOfferSpecification()`

UnsetOfferSpecification ensures that no value is present for OfferSpecification, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


