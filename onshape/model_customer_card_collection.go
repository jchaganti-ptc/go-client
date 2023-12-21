/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.27960-e195de6ac56c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// CustomerCardCollection struct for CustomerCardCollection
type CustomerCardCollection struct {
	Count          *int32                            `json:"count,omitempty"`
	Data           []Card                            `json:"data,omitempty"`
	HasMore        *bool                             `json:"hasMore,omitempty"`
	RequestOptions *RequestOptions                   `json:"requestOptions,omitempty"`
	RequestParams  map[string]map[string]interface{} `json:"requestParams,omitempty"`
	TotalCount     *int32                            `json:"totalCount,omitempty"`
	Url            *string                           `json:"url,omitempty"`
}

// NewCustomerCardCollection instantiates a new CustomerCardCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerCardCollection() *CustomerCardCollection {
	this := CustomerCardCollection{}
	return &this
}

// NewCustomerCardCollectionWithDefaults instantiates a new CustomerCardCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerCardCollectionWithDefaults() *CustomerCardCollection {
	this := CustomerCardCollection{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CustomerCardCollection) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCardCollection) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CustomerCardCollection) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CustomerCardCollection) SetCount(v int32) {
	o.Count = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CustomerCardCollection) GetData() []Card {
	if o == nil || o.Data == nil {
		var ret []Card
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCardCollection) GetDataOk() ([]Card, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CustomerCardCollection) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []Card and assigns it to the Data field.
func (o *CustomerCardCollection) SetData(v []Card) {
	o.Data = v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *CustomerCardCollection) GetHasMore() bool {
	if o == nil || o.HasMore == nil {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCardCollection) GetHasMoreOk() (*bool, bool) {
	if o == nil || o.HasMore == nil {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *CustomerCardCollection) HasHasMore() bool {
	if o != nil && o.HasMore != nil {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *CustomerCardCollection) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetRequestOptions returns the RequestOptions field value if set, zero value otherwise.
func (o *CustomerCardCollection) GetRequestOptions() RequestOptions {
	if o == nil || o.RequestOptions == nil {
		var ret RequestOptions
		return ret
	}
	return *o.RequestOptions
}

// GetRequestOptionsOk returns a tuple with the RequestOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCardCollection) GetRequestOptionsOk() (*RequestOptions, bool) {
	if o == nil || o.RequestOptions == nil {
		return nil, false
	}
	return o.RequestOptions, true
}

// HasRequestOptions returns a boolean if a field has been set.
func (o *CustomerCardCollection) HasRequestOptions() bool {
	if o != nil && o.RequestOptions != nil {
		return true
	}

	return false
}

// SetRequestOptions gets a reference to the given RequestOptions and assigns it to the RequestOptions field.
func (o *CustomerCardCollection) SetRequestOptions(v RequestOptions) {
	o.RequestOptions = &v
}

// GetRequestParams returns the RequestParams field value if set, zero value otherwise.
func (o *CustomerCardCollection) GetRequestParams() map[string]map[string]interface{} {
	if o == nil || o.RequestParams == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.RequestParams
}

// GetRequestParamsOk returns a tuple with the RequestParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCardCollection) GetRequestParamsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.RequestParams == nil {
		return nil, false
	}
	return o.RequestParams, true
}

// HasRequestParams returns a boolean if a field has been set.
func (o *CustomerCardCollection) HasRequestParams() bool {
	if o != nil && o.RequestParams != nil {
		return true
	}

	return false
}

// SetRequestParams gets a reference to the given map[string]map[string]interface{} and assigns it to the RequestParams field.
func (o *CustomerCardCollection) SetRequestParams(v map[string]map[string]interface{}) {
	o.RequestParams = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *CustomerCardCollection) GetTotalCount() int32 {
	if o == nil || o.TotalCount == nil {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCardCollection) GetTotalCountOk() (*int32, bool) {
	if o == nil || o.TotalCount == nil {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *CustomerCardCollection) HasTotalCount() bool {
	if o != nil && o.TotalCount != nil {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *CustomerCardCollection) SetTotalCount(v int32) {
	o.TotalCount = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *CustomerCardCollection) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerCardCollection) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *CustomerCardCollection) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *CustomerCardCollection) SetUrl(v string) {
	o.Url = &v
}

func (o CustomerCardCollection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.HasMore != nil {
		toSerialize["hasMore"] = o.HasMore
	}
	if o.RequestOptions != nil {
		toSerialize["requestOptions"] = o.RequestOptions
	}
	if o.RequestParams != nil {
		toSerialize["requestParams"] = o.RequestParams
	}
	if o.TotalCount != nil {
		toSerialize["totalCount"] = o.TotalCount
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerCardCollection struct {
	value *CustomerCardCollection
	isSet bool
}

func (v NullableCustomerCardCollection) Get() *CustomerCardCollection {
	return v.value
}

func (v *NullableCustomerCardCollection) Set(val *CustomerCardCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerCardCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerCardCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerCardCollection(val *CustomerCardCollection) *NullableCustomerCardCollection {
	return &NullableCustomerCardCollection{value: val, isSet: true}
}

func (v NullableCustomerCardCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerCardCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
