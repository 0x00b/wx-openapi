# BatchgetMaterialResponseItemInnerContentNewsItemInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | 图文消息的标题 | [optional] 
**ThumbMediaId** | Pointer to **string** | 图文消息的封面图片素材id（必须是永久mediaID） | [optional] 
**ShowCoverPic** | Pointer to **int32** | 是否显示封面，0为false，即不显示，1为true，即显示 | [optional] 
**Author** | Pointer to **string** | 作者 | [optional] 
**Digest** | Pointer to **string** | 图文消息的摘要，仅有单图文消息才有摘要，多图文此处为空 | [optional] 
**Content** | Pointer to **string** | 图文消息的具体内容，支持HTML标签，必须少于2万字符，小于1M，且此处会去除JS | [optional] 
**ContentSourceUrl** | Pointer to **string** | 图文消息的原文地址，即点击“阅读原文”后的URL | [optional] 
**Url** | Pointer to **string** | 图文页的URL，或者，当获取的列表是图片素材列表时，该字段是图片的URL | [optional] 
**ThumbUrl** | Pointer to **string** | 封面URL | [optional] 

## Methods

### NewBatchgetMaterialResponseItemInnerContentNewsItemInner

`func NewBatchgetMaterialResponseItemInnerContentNewsItemInner() *BatchgetMaterialResponseItemInnerContentNewsItemInner`

NewBatchgetMaterialResponseItemInnerContentNewsItemInner instantiates a new BatchgetMaterialResponseItemInnerContentNewsItemInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchgetMaterialResponseItemInnerContentNewsItemInnerWithDefaults

`func NewBatchgetMaterialResponseItemInnerContentNewsItemInnerWithDefaults() *BatchgetMaterialResponseItemInnerContentNewsItemInner`

NewBatchgetMaterialResponseItemInnerContentNewsItemInnerWithDefaults instantiates a new BatchgetMaterialResponseItemInnerContentNewsItemInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetThumbMediaId

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) GetThumbMediaId() string`

GetThumbMediaId returns the ThumbMediaId field if non-nil, zero value otherwise.

### GetThumbMediaIdOk

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) GetThumbMediaIdOk() (*string, bool)`

GetThumbMediaIdOk returns a tuple with the ThumbMediaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbMediaId

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) SetThumbMediaId(v string)`

SetThumbMediaId sets ThumbMediaId field to given value.

### HasThumbMediaId

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) HasThumbMediaId() bool`

HasThumbMediaId returns a boolean if a field has been set.

### GetShowCoverPic

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) GetShowCoverPic() int32`

GetShowCoverPic returns the ShowCoverPic field if non-nil, zero value otherwise.

### GetShowCoverPicOk

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) GetShowCoverPicOk() (*int32, bool)`

GetShowCoverPicOk returns a tuple with the ShowCoverPic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowCoverPic

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) SetShowCoverPic(v int32)`

SetShowCoverPic sets ShowCoverPic field to given value.

### HasShowCoverPic

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) HasShowCoverPic() bool`

HasShowCoverPic returns a boolean if a field has been set.

### GetAuthor

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetDigest

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetContent

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetContentSourceUrl

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) GetContentSourceUrl() string`

GetContentSourceUrl returns the ContentSourceUrl field if non-nil, zero value otherwise.

### GetContentSourceUrlOk

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) GetContentSourceUrlOk() (*string, bool)`

GetContentSourceUrlOk returns a tuple with the ContentSourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentSourceUrl

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) SetContentSourceUrl(v string)`

SetContentSourceUrl sets ContentSourceUrl field to given value.

### HasContentSourceUrl

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) HasContentSourceUrl() bool`

HasContentSourceUrl returns a boolean if a field has been set.

### GetUrl

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetThumbUrl

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) GetThumbUrl() string`

GetThumbUrl returns the ThumbUrl field if non-nil, zero value otherwise.

### GetThumbUrlOk

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) GetThumbUrlOk() (*string, bool)`

GetThumbUrlOk returns a tuple with the ThumbUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbUrl

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) SetThumbUrl(v string)`

SetThumbUrl sets ThumbUrl field to given value.

### HasThumbUrl

`func (o *BatchgetMaterialResponseItemInnerContentNewsItemInner) HasThumbUrl() bool`

HasThumbUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


