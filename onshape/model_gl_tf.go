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

// GlTF struct for GlTF
type GlTF struct {
	Accessors          []Accessor                        `json:"accessors,omitempty"`
	Animations         []Animation                       `json:"animations,omitempty"`
	Asset              *Asset                            `json:"asset,omitempty"`
	BufferViews        []BufferView                      `json:"bufferViews,omitempty"`
	Buffers            []Buffer                          `json:"buffers,omitempty"`
	Cameras            []Camera                          `json:"cameras,omitempty"`
	Extensions         map[string]map[string]interface{} `json:"extensions,omitempty"`
	ExtensionsRequired []string                          `json:"extensionsRequired,omitempty"`
	ExtensionsUsed     []string                          `json:"extensionsUsed,omitempty"`
	Extras             *map[string]interface{}           `json:"extras,omitempty"`
	Images             []Image                           `json:"images,omitempty"`
	Materials          []Material                        `json:"materials,omitempty"`
	Meshes             []Mesh                            `json:"meshes,omitempty"`
	Nodes              []Node                            `json:"nodes,omitempty"`
	Samplers           []Sampler                         `json:"samplers,omitempty"`
	Scene              *int32                            `json:"scene,omitempty"`
	Scenes             []Scene                           `json:"scenes,omitempty"`
	Skins              []Skin                            `json:"skins,omitempty"`
	Textures           []Texture                         `json:"textures,omitempty"`
}

// NewGlTF instantiates a new GlTF object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlTF() *GlTF {
	this := GlTF{}
	return &this
}

// NewGlTFWithDefaults instantiates a new GlTF object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlTFWithDefaults() *GlTF {
	this := GlTF{}
	return &this
}

// GetAccessors returns the Accessors field value if set, zero value otherwise.
func (o *GlTF) GetAccessors() []Accessor {
	if o == nil || o.Accessors == nil {
		var ret []Accessor
		return ret
	}
	return o.Accessors
}

// GetAccessorsOk returns a tuple with the Accessors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlTF) GetAccessorsOk() ([]Accessor, bool) {
	if o == nil || o.Accessors == nil {
		return nil, false
	}
	return o.Accessors, true
}

// HasAccessors returns a boolean if a field has been set.
func (o *GlTF) HasAccessors() bool {
	if o != nil && o.Accessors != nil {
		return true
	}

	return false
}

// SetAccessors gets a reference to the given []Accessor and assigns it to the Accessors field.
func (o *GlTF) SetAccessors(v []Accessor) {
	o.Accessors = v
}

// GetAnimations returns the Animations field value if set, zero value otherwise.
func (o *GlTF) GetAnimations() []Animation {
	if o == nil || o.Animations == nil {
		var ret []Animation
		return ret
	}
	return o.Animations
}

// GetAnimationsOk returns a tuple with the Animations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlTF) GetAnimationsOk() ([]Animation, bool) {
	if o == nil || o.Animations == nil {
		return nil, false
	}
	return o.Animations, true
}

// HasAnimations returns a boolean if a field has been set.
func (o *GlTF) HasAnimations() bool {
	if o != nil && o.Animations != nil {
		return true
	}

	return false
}

// SetAnimations gets a reference to the given []Animation and assigns it to the Animations field.
func (o *GlTF) SetAnimations(v []Animation) {
	o.Animations = v
}

// GetAsset returns the Asset field value if set, zero value otherwise.
func (o *GlTF) GetAsset() Asset {
	if o == nil || o.Asset == nil {
		var ret Asset
		return ret
	}
	return *o.Asset
}

// GetAssetOk returns a tuple with the Asset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlTF) GetAssetOk() (*Asset, bool) {
	if o == nil || o.Asset == nil {
		return nil, false
	}
	return o.Asset, true
}

// HasAsset returns a boolean if a field has been set.
func (o *GlTF) HasAsset() bool {
	if o != nil && o.Asset != nil {
		return true
	}

	return false
}

// SetAsset gets a reference to the given Asset and assigns it to the Asset field.
func (o *GlTF) SetAsset(v Asset) {
	o.Asset = &v
}

// GetBufferViews returns the BufferViews field value if set, zero value otherwise.
func (o *GlTF) GetBufferViews() []BufferView {
	if o == nil || o.BufferViews == nil {
		var ret []BufferView
		return ret
	}
	return o.BufferViews
}

// GetBufferViewsOk returns a tuple with the BufferViews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlTF) GetBufferViewsOk() ([]BufferView, bool) {
	if o == nil || o.BufferViews == nil {
		return nil, false
	}
	return o.BufferViews, true
}

// HasBufferViews returns a boolean if a field has been set.
func (o *GlTF) HasBufferViews() bool {
	if o != nil && o.BufferViews != nil {
		return true
	}

	return false
}

// SetBufferViews gets a reference to the given []BufferView and assigns it to the BufferViews field.
func (o *GlTF) SetBufferViews(v []BufferView) {
	o.BufferViews = v
}

// GetBuffers returns the Buffers field value if set, zero value otherwise.
func (o *GlTF) GetBuffers() []Buffer {
	if o == nil || o.Buffers == nil {
		var ret []Buffer
		return ret
	}
	return o.Buffers
}

// GetBuffersOk returns a tuple with the Buffers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlTF) GetBuffersOk() ([]Buffer, bool) {
	if o == nil || o.Buffers == nil {
		return nil, false
	}
	return o.Buffers, true
}

// HasBuffers returns a boolean if a field has been set.
func (o *GlTF) HasBuffers() bool {
	if o != nil && o.Buffers != nil {
		return true
	}

	return false
}

// SetBuffers gets a reference to the given []Buffer and assigns it to the Buffers field.
func (o *GlTF) SetBuffers(v []Buffer) {
	o.Buffers = v
}

// GetCameras returns the Cameras field value if set, zero value otherwise.
func (o *GlTF) GetCameras() []Camera {
	if o == nil || o.Cameras == nil {
		var ret []Camera
		return ret
	}
	return o.Cameras
}

// GetCamerasOk returns a tuple with the Cameras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlTF) GetCamerasOk() ([]Camera, bool) {
	if o == nil || o.Cameras == nil {
		return nil, false
	}
	return o.Cameras, true
}

// HasCameras returns a boolean if a field has been set.
func (o *GlTF) HasCameras() bool {
	if o != nil && o.Cameras != nil {
		return true
	}

	return false
}

// SetCameras gets a reference to the given []Camera and assigns it to the Cameras field.
func (o *GlTF) SetCameras(v []Camera) {
	o.Cameras = v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *GlTF) GetExtensions() map[string]map[string]interface{} {
	if o == nil || o.Extensions == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlTF) GetExtensionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *GlTF) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given map[string]map[string]interface{} and assigns it to the Extensions field.
func (o *GlTF) SetExtensions(v map[string]map[string]interface{}) {
	o.Extensions = v
}

// GetExtensionsRequired returns the ExtensionsRequired field value if set, zero value otherwise.
func (o *GlTF) GetExtensionsRequired() []string {
	if o == nil || o.ExtensionsRequired == nil {
		var ret []string
		return ret
	}
	return o.ExtensionsRequired
}

// GetExtensionsRequiredOk returns a tuple with the ExtensionsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlTF) GetExtensionsRequiredOk() ([]string, bool) {
	if o == nil || o.ExtensionsRequired == nil {
		return nil, false
	}
	return o.ExtensionsRequired, true
}

// HasExtensionsRequired returns a boolean if a field has been set.
func (o *GlTF) HasExtensionsRequired() bool {
	if o != nil && o.ExtensionsRequired != nil {
		return true
	}

	return false
}

// SetExtensionsRequired gets a reference to the given []string and assigns it to the ExtensionsRequired field.
func (o *GlTF) SetExtensionsRequired(v []string) {
	o.ExtensionsRequired = v
}

// GetExtensionsUsed returns the ExtensionsUsed field value if set, zero value otherwise.
func (o *GlTF) GetExtensionsUsed() []string {
	if o == nil || o.ExtensionsUsed == nil {
		var ret []string
		return ret
	}
	return o.ExtensionsUsed
}

// GetExtensionsUsedOk returns a tuple with the ExtensionsUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlTF) GetExtensionsUsedOk() ([]string, bool) {
	if o == nil || o.ExtensionsUsed == nil {
		return nil, false
	}
	return o.ExtensionsUsed, true
}

// HasExtensionsUsed returns a boolean if a field has been set.
func (o *GlTF) HasExtensionsUsed() bool {
	if o != nil && o.ExtensionsUsed != nil {
		return true
	}

	return false
}

// SetExtensionsUsed gets a reference to the given []string and assigns it to the ExtensionsUsed field.
func (o *GlTF) SetExtensionsUsed(v []string) {
	o.ExtensionsUsed = v
}

// GetExtras returns the Extras field value if set, zero value otherwise.
func (o *GlTF) GetExtras() map[string]interface{} {
	if o == nil || o.Extras == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Extras
}

// GetExtrasOk returns a tuple with the Extras field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlTF) GetExtrasOk() (*map[string]interface{}, bool) {
	if o == nil || o.Extras == nil {
		return nil, false
	}
	return o.Extras, true
}

// HasExtras returns a boolean if a field has been set.
func (o *GlTF) HasExtras() bool {
	if o != nil && o.Extras != nil {
		return true
	}

	return false
}

// SetExtras gets a reference to the given map[string]interface{} and assigns it to the Extras field.
func (o *GlTF) SetExtras(v map[string]interface{}) {
	o.Extras = &v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *GlTF) GetImages() []Image {
	if o == nil || o.Images == nil {
		var ret []Image
		return ret
	}
	return o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlTF) GetImagesOk() ([]Image, bool) {
	if o == nil || o.Images == nil {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *GlTF) HasImages() bool {
	if o != nil && o.Images != nil {
		return true
	}

	return false
}

// SetImages gets a reference to the given []Image and assigns it to the Images field.
func (o *GlTF) SetImages(v []Image) {
	o.Images = v
}

// GetMaterials returns the Materials field value if set, zero value otherwise.
func (o *GlTF) GetMaterials() []Material {
	if o == nil || o.Materials == nil {
		var ret []Material
		return ret
	}
	return o.Materials
}

// GetMaterialsOk returns a tuple with the Materials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlTF) GetMaterialsOk() ([]Material, bool) {
	if o == nil || o.Materials == nil {
		return nil, false
	}
	return o.Materials, true
}

// HasMaterials returns a boolean if a field has been set.
func (o *GlTF) HasMaterials() bool {
	if o != nil && o.Materials != nil {
		return true
	}

	return false
}

// SetMaterials gets a reference to the given []Material and assigns it to the Materials field.
func (o *GlTF) SetMaterials(v []Material) {
	o.Materials = v
}

// GetMeshes returns the Meshes field value if set, zero value otherwise.
func (o *GlTF) GetMeshes() []Mesh {
	if o == nil || o.Meshes == nil {
		var ret []Mesh
		return ret
	}
	return o.Meshes
}

// GetMeshesOk returns a tuple with the Meshes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlTF) GetMeshesOk() ([]Mesh, bool) {
	if o == nil || o.Meshes == nil {
		return nil, false
	}
	return o.Meshes, true
}

// HasMeshes returns a boolean if a field has been set.
func (o *GlTF) HasMeshes() bool {
	if o != nil && o.Meshes != nil {
		return true
	}

	return false
}

// SetMeshes gets a reference to the given []Mesh and assigns it to the Meshes field.
func (o *GlTF) SetMeshes(v []Mesh) {
	o.Meshes = v
}

// GetNodes returns the Nodes field value if set, zero value otherwise.
func (o *GlTF) GetNodes() []Node {
	if o == nil || o.Nodes == nil {
		var ret []Node
		return ret
	}
	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlTF) GetNodesOk() ([]Node, bool) {
	if o == nil || o.Nodes == nil {
		return nil, false
	}
	return o.Nodes, true
}

// HasNodes returns a boolean if a field has been set.
func (o *GlTF) HasNodes() bool {
	if o != nil && o.Nodes != nil {
		return true
	}

	return false
}

// SetNodes gets a reference to the given []Node and assigns it to the Nodes field.
func (o *GlTF) SetNodes(v []Node) {
	o.Nodes = v
}

// GetSamplers returns the Samplers field value if set, zero value otherwise.
func (o *GlTF) GetSamplers() []Sampler {
	if o == nil || o.Samplers == nil {
		var ret []Sampler
		return ret
	}
	return o.Samplers
}

// GetSamplersOk returns a tuple with the Samplers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlTF) GetSamplersOk() ([]Sampler, bool) {
	if o == nil || o.Samplers == nil {
		return nil, false
	}
	return o.Samplers, true
}

// HasSamplers returns a boolean if a field has been set.
func (o *GlTF) HasSamplers() bool {
	if o != nil && o.Samplers != nil {
		return true
	}

	return false
}

// SetSamplers gets a reference to the given []Sampler and assigns it to the Samplers field.
func (o *GlTF) SetSamplers(v []Sampler) {
	o.Samplers = v
}

// GetScene returns the Scene field value if set, zero value otherwise.
func (o *GlTF) GetScene() int32 {
	if o == nil || o.Scene == nil {
		var ret int32
		return ret
	}
	return *o.Scene
}

// GetSceneOk returns a tuple with the Scene field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlTF) GetSceneOk() (*int32, bool) {
	if o == nil || o.Scene == nil {
		return nil, false
	}
	return o.Scene, true
}

// HasScene returns a boolean if a field has been set.
func (o *GlTF) HasScene() bool {
	if o != nil && o.Scene != nil {
		return true
	}

	return false
}

// SetScene gets a reference to the given int32 and assigns it to the Scene field.
func (o *GlTF) SetScene(v int32) {
	o.Scene = &v
}

// GetScenes returns the Scenes field value if set, zero value otherwise.
func (o *GlTF) GetScenes() []Scene {
	if o == nil || o.Scenes == nil {
		var ret []Scene
		return ret
	}
	return o.Scenes
}

// GetScenesOk returns a tuple with the Scenes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlTF) GetScenesOk() ([]Scene, bool) {
	if o == nil || o.Scenes == nil {
		return nil, false
	}
	return o.Scenes, true
}

// HasScenes returns a boolean if a field has been set.
func (o *GlTF) HasScenes() bool {
	if o != nil && o.Scenes != nil {
		return true
	}

	return false
}

// SetScenes gets a reference to the given []Scene and assigns it to the Scenes field.
func (o *GlTF) SetScenes(v []Scene) {
	o.Scenes = v
}

// GetSkins returns the Skins field value if set, zero value otherwise.
func (o *GlTF) GetSkins() []Skin {
	if o == nil || o.Skins == nil {
		var ret []Skin
		return ret
	}
	return o.Skins
}

// GetSkinsOk returns a tuple with the Skins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlTF) GetSkinsOk() ([]Skin, bool) {
	if o == nil || o.Skins == nil {
		return nil, false
	}
	return o.Skins, true
}

// HasSkins returns a boolean if a field has been set.
func (o *GlTF) HasSkins() bool {
	if o != nil && o.Skins != nil {
		return true
	}

	return false
}

// SetSkins gets a reference to the given []Skin and assigns it to the Skins field.
func (o *GlTF) SetSkins(v []Skin) {
	o.Skins = v
}

// GetTextures returns the Textures field value if set, zero value otherwise.
func (o *GlTF) GetTextures() []Texture {
	if o == nil || o.Textures == nil {
		var ret []Texture
		return ret
	}
	return o.Textures
}

// GetTexturesOk returns a tuple with the Textures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlTF) GetTexturesOk() ([]Texture, bool) {
	if o == nil || o.Textures == nil {
		return nil, false
	}
	return o.Textures, true
}

// HasTextures returns a boolean if a field has been set.
func (o *GlTF) HasTextures() bool {
	if o != nil && o.Textures != nil {
		return true
	}

	return false
}

// SetTextures gets a reference to the given []Texture and assigns it to the Textures field.
func (o *GlTF) SetTextures(v []Texture) {
	o.Textures = v
}

func (o GlTF) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Accessors != nil {
		toSerialize["accessors"] = o.Accessors
	}
	if o.Animations != nil {
		toSerialize["animations"] = o.Animations
	}
	if o.Asset != nil {
		toSerialize["asset"] = o.Asset
	}
	if o.BufferViews != nil {
		toSerialize["bufferViews"] = o.BufferViews
	}
	if o.Buffers != nil {
		toSerialize["buffers"] = o.Buffers
	}
	if o.Cameras != nil {
		toSerialize["cameras"] = o.Cameras
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.ExtensionsRequired != nil {
		toSerialize["extensionsRequired"] = o.ExtensionsRequired
	}
	if o.ExtensionsUsed != nil {
		toSerialize["extensionsUsed"] = o.ExtensionsUsed
	}
	if o.Extras != nil {
		toSerialize["extras"] = o.Extras
	}
	if o.Images != nil {
		toSerialize["images"] = o.Images
	}
	if o.Materials != nil {
		toSerialize["materials"] = o.Materials
	}
	if o.Meshes != nil {
		toSerialize["meshes"] = o.Meshes
	}
	if o.Nodes != nil {
		toSerialize["nodes"] = o.Nodes
	}
	if o.Samplers != nil {
		toSerialize["samplers"] = o.Samplers
	}
	if o.Scene != nil {
		toSerialize["scene"] = o.Scene
	}
	if o.Scenes != nil {
		toSerialize["scenes"] = o.Scenes
	}
	if o.Skins != nil {
		toSerialize["skins"] = o.Skins
	}
	if o.Textures != nil {
		toSerialize["textures"] = o.Textures
	}
	return json.Marshal(toSerialize)
}

type NullableGlTF struct {
	value *GlTF
	isSet bool
}

func (v NullableGlTF) Get() *GlTF {
	return v.value
}

func (v *NullableGlTF) Set(val *GlTF) {
	v.value = val
	v.isSet = true
}

func (v NullableGlTF) IsSet() bool {
	return v.isSet
}

func (v *NullableGlTF) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlTF(val *GlTF) *NullableGlTF {
	return &NullableGlTF{value: val, isSet: true}
}

func (v NullableGlTF) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlTF) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
