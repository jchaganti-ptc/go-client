/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.170.23307-4b97c8644a61
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAppModificationRequestInfo struct for BTAppModificationRequestInfo
type BTAppModificationRequestInfo struct {
	DocumentId    *string `json:"documentId,omitempty"`
	ElementId     *string `json:"elementId,omitempty"`
	FailureReason *string `json:"failureReason,omitempty"`
	// URI to fetch complete information of the resource.
	Href *string `json:"href,omitempty"`
	// Id of the resource.
	Id *string `json:"id,omitempty"`
	// Name of the resource.
	Name                         *string                        `json:"name,omitempty"`
	Output                       *string                        `json:"output,omitempty"`
	ParentDocumentMicroversionId *string                        `json:"parentDocumentMicroversionId,omitempty"`
	ParentElementMicroversionId  *string                        `json:"parentElementMicroversionId,omitempty"`
	RequestState                 *BTAppModificationRequestState `json:"requestState,omitempty"`
	ResultDocumentMicroversionId *string                        `json:"resultDocumentMicroversionId,omitempty"`
	ResultElementMicroversionId  *string                        `json:"resultElementMicroversionId,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef     *string `json:"viewRef,omitempty"`
	WorkspaceId *string `json:"workspaceId,omitempty"`
}

// NewBTAppModificationRequestInfo instantiates a new BTAppModificationRequestInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAppModificationRequestInfo() *BTAppModificationRequestInfo {
	this := BTAppModificationRequestInfo{}
	return &this
}

// NewBTAppModificationRequestInfoWithDefaults instantiates a new BTAppModificationRequestInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAppModificationRequestInfoWithDefaults() *BTAppModificationRequestInfo {
	this := BTAppModificationRequestInfo{}
	return &this
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTAppModificationRequestInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppModificationRequestInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTAppModificationRequestInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTAppModificationRequestInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTAppModificationRequestInfo) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppModificationRequestInfo) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTAppModificationRequestInfo) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTAppModificationRequestInfo) SetElementId(v string) {
	o.ElementId = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *BTAppModificationRequestInfo) GetFailureReason() string {
	if o == nil || o.FailureReason == nil {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppModificationRequestInfo) GetFailureReasonOk() (*string, bool) {
	if o == nil || o.FailureReason == nil {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *BTAppModificationRequestInfo) HasFailureReason() bool {
	if o != nil && o.FailureReason != nil {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *BTAppModificationRequestInfo) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTAppModificationRequestInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppModificationRequestInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTAppModificationRequestInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTAppModificationRequestInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTAppModificationRequestInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppModificationRequestInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTAppModificationRequestInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTAppModificationRequestInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTAppModificationRequestInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppModificationRequestInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTAppModificationRequestInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTAppModificationRequestInfo) SetName(v string) {
	o.Name = &v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *BTAppModificationRequestInfo) GetOutput() string {
	if o == nil || o.Output == nil {
		var ret string
		return ret
	}
	return *o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppModificationRequestInfo) GetOutputOk() (*string, bool) {
	if o == nil || o.Output == nil {
		return nil, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *BTAppModificationRequestInfo) HasOutput() bool {
	if o != nil && o.Output != nil {
		return true
	}

	return false
}

// SetOutput gets a reference to the given string and assigns it to the Output field.
func (o *BTAppModificationRequestInfo) SetOutput(v string) {
	o.Output = &v
}

// GetParentDocumentMicroversionId returns the ParentDocumentMicroversionId field value if set, zero value otherwise.
func (o *BTAppModificationRequestInfo) GetParentDocumentMicroversionId() string {
	if o == nil || o.ParentDocumentMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.ParentDocumentMicroversionId
}

// GetParentDocumentMicroversionIdOk returns a tuple with the ParentDocumentMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppModificationRequestInfo) GetParentDocumentMicroversionIdOk() (*string, bool) {
	if o == nil || o.ParentDocumentMicroversionId == nil {
		return nil, false
	}
	return o.ParentDocumentMicroversionId, true
}

// HasParentDocumentMicroversionId returns a boolean if a field has been set.
func (o *BTAppModificationRequestInfo) HasParentDocumentMicroversionId() bool {
	if o != nil && o.ParentDocumentMicroversionId != nil {
		return true
	}

	return false
}

// SetParentDocumentMicroversionId gets a reference to the given string and assigns it to the ParentDocumentMicroversionId field.
func (o *BTAppModificationRequestInfo) SetParentDocumentMicroversionId(v string) {
	o.ParentDocumentMicroversionId = &v
}

// GetParentElementMicroversionId returns the ParentElementMicroversionId field value if set, zero value otherwise.
func (o *BTAppModificationRequestInfo) GetParentElementMicroversionId() string {
	if o == nil || o.ParentElementMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.ParentElementMicroversionId
}

// GetParentElementMicroversionIdOk returns a tuple with the ParentElementMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppModificationRequestInfo) GetParentElementMicroversionIdOk() (*string, bool) {
	if o == nil || o.ParentElementMicroversionId == nil {
		return nil, false
	}
	return o.ParentElementMicroversionId, true
}

// HasParentElementMicroversionId returns a boolean if a field has been set.
func (o *BTAppModificationRequestInfo) HasParentElementMicroversionId() bool {
	if o != nil && o.ParentElementMicroversionId != nil {
		return true
	}

	return false
}

// SetParentElementMicroversionId gets a reference to the given string and assigns it to the ParentElementMicroversionId field.
func (o *BTAppModificationRequestInfo) SetParentElementMicroversionId(v string) {
	o.ParentElementMicroversionId = &v
}

// GetRequestState returns the RequestState field value if set, zero value otherwise.
func (o *BTAppModificationRequestInfo) GetRequestState() BTAppModificationRequestState {
	if o == nil || o.RequestState == nil {
		var ret BTAppModificationRequestState
		return ret
	}
	return *o.RequestState
}

// GetRequestStateOk returns a tuple with the RequestState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppModificationRequestInfo) GetRequestStateOk() (*BTAppModificationRequestState, bool) {
	if o == nil || o.RequestState == nil {
		return nil, false
	}
	return o.RequestState, true
}

// HasRequestState returns a boolean if a field has been set.
func (o *BTAppModificationRequestInfo) HasRequestState() bool {
	if o != nil && o.RequestState != nil {
		return true
	}

	return false
}

// SetRequestState gets a reference to the given BTAppModificationRequestState and assigns it to the RequestState field.
func (o *BTAppModificationRequestInfo) SetRequestState(v BTAppModificationRequestState) {
	o.RequestState = &v
}

// GetResultDocumentMicroversionId returns the ResultDocumentMicroversionId field value if set, zero value otherwise.
func (o *BTAppModificationRequestInfo) GetResultDocumentMicroversionId() string {
	if o == nil || o.ResultDocumentMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.ResultDocumentMicroversionId
}

// GetResultDocumentMicroversionIdOk returns a tuple with the ResultDocumentMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppModificationRequestInfo) GetResultDocumentMicroversionIdOk() (*string, bool) {
	if o == nil || o.ResultDocumentMicroversionId == nil {
		return nil, false
	}
	return o.ResultDocumentMicroversionId, true
}

// HasResultDocumentMicroversionId returns a boolean if a field has been set.
func (o *BTAppModificationRequestInfo) HasResultDocumentMicroversionId() bool {
	if o != nil && o.ResultDocumentMicroversionId != nil {
		return true
	}

	return false
}

// SetResultDocumentMicroversionId gets a reference to the given string and assigns it to the ResultDocumentMicroversionId field.
func (o *BTAppModificationRequestInfo) SetResultDocumentMicroversionId(v string) {
	o.ResultDocumentMicroversionId = &v
}

// GetResultElementMicroversionId returns the ResultElementMicroversionId field value if set, zero value otherwise.
func (o *BTAppModificationRequestInfo) GetResultElementMicroversionId() string {
	if o == nil || o.ResultElementMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.ResultElementMicroversionId
}

// GetResultElementMicroversionIdOk returns a tuple with the ResultElementMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppModificationRequestInfo) GetResultElementMicroversionIdOk() (*string, bool) {
	if o == nil || o.ResultElementMicroversionId == nil {
		return nil, false
	}
	return o.ResultElementMicroversionId, true
}

// HasResultElementMicroversionId returns a boolean if a field has been set.
func (o *BTAppModificationRequestInfo) HasResultElementMicroversionId() bool {
	if o != nil && o.ResultElementMicroversionId != nil {
		return true
	}

	return false
}

// SetResultElementMicroversionId gets a reference to the given string and assigns it to the ResultElementMicroversionId field.
func (o *BTAppModificationRequestInfo) SetResultElementMicroversionId(v string) {
	o.ResultElementMicroversionId = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTAppModificationRequestInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppModificationRequestInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTAppModificationRequestInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTAppModificationRequestInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *BTAppModificationRequestInfo) GetWorkspaceId() string {
	if o == nil || o.WorkspaceId == nil {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAppModificationRequestInfo) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || o.WorkspaceId == nil {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *BTAppModificationRequestInfo) HasWorkspaceId() bool {
	if o != nil && o.WorkspaceId != nil {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *BTAppModificationRequestInfo) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

func (o BTAppModificationRequestInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.FailureReason != nil {
		toSerialize["failureReason"] = o.FailureReason
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Output != nil {
		toSerialize["output"] = o.Output
	}
	if o.ParentDocumentMicroversionId != nil {
		toSerialize["parentDocumentMicroversionId"] = o.ParentDocumentMicroversionId
	}
	if o.ParentElementMicroversionId != nil {
		toSerialize["parentElementMicroversionId"] = o.ParentElementMicroversionId
	}
	if o.RequestState != nil {
		toSerialize["requestState"] = o.RequestState
	}
	if o.ResultDocumentMicroversionId != nil {
		toSerialize["resultDocumentMicroversionId"] = o.ResultDocumentMicroversionId
	}
	if o.ResultElementMicroversionId != nil {
		toSerialize["resultElementMicroversionId"] = o.ResultElementMicroversionId
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	if o.WorkspaceId != nil {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	return json.Marshal(toSerialize)
}

type NullableBTAppModificationRequestInfo struct {
	value *BTAppModificationRequestInfo
	isSet bool
}

func (v NullableBTAppModificationRequestInfo) Get() *BTAppModificationRequestInfo {
	return v.value
}

func (v *NullableBTAppModificationRequestInfo) Set(val *BTAppModificationRequestInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAppModificationRequestInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAppModificationRequestInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAppModificationRequestInfo(val *BTAppModificationRequestInfo) *NullableBTAppModificationRequestInfo {
	return &NullableBTAppModificationRequestInfo{value: val, isSet: true}
}

func (v NullableBTAppModificationRequestInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAppModificationRequestInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
