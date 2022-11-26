# ComponentTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentAppid** | **string** | 平台型第三方平台的appid | 
**ComponentAppsecret** | **string** | 第三方平台 appsecret | 
**ComponentVerifyTicket** | **string** | 微信后台推送的 ticket | 

## Methods

### NewComponentTokenRequest

`func NewComponentTokenRequest(componentAppid string, componentAppsecret string, componentVerifyTicket string, ) *ComponentTokenRequest`

NewComponentTokenRequest instantiates a new ComponentTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentTokenRequestWithDefaults

`func NewComponentTokenRequestWithDefaults() *ComponentTokenRequest`

NewComponentTokenRequestWithDefaults instantiates a new ComponentTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentAppid

`func (o *ComponentTokenRequest) GetComponentAppid() string`

GetComponentAppid returns the ComponentAppid field if non-nil, zero value otherwise.

### GetComponentAppidOk

`func (o *ComponentTokenRequest) GetComponentAppidOk() (*string, bool)`

GetComponentAppidOk returns a tuple with the ComponentAppid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentAppid

`func (o *ComponentTokenRequest) SetComponentAppid(v string)`

SetComponentAppid sets ComponentAppid field to given value.


### GetComponentAppsecret

`func (o *ComponentTokenRequest) GetComponentAppsecret() string`

GetComponentAppsecret returns the ComponentAppsecret field if non-nil, zero value otherwise.

### GetComponentAppsecretOk

`func (o *ComponentTokenRequest) GetComponentAppsecretOk() (*string, bool)`

GetComponentAppsecretOk returns a tuple with the ComponentAppsecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentAppsecret

`func (o *ComponentTokenRequest) SetComponentAppsecret(v string)`

SetComponentAppsecret sets ComponentAppsecret field to given value.


### GetComponentVerifyTicket

`func (o *ComponentTokenRequest) GetComponentVerifyTicket() string`

GetComponentVerifyTicket returns the ComponentVerifyTicket field if non-nil, zero value otherwise.

### GetComponentVerifyTicketOk

`func (o *ComponentTokenRequest) GetComponentVerifyTicketOk() (*string, bool)`

GetComponentVerifyTicketOk returns a tuple with the ComponentVerifyTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentVerifyTicket

`func (o *ComponentTokenRequest) SetComponentVerifyTicket(v string)`

SetComponentVerifyTicket sets ComponentVerifyTicket field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


