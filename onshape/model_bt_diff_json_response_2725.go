/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6246-994505a8b5bf
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTDiffJsonResponse2725 struct for BTDiffJsonResponse2725
type BTDiffJsonResponse2725 struct {
	Change         *BTJEdit3734                 `json:"change,omitempty"`
	Patch          *BTDiffJsonResponse2725Patch `json:"patch,omitempty"`
	SourceChangeId *string                      `json:"sourceChangeId,omitempty"`
	TargetChangeId *string                      `json:"targetChangeId,omitempty"`
}

// NewBTDiffJsonResponse2725 instantiates a new BTDiffJsonResponse2725 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTDiffJsonResponse2725() *BTDiffJsonResponse2725 {
	this := BTDiffJsonResponse2725{}
	return &this
}

// NewBTDiffJsonResponse2725WithDefaults instantiates a new BTDiffJsonResponse2725 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTDiffJsonResponse2725WithDefaults() *BTDiffJsonResponse2725 {
	this := BTDiffJsonResponse2725{}
	return &this
}

// GetChange returns the Change field value if set, zero value otherwise.
func (o *BTDiffJsonResponse2725) GetChange() BTJEdit3734 {
	if o == nil || o.Change == nil {
		var ret BTJEdit3734
		return ret
	}
	return *o.Change
}

// GetChangeOk returns a tuple with the Change field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiffJsonResponse2725) GetChangeOk() (*BTJEdit3734, bool) {
	if o == nil || o.Change == nil {
		return nil, false
	}
	return o.Change, true
}

// HasChange returns a boolean if a field has been set.
func (o *BTDiffJsonResponse2725) HasChange() bool {
	if o != nil && o.Change != nil {
		return true
	}

	return false
}

// SetChange gets a reference to the given BTJEdit3734 and assigns it to the Change field.
func (o *BTDiffJsonResponse2725) SetChange(v BTJEdit3734) {
	o.Change = &v
}

// GetPatch returns the Patch field value if set, zero value otherwise.
func (o *BTDiffJsonResponse2725) GetPatch() BTDiffJsonResponse2725Patch {
	if o == nil || o.Patch == nil {
		var ret BTDiffJsonResponse2725Patch
		return ret
	}
	return *o.Patch
}

// GetPatchOk returns a tuple with the Patch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiffJsonResponse2725) GetPatchOk() (*BTDiffJsonResponse2725Patch, bool) {
	if o == nil || o.Patch == nil {
		return nil, false
	}
	return o.Patch, true
}

// HasPatch returns a boolean if a field has been set.
func (o *BTDiffJsonResponse2725) HasPatch() bool {
	if o != nil && o.Patch != nil {
		return true
	}

	return false
}

// SetPatch gets a reference to the given BTDiffJsonResponse2725Patch and assigns it to the Patch field.
func (o *BTDiffJsonResponse2725) SetPatch(v BTDiffJsonResponse2725Patch) {
	o.Patch = &v
}

// GetSourceChangeId returns the SourceChangeId field value if set, zero value otherwise.
func (o *BTDiffJsonResponse2725) GetSourceChangeId() string {
	if o == nil || o.SourceChangeId == nil {
		var ret string
		return ret
	}
	return *o.SourceChangeId
}

// GetSourceChangeIdOk returns a tuple with the SourceChangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiffJsonResponse2725) GetSourceChangeIdOk() (*string, bool) {
	if o == nil || o.SourceChangeId == nil {
		return nil, false
	}
	return o.SourceChangeId, true
}

// HasSourceChangeId returns a boolean if a field has been set.
func (o *BTDiffJsonResponse2725) HasSourceChangeId() bool {
	if o != nil && o.SourceChangeId != nil {
		return true
	}

	return false
}

// SetSourceChangeId gets a reference to the given string and assigns it to the SourceChangeId field.
func (o *BTDiffJsonResponse2725) SetSourceChangeId(v string) {
	o.SourceChangeId = &v
}

// GetTargetChangeId returns the TargetChangeId field value if set, zero value otherwise.
func (o *BTDiffJsonResponse2725) GetTargetChangeId() string {
	if o == nil || o.TargetChangeId == nil {
		var ret string
		return ret
	}
	return *o.TargetChangeId
}

// GetTargetChangeIdOk returns a tuple with the TargetChangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTDiffJsonResponse2725) GetTargetChangeIdOk() (*string, bool) {
	if o == nil || o.TargetChangeId == nil {
		return nil, false
	}
	return o.TargetChangeId, true
}

// HasTargetChangeId returns a boolean if a field has been set.
func (o *BTDiffJsonResponse2725) HasTargetChangeId() bool {
	if o != nil && o.TargetChangeId != nil {
		return true
	}

	return false
}

// SetTargetChangeId gets a reference to the given string and assigns it to the TargetChangeId field.
func (o *BTDiffJsonResponse2725) SetTargetChangeId(v string) {
	o.TargetChangeId = &v
}

func (o BTDiffJsonResponse2725) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Change != nil {
		toSerialize["change"] = o.Change
	}
	if o.Patch != nil {
		toSerialize["patch"] = o.Patch
	}
	if o.SourceChangeId != nil {
		toSerialize["sourceChangeId"] = o.SourceChangeId
	}
	if o.TargetChangeId != nil {
		toSerialize["targetChangeId"] = o.TargetChangeId
	}
	return json.Marshal(toSerialize)
}

type NullableBTDiffJsonResponse2725 struct {
	value *BTDiffJsonResponse2725
	isSet bool
}

func (v NullableBTDiffJsonResponse2725) Get() *BTDiffJsonResponse2725 {
	return v.value
}

func (v *NullableBTDiffJsonResponse2725) Set(val *BTDiffJsonResponse2725) {
	v.value = val
	v.isSet = true
}

func (v NullableBTDiffJsonResponse2725) IsSet() bool {
	return v.isSet
}

func (v *NullableBTDiffJsonResponse2725) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTDiffJsonResponse2725(val *BTDiffJsonResponse2725) *NullableBTDiffJsonResponse2725 {
	return &NullableBTDiffJsonResponse2725{value: val, isSet: true}
}

func (v NullableBTDiffJsonResponse2725) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTDiffJsonResponse2725) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
