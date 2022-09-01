/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6309-06d9e62c38f0
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPartAppearanceInfo struct for BTPartAppearanceInfo
type BTPartAppearanceInfo struct {
	Color       *BTColorInfo `json:"color,omitempty"`
	IsGenerated *bool        `json:"isGenerated,omitempty"`
	Opacity     *int32       `json:"opacity,omitempty"`
}

// NewBTPartAppearanceInfo instantiates a new BTPartAppearanceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPartAppearanceInfo() *BTPartAppearanceInfo {
	this := BTPartAppearanceInfo{}
	return &this
}

// NewBTPartAppearanceInfoWithDefaults instantiates a new BTPartAppearanceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPartAppearanceInfoWithDefaults() *BTPartAppearanceInfo {
	this := BTPartAppearanceInfo{}
	return &this
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *BTPartAppearanceInfo) GetColor() BTColorInfo {
	if o == nil || o.Color == nil {
		var ret BTColorInfo
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartAppearanceInfo) GetColorOk() (*BTColorInfo, bool) {
	if o == nil || o.Color == nil {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *BTPartAppearanceInfo) HasColor() bool {
	if o != nil && o.Color != nil {
		return true
	}

	return false
}

// SetColor gets a reference to the given BTColorInfo and assigns it to the Color field.
func (o *BTPartAppearanceInfo) SetColor(v BTColorInfo) {
	o.Color = &v
}

// GetIsGenerated returns the IsGenerated field value if set, zero value otherwise.
func (o *BTPartAppearanceInfo) GetIsGenerated() bool {
	if o == nil || o.IsGenerated == nil {
		var ret bool
		return ret
	}
	return *o.IsGenerated
}

// GetIsGeneratedOk returns a tuple with the IsGenerated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartAppearanceInfo) GetIsGeneratedOk() (*bool, bool) {
	if o == nil || o.IsGenerated == nil {
		return nil, false
	}
	return o.IsGenerated, true
}

// HasIsGenerated returns a boolean if a field has been set.
func (o *BTPartAppearanceInfo) HasIsGenerated() bool {
	if o != nil && o.IsGenerated != nil {
		return true
	}

	return false
}

// SetIsGenerated gets a reference to the given bool and assigns it to the IsGenerated field.
func (o *BTPartAppearanceInfo) SetIsGenerated(v bool) {
	o.IsGenerated = &v
}

// GetOpacity returns the Opacity field value if set, zero value otherwise.
func (o *BTPartAppearanceInfo) GetOpacity() int32 {
	if o == nil || o.Opacity == nil {
		var ret int32
		return ret
	}
	return *o.Opacity
}

// GetOpacityOk returns a tuple with the Opacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPartAppearanceInfo) GetOpacityOk() (*int32, bool) {
	if o == nil || o.Opacity == nil {
		return nil, false
	}
	return o.Opacity, true
}

// HasOpacity returns a boolean if a field has been set.
func (o *BTPartAppearanceInfo) HasOpacity() bool {
	if o != nil && o.Opacity != nil {
		return true
	}

	return false
}

// SetOpacity gets a reference to the given int32 and assigns it to the Opacity field.
func (o *BTPartAppearanceInfo) SetOpacity(v int32) {
	o.Opacity = &v
}

func (o BTPartAppearanceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Color != nil {
		toSerialize["color"] = o.Color
	}
	if o.IsGenerated != nil {
		toSerialize["isGenerated"] = o.IsGenerated
	}
	if o.Opacity != nil {
		toSerialize["opacity"] = o.Opacity
	}
	return json.Marshal(toSerialize)
}

type NullableBTPartAppearanceInfo struct {
	value *BTPartAppearanceInfo
	isSet bool
}

func (v NullableBTPartAppearanceInfo) Get() *BTPartAppearanceInfo {
	return v.value
}

func (v *NullableBTPartAppearanceInfo) Set(val *BTPartAppearanceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPartAppearanceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPartAppearanceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPartAppearanceInfo(val *BTPartAppearanceInfo) *NullableBTPartAppearanceInfo {
	return &NullableBTPartAppearanceInfo{value: val, isSet: true}
}

func (v NullableBTPartAppearanceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPartAppearanceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
