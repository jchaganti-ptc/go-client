/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5928-bd774e9c9f97
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTRevisionMessageBody struct for BTRevisionMessageBody
type BTRevisionMessageBody struct {
	AppElementSessionId *string   `json:"appElementSessionId,omitempty"`
	CommentId           *string   `json:"commentId,omitempty"`
	Data                *string   `json:"data,omitempty"`
	DocumentId          *string   `json:"documentId,omitempty"`
	DocumentState       *string   `json:"documentState,omitempty"`
	ElementId           *string   `json:"elementId,omitempty"`
	Event               *string   `json:"event,omitempty"`
	MessageId           *string   `json:"messageId,omitempty"`
	MetadataObjectType  *string   `json:"metadataObjectType,omitempty"`
	PartId              *string   `json:"partId,omitempty"`
	PartIdentity        *string   `json:"partIdentity,omitempty"`
	PartNumber          *string   `json:"partNumber,omitempty"`
	Timestamp           *JSONTime `json:"timestamp,omitempty"`
	TranslatationId     *string   `json:"translatationId,omitempty"`
	TranslationId       *string   `json:"translationId,omitempty"`
	UserId              *string   `json:"userId,omitempty"`
	VersionId           *string   `json:"versionId,omitempty"`
	WebhookId           *string   `json:"webhookId,omitempty"`
	WorkspaceId         *string   `json:"workspaceId,omitempty"`
	ElementType         *int32    `json:"elementType,omitempty"`
	IsFromInitialState  *bool     `json:"isFromInitialState,omitempty"`
	IsToTerminalState   *bool     `json:"isToTerminalState,omitempty"`
	ReleaseId           *string   `json:"releaseId,omitempty"`
	Status              *string   `json:"status,omitempty"`
	TransitionName      *string   `json:"transitionName,omitempty"`
	RevisionId          *string   `json:"revisionId,omitempty"`
}

// NewBTRevisionMessageBody instantiates a new BTRevisionMessageBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTRevisionMessageBody() *BTRevisionMessageBody {
	this := BTRevisionMessageBody{}
	return &this
}

// NewBTRevisionMessageBodyWithDefaults instantiates a new BTRevisionMessageBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTRevisionMessageBodyWithDefaults() *BTRevisionMessageBody {
	this := BTRevisionMessageBody{}
	return &this
}

// GetAppElementSessionId returns the AppElementSessionId field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetAppElementSessionId() string {
	if o == nil || o.AppElementSessionId == nil {
		var ret string
		return ret
	}
	return *o.AppElementSessionId
}

// GetAppElementSessionIdOk returns a tuple with the AppElementSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetAppElementSessionIdOk() (*string, bool) {
	if o == nil || o.AppElementSessionId == nil {
		return nil, false
	}
	return o.AppElementSessionId, true
}

// HasAppElementSessionId returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasAppElementSessionId() bool {
	if o != nil && o.AppElementSessionId != nil {
		return true
	}

	return false
}

// SetAppElementSessionId gets a reference to the given string and assigns it to the AppElementSessionId field.
func (o *BTRevisionMessageBody) SetAppElementSessionId(v string) {
	o.AppElementSessionId = &v
}

// GetCommentId returns the CommentId field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetCommentId() string {
	if o == nil || o.CommentId == nil {
		var ret string
		return ret
	}
	return *o.CommentId
}

// GetCommentIdOk returns a tuple with the CommentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetCommentIdOk() (*string, bool) {
	if o == nil || o.CommentId == nil {
		return nil, false
	}
	return o.CommentId, true
}

// HasCommentId returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasCommentId() bool {
	if o != nil && o.CommentId != nil {
		return true
	}

	return false
}

// SetCommentId gets a reference to the given string and assigns it to the CommentId field.
func (o *BTRevisionMessageBody) SetCommentId(v string) {
	o.CommentId = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *BTRevisionMessageBody) SetData(v string) {
	o.Data = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTRevisionMessageBody) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetDocumentState returns the DocumentState field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetDocumentState() string {
	if o == nil || o.DocumentState == nil {
		var ret string
		return ret
	}
	return *o.DocumentState
}

// GetDocumentStateOk returns a tuple with the DocumentState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetDocumentStateOk() (*string, bool) {
	if o == nil || o.DocumentState == nil {
		return nil, false
	}
	return o.DocumentState, true
}

// HasDocumentState returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasDocumentState() bool {
	if o != nil && o.DocumentState != nil {
		return true
	}

	return false
}

// SetDocumentState gets a reference to the given string and assigns it to the DocumentState field.
func (o *BTRevisionMessageBody) SetDocumentState(v string) {
	o.DocumentState = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTRevisionMessageBody) SetElementId(v string) {
	o.ElementId = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetEvent() string {
	if o == nil || o.Event == nil {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetEventOk() (*string, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *BTRevisionMessageBody) SetEvent(v string) {
	o.Event = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetMessageId() string {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetMessageIdOk() (*string, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasMessageId() bool {
	if o != nil && o.MessageId != nil {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *BTRevisionMessageBody) SetMessageId(v string) {
	o.MessageId = &v
}

// GetMetadataObjectType returns the MetadataObjectType field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetMetadataObjectType() string {
	if o == nil || o.MetadataObjectType == nil {
		var ret string
		return ret
	}
	return *o.MetadataObjectType
}

// GetMetadataObjectTypeOk returns a tuple with the MetadataObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetMetadataObjectTypeOk() (*string, bool) {
	if o == nil || o.MetadataObjectType == nil {
		return nil, false
	}
	return o.MetadataObjectType, true
}

// HasMetadataObjectType returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasMetadataObjectType() bool {
	if o != nil && o.MetadataObjectType != nil {
		return true
	}

	return false
}

// SetMetadataObjectType gets a reference to the given string and assigns it to the MetadataObjectType field.
func (o *BTRevisionMessageBody) SetMetadataObjectType(v string) {
	o.MetadataObjectType = &v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *BTRevisionMessageBody) SetPartId(v string) {
	o.PartId = &v
}

// GetPartIdentity returns the PartIdentity field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetPartIdentity() string {
	if o == nil || o.PartIdentity == nil {
		var ret string
		return ret
	}
	return *o.PartIdentity
}

// GetPartIdentityOk returns a tuple with the PartIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetPartIdentityOk() (*string, bool) {
	if o == nil || o.PartIdentity == nil {
		return nil, false
	}
	return o.PartIdentity, true
}

// HasPartIdentity returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasPartIdentity() bool {
	if o != nil && o.PartIdentity != nil {
		return true
	}

	return false
}

// SetPartIdentity gets a reference to the given string and assigns it to the PartIdentity field.
func (o *BTRevisionMessageBody) SetPartIdentity(v string) {
	o.PartIdentity = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *BTRevisionMessageBody) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetTimestamp() JSONTime {
	if o == nil || o.Timestamp == nil {
		var ret JSONTime
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetTimestampOk() (*JSONTime, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given JSONTime and assigns it to the Timestamp field.
func (o *BTRevisionMessageBody) SetTimestamp(v JSONTime) {
	o.Timestamp = &v
}

// GetTranslatationId returns the TranslatationId field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetTranslatationId() string {
	if o == nil || o.TranslatationId == nil {
		var ret string
		return ret
	}
	return *o.TranslatationId
}

// GetTranslatationIdOk returns a tuple with the TranslatationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetTranslatationIdOk() (*string, bool) {
	if o == nil || o.TranslatationId == nil {
		return nil, false
	}
	return o.TranslatationId, true
}

// HasTranslatationId returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasTranslatationId() bool {
	if o != nil && o.TranslatationId != nil {
		return true
	}

	return false
}

// SetTranslatationId gets a reference to the given string and assigns it to the TranslatationId field.
func (o *BTRevisionMessageBody) SetTranslatationId(v string) {
	o.TranslatationId = &v
}

// GetTranslationId returns the TranslationId field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetTranslationId() string {
	if o == nil || o.TranslationId == nil {
		var ret string
		return ret
	}
	return *o.TranslationId
}

// GetTranslationIdOk returns a tuple with the TranslationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetTranslationIdOk() (*string, bool) {
	if o == nil || o.TranslationId == nil {
		return nil, false
	}
	return o.TranslationId, true
}

// HasTranslationId returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasTranslationId() bool {
	if o != nil && o.TranslationId != nil {
		return true
	}

	return false
}

// SetTranslationId gets a reference to the given string and assigns it to the TranslationId field.
func (o *BTRevisionMessageBody) SetTranslationId(v string) {
	o.TranslationId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *BTRevisionMessageBody) SetUserId(v string) {
	o.UserId = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *BTRevisionMessageBody) SetVersionId(v string) {
	o.VersionId = &v
}

// GetWebhookId returns the WebhookId field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetWebhookId() string {
	if o == nil || o.WebhookId == nil {
		var ret string
		return ret
	}
	return *o.WebhookId
}

// GetWebhookIdOk returns a tuple with the WebhookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetWebhookIdOk() (*string, bool) {
	if o == nil || o.WebhookId == nil {
		return nil, false
	}
	return o.WebhookId, true
}

// HasWebhookId returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasWebhookId() bool {
	if o != nil && o.WebhookId != nil {
		return true
	}

	return false
}

// SetWebhookId gets a reference to the given string and assigns it to the WebhookId field.
func (o *BTRevisionMessageBody) SetWebhookId(v string) {
	o.WebhookId = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *BTRevisionMessageBody) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

// GetElementType returns the ElementType field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetElementType() int32 {
	if o == nil || o.ElementType == nil {
		var ret int32
		return ret
	}
	return *o.ElementType
}

// GetElementTypeOk returns a tuple with the ElementType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetElementTypeOk() (*int32, bool) {
	if o == nil || o.ElementType == nil {
		return nil, false
	}
	return o.ElementType, true
}

// HasElementType returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasElementType() bool {
	if o != nil && o.ElementType != nil {
		return true
	}

	return false
}

// SetElementType gets a reference to the given int32 and assigns it to the ElementType field.
func (o *BTRevisionMessageBody) SetElementType(v int32) {
	o.ElementType = &v
}

// GetIsFromInitialState returns the IsFromInitialState field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetIsFromInitialState() bool {
	if o == nil || o.IsFromInitialState == nil {
		var ret bool
		return ret
	}
	return *o.IsFromInitialState
}

// GetIsFromInitialStateOk returns a tuple with the IsFromInitialState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetIsFromInitialStateOk() (*bool, bool) {
	if o == nil || o.IsFromInitialState == nil {
		return nil, false
	}
	return o.IsFromInitialState, true
}

// HasIsFromInitialState returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasIsFromInitialState() bool {
	if o != nil && o.IsFromInitialState != nil {
		return true
	}

	return false
}

// SetIsFromInitialState gets a reference to the given bool and assigns it to the IsFromInitialState field.
func (o *BTRevisionMessageBody) SetIsFromInitialState(v bool) {
	o.IsFromInitialState = &v
}

// GetIsToTerminalState returns the IsToTerminalState field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetIsToTerminalState() bool {
	if o == nil || o.IsToTerminalState == nil {
		var ret bool
		return ret
	}
	return *o.IsToTerminalState
}

// GetIsToTerminalStateOk returns a tuple with the IsToTerminalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetIsToTerminalStateOk() (*bool, bool) {
	if o == nil || o.IsToTerminalState == nil {
		return nil, false
	}
	return o.IsToTerminalState, true
}

// HasIsToTerminalState returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasIsToTerminalState() bool {
	if o != nil && o.IsToTerminalState != nil {
		return true
	}

	return false
}

// SetIsToTerminalState gets a reference to the given bool and assigns it to the IsToTerminalState field.
func (o *BTRevisionMessageBody) SetIsToTerminalState(v bool) {
	o.IsToTerminalState = &v
}

// GetReleaseId returns the ReleaseId field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetReleaseId() string {
	if o == nil || o.ReleaseId == nil {
		var ret string
		return ret
	}
	return *o.ReleaseId
}

// GetReleaseIdOk returns a tuple with the ReleaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetReleaseIdOk() (*string, bool) {
	if o == nil || o.ReleaseId == nil {
		return nil, false
	}
	return o.ReleaseId, true
}

// HasReleaseId returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasReleaseId() bool {
	if o != nil && o.ReleaseId != nil {
		return true
	}

	return false
}

// SetReleaseId gets a reference to the given string and assigns it to the ReleaseId field.
func (o *BTRevisionMessageBody) SetReleaseId(v string) {
	o.ReleaseId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BTRevisionMessageBody) SetStatus(v string) {
	o.Status = &v
}

// GetTransitionName returns the TransitionName field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetTransitionName() string {
	if o == nil || o.TransitionName == nil {
		var ret string
		return ret
	}
	return *o.TransitionName
}

// GetTransitionNameOk returns a tuple with the TransitionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetTransitionNameOk() (*string, bool) {
	if o == nil || o.TransitionName == nil {
		return nil, false
	}
	return o.TransitionName, true
}

// HasTransitionName returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasTransitionName() bool {
	if o != nil && o.TransitionName != nil {
		return true
	}

	return false
}

// SetTransitionName gets a reference to the given string and assigns it to the TransitionName field.
func (o *BTRevisionMessageBody) SetTransitionName(v string) {
	o.TransitionName = &v
}

// GetRevisionId returns the RevisionId field value if set, zero value otherwise.
func (o *BTRevisionMessageBody) GetRevisionId() string {
	if o == nil || o.RevisionId == nil {
		var ret string
		return ret
	}
	return *o.RevisionId
}

// GetRevisionIdOk returns a tuple with the RevisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTRevisionMessageBody) GetRevisionIdOk() (*string, bool) {
	if o == nil || o.RevisionId == nil {
		return nil, false
	}
	return o.RevisionId, true
}

// HasRevisionId returns a boolean if a field has been set.
func (o *BTRevisionMessageBody) HasRevisionId() bool {
	if o != nil && o.RevisionId != nil {
		return true
	}

	return false
}

// SetRevisionId gets a reference to the given string and assigns it to the RevisionId field.
func (o *BTRevisionMessageBody) SetRevisionId(v string) {
	o.RevisionId = &v
}

func (o BTRevisionMessageBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppElementSessionId != nil {
		toSerialize["appElementSessionId"] = o.AppElementSessionId
	}
	if o.CommentId != nil {
		toSerialize["commentId"] = o.CommentId
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.DocumentState != nil {
		toSerialize["documentState"] = o.DocumentState
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	if o.MessageId != nil {
		toSerialize["messageId"] = o.MessageId
	}
	if o.MetadataObjectType != nil {
		toSerialize["metadataObjectType"] = o.MetadataObjectType
	}
	if o.PartId != nil {
		toSerialize["partId"] = o.PartId
	}
	if o.PartIdentity != nil {
		toSerialize["partIdentity"] = o.PartIdentity
	}
	if o.PartNumber != nil {
		toSerialize["partNumber"] = o.PartNumber
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.TranslatationId != nil {
		toSerialize["translatationId"] = o.TranslatationId
	}
	if o.TranslationId != nil {
		toSerialize["translationId"] = o.TranslationId
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.VersionId != nil {
		toSerialize["versionId"] = o.VersionId
	}
	if o.WebhookId != nil {
		toSerialize["webhookId"] = o.WebhookId
	}
	if o.WorkspaceId != nil {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	if o.ElementType != nil {
		toSerialize["elementType"] = o.ElementType
	}
	if o.IsFromInitialState != nil {
		toSerialize["isFromInitialState"] = o.IsFromInitialState
	}
	if o.IsToTerminalState != nil {
		toSerialize["isToTerminalState"] = o.IsToTerminalState
	}
	if o.ReleaseId != nil {
		toSerialize["releaseId"] = o.ReleaseId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TransitionName != nil {
		toSerialize["transitionName"] = o.TransitionName
	}
	if o.RevisionId != nil {
		toSerialize["revisionId"] = o.RevisionId
	}
	return json.Marshal(toSerialize)
}

type NullableBTRevisionMessageBody struct {
	value *BTRevisionMessageBody
	isSet bool
}

func (v NullableBTRevisionMessageBody) Get() *BTRevisionMessageBody {
	return v.value
}

func (v *NullableBTRevisionMessageBody) Set(val *BTRevisionMessageBody) {
	v.value = val
	v.isSet = true
}

func (v NullableBTRevisionMessageBody) IsSet() bool {
	return v.isSet
}

func (v *NullableBTRevisionMessageBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTRevisionMessageBody(val *BTRevisionMessageBody) *NullableBTRevisionMessageBody {
	return &NullableBTRevisionMessageBody{value: val, isSet: true}
}

func (v NullableBTRevisionMessageBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTRevisionMessageBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
