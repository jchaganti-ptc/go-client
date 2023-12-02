/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.173.26882-0482adeaa8aa
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// Asset struct for Asset
type Asset struct {
	Copyright  *string                           `json:"copyright,omitempty"`
	Extensions map[string]map[string]interface{} `json:"extensions,omitempty"`
	Extras     *map[string]interface{}           `json:"extras,omitempty"`
	Generator  *string                           `json:"generator,omitempty"`
	MinVersion *string                           `json:"minVersion,omitempty"`
	Version    *string                           `json:"version,omitempty"`
}

// NewAsset instantiates a new Asset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsset() *Asset {
	this := Asset{}
	return &this
}

// NewAssetWithDefaults instantiates a new Asset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWithDefaults() *Asset {
	this := Asset{}
	return &this
}

// GetCopyright returns the Copyright field value if set, zero value otherwise.
func (o *Asset) GetCopyright() string {
	if o == nil || o.Copyright == nil {
		var ret string
		return ret
	}
	return *o.Copyright
}

// GetCopyrightOk returns a tuple with the Copyright field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetCopyrightOk() (*string, bool) {
	if o == nil || o.Copyright == nil {
		return nil, false
	}
	return o.Copyright, true
}

// HasCopyright returns a boolean if a field has been set.
func (o *Asset) HasCopyright() bool {
	if o != nil && o.Copyright != nil {
		return true
	}

	return false
}

// SetCopyright gets a reference to the given string and assigns it to the Copyright field.
func (o *Asset) SetCopyright(v string) {
	o.Copyright = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *Asset) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *Asset) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *Asset) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *Asset) GetExtras() map[string]interface{} {
	if o == nil || o.Extras == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetExtrasOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *Asset) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *Asset) SetExtras(v map[string]interface{}) {
	o.Extras = &v
}

// GetGenerator returns the Generator field value if set, zero value otherwise.
func (o *Asset) GetGenerator() string {
	if o == nil || o.Generator == nil {
		var ret string
		return ret
	}
	return *o.Generator
}

// GetGeneratorOk returns a tuple with the Generator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetGeneratorOk() (*string, bool) {
	if o == nil || o.Generator == nil {
		return nil, false
	}
	return o.Generator, true
}

// HasGenerator returns a boolean if a field has been set.
func (o *Asset) HasGenerator() bool {
	if o != nil && o.Generator != nil {
		return true
	}

	return false
}

// SetGenerator gets a reference to the given string and assigns it to the Generator field.
func (o *Asset) SetGenerator(v string) {
	o.Generator = &v
}

// GetMinVersion returns the MinVersion field value if set, zero value otherwise.
func (o *Asset) GetMinVersion() string {
	if o == nil || o.MinVersion == nil {
		var ret string
		return ret
	}
	return *o.MinVersion
}

// GetMinVersionOk returns a tuple with the MinVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetMinVersionOk() (*string, bool) {
	if o == nil || o.MinVersion == nil {
		return nil, false
	}
	return o.MinVersion, true
}

// HasMinVersion returns a boolean if a field has been set.
func (o *Asset) HasMinVersion() bool {
	if o != nil && o.MinVersion != nil {
		return true
	}

	return false
}

// SetMinVersion gets a reference to the given string and assigns it to the MinVersion field.
func (o *Asset) SetMinVersion(v string) {
	o.MinVersion = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Asset) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Asset) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Asset) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Asset) SetVersion(v string) {
	o.Version = &v
}

func (o Asset) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Copyright != nil {
		toSerialize["copyright"] = o.Copyright
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	if o.Generator != nil {
		toSerialize["generator"] = o.Generator
	}
	if o.MinVersion != nil {
		toSerialize["minVersion"] = o.MinVersion
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableAsset struct {
	value *Asset
	isSet bool
}

func (v NullableAsset) Get() *Asset {
	return v.value
}

func (v *NullableAsset) Set(val *Asset) {
	v.value = val
	v.isSet = true
}

func (v NullableAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsset(val *Asset) *NullableAsset {
	return &NullableAsset{value: val, isSet: true}
}

func (v NullableAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
