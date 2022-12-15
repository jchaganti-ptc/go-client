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

// BTVersionOrWorkspaceParams struct for BTVersionOrWorkspaceParams
type BTVersionOrWorkspaceParams struct {
	ClientInteractionMode     *string `json:"clientInteractionMode,omitempty"`
	Description               *string `json:"description,omitempty"`
	DocumentId                *string `json:"documentId,omitempty"`
	FromHistory               *bool   `json:"fromHistory,omitempty"`
	IsRelease                 *bool   `json:"isRelease,omitempty"`
	MicroversionId            *string `json:"microversionId,omitempty"`
	MissingBomTableTemplateId *string `json:"missingBomTableTemplateId,omitempty"`
	Name                      *string `json:"name,omitempty"`
	Purpose                   *int32  `json:"purpose,omitempty"`
	ReadOnly                  *bool   `json:"readOnly,omitempty"`
	VersionId                 *string `json:"versionId,omitempty"`
	WorkspaceId               *string `json:"workspaceId,omitempty"`
}

// NewBTVersionOrWorkspaceParams instantiates a new BTVersionOrWorkspaceParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTVersionOrWorkspaceParams() *BTVersionOrWorkspaceParams {
	this := BTVersionOrWorkspaceParams{}
	return &this
}

// NewBTVersionOrWorkspaceParamsWithDefaults instantiates a new BTVersionOrWorkspaceParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTVersionOrWorkspaceParamsWithDefaults() *BTVersionOrWorkspaceParams {
	this := BTVersionOrWorkspaceParams{}
	return &this
}

// GetClientInteractionMode returns the ClientInteractionMode field value if set, zero value otherwise.
func (o *BTVersionOrWorkspaceParams) GetClientInteractionMode() string {
	if o == nil || o.ClientInteractionMode == nil {
		var ret string
		return ret
	}
	return *o.ClientInteractionMode
}

// GetClientInteractionModeOk returns a tuple with the ClientInteractionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionOrWorkspaceParams) GetClientInteractionModeOk() (*string, bool) {
	if o == nil || o.ClientInteractionMode == nil {
		return nil, false
	}
	return o.ClientInteractionMode, true
}

// HasClientInteractionMode returns a boolean if a field has been set.
func (o *BTVersionOrWorkspaceParams) HasClientInteractionMode() bool {
	if o != nil && o.ClientInteractionMode != nil {
		return true
	}

	return false
}

// SetClientInteractionMode gets a reference to the given string and assigns it to the ClientInteractionMode field.
func (o *BTVersionOrWorkspaceParams) SetClientInteractionMode(v string) {
	o.ClientInteractionMode = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTVersionOrWorkspaceParams) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionOrWorkspaceParams) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTVersionOrWorkspaceParams) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTVersionOrWorkspaceParams) SetDescription(v string) {
	o.Description = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTVersionOrWorkspaceParams) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionOrWorkspaceParams) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTVersionOrWorkspaceParams) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTVersionOrWorkspaceParams) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetFromHistory returns the FromHistory field value if set, zero value otherwise.
func (o *BTVersionOrWorkspaceParams) GetFromHistory() bool {
	if o == nil || o.FromHistory == nil {
		var ret bool
		return ret
	}
	return *o.FromHistory
}

// GetFromHistoryOk returns a tuple with the FromHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionOrWorkspaceParams) GetFromHistoryOk() (*bool, bool) {
	if o == nil || o.FromHistory == nil {
		return nil, false
	}
	return o.FromHistory, true
}

// HasFromHistory returns a boolean if a field has been set.
func (o *BTVersionOrWorkspaceParams) HasFromHistory() bool {
	if o != nil && o.FromHistory != nil {
		return true
	}

	return false
}

// SetFromHistory gets a reference to the given bool and assigns it to the FromHistory field.
func (o *BTVersionOrWorkspaceParams) SetFromHistory(v bool) {
	o.FromHistory = &v
}

// GetIsRelease returns the IsRelease field value if set, zero value otherwise.
func (o *BTVersionOrWorkspaceParams) GetIsRelease() bool {
	if o == nil || o.IsRelease == nil {
		var ret bool
		return ret
	}
	return *o.IsRelease
}

// GetIsReleaseOk returns a tuple with the IsRelease field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionOrWorkspaceParams) GetIsReleaseOk() (*bool, bool) {
	if o == nil || o.IsRelease == nil {
		return nil, false
	}
	return o.IsRelease, true
}

// HasIsRelease returns a boolean if a field has been set.
func (o *BTVersionOrWorkspaceParams) HasIsRelease() bool {
	if o != nil && o.IsRelease != nil {
		return true
	}

	return false
}

// SetIsRelease gets a reference to the given bool and assigns it to the IsRelease field.
func (o *BTVersionOrWorkspaceParams) SetIsRelease(v bool) {
	o.IsRelease = &v
}

// GetMicroversionId returns the MicroversionId field value if set, zero value otherwise.
func (o *BTVersionOrWorkspaceParams) GetMicroversionId() string {
	if o == nil || o.MicroversionId == nil {
		var ret string
		return ret
	}
	return *o.MicroversionId
}

// GetMicroversionIdOk returns a tuple with the MicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionOrWorkspaceParams) GetMicroversionIdOk() (*string, bool) {
	if o == nil || o.MicroversionId == nil {
		return nil, false
	}
	return o.MicroversionId, true
}

// HasMicroversionId returns a boolean if a field has been set.
func (o *BTVersionOrWorkspaceParams) HasMicroversionId() bool {
	if o != nil && o.MicroversionId != nil {
		return true
	}

	return false
}

// SetMicroversionId gets a reference to the given string and assigns it to the MicroversionId field.
func (o *BTVersionOrWorkspaceParams) SetMicroversionId(v string) {
	o.MicroversionId = &v
}

// GetMissingBomTableTemplateId returns the MissingBomTableTemplateId field value if set, zero value otherwise.
func (o *BTVersionOrWorkspaceParams) GetMissingBomTableTemplateId() string {
	if o == nil || o.MissingBomTableTemplateId == nil {
		var ret string
		return ret
	}
	return *o.MissingBomTableTemplateId
}

// GetMissingBomTableTemplateIdOk returns a tuple with the MissingBomTableTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionOrWorkspaceParams) GetMissingBomTableTemplateIdOk() (*string, bool) {
	if o == nil || o.MissingBomTableTemplateId == nil {
		return nil, false
	}
	return o.MissingBomTableTemplateId, true
}

// HasMissingBomTableTemplateId returns a boolean if a field has been set.
func (o *BTVersionOrWorkspaceParams) HasMissingBomTableTemplateId() bool {
	if o != nil && o.MissingBomTableTemplateId != nil {
		return true
	}

	return false
}

// SetMissingBomTableTemplateId gets a reference to the given string and assigns it to the MissingBomTableTemplateId field.
func (o *BTVersionOrWorkspaceParams) SetMissingBomTableTemplateId(v string) {
	o.MissingBomTableTemplateId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTVersionOrWorkspaceParams) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionOrWorkspaceParams) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTVersionOrWorkspaceParams) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTVersionOrWorkspaceParams) SetName(v string) {
	o.Name = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *BTVersionOrWorkspaceParams) GetPurpose() int32 {
	if o == nil || o.Purpose == nil {
		var ret int32
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionOrWorkspaceParams) GetPurposeOk() (*int32, bool) {
	if o == nil || o.Purpose == nil {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *BTVersionOrWorkspaceParams) HasPurpose() bool {
	if o != nil && o.Purpose != nil {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given int32 and assigns it to the Purpose field.
func (o *BTVersionOrWorkspaceParams) SetPurpose(v int32) {
	o.Purpose = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *BTVersionOrWorkspaceParams) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionOrWorkspaceParams) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *BTVersionOrWorkspaceParams) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *BTVersionOrWorkspaceParams) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *BTVersionOrWorkspaceParams) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionOrWorkspaceParams) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *BTVersionOrWorkspaceParams) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *BTVersionOrWorkspaceParams) SetVersionId(v string) {
	o.VersionId = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *BTVersionOrWorkspaceParams) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTVersionOrWorkspaceParams) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *BTVersionOrWorkspaceParams) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *BTVersionOrWorkspaceParams) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

func (o BTVersionOrWorkspaceParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientInteractionMode != nil {
		toSerialize["clientInteractionMode"] = o.ClientInteractionMode
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.FromHistory != nil {
		toSerialize["fromHistory"] = o.FromHistory
	}
	if o.IsRelease != nil {
		toSerialize["isRelease"] = o.IsRelease
	}
	if o.MicroversionId != nil {
		toSerialize["microversionId"] = o.MicroversionId
	}
	if o.MissingBomTableTemplateId != nil {
		toSerialize["missingBomTableTemplateId"] = o.MissingBomTableTemplateId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Purpose != nil {
		toSerialize["purpose"] = o.Purpose
	}
	if o.ReadOnly != nil {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if o.VersionId != nil {
		toSerialize["versionId"] = o.VersionId
	}
	if o.WorkspaceId != nil {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableBTVersionOrWorkspaceParams struct {
	value *BTVersionOrWorkspaceParams
	isSet bool
}

func (v NullableBTVersionOrWorkspaceParams) Get() *BTVersionOrWorkspaceParams {
	return v.value
}

func (v *NullableBTVersionOrWorkspaceParams) Set(val *BTVersionOrWorkspaceParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTVersionOrWorkspaceParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTVersionOrWorkspaceParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTVersionOrWorkspaceParams(val *BTVersionOrWorkspaceParams) *NullableBTVersionOrWorkspaceParams {
	return &NullableBTVersionOrWorkspaceParams{value: val, isSet: true}
}

func (v NullableBTVersionOrWorkspaceParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTVersionOrWorkspaceParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
