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

// checks if the CdrTariffsElementsRestrictions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CdrTariffsElementsRestrictions{}

// CdrTariffsElementsRestrictions struct for CdrTariffsElementsRestrictions
type CdrTariffsElementsRestrictions struct {
	StartTime *string `json:"start_time,omitempty"`
	EndTime *string `json:"end_time,omitempty"`
	StartDate *string `json:"start_date,omitempty"`
	EndDate *string `json:"end_date,omitempty"`
	MinKwh *float32 `json:"min_kwh,omitempty"`
	MaxKwh *float32 `json:"max_kwh,omitempty"`
	MinCurrent *float32 `json:"min_current,omitempty"`
	MaxCurrent *float32 `json:"max_current,omitempty"`
	MinPower *float32 `json:"min_power,omitempty"`
	MaxPower *float32 `json:"max_power,omitempty"`
	MinDuration *int32 `json:"min_duration,omitempty"`
	MaxDuration *int32 `json:"max_duration,omitempty"`
	DayOfWeek *string `json:"day_of_week,omitempty"`
	Reservation *string `json:"reservation,omitempty"`
}

// NewCdrTariffsElementsRestrictions instantiates a new CdrTariffsElementsRestrictions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCdrTariffsElementsRestrictions() *CdrTariffsElementsRestrictions {
	this := CdrTariffsElementsRestrictions{}
	return &this
}

// NewCdrTariffsElementsRestrictionsWithDefaults instantiates a new CdrTariffsElementsRestrictions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCdrTariffsElementsRestrictionsWithDefaults() *CdrTariffsElementsRestrictions {
	this := CdrTariffsElementsRestrictions{}
	return &this
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *CdrTariffsElementsRestrictions) GetStartTime() string {
	if o == nil || IsNil(o.StartTime) {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrTariffsElementsRestrictions) GetStartTimeOk() (*string, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *CdrTariffsElementsRestrictions) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *CdrTariffsElementsRestrictions) SetStartTime(v string) {
	o.StartTime = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *CdrTariffsElementsRestrictions) GetEndTime() string {
	if o == nil || IsNil(o.EndTime) {
		var ret string
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrTariffsElementsRestrictions) GetEndTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *CdrTariffsElementsRestrictions) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given string and assigns it to the EndTime field.
func (o *CdrTariffsElementsRestrictions) SetEndTime(v string) {
	o.EndTime = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CdrTariffsElementsRestrictions) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrTariffsElementsRestrictions) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CdrTariffsElementsRestrictions) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *CdrTariffsElementsRestrictions) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *CdrTariffsElementsRestrictions) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrTariffsElementsRestrictions) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *CdrTariffsElementsRestrictions) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *CdrTariffsElementsRestrictions) SetEndDate(v string) {
	o.EndDate = &v
}

// GetMinKwh returns the MinKwh field value if set, zero value otherwise.
func (o *CdrTariffsElementsRestrictions) GetMinKwh() float32 {
	if o == nil || IsNil(o.MinKwh) {
		var ret float32
		return ret
	}
	return *o.MinKwh
}

// GetMinKwhOk returns a tuple with the MinKwh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrTariffsElementsRestrictions) GetMinKwhOk() (*float32, bool) {
	if o == nil || IsNil(o.MinKwh) {
		return nil, false
	}
	return o.MinKwh, true
}

// HasMinKwh returns a boolean if a field has been set.
func (o *CdrTariffsElementsRestrictions) HasMinKwh() bool {
	if o != nil && !IsNil(o.MinKwh) {
		return true
	}

	return false
}

// SetMinKwh gets a reference to the given float32 and assigns it to the MinKwh field.
func (o *CdrTariffsElementsRestrictions) SetMinKwh(v float32) {
	o.MinKwh = &v
}

// GetMaxKwh returns the MaxKwh field value if set, zero value otherwise.
func (o *CdrTariffsElementsRestrictions) GetMaxKwh() float32 {
	if o == nil || IsNil(o.MaxKwh) {
		var ret float32
		return ret
	}
	return *o.MaxKwh
}

// GetMaxKwhOk returns a tuple with the MaxKwh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrTariffsElementsRestrictions) GetMaxKwhOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxKwh) {
		return nil, false
	}
	return o.MaxKwh, true
}

// HasMaxKwh returns a boolean if a field has been set.
func (o *CdrTariffsElementsRestrictions) HasMaxKwh() bool {
	if o != nil && !IsNil(o.MaxKwh) {
		return true
	}

	return false
}

// SetMaxKwh gets a reference to the given float32 and assigns it to the MaxKwh field.
func (o *CdrTariffsElementsRestrictions) SetMaxKwh(v float32) {
	o.MaxKwh = &v
}

// GetMinCurrent returns the MinCurrent field value if set, zero value otherwise.
func (o *CdrTariffsElementsRestrictions) GetMinCurrent() float32 {
	if o == nil || IsNil(o.MinCurrent) {
		var ret float32
		return ret
	}
	return *o.MinCurrent
}

// GetMinCurrentOk returns a tuple with the MinCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrTariffsElementsRestrictions) GetMinCurrentOk() (*float32, bool) {
	if o == nil || IsNil(o.MinCurrent) {
		return nil, false
	}
	return o.MinCurrent, true
}

// HasMinCurrent returns a boolean if a field has been set.
func (o *CdrTariffsElementsRestrictions) HasMinCurrent() bool {
	if o != nil && !IsNil(o.MinCurrent) {
		return true
	}

	return false
}

// SetMinCurrent gets a reference to the given float32 and assigns it to the MinCurrent field.
func (o *CdrTariffsElementsRestrictions) SetMinCurrent(v float32) {
	o.MinCurrent = &v
}

// GetMaxCurrent returns the MaxCurrent field value if set, zero value otherwise.
func (o *CdrTariffsElementsRestrictions) GetMaxCurrent() float32 {
	if o == nil || IsNil(o.MaxCurrent) {
		var ret float32
		return ret
	}
	return *o.MaxCurrent
}

// GetMaxCurrentOk returns a tuple with the MaxCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrTariffsElementsRestrictions) GetMaxCurrentOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxCurrent) {
		return nil, false
	}
	return o.MaxCurrent, true
}

// HasMaxCurrent returns a boolean if a field has been set.
func (o *CdrTariffsElementsRestrictions) HasMaxCurrent() bool {
	if o != nil && !IsNil(o.MaxCurrent) {
		return true
	}

	return false
}

// SetMaxCurrent gets a reference to the given float32 and assigns it to the MaxCurrent field.
func (o *CdrTariffsElementsRestrictions) SetMaxCurrent(v float32) {
	o.MaxCurrent = &v
}

// GetMinPower returns the MinPower field value if set, zero value otherwise.
func (o *CdrTariffsElementsRestrictions) GetMinPower() float32 {
	if o == nil || IsNil(o.MinPower) {
		var ret float32
		return ret
	}
	return *o.MinPower
}

// GetMinPowerOk returns a tuple with the MinPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrTariffsElementsRestrictions) GetMinPowerOk() (*float32, bool) {
	if o == nil || IsNil(o.MinPower) {
		return nil, false
	}
	return o.MinPower, true
}

// HasMinPower returns a boolean if a field has been set.
func (o *CdrTariffsElementsRestrictions) HasMinPower() bool {
	if o != nil && !IsNil(o.MinPower) {
		return true
	}

	return false
}

// SetMinPower gets a reference to the given float32 and assigns it to the MinPower field.
func (o *CdrTariffsElementsRestrictions) SetMinPower(v float32) {
	o.MinPower = &v
}

// GetMaxPower returns the MaxPower field value if set, zero value otherwise.
func (o *CdrTariffsElementsRestrictions) GetMaxPower() float32 {
	if o == nil || IsNil(o.MaxPower) {
		var ret float32
		return ret
	}
	return *o.MaxPower
}

// GetMaxPowerOk returns a tuple with the MaxPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrTariffsElementsRestrictions) GetMaxPowerOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxPower) {
		return nil, false
	}
	return o.MaxPower, true
}

// HasMaxPower returns a boolean if a field has been set.
func (o *CdrTariffsElementsRestrictions) HasMaxPower() bool {
	if o != nil && !IsNil(o.MaxPower) {
		return true
	}

	return false
}

// SetMaxPower gets a reference to the given float32 and assigns it to the MaxPower field.
func (o *CdrTariffsElementsRestrictions) SetMaxPower(v float32) {
	o.MaxPower = &v
}

// GetMinDuration returns the MinDuration field value if set, zero value otherwise.
func (o *CdrTariffsElementsRestrictions) GetMinDuration() int32 {
	if o == nil || IsNil(o.MinDuration) {
		var ret int32
		return ret
	}
	return *o.MinDuration
}

// GetMinDurationOk returns a tuple with the MinDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrTariffsElementsRestrictions) GetMinDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.MinDuration) {
		return nil, false
	}
	return o.MinDuration, true
}

// HasMinDuration returns a boolean if a field has been set.
func (o *CdrTariffsElementsRestrictions) HasMinDuration() bool {
	if o != nil && !IsNil(o.MinDuration) {
		return true
	}

	return false
}

// SetMinDuration gets a reference to the given int32 and assigns it to the MinDuration field.
func (o *CdrTariffsElementsRestrictions) SetMinDuration(v int32) {
	o.MinDuration = &v
}

// GetMaxDuration returns the MaxDuration field value if set, zero value otherwise.
func (o *CdrTariffsElementsRestrictions) GetMaxDuration() int32 {
	if o == nil || IsNil(o.MaxDuration) {
		var ret int32
		return ret
	}
	return *o.MaxDuration
}

// GetMaxDurationOk returns a tuple with the MaxDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrTariffsElementsRestrictions) GetMaxDurationOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxDuration) {
		return nil, false
	}
	return o.MaxDuration, true
}

// HasMaxDuration returns a boolean if a field has been set.
func (o *CdrTariffsElementsRestrictions) HasMaxDuration() bool {
	if o != nil && !IsNil(o.MaxDuration) {
		return true
	}

	return false
}

// SetMaxDuration gets a reference to the given int32 and assigns it to the MaxDuration field.
func (o *CdrTariffsElementsRestrictions) SetMaxDuration(v int32) {
	o.MaxDuration = &v
}

// GetDayOfWeek returns the DayOfWeek field value if set, zero value otherwise.
func (o *CdrTariffsElementsRestrictions) GetDayOfWeek() string {
	if o == nil || IsNil(o.DayOfWeek) {
		var ret string
		return ret
	}
	return *o.DayOfWeek
}

// GetDayOfWeekOk returns a tuple with the DayOfWeek field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrTariffsElementsRestrictions) GetDayOfWeekOk() (*string, bool) {
	if o == nil || IsNil(o.DayOfWeek) {
		return nil, false
	}
	return o.DayOfWeek, true
}

// HasDayOfWeek returns a boolean if a field has been set.
func (o *CdrTariffsElementsRestrictions) HasDayOfWeek() bool {
	if o != nil && !IsNil(o.DayOfWeek) {
		return true
	}

	return false
}

// SetDayOfWeek gets a reference to the given string and assigns it to the DayOfWeek field.
func (o *CdrTariffsElementsRestrictions) SetDayOfWeek(v string) {
	o.DayOfWeek = &v
}

// GetReservation returns the Reservation field value if set, zero value otherwise.
func (o *CdrTariffsElementsRestrictions) GetReservation() string {
	if o == nil || IsNil(o.Reservation) {
		var ret string
		return ret
	}
	return *o.Reservation
}

// GetReservationOk returns a tuple with the Reservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CdrTariffsElementsRestrictions) GetReservationOk() (*string, bool) {
	if o == nil || IsNil(o.Reservation) {
		return nil, false
	}
	return o.Reservation, true
}

// HasReservation returns a boolean if a field has been set.
func (o *CdrTariffsElementsRestrictions) HasReservation() bool {
	if o != nil && !IsNil(o.Reservation) {
		return true
	}

	return false
}

// SetReservation gets a reference to the given string and assigns it to the Reservation field.
func (o *CdrTariffsElementsRestrictions) SetReservation(v string) {
	o.Reservation = &v
}

func (o CdrTariffsElementsRestrictions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CdrTariffsElementsRestrictions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	if !IsNil(o.MinKwh) {
		toSerialize["min_kwh"] = o.MinKwh
	}
	if !IsNil(o.MaxKwh) {
		toSerialize["max_kwh"] = o.MaxKwh
	}
	if !IsNil(o.MinCurrent) {
		toSerialize["min_current"] = o.MinCurrent
	}
	if !IsNil(o.MaxCurrent) {
		toSerialize["max_current"] = o.MaxCurrent
	}
	if !IsNil(o.MinPower) {
		toSerialize["min_power"] = o.MinPower
	}
	if !IsNil(o.MaxPower) {
		toSerialize["max_power"] = o.MaxPower
	}
	if !IsNil(o.MinDuration) {
		toSerialize["min_duration"] = o.MinDuration
	}
	if !IsNil(o.MaxDuration) {
		toSerialize["max_duration"] = o.MaxDuration
	}
	if !IsNil(o.DayOfWeek) {
		toSerialize["day_of_week"] = o.DayOfWeek
	}
	if !IsNil(o.Reservation) {
		toSerialize["reservation"] = o.Reservation
	}
	return toSerialize, nil
}

type NullableCdrTariffsElementsRestrictions struct {
	value *CdrTariffsElementsRestrictions
	isSet bool
}

func (v NullableCdrTariffsElementsRestrictions) Get() *CdrTariffsElementsRestrictions {
	return v.value
}

func (v *NullableCdrTariffsElementsRestrictions) Set(val *CdrTariffsElementsRestrictions) {
	v.value = val
	v.isSet = true
}

func (v NullableCdrTariffsElementsRestrictions) IsSet() bool {
	return v.isSet
}

func (v *NullableCdrTariffsElementsRestrictions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCdrTariffsElementsRestrictions(val *CdrTariffsElementsRestrictions) *NullableCdrTariffsElementsRestrictions {
	return &NullableCdrTariffsElementsRestrictions{value: val, isSet: true}
}

func (v NullableCdrTariffsElementsRestrictions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCdrTariffsElementsRestrictions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


