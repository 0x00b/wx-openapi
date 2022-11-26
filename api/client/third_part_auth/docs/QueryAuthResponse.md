# QueryAuthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errcode** | Pointer to **int32** | 返回码 | [optional] 
**Errmsg** | Pointer to **string** | 对返回码的文本描述内容 | [optional] 
**AuthorizationInfo** | Pointer to [**QueryAuthResponseAuthorizationInfo**](QueryAuthResponseAuthorizationInfo.md) |  | [optional] 

## Methods

### NewQueryAuthResponse

`func NewQueryAuthResponse() *QueryAuthResponse`

NewQueryAuthResponse instantiates a new QueryAuthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryAuthResponseWithDefaults

`func NewQueryAuthResponseWithDefaults() *QueryAuthResponse`

NewQueryAuthResponseWithDefaults instantiates a new QueryAuthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrcode

`func (o *QueryAuthResponse) GetErrcode() int32`

GetErrcode returns the Errcode field if non-nil, zero value otherwise.

### GetErrcodeOk

`func (o *QueryAuthResponse) GetErrcodeOk() (*int32, bool)`

GetErrcodeOk returns a tuple with the Errcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrcode

`func (o *QueryAuthResponse) SetErrcode(v int32)`

SetErrcode sets Errcode field to given value.

### HasErrcode

`func (o *QueryAuthResponse) HasErrcode() bool`

HasErrcode returns a boolean if a field has been set.

### GetErrmsg

`func (o *QueryAuthResponse) GetErrmsg() string`

GetErrmsg returns the Errmsg field if non-nil, zero value otherwise.

### GetErrmsgOk

`func (o *QueryAuthResponse) GetErrmsgOk() (*string, bool)`

GetErrmsgOk returns a tuple with the Errmsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrmsg

`func (o *QueryAuthResponse) SetErrmsg(v string)`

SetErrmsg sets Errmsg field to given value.

### HasErrmsg

`func (o *QueryAuthResponse) HasErrmsg() bool`

HasErrmsg returns a boolean if a field has been set.

### GetAuthorizationInfo

`func (o *QueryAuthResponse) GetAuthorizationInfo() QueryAuthResponseAuthorizationInfo`

GetAuthorizationInfo returns the AuthorizationInfo field if non-nil, zero value otherwise.

### GetAuthorizationInfoOk

`func (o *QueryAuthResponse) GetAuthorizationInfoOk() (*QueryAuthResponseAuthorizationInfo, bool)`

GetAuthorizationInfoOk returns a tuple with the AuthorizationInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationInfo

`func (o *QueryAuthResponse) SetAuthorizationInfo(v QueryAuthResponseAuthorizationInfo)`

SetAuthorizationInfo sets AuthorizationInfo field to given value.

### HasAuthorizationInfo

`func (o *QueryAuthResponse) HasAuthorizationInfo() bool`

HasAuthorizationInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


