/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.170.22862-4427d042758b
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTLegacySearchHit struct for BTLegacySearchHit
type BTLegacySearchHit struct {
	DocumentId        *string                           `json:"documentId,omitempty"`
	HighlightedFields *map[string][]string              `json:"highlightedFields,omitempty"`
	HitId             *string                           `json:"hitId,omitempty"`
	Name              *string                           `json:"name,omitempty"`
	ProjectId         *string                           `json:"projectId,omitempty"`
	SourceMap         map[string]map[string]interface{} `json:"sourceMap,omitempty"`
	Type              *BTSearchEntityType               `json:"type,omitempty"`
}

// NewBTLegacySearchHit instantiates a new BTLegacySearchHit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTLegacySearchHit() *BTLegacySearchHit {
	this := BTLegacySearchHit{}
	return &this
}

// NewBTLegacySearchHitWithDefaults instantiates a new BTLegacySearchHit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTLegacySearchHitWithDefaults() *BTLegacySearchHit {
	this := BTLegacySearchHit{}
	return &this
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTLegacySearchHit) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLegacySearchHit) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTLegacySearchHit) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTLegacySearchHit) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetHighlightedFields returns the HighlightedFields field value if set, zero value otherwise.
func (o *BTLegacySearchHit) GetHighlightedFields() map[string][]string {
	if o == nil || o.HighlightedFields == nil {
		var ret map[string][]string
		return ret
	}
	return *o.HighlightedFields
}

// GetHighlightedFieldsOk returns a tuple with the HighlightedFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLegacySearchHit) GetHighlightedFieldsOk() (*map[string][]string, bool) {
	if o == nil || o.HighlightedFields == nil {
		return nil, false
	}
	return o.HighlightedFields, true
}

// HasHighlightedFields returns a boolean if a field has been set.
func (o *BTLegacySearchHit) HasHighlightedFields() bool {
	if o != nil && o.HighlightedFields != nil {
		return true
	}

	return false
}

// SetHighlightedFields gets a reference to the given map[string][]string and assigns it to the HighlightedFields field.
func (o *BTLegacySearchHit) SetHighlightedFields(v map[string][]string) {
	o.HighlightedFields = &v
}

// GetHitId returns the HitId field value if set, zero value otherwise.
func (o *BTLegacySearchHit) GetHitId() string {
	if o == nil || o.HitId == nil {
		var ret string
		return ret
	}
	return *o.HitId
}

// GetHitIdOk returns a tuple with the HitId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLegacySearchHit) GetHitIdOk() (*string, bool) {
	if o == nil || o.HitId == nil {
		return nil, false
	}
	return o.HitId, true
}

// HasHitId returns a boolean if a field has been set.
func (o *BTLegacySearchHit) HasHitId() bool {
	if o != nil && o.HitId != nil {
		return true
	}

	return false
}

// SetHitId gets a reference to the given string and assigns it to the HitId field.
func (o *BTLegacySearchHit) SetHitId(v string) {
	o.HitId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTLegacySearchHit) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLegacySearchHit) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTLegacySearchHit) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTLegacySearchHit) SetName(v string) {
	o.Name = &v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *BTLegacySearchHit) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLegacySearchHit) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *BTLegacySearchHit) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *BTLegacySearchHit) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetSourceMap returns the SourceMap field value if set, zero value otherwise.
func (o *BTLegacySearchHit) GetSourceMap() map[string]map[string]interface{} {
	if o == nil || o.SourceMap == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.SourceMap
}

// GetSourceMapOk returns a tuple with the SourceMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLegacySearchHit) GetSourceMapOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.SourceMap == nil {
		return nil, false
	}
	return o.SourceMap, true
}

// HasSourceMap returns a boolean if a field has been set.
func (o *BTLegacySearchHit) HasSourceMap() bool {
	if o != nil && o.SourceMap != nil {
		return true
	}

	return false
}

// SetSourceMap gets a reference to the given map[string]map[string]interface{} and assigns it to the SourceMap field.
func (o *BTLegacySearchHit) SetSourceMap(v map[string]map[string]interface{}) {
	o.SourceMap = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *BTLegacySearchHit) GetType() BTSearchEntityType {
	if o == nil || o.Type == nil {
		var ret BTSearchEntityType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTLegacySearchHit) GetTypeOk() (*BTSearchEntityType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *BTLegacySearchHit) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given BTSearchEntityType and assigns it to the Type field.
func (o *BTLegacySearchHit) SetType(v BTSearchEntityType) {
	o.Type = &v
}

func (o BTLegacySearchHit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.HighlightedFields != nil {
		toSerialize["highlightedFields"] = o.HighlightedFields
	}
	if o.HitId != nil {
		toSerialize["hitId"] = o.HitId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.SourceMap != nil {
		toSerialize["sourceMap"] = o.SourceMap
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableBTLegacySearchHit struct {
	value *BTLegacySearchHit
	isSet bool
}

func (v NullableBTLegacySearchHit) Get() *BTLegacySearchHit {
	return v.value
}

func (v *NullableBTLegacySearchHit) Set(val *BTLegacySearchHit) {
	v.value = val
	v.isSet = true
}

func (v NullableBTLegacySearchHit) IsSet() bool {
	return v.isSet
}

func (v *NullableBTLegacySearchHit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTLegacySearchHit(val *BTLegacySearchHit) *NullableBTLegacySearchHit {
	return &NullableBTLegacySearchHit{value: val, isSet: true}
}

func (v NullableBTLegacySearchHit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTLegacySearchHit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
