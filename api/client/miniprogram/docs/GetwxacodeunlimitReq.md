# GetwxacodeunlimitReq

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scene** | Pointer to **string** | 是 最大32个可见字符，只支持数字，大小写英文以及部分特殊字符：!#$&amp;&#39;()*+,/:;&#x3D;?@-._~，其它字符请自行编码为合法字符（因不支持%，中文无法使用 urlencode 处理，请使用其他编码方式） | [optional] 
**Page** | Pointer to **string** | 主页 否 页面 page，例如 pages/index/index，根路径前不要填加 /，不能携带参数（参数请放在scene字段里），如果不填写这个字段，默认跳主页面 | [optional] 
**CheckPath** | Pointer to **bool** | true 否 检查page 是否存在，为 true 时 page 必须是已经发布的小程序存在的页面（否则报错）；为 false 时允许小程序未发布或者 page 不存在， 但page 有数量上限（60000个）请勿滥用 | [optional] 
**EnvVersion** | Pointer to **string** | release\&quot; 否 要打开的小程序版本。正式版为 \&quot;release\&quot;，体验版为 \&quot;trial\&quot;，开发版为 \&quot;develop\&quot; | [optional] 
**Width** | Pointer to **int32** | 430 否 二维码的宽度，单位 px，最小 280px，最大 1280px | [optional] 
**AutoColor** | Pointer to **bool** | false 否 自动配置线条颜色，如果颜色依然是黑色，则说明不建议配置主色调，默认 false | [optional] 
**LineColor** | Pointer to [**GetwxacodeunlimitReqLineColor**](GetwxacodeunlimitReqLineColor.md) |  | [optional] 
**IsHyaline** | Pointer to **bool** | false 否 是否需要透明底色，为 true 时，生成透明底色的小程序 | [optional] 

## Methods

### NewGetwxacodeunlimitReq

`func NewGetwxacodeunlimitReq() *GetwxacodeunlimitReq`

NewGetwxacodeunlimitReq instantiates a new GetwxacodeunlimitReq object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetwxacodeunlimitReqWithDefaults

`func NewGetwxacodeunlimitReqWithDefaults() *GetwxacodeunlimitReq`

NewGetwxacodeunlimitReqWithDefaults instantiates a new GetwxacodeunlimitReq object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScene

`func (o *GetwxacodeunlimitReq) GetScene() string`

GetScene returns the Scene field if non-nil, zero value otherwise.

### GetSceneOk

`func (o *GetwxacodeunlimitReq) GetSceneOk() (*string, bool)`

GetSceneOk returns a tuple with the Scene field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScene

`func (o *GetwxacodeunlimitReq) SetScene(v string)`

SetScene sets Scene field to given value.

### HasScene

`func (o *GetwxacodeunlimitReq) HasScene() bool`

HasScene returns a boolean if a field has been set.

### GetPage

`func (o *GetwxacodeunlimitReq) GetPage() string`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *GetwxacodeunlimitReq) GetPageOk() (*string, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *GetwxacodeunlimitReq) SetPage(v string)`

SetPage sets Page field to given value.

### HasPage

`func (o *GetwxacodeunlimitReq) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetCheckPath

`func (o *GetwxacodeunlimitReq) GetCheckPath() bool`

GetCheckPath returns the CheckPath field if non-nil, zero value otherwise.

### GetCheckPathOk

`func (o *GetwxacodeunlimitReq) GetCheckPathOk() (*bool, bool)`

GetCheckPathOk returns a tuple with the CheckPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPath

`func (o *GetwxacodeunlimitReq) SetCheckPath(v bool)`

SetCheckPath sets CheckPath field to given value.

### HasCheckPath

`func (o *GetwxacodeunlimitReq) HasCheckPath() bool`

HasCheckPath returns a boolean if a field has been set.

### GetEnvVersion

`func (o *GetwxacodeunlimitReq) GetEnvVersion() string`

GetEnvVersion returns the EnvVersion field if non-nil, zero value otherwise.

### GetEnvVersionOk

`func (o *GetwxacodeunlimitReq) GetEnvVersionOk() (*string, bool)`

GetEnvVersionOk returns a tuple with the EnvVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvVersion

`func (o *GetwxacodeunlimitReq) SetEnvVersion(v string)`

SetEnvVersion sets EnvVersion field to given value.

### HasEnvVersion

`func (o *GetwxacodeunlimitReq) HasEnvVersion() bool`

HasEnvVersion returns a boolean if a field has been set.

### GetWidth

`func (o *GetwxacodeunlimitReq) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *GetwxacodeunlimitReq) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *GetwxacodeunlimitReq) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *GetwxacodeunlimitReq) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetAutoColor

`func (o *GetwxacodeunlimitReq) GetAutoColor() bool`

GetAutoColor returns the AutoColor field if non-nil, zero value otherwise.

### GetAutoColorOk

`func (o *GetwxacodeunlimitReq) GetAutoColorOk() (*bool, bool)`

GetAutoColorOk returns a tuple with the AutoColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoColor

`func (o *GetwxacodeunlimitReq) SetAutoColor(v bool)`

SetAutoColor sets AutoColor field to given value.

### HasAutoColor

`func (o *GetwxacodeunlimitReq) HasAutoColor() bool`

HasAutoColor returns a boolean if a field has been set.

### GetLineColor

`func (o *GetwxacodeunlimitReq) GetLineColor() GetwxacodeunlimitReqLineColor`

GetLineColor returns the LineColor field if non-nil, zero value otherwise.

### GetLineColorOk

`func (o *GetwxacodeunlimitReq) GetLineColorOk() (*GetwxacodeunlimitReqLineColor, bool)`

GetLineColorOk returns a tuple with the LineColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineColor

`func (o *GetwxacodeunlimitReq) SetLineColor(v GetwxacodeunlimitReqLineColor)`

SetLineColor sets LineColor field to given value.

### HasLineColor

`func (o *GetwxacodeunlimitReq) HasLineColor() bool`

HasLineColor returns a boolean if a field has been set.

### GetIsHyaline

`func (o *GetwxacodeunlimitReq) GetIsHyaline() bool`

GetIsHyaline returns the IsHyaline field if non-nil, zero value otherwise.

### GetIsHyalineOk

`func (o *GetwxacodeunlimitReq) GetIsHyalineOk() (*bool, bool)`

GetIsHyalineOk returns a tuple with the IsHyaline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHyaline

`func (o *GetwxacodeunlimitReq) SetIsHyaline(v bool)`

SetIsHyaline sets IsHyaline field to given value.

### HasIsHyaline

`func (o *GetwxacodeunlimitReq) HasIsHyaline() bool`

HasIsHyaline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


