/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6135-f341d0cf30dd
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTVersionOrWorkspaceMergeInfo struct for BTVersionOrWorkspaceMergeInfo
type BTVersionOrWorkspaceMergeInfo struct {
	DefaultMergeStrategy          *string            `json:"defaultMergeStrategy,omitempty"`
	Id                            *string            `json:"id,omitempty"`
	MergeStrategyElementOverrides *map[string]string `json:"mergeStrategyElementOverrides,omitempty"`
	Type                          *string            `json:"type,omitempty"`
}

// NewBTVersionOrWorkspaceMergeInfo instantiates a new BTVersionOrWorkspaceMergeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTVersionOrWorkspaceMergeInfo() *BTVersionOrWorkspaceMergeInfo {
	this := BTVersionOrWorkspaceMergeInfo{}
	return &this
}

// NewBTVersionOrWorkspaceMergeInfoWithDefaults instantiates a new BTVersionOrWorkspaceMergeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTVersionOrWorkspaceMergeInfoWithDefaults() *BTVersionOrWorkspaceMergeInfo {
	this := BTVersionOrWorkspaceMergeInfo{}
	return &this
}

// GetDefaultMergeStrategy returns the DefaultMergeStrategy field value if set, zero value otherwise.
func (o *BTVersionOrWorkspaceMergeInfo) GetDefaultMergeStrategy() string {
	if o == nil || o.DefaultMergeStrategy == nil {
		var ret string
		return ret
	}
	return *o.DefaultMergeStrategy
}

// GetDefaultMergeStrategyOk returns a tuple with the DefaultMergeStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionOrWorkspaceMergeInfo) GetDefaultMergeStrategyOk() (*string, bool) {
	if o == nil || o.DefaultMergeStrategy == nil {
		return nil, false
	}
	return o.DefaultMergeStrategy, true
}

// HasDefaultMergeStrategy returns a boolean if a field has been set.
func (o *BTVersionOrWorkspaceMergeInfo) HasDefaultMergeStrategy() bool {
	if o != nil && o.DefaultMergeStrategy != nil {
		return true
	}

	return false
}

// SetDefaultMergeStrategy gets a reference to the given string and assigns it to the DefaultMergeStrategy field.
func (o *BTVersionOrWorkspaceMergeInfo) SetDefaultMergeStrategy(v string) {
	o.DefaultMergeStrategy = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTVersionOrWorkspaceMergeInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionOrWorkspaceMergeInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTVersionOrWorkspaceMergeInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTVersionOrWorkspaceMergeInfo) SetId(v string) {
	o.Id = &v
}

// GetMergeStrategyElementOverrides returns the MergeStrategyElementOverrides field value if set, zero value otherwise.
func (o *BTVersionOrWorkspaceMergeInfo) GetMergeStrategyElementOverrides() map[string]string {
	if o == nil || o.MergeStrategyElementOverrides == nil {
		var ret map[string]string
		return ret
	}
	return *o.MergeStrategyElementOverrides
}

// GetMergeStrategyElementOverridesOk returns a tuple with the MergeStrategyElementOverrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionOrWorkspaceMergeInfo) GetMergeStrategyElementOverridesOk() (*map[string]string, bool) {
	if o == nil || o.MergeStrategyElementOverrides == nil {
		return nil, false
	}
	return o.MergeStrategyElementOverrides, true
}

// HasMergeStrategyElementOverrides returns a boolean if a field has been set.
func (o *BTVersionOrWorkspaceMergeInfo) HasMergeStrategyElementOverrides() bool {
	if o != nil && o.MergeStrategyElementOverrides != nil {
		return true
	}

	return false
}

// SetMergeStrategyElementOverrides gets a reference to the given map[string]string and assigns it to the MergeStrategyElementOverrides field.
func (o *BTVersionOrWorkspaceMergeInfo) SetMergeStrategyElementOverrides(v map[string]string) {
	o.MergeStrategyElementOverrides = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTVersionOrWorkspaceMergeInfo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionOrWorkspaceMergeInfo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTVersionOrWorkspaceMergeInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTVersionOrWorkspaceMergeInfo) SetType(v string) {
	o.Type = &v
}

func (o BTVersionOrWorkspaceMergeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultMergeStrategy != nil {
		toSerialize["defaultMergeStrategy"] = o.DefaultMergeStrategy
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MergeStrategyElementOverrides != nil {
		toSerialize["mergeStrategyElementOverrides"] = o.MergeStrategyElementOverrides
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBTVersionOrWorkspaceMergeInfo struct {
	value *BTVersionOrWorkspaceMergeInfo
	isSet bool
}

func (v NullableBTVersionOrWorkspaceMergeInfo) Get() *BTVersionOrWorkspaceMergeInfo {
	return v.value
}

func (v *NullableBTVersionOrWorkspaceMergeInfo) Set(val *BTVersionOrWorkspaceMergeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTVersionOrWorkspaceMergeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTVersionOrWorkspaceMergeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTVersionOrWorkspaceMergeInfo(val *BTVersionOrWorkspaceMergeInfo) *NullableBTVersionOrWorkspaceMergeInfo {
	return &NullableBTVersionOrWorkspaceMergeInfo{value: val, isSet: true}
}

func (v NullableBTVersionOrWorkspaceMergeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTVersionOrWorkspaceMergeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
