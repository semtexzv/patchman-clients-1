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

// CrossAccountRequestByUserId struct for CrossAccountRequestByUserId
type CrossAccountRequestByUserId struct {
	RequestId *string `json:"request_id,omitempty"`
	TargetAccount *string `json:"target_account,omitempty"`
	Status *string `json:"status,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	StartDate interface{} `json:"start_date,omitempty"`
	EndDate interface{} `json:"end_date,omitempty"`
	UserId *string `json:"user_id,omitempty"`
}

// NewCrossAccountRequestByUserId instantiates a new CrossAccountRequestByUserId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrossAccountRequestByUserId() *CrossAccountRequestByUserId {
	this := CrossAccountRequestByUserId{}
	return &this
}

// NewCrossAccountRequestByUserIdWithDefaults instantiates a new CrossAccountRequestByUserId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrossAccountRequestByUserIdWithDefaults() *CrossAccountRequestByUserId {
	this := CrossAccountRequestByUserId{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *CrossAccountRequestByUserId) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossAccountRequestByUserId) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *CrossAccountRequestByUserId) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *CrossAccountRequestByUserId) SetRequestId(v string) {
	o.RequestId = &v
}

// GetTargetAccount returns the TargetAccount field value if set, zero value otherwise.
func (o *CrossAccountRequestByUserId) GetTargetAccount() string {
	if o == nil || o.TargetAccount == nil {
		var ret string
		return ret
	}
	return *o.TargetAccount
}

// GetTargetAccountOk returns a tuple with the TargetAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossAccountRequestByUserId) GetTargetAccountOk() (*string, bool) {
	if o == nil || o.TargetAccount == nil {
		return nil, false
	}
	return o.TargetAccount, true
}

// HasTargetAccount returns a boolean if a field has been set.
func (o *CrossAccountRequestByUserId) HasTargetAccount() bool {
	if o != nil && o.TargetAccount != nil {
		return true
	}

	return false
}

// SetTargetAccount gets a reference to the given string and assigns it to the TargetAccount field.
func (o *CrossAccountRequestByUserId) SetTargetAccount(v string) {
	o.TargetAccount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CrossAccountRequestByUserId) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossAccountRequestByUserId) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CrossAccountRequestByUserId) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CrossAccountRequestByUserId) SetStatus(v string) {
	o.Status = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *CrossAccountRequestByUserId) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossAccountRequestByUserId) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CrossAccountRequestByUserId) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *CrossAccountRequestByUserId) SetCreated(v time.Time) {
	o.Created = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CrossAccountRequestByUserId) GetStartDate() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CrossAccountRequestByUserId) GetStartDateOk() (*interface{}, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CrossAccountRequestByUserId) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given interface{} and assigns it to the StartDate field.
func (o *CrossAccountRequestByUserId) SetStartDate(v interface{}) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CrossAccountRequestByUserId) GetEndDate() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CrossAccountRequestByUserId) GetEndDateOk() (*interface{}, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *CrossAccountRequestByUserId) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given interface{} and assigns it to the EndDate field.
func (o *CrossAccountRequestByUserId) SetEndDate(v interface{}) {
	o.EndDate = v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *CrossAccountRequestByUserId) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossAccountRequestByUserId) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *CrossAccountRequestByUserId) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *CrossAccountRequestByUserId) SetUserId(v string) {
	o.UserId = &v
}

func (o CrossAccountRequestByUserId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}
	if o.TargetAccount != nil {
		toSerialize["target_account"] = o.TargetAccount
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.StartDate != nil {
		toSerialize["start_date"] = o.StartDate
	}
	if o.EndDate != nil {
		toSerialize["end_date"] = o.EndDate
	}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableCrossAccountRequestByUserId struct {
	value *CrossAccountRequestByUserId
	isSet bool
}

func (v NullableCrossAccountRequestByUserId) Get() *CrossAccountRequestByUserId {
	return v.value
}

func (v *NullableCrossAccountRequestByUserId) Set(val *CrossAccountRequestByUserId) {
	v.value = val
	v.isSet = true
}

func (v NullableCrossAccountRequestByUserId) IsSet() bool {
	return v.isSet
}

func (v *NullableCrossAccountRequestByUserId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCrossAccountRequestByUserId(val *CrossAccountRequestByUserId) *NullableCrossAccountRequestByUserId {
	return &NullableCrossAccountRequestByUserId{value: val, isSet: true}
}

func (v NullableCrossAccountRequestByUserId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrossAccountRequestByUserId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


