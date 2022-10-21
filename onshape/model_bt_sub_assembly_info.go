/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7010-841c6a8f62e7
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTSubAssemblyInfo struct for BTSubAssemblyInfo
type BTSubAssemblyInfo struct {
	Configuration        *string                  `json:"configuration,omitempty"`
	DocumentId           *string                  `json:"documentId,omitempty"`
	DocumentMicroversion *string                  `json:"documentMicroversion,omitempty"`
	DocumentVersion      *string                  `json:"documentVersion,omitempty"`
	ElementId            *string                  `json:"elementId,omitempty"`
	Features             []BTAssemblyFeatureInfo  `json:"features,omitempty"`
	FullConfiguration    *string                  `json:"fullConfiguration,omitempty"`
	Instances            []BTAssemblyInstanceInfo `json:"instances,omitempty"`
	PartNumber           *string                  `json:"partNumber,omitempty"`
	Revision             *string                  `json:"revision,omitempty"`
}

// NewBTSubAssemblyInfo instantiates a new BTSubAssemblyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTSubAssemblyInfo() *BTSubAssemblyInfo {
	this := BTSubAssemblyInfo{}
	return &this
}

// NewBTSubAssemblyInfoWithDefaults instantiates a new BTSubAssemblyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTSubAssemblyInfoWithDefaults() *BTSubAssemblyInfo {
	this := BTSubAssemblyInfo{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *BTSubAssemblyInfo) GetConfiguration() string {
	if o == nil || o.Configuration == nil {
		var ret string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSubAssemblyInfo) GetConfigurationOk() (*string, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *BTSubAssemblyInfo) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given string and assigns it to the Configuration field.
func (o *BTSubAssemblyInfo) SetConfiguration(v string) {
	o.Configuration = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTSubAssemblyInfo) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSubAssemblyInfo) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTSubAssemblyInfo) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTSubAssemblyInfo) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetDocumentMicroversion returns the DocumentMicroversion field value if set, zero value otherwise.
func (o *BTSubAssemblyInfo) GetDocumentMicroversion() string {
	if o == nil || o.DocumentMicroversion == nil {
		var ret string
		return ret
	}
	return *o.DocumentMicroversion
}

// GetDocumentMicroversionOk returns a tuple with the DocumentMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSubAssemblyInfo) GetDocumentMicroversionOk() (*string, bool) {
	if o == nil || o.DocumentMicroversion == nil {
		return nil, false
	}
	return o.DocumentMicroversion, true
}

// HasDocumentMicroversion returns a boolean if a field has been set.
func (o *BTSubAssemblyInfo) HasDocumentMicroversion() bool {
	if o != nil && o.DocumentMicroversion != nil {
		return true
	}

	return false
}

// SetDocumentMicroversion gets a reference to the given string and assigns it to the DocumentMicroversion field.
func (o *BTSubAssemblyInfo) SetDocumentMicroversion(v string) {
	o.DocumentMicroversion = &v
}

// GetDocumentVersion returns the DocumentVersion field value if set, zero value otherwise.
func (o *BTSubAssemblyInfo) GetDocumentVersion() string {
	if o == nil || o.DocumentVersion == nil {
		var ret string
		return ret
	}
	return *o.DocumentVersion
}

// GetDocumentVersionOk returns a tuple with the DocumentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSubAssemblyInfo) GetDocumentVersionOk() (*string, bool) {
	if o == nil || o.DocumentVersion == nil {
		return nil, false
	}
	return o.DocumentVersion, true
}

// HasDocumentVersion returns a boolean if a field has been set.
func (o *BTSubAssemblyInfo) HasDocumentVersion() bool {
	if o != nil && o.DocumentVersion != nil {
		return true
	}

	return false
}

// SetDocumentVersion gets a reference to the given string and assigns it to the DocumentVersion field.
func (o *BTSubAssemblyInfo) SetDocumentVersion(v string) {
	o.DocumentVersion = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTSubAssemblyInfo) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSubAssemblyInfo) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTSubAssemblyInfo) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTSubAssemblyInfo) SetElementId(v string) {
	o.ElementId = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *BTSubAssemblyInfo) GetFeatures() []BTAssemblyFeatureInfo {
	if o == nil || o.Features == nil {
		var ret []BTAssemblyFeatureInfo
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSubAssemblyInfo) GetFeaturesOk() ([]BTAssemblyFeatureInfo, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *BTSubAssemblyInfo) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []BTAssemblyFeatureInfo and assigns it to the Features field.
func (o *BTSubAssemblyInfo) SetFeatures(v []BTAssemblyFeatureInfo) {
	o.Features = v
}

// GetFullConfiguration returns the FullConfiguration field value if set, zero value otherwise.
func (o *BTSubAssemblyInfo) GetFullConfiguration() string {
	if o == nil || o.FullConfiguration == nil {
		var ret string
		return ret
	}
	return *o.FullConfiguration
}

// GetFullConfigurationOk returns a tuple with the FullConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSubAssemblyInfo) GetFullConfigurationOk() (*string, bool) {
	if o == nil || o.FullConfiguration == nil {
		return nil, false
	}
	return o.FullConfiguration, true
}

// HasFullConfiguration returns a boolean if a field has been set.
func (o *BTSubAssemblyInfo) HasFullConfiguration() bool {
	if o != nil && o.FullConfiguration != nil {
		return true
	}

	return false
}

// SetFullConfiguration gets a reference to the given string and assigns it to the FullConfiguration field.
func (o *BTSubAssemblyInfo) SetFullConfiguration(v string) {
	o.FullConfiguration = &v
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *BTSubAssemblyInfo) GetInstances() []BTAssemblyInstanceInfo {
	if o == nil || o.Instances == nil {
		var ret []BTAssemblyInstanceInfo
		return ret
	}
	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSubAssemblyInfo) GetInstancesOk() ([]BTAssemblyInstanceInfo, bool) {
	if o == nil || o.Instances == nil {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *BTSubAssemblyInfo) HasInstances() bool {
	if o != nil && o.Instances != nil {
		return true
	}

	return false
}

// SetInstances gets a reference to the given []BTAssemblyInstanceInfo and assigns it to the Instances field.
func (o *BTSubAssemblyInfo) SetInstances(v []BTAssemblyInstanceInfo) {
	o.Instances = v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *BTSubAssemblyInfo) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSubAssemblyInfo) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *BTSubAssemblyInfo) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *BTSubAssemblyInfo) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *BTSubAssemblyInfo) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTSubAssemblyInfo) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *BTSubAssemblyInfo) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *BTSubAssemblyInfo) SetRevision(v string) {
	o.Revision = &v
}

func (o BTSubAssemblyInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.DocumentMicroversion != nil {
		toSerialize["documentMicroversion"] = o.DocumentMicroversion
	}
	if o.DocumentVersion != nil {
		toSerialize["documentVersion"] = o.DocumentVersion
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.FullConfiguration != nil {
		toSerialize["fullConfiguration"] = o.FullConfiguration
	}
	if o.Instances != nil {
		toSerialize["instances"] = o.Instances
	}
	if o.PartNumber != nil {
		toSerialize["partNumber"] = o.PartNumber
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	return json.Marshal(toSerialize)
}

type NullableBTSubAssemblyInfo struct {
	value *BTSubAssemblyInfo
	isSet bool
}

func (v NullableBTSubAssemblyInfo) Get() *BTSubAssemblyInfo {
	return v.value
}

func (v *NullableBTSubAssemblyInfo) Set(val *BTSubAssemblyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTSubAssemblyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTSubAssemblyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTSubAssemblyInfo(val *BTSubAssemblyInfo) *NullableBTSubAssemblyInfo {
	return &NullableBTSubAssemblyInfo{value: val, isSet: true}
}

func (v NullableBTSubAssemblyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTSubAssemblyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
