/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6489-39ccb1a53c2e
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMoveElementParams struct for BTMoveElementParams
type BTMoveElementParams struct {
	AnchorElementId         *string            `json:"anchorElementId,omitempty"`
	Description             *string            `json:"description,omitempty"`
	ElementOriginalToNewMap *map[string]string `json:"elementOriginalToNewMap,omitempty"`
	Elements                []string           `json:"elements,omitempty"`
	GenerateUnknownMessages *bool              `json:"generateUnknownMessages,omitempty"`
	ImportData              []string           `json:"importData,omitempty"`
	IsCopy                  *bool              `json:"isCopy,omitempty"`
	IsDeepCopy              *bool              `json:"isDeepCopy,omitempty"`
	IsGroupAnchor           *bool              `json:"isGroupAnchor,omitempty"`
	IsNewDocument           *bool              `json:"isNewDocument,omitempty"`
	IsPublic                *bool              `json:"isPublic,omitempty"`
	IsSelectivePartOut      *bool              `json:"isSelectivePartOut,omitempty"`
	Name                    *string            `json:"name,omitempty"`
	NeedNewVersion          *bool              `json:"needNewVersion,omitempty"`
	OwnerEmail              *string            `json:"ownerEmail,omitempty"`
	OwnerId                 *string            `json:"ownerId,omitempty"`
	OwnerType               *int32             `json:"ownerType,omitempty"`
	ParentId                *string            `json:"parentId,omitempty"`
	ProjectId               *string            `json:"projectId,omitempty"`
	SelectedGroupIds        []string           `json:"selectedGroupIds,omitempty"`
	SourceDocumentId        *string            `json:"sourceDocumentId,omitempty"`
	SourceVersionId         *string            `json:"sourceVersionId,omitempty"`
	SourceWorkspaceId       *string            `json:"sourceWorkspaceId,omitempty"`
	Tags                    []string           `json:"tags,omitempty"`
	TargetDocumentId        *string            `json:"targetDocumentId,omitempty"`
	TargetWorkspaceId       *string            `json:"targetWorkspaceId,omitempty"`
	VersionName             *string            `json:"versionName,omitempty"`
}

// NewBTMoveElementParams instantiates a new BTMoveElementParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMoveElementParams() *BTMoveElementParams {
	this := BTMoveElementParams{}
	return &this
}

// NewBTMoveElementParamsWithDefaults instantiates a new BTMoveElementParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMoveElementParamsWithDefaults() *BTMoveElementParams {
	this := BTMoveElementParams{}
	return &this
}

// GetAnchorElementId returns the AnchorElementId field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetAnchorElementId() string {
	if o == nil || o.AnchorElementId == nil {
		var ret string
		return ret
	}
	return *o.AnchorElementId
}

// GetAnchorElementIdOk returns a tuple with the AnchorElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetAnchorElementIdOk() (*string, bool) {
	if o == nil || o.AnchorElementId == nil {
		return nil, false
	}
	return o.AnchorElementId, true
}

// HasAnchorElementId returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasAnchorElementId() bool {
	if o != nil && o.AnchorElementId != nil {
		return true
	}

	return false
}

// SetAnchorElementId gets a reference to the given string and assigns it to the AnchorElementId field.
func (o *BTMoveElementParams) SetAnchorElementId(v string) {
	o.AnchorElementId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTMoveElementParams) SetDescription(v string) {
	o.Description = &v
}

// GetElementOriginalToNewMap returns the ElementOriginalToNewMap field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetElementOriginalToNewMap() map[string]string {
	if o == nil || o.ElementOriginalToNewMap == nil {
		var ret map[string]string
		return ret
	}
	return *o.ElementOriginalToNewMap
}

// GetElementOriginalToNewMapOk returns a tuple with the ElementOriginalToNewMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetElementOriginalToNewMapOk() (*map[string]string, bool) {
	if o == nil || o.ElementOriginalToNewMap == nil {
		return nil, false
	}
	return o.ElementOriginalToNewMap, true
}

// HasElementOriginalToNewMap returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasElementOriginalToNewMap() bool {
	if o != nil && o.ElementOriginalToNewMap != nil {
		return true
	}

	return false
}

// SetElementOriginalToNewMap gets a reference to the given map[string]string and assigns it to the ElementOriginalToNewMap field.
func (o *BTMoveElementParams) SetElementOriginalToNewMap(v map[string]string) {
	o.ElementOriginalToNewMap = &v
}

// GetElements returns the Elements field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetElements() []string {
	if o == nil || o.Elements == nil {
		var ret []string
		return ret
	}
	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetElementsOk() ([]string, bool) {
	if o == nil || o.Elements == nil {
		return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasElements() bool {
	if o != nil && o.Elements != nil {
		return true
	}

	return false
}

// SetElements gets a reference to the given []string and assigns it to the Elements field.
func (o *BTMoveElementParams) SetElements(v []string) {
	o.Elements = v
}

// GetGenerateUnknownMessages returns the GenerateUnknownMessages field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetGenerateUnknownMessages() bool {
	if o == nil || o.GenerateUnknownMessages == nil {
		var ret bool
		return ret
	}
	return *o.GenerateUnknownMessages
}

// GetGenerateUnknownMessagesOk returns a tuple with the GenerateUnknownMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetGenerateUnknownMessagesOk() (*bool, bool) {
	if o == nil || o.GenerateUnknownMessages == nil {
		return nil, false
	}
	return o.GenerateUnknownMessages, true
}

// HasGenerateUnknownMessages returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasGenerateUnknownMessages() bool {
	if o != nil && o.GenerateUnknownMessages != nil {
		return true
	}

	return false
}

// SetGenerateUnknownMessages gets a reference to the given bool and assigns it to the GenerateUnknownMessages field.
func (o *BTMoveElementParams) SetGenerateUnknownMessages(v bool) {
	o.GenerateUnknownMessages = &v
}

// GetImportData returns the ImportData field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetImportData() []string {
	if o == nil || o.ImportData == nil {
		var ret []string
		return ret
	}
	return o.ImportData
}

// GetImportDataOk returns a tuple with the ImportData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetImportDataOk() ([]string, bool) {
	if o == nil || o.ImportData == nil {
		return nil, false
	}
	return o.ImportData, true
}

// HasImportData returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasImportData() bool {
	if o != nil && o.ImportData != nil {
		return true
	}

	return false
}

// SetImportData gets a reference to the given []string and assigns it to the ImportData field.
func (o *BTMoveElementParams) SetImportData(v []string) {
	o.ImportData = v
}

// GetIsCopy returns the IsCopy field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetIsCopy() bool {
	if o == nil || o.IsCopy == nil {
		var ret bool
		return ret
	}
	return *o.IsCopy
}

// GetIsCopyOk returns a tuple with the IsCopy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetIsCopyOk() (*bool, bool) {
	if o == nil || o.IsCopy == nil {
		return nil, false
	}
	return o.IsCopy, true
}

// HasIsCopy returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasIsCopy() bool {
	if o != nil && o.IsCopy != nil {
		return true
	}

	return false
}

// SetIsCopy gets a reference to the given bool and assigns it to the IsCopy field.
func (o *BTMoveElementParams) SetIsCopy(v bool) {
	o.IsCopy = &v
}

// GetIsDeepCopy returns the IsDeepCopy field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetIsDeepCopy() bool {
	if o == nil || o.IsDeepCopy == nil {
		var ret bool
		return ret
	}
	return *o.IsDeepCopy
}

// GetIsDeepCopyOk returns a tuple with the IsDeepCopy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetIsDeepCopyOk() (*bool, bool) {
	if o == nil || o.IsDeepCopy == nil {
		return nil, false
	}
	return o.IsDeepCopy, true
}

// HasIsDeepCopy returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasIsDeepCopy() bool {
	if o != nil && o.IsDeepCopy != nil {
		return true
	}

	return false
}

// SetIsDeepCopy gets a reference to the given bool and assigns it to the IsDeepCopy field.
func (o *BTMoveElementParams) SetIsDeepCopy(v bool) {
	o.IsDeepCopy = &v
}

// GetIsGroupAnchor returns the IsGroupAnchor field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetIsGroupAnchor() bool {
	if o == nil || o.IsGroupAnchor == nil {
		var ret bool
		return ret
	}
	return *o.IsGroupAnchor
}

// GetIsGroupAnchorOk returns a tuple with the IsGroupAnchor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetIsGroupAnchorOk() (*bool, bool) {
	if o == nil || o.IsGroupAnchor == nil {
		return nil, false
	}
	return o.IsGroupAnchor, true
}

// HasIsGroupAnchor returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasIsGroupAnchor() bool {
	if o != nil && o.IsGroupAnchor != nil {
		return true
	}

	return false
}

// SetIsGroupAnchor gets a reference to the given bool and assigns it to the IsGroupAnchor field.
func (o *BTMoveElementParams) SetIsGroupAnchor(v bool) {
	o.IsGroupAnchor = &v
}

// GetIsNewDocument returns the IsNewDocument field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetIsNewDocument() bool {
	if o == nil || o.IsNewDocument == nil {
		var ret bool
		return ret
	}
	return *o.IsNewDocument
}

// GetIsNewDocumentOk returns a tuple with the IsNewDocument field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetIsNewDocumentOk() (*bool, bool) {
	if o == nil || o.IsNewDocument == nil {
		return nil, false
	}
	return o.IsNewDocument, true
}

// HasIsNewDocument returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasIsNewDocument() bool {
	if o != nil && o.IsNewDocument != nil {
		return true
	}

	return false
}

// SetIsNewDocument gets a reference to the given bool and assigns it to the IsNewDocument field.
func (o *BTMoveElementParams) SetIsNewDocument(v bool) {
	o.IsNewDocument = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetIsPublic() bool {
	if o == nil || o.IsPublic == nil {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetIsPublicOk() (*bool, bool) {
	if o == nil || o.IsPublic == nil {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasIsPublic() bool {
	if o != nil && o.IsPublic != nil {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *BTMoveElementParams) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetIsSelectivePartOut returns the IsSelectivePartOut field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetIsSelectivePartOut() bool {
	if o == nil || o.IsSelectivePartOut == nil {
		var ret bool
		return ret
	}
	return *o.IsSelectivePartOut
}

// GetIsSelectivePartOutOk returns a tuple with the IsSelectivePartOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetIsSelectivePartOutOk() (*bool, bool) {
	if o == nil || o.IsSelectivePartOut == nil {
		return nil, false
	}
	return o.IsSelectivePartOut, true
}

// HasIsSelectivePartOut returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasIsSelectivePartOut() bool {
	if o != nil && o.IsSelectivePartOut != nil {
		return true
	}

	return false
}

// SetIsSelectivePartOut gets a reference to the given bool and assigns it to the IsSelectivePartOut field.
func (o *BTMoveElementParams) SetIsSelectivePartOut(v bool) {
	o.IsSelectivePartOut = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTMoveElementParams) SetName(v string) {
	o.Name = &v
}

// GetNeedNewVersion returns the NeedNewVersion field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetNeedNewVersion() bool {
	if o == nil || o.NeedNewVersion == nil {
		var ret bool
		return ret
	}
	return *o.NeedNewVersion
}

// GetNeedNewVersionOk returns a tuple with the NeedNewVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetNeedNewVersionOk() (*bool, bool) {
	if o == nil || o.NeedNewVersion == nil {
		return nil, false
	}
	return o.NeedNewVersion, true
}

// HasNeedNewVersion returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasNeedNewVersion() bool {
	if o != nil && o.NeedNewVersion != nil {
		return true
	}

	return false
}

// SetNeedNewVersion gets a reference to the given bool and assigns it to the NeedNewVersion field.
func (o *BTMoveElementParams) SetNeedNewVersion(v bool) {
	o.NeedNewVersion = &v
}

// GetOwnerEmail returns the OwnerEmail field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetOwnerEmail() string {
	if o == nil || o.OwnerEmail == nil {
		var ret string
		return ret
	}
	return *o.OwnerEmail
}

// GetOwnerEmailOk returns a tuple with the OwnerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetOwnerEmailOk() (*string, bool) {
	if o == nil || o.OwnerEmail == nil {
		return nil, false
	}
	return o.OwnerEmail, true
}

// HasOwnerEmail returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasOwnerEmail() bool {
	if o != nil && o.OwnerEmail != nil {
		return true
	}

	return false
}

// SetOwnerEmail gets a reference to the given string and assigns it to the OwnerEmail field.
func (o *BTMoveElementParams) SetOwnerEmail(v string) {
	o.OwnerEmail = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *BTMoveElementParams) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOwnerType returns the OwnerType field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetOwnerType() int32 {
	if o == nil || o.OwnerType == nil {
		var ret int32
		return ret
	}
	return *o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetOwnerTypeOk() (*int32, bool) {
	if o == nil || o.OwnerType == nil {
		return nil, false
	}
	return o.OwnerType, true
}

// HasOwnerType returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasOwnerType() bool {
	if o != nil && o.OwnerType != nil {
		return true
	}

	return false
}

// SetOwnerType gets a reference to the given int32 and assigns it to the OwnerType field.
func (o *BTMoveElementParams) SetOwnerType(v int32) {
	o.OwnerType = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *BTMoveElementParams) SetParentId(v string) {
	o.ParentId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *BTMoveElementParams) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetSelectedGroupIds returns the SelectedGroupIds field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetSelectedGroupIds() []string {
	if o == nil || o.SelectedGroupIds == nil {
		var ret []string
		return ret
	}
	return o.SelectedGroupIds
}

// GetSelectedGroupIdsOk returns a tuple with the SelectedGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetSelectedGroupIdsOk() ([]string, bool) {
	if o == nil || o.SelectedGroupIds == nil {
		return nil, false
	}
	return o.SelectedGroupIds, true
}

// HasSelectedGroupIds returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasSelectedGroupIds() bool {
	if o != nil && o.SelectedGroupIds != nil {
		return true
	}

	return false
}

// SetSelectedGroupIds gets a reference to the given []string and assigns it to the SelectedGroupIds field.
func (o *BTMoveElementParams) SetSelectedGroupIds(v []string) {
	o.SelectedGroupIds = v
}

// GetSourceDocumentId returns the SourceDocumentId field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetSourceDocumentId() string {
	if o == nil || o.SourceDocumentId == nil {
		var ret string
		return ret
	}
	return *o.SourceDocumentId
}

// GetSourceDocumentIdOk returns a tuple with the SourceDocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetSourceDocumentIdOk() (*string, bool) {
	if o == nil || o.SourceDocumentId == nil {
		return nil, false
	}
	return o.SourceDocumentId, true
}

// HasSourceDocumentId returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasSourceDocumentId() bool {
	if o != nil && o.SourceDocumentId != nil {
		return true
	}

	return false
}

// SetSourceDocumentId gets a reference to the given string and assigns it to the SourceDocumentId field.
func (o *BTMoveElementParams) SetSourceDocumentId(v string) {
	o.SourceDocumentId = &v
}

// GetSourceVersionId returns the SourceVersionId field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetSourceVersionId() string {
	if o == nil || o.SourceVersionId == nil {
		var ret string
		return ret
	}
	return *o.SourceVersionId
}

// GetSourceVersionIdOk returns a tuple with the SourceVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetSourceVersionIdOk() (*string, bool) {
	if o == nil || o.SourceVersionId == nil {
		return nil, false
	}
	return o.SourceVersionId, true
}

// HasSourceVersionId returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasSourceVersionId() bool {
	if o != nil && o.SourceVersionId != nil {
		return true
	}

	return false
}

// SetSourceVersionId gets a reference to the given string and assigns it to the SourceVersionId field.
func (o *BTMoveElementParams) SetSourceVersionId(v string) {
	o.SourceVersionId = &v
}

// GetSourceWorkspaceId returns the SourceWorkspaceId field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetSourceWorkspaceId() string {
	if o == nil || o.SourceWorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.SourceWorkspaceId
}

// GetSourceWorkspaceIdOk returns a tuple with the SourceWorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetSourceWorkspaceIdOk() (*string, bool) {
	if o == nil || o.SourceWorkspaceId == nil {
		return nil, false
	}
	return o.SourceWorkspaceId, true
}

// HasSourceWorkspaceId returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasSourceWorkspaceId() bool {
	if o != nil && o.SourceWorkspaceId != nil {
		return true
	}

	return false
}

// SetSourceWorkspaceId gets a reference to the given string and assigns it to the SourceWorkspaceId field.
func (o *BTMoveElementParams) SetSourceWorkspaceId(v string) {
	o.SourceWorkspaceId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *BTMoveElementParams) SetTags(v []string) {
	o.Tags = v
}

// GetTargetDocumentId returns the TargetDocumentId field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetTargetDocumentId() string {
	if o == nil || o.TargetDocumentId == nil {
		var ret string
		return ret
	}
	return *o.TargetDocumentId
}

// GetTargetDocumentIdOk returns a tuple with the TargetDocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetTargetDocumentIdOk() (*string, bool) {
	if o == nil || o.TargetDocumentId == nil {
		return nil, false
	}
	return o.TargetDocumentId, true
}

// HasTargetDocumentId returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasTargetDocumentId() bool {
	if o != nil && o.TargetDocumentId != nil {
		return true
	}

	return false
}

// SetTargetDocumentId gets a reference to the given string and assigns it to the TargetDocumentId field.
func (o *BTMoveElementParams) SetTargetDocumentId(v string) {
	o.TargetDocumentId = &v
}

// GetTargetWorkspaceId returns the TargetWorkspaceId field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetTargetWorkspaceId() string {
	if o == nil || o.TargetWorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.TargetWorkspaceId
}

// GetTargetWorkspaceIdOk returns a tuple with the TargetWorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetTargetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.TargetWorkspaceId == nil {
		return nil, false
	}
	return o.TargetWorkspaceId, true
}

// HasTargetWorkspaceId returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasTargetWorkspaceId() bool {
	if o != nil && o.TargetWorkspaceId != nil {
		return true
	}

	return false
}

// SetTargetWorkspaceId gets a reference to the given string and assigns it to the TargetWorkspaceId field.
func (o *BTMoveElementParams) SetTargetWorkspaceId(v string) {
	o.TargetWorkspaceId = &v
}

// GetVersionName returns the VersionName field value if set, zero value otherwise.
func (o *BTMoveElementParams) GetVersionName() string {
	if o == nil || o.VersionName == nil {
		var ret string
		return ret
	}
	return *o.VersionName
}

// GetVersionNameOk returns a tuple with the VersionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMoveElementParams) GetVersionNameOk() (*string, bool) {
	if o == nil || o.VersionName == nil {
		return nil, false
	}
	return o.VersionName, true
}

// HasVersionName returns a boolean if a field has been set.
func (o *BTMoveElementParams) HasVersionName() bool {
	if o != nil && o.VersionName != nil {
		return true
	}

	return false
}

// SetVersionName gets a reference to the given string and assigns it to the VersionName field.
func (o *BTMoveElementParams) SetVersionName(v string) {
	o.VersionName = &v
}

func (o BTMoveElementParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AnchorElementId != nil {
		toSerialize["anchorElementId"] = o.AnchorElementId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ElementOriginalToNewMap != nil {
		toSerialize["elementOriginalToNewMap"] = o.ElementOriginalToNewMap
	}
	if o.Elements != nil {
		toSerialize["elements"] = o.Elements
	}
	if o.GenerateUnknownMessages != nil {
		toSerialize["generateUnknownMessages"] = o.GenerateUnknownMessages
	}
	if o.ImportData != nil {
		toSerialize["importData"] = o.ImportData
	}
	if o.IsCopy != nil {
		toSerialize["isCopy"] = o.IsCopy
	}
	if o.IsDeepCopy != nil {
		toSerialize["isDeepCopy"] = o.IsDeepCopy
	}
	if o.IsGroupAnchor != nil {
		toSerialize["isGroupAnchor"] = o.IsGroupAnchor
	}
	if o.IsNewDocument != nil {
		toSerialize["isNewDocument"] = o.IsNewDocument
	}
	if o.IsPublic != nil {
		toSerialize["isPublic"] = o.IsPublic
	}
	if o.IsSelectivePartOut != nil {
		toSerialize["isSelectivePartOut"] = o.IsSelectivePartOut
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.NeedNewVersion != nil {
		toSerialize["needNewVersion"] = o.NeedNewVersion
	}
	if o.OwnerEmail != nil {
		toSerialize["ownerEmail"] = o.OwnerEmail
	}
	if o.OwnerId != nil {
		toSerialize["ownerId"] = o.OwnerId
	}
	if o.OwnerType != nil {
		toSerialize["ownerType"] = o.OwnerType
	}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.SelectedGroupIds != nil {
		toSerialize["selectedGroupIds"] = o.SelectedGroupIds
	}
	if o.SourceDocumentId != nil {
		toSerialize["sourceDocumentId"] = o.SourceDocumentId
	}
	if o.SourceVersionId != nil {
		toSerialize["sourceVersionId"] = o.SourceVersionId
	}
	if o.SourceWorkspaceId != nil {
		toSerialize["sourceWorkspaceId"] = o.SourceWorkspaceId
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TargetDocumentId != nil {
		toSerialize["targetDocumentId"] = o.TargetDocumentId
	}
	if o.TargetWorkspaceId != nil {
		toSerialize["targetWorkspaceId"] = o.TargetWorkspaceId
	}
	if o.VersionName != nil {
		toSerialize["versionName"] = o.VersionName
	}
	return json.Marshal(toSerialize)
}

type NullableBTMoveElementParams struct {
	value *BTMoveElementParams
	isSet bool
}

func (v NullableBTMoveElementParams) Get() *BTMoveElementParams {
	return v.value
}

func (v *NullableBTMoveElementParams) Set(val *BTMoveElementParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMoveElementParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMoveElementParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMoveElementParams(val *BTMoveElementParams) *NullableBTMoveElementParams {
	return &NullableBTMoveElementParams{value: val, isSet: true}
}

func (v NullableBTMoveElementParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMoveElementParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
