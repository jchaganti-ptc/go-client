/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.28658-06d4d4923fc7
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTBillOfMaterialsTableRowMetadata1300 struct for BTBillOfMaterialsTableRowMetadata1300
type BTBillOfMaterialsTableRowMetadata1300 struct {
	BtType                  *string                                `json:"btType,omitempty"`
	CrossHighlightDataIfAny *BTTableAssemblyCrossHighlightData2675 `json:"crossHighlightDataIfAny,omitempty"`
	CrossHighlightData      *BTTableAssemblyCrossHighlightData2675 `json:"crossHighlightData,omitempty"`
}

// NewBTBillOfMaterialsTableRowMetadata1300 instantiates a new BTBillOfMaterialsTableRowMetadata1300 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBillOfMaterialsTableRowMetadata1300() *BTBillOfMaterialsTableRowMetadata1300 {
	this := BTBillOfMaterialsTableRowMetadata1300{}
	return &this
}

// NewBTBillOfMaterialsTableRowMetadata1300WithDefaults instantiates a new BTBillOfMaterialsTableRowMetadata1300 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBillOfMaterialsTableRowMetadata1300WithDefaults() *BTBillOfMaterialsTableRowMetadata1300 {
	this := BTBillOfMaterialsTableRowMetadata1300{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTableRowMetadata1300) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTableRowMetadata1300) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTableRowMetadata1300) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTBillOfMaterialsTableRowMetadata1300) SetBtType(v string) {
	o.BtType = &v
}

// GetCrossHighlightDataIfAny returns the CrossHighlightDataIfAny field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTableRowMetadata1300) GetCrossHighlightDataIfAny() BTTableAssemblyCrossHighlightData2675 {
	if o == nil || o.CrossHighlightDataIfAny == nil {
		var ret BTTableAssemblyCrossHighlightData2675
		return ret
	}
	return *o.CrossHighlightDataIfAny
}

// GetCrossHighlightDataIfAnyOk returns a tuple with the CrossHighlightDataIfAny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTableRowMetadata1300) GetCrossHighlightDataIfAnyOk() (*BTTableAssemblyCrossHighlightData2675, bool) {
	if o == nil || o.CrossHighlightDataIfAny == nil {
		return nil, false
	}
	return o.CrossHighlightDataIfAny, true
}

// HasCrossHighlightDataIfAny returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTableRowMetadata1300) HasCrossHighlightDataIfAny() bool {
	if o != nil && o.CrossHighlightDataIfAny != nil {
		return true
	}

	return false
}

// SetCrossHighlightDataIfAny gets a reference to the given BTTableAssemblyCrossHighlightData2675 and assigns it to the CrossHighlightDataIfAny field.
func (o *BTBillOfMaterialsTableRowMetadata1300) SetCrossHighlightDataIfAny(v BTTableAssemblyCrossHighlightData2675) {
	o.CrossHighlightDataIfAny = &v
}

// GetCrossHighlightData returns the CrossHighlightData field value if set, zero value otherwise.
func (o *BTBillOfMaterialsTableRowMetadata1300) GetCrossHighlightData() BTTableAssemblyCrossHighlightData2675 {
	if o == nil || o.CrossHighlightData == nil {
		var ret BTTableAssemblyCrossHighlightData2675
		return ret
	}
	return *o.CrossHighlightData
}

// GetCrossHighlightDataOk returns a tuple with the CrossHighlightData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBillOfMaterialsTableRowMetadata1300) GetCrossHighlightDataOk() (*BTTableAssemblyCrossHighlightData2675, bool) {
	if o == nil || o.CrossHighlightData == nil {
		return nil, false
	}
	return o.CrossHighlightData, true
}

// HasCrossHighlightData returns a boolean if a field has been set.
func (o *BTBillOfMaterialsTableRowMetadata1300) HasCrossHighlightData() bool {
	if o != nil && o.CrossHighlightData != nil {
		return true
	}

	return false
}

// SetCrossHighlightData gets a reference to the given BTTableAssemblyCrossHighlightData2675 and assigns it to the CrossHighlightData field.
func (o *BTBillOfMaterialsTableRowMetadata1300) SetCrossHighlightData(v BTTableAssemblyCrossHighlightData2675) {
	o.CrossHighlightData = &v
}

func (o BTBillOfMaterialsTableRowMetadata1300) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.CrossHighlightDataIfAny != nil {
		toSerialize["crossHighlightDataIfAny"] = o.CrossHighlightDataIfAny
	}
	if o.CrossHighlightData != nil {
		toSerialize["crossHighlightData"] = o.CrossHighlightData
	}
	return json.Marshal(toSerialize)
}

type NullableBTBillOfMaterialsTableRowMetadata1300 struct {
	value *BTBillOfMaterialsTableRowMetadata1300
	isSet bool
}

func (v NullableBTBillOfMaterialsTableRowMetadata1300) Get() *BTBillOfMaterialsTableRowMetadata1300 {
	return v.value
}

func (v *NullableBTBillOfMaterialsTableRowMetadata1300) Set(val *BTBillOfMaterialsTableRowMetadata1300) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBillOfMaterialsTableRowMetadata1300) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBillOfMaterialsTableRowMetadata1300) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBillOfMaterialsTableRowMetadata1300(val *BTBillOfMaterialsTableRowMetadata1300) *NullableBTBillOfMaterialsTableRowMetadata1300 {
	return &NullableBTBillOfMaterialsTableRowMetadata1300{value: val, isSet: true}
}

func (v NullableBTBillOfMaterialsTableRowMetadata1300) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBillOfMaterialsTableRowMetadata1300) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
