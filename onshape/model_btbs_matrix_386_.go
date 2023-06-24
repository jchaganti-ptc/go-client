/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.165.17983-3c4f6e999748
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTBSMatrix386 struct for BTBSMatrix386
type BTBSMatrix386 struct {
	BtType                        *string        `json:"btType,omitempty"`
	FullTransformation            []float64      `json:"fullTransformation,omitempty"`
	IsNormalized                  *bool          `json:"isNormalized,omitempty"`
	M00                           *float64       `json:"m00,omitempty"`
	M01                           *float64       `json:"m01,omitempty"`
	M02                           *float64       `json:"m02,omitempty"`
	M03                           *float64       `json:"m03,omitempty"`
	M10                           *float64       `json:"m10,omitempty"`
	M11                           *float64       `json:"m11,omitempty"`
	M12                           *float64       `json:"m12,omitempty"`
	M13                           *float64       `json:"m13,omitempty"`
	M20                           *float64       `json:"m20,omitempty"`
	M21                           *float64       `json:"m21,omitempty"`
	M22                           *float64       `json:"m22,omitempty"`
	M23                           *float64       `json:"m23,omitempty"`
	Rigid                         *bool          `json:"rigid,omitempty"`
	RigidWithinTransformTolerance *bool          `json:"rigidWithinTransformTolerance,omitempty"`
	Translation                   *BTVector3d389 `json:"translation,omitempty"`
}

// NewBTBSMatrix386 instantiates a new BTBSMatrix386 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTBSMatrix386() *BTBSMatrix386 {
	this := BTBSMatrix386{}
	return &this
}

// NewBTBSMatrix386WithDefaults instantiates a new BTBSMatrix386 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTBSMatrix386WithDefaults() *BTBSMatrix386 {
	this := BTBSMatrix386{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTBSMatrix386) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBSMatrix386) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTBSMatrix386) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTBSMatrix386) SetBtType(v string) {
	o.BtType = &v
}

// GetFullTransformation returns the FullTransformation field value if set, zero value otherwise.
func (o *BTBSMatrix386) GetFullTransformation() []float64 {
	if o == nil || o.FullTransformation == nil {
		var ret []float64
		return ret
	}
	return o.FullTransformation
}

// GetFullTransformationOk returns a tuple with the FullTransformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBSMatrix386) GetFullTransformationOk() ([]float64, bool) {
	if o == nil || o.FullTransformation == nil {
		return nil, false
	}
	return o.FullTransformation, true
}

// HasFullTransformation returns a boolean if a field has been set.
func (o *BTBSMatrix386) HasFullTransformation() bool {
	if o != nil && o.FullTransformation != nil {
		return true
	}

	return false
}

// SetFullTransformation gets a reference to the given []float64 and assigns it to the FullTransformation field.
func (o *BTBSMatrix386) SetFullTransformation(v []float64) {
	o.FullTransformation = v
}

// GetIsNormalized returns the IsNormalized field value if set, zero value otherwise.
func (o *BTBSMatrix386) GetIsNormalized() bool {
	if o == nil || o.IsNormalized == nil {
		var ret bool
		return ret
	}
	return *o.IsNormalized
}

// GetIsNormalizedOk returns a tuple with the IsNormalized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBSMatrix386) GetIsNormalizedOk() (*bool, bool) {
	if o == nil || o.IsNormalized == nil {
		return nil, false
	}
	return o.IsNormalized, true
}

// HasIsNormalized returns a boolean if a field has been set.
func (o *BTBSMatrix386) HasIsNormalized() bool {
	if o != nil && o.IsNormalized != nil {
		return true
	}

	return false
}

// SetIsNormalized gets a reference to the given bool and assigns it to the IsNormalized field.
func (o *BTBSMatrix386) SetIsNormalized(v bool) {
	o.IsNormalized = &v
}

// GetM00 returns the M00 field value if set, zero value otherwise.
func (o *BTBSMatrix386) GetM00() float64 {
	if o == nil || o.M00 == nil {
		var ret float64
		return ret
	}
	return *o.M00
}

// GetM00Ok returns a tuple with the M00 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBSMatrix386) GetM00Ok() (*float64, bool) {
	if o == nil || o.M00 == nil {
		return nil, false
	}
	return o.M00, true
}

// HasM00 returns a boolean if a field has been set.
func (o *BTBSMatrix386) HasM00() bool {
	if o != nil && o.M00 != nil {
		return true
	}

	return false
}

// SetM00 gets a reference to the given float64 and assigns it to the M00 field.
func (o *BTBSMatrix386) SetM00(v float64) {
	o.M00 = &v
}

// GetM01 returns the M01 field value if set, zero value otherwise.
func (o *BTBSMatrix386) GetM01() float64 {
	if o == nil || o.M01 == nil {
		var ret float64
		return ret
	}
	return *o.M01
}

// GetM01Ok returns a tuple with the M01 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBSMatrix386) GetM01Ok() (*float64, bool) {
	if o == nil || o.M01 == nil {
		return nil, false
	}
	return o.M01, true
}

// HasM01 returns a boolean if a field has been set.
func (o *BTBSMatrix386) HasM01() bool {
	if o != nil && o.M01 != nil {
		return true
	}

	return false
}

// SetM01 gets a reference to the given float64 and assigns it to the M01 field.
func (o *BTBSMatrix386) SetM01(v float64) {
	o.M01 = &v
}

// GetM02 returns the M02 field value if set, zero value otherwise.
func (o *BTBSMatrix386) GetM02() float64 {
	if o == nil || o.M02 == nil {
		var ret float64
		return ret
	}
	return *o.M02
}

// GetM02Ok returns a tuple with the M02 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBSMatrix386) GetM02Ok() (*float64, bool) {
	if o == nil || o.M02 == nil {
		return nil, false
	}
	return o.M02, true
}

// HasM02 returns a boolean if a field has been set.
func (o *BTBSMatrix386) HasM02() bool {
	if o != nil && o.M02 != nil {
		return true
	}

	return false
}

// SetM02 gets a reference to the given float64 and assigns it to the M02 field.
func (o *BTBSMatrix386) SetM02(v float64) {
	o.M02 = &v
}

// GetM03 returns the M03 field value if set, zero value otherwise.
func (o *BTBSMatrix386) GetM03() float64 {
	if o == nil || o.M03 == nil {
		var ret float64
		return ret
	}
	return *o.M03
}

// GetM03Ok returns a tuple with the M03 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBSMatrix386) GetM03Ok() (*float64, bool) {
	if o == nil || o.M03 == nil {
		return nil, false
	}
	return o.M03, true
}

// HasM03 returns a boolean if a field has been set.
func (o *BTBSMatrix386) HasM03() bool {
	if o != nil && o.M03 != nil {
		return true
	}

	return false
}

// SetM03 gets a reference to the given float64 and assigns it to the M03 field.
func (o *BTBSMatrix386) SetM03(v float64) {
	o.M03 = &v
}

// GetM10 returns the M10 field value if set, zero value otherwise.
func (o *BTBSMatrix386) GetM10() float64 {
	if o == nil || o.M10 == nil {
		var ret float64
		return ret
	}
	return *o.M10
}

// GetM10Ok returns a tuple with the M10 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBSMatrix386) GetM10Ok() (*float64, bool) {
	if o == nil || o.M10 == nil {
		return nil, false
	}
	return o.M10, true
}

// HasM10 returns a boolean if a field has been set.
func (o *BTBSMatrix386) HasM10() bool {
	if o != nil && o.M10 != nil {
		return true
	}

	return false
}

// SetM10 gets a reference to the given float64 and assigns it to the M10 field.
func (o *BTBSMatrix386) SetM10(v float64) {
	o.M10 = &v
}

// GetM11 returns the M11 field value if set, zero value otherwise.
func (o *BTBSMatrix386) GetM11() float64 {
	if o == nil || o.M11 == nil {
		var ret float64
		return ret
	}
	return *o.M11
}

// GetM11Ok returns a tuple with the M11 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBSMatrix386) GetM11Ok() (*float64, bool) {
	if o == nil || o.M11 == nil {
		return nil, false
	}
	return o.M11, true
}

// HasM11 returns a boolean if a field has been set.
func (o *BTBSMatrix386) HasM11() bool {
	if o != nil && o.M11 != nil {
		return true
	}

	return false
}

// SetM11 gets a reference to the given float64 and assigns it to the M11 field.
func (o *BTBSMatrix386) SetM11(v float64) {
	o.M11 = &v
}

// GetM12 returns the M12 field value if set, zero value otherwise.
func (o *BTBSMatrix386) GetM12() float64 {
	if o == nil || o.M12 == nil {
		var ret float64
		return ret
	}
	return *o.M12
}

// GetM12Ok returns a tuple with the M12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBSMatrix386) GetM12Ok() (*float64, bool) {
	if o == nil || o.M12 == nil {
		return nil, false
	}
	return o.M12, true
}

// HasM12 returns a boolean if a field has been set.
func (o *BTBSMatrix386) HasM12() bool {
	if o != nil && o.M12 != nil {
		return true
	}

	return false
}

// SetM12 gets a reference to the given float64 and assigns it to the M12 field.
func (o *BTBSMatrix386) SetM12(v float64) {
	o.M12 = &v
}

// GetM13 returns the M13 field value if set, zero value otherwise.
func (o *BTBSMatrix386) GetM13() float64 {
	if o == nil || o.M13 == nil {
		var ret float64
		return ret
	}
	return *o.M13
}

// GetM13Ok returns a tuple with the M13 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBSMatrix386) GetM13Ok() (*float64, bool) {
	if o == nil || o.M13 == nil {
		return nil, false
	}
	return o.M13, true
}

// HasM13 returns a boolean if a field has been set.
func (o *BTBSMatrix386) HasM13() bool {
	if o != nil && o.M13 != nil {
		return true
	}

	return false
}

// SetM13 gets a reference to the given float64 and assigns it to the M13 field.
func (o *BTBSMatrix386) SetM13(v float64) {
	o.M13 = &v
}

// GetM20 returns the M20 field value if set, zero value otherwise.
func (o *BTBSMatrix386) GetM20() float64 {
	if o == nil || o.M20 == nil {
		var ret float64
		return ret
	}
	return *o.M20
}

// GetM20Ok returns a tuple with the M20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBSMatrix386) GetM20Ok() (*float64, bool) {
	if o == nil || o.M20 == nil {
		return nil, false
	}
	return o.M20, true
}

// HasM20 returns a boolean if a field has been set.
func (o *BTBSMatrix386) HasM20() bool {
	if o != nil && o.M20 != nil {
		return true
	}

	return false
}

// SetM20 gets a reference to the given float64 and assigns it to the M20 field.
func (o *BTBSMatrix386) SetM20(v float64) {
	o.M20 = &v
}

// GetM21 returns the M21 field value if set, zero value otherwise.
func (o *BTBSMatrix386) GetM21() float64 {
	if o == nil || o.M21 == nil {
		var ret float64
		return ret
	}
	return *o.M21
}

// GetM21Ok returns a tuple with the M21 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBSMatrix386) GetM21Ok() (*float64, bool) {
	if o == nil || o.M21 == nil {
		return nil, false
	}
	return o.M21, true
}

// HasM21 returns a boolean if a field has been set.
func (o *BTBSMatrix386) HasM21() bool {
	if o != nil && o.M21 != nil {
		return true
	}

	return false
}

// SetM21 gets a reference to the given float64 and assigns it to the M21 field.
func (o *BTBSMatrix386) SetM21(v float64) {
	o.M21 = &v
}

// GetM22 returns the M22 field value if set, zero value otherwise.
func (o *BTBSMatrix386) GetM22() float64 {
	if o == nil || o.M22 == nil {
		var ret float64
		return ret
	}
	return *o.M22
}

// GetM22Ok returns a tuple with the M22 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBSMatrix386) GetM22Ok() (*float64, bool) {
	if o == nil || o.M22 == nil {
		return nil, false
	}
	return o.M22, true
}

// HasM22 returns a boolean if a field has been set.
func (o *BTBSMatrix386) HasM22() bool {
	if o != nil && o.M22 != nil {
		return true
	}

	return false
}

// SetM22 gets a reference to the given float64 and assigns it to the M22 field.
func (o *BTBSMatrix386) SetM22(v float64) {
	o.M22 = &v
}

// GetM23 returns the M23 field value if set, zero value otherwise.
func (o *BTBSMatrix386) GetM23() float64 {
	if o == nil || o.M23 == nil {
		var ret float64
		return ret
	}
	return *o.M23
}

// GetM23Ok returns a tuple with the M23 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBSMatrix386) GetM23Ok() (*float64, bool) {
	if o == nil || o.M23 == nil {
		return nil, false
	}
	return o.M23, true
}

// HasM23 returns a boolean if a field has been set.
func (o *BTBSMatrix386) HasM23() bool {
	if o != nil && o.M23 != nil {
		return true
	}

	return false
}

// SetM23 gets a reference to the given float64 and assigns it to the M23 field.
func (o *BTBSMatrix386) SetM23(v float64) {
	o.M23 = &v
}

// GetRigid returns the Rigid field value if set, zero value otherwise.
func (o *BTBSMatrix386) GetRigid() bool {
	if o == nil || o.Rigid == nil {
		var ret bool
		return ret
	}
	return *o.Rigid
}

// GetRigidOk returns a tuple with the Rigid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBSMatrix386) GetRigidOk() (*bool, bool) {
	if o == nil || o.Rigid == nil {
		return nil, false
	}
	return o.Rigid, true
}

// HasRigid returns a boolean if a field has been set.
func (o *BTBSMatrix386) HasRigid() bool {
	if o != nil && o.Rigid != nil {
		return true
	}

	return false
}

// SetRigid gets a reference to the given bool and assigns it to the Rigid field.
func (o *BTBSMatrix386) SetRigid(v bool) {
	o.Rigid = &v
}

// GetRigidWithinTransformTolerance returns the RigidWithinTransformTolerance field value if set, zero value otherwise.
func (o *BTBSMatrix386) GetRigidWithinTransformTolerance() bool {
	if o == nil || o.RigidWithinTransformTolerance == nil {
		var ret bool
		return ret
	}
	return *o.RigidWithinTransformTolerance
}

// GetRigidWithinTransformToleranceOk returns a tuple with the RigidWithinTransformTolerance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBSMatrix386) GetRigidWithinTransformToleranceOk() (*bool, bool) {
	if o == nil || o.RigidWithinTransformTolerance == nil {
		return nil, false
	}
	return o.RigidWithinTransformTolerance, true
}

// HasRigidWithinTransformTolerance returns a boolean if a field has been set.
func (o *BTBSMatrix386) HasRigidWithinTransformTolerance() bool {
	if o != nil && o.RigidWithinTransformTolerance != nil {
		return true
	}

	return false
}

// SetRigidWithinTransformTolerance gets a reference to the given bool and assigns it to the RigidWithinTransformTolerance field.
func (o *BTBSMatrix386) SetRigidWithinTransformTolerance(v bool) {
	o.RigidWithinTransformTolerance = &v
}

// GetTranslation returns the Translation field value if set, zero value otherwise.
func (o *BTBSMatrix386) GetTranslation() BTVector3d389 {
	if o == nil || o.Translation == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Translation
}

// GetTranslationOk returns a tuple with the Translation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTBSMatrix386) GetTranslationOk() (*BTVector3d389, bool) {
	if o == nil || o.Translation == nil {
		return nil, false
	}
	return o.Translation, true
}

// HasTranslation returns a boolean if a field has been set.
func (o *BTBSMatrix386) HasTranslation() bool {
	if o != nil && o.Translation != nil {
		return true
	}

	return false
}

// SetTranslation gets a reference to the given BTVector3d389 and assigns it to the Translation field.
func (o *BTBSMatrix386) SetTranslation(v BTVector3d389) {
	o.Translation = &v
}

func (o BTBSMatrix386) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.FullTransformation != nil {
		toSerialize["fullTransformation"] = o.FullTransformation
	}
	if o.IsNormalized != nil {
		toSerialize["isNormalized"] = o.IsNormalized
	}
	if o.M00 != nil {
		toSerialize["m00"] = o.M00
	}
	if o.M01 != nil {
		toSerialize["m01"] = o.M01
	}
	if o.M02 != nil {
		toSerialize["m02"] = o.M02
	}
	if o.M03 != nil {
		toSerialize["m03"] = o.M03
	}
	if o.M10 != nil {
		toSerialize["m10"] = o.M10
	}
	if o.M11 != nil {
		toSerialize["m11"] = o.M11
	}
	if o.M12 != nil {
		toSerialize["m12"] = o.M12
	}
	if o.M13 != nil {
		toSerialize["m13"] = o.M13
	}
	if o.M20 != nil {
		toSerialize["m20"] = o.M20
	}
	if o.M21 != nil {
		toSerialize["m21"] = o.M21
	}
	if o.M22 != nil {
		toSerialize["m22"] = o.M22
	}
	if o.M23 != nil {
		toSerialize["m23"] = o.M23
	}
	if o.Rigid != nil {
		toSerialize["rigid"] = o.Rigid
	}
	if o.RigidWithinTransformTolerance != nil {
		toSerialize["rigidWithinTransformTolerance"] = o.RigidWithinTransformTolerance
	}
	if o.Translation != nil {
		toSerialize["translation"] = o.Translation
	}
	return json.Marshal(toSerialize)
}

type NullableBTBSMatrix386 struct {
	value *BTBSMatrix386
	isSet bool
}

func (v NullableBTBSMatrix386) Get() *BTBSMatrix386 {
	return v.value
}

func (v *NullableBTBSMatrix386) Set(val *BTBSMatrix386) {
	v.value = val
	v.isSet = true
}

func (v NullableBTBSMatrix386) IsSet() bool {
	return v.isSet
}

func (v *NullableBTBSMatrix386) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTBSMatrix386(val *BTBSMatrix386) *NullableBTBSMatrix386 {
	return &NullableBTBSMatrix386{value: val, isSet: true}
}

func (v NullableBTBSMatrix386) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTBSMatrix386) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
