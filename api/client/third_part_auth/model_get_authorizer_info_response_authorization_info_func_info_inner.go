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

// checks if the GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner{}

// GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner struct for GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner
type GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner struct {
	FuncscopeCategory *GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInnerFuncscopeCategory `json:"funcscope_category,omitempty"`
}

// NewGetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner instantiates a new GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner() *GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner {
	this := GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner{}
	return &this
}

// NewGetAuthorizerInfoResponseAuthorizationInfoFuncInfoInnerWithDefaults instantiates a new GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAuthorizerInfoResponseAuthorizationInfoFuncInfoInnerWithDefaults() *GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner {
	this := GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner{}
	return &this
}

// GetFuncscopeCategory returns the FuncscopeCategory field value if set, zero value otherwise.
func (o *GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner) GetFuncscopeCategory() GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInnerFuncscopeCategory {
	if o == nil || isNil(o.FuncscopeCategory) {
		var ret GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInnerFuncscopeCategory
		return ret
	}
	return *o.FuncscopeCategory
}

// GetFuncscopeCategoryOk returns a tuple with the FuncscopeCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner) GetFuncscopeCategoryOk() (*GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInnerFuncscopeCategory, bool) {
	if o == nil || isNil(o.FuncscopeCategory) {
		return nil, false
	}
	return o.FuncscopeCategory, true
}

// HasFuncscopeCategory returns a boolean if a field has been set.
func (o *GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner) HasFuncscopeCategory() bool {
	if o != nil && !isNil(o.FuncscopeCategory) {
		return true
	}

	return false
}

// SetFuncscopeCategory gets a reference to the given GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInnerFuncscopeCategory and assigns it to the FuncscopeCategory field.
func (o *GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner) SetFuncscopeCategory(v GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInnerFuncscopeCategory) {
	o.FuncscopeCategory = &v
}

func (o GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FuncscopeCategory) {
		toSerialize["funcscope_category"] = o.FuncscopeCategory
	}
	return toSerialize, nil
}

type NullableGetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner struct {
	value *GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner
	isSet bool
}

func (v NullableGetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner) Get() *GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner {
	return v.value
}

func (v *NullableGetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner) Set(val *GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner(val *GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner) *NullableGetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner {
	return &NullableGetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner{value: val, isSet: true}
}

func (v NullableGetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
