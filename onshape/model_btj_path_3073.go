/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.167.19303-3cbf47a47fe4
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTJPath3073 Identifies a value in the json data to be operated upon.
type BTJPath3073 struct {
	BtType *string              `json:"btType,omitempty"`
	Path   []BTJPathElement2297 `json:"path,omitempty"`
	// Either empty (root) or the nodeId of a node in the tree.
	StartNode string `json:"startNode"`
}

// NewBTJPath3073 instantiates a new BTJPath3073 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTJPath3073(startNode string) *BTJPath3073 {
	this := BTJPath3073{}
	this.StartNode = startNode
	return &this
}

// NewBTJPath3073WithDefaults instantiates a new BTJPath3073 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTJPath3073WithDefaults() *BTJPath3073 {
	this := BTJPath3073{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTJPath3073) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTJPath3073) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTJPath3073) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTJPath3073) SetBtType(v string) {
	o.BtType = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *BTJPath3073) GetPath() []BTJPathElement2297 {
	if o == nil || o.Path == nil {
		var ret []BTJPathElement2297
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTJPath3073) GetPathOk() ([]BTJPathElement2297, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *BTJPath3073) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given []BTJPathElement2297 and assigns it to the Path field.
func (o *BTJPath3073) SetPath(v []BTJPathElement2297) {
	o.Path = v
}

// GetStartNode returns the StartNode field value
func (o *BTJPath3073) GetStartNode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartNode
}

// GetStartNodeOk returns a tuple with the StartNode field value
// and a boolean to check if the value has been set.
func (o *BTJPath3073) GetStartNodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartNode, true
}

// SetStartNode sets field value
func (o *BTJPath3073) SetStartNode(v string) {
	o.StartNode = v
}

func (o BTJPath3073) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["startNode"] = o.StartNode
	}
	return json.Marshal(toSerialize)
}

type NullableBTJPath3073 struct {
	value *BTJPath3073
	isSet bool
}

func (v NullableBTJPath3073) Get() *BTJPath3073 {
	return v.value
}

func (v *NullableBTJPath3073) Set(val *BTJPath3073) {
	v.value = val
	v.isSet = true
}

func (v NullableBTJPath3073) IsSet() bool {
	return v.isSet
}

func (v *NullableBTJPath3073) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTJPath3073(val *BTJPath3073) *NullableBTJPath3073 {
	return &NullableBTJPath3073{value: val, isSet: true}
}

func (v NullableBTJPath3073) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTJPath3073) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
