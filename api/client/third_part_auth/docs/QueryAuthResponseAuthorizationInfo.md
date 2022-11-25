# QueryAuthResponseAuthorizationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizerAppid** | Pointer to **string** | 授权方 appid | [optional] 
**AuthorizerAccessToken** | Pointer to **string** | 接口调用令牌（在授权的公众号/小程序具备 API 权限时，才有此返回值） | [optional] 
**ExpiresIn** | Pointer to **int32** | authorizer_access_token 的有效期（在授权的公众号/小程序具备API权限时，才有此返回值），单位：秒 | [optional] 
**AuthorizerRefreshToken** | Pointer to **string** | 刷新令牌（在授权的公众号具备API权限时，才有此返回值），刷新令牌主要用于第三方平台获取和刷新已授权用户的 authorizer_access_token。一旦丢失，只能让用户重新授权，才能再次拿到新的刷新令牌。用户重新授权后，之前的刷新令牌会失效 | [optional] 
**FuncInfo** | Pointer to [**[]GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner**](GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner.md) | 授权给开发者的权限集列表 | [optional] 

## Methods

### NewQueryAuthResponseAuthorizationInfo

`func NewQueryAuthResponseAuthorizationInfo() *QueryAuthResponseAuthorizationInfo`

NewQueryAuthResponseAuthorizationInfo instantiates a new QueryAuthResponseAuthorizationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryAuthResponseAuthorizationInfoWithDefaults

`func NewQueryAuthResponseAuthorizationInfoWithDefaults() *QueryAuthResponseAuthorizationInfo`

NewQueryAuthResponseAuthorizationInfoWithDefaults instantiates a new QueryAuthResponseAuthorizationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizerAppid

`func (o *QueryAuthResponseAuthorizationInfo) GetAuthorizerAppid() string`

GetAuthorizerAppid returns the AuthorizerAppid field if non-nil, zero value otherwise.

### GetAuthorizerAppidOk

`func (o *QueryAuthResponseAuthorizationInfo) GetAuthorizerAppidOk() (*string, bool)`

GetAuthorizerAppidOk returns a tuple with the AuthorizerAppid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizerAppid

`func (o *QueryAuthResponseAuthorizationInfo) SetAuthorizerAppid(v string)`

SetAuthorizerAppid sets AuthorizerAppid field to given value.

### HasAuthorizerAppid

`func (o *QueryAuthResponseAuthorizationInfo) HasAuthorizerAppid() bool`

HasAuthorizerAppid returns a boolean if a field has been set.

### GetAuthorizerAccessToken

`func (o *QueryAuthResponseAuthorizationInfo) GetAuthorizerAccessToken() string`

GetAuthorizerAccessToken returns the AuthorizerAccessToken field if non-nil, zero value otherwise.

### GetAuthorizerAccessTokenOk

`func (o *QueryAuthResponseAuthorizationInfo) GetAuthorizerAccessTokenOk() (*string, bool)`

GetAuthorizerAccessTokenOk returns a tuple with the AuthorizerAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizerAccessToken

`func (o *QueryAuthResponseAuthorizationInfo) SetAuthorizerAccessToken(v string)`

SetAuthorizerAccessToken sets AuthorizerAccessToken field to given value.

### HasAuthorizerAccessToken

`func (o *QueryAuthResponseAuthorizationInfo) HasAuthorizerAccessToken() bool`

HasAuthorizerAccessToken returns a boolean if a field has been set.

### GetExpiresIn

`func (o *QueryAuthResponseAuthorizationInfo) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *QueryAuthResponseAuthorizationInfo) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *QueryAuthResponseAuthorizationInfo) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *QueryAuthResponseAuthorizationInfo) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetAuthorizerRefreshToken

`func (o *QueryAuthResponseAuthorizationInfo) GetAuthorizerRefreshToken() string`

GetAuthorizerRefreshToken returns the AuthorizerRefreshToken field if non-nil, zero value otherwise.

### GetAuthorizerRefreshTokenOk

`func (o *QueryAuthResponseAuthorizationInfo) GetAuthorizerRefreshTokenOk() (*string, bool)`

GetAuthorizerRefreshTokenOk returns a tuple with the AuthorizerRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizerRefreshToken

`func (o *QueryAuthResponseAuthorizationInfo) SetAuthorizerRefreshToken(v string)`

SetAuthorizerRefreshToken sets AuthorizerRefreshToken field to given value.

### HasAuthorizerRefreshToken

`func (o *QueryAuthResponseAuthorizationInfo) HasAuthorizerRefreshToken() bool`

HasAuthorizerRefreshToken returns a boolean if a field has been set.

### GetFuncInfo

`func (o *QueryAuthResponseAuthorizationInfo) GetFuncInfo() []GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner`

GetFuncInfo returns the FuncInfo field if non-nil, zero value otherwise.

### GetFuncInfoOk

`func (o *QueryAuthResponseAuthorizationInfo) GetFuncInfoOk() (*[]GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner, bool)`

GetFuncInfoOk returns a tuple with the FuncInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuncInfo

`func (o *QueryAuthResponseAuthorizationInfo) SetFuncInfo(v []GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner)`

SetFuncInfo sets FuncInfo field to given value.

### HasFuncInfo

`func (o *QueryAuthResponseAuthorizationInfo) HasFuncInfo() bool`

HasFuncInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


