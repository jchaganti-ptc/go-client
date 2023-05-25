/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.164.16419-6916b772b99f
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTPropertiesTableTemplateInfo struct for BTPropertiesTableTemplateInfo
type BTPropertiesTableTemplateInfo struct {
	CompanyId *string `json:"companyId,omitempty"`
	// URI to fetch complete information of the resource.
	Href *string `json:"href,omitempty"`
	// Id of the resource.
	Id        *string `json:"id,omitempty"`
	IsActive  *bool   `json:"isActive,omitempty"`
	IsAllCaps *bool   `json:"isAllCaps,omitempty"`
	// Name of the resource.
	Name            *string                        `json:"name,omitempty"`
	PropertyColumns []BTSimplePropertyInfo         `json:"propertyColumns,omitempty"`
	TableType       *BTPropertiesTableTemplateType `json:"tableType,omitempty"`
	TemplateGroupId *string                        `json:"templateGroupId,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef *string `json:"viewRef,omitempty"`
}

// NewBTPropertiesTableTemplateInfo instantiates a new BTPropertiesTableTemplateInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTPropertiesTableTemplateInfo() *BTPropertiesTableTemplateInfo {
	this := BTPropertiesTableTemplateInfo{}
	return &this
}

// NewBTPropertiesTableTemplateInfoWithDefaults instantiates a new BTPropertiesTableTemplateInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTPropertiesTableTemplateInfoWithDefaults() *BTPropertiesTableTemplateInfo {
	this := BTPropertiesTableTemplateInfo{}
	return &this
}

// GetCompanyId returns the CompanyId field value if set, zero value otherwise.
func (o *BTPropertiesTableTemplateInfo) GetCompanyId() string {
	if o == nil || o.CompanyId == nil {
		var ret string
		return ret
	}
	return *o.CompanyId
}

// GetCompanyIdOk returns a tuple with the CompanyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertiesTableTemplateInfo) GetCompanyIdOk() (*string, bool) {
	if o == nil || o.CompanyId == nil {
		return nil, false
	}
	return o.CompanyId, true
}

// HasCompanyId returns a boolean if a field has been set.
func (o *BTPropertiesTableTemplateInfo) HasCompanyId() bool {
	if o != nil && o.CompanyId != nil {
		return true
	}

	return false
}

// SetCompanyId gets a reference to the given string and assigns it to the CompanyId field.
func (o *BTPropertiesTableTemplateInfo) SetCompanyId(v string) {
	o.CompanyId = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTPropertiesTableTemplateInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertiesTableTemplateInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTPropertiesTableTemplateInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTPropertiesTableTemplateInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTPropertiesTableTemplateInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertiesTableTemplateInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTPropertiesTableTemplateInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTPropertiesTableTemplateInfo) SetId(v string) {
	o.Id = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *BTPropertiesTableTemplateInfo) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertiesTableTemplateInfo) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *BTPropertiesTableTemplateInfo) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *BTPropertiesTableTemplateInfo) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetIsAllCaps returns the IsAllCaps field value if set, zero value otherwise.
func (o *BTPropertiesTableTemplateInfo) GetIsAllCaps() bool {
	if o == nil || o.IsAllCaps == nil {
		var ret bool
		return ret
	}
	return *o.IsAllCaps
}

// GetIsAllCapsOk returns a tuple with the IsAllCaps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertiesTableTemplateInfo) GetIsAllCapsOk() (*bool, bool) {
	if o == nil || o.IsAllCaps == nil {
		return nil, false
	}
	return o.IsAllCaps, true
}

// HasIsAllCaps returns a boolean if a field has been set.
func (o *BTPropertiesTableTemplateInfo) HasIsAllCaps() bool {
	if o != nil && o.IsAllCaps != nil {
		return true
	}

	return false
}

// SetIsAllCaps gets a reference to the given bool and assigns it to the IsAllCaps field.
func (o *BTPropertiesTableTemplateInfo) SetIsAllCaps(v bool) {
	o.IsAllCaps = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTPropertiesTableTemplateInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertiesTableTemplateInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTPropertiesTableTemplateInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTPropertiesTableTemplateInfo) SetName(v string) {
	o.Name = &v
}

// GetPropertyColumns returns the PropertyColumns field value if set, zero value otherwise.
func (o *BTPropertiesTableTemplateInfo) GetPropertyColumns() []BTSimplePropertyInfo {
	if o == nil || o.PropertyColumns == nil {
		var ret []BTSimplePropertyInfo
		return ret
	}
	return o.PropertyColumns
}

// GetPropertyColumnsOk returns a tuple with the PropertyColumns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertiesTableTemplateInfo) GetPropertyColumnsOk() ([]BTSimplePropertyInfo, bool) {
	if o == nil || o.PropertyColumns == nil {
		return nil, false
	}
	return o.PropertyColumns, true
}

// HasPropertyColumns returns a boolean if a field has been set.
func (o *BTPropertiesTableTemplateInfo) HasPropertyColumns() bool {
	if o != nil && o.PropertyColumns != nil {
		return true
	}

	return false
}

// SetPropertyColumns gets a reference to the given []BTSimplePropertyInfo and assigns it to the PropertyColumns field.
func (o *BTPropertiesTableTemplateInfo) SetPropertyColumns(v []BTSimplePropertyInfo) {
	o.PropertyColumns = v
}

// GetTableType returns the TableType field value if set, zero value otherwise.
func (o *BTPropertiesTableTemplateInfo) GetTableType() BTPropertiesTableTemplateType {
	if o == nil || o.TableType == nil {
		var ret BTPropertiesTableTemplateType
		return ret
	}
	return *o.TableType
}

// GetTableTypeOk returns a tuple with the TableType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertiesTableTemplateInfo) GetTableTypeOk() (*BTPropertiesTableTemplateType, bool) {
	if o == nil || o.TableType == nil {
		return nil, false
	}
	return o.TableType, true
}

// HasTableType returns a boolean if a field has been set.
func (o *BTPropertiesTableTemplateInfo) HasTableType() bool {
	if o != nil && o.TableType != nil {
		return true
	}

	return false
}

// SetTableType gets a reference to the given BTPropertiesTableTemplateType and assigns it to the TableType field.
func (o *BTPropertiesTableTemplateInfo) SetTableType(v BTPropertiesTableTemplateType) {
	o.TableType = &v
}

// GetTemplateGroupId returns the TemplateGroupId field value if set, zero value otherwise.
func (o *BTPropertiesTableTemplateInfo) GetTemplateGroupId() string {
	if o == nil || o.TemplateGroupId == nil {
		var ret string
		return ret
	}
	return *o.TemplateGroupId
}

// GetTemplateGroupIdOk returns a tuple with the TemplateGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertiesTableTemplateInfo) GetTemplateGroupIdOk() (*string, bool) {
	if o == nil || o.TemplateGroupId == nil {
		return nil, false
	}
	return o.TemplateGroupId, true
}

// HasTemplateGroupId returns a boolean if a field has been set.
func (o *BTPropertiesTableTemplateInfo) HasTemplateGroupId() bool {
	if o != nil && o.TemplateGroupId != nil {
		return true
	}

	return false
}

// SetTemplateGroupId gets a reference to the given string and assigns it to the TemplateGroupId field.
func (o *BTPropertiesTableTemplateInfo) SetTemplateGroupId(v string) {
	o.TemplateGroupId = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTPropertiesTableTemplateInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTPropertiesTableTemplateInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTPropertiesTableTemplateInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTPropertiesTableTemplateInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

func (o BTPropertiesTableTemplateInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CompanyId != nil {
		toSerialize["companyId"] = o.CompanyId
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsActive != nil {
		toSerialize["isActive"] = o.IsActive
	}
	if o.IsAllCaps != nil {
		toSerialize["isAllCaps"] = o.IsAllCaps
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PropertyColumns != nil {
		toSerialize["propertyColumns"] = o.PropertyColumns
	}
	if o.TableType != nil {
		toSerialize["tableType"] = o.TableType
	}
	if o.TemplateGroupId != nil {
		toSerialize["templateGroupId"] = o.TemplateGroupId
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	return json.Marshal(toSerialize)
}

type NullableBTPropertiesTableTemplateInfo struct {
	value *BTPropertiesTableTemplateInfo
	isSet bool
}

func (v NullableBTPropertiesTableTemplateInfo) Get() *BTPropertiesTableTemplateInfo {
	return v.value
}

func (v *NullableBTPropertiesTableTemplateInfo) Set(val *BTPropertiesTableTemplateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTPropertiesTableTemplateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTPropertiesTableTemplateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTPropertiesTableTemplateInfo(val *BTPropertiesTableTemplateInfo) *NullableBTPropertiesTableTemplateInfo {
	return &NullableBTPropertiesTableTemplateInfo{value: val, isSet: true}
}

func (v NullableBTPropertiesTableTemplateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTPropertiesTableTemplateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
