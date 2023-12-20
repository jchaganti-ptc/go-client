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

// BTIdTranslationResultInfo struct for BTIdTranslationResultInfo
type BTIdTranslationResultInfo struct {
	Source *string  `json:"source,omitempty"`
	Status *string  `json:"status,omitempty"`
	Target []string `json:"target,omitempty"`
}

// NewBTIdTranslationResultInfo instantiates a new BTIdTranslationResultInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTIdTranslationResultInfo() *BTIdTranslationResultInfo {
	this := BTIdTranslationResultInfo{}
	return &this
}

// NewBTIdTranslationResultInfoWithDefaults instantiates a new BTIdTranslationResultInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTIdTranslationResultInfoWithDefaults() *BTIdTranslationResultInfo {
	this := BTIdTranslationResultInfo{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *BTIdTranslationResultInfo) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTIdTranslationResultInfo) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *BTIdTranslationResultInfo) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *BTIdTranslationResultInfo) SetSource(v string) {
	o.Source = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BTIdTranslationResultInfo) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTIdTranslationResultInfo) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BTIdTranslationResultInfo) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BTIdTranslationResultInfo) SetStatus(v string) {
	o.Status = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *BTIdTranslationResultInfo) GetTarget() []string {
	if o == nil || o.Target == nil {
		var ret []string
		return ret
	}
	return o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTIdTranslationResultInfo) GetTargetOk() ([]string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *BTIdTranslationResultInfo) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given []string and assigns it to the Target field.
func (o *BTIdTranslationResultInfo) SetTarget(v []string) {
	o.Target = v
}

func (o BTIdTranslationResultInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableBTIdTranslationResultInfo struct {
	value *BTIdTranslationResultInfo
	isSet bool
}

func (v NullableBTIdTranslationResultInfo) Get() *BTIdTranslationResultInfo {
	return v.value
}

func (v *NullableBTIdTranslationResultInfo) Set(val *BTIdTranslationResultInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTIdTranslationResultInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTIdTranslationResultInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTIdTranslationResultInfo(val *BTIdTranslationResultInfo) *NullableBTIdTranslationResultInfo {
	return &NullableBTIdTranslationResultInfo{value: val, isSet: true}
}

func (v NullableBTIdTranslationResultInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTIdTranslationResultInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
