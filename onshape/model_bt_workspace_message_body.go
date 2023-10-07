/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.170.23604-b59b123004e9
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTWorkspaceMessageBody struct for BTWorkspaceMessageBody
type BTWorkspaceMessageBody struct {
	AppElementSessionId *string   `json:"appElementSessionId,omitempty"`
	CommentId           *string   `json:"commentId,omitempty"`
	Data                *string   `json:"data,omitempty"`
	DocumentId          *string   `json:"documentId,omitempty"`
	DocumentState       *string   `json:"documentState,omitempty"`
	DocumentType        *int32    `json:"documentType,omitempty"`
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
	// The resultant document microverion if applicable created due to workspace modification.
	DocumentMicroversionId *string `json:"documentMicroversionId,omitempty"`
}

// NewBTWorkspaceMessageBody instantiates a new BTWorkspaceMessageBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTWorkspaceMessageBody() *BTWorkspaceMessageBody {
	this := BTWorkspaceMessageBody{}
	return &this
}

// NewBTWorkspaceMessageBodyWithDefaults instantiates a new BTWorkspaceMessageBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTWorkspaceMessageBodyWithDefaults() *BTWorkspaceMessageBody {
	this := BTWorkspaceMessageBody{}
	return &this
}

// GetAppElementSessionId returns the AppElementSessionId field value if set, zero value otherwise.
func (o *BTWorkspaceMessageBody) GetAppElementSessionId() string {
	if o == nil || o.AppElementSessionId == nil {
		var ret string
		return ret
	}
	return *o.AppElementSessionId
}

// GetAppElementSessionIdOk returns a tuple with the AppElementSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceMessageBody) GetAppElementSessionIdOk() (*string, bool) {
	if o == nil || o.AppElementSessionId == nil {
		return nil, false
	}
	return o.AppElementSessionId, true
}

// HasAppElementSessionId returns a boolean if a field has been set.
func (o *BTWorkspaceMessageBody) HasAppElementSessionId() bool {
	if o != nil && o.AppElementSessionId != nil {
		return true
	}

	return false
}

// SetAppElementSessionId gets a reference to the given string and assigns it to the AppElementSessionId field.
func (o *BTWorkspaceMessageBody) SetAppElementSessionId(v string) {
	o.AppElementSessionId = &v
}

// GetCommentId returns the CommentId field value if set, zero value otherwise.
func (o *BTWorkspaceMessageBody) GetCommentId() string {
	if o == nil || o.CommentId == nil {
		var ret string
		return ret
	}
	return *o.CommentId
}

// GetCommentIdOk returns a tuple with the CommentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceMessageBody) GetCommentIdOk() (*string, bool) {
	if o == nil || o.CommentId == nil {
		return nil, false
	}
	return o.CommentId, true
}

// HasCommentId returns a boolean if a field has been set.
func (o *BTWorkspaceMessageBody) HasCommentId() bool {
	if o != nil && o.CommentId != nil {
		return true
	}

	return false
}

// SetCommentId gets a reference to the given string and assigns it to the CommentId field.
func (o *BTWorkspaceMessageBody) SetCommentId(v string) {
	o.CommentId = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BTWorkspaceMessageBody) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceMessageBody) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BTWorkspaceMessageBody) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *BTWorkspaceMessageBody) SetData(v string) {
	o.Data = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTWorkspaceMessageBody) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceMessageBody) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTWorkspaceMessageBody) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTWorkspaceMessageBody) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetDocumentState returns the DocumentState field value if set, zero value otherwise.
func (o *BTWorkspaceMessageBody) GetDocumentState() string {
	if o == nil || o.DocumentState == nil {
		var ret string
		return ret
	}
	return *o.DocumentState
}

// GetDocumentStateOk returns a tuple with the DocumentState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceMessageBody) GetDocumentStateOk() (*string, bool) {
	if o == nil || o.DocumentState == nil {
		return nil, false
	}
	return o.DocumentState, true
}

// HasDocumentState returns a boolean if a field has been set.
func (o *BTWorkspaceMessageBody) HasDocumentState() bool {
	if o != nil && o.DocumentState != nil {
		return true
	}

	return false
}

// SetDocumentState gets a reference to the given string and assigns it to the DocumentState field.
func (o *BTWorkspaceMessageBody) SetDocumentState(v string) {
	o.DocumentState = &v
}

// GetDocumentType returns the DocumentType field value if set, zero value otherwise.
func (o *BTWorkspaceMessageBody) GetDocumentType() int32 {
	if o == nil || o.DocumentType == nil {
		var ret int32
		return ret
	}
	return *o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceMessageBody) GetDocumentTypeOk() (*int32, bool) {
	if o == nil || o.DocumentType == nil {
		return nil, false
	}
	return o.DocumentType, true
}

// HasDocumentType returns a boolean if a field has been set.
func (o *BTWorkspaceMessageBody) HasDocumentType() bool {
	if o != nil && o.DocumentType != nil {
		return true
	}

	return false
}

// SetDocumentType gets a reference to the given int32 and assigns it to the DocumentType field.
func (o *BTWorkspaceMessageBody) SetDocumentType(v int32) {
	o.DocumentType = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTWorkspaceMessageBody) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceMessageBody) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTWorkspaceMessageBody) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTWorkspaceMessageBody) SetElementId(v string) {
	o.ElementId = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *BTWorkspaceMessageBody) GetEvent() string {
	if o == nil || o.Event == nil {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceMessageBody) GetEventOk() (*string, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *BTWorkspaceMessageBody) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *BTWorkspaceMessageBody) SetEvent(v string) {
	o.Event = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *BTWorkspaceMessageBody) GetMessageId() string {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceMessageBody) GetMessageIdOk() (*string, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *BTWorkspaceMessageBody) HasMessageId() bool {
	if o != nil && o.MessageId != nil {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *BTWorkspaceMessageBody) SetMessageId(v string) {
	o.MessageId = &v
}

// GetMetadataObjectType returns the MetadataObjectType field value if set, zero value otherwise.
func (o *BTWorkspaceMessageBody) GetMetadataObjectType() string {
	if o == nil || o.MetadataObjectType == nil {
		var ret string
		return ret
	}
	return *o.MetadataObjectType
}

// GetMetadataObjectTypeOk returns a tuple with the MetadataObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceMessageBody) GetMetadataObjectTypeOk() (*string, bool) {
	if o == nil || o.MetadataObjectType == nil {
		return nil, false
	}
	return o.MetadataObjectType, true
}

// HasMetadataObjectType returns a boolean if a field has been set.
func (o *BTWorkspaceMessageBody) HasMetadataObjectType() bool {
	if o != nil && o.MetadataObjectType != nil {
		return true
	}

	return false
}

// SetMetadataObjectType gets a reference to the given string and assigns it to the MetadataObjectType field.
func (o *BTWorkspaceMessageBody) SetMetadataObjectType(v string) {
	o.MetadataObjectType = &v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *BTWorkspaceMessageBody) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceMessageBody) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *BTWorkspaceMessageBody) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *BTWorkspaceMessageBody) SetPartId(v string) {
	o.PartId = &v
}

// GetPartIdentity returns the PartIdentity field value if set, zero value otherwise.
func (o *BTWorkspaceMessageBody) GetPartIdentity() string {
	if o == nil || o.PartIdentity == nil {
		var ret string
		return ret
	}
	return *o.PartIdentity
}

// GetPartIdentityOk returns a tuple with the PartIdentity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceMessageBody) GetPartIdentityOk() (*string, bool) {
	if o == nil || o.PartIdentity == nil {
		return nil, false
	}
	return o.PartIdentity, true
}

// HasPartIdentity returns a boolean if a field has been set.
func (o *BTWorkspaceMessageBody) HasPartIdentity() bool {
	if o != nil && o.PartIdentity != nil {
		return true
	}

	return false
}

// SetPartIdentity gets a reference to the given string and assigns it to the PartIdentity field.
func (o *BTWorkspaceMessageBody) SetPartIdentity(v string) {
	o.PartIdentity = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *BTWorkspaceMessageBody) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceMessageBody) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *BTWorkspaceMessageBody) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *BTWorkspaceMessageBody) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *BTWorkspaceMessageBody) GetTimestamp() JSONTime {
	if o == nil || o.Timestamp == nil {
		var ret JSONTime
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceMessageBody) GetTimestampOk() (*JSONTime, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *BTWorkspaceMessageBody) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given JSONTime and assigns it to the Timestamp field.
func (o *BTWorkspaceMessageBody) SetTimestamp(v JSONTime) {
	o.Timestamp = &v
}

// GetTranslatationId returns the TranslatationId field value if set, zero value otherwise.
func (o *BTWorkspaceMessageBody) GetTranslatationId() string {
	if o == nil || o.TranslatationId == nil {
		var ret string
		return ret
	}
	return *o.TranslatationId
}

// GetTranslatationIdOk returns a tuple with the TranslatationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceMessageBody) GetTranslatationIdOk() (*string, bool) {
	if o == nil || o.TranslatationId == nil {
		return nil, false
	}
	return o.TranslatationId, true
}

// HasTranslatationId returns a boolean if a field has been set.
func (o *BTWorkspaceMessageBody) HasTranslatationId() bool {
	if o != nil && o.TranslatationId != nil {
		return true
	}

	return false
}

// SetTranslatationId gets a reference to the given string and assigns it to the TranslatationId field.
func (o *BTWorkspaceMessageBody) SetTranslatationId(v string) {
	o.TranslatationId = &v
}

// GetTranslationId returns the TranslationId field value if set, zero value otherwise.
func (o *BTWorkspaceMessageBody) GetTranslationId() string {
	if o == nil || o.TranslationId == nil {
		var ret string
		return ret
	}
	return *o.TranslationId
}

// GetTranslationIdOk returns a tuple with the TranslationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceMessageBody) GetTranslationIdOk() (*string, bool) {
	if o == nil || o.TranslationId == nil {
		return nil, false
	}
	return o.TranslationId, true
}

// HasTranslationId returns a boolean if a field has been set.
func (o *BTWorkspaceMessageBody) HasTranslationId() bool {
	if o != nil && o.TranslationId != nil {
		return true
	}

	return false
}

// SetTranslationId gets a reference to the given string and assigns it to the TranslationId field.
func (o *BTWorkspaceMessageBody) SetTranslationId(v string) {
	o.TranslationId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *BTWorkspaceMessageBody) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceMessageBody) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *BTWorkspaceMessageBody) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *BTWorkspaceMessageBody) SetUserId(v string) {
	o.UserId = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *BTWorkspaceMessageBody) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceMessageBody) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *BTWorkspaceMessageBody) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *BTWorkspaceMessageBody) SetVersionId(v string) {
	o.VersionId = &v
}

// GetWebhookId returns the WebhookId field value if set, zero value otherwise.
func (o *BTWorkspaceMessageBody) GetWebhookId() string {
	if o == nil || o.WebhookId == nil {
		var ret string
		return ret
	}
	return *o.WebhookId
}

// GetWebhookIdOk returns a tuple with the WebhookId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceMessageBody) GetWebhookIdOk() (*string, bool) {
	if o == nil || o.WebhookId == nil {
		return nil, false
	}
	return o.WebhookId, true
}

// HasWebhookId returns a boolean if a field has been set.
func (o *BTWorkspaceMessageBody) HasWebhookId() bool {
	if o != nil && o.WebhookId != nil {
		return true
	}

	return false
}

// SetWebhookId gets a reference to the given string and assigns it to the WebhookId field.
func (o *BTWorkspaceMessageBody) SetWebhookId(v string) {
	o.WebhookId = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *BTWorkspaceMessageBody) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceMessageBody) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *BTWorkspaceMessageBody) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *BTWorkspaceMessageBody) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

// GetDocumentMicroversionId returns the DocumentMicroversionId field value if set, zero value otherwise.
func (o *BTWorkspaceMessageBody) GetDocumentMicroversionId() string {
	if o == nil || o.DocumentMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.DocumentMicroversionId
}

// GetDocumentMicroversionIdOk returns a tuple with the DocumentMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkspaceMessageBody) GetDocumentMicroversionIdOk() (*string, bool) {
	if o == nil || o.DocumentMicroversionId == nil {
		return nil, false
	}
	return o.DocumentMicroversionId, true
}

// HasDocumentMicroversionId returns a boolean if a field has been set.
func (o *BTWorkspaceMessageBody) HasDocumentMicroversionId() bool {
	if o != nil && o.DocumentMicroversionId != nil {
		return true
	}

	return false
}

// SetDocumentMicroversionId gets a reference to the given string and assigns it to the DocumentMicroversionId field.
func (o *BTWorkspaceMessageBody) SetDocumentMicroversionId(v string) {
	o.DocumentMicroversionId = &v
}

func (o BTWorkspaceMessageBody) MarshalJSON() ([]byte, error) {
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
	if o.DocumentType != nil {
		toSerialize["documentType"] = o.DocumentType
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
	if o.DocumentMicroversionId != nil {
		toSerialize["documentMicroversionId"] = o.DocumentMicroversionId
	}
	return json.Marshal(toSerialize)
}

type NullableBTWorkspaceMessageBody struct {
	value *BTWorkspaceMessageBody
	isSet bool
}

func (v NullableBTWorkspaceMessageBody) Get() *BTWorkspaceMessageBody {
	return v.value
}

func (v *NullableBTWorkspaceMessageBody) Set(val *BTWorkspaceMessageBody) {
	v.value = val
	v.isSet = true
}

func (v NullableBTWorkspaceMessageBody) IsSet() bool {
	return v.isSet
}

func (v *NullableBTWorkspaceMessageBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTWorkspaceMessageBody(val *BTWorkspaceMessageBody) *NullableBTWorkspaceMessageBody {
	return &NullableBTWorkspaceMessageBody{value: val, isSet: true}
}

func (v NullableBTWorkspaceMessageBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTWorkspaceMessageBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
