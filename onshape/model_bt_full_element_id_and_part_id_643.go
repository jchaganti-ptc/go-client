/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6108-60a91d537e42
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTFullElementIdAndPartId643 struct for BTFullElementIdAndPartId643
type BTFullElementIdAndPartId643 struct {
	BtType                         *string                               `json:"btType,omitempty"`
	Configured                     *bool                                 `json:"configured,omitempty"`
	DocumentId                     *string                               `json:"documentId,omitempty"`
	ElementId                      *string                               `json:"elementId,omitempty"`
	MicroversionId                 *BTMicroversionId366                  `json:"microversionId,omitempty"`
	MicroversionIdAndConfiguration *BTMicroversionIdAndConfiguration2338 `json:"microversionIdAndConfiguration,omitempty"`
	Target                         *BTMicroversionIdAndConfiguration2338 `json:"target,omitempty"`
	PartId                         *string                               `json:"partId,omitempty"`
}

// NewBTFullElementIdAndPartId643 instantiates a new BTFullElementIdAndPartId643 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFullElementIdAndPartId643() *BTFullElementIdAndPartId643 {
	this := BTFullElementIdAndPartId643{}
	return &this
}

// NewBTFullElementIdAndPartId643WithDefaults instantiates a new BTFullElementIdAndPartId643 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFullElementIdAndPartId643WithDefaults() *BTFullElementIdAndPartId643 {
	this := BTFullElementIdAndPartId643{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTFullElementIdAndPartId643) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFullElementIdAndPartId643) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTFullElementIdAndPartId643) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTFullElementIdAndPartId643) SetBtType(v string) {
	o.BtType = &v
}

// GetConfigured returns the Configured field value if set, zero value otherwise.
func (o *BTFullElementIdAndPartId643) GetConfigured() bool {
	if o == nil || o.Configured == nil {
		var ret bool
		return ret
	}
	return *o.Configured
}

// GetConfiguredOk returns a tuple with the Configured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFullElementIdAndPartId643) GetConfiguredOk() (*bool, bool) {
	if o == nil || o.Configured == nil {
		return nil, false
	}
	return o.Configured, true
}

// HasConfigured returns a boolean if a field has been set.
func (o *BTFullElementIdAndPartId643) HasConfigured() bool {
	if o != nil && o.Configured != nil {
		return true
	}

	return false
}

// SetConfigured gets a reference to the given bool and assigns it to the Configured field.
func (o *BTFullElementIdAndPartId643) SetConfigured(v bool) {
	o.Configured = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTFullElementIdAndPartId643) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFullElementIdAndPartId643) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTFullElementIdAndPartId643) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTFullElementIdAndPartId643) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTFullElementIdAndPartId643) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFullElementIdAndPartId643) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTFullElementIdAndPartId643) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTFullElementIdAndPartId643) SetElementId(v string) {
	o.ElementId = &v
}

// GetMicroversionId returns the MicroversionId field value if set, zero value otherwise.
func (o *BTFullElementIdAndPartId643) GetMicroversionId() BTMicroversionId366 {
	if o == nil || o.MicroversionId == nil {
		var ret BTMicroversionId366
		return ret
	}
	return *o.MicroversionId
}

// GetMicroversionIdOk returns a tuple with the MicroversionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFullElementIdAndPartId643) GetMicroversionIdOk() (*BTMicroversionId366, bool) {
	if o == nil || o.MicroversionId == nil {
		return nil, false
	}
	return o.MicroversionId, true
}

// HasMicroversionId returns a boolean if a field has been set.
func (o *BTFullElementIdAndPartId643) HasMicroversionId() bool {
	if o != nil && o.MicroversionId != nil {
		return true
	}

	return false
}

// SetMicroversionId gets a reference to the given BTMicroversionId366 and assigns it to the MicroversionId field.
func (o *BTFullElementIdAndPartId643) SetMicroversionId(v BTMicroversionId366) {
	o.MicroversionId = &v
}

// GetMicroversionIdAndConfiguration returns the MicroversionIdAndConfiguration field value if set, zero value otherwise.
func (o *BTFullElementIdAndPartId643) GetMicroversionIdAndConfiguration() BTMicroversionIdAndConfiguration2338 {
	if o == nil || o.MicroversionIdAndConfiguration == nil {
		var ret BTMicroversionIdAndConfiguration2338
		return ret
	}
	return *o.MicroversionIdAndConfiguration
}

// GetMicroversionIdAndConfigurationOk returns a tuple with the MicroversionIdAndConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFullElementIdAndPartId643) GetMicroversionIdAndConfigurationOk() (*BTMicroversionIdAndConfiguration2338, bool) {
	if o == nil || o.MicroversionIdAndConfiguration == nil {
		return nil, false
	}
	return o.MicroversionIdAndConfiguration, true
}

// HasMicroversionIdAndConfiguration returns a boolean if a field has been set.
func (o *BTFullElementIdAndPartId643) HasMicroversionIdAndConfiguration() bool {
	if o != nil && o.MicroversionIdAndConfiguration != nil {
		return true
	}

	return false
}

// SetMicroversionIdAndConfiguration gets a reference to the given BTMicroversionIdAndConfiguration2338 and assigns it to the MicroversionIdAndConfiguration field.
func (o *BTFullElementIdAndPartId643) SetMicroversionIdAndConfiguration(v BTMicroversionIdAndConfiguration2338) {
	o.MicroversionIdAndConfiguration = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *BTFullElementIdAndPartId643) GetTarget() BTMicroversionIdAndConfiguration2338 {
	if o == nil || o.Target == nil {
		var ret BTMicroversionIdAndConfiguration2338
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFullElementIdAndPartId643) GetTargetOk() (*BTMicroversionIdAndConfiguration2338, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *BTFullElementIdAndPartId643) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given BTMicroversionIdAndConfiguration2338 and assigns it to the Target field.
func (o *BTFullElementIdAndPartId643) SetTarget(v BTMicroversionIdAndConfiguration2338) {
	o.Target = &v
}

// GetPartId returns the PartId field value if set, zero value otherwise.
func (o *BTFullElementIdAndPartId643) GetPartId() string {
	if o == nil || o.PartId == nil {
		var ret string
		return ret
	}
	return *o.PartId
}

// GetPartIdOk returns a tuple with the PartId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFullElementIdAndPartId643) GetPartIdOk() (*string, bool) {
	if o == nil || o.PartId == nil {
		return nil, false
	}
	return o.PartId, true
}

// HasPartId returns a boolean if a field has been set.
func (o *BTFullElementIdAndPartId643) HasPartId() bool {
	if o != nil && o.PartId != nil {
		return true
	}

	return false
}

// SetPartId gets a reference to the given string and assigns it to the PartId field.
func (o *BTFullElementIdAndPartId643) SetPartId(v string) {
	o.PartId = &v
}

func (o BTFullElementIdAndPartId643) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Configured != nil {
		toSerialize["configured"] = o.Configured
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.MicroversionId != nil {
		toSerialize["microversionId"] = o.MicroversionId
	}
	if o.MicroversionIdAndConfiguration != nil {
		toSerialize["microversionIdAndConfiguration"] = o.MicroversionIdAndConfiguration
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.PartId != nil {
		toSerialize["partId"] = o.PartId
	}
	return json.Marshal(toSerialize)
}

type NullableBTFullElementIdAndPartId643 struct {
	value *BTFullElementIdAndPartId643
	isSet bool
}

func (v NullableBTFullElementIdAndPartId643) Get() *BTFullElementIdAndPartId643 {
	return v.value
}

func (v *NullableBTFullElementIdAndPartId643) Set(val *BTFullElementIdAndPartId643) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFullElementIdAndPartId643) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFullElementIdAndPartId643) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFullElementIdAndPartId643(val *BTFullElementIdAndPartId643) *NullableBTFullElementIdAndPartId643 {
	return &NullableBTFullElementIdAndPartId643{value: val, isSet: true}
}

func (v NullableBTFullElementIdAndPartId643) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFullElementIdAndPartId643) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
