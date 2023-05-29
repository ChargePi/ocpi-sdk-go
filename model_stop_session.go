/*
OCPI commands module

Specification for OCPIs commands handlers

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package OCPI

import (
	"encoding/json"
)

// checks if the StopSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StopSession{}

// StopSession struct for StopSession
type StopSession struct {
	ResponseUrl string `json:"response_url"`
	SessionId *string `json:"session_id,omitempty"`
}

// NewStopSession instantiates a new StopSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStopSession(responseUrl string) *StopSession {
	this := StopSession{}
	this.ResponseUrl = responseUrl
	return &this
}

// NewStopSessionWithDefaults instantiates a new StopSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStopSessionWithDefaults() *StopSession {
	this := StopSession{}
	return &this
}

// GetResponseUrl returns the ResponseUrl field value
func (o *StopSession) GetResponseUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResponseUrl
}

// GetResponseUrlOk returns a tuple with the ResponseUrl field value
// and a boolean to check if the value has been set.
func (o *StopSession) GetResponseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResponseUrl, true
}

// SetResponseUrl sets field value
func (o *StopSession) SetResponseUrl(v string) {
	o.ResponseUrl = v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *StopSession) GetSessionId() string {
	if o == nil || IsNil(o.SessionId) {
		var ret string
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopSession) GetSessionIdOk() (*string, bool) {
	if o == nil || IsNil(o.SessionId) {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *StopSession) HasSessionId() bool {
	if o != nil && !IsNil(o.SessionId) {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given string and assigns it to the SessionId field.
func (o *StopSession) SetSessionId(v string) {
	o.SessionId = &v
}

func (o StopSession) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StopSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["response_url"] = o.ResponseUrl
	if !IsNil(o.SessionId) {
		toSerialize["session_id"] = o.SessionId
	}
	return toSerialize, nil
}

type NullableStopSession struct {
	value *StopSession
	isSet bool
}

func (v NullableStopSession) Get() *StopSession {
	return v.value
}

func (v *NullableStopSession) Set(val *StopSession) {
	v.value = val
	v.isSet = true
}

func (v NullableStopSession) IsSet() bool {
	return v.isSet
}

func (v *NullableStopSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStopSession(val *StopSession) *NullableStopSession {
	return &NullableStopSession{value: val, isSet: true}
}

func (v NullableStopSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStopSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


