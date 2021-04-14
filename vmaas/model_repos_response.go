/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.8.2
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vmaas

import (
	"encoding/json"
)

// ReposResponse struct for ReposResponse
type ReposResponse struct {
	Page *float32 `json:"page,omitempty"`
	PageSize *float32 `json:"page_size,omitempty"`
	Pages *float32 `json:"pages,omitempty"`
	RepositoryList *map[string][]map[string]interface{} `json:"repository_list,omitempty"`
}

// NewReposResponse instantiates a new ReposResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReposResponse() *ReposResponse {
	this := ReposResponse{}
	return &this
}

// NewReposResponseWithDefaults instantiates a new ReposResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReposResponseWithDefaults() *ReposResponse {
	this := ReposResponse{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *ReposResponse) GetPage() float32 {
	if o == nil || o.Page == nil {
		var ret float32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposResponse) GetPageOk() (*float32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *ReposResponse) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given float32 and assigns it to the Page field.
func (o *ReposResponse) SetPage(v float32) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ReposResponse) GetPageSize() float32 {
	if o == nil || o.PageSize == nil {
		var ret float32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposResponse) GetPageSizeOk() (*float32, bool) {
	if o == nil || o.PageSize == nil {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ReposResponse) HasPageSize() bool {
	if o != nil && o.PageSize != nil {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given float32 and assigns it to the PageSize field.
func (o *ReposResponse) SetPageSize(v float32) {
	o.PageSize = &v
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *ReposResponse) GetPages() float32 {
	if o == nil || o.Pages == nil {
		var ret float32
		return ret
	}
	return *o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposResponse) GetPagesOk() (*float32, bool) {
	if o == nil || o.Pages == nil {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *ReposResponse) HasPages() bool {
	if o != nil && o.Pages != nil {
		return true
	}

	return false
}

// SetPages gets a reference to the given float32 and assigns it to the Pages field.
func (o *ReposResponse) SetPages(v float32) {
	o.Pages = &v
}

// GetRepositoryList returns the RepositoryList field value if set, zero value otherwise.
func (o *ReposResponse) GetRepositoryList() map[string][]map[string]interface{} {
	if o == nil || o.RepositoryList == nil {
		var ret map[string][]map[string]interface{}
		return ret
	}
	return *o.RepositoryList
}

// GetRepositoryListOk returns a tuple with the RepositoryList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReposResponse) GetRepositoryListOk() (*map[string][]map[string]interface{}, bool) {
	if o == nil || o.RepositoryList == nil {
		return nil, false
	}
	return o.RepositoryList, true
}

// HasRepositoryList returns a boolean if a field has been set.
func (o *ReposResponse) HasRepositoryList() bool {
	if o != nil && o.RepositoryList != nil {
		return true
	}

	return false
}

// SetRepositoryList gets a reference to the given map[string][]map[string]interface{} and assigns it to the RepositoryList field.
func (o *ReposResponse) SetRepositoryList(v map[string][]map[string]interface{}) {
	o.RepositoryList = &v
}

func (o ReposResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.PageSize != nil {
		toSerialize["page_size"] = o.PageSize
	}
	if o.Pages != nil {
		toSerialize["pages"] = o.Pages
	}
	if o.RepositoryList != nil {
		toSerialize["repository_list"] = o.RepositoryList
	}
	return json.Marshal(toSerialize)
}

type NullableReposResponse struct {
	value *ReposResponse
	isSet bool
}

func (v NullableReposResponse) Get() *ReposResponse {
	return v.value
}

func (v *NullableReposResponse) Set(val *ReposResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableReposResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableReposResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReposResponse(val *ReposResponse) *NullableReposResponse {
	return &NullableReposResponse{value: val, isSet: true}
}

func (v NullableReposResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReposResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


