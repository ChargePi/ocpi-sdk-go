/*
OCPI tokens module

Specification for OCPIs tokens handlers

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package OCPI

import (
	"encoding/json"
)

// checks if the TokenEnergyContract type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenEnergyContract{}

// TokenEnergyContract struct for TokenEnergyContract
type TokenEnergyContract struct {
	SupplierName string `json:"supplier_name"`
	ContractId *string `json:"contract_id,omitempty"`
}

// NewTokenEnergyContract instantiates a new TokenEnergyContract object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenEnergyContract(supplierName string) *TokenEnergyContract {
	this := TokenEnergyContract{}
	this.SupplierName = supplierName
	return &this
}

// NewTokenEnergyContractWithDefaults instantiates a new TokenEnergyContract object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenEnergyContractWithDefaults() *TokenEnergyContract {
	this := TokenEnergyContract{}
	return &this
}

// GetSupplierName returns the SupplierName field value
func (o *TokenEnergyContract) GetSupplierName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupplierName
}

// GetSupplierNameOk returns a tuple with the SupplierName field value
// and a boolean to check if the value has been set.
func (o *TokenEnergyContract) GetSupplierNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupplierName, true
}

// SetSupplierName sets field value
func (o *TokenEnergyContract) SetSupplierName(v string) {
	o.SupplierName = v
}

// GetContractId returns the ContractId field value if set, zero value otherwise.
func (o *TokenEnergyContract) GetContractId() string {
	if o == nil || IsNil(o.ContractId) {
		var ret string
		return ret
	}
	return *o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenEnergyContract) GetContractIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContractId) {
		return nil, false
	}
	return o.ContractId, true
}

// HasContractId returns a boolean if a field has been set.
func (o *TokenEnergyContract) HasContractId() bool {
	if o != nil && !IsNil(o.ContractId) {
		return true
	}

	return false
}

// SetContractId gets a reference to the given string and assigns it to the ContractId field.
func (o *TokenEnergyContract) SetContractId(v string) {
	o.ContractId = &v
}

func (o TokenEnergyContract) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenEnergyContract) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["supplier_name"] = o.SupplierName
	if !IsNil(o.ContractId) {
		toSerialize["contract_id"] = o.ContractId
	}
	return toSerialize, nil
}

type NullableTokenEnergyContract struct {
	value *TokenEnergyContract
	isSet bool
}

func (v NullableTokenEnergyContract) Get() *TokenEnergyContract {
	return v.value
}

func (v *NullableTokenEnergyContract) Set(val *TokenEnergyContract) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenEnergyContract) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenEnergyContract) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenEnergyContract(val *TokenEnergyContract) *NullableTokenEnergyContract {
	return &NullableTokenEnergyContract{value: val, isSet: true}
}

func (v NullableTokenEnergyContract) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenEnergyContract) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


