/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.27960-e195de6ac56c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAssemblyInstanceOccurrenceInfo struct for BTAssemblyInstanceOccurrenceInfo
type BTAssemblyInstanceOccurrenceInfo struct {
	Occurrences []BTAssemblyOccurrenceInfo `json:"occurrences,omitempty"`
}

// NewBTAssemblyInstanceOccurrenceInfo instantiates a new BTAssemblyInstanceOccurrenceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyInstanceOccurrenceInfo() *BTAssemblyInstanceOccurrenceInfo {
	this := BTAssemblyInstanceOccurrenceInfo{}
	return &this
}

// NewBTAssemblyInstanceOccurrenceInfoWithDefaults instantiates a new BTAssemblyInstanceOccurrenceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyInstanceOccurrenceInfoWithDefaults() *BTAssemblyInstanceOccurrenceInfo {
	this := BTAssemblyInstanceOccurrenceInfo{}
	return &this
}

// GetOccurrences returns the Occurrences field value if set, zero value otherwise.
func (o *BTAssemblyInstanceOccurrenceInfo) GetOccurrences() []BTAssemblyOccurrenceInfo {
	if o == nil || o.Occurrences == nil {
		var ret []BTAssemblyOccurrenceInfo
		return ret
	}
	return o.Occurrences
}

// GetOccurrencesOk returns a tuple with the Occurrences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyInstanceOccurrenceInfo) GetOccurrencesOk() ([]BTAssemblyOccurrenceInfo, bool) {
	if o == nil || o.Occurrences == nil {
		return nil, false
	}
	return o.Occurrences, true
}

// HasOccurrences returns a boolean if a field has been set.
func (o *BTAssemblyInstanceOccurrenceInfo) HasOccurrences() bool {
	if o != nil && o.Occurrences != nil {
		return true
	}

	return false
}

// SetOccurrences gets a reference to the given []BTAssemblyOccurrenceInfo and assigns it to the Occurrences field.
func (o *BTAssemblyInstanceOccurrenceInfo) SetOccurrences(v []BTAssemblyOccurrenceInfo) {
	o.Occurrences = v
}

func (o BTAssemblyInstanceOccurrenceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Occurrences != nil {
		toSerialize["occurrences"] = o.Occurrences
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyInstanceOccurrenceInfo struct {
	value *BTAssemblyInstanceOccurrenceInfo
	isSet bool
}

func (v NullableBTAssemblyInstanceOccurrenceInfo) Get() *BTAssemblyInstanceOccurrenceInfo {
	return v.value
}

func (v *NullableBTAssemblyInstanceOccurrenceInfo) Set(val *BTAssemblyInstanceOccurrenceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyInstanceOccurrenceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyInstanceOccurrenceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyInstanceOccurrenceInfo(val *BTAssemblyInstanceOccurrenceInfo) *NullableBTAssemblyInstanceOccurrenceInfo {
	return &NullableBTAssemblyInstanceOccurrenceInfo{value: val, isSet: true}
}

func (v NullableBTAssemblyInstanceOccurrenceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyInstanceOccurrenceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
