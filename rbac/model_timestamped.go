/*
 * Role Based Access Control
 *
 * The API for Role Based Access Control.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rbac

import (
	"encoding/json"
	"time"
)

// Timestamped struct for Timestamped
type Timestamped struct {
	Created time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

// NewTimestamped instantiates a new Timestamped object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimestamped(created time.Time, modified time.Time, ) *Timestamped {
	this := Timestamped{}
	this.Created = created
	this.Modified = modified
	return &this
}

// NewTimestampedWithDefaults instantiates a new Timestamped object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimestampedWithDefaults() *Timestamped {
	this := Timestamped{}
	return &this
}

// GetCreated returns the Created field value
func (o *Timestamped) GetCreated() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Timestamped) GetCreatedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Timestamped) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *Timestamped) GetModified() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *Timestamped) GetModifiedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *Timestamped) SetModified(v time.Time) {
	o.Modified = v
}

func (o Timestamped) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["modified"] = o.Modified
	}
	return json.Marshal(toSerialize)
}

type NullableTimestamped struct {
	value *Timestamped
	isSet bool
}

func (v NullableTimestamped) Get() *Timestamped {
	return v.value
}

func (v *NullableTimestamped) Set(val *Timestamped) {
	v.value = val
	v.isSet = true
}

func (v NullableTimestamped) IsSet() bool {
	return v.isSet
}

func (v *NullableTimestamped) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimestamped(val *Timestamped) *NullableTimestamped {
	return &NullableTimestamped{value: val, isSet: true}
}

func (v NullableTimestamped) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimestamped) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


