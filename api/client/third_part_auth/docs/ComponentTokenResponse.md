# ComponentTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errcode** | Pointer to **int32** | 返回码 | [optional] 
**Errmsg** | Pointer to **string** | 对返回码的文本描述内容 | [optional] 
**ExpiresIn** | Pointer to **int32** | 有效期，单位：秒 | [optional] 
**ComponentAccessToken** | Pointer to **string** | 第三方平台 access_token | [optional] 

## Methods

### NewComponentTokenResponse

`func NewComponentTokenResponse() *ComponentTokenResponse`

NewComponentTokenResponse instantiates a new ComponentTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentTokenResponseWithDefaults

`func NewComponentTokenResponseWithDefaults() *ComponentTokenResponse`

NewComponentTokenResponseWithDefaults instantiates a new ComponentTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrcode

`func (o *ComponentTokenResponse) GetErrcode() int32`

GetErrcode returns the Errcode field if non-nil, zero value otherwise.

### GetErrcodeOk

`func (o *ComponentTokenResponse) GetErrcodeOk() (*int32, bool)`

GetErrcodeOk returns a tuple with the Errcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrcode

`func (o *ComponentTokenResponse) SetErrcode(v int32)`

SetErrcode sets Errcode field to given value.

### HasErrcode

`func (o *ComponentTokenResponse) HasErrcode() bool`

HasErrcode returns a boolean if a field has been set.

### GetErrmsg

`func (o *ComponentTokenResponse) GetErrmsg() string`

GetErrmsg returns the Errmsg field if non-nil, zero value otherwise.

### GetErrmsgOk

`func (o *ComponentTokenResponse) GetErrmsgOk() (*string, bool)`

GetErrmsgOk returns a tuple with the Errmsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrmsg

`func (o *ComponentTokenResponse) SetErrmsg(v string)`

SetErrmsg sets Errmsg field to given value.

### HasErrmsg

`func (o *ComponentTokenResponse) HasErrmsg() bool`

HasErrmsg returns a boolean if a field has been set.

### GetExpiresIn

`func (o *ComponentTokenResponse) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *ComponentTokenResponse) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *ComponentTokenResponse) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *ComponentTokenResponse) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetComponentAccessToken

`func (o *ComponentTokenResponse) GetComponentAccessToken() string`

GetComponentAccessToken returns the ComponentAccessToken field if non-nil, zero value otherwise.

### GetComponentAccessTokenOk

`func (o *ComponentTokenResponse) GetComponentAccessTokenOk() (*string, bool)`

GetComponentAccessTokenOk returns a tuple with the ComponentAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentAccessToken

`func (o *ComponentTokenResponse) SetComponentAccessToken(v string)`

SetComponentAccessToken sets ComponentAccessToken field to given value.

### HasComponentAccessToken

`func (o *ComponentTokenResponse) HasComponentAccessToken() bool`

HasComponentAccessToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


