/*
OCPI sessions module

Specification for OCPIs sessions handlers

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package OCPI

import (
	"encoding/json"
)

// checks if the SessionCdrToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionCdrToken{}

// SessionCdrToken struct for SessionCdrToken
type SessionCdrToken struct {
	CountryCode string `json:"country_code"`
	PartyId string `json:"party_id"`
	Uid string `json:"uid"`
	Type string `json:"type"`
	ContractId string `json:"contract_id"`
}

// NewSessionCdrToken instantiates a new SessionCdrToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionCdrToken(countryCode string, partyId string, uid string, type_ string, contractId string) *SessionCdrToken {
	this := SessionCdrToken{}
	this.CountryCode = countryCode
	this.PartyId = partyId
	this.Uid = uid
	this.Type = type_
	this.ContractId = contractId
	return &this
}

// NewSessionCdrTokenWithDefaults instantiates a new SessionCdrToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionCdrTokenWithDefaults() *SessionCdrToken {
	this := SessionCdrToken{}
	return &this
}

// GetCountryCode returns the CountryCode field value
func (o *SessionCdrToken) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *SessionCdrToken) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *SessionCdrToken) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetPartyId returns the PartyId field value
func (o *SessionCdrToken) GetPartyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PartyId
}

// GetPartyIdOk returns a tuple with the PartyId field value
// and a boolean to check if the value has been set.
func (o *SessionCdrToken) GetPartyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartyId, true
}

// SetPartyId sets field value
func (o *SessionCdrToken) SetPartyId(v string) {
	o.PartyId = v
}

// GetUid returns the Uid field value
func (o *SessionCdrToken) GetUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uid
}

// GetUidOk returns a tuple with the Uid field value
// and a boolean to check if the value has been set.
func (o *SessionCdrToken) GetUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uid, true
}

// SetUid sets field value
func (o *SessionCdrToken) SetUid(v string) {
	o.Uid = v
}

// GetType returns the Type field value
func (o *SessionCdrToken) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SessionCdrToken) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SessionCdrToken) SetType(v string) {
	o.Type = v
}

// GetContractId returns the ContractId field value
func (o *SessionCdrToken) GetContractId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value
// and a boolean to check if the value has been set.
func (o *SessionCdrToken) GetContractIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractId, true
}

// SetContractId sets field value
func (o *SessionCdrToken) SetContractId(v string) {
	o.ContractId = v
}

func (o SessionCdrToken) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionCdrToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["country_code"] = o.CountryCode
	toSerialize["party_id"] = o.PartyId
	toSerialize["uid"] = o.Uid
	toSerialize["type"] = o.Type
	toSerialize["contract_id"] = o.ContractId
	return toSerialize, nil
}

type NullableSessionCdrToken struct {
	value *SessionCdrToken
	isSet bool
}

func (v NullableSessionCdrToken) Get() *SessionCdrToken {
	return v.value
}

func (v *NullableSessionCdrToken) Set(val *SessionCdrToken) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionCdrToken) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionCdrToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionCdrToken(val *SessionCdrToken) *NullableSessionCdrToken {
	return &NullableSessionCdrToken{value: val, isSet: true}
}

func (v NullableSessionCdrToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionCdrToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


