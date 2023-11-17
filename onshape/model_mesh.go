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

// Mesh struct for Mesh
type Mesh struct {
	Extensions map[string]map[string]interface{} `json:"extensions,omitempty"`
	Extras     *map[string]interface{}           `json:"extras,omitempty"`
	Name       *string                           `json:"name,omitempty"`
	Primitives []MeshPrimitive                   `json:"primitives,omitempty"`
	Weights    []float32                         `json:"weights,omitempty"`
}

// NewMesh instantiates a new Mesh object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMesh() *Mesh {
	this := Mesh{}
	return &this
}

// NewMeshWithDefaults instantiates a new Mesh object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeshWithDefaults() *Mesh {
	this := Mesh{}
	return &this
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *Mesh) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mesh) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *Mesh) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *Mesh) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *Mesh) GetExtras() map[string]interface{} {
	if o == nil || o.Extras == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mesh) GetExtrasOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *Mesh) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *Mesh) SetExtras(v map[string]interface{}) {
	o.Extras = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Mesh) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mesh) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Mesh) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Mesh) SetName(v string) {
	o.Name = &v
}

// GetPrimitives returns the Primitives field value if set, zero value otherwise.
func (o *Mesh) GetPrimitives() []MeshPrimitive {
	if o == nil || o.Primitives == nil {
		var ret []MeshPrimitive
		return ret
	}
	return o.Primitives
}

// GetPrimitivesOk returns a tuple with the Primitives field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mesh) GetPrimitivesOk() ([]MeshPrimitive, bool) {
	if o == nil || o.Primitives == nil {
		return nil, false
	}
	return o.Primitives, true
}

// HasPrimitives returns a boolean if a field has been set.
func (o *Mesh) HasPrimitives() bool {
	if o != nil && o.Primitives != nil {
		return true
	}

	return false
}

// SetPrimitives gets a reference to the given []MeshPrimitive and assigns it to the Primitives field.
func (o *Mesh) SetPrimitives(v []MeshPrimitive) {
	o.Primitives = v
}

// GetWeights returns the Weights field value if set, zero value otherwise.
func (o *Mesh) GetWeights() []float32 {
	if o == nil || o.Weights == nil {
		var ret []float32
		return ret
	}
	return o.Weights
}

// GetWeightsOk returns a tuple with the Weights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Mesh) GetWeightsOk() ([]float32, bool) {
	if o == nil || o.Weights == nil {
		return nil, false
	}
	return o.Weights, true
}

// HasWeights returns a boolean if a field has been set.
func (o *Mesh) HasWeights() bool {
	if o != nil && o.Weights != nil {
		return true
	}

	return false
}

// SetWeights gets a reference to the given []float32 and assigns it to the Weights field.
func (o *Mesh) SetWeights(v []float32) {
	o.Weights = v
}

func (o Mesh) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Primitives != nil {
		toSerialize["primitives"] = o.Primitives
	}
	if o.Weights != nil {
		toSerialize["weights"] = o.Weights
	}
	return json.Marshal(toSerialize)
}

type NullableMesh struct {
	value *Mesh
	isSet bool
}

func (v NullableMesh) Get() *Mesh {
	return v.value
}

func (v *NullableMesh) Set(val *Mesh) {
	v.value = val
	v.isSet = true
}

func (v NullableMesh) IsSet() bool {
	return v.isSet
}

func (v *NullableMesh) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMesh(val *Mesh) *NullableMesh {
	return &NullableMesh{value: val, isSet: true}
}

func (v NullableMesh) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMesh) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}