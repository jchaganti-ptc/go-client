/*
Onshape REST API

The Onshape REST API consumed by all client. # Authorization The simplest way to authorize and enable the **Try it out** functionality is to sign in to Onshape and use the current session. The **Authorize** button enables other authorization techniques. To ensure the current session isn't used when trying other authentication techniques, make sure to remove the Onshape cookie as per the instructions for your particular browser. Alternatively, a private or incognito window may be used. Here's [how to remove a specific cookie on Chrome](https://support.google.com/chrome/answer/95647#zippy=%2Cdelete-cookies-from-a-site). - **Current Session** authorization is enabled by default if the browser is already signed in to [Onshape](/). - **OAuth2** authorization uses an Onshape OAuth2 app created on the [Onshape Developer Portal](https://dev-portal.onshape.com/oauthApps). The redirect URL field should include `https://cad.onshape.com/glassworks/explorer/oauth2-redirect.html`. - **API Key** authorization using basic authentication is also available. The keys can be generated in the [Onshape Developer Portal](https://dev-portal.onshape.com/keys). In the authentication dialog, enter the access key in the `Username` field, and enter the secret key in the `Password` field. Basic authentication should only be used during the development process since sharing API Keys provides the same level of access as a username and password.

API version: 1.152.6228-2857ab16a033
Contact: api-support@onshape.zendesk.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package onshape

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/oauth2"
)

var (
	jsonCheck = regexp.MustCompile(`(?i:(?:application|text)/(?:vnd\.[^;]+\+)?json)`)
	xmlCheck  = regexp.MustCompile(`(?i:(?:application|text)/xml)`)
)

// APIClient manages communication with the Onshape REST API API v1.152.6228-2857ab16a033
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	APIApplicationApi *APIApplicationApiService

	AccountApi *AccountApiService

	AliasApi *AliasApiService

	AppAssociativeDataApi *AppAssociativeDataApiService

	AppElementApi *AppElementApiService

	AssemblyApi *AssemblyApiService

	BillingApi *BillingApiService

	BlobElementApi *BlobElementApiService

	CommentApi *CommentApiService

	CompanyApi *CompanyApiService

	DocumentApi *DocumentApiService

	DrawingApi *DrawingApiService

	ElementApi *ElementApiService

	EventApi *EventApiService

	ExportRuleApi *ExportRuleApiService

	FeatureStudioApi *FeatureStudioApiService

	FolderApi *FolderApiService

	InsertableApi *InsertableApiService

	MetadataApi *MetadataApiService

	MetadataCategoryApi *MetadataCategoryApiService

	OpenApiApi *OpenApiApiService

	PartApi *PartApiService

	PartNumberApi *PartNumberApiService

	PartStudioApi *PartStudioApiService

	PublicationApi *PublicationApiService

	ReleasePackageApi *ReleasePackageApiService

	RevisionApi *RevisionApiService

	SketchApi *SketchApiService

	TeamApi *TeamApiService

	ThumbnailApi *ThumbnailApiService

	TranslationApi *TranslationApiService

	UserApi *UserApiService

	VariablesApi *VariablesApiService

	VersionApi *VersionApiService

	WebhookApi *WebhookApiService

	WorkflowApi *WorkflowApiService

	WorkflowableTestObjectApi *WorkflowableTestObjectApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.APIApplicationApi = (*APIApplicationApiService)(&c.common)
	c.AccountApi = (*AccountApiService)(&c.common)
	c.AliasApi = (*AliasApiService)(&c.common)
	c.AppAssociativeDataApi = (*AppAssociativeDataApiService)(&c.common)
	c.AppElementApi = (*AppElementApiService)(&c.common)
	c.AssemblyApi = (*AssemblyApiService)(&c.common)
	c.BillingApi = (*BillingApiService)(&c.common)
	c.BlobElementApi = (*BlobElementApiService)(&c.common)
	c.CommentApi = (*CommentApiService)(&c.common)
	c.CompanyApi = (*CompanyApiService)(&c.common)
	c.DocumentApi = (*DocumentApiService)(&c.common)
	c.DrawingApi = (*DrawingApiService)(&c.common)
	c.ElementApi = (*ElementApiService)(&c.common)
	c.EventApi = (*EventApiService)(&c.common)
	c.ExportRuleApi = (*ExportRuleApiService)(&c.common)
	c.FeatureStudioApi = (*FeatureStudioApiService)(&c.common)
	c.FolderApi = (*FolderApiService)(&c.common)
	c.InsertableApi = (*InsertableApiService)(&c.common)
	c.MetadataApi = (*MetadataApiService)(&c.common)
	c.MetadataCategoryApi = (*MetadataCategoryApiService)(&c.common)
	c.OpenApiApi = (*OpenApiApiService)(&c.common)
	c.PartApi = (*PartApiService)(&c.common)
	c.PartNumberApi = (*PartNumberApiService)(&c.common)
	c.PartStudioApi = (*PartStudioApiService)(&c.common)
	c.PublicationApi = (*PublicationApiService)(&c.common)
	c.ReleasePackageApi = (*ReleasePackageApiService)(&c.common)
	c.RevisionApi = (*RevisionApiService)(&c.common)
	c.SketchApi = (*SketchApiService)(&c.common)
	c.TeamApi = (*TeamApiService)(&c.common)
	c.ThumbnailApi = (*ThumbnailApiService)(&c.common)
	c.TranslationApi = (*TranslationApiService)(&c.common)
	c.UserApi = (*UserApiService)(&c.common)
	c.VariablesApi = (*VariablesApiService)(&c.common)
	c.VersionApi = (*VersionApiService)(&c.common)
	c.WebhookApi = (*WebhookApiService)(&c.common)
	c.WorkflowApi = (*WorkflowApiService)(&c.common)
	c.WorkflowableTestObjectApi = (*WorkflowableTestObjectApiService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insensitive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	} else if t, ok := obj.(time.Time); ok {
		return t.Format(time.RFC3339)
	}

	return fmt.Sprintf("%v", obj)
}

// helper for converting interface{} parameters to json strings
func parameterToJson(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	if c.cfg.Debug {
		dump, err := httputil.DumpRequestOut(request, true)
		if err != nil {
			return nil, err
		}
		log.Printf("\n%s\n", string(dump))
	}

	resp, err := c.cfg.HTTPClient.Do(request)
	if err != nil {
		return resp, err
	}

	if c.cfg.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
		}
		log.Printf("\n%s\n", string(dump))
	}
	return resp, err
}

// Allow modification of underlying config for alternate implementations and testing
// Caution: modifying the configuration while live can cause data races and potentially unwanted behavior
func (c *APIClient) GetConfig() *Configuration {
	return c.cfg
}

type formFile struct {
	fileData     io.Reader
	fileName     string
	formFileName string
}

type multipartStreamer struct {
	buffer      *bytes.Buffer
	reader      io.Reader
	multiwriter *multipart.Writer
	last        int
}

func newMultipartStreamer() *multipartStreamer {
	buffer := &bytes.Buffer{}
	return &multipartStreamer{
		buffer:      buffer,
		reader:      bytes.NewReader(buffer.Bytes()[0:0]),
		multiwriter: multipart.NewWriter(buffer),
		last:        0,
	}
}

func (ms *multipartStreamer) appendStream(reader io.Reader) {
	ms.reader = io.MultiReader(
		ms.reader,
		reader)
}

func (ms *multipartStreamer) partition() {
	ms.reader = io.MultiReader(
		ms.reader,
		bytes.NewReader(ms.buffer.Bytes()[ms.last:]))
	ms.last = len(ms.buffer.Bytes())
}

func (ms *multipartStreamer) Close() {
	ms.multiwriter.Close()
	ms.partition()
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	formFiles []formFile) (localVarRequest *http.Request, err error) {

	var body io.Reader

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(formFiles) > 0) {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		w := newMultipartStreamer()

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w.multiwriter, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.multiwriter.WriteField(k, iv)
				}
			}
		}
		for _, formFile := range formFiles {
			if formFile.fileName != "" {
				w.multiwriter.Boundary()
				_, err := w.multiwriter.CreateFormFile(formFile.formFileName, filepath.Base(formFile.fileName))
				if err != nil {
					return nil, err
				}
				w.partition()
				w.appendStream(formFile.fileData)
			}
		}

		// Set the Boundary in the Content-Type
		headerParams["Content-Type"] = w.multiwriter.FormDataContentType()

		w.Close()
		body = w.reader
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.(*bytes.Buffer).WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.(*bytes.Buffer).Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		url.Host = c.cfg.Host
	}

	// Override request scheme, if applicable
	if c.cfg.Scheme != "" {
		url.Scheme = c.cfg.Scheme
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers[h] = []string{v}
		}
		localVarRequest.Header = headers
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// OAuth2 authentication
		if tok, ok := ctx.Value(ContextOAuth2).(oauth2.TokenSource); ok {
			// We were able to grab an oauth2 token from the context
			var latestToken *oauth2.Token
			if latestToken, err = tok.Token(); err != nil {
				return nil, err
			}

			latestToken.SetAuthHeader(localVarRequest)
		}

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(ContextBasicAuth).(BasicAuth); ok {
			localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
		}

		// AccessToken Authentication
		if auth, ok := ctx.Value(ContextAccessToken).(string); ok {
			localVarRequest.Header.Add("Authorization", "Bearer "+auth)
		}

	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}
	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, br *io.ReadCloser, contentType string) (err error) {
	if f, ok := v.(*HttpFile); ok {
		*f = NewHttpFileFromReader("", *br)
		return
	}
	if f, ok := v.(**HttpFile); ok {
		fil := NewHttpFileFromReader("", *br)
		*f = &fil
		return
	}

	b, err := ioutil.ReadAll(*br)
	(*br).Close()

	if err != nil {
		return
	}

	*br = ioutil.NopCloser(bytes.NewBuffer(b))

	if len(b) == 0 {
		return nil
	}
	if s, ok := v.(*string); ok {
		*s = string(b)
		return nil
	}
	if xmlCheck.MatchString(contentType) {
		if err = xml.Unmarshal(b, v); err != nil {
			return err
		}
		return nil
	}
	if jsonCheck.MatchString(contentType) {
		if actualObj, ok := v.(interface{ GetActualInstance() interface{} }); ok { // oneOf, anyOf schemas
			if unmarshalObj, ok := actualObj.(interface{ UnmarshalJSON([]byte) error }); ok { // make sure it has UnmarshalJSON defined
				if err = unmarshalObj.UnmarshalJSON(b); err != nil {
					return err
				}
			} else {
				return errors.New("Unknown type with GetActualInstance but no unmarshalObj.UnmarshalJSON defined")
			}
		} else if err = json.Unmarshal(b, v); err != nil { // simple model
			return err
		}
		return nil
	}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// A wrapper for strict JSON decoding
func newStrictDecoder(data []byte) *json.Decoder {
	dec := json.NewDecoder(bytes.NewBuffer(data))
	dec.DisallowUnknownFields()
	return dec
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf io.Reader, err error) {
	bbf := &bytes.Buffer{}

	if reader, ok := body.(io.Reader); ok {
		_, err = bbf.ReadFrom(reader)
	} else if fp, ok := body.(*HttpFile); ok {
		return fp.Data, nil
	} else if fp, ok := body.(**HttpFile); ok {
		return (*fp).Data, nil
	} else if b, ok := body.([]byte); ok {
		_, err = bbf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bbf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bbf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bbf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		err = xml.NewEncoder(bbf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bbf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}
	bodyBuf = bbf
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		} else {
			expires = now.Add(lifetime)
		}
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

type HttpFile struct {
	Name string
	Data io.ReadCloser
}

func NewHttpFile(name string, data []byte) HttpFile {
	buf := io.NopCloser(bytes.NewBuffer(data))
	return HttpFile{name, buf}
}

func NewHttpFileFromReader(name string, data io.ReadCloser) HttpFile {
	return HttpFile{name, data}
}

func NewHttpFileFromOsFile(file *os.File) HttpFile {
	name := file.Name()
	return HttpFile{name, file}
}

func (h *HttpFile) Close() {
	h.Data.Close()
}

// GenericOpenAPIError Provides access to the body, error and model on returned errors.
type GenericOpenAPIError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericOpenAPIError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericOpenAPIError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericOpenAPIError) Model() interface{} {
	return e.model
}
