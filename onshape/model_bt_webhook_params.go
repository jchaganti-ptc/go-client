/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.165.17983-3c4f6e999748
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTWebhookParams struct for BTWebhookParams
type BTWebhookParams struct {
	ClientId *string `json:"clientId,omitempty"`
	// Company admins can register webhooks to listen to all company events.
	CompanyId   *string `json:"companyId,omitempty"`
	Data        *string `json:"data,omitempty"`
	Description *string `json:"description,omitempty"`
	DocumentId  *string `json:"documentId,omitempty"`
	ElementId   *string `json:"elementId,omitempty"`
	// List of events for which webhook callback is invoked.
	Events []string `json:"events,omitempty"`
	// Applications can pass this parameter as X-Session-ID with every REST call to distinguish webhooks triggered by self.
	ExternalSessionId *string `json:"externalSessionId,omitempty"`
	Filter            *string `json:"filter,omitempty"`
	FolderId          *string `json:"folderId,omitempty"`
	Id                *string `json:"id,omitempty"`
	// Transient webhooks are automatically cleaned up after a period of inactivity.
	IsTransient    *bool             `json:"isTransient,omitempty"`
	LinkDocumentId *string           `json:"linkDocumentId,omitempty"`
	Options        *BTWebhookOptions `json:"options,omitempty"`
	PartId         *string           `json:"partId,omitempty"`
	ProjectId      *string           `json:"projectId,omitempty"`
	Url            *string           `json:"url,omitempty"`
	UserId         *string           `json:"userId,omitempty"`
	VersionId      *string           `json:"versionId,omitempty"`
	WorkspaceId    *string           `json:"workspaceId,omitempty"`
}

// NewBTWebhookParams instantiates a new BTWebhookParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTWebhookParams() *BTWebhookParams {
	this := BTWebhookParams{}
	return &this
}

// NewBTWebhookParamsWithDefaults instantiates a new BTWebhookParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTWebhookParamsWithDefaults() *BTWebhookParams {
	this := BTWebhookParams{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *BTWebhookParams) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookParams) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *BTWebhookParams) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *BTWebhookParams) SetClientId(v string) {
	o.ClientId = &v
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *BTWebhookParams) GetCompanyId() string {
	if o == nil || o.CompanyId == nil {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookParams) GetCompanyIdOk() (*string, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *BTWebhookParams) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *BTWebhookParams) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BTWebhookParams) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookParams) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BTWebhookParams) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *BTWebhookParams) SetData(v string) {
	o.Data = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BTWebhookParams) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookParams) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BTWebhookParams) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BTWebhookParams) SetDescription(v string) {
	o.Description = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTWebhookParams) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookParams) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTWebhookParams) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTWebhookParams) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTWebhookParams) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookParams) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTWebhookParams) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTWebhookParams) SetElementId(v string) {
	o.ElementId = &v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *BTWebhookParams) GetEvents() []string {
	if o == nil || o.Events == nil {
		var ret []string
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookParams) GetEventsOk() ([]string, bool) {
	if o == nil || o.Events == nil {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *BTWebhookParams) HasEvents() bool {
	if o != nil && o.Events != nil {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []string and assigns it to the Events field.
func (o *BTWebhookParams) SetEvents(v []string) {
	o.Events = v
}

// GetExternalSessionId returns the ExternalSessionId field value if set, zero value otherwise.
func (o *BTWebhookParams) GetExternalSessionId() string {
	if o == nil || o.ExternalSessionId == nil {
		var ret string
		return ret
	}
	return *o.ExternalSessionId
}

// GetExternalSessionIdOk returns a tuple with the ExternalSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookParams) GetExternalSessionIdOk() (*string, bool) {
	if o == nil || o.ExternalSessionId == nil {
		return nil, false
	}
	return o.ExternalSessionId, true
}

// HasExternalSessionId returns a boolean if a field has been set.
func (o *BTWebhookParams) HasExternalSessionId() bool {
	if o != nil && o.ExternalSessionId != nil {
		return true
	}

	return false
}

// SetExternalSessionId gets a reference to the given string and assigns it to the ExternalSessionId field.
func (o *BTWebhookParams) SetExternalSessionId(v string) {
	o.ExternalSessionId = &v
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *BTWebhookParams) GetFilter() string {
	if o == nil || o.Filter == nil {
		var ret string
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookParams) GetFilterOk() (*string, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *BTWebhookParams) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given string and assigns it to the Filter field.
func (o *BTWebhookParams) SetFilter(v string) {
	o.Filter = &v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise.
func (o *BTWebhookParams) GetFolderId() string {
	if o == nil || o.FolderId == nil {
		var ret string
		return ret
	}
	return *o.FolderId
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookParams) GetFolderIdOk() (*string, bool) {
	if o == nil || o.FolderId == nil {
		return nil, false
	}
	return o.FolderId, true
}

// HasFolderId returns a boolean if a field has been set.
func (o *BTWebhookParams) HasFolderId() bool {
	if o != nil && o.FolderId != nil {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given string and assigns it to the FolderId field.
func (o *BTWebhookParams) SetFolderId(v string) {
	o.FolderId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTWebhookParams) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookParams) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTWebhookParams) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTWebhookParams) SetId(v string) {
	o.Id = &v
}

// GetIsTransient returns the IsTransient field value if set, zero value otherwise.
func (o *BTWebhookParams) GetIsTransient() bool {
	if o == nil || o.IsTransient == nil {
		var ret bool
		return ret
	}
	return *o.IsTransient
}

// GetIsTransientOk returns a tuple with the IsTransient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookParams) GetIsTransientOk() (*bool, bool) {
	if o == nil || o.IsTransient == nil {
		return nil, false
	}
	return o.IsTransient, true
}

// HasIsTransient returns a boolean if a field has been set.
func (o *BTWebhookParams) HasIsTransient() bool {
	if o != nil && o.IsTransient != nil {
		return true
	}

	return false
}

// SetIsTransient gets a reference to the given bool and assigns it to the IsTransient field.
func (o *BTWebhookParams) SetIsTransient(v bool) {
	o.IsTransient = &v
}

// GetLinkDocumentId returns the LinkDocumentId field value if set, zero value otherwise.
func (o *BTWebhookParams) GetLinkDocumentId() string {
	if o == nil || o.LinkDocumentId == nil {
		var ret string
		return ret
	}
	return *o.LinkDocumentId
}

// GetLinkDocumentIdOk returns a tuple with the LinkDocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookParams) GetLinkDocumentIdOk() (*string, bool) {
	if o == nil || o.LinkDocumentId == nil {
		return nil, false
	}
	return o.LinkDocumentId, true
}

// HasLinkDocumentId returns a boolean if a field has been set.
func (o *BTWebhookParams) HasLinkDocumentId() bool {
	if o != nil && o.LinkDocumentId != nil {
		return true
	}

	return false
}

// SetLinkDocumentId gets a reference to the given string and assigns it to the LinkDocumentId field.
func (o *BTWebhookParams) SetLinkDocumentId(v string) {
	o.LinkDocumentId = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *BTWebhookParams) GetOptions() BTWebhookOptions {
	if o == nil || o.Options == nil {
		var ret BTWebhookOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookParams) GetOptionsOk() (*BTWebhookOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *BTWebhookParams) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given BTWebhookOptions and assigns it to the Options field.
func (o *BTWebhookParams) SetOptions(v BTWebhookOptions) {
	o.Options = &v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *BTWebhookParams) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookParams) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *BTWebhookParams) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *BTWebhookParams) SetPartId(v string) {
	o.PartId = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *BTWebhookParams) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookParams) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *BTWebhookParams) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *BTWebhookParams) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *BTWebhookParams) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookParams) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *BTWebhookParams) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *BTWebhookParams) SetUrl(v string) {
	o.Url = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *BTWebhookParams) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookParams) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *BTWebhookParams) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *BTWebhookParams) SetUserId(v string) {
	o.UserId = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *BTWebhookParams) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookParams) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *BTWebhookParams) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *BTWebhookParams) SetVersionId(v string) {
	o.VersionId = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *BTWebhookParams) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWebhookParams) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *BTWebhookParams) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *BTWebhookParams) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

func (o BTWebhookParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.CompanyId != nil {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.Events != nil {
		toSerialize["events"] = o.Events
	}
	if o.ExternalSessionId != nil {
		toSerialize["externalSessionId"] = o.ExternalSessionId
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if o.FolderId != nil {
		toSerialize["folderId"] = o.FolderId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsTransient != nil {
		toSerialize["isTransient"] = o.IsTransient
	}
	if o.LinkDocumentId != nil {
		toSerialize["linkDocumentId"] = o.LinkDocumentId
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.PartId != nil {
		toSerialize["partId"] = o.PartId
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.VersionId != nil {
		toSerialize["versionId"] = o.VersionId
	}
	if o.WorkspaceId != nil {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableBTWebhookParams struct {
	value *BTWebhookParams
	isSet bool
}

func (v NullableBTWebhookParams) Get() *BTWebhookParams {
	return v.value
}

func (v *NullableBTWebhookParams) Set(val *BTWebhookParams) {
	v.value = val
	v.isSet = true
}

func (v NullableBTWebhookParams) IsSet() bool {
	return v.isSet
}

func (v *NullableBTWebhookParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTWebhookParams(val *BTWebhookParams) *NullableBTWebhookParams {
	return &NullableBTWebhookParams{value: val, isSet: true}
}

func (v NullableBTWebhookParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTWebhookParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
