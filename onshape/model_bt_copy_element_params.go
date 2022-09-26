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

// BTCopyElementParams struct for BTCopyElementParams
type BTCopyElementParams struct {
	AnchorElementId   *string `json:"anchorElementId,omitempty"`
	DocumentIdSource  *string `json:"documentIdSource,omitempty"`
	ElementIdSource   *string `json:"elementIdSource,omitempty"`
	IsGroupAnchor     *bool   `json:"isGroupAnchor,omitempty"`
	WorkspaceIdSource *string `json:"workspaceIdSource,omitempty"`
}

// NewBTCopyElementParams instantiates a new BTCopyElementParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTCopyElementParams() *BTCopyElementParams {
	this := BTCopyElementParams{}
	return &this
}

// NewBTCopyElementParamsWithDefaults instantiates a new BTCopyElementParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTCopyElementParamsWithDefaults() *BTCopyElementParams {
	this := BTCopyElementParams{}
	return &this
}

// GetAnchorElementId returns the AnchorElementId field value if set, zero value otherwise.
func (o *BTCopyElementParams) GetAnchorElementId() string {
	if o == nil || o.AnchorElementId == nil {
		var ret string
		return ret
	}
	return *o.AnchorElementId
}

// GetAnchorElementIdOk returns a tuple with the AnchorElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCopyElementParams) GetAnchorElementIdOk() (*string, bool) {
	if o == nil || o.AnchorElementId == nil {
		return nil, false
	}
	return o.AnchorElementId, true
}

// HasAnchorElementId returns a boolean if a field has been set.
func (o *BTCopyElementParams) HasAnchorElementId() bool {
	if o != nil && o.AnchorElementId != nil {
		return true
	}

	return false
}

// SetAnchorElementId gets a reference to the given string and assigns it to the AnchorElementId field.
func (o *BTCopyElementParams) SetAnchorElementId(v string) {
	o.AnchorElementId = &v
}

// GetDocumentIdSource returns the DocumentIdSource field value if set, zero value otherwise.
func (o *BTCopyElementParams) GetDocumentIdSource() string {
	if o == nil || o.DocumentIdSource == nil {
		var ret string
		return ret
	}
	return *o.DocumentIdSource
}

// GetDocumentIdSourceOk returns a tuple with the DocumentIdSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCopyElementParams) GetDocumentIdSourceOk() (*string, bool) {
	if o == nil || o.DocumentIdSource == nil {
		return nil, false
	}
	return o.DocumentIdSource, true
}

// HasDocumentIdSource returns a boolean if a field has been set.
func (o *BTCopyElementParams) HasDocumentIdSource() bool {
	if o != nil && o.DocumentIdSource != nil {
		return true
	}

	return false
}

// SetDocumentIdSource gets a reference to the given string and assigns it to the DocumentIdSource field.
func (o *BTCopyElementParams) SetDocumentIdSource(v string) {
	o.DocumentIdSource = &v
}

// GetElementIdSource returns the ElementIdSource field value if set, zero value otherwise.
func (o *BTCopyElementParams) GetElementIdSource() string {
	if o == nil || o.ElementIdSource == nil {
		var ret string
		return ret
	}
	return *o.ElementIdSource
}

// GetElementIdSourceOk returns a tuple with the ElementIdSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCopyElementParams) GetElementIdSourceOk() (*string, bool) {
	if o == nil || o.ElementIdSource == nil {
		return nil, false
	}
	return o.ElementIdSource, true
}

// HasElementIdSource returns a boolean if a field has been set.
func (o *BTCopyElementParams) HasElementIdSource() bool {
	if o != nil && o.ElementIdSource != nil {
		return true
	}

	return false
}

// SetElementIdSource gets a reference to the given string and assigns it to the ElementIdSource field.
func (o *BTCopyElementParams) SetElementIdSource(v string) {
	o.ElementIdSource = &v
}

// GetIsGroupAnchor returns the IsGroupAnchor field value if set, zero value otherwise.
func (o *BTCopyElementParams) GetIsGroupAnchor() bool {
	if o == nil || o.IsGroupAnchor == nil {
		var ret bool
		return ret
	}
	return *o.IsGroupAnchor
}

// GetIsGroupAnchorOk returns a tuple with the IsGroupAnchor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCopyElementParams) GetIsGroupAnchorOk() (*bool, bool) {
	if o == nil || o.IsGroupAnchor == nil {
		return nil, false
	}
	return o.IsGroupAnchor, true
}

// HasIsGroupAnchor returns a boolean if a field has been set.
func (o *BTCopyElementParams) HasIsGroupAnchor() bool {
	if o != nil && o.IsGroupAnchor != nil {
		return true
	}

	return false
}

// SetIsGroupAnchor gets a reference to the given bool and assigns it to the IsGroupAnchor field.
func (o *BTCopyElementParams) SetIsGroupAnchor(v bool) {
	o.IsGroupAnchor = &v
}

// GetWorkspaceIdSource returns the WorkspaceIdSource field value if set, zero value otherwise.
func (o *BTCopyElementParams) GetWorkspaceIdSource() string {
	if o == nil || o.WorkspaceIdSource == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceIdSource
}

// GetWorkspaceIdSourceOk returns a tuple with the WorkspaceIdSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTCopyElementParams) GetWorkspaceIdSourceOk() (*string, bool) {
	if o == nil || o.WorkspaceIdSource == nil {
		return nil, false
	}
	return o.WorkspaceIdSource, true
}

// HasWorkspaceIdSource returns a boolean if a field has been set.
func (o *BTCopyElementParams) HasWorkspaceIdSource() bool {
	if o != nil && o.WorkspaceIdSource != nil {
		return true
	}

	return false
}

// SetWorkspaceIdSource gets a reference to the given string and assigns it to the WorkspaceIdSource field.
func (o *BTCopyElementParams) SetWorkspaceIdSource(v string) {
	o.WorkspaceIdSource = &v
}

func (o BTCopyElementParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AnchorElementId != nil {
		toSerialize["anchorElementId"] = o.AnchorElementId
	}
	if o.DocumentIdSource != nil {
		toSerialize["documentIdSource"] = o.DocumentIdSource
	}
	if o.ElementIdSource != nil {
		toSerialize["elementIdSource"] = o.ElementIdSource
	}
	if o.IsGroupAnchor != nil {
		toSerialize["isGroupAnchor"] = o.IsGroupAnchor
	}
	if o.WorkspaceIdSource != nil {
		toSerialize["workspaceIdSource"] = o.WorkspaceIdSource
	}
	return json.Marshal(toSerialize)
}

type NullableBTCopyElementParams struct {
	value *BTCopyElementParams
	isSet bool
}

func (v NullableBTCopyElementParams) Get() *BTCopyElementParams {
	return v.value
}

func (v *NullableBTCopyElementParams) Set(val *BTCopyElementParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTCopyElementParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTCopyElementParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTCopyElementParams(val *BTCopyElementParams) *NullableBTCopyElementParams {
	return &NullableBTCopyElementParams{value: val, isSet: true}
}

func (v NullableBTCopyElementParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTCopyElementParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
