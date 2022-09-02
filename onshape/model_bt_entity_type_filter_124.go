/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6326-97b3616ccba2
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTEntityTypeFilter124 struct for BTEntityTypeFilter124
type BTEntityTypeFilter124 struct {
	BtType     *string `json:"btType,omitempty"`
	EntityType *string `json:"entityType,omitempty"`
}

// NewBTEntityTypeFilter124 instantiates a new BTEntityTypeFilter124 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTEntityTypeFilter124() *BTEntityTypeFilter124 {
	this := BTEntityTypeFilter124{}
	return &this
}

// NewBTEntityTypeFilter124WithDefaults instantiates a new BTEntityTypeFilter124 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTEntityTypeFilter124WithDefaults() *BTEntityTypeFilter124 {
	this := BTEntityTypeFilter124{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTEntityTypeFilter124) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityTypeFilter124) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTEntityTypeFilter124) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTEntityTypeFilter124) SetBtType(v string) {
	o.BtType = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *BTEntityTypeFilter124) GetEntityType() string {
	if o == nil || o.EntityType == nil {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEntityTypeFilter124) GetEntityTypeOk() (*string, bool) {
	if o == nil || o.EntityType == nil {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *BTEntityTypeFilter124) HasEntityType() bool {
	if o != nil && o.EntityType != nil {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *BTEntityTypeFilter124) SetEntityType(v string) {
	o.EntityType = &v
}

func (o BTEntityTypeFilter124) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.EntityType != nil {
		toSerialize["entityType"] = o.EntityType
	}
	return json.Marshal(toSerialize)
}

type NullableBTEntityTypeFilter124 struct {
	value *BTEntityTypeFilter124
	isSet bool
}

func (v NullableBTEntityTypeFilter124) Get() *BTEntityTypeFilter124 {
	return v.value
}

func (v *NullableBTEntityTypeFilter124) Set(val *BTEntityTypeFilter124) {
	v.value = val
	v.isSet = true
}

func (v NullableBTEntityTypeFilter124) IsSet() bool {
	return v.isSet
}

func (v *NullableBTEntityTypeFilter124) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTEntityTypeFilter124(val *BTEntityTypeFilter124) *NullableBTEntityTypeFilter124 {
	return &NullableBTEntityTypeFilter124{value: val, isSet: true}
}

func (v NullableBTEntityTypeFilter124) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTEntityTypeFilter124) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
