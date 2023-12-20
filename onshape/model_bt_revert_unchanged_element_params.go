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

// BTRevertUnchangedElementParams struct for BTRevertUnchangedElementParams
type BTRevertUnchangedElementParams struct {
	Configuration *string  `json:"configuration,omitempty"`
	ElementId     *string  `json:"elementId,omitempty"`
	ReferenceIds  []string `json:"referenceIds,omitempty"`
}

// NewBTRevertUnchangedElementParams instantiates a new BTRevertUnchangedElementParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTRevertUnchangedElementParams() *BTRevertUnchangedElementParams {
	this := BTRevertUnchangedElementParams{}
	return &this
}

// NewBTRevertUnchangedElementParamsWithDefaults instantiates a new BTRevertUnchangedElementParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTRevertUnchangedElementParamsWithDefaults() *BTRevertUnchangedElementParams {
	this := BTRevertUnchangedElementParams{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *BTRevertUnchangedElementParams) GetConfiguration() string {
	if o == nil || o.Configuration == nil {
		var ret string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevertUnchangedElementParams) GetConfigurationOk() (*string, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *BTRevertUnchangedElementParams) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given string and assigns it to the Configuration field.
func (o *BTRevertUnchangedElementParams) SetConfiguration(v string) {
	o.Configuration = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTRevertUnchangedElementParams) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevertUnchangedElementParams) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTRevertUnchangedElementParams) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTRevertUnchangedElementParams) SetElementId(v string) {
	o.ElementId = &v
}

// GetReferenceIds returns the ReferenceIds field value if set, zero value otherwise.
func (o *BTRevertUnchangedElementParams) GetReferenceIds() []string {
	if o == nil || o.ReferenceIds == nil {
		var ret []string
		return ret
	}
	return o.ReferenceIds
}

// GetReferenceIdsOk returns a tuple with the ReferenceIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevertUnchangedElementParams) GetReferenceIdsOk() ([]string, bool) {
	if o == nil || o.ReferenceIds == nil {
		return nil, false
	}
	return o.ReferenceIds, true
}

// HasReferenceIds returns a boolean if a field has been set.
func (o *BTRevertUnchangedElementParams) HasReferenceIds() bool {
	if o != nil && o.ReferenceIds != nil {
		return true
	}

	return false
}

// SetReferenceIds gets a reference to the given []string and assigns it to the ReferenceIds field.
func (o *BTRevertUnchangedElementParams) SetReferenceIds(v []string) {
	o.ReferenceIds = v
}

func (o BTRevertUnchangedElementParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.ReferenceIds != nil {
		toSerialize["referenceIds"] = o.ReferenceIds
	}
	return json.Marshal(toSerialize)
}

type NullableBTRevertUnchangedElementParams struct {
	value *BTRevertUnchangedElementParams
	isSet bool
}

func (v NullableBTRevertUnchangedElementParams) Get() *BTRevertUnchangedElementParams {
	return v.value
}

func (v *NullableBTRevertUnchangedElementParams) Set(val *BTRevertUnchangedElementParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTRevertUnchangedElementParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTRevertUnchangedElementParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTRevertUnchangedElementParams(val *BTRevertUnchangedElementParams) *NullableBTRevertUnchangedElementParams {
	return &NullableBTRevertUnchangedElementParams{value: val, isSet: true}
}

func (v NullableBTRevertUnchangedElementParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTRevertUnchangedElementParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
