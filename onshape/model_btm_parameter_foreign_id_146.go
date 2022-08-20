/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6126-5c3a878ad24b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMParameterForeignId146 struct for BTMParameterForeignId146
type BTMParameterForeignId146 struct {
	ImportMicroversion *string                    `json:"importMicroversion,omitempty"`
	NodeId             *string                    `json:"nodeId,omitempty"`
	ParameterId        *string                    `json:"parameterId,omitempty"`
	BtType             *string                    `json:"btType,omitempty"`
	ForeignId          *string                    `json:"foreignId,omitempty"`
	ForeignName        *string                    `json:"foreignName,omitempty"`
	LocationInfo       *BTForeignDataResponse1070 `json:"locationInfo,omitempty"`
}

// NewBTMParameterForeignId146 instantiates a new BTMParameterForeignId146 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMParameterForeignId146() *BTMParameterForeignId146 {
	this := BTMParameterForeignId146{}
	return &this
}

// NewBTMParameterForeignId146WithDefaults instantiates a new BTMParameterForeignId146 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMParameterForeignId146WithDefaults() *BTMParameterForeignId146 {
	this := BTMParameterForeignId146{}
	return &this
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMParameterForeignId146) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterForeignId146) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMParameterForeignId146) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMParameterForeignId146) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMParameterForeignId146) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterForeignId146) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMParameterForeignId146) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMParameterForeignId146) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTMParameterForeignId146) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterForeignId146) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTMParameterForeignId146) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTMParameterForeignId146) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMParameterForeignId146) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterForeignId146) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMParameterForeignId146) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMParameterForeignId146) SetBtType(v string) {
	o.BtType = &v
}

// GetForeignId returns the ForeignId field value if set, zero value otherwise.
func (o *BTMParameterForeignId146) GetForeignId() string {
	if o == nil || o.ForeignId == nil {
		var ret string
		return ret
	}
	return *o.ForeignId
}

// GetForeignIdOk returns a tuple with the ForeignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterForeignId146) GetForeignIdOk() (*string, bool) {
	if o == nil || o.ForeignId == nil {
		return nil, false
	}
	return o.ForeignId, true
}

// HasForeignId returns a boolean if a field has been set.
func (o *BTMParameterForeignId146) HasForeignId() bool {
	if o != nil && o.ForeignId != nil {
		return true
	}

	return false
}

// SetForeignId gets a reference to the given string and assigns it to the ForeignId field.
func (o *BTMParameterForeignId146) SetForeignId(v string) {
	o.ForeignId = &v
}

// GetForeignName returns the ForeignName field value if set, zero value otherwise.
func (o *BTMParameterForeignId146) GetForeignName() string {
	if o == nil || o.ForeignName == nil {
		var ret string
		return ret
	}
	return *o.ForeignName
}

// GetForeignNameOk returns a tuple with the ForeignName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterForeignId146) GetForeignNameOk() (*string, bool) {
	if o == nil || o.ForeignName == nil {
		return nil, false
	}
	return o.ForeignName, true
}

// HasForeignName returns a boolean if a field has been set.
func (o *BTMParameterForeignId146) HasForeignName() bool {
	if o != nil && o.ForeignName != nil {
		return true
	}

	return false
}

// SetForeignName gets a reference to the given string and assigns it to the ForeignName field.
func (o *BTMParameterForeignId146) SetForeignName(v string) {
	o.ForeignName = &v
}

// GetLocationInfo returns the LocationInfo field value if set, zero value otherwise.
func (o *BTMParameterForeignId146) GetLocationInfo() BTForeignDataResponse1070 {
	if o == nil || o.LocationInfo == nil {
		var ret BTForeignDataResponse1070
		return ret
	}
	return *o.LocationInfo
}

// GetLocationInfoOk returns a tuple with the LocationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterForeignId146) GetLocationInfoOk() (*BTForeignDataResponse1070, bool) {
	if o == nil || o.LocationInfo == nil {
		return nil, false
	}
	return o.LocationInfo, true
}

// HasLocationInfo returns a boolean if a field has been set.
func (o *BTMParameterForeignId146) HasLocationInfo() bool {
	if o != nil && o.LocationInfo != nil {
		return true
	}

	return false
}

// SetLocationInfo gets a reference to the given BTForeignDataResponse1070 and assigns it to the LocationInfo field.
func (o *BTMParameterForeignId146) SetLocationInfo(v BTForeignDataResponse1070) {
	o.LocationInfo = &v
}

func (o BTMParameterForeignId146) MarshalJSON() ([]byte, error) {
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
	if o.ForeignId != nil {
		toSerialize["foreignId"] = o.ForeignId
	}
	if o.ForeignName != nil {
		toSerialize["foreignName"] = o.ForeignName
	}
	if o.LocationInfo != nil {
		toSerialize["locationInfo"] = o.LocationInfo
	}
	return json.Marshal(toSerialize)
}

type NullableBTMParameterForeignId146 struct {
	value *BTMParameterForeignId146
	isSet bool
}

func (v NullableBTMParameterForeignId146) Get() *BTMParameterForeignId146 {
	return v.value
}

func (v *NullableBTMParameterForeignId146) Set(val *BTMParameterForeignId146) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMParameterForeignId146) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMParameterForeignId146) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMParameterForeignId146(val *BTMParameterForeignId146) *NullableBTMParameterForeignId146 {
	return &NullableBTMParameterForeignId146{value: val, isSet: true}
}

func (v NullableBTMParameterForeignId146) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMParameterForeignId146) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
