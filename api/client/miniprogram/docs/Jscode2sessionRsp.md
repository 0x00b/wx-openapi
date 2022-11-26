# Jscode2sessionRsp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errcode** | Pointer to **int32** | 返回码 | [optional] 
**Errmsg** | Pointer to **string** | 对返回码的文本描述内容 | [optional] 
**Openid** | Pointer to **string** | 用户唯一标识 | [optional] 
**SessionKey** | Pointer to **string** | 会话密钥 | [optional] 
**Unionid** | Pointer to **string** | 用户在开放平台的唯一标识符，若当前小程序已绑定到微信开放平台帐号下会返回，详见UnionID机制说明。 | [optional] 

## Methods

### NewJscode2sessionRsp

`func NewJscode2sessionRsp() *Jscode2sessionRsp`

NewJscode2sessionRsp instantiates a new Jscode2sessionRsp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJscode2sessionRspWithDefaults

`func NewJscode2sessionRspWithDefaults() *Jscode2sessionRsp`

NewJscode2sessionRspWithDefaults instantiates a new Jscode2sessionRsp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrcode

`func (o *Jscode2sessionRsp) GetErrcode() int32`

GetErrcode returns the Errcode field if non-nil, zero value otherwise.

### GetErrcodeOk

`func (o *Jscode2sessionRsp) GetErrcodeOk() (*int32, bool)`

GetErrcodeOk returns a tuple with the Errcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrcode

`func (o *Jscode2sessionRsp) SetErrcode(v int32)`

SetErrcode sets Errcode field to given value.

### HasErrcode

`func (o *Jscode2sessionRsp) HasErrcode() bool`

HasErrcode returns a boolean if a field has been set.

### GetErrmsg

`func (o *Jscode2sessionRsp) GetErrmsg() string`

GetErrmsg returns the Errmsg field if non-nil, zero value otherwise.

### GetErrmsgOk

`func (o *Jscode2sessionRsp) GetErrmsgOk() (*string, bool)`

GetErrmsgOk returns a tuple with the Errmsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrmsg

`func (o *Jscode2sessionRsp) SetErrmsg(v string)`

SetErrmsg sets Errmsg field to given value.

### HasErrmsg

`func (o *Jscode2sessionRsp) HasErrmsg() bool`

HasErrmsg returns a boolean if a field has been set.

### GetOpenid

`func (o *Jscode2sessionRsp) GetOpenid() string`

GetOpenid returns the Openid field if non-nil, zero value otherwise.

### GetOpenidOk

`func (o *Jscode2sessionRsp) GetOpenidOk() (*string, bool)`

GetOpenidOk returns a tuple with the Openid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenid

`func (o *Jscode2sessionRsp) SetOpenid(v string)`

SetOpenid sets Openid field to given value.

### HasOpenid

`func (o *Jscode2sessionRsp) HasOpenid() bool`

HasOpenid returns a boolean if a field has been set.

### GetSessionKey

`func (o *Jscode2sessionRsp) GetSessionKey() string`

GetSessionKey returns the SessionKey field if non-nil, zero value otherwise.

### GetSessionKeyOk

`func (o *Jscode2sessionRsp) GetSessionKeyOk() (*string, bool)`

GetSessionKeyOk returns a tuple with the SessionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionKey

`func (o *Jscode2sessionRsp) SetSessionKey(v string)`

SetSessionKey sets SessionKey field to given value.

### HasSessionKey

`func (o *Jscode2sessionRsp) HasSessionKey() bool`

HasSessionKey returns a boolean if a field has been set.

### GetUnionid

`func (o *Jscode2sessionRsp) GetUnionid() string`

GetUnionid returns the Unionid field if non-nil, zero value otherwise.

### GetUnionidOk

`func (o *Jscode2sessionRsp) GetUnionidOk() (*string, bool)`

GetUnionidOk returns a tuple with the Unionid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnionid

`func (o *Jscode2sessionRsp) SetUnionid(v string)`

SetUnionid sets Unionid field to given value.

### HasUnionid

`func (o *Jscode2sessionRsp) HasUnionid() bool`

HasUnionid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


