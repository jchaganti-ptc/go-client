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

// BTAssociativeDataInfo struct for BTAssociativeDataInfo
type BTAssociativeDataInfo struct {
	AssociativeDataId    *string           `json:"associativeDataId,omitempty"`
	Data                 []BTNameValuePair `json:"data,omitempty"`
	DocumentId           *string           `json:"documentId,omitempty"`
	DocumentMicroversion *string           `json:"documentMicroversion,omitempty"`
	ElementId            *string           `json:"elementId,omitempty"`
	Error                *string           `json:"error,omitempty"`
	IdTag                *string           `json:"idTag,omitempty"`
	MicroversionId       *string           `json:"microversionId,omitempty"`
	OccurrenceId         *string           `json:"occurrenceId,omitempty"`
	Type                 *string           `json:"type,omitempty"`
	VersionId            *string           `json:"versionId,omitempty"`
}

// NewBTAssociativeDataInfo instantiates a new BTAssociativeDataInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssociativeDataInfo() *BTAssociativeDataInfo {
	this := BTAssociativeDataInfo{}
	return &this
}

// NewBTAssociativeDataInfoWithDefaults instantiates a new BTAssociativeDataInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssociativeDataInfoWithDefaults() *BTAssociativeDataInfo {
	this := BTAssociativeDataInfo{}
	return &this
}

// GetAssociativeDataId returns the AssociativeDataId field value if set, zero value otherwise.
func (o *BTAssociativeDataInfo) GetAssociativeDataId() string {
	if o == nil || o.AssociativeDataId == nil {
		var ret string
		return ret
	}
	return *o.AssociativeDataId
}

// GetAssociativeDataIdOk returns a tuple with the AssociativeDataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssociativeDataInfo) GetAssociativeDataIdOk() (*string, bool) {
	if o == nil || o.AssociativeDataId == nil {
		return nil, false
	}
	return o.AssociativeDataId, true
}

// HasAssociativeDataId returns a boolean if a field has been set.
func (o *BTAssociativeDataInfo) HasAssociativeDataId() bool {
	if o != nil && o.AssociativeDataId != nil {
		return true
	}

	return false
}

// SetAssociativeDataId gets a reference to the given string and assigns it to the AssociativeDataId field.
func (o *BTAssociativeDataInfo) SetAssociativeDataId(v string) {
	o.AssociativeDataId = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BTAssociativeDataInfo) GetData() []BTNameValuePair {
	if o == nil || o.Data == nil {
		var ret []BTNameValuePair
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssociativeDataInfo) GetDataOk() ([]BTNameValuePair, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BTAssociativeDataInfo) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []BTNameValuePair and assigns it to the Data field.
func (o *BTAssociativeDataInfo) SetData(v []BTNameValuePair) {
	o.Data = v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTAssociativeDataInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssociativeDataInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTAssociativeDataInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTAssociativeDataInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetDocumentMicroversion returns the DocumentMicroversion field value if set, zero value otherwise.
func (o *BTAssociativeDataInfo) GetDocumentMicroversion() string {
	if o == nil || o.DocumentMicroversion == nil {
		var ret string
		return ret
	}
	return *o.DocumentMicroversion
}

// GetDocumentMicroversionOk returns a tuple with the DocumentMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssociativeDataInfo) GetDocumentMicroversionOk() (*string, bool) {
	if o == nil || o.DocumentMicroversion == nil {
		return nil, false
	}
	return o.DocumentMicroversion, true
}

// HasDocumentMicroversion returns a boolean if a field has been set.
func (o *BTAssociativeDataInfo) HasDocumentMicroversion() bool {
	if o != nil && o.DocumentMicroversion != nil {
		return true
	}

	return false
}

// SetDocumentMicroversion gets a reference to the given string and assigns it to the DocumentMicroversion field.
func (o *BTAssociativeDataInfo) SetDocumentMicroversion(v string) {
	o.DocumentMicroversion = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTAssociativeDataInfo) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssociativeDataInfo) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTAssociativeDataInfo) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTAssociativeDataInfo) SetElementId(v string) {
	o.ElementId = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *BTAssociativeDataInfo) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssociativeDataInfo) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *BTAssociativeDataInfo) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *BTAssociativeDataInfo) SetError(v string) {
	o.Error = &v
}

// GetIdTag returns the IdTag field value if set, zero value otherwise.
func (o *BTAssociativeDataInfo) GetIdTag() string {
	if o == nil || o.IdTag == nil {
		var ret string
		return ret
	}
	return *o.IdTag
}

// GetIdTagOk returns a tuple with the IdTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssociativeDataInfo) GetIdTagOk() (*string, bool) {
	if o == nil || o.IdTag == nil {
		return nil, false
	}
	return o.IdTag, true
}

// HasIdTag returns a boolean if a field has been set.
func (o *BTAssociativeDataInfo) HasIdTag() bool {
	if o != nil && o.IdTag != nil {
		return true
	}

	return false
}

// SetIdTag gets a reference to the given string and assigns it to the IdTag field.
func (o *BTAssociativeDataInfo) SetIdTag(v string) {
	o.IdTag = &v
}

// GetMicroversionId returns the MicroversionId field value if set, zero value otherwise.
func (o *BTAssociativeDataInfo) GetMicroversionId() string {
	if o == nil || o.MicroversionId == nil {
		var ret string
		return ret
	}
	return *o.MicroversionId
}

// GetMicroversionIdOk returns a tuple with the MicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssociativeDataInfo) GetMicroversionIdOk() (*string, bool) {
	if o == nil || o.MicroversionId == nil {
		return nil, false
	}
	return o.MicroversionId, true
}

// HasMicroversionId returns a boolean if a field has been set.
func (o *BTAssociativeDataInfo) HasMicroversionId() bool {
	if o != nil && o.MicroversionId != nil {
		return true
	}

	return false
}

// SetMicroversionId gets a reference to the given string and assigns it to the MicroversionId field.
func (o *BTAssociativeDataInfo) SetMicroversionId(v string) {
	o.MicroversionId = &v
}

// GetOccurrenceId returns the OccurrenceId field value if set, zero value otherwise.
func (o *BTAssociativeDataInfo) GetOccurrenceId() string {
	if o == nil || o.OccurrenceId == nil {
		var ret string
		return ret
	}
	return *o.OccurrenceId
}

// GetOccurrenceIdOk returns a tuple with the OccurrenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssociativeDataInfo) GetOccurrenceIdOk() (*string, bool) {
	if o == nil || o.OccurrenceId == nil {
		return nil, false
	}
	return o.OccurrenceId, true
}

// HasOccurrenceId returns a boolean if a field has been set.
func (o *BTAssociativeDataInfo) HasOccurrenceId() bool {
	if o != nil && o.OccurrenceId != nil {
		return true
	}

	return false
}

// SetOccurrenceId gets a reference to the given string and assigns it to the OccurrenceId field.
func (o *BTAssociativeDataInfo) SetOccurrenceId(v string) {
	o.OccurrenceId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTAssociativeDataInfo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssociativeDataInfo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTAssociativeDataInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *BTAssociativeDataInfo) SetType(v string) {
	o.Type = &v
}

// GetVersionId returns the VersionId field value if set, zero value otherwise.
func (o *BTAssociativeDataInfo) GetVersionId() string {
	if o == nil || o.VersionId == nil {
		var ret string
		return ret
	}
	return *o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssociativeDataInfo) GetVersionIdOk() (*string, bool) {
	if o == nil || o.VersionId == nil {
		return nil, false
	}
	return o.VersionId, true
}

// HasVersionId returns a boolean if a field has been set.
func (o *BTAssociativeDataInfo) HasVersionId() bool {
	if o != nil && o.VersionId != nil {
		return true
	}

	return false
}

// SetVersionId gets a reference to the given string and assigns it to the VersionId field.
func (o *BTAssociativeDataInfo) SetVersionId(v string) {
	o.VersionId = &v
}

func (o BTAssociativeDataInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AssociativeDataId != nil {
		toSerialize["associativeDataId"] = o.AssociativeDataId
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.DocumentMicroversion != nil {
		toSerialize["documentMicroversion"] = o.DocumentMicroversion
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	if o.IdTag != nil {
		toSerialize["idTag"] = o.IdTag
	}
	if o.MicroversionId != nil {
		toSerialize["microversionId"] = o.MicroversionId
	}
	if o.OccurrenceId != nil {
		toSerialize["occurrenceId"] = o.OccurrenceId
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.VersionId != nil {
		toSerialize["versionId"] = o.VersionId
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssociativeDataInfo struct {
	value *BTAssociativeDataInfo
	isSet bool
}

func (v NullableBTAssociativeDataInfo) Get() *BTAssociativeDataInfo {
	return v.value
}

func (v *NullableBTAssociativeDataInfo) Set(val *BTAssociativeDataInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssociativeDataInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssociativeDataInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssociativeDataInfo(val *BTAssociativeDataInfo) *NullableBTAssociativeDataInfo {
	return &NullableBTAssociativeDataInfo{value: val, isSet: true}
}

func (v NullableBTAssociativeDataInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssociativeDataInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
