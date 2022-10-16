/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.6926-35d1d6168263
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTProjectInfo struct for BTProjectInfo
type BTProjectInfo struct {
	CanMove           *bool                       `json:"canMove,omitempty"`
	CreatedAt         *JSONTime                   `json:"createdAt,omitempty"`
	CreatedBy         *BTUserBasicSummaryInfo     `json:"createdBy,omitempty"`
	Description       *string                     `json:"description,omitempty"`
	Href              *string                     `json:"href,omitempty"`
	Id                *string                     `json:"id,omitempty"`
	IsContainer       *bool                       `json:"isContainer,omitempty"`
	IsEnterpriseOwned *bool                       `json:"isEnterpriseOwned,omitempty"`
	IsMutable         *bool                       `json:"isMutable,omitempty"`
	ModifiedAt        *JSONTime                   `json:"modifiedAt,omitempty"`
	ModifiedBy        *BTUserBasicSummaryInfo     `json:"modifiedBy,omitempty"`
	Name              *string                     `json:"name,omitempty"`
	Owner             *BTOwnerInfo                `json:"owner,omitempty"`
	PermissionScheme  *BTRbacPermissionSchemeInfo `json:"permissionScheme,omitempty"`
	PermissionSet     []string                    `json:"permissionSet,omitempty"`
	ProjectId         *string                     `json:"projectId,omitempty"`
	ResourceType      *string                     `json:"resourceType,omitempty"`
	RoleMapEntries    []RoleMapEntry              `json:"roleMapEntries,omitempty"`
	Trash             *bool                       `json:"trash,omitempty"`
	TreeHref          *string                     `json:"treeHref,omitempty"`
	UnparentHref      *string                     `json:"unparentHref,omitempty"`
	ViewRef           *string                     `json:"viewRef,omitempty"`
}

// NewBTProjectInfo instantiates a new BTProjectInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTProjectInfo() *BTProjectInfo {
	this := BTProjectInfo{}
	return &this
}

// NewBTProjectInfoWithDefaults instantiates a new BTProjectInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTProjectInfoWithDefaults() *BTProjectInfo {
	this := BTProjectInfo{}
	return &this
}

// GetCanMove returns the CanMove field value if set, zero value otherwise.
func (o *BTProjectInfo) GetCanMove() bool {
	if o == nil || o.CanMove == nil {
		var ret bool
		return ret
	}
	return *o.CanMove
}

// GetCanMoveOk returns a tuple with the CanMove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTProjectInfo) GetCanMoveOk() (*bool, bool) {
	if o == nil || o.CanMove == nil {
		return nil, false
	}
	return o.CanMove, true
}

// HasCanMove returns a boolean if a field has been set.
func (o *BTProjectInfo) HasCanMove() bool {
	if o != nil && o.CanMove != nil {
		return true
	}

	return false
}

// SetCanMove gets a reference to the given bool and assigns it to the CanMove field.
func (o *BTProjectInfo) SetCanMove(v bool) {
	o.CanMove = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BTProjectInfo) GetCreatedAt() JSONTime {
	if o == nil || o.CreatedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTProjectInfo) GetCreatedAtOk() (*JSONTime, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BTProjectInfo) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given JSONTime and assigns it to the CreatedAt field.
func (o *BTProjectInfo) SetCreatedAt(v JSONTime) {
	o.CreatedAt = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *BTProjectInfo) GetCreatedBy() BTUserBasicSummaryInfo {
	if o == nil || o.CreatedBy == nil {
		var ret BTUserBasicSummaryInfo
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTProjectInfo) GetCreatedByOk() (*BTUserBasicSummaryInfo, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *BTProjectInfo) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given BTUserBasicSummaryInfo and assigns it to the CreatedBy field.
func (o *BTProjectInfo) SetCreatedBy(v BTUserBasicSummaryInfo) {
	o.CreatedBy = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTProjectInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTProjectInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTProjectInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTProjectInfo) SetDescription(v string) {
	o.Description = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTProjectInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTProjectInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTProjectInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTProjectInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTProjectInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTProjectInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTProjectInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTProjectInfo) SetId(v string) {
	o.Id = &v
}

// GetIsContainer returns the IsContainer field value if set, zero value otherwise.
func (o *BTProjectInfo) GetIsContainer() bool {
	if o == nil || o.IsContainer == nil {
		var ret bool
		return ret
	}
	return *o.IsContainer
}

// GetIsContainerOk returns a tuple with the IsContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTProjectInfo) GetIsContainerOk() (*bool, bool) {
	if o == nil || o.IsContainer == nil {
		return nil, false
	}
	return o.IsContainer, true
}

// HasIsContainer returns a boolean if a field has been set.
func (o *BTProjectInfo) HasIsContainer() bool {
	if o != nil && o.IsContainer != nil {
		return true
	}

	return false
}

// SetIsContainer gets a reference to the given bool and assigns it to the IsContainer field.
func (o *BTProjectInfo) SetIsContainer(v bool) {
	o.IsContainer = &v
}

// GetIsEnterpriseOwned returns the IsEnterpriseOwned field value if set, zero value otherwise.
func (o *BTProjectInfo) GetIsEnterpriseOwned() bool {
	if o == nil || o.IsEnterpriseOwned == nil {
		var ret bool
		return ret
	}
	return *o.IsEnterpriseOwned
}

// GetIsEnterpriseOwnedOk returns a tuple with the IsEnterpriseOwned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTProjectInfo) GetIsEnterpriseOwnedOk() (*bool, bool) {
	if o == nil || o.IsEnterpriseOwned == nil {
		return nil, false
	}
	return o.IsEnterpriseOwned, true
}

// HasIsEnterpriseOwned returns a boolean if a field has been set.
func (o *BTProjectInfo) HasIsEnterpriseOwned() bool {
	if o != nil && o.IsEnterpriseOwned != nil {
		return true
	}

	return false
}

// SetIsEnterpriseOwned gets a reference to the given bool and assigns it to the IsEnterpriseOwned field.
func (o *BTProjectInfo) SetIsEnterpriseOwned(v bool) {
	o.IsEnterpriseOwned = &v
}

// GetIsMutable returns the IsMutable field value if set, zero value otherwise.
func (o *BTProjectInfo) GetIsMutable() bool {
	if o == nil || o.IsMutable == nil {
		var ret bool
		return ret
	}
	return *o.IsMutable
}

// GetIsMutableOk returns a tuple with the IsMutable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTProjectInfo) GetIsMutableOk() (*bool, bool) {
	if o == nil || o.IsMutable == nil {
		return nil, false
	}
	return o.IsMutable, true
}

// HasIsMutable returns a boolean if a field has been set.
func (o *BTProjectInfo) HasIsMutable() bool {
	if o != nil && o.IsMutable != nil {
		return true
	}

	return false
}

// SetIsMutable gets a reference to the given bool and assigns it to the IsMutable field.
func (o *BTProjectInfo) SetIsMutable(v bool) {
	o.IsMutable = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *BTProjectInfo) GetModifiedAt() JSONTime {
	if o == nil || o.ModifiedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTProjectInfo) GetModifiedAtOk() (*JSONTime, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *BTProjectInfo) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given JSONTime and assigns it to the ModifiedAt field.
func (o *BTProjectInfo) SetModifiedAt(v JSONTime) {
	o.ModifiedAt = &v
}

// GetModifiedBy returns the ModifiedBy field value if set, zero value otherwise.
func (o *BTProjectInfo) GetModifiedBy() BTUserBasicSummaryInfo {
	if o == nil || o.ModifiedBy == nil {
		var ret BTUserBasicSummaryInfo
		return ret
	}
	return *o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTProjectInfo) GetModifiedByOk() (*BTUserBasicSummaryInfo, bool) {
	if o == nil || o.ModifiedBy == nil {
		return nil, false
	}
	return o.ModifiedBy, true
}

// HasModifiedBy returns a boolean if a field has been set.
func (o *BTProjectInfo) HasModifiedBy() bool {
	if o != nil && o.ModifiedBy != nil {
		return true
	}

	return false
}

// SetModifiedBy gets a reference to the given BTUserBasicSummaryInfo and assigns it to the ModifiedBy field.
func (o *BTProjectInfo) SetModifiedBy(v BTUserBasicSummaryInfo) {
	o.ModifiedBy = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTProjectInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTProjectInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTProjectInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTProjectInfo) SetName(v string) {
	o.Name = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *BTProjectInfo) GetOwner() BTOwnerInfo {
	if o == nil || o.Owner == nil {
		var ret BTOwnerInfo
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTProjectInfo) GetOwnerOk() (*BTOwnerInfo, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *BTProjectInfo) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given BTOwnerInfo and assigns it to the Owner field.
func (o *BTProjectInfo) SetOwner(v BTOwnerInfo) {
	o.Owner = &v
}

// GetPermissionScheme returns the PermissionScheme field value if set, zero value otherwise.
func (o *BTProjectInfo) GetPermissionScheme() BTRbacPermissionSchemeInfo {
	if o == nil || o.PermissionScheme == nil {
		var ret BTRbacPermissionSchemeInfo
		return ret
	}
	return *o.PermissionScheme
}

// GetPermissionSchemeOk returns a tuple with the PermissionScheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTProjectInfo) GetPermissionSchemeOk() (*BTRbacPermissionSchemeInfo, bool) {
	if o == nil || o.PermissionScheme == nil {
		return nil, false
	}
	return o.PermissionScheme, true
}

// HasPermissionScheme returns a boolean if a field has been set.
func (o *BTProjectInfo) HasPermissionScheme() bool {
	if o != nil && o.PermissionScheme != nil {
		return true
	}

	return false
}

// SetPermissionScheme gets a reference to the given BTRbacPermissionSchemeInfo and assigns it to the PermissionScheme field.
func (o *BTProjectInfo) SetPermissionScheme(v BTRbacPermissionSchemeInfo) {
	o.PermissionScheme = &v
}

// GetPermissionSet returns the PermissionSet field value if set, zero value otherwise.
func (o *BTProjectInfo) GetPermissionSet() []string {
	if o == nil || o.PermissionSet == nil {
		var ret []string
		return ret
	}
	return o.PermissionSet
}

// GetPermissionSetOk returns a tuple with the PermissionSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTProjectInfo) GetPermissionSetOk() ([]string, bool) {
	if o == nil || o.PermissionSet == nil {
		return nil, false
	}
	return o.PermissionSet, true
}

// HasPermissionSet returns a boolean if a field has been set.
func (o *BTProjectInfo) HasPermissionSet() bool {
	if o != nil && o.PermissionSet != nil {
		return true
	}

	return false
}

// SetPermissionSet gets a reference to the given []string and assigns it to the PermissionSet field.
func (o *BTProjectInfo) SetPermissionSet(v []string) {
	o.PermissionSet = v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *BTProjectInfo) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTProjectInfo) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *BTProjectInfo) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *BTProjectInfo) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *BTProjectInfo) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTProjectInfo) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *BTProjectInfo) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *BTProjectInfo) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetRoleMapEntries returns the RoleMapEntries field value if set, zero value otherwise.
func (o *BTProjectInfo) GetRoleMapEntries() []RoleMapEntry {
	if o == nil || o.RoleMapEntries == nil {
		var ret []RoleMapEntry
		return ret
	}
	return o.RoleMapEntries
}

// GetRoleMapEntriesOk returns a tuple with the RoleMapEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTProjectInfo) GetRoleMapEntriesOk() ([]RoleMapEntry, bool) {
	if o == nil || o.RoleMapEntries == nil {
		return nil, false
	}
	return o.RoleMapEntries, true
}

// HasRoleMapEntries returns a boolean if a field has been set.
func (o *BTProjectInfo) HasRoleMapEntries() bool {
	if o != nil && o.RoleMapEntries != nil {
		return true
	}

	return false
}

// SetRoleMapEntries gets a reference to the given []RoleMapEntry and assigns it to the RoleMapEntries field.
func (o *BTProjectInfo) SetRoleMapEntries(v []RoleMapEntry) {
	o.RoleMapEntries = v
}

// GetTrash returns the Trash field value if set, zero value otherwise.
func (o *BTProjectInfo) GetTrash() bool {
	if o == nil || o.Trash == nil {
		var ret bool
		return ret
	}
	return *o.Trash
}

// GetTrashOk returns a tuple with the Trash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTProjectInfo) GetTrashOk() (*bool, bool) {
	if o == nil || o.Trash == nil {
		return nil, false
	}
	return o.Trash, true
}

// HasTrash returns a boolean if a field has been set.
func (o *BTProjectInfo) HasTrash() bool {
	if o != nil && o.Trash != nil {
		return true
	}

	return false
}

// SetTrash gets a reference to the given bool and assigns it to the Trash field.
func (o *BTProjectInfo) SetTrash(v bool) {
	o.Trash = &v
}

// GetTreeHref returns the TreeHref field value if set, zero value otherwise.
func (o *BTProjectInfo) GetTreeHref() string {
	if o == nil || o.TreeHref == nil {
		var ret string
		return ret
	}
	return *o.TreeHref
}

// GetTreeHrefOk returns a tuple with the TreeHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTProjectInfo) GetTreeHrefOk() (*string, bool) {
	if o == nil || o.TreeHref == nil {
		return nil, false
	}
	return o.TreeHref, true
}

// HasTreeHref returns a boolean if a field has been set.
func (o *BTProjectInfo) HasTreeHref() bool {
	if o != nil && o.TreeHref != nil {
		return true
	}

	return false
}

// SetTreeHref gets a reference to the given string and assigns it to the TreeHref field.
func (o *BTProjectInfo) SetTreeHref(v string) {
	o.TreeHref = &v
}

// GetUnparentHref returns the UnparentHref field value if set, zero value otherwise.
func (o *BTProjectInfo) GetUnparentHref() string {
	if o == nil || o.UnparentHref == nil {
		var ret string
		return ret
	}
	return *o.UnparentHref
}

// GetUnparentHrefOk returns a tuple with the UnparentHref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTProjectInfo) GetUnparentHrefOk() (*string, bool) {
	if o == nil || o.UnparentHref == nil {
		return nil, false
	}
	return o.UnparentHref, true
}

// HasUnparentHref returns a boolean if a field has been set.
func (o *BTProjectInfo) HasUnparentHref() bool {
	if o != nil && o.UnparentHref != nil {
		return true
	}

	return false
}

// SetUnparentHref gets a reference to the given string and assigns it to the UnparentHref field.
func (o *BTProjectInfo) SetUnparentHref(v string) {
	o.UnparentHref = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTProjectInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTProjectInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTProjectInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTProjectInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTProjectInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanMove != nil {
		toSerialize["canMove"] = o.CanMove
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsContainer != nil {
		toSerialize["isContainer"] = o.IsContainer
	}
	if o.IsEnterpriseOwned != nil {
		toSerialize["isEnterpriseOwned"] = o.IsEnterpriseOwned
	}
	if o.IsMutable != nil {
		toSerialize["isMutable"] = o.IsMutable
	}
	if o.ModifiedAt != nil {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if o.ModifiedBy != nil {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.PermissionScheme != nil {
		toSerialize["permissionScheme"] = o.PermissionScheme
	}
	if o.PermissionSet != nil {
		toSerialize["permissionSet"] = o.PermissionSet
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.ResourceType != nil {
		toSerialize["resourceType"] = o.ResourceType
	}
	if o.RoleMapEntries != nil {
		toSerialize["roleMapEntries"] = o.RoleMapEntries
	}
	if o.Trash != nil {
		toSerialize["trash"] = o.Trash
	}
	if o.TreeHref != nil {
		toSerialize["treeHref"] = o.TreeHref
	}
	if o.UnparentHref != nil {
		toSerialize["unparentHref"] = o.UnparentHref
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTProjectInfo struct {
	value *BTProjectInfo
	isSet bool
}

func (v NullableBTProjectInfo) Get() *BTProjectInfo {
	return v.value
}

func (v *NullableBTProjectInfo) Set(val *BTProjectInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTProjectInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTProjectInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTProjectInfo(val *BTProjectInfo) *NullableBTProjectInfo {
	return &NullableBTProjectInfo{value: val, isSet: true}
}

func (v NullableBTProjectInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTProjectInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
