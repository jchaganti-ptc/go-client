/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5928-bd774e9c9f97
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAppElementReferenceInfo struct for BTAppElementReferenceInfo
type BTAppElementReferenceInfo struct {
	// The latest change id for the element, after the edit was committed.
	ChangeId string `json:"changeId"`
	// The numeric code identifying the error that occurred, if one occurred.
	ErrorCode *int32 `json:"errorCode,omitempty"`
	// A human-readable value for the error that occurred, if one occurred.
	ErrorDescription *string `json:"errorDescription,omitempty"`
	ErrorValue       *string `json:"errorValue,omitempty"`
	// The latest change id for the element, before the edit was made.
	ParentChangeId *string `json:"parentChangeId,omitempty"`
	ReferenceId    *string `json:"referenceId,omitempty"`
	// The id of the transaction in which the edit was applied.
	TransactionId *string `json:"transactionId,omitempty"`
}

// NewBTAppElementReferenceInfo instantiates a new BTAppElementReferenceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAppElementReferenceInfo(changeId string) *BTAppElementReferenceInfo {
	this := BTAppElementReferenceInfo{}
	this.ChangeId = changeId
	return &this
}

// NewBTAppElementReferenceInfoWithDefaults instantiates a new BTAppElementReferenceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAppElementReferenceInfoWithDefaults() *BTAppElementReferenceInfo {
	this := BTAppElementReferenceInfo{}
	return &this
}

// GetChangeId returns the ChangeId field value
func (o *BTAppElementReferenceInfo) GetChangeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChangeId
}

// GetChangeIdOk returns a tuple with the ChangeId field value
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceInfo) GetChangeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChangeId, true
}

// SetChangeId sets field value
func (o *BTAppElementReferenceInfo) SetChangeId(v string) {
	o.ChangeId = v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *BTAppElementReferenceInfo) GetErrorCode() int32 {
	if o == nil || o.ErrorCode == nil {
		var ret int32
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceInfo) GetErrorCodeOk() (*int32, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *BTAppElementReferenceInfo) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given int32 and assigns it to the ErrorCode field.
func (o *BTAppElementReferenceInfo) SetErrorCode(v int32) {
	o.ErrorCode = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *BTAppElementReferenceInfo) GetErrorDescription() string {
	if o == nil || o.ErrorDescription == nil {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceInfo) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || o.ErrorDescription == nil {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *BTAppElementReferenceInfo) HasErrorDescription() bool {
	if o != nil && o.ErrorDescription != nil {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *BTAppElementReferenceInfo) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetErrorValue returns the ErrorValue field value if set, zero value otherwise.
func (o *BTAppElementReferenceInfo) GetErrorValue() string {
	if o == nil || o.ErrorValue == nil {
		var ret string
		return ret
	}
	return *o.ErrorValue
}

// GetErrorValueOk returns a tuple with the ErrorValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceInfo) GetErrorValueOk() (*string, bool) {
	if o == nil || o.ErrorValue == nil {
		return nil, false
	}
	return o.ErrorValue, true
}

// HasErrorValue returns a boolean if a field has been set.
func (o *BTAppElementReferenceInfo) HasErrorValue() bool {
	if o != nil && o.ErrorValue != nil {
		return true
	}

	return false
}

// SetErrorValue gets a reference to the given string and assigns it to the ErrorValue field.
func (o *BTAppElementReferenceInfo) SetErrorValue(v string) {
	o.ErrorValue = &v
}

// GetParentChangeId returns the ParentChangeId field value if set, zero value otherwise.
func (o *BTAppElementReferenceInfo) GetParentChangeId() string {
	if o == nil || o.ParentChangeId == nil {
		var ret string
		return ret
	}
	return *o.ParentChangeId
}

// GetParentChangeIdOk returns a tuple with the ParentChangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceInfo) GetParentChangeIdOk() (*string, bool) {
	if o == nil || o.ParentChangeId == nil {
		return nil, false
	}
	return o.ParentChangeId, true
}

// HasParentChangeId returns a boolean if a field has been set.
func (o *BTAppElementReferenceInfo) HasParentChangeId() bool {
	if o != nil && o.ParentChangeId != nil {
		return true
	}

	return false
}

// SetParentChangeId gets a reference to the given string and assigns it to the ParentChangeId field.
func (o *BTAppElementReferenceInfo) SetParentChangeId(v string) {
	o.ParentChangeId = &v
}

// GetReferenceId returns the ReferenceId field value if set, zero value otherwise.
func (o *BTAppElementReferenceInfo) GetReferenceId() string {
	if o == nil || o.ReferenceId == nil {
		var ret string
		return ret
	}
	return *o.ReferenceId
}

// GetReferenceIdOk returns a tuple with the ReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceInfo) GetReferenceIdOk() (*string, bool) {
	if o == nil || o.ReferenceId == nil {
		return nil, false
	}
	return o.ReferenceId, true
}

// HasReferenceId returns a boolean if a field has been set.
func (o *BTAppElementReferenceInfo) HasReferenceId() bool {
	if o != nil && o.ReferenceId != nil {
		return true
	}

	return false
}

// SetReferenceId gets a reference to the given string and assigns it to the ReferenceId field.
func (o *BTAppElementReferenceInfo) SetReferenceId(v string) {
	o.ReferenceId = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *BTAppElementReferenceInfo) GetTransactionId() string {
	if o == nil || o.TransactionId == nil {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementReferenceInfo) GetTransactionIdOk() (*string, bool) {
	if o == nil || o.TransactionId == nil {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *BTAppElementReferenceInfo) HasTransactionId() bool {
	if o != nil && o.TransactionId != nil {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *BTAppElementReferenceInfo) SetTransactionId(v string) {
	o.TransactionId = &v
}

func (o BTAppElementReferenceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["changeId"] = o.ChangeId
	}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.ErrorDescription != nil {
		toSerialize["errorDescription"] = o.ErrorDescription
	}
	if o.ErrorValue != nil {
		toSerialize["errorValue"] = o.ErrorValue
	}
	if o.ParentChangeId != nil {
		toSerialize["parentChangeId"] = o.ParentChangeId
	}
	if o.ReferenceId != nil {
		toSerialize["referenceId"] = o.ReferenceId
	}
	if o.TransactionId != nil {
		toSerialize["transactionId"] = o.TransactionId
	}
	return json.Marshal(toSerialize)
}

type NullableBTAppElementReferenceInfo struct {
	value *BTAppElementReferenceInfo
	isSet bool
}

func (v NullableBTAppElementReferenceInfo) Get() *BTAppElementReferenceInfo {
	return v.value
}

func (v *NullableBTAppElementReferenceInfo) Set(val *BTAppElementReferenceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppElementReferenceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppElementReferenceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppElementReferenceInfo(val *BTAppElementReferenceInfo) *NullableBTAppElementReferenceInfo {
	return &NullableBTAppElementReferenceInfo{value: val, isSet: true}
}

func (v NullableBTAppElementReferenceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppElementReferenceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
