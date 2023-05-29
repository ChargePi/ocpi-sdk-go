/*
OCPI credentials module

Specification for OCPIs credentials handlers

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package OCPI

import (
	"encoding/json"
)

// checks if the CredentialsDataRoles type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialsDataRoles{}

// CredentialsDataRoles struct for CredentialsDataRoles
type CredentialsDataRoles struct {
	Role string `json:"role"`
	BusinessDetails CredentialsDataRolesBusinessDetails `json:"business_details"`
	PartyId string `json:"party_id"`
	CountryCode string `json:"country_code"`
}

// NewCredentialsDataRoles instantiates a new CredentialsDataRoles object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialsDataRoles(role string, businessDetails CredentialsDataRolesBusinessDetails, partyId string, countryCode string) *CredentialsDataRoles {
	this := CredentialsDataRoles{}
	this.Role = role
	this.BusinessDetails = businessDetails
	this.PartyId = partyId
	this.CountryCode = countryCode
	return &this
}

// NewCredentialsDataRolesWithDefaults instantiates a new CredentialsDataRoles object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialsDataRolesWithDefaults() *CredentialsDataRoles {
	this := CredentialsDataRoles{}
	return &this
}

// GetRole returns the Role field value
func (o *CredentialsDataRoles) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *CredentialsDataRoles) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *CredentialsDataRoles) SetRole(v string) {
	o.Role = v
}

// GetBusinessDetails returns the BusinessDetails field value
func (o *CredentialsDataRoles) GetBusinessDetails() CredentialsDataRolesBusinessDetails {
	if o == nil {
		var ret CredentialsDataRolesBusinessDetails
		return ret
	}

	return o.BusinessDetails
}

// GetBusinessDetailsOk returns a tuple with the BusinessDetails field value
// and a boolean to check if the value has been set.
func (o *CredentialsDataRoles) GetBusinessDetailsOk() (*CredentialsDataRolesBusinessDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessDetails, true
}

// SetBusinessDetails sets field value
func (o *CredentialsDataRoles) SetBusinessDetails(v CredentialsDataRolesBusinessDetails) {
	o.BusinessDetails = v
}

// GetPartyId returns the PartyId field value
func (o *CredentialsDataRoles) GetPartyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PartyId
}

// GetPartyIdOk returns a tuple with the PartyId field value
// and a boolean to check if the value has been set.
func (o *CredentialsDataRoles) GetPartyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartyId, true
}

// SetPartyId sets field value
func (o *CredentialsDataRoles) SetPartyId(v string) {
	o.PartyId = v
}

// GetCountryCode returns the CountryCode field value
func (o *CredentialsDataRoles) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *CredentialsDataRoles) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *CredentialsDataRoles) SetCountryCode(v string) {
	o.CountryCode = v
}

func (o CredentialsDataRoles) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialsDataRoles) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["role"] = o.Role
	toSerialize["business_details"] = o.BusinessDetails
	toSerialize["party_id"] = o.PartyId
	toSerialize["country_code"] = o.CountryCode
	return toSerialize, nil
}

type NullableCredentialsDataRoles struct {
	value *CredentialsDataRoles
	isSet bool
}

func (v NullableCredentialsDataRoles) Get() *CredentialsDataRoles {
	return v.value
}

func (v *NullableCredentialsDataRoles) Set(val *CredentialsDataRoles) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialsDataRoles) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialsDataRoles) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialsDataRoles(val *CredentialsDataRoles) *NullableCredentialsDataRoles {
	return &NullableCredentialsDataRoles{value: val, isSet: true}
}

func (v NullableCredentialsDataRoles) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialsDataRoles) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


