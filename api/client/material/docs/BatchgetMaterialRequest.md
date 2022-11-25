# BatchgetMaterialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | 素材的类型，图片（image）、视频（video）、语音 （voice）、图文（news） | [optional] 
**Offset** | Pointer to **int32** | 从全部素材的该偏移位置开始返回，0表示从第一个素材 返回 | [optional] 
**Count** | Pointer to **int32** | 返回素材的数量，取值在1到20之间 | [optional] 

## Methods

### NewBatchgetMaterialRequest

`func NewBatchgetMaterialRequest() *BatchgetMaterialRequest`

NewBatchgetMaterialRequest instantiates a new BatchgetMaterialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchgetMaterialRequestWithDefaults

`func NewBatchgetMaterialRequestWithDefaults() *BatchgetMaterialRequest`

NewBatchgetMaterialRequestWithDefaults instantiates a new BatchgetMaterialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BatchgetMaterialRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BatchgetMaterialRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BatchgetMaterialRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BatchgetMaterialRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOffset

`func (o *BatchgetMaterialRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *BatchgetMaterialRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *BatchgetMaterialRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *BatchgetMaterialRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetCount

`func (o *BatchgetMaterialRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *BatchgetMaterialRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *BatchgetMaterialRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *BatchgetMaterialRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


