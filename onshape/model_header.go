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

// Header struct for Header
type Header struct {
	Deprecated  *bool                             `json:"deprecated,omitempty"`
	Description *string                           `json:"description,omitempty"`
	Example     *map[string]interface{}           `json:"example,omitempty"`
	Examples    *map[string]Example               `json:"examples,omitempty"`
	Explode     *bool                             `json:"explode,omitempty"`
	Extensions  map[string]map[string]interface{} `json:"extensions,omitempty"`
	Getref      *string                           `json:"get$ref,omitempty"`
	Required    *bool                             `json:"required,omitempty"`
	Schema      *Schema                           `json:"schema,omitempty"`
	Style       *StyleEnum                        `json:"style,omitempty"`
}

// NewHeader instantiates a new Header object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHeader() *Header {
	this := Header{}
	return &this
}

// NewHeaderWithDefaults instantiates a new Header object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHeaderWithDefaults() *Header {
	this := Header{}
	return &this
}

// GetDeprecated returns the Deprecated field value if set, zero value otherwise.
func (o *Header) GetDeprecated() bool {
	if o == nil || o.Deprecated == nil {
		var ret bool
		return ret
	}
	return *o.Deprecated
}

// GetDeprecatedOk returns a tuple with the Deprecated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Header) GetDeprecatedOk() (*bool, bool) {
	if o == nil || o.Deprecated == nil {
		return nil, false
	}
	return o.Deprecated, true
}

// HasDeprecated returns a boolean if a field has been set.
func (o *Header) HasDeprecated() bool {
	if o != nil && o.Deprecated != nil {
		return true
	}

	return false
}

// SetDeprecated gets a reference to the given bool and assigns it to the Deprecated field.
func (o *Header) SetDeprecated(v bool) {
	o.Deprecated = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Header) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Header) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Header) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Header) SetDescription(v string) {
	o.Description = &v
}

// GetExample returns the Example field value if set, zero value otherwise.
func (o *Header) GetExample() map[string]interface{} {
	if o == nil || o.Example == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Example
}

// GetExampleOk returns a tuple with the Example field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Header) GetExampleOk() (*map[string]interface{}, bool) {
	if o == nil || o.Example == nil {
		return nil, false
	}
	return o.Example, true
}

// HasExample returns a boolean if a field has been set.
func (o *Header) HasExample() bool {
	if o != nil && o.Example != nil {
		return true
	}

	return false
}

// SetExample gets a reference to the given map[string]interface{} and assigns it to the Example field.
func (o *Header) SetExample(v map[string]interface{}) {
	o.Example = &v
}

// GetExamples returns the Examples field value if set, zero value otherwise.
func (o *Header) GetExamples() map[string]Example {
	if o == nil || o.Examples == nil {
		var ret map[string]Example
		return ret
	}
	return *o.Examples
}

// GetExamplesOk returns a tuple with the Examples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Header) GetExamplesOk() (*map[string]Example, bool) {
	if o == nil || o.Examples == nil {
		return nil, false
	}
	return o.Examples, true
}

// HasExamples returns a boolean if a field has been set.
func (o *Header) HasExamples() bool {
	if o != nil && o.Examples != nil {
		return true
	}

	return false
}

// SetExamples gets a reference to the given map[string]Example and assigns it to the Examples field.
func (o *Header) SetExamples(v map[string]Example) {
	o.Examples = &v
}

// GetExplode returns the Explode field value if set, zero value otherwise.
func (o *Header) GetExplode() bool {
	if o == nil || o.Explode == nil {
		var ret bool
		return ret
	}
	return *o.Explode
}

// GetExplodeOk returns a tuple with the Explode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Header) GetExplodeOk() (*bool, bool) {
	if o == nil || o.Explode == nil {
		return nil, false
	}
	return o.Explode, true
}

// HasExplode returns a boolean if a field has been set.
func (o *Header) HasExplode() bool {
	if o != nil && o.Explode != nil {
		return true
	}

	return false
}

// SetExplode gets a reference to the given bool and assigns it to the Explode field.
func (o *Header) SetExplode(v bool) {
	o.Explode = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *Header) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Header) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *Header) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *Header) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetGetref returns the Getref field value if set, zero value otherwise.
func (o *Header) GetGetref() string {
	if o == nil || o.Getref == nil {
		var ret string
		return ret
	}
	return *o.Getref
}

// GetGetrefOk returns a tuple with the Getref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Header) GetGetrefOk() (*string, bool) {
	if o == nil || o.Getref == nil {
		return nil, false
	}
	return o.Getref, true
}

// HasGetref returns a boolean if a field has been set.
func (o *Header) HasGetref() bool {
	if o != nil && o.Getref != nil {
		return true
	}

	return false
}

// SetGetref gets a reference to the given string and assigns it to the Getref field.
func (o *Header) SetGetref(v string) {
	o.Getref = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *Header) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Header) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *Header) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *Header) SetRequired(v bool) {
	o.Required = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *Header) GetSchema() Schema {
	if o == nil || o.Schema == nil {
		var ret Schema
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Header) GetSchemaOk() (*Schema, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *Header) HasSchema() bool {
	if o != nil && o.Schema != nil {
		return true
	}

	return false
}

// SetSchema gets a reference to the given Schema and assigns it to the Schema field.
func (o *Header) SetSchema(v Schema) {
	o.Schema = &v
}

// GetStyle returns the Style field value if set, zero value otherwise.
func (o *Header) GetStyle() StyleEnum {
	if o == nil || o.Style == nil {
		var ret StyleEnum
		return ret
	}
	return *o.Style
}

// GetStyleOk returns a tuple with the Style field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Header) GetStyleOk() (*StyleEnum, bool) {
	if o == nil || o.Style == nil {
		return nil, false
	}
	return o.Style, true
}

// HasStyle returns a boolean if a field has been set.
func (o *Header) HasStyle() bool {
	if o != nil && o.Style != nil {
		return true
	}

	return false
}

// SetStyle gets a reference to the given StyleEnum and assigns it to the Style field.
func (o *Header) SetStyle(v StyleEnum) {
	o.Style = &v
}

func (o Header) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Deprecated != nil {
		toSerialize["deprecated"] = o.Deprecated
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Example != nil {
		toSerialize["example"] = o.Example
	}
	if o.Examples != nil {
		toSerialize["examples"] = o.Examples
	}
	if o.Explode != nil {
		toSerialize["explode"] = o.Explode
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Getref != nil {
		toSerialize["get$ref"] = o.Getref
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	if o.Style != nil {
		toSerialize["style"] = o.Style
	}
	return json.Marshal(toSerialize)
}

type NullableHeader struct {
	value *Header
	isSet bool
}

func (v NullableHeader) Get() *Header {
	return v.value
}

func (v *NullableHeader) Set(val *Header) {
	v.value = val
	v.isSet = true
}

func (v NullableHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHeader(val *Header) *NullableHeader {
	return &NullableHeader{value: val, isSet: true}
}

func (v NullableHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
