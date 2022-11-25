# GetAuthorizerInfoResponseAuthorizerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NickName** | Pointer to **string** | 昵称 | [optional] 
**HeadImg** | Pointer to **string** | 头像 | [optional] 
**ServiceTypeInfo** | Pointer to [**GetAuthorizerInfoResponseAuthorizerInfoServiceTypeInfo**](GetAuthorizerInfoResponseAuthorizerInfoServiceTypeInfo.md) |  | [optional] 
**VerifyTypeInfo** | Pointer to [**GetAuthorizerInfoResponseAuthorizerInfoVerifyTypeInfo**](GetAuthorizerInfoResponseAuthorizerInfoVerifyTypeInfo.md) |  | [optional] 
**UserName** | Pointer to **string** | 原始 ID | [optional] 
**PrincipalName** | Pointer to **string** | 主体名称 | [optional] 
**Alias** | Pointer to **string** | 公众号所设置的微信号，可能为空 | [optional] 
**BusinessInfo** | Pointer to [**GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo**](GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo.md) |  | [optional] 
**QrcodeUrl** | Pointer to **string** | 二维码图片的 URL，开发者最好自行也进行保存 | [optional] 
**Signature** | Pointer to **string** | 帐号介绍 | [optional] 
**MiniProgramInfo** | Pointer to [**GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo**](GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo.md) |  | [optional] 

## Methods

### NewGetAuthorizerInfoResponseAuthorizerInfo

`func NewGetAuthorizerInfoResponseAuthorizerInfo() *GetAuthorizerInfoResponseAuthorizerInfo`

NewGetAuthorizerInfoResponseAuthorizerInfo instantiates a new GetAuthorizerInfoResponseAuthorizerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAuthorizerInfoResponseAuthorizerInfoWithDefaults

`func NewGetAuthorizerInfoResponseAuthorizerInfoWithDefaults() *GetAuthorizerInfoResponseAuthorizerInfo`

NewGetAuthorizerInfoResponseAuthorizerInfoWithDefaults instantiates a new GetAuthorizerInfoResponseAuthorizerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNickName

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) GetNickName() string`

GetNickName returns the NickName field if non-nil, zero value otherwise.

### GetNickNameOk

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) GetNickNameOk() (*string, bool)`

GetNickNameOk returns a tuple with the NickName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNickName

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) SetNickName(v string)`

SetNickName sets NickName field to given value.

### HasNickName

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) HasNickName() bool`

HasNickName returns a boolean if a field has been set.

### GetHeadImg

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) GetHeadImg() string`

GetHeadImg returns the HeadImg field if non-nil, zero value otherwise.

### GetHeadImgOk

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) GetHeadImgOk() (*string, bool)`

GetHeadImgOk returns a tuple with the HeadImg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadImg

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) SetHeadImg(v string)`

SetHeadImg sets HeadImg field to given value.

### HasHeadImg

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) HasHeadImg() bool`

HasHeadImg returns a boolean if a field has been set.

### GetServiceTypeInfo

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) GetServiceTypeInfo() GetAuthorizerInfoResponseAuthorizerInfoServiceTypeInfo`

GetServiceTypeInfo returns the ServiceTypeInfo field if non-nil, zero value otherwise.

### GetServiceTypeInfoOk

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) GetServiceTypeInfoOk() (*GetAuthorizerInfoResponseAuthorizerInfoServiceTypeInfo, bool)`

GetServiceTypeInfoOk returns a tuple with the ServiceTypeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTypeInfo

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) SetServiceTypeInfo(v GetAuthorizerInfoResponseAuthorizerInfoServiceTypeInfo)`

SetServiceTypeInfo sets ServiceTypeInfo field to given value.

### HasServiceTypeInfo

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) HasServiceTypeInfo() bool`

HasServiceTypeInfo returns a boolean if a field has been set.

### GetVerifyTypeInfo

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) GetVerifyTypeInfo() GetAuthorizerInfoResponseAuthorizerInfoVerifyTypeInfo`

GetVerifyTypeInfo returns the VerifyTypeInfo field if non-nil, zero value otherwise.

### GetVerifyTypeInfoOk

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) GetVerifyTypeInfoOk() (*GetAuthorizerInfoResponseAuthorizerInfoVerifyTypeInfo, bool)`

GetVerifyTypeInfoOk returns a tuple with the VerifyTypeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyTypeInfo

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) SetVerifyTypeInfo(v GetAuthorizerInfoResponseAuthorizerInfoVerifyTypeInfo)`

SetVerifyTypeInfo sets VerifyTypeInfo field to given value.

### HasVerifyTypeInfo

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) HasVerifyTypeInfo() bool`

HasVerifyTypeInfo returns a boolean if a field has been set.

### GetUserName

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetPrincipalName

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) GetPrincipalName() string`

GetPrincipalName returns the PrincipalName field if non-nil, zero value otherwise.

### GetPrincipalNameOk

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) GetPrincipalNameOk() (*string, bool)`

GetPrincipalNameOk returns a tuple with the PrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalName

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) SetPrincipalName(v string)`

SetPrincipalName sets PrincipalName field to given value.

### HasPrincipalName

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) HasPrincipalName() bool`

HasPrincipalName returns a boolean if a field has been set.

### GetAlias

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetBusinessInfo

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) GetBusinessInfo() GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo`

GetBusinessInfo returns the BusinessInfo field if non-nil, zero value otherwise.

### GetBusinessInfoOk

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) GetBusinessInfoOk() (*GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo, bool)`

GetBusinessInfoOk returns a tuple with the BusinessInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessInfo

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) SetBusinessInfo(v GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo)`

SetBusinessInfo sets BusinessInfo field to given value.

### HasBusinessInfo

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) HasBusinessInfo() bool`

HasBusinessInfo returns a boolean if a field has been set.

### GetQrcodeUrl

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) GetQrcodeUrl() string`

GetQrcodeUrl returns the QrcodeUrl field if non-nil, zero value otherwise.

### GetQrcodeUrlOk

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) GetQrcodeUrlOk() (*string, bool)`

GetQrcodeUrlOk returns a tuple with the QrcodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrcodeUrl

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) SetQrcodeUrl(v string)`

SetQrcodeUrl sets QrcodeUrl field to given value.

### HasQrcodeUrl

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) HasQrcodeUrl() bool`

HasQrcodeUrl returns a boolean if a field has been set.

### GetSignature

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetMiniProgramInfo

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) GetMiniProgramInfo() GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo`

GetMiniProgramInfo returns the MiniProgramInfo field if non-nil, zero value otherwise.

### GetMiniProgramInfoOk

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) GetMiniProgramInfoOk() (*GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo, bool)`

GetMiniProgramInfoOk returns a tuple with the MiniProgramInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiniProgramInfo

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) SetMiniProgramInfo(v GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo)`

SetMiniProgramInfo sets MiniProgramInfo field to given value.

### HasMiniProgramInfo

`func (o *GetAuthorizerInfoResponseAuthorizerInfo) HasMiniProgramInfo() bool`

HasMiniProgramInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


