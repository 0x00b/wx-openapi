# BatchgetMaterialResponseItemInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MediaId** | Pointer to **string** |  | [optional] 
**Content** | Pointer to [**BatchgetMaterialResponseItemInnerContent**](BatchgetMaterialResponseItemInnerContent.md) |  | [optional] 
**Url** | Pointer to **string** | 图文页的URL，或者，当获取的列表是图片素材列表时，该字段是图片的URL | [optional] 
**UpdateTime** | Pointer to **int64** | 这篇图文消息素材的最后更新时间 | [optional] 
**Name** | Pointer to **string** | 文件名称 | [optional] 

## Methods

### NewBatchgetMaterialResponseItemInner

`func NewBatchgetMaterialResponseItemInner() *BatchgetMaterialResponseItemInner`

NewBatchgetMaterialResponseItemInner instantiates a new BatchgetMaterialResponseItemInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchgetMaterialResponseItemInnerWithDefaults

`func NewBatchgetMaterialResponseItemInnerWithDefaults() *BatchgetMaterialResponseItemInner`

NewBatchgetMaterialResponseItemInnerWithDefaults instantiates a new BatchgetMaterialResponseItemInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMediaId

`func (o *BatchgetMaterialResponseItemInner) GetMediaId() string`

GetMediaId returns the MediaId field if non-nil, zero value otherwise.

### GetMediaIdOk

`func (o *BatchgetMaterialResponseItemInner) GetMediaIdOk() (*string, bool)`

GetMediaIdOk returns a tuple with the MediaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaId

`func (o *BatchgetMaterialResponseItemInner) SetMediaId(v string)`

SetMediaId sets MediaId field to given value.

### HasMediaId

`func (o *BatchgetMaterialResponseItemInner) HasMediaId() bool`

HasMediaId returns a boolean if a field has been set.

### GetContent

`func (o *BatchgetMaterialResponseItemInner) GetContent() BatchgetMaterialResponseItemInnerContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *BatchgetMaterialResponseItemInner) GetContentOk() (*BatchgetMaterialResponseItemInnerContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *BatchgetMaterialResponseItemInner) SetContent(v BatchgetMaterialResponseItemInnerContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *BatchgetMaterialResponseItemInner) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetUrl

`func (o *BatchgetMaterialResponseItemInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BatchgetMaterialResponseItemInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BatchgetMaterialResponseItemInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *BatchgetMaterialResponseItemInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUpdateTime

`func (o *BatchgetMaterialResponseItemInner) GetUpdateTime() int64`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *BatchgetMaterialResponseItemInner) GetUpdateTimeOk() (*int64, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *BatchgetMaterialResponseItemInner) SetUpdateTime(v int64)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *BatchgetMaterialResponseItemInner) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetName

`func (o *BatchgetMaterialResponseItemInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BatchgetMaterialResponseItemInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BatchgetMaterialResponseItemInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BatchgetMaterialResponseItemInner) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


