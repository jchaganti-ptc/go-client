/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.173.27313-c3b730633f3c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// Components struct for Components
type Components struct {
	Callbacks       *map[string]Callback              `json:"callbacks,omitempty"`
	Examples        *map[string]Example               `json:"examples,omitempty"`
	Extensions      map[string]map[string]interface{} `json:"extensions,omitempty"`
	Headers         *map[string]Header                `json:"headers,omitempty"`
	Links           *map[string]Link                  `json:"links,omitempty"`
	Parameters      *map[string]Parameter             `json:"parameters,omitempty"`
	PathItems       *map[string]PathItem              `json:"pathItems,omitempty"`
	RequestBodies   *map[string]RequestBody           `json:"requestBodies,omitempty"`
	Responses       *map[string]ApiResponse           `json:"responses,omitempty"`
	Schemas         *map[string]Schema                `json:"schemas,omitempty"`
	SecuritySchemes *map[string]SecurityScheme        `json:"securitySchemes,omitempty"`
}

// NewComponents instantiates a new Components object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponents() *Components {
	this := Components{}
	return &this
}

// NewComponentsWithDefaults instantiates a new Components object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentsWithDefaults() *Components {
	this := Components{}
	return &this
}

// GetCallbacks returns the Callbacks field value if set, zero value otherwise.
func (o *Components) GetCallbacks() map[string]Callback {
	if o == nil || o.Callbacks == nil {
		var ret map[string]Callback
		return ret
	}
	return *o.Callbacks
}

// GetCallbacksOk returns a tuple with the Callbacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Components) GetCallbacksOk() (*map[string]Callback, bool) {
	if o == nil || o.Callbacks == nil {
		return nil, false
	}
	return o.Callbacks, true
}

// HasCallbacks returns a boolean if a field has been set.
func (o *Components) HasCallbacks() bool {
	if o != nil && o.Callbacks != nil {
		return true
	}

	return false
}

// SetCallbacks gets a reference to the given map[string]Callback and assigns it to the Callbacks field.
func (o *Components) SetCallbacks(v map[string]Callback) {
	o.Callbacks = &v
}

// GetExamples returns the Examples field value if set, zero value otherwise.
func (o *Components) GetExamples() map[string]Example {
	if o == nil || o.Examples == nil {
		var ret map[string]Example
		return ret
	}
	return *o.Examples
}

// GetExamplesOk returns a tuple with the Examples field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Components) GetExamplesOk() (*map[string]Example, bool) {
	if o == nil || o.Examples == nil {
		return nil, false
	}
	return o.Examples, true
}

// HasExamples returns a boolean if a field has been set.
func (o *Components) HasExamples() bool {
	if o != nil && o.Examples != nil {
		return true
	}

	return false
}

// SetExamples gets a reference to the given map[string]Example and assigns it to the Examples field.
func (o *Components) SetExamples(v map[string]Example) {
	o.Examples = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *Components) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Components) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *Components) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *Components) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *Components) GetHeaders() map[string]Header {
	if o == nil || o.Headers == nil {
		var ret map[string]Header
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Components) GetHeadersOk() (*map[string]Header, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *Components) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]Header and assigns it to the Headers field.
func (o *Components) SetHeaders(v map[string]Header) {
	o.Headers = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Components) GetLinks() map[string]Link {
	if o == nil || o.Links == nil {
		var ret map[string]Link
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Components) GetLinksOk() (*map[string]Link, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Components) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]Link and assigns it to the Links field.
func (o *Components) SetLinks(v map[string]Link) {
	o.Links = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *Components) GetParameters() map[string]Parameter {
	if o == nil || o.Parameters == nil {
		var ret map[string]Parameter
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Components) GetParametersOk() (*map[string]Parameter, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *Components) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]Parameter and assigns it to the Parameters field.
func (o *Components) SetParameters(v map[string]Parameter) {
	o.Parameters = &v
}

// GetPathItems returns the PathItems field value if set, zero value otherwise.
func (o *Components) GetPathItems() map[string]PathItem {
	if o == nil || o.PathItems == nil {
		var ret map[string]PathItem
		return ret
	}
	return *o.PathItems
}

// GetPathItemsOk returns a tuple with the PathItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Components) GetPathItemsOk() (*map[string]PathItem, bool) {
	if o == nil || o.PathItems == nil {
		return nil, false
	}
	return o.PathItems, true
}

// HasPathItems returns a boolean if a field has been set.
func (o *Components) HasPathItems() bool {
	if o != nil && o.PathItems != nil {
		return true
	}

	return false
}

// SetPathItems gets a reference to the given map[string]PathItem and assigns it to the PathItems field.
func (o *Components) SetPathItems(v map[string]PathItem) {
	o.PathItems = &v
}

// GetRequestBodies returns the RequestBodies field value if set, zero value otherwise.
func (o *Components) GetRequestBodies() map[string]RequestBody {
	if o == nil || o.RequestBodies == nil {
		var ret map[string]RequestBody
		return ret
	}
	return *o.RequestBodies
}

// GetRequestBodiesOk returns a tuple with the RequestBodies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Components) GetRequestBodiesOk() (*map[string]RequestBody, bool) {
	if o == nil || o.RequestBodies == nil {
		return nil, false
	}
	return o.RequestBodies, true
}

// HasRequestBodies returns a boolean if a field has been set.
func (o *Components) HasRequestBodies() bool {
	if o != nil && o.RequestBodies != nil {
		return true
	}

	return false
}

// SetRequestBodies gets a reference to the given map[string]RequestBody and assigns it to the RequestBodies field.
func (o *Components) SetRequestBodies(v map[string]RequestBody) {
	o.RequestBodies = &v
}

// GetResponses returns the Responses field value if set, zero value otherwise.
func (o *Components) GetResponses() map[string]ApiResponse {
	if o == nil || o.Responses == nil {
		var ret map[string]ApiResponse
		return ret
	}
	return *o.Responses
}

// GetResponsesOk returns a tuple with the Responses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Components) GetResponsesOk() (*map[string]ApiResponse, bool) {
	if o == nil || o.Responses == nil {
		return nil, false
	}
	return o.Responses, true
}

// HasResponses returns a boolean if a field has been set.
func (o *Components) HasResponses() bool {
	if o != nil && o.Responses != nil {
		return true
	}

	return false
}

// SetResponses gets a reference to the given map[string]ApiResponse and assigns it to the Responses field.
func (o *Components) SetResponses(v map[string]ApiResponse) {
	o.Responses = &v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *Components) GetSchemas() map[string]Schema {
	if o == nil || o.Schemas == nil {
		var ret map[string]Schema
		return ret
	}
	return *o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Components) GetSchemasOk() (*map[string]Schema, bool) {
	if o == nil || o.Schemas == nil {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *Components) HasSchemas() bool {
	if o != nil && o.Schemas != nil {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given map[string]Schema and assigns it to the Schemas field.
func (o *Components) SetSchemas(v map[string]Schema) {
	o.Schemas = &v
}

// GetSecuritySchemes returns the SecuritySchemes field value if set, zero value otherwise.
func (o *Components) GetSecuritySchemes() map[string]SecurityScheme {
	if o == nil || o.SecuritySchemes == nil {
		var ret map[string]SecurityScheme
		return ret
	}
	return *o.SecuritySchemes
}

// GetSecuritySchemesOk returns a tuple with the SecuritySchemes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Components) GetSecuritySchemesOk() (*map[string]SecurityScheme, bool) {
	if o == nil || o.SecuritySchemes == nil {
		return nil, false
	}
	return o.SecuritySchemes, true
}

// HasSecuritySchemes returns a boolean if a field has been set.
func (o *Components) HasSecuritySchemes() bool {
	if o != nil && o.SecuritySchemes != nil {
		return true
	}

	return false
}

// SetSecuritySchemes gets a reference to the given map[string]SecurityScheme and assigns it to the SecuritySchemes field.
func (o *Components) SetSecuritySchemes(v map[string]SecurityScheme) {
	o.SecuritySchemes = &v
}

func (o Components) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Callbacks != nil {
		toSerialize["callbacks"] = o.Callbacks
	}
	if o.Examples != nil {
		toSerialize["examples"] = o.Examples
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.PathItems != nil {
		toSerialize["pathItems"] = o.PathItems
	}
	if o.RequestBodies != nil {
		toSerialize["requestBodies"] = o.RequestBodies
	}
	if o.Responses != nil {
		toSerialize["responses"] = o.Responses
	}
	if o.Schemas != nil {
		toSerialize["schemas"] = o.Schemas
	}
	if o.SecuritySchemes != nil {
		toSerialize["securitySchemes"] = o.SecuritySchemes
	}
	return json.Marshal(toSerialize)
}

type NullableComponents struct {
	value *Components
	isSet bool
}

func (v NullableComponents) Get() *Components {
	return v.value
}

func (v *NullableComponents) Set(val *Components) {
	v.value = val
	v.isSet = true
}

func (v NullableComponents) IsSet() bool {
	return v.isSet
}

func (v *NullableComponents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponents(val *Components) *NullableComponents {
	return &NullableComponents{value: val, isSet: true}
}

func (v NullableComponents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
