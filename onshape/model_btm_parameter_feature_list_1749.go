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

// BTMParameterFeatureList1749 struct for BTMParameterFeatureList1749
type BTMParameterFeatureList1749 struct {
	BtType             *string  `json:"btType,omitempty"`
	ImportMicroversion *string  `json:"importMicroversion,omitempty"`
	NodeId             *string  `json:"nodeId,omitempty"`
	ParameterId        *string  `json:"parameterId,omitempty"`
	FeatureIds         []string `json:"featureIds,omitempty"`
}

// NewBTMParameterFeatureList1749 instantiates a new BTMParameterFeatureList1749 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMParameterFeatureList1749() *BTMParameterFeatureList1749 {
	this := BTMParameterFeatureList1749{}
	return &this
}

// NewBTMParameterFeatureList1749WithDefaults instantiates a new BTMParameterFeatureList1749 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMParameterFeatureList1749WithDefaults() *BTMParameterFeatureList1749 {
	this := BTMParameterFeatureList1749{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMParameterFeatureList1749) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterFeatureList1749) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMParameterFeatureList1749) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMParameterFeatureList1749) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMParameterFeatureList1749) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterFeatureList1749) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMParameterFeatureList1749) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMParameterFeatureList1749) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMParameterFeatureList1749) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterFeatureList1749) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMParameterFeatureList1749) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMParameterFeatureList1749) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTMParameterFeatureList1749) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterFeatureList1749) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTMParameterFeatureList1749) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTMParameterFeatureList1749) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetFeatureIds returns the FeatureIds field value if set, zero value otherwise.
func (o *BTMParameterFeatureList1749) GetFeatureIds() []string {
	if o == nil || o.FeatureIds == nil {
		var ret []string
		return ret
	}
	return o.FeatureIds
}

// GetFeatureIdsOk returns a tuple with the FeatureIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterFeatureList1749) GetFeatureIdsOk() ([]string, bool) {
	if o == nil || o.FeatureIds == nil {
		return nil, false
	}
	return o.FeatureIds, true
}

// HasFeatureIds returns a boolean if a field has been set.
func (o *BTMParameterFeatureList1749) HasFeatureIds() bool {
	if o != nil && o.FeatureIds != nil {
		return true
	}

	return false
}

// SetFeatureIds gets a reference to the given []string and assigns it to the FeatureIds field.
func (o *BTMParameterFeatureList1749) SetFeatureIds(v []string) {
	o.FeatureIds = v
}

func (o BTMParameterFeatureList1749) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.FeatureIds != nil {
		toSerialize["featureIds"] = o.FeatureIds
	}
	return json.Marshal(toSerialize)
}

type NullableBTMParameterFeatureList1749 struct {
	value *BTMParameterFeatureList1749
	isSet bool
}

func (v NullableBTMParameterFeatureList1749) Get() *BTMParameterFeatureList1749 {
	return v.value
}

func (v *NullableBTMParameterFeatureList1749) Set(val *BTMParameterFeatureList1749) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMParameterFeatureList1749) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMParameterFeatureList1749) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMParameterFeatureList1749(val *BTMParameterFeatureList1749) *NullableBTMParameterFeatureList1749 {
	return &NullableBTMParameterFeatureList1749{value: val, isSet: true}
}

func (v NullableBTMParameterFeatureList1749) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMParameterFeatureList1749) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
