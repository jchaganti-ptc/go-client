/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.155.7232-a44b68534e12
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTExportTessellatedFacesResponse898 struct for BTExportTessellatedFacesResponse898
type BTExportTessellatedFacesResponse898 struct {
	Bodies                           []BTExportTessellatedBody3398   `json:"bodies,omitempty"`
	BodiesInfo                       *BTExportModelBodiesResponse734 `json:"bodiesInfo,omitempty"`
	CombineCompositePartConstituents *bool                           `json:"combineCompositePartConstituents,omitempty"`
	DisplayData                      *BTPartStudioDisplayData346     `json:"displayData,omitempty"`
	DocumentId                       *string                         `json:"documentId,omitempty"`
	ElementId                        *string                         `json:"elementId,omitempty"`
	ErrorEnum                        *string                         `json:"errorEnum,omitempty"`
	FacetPoints                      []BTVector3d389                 `json:"facetPoints,omitempty"`
	FullElementId                    *BTFullElementId756             `json:"fullElementId,omitempty"`
	OutputFaceAppearances            *bool                           `json:"outputFaceAppearances,omitempty"`
	OutputSeparateFaceNodes          *bool                           `json:"outputSeparateFaceNodes,omitempty"`
}

// NewBTExportTessellatedFacesResponse898 instantiates a new BTExportTessellatedFacesResponse898 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExportTessellatedFacesResponse898() *BTExportTessellatedFacesResponse898 {
	this := BTExportTessellatedFacesResponse898{}
	return &this
}

// NewBTExportTessellatedFacesResponse898WithDefaults instantiates a new BTExportTessellatedFacesResponse898 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExportTessellatedFacesResponse898WithDefaults() *BTExportTessellatedFacesResponse898 {
	this := BTExportTessellatedFacesResponse898{}
	return &this
}

// GetBodies returns the Bodies field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesResponse898) GetBodies() []BTExportTessellatedBody3398 {
	if o == nil || o.Bodies == nil {
		var ret []BTExportTessellatedBody3398
		return ret
	}
	return o.Bodies
}

// GetBodiesOk returns a tuple with the Bodies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesResponse898) GetBodiesOk() ([]BTExportTessellatedBody3398, bool) {
	if o == nil || o.Bodies == nil {
		return nil, false
	}
	return o.Bodies, true
}

// HasBodies returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesResponse898) HasBodies() bool {
	if o != nil && o.Bodies != nil {
		return true
	}

	return false
}

// SetBodies gets a reference to the given []BTExportTessellatedBody3398 and assigns it to the Bodies field.
func (o *BTExportTessellatedFacesResponse898) SetBodies(v []BTExportTessellatedBody3398) {
	o.Bodies = v
}

// GetBodiesInfo returns the BodiesInfo field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesResponse898) GetBodiesInfo() BTExportModelBodiesResponse734 {
	if o == nil || o.BodiesInfo == nil {
		var ret BTExportModelBodiesResponse734
		return ret
	}
	return *o.BodiesInfo
}

// GetBodiesInfoOk returns a tuple with the BodiesInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesResponse898) GetBodiesInfoOk() (*BTExportModelBodiesResponse734, bool) {
	if o == nil || o.BodiesInfo == nil {
		return nil, false
	}
	return o.BodiesInfo, true
}

// HasBodiesInfo returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesResponse898) HasBodiesInfo() bool {
	if o != nil && o.BodiesInfo != nil {
		return true
	}

	return false
}

// SetBodiesInfo gets a reference to the given BTExportModelBodiesResponse734 and assigns it to the BodiesInfo field.
func (o *BTExportTessellatedFacesResponse898) SetBodiesInfo(v BTExportModelBodiesResponse734) {
	o.BodiesInfo = &v
}

// GetCombineCompositePartConstituents returns the CombineCompositePartConstituents field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesResponse898) GetCombineCompositePartConstituents() bool {
	if o == nil || o.CombineCompositePartConstituents == nil {
		var ret bool
		return ret
	}
	return *o.CombineCompositePartConstituents
}

// GetCombineCompositePartConstituentsOk returns a tuple with the CombineCompositePartConstituents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesResponse898) GetCombineCompositePartConstituentsOk() (*bool, bool) {
	if o == nil || o.CombineCompositePartConstituents == nil {
		return nil, false
	}
	return o.CombineCompositePartConstituents, true
}

// HasCombineCompositePartConstituents returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesResponse898) HasCombineCompositePartConstituents() bool {
	if o != nil && o.CombineCompositePartConstituents != nil {
		return true
	}

	return false
}

// SetCombineCompositePartConstituents gets a reference to the given bool and assigns it to the CombineCompositePartConstituents field.
func (o *BTExportTessellatedFacesResponse898) SetCombineCompositePartConstituents(v bool) {
	o.CombineCompositePartConstituents = &v
}

// GetDisplayData returns the DisplayData field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesResponse898) GetDisplayData() BTPartStudioDisplayData346 {
	if o == nil || o.DisplayData == nil {
		var ret BTPartStudioDisplayData346
		return ret
	}
	return *o.DisplayData
}

// GetDisplayDataOk returns a tuple with the DisplayData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesResponse898) GetDisplayDataOk() (*BTPartStudioDisplayData346, bool) {
	if o == nil || o.DisplayData == nil {
		return nil, false
	}
	return o.DisplayData, true
}

// HasDisplayData returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesResponse898) HasDisplayData() bool {
	if o != nil && o.DisplayData != nil {
		return true
	}

	return false
}

// SetDisplayData gets a reference to the given BTPartStudioDisplayData346 and assigns it to the DisplayData field.
func (o *BTExportTessellatedFacesResponse898) SetDisplayData(v BTPartStudioDisplayData346) {
	o.DisplayData = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesResponse898) GetDocumentId() string {
	if o == nil || o.DocumentId == nil {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesResponse898) GetDocumentIdOk() (*string, bool) {
	if o == nil || o.DocumentId == nil {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesResponse898) HasDocumentId() bool {
	if o != nil && o.DocumentId != nil {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *BTExportTessellatedFacesResponse898) SetDocumentId(v string) {
	o.DocumentId = &v
}

// GetElementId returns the ElementId field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesResponse898) GetElementId() string {
	if o == nil || o.ElementId == nil {
		var ret string
		return ret
	}
	return *o.ElementId
}

// GetElementIdOk returns a tuple with the ElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesResponse898) GetElementIdOk() (*string, bool) {
	if o == nil || o.ElementId == nil {
		return nil, false
	}
	return o.ElementId, true
}

// HasElementId returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesResponse898) HasElementId() bool {
	if o != nil && o.ElementId != nil {
		return true
	}

	return false
}

// SetElementId gets a reference to the given string and assigns it to the ElementId field.
func (o *BTExportTessellatedFacesResponse898) SetElementId(v string) {
	o.ElementId = &v
}

// GetErrorEnum returns the ErrorEnum field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesResponse898) GetErrorEnum() string {
	if o == nil || o.ErrorEnum == nil {
		var ret string
		return ret
	}
	return *o.ErrorEnum
}

// GetErrorEnumOk returns a tuple with the ErrorEnum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesResponse898) GetErrorEnumOk() (*string, bool) {
	if o == nil || o.ErrorEnum == nil {
		return nil, false
	}
	return o.ErrorEnum, true
}

// HasErrorEnum returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesResponse898) HasErrorEnum() bool {
	if o != nil && o.ErrorEnum != nil {
		return true
	}

	return false
}

// SetErrorEnum gets a reference to the given string and assigns it to the ErrorEnum field.
func (o *BTExportTessellatedFacesResponse898) SetErrorEnum(v string) {
	o.ErrorEnum = &v
}

// GetFacetPoints returns the FacetPoints field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesResponse898) GetFacetPoints() []BTVector3d389 {
	if o == nil || o.FacetPoints == nil {
		var ret []BTVector3d389
		return ret
	}
	return o.FacetPoints
}

// GetFacetPointsOk returns a tuple with the FacetPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesResponse898) GetFacetPointsOk() ([]BTVector3d389, bool) {
	if o == nil || o.FacetPoints == nil {
		return nil, false
	}
	return o.FacetPoints, true
}

// HasFacetPoints returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesResponse898) HasFacetPoints() bool {
	if o != nil && o.FacetPoints != nil {
		return true
	}

	return false
}

// SetFacetPoints gets a reference to the given []BTVector3d389 and assigns it to the FacetPoints field.
func (o *BTExportTessellatedFacesResponse898) SetFacetPoints(v []BTVector3d389) {
	o.FacetPoints = v
}

// GetFullElementId returns the FullElementId field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesResponse898) GetFullElementId() BTFullElementId756 {
	if o == nil || o.FullElementId == nil {
		var ret BTFullElementId756
		return ret
	}
	return *o.FullElementId
}

// GetFullElementIdOk returns a tuple with the FullElementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesResponse898) GetFullElementIdOk() (*BTFullElementId756, bool) {
	if o == nil || o.FullElementId == nil {
		return nil, false
	}
	return o.FullElementId, true
}

// HasFullElementId returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesResponse898) HasFullElementId() bool {
	if o != nil && o.FullElementId != nil {
		return true
	}

	return false
}

// SetFullElementId gets a reference to the given BTFullElementId756 and assigns it to the FullElementId field.
func (o *BTExportTessellatedFacesResponse898) SetFullElementId(v BTFullElementId756) {
	o.FullElementId = &v
}

// GetOutputFaceAppearances returns the OutputFaceAppearances field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesResponse898) GetOutputFaceAppearances() bool {
	if o == nil || o.OutputFaceAppearances == nil {
		var ret bool
		return ret
	}
	return *o.OutputFaceAppearances
}

// GetOutputFaceAppearancesOk returns a tuple with the OutputFaceAppearances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesResponse898) GetOutputFaceAppearancesOk() (*bool, bool) {
	if o == nil || o.OutputFaceAppearances == nil {
		return nil, false
	}
	return o.OutputFaceAppearances, true
}

// HasOutputFaceAppearances returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesResponse898) HasOutputFaceAppearances() bool {
	if o != nil && o.OutputFaceAppearances != nil {
		return true
	}

	return false
}

// SetOutputFaceAppearances gets a reference to the given bool and assigns it to the OutputFaceAppearances field.
func (o *BTExportTessellatedFacesResponse898) SetOutputFaceAppearances(v bool) {
	o.OutputFaceAppearances = &v
}

// GetOutputSeparateFaceNodes returns the OutputSeparateFaceNodes field value if set, zero value otherwise.
func (o *BTExportTessellatedFacesResponse898) GetOutputSeparateFaceNodes() bool {
	if o == nil || o.OutputSeparateFaceNodes == nil {
		var ret bool
		return ret
	}
	return *o.OutputSeparateFaceNodes
}

// GetOutputSeparateFaceNodesOk returns a tuple with the OutputSeparateFaceNodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportTessellatedFacesResponse898) GetOutputSeparateFaceNodesOk() (*bool, bool) {
	if o == nil || o.OutputSeparateFaceNodes == nil {
		return nil, false
	}
	return o.OutputSeparateFaceNodes, true
}

// HasOutputSeparateFaceNodes returns a boolean if a field has been set.
func (o *BTExportTessellatedFacesResponse898) HasOutputSeparateFaceNodes() bool {
	if o != nil && o.OutputSeparateFaceNodes != nil {
		return true
	}

	return false
}

// SetOutputSeparateFaceNodes gets a reference to the given bool and assigns it to the OutputSeparateFaceNodes field.
func (o *BTExportTessellatedFacesResponse898) SetOutputSeparateFaceNodes(v bool) {
	o.OutputSeparateFaceNodes = &v
}

func (o BTExportTessellatedFacesResponse898) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bodies != nil {
		toSerialize["bodies"] = o.Bodies
	}
	if o.BodiesInfo != nil {
		toSerialize["bodiesInfo"] = o.BodiesInfo
	}
	if o.CombineCompositePartConstituents != nil {
		toSerialize["combineCompositePartConstituents"] = o.CombineCompositePartConstituents
	}
	if o.DisplayData != nil {
		toSerialize["displayData"] = o.DisplayData
	}
	if o.DocumentId != nil {
		toSerialize["documentId"] = o.DocumentId
	}
	if o.ElementId != nil {
		toSerialize["elementId"] = o.ElementId
	}
	if o.ErrorEnum != nil {
		toSerialize["errorEnum"] = o.ErrorEnum
	}
	if o.FacetPoints != nil {
		toSerialize["facetPoints"] = o.FacetPoints
	}
	if o.FullElementId != nil {
		toSerialize["fullElementId"] = o.FullElementId
	}
	if o.OutputFaceAppearances != nil {
		toSerialize["outputFaceAppearances"] = o.OutputFaceAppearances
	}
	if o.OutputSeparateFaceNodes != nil {
		toSerialize["outputSeparateFaceNodes"] = o.OutputSeparateFaceNodes
	}
	return json.Marshal(toSerialize)
}

type NullableBTExportTessellatedFacesResponse898 struct {
	value *BTExportTessellatedFacesResponse898
	isSet bool
}

func (v NullableBTExportTessellatedFacesResponse898) Get() *BTExportTessellatedFacesResponse898 {
	return v.value
}

func (v *NullableBTExportTessellatedFacesResponse898) Set(val *BTExportTessellatedFacesResponse898) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExportTessellatedFacesResponse898) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExportTessellatedFacesResponse898) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExportTessellatedFacesResponse898(val *BTExportTessellatedFacesResponse898) *NullableBTExportTessellatedFacesResponse898 {
	return &NullableBTExportTessellatedFacesResponse898{value: val, isSet: true}
}

func (v NullableBTExportTessellatedFacesResponse898) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExportTessellatedFacesResponse898) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
