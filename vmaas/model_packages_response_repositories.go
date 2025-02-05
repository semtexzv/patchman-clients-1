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

// PackagesResponseRepositories struct for PackagesResponseRepositories
type PackagesResponseRepositories struct {
	Label *string `json:"label,omitempty"`
	Name *string `json:"name,omitempty"`
	Basearch *string `json:"basearch,omitempty"`
	Releasever *string `json:"releasever,omitempty"`
}

// NewPackagesResponseRepositories instantiates a new PackagesResponseRepositories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackagesResponseRepositories() *PackagesResponseRepositories {
	this := PackagesResponseRepositories{}
	return &this
}

// NewPackagesResponseRepositoriesWithDefaults instantiates a new PackagesResponseRepositories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackagesResponseRepositoriesWithDefaults() *PackagesResponseRepositories {
	this := PackagesResponseRepositories{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PackagesResponseRepositories) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesResponseRepositories) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PackagesResponseRepositories) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PackagesResponseRepositories) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PackagesResponseRepositories) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesResponseRepositories) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PackagesResponseRepositories) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PackagesResponseRepositories) SetName(v string) {
	o.Name = &v
}

// GetBasearch returns the Basearch field value if set, zero value otherwise.
func (o *PackagesResponseRepositories) GetBasearch() string {
	if o == nil || o.Basearch == nil {
		var ret string
		return ret
	}
	return *o.Basearch
}

// GetBasearchOk returns a tuple with the Basearch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesResponseRepositories) GetBasearchOk() (*string, bool) {
	if o == nil || o.Basearch == nil {
		return nil, false
	}
	return o.Basearch, true
}

// HasBasearch returns a boolean if a field has been set.
func (o *PackagesResponseRepositories) HasBasearch() bool {
	if o != nil && o.Basearch != nil {
		return true
	}

	return false
}

// SetBasearch gets a reference to the given string and assigns it to the Basearch field.
func (o *PackagesResponseRepositories) SetBasearch(v string) {
	o.Basearch = &v
}

// GetReleasever returns the Releasever field value if set, zero value otherwise.
func (o *PackagesResponseRepositories) GetReleasever() string {
	if o == nil || o.Releasever == nil {
		var ret string
		return ret
	}
	return *o.Releasever
}

// GetReleaseverOk returns a tuple with the Releasever field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackagesResponseRepositories) GetReleaseverOk() (*string, bool) {
	if o == nil || o.Releasever == nil {
		return nil, false
	}
	return o.Releasever, true
}

// HasReleasever returns a boolean if a field has been set.
func (o *PackagesResponseRepositories) HasReleasever() bool {
	if o != nil && o.Releasever != nil {
		return true
	}

	return false
}

// SetReleasever gets a reference to the given string and assigns it to the Releasever field.
func (o *PackagesResponseRepositories) SetReleasever(v string) {
	o.Releasever = &v
}

func (o PackagesResponseRepositories) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Basearch != nil {
		toSerialize["basearch"] = o.Basearch
	}
	if o.Releasever != nil {
		toSerialize["releasever"] = o.Releasever
	}
	return json.Marshal(toSerialize)
}

type NullablePackagesResponseRepositories struct {
	value *PackagesResponseRepositories
	isSet bool
}

func (v NullablePackagesResponseRepositories) Get() *PackagesResponseRepositories {
	return v.value
}

func (v *NullablePackagesResponseRepositories) Set(val *PackagesResponseRepositories) {
	v.value = val
	v.isSet = true
}

func (v NullablePackagesResponseRepositories) IsSet() bool {
	return v.isSet
}

func (v *NullablePackagesResponseRepositories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackagesResponseRepositories(val *PackagesResponseRepositories) *NullablePackagesResponseRepositories {
	return &NullablePackagesResponseRepositories{value: val, isSet: true}
}

func (v NullablePackagesResponseRepositories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackagesResponseRepositories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


