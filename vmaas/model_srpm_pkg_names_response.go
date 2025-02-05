/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vmaas

import (
	"encoding/json"
)

// SRPMPkgNamesResponse struct for SRPMPkgNamesResponse
type SRPMPkgNamesResponse struct {
	LastChange interface{} `json:"last_change,omitempty"`
	SrpmNameList *map[string]map[string][]string `json:"srpm_name_list,omitempty"`
}

// NewSRPMPkgNamesResponse instantiates a new SRPMPkgNamesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSRPMPkgNamesResponse() *SRPMPkgNamesResponse {
	this := SRPMPkgNamesResponse{}
	return &this
}

// NewSRPMPkgNamesResponseWithDefaults instantiates a new SRPMPkgNamesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSRPMPkgNamesResponseWithDefaults() *SRPMPkgNamesResponse {
	this := SRPMPkgNamesResponse{}
	return &this
}

// GetLastChange returns the LastChange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SRPMPkgNamesResponse) GetLastChange() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.LastChange
}

// GetLastChangeOk returns a tuple with the LastChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SRPMPkgNamesResponse) GetLastChangeOk() (*interface{}, bool) {
	if o == nil || o.LastChange == nil {
		return nil, false
	}
	return &o.LastChange, true
}

// HasLastChange returns a boolean if a field has been set.
func (o *SRPMPkgNamesResponse) HasLastChange() bool {
	if o != nil && o.LastChange != nil {
		return true
	}

	return false
}

// SetLastChange gets a reference to the given interface{} and assigns it to the LastChange field.
func (o *SRPMPkgNamesResponse) SetLastChange(v interface{}) {
	o.LastChange = v
}

// GetSrpmNameList returns the SrpmNameList field value if set, zero value otherwise.
func (o *SRPMPkgNamesResponse) GetSrpmNameList() map[string]map[string][]string {
	if o == nil || o.SrpmNameList == nil {
		var ret map[string]map[string][]string
		return ret
	}
	return *o.SrpmNameList
}

// GetSrpmNameListOk returns a tuple with the SrpmNameList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SRPMPkgNamesResponse) GetSrpmNameListOk() (*map[string]map[string][]string, bool) {
	if o == nil || o.SrpmNameList == nil {
		return nil, false
	}
	return o.SrpmNameList, true
}

// HasSrpmNameList returns a boolean if a field has been set.
func (o *SRPMPkgNamesResponse) HasSrpmNameList() bool {
	if o != nil && o.SrpmNameList != nil {
		return true
	}

	return false
}

// SetSrpmNameList gets a reference to the given map[string]map[string][]string and assigns it to the SrpmNameList field.
func (o *SRPMPkgNamesResponse) SetSrpmNameList(v map[string]map[string][]string) {
	o.SrpmNameList = &v
}

func (o SRPMPkgNamesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastChange != nil {
		toSerialize["last_change"] = o.LastChange
	}
	if o.SrpmNameList != nil {
		toSerialize["srpm_name_list"] = o.SrpmNameList
	}
	return json.Marshal(toSerialize)
}

type NullableSRPMPkgNamesResponse struct {
	value *SRPMPkgNamesResponse
	isSet bool
}

func (v NullableSRPMPkgNamesResponse) Get() *SRPMPkgNamesResponse {
	return v.value
}

func (v *NullableSRPMPkgNamesResponse) Set(val *SRPMPkgNamesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSRPMPkgNamesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSRPMPkgNamesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSRPMPkgNamesResponse(val *SRPMPkgNamesResponse) *NullableSRPMPkgNamesResponse {
	return &NullableSRPMPkgNamesResponse{value: val, isSet: true}
}

func (v NullableSRPMPkgNamesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSRPMPkgNamesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


