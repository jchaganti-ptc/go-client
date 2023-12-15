/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.173.27678-64d64396ca66
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTWorkflowAuditLogEntryInfo struct for BTWorkflowAuditLogEntryInfo
type BTWorkflowAuditLogEntryInfo struct {
	ApprovalOverride      *bool                   `json:"approvalOverride,omitempty"`
	ApproverIds           []string                `json:"approverIds,omitempty"`
	CommentId             *string                 `json:"commentId,omitempty"`
	Date                  *JSONTime               `json:"date,omitempty"`
	EntryType             *int32                  `json:"entryType,omitempty"`
	ErrorMessage          *string                 `json:"errorMessage,omitempty"`
	FeatureScriptConsole  *string                 `json:"featureScriptConsole,omitempty"`
	FeatureScriptNotices  []string                `json:"featureScriptNotices,omitempty"`
	FeatureScriptResponse *map[string]interface{} `json:"featureScriptResponse,omitempty"`
	Id                    *string                 `json:"id,omitempty"`
	ObjectId              *string                 `json:"objectId,omitempty"`
	PropertyUpdates       []BTPropertyUpdateInfo  `json:"propertyUpdates,omitempty"`
	SupportCode           *string                 `json:"supportCode,omitempty"`
	UserId                *string                 `json:"userId,omitempty"`
	WorkflowAction        *string                 `json:"workflowAction,omitempty"`
	WorkflowState         *string                 `json:"workflowState,omitempty"`
	WorkflowTransition    *string                 `json:"workflowTransition,omitempty"`
}

// NewBTWorkflowAuditLogEntryInfo instantiates a new BTWorkflowAuditLogEntryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTWorkflowAuditLogEntryInfo() *BTWorkflowAuditLogEntryInfo {
	this := BTWorkflowAuditLogEntryInfo{}
	return &this
}

// NewBTWorkflowAuditLogEntryInfoWithDefaults instantiates a new BTWorkflowAuditLogEntryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTWorkflowAuditLogEntryInfoWithDefaults() *BTWorkflowAuditLogEntryInfo {
	this := BTWorkflowAuditLogEntryInfo{}
	return &this
}

// GetApprovalOverride returns the ApprovalOverride field value if set, zero value otherwise.
func (o *BTWorkflowAuditLogEntryInfo) GetApprovalOverride() bool {
	if o == nil || o.ApprovalOverride == nil {
		var ret bool
		return ret
	}
	return *o.ApprovalOverride
}

// GetApprovalOverrideOk returns a tuple with the ApprovalOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowAuditLogEntryInfo) GetApprovalOverrideOk() (*bool, bool) {
	if o == nil || o.ApprovalOverride == nil {
		return nil, false
	}
	return o.ApprovalOverride, true
}

// HasApprovalOverride returns a boolean if a field has been set.
func (o *BTWorkflowAuditLogEntryInfo) HasApprovalOverride() bool {
	if o != nil && o.ApprovalOverride != nil {
		return true
	}

	return false
}

// SetApprovalOverride gets a reference to the given bool and assigns it to the ApprovalOverride field.
func (o *BTWorkflowAuditLogEntryInfo) SetApprovalOverride(v bool) {
	o.ApprovalOverride = &v
}

// GetApproverIds returns the ApproverIds field value if set, zero value otherwise.
func (o *BTWorkflowAuditLogEntryInfo) GetApproverIds() []string {
	if o == nil || o.ApproverIds == nil {
		var ret []string
		return ret
	}
	return o.ApproverIds
}

// GetApproverIdsOk returns a tuple with the ApproverIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowAuditLogEntryInfo) GetApproverIdsOk() ([]string, bool) {
	if o == nil || o.ApproverIds == nil {
		return nil, false
	}
	return o.ApproverIds, true
}

// HasApproverIds returns a boolean if a field has been set.
func (o *BTWorkflowAuditLogEntryInfo) HasApproverIds() bool {
	if o != nil && o.ApproverIds != nil {
		return true
	}

	return false
}

// SetApproverIds gets a reference to the given []string and assigns it to the ApproverIds field.
func (o *BTWorkflowAuditLogEntryInfo) SetApproverIds(v []string) {
	o.ApproverIds = v
}

// GetCommentId returns the CommentId field value if set, zero value otherwise.
func (o *BTWorkflowAuditLogEntryInfo) GetCommentId() string {
	if o == nil || o.CommentId == nil {
		var ret string
		return ret
	}
	return *o.CommentId
}

// GetCommentIdOk returns a tuple with the CommentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowAuditLogEntryInfo) GetCommentIdOk() (*string, bool) {
	if o == nil || o.CommentId == nil {
		return nil, false
	}
	return o.CommentId, true
}

// HasCommentId returns a boolean if a field has been set.
func (o *BTWorkflowAuditLogEntryInfo) HasCommentId() bool {
	if o != nil && o.CommentId != nil {
		return true
	}

	return false
}

// SetCommentId gets a reference to the given string and assigns it to the CommentId field.
func (o *BTWorkflowAuditLogEntryInfo) SetCommentId(v string) {
	o.CommentId = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *BTWorkflowAuditLogEntryInfo) GetDate() JSONTime {
	if o == nil || o.Date == nil {
		var ret JSONTime
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowAuditLogEntryInfo) GetDateOk() (*JSONTime, bool) {
	if o == nil || o.Date == nil {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *BTWorkflowAuditLogEntryInfo) HasDate() bool {
	if o != nil && o.Date != nil {
		return true
	}

	return false
}

// SetDate gets a reference to the given JSONTime and assigns it to the Date field.
func (o *BTWorkflowAuditLogEntryInfo) SetDate(v JSONTime) {
	o.Date = &v
}

// GetEntryType returns the EntryType field value if set, zero value otherwise.
func (o *BTWorkflowAuditLogEntryInfo) GetEntryType() int32 {
	if o == nil || o.EntryType == nil {
		var ret int32
		return ret
	}
	return *o.EntryType
}

// GetEntryTypeOk returns a tuple with the EntryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowAuditLogEntryInfo) GetEntryTypeOk() (*int32, bool) {
	if o == nil || o.EntryType == nil {
		return nil, false
	}
	return o.EntryType, true
}

// HasEntryType returns a boolean if a field has been set.
func (o *BTWorkflowAuditLogEntryInfo) HasEntryType() bool {
	if o != nil && o.EntryType != nil {
		return true
	}

	return false
}

// SetEntryType gets a reference to the given int32 and assigns it to the EntryType field.
func (o *BTWorkflowAuditLogEntryInfo) SetEntryType(v int32) {
	o.EntryType = &v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *BTWorkflowAuditLogEntryInfo) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowAuditLogEntryInfo) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *BTWorkflowAuditLogEntryInfo) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *BTWorkflowAuditLogEntryInfo) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

// GetFeatureScriptConsole returns the FeatureScriptConsole field value if set, zero value otherwise.
func (o *BTWorkflowAuditLogEntryInfo) GetFeatureScriptConsole() string {
	if o == nil || o.FeatureScriptConsole == nil {
		var ret string
		return ret
	}
	return *o.FeatureScriptConsole
}

// GetFeatureScriptConsoleOk returns a tuple with the FeatureScriptConsole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowAuditLogEntryInfo) GetFeatureScriptConsoleOk() (*string, bool) {
	if o == nil || o.FeatureScriptConsole == nil {
		return nil, false
	}
	return o.FeatureScriptConsole, true
}

// HasFeatureScriptConsole returns a boolean if a field has been set.
func (o *BTWorkflowAuditLogEntryInfo) HasFeatureScriptConsole() bool {
	if o != nil && o.FeatureScriptConsole != nil {
		return true
	}

	return false
}

// SetFeatureScriptConsole gets a reference to the given string and assigns it to the FeatureScriptConsole field.
func (o *BTWorkflowAuditLogEntryInfo) SetFeatureScriptConsole(v string) {
	o.FeatureScriptConsole = &v
}

// GetFeatureScriptNotices returns the FeatureScriptNotices field value if set, zero value otherwise.
func (o *BTWorkflowAuditLogEntryInfo) GetFeatureScriptNotices() []string {
	if o == nil || o.FeatureScriptNotices == nil {
		var ret []string
		return ret
	}
	return o.FeatureScriptNotices
}

// GetFeatureScriptNoticesOk returns a tuple with the FeatureScriptNotices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowAuditLogEntryInfo) GetFeatureScriptNoticesOk() ([]string, bool) {
	if o == nil || o.FeatureScriptNotices == nil {
		return nil, false
	}
	return o.FeatureScriptNotices, true
}

// HasFeatureScriptNotices returns a boolean if a field has been set.
func (o *BTWorkflowAuditLogEntryInfo) HasFeatureScriptNotices() bool {
	if o != nil && o.FeatureScriptNotices != nil {
		return true
	}

	return false
}

// SetFeatureScriptNotices gets a reference to the given []string and assigns it to the FeatureScriptNotices field.
func (o *BTWorkflowAuditLogEntryInfo) SetFeatureScriptNotices(v []string) {
	o.FeatureScriptNotices = v
}

// GetFeatureScriptResponse returns the FeatureScriptResponse field value if set, zero value otherwise.
func (o *BTWorkflowAuditLogEntryInfo) GetFeatureScriptResponse() map[string]interface{} {
	if o == nil || o.FeatureScriptResponse == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.FeatureScriptResponse
}

// GetFeatureScriptResponseOk returns a tuple with the FeatureScriptResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowAuditLogEntryInfo) GetFeatureScriptResponseOk() (*map[string]interface{}, bool) {
	if o == nil || o.FeatureScriptResponse == nil {
		return nil, false
	}
	return o.FeatureScriptResponse, true
}

// HasFeatureScriptResponse returns a boolean if a field has been set.
func (o *BTWorkflowAuditLogEntryInfo) HasFeatureScriptResponse() bool {
	if o != nil && o.FeatureScriptResponse != nil {
		return true
	}

	return false
}

// SetFeatureScriptResponse gets a reference to the given map[string]interface{} and assigns it to the FeatureScriptResponse field.
func (o *BTWorkflowAuditLogEntryInfo) SetFeatureScriptResponse(v map[string]interface{}) {
	o.FeatureScriptResponse = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTWorkflowAuditLogEntryInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowAuditLogEntryInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTWorkflowAuditLogEntryInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTWorkflowAuditLogEntryInfo) SetId(v string) {
	o.Id = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *BTWorkflowAuditLogEntryInfo) GetObjectId() string {
	if o == nil || o.ObjectId == nil {
		var ret string
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowAuditLogEntryInfo) GetObjectIdOk() (*string, bool) {
	if o == nil || o.ObjectId == nil {
		return nil, false
	}
	return o.ObjectId, true
}

// HasObjectId returns a boolean if a field has been set.
func (o *BTWorkflowAuditLogEntryInfo) HasObjectId() bool {
	if o != nil && o.ObjectId != nil {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given string and assigns it to the ObjectId field.
func (o *BTWorkflowAuditLogEntryInfo) SetObjectId(v string) {
	o.ObjectId = &v
}

// GetPropertyUpdates returns the PropertyUpdates field value if set, zero value otherwise.
func (o *BTWorkflowAuditLogEntryInfo) GetPropertyUpdates() []BTPropertyUpdateInfo {
	if o == nil || o.PropertyUpdates == nil {
		var ret []BTPropertyUpdateInfo
		return ret
	}
	return o.PropertyUpdates
}

// GetPropertyUpdatesOk returns a tuple with the PropertyUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowAuditLogEntryInfo) GetPropertyUpdatesOk() ([]BTPropertyUpdateInfo, bool) {
	if o == nil || o.PropertyUpdates == nil {
		return nil, false
	}
	return o.PropertyUpdates, true
}

// HasPropertyUpdates returns a boolean if a field has been set.
func (o *BTWorkflowAuditLogEntryInfo) HasPropertyUpdates() bool {
	if o != nil && o.PropertyUpdates != nil {
		return true
	}

	return false
}

// SetPropertyUpdates gets a reference to the given []BTPropertyUpdateInfo and assigns it to the PropertyUpdates field.
func (o *BTWorkflowAuditLogEntryInfo) SetPropertyUpdates(v []BTPropertyUpdateInfo) {
	o.PropertyUpdates = v
}

// GetSupportCode returns the SupportCode field value if set, zero value otherwise.
func (o *BTWorkflowAuditLogEntryInfo) GetSupportCode() string {
	if o == nil || o.SupportCode == nil {
		var ret string
		return ret
	}
	return *o.SupportCode
}

// GetSupportCodeOk returns a tuple with the SupportCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowAuditLogEntryInfo) GetSupportCodeOk() (*string, bool) {
	if o == nil || o.SupportCode == nil {
		return nil, false
	}
	return o.SupportCode, true
}

// HasSupportCode returns a boolean if a field has been set.
func (o *BTWorkflowAuditLogEntryInfo) HasSupportCode() bool {
	if o != nil && o.SupportCode != nil {
		return true
	}

	return false
}

// SetSupportCode gets a reference to the given string and assigns it to the SupportCode field.
func (o *BTWorkflowAuditLogEntryInfo) SetSupportCode(v string) {
	o.SupportCode = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *BTWorkflowAuditLogEntryInfo) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowAuditLogEntryInfo) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *BTWorkflowAuditLogEntryInfo) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *BTWorkflowAuditLogEntryInfo) SetUserId(v string) {
	o.UserId = &v
}

// GetWorkflowAction returns the WorkflowAction field value if set, zero value otherwise.
func (o *BTWorkflowAuditLogEntryInfo) GetWorkflowAction() string {
	if o == nil || o.WorkflowAction == nil {
		var ret string
		return ret
	}
	return *o.WorkflowAction
}

// GetWorkflowActionOk returns a tuple with the WorkflowAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowAuditLogEntryInfo) GetWorkflowActionOk() (*string, bool) {
	if o == nil || o.WorkflowAction == nil {
		return nil, false
	}
	return o.WorkflowAction, true
}

// HasWorkflowAction returns a boolean if a field has been set.
func (o *BTWorkflowAuditLogEntryInfo) HasWorkflowAction() bool {
	if o != nil && o.WorkflowAction != nil {
		return true
	}

	return false
}

// SetWorkflowAction gets a reference to the given string and assigns it to the WorkflowAction field.
func (o *BTWorkflowAuditLogEntryInfo) SetWorkflowAction(v string) {
	o.WorkflowAction = &v
}

// GetWorkflowState returns the WorkflowState field value if set, zero value otherwise.
func (o *BTWorkflowAuditLogEntryInfo) GetWorkflowState() string {
	if o == nil || o.WorkflowState == nil {
		var ret string
		return ret
	}
	return *o.WorkflowState
}

// GetWorkflowStateOk returns a tuple with the WorkflowState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowAuditLogEntryInfo) GetWorkflowStateOk() (*string, bool) {
	if o == nil || o.WorkflowState == nil {
		return nil, false
	}
	return o.WorkflowState, true
}

// HasWorkflowState returns a boolean if a field has been set.
func (o *BTWorkflowAuditLogEntryInfo) HasWorkflowState() bool {
	if o != nil && o.WorkflowState != nil {
		return true
	}

	return false
}

// SetWorkflowState gets a reference to the given string and assigns it to the WorkflowState field.
func (o *BTWorkflowAuditLogEntryInfo) SetWorkflowState(v string) {
	o.WorkflowState = &v
}

// GetWorkflowTransition returns the WorkflowTransition field value if set, zero value otherwise.
func (o *BTWorkflowAuditLogEntryInfo) GetWorkflowTransition() string {
	if o == nil || o.WorkflowTransition == nil {
		var ret string
		return ret
	}
	return *o.WorkflowTransition
}

// GetWorkflowTransitionOk returns a tuple with the WorkflowTransition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTWorkflowAuditLogEntryInfo) GetWorkflowTransitionOk() (*string, bool) {
	if o == nil || o.WorkflowTransition == nil {
		return nil, false
	}
	return o.WorkflowTransition, true
}

// HasWorkflowTransition returns a boolean if a field has been set.
func (o *BTWorkflowAuditLogEntryInfo) HasWorkflowTransition() bool {
	if o != nil && o.WorkflowTransition != nil {
		return true
	}

	return false
}

// SetWorkflowTransition gets a reference to the given string and assigns it to the WorkflowTransition field.
func (o *BTWorkflowAuditLogEntryInfo) SetWorkflowTransition(v string) {
	o.WorkflowTransition = &v
}

func (o BTWorkflowAuditLogEntryInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApprovalOverride != nil {
		toSerialize["approvalOverride"] = o.ApprovalOverride
	}
	if o.ApproverIds != nil {
		toSerialize["approverIds"] = o.ApproverIds
	}
	if o.CommentId != nil {
		toSerialize["commentId"] = o.CommentId
	}
	if o.Date != nil {
		toSerialize["date"] = o.Date
	}
	if o.EntryType != nil {
		toSerialize["entryType"] = o.EntryType
	}
	if o.ErrorMessage != nil {
		toSerialize["errorMessage"] = o.ErrorMessage
	}
	if o.FeatureScriptConsole != nil {
		toSerialize["featureScriptConsole"] = o.FeatureScriptConsole
	}
	if o.FeatureScriptNotices != nil {
		toSerialize["featureScriptNotices"] = o.FeatureScriptNotices
	}
	if o.FeatureScriptResponse != nil {
		toSerialize["featureScriptResponse"] = o.FeatureScriptResponse
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ObjectId != nil {
		toSerialize["objectId"] = o.ObjectId
	}
	if o.PropertyUpdates != nil {
		toSerialize["propertyUpdates"] = o.PropertyUpdates
	}
	if o.SupportCode != nil {
		toSerialize["supportCode"] = o.SupportCode
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	if o.WorkflowAction != nil {
		toSerialize["workflowAction"] = o.WorkflowAction
	}
	if o.WorkflowState != nil {
		toSerialize["workflowState"] = o.WorkflowState
	}
	if o.WorkflowTransition != nil {
		toSerialize["workflowTransition"] = o.WorkflowTransition
	}
	return json.Marshal(toSerialize)
}

type NullableBTWorkflowAuditLogEntryInfo struct {
	value *BTWorkflowAuditLogEntryInfo
	isSet bool
}

func (v NullableBTWorkflowAuditLogEntryInfo) Get() *BTWorkflowAuditLogEntryInfo {
	return v.value
}

func (v *NullableBTWorkflowAuditLogEntryInfo) Set(val *BTWorkflowAuditLogEntryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTWorkflowAuditLogEntryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTWorkflowAuditLogEntryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTWorkflowAuditLogEntryInfo(val *BTWorkflowAuditLogEntryInfo) *NullableBTWorkflowAuditLogEntryInfo {
	return &NullableBTWorkflowAuditLogEntryInfo{value: val, isSet: true}
}

func (v NullableBTWorkflowAuditLogEntryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTWorkflowAuditLogEntryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
