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
	"fmt"
)

// BTMIndividualQuery138 - struct for BTMIndividualQuery138
type BTMIndividualQuery138 struct {
	implBTMIndividualQuery138 interface{}
}

// BTMIndividualCreatedByQuery137AsBTMIndividualQuery138 is a convenience function that returns BTMIndividualCreatedByQuery137 wrapped in BTMIndividualQuery138
func (o *BTMIndividualCreatedByQuery137) AsBTMIndividualQuery138() *BTMIndividualQuery138 {
	return &BTMIndividualQuery138{o}
}

// BTMIndividualSketchRegionQuery140AsBTMIndividualQuery138 is a convenience function that returns BTMIndividualSketchRegionQuery140 wrapped in BTMIndividualQuery138
func (o *BTMIndividualSketchRegionQuery140) AsBTMIndividualQuery138() *BTMIndividualQuery138 {
	return &BTMIndividualQuery138{o}
}

// BTMIndividualCoEdgeQuery1332AsBTMIndividualQuery138 is a convenience function that returns BTMIndividualCoEdgeQuery1332 wrapped in BTMIndividualQuery138
func (o *BTMIndividualCoEdgeQuery1332) AsBTMIndividualQuery138() *BTMIndividualQuery138 {
	return &BTMIndividualQuery138{o}
}

// BTMIndividualSketchUniqueVerticesQuery1472AsBTMIndividualQuery138 is a convenience function that returns BTMIndividualSketchUniqueVerticesQuery1472 wrapped in BTMIndividualQuery138
func (o *BTMIndividualSketchUniqueVerticesQuery1472) AsBTMIndividualQuery138() *BTMIndividualQuery138 {
	return &BTMIndividualQuery138{o}
}

// NewBTMIndividualQuery138 instantiates a new BTMIndividualQuery138 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMIndividualQuery138() *BTMIndividualQuery138 {
	this := BTMIndividualQuery138{Newbase_BTMIndividualQuery138()}
	return &this
}

// NewBTMIndividualQuery138WithDefaults instantiates a new BTMIndividualQuery138 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMIndividualQuery138WithDefaults() *BTMIndividualQuery138 {
	this := BTMIndividualQuery138{Newbase_BTMIndividualQuery138WithDefaults()}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMIndividualQuery138) GetBtType() string {
	type getResult interface {
		GetBtType() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtType()
	} else {
		var de string
		return de
	}
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQuery138) GetBtTypeOk() (*string, bool) {
	type getResult interface {
		GetBtTypeOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetBtTypeOk()
	} else {
		return nil, false
	}
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMIndividualQuery138) HasBtType() bool {
	type getResult interface {
		HasBtType() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasBtType()
	} else {
		return false
	}
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMIndividualQuery138) SetBtType(v string) {
	type getResult interface {
		SetBtType(v string)
	}

	o.GetActualInstance().(getResult).SetBtType(v)
}

// GetDeterministicIdList returns the DeterministicIdList field value if set, zero value otherwise.
func (o *BTMIndividualQuery138) GetDeterministicIdList() BTMIndividualQueryBase139 {
	type getResult interface {
		GetDeterministicIdList() BTMIndividualQueryBase139
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDeterministicIdList()
	} else {
		var de BTMIndividualQueryBase139
		return de
	}
}

// GetDeterministicIdListOk returns a tuple with the DeterministicIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQuery138) GetDeterministicIdListOk() (*BTMIndividualQueryBase139, bool) {
	type getResult interface {
		GetDeterministicIdListOk() (*BTMIndividualQueryBase139, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDeterministicIdListOk()
	} else {
		return nil, false
	}
}

// HasDeterministicIdList returns a boolean if a field has been set.
func (o *BTMIndividualQuery138) HasDeterministicIdList() bool {
	type getResult interface {
		HasDeterministicIdList() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDeterministicIdList()
	} else {
		return false
	}
}

// SetDeterministicIdList gets a reference to the given BTMIndividualQueryBase139 and assigns it to the DeterministicIdList field.
func (o *BTMIndividualQuery138) SetDeterministicIdList(v BTMIndividualQueryBase139) {
	type getResult interface {
		SetDeterministicIdList(v BTMIndividualQueryBase139)
	}

	o.GetActualInstance().(getResult).SetDeterministicIdList(v)
}

// GetDeterministicIds returns the DeterministicIds field value if set, zero value otherwise.
func (o *BTMIndividualQuery138) GetDeterministicIds() []string {
	type getResult interface {
		GetDeterministicIds() []string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDeterministicIds()
	} else {
		var de []string
		return de
	}
}

// GetDeterministicIdsOk returns a tuple with the DeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQuery138) GetDeterministicIdsOk() ([]string, bool) {
	type getResult interface {
		GetDeterministicIdsOk() ([]string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetDeterministicIdsOk()
	} else {
		return nil, false
	}
}

// HasDeterministicIds returns a boolean if a field has been set.
func (o *BTMIndividualQuery138) HasDeterministicIds() bool {
	type getResult interface {
		HasDeterministicIds() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasDeterministicIds()
	} else {
		return false
	}
}

// SetDeterministicIds gets a reference to the given []string and assigns it to the DeterministicIds field.
func (o *BTMIndividualQuery138) SetDeterministicIds(v []string) {
	type getResult interface {
		SetDeterministicIds(v []string)
	}

	o.GetActualInstance().(getResult).SetDeterministicIds(v)
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMIndividualQuery138) GetImportMicroversion() string {
	type getResult interface {
		GetImportMicroversion() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetImportMicroversion()
	} else {
		var de string
		return de
	}
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQuery138) GetImportMicroversionOk() (*string, bool) {
	type getResult interface {
		GetImportMicroversionOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetImportMicroversionOk()
	} else {
		return nil, false
	}
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMIndividualQuery138) HasImportMicroversion() bool {
	type getResult interface {
		HasImportMicroversion() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasImportMicroversion()
	} else {
		return false
	}
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMIndividualQuery138) SetImportMicroversion(v string) {
	type getResult interface {
		SetImportMicroversion(v string)
	}

	o.GetActualInstance().(getResult).SetImportMicroversion(v)
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMIndividualQuery138) GetNodeId() string {
	type getResult interface {
		GetNodeId() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeId()
	} else {
		var de string
		return de
	}
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQuery138) GetNodeIdOk() (*string, bool) {
	type getResult interface {
		GetNodeIdOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetNodeIdOk()
	} else {
		return nil, false
	}
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMIndividualQuery138) HasNodeId() bool {
	type getResult interface {
		HasNodeId() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasNodeId()
	} else {
		return false
	}
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMIndividualQuery138) SetNodeId(v string) {
	type getResult interface {
		SetNodeId(v string)
	}

	o.GetActualInstance().(getResult).SetNodeId(v)
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *BTMIndividualQuery138) GetQuery() BTMIndividualQueryBase139 {
	type getResult interface {
		GetQuery() BTMIndividualQueryBase139
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetQuery()
	} else {
		var de BTMIndividualQueryBase139
		return de
	}
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQuery138) GetQueryOk() (*BTMIndividualQueryBase139, bool) {
	type getResult interface {
		GetQueryOk() (*BTMIndividualQueryBase139, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetQueryOk()
	} else {
		return nil, false
	}
}

// HasQuery returns a boolean if a field has been set.
func (o *BTMIndividualQuery138) HasQuery() bool {
	type getResult interface {
		HasQuery() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasQuery()
	} else {
		return false
	}
}

// SetQuery gets a reference to the given BTMIndividualQueryBase139 and assigns it to the Query field.
func (o *BTMIndividualQuery138) SetQuery(v BTMIndividualQueryBase139) {
	type getResult interface {
		SetQuery(v BTMIndividualQueryBase139)
	}

	o.GetActualInstance().(getResult).SetQuery(v)
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *BTMIndividualQuery138) GetQueryString() string {
	type getResult interface {
		GetQueryString() string
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetQueryString()
	} else {
		var de string
		return de
	}
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQuery138) GetQueryStringOk() (*string, bool) {
	type getResult interface {
		GetQueryStringOk() (*string, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetQueryStringOk()
	} else {
		return nil, false
	}
}

// HasQueryString returns a boolean if a field has been set.
func (o *BTMIndividualQuery138) HasQueryString() bool {
	type getResult interface {
		HasQueryString() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasQueryString()
	} else {
		return false
	}
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *BTMIndividualQuery138) SetQueryString(v string) {
	type getResult interface {
		SetQueryString(v string)
	}

	o.GetActualInstance().(getResult).SetQueryString(v)
}

// GetPersistentQuery returns the PersistentQuery field value if set, zero value otherwise.
func (o *BTMIndividualQuery138) GetPersistentQuery() BTPStatement269 {
	type getResult interface {
		GetPersistentQuery() BTPStatement269
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetPersistentQuery()
	} else {
		var de BTPStatement269
		return de
	}
}

// GetPersistentQueryOk returns a tuple with the PersistentQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQuery138) GetPersistentQueryOk() (*BTPStatement269, bool) {
	type getResult interface {
		GetPersistentQueryOk() (*BTPStatement269, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetPersistentQueryOk()
	} else {
		return nil, false
	}
}

// HasPersistentQuery returns a boolean if a field has been set.
func (o *BTMIndividualQuery138) HasPersistentQuery() bool {
	type getResult interface {
		HasPersistentQuery() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasPersistentQuery()
	} else {
		return false
	}
}

// SetPersistentQuery gets a reference to the given BTPStatement269 and assigns it to the PersistentQuery field.
func (o *BTMIndividualQuery138) SetPersistentQuery(v BTPStatement269) {
	type getResult interface {
		SetPersistentQuery(v BTPStatement269)
	}

	o.GetActualInstance().(getResult).SetPersistentQuery(v)
}

// GetQueryStatement returns the QueryStatement field value if set, zero value otherwise.
func (o *BTMIndividualQuery138) GetQueryStatement() BTPStatement269 {
	type getResult interface {
		GetQueryStatement() BTPStatement269
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetQueryStatement()
	} else {
		var de BTPStatement269
		return de
	}
}

// GetQueryStatementOk returns a tuple with the QueryStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQuery138) GetQueryStatementOk() (*BTPStatement269, bool) {
	type getResult interface {
		GetQueryStatementOk() (*BTPStatement269, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetQueryStatementOk()
	} else {
		return nil, false
	}
}

// HasQueryStatement returns a boolean if a field has been set.
func (o *BTMIndividualQuery138) HasQueryStatement() bool {
	type getResult interface {
		HasQueryStatement() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasQueryStatement()
	} else {
		return false
	}
}

// SetQueryStatement gets a reference to the given BTPStatement269 and assigns it to the QueryStatement field.
func (o *BTMIndividualQuery138) SetQueryStatement(v BTPStatement269) {
	type getResult interface {
		SetQueryStatement(v BTPStatement269)
	}

	o.GetActualInstance().(getResult).SetQueryStatement(v)
}

// GetVariableName returns the VariableName field value if set, zero value otherwise.
func (o *BTMIndividualQuery138) GetVariableName() BTMIndividualQuery138 {
	type getResult interface {
		GetVariableName() BTMIndividualQuery138
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetVariableName()
	} else {
		var de BTMIndividualQuery138
		return de
	}
}

// GetVariableNameOk returns a tuple with the VariableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMIndividualQuery138) GetVariableNameOk() (*BTMIndividualQuery138, bool) {
	type getResult interface {
		GetVariableNameOk() (*BTMIndividualQuery138, bool)
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.GetVariableNameOk()
	} else {
		return nil, false
	}
}

// HasVariableName returns a boolean if a field has been set.
func (o *BTMIndividualQuery138) HasVariableName() bool {
	type getResult interface {
		HasVariableName() bool
	}

	if tx, ok := o.GetActualInstance().(getResult); ok {
		return tx.HasVariableName()
	} else {
		return false
	}
}

// SetVariableName gets a reference to the given BTMIndividualQuery138 and assigns it to the VariableName field.
func (o *BTMIndividualQuery138) SetVariableName(v BTMIndividualQuery138) {
	type getResult interface {
		SetVariableName(v BTMIndividualQuery138)
	}

	o.GetActualInstance().(getResult).SetVariableName(v)
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *BTMIndividualQuery138) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'BTMIndividualCoEdgeQuery-1332'
	if jsonDict["btType"] == "BTMIndividualCoEdgeQuery-1332" {
		// try to unmarshal JSON data into BTMIndividualCoEdgeQuery1332
		var qr *BTMIndividualCoEdgeQuery1332
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMIndividualQuery138 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMIndividualQuery138 = nil
			return fmt.Errorf("Failed to unmarshal BTMIndividualQuery138 as BTMIndividualCoEdgeQuery1332: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMIndividualCreatedByQuery-137'
	if jsonDict["btType"] == "BTMIndividualCreatedByQuery-137" {
		// try to unmarshal JSON data into BTMIndividualCreatedByQuery137
		var qr *BTMIndividualCreatedByQuery137
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMIndividualQuery138 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMIndividualQuery138 = nil
			return fmt.Errorf("Failed to unmarshal BTMIndividualQuery138 as BTMIndividualCreatedByQuery137: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMIndividualSketchRegionQuery-140'
	if jsonDict["btType"] == "BTMIndividualSketchRegionQuery-140" {
		// try to unmarshal JSON data into BTMIndividualSketchRegionQuery140
		var qr *BTMIndividualSketchRegionQuery140
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMIndividualQuery138 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMIndividualQuery138 = nil
			return fmt.Errorf("Failed to unmarshal BTMIndividualQuery138 as BTMIndividualSketchRegionQuery140: %s", err.Error())
		}
	}

	// check if the discriminator value is 'BTMIndividualSketchUniqueVerticesQuery-1472'
	if jsonDict["btType"] == "BTMIndividualSketchUniqueVerticesQuery-1472" {
		// try to unmarshal JSON data into BTMIndividualSketchUniqueVerticesQuery1472
		var qr *BTMIndividualSketchUniqueVerticesQuery1472
		err = json.Unmarshal(data, &qr)
		if err == nil {
			dst.implBTMIndividualQuery138 = qr
			return nil // data stored, return on the first match
		} else {
			dst.implBTMIndividualQuery138 = nil
			return fmt.Errorf("Failed to unmarshal BTMIndividualQuery138 as BTMIndividualSketchUniqueVerticesQuery1472: %s", err.Error())
		}
	}

	var qtx *base_BTMIndividualQuery138
	err = json.Unmarshal(data, &qtx)
	if err == nil {
		dst.implBTMIndividualQuery138 = qtx
		return nil // data stored in dst.base_BTMIndividualQuery138, return on the first match
	} else {
		dst.implBTMIndividualQuery138 = nil
		return fmt.Errorf("Failed to unmarshal BTMIndividualQuery138 as base_BTMIndividualQuery138: %s", err.Error())
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BTMIndividualQuery138) MarshalJSON() ([]byte, error) {
	ret := src.GetActualInstance()
	if ret == nil {
		return nil, nil // no data in oneOf schemas
	} else {
		return json.Marshal(&ret)
	}
}

// Get the actual instance
func (obj *BTMIndividualQuery138) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	return obj.implBTMIndividualQuery138
}

type NullableBTMIndividualQuery138 struct {
	value *BTMIndividualQuery138
	isSet bool
}

func (v NullableBTMIndividualQuery138) Get() *BTMIndividualQuery138 {
	return v.value
}

func (v *NullableBTMIndividualQuery138) Set(val *BTMIndividualQuery138) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMIndividualQuery138) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMIndividualQuery138) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMIndividualQuery138(val *BTMIndividualQuery138) *NullableBTMIndividualQuery138 {
	return &NullableBTMIndividualQuery138{value: val, isSet: true}
}

func (v NullableBTMIndividualQuery138) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMIndividualQuery138) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

type base_BTMIndividualQuery138 struct {
	BtType              *string                    `json:"btType,omitempty"`
	DeterministicIdList *BTMIndividualQueryBase139 `json:"deterministicIdList,omitempty"`
	DeterministicIds    []string                   `json:"deterministicIds,omitempty"`
	ImportMicroversion  *string                    `json:"importMicroversion,omitempty"`
	NodeId              *string                    `json:"nodeId,omitempty"`
	Query               *BTMIndividualQueryBase139 `json:"query,omitempty"`
	QueryString         *string                    `json:"queryString,omitempty"`
	PersistentQuery     *BTPStatement269           `json:"persistentQuery,omitempty"`
	QueryStatement      *BTPStatement269           `json:"queryStatement,omitempty"`
	VariableName        *BTMIndividualQuery138     `json:"variableName,omitempty"`
}

// Newbase_BTMIndividualQuery138 instantiates a new base_BTMIndividualQuery138 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func Newbase_BTMIndividualQuery138() *base_BTMIndividualQuery138 {
	this := base_BTMIndividualQuery138{}
	return &this
}

// Newbase_BTMIndividualQuery138WithDefaults instantiates a new base_BTMIndividualQuery138 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func Newbase_BTMIndividualQuery138WithDefaults() *base_BTMIndividualQuery138 {
	this := base_BTMIndividualQuery138{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *base_BTMIndividualQuery138) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQuery138) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *base_BTMIndividualQuery138) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *base_BTMIndividualQuery138) SetBtType(v string) {
	o.BtType = &v
}

// GetDeterministicIdList returns the DeterministicIdList field value if set, zero value otherwise.
func (o *base_BTMIndividualQuery138) GetDeterministicIdList() BTMIndividualQueryBase139 {
	if o == nil || o.DeterministicIdList == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.DeterministicIdList
}

// GetDeterministicIdListOk returns a tuple with the DeterministicIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQuery138) GetDeterministicIdListOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.DeterministicIdList == nil {
		return nil, false
	}
	return o.DeterministicIdList, true
}

// HasDeterministicIdList returns a boolean if a field has been set.
func (o *base_BTMIndividualQuery138) HasDeterministicIdList() bool {
	if o != nil && o.DeterministicIdList != nil {
		return true
	}

	return false
}

// SetDeterministicIdList gets a reference to the given BTMIndividualQueryBase139 and assigns it to the DeterministicIdList field.
func (o *base_BTMIndividualQuery138) SetDeterministicIdList(v BTMIndividualQueryBase139) {
	o.DeterministicIdList = &v
}

// GetDeterministicIds returns the DeterministicIds field value if set, zero value otherwise.
func (o *base_BTMIndividualQuery138) GetDeterministicIds() []string {
	if o == nil || o.DeterministicIds == nil {
		var ret []string
		return ret
	}
	return o.DeterministicIds
}

// GetDeterministicIdsOk returns a tuple with the DeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQuery138) GetDeterministicIdsOk() ([]string, bool) {
	if o == nil || o.DeterministicIds == nil {
		return nil, false
	}
	return o.DeterministicIds, true
}

// HasDeterministicIds returns a boolean if a field has been set.
func (o *base_BTMIndividualQuery138) HasDeterministicIds() bool {
	if o != nil && o.DeterministicIds != nil {
		return true
	}

	return false
}

// SetDeterministicIds gets a reference to the given []string and assigns it to the DeterministicIds field.
func (o *base_BTMIndividualQuery138) SetDeterministicIds(v []string) {
	o.DeterministicIds = v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *base_BTMIndividualQuery138) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQuery138) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *base_BTMIndividualQuery138) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *base_BTMIndividualQuery138) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *base_BTMIndividualQuery138) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQuery138) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *base_BTMIndividualQuery138) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *base_BTMIndividualQuery138) SetNodeId(v string) {
	o.NodeId = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *base_BTMIndividualQuery138) GetQuery() BTMIndividualQueryBase139 {
	if o == nil || o.Query == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQuery138) GetQueryOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *base_BTMIndividualQuery138) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given BTMIndividualQueryBase139 and assigns it to the Query field.
func (o *base_BTMIndividualQuery138) SetQuery(v BTMIndividualQueryBase139) {
	o.Query = &v
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *base_BTMIndividualQuery138) GetQueryString() string {
	if o == nil || o.QueryString == nil {
		var ret string
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQuery138) GetQueryStringOk() (*string, bool) {
	if o == nil || o.QueryString == nil {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *base_BTMIndividualQuery138) HasQueryString() bool {
	if o != nil && o.QueryString != nil {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *base_BTMIndividualQuery138) SetQueryString(v string) {
	o.QueryString = &v
}

// GetPersistentQuery returns the PersistentQuery field value if set, zero value otherwise.
func (o *base_BTMIndividualQuery138) GetPersistentQuery() BTPStatement269 {
	if o == nil || o.PersistentQuery == nil {
		var ret BTPStatement269
		return ret
	}
	return *o.PersistentQuery
}

// GetPersistentQueryOk returns a tuple with the PersistentQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQuery138) GetPersistentQueryOk() (*BTPStatement269, bool) {
	if o == nil || o.PersistentQuery == nil {
		return nil, false
	}
	return o.PersistentQuery, true
}

// HasPersistentQuery returns a boolean if a field has been set.
func (o *base_BTMIndividualQuery138) HasPersistentQuery() bool {
	if o != nil && o.PersistentQuery != nil {
		return true
	}

	return false
}

// SetPersistentQuery gets a reference to the given BTPStatement269 and assigns it to the PersistentQuery field.
func (o *base_BTMIndividualQuery138) SetPersistentQuery(v BTPStatement269) {
	o.PersistentQuery = &v
}

// GetQueryStatement returns the QueryStatement field value if set, zero value otherwise.
func (o *base_BTMIndividualQuery138) GetQueryStatement() BTPStatement269 {
	if o == nil || o.QueryStatement == nil {
		var ret BTPStatement269
		return ret
	}
	return *o.QueryStatement
}

// GetQueryStatementOk returns a tuple with the QueryStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQuery138) GetQueryStatementOk() (*BTPStatement269, bool) {
	if o == nil || o.QueryStatement == nil {
		return nil, false
	}
	return o.QueryStatement, true
}

// HasQueryStatement returns a boolean if a field has been set.
func (o *base_BTMIndividualQuery138) HasQueryStatement() bool {
	if o != nil && o.QueryStatement != nil {
		return true
	}

	return false
}

// SetQueryStatement gets a reference to the given BTPStatement269 and assigns it to the QueryStatement field.
func (o *base_BTMIndividualQuery138) SetQueryStatement(v BTPStatement269) {
	o.QueryStatement = &v
}

// GetVariableName returns the VariableName field value if set, zero value otherwise.
func (o *base_BTMIndividualQuery138) GetVariableName() BTMIndividualQuery138 {
	if o == nil || o.VariableName == nil {
		var ret BTMIndividualQuery138
		return ret
	}
	return *o.VariableName
}

// GetVariableNameOk returns a tuple with the VariableName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *base_BTMIndividualQuery138) GetVariableNameOk() (*BTMIndividualQuery138, bool) {
	if o == nil || o.VariableName == nil {
		return nil, false
	}
	return o.VariableName, true
}

// HasVariableName returns a boolean if a field has been set.
func (o *base_BTMIndividualQuery138) HasVariableName() bool {
	if o != nil && o.VariableName != nil {
		return true
	}

	return false
}

// SetVariableName gets a reference to the given BTMIndividualQuery138 and assigns it to the VariableName field.
func (o *base_BTMIndividualQuery138) SetVariableName(v BTMIndividualQuery138) {
	o.VariableName = &v
}

func (o base_BTMIndividualQuery138) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BtType != nil {
		toSerialize["btType"] = o.BtType
	}
	if o.DeterministicIdList != nil {
		toSerialize["deterministicIdList"] = o.DeterministicIdList
	}
	if o.DeterministicIds != nil {
		toSerialize["deterministicIds"] = o.DeterministicIds
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.QueryString != nil {
		toSerialize["queryString"] = o.QueryString
	}
	if o.PersistentQuery != nil {
		toSerialize["persistentQuery"] = o.PersistentQuery
	}
	if o.QueryStatement != nil {
		toSerialize["queryStatement"] = o.QueryStatement
	}
	if o.VariableName != nil {
		toSerialize["variableName"] = o.VariableName
	}
	return json.Marshal(toSerialize)
}
