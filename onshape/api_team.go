/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.173.27313-c3b730633f3c
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// TeamApiService TeamApi service
type TeamApiService service

type ApiFindRequest struct {
	ctx                      context.Context
	ApiService               *TeamApiService
	prefix                   *string
	uid                      *string
	companyId                *string
	includeCompanyOwnedTeams *bool
}

func (r ApiFindRequest) Prefix(prefix string) ApiFindRequest {
	r.prefix = &prefix
	return r
}

func (r ApiFindRequest) Uid(uid string) ApiFindRequest {
	r.uid = &uid
	return r
}

func (r ApiFindRequest) CompanyId(companyId string) ApiFindRequest {
	r.companyId = &companyId
	return r
}

func (r ApiFindRequest) IncludeCompanyOwnedTeams(includeCompanyOwnedTeams bool) ApiFindRequest {
	r.includeCompanyOwnedTeams = &includeCompanyOwnedTeams
	return r
}

func (r ApiFindRequest) Execute() (*BTGlobalTreeNodeListResponseBTTeamInfo, *http.Response, error) {
	return r.ApiService.FindExecute(r)
}

/*
Find Get a list of all teams the current user belongs to.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFindRequest
*/
func (a *TeamApiService) Find(ctx context.Context) ApiFindRequest {
	return ApiFindRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return BTGlobalTreeNodeListResponseBTTeamInfo
func (a *TeamApiService) FindExecute(r ApiFindRequest) (*BTGlobalTreeNodeListResponseBTTeamInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTGlobalTreeNodeListResponseBTTeamInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamApiService.Find")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.prefix != nil {
		localVarQueryParams.Add("prefix", parameterToString(*r.prefix, ""))
	}
	if r.uid != nil {
		localVarQueryParams.Add("uid", parameterToString(*r.uid, ""))
	}
	if r.companyId != nil {
		localVarQueryParams.Add("companyId", parameterToString(*r.companyId, ""))
	}
	if r.includeCompanyOwnedTeams != nil {
		localVarQueryParams.Add("includeCompanyOwnedTeams", parameterToString(*r.includeCompanyOwnedTeams, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTGlobalTreeNodeListResponseBTTeamInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetMembersRequest struct {
	ctx        context.Context
	ApiService *TeamApiService
	tid        string
	sortColumn *string
	sortOrder  *string
	offset     *int32
	limit      *int32
	q          *string
}

func (r ApiGetMembersRequest) SortColumn(sortColumn string) ApiGetMembersRequest {
	r.sortColumn = &sortColumn
	return r
}

func (r ApiGetMembersRequest) SortOrder(sortOrder string) ApiGetMembersRequest {
	r.sortOrder = &sortOrder
	return r
}

func (r ApiGetMembersRequest) Offset(offset int32) ApiGetMembersRequest {
	r.offset = &offset
	return r
}

func (r ApiGetMembersRequest) Limit(limit int32) ApiGetMembersRequest {
	r.limit = &limit
	return r
}

func (r ApiGetMembersRequest) Q(q string) ApiGetMembersRequest {
	r.q = &q
	return r
}

func (r ApiGetMembersRequest) Execute() (*BTListResponseBTTeamMemberInfo, *http.Response, error) {
	return r.ApiService.GetMembersExecute(r)
}

/*
GetMembers Get a list of a team's members.

Returns a maximum of 20 per page.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tid
 @return ApiGetMembersRequest
*/
func (a *TeamApiService) GetMembers(ctx context.Context, tid string) ApiGetMembersRequest {
	return ApiGetMembersRequest{
		ApiService: a,
		ctx:        ctx,
		tid:        tid,
	}
}

// Execute executes the request
//  @return BTListResponseBTTeamMemberInfo
func (a *TeamApiService) GetMembersExecute(r ApiGetMembersRequest) (*BTListResponseBTTeamMemberInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTListResponseBTTeamMemberInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamApiService.GetMembers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams/{tid}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"tid"+"}", url.PathEscape(parameterToString(r.tid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.sortColumn != nil {
		localVarQueryParams.Add("sortColumn", parameterToString(*r.sortColumn, ""))
	}
	if r.sortOrder != nil {
		localVarQueryParams.Add("sortOrder", parameterToString(*r.sortOrder, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.q != nil {
		localVarQueryParams.Add("q", parameterToString(*r.q, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTListResponseBTTeamMemberInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetTeamRequest struct {
	ctx        context.Context
	ApiService *TeamApiService
	tid        string
}

func (r ApiGetTeamRequest) Execute() (*BTTeamInfo, *http.Response, error) {
	return r.ApiService.GetTeamExecute(r)
}

/*
GetTeam Get team information by team ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tid
 @return ApiGetTeamRequest
*/
func (a *TeamApiService) GetTeam(ctx context.Context, tid string) ApiGetTeamRequest {
	return ApiGetTeamRequest{
		ApiService: a,
		ctx:        ctx,
		tid:        tid,
	}
}

// Execute executes the request
//  @return BTTeamInfo
func (a *TeamApiService) GetTeamExecute(r ApiGetTeamRequest) (*BTTeamInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTTeamInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamApiService.GetTeam")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/teams/{tid}"
	localVarPath = strings.Replace(localVarPath, "{"+"tid"+"}", url.PathEscape(parameterToString(r.tid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json;charset=UTF-8; qs=0.09"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	var _ io.Reader

	if localVarHTTPResponse.StatusCode >= 300 {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v BTTeamInfo
		err = a.client.decode(&v, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, &localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))

	if err != nil {
		localVarBody, _ := ioutil.ReadAll(localVarHTTPResponse.Body)

		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
