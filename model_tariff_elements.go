/*
OCPI tariffs module

Specification for OCPIs tariffs handlers

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package OCPI

import (
	"encoding/json"
)

// checks if the TariffElements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TariffElements{}

// TariffElements struct for TariffElements
type TariffElements struct {
	PriceComponents *string `json:"price_components,omitempty"`
	Restrictions *TariffElementsRestrictions `json:"restrictions,omitempty"`
}

// NewTariffElements instantiates a new TariffElements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTariffElements() *TariffElements {
	this := TariffElements{}
	return &this
}

// NewTariffElementsWithDefaults instantiates a new TariffElements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTariffElementsWithDefaults() *TariffElements {
	this := TariffElements{}
	return &this
}

// GetPriceComponents returns the PriceComponents field value if set, zero value otherwise.
func (o *TariffElements) GetPriceComponents() string {
	if o == nil || IsNil(o.PriceComponents) {
		var ret string
		return ret
	}
	return *o.PriceComponents
}

// GetPriceComponentsOk returns a tuple with the PriceComponents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TariffElements) GetPriceComponentsOk() (*string, bool) {
	if o == nil || IsNil(o.PriceComponents) {
		return nil, false
	}
	return o.PriceComponents, true
}

// HasPriceComponents returns a boolean if a field has been set.
func (o *TariffElements) HasPriceComponents() bool {
	if o != nil && !IsNil(o.PriceComponents) {
		return true
	}

	return false
}

// SetPriceComponents gets a reference to the given string and assigns it to the PriceComponents field.
func (o *TariffElements) SetPriceComponents(v string) {
	o.PriceComponents = &v
}

// GetRestrictions returns the Restrictions field value if set, zero value otherwise.
func (o *TariffElements) GetRestrictions() TariffElementsRestrictions {
	if o == nil || IsNil(o.Restrictions) {
		var ret TariffElementsRestrictions
		return ret
	}
	return *o.Restrictions
}

// GetRestrictionsOk returns a tuple with the Restrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TariffElements) GetRestrictionsOk() (*TariffElementsRestrictions, bool) {
	if o == nil || IsNil(o.Restrictions) {
		return nil, false
	}
	return o.Restrictions, true
}

// HasRestrictions returns a boolean if a field has been set.
func (o *TariffElements) HasRestrictions() bool {
	if o != nil && !IsNil(o.Restrictions) {
		return true
	}

	return false
}

// SetRestrictions gets a reference to the given TariffElementsRestrictions and assigns it to the Restrictions field.
func (o *TariffElements) SetRestrictions(v TariffElementsRestrictions) {
	o.Restrictions = &v
}

func (o TariffElements) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TariffElements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PriceComponents) {
		toSerialize["price_components"] = o.PriceComponents
	}
	if !IsNil(o.Restrictions) {
		toSerialize["restrictions"] = o.Restrictions
	}
	return toSerialize, nil
}

type NullableTariffElements struct {
	value *TariffElements
	isSet bool
}

func (v NullableTariffElements) Get() *TariffElements {
	return v.value
}

func (v *NullableTariffElements) Set(val *TariffElements) {
	v.value = val
	v.isSet = true
}

func (v NullableTariffElements) IsSet() bool {
	return v.isSet
}

func (v *NullableTariffElements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTariffElements(val *TariffElements) *NullableTariffElements {
	return &NullableTariffElements{value: val, isSet: true}
}

func (v NullableTariffElements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTariffElements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


