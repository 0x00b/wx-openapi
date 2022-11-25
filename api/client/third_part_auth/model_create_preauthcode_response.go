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

// checks if the CreatePreauthcodeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePreauthcodeResponse{}

// CreatePreauthcodeResponse struct for CreatePreauthcodeResponse
type CreatePreauthcodeResponse struct {
	// 返回码
	Errcode *int32 `json:"errcode,omitempty"`
	// 对返回码的文本描述内容
	Errmsg *string `json:"errmsg,omitempty"`
	// 有效期，单位：秒
	ExpiresIn *int32 `json:"expires_in,omitempty"`
	// 预授权码
	PreAuthCode *string `json:"pre_auth_code,omitempty"`
}

// NewCreatePreauthcodeResponse instantiates a new CreatePreauthcodeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePreauthcodeResponse() *CreatePreauthcodeResponse {
	this := CreatePreauthcodeResponse{}
	return &this
}

// NewCreatePreauthcodeResponseWithDefaults instantiates a new CreatePreauthcodeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePreauthcodeResponseWithDefaults() *CreatePreauthcodeResponse {
	this := CreatePreauthcodeResponse{}
	return &this
}

// GetErrcode returns the Errcode field value if set, zero value otherwise.
func (o *CreatePreauthcodeResponse) GetErrcode() int32 {
	if o == nil || isNil(o.Errcode) {
		var ret int32
		return ret
	}
	return *o.Errcode
}

// GetErrcodeOk returns a tuple with the Errcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePreauthcodeResponse) GetErrcodeOk() (*int32, bool) {
	if o == nil || isNil(o.Errcode) {
		return nil, false
	}
	return o.Errcode, true
}

// HasErrcode returns a boolean if a field has been set.
func (o *CreatePreauthcodeResponse) HasErrcode() bool {
	if o != nil && !isNil(o.Errcode) {
		return true
	}

	return false
}

// SetErrcode gets a reference to the given int32 and assigns it to the Errcode field.
func (o *CreatePreauthcodeResponse) SetErrcode(v int32) {
	o.Errcode = &v
}

// GetErrmsg returns the Errmsg field value if set, zero value otherwise.
func (o *CreatePreauthcodeResponse) GetErrmsg() string {
	if o == nil || isNil(o.Errmsg) {
		var ret string
		return ret
	}
	return *o.Errmsg
}

// GetErrmsgOk returns a tuple with the Errmsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePreauthcodeResponse) GetErrmsgOk() (*string, bool) {
	if o == nil || isNil(o.Errmsg) {
		return nil, false
	}
	return o.Errmsg, true
}

// HasErrmsg returns a boolean if a field has been set.
func (o *CreatePreauthcodeResponse) HasErrmsg() bool {
	if o != nil && !isNil(o.Errmsg) {
		return true
	}

	return false
}

// SetErrmsg gets a reference to the given string and assigns it to the Errmsg field.
func (o *CreatePreauthcodeResponse) SetErrmsg(v string) {
	o.Errmsg = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *CreatePreauthcodeResponse) GetExpiresIn() int32 {
	if o == nil || isNil(o.ExpiresIn) {
		var ret int32
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePreauthcodeResponse) GetExpiresInOk() (*int32, bool) {
	if o == nil || isNil(o.ExpiresIn) {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *CreatePreauthcodeResponse) HasExpiresIn() bool {
	if o != nil && !isNil(o.ExpiresIn) {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int32 and assigns it to the ExpiresIn field.
func (o *CreatePreauthcodeResponse) SetExpiresIn(v int32) {
	o.ExpiresIn = &v
}

// GetPreAuthCode returns the PreAuthCode field value if set, zero value otherwise.
func (o *CreatePreauthcodeResponse) GetPreAuthCode() string {
	if o == nil || isNil(o.PreAuthCode) {
		var ret string
		return ret
	}
	return *o.PreAuthCode
}

// GetPreAuthCodeOk returns a tuple with the PreAuthCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatePreauthcodeResponse) GetPreAuthCodeOk() (*string, bool) {
	if o == nil || isNil(o.PreAuthCode) {
		return nil, false
	}
	return o.PreAuthCode, true
}

// HasPreAuthCode returns a boolean if a field has been set.
func (o *CreatePreauthcodeResponse) HasPreAuthCode() bool {
	if o != nil && !isNil(o.PreAuthCode) {
		return true
	}

	return false
}

// SetPreAuthCode gets a reference to the given string and assigns it to the PreAuthCode field.
func (o *CreatePreauthcodeResponse) SetPreAuthCode(v string) {
	o.PreAuthCode = &v
}

func (o CreatePreauthcodeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePreauthcodeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Errcode) {
		toSerialize["errcode"] = o.Errcode
	}
	if !isNil(o.Errmsg) {
		toSerialize["errmsg"] = o.Errmsg
	}
	if !isNil(o.ExpiresIn) {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	if !isNil(o.PreAuthCode) {
		toSerialize["pre_auth_code"] = o.PreAuthCode
	}
	return toSerialize, nil
}

type NullableCreatePreauthcodeResponse struct {
	value *CreatePreauthcodeResponse
	isSet bool
}

func (v NullableCreatePreauthcodeResponse) Get() *CreatePreauthcodeResponse {
	return v.value
}

func (v *NullableCreatePreauthcodeResponse) Set(val *CreatePreauthcodeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePreauthcodeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePreauthcodeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePreauthcodeResponse(val *CreatePreauthcodeResponse) *NullableCreatePreauthcodeResponse {
	return &NullableCreatePreauthcodeResponse{value: val, isSet: true}
}

func (v NullableCreatePreauthcodeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePreauthcodeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}