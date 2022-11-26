# GetAuthorizerInfoResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errcode** | Pointer to **int32** | 返回码 | [optional] 
**Errmsg** | Pointer to **string** | 对返回码的文本描述内容 | [optional] 
**AuthorizerInfo** | Pointer to [**GetAuthorizerInfoResponseAuthorizerInfo**](GetAuthorizerInfoResponseAuthorizerInfo.md) |  | [optional] 
**AuthorizationInfo** | Pointer to [**GetAuthorizerInfoResponseAuthorizationInfo**](GetAuthorizerInfoResponseAuthorizationInfo.md) |  | [optional] 

## Methods

### NewGetAuthorizerInfoResponse

`func NewGetAuthorizerInfoResponse() *GetAuthorizerInfoResponse`

NewGetAuthorizerInfoResponse instantiates a new GetAuthorizerInfoResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAuthorizerInfoResponseWithDefaults

`func NewGetAuthorizerInfoResponseWithDefaults() *GetAuthorizerInfoResponse`

NewGetAuthorizerInfoResponseWithDefaults instantiates a new GetAuthorizerInfoResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrcode

`func (o *GetAuthorizerInfoResponse) GetErrcode() int32`

GetErrcode returns the Errcode field if non-nil, zero value otherwise.

### GetErrcodeOk

`func (o *GetAuthorizerInfoResponse) GetErrcodeOk() (*int32, bool)`

GetErrcodeOk returns a tuple with the Errcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrcode

`func (o *GetAuthorizerInfoResponse) SetErrcode(v int32)`

SetErrcode sets Errcode field to given value.

### HasErrcode

`func (o *GetAuthorizerInfoResponse) HasErrcode() bool`

HasErrcode returns a boolean if a field has been set.

### GetErrmsg

`func (o *GetAuthorizerInfoResponse) GetErrmsg() string`

GetErrmsg returns the Errmsg field if non-nil, zero value otherwise.

### GetErrmsgOk

`func (o *GetAuthorizerInfoResponse) GetErrmsgOk() (*string, bool)`

GetErrmsgOk returns a tuple with the Errmsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrmsg

`func (o *GetAuthorizerInfoResponse) SetErrmsg(v string)`

SetErrmsg sets Errmsg field to given value.

### HasErrmsg

`func (o *GetAuthorizerInfoResponse) HasErrmsg() bool`

HasErrmsg returns a boolean if a field has been set.

### GetAuthorizerInfo

`func (o *GetAuthorizerInfoResponse) GetAuthorizerInfo() GetAuthorizerInfoResponseAuthorizerInfo`

GetAuthorizerInfo returns the AuthorizerInfo field if non-nil, zero value otherwise.

### GetAuthorizerInfoOk

`func (o *GetAuthorizerInfoResponse) GetAuthorizerInfoOk() (*GetAuthorizerInfoResponseAuthorizerInfo, bool)`

GetAuthorizerInfoOk returns a tuple with the AuthorizerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizerInfo

`func (o *GetAuthorizerInfoResponse) SetAuthorizerInfo(v GetAuthorizerInfoResponseAuthorizerInfo)`

SetAuthorizerInfo sets AuthorizerInfo field to given value.

### HasAuthorizerInfo

`func (o *GetAuthorizerInfoResponse) HasAuthorizerInfo() bool`

HasAuthorizerInfo returns a boolean if a field has been set.

### GetAuthorizationInfo

`func (o *GetAuthorizerInfoResponse) GetAuthorizationInfo() GetAuthorizerInfoResponseAuthorizationInfo`

GetAuthorizationInfo returns the AuthorizationInfo field if non-nil, zero value otherwise.

### GetAuthorizationInfoOk

`func (o *GetAuthorizerInfoResponse) GetAuthorizationInfoOk() (*GetAuthorizerInfoResponseAuthorizationInfo, bool)`

GetAuthorizationInfoOk returns a tuple with the AuthorizationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationInfo

`func (o *GetAuthorizerInfoResponse) SetAuthorizationInfo(v GetAuthorizerInfoResponseAuthorizationInfo)`

SetAuthorizationInfo sets AuthorizationInfo field to given value.

### HasAuthorizationInfo

`func (o *GetAuthorizerInfoResponse) HasAuthorizationInfo() bool`

HasAuthorizationInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


