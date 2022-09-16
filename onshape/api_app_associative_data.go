/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.153.6489-39ccb1a53c2e
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
	"reflect"
	"strings"
)

// AppAssociativeDataApiService AppAssociativeDataApi service
type AppAssociativeDataApiService service

type ApiCopyAssociativeDataRequest struct {
	ctx                                                    context.Context
	ApiService                                             *AppAssociativeDataApiService
	did                                                    string
	wid                                                    string
	eid                                                    string
	bTAppElementParamsArrayBTCopyViewAssociativeDataParams *BTAppElementParamsArrayBTCopyViewAssociativeDataParams
}

func (r ApiCopyAssociativeDataRequest) BTAppElementParamsArrayBTCopyViewAssociativeDataParams(bTAppElementParamsArrayBTCopyViewAssociativeDataParams BTAppElementParamsArrayBTCopyViewAssociativeDataParams) ApiCopyAssociativeDataRequest {
	r.bTAppElementParamsArrayBTCopyViewAssociativeDataParams = &bTAppElementParamsArrayBTCopyViewAssociativeDataParams
	return r
}

func (r ApiCopyAssociativeDataRequest) Execute() (*BTAppAssociativeDataArrayInfo, *http.Response, error) {
	return r.ApiService.CopyAssociativeDataExecute(r)
}

/*
CopyAssociativeData Copy associative data between sub-views inside this application tab by document ID, workspace ID, and tab ID. Useful if the application has multiple sub-components; for example, Drawing views.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did
 @param wid
 @param eid
 @return ApiCopyAssociativeDataRequest
*/
func (a *AppAssociativeDataApiService) CopyAssociativeData(ctx context.Context, did string, wid string, eid string) ApiCopyAssociativeDataRequest {
	return ApiCopyAssociativeDataRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wid:        wid,
		eid:        eid,
	}
}

// Execute executes the request
//  @return BTAppAssociativeDataArrayInfo
func (a *AppAssociativeDataApiService) CopyAssociativeDataExecute(r ApiCopyAssociativeDataRequest) (*BTAppAssociativeDataArrayInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTAppAssociativeDataArrayInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppAssociativeDataApiService.CopyAssociativeData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/appelements/d/{did}/w/{wid}/e/{eid}/copyassociativedata"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wid"+"}", url.PathEscape(parameterToString(r.wid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;charset=UTF-8; qs=0.09"}

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
	// body params
	localVarPostBody = r.bTAppElementParamsArrayBTCopyViewAssociativeDataParams
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
		var v BTAppAssociativeDataArrayInfo
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

type ApiDeleteAssociativeDataRequest struct {
	ctx                  context.Context
	ApiService           *AppAssociativeDataApiService
	did                  string
	eid                  string
	wvm                  string
	wvmid                string
	transactionId        *string
	parentChangeId       *string
	associativeDataId    *[]string
	elementId            *string
	viewId               *string
	microversionId       *string
	documentMicroversion *string
	deterministicId      *string
	featureId            *string
	entityId             *string
	occurrenceId         *string
}

func (r ApiDeleteAssociativeDataRequest) TransactionId(transactionId string) ApiDeleteAssociativeDataRequest {
	r.transactionId = &transactionId
	return r
}

func (r ApiDeleteAssociativeDataRequest) ParentChangeId(parentChangeId string) ApiDeleteAssociativeDataRequest {
	r.parentChangeId = &parentChangeId
	return r
}

func (r ApiDeleteAssociativeDataRequest) AssociativeDataId(associativeDataId []string) ApiDeleteAssociativeDataRequest {
	r.associativeDataId = &associativeDataId
	return r
}

func (r ApiDeleteAssociativeDataRequest) ElementId(elementId string) ApiDeleteAssociativeDataRequest {
	r.elementId = &elementId
	return r
}

func (r ApiDeleteAssociativeDataRequest) ViewId(viewId string) ApiDeleteAssociativeDataRequest {
	r.viewId = &viewId
	return r
}

func (r ApiDeleteAssociativeDataRequest) MicroversionId(microversionId string) ApiDeleteAssociativeDataRequest {
	r.microversionId = &microversionId
	return r
}

func (r ApiDeleteAssociativeDataRequest) DocumentMicroversion(documentMicroversion string) ApiDeleteAssociativeDataRequest {
	r.documentMicroversion = &documentMicroversion
	return r
}

func (r ApiDeleteAssociativeDataRequest) DeterministicId(deterministicId string) ApiDeleteAssociativeDataRequest {
	r.deterministicId = &deterministicId
	return r
}

func (r ApiDeleteAssociativeDataRequest) FeatureId(featureId string) ApiDeleteAssociativeDataRequest {
	r.featureId = &featureId
	return r
}

func (r ApiDeleteAssociativeDataRequest) EntityId(entityId string) ApiDeleteAssociativeDataRequest {
	r.entityId = &entityId
	return r
}

func (r ApiDeleteAssociativeDataRequest) OccurrenceId(occurrenceId string) ApiDeleteAssociativeDataRequest {
	r.occurrenceId = &occurrenceId
	return r
}

func (r ApiDeleteAssociativeDataRequest) Execute() (*BTAppElementBasicInfo, *http.Response, error) {
	return r.ApiService.DeleteAssociativeDataExecute(r)
}

/*
DeleteAssociativeData Delete associative data for this application tab by document ID, workspace or version or microversion ID, and tab ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did
 @param eid
 @param wvm
 @param wvmid
 @return ApiDeleteAssociativeDataRequest
*/
func (a *AppAssociativeDataApiService) DeleteAssociativeData(ctx context.Context, did string, eid string, wvm string, wvmid string) ApiDeleteAssociativeDataRequest {
	return ApiDeleteAssociativeDataRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		eid:        eid,
		wvm:        wvm,
		wvmid:      wvmid,
	}
}

// Execute executes the request
//  @return BTAppElementBasicInfo
func (a *AppAssociativeDataApiService) DeleteAssociativeDataExecute(r ApiDeleteAssociativeDataRequest) (*BTAppElementBasicInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodDelete
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTAppElementBasicInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppAssociativeDataApiService.DeleteAssociativeData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/associativedata"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvm"+"}", url.PathEscape(parameterToString(r.wvm, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvmid"+"}", url.PathEscape(parameterToString(r.wvmid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.transactionId != nil {
		localVarQueryParams.Add("transactionId", parameterToString(*r.transactionId, ""))
	}
	if r.parentChangeId != nil {
		localVarQueryParams.Add("parentChangeId", parameterToString(*r.parentChangeId, ""))
	}
	if r.associativeDataId != nil {
		t := *r.associativeDataId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("associativeDataId", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("associativeDataId", parameterToString(t, "multi"))
		}
	}
	if r.elementId != nil {
		localVarQueryParams.Add("elementId", parameterToString(*r.elementId, ""))
	}
	if r.viewId != nil {
		localVarQueryParams.Add("viewId", parameterToString(*r.viewId, ""))
	}
	if r.microversionId != nil {
		localVarQueryParams.Add("microversionId", parameterToString(*r.microversionId, ""))
	}
	if r.documentMicroversion != nil {
		localVarQueryParams.Add("documentMicroversion", parameterToString(*r.documentMicroversion, ""))
	}
	if r.deterministicId != nil {
		localVarQueryParams.Add("deterministicId", parameterToString(*r.deterministicId, ""))
	}
	if r.featureId != nil {
		localVarQueryParams.Add("featureId", parameterToString(*r.featureId, ""))
	}
	if r.entityId != nil {
		localVarQueryParams.Add("entityId", parameterToString(*r.entityId, ""))
	}
	if r.occurrenceId != nil {
		localVarQueryParams.Add("occurrenceId", parameterToString(*r.occurrenceId, ""))
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
		var v BTAppElementBasicInfo
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

type ApiGetAssociativeDataRequest struct {
	ctx                  context.Context
	ApiService           *AppAssociativeDataApiService
	did                  string
	wvm                  string
	wvmid                string
	eid                  string
	linkDocumentId       *string
	transactionId        *string
	changeId             *string
	associativeDataId    *[]string
	elementId            *string
	viewId               *string
	microversionId       *string
	documentMicroversion *string
	deterministicId      *string
	featureId            *string
	entityId             *string
	occurrenceId         *string
	returnIdTags         *bool
}

// The id of the document through which the above document should be accessed; only applicable when accessing a version of the document. This allows a user who has access to document a to see data from document b, as long as document b has been linked to document a by a user who has permission to both.
func (r ApiGetAssociativeDataRequest) LinkDocumentId(linkDocumentId string) ApiGetAssociativeDataRequest {
	r.linkDocumentId = &linkDocumentId
	return r
}

func (r ApiGetAssociativeDataRequest) TransactionId(transactionId string) ApiGetAssociativeDataRequest {
	r.transactionId = &transactionId
	return r
}

func (r ApiGetAssociativeDataRequest) ChangeId(changeId string) ApiGetAssociativeDataRequest {
	r.changeId = &changeId
	return r
}

func (r ApiGetAssociativeDataRequest) AssociativeDataId(associativeDataId []string) ApiGetAssociativeDataRequest {
	r.associativeDataId = &associativeDataId
	return r
}

func (r ApiGetAssociativeDataRequest) ElementId(elementId string) ApiGetAssociativeDataRequest {
	r.elementId = &elementId
	return r
}

func (r ApiGetAssociativeDataRequest) ViewId(viewId string) ApiGetAssociativeDataRequest {
	r.viewId = &viewId
	return r
}

func (r ApiGetAssociativeDataRequest) MicroversionId(microversionId string) ApiGetAssociativeDataRequest {
	r.microversionId = &microversionId
	return r
}

func (r ApiGetAssociativeDataRequest) DocumentMicroversion(documentMicroversion string) ApiGetAssociativeDataRequest {
	r.documentMicroversion = &documentMicroversion
	return r
}

func (r ApiGetAssociativeDataRequest) DeterministicId(deterministicId string) ApiGetAssociativeDataRequest {
	r.deterministicId = &deterministicId
	return r
}

func (r ApiGetAssociativeDataRequest) FeatureId(featureId string) ApiGetAssociativeDataRequest {
	r.featureId = &featureId
	return r
}

func (r ApiGetAssociativeDataRequest) EntityId(entityId string) ApiGetAssociativeDataRequest {
	r.entityId = &entityId
	return r
}

func (r ApiGetAssociativeDataRequest) OccurrenceId(occurrenceId string) ApiGetAssociativeDataRequest {
	r.occurrenceId = &occurrenceId
	return r
}

func (r ApiGetAssociativeDataRequest) ReturnIdTags(returnIdTags bool) ApiGetAssociativeDataRequest {
	r.returnIdTags = &returnIdTags
	return r
}

func (r ApiGetAssociativeDataRequest) Execute() (*BTAppAssociativeDataArrayInfo, *http.Response, error) {
	return r.ApiService.GetAssociativeDataExecute(r)
}

/*
GetAssociativeData Retrieve associative data for the application tab by document ID, workspace or version or microversion ID, and tab ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did The id of the document in which to perform the operation.
 @param wvm Indicates which of workspace id, version id, or document microversion id is specified below.
 @param wvmid The id of the workspace, version, or document microversion in which the operation should be performed.
 @param eid The id of the element in which to perform the operation.
 @return ApiGetAssociativeDataRequest
*/
func (a *AppAssociativeDataApiService) GetAssociativeData(ctx context.Context, did string, wvm string, wvmid string, eid string) ApiGetAssociativeDataRequest {
	return ApiGetAssociativeDataRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wvm:        wvm,
		wvmid:      wvmid,
		eid:        eid,
	}
}

// Execute executes the request
//  @return BTAppAssociativeDataArrayInfo
func (a *AppAssociativeDataApiService) GetAssociativeDataExecute(r ApiGetAssociativeDataRequest) (*BTAppAssociativeDataArrayInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTAppAssociativeDataArrayInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppAssociativeDataApiService.GetAssociativeData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/associativedata"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvm"+"}", url.PathEscape(parameterToString(r.wvm, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvmid"+"}", url.PathEscape(parameterToString(r.wvmid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.linkDocumentId != nil {
		localVarQueryParams.Add("linkDocumentId", parameterToString(*r.linkDocumentId, ""))
	}
	if r.transactionId != nil {
		localVarQueryParams.Add("transactionId", parameterToString(*r.transactionId, ""))
	}
	if r.changeId != nil {
		localVarQueryParams.Add("changeId", parameterToString(*r.changeId, ""))
	}
	if r.associativeDataId != nil {
		t := *r.associativeDataId
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("associativeDataId", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("associativeDataId", parameterToString(t, "multi"))
		}
	}
	if r.elementId != nil {
		localVarQueryParams.Add("elementId", parameterToString(*r.elementId, ""))
	}
	if r.viewId != nil {
		localVarQueryParams.Add("viewId", parameterToString(*r.viewId, ""))
	}
	if r.microversionId != nil {
		localVarQueryParams.Add("microversionId", parameterToString(*r.microversionId, ""))
	}
	if r.documentMicroversion != nil {
		localVarQueryParams.Add("documentMicroversion", parameterToString(*r.documentMicroversion, ""))
	}
	if r.deterministicId != nil {
		localVarQueryParams.Add("deterministicId", parameterToString(*r.deterministicId, ""))
	}
	if r.featureId != nil {
		localVarQueryParams.Add("featureId", parameterToString(*r.featureId, ""))
	}
	if r.entityId != nil {
		localVarQueryParams.Add("entityId", parameterToString(*r.entityId, ""))
	}
	if r.occurrenceId != nil {
		localVarQueryParams.Add("occurrenceId", parameterToString(*r.occurrenceId, ""))
	}
	if r.returnIdTags != nil {
		localVarQueryParams.Add("returnIdTags", parameterToString(*r.returnIdTags, ""))
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
		var v BTAppAssociativeDataArrayInfo
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

type ApiPostAssociativeDataRequest struct {
	ctx        context.Context
	ApiService *AppAssociativeDataApiService
	did        string
	eid        string
	wvm        string
	wvmid      string
	body       *string
}

func (r ApiPostAssociativeDataRequest) Body(body string) ApiPostAssociativeDataRequest {
	r.body = &body
	return r
}

func (r ApiPostAssociativeDataRequest) Execute() (*BTAppAssociativeDataArrayInfo, *http.Response, error) {
	return r.ApiService.PostAssociativeDataExecute(r)
}

/*
PostAssociativeData Update associative data for an application tab by document ID, workspace or version or microversion ID, and tab ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did
 @param eid
 @param wvm
 @param wvmid
 @return ApiPostAssociativeDataRequest
*/
func (a *AppAssociativeDataApiService) PostAssociativeData(ctx context.Context, did string, eid string, wvm string, wvmid string) ApiPostAssociativeDataRequest {
	return ApiPostAssociativeDataRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		eid:        eid,
		wvm:        wvm,
		wvmid:      wvmid,
	}
}

// Execute executes the request
//  @return BTAppAssociativeDataArrayInfo
func (a *AppAssociativeDataApiService) PostAssociativeDataExecute(r ApiPostAssociativeDataRequest) (*BTAppAssociativeDataArrayInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTAppAssociativeDataArrayInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AppAssociativeDataApiService.PostAssociativeData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/appelements/d/{did}/{wvm}/{wvmid}/e/{eid}/associativedata"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvm"+"}", url.PathEscape(parameterToString(r.wvm, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvmid"+"}", url.PathEscape(parameterToString(r.wvmid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json;charset=UTF-8; qs=0.09"}

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
	// body params
	localVarPostBody = r.body
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
		var v BTAppAssociativeDataArrayInfo
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
