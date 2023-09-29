/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.170.23149-3610d8cd750e
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTOriginDisplayData934 struct for BTOriginDisplayData934
type BTOriginDisplayData934 struct {
	BtType *string `json:"btType,omitempty"`
	Hidden *bool   `json:"hidden,omitempty"`
}

// NewBTOriginDisplayData934 instantiates a new BTOriginDisplayData934 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTOriginDisplayData934() *BTOriginDisplayData934 {
	this := BTOriginDisplayData934{}
	return &this
}

// NewBTOriginDisplayData934WithDefaults instantiates a new BTOriginDisplayData934 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTOriginDisplayData934WithDefaults() *BTOriginDisplayData934 {
	this := BTOriginDisplayData934{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTOriginDisplayData934) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOriginDisplayData934) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTOriginDisplayData934) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTOriginDisplayData934) SetBtType(v string) {
	o.BtType = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *BTOriginDisplayData934) GetHidden() bool {
	if o == nil || o.Hidden == nil {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOriginDisplayData934) GetHiddenOk() (*bool, bool) {
	if o == nil || o.Hidden == nil {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *BTOriginDisplayData934) HasHidden() bool {
	if o != nil && o.Hidden != nil {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *BTOriginDisplayData934) SetHidden(v bool) {
	o.Hidden = &v
}

func (o BTOriginDisplayData934) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Hidden != nil {
		toSerialize["hidden"] = o.Hidden
	}
	return json.Marshal(toSerialize)
}

type NullableBTOriginDisplayData934 struct {
	value *BTOriginDisplayData934
	isSet bool
}

func (v NullableBTOriginDisplayData934) Get() *BTOriginDisplayData934 {
	return v.value
}

func (v *NullableBTOriginDisplayData934) Set(val *BTOriginDisplayData934) {
	v.value = val
	v.isSet = true
}

func (v NullableBTOriginDisplayData934) IsSet() bool {
	return v.isSet
}

func (v *NullableBTOriginDisplayData934) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTOriginDisplayData934(val *BTOriginDisplayData934) *NullableBTOriginDisplayData934 {
	return &NullableBTOriginDisplayData934{value: val, isSet: true}
}

func (v NullableBTOriginDisplayData934) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTOriginDisplayData934) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
