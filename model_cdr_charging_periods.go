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

// checks if the CdrChargingPeriods type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CdrChargingPeriods{}

// CdrChargingPeriods struct for CdrChargingPeriods
type CdrChargingPeriods struct {
	StartDateTime string `json:"start_date_time"`
	Dimensions *CdrChargingPeriodsDimensions `json:"dimensions,omitempty"`
	TariffId *string `json:"tariff_id,omitempty"`
}

// NewCdrChargingPeriods instantiates a new CdrChargingPeriods object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdrChargingPeriods(startDateTime string) *CdrChargingPeriods {
	this := CdrChargingPeriods{}
	this.StartDateTime = startDateTime
	return &this
}

// NewCdrChargingPeriodsWithDefaults instantiates a new CdrChargingPeriods object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdrChargingPeriodsWithDefaults() *CdrChargingPeriods {
	this := CdrChargingPeriods{}
	return &this
}

// GetStartDateTime returns the StartDateTime field value
func (o *CdrChargingPeriods) GetStartDateTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value
// and a boolean to check if the value has been set.
func (o *CdrChargingPeriods) GetStartDateTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDateTime, true
}

// SetStartDateTime sets field value
func (o *CdrChargingPeriods) SetStartDateTime(v string) {
	o.StartDateTime = v
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *CdrChargingPeriods) GetDimensions() CdrChargingPeriodsDimensions {
	if o == nil || IsNil(o.Dimensions) {
		var ret CdrChargingPeriodsDimensions
		return ret
	}
	return *o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrChargingPeriods) GetDimensionsOk() (*CdrChargingPeriodsDimensions, bool) {
	if o == nil || IsNil(o.Dimensions) {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *CdrChargingPeriods) HasDimensions() bool {
	if o != nil && !IsNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given CdrChargingPeriodsDimensions and assigns it to the Dimensions field.
func (o *CdrChargingPeriods) SetDimensions(v CdrChargingPeriodsDimensions) {
	o.Dimensions = &v
}

// GetTariffId returns the TariffId field value if set, zero value otherwise.
func (o *CdrChargingPeriods) GetTariffId() string {
	if o == nil || IsNil(o.TariffId) {
		var ret string
		return ret
	}
	return *o.TariffId
}

// GetTariffIdOk returns a tuple with the TariffId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrChargingPeriods) GetTariffIdOk() (*string, bool) {
	if o == nil || IsNil(o.TariffId) {
		return nil, false
	}
	return o.TariffId, true
}

// HasTariffId returns a boolean if a field has been set.
func (o *CdrChargingPeriods) HasTariffId() bool {
	if o != nil && !IsNil(o.TariffId) {
		return true
	}

	return false
}

// SetTariffId gets a reference to the given string and assigns it to the TariffId field.
func (o *CdrChargingPeriods) SetTariffId(v string) {
	o.TariffId = &v
}

func (o CdrChargingPeriods) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CdrChargingPeriods) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start_date_time"] = o.StartDateTime
	if !IsNil(o.Dimensions) {
		toSerialize["dimensions"] = o.Dimensions
	}
	if !IsNil(o.TariffId) {
		toSerialize["tariff_id"] = o.TariffId
	}
	return toSerialize, nil
}

type NullableCdrChargingPeriods struct {
	value *CdrChargingPeriods
	isSet bool
}

func (v NullableCdrChargingPeriods) Get() *CdrChargingPeriods {
	return v.value
}

func (v *NullableCdrChargingPeriods) Set(val *CdrChargingPeriods) {
	v.value = val
	v.isSet = true
}

func (v NullableCdrChargingPeriods) IsSet() bool {
	return v.isSet
}

func (v *NullableCdrChargingPeriods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdrChargingPeriods(val *CdrChargingPeriods) *NullableCdrChargingPeriods {
	return &NullableCdrChargingPeriods{value: val, isSet: true}
}

func (v NullableCdrChargingPeriods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdrChargingPeriods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


