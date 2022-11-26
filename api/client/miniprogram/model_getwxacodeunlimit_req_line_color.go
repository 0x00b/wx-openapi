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

// checks if the GetwxacodeunlimitReqLineColor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetwxacodeunlimitReqLineColor{}

// GetwxacodeunlimitReqLineColor {\"r\":0,\"g\":0,\"b\":0} 否 auto_color 为 false 时生效，使用 rgb 设置颜色 例如 {\"r\":\"xxx\",\"g\":\"xxx\",\"b\":\"xxx\"} 十进制表示
type GetwxacodeunlimitReqLineColor struct {
	R *int32 `json:"r,omitempty"`
	G *int32 `json:"g,omitempty"`
	B *int32 `json:"b,omitempty"`
}

// NewGetwxacodeunlimitReqLineColor instantiates a new GetwxacodeunlimitReqLineColor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetwxacodeunlimitReqLineColor() *GetwxacodeunlimitReqLineColor {
	this := GetwxacodeunlimitReqLineColor{}
	return &this
}

// NewGetwxacodeunlimitReqLineColorWithDefaults instantiates a new GetwxacodeunlimitReqLineColor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetwxacodeunlimitReqLineColorWithDefaults() *GetwxacodeunlimitReqLineColor {
	this := GetwxacodeunlimitReqLineColor{}
	return &this
}

// GetR returns the R field value if set, zero value otherwise.
func (o *GetwxacodeunlimitReqLineColor) GetR() int32 {
	if o == nil || isNil(o.R) {
		var ret int32
		return ret
	}
	return *o.R
}

// GetROk returns a tuple with the R field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetwxacodeunlimitReqLineColor) GetROk() (*int32, bool) {
	if o == nil || isNil(o.R) {
		return nil, false
	}
	return o.R, true
}

// HasR returns a boolean if a field has been set.
func (o *GetwxacodeunlimitReqLineColor) HasR() bool {
	if o != nil && !isNil(o.R) {
		return true
	}

	return false
}

// SetR gets a reference to the given int32 and assigns it to the R field.
func (o *GetwxacodeunlimitReqLineColor) SetR(v int32) {
	o.R = &v
}

// GetG returns the G field value if set, zero value otherwise.
func (o *GetwxacodeunlimitReqLineColor) GetG() int32 {
	if o == nil || isNil(o.G) {
		var ret int32
		return ret
	}
	return *o.G
}

// GetGOk returns a tuple with the G field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetwxacodeunlimitReqLineColor) GetGOk() (*int32, bool) {
	if o == nil || isNil(o.G) {
		return nil, false
	}
	return o.G, true
}

// HasG returns a boolean if a field has been set.
func (o *GetwxacodeunlimitReqLineColor) HasG() bool {
	if o != nil && !isNil(o.G) {
		return true
	}

	return false
}

// SetG gets a reference to the given int32 and assigns it to the G field.
func (o *GetwxacodeunlimitReqLineColor) SetG(v int32) {
	o.G = &v
}

// GetB returns the B field value if set, zero value otherwise.
func (o *GetwxacodeunlimitReqLineColor) GetB() int32 {
	if o == nil || isNil(o.B) {
		var ret int32
		return ret
	}
	return *o.B
}

// GetBOk returns a tuple with the B field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetwxacodeunlimitReqLineColor) GetBOk() (*int32, bool) {
	if o == nil || isNil(o.B) {
		return nil, false
	}
	return o.B, true
}

// HasB returns a boolean if a field has been set.
func (o *GetwxacodeunlimitReqLineColor) HasB() bool {
	if o != nil && !isNil(o.B) {
		return true
	}

	return false
}

// SetB gets a reference to the given int32 and assigns it to the B field.
func (o *GetwxacodeunlimitReqLineColor) SetB(v int32) {
	o.B = &v
}

func (o GetwxacodeunlimitReqLineColor) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetwxacodeunlimitReqLineColor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.R) {
		toSerialize["r"] = o.R
	}
	if !isNil(o.G) {
		toSerialize["g"] = o.G
	}
	if !isNil(o.B) {
		toSerialize["b"] = o.B
	}
	return toSerialize, nil
}

type NullableGetwxacodeunlimitReqLineColor struct {
	value *GetwxacodeunlimitReqLineColor
	isSet bool
}

func (v NullableGetwxacodeunlimitReqLineColor) Get() *GetwxacodeunlimitReqLineColor {
	return v.value
}

func (v *NullableGetwxacodeunlimitReqLineColor) Set(val *GetwxacodeunlimitReqLineColor) {
	v.value = val
	v.isSet = true
}

func (v NullableGetwxacodeunlimitReqLineColor) IsSet() bool {
	return v.isSet
}

func (v *NullableGetwxacodeunlimitReqLineColor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetwxacodeunlimitReqLineColor(val *GetwxacodeunlimitReqLineColor) *NullableGetwxacodeunlimitReqLineColor {
	return &NullableGetwxacodeunlimitReqLineColor{value: val, isSet: true}
}

func (v NullableGetwxacodeunlimitReqLineColor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetwxacodeunlimitReqLineColor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
