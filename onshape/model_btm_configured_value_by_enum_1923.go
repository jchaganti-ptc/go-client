/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.154.6621-03f431879e4b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMConfiguredValueByEnum1923 struct for BTMConfiguredValueByEnum1923
type BTMConfiguredValueByEnum1923 struct {
	BtType                   *string        `json:"btType,omitempty"`
	ConfigurationValueString *string        `json:"configurationValueString,omitempty"`
	ImportMicroversion       *string        `json:"importMicroversion,omitempty"`
	NodeId                   *string        `json:"nodeId,omitempty"`
	Value                    *BTMParameter1 `json:"value,omitempty"`
	EnumName                 *string        `json:"enumName,omitempty"`
	EnumValue                *string        `json:"enumValue,omitempty"`
	Namespace                *string        `json:"namespace,omitempty"`
}

// NewBTMConfiguredValueByEnum1923 instantiates a new BTMConfiguredValueByEnum1923 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMConfiguredValueByEnum1923() *BTMConfiguredValueByEnum1923 {
	this := BTMConfiguredValueByEnum1923{}
	return &this
}

// NewBTMConfiguredValueByEnum1923WithDefaults instantiates a new BTMConfiguredValueByEnum1923 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMConfiguredValueByEnum1923WithDefaults() *BTMConfiguredValueByEnum1923 {
	this := BTMConfiguredValueByEnum1923{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMConfiguredValueByEnum1923) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValueByEnum1923) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMConfiguredValueByEnum1923) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMConfiguredValueByEnum1923) SetBtType(v string) {
	o.BtType = &v
}

// GetConfigurationValueString returns the ConfigurationValueString field value if set, zero value otherwise.
func (o *BTMConfiguredValueByEnum1923) GetConfigurationValueString() string {
	if o == nil || o.ConfigurationValueString == nil {
		var ret string
		return ret
	}
	return *o.ConfigurationValueString
}

// GetConfigurationValueStringOk returns a tuple with the ConfigurationValueString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValueByEnum1923) GetConfigurationValueStringOk() (*string, bool) {
	if o == nil || o.ConfigurationValueString == nil {
		return nil, false
	}
	return o.ConfigurationValueString, true
}

// HasConfigurationValueString returns a boolean if a field has been set.
func (o *BTMConfiguredValueByEnum1923) HasConfigurationValueString() bool {
	if o != nil && o.ConfigurationValueString != nil {
		return true
	}

	return false
}

// SetConfigurationValueString gets a reference to the given string and assigns it to the ConfigurationValueString field.
func (o *BTMConfiguredValueByEnum1923) SetConfigurationValueString(v string) {
	o.ConfigurationValueString = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMConfiguredValueByEnum1923) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValueByEnum1923) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMConfiguredValueByEnum1923) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMConfiguredValueByEnum1923) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMConfiguredValueByEnum1923) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValueByEnum1923) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMConfiguredValueByEnum1923) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMConfiguredValueByEnum1923) SetNodeId(v string) {
	o.NodeId = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BTMConfiguredValueByEnum1923) GetValue() BTMParameter1 {
	if o == nil || o.Value == nil {
		var ret BTMParameter1
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValueByEnum1923) GetValueOk() (*BTMParameter1, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BTMConfiguredValueByEnum1923) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given BTMParameter1 and assigns it to the Value field.
func (o *BTMConfiguredValueByEnum1923) SetValue(v BTMParameter1) {
	o.Value = &v
}

// GetEnumName returns the EnumName field value if set, zero value otherwise.
func (o *BTMConfiguredValueByEnum1923) GetEnumName() string {
	if o == nil || o.EnumName == nil {
		var ret string
		return ret
	}
	return *o.EnumName
}

// GetEnumNameOk returns a tuple with the EnumName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValueByEnum1923) GetEnumNameOk() (*string, bool) {
	if o == nil || o.EnumName == nil {
		return nil, false
	}
	return o.EnumName, true
}

// HasEnumName returns a boolean if a field has been set.
func (o *BTMConfiguredValueByEnum1923) HasEnumName() bool {
	if o != nil && o.EnumName != nil {
		return true
	}

	return false
}

// SetEnumName gets a reference to the given string and assigns it to the EnumName field.
func (o *BTMConfiguredValueByEnum1923) SetEnumName(v string) {
	o.EnumName = &v
}

// GetEnumValue returns the EnumValue field value if set, zero value otherwise.
func (o *BTMConfiguredValueByEnum1923) GetEnumValue() string {
	if o == nil || o.EnumValue == nil {
		var ret string
		return ret
	}
	return *o.EnumValue
}

// GetEnumValueOk returns a tuple with the EnumValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValueByEnum1923) GetEnumValueOk() (*string, bool) {
	if o == nil || o.EnumValue == nil {
		return nil, false
	}
	return o.EnumValue, true
}

// HasEnumValue returns a boolean if a field has been set.
func (o *BTMConfiguredValueByEnum1923) HasEnumValue() bool {
	if o != nil && o.EnumValue != nil {
		return true
	}

	return false
}

// SetEnumValue gets a reference to the given string and assigns it to the EnumValue field.
func (o *BTMConfiguredValueByEnum1923) SetEnumValue(v string) {
	o.EnumValue = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTMConfiguredValueByEnum1923) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMConfiguredValueByEnum1923) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTMConfiguredValueByEnum1923) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *BTMConfiguredValueByEnum1923) SetNamespace(v string) {
	o.Namespace = &v
}

func (o BTMConfiguredValueByEnum1923) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ConfigurationValueString != nil {
		toSerialize["configurationValueString"] = o.ConfigurationValueString
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.EnumName != nil {
		toSerialize["enumName"] = o.EnumName
	}
	if o.EnumValue != nil {
		toSerialize["enumValue"] = o.EnumValue
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	return json.Marshal(toSerialize)
}

type NullableBTMConfiguredValueByEnum1923 struct {
	value *BTMConfiguredValueByEnum1923
	isSet bool
}

func (v NullableBTMConfiguredValueByEnum1923) Get() *BTMConfiguredValueByEnum1923 {
	return v.value
}

func (v *NullableBTMConfiguredValueByEnum1923) Set(val *BTMConfiguredValueByEnum1923) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMConfiguredValueByEnum1923) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMConfiguredValueByEnum1923) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMConfiguredValueByEnum1923(val *BTMConfiguredValueByEnum1923) *NullableBTMConfiguredValueByEnum1923 {
	return &NullableBTMConfiguredValueByEnum1923{value: val, isSet: true}
}

func (v NullableBTMConfiguredValueByEnum1923) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMConfiguredValueByEnum1923) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
