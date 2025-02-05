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

// CrossAccountRequestByAccount struct for CrossAccountRequestByAccount
type CrossAccountRequestByAccount struct {
	RequestId *string `json:"request_id,omitempty"`
	TargetAccount *string `json:"target_account,omitempty"`
	Status *string `json:"status,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	StartDate interface{} `json:"start_date,omitempty"`
	EndDate interface{} `json:"end_date,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	Email *string `json:"email,omitempty"`
}

// NewCrossAccountRequestByAccount instantiates a new CrossAccountRequestByAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCrossAccountRequestByAccount() *CrossAccountRequestByAccount {
	this := CrossAccountRequestByAccount{}
	return &this
}

// NewCrossAccountRequestByAccountWithDefaults instantiates a new CrossAccountRequestByAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCrossAccountRequestByAccountWithDefaults() *CrossAccountRequestByAccount {
	this := CrossAccountRequestByAccount{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *CrossAccountRequestByAccount) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossAccountRequestByAccount) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *CrossAccountRequestByAccount) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *CrossAccountRequestByAccount) SetRequestId(v string) {
	o.RequestId = &v
}

// GetTargetAccount returns the TargetAccount field value if set, zero value otherwise.
func (o *CrossAccountRequestByAccount) GetTargetAccount() string {
	if o == nil || o.TargetAccount == nil {
		var ret string
		return ret
	}
	return *o.TargetAccount
}

// GetTargetAccountOk returns a tuple with the TargetAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossAccountRequestByAccount) GetTargetAccountOk() (*string, bool) {
	if o == nil || o.TargetAccount == nil {
		return nil, false
	}
	return o.TargetAccount, true
}

// HasTargetAccount returns a boolean if a field has been set.
func (o *CrossAccountRequestByAccount) HasTargetAccount() bool {
	if o != nil && o.TargetAccount != nil {
		return true
	}

	return false
}

// SetTargetAccount gets a reference to the given string and assigns it to the TargetAccount field.
func (o *CrossAccountRequestByAccount) SetTargetAccount(v string) {
	o.TargetAccount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CrossAccountRequestByAccount) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossAccountRequestByAccount) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CrossAccountRequestByAccount) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CrossAccountRequestByAccount) SetStatus(v string) {
	o.Status = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *CrossAccountRequestByAccount) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossAccountRequestByAccount) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CrossAccountRequestByAccount) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *CrossAccountRequestByAccount) SetCreated(v time.Time) {
	o.Created = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CrossAccountRequestByAccount) GetStartDate() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CrossAccountRequestByAccount) GetStartDateOk() (*interface{}, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CrossAccountRequestByAccount) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given interface{} and assigns it to the StartDate field.
func (o *CrossAccountRequestByAccount) SetStartDate(v interface{}) {
	o.StartDate = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CrossAccountRequestByAccount) GetEndDate() interface{} {
	if o == nil  {
		var ret interface{}
		return ret
	}
	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CrossAccountRequestByAccount) GetEndDateOk() (*interface{}, bool) {
	if o == nil || o.EndDate == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *CrossAccountRequestByAccount) HasEndDate() bool {
	if o != nil && o.EndDate != nil {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given interface{} and assigns it to the EndDate field.
func (o *CrossAccountRequestByAccount) SetEndDate(v interface{}) {
	o.EndDate = v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *CrossAccountRequestByAccount) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossAccountRequestByAccount) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *CrossAccountRequestByAccount) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *CrossAccountRequestByAccount) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *CrossAccountRequestByAccount) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossAccountRequestByAccount) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *CrossAccountRequestByAccount) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *CrossAccountRequestByAccount) SetLastName(v string) {
	o.LastName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *CrossAccountRequestByAccount) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CrossAccountRequestByAccount) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *CrossAccountRequestByAccount) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *CrossAccountRequestByAccount) SetEmail(v string) {
	o.Email = &v
}

func (o CrossAccountRequestByAccount) MarshalJSON() ([]byte, error) {
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
	if o.FirstName != nil {
		toSerialize["first_name"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["last_name"] = o.LastName
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableCrossAccountRequestByAccount struct {
	value *CrossAccountRequestByAccount
	isSet bool
}

func (v NullableCrossAccountRequestByAccount) Get() *CrossAccountRequestByAccount {
	return v.value
}

func (v *NullableCrossAccountRequestByAccount) Set(val *CrossAccountRequestByAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableCrossAccountRequestByAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableCrossAccountRequestByAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCrossAccountRequestByAccount(val *CrossAccountRequestByAccount) *NullableCrossAccountRequestByAccount {
	return &NullableCrossAccountRequestByAccount{value: val, isSet: true}
}

func (v NullableCrossAccountRequestByAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCrossAccountRequestByAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


