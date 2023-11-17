/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.172.26043-b28d7068bd76
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// Encoding struct for Encoding
type Encoding struct {
	AllowReserved *bool                             `json:"allowReserved,omitempty"`
	ContentType   *string                           `json:"contentType,omitempty"`
	Explode       *bool                             `json:"explode,omitempty"`
	Extensions    map[string]map[string]interface{} `json:"extensions,omitempty"`
	Headers       *map[string]Header                `json:"headers,omitempty"`
	Style         *StyleEnum                        `json:"style,omitempty"`
}

// NewEncoding instantiates a new Encoding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncoding() *Encoding {
	this := Encoding{}
	return &this
}

// NewEncodingWithDefaults instantiates a new Encoding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncodingWithDefaults() *Encoding {
	this := Encoding{}
	return &this
}

// GetAllowReserved returns the AllowReserved field value if set, zero value otherwise.
func (o *Encoding) GetAllowReserved() bool {
	if o == nil || o.AllowReserved == nil {
		var ret bool
		return ret
	}
	return *o.AllowReserved
}

// GetAllowReservedOk returns a tuple with the AllowReserved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encoding) GetAllowReservedOk() (*bool, bool) {
	if o == nil || o.AllowReserved == nil {
		return nil, false
	}
	return o.AllowReserved, true
}

// HasAllowReserved returns a boolean if a field has been set.
func (o *Encoding) HasAllowReserved() bool {
	if o != nil && o.AllowReserved != nil {
		return true
	}

	return false
}

// SetAllowReserved gets a reference to the given bool and assigns it to the AllowReserved field.
func (o *Encoding) SetAllowReserved(v bool) {
	o.AllowReserved = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *Encoding) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encoding) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *Encoding) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *Encoding) SetContentType(v string) {
	o.ContentType = &v
}

// GetExplode returns the Explode field value if set, zero value otherwise.
func (o *Encoding) GetExplode() bool {
	if o == nil || o.Explode == nil {
		var ret bool
		return ret
	}
	return *o.Explode
}

// GetExplodeOk returns a tuple with the Explode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encoding) GetExplodeOk() (*bool, bool) {
	if o == nil || o.Explode == nil {
		return nil, false
	}
	return o.Explode, true
}

// HasExplode returns a boolean if a field has been set.
func (o *Encoding) HasExplode() bool {
	if o != nil && o.Explode != nil {
		return true
	}

	return false
}

// SetExplode gets a reference to the given bool and assigns it to the Explode field.
func (o *Encoding) SetExplode(v bool) {
	o.Explode = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *Encoding) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encoding) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *Encoding) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *Encoding) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *Encoding) GetHeaders() map[string]Header {
	if o == nil || o.Headers == nil {
		var ret map[string]Header
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encoding) GetHeadersOk() (*map[string]Header, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *Encoding) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]Header and assigns it to the Headers field.
func (o *Encoding) SetHeaders(v map[string]Header) {
	o.Headers = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *Encoding) GetStyle() StyleEnum {
	if o == nil || o.Style == nil {
		var ret StyleEnum
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Encoding) GetStyleOk() (*StyleEnum, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *Encoding) HasStyle() bool {
	if o != nil && o.Style != nil {
		return true
	}

	return false
}

// SetStyle gets a reference to the given StyleEnum and assigns it to the Style field.
func (o *Encoding) SetStyle(v StyleEnum) {
	o.Style = &v
}

func (o Encoding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowReserved != nil {
		toSerialize["allowReserved"] = o.AllowReserved
	}
	if o.ContentType != nil {
		toSerialize["contentType"] = o.ContentType
	}
	if o.Explode != nil {
		toSerialize["explode"] = o.Explode
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return json.Marshal(toSerialize)
}

type NullableEncoding struct {
	value *Encoding
	isSet bool
}

func (v NullableEncoding) Get() *Encoding {
	return v.value
}

func (v *NullableEncoding) Set(val *Encoding) {
	v.value = val
	v.isSet = true
}

func (v NullableEncoding) IsSet() bool {
	return v.isSet
}

func (v *NullableEncoding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncoding(val *Encoding) *NullableEncoding {
	return &NullableEncoding{value: val, isSet: true}
}

func (v NullableEncoding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncoding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}