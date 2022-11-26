# SubscribeMessageSendRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Touser** | Pointer to **string** | 是 接收者（用户）的 openid | [optional] 
**TemplateId** | Pointer to **string** | 是 所需下发的订阅模板id | [optional] 
**Page** | Pointer to **string** | 否 点击模板卡片后的跳转页面，仅限本小程序内的页面。支持带参数,（示例index?foo&#x3D;bar）。该字段不填则模板无跳转。 | [optional] 
**Data** | Pointer to **map[string]interface{}** | 是 模板内容，格式形如 { \&quot;key1\&quot;: { \&quot;value\&quot;: any }, \&quot;key2\&quot;: { \&quot;value\&quot;: any } } | [optional] 
**MiniprogramState** | Pointer to **string** | 否 跳转小程序类型：developer为开发版；trial为体验版；formal为正式版；默认为正式版 | [optional] 
**Lang** | Pointer to **string** | 否 进入小程序查看”的语言类型，支持zh_CN(简体中文)、en_US(英文)、zh_HK(繁体中文)、zh_TW(繁体中文)，默认为zh_CN | [optional] 

## Methods

### NewSubscribeMessageSendRequest

`func NewSubscribeMessageSendRequest() *SubscribeMessageSendRequest`

NewSubscribeMessageSendRequest instantiates a new SubscribeMessageSendRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscribeMessageSendRequestWithDefaults

`func NewSubscribeMessageSendRequestWithDefaults() *SubscribeMessageSendRequest`

NewSubscribeMessageSendRequestWithDefaults instantiates a new SubscribeMessageSendRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTouser

`func (o *SubscribeMessageSendRequest) GetTouser() string`

GetTouser returns the Touser field if non-nil, zero value otherwise.

### GetTouserOk

`func (o *SubscribeMessageSendRequest) GetTouserOk() (*string, bool)`

GetTouserOk returns a tuple with the Touser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTouser

`func (o *SubscribeMessageSendRequest) SetTouser(v string)`

SetTouser sets Touser field to given value.

### HasTouser

`func (o *SubscribeMessageSendRequest) HasTouser() bool`

HasTouser returns a boolean if a field has been set.

### GetTemplateId

`func (o *SubscribeMessageSendRequest) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *SubscribeMessageSendRequest) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *SubscribeMessageSendRequest) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *SubscribeMessageSendRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetPage

`func (o *SubscribeMessageSendRequest) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *SubscribeMessageSendRequest) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *SubscribeMessageSendRequest) SetPage(v string)`

SetPage sets Page field to given value.

### HasPage

`func (o *SubscribeMessageSendRequest) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetData

`func (o *SubscribeMessageSendRequest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SubscribeMessageSendRequest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SubscribeMessageSendRequest) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *SubscribeMessageSendRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMiniprogramState

`func (o *SubscribeMessageSendRequest) GetMiniprogramState() string`

GetMiniprogramState returns the MiniprogramState field if non-nil, zero value otherwise.

### GetMiniprogramStateOk

`func (o *SubscribeMessageSendRequest) GetMiniprogramStateOk() (*string, bool)`

GetMiniprogramStateOk returns a tuple with the MiniprogramState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiniprogramState

`func (o *SubscribeMessageSendRequest) SetMiniprogramState(v string)`

SetMiniprogramState sets MiniprogramState field to given value.

### HasMiniprogramState

`func (o *SubscribeMessageSendRequest) HasMiniprogramState() bool`

HasMiniprogramState returns a boolean if a field has been set.

### GetLang

`func (o *SubscribeMessageSendRequest) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *SubscribeMessageSendRequest) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *SubscribeMessageSendRequest) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *SubscribeMessageSendRequest) HasLang() bool`

HasLang returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


