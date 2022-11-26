# AuthorizerTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentAppid** | **string** | 平台型第三方平台的appid | 
**AuthorizerAppid** | **string** | 授权方 appid | 
**AuthorizerRefreshToken** | **string** | 刷新令牌，获取授权信息时得到 | 

## Methods

### NewAuthorizerTokenRequest

`func NewAuthorizerTokenRequest(componentAppid string, authorizerAppid string, authorizerRefreshToken string, ) *AuthorizerTokenRequest`

NewAuthorizerTokenRequest instantiates a new AuthorizerTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizerTokenRequestWithDefaults

`func NewAuthorizerTokenRequestWithDefaults() *AuthorizerTokenRequest`

NewAuthorizerTokenRequestWithDefaults instantiates a new AuthorizerTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentAppid

`func (o *AuthorizerTokenRequest) GetComponentAppid() string`

GetComponentAppid returns the ComponentAppid field if non-nil, zero value otherwise.

### GetComponentAppidOk

`func (o *AuthorizerTokenRequest) GetComponentAppidOk() (*string, bool)`

GetComponentAppidOk returns a tuple with the ComponentAppid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentAppid

`func (o *AuthorizerTokenRequest) SetComponentAppid(v string)`

SetComponentAppid sets ComponentAppid field to given value.


### GetAuthorizerAppid

`func (o *AuthorizerTokenRequest) GetAuthorizerAppid() string`

GetAuthorizerAppid returns the AuthorizerAppid field if non-nil, zero value otherwise.

### GetAuthorizerAppidOk

`func (o *AuthorizerTokenRequest) GetAuthorizerAppidOk() (*string, bool)`

GetAuthorizerAppidOk returns a tuple with the AuthorizerAppid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizerAppid

`func (o *AuthorizerTokenRequest) SetAuthorizerAppid(v string)`

SetAuthorizerAppid sets AuthorizerAppid field to given value.


### GetAuthorizerRefreshToken

`func (o *AuthorizerTokenRequest) GetAuthorizerRefreshToken() string`

GetAuthorizerRefreshToken returns the AuthorizerRefreshToken field if non-nil, zero value otherwise.

### GetAuthorizerRefreshTokenOk

`func (o *AuthorizerTokenRequest) GetAuthorizerRefreshTokenOk() (*string, bool)`

GetAuthorizerRefreshTokenOk returns a tuple with the AuthorizerRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizerRefreshToken

`func (o *AuthorizerTokenRequest) SetAuthorizerRefreshToken(v string)`

SetAuthorizerRefreshToken sets AuthorizerRefreshToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


