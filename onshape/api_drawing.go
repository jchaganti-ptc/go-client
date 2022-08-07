/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.151.5931-1b7b413e7f30
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

// DrawingApiService DrawingApi service
type DrawingApiService service

type ApiCreateDrawingAppElementRequest struct {
	ctx             context.Context
	ApiService      *DrawingApiService
	bTDrawingParams *BTDrawingParams
}

func (r ApiCreateDrawingAppElementRequest) BTDrawingParams(bTDrawingParams BTDrawingParams) ApiCreateDrawingAppElementRequest {
	r.bTDrawingParams = &bTDrawingParams
	return r
}

func (r ApiCreateDrawingAppElementRequest) Execute() (*BTDocumentElementInfo, *http.Response, error) {
	return r.ApiService.CreateDrawingAppElementExecute(r)
}

/*
CreateDrawingAppElement Method for CreateDrawingAppElement

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateDrawingAppElementRequest
*/
func (a *DrawingApiService) CreateDrawingAppElement(ctx context.Context) ApiCreateDrawingAppElementRequest {
	return ApiCreateDrawingAppElementRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return BTDocumentElementInfo
func (a *DrawingApiService) CreateDrawingAppElementExecute(r ApiCreateDrawingAppElementRequest) (*BTDocumentElementInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTDocumentElementInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrawingApiService.CreateDrawingAppElement")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/drawings/create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bTDrawingParams == nil {
		return localVarReturnValue, nil, reportError("bTDrawingParams is required and must be specified")
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
	localVarPostBody = r.bTDrawingParams
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
		var v BTDocumentElementInfo
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

type ApiCreateDrawingTranslationRequest struct {
	ctx                     context.Context
	ApiService              *DrawingApiService
	did                     string
	wv                      string
	wvid                    string
	eid                     string
	bTTranslateFormatParams *BTTranslateFormatParams
}

func (r ApiCreateDrawingTranslationRequest) BTTranslateFormatParams(bTTranslateFormatParams BTTranslateFormatParams) ApiCreateDrawingTranslationRequest {
	r.bTTranslateFormatParams = &bTTranslateFormatParams
	return r
}

func (r ApiCreateDrawingTranslationRequest) Execute() (*BTTranslationRequestInfo, *http.Response, error) {
	return r.ApiService.CreateDrawingTranslationExecute(r)
}

/*
CreateDrawingTranslation Method for CreateDrawingTranslation

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did
 @param wv
 @param wvid
 @param eid
 @return ApiCreateDrawingTranslationRequest
*/
func (a *DrawingApiService) CreateDrawingTranslation(ctx context.Context, did string, wv string, wvid string, eid string) ApiCreateDrawingTranslationRequest {
	return ApiCreateDrawingTranslationRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wv:         wv,
		wvid:       wvid,
		eid:        eid,
	}
}

// Execute executes the request
//  @return BTTranslationRequestInfo
func (a *DrawingApiService) CreateDrawingTranslationExecute(r ApiCreateDrawingTranslationRequest) (*BTTranslationRequestInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *BTTranslationRequestInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrawingApiService.CreateDrawingTranslation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/drawings/d/{did}/{wv}/{wvid}/e/{eid}/translations"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wv"+"}", url.PathEscape(parameterToString(r.wv, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wvid"+"}", url.PathEscape(parameterToString(r.wvid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bTTranslateFormatParams == nil {
		return localVarReturnValue, nil, reportError("bTTranslateFormatParams is required and must be specified")
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
	localVarPostBody = r.bTTranslateFormatParams
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
		var v BTTranslationRequestInfo
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

type ApiGetDrawingTranslatorFormatsRequest struct {
	ctx        context.Context
	ApiService *DrawingApiService
	did        string
	wid        string
	eid        string
}

func (r ApiGetDrawingTranslatorFormatsRequest) Execute() ([]BTModelFormatInfo, *http.Response, error) {
	return r.ApiService.GetDrawingTranslatorFormatsExecute(r)
}

/*
GetDrawingTranslatorFormats Retrieve translation formats by document ID, workspace ID, and tab ID.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param did
 @param wid
 @param eid
 @return ApiGetDrawingTranslatorFormatsRequest
*/
func (a *DrawingApiService) GetDrawingTranslatorFormats(ctx context.Context, did string, wid string, eid string) ApiGetDrawingTranslatorFormatsRequest {
	return ApiGetDrawingTranslatorFormatsRequest{
		ApiService: a,
		ctx:        ctx,
		did:        did,
		wid:        wid,
		eid:        eid,
	}
}

// Execute executes the request
//  @return []BTModelFormatInfo
func (a *DrawingApiService) GetDrawingTranslatorFormatsExecute(r ApiGetDrawingTranslatorFormatsRequest) ([]BTModelFormatInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []BTModelFormatInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DrawingApiService.GetDrawingTranslatorFormats")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/drawings/d/{did}/w/{wid}/e/{eid}/translationformats"
	localVarPath = strings.Replace(localVarPath, "{"+"did"+"}", url.PathEscape(parameterToString(r.did, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"wid"+"}", url.PathEscape(parameterToString(r.wid, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"eid"+"}", url.PathEscape(parameterToString(r.eid, "")), -1)

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
		var v []BTModelFormatInfo
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
