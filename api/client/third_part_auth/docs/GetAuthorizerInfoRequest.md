# GetAuthorizerInfoRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentAppid** | **string** | 平台型第三方平台的appid | 
**AuthorizerAppid** | **string** | 授权方 appid | 

## Methods

### NewGetAuthorizerInfoRequest

`func NewGetAuthorizerInfoRequest(componentAppid string, authorizerAppid string, ) *GetAuthorizerInfoRequest`

NewGetAuthorizerInfoRequest instantiates a new GetAuthorizerInfoRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAuthorizerInfoRequestWithDefaults

`func NewGetAuthorizerInfoRequestWithDefaults() *GetAuthorizerInfoRequest`

NewGetAuthorizerInfoRequestWithDefaults instantiates a new GetAuthorizerInfoRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentAppid

`func (o *GetAuthorizerInfoRequest) GetComponentAppid() string`

GetComponentAppid returns the ComponentAppid field if non-nil, zero value otherwise.

### GetComponentAppidOk

`func (o *GetAuthorizerInfoRequest) GetComponentAppidOk() (*string, bool)`

GetComponentAppidOk returns a tuple with the ComponentAppid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentAppid

`func (o *GetAuthorizerInfoRequest) SetComponentAppid(v string)`

SetComponentAppid sets ComponentAppid field to given value.


### GetAuthorizerAppid

`func (o *GetAuthorizerInfoRequest) GetAuthorizerAppid() string`

GetAuthorizerAppid returns the AuthorizerAppid field if non-nil, zero value otherwise.

### GetAuthorizerAppidOk

`func (o *GetAuthorizerInfoRequest) GetAuthorizerAppidOk() (*string, bool)`

GetAuthorizerAppidOk returns a tuple with the AuthorizerAppid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizerAppid

`func (o *GetAuthorizerInfoRequest) SetAuthorizerAppid(v string)`

SetAuthorizerAppid sets AuthorizerAppid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


