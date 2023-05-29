/*
OCPI tariffs module

Specification for OCPIs tariffs handlers

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package OCPI

import (
	"encoding/json"
	"fmt"
)

// ReservationRestrictionType the model 'ReservationRestrictionType'
type ReservationRestrictionType string

// List of reservationRestrictionType
const (
	RESERVATION ReservationRestrictionType = "RESERVATION"
	RESERVATION_EXPIRES ReservationRestrictionType = "RESERVATION_EXPIRES"
)

// All allowed values of ReservationRestrictionType enum
var AllowedReservationRestrictionTypeEnumValues = []ReservationRestrictionType{
	"RESERVATION",
	"RESERVATION_EXPIRES",
}

func (v *ReservationRestrictionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReservationRestrictionType(value)
	for _, existing := range AllowedReservationRestrictionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReservationRestrictionType", value)
}

// NewReservationRestrictionTypeFromValue returns a pointer to a valid ReservationRestrictionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReservationRestrictionTypeFromValue(v string) (*ReservationRestrictionType, error) {
	ev := ReservationRestrictionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReservationRestrictionType: valid values are %v", v, AllowedReservationRestrictionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReservationRestrictionType) IsValid() bool {
	for _, existing := range AllowedReservationRestrictionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to reservationRestrictionType value
func (v ReservationRestrictionType) Ptr() *ReservationRestrictionType {
	return &v
}

type NullableReservationRestrictionType struct {
	value *ReservationRestrictionType
	isSet bool
}

func (v NullableReservationRestrictionType) Get() *ReservationRestrictionType {
	return v.value
}

func (v *NullableReservationRestrictionType) Set(val *ReservationRestrictionType) {
	v.value = val
	v.isSet = true
}

func (v NullableReservationRestrictionType) IsSet() bool {
	return v.isSet
}

func (v *NullableReservationRestrictionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservationRestrictionType(val *ReservationRestrictionType) *NullableReservationRestrictionType {
	return &NullableReservationRestrictionType{value: val, isSet: true}
}

func (v NullableReservationRestrictionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservationRestrictionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

