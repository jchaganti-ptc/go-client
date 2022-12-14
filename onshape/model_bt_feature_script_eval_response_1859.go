/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.157.8692-e40d034a55a7
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTFeatureScriptEvalResponse1859 struct for BTFeatureScriptEvalResponse1859
type BTFeatureScriptEvalResponse1859 struct {
	Console                *string        `json:"console,omitempty"`
	LibraryVersion         *int32         `json:"libraryVersion,omitempty"`
	MicroversionSkew       *bool          `json:"microversionSkew,omitempty"`
	Notices                []BTNotice227  `json:"notices,omitempty"`
	RejectMicroversionSkew *bool          `json:"rejectMicroversionSkew,omitempty"`
	Result                 *BTFSValue1888 `json:"result,omitempty"`
	SerializationVersion   *string        `json:"serializationVersion,omitempty"`
	SourceMicroversion     *string        `json:"sourceMicroversion,omitempty"`
}

// NewBTFeatureScriptEvalResponse1859 instantiates a new BTFeatureScriptEvalResponse1859 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTFeatureScriptEvalResponse1859() *BTFeatureScriptEvalResponse1859 {
	this := BTFeatureScriptEvalResponse1859{}
	return &this
}

// NewBTFeatureScriptEvalResponse1859WithDefaults instantiates a new BTFeatureScriptEvalResponse1859 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTFeatureScriptEvalResponse1859WithDefaults() *BTFeatureScriptEvalResponse1859 {
	this := BTFeatureScriptEvalResponse1859{}
	return &this
}

// GetConsole returns the Console field value if set, zero value otherwise.
func (o *BTFeatureScriptEvalResponse1859) GetConsole() string {
	if o == nil || o.Console == nil {
		var ret string
		return ret
	}
	return *o.Console
}

// GetConsoleOk returns a tuple with the Console field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureScriptEvalResponse1859) GetConsoleOk() (*string, bool) {
	if o == nil || o.Console == nil {
		return nil, false
	}
	return o.Console, true
}

// HasConsole returns a boolean if a field has been set.
func (o *BTFeatureScriptEvalResponse1859) HasConsole() bool {
	if o != nil && o.Console != nil {
		return true
	}

	return false
}

// SetConsole gets a reference to the given string and assigns it to the Console field.
func (o *BTFeatureScriptEvalResponse1859) SetConsole(v string) {
	o.Console = &v
}

// GetLibraryVersion returns the LibraryVersion field value if set, zero value otherwise.
func (o *BTFeatureScriptEvalResponse1859) GetLibraryVersion() int32 {
	if o == nil || o.LibraryVersion == nil {
		var ret int32
		return ret
	}
	return *o.LibraryVersion
}

// GetLibraryVersionOk returns a tuple with the LibraryVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureScriptEvalResponse1859) GetLibraryVersionOk() (*int32, bool) {
	if o == nil || o.LibraryVersion == nil {
		return nil, false
	}
	return o.LibraryVersion, true
}

// HasLibraryVersion returns a boolean if a field has been set.
func (o *BTFeatureScriptEvalResponse1859) HasLibraryVersion() bool {
	if o != nil && o.LibraryVersion != nil {
		return true
	}

	return false
}

// SetLibraryVersion gets a reference to the given int32 and assigns it to the LibraryVersion field.
func (o *BTFeatureScriptEvalResponse1859) SetLibraryVersion(v int32) {
	o.LibraryVersion = &v
}

// GetMicroversionSkew returns the MicroversionSkew field value if set, zero value otherwise.
func (o *BTFeatureScriptEvalResponse1859) GetMicroversionSkew() bool {
	if o == nil || o.MicroversionSkew == nil {
		var ret bool
		return ret
	}
	return *o.MicroversionSkew
}

// GetMicroversionSkewOk returns a tuple with the MicroversionSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureScriptEvalResponse1859) GetMicroversionSkewOk() (*bool, bool) {
	if o == nil || o.MicroversionSkew == nil {
		return nil, false
	}
	return o.MicroversionSkew, true
}

// HasMicroversionSkew returns a boolean if a field has been set.
func (o *BTFeatureScriptEvalResponse1859) HasMicroversionSkew() bool {
	if o != nil && o.MicroversionSkew != nil {
		return true
	}

	return false
}

// SetMicroversionSkew gets a reference to the given bool and assigns it to the MicroversionSkew field.
func (o *BTFeatureScriptEvalResponse1859) SetMicroversionSkew(v bool) {
	o.MicroversionSkew = &v
}

// GetNotices returns the Notices field value if set, zero value otherwise.
func (o *BTFeatureScriptEvalResponse1859) GetNotices() []BTNotice227 {
	if o == nil || o.Notices == nil {
		var ret []BTNotice227
		return ret
	}
	return o.Notices
}

// GetNoticesOk returns a tuple with the Notices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureScriptEvalResponse1859) GetNoticesOk() ([]BTNotice227, bool) {
	if o == nil || o.Notices == nil {
		return nil, false
	}
	return o.Notices, true
}

// HasNotices returns a boolean if a field has been set.
func (o *BTFeatureScriptEvalResponse1859) HasNotices() bool {
	if o != nil && o.Notices != nil {
		return true
	}

	return false
}

// SetNotices gets a reference to the given []BTNotice227 and assigns it to the Notices field.
func (o *BTFeatureScriptEvalResponse1859) SetNotices(v []BTNotice227) {
	o.Notices = v
}

// GetRejectMicroversionSkew returns the RejectMicroversionSkew field value if set, zero value otherwise.
func (o *BTFeatureScriptEvalResponse1859) GetRejectMicroversionSkew() bool {
	if o == nil || o.RejectMicroversionSkew == nil {
		var ret bool
		return ret
	}
	return *o.RejectMicroversionSkew
}

// GetRejectMicroversionSkewOk returns a tuple with the RejectMicroversionSkew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureScriptEvalResponse1859) GetRejectMicroversionSkewOk() (*bool, bool) {
	if o == nil || o.RejectMicroversionSkew == nil {
		return nil, false
	}
	return o.RejectMicroversionSkew, true
}

// HasRejectMicroversionSkew returns a boolean if a field has been set.
func (o *BTFeatureScriptEvalResponse1859) HasRejectMicroversionSkew() bool {
	if o != nil && o.RejectMicroversionSkew != nil {
		return true
	}

	return false
}

// SetRejectMicroversionSkew gets a reference to the given bool and assigns it to the RejectMicroversionSkew field.
func (o *BTFeatureScriptEvalResponse1859) SetRejectMicroversionSkew(v bool) {
	o.RejectMicroversionSkew = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *BTFeatureScriptEvalResponse1859) GetResult() BTFSValue1888 {
	if o == nil || o.Result == nil {
		var ret BTFSValue1888
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureScriptEvalResponse1859) GetResultOk() (*BTFSValue1888, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *BTFeatureScriptEvalResponse1859) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given BTFSValue1888 and assigns it to the Result field.
func (o *BTFeatureScriptEvalResponse1859) SetResult(v BTFSValue1888) {
	o.Result = &v
}

// GetSerializationVersion returns the SerializationVersion field value if set, zero value otherwise.
func (o *BTFeatureScriptEvalResponse1859) GetSerializationVersion() string {
	if o == nil || o.SerializationVersion == nil {
		var ret string
		return ret
	}
	return *o.SerializationVersion
}

// GetSerializationVersionOk returns a tuple with the SerializationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureScriptEvalResponse1859) GetSerializationVersionOk() (*string, bool) {
	if o == nil || o.SerializationVersion == nil {
		return nil, false
	}
	return o.SerializationVersion, true
}

// HasSerializationVersion returns a boolean if a field has been set.
func (o *BTFeatureScriptEvalResponse1859) HasSerializationVersion() bool {
	if o != nil && o.SerializationVersion != nil {
		return true
	}

	return false
}

// SetSerializationVersion gets a reference to the given string and assigns it to the SerializationVersion field.
func (o *BTFeatureScriptEvalResponse1859) SetSerializationVersion(v string) {
	o.SerializationVersion = &v
}

// GetSourceMicroversion returns the SourceMicroversion field value if set, zero value otherwise.
func (o *BTFeatureScriptEvalResponse1859) GetSourceMicroversion() string {
	if o == nil || o.SourceMicroversion == nil {
		var ret string
		return ret
	}
	return *o.SourceMicroversion
}

// GetSourceMicroversionOk returns a tuple with the SourceMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTFeatureScriptEvalResponse1859) GetSourceMicroversionOk() (*string, bool) {
	if o == nil || o.SourceMicroversion == nil {
		return nil, false
	}
	return o.SourceMicroversion, true
}

// HasSourceMicroversion returns a boolean if a field has been set.
func (o *BTFeatureScriptEvalResponse1859) HasSourceMicroversion() bool {
	if o != nil && o.SourceMicroversion != nil {
		return true
	}

	return false
}

// SetSourceMicroversion gets a reference to the given string and assigns it to the SourceMicroversion field.
func (o *BTFeatureScriptEvalResponse1859) SetSourceMicroversion(v string) {
	o.SourceMicroversion = &v
}

func (o BTFeatureScriptEvalResponse1859) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Console != nil {
		toSerialize["console"] = o.Console
	}
	if o.LibraryVersion != nil {
		toSerialize["libraryVersion"] = o.LibraryVersion
	}
	if o.MicroversionSkew != nil {
		toSerialize["microversionSkew"] = o.MicroversionSkew
	}
	if o.Notices != nil {
		toSerialize["notices"] = o.Notices
	}
	if o.RejectMicroversionSkew != nil {
		toSerialize["rejectMicroversionSkew"] = o.RejectMicroversionSkew
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.SerializationVersion != nil {
		toSerialize["serializationVersion"] = o.SerializationVersion
	}
	if o.SourceMicroversion != nil {
		toSerialize["sourceMicroversion"] = o.SourceMicroversion
	}
	return json.Marshal(toSerialize)
}

type NullableBTFeatureScriptEvalResponse1859 struct {
	value *BTFeatureScriptEvalResponse1859
	isSet bool
}

func (v NullableBTFeatureScriptEvalResponse1859) Get() *BTFeatureScriptEvalResponse1859 {
	return v.value
}

func (v *NullableBTFeatureScriptEvalResponse1859) Set(val *BTFeatureScriptEvalResponse1859) {
	v.value = val
	v.isSet = true
}

func (v NullableBTFeatureScriptEvalResponse1859) IsSet() bool {
	return v.isSet
}

func (v *NullableBTFeatureScriptEvalResponse1859) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTFeatureScriptEvalResponse1859(val *BTFeatureScriptEvalResponse1859) *NullableBTFeatureScriptEvalResponse1859 {
	return &NullableBTFeatureScriptEvalResponse1859{value: val, isSet: true}
}

func (v NullableBTFeatureScriptEvalResponse1859) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTFeatureScriptEvalResponse1859) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
