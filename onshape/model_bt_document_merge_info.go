/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.157.9191-43c781405890
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTDocumentMergeInfo struct for BTDocumentMergeInfo
type BTDocumentMergeInfo struct {
	LibraryVersionMismatch       *bool                   `json:"libraryVersionMismatch,omitempty"`
	OverwrittenElements          []BTDocumentElementInfo `json:"overwrittenElements,omitempty"`
	ParentDocumentMicroversionId *string                 `json:"parentDocumentMicroversionId,omitempty"`
}

// NewBTDocumentMergeInfo instantiates a new BTDocumentMergeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDocumentMergeInfo() *BTDocumentMergeInfo {
	this := BTDocumentMergeInfo{}
	return &this
}

// NewBTDocumentMergeInfoWithDefaults instantiates a new BTDocumentMergeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDocumentMergeInfoWithDefaults() *BTDocumentMergeInfo {
	this := BTDocumentMergeInfo{}
	return &this
}

// GetLibraryVersionMismatch returns the LibraryVersionMismatch field value if set, zero value otherwise.
func (o *BTDocumentMergeInfo) GetLibraryVersionMismatch() bool {
	if o == nil || o.LibraryVersionMismatch == nil {
		var ret bool
		return ret
	}
	return *o.LibraryVersionMismatch
}

// GetLibraryVersionMismatchOk returns a tuple with the LibraryVersionMismatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentMergeInfo) GetLibraryVersionMismatchOk() (*bool, bool) {
	if o == nil || o.LibraryVersionMismatch == nil {
		return nil, false
	}
	return o.LibraryVersionMismatch, true
}

// HasLibraryVersionMismatch returns a boolean if a field has been set.
func (o *BTDocumentMergeInfo) HasLibraryVersionMismatch() bool {
	if o != nil && o.LibraryVersionMismatch != nil {
		return true
	}

	return false
}

// SetLibraryVersionMismatch gets a reference to the given bool and assigns it to the LibraryVersionMismatch field.
func (o *BTDocumentMergeInfo) SetLibraryVersionMismatch(v bool) {
	o.LibraryVersionMismatch = &v
}

// GetOverwrittenElements returns the OverwrittenElements field value if set, zero value otherwise.
func (o *BTDocumentMergeInfo) GetOverwrittenElements() []BTDocumentElementInfo {
	if o == nil || o.OverwrittenElements == nil {
		var ret []BTDocumentElementInfo
		return ret
	}
	return o.OverwrittenElements
}

// GetOverwrittenElementsOk returns a tuple with the OverwrittenElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentMergeInfo) GetOverwrittenElementsOk() ([]BTDocumentElementInfo, bool) {
	if o == nil || o.OverwrittenElements == nil {
		return nil, false
	}
	return o.OverwrittenElements, true
}

// HasOverwrittenElements returns a boolean if a field has been set.
func (o *BTDocumentMergeInfo) HasOverwrittenElements() bool {
	if o != nil && o.OverwrittenElements != nil {
		return true
	}

	return false
}

// SetOverwrittenElements gets a reference to the given []BTDocumentElementInfo and assigns it to the OverwrittenElements field.
func (o *BTDocumentMergeInfo) SetOverwrittenElements(v []BTDocumentElementInfo) {
	o.OverwrittenElements = v
}

// GetParentDocumentMicroversionId returns the ParentDocumentMicroversionId field value if set, zero value otherwise.
func (o *BTDocumentMergeInfo) GetParentDocumentMicroversionId() string {
	if o == nil || o.ParentDocumentMicroversionId == nil {
		var ret string
		return ret
	}
	return *o.ParentDocumentMicroversionId
}

// GetParentDocumentMicroversionIdOk returns a tuple with the ParentDocumentMicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentMergeInfo) GetParentDocumentMicroversionIdOk() (*string, bool) {
	if o == nil || o.ParentDocumentMicroversionId == nil {
		return nil, false
	}
	return o.ParentDocumentMicroversionId, true
}

// HasParentDocumentMicroversionId returns a boolean if a field has been set.
func (o *BTDocumentMergeInfo) HasParentDocumentMicroversionId() bool {
	if o != nil && o.ParentDocumentMicroversionId != nil {
		return true
	}

	return false
}

// SetParentDocumentMicroversionId gets a reference to the given string and assigns it to the ParentDocumentMicroversionId field.
func (o *BTDocumentMergeInfo) SetParentDocumentMicroversionId(v string) {
	o.ParentDocumentMicroversionId = &v
}

func (o BTDocumentMergeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LibraryVersionMismatch != nil {
		toSerialize["libraryVersionMismatch"] = o.LibraryVersionMismatch
	}
	if o.OverwrittenElements != nil {
		toSerialize["overwrittenElements"] = o.OverwrittenElements
	}
	if o.ParentDocumentMicroversionId != nil {
		toSerialize["parentDocumentMicroversionId"] = o.ParentDocumentMicroversionId
	}
	return json.Marshal(toSerialize)
}

type NullableBTDocumentMergeInfo struct {
	value *BTDocumentMergeInfo
	isSet bool
}

func (v NullableBTDocumentMergeInfo) Get() *BTDocumentMergeInfo {
	return v.value
}

func (v *NullableBTDocumentMergeInfo) Set(val *BTDocumentMergeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDocumentMergeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDocumentMergeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDocumentMergeInfo(val *BTDocumentMergeInfo) *NullableBTDocumentMergeInfo {
	return &NullableBTDocumentMergeInfo{value: val, isSet: true}
}

func (v NullableBTDocumentMergeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDocumentMergeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
