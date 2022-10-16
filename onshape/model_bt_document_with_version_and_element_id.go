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

// BTDocumentWithVersionAndElementId struct for BTDocumentWithVersionAndElementId
type BTDocumentWithVersionAndElementId struct {
	DocumentId                   *string   `json:"documentId,omitempty"`
	DocumentVersionId            *string   `json:"documentVersionId,omitempty"`
	ElementId                    *string   `json:"elementId,omitempty"`
	ElementLibraryId             *ObjectId `json:"elementLibraryId,omitempty"`
	ElementLibraryVersion        *ObjectId `json:"elementLibraryVersion,omitempty"`
	PartNumber                   *string   `json:"partNumber,omitempty"`
	Revision                     *string   `json:"revision,omitempty"`
	UniqueVersionId              *string   `json:"uniqueVersionId,omitempty"`
	ValidElementLibraryReference *bool     `json:"validElementLibraryReference,omitempty"`
	ValidRevisionReference       *bool     `json:"validRevisionReference,omitempty"`
}

// NewBTDocumentWithVersionAndElementId instantiates a new BTDocumentWithVersionAndElementId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDocumentWithVersionAndElementId() *BTDocumentWithVersionAndElementId {
	this := BTDocumentWithVersionAndElementId{}
	return &this
}

// NewBTDocumentWithVersionAndElementIdWithDefaults instantiates a new BTDocumentWithVersionAndElementId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDocumentWithVersionAndElementIdWithDefaults() *BTDocumentWithVersionAndElementId {
	this := BTDocumentWithVersionAndElementId{}
	return &this
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTDocumentWithVersionAndElementId) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentWithVersionAndElementId) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTDocumentWithVersionAndElementId) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTDocumentWithVersionAndElementId) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetDocumentVersionId returns the DocumentVersionId field value if set, zero value otherwise.
func (o *BTDocumentWithVersionAndElementId) GetDocumentVersionId() string {
	if o == nil || o.DocumentVersionId == nil {
		var ret string
		return ret
	}
	return *o.DocumentVersionId
}

// GetDocumentVersionIdOk returns a tuple with the DocumentVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentWithVersionAndElementId) GetDocumentVersionIdOk() (*string, bool) {
	if o == nil || o.DocumentVersionId == nil {
		return nil, false
	}
	return o.DocumentVersionId, true
}

// HasDocumentVersionId returns a boolean if a field has been set.
func (o *BTDocumentWithVersionAndElementId) HasDocumentVersionId() bool {
	if o != nil && o.DocumentVersionId != nil {
		return true
	}

	return false
}

// SetDocumentVersionId gets a reference to the given string and assigns it to the DocumentVersionId field.
func (o *BTDocumentWithVersionAndElementId) SetDocumentVersionId(v string) {
	o.DocumentVersionId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTDocumentWithVersionAndElementId) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentWithVersionAndElementId) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTDocumentWithVersionAndElementId) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTDocumentWithVersionAndElementId) SetElementId(v string) {
	o.ElementId = &v
}

// GetElementLibraryId returns the ElementLibraryId field value if set, zero value otherwise.
func (o *BTDocumentWithVersionAndElementId) GetElementLibraryId() ObjectId {
	if o == nil || o.ElementLibraryId == nil {
		var ret ObjectId
		return ret
	}
	return *o.ElementLibraryId
}

// GetElementLibraryIdOk returns a tuple with the ElementLibraryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentWithVersionAndElementId) GetElementLibraryIdOk() (*ObjectId, bool) {
	if o == nil || o.ElementLibraryId == nil {
		return nil, false
	}
	return o.ElementLibraryId, true
}

// HasElementLibraryId returns a boolean if a field has been set.
func (o *BTDocumentWithVersionAndElementId) HasElementLibraryId() bool {
	if o != nil && o.ElementLibraryId != nil {
		return true
	}

	return false
}

// SetElementLibraryId gets a reference to the given ObjectId and assigns it to the ElementLibraryId field.
func (o *BTDocumentWithVersionAndElementId) SetElementLibraryId(v ObjectId) {
	o.ElementLibraryId = &v
}

// GetElementLibraryVersion returns the ElementLibraryVersion field value if set, zero value otherwise.
func (o *BTDocumentWithVersionAndElementId) GetElementLibraryVersion() ObjectId {
	if o == nil || o.ElementLibraryVersion == nil {
		var ret ObjectId
		return ret
	}
	return *o.ElementLibraryVersion
}

// GetElementLibraryVersionOk returns a tuple with the ElementLibraryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentWithVersionAndElementId) GetElementLibraryVersionOk() (*ObjectId, bool) {
	if o == nil || o.ElementLibraryVersion == nil {
		return nil, false
	}
	return o.ElementLibraryVersion, true
}

// HasElementLibraryVersion returns a boolean if a field has been set.
func (o *BTDocumentWithVersionAndElementId) HasElementLibraryVersion() bool {
	if o != nil && o.ElementLibraryVersion != nil {
		return true
	}

	return false
}

// SetElementLibraryVersion gets a reference to the given ObjectId and assigns it to the ElementLibraryVersion field.
func (o *BTDocumentWithVersionAndElementId) SetElementLibraryVersion(v ObjectId) {
	o.ElementLibraryVersion = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *BTDocumentWithVersionAndElementId) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentWithVersionAndElementId) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *BTDocumentWithVersionAndElementId) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *BTDocumentWithVersionAndElementId) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *BTDocumentWithVersionAndElementId) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentWithVersionAndElementId) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *BTDocumentWithVersionAndElementId) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *BTDocumentWithVersionAndElementId) SetRevision(v string) {
	o.Revision = &v
}

// GetUniqueVersionId returns the UniqueVersionId field value if set, zero value otherwise.
func (o *BTDocumentWithVersionAndElementId) GetUniqueVersionId() string {
	if o == nil || o.UniqueVersionId == nil {
		var ret string
		return ret
	}
	return *o.UniqueVersionId
}

// GetUniqueVersionIdOk returns a tuple with the UniqueVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentWithVersionAndElementId) GetUniqueVersionIdOk() (*string, bool) {
	if o == nil || o.UniqueVersionId == nil {
		return nil, false
	}
	return o.UniqueVersionId, true
}

// HasUniqueVersionId returns a boolean if a field has been set.
func (o *BTDocumentWithVersionAndElementId) HasUniqueVersionId() bool {
	if o != nil && o.UniqueVersionId != nil {
		return true
	}

	return false
}

// SetUniqueVersionId gets a reference to the given string and assigns it to the UniqueVersionId field.
func (o *BTDocumentWithVersionAndElementId) SetUniqueVersionId(v string) {
	o.UniqueVersionId = &v
}

// GetValidElementLibraryReference returns the ValidElementLibraryReference field value if set, zero value otherwise.
func (o *BTDocumentWithVersionAndElementId) GetValidElementLibraryReference() bool {
	if o == nil || o.ValidElementLibraryReference == nil {
		var ret bool
		return ret
	}
	return *o.ValidElementLibraryReference
}

// GetValidElementLibraryReferenceOk returns a tuple with the ValidElementLibraryReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentWithVersionAndElementId) GetValidElementLibraryReferenceOk() (*bool, bool) {
	if o == nil || o.ValidElementLibraryReference == nil {
		return nil, false
	}
	return o.ValidElementLibraryReference, true
}

// HasValidElementLibraryReference returns a boolean if a field has been set.
func (o *BTDocumentWithVersionAndElementId) HasValidElementLibraryReference() bool {
	if o != nil && o.ValidElementLibraryReference != nil {
		return true
	}

	return false
}

// SetValidElementLibraryReference gets a reference to the given bool and assigns it to the ValidElementLibraryReference field.
func (o *BTDocumentWithVersionAndElementId) SetValidElementLibraryReference(v bool) {
	o.ValidElementLibraryReference = &v
}

// GetValidRevisionReference returns the ValidRevisionReference field value if set, zero value otherwise.
func (o *BTDocumentWithVersionAndElementId) GetValidRevisionReference() bool {
	if o == nil || o.ValidRevisionReference == nil {
		var ret bool
		return ret
	}
	return *o.ValidRevisionReference
}

// GetValidRevisionReferenceOk returns a tuple with the ValidRevisionReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDocumentWithVersionAndElementId) GetValidRevisionReferenceOk() (*bool, bool) {
	if o == nil || o.ValidRevisionReference == nil {
		return nil, false
	}
	return o.ValidRevisionReference, true
}

// HasValidRevisionReference returns a boolean if a field has been set.
func (o *BTDocumentWithVersionAndElementId) HasValidRevisionReference() bool {
	if o != nil && o.ValidRevisionReference != nil {
		return true
	}

	return false
}

// SetValidRevisionReference gets a reference to the given bool and assigns it to the ValidRevisionReference field.
func (o *BTDocumentWithVersionAndElementId) SetValidRevisionReference(v bool) {
	o.ValidRevisionReference = &v
}

func (o BTDocumentWithVersionAndElementId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.DocumentVersionId != nil {
		toSerialize["documentVersionId"] = o.DocumentVersionId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.ElementLibraryId != nil {
		toSerialize["elementLibraryId"] = o.ElementLibraryId
	}
	if o.ElementLibraryVersion != nil {
		toSerialize["elementLibraryVersion"] = o.ElementLibraryVersion
	}
	if o.PartNumber != nil {
		toSerialize["partNumber"] = o.PartNumber
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	if o.UniqueVersionId != nil {
		toSerialize["uniqueVersionId"] = o.UniqueVersionId
	}
	if o.ValidElementLibraryReference != nil {
		toSerialize["validElementLibraryReference"] = o.ValidElementLibraryReference
	}
	if o.ValidRevisionReference != nil {
		toSerialize["validRevisionReference"] = o.ValidRevisionReference
	}
	return json.Marshal(toSerialize)
}

type NullableBTDocumentWithVersionAndElementId struct {
	value *BTDocumentWithVersionAndElementId
	isSet bool
}

func (v NullableBTDocumentWithVersionAndElementId) Get() *BTDocumentWithVersionAndElementId {
	return v.value
}

func (v *NullableBTDocumentWithVersionAndElementId) Set(val *BTDocumentWithVersionAndElementId) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDocumentWithVersionAndElementId) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDocumentWithVersionAndElementId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDocumentWithVersionAndElementId(val *BTDocumentWithVersionAndElementId) *NullableBTDocumentWithVersionAndElementId {
	return &NullableBTDocumentWithVersionAndElementId{value: val, isSet: true}
}

func (v NullableBTDocumentWithVersionAndElementId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDocumentWithVersionAndElementId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
