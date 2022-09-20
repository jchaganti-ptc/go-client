/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6524-ea93b01144ea
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAppElementHistoryEntryInfo struct for BTAppElementHistoryEntryInfo
type BTAppElementHistoryEntryInfo struct {
	ChangeId    *string   `json:"changeId,omitempty"`
	CreatedAt   *JSONTime `json:"createdAt,omitempty"`
	Description *string   `json:"description,omitempty"`
}

// NewBTAppElementHistoryEntryInfo instantiates a new BTAppElementHistoryEntryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAppElementHistoryEntryInfo() *BTAppElementHistoryEntryInfo {
	this := BTAppElementHistoryEntryInfo{}
	return &this
}

// NewBTAppElementHistoryEntryInfoWithDefaults instantiates a new BTAppElementHistoryEntryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAppElementHistoryEntryInfoWithDefaults() *BTAppElementHistoryEntryInfo {
	this := BTAppElementHistoryEntryInfo{}
	return &this
}

// GetChangeId returns the ChangeId field value if set, zero value otherwise.
func (o *BTAppElementHistoryEntryInfo) GetChangeId() string {
	if o == nil || o.ChangeId == nil {
		var ret string
		return ret
	}
	return *o.ChangeId
}

// GetChangeIdOk returns a tuple with the ChangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementHistoryEntryInfo) GetChangeIdOk() (*string, bool) {
	if o == nil || o.ChangeId == nil {
		return nil, false
	}
	return o.ChangeId, true
}

// HasChangeId returns a boolean if a field has been set.
func (o *BTAppElementHistoryEntryInfo) HasChangeId() bool {
	if o != nil && o.ChangeId != nil {
		return true
	}

	return false
}

// SetChangeId gets a reference to the given string and assigns it to the ChangeId field.
func (o *BTAppElementHistoryEntryInfo) SetChangeId(v string) {
	o.ChangeId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BTAppElementHistoryEntryInfo) GetCreatedAt() JSONTime {
	if o == nil || o.CreatedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementHistoryEntryInfo) GetCreatedAtOk() (*JSONTime, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BTAppElementHistoryEntryInfo) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given JSONTime and assigns it to the CreatedAt field.
func (o *BTAppElementHistoryEntryInfo) SetCreatedAt(v JSONTime) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTAppElementHistoryEntryInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppElementHistoryEntryInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTAppElementHistoryEntryInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTAppElementHistoryEntryInfo) SetDescription(v string) {
	o.Description = &v
}

func (o BTAppElementHistoryEntryInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangeId != nil {
		toSerialize["changeId"] = o.ChangeId
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableBTAppElementHistoryEntryInfo struct {
	value *BTAppElementHistoryEntryInfo
	isSet bool
}

func (v NullableBTAppElementHistoryEntryInfo) Get() *BTAppElementHistoryEntryInfo {
	return v.value
}

func (v *NullableBTAppElementHistoryEntryInfo) Set(val *BTAppElementHistoryEntryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppElementHistoryEntryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppElementHistoryEntryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppElementHistoryEntryInfo(val *BTAppElementHistoryEntryInfo) *NullableBTAppElementHistoryEntryInfo {
	return &NullableBTAppElementHistoryEntryInfo{value: val, isSet: true}
}

func (v NullableBTAppElementHistoryEntryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppElementHistoryEntryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
