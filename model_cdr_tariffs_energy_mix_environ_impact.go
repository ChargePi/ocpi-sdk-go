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

// checks if the CdrTariffsEnergyMixEnvironImpact type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CdrTariffsEnergyMixEnvironImpact{}

// CdrTariffsEnergyMixEnvironImpact struct for CdrTariffsEnergyMixEnvironImpact
type CdrTariffsEnergyMixEnvironImpact struct {
	Category string `json:"category"`
	Amount float32 `json:"amount"`
}

// NewCdrTariffsEnergyMixEnvironImpact instantiates a new CdrTariffsEnergyMixEnvironImpact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdrTariffsEnergyMixEnvironImpact(category string, amount float32) *CdrTariffsEnergyMixEnvironImpact {
	this := CdrTariffsEnergyMixEnvironImpact{}
	this.Category = category
	this.Amount = amount
	return &this
}

// NewCdrTariffsEnergyMixEnvironImpactWithDefaults instantiates a new CdrTariffsEnergyMixEnvironImpact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdrTariffsEnergyMixEnvironImpactWithDefaults() *CdrTariffsEnergyMixEnvironImpact {
	this := CdrTariffsEnergyMixEnvironImpact{}
	return &this
}

// GetCategory returns the Category field value
func (o *CdrTariffsEnergyMixEnvironImpact) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *CdrTariffsEnergyMixEnvironImpact) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *CdrTariffsEnergyMixEnvironImpact) SetCategory(v string) {
	o.Category = v
}

// GetAmount returns the Amount field value
func (o *CdrTariffsEnergyMixEnvironImpact) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CdrTariffsEnergyMixEnvironImpact) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CdrTariffsEnergyMixEnvironImpact) SetAmount(v float32) {
	o.Amount = v
}

func (o CdrTariffsEnergyMixEnvironImpact) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CdrTariffsEnergyMixEnvironImpact) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["category"] = o.Category
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

type NullableCdrTariffsEnergyMixEnvironImpact struct {
	value *CdrTariffsEnergyMixEnvironImpact
	isSet bool
}

func (v NullableCdrTariffsEnergyMixEnvironImpact) Get() *CdrTariffsEnergyMixEnvironImpact {
	return v.value
}

func (v *NullableCdrTariffsEnergyMixEnvironImpact) Set(val *CdrTariffsEnergyMixEnvironImpact) {
	v.value = val
	v.isSet = true
}

func (v NullableCdrTariffsEnergyMixEnvironImpact) IsSet() bool {
	return v.isSet
}

func (v *NullableCdrTariffsEnergyMixEnvironImpact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdrTariffsEnergyMixEnvironImpact(val *CdrTariffsEnergyMixEnvironImpact) *NullableCdrTariffsEnergyMixEnvironImpact {
	return &NullableCdrTariffsEnergyMixEnvironImpact{value: val, isSet: true}
}

func (v NullableCdrTariffsEnergyMixEnvironImpact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdrTariffsEnergyMixEnvironImpact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


