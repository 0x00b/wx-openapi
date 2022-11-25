/*
微信第三方平台授权相关接口

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package third_part_auth

import (
	"encoding/json"
)

// checks if the QueryAuthResponseAuthorizationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueryAuthResponseAuthorizationInfo{}

// QueryAuthResponseAuthorizationInfo struct for QueryAuthResponseAuthorizationInfo
type QueryAuthResponseAuthorizationInfo struct {
	// 授权方 appid
	AuthorizerAppid *string `json:"authorizer_appid,omitempty"`
	// 接口调用令牌（在授权的公众号/小程序具备 API 权限时，才有此返回值）
	AuthorizerAccessToken *string `json:"authorizer_access_token,omitempty"`
	// authorizer_access_token 的有效期（在授权的公众号/小程序具备API权限时，才有此返回值），单位：秒
	ExpiresIn *int32 `json:"expires_in,omitempty"`
	// 刷新令牌（在授权的公众号具备API权限时，才有此返回值），刷新令牌主要用于第三方平台获取和刷新已授权用户的 authorizer_access_token。一旦丢失，只能让用户重新授权，才能再次拿到新的刷新令牌。用户重新授权后，之前的刷新令牌会失效
	AuthorizerRefreshToken *string `json:"authorizer_refresh_token,omitempty"`
	// 授权给开发者的权限集列表
	FuncInfo []GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner `json:"func_info,omitempty"`
}

// NewQueryAuthResponseAuthorizationInfo instantiates a new QueryAuthResponseAuthorizationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryAuthResponseAuthorizationInfo() *QueryAuthResponseAuthorizationInfo {
	this := QueryAuthResponseAuthorizationInfo{}
	return &this
}

// NewQueryAuthResponseAuthorizationInfoWithDefaults instantiates a new QueryAuthResponseAuthorizationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryAuthResponseAuthorizationInfoWithDefaults() *QueryAuthResponseAuthorizationInfo {
	this := QueryAuthResponseAuthorizationInfo{}
	return &this
}

// GetAuthorizerAppid returns the AuthorizerAppid field value if set, zero value otherwise.
func (o *QueryAuthResponseAuthorizationInfo) GetAuthorizerAppid() string {
	if o == nil || isNil(o.AuthorizerAppid) {
		var ret string
		return ret
	}
	return *o.AuthorizerAppid
}

// GetAuthorizerAppidOk returns a tuple with the AuthorizerAppid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAuthResponseAuthorizationInfo) GetAuthorizerAppidOk() (*string, bool) {
	if o == nil || isNil(o.AuthorizerAppid) {
		return nil, false
	}
	return o.AuthorizerAppid, true
}

// HasAuthorizerAppid returns a boolean if a field has been set.
func (o *QueryAuthResponseAuthorizationInfo) HasAuthorizerAppid() bool {
	if o != nil && !isNil(o.AuthorizerAppid) {
		return true
	}

	return false
}

// SetAuthorizerAppid gets a reference to the given string and assigns it to the AuthorizerAppid field.
func (o *QueryAuthResponseAuthorizationInfo) SetAuthorizerAppid(v string) {
	o.AuthorizerAppid = &v
}

// GetAuthorizerAccessToken returns the AuthorizerAccessToken field value if set, zero value otherwise.
func (o *QueryAuthResponseAuthorizationInfo) GetAuthorizerAccessToken() string {
	if o == nil || isNil(o.AuthorizerAccessToken) {
		var ret string
		return ret
	}
	return *o.AuthorizerAccessToken
}

// GetAuthorizerAccessTokenOk returns a tuple with the AuthorizerAccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAuthResponseAuthorizationInfo) GetAuthorizerAccessTokenOk() (*string, bool) {
	if o == nil || isNil(o.AuthorizerAccessToken) {
		return nil, false
	}
	return o.AuthorizerAccessToken, true
}

// HasAuthorizerAccessToken returns a boolean if a field has been set.
func (o *QueryAuthResponseAuthorizationInfo) HasAuthorizerAccessToken() bool {
	if o != nil && !isNil(o.AuthorizerAccessToken) {
		return true
	}

	return false
}

// SetAuthorizerAccessToken gets a reference to the given string and assigns it to the AuthorizerAccessToken field.
func (o *QueryAuthResponseAuthorizationInfo) SetAuthorizerAccessToken(v string) {
	o.AuthorizerAccessToken = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *QueryAuthResponseAuthorizationInfo) GetExpiresIn() int32 {
	if o == nil || isNil(o.ExpiresIn) {
		var ret int32
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAuthResponseAuthorizationInfo) GetExpiresInOk() (*int32, bool) {
	if o == nil || isNil(o.ExpiresIn) {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *QueryAuthResponseAuthorizationInfo) HasExpiresIn() bool {
	if o != nil && !isNil(o.ExpiresIn) {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int32 and assigns it to the ExpiresIn field.
func (o *QueryAuthResponseAuthorizationInfo) SetExpiresIn(v int32) {
	o.ExpiresIn = &v
}

// GetAuthorizerRefreshToken returns the AuthorizerRefreshToken field value if set, zero value otherwise.
func (o *QueryAuthResponseAuthorizationInfo) GetAuthorizerRefreshToken() string {
	if o == nil || isNil(o.AuthorizerRefreshToken) {
		var ret string
		return ret
	}
	return *o.AuthorizerRefreshToken
}

// GetAuthorizerRefreshTokenOk returns a tuple with the AuthorizerRefreshToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAuthResponseAuthorizationInfo) GetAuthorizerRefreshTokenOk() (*string, bool) {
	if o == nil || isNil(o.AuthorizerRefreshToken) {
		return nil, false
	}
	return o.AuthorizerRefreshToken, true
}

// HasAuthorizerRefreshToken returns a boolean if a field has been set.
func (o *QueryAuthResponseAuthorizationInfo) HasAuthorizerRefreshToken() bool {
	if o != nil && !isNil(o.AuthorizerRefreshToken) {
		return true
	}

	return false
}

// SetAuthorizerRefreshToken gets a reference to the given string and assigns it to the AuthorizerRefreshToken field.
func (o *QueryAuthResponseAuthorizationInfo) SetAuthorizerRefreshToken(v string) {
	o.AuthorizerRefreshToken = &v
}

// GetFuncInfo returns the FuncInfo field value if set, zero value otherwise.
func (o *QueryAuthResponseAuthorizationInfo) GetFuncInfo() []GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner {
	if o == nil || isNil(o.FuncInfo) {
		var ret []GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner
		return ret
	}
	return o.FuncInfo
}

// GetFuncInfoOk returns a tuple with the FuncInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueryAuthResponseAuthorizationInfo) GetFuncInfoOk() ([]GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner, bool) {
	if o == nil || isNil(o.FuncInfo) {
		return nil, false
	}
	return o.FuncInfo, true
}

// HasFuncInfo returns a boolean if a field has been set.
func (o *QueryAuthResponseAuthorizationInfo) HasFuncInfo() bool {
	if o != nil && !isNil(o.FuncInfo) {
		return true
	}

	return false
}

// SetFuncInfo gets a reference to the given []GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner and assigns it to the FuncInfo field.
func (o *QueryAuthResponseAuthorizationInfo) SetFuncInfo(v []GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner) {
	o.FuncInfo = v
}

func (o QueryAuthResponseAuthorizationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueryAuthResponseAuthorizationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AuthorizerAppid) {
		toSerialize["authorizer_appid"] = o.AuthorizerAppid
	}
	if !isNil(o.AuthorizerAccessToken) {
		toSerialize["authorizer_access_token"] = o.AuthorizerAccessToken
	}
	if !isNil(o.ExpiresIn) {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if !isNil(o.AuthorizerRefreshToken) {
		toSerialize["authorizer_refresh_token"] = o.AuthorizerRefreshToken
	}
	if !isNil(o.FuncInfo) {
		toSerialize["func_info"] = o.FuncInfo
	}
	return toSerialize, nil
}

type NullableQueryAuthResponseAuthorizationInfo struct {
	value *QueryAuthResponseAuthorizationInfo
	isSet bool
}

func (v NullableQueryAuthResponseAuthorizationInfo) Get() *QueryAuthResponseAuthorizationInfo {
	return v.value
}

func (v *NullableQueryAuthResponseAuthorizationInfo) Set(val *QueryAuthResponseAuthorizationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryAuthResponseAuthorizationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryAuthResponseAuthorizationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryAuthResponseAuthorizationInfo(val *QueryAuthResponseAuthorizationInfo) *NullableQueryAuthResponseAuthorizationInfo {
	return &NullableQueryAuthResponseAuthorizationInfo{value: val, isSet: true}
}

func (v NullableQueryAuthResponseAuthorizationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryAuthResponseAuthorizationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
