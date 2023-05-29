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

// checks if the CdrSignedData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CdrSignedData{}

// CdrSignedData struct for CdrSignedData
type CdrSignedData struct {
	EncodingMethod string `json:"encoding_method"`
	EncodingMethodVersion *int32 `json:"encoding_method_version,omitempty"`
	PublicKey *string `json:"public_key,omitempty"`
	SignedValues *CdrSignedDataSignedValues `json:"signed_values,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewCdrSignedData instantiates a new CdrSignedData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdrSignedData(encodingMethod string) *CdrSignedData {
	this := CdrSignedData{}
	this.EncodingMethod = encodingMethod
	return &this
}

// NewCdrSignedDataWithDefaults instantiates a new CdrSignedData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdrSignedDataWithDefaults() *CdrSignedData {
	this := CdrSignedData{}
	return &this
}

// GetEncodingMethod returns the EncodingMethod field value
func (o *CdrSignedData) GetEncodingMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EncodingMethod
}

// GetEncodingMethodOk returns a tuple with the EncodingMethod field value
// and a boolean to check if the value has been set.
func (o *CdrSignedData) GetEncodingMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EncodingMethod, true
}

// SetEncodingMethod sets field value
func (o *CdrSignedData) SetEncodingMethod(v string) {
	o.EncodingMethod = v
}

// GetEncodingMethodVersion returns the EncodingMethodVersion field value if set, zero value otherwise.
func (o *CdrSignedData) GetEncodingMethodVersion() int32 {
	if o == nil || IsNil(o.EncodingMethodVersion) {
		var ret int32
		return ret
	}
	return *o.EncodingMethodVersion
}

// GetEncodingMethodVersionOk returns a tuple with the EncodingMethodVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrSignedData) GetEncodingMethodVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.EncodingMethodVersion) {
		return nil, false
	}
	return o.EncodingMethodVersion, true
}

// HasEncodingMethodVersion returns a boolean if a field has been set.
func (o *CdrSignedData) HasEncodingMethodVersion() bool {
	if o != nil && !IsNil(o.EncodingMethodVersion) {
		return true
	}

	return false
}

// SetEncodingMethodVersion gets a reference to the given int32 and assigns it to the EncodingMethodVersion field.
func (o *CdrSignedData) SetEncodingMethodVersion(v int32) {
	o.EncodingMethodVersion = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *CdrSignedData) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey) {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrSignedData) GetPublicKeyOk() (*string, bool) {
	if o == nil || IsNil(o.PublicKey) {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *CdrSignedData) HasPublicKey() bool {
	if o != nil && !IsNil(o.PublicKey) {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *CdrSignedData) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetSignedValues returns the SignedValues field value if set, zero value otherwise.
func (o *CdrSignedData) GetSignedValues() CdrSignedDataSignedValues {
	if o == nil || IsNil(o.SignedValues) {
		var ret CdrSignedDataSignedValues
		return ret
	}
	return *o.SignedValues
}

// GetSignedValuesOk returns a tuple with the SignedValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrSignedData) GetSignedValuesOk() (*CdrSignedDataSignedValues, bool) {
	if o == nil || IsNil(o.SignedValues) {
		return nil, false
	}
	return o.SignedValues, true
}

// HasSignedValues returns a boolean if a field has been set.
func (o *CdrSignedData) HasSignedValues() bool {
	if o != nil && !IsNil(o.SignedValues) {
		return true
	}

	return false
}

// SetSignedValues gets a reference to the given CdrSignedDataSignedValues and assigns it to the SignedValues field.
func (o *CdrSignedData) SetSignedValues(v CdrSignedDataSignedValues) {
	o.SignedValues = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CdrSignedData) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrSignedData) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CdrSignedData) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CdrSignedData) SetUrl(v string) {
	o.Url = &v
}

func (o CdrSignedData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CdrSignedData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["encoding_method"] = o.EncodingMethod
	if !IsNil(o.EncodingMethodVersion) {
		toSerialize["encoding_method_version"] = o.EncodingMethodVersion
	}
	if !IsNil(o.PublicKey) {
		toSerialize["public_key"] = o.PublicKey
	}
	if !IsNil(o.SignedValues) {
		toSerialize["signed_values"] = o.SignedValues
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableCdrSignedData struct {
	value *CdrSignedData
	isSet bool
}

func (v NullableCdrSignedData) Get() *CdrSignedData {
	return v.value
}

func (v *NullableCdrSignedData) Set(val *CdrSignedData) {
	v.value = val
	v.isSet = true
}

func (v NullableCdrSignedData) IsSet() bool {
	return v.isSet
}

func (v *NullableCdrSignedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdrSignedData(val *CdrSignedData) *NullableCdrSignedData {
	return &NullableCdrSignedData{value: val, isSet: true}
}

func (v NullableCdrSignedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdrSignedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


