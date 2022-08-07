/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5931-1b7b413e7f30
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMParameterProgress3232 struct for BTMParameterProgress3232
type BTMParameterProgress3232 struct {
	ImportMicroversion *string  `json:"importMicroversion,omitempty"`
	NodeId             *string  `json:"nodeId,omitempty"`
	ParameterId        *string  `json:"parameterId,omitempty"`
	BtType             *string  `json:"btType,omitempty"`
	PercentDone        *float64 `json:"percentDone,omitempty"`
	Status             *string  `json:"status,omitempty"`
	StatusMessage      *string  `json:"statusMessage,omitempty"`
}

// NewBTMParameterProgress3232 instantiates a new BTMParameterProgress3232 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMParameterProgress3232() *BTMParameterProgress3232 {
	this := BTMParameterProgress3232{}
	return &this
}

// NewBTMParameterProgress3232WithDefaults instantiates a new BTMParameterProgress3232 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMParameterProgress3232WithDefaults() *BTMParameterProgress3232 {
	this := BTMParameterProgress3232{}
	return &this
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMParameterProgress3232) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterProgress3232) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMParameterProgress3232) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMParameterProgress3232) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMParameterProgress3232) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterProgress3232) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMParameterProgress3232) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMParameterProgress3232) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameterId returns the ParameterId field value if set, zero value otherwise.
func (o *BTMParameterProgress3232) GetParameterId() string {
	if o == nil || o.ParameterId == nil {
		var ret string
		return ret
	}
	return *o.ParameterId
}

// GetParameterIdOk returns a tuple with the ParameterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterProgress3232) GetParameterIdOk() (*string, bool) {
	if o == nil || o.ParameterId == nil {
		return nil, false
	}
	return o.ParameterId, true
}

// HasParameterId returns a boolean if a field has been set.
func (o *BTMParameterProgress3232) HasParameterId() bool {
	if o != nil && o.ParameterId != nil {
		return true
	}

	return false
}

// SetParameterId gets a reference to the given string and assigns it to the ParameterId field.
func (o *BTMParameterProgress3232) SetParameterId(v string) {
	o.ParameterId = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMParameterProgress3232) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterProgress3232) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMParameterProgress3232) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMParameterProgress3232) SetBtType(v string) {
	o.BtType = &v
}

// GetPercentDone returns the PercentDone field value if set, zero value otherwise.
func (o *BTMParameterProgress3232) GetPercentDone() float64 {
	if o == nil || o.PercentDone == nil {
		var ret float64
		return ret
	}
	return *o.PercentDone
}

// GetPercentDoneOk returns a tuple with the PercentDone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterProgress3232) GetPercentDoneOk() (*float64, bool) {
	if o == nil || o.PercentDone == nil {
		return nil, false
	}
	return o.PercentDone, true
}

// HasPercentDone returns a boolean if a field has been set.
func (o *BTMParameterProgress3232) HasPercentDone() bool {
	if o != nil && o.PercentDone != nil {
		return true
	}

	return false
}

// SetPercentDone gets a reference to the given float64 and assigns it to the PercentDone field.
func (o *BTMParameterProgress3232) SetPercentDone(v float64) {
	o.PercentDone = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BTMParameterProgress3232) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterProgress3232) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BTMParameterProgress3232) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BTMParameterProgress3232) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *BTMParameterProgress3232) GetStatusMessage() string {
	if o == nil || o.StatusMessage == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMParameterProgress3232) GetStatusMessageOk() (*string, bool) {
	if o == nil || o.StatusMessage == nil {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *BTMParameterProgress3232) HasStatusMessage() bool {
	if o != nil && o.StatusMessage != nil {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *BTMParameterProgress3232) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

func (o BTMParameterProgress3232) MarshalJSON() ([]byte, error) {
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
	if o.PercentDone != nil {
		toSerialize["percentDone"] = o.PercentDone
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusMessage != nil {
		toSerialize["statusMessage"] = o.StatusMessage
	}
	return json.Marshal(toSerialize)
}

type NullableBTMParameterProgress3232 struct {
	value *BTMParameterProgress3232
	isSet bool
}

func (v NullableBTMParameterProgress3232) Get() *BTMParameterProgress3232 {
	return v.value
}

func (v *NullableBTMParameterProgress3232) Set(val *BTMParameterProgress3232) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMParameterProgress3232) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMParameterProgress3232) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMParameterProgress3232(val *BTMParameterProgress3232) *NullableBTMParameterProgress3232 {
	return &NullableBTMParameterProgress3232{value: val, isSet: true}
}

func (v NullableBTMParameterProgress3232) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMParameterProgress3232) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
