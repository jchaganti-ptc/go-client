/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.168.21206-13add30fbba2
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTMMeshPointQuery1183 struct for BTMMeshPointQuery1183
type BTMMeshPointQuery1183 struct {
	BtType              *string                    `json:"btType,omitempty"`
	DeterministicIdList *BTMIndividualQueryBase139 `json:"deterministicIdList,omitempty"`
	DeterministicIds    []string                   `json:"deterministicIds,omitempty"`
	FullPathAsString    *string                    `json:"fullPathAsString,omitempty"`
	ImportMicroversion  *string                    `json:"importMicroversion,omitempty"`
	NodeId              *string                    `json:"nodeId,omitempty"`
	Occurrence          *BTOccurrence74            `json:"occurrence,omitempty"`
	Path                []string                   `json:"path,omitempty"`
	Query               *BTMIndividualQueryBase139 `json:"query,omitempty"`
	QueryString         *string                    `json:"queryString,omitempty"`
	Origin              *BTVector3d389             `json:"origin,omitempty"`
}

// NewBTMMeshPointQuery1183 instantiates a new BTMMeshPointQuery1183 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTMMeshPointQuery1183() *BTMMeshPointQuery1183 {
	this := BTMMeshPointQuery1183{}
	return &this
}

// NewBTMMeshPointQuery1183WithDefaults instantiates a new BTMMeshPointQuery1183 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTMMeshPointQuery1183WithDefaults() *BTMMeshPointQuery1183 {
	this := BTMMeshPointQuery1183{}
	return &this
}

// GetBtType returns the BtType field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetBtType() string {
	if o == nil || o.BtType == nil {
		var ret string
		return ret
	}
	return *o.BtType
}

// GetBtTypeOk returns a tuple with the BtType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetBtTypeOk() (*string, bool) {
	if o == nil || o.BtType == nil {
		return nil, false
	}
	return o.BtType, true
}

// HasBtType returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasBtType() bool {
	if o != nil && o.BtType != nil {
		return true
	}

	return false
}

// SetBtType gets a reference to the given string and assigns it to the BtType field.
func (o *BTMMeshPointQuery1183) SetBtType(v string) {
	o.BtType = &v
}

// GetDeterministicIdList returns the DeterministicIdList field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetDeterministicIdList() BTMIndividualQueryBase139 {
	if o == nil || o.DeterministicIdList == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.DeterministicIdList
}

// GetDeterministicIdListOk returns a tuple with the DeterministicIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetDeterministicIdListOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.DeterministicIdList == nil {
		return nil, false
	}
	return o.DeterministicIdList, true
}

// HasDeterministicIdList returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasDeterministicIdList() bool {
	if o != nil && o.DeterministicIdList != nil {
		return true
	}

	return false
}

// SetDeterministicIdList gets a reference to the given BTMIndividualQueryBase139 and assigns it to the DeterministicIdList field.
func (o *BTMMeshPointQuery1183) SetDeterministicIdList(v BTMIndividualQueryBase139) {
	o.DeterministicIdList = &v
}

// GetDeterministicIds returns the DeterministicIds field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetDeterministicIds() []string {
	if o == nil || o.DeterministicIds == nil {
		var ret []string
		return ret
	}
	return o.DeterministicIds
}

// GetDeterministicIdsOk returns a tuple with the DeterministicIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetDeterministicIdsOk() ([]string, bool) {
	if o == nil || o.DeterministicIds == nil {
		return nil, false
	}
	return o.DeterministicIds, true
}

// HasDeterministicIds returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasDeterministicIds() bool {
	if o != nil && o.DeterministicIds != nil {
		return true
	}

	return false
}

// SetDeterministicIds gets a reference to the given []string and assigns it to the DeterministicIds field.
func (o *BTMMeshPointQuery1183) SetDeterministicIds(v []string) {
	o.DeterministicIds = v
}

// GetFullPathAsString returns the FullPathAsString field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetFullPathAsString() string {
	if o == nil || o.FullPathAsString == nil {
		var ret string
		return ret
	}
	return *o.FullPathAsString
}

// GetFullPathAsStringOk returns a tuple with the FullPathAsString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetFullPathAsStringOk() (*string, bool) {
	if o == nil || o.FullPathAsString == nil {
		return nil, false
	}
	return o.FullPathAsString, true
}

// HasFullPathAsString returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasFullPathAsString() bool {
	if o != nil && o.FullPathAsString != nil {
		return true
	}

	return false
}

// SetFullPathAsString gets a reference to the given string and assigns it to the FullPathAsString field.
func (o *BTMMeshPointQuery1183) SetFullPathAsString(v string) {
	o.FullPathAsString = &v
}

// GetImportMicroversion returns the ImportMicroversion field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetImportMicroversion() string {
	if o == nil || o.ImportMicroversion == nil {
		var ret string
		return ret
	}
	return *o.ImportMicroversion
}

// GetImportMicroversionOk returns a tuple with the ImportMicroversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetImportMicroversionOk() (*string, bool) {
	if o == nil || o.ImportMicroversion == nil {
		return nil, false
	}
	return o.ImportMicroversion, true
}

// HasImportMicroversion returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasImportMicroversion() bool {
	if o != nil && o.ImportMicroversion != nil {
		return true
	}

	return false
}

// SetImportMicroversion gets a reference to the given string and assigns it to the ImportMicroversion field.
func (o *BTMMeshPointQuery1183) SetImportMicroversion(v string) {
	o.ImportMicroversion = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *BTMMeshPointQuery1183) SetNodeId(v string) {
	o.NodeId = &v
}

// GetOccurrence returns the Occurrence field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetOccurrence() BTOccurrence74 {
	if o == nil || o.Occurrence == nil {
		var ret BTOccurrence74
		return ret
	}
	return *o.Occurrence
}

// GetOccurrenceOk returns a tuple with the Occurrence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetOccurrenceOk() (*BTOccurrence74, bool) {
	if o == nil || o.Occurrence == nil {
		return nil, false
	}
	return o.Occurrence, true
}

// HasOccurrence returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasOccurrence() bool {
	if o != nil && o.Occurrence != nil {
		return true
	}

	return false
}

// SetOccurrence gets a reference to the given BTOccurrence74 and assigns it to the Occurrence field.
func (o *BTMMeshPointQuery1183) SetOccurrence(v BTOccurrence74) {
	o.Occurrence = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetPath() []string {
	if o == nil || o.Path == nil {
		var ret []string
		return ret
	}
	return o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetPathOk() ([]string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given []string and assigns it to the Path field.
func (o *BTMMeshPointQuery1183) SetPath(v []string) {
	o.Path = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetQuery() BTMIndividualQueryBase139 {
	if o == nil || o.Query == nil {
		var ret BTMIndividualQueryBase139
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetQueryOk() (*BTMIndividualQueryBase139, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasQuery() bool {
	if o != nil && o.Query != nil {
		return true
	}

	return false
}

// SetQuery gets a reference to the given BTMIndividualQueryBase139 and assigns it to the Query field.
func (o *BTMMeshPointQuery1183) SetQuery(v BTMIndividualQueryBase139) {
	o.Query = &v
}

// GetQueryString returns the QueryString field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetQueryString() string {
	if o == nil || o.QueryString == nil {
		var ret string
		return ret
	}
	return *o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetQueryStringOk() (*string, bool) {
	if o == nil || o.QueryString == nil {
		return nil, false
	}
	return o.QueryString, true
}

// HasQueryString returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasQueryString() bool {
	if o != nil && o.QueryString != nil {
		return true
	}

	return false
}

// SetQueryString gets a reference to the given string and assigns it to the QueryString field.
func (o *BTMMeshPointQuery1183) SetQueryString(v string) {
	o.QueryString = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *BTMMeshPointQuery1183) GetOrigin() BTVector3d389 {
	if o == nil || o.Origin == nil {
		var ret BTVector3d389
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTMMeshPointQuery1183) GetOriginOk() (*BTVector3d389, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *BTMMeshPointQuery1183) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given BTVector3d389 and assigns it to the Origin field.
func (o *BTMMeshPointQuery1183) SetOrigin(v BTVector3d389) {
	o.Origin = &v
}

func (o BTMMeshPointQuery1183) MarshalJSON() ([]byte, error) {
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
	if o.FullPathAsString != nil {
		toSerialize["fullPathAsString"] = o.FullPathAsString
	}
	if o.ImportMicroversion != nil {
		toSerialize["importMicroversion"] = o.ImportMicroversion
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Occurrence != nil {
		toSerialize["occurrence"] = o.Occurrence
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.QueryString != nil {
		toSerialize["queryString"] = o.QueryString
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	return json.Marshal(toSerialize)
}

type NullableBTMMeshPointQuery1183 struct {
	value *BTMMeshPointQuery1183
	isSet bool
}

func (v NullableBTMMeshPointQuery1183) Get() *BTMMeshPointQuery1183 {
	return v.value
}

func (v *NullableBTMMeshPointQuery1183) Set(val *BTMMeshPointQuery1183) {
	v.value = val
	v.isSet = true
}

func (v NullableBTMMeshPointQuery1183) IsSet() bool {
	return v.isSet
}

func (v *NullableBTMMeshPointQuery1183) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTMMeshPointQuery1183(val *BTMMeshPointQuery1183) *NullableBTMMeshPointQuery1183 {
	return &NullableBTMMeshPointQuery1183{value: val, isSet: true}
}

func (v NullableBTMMeshPointQuery1183) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTMMeshPointQuery1183) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
