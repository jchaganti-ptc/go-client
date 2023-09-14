/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.169.22266-e2d421ffb3ea
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTUnchangedElementInfo struct for BTUnchangedElementInfo
type BTUnchangedElementInfo struct {
	ConnectionId        *string                    `json:"connectionId,omitempty"`
	ElementId           *string                    `json:"elementId,omitempty"`
	UnchangedReferences []BTUnchangedReferenceInfo `json:"unchangedReferences,omitempty"`
}

// NewBTUnchangedElementInfo instantiates a new BTUnchangedElementInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUnchangedElementInfo() *BTUnchangedElementInfo {
	this := BTUnchangedElementInfo{}
	return &this
}

// NewBTUnchangedElementInfoWithDefaults instantiates a new BTUnchangedElementInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUnchangedElementInfoWithDefaults() *BTUnchangedElementInfo {
	this := BTUnchangedElementInfo{}
	return &this
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise.
func (o *BTUnchangedElementInfo) GetConnectionId() string {
	if o == nil || o.ConnectionId == nil {
		var ret string
		return ret
	}
	return *o.ConnectionId
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUnchangedElementInfo) GetConnectionIdOk() (*string, bool) {
	if o == nil || o.ConnectionId == nil {
		return nil, false
	}
	return o.ConnectionId, true
}

// HasConnectionId returns a boolean if a field has been set.
func (o *BTUnchangedElementInfo) HasConnectionId() bool {
	if o != nil && o.ConnectionId != nil {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given string and assigns it to the ConnectionId field.
func (o *BTUnchangedElementInfo) SetConnectionId(v string) {
	o.ConnectionId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTUnchangedElementInfo) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUnchangedElementInfo) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTUnchangedElementInfo) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTUnchangedElementInfo) SetElementId(v string) {
	o.ElementId = &v
}

// GetUnchangedReferences returns the UnchangedReferences field value if set, zero value otherwise.
func (o *BTUnchangedElementInfo) GetUnchangedReferences() []BTUnchangedReferenceInfo {
	if o == nil || o.UnchangedReferences == nil {
		var ret []BTUnchangedReferenceInfo
		return ret
	}
	return o.UnchangedReferences
}

// GetUnchangedReferencesOk returns a tuple with the UnchangedReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUnchangedElementInfo) GetUnchangedReferencesOk() ([]BTUnchangedReferenceInfo, bool) {
	if o == nil || o.UnchangedReferences == nil {
		return nil, false
	}
	return o.UnchangedReferences, true
}

// HasUnchangedReferences returns a boolean if a field has been set.
func (o *BTUnchangedElementInfo) HasUnchangedReferences() bool {
	if o != nil && o.UnchangedReferences != nil {
		return true
	}

	return false
}

// SetUnchangedReferences gets a reference to the given []BTUnchangedReferenceInfo and assigns it to the UnchangedReferences field.
func (o *BTUnchangedElementInfo) SetUnchangedReferences(v []BTUnchangedReferenceInfo) {
	o.UnchangedReferences = v
}

func (o BTUnchangedElementInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConnectionId != nil {
		toSerialize["connectionId"] = o.ConnectionId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.UnchangedReferences != nil {
		toSerialize["unchangedReferences"] = o.UnchangedReferences
	}
	return json.Marshal(toSerialize)
}

type NullableBTUnchangedElementInfo struct {
	value *BTUnchangedElementInfo
	isSet bool
}

func (v NullableBTUnchangedElementInfo) Get() *BTUnchangedElementInfo {
	return v.value
}

func (v *NullableBTUnchangedElementInfo) Set(val *BTUnchangedElementInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUnchangedElementInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUnchangedElementInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUnchangedElementInfo(val *BTUnchangedElementInfo) *NullableBTUnchangedElementInfo {
	return &NullableBTUnchangedElementInfo{value: val, isSet: true}
}

func (v NullableBTUnchangedElementInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUnchangedElementInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
