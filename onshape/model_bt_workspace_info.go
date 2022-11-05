/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.156.7192-0ed4c121c7d8
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTWorkspaceInfo struct for BTWorkspaceInfo
type BTWorkspaceInfo struct {
	CanDelete    *bool                   `json:"canDelete,omitempty"`
	CreatedAt    *JSONTime               `json:"createdAt,omitempty"`
	Creator      *BTUserBasicSummaryInfo `json:"creator,omitempty"`
	Description  *string                 `json:"description,omitempty"`
	DocumentId   *string                 `json:"documentId,omitempty"`
	Href         *string                 `json:"href,omitempty"`
	Id           *string                 `json:"id,omitempty"`
	IsReadOnly   *bool                   `json:"isReadOnly,omitempty"`
	LastModifier *BTUserBasicSummaryInfo `json:"lastModifier,omitempty"`
	Microversion *string                 `json:"microversion,omitempty"`
	ModifiedAt   *JSONTime               `json:"modifiedAt,omitempty"`
	Name         *string                 `json:"name,omitempty"`
	OverrideDate *JSONTime               `json:"overrideDate,omitempty"`
	Parent       *string                 `json:"parent,omitempty"`
	Parents      []BTVersionInfo         `json:"parents,omitempty"`
	Thumbnail    *BTThumbnailInfo        `json:"thumbnail,omitempty"`
	Type         *string                 `json:"type,omitempty"`
	ViewRef      *string                 `json:"viewRef,omitempty"`
}

// NewBTWorkspaceInfo instantiates a new BTWorkspaceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTWorkspaceInfo() *BTWorkspaceInfo {
	this := BTWorkspaceInfo{}
	return &this
}

// NewBTWorkspaceInfoWithDefaults instantiates a new BTWorkspaceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTWorkspaceInfoWithDefaults() *BTWorkspaceInfo {
	this := BTWorkspaceInfo{}
	return &this
}

// GetCanDelete returns the CanDelete field value if set, zero value otherwise.
func (o *BTWorkspaceInfo) GetCanDelete() bool {
	if o == nil || o.CanDelete == nil {
		var ret bool
		return ret
	}
	return *o.CanDelete
}

// GetCanDeleteOk returns a tuple with the CanDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceInfo) GetCanDeleteOk() (*bool, bool) {
	if o == nil || o.CanDelete == nil {
		return nil, false
	}
	return o.CanDelete, true
}

// HasCanDelete returns a boolean if a field has been set.
func (o *BTWorkspaceInfo) HasCanDelete() bool {
	if o != nil && o.CanDelete != nil {
		return true
	}

	return false
}

// SetCanDelete gets a reference to the given bool and assigns it to the CanDelete field.
func (o *BTWorkspaceInfo) SetCanDelete(v bool) {
	o.CanDelete = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BTWorkspaceInfo) GetCreatedAt() JSONTime {
	if o == nil || o.CreatedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceInfo) GetCreatedAtOk() (*JSONTime, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BTWorkspaceInfo) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given JSONTime and assigns it to the CreatedAt field.
func (o *BTWorkspaceInfo) SetCreatedAt(v JSONTime) {
	o.CreatedAt = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *BTWorkspaceInfo) GetCreator() BTUserBasicSummaryInfo {
	if o == nil || o.Creator == nil {
		var ret BTUserBasicSummaryInfo
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceInfo) GetCreatorOk() (*BTUserBasicSummaryInfo, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *BTWorkspaceInfo) HasCreator() bool {
	if o != nil && o.Creator != nil {
		return true
	}

	return false
}

// SetCreator gets a reference to the given BTUserBasicSummaryInfo and assigns it to the Creator field.
func (o *BTWorkspaceInfo) SetCreator(v BTUserBasicSummaryInfo) {
	o.Creator = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTWorkspaceInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTWorkspaceInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTWorkspaceInfo) SetDescription(v string) {
	o.Description = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTWorkspaceInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTWorkspaceInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTWorkspaceInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTWorkspaceInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTWorkspaceInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTWorkspaceInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTWorkspaceInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTWorkspaceInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTWorkspaceInfo) SetId(v string) {
	o.Id = &v
}

// GetIsReadOnly returns the IsReadOnly field value if set, zero value otherwise.
func (o *BTWorkspaceInfo) GetIsReadOnly() bool {
	if o == nil || o.IsReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.IsReadOnly
}

// GetIsReadOnlyOk returns a tuple with the IsReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceInfo) GetIsReadOnlyOk() (*bool, bool) {
	if o == nil || o.IsReadOnly == nil {
		return nil, false
	}
	return o.IsReadOnly, true
}

// HasIsReadOnly returns a boolean if a field has been set.
func (o *BTWorkspaceInfo) HasIsReadOnly() bool {
	if o != nil && o.IsReadOnly != nil {
		return true
	}

	return false
}

// SetIsReadOnly gets a reference to the given bool and assigns it to the IsReadOnly field.
func (o *BTWorkspaceInfo) SetIsReadOnly(v bool) {
	o.IsReadOnly = &v
}

// GetLastModifier returns the LastModifier field value if set, zero value otherwise.
func (o *BTWorkspaceInfo) GetLastModifier() BTUserBasicSummaryInfo {
	if o == nil || o.LastModifier == nil {
		var ret BTUserBasicSummaryInfo
		return ret
	}
	return *o.LastModifier
}

// GetLastModifierOk returns a tuple with the LastModifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceInfo) GetLastModifierOk() (*BTUserBasicSummaryInfo, bool) {
	if o == nil || o.LastModifier == nil {
		return nil, false
	}
	return o.LastModifier, true
}

// HasLastModifier returns a boolean if a field has been set.
func (o *BTWorkspaceInfo) HasLastModifier() bool {
	if o != nil && o.LastModifier != nil {
		return true
	}

	return false
}

// SetLastModifier gets a reference to the given BTUserBasicSummaryInfo and assigns it to the LastModifier field.
func (o *BTWorkspaceInfo) SetLastModifier(v BTUserBasicSummaryInfo) {
	o.LastModifier = &v
}

// GetMicroversion returns the Microversion field value if set, zero value otherwise.
func (o *BTWorkspaceInfo) GetMicroversion() string {
	if o == nil || o.Microversion == nil {
		var ret string
		return ret
	}
	return *o.Microversion
}

// GetMicroversionOk returns a tuple with the Microversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceInfo) GetMicroversionOk() (*string, bool) {
	if o == nil || o.Microversion == nil {
		return nil, false
	}
	return o.Microversion, true
}

// HasMicroversion returns a boolean if a field has been set.
func (o *BTWorkspaceInfo) HasMicroversion() bool {
	if o != nil && o.Microversion != nil {
		return true
	}

	return false
}

// SetMicroversion gets a reference to the given string and assigns it to the Microversion field.
func (o *BTWorkspaceInfo) SetMicroversion(v string) {
	o.Microversion = &v
}

// GetModifiedAt returns the ModifiedAt field value if set, zero value otherwise.
func (o *BTWorkspaceInfo) GetModifiedAt() JSONTime {
	if o == nil || o.ModifiedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceInfo) GetModifiedAtOk() (*JSONTime, bool) {
	if o == nil || o.ModifiedAt == nil {
		return nil, false
	}
	return o.ModifiedAt, true
}

// HasModifiedAt returns a boolean if a field has been set.
func (o *BTWorkspaceInfo) HasModifiedAt() bool {
	if o != nil && o.ModifiedAt != nil {
		return true
	}

	return false
}

// SetModifiedAt gets a reference to the given JSONTime and assigns it to the ModifiedAt field.
func (o *BTWorkspaceInfo) SetModifiedAt(v JSONTime) {
	o.ModifiedAt = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTWorkspaceInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTWorkspaceInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTWorkspaceInfo) SetName(v string) {
	o.Name = &v
}

// GetOverrideDate returns the OverrideDate field value if set, zero value otherwise.
func (o *BTWorkspaceInfo) GetOverrideDate() JSONTime {
	if o == nil || o.OverrideDate == nil {
		var ret JSONTime
		return ret
	}
	return *o.OverrideDate
}

// GetOverrideDateOk returns a tuple with the OverrideDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceInfo) GetOverrideDateOk() (*JSONTime, bool) {
	if o == nil || o.OverrideDate == nil {
		return nil, false
	}
	return o.OverrideDate, true
}

// HasOverrideDate returns a boolean if a field has been set.
func (o *BTWorkspaceInfo) HasOverrideDate() bool {
	if o != nil && o.OverrideDate != nil {
		return true
	}

	return false
}

// SetOverrideDate gets a reference to the given JSONTime and assigns it to the OverrideDate field.
func (o *BTWorkspaceInfo) SetOverrideDate(v JSONTime) {
	o.OverrideDate = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *BTWorkspaceInfo) GetParent() string {
	if o == nil || o.Parent == nil {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceInfo) GetParentOk() (*string, bool) {
	if o == nil || o.Parent == nil {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *BTWorkspaceInfo) HasParent() bool {
	if o != nil && o.Parent != nil {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *BTWorkspaceInfo) SetParent(v string) {
	o.Parent = &v
}

// GetParents returns the Parents field value if set, zero value otherwise.
func (o *BTWorkspaceInfo) GetParents() []BTVersionInfo {
	if o == nil || o.Parents == nil {
		var ret []BTVersionInfo
		return ret
	}
	return o.Parents
}

// GetParentsOk returns a tuple with the Parents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceInfo) GetParentsOk() ([]BTVersionInfo, bool) {
	if o == nil || o.Parents == nil {
		return nil, false
	}
	return o.Parents, true
}

// HasParents returns a boolean if a field has been set.
func (o *BTWorkspaceInfo) HasParents() bool {
	if o != nil && o.Parents != nil {
		return true
	}

	return false
}

// SetParents gets a reference to the given []BTVersionInfo and assigns it to the Parents field.
func (o *BTWorkspaceInfo) SetParents(v []BTVersionInfo) {
	o.Parents = v
}

// GetThumbnail returns the Thumbnail field value if set, zero value otherwise.
func (o *BTWorkspaceInfo) GetThumbnail() BTThumbnailInfo {
	if o == nil || o.Thumbnail == nil {
		var ret BTThumbnailInfo
		return ret
	}
	return *o.Thumbnail
}

// GetThumbnailOk returns a tuple with the Thumbnail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceInfo) GetThumbnailOk() (*BTThumbnailInfo, bool) {
	if o == nil || o.Thumbnail == nil {
		return nil, false
	}
	return o.Thumbnail, true
}

// HasThumbnail returns a boolean if a field has been set.
func (o *BTWorkspaceInfo) HasThumbnail() bool {
	if o != nil && o.Thumbnail != nil {
		return true
	}

	return false
}

// SetThumbnail gets a reference to the given BTThumbnailInfo and assigns it to the Thumbnail field.
func (o *BTWorkspaceInfo) SetThumbnail(v BTThumbnailInfo) {
	o.Thumbnail = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTWorkspaceInfo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceInfo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTWorkspaceInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTWorkspaceInfo) SetType(v string) {
	o.Type = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTWorkspaceInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTWorkspaceInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTWorkspaceInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTWorkspaceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanDelete != nil {
		toSerialize["canDelete"] = o.CanDelete
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsReadOnly != nil {
		toSerialize["isReadOnly"] = o.IsReadOnly
	}
	if o.LastModifier != nil {
		toSerialize["lastModifier"] = o.LastModifier
	}
	if o.Microversion != nil {
		toSerialize["microversion"] = o.Microversion
	}
	if o.ModifiedAt != nil {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OverrideDate != nil {
		toSerialize["overrideDate"] = o.OverrideDate
	}
	if o.Parent != nil {
		toSerialize["parent"] = o.Parent
	}
	if o.Parents != nil {
		toSerialize["parents"] = o.Parents
	}
	if o.Thumbnail != nil {
		toSerialize["thumbnail"] = o.Thumbnail
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTWorkspaceInfo struct {
	value *BTWorkspaceInfo
	isSet bool
}

func (v NullableBTWorkspaceInfo) Get() *BTWorkspaceInfo {
	return v.value
}

func (v *NullableBTWorkspaceInfo) Set(val *BTWorkspaceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTWorkspaceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTWorkspaceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTWorkspaceInfo(val *BTWorkspaceInfo) *NullableBTWorkspaceInfo {
	return &NullableBTWorkspaceInfo{value: val, isSet: true}
}

func (v NullableBTWorkspaceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTWorkspaceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
