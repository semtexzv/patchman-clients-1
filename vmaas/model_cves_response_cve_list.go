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
	"time"
)

// CvesResponseCveList struct for CvesResponseCveList
type CvesResponseCveList struct {
	Impact *string `json:"impact,omitempty"`
	PublicDate *time.Time `json:"public_date,omitempty"`
	Synopsis *string `json:"synopsis,omitempty"`
	Description *string `json:"description,omitempty"`
	ModifiedDate *time.Time `json:"modified_date,omitempty"`
	RedhatUrl *string `json:"redhat_url,omitempty"`
	SecondaryUrl *string `json:"secondary_url,omitempty"`
	Cvss2Score *string `json:"cvss2_score,omitempty"`
	Cvss2Metrics *string `json:"cvss2_metrics,omitempty"`
	Cvss3Score *string `json:"cvss3_score,omitempty"`
	Cvss3Metrics *string `json:"cvss3_metrics,omitempty"`
	CweList *[]string `json:"cwe_list,omitempty"`
	ErrataList *[]string `json:"errata_list,omitempty"`
	PackageList *[]string `json:"package_list,omitempty"`
	SourcePackageList *[]string `json:"source_package_list,omitempty"`
}

// NewCvesResponseCveList instantiates a new CvesResponseCveList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCvesResponseCveList() *CvesResponseCveList {
	this := CvesResponseCveList{}
	return &this
}

// NewCvesResponseCveListWithDefaults instantiates a new CvesResponseCveList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCvesResponseCveListWithDefaults() *CvesResponseCveList {
	this := CvesResponseCveList{}
	return &this
}

// GetImpact returns the Impact field value if set, zero value otherwise.
func (o *CvesResponseCveList) GetImpact() string {
	if o == nil || o.Impact == nil {
		var ret string
		return ret
	}
	return *o.Impact
}

// GetImpactOk returns a tuple with the Impact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesResponseCveList) GetImpactOk() (*string, bool) {
	if o == nil || o.Impact == nil {
		return nil, false
	}
	return o.Impact, true
}

// HasImpact returns a boolean if a field has been set.
func (o *CvesResponseCveList) HasImpact() bool {
	if o != nil && o.Impact != nil {
		return true
	}

	return false
}

// SetImpact gets a reference to the given string and assigns it to the Impact field.
func (o *CvesResponseCveList) SetImpact(v string) {
	o.Impact = &v
}

// GetPublicDate returns the PublicDate field value if set, zero value otherwise.
func (o *CvesResponseCveList) GetPublicDate() time.Time {
	if o == nil || o.PublicDate == nil {
		var ret time.Time
		return ret
	}
	return *o.PublicDate
}

// GetPublicDateOk returns a tuple with the PublicDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesResponseCveList) GetPublicDateOk() (*time.Time, bool) {
	if o == nil || o.PublicDate == nil {
		return nil, false
	}
	return o.PublicDate, true
}

// HasPublicDate returns a boolean if a field has been set.
func (o *CvesResponseCveList) HasPublicDate() bool {
	if o != nil && o.PublicDate != nil {
		return true
	}

	return false
}

// SetPublicDate gets a reference to the given time.Time and assigns it to the PublicDate field.
func (o *CvesResponseCveList) SetPublicDate(v time.Time) {
	o.PublicDate = &v
}

// GetSynopsis returns the Synopsis field value if set, zero value otherwise.
func (o *CvesResponseCveList) GetSynopsis() string {
	if o == nil || o.Synopsis == nil {
		var ret string
		return ret
	}
	return *o.Synopsis
}

// GetSynopsisOk returns a tuple with the Synopsis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesResponseCveList) GetSynopsisOk() (*string, bool) {
	if o == nil || o.Synopsis == nil {
		return nil, false
	}
	return o.Synopsis, true
}

// HasSynopsis returns a boolean if a field has been set.
func (o *CvesResponseCveList) HasSynopsis() bool {
	if o != nil && o.Synopsis != nil {
		return true
	}

	return false
}

// SetSynopsis gets a reference to the given string and assigns it to the Synopsis field.
func (o *CvesResponseCveList) SetSynopsis(v string) {
	o.Synopsis = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CvesResponseCveList) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesResponseCveList) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CvesResponseCveList) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CvesResponseCveList) SetDescription(v string) {
	o.Description = &v
}

// GetModifiedDate returns the ModifiedDate field value if set, zero value otherwise.
func (o *CvesResponseCveList) GetModifiedDate() time.Time {
	if o == nil || o.ModifiedDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDate
}

// GetModifiedDateOk returns a tuple with the ModifiedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesResponseCveList) GetModifiedDateOk() (*time.Time, bool) {
	if o == nil || o.ModifiedDate == nil {
		return nil, false
	}
	return o.ModifiedDate, true
}

// HasModifiedDate returns a boolean if a field has been set.
func (o *CvesResponseCveList) HasModifiedDate() bool {
	if o != nil && o.ModifiedDate != nil {
		return true
	}

	return false
}

// SetModifiedDate gets a reference to the given time.Time and assigns it to the ModifiedDate field.
func (o *CvesResponseCveList) SetModifiedDate(v time.Time) {
	o.ModifiedDate = &v
}

// GetRedhatUrl returns the RedhatUrl field value if set, zero value otherwise.
func (o *CvesResponseCveList) GetRedhatUrl() string {
	if o == nil || o.RedhatUrl == nil {
		var ret string
		return ret
	}
	return *o.RedhatUrl
}

// GetRedhatUrlOk returns a tuple with the RedhatUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesResponseCveList) GetRedhatUrlOk() (*string, bool) {
	if o == nil || o.RedhatUrl == nil {
		return nil, false
	}
	return o.RedhatUrl, true
}

// HasRedhatUrl returns a boolean if a field has been set.
func (o *CvesResponseCveList) HasRedhatUrl() bool {
	if o != nil && o.RedhatUrl != nil {
		return true
	}

	return false
}

// SetRedhatUrl gets a reference to the given string and assigns it to the RedhatUrl field.
func (o *CvesResponseCveList) SetRedhatUrl(v string) {
	o.RedhatUrl = &v
}

// GetSecondaryUrl returns the SecondaryUrl field value if set, zero value otherwise.
func (o *CvesResponseCveList) GetSecondaryUrl() string {
	if o == nil || o.SecondaryUrl == nil {
		var ret string
		return ret
	}
	return *o.SecondaryUrl
}

// GetSecondaryUrlOk returns a tuple with the SecondaryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesResponseCveList) GetSecondaryUrlOk() (*string, bool) {
	if o == nil || o.SecondaryUrl == nil {
		return nil, false
	}
	return o.SecondaryUrl, true
}

// HasSecondaryUrl returns a boolean if a field has been set.
func (o *CvesResponseCveList) HasSecondaryUrl() bool {
	if o != nil && o.SecondaryUrl != nil {
		return true
	}

	return false
}

// SetSecondaryUrl gets a reference to the given string and assigns it to the SecondaryUrl field.
func (o *CvesResponseCveList) SetSecondaryUrl(v string) {
	o.SecondaryUrl = &v
}

// GetCvss2Score returns the Cvss2Score field value if set, zero value otherwise.
func (o *CvesResponseCveList) GetCvss2Score() string {
	if o == nil || o.Cvss2Score == nil {
		var ret string
		return ret
	}
	return *o.Cvss2Score
}

// GetCvss2ScoreOk returns a tuple with the Cvss2Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesResponseCveList) GetCvss2ScoreOk() (*string, bool) {
	if o == nil || o.Cvss2Score == nil {
		return nil, false
	}
	return o.Cvss2Score, true
}

// HasCvss2Score returns a boolean if a field has been set.
func (o *CvesResponseCveList) HasCvss2Score() bool {
	if o != nil && o.Cvss2Score != nil {
		return true
	}

	return false
}

// SetCvss2Score gets a reference to the given string and assigns it to the Cvss2Score field.
func (o *CvesResponseCveList) SetCvss2Score(v string) {
	o.Cvss2Score = &v
}

// GetCvss2Metrics returns the Cvss2Metrics field value if set, zero value otherwise.
func (o *CvesResponseCveList) GetCvss2Metrics() string {
	if o == nil || o.Cvss2Metrics == nil {
		var ret string
		return ret
	}
	return *o.Cvss2Metrics
}

// GetCvss2MetricsOk returns a tuple with the Cvss2Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesResponseCveList) GetCvss2MetricsOk() (*string, bool) {
	if o == nil || o.Cvss2Metrics == nil {
		return nil, false
	}
	return o.Cvss2Metrics, true
}

// HasCvss2Metrics returns a boolean if a field has been set.
func (o *CvesResponseCveList) HasCvss2Metrics() bool {
	if o != nil && o.Cvss2Metrics != nil {
		return true
	}

	return false
}

// SetCvss2Metrics gets a reference to the given string and assigns it to the Cvss2Metrics field.
func (o *CvesResponseCveList) SetCvss2Metrics(v string) {
	o.Cvss2Metrics = &v
}

// GetCvss3Score returns the Cvss3Score field value if set, zero value otherwise.
func (o *CvesResponseCveList) GetCvss3Score() string {
	if o == nil || o.Cvss3Score == nil {
		var ret string
		return ret
	}
	return *o.Cvss3Score
}

// GetCvss3ScoreOk returns a tuple with the Cvss3Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesResponseCveList) GetCvss3ScoreOk() (*string, bool) {
	if o == nil || o.Cvss3Score == nil {
		return nil, false
	}
	return o.Cvss3Score, true
}

// HasCvss3Score returns a boolean if a field has been set.
func (o *CvesResponseCveList) HasCvss3Score() bool {
	if o != nil && o.Cvss3Score != nil {
		return true
	}

	return false
}

// SetCvss3Score gets a reference to the given string and assigns it to the Cvss3Score field.
func (o *CvesResponseCveList) SetCvss3Score(v string) {
	o.Cvss3Score = &v
}

// GetCvss3Metrics returns the Cvss3Metrics field value if set, zero value otherwise.
func (o *CvesResponseCveList) GetCvss3Metrics() string {
	if o == nil || o.Cvss3Metrics == nil {
		var ret string
		return ret
	}
	return *o.Cvss3Metrics
}

// GetCvss3MetricsOk returns a tuple with the Cvss3Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesResponseCveList) GetCvss3MetricsOk() (*string, bool) {
	if o == nil || o.Cvss3Metrics == nil {
		return nil, false
	}
	return o.Cvss3Metrics, true
}

// HasCvss3Metrics returns a boolean if a field has been set.
func (o *CvesResponseCveList) HasCvss3Metrics() bool {
	if o != nil && o.Cvss3Metrics != nil {
		return true
	}

	return false
}

// SetCvss3Metrics gets a reference to the given string and assigns it to the Cvss3Metrics field.
func (o *CvesResponseCveList) SetCvss3Metrics(v string) {
	o.Cvss3Metrics = &v
}

// GetCweList returns the CweList field value if set, zero value otherwise.
func (o *CvesResponseCveList) GetCweList() []string {
	if o == nil || o.CweList == nil {
		var ret []string
		return ret
	}
	return *o.CweList
}

// GetCweListOk returns a tuple with the CweList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesResponseCveList) GetCweListOk() (*[]string, bool) {
	if o == nil || o.CweList == nil {
		return nil, false
	}
	return o.CweList, true
}

// HasCweList returns a boolean if a field has been set.
func (o *CvesResponseCveList) HasCweList() bool {
	if o != nil && o.CweList != nil {
		return true
	}

	return false
}

// SetCweList gets a reference to the given []string and assigns it to the CweList field.
func (o *CvesResponseCveList) SetCweList(v []string) {
	o.CweList = &v
}

// GetErrataList returns the ErrataList field value if set, zero value otherwise.
func (o *CvesResponseCveList) GetErrataList() []string {
	if o == nil || o.ErrataList == nil {
		var ret []string
		return ret
	}
	return *o.ErrataList
}

// GetErrataListOk returns a tuple with the ErrataList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesResponseCveList) GetErrataListOk() (*[]string, bool) {
	if o == nil || o.ErrataList == nil {
		return nil, false
	}
	return o.ErrataList, true
}

// HasErrataList returns a boolean if a field has been set.
func (o *CvesResponseCveList) HasErrataList() bool {
	if o != nil && o.ErrataList != nil {
		return true
	}

	return false
}

// SetErrataList gets a reference to the given []string and assigns it to the ErrataList field.
func (o *CvesResponseCveList) SetErrataList(v []string) {
	o.ErrataList = &v
}

// GetPackageList returns the PackageList field value if set, zero value otherwise.
func (o *CvesResponseCveList) GetPackageList() []string {
	if o == nil || o.PackageList == nil {
		var ret []string
		return ret
	}
	return *o.PackageList
}

// GetPackageListOk returns a tuple with the PackageList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesResponseCveList) GetPackageListOk() (*[]string, bool) {
	if o == nil || o.PackageList == nil {
		return nil, false
	}
	return o.PackageList, true
}

// HasPackageList returns a boolean if a field has been set.
func (o *CvesResponseCveList) HasPackageList() bool {
	if o != nil && o.PackageList != nil {
		return true
	}

	return false
}

// SetPackageList gets a reference to the given []string and assigns it to the PackageList field.
func (o *CvesResponseCveList) SetPackageList(v []string) {
	o.PackageList = &v
}

// GetSourcePackageList returns the SourcePackageList field value if set, zero value otherwise.
func (o *CvesResponseCveList) GetSourcePackageList() []string {
	if o == nil || o.SourcePackageList == nil {
		var ret []string
		return ret
	}
	return *o.SourcePackageList
}

// GetSourcePackageListOk returns a tuple with the SourcePackageList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CvesResponseCveList) GetSourcePackageListOk() (*[]string, bool) {
	if o == nil || o.SourcePackageList == nil {
		return nil, false
	}
	return o.SourcePackageList, true
}

// HasSourcePackageList returns a boolean if a field has been set.
func (o *CvesResponseCveList) HasSourcePackageList() bool {
	if o != nil && o.SourcePackageList != nil {
		return true
	}

	return false
}

// SetSourcePackageList gets a reference to the given []string and assigns it to the SourcePackageList field.
func (o *CvesResponseCveList) SetSourcePackageList(v []string) {
	o.SourcePackageList = &v
}

func (o CvesResponseCveList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Impact != nil {
		toSerialize["impact"] = o.Impact
	}
	if o.PublicDate != nil {
		toSerialize["public_date"] = o.PublicDate
	}
	if o.Synopsis != nil {
		toSerialize["synopsis"] = o.Synopsis
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ModifiedDate != nil {
		toSerialize["modified_date"] = o.ModifiedDate
	}
	if o.RedhatUrl != nil {
		toSerialize["redhat_url"] = o.RedhatUrl
	}
	if o.SecondaryUrl != nil {
		toSerialize["secondary_url"] = o.SecondaryUrl
	}
	if o.Cvss2Score != nil {
		toSerialize["cvss2_score"] = o.Cvss2Score
	}
	if o.Cvss2Metrics != nil {
		toSerialize["cvss2_metrics"] = o.Cvss2Metrics
	}
	if o.Cvss3Score != nil {
		toSerialize["cvss3_score"] = o.Cvss3Score
	}
	if o.Cvss3Metrics != nil {
		toSerialize["cvss3_metrics"] = o.Cvss3Metrics
	}
	if o.CweList != nil {
		toSerialize["cwe_list"] = o.CweList
	}
	if o.ErrataList != nil {
		toSerialize["errata_list"] = o.ErrataList
	}
	if o.PackageList != nil {
		toSerialize["package_list"] = o.PackageList
	}
	if o.SourcePackageList != nil {
		toSerialize["source_package_list"] = o.SourcePackageList
	}
	return json.Marshal(toSerialize)
}

type NullableCvesResponseCveList struct {
	value *CvesResponseCveList
	isSet bool
}

func (v NullableCvesResponseCveList) Get() *CvesResponseCveList {
	return v.value
}

func (v *NullableCvesResponseCveList) Set(val *CvesResponseCveList) {
	v.value = val
	v.isSet = true
}

func (v NullableCvesResponseCveList) IsSet() bool {
	return v.isSet
}

func (v *NullableCvesResponseCveList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCvesResponseCveList(val *CvesResponseCveList) *NullableCvesResponseCveList {
	return &NullableCvesResponseCveList{value: val, isSet: true}
}

func (v NullableCvesResponseCveList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCvesResponseCveList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


