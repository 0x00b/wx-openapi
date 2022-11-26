# GetAuthorizerInfoResponseAuthorizationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationAppid** | Pointer to **string** | 授权方 appid | [optional] 
**FuncInfo** | Pointer to [**[]GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner**](GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner.md) | 授权给开发者的权限集列表 | [optional] 

## Methods

### NewGetAuthorizerInfoResponseAuthorizationInfo

`func NewGetAuthorizerInfoResponseAuthorizationInfo() *GetAuthorizerInfoResponseAuthorizationInfo`

NewGetAuthorizerInfoResponseAuthorizationInfo instantiates a new GetAuthorizerInfoResponseAuthorizationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAuthorizerInfoResponseAuthorizationInfoWithDefaults

`func NewGetAuthorizerInfoResponseAuthorizationInfoWithDefaults() *GetAuthorizerInfoResponseAuthorizationInfo`

NewGetAuthorizerInfoResponseAuthorizationInfoWithDefaults instantiates a new GetAuthorizerInfoResponseAuthorizationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationAppid

`func (o *GetAuthorizerInfoResponseAuthorizationInfo) GetAuthorizationAppid() string`

GetAuthorizationAppid returns the AuthorizationAppid field if non-nil, zero value otherwise.

### GetAuthorizationAppidOk

`func (o *GetAuthorizerInfoResponseAuthorizationInfo) GetAuthorizationAppidOk() (*string, bool)`

GetAuthorizationAppidOk returns a tuple with the AuthorizationAppid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationAppid

`func (o *GetAuthorizerInfoResponseAuthorizationInfo) SetAuthorizationAppid(v string)`

SetAuthorizationAppid sets AuthorizationAppid field to given value.

### HasAuthorizationAppid

`func (o *GetAuthorizerInfoResponseAuthorizationInfo) HasAuthorizationAppid() bool`

HasAuthorizationAppid returns a boolean if a field has been set.

### GetFuncInfo

`func (o *GetAuthorizerInfoResponseAuthorizationInfo) GetFuncInfo() []GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner`

GetFuncInfo returns the FuncInfo field if non-nil, zero value otherwise.

### GetFuncInfoOk

`func (o *GetAuthorizerInfoResponseAuthorizationInfo) GetFuncInfoOk() (*[]GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner, bool)`

GetFuncInfoOk returns a tuple with the FuncInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncInfo

`func (o *GetAuthorizerInfoResponseAuthorizationInfo) SetFuncInfo(v []GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner)`

SetFuncInfo sets FuncInfo field to given value.

### HasFuncInfo

`func (o *GetAuthorizerInfoResponseAuthorizationInfo) HasFuncInfo() bool`

HasFuncInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


