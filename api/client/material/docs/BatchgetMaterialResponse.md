# BatchgetMaterialResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errcode** | Pointer to **int32** | 返回码 | [optional] 
**Errmsg** | Pointer to **string** | 对返回码的文本描述内容 | [optional] 
**TotalCount** | Pointer to **int32** | 该类型的素材的总数 | [optional] 
**ItemCount** | Pointer to **int32** | 本次调用获取的素材的数量 | [optional] 
**Item** | Pointer to [**[]BatchgetMaterialResponseItemInner**](BatchgetMaterialResponseItemInner.md) |  | [optional] 

## Methods

### NewBatchgetMaterialResponse

`func NewBatchgetMaterialResponse() *BatchgetMaterialResponse`

NewBatchgetMaterialResponse instantiates a new BatchgetMaterialResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchgetMaterialResponseWithDefaults

`func NewBatchgetMaterialResponseWithDefaults() *BatchgetMaterialResponse`

NewBatchgetMaterialResponseWithDefaults instantiates a new BatchgetMaterialResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrcode

`func (o *BatchgetMaterialResponse) GetErrcode() int32`

GetErrcode returns the Errcode field if non-nil, zero value otherwise.

### GetErrcodeOk

`func (o *BatchgetMaterialResponse) GetErrcodeOk() (*int32, bool)`

GetErrcodeOk returns a tuple with the Errcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrcode

`func (o *BatchgetMaterialResponse) SetErrcode(v int32)`

SetErrcode sets Errcode field to given value.

### HasErrcode

`func (o *BatchgetMaterialResponse) HasErrcode() bool`

HasErrcode returns a boolean if a field has been set.

### GetErrmsg

`func (o *BatchgetMaterialResponse) GetErrmsg() string`

GetErrmsg returns the Errmsg field if non-nil, zero value otherwise.

### GetErrmsgOk

`func (o *BatchgetMaterialResponse) GetErrmsgOk() (*string, bool)`

GetErrmsgOk returns a tuple with the Errmsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrmsg

`func (o *BatchgetMaterialResponse) SetErrmsg(v string)`

SetErrmsg sets Errmsg field to given value.

### HasErrmsg

`func (o *BatchgetMaterialResponse) HasErrmsg() bool`

HasErrmsg returns a boolean if a field has been set.

### GetTotalCount

`func (o *BatchgetMaterialResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *BatchgetMaterialResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *BatchgetMaterialResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *BatchgetMaterialResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetItemCount

`func (o *BatchgetMaterialResponse) GetItemCount() int32`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *BatchgetMaterialResponse) GetItemCountOk() (*int32, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *BatchgetMaterialResponse) SetItemCount(v int32)`

SetItemCount sets ItemCount field to given value.

### HasItemCount

`func (o *BatchgetMaterialResponse) HasItemCount() bool`

HasItemCount returns a boolean if a field has been set.

### GetItem

`func (o *BatchgetMaterialResponse) GetItem() []BatchgetMaterialResponseItemInner`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *BatchgetMaterialResponse) GetItemOk() (*[]BatchgetMaterialResponseItemInner, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *BatchgetMaterialResponse) SetItem(v []BatchgetMaterialResponseItemInner)`

SetItem sets Item field to given value.

### HasItem

`func (o *BatchgetMaterialResponse) HasItem() bool`

HasItem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


