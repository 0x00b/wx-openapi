/*
微信小程序相关接口

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package miniprogram

import (
	"encoding/json"
)

// checks if the TokenRsp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenRsp{}

// TokenRsp struct for TokenRsp
type TokenRsp struct {
	// 返回码
	Errcode *int32 `json:"errcode,omitempty"`
	// 对返回码的文本描述内容
	Errmsg *string `json:"errmsg,omitempty"`
	// access_token
	AccessToken *string `json:"access_token,omitempty"`
	// expires_in
	ExpiresIn *int32 `json:"expires_in,omitempty"`
}

// NewTokenRsp instantiates a new TokenRsp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenRsp() *TokenRsp {
	this := TokenRsp{}
	return &this
}

// NewTokenRspWithDefaults instantiates a new TokenRsp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRspWithDefaults() *TokenRsp {
	this := TokenRsp{}
	return &this
}

// GetErrcode returns the Errcode field value if set, zero value otherwise.
func (o *TokenRsp) GetErrcode() int32 {
	if o == nil || isNil(o.Errcode) {
		var ret int32
		return ret
	}
	return *o.Errcode
}

// GetErrcodeOk returns a tuple with the Errcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRsp) GetErrcodeOk() (*int32, bool) {
	if o == nil || isNil(o.Errcode) {
		return nil, false
	}
	return o.Errcode, true
}

// HasErrcode returns a boolean if a field has been set.
func (o *TokenRsp) HasErrcode() bool {
	if o != nil && !isNil(o.Errcode) {
		return true
	}

	return false
}

// SetErrcode gets a reference to the given int32 and assigns it to the Errcode field.
func (o *TokenRsp) SetErrcode(v int32) {
	o.Errcode = &v
}

// GetErrmsg returns the Errmsg field value if set, zero value otherwise.
func (o *TokenRsp) GetErrmsg() string {
	if o == nil || isNil(o.Errmsg) {
		var ret string
		return ret
	}
	return *o.Errmsg
}

// GetErrmsgOk returns a tuple with the Errmsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRsp) GetErrmsgOk() (*string, bool) {
	if o == nil || isNil(o.Errmsg) {
		return nil, false
	}
	return o.Errmsg, true
}

// HasErrmsg returns a boolean if a field has been set.
func (o *TokenRsp) HasErrmsg() bool {
	if o != nil && !isNil(o.Errmsg) {
		return true
	}

	return false
}

// SetErrmsg gets a reference to the given string and assigns it to the Errmsg field.
func (o *TokenRsp) SetErrmsg(v string) {
	o.Errmsg = &v
}

// GetAccessToken returns the AccessToken field value if set, zero value otherwise.
func (o *TokenRsp) GetAccessToken() string {
	if o == nil || isNil(o.AccessToken) {
		var ret string
		return ret
	}
	return *o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRsp) GetAccessTokenOk() (*string, bool) {
	if o == nil || isNil(o.AccessToken) {
		return nil, false
	}
	return o.AccessToken, true
}

// HasAccessToken returns a boolean if a field has been set.
func (o *TokenRsp) HasAccessToken() bool {
	if o != nil && !isNil(o.AccessToken) {
		return true
	}

	return false
}

// SetAccessToken gets a reference to the given string and assigns it to the AccessToken field.
func (o *TokenRsp) SetAccessToken(v string) {
	o.AccessToken = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *TokenRsp) GetExpiresIn() int32 {
	if o == nil || isNil(o.ExpiresIn) {
		var ret int32
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenRsp) GetExpiresInOk() (*int32, bool) {
	if o == nil || isNil(o.ExpiresIn) {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *TokenRsp) HasExpiresIn() bool {
	if o != nil && !isNil(o.ExpiresIn) {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int32 and assigns it to the ExpiresIn field.
func (o *TokenRsp) SetExpiresIn(v int32) {
	o.ExpiresIn = &v
}

func (o TokenRsp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenRsp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Errcode) {
		toSerialize["errcode"] = o.Errcode
	}
	if !isNil(o.Errmsg) {
		toSerialize["errmsg"] = o.Errmsg
	}
	if !isNil(o.AccessToken) {
		toSerialize["access_token"] = o.AccessToken
	}
	if !isNil(o.ExpiresIn) {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	return toSerialize, nil
}

type NullableTokenRsp struct {
	value *TokenRsp
	isSet bool
}

func (v NullableTokenRsp) Get() *TokenRsp {
	return v.value
}

func (v *NullableTokenRsp) Set(val *TokenRsp) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenRsp) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenRsp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenRsp(val *TokenRsp) *NullableTokenRsp {
	return &NullableTokenRsp{value: val, isSet: true}
}

func (v NullableTokenRsp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenRsp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}