/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5901-446ed8116a32
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTOtherDocumentSelectorParametersInfo struct for BTOtherDocumentSelectorParametersInfo
type BTOtherDocumentSelectorParametersInfo struct {
	JsonType           string              `json:"jsonType"`
	FolderStatePath    []BTFolderStateInfo `json:"folderStatePath,omitempty"`
	SelectedDocumentId *string             `json:"selectedDocumentId,omitempty"`
	SelectedVersionId  *string             `json:"selectedVersionId,omitempty"`
}

// NewBTOtherDocumentSelectorParametersInfo instantiates a new BTOtherDocumentSelectorParametersInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTOtherDocumentSelectorParametersInfo(jsonType string) *BTOtherDocumentSelectorParametersInfo {
	this := BTOtherDocumentSelectorParametersInfo{}
	this.JsonType = jsonType
	return &this
}

// NewBTOtherDocumentSelectorParametersInfoWithDefaults instantiates a new BTOtherDocumentSelectorParametersInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTOtherDocumentSelectorParametersInfoWithDefaults() *BTOtherDocumentSelectorParametersInfo {
	this := BTOtherDocumentSelectorParametersInfo{}
	return &this
}

// GetJsonType returns the JsonType field value
func (o *BTOtherDocumentSelectorParametersInfo) GetJsonType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JsonType
}

// GetJsonTypeOk returns a tuple with the JsonType field value
// and a boolean to check if the value has been set.
func (o *BTOtherDocumentSelectorParametersInfo) GetJsonTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JsonType, true
}

// SetJsonType sets field value
func (o *BTOtherDocumentSelectorParametersInfo) SetJsonType(v string) {
	o.JsonType = v
}

// GetFolderStatePath returns the FolderStatePath field value if set, zero value otherwise.
func (o *BTOtherDocumentSelectorParametersInfo) GetFolderStatePath() []BTFolderStateInfo {
	if o == nil || o.FolderStatePath == nil {
		var ret []BTFolderStateInfo
		return ret
	}
	return o.FolderStatePath
}

// GetFolderStatePathOk returns a tuple with the FolderStatePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOtherDocumentSelectorParametersInfo) GetFolderStatePathOk() ([]BTFolderStateInfo, bool) {
	if o == nil || o.FolderStatePath == nil {
		return nil, false
	}
	return o.FolderStatePath, true
}

// HasFolderStatePath returns a boolean if a field has been set.
func (o *BTOtherDocumentSelectorParametersInfo) HasFolderStatePath() bool {
	if o != nil && o.FolderStatePath != nil {
		return true
	}

	return false
}

// SetFolderStatePath gets a reference to the given []BTFolderStateInfo and assigns it to the FolderStatePath field.
func (o *BTOtherDocumentSelectorParametersInfo) SetFolderStatePath(v []BTFolderStateInfo) {
	o.FolderStatePath = v
}

// GetSelectedDocumentId returns the SelectedDocumentId field value if set, zero value otherwise.
func (o *BTOtherDocumentSelectorParametersInfo) GetSelectedDocumentId() string {
	if o == nil || o.SelectedDocumentId == nil {
		var ret string
		return ret
	}
	return *o.SelectedDocumentId
}

// GetSelectedDocumentIdOk returns a tuple with the SelectedDocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOtherDocumentSelectorParametersInfo) GetSelectedDocumentIdOk() (*string, bool) {
	if o == nil || o.SelectedDocumentId == nil {
		return nil, false
	}
	return o.SelectedDocumentId, true
}

// HasSelectedDocumentId returns a boolean if a field has been set.
func (o *BTOtherDocumentSelectorParametersInfo) HasSelectedDocumentId() bool {
	if o != nil && o.SelectedDocumentId != nil {
		return true
	}

	return false
}

// SetSelectedDocumentId gets a reference to the given string and assigns it to the SelectedDocumentId field.
func (o *BTOtherDocumentSelectorParametersInfo) SetSelectedDocumentId(v string) {
	o.SelectedDocumentId = &v
}

// GetSelectedVersionId returns the SelectedVersionId field value if set, zero value otherwise.
func (o *BTOtherDocumentSelectorParametersInfo) GetSelectedVersionId() string {
	if o == nil || o.SelectedVersionId == nil {
		var ret string
		return ret
	}
	return *o.SelectedVersionId
}

// GetSelectedVersionIdOk returns a tuple with the SelectedVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTOtherDocumentSelectorParametersInfo) GetSelectedVersionIdOk() (*string, bool) {
	if o == nil || o.SelectedVersionId == nil {
		return nil, false
	}
	return o.SelectedVersionId, true
}

// HasSelectedVersionId returns a boolean if a field has been set.
func (o *BTOtherDocumentSelectorParametersInfo) HasSelectedVersionId() bool {
	if o != nil && o.SelectedVersionId != nil {
		return true
	}

	return false
}

// SetSelectedVersionId gets a reference to the given string and assigns it to the SelectedVersionId field.
func (o *BTOtherDocumentSelectorParametersInfo) SetSelectedVersionId(v string) {
	o.SelectedVersionId = &v
}

func (o BTOtherDocumentSelectorParametersInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["jsonType"] = o.JsonType
	}
	if o.FolderStatePath != nil {
		toSerialize["folderStatePath"] = o.FolderStatePath
	}
	if o.SelectedDocumentId != nil {
		toSerialize["selectedDocumentId"] = o.SelectedDocumentId
	}
	if o.SelectedVersionId != nil {
		toSerialize["selectedVersionId"] = o.SelectedVersionId
	}
	return json.Marshal(toSerialize)
}

type NullableBTOtherDocumentSelectorParametersInfo struct {
	value *BTOtherDocumentSelectorParametersInfo
	isSet bool
}

func (v NullableBTOtherDocumentSelectorParametersInfo) Get() *BTOtherDocumentSelectorParametersInfo {
	return v.value
}

func (v *NullableBTOtherDocumentSelectorParametersInfo) Set(val *BTOtherDocumentSelectorParametersInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTOtherDocumentSelectorParametersInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTOtherDocumentSelectorParametersInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTOtherDocumentSelectorParametersInfo(val *BTOtherDocumentSelectorParametersInfo) *NullableBTOtherDocumentSelectorParametersInfo {
	return &NullableBTOtherDocumentSelectorParametersInfo{value: val, isSet: true}
}

func (v NullableBTOtherDocumentSelectorParametersInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTOtherDocumentSelectorParametersInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
