/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.156.7192-0ed4c121c7d8
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTSMDefinitionEntityTypeFilter1651 struct for BTSMDefinitionEntityTypeFilter1651
type BTSMDefinitionEntityTypeFilter1651 struct {
	BtType                 *string `json:"btType,omitempty"`
	SmDefinitionEntityType *string `json:"smDefinitionEntityType,omitempty"`
}

// NewBTSMDefinitionEntityTypeFilter1651 instantiates a new BTSMDefinitionEntityTypeFilter1651 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSMDefinitionEntityTypeFilter1651() *BTSMDefinitionEntityTypeFilter1651 {
	this := BTSMDefinitionEntityTypeFilter1651{}
	return &this
}

// NewBTSMDefinitionEntityTypeFilter1651WithDefaults instantiates a new BTSMDefinitionEntityTypeFilter1651 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSMDefinitionEntityTypeFilter1651WithDefaults() *BTSMDefinitionEntityTypeFilter1651 {
	this := BTSMDefinitionEntityTypeFilter1651{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTSMDefinitionEntityTypeFilter1651) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSMDefinitionEntityTypeFilter1651) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTSMDefinitionEntityTypeFilter1651) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTSMDefinitionEntityTypeFilter1651) SetBtType(v string) {
	o.BtType = &v
}

// GetSmDefinitionEntityType returns the SmDefinitionEntityType field value if set, zero value otherwise.
func (o *BTSMDefinitionEntityTypeFilter1651) GetSmDefinitionEntityType() string {
	if o == nil || o.SmDefinitionEntityType == nil {
		var ret string
		return ret
	}
	return *o.SmDefinitionEntityType
}

// GetSmDefinitionEntityTypeOk returns a tuple with the SmDefinitionEntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSMDefinitionEntityTypeFilter1651) GetSmDefinitionEntityTypeOk() (*string, bool) {
	if o == nil || o.SmDefinitionEntityType == nil {
		return nil, false
	}
	return o.SmDefinitionEntityType, true
}

// HasSmDefinitionEntityType returns a boolean if a field has been set.
func (o *BTSMDefinitionEntityTypeFilter1651) HasSmDefinitionEntityType() bool {
	if o != nil && o.SmDefinitionEntityType != nil {
		return true
	}

	return false
}

// SetSmDefinitionEntityType gets a reference to the given string and assigns it to the SmDefinitionEntityType field.
func (o *BTSMDefinitionEntityTypeFilter1651) SetSmDefinitionEntityType(v string) {
	o.SmDefinitionEntityType = &v
}

func (o BTSMDefinitionEntityTypeFilter1651) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.SmDefinitionEntityType != nil {
		toSerialize["smDefinitionEntityType"] = o.SmDefinitionEntityType
	}
	return json.Marshal(toSerialize)
}

type NullableBTSMDefinitionEntityTypeFilter1651 struct {
	value *BTSMDefinitionEntityTypeFilter1651
	isSet bool
}

func (v NullableBTSMDefinitionEntityTypeFilter1651) Get() *BTSMDefinitionEntityTypeFilter1651 {
	return v.value
}

func (v *NullableBTSMDefinitionEntityTypeFilter1651) Set(val *BTSMDefinitionEntityTypeFilter1651) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSMDefinitionEntityTypeFilter1651) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSMDefinitionEntityTypeFilter1651) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSMDefinitionEntityTypeFilter1651(val *BTSMDefinitionEntityTypeFilter1651) *NullableBTSMDefinitionEntityTypeFilter1651 {
	return &NullableBTSMDefinitionEntityTypeFilter1651{value: val, isSet: true}
}

func (v NullableBTSMDefinitionEntityTypeFilter1651) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSMDefinitionEntityTypeFilter1651) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
