/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.164.16419-6916b772b99f
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTEdgeTopologyFilter122 struct for BTEdgeTopologyFilter122
type BTEdgeTopologyFilter122 struct {
	BtType         *string          `json:"btType,omitempty"`
	EdgeTopology   *GBTEdgeTopology `json:"edgeTopology,omitempty"`
	IsInternalEdge *bool            `json:"isInternalEdge,omitempty"`
}

// NewBTEdgeTopologyFilter122 instantiates a new BTEdgeTopologyFilter122 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTEdgeTopologyFilter122() *BTEdgeTopologyFilter122 {
	this := BTEdgeTopologyFilter122{}
	return &this
}

// NewBTEdgeTopologyFilter122WithDefaults instantiates a new BTEdgeTopologyFilter122 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTEdgeTopologyFilter122WithDefaults() *BTEdgeTopologyFilter122 {
	this := BTEdgeTopologyFilter122{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTEdgeTopologyFilter122) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEdgeTopologyFilter122) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTEdgeTopologyFilter122) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTEdgeTopologyFilter122) SetBtType(v string) {
	o.BtType = &v
}

// GetEdgeTopology returns the EdgeTopology field value if set, zero value otherwise.
func (o *BTEdgeTopologyFilter122) GetEdgeTopology() GBTEdgeTopology {
	if o == nil || o.EdgeTopology == nil {
		var ret GBTEdgeTopology
		return ret
	}
	return *o.EdgeTopology
}

// GetEdgeTopologyOk returns a tuple with the EdgeTopology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEdgeTopologyFilter122) GetEdgeTopologyOk() (*GBTEdgeTopology, bool) {
	if o == nil || o.EdgeTopology == nil {
		return nil, false
	}
	return o.EdgeTopology, true
}

// HasEdgeTopology returns a boolean if a field has been set.
func (o *BTEdgeTopologyFilter122) HasEdgeTopology() bool {
	if o != nil && o.EdgeTopology != nil {
		return true
	}

	return false
}

// SetEdgeTopology gets a reference to the given GBTEdgeTopology and assigns it to the EdgeTopology field.
func (o *BTEdgeTopologyFilter122) SetEdgeTopology(v GBTEdgeTopology) {
	o.EdgeTopology = &v
}

// GetIsInternalEdge returns the IsInternalEdge field value if set, zero value otherwise.
func (o *BTEdgeTopologyFilter122) GetIsInternalEdge() bool {
	if o == nil || o.IsInternalEdge == nil {
		var ret bool
		return ret
	}
	return *o.IsInternalEdge
}

// GetIsInternalEdgeOk returns a tuple with the IsInternalEdge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTEdgeTopologyFilter122) GetIsInternalEdgeOk() (*bool, bool) {
	if o == nil || o.IsInternalEdge == nil {
		return nil, false
	}
	return o.IsInternalEdge, true
}

// HasIsInternalEdge returns a boolean if a field has been set.
func (o *BTEdgeTopologyFilter122) HasIsInternalEdge() bool {
	if o != nil && o.IsInternalEdge != nil {
		return true
	}

	return false
}

// SetIsInternalEdge gets a reference to the given bool and assigns it to the IsInternalEdge field.
func (o *BTEdgeTopologyFilter122) SetIsInternalEdge(v bool) {
	o.IsInternalEdge = &v
}

func (o BTEdgeTopologyFilter122) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.EdgeTopology != nil {
		toSerialize["edgeTopology"] = o.EdgeTopology
	}
	if o.IsInternalEdge != nil {
		toSerialize["isInternalEdge"] = o.IsInternalEdge
	}
	return json.Marshal(toSerialize)
}

type NullableBTEdgeTopologyFilter122 struct {
	value *BTEdgeTopologyFilter122
	isSet bool
}

func (v NullableBTEdgeTopologyFilter122) Get() *BTEdgeTopologyFilter122 {
	return v.value
}

func (v *NullableBTEdgeTopologyFilter122) Set(val *BTEdgeTopologyFilter122) {
	v.value = val
	v.isSet = true
}

func (v NullableBTEdgeTopologyFilter122) IsSet() bool {
	return v.isSet
}

func (v *NullableBTEdgeTopologyFilter122) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTEdgeTopologyFilter122(val *BTEdgeTopologyFilter122) *NullableBTEdgeTopologyFilter122 {
	return &NullableBTEdgeTopologyFilter122{value: val, isSet: true}
}

func (v NullableBTEdgeTopologyFilter122) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTEdgeTopologyFilter122) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
