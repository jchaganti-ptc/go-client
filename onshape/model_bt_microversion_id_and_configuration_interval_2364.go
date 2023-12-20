/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.27881-5dbbe8053cdf
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMicroversionIdAndConfigurationInterval2364 struct for BTMicroversionIdAndConfigurationInterval2364
type BTMicroversionIdAndConfigurationInterval2364 struct {
	BtType                 *string                               `json:"btType,omitempty"`
	From                   *BTMicroversionIdAndConfiguration2338 `json:"from,omitempty"`
	MicroversionIdInterval *BTMicroversionIdInterval367          `json:"microversionIdInterval,omitempty"`
	To                     *BTMicroversionIdAndConfiguration2338 `json:"to,omitempty"`
}

// NewBTMicroversionIdAndConfigurationInterval2364 instantiates a new BTMicroversionIdAndConfigurationInterval2364 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMicroversionIdAndConfigurationInterval2364() *BTMicroversionIdAndConfigurationInterval2364 {
	this := BTMicroversionIdAndConfigurationInterval2364{}
	return &this
}

// NewBTMicroversionIdAndConfigurationInterval2364WithDefaults instantiates a new BTMicroversionIdAndConfigurationInterval2364 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMicroversionIdAndConfigurationInterval2364WithDefaults() *BTMicroversionIdAndConfigurationInterval2364 {
	this := BTMicroversionIdAndConfigurationInterval2364{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMicroversionIdAndConfigurationInterval2364) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMicroversionIdAndConfigurationInterval2364) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMicroversionIdAndConfigurationInterval2364) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMicroversionIdAndConfigurationInterval2364) SetBtType(v string) {
	o.BtType = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *BTMicroversionIdAndConfigurationInterval2364) GetFrom() BTMicroversionIdAndConfiguration2338 {
	if o == nil || o.From == nil {
		var ret BTMicroversionIdAndConfiguration2338
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMicroversionIdAndConfigurationInterval2364) GetFromOk() (*BTMicroversionIdAndConfiguration2338, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *BTMicroversionIdAndConfigurationInterval2364) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given BTMicroversionIdAndConfiguration2338 and assigns it to the From field.
func (o *BTMicroversionIdAndConfigurationInterval2364) SetFrom(v BTMicroversionIdAndConfiguration2338) {
	o.From = &v
}

// GetMicroversionIdInterval returns the MicroversionIdInterval field value if set, zero value otherwise.
func (o *BTMicroversionIdAndConfigurationInterval2364) GetMicroversionIdInterval() BTMicroversionIdInterval367 {
	if o == nil || o.MicroversionIdInterval == nil {
		var ret BTMicroversionIdInterval367
		return ret
	}
	return *o.MicroversionIdInterval
}

// GetMicroversionIdIntervalOk returns a tuple with the MicroversionIdInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMicroversionIdAndConfigurationInterval2364) GetMicroversionIdIntervalOk() (*BTMicroversionIdInterval367, bool) {
	if o == nil || o.MicroversionIdInterval == nil {
		return nil, false
	}
	return o.MicroversionIdInterval, true
}

// HasMicroversionIdInterval returns a boolean if a field has been set.
func (o *BTMicroversionIdAndConfigurationInterval2364) HasMicroversionIdInterval() bool {
	if o != nil && o.MicroversionIdInterval != nil {
		return true
	}

	return false
}

// SetMicroversionIdInterval gets a reference to the given BTMicroversionIdInterval367 and assigns it to the MicroversionIdInterval field.
func (o *BTMicroversionIdAndConfigurationInterval2364) SetMicroversionIdInterval(v BTMicroversionIdInterval367) {
	o.MicroversionIdInterval = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *BTMicroversionIdAndConfigurationInterval2364) GetTo() BTMicroversionIdAndConfiguration2338 {
	if o == nil || o.To == nil {
		var ret BTMicroversionIdAndConfiguration2338
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMicroversionIdAndConfigurationInterval2364) GetToOk() (*BTMicroversionIdAndConfiguration2338, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *BTMicroversionIdAndConfigurationInterval2364) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given BTMicroversionIdAndConfiguration2338 and assigns it to the To field.
func (o *BTMicroversionIdAndConfigurationInterval2364) SetTo(v BTMicroversionIdAndConfiguration2338) {
	o.To = &v
}

func (o BTMicroversionIdAndConfigurationInterval2364) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.MicroversionIdInterval != nil {
		toSerialize["microversionIdInterval"] = o.MicroversionIdInterval
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableBTMicroversionIdAndConfigurationInterval2364 struct {
	value *BTMicroversionIdAndConfigurationInterval2364
	isSet bool
}

func (v NullableBTMicroversionIdAndConfigurationInterval2364) Get() *BTMicroversionIdAndConfigurationInterval2364 {
	return v.value
}

func (v *NullableBTMicroversionIdAndConfigurationInterval2364) Set(val *BTMicroversionIdAndConfigurationInterval2364) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMicroversionIdAndConfigurationInterval2364) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMicroversionIdAndConfigurationInterval2364) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMicroversionIdAndConfigurationInterval2364(val *BTMicroversionIdAndConfigurationInterval2364) *NullableBTMicroversionIdAndConfigurationInterval2364 {
	return &NullableBTMicroversionIdAndConfigurationInterval2364{value: val, isSet: true}
}

func (v NullableBTMicroversionIdAndConfigurationInterval2364) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMicroversionIdAndConfigurationInterval2364) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
