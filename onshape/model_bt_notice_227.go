/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.157.8712-62a9cfa549d9
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTNotice227 struct for BTNotice227
type BTNotice227 struct {
	Level         *string             `json:"level,omitempty"`
	LocationInfos []BTLocationInfo226 `json:"locationInfos,omitempty"`
	Message       *string             `json:"message,omitempty"`
	NodeId        *string             `json:"nodeId,omitempty"`
	ParameterId   *string             `json:"parameterId,omitempty"`
	StackTrace    []BTLocationInfo226 `json:"stackTrace,omitempty"`
	TryNode       *BTNodeReference21  `json:"tryNode,omitempty"`
	Type          *string             `json:"type,omitempty"`
}

// NewBTNotice227 instantiates a new BTNotice227 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTNotice227() *BTNotice227 {
	this := BTNotice227{}
	return &this
}

// NewBTNotice227WithDefaults instantiates a new BTNotice227 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTNotice227WithDefaults() *BTNotice227 {
	this := BTNotice227{}
	return &this
}

// GetLevel returns the Level field value if set, zero value otherwise.
func (o *BTNotice227) GetLevel() string {
	if o == nil || o.Level == nil {
		var ret string
		return ret
	}
	return *o.Level
}

// GetLevelOk returns a tuple with the Level field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNotice227) GetLevelOk() (*string, bool) {
	if o == nil || o.Level == nil {
		return nil, false
	}
	return o.Level, true
}

// HasLevel returns a boolean if a field has been set.
func (o *BTNotice227) HasLevel() bool {
	if o != nil && o.Level != nil {
		return true
	}

	return false
}

// SetLevel gets a reference to the given string and assigns it to the Level field.
func (o *BTNotice227) SetLevel(v string) {
	o.Level = &v
}

// GetLocationInfos returns the LocationInfos field value if set, zero value otherwise.
func (o *BTNotice227) GetLocationInfos() []BTLocationInfo226 {
	if o == nil || o.LocationInfos == nil {
		var ret []BTLocationInfo226
		return ret
	}
	return o.LocationInfos
}

// GetLocationInfosOk returns a tuple with the LocationInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNotice227) GetLocationInfosOk() ([]BTLocationInfo226, bool) {
	if o == nil || o.LocationInfos == nil {
		return nil, false
	}
	return o.LocationInfos, true
}

// HasLocationInfos returns a boolean if a field has been set.
func (o *BTNotice227) HasLocationInfos() bool {
	if o != nil && o.LocationInfos != nil {
		return true
	}

	return false
}

// SetLocationInfos gets a reference to the given []BTLocationInfo226 and assigns it to the LocationInfos field.
func (o *BTNotice227) SetLocationInfos(v []BTLocationInfo226) {
	o.LocationInfos = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BTNotice227) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNotice227) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BTNotice227) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BTNotice227) SetMessage(v string) {
	o.Message = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTNotice227) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNotice227) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTNotice227) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTNotice227) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTNotice227) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNotice227) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTNotice227) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTNotice227) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetStackTrace returns the StackTrace field value if set, zero value otherwise.
func (o *BTNotice227) GetStackTrace() []BTLocationInfo226 {
	if o == nil || o.StackTrace == nil {
		var ret []BTLocationInfo226
		return ret
	}
	return o.StackTrace
}

// GetStackTraceOk returns a tuple with the StackTrace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNotice227) GetStackTraceOk() ([]BTLocationInfo226, bool) {
	if o == nil || o.StackTrace == nil {
		return nil, false
	}
	return o.StackTrace, true
}

// HasStackTrace returns a boolean if a field has been set.
func (o *BTNotice227) HasStackTrace() bool {
	if o != nil && o.StackTrace != nil {
		return true
	}

	return false
}

// SetStackTrace gets a reference to the given []BTLocationInfo226 and assigns it to the StackTrace field.
func (o *BTNotice227) SetStackTrace(v []BTLocationInfo226) {
	o.StackTrace = v
}

// GetTryNode returns the TryNode field value if set, zero value otherwise.
func (o *BTNotice227) GetTryNode() BTNodeReference21 {
	if o == nil || o.TryNode == nil {
		var ret BTNodeReference21
		return ret
	}
	return *o.TryNode
}

// GetTryNodeOk returns a tuple with the TryNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNotice227) GetTryNodeOk() (*BTNodeReference21, bool) {
	if o == nil || o.TryNode == nil {
		return nil, false
	}
	return o.TryNode, true
}

// HasTryNode returns a boolean if a field has been set.
func (o *BTNotice227) HasTryNode() bool {
	if o != nil && o.TryNode != nil {
		return true
	}

	return false
}

// SetTryNode gets a reference to the given BTNodeReference21 and assigns it to the TryNode field.
func (o *BTNotice227) SetTryNode(v BTNodeReference21) {
	o.TryNode = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTNotice227) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTNotice227) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTNotice227) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTNotice227) SetType(v string) {
	o.Type = &v
}

func (o BTNotice227) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Level != nil {
		toSerialize["level"] = o.Level
	}
	if o.LocationInfos != nil {
		toSerialize["locationInfos"] = o.LocationInfos
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.StackTrace != nil {
		toSerialize["stackTrace"] = o.StackTrace
	}
	if o.TryNode != nil {
		toSerialize["tryNode"] = o.TryNode
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBTNotice227 struct {
	value *BTNotice227
	isSet bool
}

func (v NullableBTNotice227) Get() *BTNotice227 {
	return v.value
}

func (v *NullableBTNotice227) Set(val *BTNotice227) {
	v.value = val
	v.isSet = true
}

func (v NullableBTNotice227) IsSet() bool {
	return v.isSet
}

func (v *NullableBTNotice227) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTNotice227(val *BTNotice227) *NullableBTNotice227 {
	return &NullableBTNotice227{value: val, isSet: true}
}

func (v NullableBTNotice227) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTNotice227) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
