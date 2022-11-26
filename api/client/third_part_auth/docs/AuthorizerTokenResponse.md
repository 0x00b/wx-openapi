# AuthorizerTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errcode** | Pointer to **int32** | 返回码 | [optional] 
**Errmsg** | Pointer to **string** | 对返回码的文本描述内容 | [optional] 
**ExpiresIn** | Pointer to **int32** | 有效期，单位：秒 | [optional] 
**AuthorizerAccessToken** | Pointer to **string** | 授权方令牌 | [optional] 
**AuthorizerRefreshToken** | Pointer to **string** | 刷新令牌 | [optional] 

## Methods

### NewAuthorizerTokenResponse

`func NewAuthorizerTokenResponse() *AuthorizerTokenResponse`

NewAuthorizerTokenResponse instantiates a new AuthorizerTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizerTokenResponseWithDefaults

`func NewAuthorizerTokenResponseWithDefaults() *AuthorizerTokenResponse`

NewAuthorizerTokenResponseWithDefaults instantiates a new AuthorizerTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrcode

`func (o *AuthorizerTokenResponse) GetErrcode() int32`

GetErrcode returns the Errcode field if non-nil, zero value otherwise.

### GetErrcodeOk

`func (o *AuthorizerTokenResponse) GetErrcodeOk() (*int32, bool)`

GetErrcodeOk returns a tuple with the Errcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrcode

`func (o *AuthorizerTokenResponse) SetErrcode(v int32)`

SetErrcode sets Errcode field to given value.

### HasErrcode

`func (o *AuthorizerTokenResponse) HasErrcode() bool`

HasErrcode returns a boolean if a field has been set.

### GetErrmsg

`func (o *AuthorizerTokenResponse) GetErrmsg() string`

GetErrmsg returns the Errmsg field if non-nil, zero value otherwise.

### GetErrmsgOk

`func (o *AuthorizerTokenResponse) GetErrmsgOk() (*string, bool)`

GetErrmsgOk returns a tuple with the Errmsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrmsg

`func (o *AuthorizerTokenResponse) SetErrmsg(v string)`

SetErrmsg sets Errmsg field to given value.

### HasErrmsg

`func (o *AuthorizerTokenResponse) HasErrmsg() bool`

HasErrmsg returns a boolean if a field has been set.

### GetExpiresIn

`func (o *AuthorizerTokenResponse) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *AuthorizerTokenResponse) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *AuthorizerTokenResponse) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *AuthorizerTokenResponse) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetAuthorizerAccessToken

`func (o *AuthorizerTokenResponse) GetAuthorizerAccessToken() string`

GetAuthorizerAccessToken returns the AuthorizerAccessToken field if non-nil, zero value otherwise.

### GetAuthorizerAccessTokenOk

`func (o *AuthorizerTokenResponse) GetAuthorizerAccessTokenOk() (*string, bool)`

GetAuthorizerAccessTokenOk returns a tuple with the AuthorizerAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizerAccessToken

`func (o *AuthorizerTokenResponse) SetAuthorizerAccessToken(v string)`

SetAuthorizerAccessToken sets AuthorizerAccessToken field to given value.

### HasAuthorizerAccessToken

`func (o *AuthorizerTokenResponse) HasAuthorizerAccessToken() bool`

HasAuthorizerAccessToken returns a boolean if a field has been set.

### GetAuthorizerRefreshToken

`func (o *AuthorizerTokenResponse) GetAuthorizerRefreshToken() string`

GetAuthorizerRefreshToken returns the AuthorizerRefreshToken field if non-nil, zero value otherwise.

### GetAuthorizerRefreshTokenOk

`func (o *AuthorizerTokenResponse) GetAuthorizerRefreshTokenOk() (*string, bool)`

GetAuthorizerRefreshTokenOk returns a tuple with the AuthorizerRefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizerRefreshToken

`func (o *AuthorizerTokenResponse) SetAuthorizerRefreshToken(v string)`

SetAuthorizerRefreshToken sets AuthorizerRefreshToken field to given value.

### HasAuthorizerRefreshToken

`func (o *AuthorizerTokenResponse) HasAuthorizerRefreshToken() bool`

HasAuthorizerRefreshToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


