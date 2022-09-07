/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6358-16771368e3d1
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTApiTable2300 struct for BTApiTable2300
type BTApiTable2300 struct {
	Columns   []BTApiTableColumn3141 `json:"columns,omitempty"`
	EntityIds []string               `json:"entityIds,omitempty"`
	Id        *string                `json:"id,omitempty"`
	Rows      []BTApiTableRow2915    `json:"rows,omitempty"`
	Title     *string                `json:"title,omitempty"`
}

// NewBTApiTable2300 instantiates a new BTApiTable2300 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTApiTable2300() *BTApiTable2300 {
	this := BTApiTable2300{}
	return &this
}

// NewBTApiTable2300WithDefaults instantiates a new BTApiTable2300 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTApiTable2300WithDefaults() *BTApiTable2300 {
	this := BTApiTable2300{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *BTApiTable2300) GetColumns() []BTApiTableColumn3141 {
	if o == nil || o.Columns == nil {
		var ret []BTApiTableColumn3141
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiTable2300) GetColumnsOk() ([]BTApiTableColumn3141, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *BTApiTable2300) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []BTApiTableColumn3141 and assigns it to the Columns field.
func (o *BTApiTable2300) SetColumns(v []BTApiTableColumn3141) {
	o.Columns = v
}

// GetEntityIds returns the EntityIds field value if set, zero value otherwise.
func (o *BTApiTable2300) GetEntityIds() []string {
	if o == nil || o.EntityIds == nil {
		var ret []string
		return ret
	}
	return o.EntityIds
}

// GetEntityIdsOk returns a tuple with the EntityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiTable2300) GetEntityIdsOk() ([]string, bool) {
	if o == nil || o.EntityIds == nil {
		return nil, false
	}
	return o.EntityIds, true
}

// HasEntityIds returns a boolean if a field has been set.
func (o *BTApiTable2300) HasEntityIds() bool {
	if o != nil && o.EntityIds != nil {
		return true
	}

	return false
}

// SetEntityIds gets a reference to the given []string and assigns it to the EntityIds field.
func (o *BTApiTable2300) SetEntityIds(v []string) {
	o.EntityIds = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTApiTable2300) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiTable2300) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTApiTable2300) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTApiTable2300) SetId(v string) {
	o.Id = &v
}

// GetRows returns the Rows field value if set, zero value otherwise.
func (o *BTApiTable2300) GetRows() []BTApiTableRow2915 {
	if o == nil || o.Rows == nil {
		var ret []BTApiTableRow2915
		return ret
	}
	return o.Rows
}

// GetRowsOk returns a tuple with the Rows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiTable2300) GetRowsOk() ([]BTApiTableRow2915, bool) {
	if o == nil || o.Rows == nil {
		return nil, false
	}
	return o.Rows, true
}

// HasRows returns a boolean if a field has been set.
func (o *BTApiTable2300) HasRows() bool {
	if o != nil && o.Rows != nil {
		return true
	}

	return false
}

// SetRows gets a reference to the given []BTApiTableRow2915 and assigns it to the Rows field.
func (o *BTApiTable2300) SetRows(v []BTApiTableRow2915) {
	o.Rows = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *BTApiTable2300) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTApiTable2300) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *BTApiTable2300) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *BTApiTable2300) SetTitle(v string) {
	o.Title = &v
}

func (o BTApiTable2300) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.EntityIds != nil {
		toSerialize["entityIds"] = o.EntityIds
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Rows != nil {
		toSerialize["rows"] = o.Rows
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableBTApiTable2300 struct {
	value *BTApiTable2300
	isSet bool
}

func (v NullableBTApiTable2300) Get() *BTApiTable2300 {
	return v.value
}

func (v *NullableBTApiTable2300) Set(val *BTApiTable2300) {
	v.value = val
	v.isSet = true
}

func (v NullableBTApiTable2300) IsSet() bool {
	return v.isSet
}

func (v *NullableBTApiTable2300) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTApiTable2300(val *BTApiTable2300) *NullableBTApiTable2300 {
	return &NullableBTApiTable2300{value: val, isSet: true}
}

func (v NullableBTApiTable2300) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTApiTable2300) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
