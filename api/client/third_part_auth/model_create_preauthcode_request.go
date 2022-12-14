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

// checks if the CreatePreauthcodeRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatePreauthcodeRequest{}

// CreatePreauthcodeRequest struct for CreatePreauthcodeRequest
type CreatePreauthcodeRequest struct {
	// 平台型第三方平台的appid
	ComponentAppid string `json:"component_appid"`
}

// NewCreatePreauthcodeRequest instantiates a new CreatePreauthcodeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatePreauthcodeRequest(componentAppid string) *CreatePreauthcodeRequest {
	this := CreatePreauthcodeRequest{}
	this.ComponentAppid = componentAppid
	return &this
}

// NewCreatePreauthcodeRequestWithDefaults instantiates a new CreatePreauthcodeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatePreauthcodeRequestWithDefaults() *CreatePreauthcodeRequest {
	this := CreatePreauthcodeRequest{}
	return &this
}

// GetComponentAppid returns the ComponentAppid field value
func (o *CreatePreauthcodeRequest) GetComponentAppid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComponentAppid
}

// GetComponentAppidOk returns a tuple with the ComponentAppid field value
// and a boolean to check if the value has been set.
func (o *CreatePreauthcodeRequest) GetComponentAppidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComponentAppid, true
}

// SetComponentAppid sets field value
func (o *CreatePreauthcodeRequest) SetComponentAppid(v string) {
	o.ComponentAppid = v
}

func (o CreatePreauthcodeRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatePreauthcodeRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["component_appid"] = o.ComponentAppid
	return toSerialize, nil
}

type NullableCreatePreauthcodeRequest struct {
	value *CreatePreauthcodeRequest
	isSet bool
}

func (v NullableCreatePreauthcodeRequest) Get() *CreatePreauthcodeRequest {
	return v.value
}

func (v *NullableCreatePreauthcodeRequest) Set(val *CreatePreauthcodeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatePreauthcodeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatePreauthcodeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatePreauthcodeRequest(val *CreatePreauthcodeRequest) *NullableCreatePreauthcodeRequest {
	return &NullableCreatePreauthcodeRequest{value: val, isSet: true}
}

func (v NullableCreatePreauthcodeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatePreauthcodeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
