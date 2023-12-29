/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.28126-700645382199
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTStringFormatBlockPattern1755 struct for BTStringFormatBlockPattern1755
type BTStringFormatBlockPattern1755 struct {
	BtType                        *string `json:"btType,omitempty"`
	ErrorMessage                  *string `json:"errorMessage,omitempty"`
	ShouldResetValueWhenConfirmed *bool   `json:"shouldResetValueWhenConfirmed,omitempty"`
	RegExpToBlock                 *string `json:"regExpToBlock,omitempty"`
}

// NewBTStringFormatBlockPattern1755 instantiates a new BTStringFormatBlockPattern1755 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTStringFormatBlockPattern1755() *BTStringFormatBlockPattern1755 {
	this := BTStringFormatBlockPattern1755{}
	return &this
}

// NewBTStringFormatBlockPattern1755WithDefaults instantiates a new BTStringFormatBlockPattern1755 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTStringFormatBlockPattern1755WithDefaults() *BTStringFormatBlockPattern1755 {
	this := BTStringFormatBlockPattern1755{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTStringFormatBlockPattern1755) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringFormatBlockPattern1755) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTStringFormatBlockPattern1755) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTStringFormatBlockPattern1755) SetBtType(v string) {
	o.BtType = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *BTStringFormatBlockPattern1755) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringFormatBlockPattern1755) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *BTStringFormatBlockPattern1755) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *BTStringFormatBlockPattern1755) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetShouldResetValueWhenConfirmed returns the ShouldResetValueWhenConfirmed field value if set, zero value otherwise.
func (o *BTStringFormatBlockPattern1755) GetShouldResetValueWhenConfirmed() bool {
	if o == nil || o.ShouldResetValueWhenConfirmed == nil {
		var ret bool
		return ret
	}
	return *o.ShouldResetValueWhenConfirmed
}

// GetShouldResetValueWhenConfirmedOk returns a tuple with the ShouldResetValueWhenConfirmed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringFormatBlockPattern1755) GetShouldResetValueWhenConfirmedOk() (*bool, bool) {
	if o == nil || o.ShouldResetValueWhenConfirmed == nil {
		return nil, false
	}
	return o.ShouldResetValueWhenConfirmed, true
}

// HasShouldResetValueWhenConfirmed returns a boolean if a field has been set.
func (o *BTStringFormatBlockPattern1755) HasShouldResetValueWhenConfirmed() bool {
	if o != nil && o.ShouldResetValueWhenConfirmed != nil {
		return true
	}

	return false
}

// SetShouldResetValueWhenConfirmed gets a reference to the given bool and assigns it to the ShouldResetValueWhenConfirmed field.
func (o *BTStringFormatBlockPattern1755) SetShouldResetValueWhenConfirmed(v bool) {
	o.ShouldResetValueWhenConfirmed = &v
}

// GetRegExpToBlock returns the RegExpToBlock field value if set, zero value otherwise.
func (o *BTStringFormatBlockPattern1755) GetRegExpToBlock() string {
	if o == nil || o.RegExpToBlock == nil {
		var ret string
		return ret
	}
	return *o.RegExpToBlock
}

// GetRegExpToBlockOk returns a tuple with the RegExpToBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTStringFormatBlockPattern1755) GetRegExpToBlockOk() (*string, bool) {
	if o == nil || o.RegExpToBlock == nil {
		return nil, false
	}
	return o.RegExpToBlock, true
}

// HasRegExpToBlock returns a boolean if a field has been set.
func (o *BTStringFormatBlockPattern1755) HasRegExpToBlock() bool {
	if o != nil && o.RegExpToBlock != nil {
		return true
	}

	return false
}

// SetRegExpToBlock gets a reference to the given string and assigns it to the RegExpToBlock field.
func (o *BTStringFormatBlockPattern1755) SetRegExpToBlock(v string) {
	o.RegExpToBlock = &v
}

func (o BTStringFormatBlockPattern1755) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.ShouldResetValueWhenConfirmed != nil {
		toSerialize["shouldResetValueWhenConfirmed"] = o.ShouldResetValueWhenConfirmed
	}
	if o.RegExpToBlock != nil {
		toSerialize["regExpToBlock"] = o.RegExpToBlock
	}
	return json.Marshal(toSerialize)
}

type NullableBTStringFormatBlockPattern1755 struct {
	value *BTStringFormatBlockPattern1755
	isSet bool
}

func (v NullableBTStringFormatBlockPattern1755) Get() *BTStringFormatBlockPattern1755 {
	return v.value
}

func (v *NullableBTStringFormatBlockPattern1755) Set(val *BTStringFormatBlockPattern1755) {
	v.value = val
	v.isSet = true
}

func (v NullableBTStringFormatBlockPattern1755) IsSet() bool {
	return v.isSet
}

func (v *NullableBTStringFormatBlockPattern1755) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTStringFormatBlockPattern1755(val *BTStringFormatBlockPattern1755) *NullableBTStringFormatBlockPattern1755 {
	return &NullableBTStringFormatBlockPattern1755{value: val, isSet: true}
}

func (v NullableBTStringFormatBlockPattern1755) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTStringFormatBlockPattern1755) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
