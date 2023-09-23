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

// BTUpdateFeaturesCall1748 struct for BTUpdateFeaturesCall1748
type BTUpdateFeaturesCall1748 struct {
	BtType                      *string         `json:"btType,omitempty"`
	Features                    []BTMFeature134 `json:"features,omitempty"`
	LibraryVersion              *int32          `json:"libraryVersion,omitempty"`
	MicroversionSkew            *bool           `json:"microversionSkew,omitempty"`
	RejectMicroversionSkew      *bool           `json:"rejectMicroversionSkew,omitempty"`
	SerializationVersion        *string         `json:"serializationVersion,omitempty"`
	SourceMicroversion          *string         `json:"sourceMicroversion,omitempty"`
	UpdateSuppressionAttributes *bool           `json:"updateSuppressionAttributes,omitempty"`
}

// NewBTUpdateFeaturesCall1748 instantiates a new BTUpdateFeaturesCall1748 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUpdateFeaturesCall1748() *BTUpdateFeaturesCall1748 {
	this := BTUpdateFeaturesCall1748{}
	return &this
}

// NewBTUpdateFeaturesCall1748WithDefaults instantiates a new BTUpdateFeaturesCall1748 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUpdateFeaturesCall1748WithDefaults() *BTUpdateFeaturesCall1748 {
	this := BTUpdateFeaturesCall1748{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTUpdateFeaturesCall1748) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUpdateFeaturesCall1748) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTUpdateFeaturesCall1748) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTUpdateFeaturesCall1748) SetBtType(v string) {
	o.BtType = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *BTUpdateFeaturesCall1748) GetFeatures() []BTMFeature134 {
	if o == nil || o.Features == nil {
		var ret []BTMFeature134
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUpdateFeaturesCall1748) GetFeaturesOk() ([]BTMFeature134, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *BTUpdateFeaturesCall1748) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []BTMFeature134 and assigns it to the Features field.
func (o *BTUpdateFeaturesCall1748) SetFeatures(v []BTMFeature134) {
	o.Features = v
}

// GetLibraryVersion returns the LibraryVersion field value if set, zero value otherwise.
func (o *BTUpdateFeaturesCall1748) GetLibraryVersion() int32 {
	if o == nil || o.LibraryVersion == nil {
		var ret int32
		return ret
	}
	return *o.LibraryVersion
}

// GetLibraryVersionOk returns a tuple with the LibraryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUpdateFeaturesCall1748) GetLibraryVersionOk() (*int32, bool) {
	if o == nil || o.LibraryVersion == nil {
		return nil, false
	}
	return o.LibraryVersion, true
}

// HasLibraryVersion returns a boolean if a field has been set.
func (o *BTUpdateFeaturesCall1748) HasLibraryVersion() bool {
	if o != nil && o.LibraryVersion != nil {
		return true
	}

	return false
}

// SetLibraryVersion gets a reference to the given int32 and assigns it to the LibraryVersion field.
func (o *BTUpdateFeaturesCall1748) SetLibraryVersion(v int32) {
	o.LibraryVersion = &v
}

// GetMicroversionSkew returns the MicroversionSkew field value if set, zero value otherwise.
func (o *BTUpdateFeaturesCall1748) GetMicroversionSkew() bool {
	if o == nil || o.MicroversionSkew == nil {
		var ret bool
		return ret
	}
	return *o.MicroversionSkew
}

// GetMicroversionSkewOk returns a tuple with the MicroversionSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUpdateFeaturesCall1748) GetMicroversionSkewOk() (*bool, bool) {
	if o == nil || o.MicroversionSkew == nil {
		return nil, false
	}
	return o.MicroversionSkew, true
}

// HasMicroversionSkew returns a boolean if a field has been set.
func (o *BTUpdateFeaturesCall1748) HasMicroversionSkew() bool {
	if o != nil && o.MicroversionSkew != nil {
		return true
	}

	return false
}

// SetMicroversionSkew gets a reference to the given bool and assigns it to the MicroversionSkew field.
func (o *BTUpdateFeaturesCall1748) SetMicroversionSkew(v bool) {
	o.MicroversionSkew = &v
}

// GetRejectMicroversionSkew returns the RejectMicroversionSkew field value if set, zero value otherwise.
func (o *BTUpdateFeaturesCall1748) GetRejectMicroversionSkew() bool {
	if o == nil || o.RejectMicroversionSkew == nil {
		var ret bool
		return ret
	}
	return *o.RejectMicroversionSkew
}

// GetRejectMicroversionSkewOk returns a tuple with the RejectMicroversionSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUpdateFeaturesCall1748) GetRejectMicroversionSkewOk() (*bool, bool) {
	if o == nil || o.RejectMicroversionSkew == nil {
		return nil, false
	}
	return o.RejectMicroversionSkew, true
}

// HasRejectMicroversionSkew returns a boolean if a field has been set.
func (o *BTUpdateFeaturesCall1748) HasRejectMicroversionSkew() bool {
	if o != nil && o.RejectMicroversionSkew != nil {
		return true
	}

	return false
}

// SetRejectMicroversionSkew gets a reference to the given bool and assigns it to the RejectMicroversionSkew field.
func (o *BTUpdateFeaturesCall1748) SetRejectMicroversionSkew(v bool) {
	o.RejectMicroversionSkew = &v
}

// GetSerializationVersion returns the SerializationVersion field value if set, zero value otherwise.
func (o *BTUpdateFeaturesCall1748) GetSerializationVersion() string {
	if o == nil || o.SerializationVersion == nil {
		var ret string
		return ret
	}
	return *o.SerializationVersion
}

// GetSerializationVersionOk returns a tuple with the SerializationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUpdateFeaturesCall1748) GetSerializationVersionOk() (*string, bool) {
	if o == nil || o.SerializationVersion == nil {
		return nil, false
	}
	return o.SerializationVersion, true
}

// HasSerializationVersion returns a boolean if a field has been set.
func (o *BTUpdateFeaturesCall1748) HasSerializationVersion() bool {
	if o != nil && o.SerializationVersion != nil {
		return true
	}

	return false
}

// SetSerializationVersion gets a reference to the given string and assigns it to the SerializationVersion field.
func (o *BTUpdateFeaturesCall1748) SetSerializationVersion(v string) {
	o.SerializationVersion = &v
}

// GetSourceMicroversion returns the SourceMicroversion field value if set, zero value otherwise.
func (o *BTUpdateFeaturesCall1748) GetSourceMicroversion() string {
	if o == nil || o.SourceMicroversion == nil {
		var ret string
		return ret
	}
	return *o.SourceMicroversion
}

// GetSourceMicroversionOk returns a tuple with the SourceMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUpdateFeaturesCall1748) GetSourceMicroversionOk() (*string, bool) {
	if o == nil || o.SourceMicroversion == nil {
		return nil, false
	}
	return o.SourceMicroversion, true
}

// HasSourceMicroversion returns a boolean if a field has been set.
func (o *BTUpdateFeaturesCall1748) HasSourceMicroversion() bool {
	if o != nil && o.SourceMicroversion != nil {
		return true
	}

	return false
}

// SetSourceMicroversion gets a reference to the given string and assigns it to the SourceMicroversion field.
func (o *BTUpdateFeaturesCall1748) SetSourceMicroversion(v string) {
	o.SourceMicroversion = &v
}

// GetUpdateSuppressionAttributes returns the UpdateSuppressionAttributes field value if set, zero value otherwise.
func (o *BTUpdateFeaturesCall1748) GetUpdateSuppressionAttributes() bool {
	if o == nil || o.UpdateSuppressionAttributes == nil {
		var ret bool
		return ret
	}
	return *o.UpdateSuppressionAttributes
}

// GetUpdateSuppressionAttributesOk returns a tuple with the UpdateSuppressionAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUpdateFeaturesCall1748) GetUpdateSuppressionAttributesOk() (*bool, bool) {
	if o == nil || o.UpdateSuppressionAttributes == nil {
		return nil, false
	}
	return o.UpdateSuppressionAttributes, true
}

// HasUpdateSuppressionAttributes returns a boolean if a field has been set.
func (o *BTUpdateFeaturesCall1748) HasUpdateSuppressionAttributes() bool {
	if o != nil && o.UpdateSuppressionAttributes != nil {
		return true
	}

	return false
}

// SetUpdateSuppressionAttributes gets a reference to the given bool and assigns it to the UpdateSuppressionAttributes field.
func (o *BTUpdateFeaturesCall1748) SetUpdateSuppressionAttributes(v bool) {
	o.UpdateSuppressionAttributes = &v
}

func (o BTUpdateFeaturesCall1748) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.LibraryVersion != nil {
		toSerialize["libraryVersion"] = o.LibraryVersion
	}
	if o.MicroversionSkew != nil {
		toSerialize["microversionSkew"] = o.MicroversionSkew
	}
	if o.RejectMicroversionSkew != nil {
		toSerialize["rejectMicroversionSkew"] = o.RejectMicroversionSkew
	}
	if o.SerializationVersion != nil {
		toSerialize["serializationVersion"] = o.SerializationVersion
	}
	if o.SourceMicroversion != nil {
		toSerialize["sourceMicroversion"] = o.SourceMicroversion
	}
	if o.UpdateSuppressionAttributes != nil {
		toSerialize["updateSuppressionAttributes"] = o.UpdateSuppressionAttributes
	}
	return json.Marshal(toSerialize)
}

type NullableBTUpdateFeaturesCall1748 struct {
	value *BTUpdateFeaturesCall1748
	isSet bool
}

func (v NullableBTUpdateFeaturesCall1748) Get() *BTUpdateFeaturesCall1748 {
	return v.value
}

func (v *NullableBTUpdateFeaturesCall1748) Set(val *BTUpdateFeaturesCall1748) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUpdateFeaturesCall1748) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUpdateFeaturesCall1748) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUpdateFeaturesCall1748(val *BTUpdateFeaturesCall1748) *NullableBTUpdateFeaturesCall1748 {
	return &NullableBTUpdateFeaturesCall1748{value: val, isSet: true}
}

func (v NullableBTUpdateFeaturesCall1748) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUpdateFeaturesCall1748) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
