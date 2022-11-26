# StartPushTicketResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errcode** | Pointer to **int32** | 返回码 | [optional] 
**Errmsg** | Pointer to **string** | 对返回码的文本描述内容 | [optional] 

## Methods

### NewStartPushTicketResponse

`func NewStartPushTicketResponse() *StartPushTicketResponse`

NewStartPushTicketResponse instantiates a new StartPushTicketResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartPushTicketResponseWithDefaults

`func NewStartPushTicketResponseWithDefaults() *StartPushTicketResponse`

NewStartPushTicketResponseWithDefaults instantiates a new StartPushTicketResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrcode

`func (o *StartPushTicketResponse) GetErrcode() int32`

GetErrcode returns the Errcode field if non-nil, zero value otherwise.

### GetErrcodeOk

`func (o *StartPushTicketResponse) GetErrcodeOk() (*int32, bool)`

GetErrcodeOk returns a tuple with the Errcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrcode

`func (o *StartPushTicketResponse) SetErrcode(v int32)`

SetErrcode sets Errcode field to given value.

### HasErrcode

`func (o *StartPushTicketResponse) HasErrcode() bool`

HasErrcode returns a boolean if a field has been set.

### GetErrmsg

`func (o *StartPushTicketResponse) GetErrmsg() string`

GetErrmsg returns the Errmsg field if non-nil, zero value otherwise.

### GetErrmsgOk

`func (o *StartPushTicketResponse) GetErrmsgOk() (*string, bool)`

GetErrmsgOk returns a tuple with the Errmsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrmsg

`func (o *StartPushTicketResponse) SetErrmsg(v string)`

SetErrmsg sets Errmsg field to given value.

### HasErrmsg

`func (o *StartPushTicketResponse) HasErrmsg() bool`

HasErrmsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


