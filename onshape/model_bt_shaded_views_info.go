/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6489-39ccb1a53c2e
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTShadedViewsInfo struct for BTShadedViewsInfo
type BTShadedViewsInfo struct {
	Images [][]string `json:"images,omitempty"`
}

// NewBTShadedViewsInfo instantiates a new BTShadedViewsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTShadedViewsInfo() *BTShadedViewsInfo {
	this := BTShadedViewsInfo{}
	return &this
}

// NewBTShadedViewsInfoWithDefaults instantiates a new BTShadedViewsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTShadedViewsInfoWithDefaults() *BTShadedViewsInfo {
	this := BTShadedViewsInfo{}
	return &this
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *BTShadedViewsInfo) GetImages() [][]string {
	if o == nil || o.Images == nil {
		var ret [][]string
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTShadedViewsInfo) GetImagesOk() ([][]string, bool) {
	if o == nil || o.Images == nil {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *BTShadedViewsInfo) HasImages() bool {
	if o != nil && o.Images != nil {
		return true
	}

	return false
}

// SetImages gets a reference to the given [][]string and assigns it to the Images field.
func (o *BTShadedViewsInfo) SetImages(v [][]string) {
	o.Images = v
}

func (o BTShadedViewsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Images != nil {
		toSerialize["images"] = o.Images
	}
	return json.Marshal(toSerialize)
}

type NullableBTShadedViewsInfo struct {
	value *BTShadedViewsInfo
	isSet bool
}

func (v NullableBTShadedViewsInfo) Get() *BTShadedViewsInfo {
	return v.value
}

func (v *NullableBTShadedViewsInfo) Set(val *BTShadedViewsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTShadedViewsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTShadedViewsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTShadedViewsInfo(val *BTShadedViewsInfo) *NullableBTShadedViewsInfo {
	return &NullableBTShadedViewsInfo{value: val, isSet: true}
}

func (v NullableBTShadedViewsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTShadedViewsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
