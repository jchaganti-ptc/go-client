/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5884-a8034368dd2e
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTExportModelEdge1782 struct for BTExportModelEdge1782
type BTExportModelEdge1782 struct {
	Curve    *BTCurveDescription1583        `json:"curve,omitempty"`
	Geometry *BTExportModelEdgeGeometry1125 `json:"geometry,omitempty"`
	Id       *string                        `json:"id,omitempty"`
	Vertices []string                       `json:"vertices,omitempty"`
}

// NewBTExportModelEdge1782 instantiates a new BTExportModelEdge1782 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTExportModelEdge1782() *BTExportModelEdge1782 {
	this := BTExportModelEdge1782{}
	return &this
}

// NewBTExportModelEdge1782WithDefaults instantiates a new BTExportModelEdge1782 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTExportModelEdge1782WithDefaults() *BTExportModelEdge1782 {
	this := BTExportModelEdge1782{}
	return &this
}

// GetCurve returns the Curve field value if set, zero value otherwise.
func (o *BTExportModelEdge1782) GetCurve() BTCurveDescription1583 {
	if o == nil || o.Curve == nil {
		var ret BTCurveDescription1583
		return ret
	}
	return *o.Curve
}

// GetCurveOk returns a tuple with the Curve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelEdge1782) GetCurveOk() (*BTCurveDescription1583, bool) {
	if o == nil || o.Curve == nil {
		return nil, false
	}
	return o.Curve, true
}

// HasCurve returns a boolean if a field has been set.
func (o *BTExportModelEdge1782) HasCurve() bool {
	if o != nil && o.Curve != nil {
		return true
	}

	return false
}

// SetCurve gets a reference to the given BTCurveDescription1583 and assigns it to the Curve field.
func (o *BTExportModelEdge1782) SetCurve(v BTCurveDescription1583) {
	o.Curve = &v
}

// GetGeometry returns the Geometry field value if set, zero value otherwise.
func (o *BTExportModelEdge1782) GetGeometry() BTExportModelEdgeGeometry1125 {
	if o == nil || o.Geometry == nil {
		var ret BTExportModelEdgeGeometry1125
		return ret
	}
	return *o.Geometry
}

// GetGeometryOk returns a tuple with the Geometry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelEdge1782) GetGeometryOk() (*BTExportModelEdgeGeometry1125, bool) {
	if o == nil || o.Geometry == nil {
		return nil, false
	}
	return o.Geometry, true
}

// HasGeometry returns a boolean if a field has been set.
func (o *BTExportModelEdge1782) HasGeometry() bool {
	if o != nil && o.Geometry != nil {
		return true
	}

	return false
}

// SetGeometry gets a reference to the given BTExportModelEdgeGeometry1125 and assigns it to the Geometry field.
func (o *BTExportModelEdge1782) SetGeometry(v BTExportModelEdgeGeometry1125) {
	o.Geometry = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTExportModelEdge1782) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelEdge1782) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTExportModelEdge1782) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTExportModelEdge1782) SetId(v string) {
	o.Id = &v
}

// GetVertices returns the Vertices field value if set, zero value otherwise.
func (o *BTExportModelEdge1782) GetVertices() []string {
	if o == nil || o.Vertices == nil {
		var ret []string
		return ret
	}
	return o.Vertices
}

// GetVerticesOk returns a tuple with the Vertices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTExportModelEdge1782) GetVerticesOk() ([]string, bool) {
	if o == nil || o.Vertices == nil {
		return nil, false
	}
	return o.Vertices, true
}

// HasVertices returns a boolean if a field has been set.
func (o *BTExportModelEdge1782) HasVertices() bool {
	if o != nil && o.Vertices != nil {
		return true
	}

	return false
}

// SetVertices gets a reference to the given []string and assigns it to the Vertices field.
func (o *BTExportModelEdge1782) SetVertices(v []string) {
	o.Vertices = v
}

func (o BTExportModelEdge1782) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Curve != nil {
		toSerialize["curve"] = o.Curve
	}
	if o.Geometry != nil {
		toSerialize["geometry"] = o.Geometry
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Vertices != nil {
		toSerialize["vertices"] = o.Vertices
	}
	return json.Marshal(toSerialize)
}

type NullableBTExportModelEdge1782 struct {
	value *BTExportModelEdge1782
	isSet bool
}

func (v NullableBTExportModelEdge1782) Get() *BTExportModelEdge1782 {
	return v.value
}

func (v *NullableBTExportModelEdge1782) Set(val *BTExportModelEdge1782) {
	v.value = val
	v.isSet = true
}

func (v NullableBTExportModelEdge1782) IsSet() bool {
	return v.isSet
}

func (v *NullableBTExportModelEdge1782) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTExportModelEdge1782(val *BTExportModelEdge1782) *NullableBTExportModelEdge1782 {
	return &NullableBTExportModelEdge1782{value: val, isSet: true}
}

func (v NullableBTExportModelEdge1782) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTExportModelEdge1782) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
