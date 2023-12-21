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

// BTFSTableRowMetadata2262 struct for BTFSTableRowMetadata2262
type BTFSTableRowMetadata2262 struct {
	BtType                  *string                            `json:"btType,omitempty"`
	CrossHighlightDataIfAny *BTTableBaseCrossHighlightData2609 `json:"crossHighlightDataIfAny,omitempty"`
	Callout                 *string                            `json:"callout,omitempty"`
	CrossHighlightData      *BTTableBaseCrossHighlightData2609 `json:"crossHighlightData,omitempty"`
}

// NewBTFSTableRowMetadata2262 instantiates a new BTFSTableRowMetadata2262 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFSTableRowMetadata2262() *BTFSTableRowMetadata2262 {
	this := BTFSTableRowMetadata2262{}
	return &this
}

// NewBTFSTableRowMetadata2262WithDefaults instantiates a new BTFSTableRowMetadata2262 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFSTableRowMetadata2262WithDefaults() *BTFSTableRowMetadata2262 {
	this := BTFSTableRowMetadata2262{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTFSTableRowMetadata2262) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSTableRowMetadata2262) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTFSTableRowMetadata2262) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTFSTableRowMetadata2262) SetBtType(v string) {
	o.BtType = &v
}

// GetCrossHighlightDataIfAny returns the CrossHighlightDataIfAny field value if set, zero value otherwise.
func (o *BTFSTableRowMetadata2262) GetCrossHighlightDataIfAny() BTTableBaseCrossHighlightData2609 {
	if o == nil || o.CrossHighlightDataIfAny == nil {
		var ret BTTableBaseCrossHighlightData2609
		return ret
	}
	return *o.CrossHighlightDataIfAny
}

// GetCrossHighlightDataIfAnyOk returns a tuple with the CrossHighlightDataIfAny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSTableRowMetadata2262) GetCrossHighlightDataIfAnyOk() (*BTTableBaseCrossHighlightData2609, bool) {
	if o == nil || o.CrossHighlightDataIfAny == nil {
		return nil, false
	}
	return o.CrossHighlightDataIfAny, true
}

// HasCrossHighlightDataIfAny returns a boolean if a field has been set.
func (o *BTFSTableRowMetadata2262) HasCrossHighlightDataIfAny() bool {
	if o != nil && o.CrossHighlightDataIfAny != nil {
		return true
	}

	return false
}

// SetCrossHighlightDataIfAny gets a reference to the given BTTableBaseCrossHighlightData2609 and assigns it to the CrossHighlightDataIfAny field.
func (o *BTFSTableRowMetadata2262) SetCrossHighlightDataIfAny(v BTTableBaseCrossHighlightData2609) {
	o.CrossHighlightDataIfAny = &v
}

// GetCallout returns the Callout field value if set, zero value otherwise.
func (o *BTFSTableRowMetadata2262) GetCallout() string {
	if o == nil || o.Callout == nil {
		var ret string
		return ret
	}
	return *o.Callout
}

// GetCalloutOk returns a tuple with the Callout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSTableRowMetadata2262) GetCalloutOk() (*string, bool) {
	if o == nil || o.Callout == nil {
		return nil, false
	}
	return o.Callout, true
}

// HasCallout returns a boolean if a field has been set.
func (o *BTFSTableRowMetadata2262) HasCallout() bool {
	if o != nil && o.Callout != nil {
		return true
	}

	return false
}

// SetCallout gets a reference to the given string and assigns it to the Callout field.
func (o *BTFSTableRowMetadata2262) SetCallout(v string) {
	o.Callout = &v
}

// GetCrossHighlightData returns the CrossHighlightData field value if set, zero value otherwise.
func (o *BTFSTableRowMetadata2262) GetCrossHighlightData() BTTableBaseCrossHighlightData2609 {
	if o == nil || o.CrossHighlightData == nil {
		var ret BTTableBaseCrossHighlightData2609
		return ret
	}
	return *o.CrossHighlightData
}

// GetCrossHighlightDataOk returns a tuple with the CrossHighlightData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFSTableRowMetadata2262) GetCrossHighlightDataOk() (*BTTableBaseCrossHighlightData2609, bool) {
	if o == nil || o.CrossHighlightData == nil {
		return nil, false
	}
	return o.CrossHighlightData, true
}

// HasCrossHighlightData returns a boolean if a field has been set.
func (o *BTFSTableRowMetadata2262) HasCrossHighlightData() bool {
	if o != nil && o.CrossHighlightData != nil {
		return true
	}

	return false
}

// SetCrossHighlightData gets a reference to the given BTTableBaseCrossHighlightData2609 and assigns it to the CrossHighlightData field.
func (o *BTFSTableRowMetadata2262) SetCrossHighlightData(v BTTableBaseCrossHighlightData2609) {
	o.CrossHighlightData = &v
}

func (o BTFSTableRowMetadata2262) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.CrossHighlightDataIfAny != nil {
		toSerialize["crossHighlightDataIfAny"] = o.CrossHighlightDataIfAny
	}
	if o.Callout != nil {
		toSerialize["callout"] = o.Callout
	}
	if o.CrossHighlightData != nil {
		toSerialize["crossHighlightData"] = o.CrossHighlightData
	}
	return json.Marshal(toSerialize)
}

type NullableBTFSTableRowMetadata2262 struct {
	value *BTFSTableRowMetadata2262
	isSet bool
}

func (v NullableBTFSTableRowMetadata2262) Get() *BTFSTableRowMetadata2262 {
	return v.value
}

func (v *NullableBTFSTableRowMetadata2262) Set(val *BTFSTableRowMetadata2262) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFSTableRowMetadata2262) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFSTableRowMetadata2262) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFSTableRowMetadata2262(val *BTFSTableRowMetadata2262) *NullableBTFSTableRowMetadata2262 {
	return &NullableBTFSTableRowMetadata2262{value: val, isSet: true}
}

func (v NullableBTFSTableRowMetadata2262) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFSTableRowMetadata2262) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
