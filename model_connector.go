/*
OCPI locations module

Specification for OCPIs locations handlers

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package OCPI

import (
	"encoding/json"
)

// checks if the Connector type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Connector{}

// Connector struct for Connector
type Connector struct {
	Id string `json:"id"`
	Standard string `json:"standard"`
	Format string `json:"format"`
	PowerType string `json:"power_type"`
	MaxVoltage int32 `json:"max_voltage"`
	MaxAmperage int32 `json:"max_amperage"`
	MaxElectricPower *int32 `json:"max_electric_power,omitempty"`
	TariffIds *string `json:"tariff_ids,omitempty"`
	TermsAndConditions *string `json:"terms_and_conditions,omitempty"`
	LastUpdated string `json:"last_updated"`
}

// NewConnector instantiates a new Connector object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnector(id string, standard string, format string, powerType string, maxVoltage int32, maxAmperage int32, lastUpdated string) *Connector {
	this := Connector{}
	this.Id = id
	this.Standard = standard
	this.Format = format
	this.PowerType = powerType
	this.MaxVoltage = maxVoltage
	this.MaxAmperage = maxAmperage
	this.LastUpdated = lastUpdated
	return &this
}

// NewConnectorWithDefaults instantiates a new Connector object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorWithDefaults() *Connector {
	this := Connector{}
	return &this
}

// GetId returns the Id field value
func (o *Connector) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Connector) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Connector) SetId(v string) {
	o.Id = v
}

// GetStandard returns the Standard field value
func (o *Connector) GetStandard() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Standard
}

// GetStandardOk returns a tuple with the Standard field value
// and a boolean to check if the value has been set.
func (o *Connector) GetStandardOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Standard, true
}

// SetStandard sets field value
func (o *Connector) SetStandard(v string) {
	o.Standard = v
}

// GetFormat returns the Format field value
func (o *Connector) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *Connector) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *Connector) SetFormat(v string) {
	o.Format = v
}

// GetPowerType returns the PowerType field value
func (o *Connector) GetPowerType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PowerType
}

// GetPowerTypeOk returns a tuple with the PowerType field value
// and a boolean to check if the value has been set.
func (o *Connector) GetPowerTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PowerType, true
}

// SetPowerType sets field value
func (o *Connector) SetPowerType(v string) {
	o.PowerType = v
}

// GetMaxVoltage returns the MaxVoltage field value
func (o *Connector) GetMaxVoltage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxVoltage
}

// GetMaxVoltageOk returns a tuple with the MaxVoltage field value
// and a boolean to check if the value has been set.
func (o *Connector) GetMaxVoltageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxVoltage, true
}

// SetMaxVoltage sets field value
func (o *Connector) SetMaxVoltage(v int32) {
	o.MaxVoltage = v
}

// GetMaxAmperage returns the MaxAmperage field value
func (o *Connector) GetMaxAmperage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxAmperage
}

// GetMaxAmperageOk returns a tuple with the MaxAmperage field value
// and a boolean to check if the value has been set.
func (o *Connector) GetMaxAmperageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxAmperage, true
}

// SetMaxAmperage sets field value
func (o *Connector) SetMaxAmperage(v int32) {
	o.MaxAmperage = v
}

// GetMaxElectricPower returns the MaxElectricPower field value if set, zero value otherwise.
func (o *Connector) GetMaxElectricPower() int32 {
	if o == nil || IsNil(o.MaxElectricPower) {
		var ret int32
		return ret
	}
	return *o.MaxElectricPower
}

// GetMaxElectricPowerOk returns a tuple with the MaxElectricPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connector) GetMaxElectricPowerOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxElectricPower) {
		return nil, false
	}
	return o.MaxElectricPower, true
}

// HasMaxElectricPower returns a boolean if a field has been set.
func (o *Connector) HasMaxElectricPower() bool {
	if o != nil && !IsNil(o.MaxElectricPower) {
		return true
	}

	return false
}

// SetMaxElectricPower gets a reference to the given int32 and assigns it to the MaxElectricPower field.
func (o *Connector) SetMaxElectricPower(v int32) {
	o.MaxElectricPower = &v
}

// GetTariffIds returns the TariffIds field value if set, zero value otherwise.
func (o *Connector) GetTariffIds() string {
	if o == nil || IsNil(o.TariffIds) {
		var ret string
		return ret
	}
	return *o.TariffIds
}

// GetTariffIdsOk returns a tuple with the TariffIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connector) GetTariffIdsOk() (*string, bool) {
	if o == nil || IsNil(o.TariffIds) {
		return nil, false
	}
	return o.TariffIds, true
}

// HasTariffIds returns a boolean if a field has been set.
func (o *Connector) HasTariffIds() bool {
	if o != nil && !IsNil(o.TariffIds) {
		return true
	}

	return false
}

// SetTariffIds gets a reference to the given string and assigns it to the TariffIds field.
func (o *Connector) SetTariffIds(v string) {
	o.TariffIds = &v
}

// GetTermsAndConditions returns the TermsAndConditions field value if set, zero value otherwise.
func (o *Connector) GetTermsAndConditions() string {
	if o == nil || IsNil(o.TermsAndConditions) {
		var ret string
		return ret
	}
	return *o.TermsAndConditions
}

// GetTermsAndConditionsOk returns a tuple with the TermsAndConditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connector) GetTermsAndConditionsOk() (*string, bool) {
	if o == nil || IsNil(o.TermsAndConditions) {
		return nil, false
	}
	return o.TermsAndConditions, true
}

// HasTermsAndConditions returns a boolean if a field has been set.
func (o *Connector) HasTermsAndConditions() bool {
	if o != nil && !IsNil(o.TermsAndConditions) {
		return true
	}

	return false
}

// SetTermsAndConditions gets a reference to the given string and assigns it to the TermsAndConditions field.
func (o *Connector) SetTermsAndConditions(v string) {
	o.TermsAndConditions = &v
}

// GetLastUpdated returns the LastUpdated field value
func (o *Connector) GetLastUpdated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *Connector) GetLastUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *Connector) SetLastUpdated(v string) {
	o.LastUpdated = v
}

func (o Connector) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Connector) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["standard"] = o.Standard
	toSerialize["format"] = o.Format
	toSerialize["power_type"] = o.PowerType
	toSerialize["max_voltage"] = o.MaxVoltage
	toSerialize["max_amperage"] = o.MaxAmperage
	if !IsNil(o.MaxElectricPower) {
		toSerialize["max_electric_power"] = o.MaxElectricPower
	}
	if !IsNil(o.TariffIds) {
		toSerialize["tariff_ids"] = o.TariffIds
	}
	if !IsNil(o.TermsAndConditions) {
		toSerialize["terms_and_conditions"] = o.TermsAndConditions
	}
	toSerialize["last_updated"] = o.LastUpdated
	return toSerialize, nil
}

type NullableConnector struct {
	value *Connector
	isSet bool
}

func (v NullableConnector) Get() *Connector {
	return v.value
}

func (v *NullableConnector) Set(val *Connector) {
	v.value = val
	v.isSet = true
}

func (v NullableConnector) IsSet() bool {
	return v.isSet
}

func (v *NullableConnector) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnector(val *Connector) *NullableConnector {
	return &NullableConnector{value: val, isSet: true}
}

func (v NullableConnector) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnector) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


