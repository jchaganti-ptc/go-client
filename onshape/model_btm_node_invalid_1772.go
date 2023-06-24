/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.165.17983-3c4f6e999748
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMNodeInvalid1772 struct for BTMNodeInvalid1772
type BTMNodeInvalid1772 struct {
	BtType             *string `json:"btType,omitempty"`
	ImportMicroversion *string `json:"importMicroversion,omitempty"`
	NodeId             *string `json:"nodeId,omitempty"`
}

// NewBTMNodeInvalid1772 instantiates a new BTMNodeInvalid1772 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMNodeInvalid1772() *BTMNodeInvalid1772 {
	this := BTMNodeInvalid1772{}
	return &this
}

// NewBTMNodeInvalid1772WithDefaults instantiates a new BTMNodeInvalid1772 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMNodeInvalid1772WithDefaults() *BTMNodeInvalid1772 {
	this := BTMNodeInvalid1772{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMNodeInvalid1772) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNodeInvalid1772) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMNodeInvalid1772) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMNodeInvalid1772) SetBtType(v string) {
	o.BtType = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMNodeInvalid1772) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNodeInvalid1772) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMNodeInvalid1772) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMNodeInvalid1772) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMNodeInvalid1772) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMNodeInvalid1772) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMNodeInvalid1772) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMNodeInvalid1772) SetNodeId(v string) {
	o.NodeId = &v
}

func (o BTMNodeInvalid1772) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableBTMNodeInvalid1772 struct {
	value *BTMNodeInvalid1772
	isSet bool
}

func (v NullableBTMNodeInvalid1772) Get() *BTMNodeInvalid1772 {
	return v.value
}

func (v *NullableBTMNodeInvalid1772) Set(val *BTMNodeInvalid1772) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMNodeInvalid1772) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMNodeInvalid1772) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMNodeInvalid1772(val *BTMNodeInvalid1772) *NullableBTMNodeInvalid1772 {
	return &NullableBTMNodeInvalid1772{value: val, isSet: true}
}

func (v NullableBTMNodeInvalid1772) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMNodeInvalid1772) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
