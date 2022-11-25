# GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VisitStatus** | Pointer to **int32** |  | [optional] 
**Network** | Pointer to [**GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfoNetwork**](GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfoNetwork.md) |  | [optional] 
**Categories** | Pointer to [**[]GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfoCategoriesInner**](GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfoCategoriesInner.md) |  | [optional] 

## Methods

### NewGetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo

`func NewGetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo() *GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo`

NewGetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo instantiates a new GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfoWithDefaults

`func NewGetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfoWithDefaults() *GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo`

NewGetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfoWithDefaults instantiates a new GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVisitStatus

`func (o *GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo) GetVisitStatus() int32`

GetVisitStatus returns the VisitStatus field if non-nil, zero value otherwise.

### GetVisitStatusOk

`func (o *GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo) GetVisitStatusOk() (*int32, bool)`

GetVisitStatusOk returns a tuple with the VisitStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisitStatus

`func (o *GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo) SetVisitStatus(v int32)`

SetVisitStatus sets VisitStatus field to given value.

### HasVisitStatus

`func (o *GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo) HasVisitStatus() bool`

HasVisitStatus returns a boolean if a field has been set.

### GetNetwork

`func (o *GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo) GetNetwork() GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfoNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo) GetNetworkOk() (*GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfoNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo) SetNetwork(v GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfoNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetCategories

`func (o *GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo) GetCategories() []GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfoCategoriesInner`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo) GetCategoriesOk() (*[]GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfoCategoriesInner, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo) SetCategories(v []GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfoCategoriesInner)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


