/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6108-60a91d537e42
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTDiffInfo struct for BTDiffInfo
type BTDiffInfo struct {
	Changes                *map[string]BTDiffInfo `json:"changes,omitempty"`
	EntityType             *string                `json:"entityType,omitempty"`
	GeometryChangeMessages []string               `json:"geometryChangeMessages,omitempty"`
	SourceId               *string                `json:"sourceId,omitempty"`
	SourceValue            *string                `json:"sourceValue,omitempty"`
	TargetId               *string                `json:"targetId,omitempty"`
	TargetValue            *string                `json:"targetValue,omitempty"`
	Type                   *string                `json:"type,omitempty"`
}

// NewBTDiffInfo instantiates a new BTDiffInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDiffInfo() *BTDiffInfo {
	this := BTDiffInfo{}
	return &this
}

// NewBTDiffInfoWithDefaults instantiates a new BTDiffInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDiffInfoWithDefaults() *BTDiffInfo {
	this := BTDiffInfo{}
	return &this
}

// GetChanges returns the Changes field value if set, zero value otherwise.
func (o *BTDiffInfo) GetChanges() map[string]BTDiffInfo {
	if o == nil || o.Changes == nil {
		var ret map[string]BTDiffInfo
		return ret
	}
	return *o.Changes
}

// GetChangesOk returns a tuple with the Changes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiffInfo) GetChangesOk() (*map[string]BTDiffInfo, bool) {
	if o == nil || o.Changes == nil {
		return nil, false
	}
	return o.Changes, true
}

// HasChanges returns a boolean if a field has been set.
func (o *BTDiffInfo) HasChanges() bool {
	if o != nil && o.Changes != nil {
		return true
	}

	return false
}

// SetChanges gets a reference to the given map[string]BTDiffInfo and assigns it to the Changes field.
func (o *BTDiffInfo) SetChanges(v map[string]BTDiffInfo) {
	o.Changes = &v
}

// GetEntityType returns the EntityType field value if set, zero value otherwise.
func (o *BTDiffInfo) GetEntityType() string {
	if o == nil || o.EntityType == nil {
		var ret string
		return ret
	}
	return *o.EntityType
}

// GetEntityTypeOk returns a tuple with the EntityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiffInfo) GetEntityTypeOk() (*string, bool) {
	if o == nil || o.EntityType == nil {
		return nil, false
	}
	return o.EntityType, true
}

// HasEntityType returns a boolean if a field has been set.
func (o *BTDiffInfo) HasEntityType() bool {
	if o != nil && o.EntityType != nil {
		return true
	}

	return false
}

// SetEntityType gets a reference to the given string and assigns it to the EntityType field.
func (o *BTDiffInfo) SetEntityType(v string) {
	o.EntityType = &v
}

// GetGeometryChangeMessages returns the GeometryChangeMessages field value if set, zero value otherwise.
func (o *BTDiffInfo) GetGeometryChangeMessages() []string {
	if o == nil || o.GeometryChangeMessages == nil {
		var ret []string
		return ret
	}
	return o.GeometryChangeMessages
}

// GetGeometryChangeMessagesOk returns a tuple with the GeometryChangeMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiffInfo) GetGeometryChangeMessagesOk() ([]string, bool) {
	if o == nil || o.GeometryChangeMessages == nil {
		return nil, false
	}
	return o.GeometryChangeMessages, true
}

// HasGeometryChangeMessages returns a boolean if a field has been set.
func (o *BTDiffInfo) HasGeometryChangeMessages() bool {
	if o != nil && o.GeometryChangeMessages != nil {
		return true
	}

	return false
}

// SetGeometryChangeMessages gets a reference to the given []string and assigns it to the GeometryChangeMessages field.
func (o *BTDiffInfo) SetGeometryChangeMessages(v []string) {
	o.GeometryChangeMessages = v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise.
func (o *BTDiffInfo) GetSourceId() string {
	if o == nil || o.SourceId == nil {
		var ret string
		return ret
	}
	return *o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiffInfo) GetSourceIdOk() (*string, bool) {
	if o == nil || o.SourceId == nil {
		return nil, false
	}
	return o.SourceId, true
}

// HasSourceId returns a boolean if a field has been set.
func (o *BTDiffInfo) HasSourceId() bool {
	if o != nil && o.SourceId != nil {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given string and assigns it to the SourceId field.
func (o *BTDiffInfo) SetSourceId(v string) {
	o.SourceId = &v
}

// GetSourceValue returns the SourceValue field value if set, zero value otherwise.
func (o *BTDiffInfo) GetSourceValue() string {
	if o == nil || o.SourceValue == nil {
		var ret string
		return ret
	}
	return *o.SourceValue
}

// GetSourceValueOk returns a tuple with the SourceValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiffInfo) GetSourceValueOk() (*string, bool) {
	if o == nil || o.SourceValue == nil {
		return nil, false
	}
	return o.SourceValue, true
}

// HasSourceValue returns a boolean if a field has been set.
func (o *BTDiffInfo) HasSourceValue() bool {
	if o != nil && o.SourceValue != nil {
		return true
	}

	return false
}

// SetSourceValue gets a reference to the given string and assigns it to the SourceValue field.
func (o *BTDiffInfo) SetSourceValue(v string) {
	o.SourceValue = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *BTDiffInfo) GetTargetId() string {
	if o == nil || o.TargetId == nil {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiffInfo) GetTargetIdOk() (*string, bool) {
	if o == nil || o.TargetId == nil {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *BTDiffInfo) HasTargetId() bool {
	if o != nil && o.TargetId != nil {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *BTDiffInfo) SetTargetId(v string) {
	o.TargetId = &v
}

// GetTargetValue returns the TargetValue field value if set, zero value otherwise.
func (o *BTDiffInfo) GetTargetValue() string {
	if o == nil || o.TargetValue == nil {
		var ret string
		return ret
	}
	return *o.TargetValue
}

// GetTargetValueOk returns a tuple with the TargetValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiffInfo) GetTargetValueOk() (*string, bool) {
	if o == nil || o.TargetValue == nil {
		return nil, false
	}
	return o.TargetValue, true
}

// HasTargetValue returns a boolean if a field has been set.
func (o *BTDiffInfo) HasTargetValue() bool {
	if o != nil && o.TargetValue != nil {
		return true
	}

	return false
}

// SetTargetValue gets a reference to the given string and assigns it to the TargetValue field.
func (o *BTDiffInfo) SetTargetValue(v string) {
	o.TargetValue = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTDiffInfo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiffInfo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTDiffInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTDiffInfo) SetType(v string) {
	o.Type = &v
}

func (o BTDiffInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Changes != nil {
		toSerialize["changes"] = o.Changes
	}
	if o.EntityType != nil {
		toSerialize["entityType"] = o.EntityType
	}
	if o.GeometryChangeMessages != nil {
		toSerialize["geometryChangeMessages"] = o.GeometryChangeMessages
	}
	if o.SourceId != nil {
		toSerialize["sourceId"] = o.SourceId
	}
	if o.SourceValue != nil {
		toSerialize["sourceValue"] = o.SourceValue
	}
	if o.TargetId != nil {
		toSerialize["targetId"] = o.TargetId
	}
	if o.TargetValue != nil {
		toSerialize["targetValue"] = o.TargetValue
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBTDiffInfo struct {
	value *BTDiffInfo
	isSet bool
}

func (v NullableBTDiffInfo) Get() *BTDiffInfo {
	return v.value
}

func (v *NullableBTDiffInfo) Set(val *BTDiffInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDiffInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDiffInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDiffInfo(val *BTDiffInfo) *NullableBTDiffInfo {
	return &NullableBTDiffInfo{value: val, isSet: true}
}

func (v NullableBTDiffInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDiffInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
