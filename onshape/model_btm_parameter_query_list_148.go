/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5809-984b9f882afe
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMParameterQueryList148 struct for BTMParameterQueryList148
type BTMParameterQueryList148 struct {
	ImportMicroversion *string                     `json:"importMicroversion,omitempty"`
	NodeId             *string                     `json:"nodeId,omitempty"`
	ParameterId        *string                     `json:"parameterId,omitempty"`
	BtType             *string                     `json:"btType,omitempty"`
	Queries            []BTMIndividualQueryBase139 `json:"queries,omitempty"`
}

// NewBTMParameterQueryList148 instantiates a new BTMParameterQueryList148 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMParameterQueryList148() *BTMParameterQueryList148 {
	this := BTMParameterQueryList148{}
	return &this
}

// NewBTMParameterQueryList148WithDefaults instantiates a new BTMParameterQueryList148 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMParameterQueryList148WithDefaults() *BTMParameterQueryList148 {
	this := BTMParameterQueryList148{}
	return &this
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMParameterQueryList148) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterQueryList148) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMParameterQueryList148) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMParameterQueryList148) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMParameterQueryList148) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterQueryList148) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMParameterQueryList148) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMParameterQueryList148) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTMParameterQueryList148) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterQueryList148) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTMParameterQueryList148) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTMParameterQueryList148) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMParameterQueryList148) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterQueryList148) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMParameterQueryList148) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMParameterQueryList148) SetBtType(v string) {
	o.BtType = &v
}

// GetQueries returns the Queries field value if set, zero value otherwise.
func (o *BTMParameterQueryList148) GetQueries() []BTMIndividualQueryBase139 {
	if o == nil || o.Queries == nil {
		var ret []BTMIndividualQueryBase139
		return ret
	}
	return o.Queries
}

// GetQueriesOk returns a tuple with the Queries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterQueryList148) GetQueriesOk() ([]BTMIndividualQueryBase139, bool) {
	if o == nil || o.Queries == nil {
		return nil, false
	}
	return o.Queries, true
}

// HasQueries returns a boolean if a field has been set.
func (o *BTMParameterQueryList148) HasQueries() bool {
	if o != nil && o.Queries != nil {
		return true
	}

	return false
}

// SetQueries gets a reference to the given []BTMIndividualQueryBase139 and assigns it to the Queries field.
func (o *BTMParameterQueryList148) SetQueries(v []BTMIndividualQueryBase139) {
	o.Queries = v
}

func (o BTMParameterQueryList148) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.ParameterId != nil {
		toSerialize["parameterId"] = o.ParameterId
	}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Queries != nil {
		toSerialize["queries"] = o.Queries
	}
	return json.Marshal(toSerialize)
}

type NullableBTMParameterQueryList148 struct {
	value *BTMParameterQueryList148
	isSet bool
}

func (v NullableBTMParameterQueryList148) Get() *BTMParameterQueryList148 {
	return v.value
}

func (v *NullableBTMParameterQueryList148) Set(val *BTMParameterQueryList148) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMParameterQueryList148) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMParameterQueryList148) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMParameterQueryList148(val *BTMParameterQueryList148) *NullableBTMParameterQueryList148 {
	return &NullableBTMParameterQueryList148{value: val, isSet: true}
}

func (v NullableBTMParameterQueryList148) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMParameterQueryList148) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
