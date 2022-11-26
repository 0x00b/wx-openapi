# CreatePreauthcodeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errcode** | Pointer to **int32** | 返回码 | [optional] 
**Errmsg** | Pointer to **string** | 对返回码的文本描述内容 | [optional] 
**ExpiresIn** | Pointer to **int32** | 有效期，单位：秒 | [optional] 
**PreAuthCode** | Pointer to **string** | 预授权码 | [optional] 

## Methods

### NewCreatePreauthcodeResponse

`func NewCreatePreauthcodeResponse() *CreatePreauthcodeResponse`

NewCreatePreauthcodeResponse instantiates a new CreatePreauthcodeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePreauthcodeResponseWithDefaults

`func NewCreatePreauthcodeResponseWithDefaults() *CreatePreauthcodeResponse`

NewCreatePreauthcodeResponseWithDefaults instantiates a new CreatePreauthcodeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrcode

`func (o *CreatePreauthcodeResponse) GetErrcode() int32`

GetErrcode returns the Errcode field if non-nil, zero value otherwise.

### GetErrcodeOk

`func (o *CreatePreauthcodeResponse) GetErrcodeOk() (*int32, bool)`

GetErrcodeOk returns a tuple with the Errcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrcode

`func (o *CreatePreauthcodeResponse) SetErrcode(v int32)`

SetErrcode sets Errcode field to given value.

### HasErrcode

`func (o *CreatePreauthcodeResponse) HasErrcode() bool`

HasErrcode returns a boolean if a field has been set.

### GetErrmsg

`func (o *CreatePreauthcodeResponse) GetErrmsg() string`

GetErrmsg returns the Errmsg field if non-nil, zero value otherwise.

### GetErrmsgOk

`func (o *CreatePreauthcodeResponse) GetErrmsgOk() (*string, bool)`

GetErrmsgOk returns a tuple with the Errmsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrmsg

`func (o *CreatePreauthcodeResponse) SetErrmsg(v string)`

SetErrmsg sets Errmsg field to given value.

### HasErrmsg

`func (o *CreatePreauthcodeResponse) HasErrmsg() bool`

HasErrmsg returns a boolean if a field has been set.

### GetExpiresIn

`func (o *CreatePreauthcodeResponse) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *CreatePreauthcodeResponse) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *CreatePreauthcodeResponse) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *CreatePreauthcodeResponse) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetPreAuthCode

`func (o *CreatePreauthcodeResponse) GetPreAuthCode() string`

GetPreAuthCode returns the PreAuthCode field if non-nil, zero value otherwise.

### GetPreAuthCodeOk

`func (o *CreatePreauthcodeResponse) GetPreAuthCodeOk() (*string, bool)`

GetPreAuthCodeOk returns a tuple with the PreAuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthCode

`func (o *CreatePreauthcodeResponse) SetPreAuthCode(v string)`

SetPreAuthCode sets PreAuthCode field to given value.

### HasPreAuthCode

`func (o *CreatePreauthcodeResponse) HasPreAuthCode() bool`

HasPreAuthCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


