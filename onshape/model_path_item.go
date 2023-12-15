/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.173.27678-64d64396ca66
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// PathItem struct for PathItem
type PathItem struct {
	Delete      *Operation                        `json:"delete,omitempty"`
	Description *string                           `json:"description,omitempty"`
	Extensions  map[string]map[string]interface{} `json:"extensions,omitempty"`
	Get         *Operation                        `json:"get,omitempty"`
	Getref      *string                           `json:"get$ref,omitempty"`
	Head        *Operation                        `json:"head,omitempty"`
	Options     *Operation                        `json:"options,omitempty"`
	Parameters  []Parameter                       `json:"parameters,omitempty"`
	Patch       *Operation                        `json:"patch,omitempty"`
	Post        *Operation                        `json:"post,omitempty"`
	Put         *Operation                        `json:"put,omitempty"`
	Servers     []Server                          `json:"servers,omitempty"`
	Summary     *string                           `json:"summary,omitempty"`
	Trace       *Operation                        `json:"trace,omitempty"`
}

// NewPathItem instantiates a new PathItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPathItem() *PathItem {
	this := PathItem{}
	return &this
}

// NewPathItemWithDefaults instantiates a new PathItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPathItemWithDefaults() *PathItem {
	this := PathItem{}
	return &this
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *PathItem) GetDelete() Operation {
	if o == nil || o.Delete == nil {
		var ret Operation
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathItem) GetDeleteOk() (*Operation, bool) {
	if o == nil || o.Delete == nil {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *PathItem) HasDelete() bool {
	if o != nil && o.Delete != nil {
		return true
	}

	return false
}

// SetDelete gets a reference to the given Operation and assigns it to the Delete field.
func (o *PathItem) SetDelete(v Operation) {
	o.Delete = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PathItem) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathItem) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PathItem) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PathItem) SetDescription(v string) {
	o.Description = &v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *PathItem) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathItem) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *PathItem) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *PathItem) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetGet returns the Get field value if set, zero value otherwise.
func (o *PathItem) GetGet() Operation {
	if o == nil || o.Get == nil {
		var ret Operation
		return ret
	}
	return *o.Get
}

// GetGetOk returns a tuple with the Get field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathItem) GetGetOk() (*Operation, bool) {
	if o == nil || o.Get == nil {
		return nil, false
	}
	return o.Get, true
}

// HasGet returns a boolean if a field has been set.
func (o *PathItem) HasGet() bool {
	if o != nil && o.Get != nil {
		return true
	}

	return false
}

// SetGet gets a reference to the given Operation and assigns it to the Get field.
func (o *PathItem) SetGet(v Operation) {
	o.Get = &v
}

// GetGetref returns the Getref field value if set, zero value otherwise.
func (o *PathItem) GetGetref() string {
	if o == nil || o.Getref == nil {
		var ret string
		return ret
	}
	return *o.Getref
}

// GetGetrefOk returns a tuple with the Getref field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathItem) GetGetrefOk() (*string, bool) {
	if o == nil || o.Getref == nil {
		return nil, false
	}
	return o.Getref, true
}

// HasGetref returns a boolean if a field has been set.
func (o *PathItem) HasGetref() bool {
	if o != nil && o.Getref != nil {
		return true
	}

	return false
}

// SetGetref gets a reference to the given string and assigns it to the Getref field.
func (o *PathItem) SetGetref(v string) {
	o.Getref = &v
}

// GetHead returns the Head field value if set, zero value otherwise.
func (o *PathItem) GetHead() Operation {
	if o == nil || o.Head == nil {
		var ret Operation
		return ret
	}
	return *o.Head
}

// GetHeadOk returns a tuple with the Head field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathItem) GetHeadOk() (*Operation, bool) {
	if o == nil || o.Head == nil {
		return nil, false
	}
	return o.Head, true
}

// HasHead returns a boolean if a field has been set.
func (o *PathItem) HasHead() bool {
	if o != nil && o.Head != nil {
		return true
	}

	return false
}

// SetHead gets a reference to the given Operation and assigns it to the Head field.
func (o *PathItem) SetHead(v Operation) {
	o.Head = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *PathItem) GetOptions() Operation {
	if o == nil || o.Options == nil {
		var ret Operation
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathItem) GetOptionsOk() (*Operation, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *PathItem) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given Operation and assigns it to the Options field.
func (o *PathItem) SetOptions(v Operation) {
	o.Options = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *PathItem) GetParameters() []Parameter {
	if o == nil || o.Parameters == nil {
		var ret []Parameter
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathItem) GetParametersOk() ([]Parameter, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *PathItem) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []Parameter and assigns it to the Parameters field.
func (o *PathItem) SetParameters(v []Parameter) {
	o.Parameters = v
}

// GetPatch returns the Patch field value if set, zero value otherwise.
func (o *PathItem) GetPatch() Operation {
	if o == nil || o.Patch == nil {
		var ret Operation
		return ret
	}
	return *o.Patch
}

// GetPatchOk returns a tuple with the Patch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathItem) GetPatchOk() (*Operation, bool) {
	if o == nil || o.Patch == nil {
		return nil, false
	}
	return o.Patch, true
}

// HasPatch returns a boolean if a field has been set.
func (o *PathItem) HasPatch() bool {
	if o != nil && o.Patch != nil {
		return true
	}

	return false
}

// SetPatch gets a reference to the given Operation and assigns it to the Patch field.
func (o *PathItem) SetPatch(v Operation) {
	o.Patch = &v
}

// GetPost returns the Post field value if set, zero value otherwise.
func (o *PathItem) GetPost() Operation {
	if o == nil || o.Post == nil {
		var ret Operation
		return ret
	}
	return *o.Post
}

// GetPostOk returns a tuple with the Post field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathItem) GetPostOk() (*Operation, bool) {
	if o == nil || o.Post == nil {
		return nil, false
	}
	return o.Post, true
}

// HasPost returns a boolean if a field has been set.
func (o *PathItem) HasPost() bool {
	if o != nil && o.Post != nil {
		return true
	}

	return false
}

// SetPost gets a reference to the given Operation and assigns it to the Post field.
func (o *PathItem) SetPost(v Operation) {
	o.Post = &v
}

// GetPut returns the Put field value if set, zero value otherwise.
func (o *PathItem) GetPut() Operation {
	if o == nil || o.Put == nil {
		var ret Operation
		return ret
	}
	return *o.Put
}

// GetPutOk returns a tuple with the Put field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathItem) GetPutOk() (*Operation, bool) {
	if o == nil || o.Put == nil {
		return nil, false
	}
	return o.Put, true
}

// HasPut returns a boolean if a field has been set.
func (o *PathItem) HasPut() bool {
	if o != nil && o.Put != nil {
		return true
	}

	return false
}

// SetPut gets a reference to the given Operation and assigns it to the Put field.
func (o *PathItem) SetPut(v Operation) {
	o.Put = &v
}

// GetServers returns the Servers field value if set, zero value otherwise.
func (o *PathItem) GetServers() []Server {
	if o == nil || o.Servers == nil {
		var ret []Server
		return ret
	}
	return o.Servers
}

// GetServersOk returns a tuple with the Servers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathItem) GetServersOk() ([]Server, bool) {
	if o == nil || o.Servers == nil {
		return nil, false
	}
	return o.Servers, true
}

// HasServers returns a boolean if a field has been set.
func (o *PathItem) HasServers() bool {
	if o != nil && o.Servers != nil {
		return true
	}

	return false
}

// SetServers gets a reference to the given []Server and assigns it to the Servers field.
func (o *PathItem) SetServers(v []Server) {
	o.Servers = v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *PathItem) GetSummary() string {
	if o == nil || o.Summary == nil {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathItem) GetSummaryOk() (*string, bool) {
	if o == nil || o.Summary == nil {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *PathItem) HasSummary() bool {
	if o != nil && o.Summary != nil {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *PathItem) SetSummary(v string) {
	o.Summary = &v
}

// GetTrace returns the Trace field value if set, zero value otherwise.
func (o *PathItem) GetTrace() Operation {
	if o == nil || o.Trace == nil {
		var ret Operation
		return ret
	}
	return *o.Trace
}

// GetTraceOk returns a tuple with the Trace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PathItem) GetTraceOk() (*Operation, bool) {
	if o == nil || o.Trace == nil {
		return nil, false
	}
	return o.Trace, true
}

// HasTrace returns a boolean if a field has been set.
func (o *PathItem) HasTrace() bool {
	if o != nil && o.Trace != nil {
		return true
	}

	return false
}

// SetTrace gets a reference to the given Operation and assigns it to the Trace field.
func (o *PathItem) SetTrace(v Operation) {
	o.Trace = &v
}

func (o PathItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Delete != nil {
		toSerialize["delete"] = o.Delete
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Get != nil {
		toSerialize["get"] = o.Get
	}
	if o.Getref != nil {
		toSerialize["get$ref"] = o.Getref
	}
	if o.Head != nil {
		toSerialize["head"] = o.Head
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.Patch != nil {
		toSerialize["patch"] = o.Patch
	}
	if o.Post != nil {
		toSerialize["post"] = o.Post
	}
	if o.Put != nil {
		toSerialize["put"] = o.Put
	}
	if o.Servers != nil {
		toSerialize["servers"] = o.Servers
	}
	if o.Summary != nil {
		toSerialize["summary"] = o.Summary
	}
	if o.Trace != nil {
		toSerialize["trace"] = o.Trace
	}
	return json.Marshal(toSerialize)
}

type NullablePathItem struct {
	value *PathItem
	isSet bool
}

func (v NullablePathItem) Get() *PathItem {
	return v.value
}

func (v *NullablePathItem) Set(val *PathItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePathItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePathItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePathItem(val *PathItem) *NullablePathItem {
	return &NullablePathItem{value: val, isSet: true}
}

func (v NullablePathItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePathItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
