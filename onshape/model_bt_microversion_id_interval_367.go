/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7010-841c6a8f62e7
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMicroversionIdInterval367 struct for BTMicroversionIdInterval367
type BTMicroversionIdInterval367 struct {
	BtType  *string              `json:"btType,omitempty"`
	From    *BTMicroversionId366 `json:"from,omitempty"`
	To      *BTMicroversionId366 `json:"to,omitempty"`
	Trivial *bool                `json:"trivial,omitempty"`
}

// NewBTMicroversionIdInterval367 instantiates a new BTMicroversionIdInterval367 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMicroversionIdInterval367() *BTMicroversionIdInterval367 {
	this := BTMicroversionIdInterval367{}
	return &this
}

// NewBTMicroversionIdInterval367WithDefaults instantiates a new BTMicroversionIdInterval367 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMicroversionIdInterval367WithDefaults() *BTMicroversionIdInterval367 {
	this := BTMicroversionIdInterval367{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMicroversionIdInterval367) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMicroversionIdInterval367) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMicroversionIdInterval367) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMicroversionIdInterval367) SetBtType(v string) {
	o.BtType = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *BTMicroversionIdInterval367) GetFrom() BTMicroversionId366 {
	if o == nil || o.From == nil {
		var ret BTMicroversionId366
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMicroversionIdInterval367) GetFromOk() (*BTMicroversionId366, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *BTMicroversionIdInterval367) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given BTMicroversionId366 and assigns it to the From field.
func (o *BTMicroversionIdInterval367) SetFrom(v BTMicroversionId366) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *BTMicroversionIdInterval367) GetTo() BTMicroversionId366 {
	if o == nil || o.To == nil {
		var ret BTMicroversionId366
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMicroversionIdInterval367) GetToOk() (*BTMicroversionId366, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *BTMicroversionIdInterval367) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given BTMicroversionId366 and assigns it to the To field.
func (o *BTMicroversionIdInterval367) SetTo(v BTMicroversionId366) {
	o.To = &v
}

// GetTrivial returns the Trivial field value if set, zero value otherwise.
func (o *BTMicroversionIdInterval367) GetTrivial() bool {
	if o == nil || o.Trivial == nil {
		var ret bool
		return ret
	}
	return *o.Trivial
}

// GetTrivialOk returns a tuple with the Trivial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMicroversionIdInterval367) GetTrivialOk() (*bool, bool) {
	if o == nil || o.Trivial == nil {
		return nil, false
	}
	return o.Trivial, true
}

// HasTrivial returns a boolean if a field has been set.
func (o *BTMicroversionIdInterval367) HasTrivial() bool {
	if o != nil && o.Trivial != nil {
		return true
	}

	return false
}

// SetTrivial gets a reference to the given bool and assigns it to the Trivial field.
func (o *BTMicroversionIdInterval367) SetTrivial(v bool) {
	o.Trivial = &v
}

func (o BTMicroversionIdInterval367) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	if o.Trivial != nil {
		toSerialize["trivial"] = o.Trivial
	}
	return json.Marshal(toSerialize)
}

type NullableBTMicroversionIdInterval367 struct {
	value *BTMicroversionIdInterval367
	isSet bool
}

func (v NullableBTMicroversionIdInterval367) Get() *BTMicroversionIdInterval367 {
	return v.value
}

func (v *NullableBTMicroversionIdInterval367) Set(val *BTMicroversionIdInterval367) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMicroversionIdInterval367) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMicroversionIdInterval367) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMicroversionIdInterval367(val *BTMicroversionIdInterval367) *NullableBTMicroversionIdInterval367 {
	return &NullableBTMicroversionIdInterval367{value: val, isSet: true}
}

func (v NullableBTMicroversionIdInterval367) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMicroversionIdInterval367) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
