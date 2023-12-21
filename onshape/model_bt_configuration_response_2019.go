/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.27960-e195de6ac56c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTConfigurationResponse2019 struct for BTConfigurationResponse2019
type BTConfigurationResponse2019 struct {
	BtType                  *string                        `json:"btType,omitempty"`
	ConfigurationParameters []BTMConfigurationParameter819 `json:"configurationParameters,omitempty"`
	CurrentConfiguration    []BTMParameter1                `json:"currentConfiguration,omitempty"`
	LibraryVersion          *int32                         `json:"libraryVersion,omitempty"`
	MicroversionSkew        *bool                          `json:"microversionSkew,omitempty"`
	RejectMicroversionSkew  *bool                          `json:"rejectMicroversionSkew,omitempty"`
	SerializationVersion    *string                        `json:"serializationVersion,omitempty"`
	SourceMicroversion      *string                        `json:"sourceMicroversion,omitempty"`
}

// NewBTConfigurationResponse2019 instantiates a new BTConfigurationResponse2019 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTConfigurationResponse2019() *BTConfigurationResponse2019 {
	this := BTConfigurationResponse2019{}
	return &this
}

// NewBTConfigurationResponse2019WithDefaults instantiates a new BTConfigurationResponse2019 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTConfigurationResponse2019WithDefaults() *BTConfigurationResponse2019 {
	this := BTConfigurationResponse2019{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTConfigurationResponse2019) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationResponse2019) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTConfigurationResponse2019) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTConfigurationResponse2019) SetBtType(v string) {
	o.BtType = &v
}

// GetConfigurationParameters returns the ConfigurationParameters field value if set, zero value otherwise.
func (o *BTConfigurationResponse2019) GetConfigurationParameters() []BTMConfigurationParameter819 {
	if o == nil || o.ConfigurationParameters == nil {
		var ret []BTMConfigurationParameter819
		return ret
	}
	return o.ConfigurationParameters
}

// GetConfigurationParametersOk returns a tuple with the ConfigurationParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationResponse2019) GetConfigurationParametersOk() ([]BTMConfigurationParameter819, bool) {
	if o == nil || o.ConfigurationParameters == nil {
		return nil, false
	}
	return o.ConfigurationParameters, true
}

// HasConfigurationParameters returns a boolean if a field has been set.
func (o *BTConfigurationResponse2019) HasConfigurationParameters() bool {
	if o != nil && o.ConfigurationParameters != nil {
		return true
	}

	return false
}

// SetConfigurationParameters gets a reference to the given []BTMConfigurationParameter819 and assigns it to the ConfigurationParameters field.
func (o *BTConfigurationResponse2019) SetConfigurationParameters(v []BTMConfigurationParameter819) {
	o.ConfigurationParameters = v
}

// GetCurrentConfiguration returns the CurrentConfiguration field value if set, zero value otherwise.
func (o *BTConfigurationResponse2019) GetCurrentConfiguration() []BTMParameter1 {
	if o == nil || o.CurrentConfiguration == nil {
		var ret []BTMParameter1
		return ret
	}
	return o.CurrentConfiguration
}

// GetCurrentConfigurationOk returns a tuple with the CurrentConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationResponse2019) GetCurrentConfigurationOk() ([]BTMParameter1, bool) {
	if o == nil || o.CurrentConfiguration == nil {
		return nil, false
	}
	return o.CurrentConfiguration, true
}

// HasCurrentConfiguration returns a boolean if a field has been set.
func (o *BTConfigurationResponse2019) HasCurrentConfiguration() bool {
	if o != nil && o.CurrentConfiguration != nil {
		return true
	}

	return false
}

// SetCurrentConfiguration gets a reference to the given []BTMParameter1 and assigns it to the CurrentConfiguration field.
func (o *BTConfigurationResponse2019) SetCurrentConfiguration(v []BTMParameter1) {
	o.CurrentConfiguration = v
}

// GetLibraryVersion returns the LibraryVersion field value if set, zero value otherwise.
func (o *BTConfigurationResponse2019) GetLibraryVersion() int32 {
	if o == nil || o.LibraryVersion == nil {
		var ret int32
		return ret
	}
	return *o.LibraryVersion
}

// GetLibraryVersionOk returns a tuple with the LibraryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationResponse2019) GetLibraryVersionOk() (*int32, bool) {
	if o == nil || o.LibraryVersion == nil {
		return nil, false
	}
	return o.LibraryVersion, true
}

// HasLibraryVersion returns a boolean if a field has been set.
func (o *BTConfigurationResponse2019) HasLibraryVersion() bool {
	if o != nil && o.LibraryVersion != nil {
		return true
	}

	return false
}

// SetLibraryVersion gets a reference to the given int32 and assigns it to the LibraryVersion field.
func (o *BTConfigurationResponse2019) SetLibraryVersion(v int32) {
	o.LibraryVersion = &v
}

// GetMicroversionSkew returns the MicroversionSkew field value if set, zero value otherwise.
func (o *BTConfigurationResponse2019) GetMicroversionSkew() bool {
	if o == nil || o.MicroversionSkew == nil {
		var ret bool
		return ret
	}
	return *o.MicroversionSkew
}

// GetMicroversionSkewOk returns a tuple with the MicroversionSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationResponse2019) GetMicroversionSkewOk() (*bool, bool) {
	if o == nil || o.MicroversionSkew == nil {
		return nil, false
	}
	return o.MicroversionSkew, true
}

// HasMicroversionSkew returns a boolean if a field has been set.
func (o *BTConfigurationResponse2019) HasMicroversionSkew() bool {
	if o != nil && o.MicroversionSkew != nil {
		return true
	}

	return false
}

// SetMicroversionSkew gets a reference to the given bool and assigns it to the MicroversionSkew field.
func (o *BTConfigurationResponse2019) SetMicroversionSkew(v bool) {
	o.MicroversionSkew = &v
}

// GetRejectMicroversionSkew returns the RejectMicroversionSkew field value if set, zero value otherwise.
func (o *BTConfigurationResponse2019) GetRejectMicroversionSkew() bool {
	if o == nil || o.RejectMicroversionSkew == nil {
		var ret bool
		return ret
	}
	return *o.RejectMicroversionSkew
}

// GetRejectMicroversionSkewOk returns a tuple with the RejectMicroversionSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationResponse2019) GetRejectMicroversionSkewOk() (*bool, bool) {
	if o == nil || o.RejectMicroversionSkew == nil {
		return nil, false
	}
	return o.RejectMicroversionSkew, true
}

// HasRejectMicroversionSkew returns a boolean if a field has been set.
func (o *BTConfigurationResponse2019) HasRejectMicroversionSkew() bool {
	if o != nil && o.RejectMicroversionSkew != nil {
		return true
	}

	return false
}

// SetRejectMicroversionSkew gets a reference to the given bool and assigns it to the RejectMicroversionSkew field.
func (o *BTConfigurationResponse2019) SetRejectMicroversionSkew(v bool) {
	o.RejectMicroversionSkew = &v
}

// GetSerializationVersion returns the SerializationVersion field value if set, zero value otherwise.
func (o *BTConfigurationResponse2019) GetSerializationVersion() string {
	if o == nil || o.SerializationVersion == nil {
		var ret string
		return ret
	}
	return *o.SerializationVersion
}

// GetSerializationVersionOk returns a tuple with the SerializationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationResponse2019) GetSerializationVersionOk() (*string, bool) {
	if o == nil || o.SerializationVersion == nil {
		return nil, false
	}
	return o.SerializationVersion, true
}

// HasSerializationVersion returns a boolean if a field has been set.
func (o *BTConfigurationResponse2019) HasSerializationVersion() bool {
	if o != nil && o.SerializationVersion != nil {
		return true
	}

	return false
}

// SetSerializationVersion gets a reference to the given string and assigns it to the SerializationVersion field.
func (o *BTConfigurationResponse2019) SetSerializationVersion(v string) {
	o.SerializationVersion = &v
}

// GetSourceMicroversion returns the SourceMicroversion field value if set, zero value otherwise.
func (o *BTConfigurationResponse2019) GetSourceMicroversion() string {
	if o == nil || o.SourceMicroversion == nil {
		var ret string
		return ret
	}
	return *o.SourceMicroversion
}

// GetSourceMicroversionOk returns a tuple with the SourceMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTConfigurationResponse2019) GetSourceMicroversionOk() (*string, bool) {
	if o == nil || o.SourceMicroversion == nil {
		return nil, false
	}
	return o.SourceMicroversion, true
}

// HasSourceMicroversion returns a boolean if a field has been set.
func (o *BTConfigurationResponse2019) HasSourceMicroversion() bool {
	if o != nil && o.SourceMicroversion != nil {
		return true
	}

	return false
}

// SetSourceMicroversion gets a reference to the given string and assigns it to the SourceMicroversion field.
func (o *BTConfigurationResponse2019) SetSourceMicroversion(v string) {
	o.SourceMicroversion = &v
}

func (o BTConfigurationResponse2019) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ConfigurationParameters != nil {
		toSerialize["configurationParameters"] = o.ConfigurationParameters
	}
	if o.CurrentConfiguration != nil {
		toSerialize["currentConfiguration"] = o.CurrentConfiguration
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
	return json.Marshal(toSerialize)
}

type NullableBTConfigurationResponse2019 struct {
	value *BTConfigurationResponse2019
	isSet bool
}

func (v NullableBTConfigurationResponse2019) Get() *BTConfigurationResponse2019 {
	return v.value
}

func (v *NullableBTConfigurationResponse2019) Set(val *BTConfigurationResponse2019) {
	v.value = val
	v.isSet = true
}

func (v NullableBTConfigurationResponse2019) IsSet() bool {
	return v.isSet
}

func (v *NullableBTConfigurationResponse2019) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTConfigurationResponse2019(val *BTConfigurationResponse2019) *NullableBTConfigurationResponse2019 {
	return &NullableBTConfigurationResponse2019{value: val, isSet: true}
}

func (v NullableBTConfigurationResponse2019) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTConfigurationResponse2019) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
