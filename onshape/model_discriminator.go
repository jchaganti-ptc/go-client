/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.172.25652-944cf1af37c9
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// Discriminator struct for Discriminator
type Discriminator struct {
	Extensions   map[string]map[string]interface{} `json:"extensions,omitempty"`
	Mapping      *map[string]string                `json:"mapping,omitempty"`
	PropertyName *string                           `json:"propertyName,omitempty"`
}

// NewDiscriminator instantiates a new Discriminator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscriminator() *Discriminator {
	this := Discriminator{}
	return &this
}

// NewDiscriminatorWithDefaults instantiates a new Discriminator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscriminatorWithDefaults() *Discriminator {
	this := Discriminator{}
	return &this
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *Discriminator) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discriminator) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *Discriminator) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *Discriminator) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetMapping returns the Mapping field value if set, zero value otherwise.
func (o *Discriminator) GetMapping() map[string]string {
	if o == nil || o.Mapping == nil {
		var ret map[string]string
		return ret
	}
	return *o.Mapping
}

// GetMappingOk returns a tuple with the Mapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discriminator) GetMappingOk() (*map[string]string, bool) {
	if o == nil || o.Mapping == nil {
		return nil, false
	}
	return o.Mapping, true
}

// HasMapping returns a boolean if a field has been set.
func (o *Discriminator) HasMapping() bool {
	if o != nil && o.Mapping != nil {
		return true
	}

	return false
}

// SetMapping gets a reference to the given map[string]string and assigns it to the Mapping field.
func (o *Discriminator) SetMapping(v map[string]string) {
	o.Mapping = &v
}

// GetPropertyName returns the PropertyName field value if set, zero value otherwise.
func (o *Discriminator) GetPropertyName() string {
	if o == nil || o.PropertyName == nil {
		var ret string
		return ret
	}
	return *o.PropertyName
}

// GetPropertyNameOk returns a tuple with the PropertyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Discriminator) GetPropertyNameOk() (*string, bool) {
	if o == nil || o.PropertyName == nil {
		return nil, false
	}
	return o.PropertyName, true
}

// HasPropertyName returns a boolean if a field has been set.
func (o *Discriminator) HasPropertyName() bool {
	if o != nil && o.PropertyName != nil {
		return true
	}

	return false
}

// SetPropertyName gets a reference to the given string and assigns it to the PropertyName field.
func (o *Discriminator) SetPropertyName(v string) {
	o.PropertyName = &v
}

func (o Discriminator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Mapping != nil {
		toSerialize["mapping"] = o.Mapping
	}
	if o.PropertyName != nil {
		toSerialize["propertyName"] = o.PropertyName
	}
	return json.Marshal(toSerialize)
}

type NullableDiscriminator struct {
	value *Discriminator
	isSet bool
}

func (v NullableDiscriminator) Get() *Discriminator {
	return v.value
}

func (v *NullableDiscriminator) Set(val *Discriminator) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscriminator) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscriminator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscriminator(val *Discriminator) *NullableDiscriminator {
	return &NullableDiscriminator{value: val, isSet: true}
}

func (v NullableDiscriminator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscriminator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
