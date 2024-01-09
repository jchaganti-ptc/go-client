/*
Onshape REST API

## Welcome to the Onshape REST API Explorer  To use this API explorer, sign in to your [Onshape](https://cad.onshape.com) account in another tab, then click the **Try it out** button below (it toggles to a **Cancel** button when selected).  See the **[API Explorer Guide](https://onshape-public.github.io/docs/api-intro/explorer/)** for help navigating this API Explorer, including **[authentication](https://onshape-public.github.io/docs/api-intro/explorer/#authentication)**.  **Tip:** To ensure the current session isn't used when trying other authentication techniques, make sure to [remove the Onshape cookie](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site) as per the instructions for your browser. Alternatively, you can use a private or incognito window.  ## See Also  * [Onshape API Guide](https://onshape-public.github.io/docs/): Our full suite of developer guides, to be used as an accompaniment to this API Explorer. * [Onshape Developer Portal](https://dev-portal.onshape.com/): The Onshape portal for managing your API keys, OAuth2 credentials, your Onshape applications, and your Onshape App Store entries. * [Authentication Guide](https://onshape-public.github.io/docs/auth/): Our guide to using API keys, request signatures, and OAuth2 in  your Onshape applications.

API version: 1.174.28658-06d4d4923fc7
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"encoding/json"
)

// BTUserAdminSummaryInfo struct for BTUserAdminSummaryInfo
type BTUserAdminSummaryInfo struct {
	// URI to fetch complete information of the resource.
	Href *string `json:"href,omitempty"`
	// Id of the resource.
	Id *string `json:"id,omitempty"`
	// Name of the resource.
	Name *string `json:"name,omitempty"`
	// URI to visualize the resource in a webclient if applicable.
	ViewRef                   *string               `json:"viewRef,omitempty"`
	Image                     *string               `json:"image,omitempty"`
	State                     *int32                `json:"state,omitempty"`
	DocumentationName         *string               `json:"documentationName,omitempty"`
	Email                     *string               `json:"email,omitempty"`
	FirstName                 *string               `json:"firstName,omitempty"`
	LastName                  *string               `json:"lastName,omitempty"`
	Company                   *BTCompanySummaryInfo `json:"company,omitempty"`
	DocumentationNameOverride *string               `json:"documentationNameOverride,omitempty"`
	GlobalPermissions         *GlobalPermissionInfo `json:"globalPermissions,omitempty"`
	InvitationState           *int32                `json:"invitationState,omitempty"`
	IsGuest                   *bool                 `json:"isGuest,omitempty"`
	IsLight                   *bool                 `json:"isLight,omitempty"`
	LastLoginTime             *JSONTime             `json:"lastLoginTime,omitempty"`
	PersonalMessageAllowed    *bool                 `json:"personalMessageAllowed,omitempty"`
	Source                    *int32                `json:"source,omitempty"`
	ActivePlanId              *string               `json:"activePlanId,omitempty"`
	BillingUpdateRequired     *bool                 `json:"billingUpdateRequired,omitempty"`
	CompanyRoles              []CompanyRole         `json:"companyRoles,omitempty"`
	CreatedAt                 *JSONTime             `json:"createdAt,omitempty"`
	ForumId                   *string               `json:"forumId,omitempty"`
	SystemUser                *bool                 `json:"systemUser,omitempty"`
	TotpEnabled               *bool                 `json:"totpEnabled,omitempty"`
}

// NewBTUserAdminSummaryInfo instantiates a new BTUserAdminSummaryInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBTUserAdminSummaryInfo() *BTUserAdminSummaryInfo {
	this := BTUserAdminSummaryInfo{}
	return &this
}

// NewBTUserAdminSummaryInfoWithDefaults instantiates a new BTUserAdminSummaryInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBTUserAdminSummaryInfoWithDefaults() *BTUserAdminSummaryInfo {
	this := BTUserAdminSummaryInfo{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *BTUserAdminSummaryInfo) SetHref(v string) {
	o.Href = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BTUserAdminSummaryInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BTUserAdminSummaryInfo) SetName(v string) {
	o.Name = &v
}

// GetViewRef returns the ViewRef field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetViewRef() string {
	if o == nil || o.ViewRef == nil {
		var ret string
		return ret
	}
	return *o.ViewRef
}

// GetViewRefOk returns a tuple with the ViewRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetViewRefOk() (*string, bool) {
	if o == nil || o.ViewRef == nil {
		return nil, false
	}
	return o.ViewRef, true
}

// HasViewRef returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasViewRef() bool {
	if o != nil && o.ViewRef != nil {
		return true
	}

	return false
}

// SetViewRef gets a reference to the given string and assigns it to the ViewRef field.
func (o *BTUserAdminSummaryInfo) SetViewRef(v string) {
	o.ViewRef = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetImage() string {
	if o == nil || o.Image == nil {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetImageOk() (*string, bool) {
	if o == nil || o.Image == nil {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasImage() bool {
	if o != nil && o.Image != nil {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *BTUserAdminSummaryInfo) SetImage(v string) {
	o.Image = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetState() int32 {
	if o == nil || o.State == nil {
		var ret int32
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetStateOk() (*int32, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given int32 and assigns it to the State field.
func (o *BTUserAdminSummaryInfo) SetState(v int32) {
	o.State = &v
}

// GetDocumentationName returns the DocumentationName field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetDocumentationName() string {
	if o == nil || o.DocumentationName == nil {
		var ret string
		return ret
	}
	return *o.DocumentationName
}

// GetDocumentationNameOk returns a tuple with the DocumentationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetDocumentationNameOk() (*string, bool) {
	if o == nil || o.DocumentationName == nil {
		return nil, false
	}
	return o.DocumentationName, true
}

// HasDocumentationName returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasDocumentationName() bool {
	if o != nil && o.DocumentationName != nil {
		return true
	}

	return false
}

// SetDocumentationName gets a reference to the given string and assigns it to the DocumentationName field.
func (o *BTUserAdminSummaryInfo) SetDocumentationName(v string) {
	o.DocumentationName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *BTUserAdminSummaryInfo) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetFirstName() string {
	if o == nil || o.FirstName == nil {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetFirstNameOk() (*string, bool) {
	if o == nil || o.FirstName == nil {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasFirstName() bool {
	if o != nil && o.FirstName != nil {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *BTUserAdminSummaryInfo) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetLastName() string {
	if o == nil || o.LastName == nil {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetLastNameOk() (*string, bool) {
	if o == nil || o.LastName == nil {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasLastName() bool {
	if o != nil && o.LastName != nil {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *BTUserAdminSummaryInfo) SetLastName(v string) {
	o.LastName = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetCompany() BTCompanySummaryInfo {
	if o == nil || o.Company == nil {
		var ret BTCompanySummaryInfo
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetCompanyOk() (*BTCompanySummaryInfo, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given BTCompanySummaryInfo and assigns it to the Company field.
func (o *BTUserAdminSummaryInfo) SetCompany(v BTCompanySummaryInfo) {
	o.Company = &v
}

// GetDocumentationNameOverride returns the DocumentationNameOverride field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetDocumentationNameOverride() string {
	if o == nil || o.DocumentationNameOverride == nil {
		var ret string
		return ret
	}
	return *o.DocumentationNameOverride
}

// GetDocumentationNameOverrideOk returns a tuple with the DocumentationNameOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetDocumentationNameOverrideOk() (*string, bool) {
	if o == nil || o.DocumentationNameOverride == nil {
		return nil, false
	}
	return o.DocumentationNameOverride, true
}

// HasDocumentationNameOverride returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasDocumentationNameOverride() bool {
	if o != nil && o.DocumentationNameOverride != nil {
		return true
	}

	return false
}

// SetDocumentationNameOverride gets a reference to the given string and assigns it to the DocumentationNameOverride field.
func (o *BTUserAdminSummaryInfo) SetDocumentationNameOverride(v string) {
	o.DocumentationNameOverride = &v
}

// GetGlobalPermissions returns the GlobalPermissions field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetGlobalPermissions() GlobalPermissionInfo {
	if o == nil || o.GlobalPermissions == nil {
		var ret GlobalPermissionInfo
		return ret
	}
	return *o.GlobalPermissions
}

// GetGlobalPermissionsOk returns a tuple with the GlobalPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetGlobalPermissionsOk() (*GlobalPermissionInfo, bool) {
	if o == nil || o.GlobalPermissions == nil {
		return nil, false
	}
	return o.GlobalPermissions, true
}

// HasGlobalPermissions returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasGlobalPermissions() bool {
	if o != nil && o.GlobalPermissions != nil {
		return true
	}

	return false
}

// SetGlobalPermissions gets a reference to the given GlobalPermissionInfo and assigns it to the GlobalPermissions field.
func (o *BTUserAdminSummaryInfo) SetGlobalPermissions(v GlobalPermissionInfo) {
	o.GlobalPermissions = &v
}

// GetInvitationState returns the InvitationState field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetInvitationState() int32 {
	if o == nil || o.InvitationState == nil {
		var ret int32
		return ret
	}
	return *o.InvitationState
}

// GetInvitationStateOk returns a tuple with the InvitationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetInvitationStateOk() (*int32, bool) {
	if o == nil || o.InvitationState == nil {
		return nil, false
	}
	return o.InvitationState, true
}

// HasInvitationState returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasInvitationState() bool {
	if o != nil && o.InvitationState != nil {
		return true
	}

	return false
}

// SetInvitationState gets a reference to the given int32 and assigns it to the InvitationState field.
func (o *BTUserAdminSummaryInfo) SetInvitationState(v int32) {
	o.InvitationState = &v
}

// GetIsGuest returns the IsGuest field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetIsGuest() bool {
	if o == nil || o.IsGuest == nil {
		var ret bool
		return ret
	}
	return *o.IsGuest
}

// GetIsGuestOk returns a tuple with the IsGuest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetIsGuestOk() (*bool, bool) {
	if o == nil || o.IsGuest == nil {
		return nil, false
	}
	return o.IsGuest, true
}

// HasIsGuest returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasIsGuest() bool {
	if o != nil && o.IsGuest != nil {
		return true
	}

	return false
}

// SetIsGuest gets a reference to the given bool and assigns it to the IsGuest field.
func (o *BTUserAdminSummaryInfo) SetIsGuest(v bool) {
	o.IsGuest = &v
}

// GetIsLight returns the IsLight field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetIsLight() bool {
	if o == nil || o.IsLight == nil {
		var ret bool
		return ret
	}
	return *o.IsLight
}

// GetIsLightOk returns a tuple with the IsLight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetIsLightOk() (*bool, bool) {
	if o == nil || o.IsLight == nil {
		return nil, false
	}
	return o.IsLight, true
}

// HasIsLight returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasIsLight() bool {
	if o != nil && o.IsLight != nil {
		return true
	}

	return false
}

// SetIsLight gets a reference to the given bool and assigns it to the IsLight field.
func (o *BTUserAdminSummaryInfo) SetIsLight(v bool) {
	o.IsLight = &v
}

// GetLastLoginTime returns the LastLoginTime field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetLastLoginTime() JSONTime {
	if o == nil || o.LastLoginTime == nil {
		var ret JSONTime
		return ret
	}
	return *o.LastLoginTime
}

// GetLastLoginTimeOk returns a tuple with the LastLoginTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetLastLoginTimeOk() (*JSONTime, bool) {
	if o == nil || o.LastLoginTime == nil {
		return nil, false
	}
	return o.LastLoginTime, true
}

// HasLastLoginTime returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasLastLoginTime() bool {
	if o != nil && o.LastLoginTime != nil {
		return true
	}

	return false
}

// SetLastLoginTime gets a reference to the given JSONTime and assigns it to the LastLoginTime field.
func (o *BTUserAdminSummaryInfo) SetLastLoginTime(v JSONTime) {
	o.LastLoginTime = &v
}

// GetPersonalMessageAllowed returns the PersonalMessageAllowed field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetPersonalMessageAllowed() bool {
	if o == nil || o.PersonalMessageAllowed == nil {
		var ret bool
		return ret
	}
	return *o.PersonalMessageAllowed
}

// GetPersonalMessageAllowedOk returns a tuple with the PersonalMessageAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetPersonalMessageAllowedOk() (*bool, bool) {
	if o == nil || o.PersonalMessageAllowed == nil {
		return nil, false
	}
	return o.PersonalMessageAllowed, true
}

// HasPersonalMessageAllowed returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasPersonalMessageAllowed() bool {
	if o != nil && o.PersonalMessageAllowed != nil {
		return true
	}

	return false
}

// SetPersonalMessageAllowed gets a reference to the given bool and assigns it to the PersonalMessageAllowed field.
func (o *BTUserAdminSummaryInfo) SetPersonalMessageAllowed(v bool) {
	o.PersonalMessageAllowed = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetSource() int32 {
	if o == nil || o.Source == nil {
		var ret int32
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetSourceOk() (*int32, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given int32 and assigns it to the Source field.
func (o *BTUserAdminSummaryInfo) SetSource(v int32) {
	o.Source = &v
}

// GetActivePlanId returns the ActivePlanId field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetActivePlanId() string {
	if o == nil || o.ActivePlanId == nil {
		var ret string
		return ret
	}
	return *o.ActivePlanId
}

// GetActivePlanIdOk returns a tuple with the ActivePlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetActivePlanIdOk() (*string, bool) {
	if o == nil || o.ActivePlanId == nil {
		return nil, false
	}
	return o.ActivePlanId, true
}

// HasActivePlanId returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasActivePlanId() bool {
	if o != nil && o.ActivePlanId != nil {
		return true
	}

	return false
}

// SetActivePlanId gets a reference to the given string and assigns it to the ActivePlanId field.
func (o *BTUserAdminSummaryInfo) SetActivePlanId(v string) {
	o.ActivePlanId = &v
}

// GetBillingUpdateRequired returns the BillingUpdateRequired field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetBillingUpdateRequired() bool {
	if o == nil || o.BillingUpdateRequired == nil {
		var ret bool
		return ret
	}
	return *o.BillingUpdateRequired
}

// GetBillingUpdateRequiredOk returns a tuple with the BillingUpdateRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetBillingUpdateRequiredOk() (*bool, bool) {
	if o == nil || o.BillingUpdateRequired == nil {
		return nil, false
	}
	return o.BillingUpdateRequired, true
}

// HasBillingUpdateRequired returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasBillingUpdateRequired() bool {
	if o != nil && o.BillingUpdateRequired != nil {
		return true
	}

	return false
}

// SetBillingUpdateRequired gets a reference to the given bool and assigns it to the BillingUpdateRequired field.
func (o *BTUserAdminSummaryInfo) SetBillingUpdateRequired(v bool) {
	o.BillingUpdateRequired = &v
}

// GetCompanyRoles returns the CompanyRoles field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetCompanyRoles() []CompanyRole {
	if o == nil || o.CompanyRoles == nil {
		var ret []CompanyRole
		return ret
	}
	return o.CompanyRoles
}

// GetCompanyRolesOk returns a tuple with the CompanyRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetCompanyRolesOk() ([]CompanyRole, bool) {
	if o == nil || o.CompanyRoles == nil {
		return nil, false
	}
	return o.CompanyRoles, true
}

// HasCompanyRoles returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasCompanyRoles() bool {
	if o != nil && o.CompanyRoles != nil {
		return true
	}

	return false
}

// SetCompanyRoles gets a reference to the given []CompanyRole and assigns it to the CompanyRoles field.
func (o *BTUserAdminSummaryInfo) SetCompanyRoles(v []CompanyRole) {
	o.CompanyRoles = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetCreatedAt() JSONTime {
	if o == nil || o.CreatedAt == nil {
		var ret JSONTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetCreatedAtOk() (*JSONTime, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given JSONTime and assigns it to the CreatedAt field.
func (o *BTUserAdminSummaryInfo) SetCreatedAt(v JSONTime) {
	o.CreatedAt = &v
}

// GetForumId returns the ForumId field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetForumId() string {
	if o == nil || o.ForumId == nil {
		var ret string
		return ret
	}
	return *o.ForumId
}

// GetForumIdOk returns a tuple with the ForumId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetForumIdOk() (*string, bool) {
	if o == nil || o.ForumId == nil {
		return nil, false
	}
	return o.ForumId, true
}

// HasForumId returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasForumId() bool {
	if o != nil && o.ForumId != nil {
		return true
	}

	return false
}

// SetForumId gets a reference to the given string and assigns it to the ForumId field.
func (o *BTUserAdminSummaryInfo) SetForumId(v string) {
	o.ForumId = &v
}

// GetSystemUser returns the SystemUser field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetSystemUser() bool {
	if o == nil || o.SystemUser == nil {
		var ret bool
		return ret
	}
	return *o.SystemUser
}

// GetSystemUserOk returns a tuple with the SystemUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetSystemUserOk() (*bool, bool) {
	if o == nil || o.SystemUser == nil {
		return nil, false
	}
	return o.SystemUser, true
}

// HasSystemUser returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasSystemUser() bool {
	if o != nil && o.SystemUser != nil {
		return true
	}

	return false
}

// SetSystemUser gets a reference to the given bool and assigns it to the SystemUser field.
func (o *BTUserAdminSummaryInfo) SetSystemUser(v bool) {
	o.SystemUser = &v
}

// GetTotpEnabled returns the TotpEnabled field value if set, zero value otherwise.
func (o *BTUserAdminSummaryInfo) GetTotpEnabled() bool {
	if o == nil || o.TotpEnabled == nil {
		var ret bool
		return ret
	}
	return *o.TotpEnabled
}

// GetTotpEnabledOk returns a tuple with the TotpEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BTUserAdminSummaryInfo) GetTotpEnabledOk() (*bool, bool) {
	if o == nil || o.TotpEnabled == nil {
		return nil, false
	}
	return o.TotpEnabled, true
}

// HasTotpEnabled returns a boolean if a field has been set.
func (o *BTUserAdminSummaryInfo) HasTotpEnabled() bool {
	if o != nil && o.TotpEnabled != nil {
		return true
	}

	return false
}

// SetTotpEnabled gets a reference to the given bool and assigns it to the TotpEnabled field.
func (o *BTUserAdminSummaryInfo) SetTotpEnabled(v bool) {
	o.TotpEnabled = &v
}

func (o BTUserAdminSummaryInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ViewRef != nil {
		toSerialize["viewRef"] = o.ViewRef
	}
	if o.Image != nil {
		toSerialize["image"] = o.Image
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.DocumentationName != nil {
		toSerialize["documentationName"] = o.DocumentationName
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.FirstName != nil {
		toSerialize["firstName"] = o.FirstName
	}
	if o.LastName != nil {
		toSerialize["lastName"] = o.LastName
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.DocumentationNameOverride != nil {
		toSerialize["documentationNameOverride"] = o.DocumentationNameOverride
	}
	if o.GlobalPermissions != nil {
		toSerialize["globalPermissions"] = o.GlobalPermissions
	}
	if o.InvitationState != nil {
		toSerialize["invitationState"] = o.InvitationState
	}
	if o.IsGuest != nil {
		toSerialize["isGuest"] = o.IsGuest
	}
	if o.IsLight != nil {
		toSerialize["isLight"] = o.IsLight
	}
	if o.LastLoginTime != nil {
		toSerialize["lastLoginTime"] = o.LastLoginTime
	}
	if o.PersonalMessageAllowed != nil {
		toSerialize["personalMessageAllowed"] = o.PersonalMessageAllowed
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.ActivePlanId != nil {
		toSerialize["activePlanId"] = o.ActivePlanId
	}
	if o.BillingUpdateRequired != nil {
		toSerialize["billingUpdateRequired"] = o.BillingUpdateRequired
	}
	if o.CompanyRoles != nil {
		toSerialize["companyRoles"] = o.CompanyRoles
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.ForumId != nil {
		toSerialize["forumId"] = o.ForumId
	}
	if o.SystemUser != nil {
		toSerialize["systemUser"] = o.SystemUser
	}
	if o.TotpEnabled != nil {
		toSerialize["totpEnabled"] = o.TotpEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableBTUserAdminSummaryInfo struct {
	value *BTUserAdminSummaryInfo
	isSet bool
}

func (v NullableBTUserAdminSummaryInfo) Get() *BTUserAdminSummaryInfo {
	return v.value
}

func (v *NullableBTUserAdminSummaryInfo) Set(val *BTUserAdminSummaryInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBTUserAdminSummaryInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBTUserAdminSummaryInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBTUserAdminSummaryInfo(val *BTUserAdminSummaryInfo) *NullableBTUserAdminSummaryInfo {
	return &NullableBTUserAdminSummaryInfo{value: val, isSet: true}
}

func (v NullableBTUserAdminSummaryInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBTUserAdminSummaryInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
