/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.173.27313-c3b730633f3c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTBaseEntityAppearanceSettings1391 struct for BTBaseEntityAppearanceSettings1391
type BTBaseEntityAppearanceSettings1391 struct {
	BtType                             *string                                     `json:"btType,omitempty"`
	ColorIdToBaseEntityAppearanceEntry *map[string]BTBaseEntityAppearanceEntry3607 `json:"colorIdToBaseEntityAppearanceEntry,omitempty"`
	IsNoop                             *bool                                       `json:"isNoop,omitempty"`
}

// NewBTBaseEntityAppearanceSettings1391 instantiates a new BTBaseEntityAppearanceSettings1391 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBaseEntityAppearanceSettings1391() *BTBaseEntityAppearanceSettings1391 {
	this := BTBaseEntityAppearanceSettings1391{}
	return &this
}

// NewBTBaseEntityAppearanceSettings1391WithDefaults instantiates a new BTBaseEntityAppearanceSettings1391 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBaseEntityAppearanceSettings1391WithDefaults() *BTBaseEntityAppearanceSettings1391 {
	this := BTBaseEntityAppearanceSettings1391{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTBaseEntityAppearanceSettings1391) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBaseEntityAppearanceSettings1391) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTBaseEntityAppearanceSettings1391) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTBaseEntityAppearanceSettings1391) SetBtType(v string) {
	o.BtType = &v
}

// GetColorIdToBaseEntityAppearanceEntry returns the ColorIdToBaseEntityAppearanceEntry field value if set, zero value otherwise.
func (o *BTBaseEntityAppearanceSettings1391) GetColorIdToBaseEntityAppearanceEntry() map[string]BTBaseEntityAppearanceEntry3607 {
	if o == nil || o.ColorIdToBaseEntityAppearanceEntry == nil {
		var ret map[string]BTBaseEntityAppearanceEntry3607
		return ret
	}
	return *o.ColorIdToBaseEntityAppearanceEntry
}

// GetColorIdToBaseEntityAppearanceEntryOk returns a tuple with the ColorIdToBaseEntityAppearanceEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBaseEntityAppearanceSettings1391) GetColorIdToBaseEntityAppearanceEntryOk() (*map[string]BTBaseEntityAppearanceEntry3607, bool) {
	if o == nil || o.ColorIdToBaseEntityAppearanceEntry == nil {
		return nil, false
	}
	return o.ColorIdToBaseEntityAppearanceEntry, true
}

// HasColorIdToBaseEntityAppearanceEntry returns a boolean if a field has been set.
func (o *BTBaseEntityAppearanceSettings1391) HasColorIdToBaseEntityAppearanceEntry() bool {
	if o != nil && o.ColorIdToBaseEntityAppearanceEntry != nil {
		return true
	}

	return false
}

// SetColorIdToBaseEntityAppearanceEntry gets a reference to the given map[string]BTBaseEntityAppearanceEntry3607 and assigns it to the ColorIdToBaseEntityAppearanceEntry field.
func (o *BTBaseEntityAppearanceSettings1391) SetColorIdToBaseEntityAppearanceEntry(v map[string]BTBaseEntityAppearanceEntry3607) {
	o.ColorIdToBaseEntityAppearanceEntry = &v
}

// GetIsNoop returns the IsNoop field value if set, zero value otherwise.
func (o *BTBaseEntityAppearanceSettings1391) GetIsNoop() bool {
	if o == nil || o.IsNoop == nil {
		var ret bool
		return ret
	}
	return *o.IsNoop
}

// GetIsNoopOk returns a tuple with the IsNoop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBaseEntityAppearanceSettings1391) GetIsNoopOk() (*bool, bool) {
	if o == nil || o.IsNoop == nil {
		return nil, false
	}
	return o.IsNoop, true
}

// HasIsNoop returns a boolean if a field has been set.
func (o *BTBaseEntityAppearanceSettings1391) HasIsNoop() bool {
	if o != nil && o.IsNoop != nil {
		return true
	}

	return false
}

// SetIsNoop gets a reference to the given bool and assigns it to the IsNoop field.
func (o *BTBaseEntityAppearanceSettings1391) SetIsNoop(v bool) {
	o.IsNoop = &v
}

func (o BTBaseEntityAppearanceSettings1391) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ColorIdToBaseEntityAppearanceEntry != nil {
		toSerialize["colorIdToBaseEntityAppearanceEntry"] = o.ColorIdToBaseEntityAppearanceEntry
	}
	if o.IsNoop != nil {
		toSerialize["isNoop"] = o.IsNoop
	}
	return json.Marshal(toSerialize)
}

type NullableBTBaseEntityAppearanceSettings1391 struct {
	value *BTBaseEntityAppearanceSettings1391
	isSet bool
}

func (v NullableBTBaseEntityAppearanceSettings1391) Get() *BTBaseEntityAppearanceSettings1391 {
	return v.value
}

func (v *NullableBTBaseEntityAppearanceSettings1391) Set(val *BTBaseEntityAppearanceSettings1391) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBaseEntityAppearanceSettings1391) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBaseEntityAppearanceSettings1391) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBaseEntityAppearanceSettings1391(val *BTBaseEntityAppearanceSettings1391) *NullableBTBaseEntityAppearanceSettings1391 {
	return &NullableBTBaseEntityAppearanceSettings1391{value: val, isSet: true}
}

func (v NullableBTBaseEntityAppearanceSettings1391) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBaseEntityAppearanceSettings1391) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
