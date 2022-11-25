# QueryAuthRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentAppid** | **string** | 平台型第三方平台的appid | 
**AuthorizationCode** | **string** | 授权码, 会在授权成功时返回给第三方平台，详见第三方平台授权流程说明 | 

## Methods

### NewQueryAuthRequest

`func NewQueryAuthRequest(componentAppid string, authorizationCode string, ) *QueryAuthRequest`

NewQueryAuthRequest instantiates a new QueryAuthRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryAuthRequestWithDefaults

`func NewQueryAuthRequestWithDefaults() *QueryAuthRequest`

NewQueryAuthRequestWithDefaults instantiates a new QueryAuthRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentAppid

`func (o *QueryAuthRequest) GetComponentAppid() string`

GetComponentAppid returns the ComponentAppid field if non-nil, zero value otherwise.

### GetComponentAppidOk

`func (o *QueryAuthRequest) GetComponentAppidOk() (*string, bool)`

GetComponentAppidOk returns a tuple with the ComponentAppid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentAppid

`func (o *QueryAuthRequest) SetComponentAppid(v string)`

SetComponentAppid sets ComponentAppid field to given value.


### GetAuthorizationCode

`func (o *QueryAuthRequest) GetAuthorizationCode() string`

GetAuthorizationCode returns the AuthorizationCode field if non-nil, zero value otherwise.

### GetAuthorizationCodeOk

`func (o *QueryAuthRequest) GetAuthorizationCodeOk() (*string, bool)`

GetAuthorizationCodeOk returns a tuple with the AuthorizationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationCode

`func (o *QueryAuthRequest) SetAuthorizationCode(v string)`

SetAuthorizationCode sets AuthorizationCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


