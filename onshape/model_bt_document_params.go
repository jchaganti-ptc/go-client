/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.28658-06d4d4923fc7
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTDocumentParams Parameters for creating and updating documents.
type BTDocumentParams struct {
	// Document description.
	Description string `json:"description"`
	// List of element IDs to include in the document.
	Elements []BTDocumentElementCreationDescriptor `json:"elements,omitempty"`
	// `true` if the current user can toggle the Force Export Rule flag on a document.
	ForceExportRules *bool `json:"forceExportRules,omitempty"`
	// Set to `true` for debugging.
	GenerateUnknownMessages *bool `json:"generateUnknownMessages,omitempty"`
	// Set to `true` to generate an empty document.
	IsEmptyContent *bool `json:"isEmptyContent,omitempty"`
	// Set to `true` to make the document public.
	IsPublic *bool `json:"isPublic,omitempty"`
	// Document name.
	Name string `json:"name"`
	// Set to `true` to indicate that revisions are not managed for this document.
	NotRevisionManaged *bool `json:"notRevisionManaged,omitempty"`
	// The document owner's email address.
	OwnerEmail *string `json:"ownerEmail,omitempty"`
	// If `ownerType=USER`, this is the user ID. If `ownerType=COMPANY`, this is the company ID.
	OwnerId *string `json:"ownerId,omitempty"`
	// The document's owner type. `USER=0` | `COMPANY=1` | `ONSHAPE=2`
	OwnerType *int32 `json:"ownerType,omitempty"`
	// Document ID of this document's parent.
	ParentId *string `json:"parentId,omitempty"`
	// ID of the project this document belongs to.
	ProjectId *string `json:"projectId,omitempty"`
	// Array of strings to set as tags for the document.
	Tags []string `json:"tags,omitempty"`
}

// NewBTDocumentParams instantiates a new BTDocumentParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDocumentParams(description string, name string) *BTDocumentParams {
	this := BTDocumentParams{}
	this.Description = description
	this.Name = name
	return &this
}

// NewBTDocumentParamsWithDefaults instantiates a new BTDocumentParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDocumentParamsWithDefaults() *BTDocumentParams {
	this := BTDocumentParams{}
	return &this
}

// GetDescription returns the Description field value
func (o *BTDocumentParams) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *BTDocumentParams) SetDescription(v string) {
	o.Description = v
}

// GetElements returns the Elements field value if set, zero value otherwise.
func (o *BTDocumentParams) GetElements() []BTDocumentElementCreationDescriptor {
	if o == nil || o.Elements == nil {
		var ret []BTDocumentElementCreationDescriptor
		return ret
	}
	return o.Elements
}

// GetElementsOk returns a tuple with the Elements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetElementsOk() ([]BTDocumentElementCreationDescriptor, bool) {
	if o == nil || o.Elements == nil {
		return nil, false
	}
	return o.Elements, true
}

// HasElements returns a boolean if a field has been set.
func (o *BTDocumentParams) HasElements() bool {
	if o != nil && o.Elements != nil {
		return true
	}

	return false
}

// SetElements gets a reference to the given []BTDocumentElementCreationDescriptor and assigns it to the Elements field.
func (o *BTDocumentParams) SetElements(v []BTDocumentElementCreationDescriptor) {
	o.Elements = v
}

// GetForceExportRules returns the ForceExportRules field value if set, zero value otherwise.
func (o *BTDocumentParams) GetForceExportRules() bool {
	if o == nil || o.ForceExportRules == nil {
		var ret bool
		return ret
	}
	return *o.ForceExportRules
}

// GetForceExportRulesOk returns a tuple with the ForceExportRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetForceExportRulesOk() (*bool, bool) {
	if o == nil || o.ForceExportRules == nil {
		return nil, false
	}
	return o.ForceExportRules, true
}

// HasForceExportRules returns a boolean if a field has been set.
func (o *BTDocumentParams) HasForceExportRules() bool {
	if o != nil && o.ForceExportRules != nil {
		return true
	}

	return false
}

// SetForceExportRules gets a reference to the given bool and assigns it to the ForceExportRules field.
func (o *BTDocumentParams) SetForceExportRules(v bool) {
	o.ForceExportRules = &v
}

// GetGenerateUnknownMessages returns the GenerateUnknownMessages field value if set, zero value otherwise.
func (o *BTDocumentParams) GetGenerateUnknownMessages() bool {
	if o == nil || o.GenerateUnknownMessages == nil {
		var ret bool
		return ret
	}
	return *o.GenerateUnknownMessages
}

// GetGenerateUnknownMessagesOk returns a tuple with the GenerateUnknownMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetGenerateUnknownMessagesOk() (*bool, bool) {
	if o == nil || o.GenerateUnknownMessages == nil {
		return nil, false
	}
	return o.GenerateUnknownMessages, true
}

// HasGenerateUnknownMessages returns a boolean if a field has been set.
func (o *BTDocumentParams) HasGenerateUnknownMessages() bool {
	if o != nil && o.GenerateUnknownMessages != nil {
		return true
	}

	return false
}

// SetGenerateUnknownMessages gets a reference to the given bool and assigns it to the GenerateUnknownMessages field.
func (o *BTDocumentParams) SetGenerateUnknownMessages(v bool) {
	o.GenerateUnknownMessages = &v
}

// GetIsEmptyContent returns the IsEmptyContent field value if set, zero value otherwise.
func (o *BTDocumentParams) GetIsEmptyContent() bool {
	if o == nil || o.IsEmptyContent == nil {
		var ret bool
		return ret
	}
	return *o.IsEmptyContent
}

// GetIsEmptyContentOk returns a tuple with the IsEmptyContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetIsEmptyContentOk() (*bool, bool) {
	if o == nil || o.IsEmptyContent == nil {
		return nil, false
	}
	return o.IsEmptyContent, true
}

// HasIsEmptyContent returns a boolean if a field has been set.
func (o *BTDocumentParams) HasIsEmptyContent() bool {
	if o != nil && o.IsEmptyContent != nil {
		return true
	}

	return false
}

// SetIsEmptyContent gets a reference to the given bool and assigns it to the IsEmptyContent field.
func (o *BTDocumentParams) SetIsEmptyContent(v bool) {
	o.IsEmptyContent = &v
}

// GetIsPublic returns the IsPublic field value if set, zero value otherwise.
func (o *BTDocumentParams) GetIsPublic() bool {
	if o == nil || o.IsPublic == nil {
		var ret bool
		return ret
	}
	return *o.IsPublic
}

// GetIsPublicOk returns a tuple with the IsPublic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetIsPublicOk() (*bool, bool) {
	if o == nil || o.IsPublic == nil {
		return nil, false
	}
	return o.IsPublic, true
}

// HasIsPublic returns a boolean if a field has been set.
func (o *BTDocumentParams) HasIsPublic() bool {
	if o != nil && o.IsPublic != nil {
		return true
	}

	return false
}

// SetIsPublic gets a reference to the given bool and assigns it to the IsPublic field.
func (o *BTDocumentParams) SetIsPublic(v bool) {
	o.IsPublic = &v
}

// GetName returns the Name field value
func (o *BTDocumentParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BTDocumentParams) SetName(v string) {
	o.Name = v
}

// GetNotRevisionManaged returns the NotRevisionManaged field value if set, zero value otherwise.
func (o *BTDocumentParams) GetNotRevisionManaged() bool {
	if o == nil || o.NotRevisionManaged == nil {
		var ret bool
		return ret
	}
	return *o.NotRevisionManaged
}

// GetNotRevisionManagedOk returns a tuple with the NotRevisionManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetNotRevisionManagedOk() (*bool, bool) {
	if o == nil || o.NotRevisionManaged == nil {
		return nil, false
	}
	return o.NotRevisionManaged, true
}

// HasNotRevisionManaged returns a boolean if a field has been set.
func (o *BTDocumentParams) HasNotRevisionManaged() bool {
	if o != nil && o.NotRevisionManaged != nil {
		return true
	}

	return false
}

// SetNotRevisionManaged gets a reference to the given bool and assigns it to the NotRevisionManaged field.
func (o *BTDocumentParams) SetNotRevisionManaged(v bool) {
	o.NotRevisionManaged = &v
}

// GetOwnerEmail returns the OwnerEmail field value if set, zero value otherwise.
func (o *BTDocumentParams) GetOwnerEmail() string {
	if o == nil || o.OwnerEmail == nil {
		var ret string
		return ret
	}
	return *o.OwnerEmail
}

// GetOwnerEmailOk returns a tuple with the OwnerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetOwnerEmailOk() (*string, bool) {
	if o == nil || o.OwnerEmail == nil {
		return nil, false
	}
	return o.OwnerEmail, true
}

// HasOwnerEmail returns a boolean if a field has been set.
func (o *BTDocumentParams) HasOwnerEmail() bool {
	if o != nil && o.OwnerEmail != nil {
		return true
	}

	return false
}

// SetOwnerEmail gets a reference to the given string and assigns it to the OwnerEmail field.
func (o *BTDocumentParams) SetOwnerEmail(v string) {
	o.OwnerEmail = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *BTDocumentParams) GetOwnerId() string {
	if o == nil || o.OwnerId == nil {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetOwnerIdOk() (*string, bool) {
	if o == nil || o.OwnerId == nil {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *BTDocumentParams) HasOwnerId() bool {
	if o != nil && o.OwnerId != nil {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *BTDocumentParams) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetOwnerType returns the OwnerType field value if set, zero value otherwise.
func (o *BTDocumentParams) GetOwnerType() int32 {
	if o == nil || o.OwnerType == nil {
		var ret int32
		return ret
	}
	return *o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetOwnerTypeOk() (*int32, bool) {
	if o == nil || o.OwnerType == nil {
		return nil, false
	}
	return o.OwnerType, true
}

// HasOwnerType returns a boolean if a field has been set.
func (o *BTDocumentParams) HasOwnerType() bool {
	if o != nil && o.OwnerType != nil {
		return true
	}

	return false
}

// SetOwnerType gets a reference to the given int32 and assigns it to the OwnerType field.
func (o *BTDocumentParams) SetOwnerType(v int32) {
	o.OwnerType = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *BTDocumentParams) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *BTDocumentParams) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *BTDocumentParams) SetParentId(v string) {
	o.ParentId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *BTDocumentParams) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *BTDocumentParams) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *BTDocumentParams) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *BTDocumentParams) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentParams) GetTagsOk() ([]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *BTDocumentParams) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *BTDocumentParams) SetTags(v []string) {
	o.Tags = v
}

func (o BTDocumentParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.Elements != nil {
		toSerialize["elements"] = o.Elements
	}
	if o.ForceExportRules != nil {
		toSerialize["forceExportRules"] = o.ForceExportRules
	}
	if o.GenerateUnknownMessages != nil {
		toSerialize["generateUnknownMessages"] = o.GenerateUnknownMessages
	}
	if o.IsEmptyContent != nil {
		toSerialize["isEmptyContent"] = o.IsEmptyContent
	}
	if o.IsPublic != nil {
		toSerialize["isPublic"] = o.IsPublic
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.NotRevisionManaged != nil {
		toSerialize["notRevisionManaged"] = o.NotRevisionManaged
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
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableBTDocumentParams struct {
	value *BTDocumentParams
	isSet bool
}

func (v NullableBTDocumentParams) Get() *BTDocumentParams {
	return v.value
}

func (v *NullableBTDocumentParams) Set(val *BTDocumentParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDocumentParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDocumentParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDocumentParams(val *BTDocumentParams) *NullableBTDocumentParams {
	return &NullableBTDocumentParams{value: val, isSet: true}
}

func (v NullableBTDocumentParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDocumentParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
