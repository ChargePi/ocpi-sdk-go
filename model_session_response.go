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

// checks if the SessionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionResponse{}

// SessionResponse struct for SessionResponse
type SessionResponse struct {
	Session Session `json:"session"`
	StatusCode float32 `json:"status_code"`
	StatusMessage *string `json:"status_message,omitempty"`
	TimeStamp *string `json:"timeStamp,omitempty"`
}

// NewSessionResponse instantiates a new SessionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionResponse(session Session, statusCode float32) *SessionResponse {
	this := SessionResponse{}
	this.Session = session
	this.StatusCode = statusCode
	return &this
}

// NewSessionResponseWithDefaults instantiates a new SessionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionResponseWithDefaults() *SessionResponse {
	this := SessionResponse{}
	return &this
}

// GetSession returns the Session field value
func (o *SessionResponse) GetSession() Session {
	if o == nil {
		var ret Session
		return ret
	}

	return o.Session
}

// GetSessionOk returns a tuple with the Session field value
// and a boolean to check if the value has been set.
func (o *SessionResponse) GetSessionOk() (*Session, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Session, true
}

// SetSession sets field value
func (o *SessionResponse) SetSession(v Session) {
	o.Session = v
}

// GetStatusCode returns the StatusCode field value
func (o *SessionResponse) GetStatusCode() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *SessionResponse) GetStatusCodeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *SessionResponse) SetStatusCode(v float32) {
	o.StatusCode = v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *SessionResponse) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionResponse) GetStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMessage) {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *SessionResponse) HasStatusMessage() bool {
	if o != nil && !IsNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *SessionResponse) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetTimeStamp returns the TimeStamp field value if set, zero value otherwise.
func (o *SessionResponse) GetTimeStamp() string {
	if o == nil || IsNil(o.TimeStamp) {
		var ret string
		return ret
	}
	return *o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionResponse) GetTimeStampOk() (*string, bool) {
	if o == nil || IsNil(o.TimeStamp) {
		return nil, false
	}
	return o.TimeStamp, true
}

// HasTimeStamp returns a boolean if a field has been set.
func (o *SessionResponse) HasTimeStamp() bool {
	if o != nil && !IsNil(o.TimeStamp) {
		return true
	}

	return false
}

// SetTimeStamp gets a reference to the given string and assigns it to the TimeStamp field.
func (o *SessionResponse) SetTimeStamp(v string) {
	o.TimeStamp = &v
}

func (o SessionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["session"] = o.Session
	toSerialize["status_code"] = o.StatusCode
	if !IsNil(o.StatusMessage) {
		toSerialize["status_message"] = o.StatusMessage
	}
	if !IsNil(o.TimeStamp) {
		toSerialize["timeStamp"] = o.TimeStamp
	}
	return toSerialize, nil
}

type NullableSessionResponse struct {
	value *SessionResponse
	isSet bool
}

func (v NullableSessionResponse) Get() *SessionResponse {
	return v.value
}

func (v *NullableSessionResponse) Set(val *SessionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionResponse(val *SessionResponse) *NullableSessionResponse {
	return &NullableSessionResponse{value: val, isSet: true}
}

func (v NullableSessionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


