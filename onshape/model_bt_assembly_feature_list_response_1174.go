/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6415-48a6b2252b8c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTAssemblyFeatureListResponse1174 struct for BTAssemblyFeatureListResponse1174
type BTAssemblyFeatureListResponse1174 struct {
	LibraryVersion         *int32                         `json:"libraryVersion,omitempty"`
	MicroversionSkew       *bool                          `json:"microversionSkew,omitempty"`
	RejectMicroversionSkew *bool                          `json:"rejectMicroversionSkew,omitempty"`
	SerializationVersion   *string                        `json:"serializationVersion,omitempty"`
	SourceMicroversion     *string                        `json:"sourceMicroversion,omitempty"`
	BtType                 *string                        `json:"btType,omitempty"`
	FeatureStates          *map[string]BTFeatureState1688 `json:"featureStates,omitempty"`
	Features               []BTMAssemblyFeature887        `json:"features,omitempty"`
	IsComplete             *bool                          `json:"isComplete,omitempty"`
}

// NewBTAssemblyFeatureListResponse1174 instantiates a new BTAssemblyFeatureListResponse1174 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTAssemblyFeatureListResponse1174() *BTAssemblyFeatureListResponse1174 {
	this := BTAssemblyFeatureListResponse1174{}
	return &this
}

// NewBTAssemblyFeatureListResponse1174WithDefaults instantiates a new BTAssemblyFeatureListResponse1174 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTAssemblyFeatureListResponse1174WithDefaults() *BTAssemblyFeatureListResponse1174 {
	this := BTAssemblyFeatureListResponse1174{}
	return &this
}

// GetLibraryVersion returns the LibraryVersion field value if set, zero value otherwise.
func (o *BTAssemblyFeatureListResponse1174) GetLibraryVersion() int32 {
	if o == nil || o.LibraryVersion == nil {
		var ret int32
		return ret
	}
	return *o.LibraryVersion
}

// GetLibraryVersionOk returns a tuple with the LibraryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyFeatureListResponse1174) GetLibraryVersionOk() (*int32, bool) {
	if o == nil || o.LibraryVersion == nil {
		return nil, false
	}
	return o.LibraryVersion, true
}

// HasLibraryVersion returns a boolean if a field has been set.
func (o *BTAssemblyFeatureListResponse1174) HasLibraryVersion() bool {
	if o != nil && o.LibraryVersion != nil {
		return true
	}

	return false
}

// SetLibraryVersion gets a reference to the given int32 and assigns it to the LibraryVersion field.
func (o *BTAssemblyFeatureListResponse1174) SetLibraryVersion(v int32) {
	o.LibraryVersion = &v
}

// GetMicroversionSkew returns the MicroversionSkew field value if set, zero value otherwise.
func (o *BTAssemblyFeatureListResponse1174) GetMicroversionSkew() bool {
	if o == nil || o.MicroversionSkew == nil {
		var ret bool
		return ret
	}
	return *o.MicroversionSkew
}

// GetMicroversionSkewOk returns a tuple with the MicroversionSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyFeatureListResponse1174) GetMicroversionSkewOk() (*bool, bool) {
	if o == nil || o.MicroversionSkew == nil {
		return nil, false
	}
	return o.MicroversionSkew, true
}

// HasMicroversionSkew returns a boolean if a field has been set.
func (o *BTAssemblyFeatureListResponse1174) HasMicroversionSkew() bool {
	if o != nil && o.MicroversionSkew != nil {
		return true
	}

	return false
}

// SetMicroversionSkew gets a reference to the given bool and assigns it to the MicroversionSkew field.
func (o *BTAssemblyFeatureListResponse1174) SetMicroversionSkew(v bool) {
	o.MicroversionSkew = &v
}

// GetRejectMicroversionSkew returns the RejectMicroversionSkew field value if set, zero value otherwise.
func (o *BTAssemblyFeatureListResponse1174) GetRejectMicroversionSkew() bool {
	if o == nil || o.RejectMicroversionSkew == nil {
		var ret bool
		return ret
	}
	return *o.RejectMicroversionSkew
}

// GetRejectMicroversionSkewOk returns a tuple with the RejectMicroversionSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyFeatureListResponse1174) GetRejectMicroversionSkewOk() (*bool, bool) {
	if o == nil || o.RejectMicroversionSkew == nil {
		return nil, false
	}
	return o.RejectMicroversionSkew, true
}

// HasRejectMicroversionSkew returns a boolean if a field has been set.
func (o *BTAssemblyFeatureListResponse1174) HasRejectMicroversionSkew() bool {
	if o != nil && o.RejectMicroversionSkew != nil {
		return true
	}

	return false
}

// SetRejectMicroversionSkew gets a reference to the given bool and assigns it to the RejectMicroversionSkew field.
func (o *BTAssemblyFeatureListResponse1174) SetRejectMicroversionSkew(v bool) {
	o.RejectMicroversionSkew = &v
}

// GetSerializationVersion returns the SerializationVersion field value if set, zero value otherwise.
func (o *BTAssemblyFeatureListResponse1174) GetSerializationVersion() string {
	if o == nil || o.SerializationVersion == nil {
		var ret string
		return ret
	}
	return *o.SerializationVersion
}

// GetSerializationVersionOk returns a tuple with the SerializationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyFeatureListResponse1174) GetSerializationVersionOk() (*string, bool) {
	if o == nil || o.SerializationVersion == nil {
		return nil, false
	}
	return o.SerializationVersion, true
}

// HasSerializationVersion returns a boolean if a field has been set.
func (o *BTAssemblyFeatureListResponse1174) HasSerializationVersion() bool {
	if o != nil && o.SerializationVersion != nil {
		return true
	}

	return false
}

// SetSerializationVersion gets a reference to the given string and assigns it to the SerializationVersion field.
func (o *BTAssemblyFeatureListResponse1174) SetSerializationVersion(v string) {
	o.SerializationVersion = &v
}

// GetSourceMicroversion returns the SourceMicroversion field value if set, zero value otherwise.
func (o *BTAssemblyFeatureListResponse1174) GetSourceMicroversion() string {
	if o == nil || o.SourceMicroversion == nil {
		var ret string
		return ret
	}
	return *o.SourceMicroversion
}

// GetSourceMicroversionOk returns a tuple with the SourceMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyFeatureListResponse1174) GetSourceMicroversionOk() (*string, bool) {
	if o == nil || o.SourceMicroversion == nil {
		return nil, false
	}
	return o.SourceMicroversion, true
}

// HasSourceMicroversion returns a boolean if a field has been set.
func (o *BTAssemblyFeatureListResponse1174) HasSourceMicroversion() bool {
	if o != nil && o.SourceMicroversion != nil {
		return true
	}

	return false
}

// SetSourceMicroversion gets a reference to the given string and assigns it to the SourceMicroversion field.
func (o *BTAssemblyFeatureListResponse1174) SetSourceMicroversion(v string) {
	o.SourceMicroversion = &v
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTAssemblyFeatureListResponse1174) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyFeatureListResponse1174) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTAssemblyFeatureListResponse1174) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTAssemblyFeatureListResponse1174) SetBtType(v string) {
	o.BtType = &v
}

// GetFeatureStates returns the FeatureStates field value if set, zero value otherwise.
func (o *BTAssemblyFeatureListResponse1174) GetFeatureStates() map[string]BTFeatureState1688 {
	if o == nil || o.FeatureStates == nil {
		var ret map[string]BTFeatureState1688
		return ret
	}
	return *o.FeatureStates
}

// GetFeatureStatesOk returns a tuple with the FeatureStates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyFeatureListResponse1174) GetFeatureStatesOk() (*map[string]BTFeatureState1688, bool) {
	if o == nil || o.FeatureStates == nil {
		return nil, false
	}
	return o.FeatureStates, true
}

// HasFeatureStates returns a boolean if a field has been set.
func (o *BTAssemblyFeatureListResponse1174) HasFeatureStates() bool {
	if o != nil && o.FeatureStates != nil {
		return true
	}

	return false
}

// SetFeatureStates gets a reference to the given map[string]BTFeatureState1688 and assigns it to the FeatureStates field.
func (o *BTAssemblyFeatureListResponse1174) SetFeatureStates(v map[string]BTFeatureState1688) {
	o.FeatureStates = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *BTAssemblyFeatureListResponse1174) GetFeatures() []BTMAssemblyFeature887 {
	if o == nil || o.Features == nil {
		var ret []BTMAssemblyFeature887
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyFeatureListResponse1174) GetFeaturesOk() ([]BTMAssemblyFeature887, bool) {
	if o == nil || o.Features == nil {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *BTAssemblyFeatureListResponse1174) HasFeatures() bool {
	if o != nil && o.Features != nil {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []BTMAssemblyFeature887 and assigns it to the Features field.
func (o *BTAssemblyFeatureListResponse1174) SetFeatures(v []BTMAssemblyFeature887) {
	o.Features = v
}

// GetIsComplete returns the IsComplete field value if set, zero value otherwise.
func (o *BTAssemblyFeatureListResponse1174) GetIsComplete() bool {
	if o == nil || o.IsComplete == nil {
		var ret bool
		return ret
	}
	return *o.IsComplete
}

// GetIsCompleteOk returns a tuple with the IsComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTAssemblyFeatureListResponse1174) GetIsCompleteOk() (*bool, bool) {
	if o == nil || o.IsComplete == nil {
		return nil, false
	}
	return o.IsComplete, true
}

// HasIsComplete returns a boolean if a field has been set.
func (o *BTAssemblyFeatureListResponse1174) HasIsComplete() bool {
	if o != nil && o.IsComplete != nil {
		return true
	}

	return false
}

// SetIsComplete gets a reference to the given bool and assigns it to the IsComplete field.
func (o *BTAssemblyFeatureListResponse1174) SetIsComplete(v bool) {
	o.IsComplete = &v
}

func (o BTAssemblyFeatureListResponse1174) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.FeatureStates != nil {
		toSerialize["featureStates"] = o.FeatureStates
	}
	if o.Features != nil {
		toSerialize["features"] = o.Features
	}
	if o.IsComplete != nil {
		toSerialize["isComplete"] = o.IsComplete
	}
	return json.Marshal(toSerialize)
}

type NullableBTAssemblyFeatureListResponse1174 struct {
	value *BTAssemblyFeatureListResponse1174
	isSet bool
}

func (v NullableBTAssemblyFeatureListResponse1174) Get() *BTAssemblyFeatureListResponse1174 {
	return v.value
}

func (v *NullableBTAssemblyFeatureListResponse1174) Set(val *BTAssemblyFeatureListResponse1174) {
	v.value = val
	v.isSet = true
}

func (v NullableBTAssemblyFeatureListResponse1174) IsSet() bool {
	return v.isSet
}

func (v *NullableBTAssemblyFeatureListResponse1174) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTAssemblyFeatureListResponse1174(val *BTAssemblyFeatureListResponse1174) *NullableBTAssemblyFeatureListResponse1174 {
	return &NullableBTAssemblyFeatureListResponse1174{value: val, isSet: true}
}

func (v NullableBTAssemblyFeatureListResponse1174) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTAssemblyFeatureListResponse1174) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
