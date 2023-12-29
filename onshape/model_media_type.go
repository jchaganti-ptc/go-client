/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.28126-700645382199
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// MediaType struct for MediaType
type MediaType struct {
	Encoding       *map[string]Encoding              `json:"encoding,omitempty"`
	Example        *map[string]interface{}           `json:"example,omitempty"`
	ExampleSetFlag *bool                             `json:"exampleSetFlag,omitempty"`
	Examples       *map[string]Example               `json:"examples,omitempty"`
	Extensions     map[string]map[string]interface{} `json:"extensions,omitempty"`
	Schema         *Schema                           `json:"schema,omitempty"`
}

// NewMediaType instantiates a new MediaType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMediaType() *MediaType {
	this := MediaType{}
	return &this
}

// NewMediaTypeWithDefaults instantiates a new MediaType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMediaTypeWithDefaults() *MediaType {
	this := MediaType{}
	return &this
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *MediaType) GetEncoding() map[string]Encoding {
	if o == nil || o.Encoding == nil {
		var ret map[string]Encoding
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaType) GetEncodingOk() (*map[string]Encoding, bool) {
	if o == nil || o.Encoding == nil {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *MediaType) HasEncoding() bool {
	if o != nil && o.Encoding != nil {
		return true
	}

	return false
}

// SetEncoding gets a reference to the given map[string]Encoding and assigns it to the Encoding field.
func (o *MediaType) SetEncoding(v map[string]Encoding) {
	o.Encoding = &v
}

// GetExample returns the Example field value if set, zero value otherwise.
func (o *MediaType) GetExample() map[string]interface{} {
	if o == nil || o.Example == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Example
}

// GetExampleOk returns a tuple with the Example field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaType) GetExampleOk() (*map[string]interface{}, bool) {
	if o == nil || o.Example == nil {
		return nil, false
	}
	return o.Example, true
}

// HasExample returns a boolean if a field has been set.
func (o *MediaType) HasExample() bool {
	if o != nil && o.Example != nil {
		return true
	}

	return false
}

// SetExample gets a reference to the given map[string]interface{} and assigns it to the Example field.
func (o *MediaType) SetExample(v map[string]interface{}) {
	o.Example = &v
}

// GetExampleSetFlag returns the ExampleSetFlag field value if set, zero value otherwise.
func (o *MediaType) GetExampleSetFlag() bool {
	if o == nil || o.ExampleSetFlag == nil {
		var ret bool
		return ret
	}
	return *o.ExampleSetFlag
}

// GetExampleSetFlagOk returns a tuple with the ExampleSetFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaType) GetExampleSetFlagOk() (*bool, bool) {
	if o == nil || o.ExampleSetFlag == nil {
		return nil, false
	}
	return o.ExampleSetFlag, true
}

// HasExampleSetFlag returns a boolean if a field has been set.
func (o *MediaType) HasExampleSetFlag() bool {
	if o != nil && o.ExampleSetFlag != nil {
		return true
	}

	return false
}

// SetExampleSetFlag gets a reference to the given bool and assigns it to the ExampleSetFlag field.
func (o *MediaType) SetExampleSetFlag(v bool) {
	o.ExampleSetFlag = &v
}

// GetExamples returns the Examples field value if set, zero value otherwise.
func (o *MediaType) GetExamples() map[string]Example {
	if o == nil || o.Examples == nil {
		var ret map[string]Example
		return ret
	}
	return *o.Examples
}

// GetExamplesOk returns a tuple with the Examples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaType) GetExamplesOk() (*map[string]Example, bool) {
	if o == nil || o.Examples == nil {
		return nil, false
	}
	return o.Examples, true
}

// HasExamples returns a boolean if a field has been set.
func (o *MediaType) HasExamples() bool {
	if o != nil && o.Examples != nil {
		return true
	}

	return false
}

// SetExamples gets a reference to the given map[string]Example and assigns it to the Examples field.
func (o *MediaType) SetExamples(v map[string]Example) {
	o.Examples = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *MediaType) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaType) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *MediaType) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *MediaType) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *MediaType) GetSchema() Schema {
	if o == nil || o.Schema == nil {
		var ret Schema
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MediaType) GetSchemaOk() (*Schema, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *MediaType) HasSchema() bool {
	if o != nil && o.Schema != nil {
		return true
	}

	return false
}

// SetSchema gets a reference to the given Schema and assigns it to the Schema field.
func (o *MediaType) SetSchema(v Schema) {
	o.Schema = &v
}

func (o MediaType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Encoding != nil {
		toSerialize["encoding"] = o.Encoding
	}
	if o.Example != nil {
		toSerialize["example"] = o.Example
	}
	if o.ExampleSetFlag != nil {
		toSerialize["exampleSetFlag"] = o.ExampleSetFlag
	}
	if o.Examples != nil {
		toSerialize["examples"] = o.Examples
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	return json.Marshal(toSerialize)
}

type NullableMediaType struct {
	value *MediaType
	isSet bool
}

func (v NullableMediaType) Get() *MediaType {
	return v.value
}

func (v *NullableMediaType) Set(val *MediaType) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaType) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaType(val *MediaType) *NullableMediaType {
	return &NullableMediaType{value: val, isSet: true}
}

func (v NullableMediaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
