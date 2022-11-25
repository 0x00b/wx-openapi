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

// checks if the GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo{}

// GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo 用以了解功能的开通状况（0代表未开通，1代表已开通），详见business_info 说明
type GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo struct {
	// 是否开通微信门店功能
	OpenStore *int32 `json:"open_store,omitempty"`
	// 是否开通微信扫商品功能
	OpenScan *int32 `json:"open_scan,omitempty"`
	// 是否开通微信支付功能
	OpenPay *int32 `json:"open_pay,omitempty"`
	// 是否开通微信卡券功能
	OpenCard *int32 `json:"open_card,omitempty"`
	// 是否开通微信摇一摇功能
	OpenShake *int32 `json:"open_shake,omitempty"`
}

// NewGetAuthorizerInfoResponseAuthorizerInfoBusinessInfo instantiates a new GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAuthorizerInfoResponseAuthorizerInfoBusinessInfo() *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo {
	this := GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo{}
	return &this
}

// NewGetAuthorizerInfoResponseAuthorizerInfoBusinessInfoWithDefaults instantiates a new GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAuthorizerInfoResponseAuthorizerInfoBusinessInfoWithDefaults() *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo {
	this := GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo{}
	return &this
}

// GetOpenStore returns the OpenStore field value if set, zero value otherwise.
func (o *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) GetOpenStore() int32 {
	if o == nil || isNil(o.OpenStore) {
		var ret int32
		return ret
	}
	return *o.OpenStore
}

// GetOpenStoreOk returns a tuple with the OpenStore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) GetOpenStoreOk() (*int32, bool) {
	if o == nil || isNil(o.OpenStore) {
		return nil, false
	}
	return o.OpenStore, true
}

// HasOpenStore returns a boolean if a field has been set.
func (o *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) HasOpenStore() bool {
	if o != nil && !isNil(o.OpenStore) {
		return true
	}

	return false
}

// SetOpenStore gets a reference to the given int32 and assigns it to the OpenStore field.
func (o *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) SetOpenStore(v int32) {
	o.OpenStore = &v
}

// GetOpenScan returns the OpenScan field value if set, zero value otherwise.
func (o *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) GetOpenScan() int32 {
	if o == nil || isNil(o.OpenScan) {
		var ret int32
		return ret
	}
	return *o.OpenScan
}

// GetOpenScanOk returns a tuple with the OpenScan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) GetOpenScanOk() (*int32, bool) {
	if o == nil || isNil(o.OpenScan) {
		return nil, false
	}
	return o.OpenScan, true
}

// HasOpenScan returns a boolean if a field has been set.
func (o *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) HasOpenScan() bool {
	if o != nil && !isNil(o.OpenScan) {
		return true
	}

	return false
}

// SetOpenScan gets a reference to the given int32 and assigns it to the OpenScan field.
func (o *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) SetOpenScan(v int32) {
	o.OpenScan = &v
}

// GetOpenPay returns the OpenPay field value if set, zero value otherwise.
func (o *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) GetOpenPay() int32 {
	if o == nil || isNil(o.OpenPay) {
		var ret int32
		return ret
	}
	return *o.OpenPay
}

// GetOpenPayOk returns a tuple with the OpenPay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) GetOpenPayOk() (*int32, bool) {
	if o == nil || isNil(o.OpenPay) {
		return nil, false
	}
	return o.OpenPay, true
}

// HasOpenPay returns a boolean if a field has been set.
func (o *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) HasOpenPay() bool {
	if o != nil && !isNil(o.OpenPay) {
		return true
	}

	return false
}

// SetOpenPay gets a reference to the given int32 and assigns it to the OpenPay field.
func (o *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) SetOpenPay(v int32) {
	o.OpenPay = &v
}

// GetOpenCard returns the OpenCard field value if set, zero value otherwise.
func (o *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) GetOpenCard() int32 {
	if o == nil || isNil(o.OpenCard) {
		var ret int32
		return ret
	}
	return *o.OpenCard
}

// GetOpenCardOk returns a tuple with the OpenCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) GetOpenCardOk() (*int32, bool) {
	if o == nil || isNil(o.OpenCard) {
		return nil, false
	}
	return o.OpenCard, true
}

// HasOpenCard returns a boolean if a field has been set.
func (o *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) HasOpenCard() bool {
	if o != nil && !isNil(o.OpenCard) {
		return true
	}

	return false
}

// SetOpenCard gets a reference to the given int32 and assigns it to the OpenCard field.
func (o *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) SetOpenCard(v int32) {
	o.OpenCard = &v
}

// GetOpenShake returns the OpenShake field value if set, zero value otherwise.
func (o *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) GetOpenShake() int32 {
	if o == nil || isNil(o.OpenShake) {
		var ret int32
		return ret
	}
	return *o.OpenShake
}

// GetOpenShakeOk returns a tuple with the OpenShake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) GetOpenShakeOk() (*int32, bool) {
	if o == nil || isNil(o.OpenShake) {
		return nil, false
	}
	return o.OpenShake, true
}

// HasOpenShake returns a boolean if a field has been set.
func (o *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) HasOpenShake() bool {
	if o != nil && !isNil(o.OpenShake) {
		return true
	}

	return false
}

// SetOpenShake gets a reference to the given int32 and assigns it to the OpenShake field.
func (o *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) SetOpenShake(v int32) {
	o.OpenShake = &v
}

func (o GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OpenStore) {
		toSerialize["open_store"] = o.OpenStore
	}
	if !isNil(o.OpenScan) {
		toSerialize["open_scan"] = o.OpenScan
	}
	if !isNil(o.OpenPay) {
		toSerialize["open_pay"] = o.OpenPay
	}
	if !isNil(o.OpenCard) {
		toSerialize["open_card"] = o.OpenCard
	}
	if !isNil(o.OpenShake) {
		toSerialize["open_shake"] = o.OpenShake
	}
	return toSerialize, nil
}

type NullableGetAuthorizerInfoResponseAuthorizerInfoBusinessInfo struct {
	value *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo
	isSet bool
}

func (v NullableGetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) Get() *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo {
	return v.value
}

func (v *NullableGetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) Set(val *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAuthorizerInfoResponseAuthorizerInfoBusinessInfo(val *GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) *NullableGetAuthorizerInfoResponseAuthorizerInfoBusinessInfo {
	return &NullableGetAuthorizerInfoResponseAuthorizerInfoBusinessInfo{value: val, isSet: true}
}

func (v NullableGetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAuthorizerInfoResponseAuthorizerInfoBusinessInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}