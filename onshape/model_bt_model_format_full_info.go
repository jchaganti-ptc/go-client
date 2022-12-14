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

// BTModelFormatFullInfo struct for BTModelFormatFullInfo
type BTModelFormatFullInfo struct {
	CouldBeAssembly        *bool   `json:"couldBeAssembly,omitempty"`
	Name                   *string `json:"name,omitempty"`
	TranslatorName         *string `json:"translatorName,omitempty"`
	ValidDestinationFormat *bool   `json:"validDestinationFormat,omitempty"`
	ValidSourceFormat      *bool   `json:"validSourceFormat,omitempty"`
}

// NewBTModelFormatFullInfo instantiates a new BTModelFormatFullInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTModelFormatFullInfo() *BTModelFormatFullInfo {
	this := BTModelFormatFullInfo{}
	return &this
}

// NewBTModelFormatFullInfoWithDefaults instantiates a new BTModelFormatFullInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTModelFormatFullInfoWithDefaults() *BTModelFormatFullInfo {
	this := BTModelFormatFullInfo{}
	return &this
}

// GetCouldBeAssembly returns the CouldBeAssembly field value if set, zero value otherwise.
func (o *BTModelFormatFullInfo) GetCouldBeAssembly() bool {
	if o == nil || o.CouldBeAssembly == nil {
		var ret bool
		return ret
	}
	return *o.CouldBeAssembly
}

// GetCouldBeAssemblyOk returns a tuple with the CouldBeAssembly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTModelFormatFullInfo) GetCouldBeAssemblyOk() (*bool, bool) {
	if o == nil || o.CouldBeAssembly == nil {
		return nil, false
	}
	return o.CouldBeAssembly, true
}

// HasCouldBeAssembly returns a boolean if a field has been set.
func (o *BTModelFormatFullInfo) HasCouldBeAssembly() bool {
	if o != nil && o.CouldBeAssembly != nil {
		return true
	}

	return false
}

// SetCouldBeAssembly gets a reference to the given bool and assigns it to the CouldBeAssembly field.
func (o *BTModelFormatFullInfo) SetCouldBeAssembly(v bool) {
	o.CouldBeAssembly = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTModelFormatFullInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTModelFormatFullInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTModelFormatFullInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTModelFormatFullInfo) SetName(v string) {
	o.Name = &v
}

// GetTranslatorName returns the TranslatorName field value if set, zero value otherwise.
func (o *BTModelFormatFullInfo) GetTranslatorName() string {
	if o == nil || o.TranslatorName == nil {
		var ret string
		return ret
	}
	return *o.TranslatorName
}

// GetTranslatorNameOk returns a tuple with the TranslatorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTModelFormatFullInfo) GetTranslatorNameOk() (*string, bool) {
	if o == nil || o.TranslatorName == nil {
		return nil, false
	}
	return o.TranslatorName, true
}

// HasTranslatorName returns a boolean if a field has been set.
func (o *BTModelFormatFullInfo) HasTranslatorName() bool {
	if o != nil && o.TranslatorName != nil {
		return true
	}

	return false
}

// SetTranslatorName gets a reference to the given string and assigns it to the TranslatorName field.
func (o *BTModelFormatFullInfo) SetTranslatorName(v string) {
	o.TranslatorName = &v
}

// GetValidDestinationFormat returns the ValidDestinationFormat field value if set, zero value otherwise.
func (o *BTModelFormatFullInfo) GetValidDestinationFormat() bool {
	if o == nil || o.ValidDestinationFormat == nil {
		var ret bool
		return ret
	}
	return *o.ValidDestinationFormat
}

// GetValidDestinationFormatOk returns a tuple with the ValidDestinationFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTModelFormatFullInfo) GetValidDestinationFormatOk() (*bool, bool) {
	if o == nil || o.ValidDestinationFormat == nil {
		return nil, false
	}
	return o.ValidDestinationFormat, true
}

// HasValidDestinationFormat returns a boolean if a field has been set.
func (o *BTModelFormatFullInfo) HasValidDestinationFormat() bool {
	if o != nil && o.ValidDestinationFormat != nil {
		return true
	}

	return false
}

// SetValidDestinationFormat gets a reference to the given bool and assigns it to the ValidDestinationFormat field.
func (o *BTModelFormatFullInfo) SetValidDestinationFormat(v bool) {
	o.ValidDestinationFormat = &v
}

// GetValidSourceFormat returns the ValidSourceFormat field value if set, zero value otherwise.
func (o *BTModelFormatFullInfo) GetValidSourceFormat() bool {
	if o == nil || o.ValidSourceFormat == nil {
		var ret bool
		return ret
	}
	return *o.ValidSourceFormat
}

// GetValidSourceFormatOk returns a tuple with the ValidSourceFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTModelFormatFullInfo) GetValidSourceFormatOk() (*bool, bool) {
	if o == nil || o.ValidSourceFormat == nil {
		return nil, false
	}
	return o.ValidSourceFormat, true
}

// HasValidSourceFormat returns a boolean if a field has been set.
func (o *BTModelFormatFullInfo) HasValidSourceFormat() bool {
	if o != nil && o.ValidSourceFormat != nil {
		return true
	}

	return false
}

// SetValidSourceFormat gets a reference to the given bool and assigns it to the ValidSourceFormat field.
func (o *BTModelFormatFullInfo) SetValidSourceFormat(v bool) {
	o.ValidSourceFormat = &v
}

func (o BTModelFormatFullInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CouldBeAssembly != nil {
		toSerialize["couldBeAssembly"] = o.CouldBeAssembly
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.TranslatorName != nil {
		toSerialize["translatorName"] = o.TranslatorName
	}
	if o.ValidDestinationFormat != nil {
		toSerialize["validDestinationFormat"] = o.ValidDestinationFormat
	}
	if o.ValidSourceFormat != nil {
		toSerialize["validSourceFormat"] = o.ValidSourceFormat
	}
	return json.Marshal(toSerialize)
}

type NullableBTModelFormatFullInfo struct {
	value *BTModelFormatFullInfo
	isSet bool
}

func (v NullableBTModelFormatFullInfo) Get() *BTModelFormatFullInfo {
	return v.value
}

func (v *NullableBTModelFormatFullInfo) Set(val *BTModelFormatFullInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTModelFormatFullInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTModelFormatFullInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTModelFormatFullInfo(val *BTModelFormatFullInfo) *NullableBTModelFormatFullInfo {
	return &NullableBTModelFormatFullInfo{value: val, isSet: true}
}

func (v NullableBTModelFormatFullInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTModelFormatFullInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
