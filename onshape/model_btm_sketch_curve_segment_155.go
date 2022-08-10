/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5959-c304a5ae37eb
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMSketchCurveSegment155 struct for BTMSketchCurveSegment155
type BTMSketchCurveSegment155 struct {
	BtType                              *string             `json:"btType,omitempty"`
	ControlBoxIds                       []string            `json:"controlBoxIds,omitempty"`
	EntityId                            *string             `json:"entityId,omitempty"`
	EntityIdAndReplaceInDependentFields *string             `json:"entityIdAndReplaceInDependentFields,omitempty"`
	ImportMicroversion                  *string             `json:"importMicroversion,omitempty"`
	IsConstruction                      *bool               `json:"isConstruction,omitempty"`
	IsFromEndpointSplineHandle          *bool               `json:"isFromEndpointSplineHandle,omitempty"`
	IsFromSplineControlPolygon          *bool               `json:"isFromSplineControlPolygon,omitempty"`
	IsFromSplineHandle                  *bool               `json:"isFromSplineHandle,omitempty"`
	Namespace                           *string             `json:"namespace,omitempty"`
	NodeId                              *string             `json:"nodeId,omitempty"`
	Parameters                          []BTMParameter1     `json:"parameters,omitempty"`
	CenterId                            *string             `json:"centerId,omitempty"`
	Geometry                            *BTCurveGeometry114 `json:"geometry,omitempty"`
	InternalIds                         []string            `json:"internalIds,omitempty"`
	EndParam                            *float64            `json:"endParam,omitempty"`
	EndPointId                          *string             `json:"endPointId,omitempty"`
	StartParam                          *float64            `json:"startParam,omitempty"`
	StartPointId                        *string             `json:"startPointId,omitempty"`
}

// NewBTMSketchCurveSegment155 instantiates a new BTMSketchCurveSegment155 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMSketchCurveSegment155() *BTMSketchCurveSegment155 {
	this := BTMSketchCurveSegment155{}
	return &this
}

// NewBTMSketchCurveSegment155WithDefaults instantiates a new BTMSketchCurveSegment155 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMSketchCurveSegment155WithDefaults() *BTMSketchCurveSegment155 {
	this := BTMSketchCurveSegment155{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMSketchCurveSegment155) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCurveSegment155) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMSketchCurveSegment155) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMSketchCurveSegment155) SetBtType(v string) {
	o.BtType = &v
}

// GetControlBoxIds returns the ControlBoxIds field value if set, zero value otherwise.
func (o *BTMSketchCurveSegment155) GetControlBoxIds() []string {
	if o == nil || o.ControlBoxIds == nil {
		var ret []string
		return ret
	}
	return o.ControlBoxIds
}

// GetControlBoxIdsOk returns a tuple with the ControlBoxIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCurveSegment155) GetControlBoxIdsOk() ([]string, bool) {
	if o == nil || o.ControlBoxIds == nil {
		return nil, false
	}
	return o.ControlBoxIds, true
}

// HasControlBoxIds returns a boolean if a field has been set.
func (o *BTMSketchCurveSegment155) HasControlBoxIds() bool {
	if o != nil && o.ControlBoxIds != nil {
		return true
	}

	return false
}

// SetControlBoxIds gets a reference to the given []string and assigns it to the ControlBoxIds field.
func (o *BTMSketchCurveSegment155) SetControlBoxIds(v []string) {
	o.ControlBoxIds = v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *BTMSketchCurveSegment155) GetEntityId() string {
	if o == nil || o.EntityId == nil {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCurveSegment155) GetEntityIdOk() (*string, bool) {
	if o == nil || o.EntityId == nil {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *BTMSketchCurveSegment155) HasEntityId() bool {
	if o != nil && o.EntityId != nil {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *BTMSketchCurveSegment155) SetEntityId(v string) {
	o.EntityId = &v
}

// GetEntityIdAndReplaceInDependentFields returns the EntityIdAndReplaceInDependentFields field value if set, zero value otherwise.
func (o *BTMSketchCurveSegment155) GetEntityIdAndReplaceInDependentFields() string {
	if o == nil || o.EntityIdAndReplaceInDependentFields == nil {
		var ret string
		return ret
	}
	return *o.EntityIdAndReplaceInDependentFields
}

// GetEntityIdAndReplaceInDependentFieldsOk returns a tuple with the EntityIdAndReplaceInDependentFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCurveSegment155) GetEntityIdAndReplaceInDependentFieldsOk() (*string, bool) {
	if o == nil || o.EntityIdAndReplaceInDependentFields == nil {
		return nil, false
	}
	return o.EntityIdAndReplaceInDependentFields, true
}

// HasEntityIdAndReplaceInDependentFields returns a boolean if a field has been set.
func (o *BTMSketchCurveSegment155) HasEntityIdAndReplaceInDependentFields() bool {
	if o != nil && o.EntityIdAndReplaceInDependentFields != nil {
		return true
	}

	return false
}

// SetEntityIdAndReplaceInDependentFields gets a reference to the given string and assigns it to the EntityIdAndReplaceInDependentFields field.
func (o *BTMSketchCurveSegment155) SetEntityIdAndReplaceInDependentFields(v string) {
	o.EntityIdAndReplaceInDependentFields = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMSketchCurveSegment155) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCurveSegment155) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMSketchCurveSegment155) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMSketchCurveSegment155) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetIsConstruction returns the IsConstruction field value if set, zero value otherwise.
func (o *BTMSketchCurveSegment155) GetIsConstruction() bool {
	if o == nil || o.IsConstruction == nil {
		var ret bool
		return ret
	}
	return *o.IsConstruction
}

// GetIsConstructionOk returns a tuple with the IsConstruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCurveSegment155) GetIsConstructionOk() (*bool, bool) {
	if o == nil || o.IsConstruction == nil {
		return nil, false
	}
	return o.IsConstruction, true
}

// HasIsConstruction returns a boolean if a field has been set.
func (o *BTMSketchCurveSegment155) HasIsConstruction() bool {
	if o != nil && o.IsConstruction != nil {
		return true
	}

	return false
}

// SetIsConstruction gets a reference to the given bool and assigns it to the IsConstruction field.
func (o *BTMSketchCurveSegment155) SetIsConstruction(v bool) {
	o.IsConstruction = &v
}

// GetIsFromEndpointSplineHandle returns the IsFromEndpointSplineHandle field value if set, zero value otherwise.
func (o *BTMSketchCurveSegment155) GetIsFromEndpointSplineHandle() bool {
	if o == nil || o.IsFromEndpointSplineHandle == nil {
		var ret bool
		return ret
	}
	return *o.IsFromEndpointSplineHandle
}

// GetIsFromEndpointSplineHandleOk returns a tuple with the IsFromEndpointSplineHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCurveSegment155) GetIsFromEndpointSplineHandleOk() (*bool, bool) {
	if o == nil || o.IsFromEndpointSplineHandle == nil {
		return nil, false
	}
	return o.IsFromEndpointSplineHandle, true
}

// HasIsFromEndpointSplineHandle returns a boolean if a field has been set.
func (o *BTMSketchCurveSegment155) HasIsFromEndpointSplineHandle() bool {
	if o != nil && o.IsFromEndpointSplineHandle != nil {
		return true
	}

	return false
}

// SetIsFromEndpointSplineHandle gets a reference to the given bool and assigns it to the IsFromEndpointSplineHandle field.
func (o *BTMSketchCurveSegment155) SetIsFromEndpointSplineHandle(v bool) {
	o.IsFromEndpointSplineHandle = &v
}

// GetIsFromSplineControlPolygon returns the IsFromSplineControlPolygon field value if set, zero value otherwise.
func (o *BTMSketchCurveSegment155) GetIsFromSplineControlPolygon() bool {
	if o == nil || o.IsFromSplineControlPolygon == nil {
		var ret bool
		return ret
	}
	return *o.IsFromSplineControlPolygon
}

// GetIsFromSplineControlPolygonOk returns a tuple with the IsFromSplineControlPolygon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCurveSegment155) GetIsFromSplineControlPolygonOk() (*bool, bool) {
	if o == nil || o.IsFromSplineControlPolygon == nil {
		return nil, false
	}
	return o.IsFromSplineControlPolygon, true
}

// HasIsFromSplineControlPolygon returns a boolean if a field has been set.
func (o *BTMSketchCurveSegment155) HasIsFromSplineControlPolygon() bool {
	if o != nil && o.IsFromSplineControlPolygon != nil {
		return true
	}

	return false
}

// SetIsFromSplineControlPolygon gets a reference to the given bool and assigns it to the IsFromSplineControlPolygon field.
func (o *BTMSketchCurveSegment155) SetIsFromSplineControlPolygon(v bool) {
	o.IsFromSplineControlPolygon = &v
}

// GetIsFromSplineHandle returns the IsFromSplineHandle field value if set, zero value otherwise.
func (o *BTMSketchCurveSegment155) GetIsFromSplineHandle() bool {
	if o == nil || o.IsFromSplineHandle == nil {
		var ret bool
		return ret
	}
	return *o.IsFromSplineHandle
}

// GetIsFromSplineHandleOk returns a tuple with the IsFromSplineHandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCurveSegment155) GetIsFromSplineHandleOk() (*bool, bool) {
	if o == nil || o.IsFromSplineHandle == nil {
		return nil, false
	}
	return o.IsFromSplineHandle, true
}

// HasIsFromSplineHandle returns a boolean if a field has been set.
func (o *BTMSketchCurveSegment155) HasIsFromSplineHandle() bool {
	if o != nil && o.IsFromSplineHandle != nil {
		return true
	}

	return false
}

// SetIsFromSplineHandle gets a reference to the given bool and assigns it to the IsFromSplineHandle field.
func (o *BTMSketchCurveSegment155) SetIsFromSplineHandle(v bool) {
	o.IsFromSplineHandle = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *BTMSketchCurveSegment155) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCurveSegment155) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *BTMSketchCurveSegment155) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *BTMSketchCurveSegment155) SetNamespace(v string) {
	o.Namespace = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMSketchCurveSegment155) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCurveSegment155) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMSketchCurveSegment155) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMSketchCurveSegment155) SetNodeId(v string) {
	o.NodeId = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *BTMSketchCurveSegment155) GetParameters() []BTMParameter1 {
	if o == nil || o.Parameters == nil {
		var ret []BTMParameter1
		return ret
	}
	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCurveSegment155) GetParametersOk() ([]BTMParameter1, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *BTMSketchCurveSegment155) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []BTMParameter1 and assigns it to the Parameters field.
func (o *BTMSketchCurveSegment155) SetParameters(v []BTMParameter1) {
	o.Parameters = v
}

// GetCenterId returns the CenterId field value if set, zero value otherwise.
func (o *BTMSketchCurveSegment155) GetCenterId() string {
	if o == nil || o.CenterId == nil {
		var ret string
		return ret
	}
	return *o.CenterId
}

// GetCenterIdOk returns a tuple with the CenterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCurveSegment155) GetCenterIdOk() (*string, bool) {
	if o == nil || o.CenterId == nil {
		return nil, false
	}
	return o.CenterId, true
}

// HasCenterId returns a boolean if a field has been set.
func (o *BTMSketchCurveSegment155) HasCenterId() bool {
	if o != nil && o.CenterId != nil {
		return true
	}

	return false
}

// SetCenterId gets a reference to the given string and assigns it to the CenterId field.
func (o *BTMSketchCurveSegment155) SetCenterId(v string) {
	o.CenterId = &v
}

// GetGeometry returns the Geometry field value if set, zero value otherwise.
func (o *BTMSketchCurveSegment155) GetGeometry() BTCurveGeometry114 {
	if o == nil || o.Geometry == nil {
		var ret BTCurveGeometry114
		return ret
	}
	return *o.Geometry
}

// GetGeometryOk returns a tuple with the Geometry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCurveSegment155) GetGeometryOk() (*BTCurveGeometry114, bool) {
	if o == nil || o.Geometry == nil {
		return nil, false
	}
	return o.Geometry, true
}

// HasGeometry returns a boolean if a field has been set.
func (o *BTMSketchCurveSegment155) HasGeometry() bool {
	if o != nil && o.Geometry != nil {
		return true
	}

	return false
}

// SetGeometry gets a reference to the given BTCurveGeometry114 and assigns it to the Geometry field.
func (o *BTMSketchCurveSegment155) SetGeometry(v BTCurveGeometry114) {
	o.Geometry = &v
}

// GetInternalIds returns the InternalIds field value if set, zero value otherwise.
func (o *BTMSketchCurveSegment155) GetInternalIds() []string {
	if o == nil || o.InternalIds == nil {
		var ret []string
		return ret
	}
	return o.InternalIds
}

// GetInternalIdsOk returns a tuple with the InternalIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCurveSegment155) GetInternalIdsOk() ([]string, bool) {
	if o == nil || o.InternalIds == nil {
		return nil, false
	}
	return o.InternalIds, true
}

// HasInternalIds returns a boolean if a field has been set.
func (o *BTMSketchCurveSegment155) HasInternalIds() bool {
	if o != nil && o.InternalIds != nil {
		return true
	}

	return false
}

// SetInternalIds gets a reference to the given []string and assigns it to the InternalIds field.
func (o *BTMSketchCurveSegment155) SetInternalIds(v []string) {
	o.InternalIds = v
}

// GetEndParam returns the EndParam field value if set, zero value otherwise.
func (o *BTMSketchCurveSegment155) GetEndParam() float64 {
	if o == nil || o.EndParam == nil {
		var ret float64
		return ret
	}
	return *o.EndParam
}

// GetEndParamOk returns a tuple with the EndParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCurveSegment155) GetEndParamOk() (*float64, bool) {
	if o == nil || o.EndParam == nil {
		return nil, false
	}
	return o.EndParam, true
}

// HasEndParam returns a boolean if a field has been set.
func (o *BTMSketchCurveSegment155) HasEndParam() bool {
	if o != nil && o.EndParam != nil {
		return true
	}

	return false
}

// SetEndParam gets a reference to the given float64 and assigns it to the EndParam field.
func (o *BTMSketchCurveSegment155) SetEndParam(v float64) {
	o.EndParam = &v
}

// GetEndPointId returns the EndPointId field value if set, zero value otherwise.
func (o *BTMSketchCurveSegment155) GetEndPointId() string {
	if o == nil || o.EndPointId == nil {
		var ret string
		return ret
	}
	return *o.EndPointId
}

// GetEndPointIdOk returns a tuple with the EndPointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCurveSegment155) GetEndPointIdOk() (*string, bool) {
	if o == nil || o.EndPointId == nil {
		return nil, false
	}
	return o.EndPointId, true
}

// HasEndPointId returns a boolean if a field has been set.
func (o *BTMSketchCurveSegment155) HasEndPointId() bool {
	if o != nil && o.EndPointId != nil {
		return true
	}

	return false
}

// SetEndPointId gets a reference to the given string and assigns it to the EndPointId field.
func (o *BTMSketchCurveSegment155) SetEndPointId(v string) {
	o.EndPointId = &v
}

// GetStartParam returns the StartParam field value if set, zero value otherwise.
func (o *BTMSketchCurveSegment155) GetStartParam() float64 {
	if o == nil || o.StartParam == nil {
		var ret float64
		return ret
	}
	return *o.StartParam
}

// GetStartParamOk returns a tuple with the StartParam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCurveSegment155) GetStartParamOk() (*float64, bool) {
	if o == nil || o.StartParam == nil {
		return nil, false
	}
	return o.StartParam, true
}

// HasStartParam returns a boolean if a field has been set.
func (o *BTMSketchCurveSegment155) HasStartParam() bool {
	if o != nil && o.StartParam != nil {
		return true
	}

	return false
}

// SetStartParam gets a reference to the given float64 and assigns it to the StartParam field.
func (o *BTMSketchCurveSegment155) SetStartParam(v float64) {
	o.StartParam = &v
}

// GetStartPointId returns the StartPointId field value if set, zero value otherwise.
func (o *BTMSketchCurveSegment155) GetStartPointId() string {
	if o == nil || o.StartPointId == nil {
		var ret string
		return ret
	}
	return *o.StartPointId
}

// GetStartPointIdOk returns a tuple with the StartPointId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMSketchCurveSegment155) GetStartPointIdOk() (*string, bool) {
	if o == nil || o.StartPointId == nil {
		return nil, false
	}
	return o.StartPointId, true
}

// HasStartPointId returns a boolean if a field has been set.
func (o *BTMSketchCurveSegment155) HasStartPointId() bool {
	if o != nil && o.StartPointId != nil {
		return true
	}

	return false
}

// SetStartPointId gets a reference to the given string and assigns it to the StartPointId field.
func (o *BTMSketchCurveSegment155) SetStartPointId(v string) {
	o.StartPointId = &v
}

func (o BTMSketchCurveSegment155) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.ControlBoxIds != nil {
		toSerialize["controlBoxIds"] = o.ControlBoxIds
	}
	if o.EntityId != nil {
		toSerialize["entityId"] = o.EntityId
	}
	if o.EntityIdAndReplaceInDependentFields != nil {
		toSerialize["entityIdAndReplaceInDependentFields"] = o.EntityIdAndReplaceInDependentFields
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.IsConstruction != nil {
		toSerialize["isConstruction"] = o.IsConstruction
	}
	if o.IsFromEndpointSplineHandle != nil {
		toSerialize["isFromEndpointSplineHandle"] = o.IsFromEndpointSplineHandle
	}
	if o.IsFromSplineControlPolygon != nil {
		toSerialize["isFromSplineControlPolygon"] = o.IsFromSplineControlPolygon
	}
	if o.IsFromSplineHandle != nil {
		toSerialize["isFromSplineHandle"] = o.IsFromSplineHandle
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.CenterId != nil {
		toSerialize["centerId"] = o.CenterId
	}
	if o.Geometry != nil {
		toSerialize["geometry"] = o.Geometry
	}
	if o.InternalIds != nil {
		toSerialize["internalIds"] = o.InternalIds
	}
	if o.EndParam != nil {
		toSerialize["endParam"] = o.EndParam
	}
	if o.EndPointId != nil {
		toSerialize["endPointId"] = o.EndPointId
	}
	if o.StartParam != nil {
		toSerialize["startParam"] = o.StartParam
	}
	if o.StartPointId != nil {
		toSerialize["startPointId"] = o.StartPointId
	}
	return json.Marshal(toSerialize)
}

type NullableBTMSketchCurveSegment155 struct {
	value *BTMSketchCurveSegment155
	isSet bool
}

func (v NullableBTMSketchCurveSegment155) Get() *BTMSketchCurveSegment155 {
	return v.value
}

func (v *NullableBTMSketchCurveSegment155) Set(val *BTMSketchCurveSegment155) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMSketchCurveSegment155) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMSketchCurveSegment155) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMSketchCurveSegment155(val *BTMSketchCurveSegment155) *NullableBTMSketchCurveSegment155 {
	return &NullableBTMSketchCurveSegment155{value: val, isSet: true}
}

func (v NullableBTMSketchCurveSegment155) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMSketchCurveSegment155) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
