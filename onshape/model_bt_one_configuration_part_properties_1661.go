/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6633-e6f3d20f07a4
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTOneConfigurationPartProperties1661 struct for BTOneConfigurationPartProperties1661
type BTOneConfigurationPartProperties1661 struct {
	BtType        *string                   `json:"btType,omitempty"`
	Configuration *map[string]BTFSValue1888 `json:"configuration,omitempty"`
	Merged        *BTOnePartProperties230   `json:"merged,omitempty"`
	NodeId        *string                   `json:"nodeId,omitempty"`
	Properties    []BTOnePartProperties230  `json:"properties,omitempty"`
	PropertyIds   []string                  `json:"propertyIds,omitempty"`
}

// NewBTOneConfigurationPartProperties1661 instantiates a new BTOneConfigurationPartProperties1661 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTOneConfigurationPartProperties1661() *BTOneConfigurationPartProperties1661 {
	this := BTOneConfigurationPartProperties1661{}
	return &this
}

// NewBTOneConfigurationPartProperties1661WithDefaults instantiates a new BTOneConfigurationPartProperties1661 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTOneConfigurationPartProperties1661WithDefaults() *BTOneConfigurationPartProperties1661 {
	this := BTOneConfigurationPartProperties1661{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTOneConfigurationPartProperties1661) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOneConfigurationPartProperties1661) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTOneConfigurationPartProperties1661) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTOneConfigurationPartProperties1661) SetBtType(v string) {
	o.BtType = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *BTOneConfigurationPartProperties1661) GetConfiguration() map[string]BTFSValue1888 {
	if o == nil || o.Configuration == nil {
		var ret map[string]BTFSValue1888
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOneConfigurationPartProperties1661) GetConfigurationOk() (*map[string]BTFSValue1888, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *BTOneConfigurationPartProperties1661) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]BTFSValue1888 and assigns it to the Configuration field.
func (o *BTOneConfigurationPartProperties1661) SetConfiguration(v map[string]BTFSValue1888) {
	o.Configuration = &v
}

// GetMerged returns the Merged field value if set, zero value otherwise.
func (o *BTOneConfigurationPartProperties1661) GetMerged() BTOnePartProperties230 {
	if o == nil || o.Merged == nil {
		var ret BTOnePartProperties230
		return ret
	}
	return *o.Merged
}

// GetMergedOk returns a tuple with the Merged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOneConfigurationPartProperties1661) GetMergedOk() (*BTOnePartProperties230, bool) {
	if o == nil || o.Merged == nil {
		return nil, false
	}
	return o.Merged, true
}

// HasMerged returns a boolean if a field has been set.
func (o *BTOneConfigurationPartProperties1661) HasMerged() bool {
	if o != nil && o.Merged != nil {
		return true
	}

	return false
}

// SetMerged gets a reference to the given BTOnePartProperties230 and assigns it to the Merged field.
func (o *BTOneConfigurationPartProperties1661) SetMerged(v BTOnePartProperties230) {
	o.Merged = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTOneConfigurationPartProperties1661) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOneConfigurationPartProperties1661) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTOneConfigurationPartProperties1661) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTOneConfigurationPartProperties1661) SetNodeId(v string) {
	o.NodeId = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *BTOneConfigurationPartProperties1661) GetProperties() []BTOnePartProperties230 {
	if o == nil || o.Properties == nil {
		var ret []BTOnePartProperties230
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOneConfigurationPartProperties1661) GetPropertiesOk() ([]BTOnePartProperties230, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *BTOneConfigurationPartProperties1661) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []BTOnePartProperties230 and assigns it to the Properties field.
func (o *BTOneConfigurationPartProperties1661) SetProperties(v []BTOnePartProperties230) {
	o.Properties = v
}

// GetPropertyIds returns the PropertyIds field value if set, zero value otherwise.
func (o *BTOneConfigurationPartProperties1661) GetPropertyIds() []string {
	if o == nil || o.PropertyIds == nil {
		var ret []string
		return ret
	}
	return o.PropertyIds
}

// GetPropertyIdsOk returns a tuple with the PropertyIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOneConfigurationPartProperties1661) GetPropertyIdsOk() ([]string, bool) {
	if o == nil || o.PropertyIds == nil {
		return nil, false
	}
	return o.PropertyIds, true
}

// HasPropertyIds returns a boolean if a field has been set.
func (o *BTOneConfigurationPartProperties1661) HasPropertyIds() bool {
	if o != nil && o.PropertyIds != nil {
		return true
	}

	return false
}

// SetPropertyIds gets a reference to the given []string and assigns it to the PropertyIds field.
func (o *BTOneConfigurationPartProperties1661) SetPropertyIds(v []string) {
	o.PropertyIds = v
}

func (o BTOneConfigurationPartProperties1661) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.Merged != nil {
		toSerialize["merged"] = o.Merged
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.PropertyIds != nil {
		toSerialize["propertyIds"] = o.PropertyIds
	}
	return json.Marshal(toSerialize)
}

type NullableBTOneConfigurationPartProperties1661 struct {
	value *BTOneConfigurationPartProperties1661
	isSet bool
}

func (v NullableBTOneConfigurationPartProperties1661) Get() *BTOneConfigurationPartProperties1661 {
	return v.value
}

func (v *NullableBTOneConfigurationPartProperties1661) Set(val *BTOneConfigurationPartProperties1661) {
	v.value = val
	v.isSet = true
}

func (v NullableBTOneConfigurationPartProperties1661) IsSet() bool {
	return v.isSet
}

func (v *NullableBTOneConfigurationPartProperties1661) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTOneConfigurationPartProperties1661(val *BTOneConfigurationPartProperties1661) *NullableBTOneConfigurationPartProperties1661 {
	return &NullableBTOneConfigurationPartProperties1661{value: val, isSet: true}
}

func (v NullableBTOneConfigurationPartProperties1661) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTOneConfigurationPartProperties1661) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
