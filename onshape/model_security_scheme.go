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

// SecurityScheme struct for SecurityScheme
type SecurityScheme struct {
	BearerFormat     *string                           `json:"bearerFormat,omitempty"`
	Description      *string                           `json:"description,omitempty"`
	Extensions       map[string]map[string]interface{} `json:"extensions,omitempty"`
	Flows            *OAuthFlows                       `json:"flows,omitempty"`
	Getref           *string                           `json:"get$ref,omitempty"`
	In               *In                               `json:"in,omitempty"`
	Name             *string                           `json:"name,omitempty"`
	OpenIdConnectUrl *string                           `json:"openIdConnectUrl,omitempty"`
	Scheme           *string                           `json:"scheme,omitempty"`
	Type             *Type                             `json:"type,omitempty"`
}

// NewSecurityScheme instantiates a new SecurityScheme object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityScheme() *SecurityScheme {
	this := SecurityScheme{}
	return &this
}

// NewSecuritySchemeWithDefaults instantiates a new SecurityScheme object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecuritySchemeWithDefaults() *SecurityScheme {
	this := SecurityScheme{}
	return &this
}

// GetBearerFormat returns the BearerFormat field value if set, zero value otherwise.
func (o *SecurityScheme) GetBearerFormat() string {
	if o == nil || o.BearerFormat == nil {
		var ret string
		return ret
	}
	return *o.BearerFormat
}

// GetBearerFormatOk returns a tuple with the BearerFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScheme) GetBearerFormatOk() (*string, bool) {
	if o == nil || o.BearerFormat == nil {
		return nil, false
	}
	return o.BearerFormat, true
}

// HasBearerFormat returns a boolean if a field has been set.
func (o *SecurityScheme) HasBearerFormat() bool {
	if o != nil && o.BearerFormat != nil {
		return true
	}

	return false
}

// SetBearerFormat gets a reference to the given string and assigns it to the BearerFormat field.
func (o *SecurityScheme) SetBearerFormat(v string) {
	o.BearerFormat = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SecurityScheme) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScheme) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SecurityScheme) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SecurityScheme) SetDescription(v string) {
	o.Description = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *SecurityScheme) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScheme) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *SecurityScheme) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *SecurityScheme) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetFlows returns the Flows field value if set, zero value otherwise.
func (o *SecurityScheme) GetFlows() OAuthFlows {
	if o == nil || o.Flows == nil {
		var ret OAuthFlows
		return ret
	}
	return *o.Flows
}

// GetFlowsOk returns a tuple with the Flows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScheme) GetFlowsOk() (*OAuthFlows, bool) {
	if o == nil || o.Flows == nil {
		return nil, false
	}
	return o.Flows, true
}

// HasFlows returns a boolean if a field has been set.
func (o *SecurityScheme) HasFlows() bool {
	if o != nil && o.Flows != nil {
		return true
	}

	return false
}

// SetFlows gets a reference to the given OAuthFlows and assigns it to the Flows field.
func (o *SecurityScheme) SetFlows(v OAuthFlows) {
	o.Flows = &v
}

// GetGetref returns the Getref field value if set, zero value otherwise.
func (o *SecurityScheme) GetGetref() string {
	if o == nil || o.Getref == nil {
		var ret string
		return ret
	}
	return *o.Getref
}

// GetGetrefOk returns a tuple with the Getref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScheme) GetGetrefOk() (*string, bool) {
	if o == nil || o.Getref == nil {
		return nil, false
	}
	return o.Getref, true
}

// HasGetref returns a boolean if a field has been set.
func (o *SecurityScheme) HasGetref() bool {
	if o != nil && o.Getref != nil {
		return true
	}

	return false
}

// SetGetref gets a reference to the given string and assigns it to the Getref field.
func (o *SecurityScheme) SetGetref(v string) {
	o.Getref = &v
}

// GetIn returns the In field value if set, zero value otherwise.
func (o *SecurityScheme) GetIn() In {
	if o == nil || o.In == nil {
		var ret In
		return ret
	}
	return *o.In
}

// GetInOk returns a tuple with the In field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScheme) GetInOk() (*In, bool) {
	if o == nil || o.In == nil {
		return nil, false
	}
	return o.In, true
}

// HasIn returns a boolean if a field has been set.
func (o *SecurityScheme) HasIn() bool {
	if o != nil && o.In != nil {
		return true
	}

	return false
}

// SetIn gets a reference to the given In and assigns it to the In field.
func (o *SecurityScheme) SetIn(v In) {
	o.In = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SecurityScheme) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScheme) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SecurityScheme) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SecurityScheme) SetName(v string) {
	o.Name = &v
}

// GetOpenIdConnectUrl returns the OpenIdConnectUrl field value if set, zero value otherwise.
func (o *SecurityScheme) GetOpenIdConnectUrl() string {
	if o == nil || o.OpenIdConnectUrl == nil {
		var ret string
		return ret
	}
	return *o.OpenIdConnectUrl
}

// GetOpenIdConnectUrlOk returns a tuple with the OpenIdConnectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScheme) GetOpenIdConnectUrlOk() (*string, bool) {
	if o == nil || o.OpenIdConnectUrl == nil {
		return nil, false
	}
	return o.OpenIdConnectUrl, true
}

// HasOpenIdConnectUrl returns a boolean if a field has been set.
func (o *SecurityScheme) HasOpenIdConnectUrl() bool {
	if o != nil && o.OpenIdConnectUrl != nil {
		return true
	}

	return false
}

// SetOpenIdConnectUrl gets a reference to the given string and assigns it to the OpenIdConnectUrl field.
func (o *SecurityScheme) SetOpenIdConnectUrl(v string) {
	o.OpenIdConnectUrl = &v
}

// GetScheme returns the Scheme field value if set, zero value otherwise.
func (o *SecurityScheme) GetScheme() string {
	if o == nil || o.Scheme == nil {
		var ret string
		return ret
	}
	return *o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScheme) GetSchemeOk() (*string, bool) {
	if o == nil || o.Scheme == nil {
		return nil, false
	}
	return o.Scheme, true
}

// HasScheme returns a boolean if a field has been set.
func (o *SecurityScheme) HasScheme() bool {
	if o != nil && o.Scheme != nil {
		return true
	}

	return false
}

// SetScheme gets a reference to the given string and assigns it to the Scheme field.
func (o *SecurityScheme) SetScheme(v string) {
	o.Scheme = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SecurityScheme) GetType() Type {
	if o == nil || o.Type == nil {
		var ret Type
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityScheme) GetTypeOk() (*Type, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SecurityScheme) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given Type and assigns it to the Type field.
func (o *SecurityScheme) SetType(v Type) {
	o.Type = &v
}

func (o SecurityScheme) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BearerFormat != nil {
		toSerialize["bearerFormat"] = o.BearerFormat
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Flows != nil {
		toSerialize["flows"] = o.Flows
	}
	if o.Getref != nil {
		toSerialize["get$ref"] = o.Getref
	}
	if o.In != nil {
		toSerialize["in"] = o.In
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OpenIdConnectUrl != nil {
		toSerialize["openIdConnectUrl"] = o.OpenIdConnectUrl
	}
	if o.Scheme != nil {
		toSerialize["scheme"] = o.Scheme
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSecurityScheme struct {
	value *SecurityScheme
	isSet bool
}

func (v NullableSecurityScheme) Get() *SecurityScheme {
	return v.value
}

func (v *NullableSecurityScheme) Set(val *SecurityScheme) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityScheme) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityScheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityScheme(val *SecurityScheme) *NullableSecurityScheme {
	return &NullableSecurityScheme{value: val, isSet: true}
}

func (v NullableSecurityScheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityScheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
