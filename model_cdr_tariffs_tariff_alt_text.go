/*
OCPI CDRs module

Specification for OCPIs CDRs handlers

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package OCPI

import (
	"encoding/json"
)

// checks if the CdrTariffsTariffAltText type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CdrTariffsTariffAltText{}

// CdrTariffsTariffAltText struct for CdrTariffsTariffAltText
type CdrTariffsTariffAltText struct {
	Language string `json:"language"`
	Text string `json:"text"`
}

// NewCdrTariffsTariffAltText instantiates a new CdrTariffsTariffAltText object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdrTariffsTariffAltText(language string, text string) *CdrTariffsTariffAltText {
	this := CdrTariffsTariffAltText{}
	this.Language = language
	this.Text = text
	return &this
}

// NewCdrTariffsTariffAltTextWithDefaults instantiates a new CdrTariffsTariffAltText object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdrTariffsTariffAltTextWithDefaults() *CdrTariffsTariffAltText {
	this := CdrTariffsTariffAltText{}
	return &this
}

// GetLanguage returns the Language field value
func (o *CdrTariffsTariffAltText) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *CdrTariffsTariffAltText) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *CdrTariffsTariffAltText) SetLanguage(v string) {
	o.Language = v
}

// GetText returns the Text field value
func (o *CdrTariffsTariffAltText) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *CdrTariffsTariffAltText) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Text, true
}

// SetText sets field value
func (o *CdrTariffsTariffAltText) SetText(v string) {
	o.Text = v
}

func (o CdrTariffsTariffAltText) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CdrTariffsTariffAltText) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["language"] = o.Language
	toSerialize["text"] = o.Text
	return toSerialize, nil
}

type NullableCdrTariffsTariffAltText struct {
	value *CdrTariffsTariffAltText
	isSet bool
}

func (v NullableCdrTariffsTariffAltText) Get() *CdrTariffsTariffAltText {
	return v.value
}

func (v *NullableCdrTariffsTariffAltText) Set(val *CdrTariffsTariffAltText) {
	v.value = val
	v.isSet = true
}

func (v NullableCdrTariffsTariffAltText) IsSet() bool {
	return v.isSet
}

func (v *NullableCdrTariffsTariffAltText) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdrTariffsTariffAltText(val *CdrTariffsTariffAltText) *NullableCdrTariffsTariffAltText {
	return &NullableCdrTariffsTariffAltText{value: val, isSet: true}
}

func (v NullableCdrTariffsTariffAltText) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdrTariffsTariffAltText) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


