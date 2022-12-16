/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type APIObservation struct {

	// ID of the APIG API.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Time when the API is registered, in UTC format.
	RegisterTime *string `json:"registerTime,omitempty" tf:"register_time,omitempty"`

	// Time when the API was last modified, in UTC format.
	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type APIParameters struct {

	// Specifies ID of the front-end custom authorizer.
	// +kubebuilder:validation:Optional
	AuthorizerID *string `json:"authorizerId,omitempty" tf:"authorizer_id,omitempty"`

	// Specifies an array of one or more backend parameters.
	// The object structure is documented below. The maximum of request parameters is 50.
	// +kubebuilder:validation:Optional
	BackendParams []BackendParamsParameters `json:"backendParams,omitempty" tf:"backend_params,omitempty"`

	// Specifies the description of the API request body, which can be an example
	// request body, media type or parameters. The request body does not exceed 20,480 characters. Chinese characters must be
	// in UTF-8 or Unicode format.
	// +kubebuilder:validation:Optional
	BodyDescription *string `json:"bodyDescription,omitempty" tf:"body_description,omitempty"`

	// Specifies whether CORS is supported, default to false.
	// +kubebuilder:validation:Optional
	Cors *bool `json:"cors,omitempty" tf:"cors,omitempty"`

	// Specifies the API description, which can contain a maximum of 255 characters. The
	// Chinese characters must be in UTF-8 or Unicode format.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the example response for a successful request. Ensure that the
	// response does not exceed 20,480 characters. Chinese characters must be in UTF-8 or Unicode format.
	// +kubebuilder:validation:Optional
	FailureResponse *string `json:"failureResponse,omitempty" tf:"failure_response,omitempty"`

	// Specifies the function graph backend details. The object
	// structure is documented below. Changing this will create a new API resource.
	// +kubebuilder:validation:Optional
	FuncGraph []FuncGraphParameters `json:"funcGraph,omitempty" tf:"func_graph,omitempty"`

	// Specifies the Mock policy backends. The maximum of the policy is 5.
	// The object structure is documented below.
	// +kubebuilder:validation:Optional
	FuncGraphPolicy []FuncGraphPolicyParameters `json:"funcGraphPolicy,omitempty" tf:"func_graph_policy,omitempty"`

	// Specifies an ID of the APIG group to which the API belongs to.
	// +crossplane:generate:reference:type=Group
	// +kubebuilder:validation:Optional
	GroupID *string `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Reference to a Group to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDRef *v1.Reference `json:"groupIdRef,omitempty" tf:"-"`

	// Selector for a Group to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDSelector *v1.Selector `json:"groupIdSelector,omitempty" tf:"-"`

	// Specifies an ID of the APIG dedicated instance to which the API belongs
	// to. Changing this will create a new API resource.
	// +crossplane:generate:reference:type=Instance
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the route matching mode. The valid value are Exact and Prefix,
	// default to Exact.
	// +kubebuilder:validation:Optional
	Matching *string `json:"matching,omitempty" tf:"matching,omitempty"`

	// Specifies the mock backend details. The object structure is documented
	// below. Changing this will create a new API resource.
	// +kubebuilder:validation:Optional
	Mock []MockParameters `json:"mock,omitempty" tf:"mock,omitempty"`

	// Specifies the Mock policy backends. The maximum of the policy is 5.
	// The object structure is documented below.
	// +kubebuilder:validation:Optional
	MockPolicy []MockPolicyParameters `json:"mockPolicy,omitempty" tf:"mock_policy,omitempty"`

	// Specifies the API name, which can consists of 3 to 64 characters, starting with a letter.
	// Only letters, digits and underscores (_) are allowed. Chinese characters must be in UTF-8 or Unicode format.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the region in which to create the API resource. If omitted, the
	// provider-level region will be used. Changing this will create a new API resource.
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// Specifies the request method of the API. The valid values are GET, POST
	// , PUT, DELETE, HEAD, PATCH, OPTIONS and ANY.
	// +kubebuilder:validation:Required
	RequestMethod *string `json:"requestMethod" tf:"request_method,omitempty"`

	// Specifies an array of one or more request parameters of the front-end. The maximum
	// of request parameters is 50. The object structure is documented below.
	// +kubebuilder:validation:Optional
	RequestParams []RequestParamsParameters `json:"requestParams,omitempty" tf:"request_params,omitempty"`

	// Specifies the request address, which can contain a maximum of 512 characters
	// request parameters enclosed with brackets ({}).
	// +kubebuilder:validation:Required
	RequestPath *string `json:"requestPath" tf:"request_path,omitempty"`

	// Specifies the request protocol of the API. The valid value are
	// HTTP, HTTPS and BOTH.
	// +kubebuilder:validation:Required
	RequestProtocol *string `json:"requestProtocol" tf:"request_protocol,omitempty"`

	// Specifies the APIG group response ID.
	// +kubebuilder:validation:Optional
	ResponseID *string `json:"responseId,omitempty" tf:"response_id,omitempty"`

	// Specifies the security authentication mode. The valid values are
	// NONE, APP and IAM, default to NONE.
	// +kubebuilder:validation:Optional
	SecurityAuthentication *string `json:"securityAuthentication,omitempty" tf:"security_authentication,omitempty"`

	// Specifies whether AppCode authentication is enabled. The applicaiton code
	// must located in the header when simple_authentication is true.
	// +kubebuilder:validation:Optional
	SimpleAuthentication *bool `json:"simpleAuthentication,omitempty" tf:"simple_authentication,omitempty"`

	// Specifies the example response for a successful request. Ensure that the
	// response does not exceed 20,480 characters. Chinese characters must be in UTF-8 or Unicode format.
	// +kubebuilder:validation:Optional
	SuccessResponse *string `json:"successResponse,omitempty" tf:"success_response,omitempty"`

	// Specifies the API type. The valid values are Public and Private.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Specifies the web backend details. The object structure is documented
	// below. Changing this will create a new API resource.
	// +kubebuilder:validation:Optional
	Web []WebParameters `json:"web,omitempty" tf:"web,omitempty"`

	// Specifies the example response for a failed request. The maximum of the policy is 5.
	// The object structure is documented below.
	// +kubebuilder:validation:Optional
	WebPolicy []WebPolicyParameters `json:"webPolicy,omitempty" tf:"web_policy,omitempty"`
}

type BackendParamsObservation struct {
}

type BackendParamsParameters struct {

	// Specifies the description of the constant or system parameter, which contain a
	// maximum of 255 characters, and the angle brackets (< and >) are not allowed.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the location of the backend parameter. The valid values are PATH,
	// QUERY and HEADER.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Specifies the backend parameter name, which contain of 1 to 32 characters and start with a
	// letter. Only letters, digits, hyphens (-), underscores (_) and periods (.) are allowed. The parameter name is not
	// case-sensitive. It cannot start with 'x-apig-' or 'x-sdk-' and cannot be 'x-stage'. If the location is specified as
	// HEADER, the name cannot contain underscores.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the backend parameter type. The valid values are REQUEST, CONSTANT
	// and SYSTEM.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Specifies the request parameter name corresponding to the request parameter name of the
	// back-end parameter.
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type ConditionsObservation struct {
}

type ConditionsParameters struct {

	// Specifies the request parameter name. This parameter is required if the policy type
	// is param.
	// +kubebuilder:validation:Optional
	ParamName *string `json:"paramName,omitempty" tf:"param_name,omitempty"`

	// Specifies the policy type. The valid values are param and source, default to
	// source.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Specifies the API type. The valid values are Public and Private.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the condition type. For a condition with the input parameter source:
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type FuncGraphObservation struct {
}

type FuncGraphParameters struct {

	// Specifies the ID of the backend custom authorization.
	// +kubebuilder:validation:Optional
	AuthorizerID *string `json:"authorizerId,omitempty" tf:"authorizer_id,omitempty"`

	// Specifies the function graph URN.
	// +kubebuilder:validation:Required
	FunctionUrn *string `json:"functionUrn" tf:"function_urn,omitempty"`

	// Specifies the invocation mode. The valid values are async and sync,
	// default to sync.
	// +kubebuilder:validation:Optional
	InvocationType *string `json:"invocationType,omitempty" tf:"invocation_type,omitempty"`

	// Specifies the location of the backend parameter. The valid value is range form 1 to
	// 600,000, default to 5,000.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Specifies the version of the function graph.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type FuncGraphPolicyBackendParamsObservation struct {
}

type FuncGraphPolicyBackendParamsParameters struct {

	// Specifies the description of the request parameter, which contain a maximum of 255
	// characters, and the angle brackets (< and >) are not allowed.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the location of the request parameter. The valid values are PATH,
	// QUERY and HEADER, default to PATH.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Specifies the backend policy name, which can contains of 3 to 64 characters and start with
	// a letter. Only letters, digits, and underscores (_) are allowed.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the API type. The valid values are Public and Private.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Specifies the condition type. For a condition with the input parameter source:
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type FuncGraphPolicyObservation struct {
}

type FuncGraphPolicyParameters struct {

	// Specifies the ID of the backend custom authorization.
	// +kubebuilder:validation:Optional
	AuthorizerID *string `json:"authorizerId,omitempty" tf:"authorizer_id,omitempty"`

	// Specifies an array of one or more backend parameters. The maximum of request
	// parameters is 50. The object structure is documented above.
	// +kubebuilder:validation:Optional
	BackendParams []FuncGraphPolicyBackendParamsParameters `json:"backendParams,omitempty" tf:"backend_params,omitempty"`

	// Specifies an array of one or more policy conditions. Up to five conditions can be set.
	// The object structure is documented below.
	// +kubebuilder:validation:Required
	Conditions []ConditionsParameters `json:"conditions" tf:"conditions,omitempty"`

	// Specifies the effective mode of the backend policy. The valid values are ALL
	// and ANY, default to ANY.
	// +kubebuilder:validation:Optional
	EffectiveMode *string `json:"effectiveMode,omitempty" tf:"effective_mode,omitempty"`

	// Specifies the URN of the function graph.
	// +kubebuilder:validation:Required
	FunctionUrn *string `json:"functionUrn" tf:"function_urn,omitempty"`

	// Specifies the invocation mode of the function graph. The valid values are
	// async and sync, default to sync.
	// +kubebuilder:validation:Optional
	InvocationMode *string `json:"invocationMode,omitempty" tf:"invocation_mode,omitempty"`

	// Specifies the backend policy name, which can contains of 3 to 64 characters and start with
	// a letter. Only letters, digits, and underscores (_) are allowed.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the timeout, in ms, which allowed for APIG to request the backend service. The
	// valid value is range from 1 to 600,000, default to 5,000.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Specifies the version of the function graph.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type MockObservation struct {
}

type MockParameters struct {

	// Specifies the ID of the backend custom authorization.
	// +kubebuilder:validation:Optional
	AuthorizerID *string `json:"authorizerId,omitempty" tf:"authorizer_id,omitempty"`

	// Specifies the response of the backend policy, which contain a maximum of 2,048
	// characters, and the angle brackets (< and >) are not allowed.
	// +kubebuilder:validation:Optional
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type MockPolicyBackendParamsObservation struct {
}

type MockPolicyBackendParamsParameters struct {

	// Specifies the description of the request parameter, which contain a maximum of 255
	// characters, and the angle brackets (< and >) are not allowed.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the location of the request parameter. The valid values are PATH,
	// QUERY and HEADER, default to PATH.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Specifies the backend policy name, which can contains of 3 to 64 characters and start with
	// a letter. Only letters, digits, and underscores (_) are allowed.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the API type. The valid values are Public and Private.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Specifies the condition type. For a condition with the input parameter source:
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type MockPolicyConditionsObservation struct {
}

type MockPolicyConditionsParameters struct {

	// Specifies the request parameter name. This parameter is required if the policy type
	// is param.
	// +kubebuilder:validation:Optional
	ParamName *string `json:"paramName,omitempty" tf:"param_name,omitempty"`

	// Specifies the policy type. The valid values are param and source, default to
	// source.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Specifies the API type. The valid values are Public and Private.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the condition type. For a condition with the input parameter source:
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type MockPolicyObservation struct {
}

type MockPolicyParameters struct {

	// Specifies the ID of the backend custom authorization.
	// +kubebuilder:validation:Optional
	AuthorizerID *string `json:"authorizerId,omitempty" tf:"authorizer_id,omitempty"`

	// Specifies an array of one or more backend parameters. The maximum of request
	// parameters is 50. The object structure is documented above.
	// +kubebuilder:validation:Optional
	BackendParams []MockPolicyBackendParamsParameters `json:"backendParams,omitempty" tf:"backend_params,omitempty"`

	// Specifies an array of one or more policy conditions. Up to five conditions can be set.
	// The object structure is documented below.
	// +kubebuilder:validation:Required
	Conditions []MockPolicyConditionsParameters `json:"conditions" tf:"conditions,omitempty"`

	// Specifies the effective mode of the backend policy. The valid values are ALL
	// and ANY, default to ANY.
	// +kubebuilder:validation:Optional
	EffectiveMode *string `json:"effectiveMode,omitempty" tf:"effective_mode,omitempty"`

	// Specifies the backend policy name, which can contains of 3 to 64 characters and start with
	// a letter. Only letters, digits, and underscores (_) are allowed.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the response of the backend policy, which contain a maximum of 2,048
	// characters, and the angle brackets (< and >) are not allowed.
	// +kubebuilder:validation:Optional
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type RequestParamsObservation struct {
}

type RequestParamsParameters struct {

	// Specifies the default value of the request parameter, which contain a maximum of 255
	// characters, and the angle brackets (< and >) are not allowed.
	// +kubebuilder:validation:Optional
	Default *string `json:"default,omitempty" tf:"default,omitempty"`

	// Specifies the description of the request parameter, which contain a maximum of 255
	// characters, and the angle brackets (< and >) are not allowed.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the example value of the request parameter, which contain a maximum of 255
	// characters, and the angle brackets (< and >) are not allowed.
	// +kubebuilder:validation:Optional
	Example *string `json:"example,omitempty" tf:"example,omitempty"`

	// Specifies the location of the request parameter. The valid values are PATH,
	// QUERY and HEADER, default to PATH.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the maximum value or size of the request parameter.
	// +kubebuilder:validation:Optional
	Maximum *float64 `json:"maximum,omitempty" tf:"maximum,omitempty"`

	// Specifies the minimum value or size of the request parameter. For string type,
	// The maximum and minimum means size. For number type, they means value.
	// +kubebuilder:validation:Optional
	Minimum *float64 `json:"minimum,omitempty" tf:"minimum,omitempty"`

	// Specifies the request parameter name, which contain of 1 to 32 characters and start with a
	// letter. Only letters, digits, hyphens (-), underscores (_) and periods (.) are allowed. If Location is specified as
	// HEADER and security_authentication is specified as APP, the parameter name is not 'Authorization' (
	// case-insensitive) and cannot contain underscores.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies whether the request parameter is required.
	// +kubebuilder:validation:Required
	Required *bool `json:"required" tf:"required,omitempty"`

	// Specifies the request parameter type. The valid values are STRING and NUMBER,
	// default to STRING.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type WebObservation struct {
}

type WebParameters struct {

	// Specifies the ID of the backend custom authorization.
	// +kubebuilder:validation:Optional
	AuthorizerID *string `json:"authorizerId,omitempty" tf:"authorizer_id,omitempty"`

	// Specifies the backend service address, which consists of a domain name or IP
	// address, and a port number, with not more than 255 characters. The backend service address must be in the format "Host
	// name:Port number", for example, apig.example.com:7443. If the port number is not specified, the default HTTPS port
	// 443, or the default HTTP port 80 is used. The backend service address can contain environment variables, each starting
	// with a letter and consisting of 3 to 32 characters. Only letters, digits, hyphens (-), and underscores (_) are
	// allowed.
	// +kubebuilder:validation:Optional
	BackendAddress *string `json:"backendAddress,omitempty" tf:"backend_address,omitempty"`

	// Specifies the proxy host header. The host header can be customized for requests to
	// be forwarded to cloud servers through the VPC channel. By default, the original host header of the request is used.
	// +kubebuilder:validation:Optional
	HostHeader *string `json:"hostHeader,omitempty" tf:"host_header,omitempty"`

	// Specifies the backend request address, which can contain a maximum of 512 characters and
	// must comply with URI specifications.
	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`

	// Specifies the backend request method of the API. The valid types are GET,
	// POST, PUT, DELETE, HEAD, PATCH, OPTIONS and ANY.
	// +kubebuilder:validation:Optional
	RequestMethod *string `json:"requestMethod,omitempty" tf:"request_method,omitempty"`

	// Specifies the backend request protocol. The valid values are HTTP and
	// HTTPS, default to HTTPS.
	// +kubebuilder:validation:Optional
	RequestProtocol *string `json:"requestProtocol,omitempty" tf:"request_protocol,omitempty"`

	// Specifies the indicates whether to enable two-way authentication, default to false.
	// +kubebuilder:validation:Optional
	SSLEnable *bool `json:"sslEnable,omitempty" tf:"ssl_enable,omitempty"`

	// Specifies the timeout, in ms, which allowed for APIG to request the backend service. The
	// valid value is range from 1 to 600,000, default to 5,000.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Specifies the VPC channel ID. This parameter and backend_address are
	// alternative.
	// +kubebuilder:validation:Optional
	VPCChannelID *string `json:"vpcChannelId,omitempty" tf:"vpc_channel_id,omitempty"`
}

type WebPolicyBackendParamsObservation struct {
}

type WebPolicyBackendParamsParameters struct {

	// Specifies the description of the request parameter, which contain a maximum of 255
	// characters, and the angle brackets (< and >) are not allowed.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Specifies the location of the request parameter. The valid values are PATH,
	// QUERY and HEADER, default to PATH.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Specifies the backend policy name, which can contains of 3 to 64 characters and start with
	// a letter. Only letters, digits, and underscores (_) are allowed.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the API type. The valid values are Public and Private.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`

	// Specifies the condition type. For a condition with the input parameter source:
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type WebPolicyConditionsObservation struct {
}

type WebPolicyConditionsParameters struct {

	// Specifies the request parameter name. This parameter is required if the policy type
	// is param.
	// +kubebuilder:validation:Optional
	ParamName *string `json:"paramName,omitempty" tf:"param_name,omitempty"`

	// Specifies the policy type. The valid values are param and source, default to
	// source.
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`

	// Specifies the API type. The valid values are Public and Private.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies the condition type. For a condition with the input parameter source:
	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type WebPolicyObservation struct {
}

type WebPolicyParameters struct {

	// Specifies the ID of the backend custom authorization.
	// +kubebuilder:validation:Optional
	AuthorizerID *string `json:"authorizerId,omitempty" tf:"authorizer_id,omitempty"`

	// Specifies the backend service address, which consists of a domain name or IP
	// address, and a port number, with not more than 255 characters. The backend service address must be in the format "Host
	// name:Port number", for example, apig.example.com:7443. If the port number is not specified, the default HTTPS port 443
	// or the default HTTP port 80 is used. The backend service address can contain environment variables, each starting with
	// a letter and consisting of 3 to 32 characters. Only letters, digits, hyphens (-), and underscores (_) are allowed.
	// +kubebuilder:validation:Optional
	BackendAddress *string `json:"backendAddress,omitempty" tf:"backend_address,omitempty"`

	// Specifies an array of one or more backend parameters. The maximum of request
	// parameters is 50. The object structure is documented above.
	// +kubebuilder:validation:Optional
	BackendParams []WebPolicyBackendParamsParameters `json:"backendParams,omitempty" tf:"backend_params,omitempty"`

	// Specifies an array of one or more policy conditions. Up to five conditions can be set.
	// The object structure is documented below.
	// +kubebuilder:validation:Required
	Conditions []WebPolicyConditionsParameters `json:"conditions" tf:"conditions,omitempty"`

	// Specifies the effective mode of the backend policy. The valid values are ALL
	// and ANY, default to ANY.
	// +kubebuilder:validation:Optional
	EffectiveMode *string `json:"effectiveMode,omitempty" tf:"effective_mode,omitempty"`

	// Specifies the proxy host header. The host header can be customized for requests to
	// be forwarded to cloud servers through the VPC channel. By default, the original host header of the request is used.
	// +kubebuilder:validation:Optional
	HostHeader *string `json:"hostHeader,omitempty" tf:"host_header,omitempty"`

	// Specifies the backend policy name, which can contains of 3 to 64 characters and start with
	// a letter. Only letters, digits, and underscores (_) are allowed.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Specifies the backend request address, which can contain a maximum of 512 characters and
	// must comply with URI specifications.
	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`

	// Specifies the backend request method of the API. The valid types are GET,
	// POST, PUT, DELETE, HEAD, PATCH, OPTIONS and ANY.
	// +kubebuilder:validation:Required
	RequestMethod *string `json:"requestMethod" tf:"request_method,omitempty"`

	// Specifies the backend request protocol. The valid values are HTTP and
	// HTTPS, default to HTTPS.
	// +kubebuilder:validation:Optional
	RequestProtocol *string `json:"requestProtocol,omitempty" tf:"request_protocol,omitempty"`

	// Specifies the timeout, in ms, which allowed for APIG to request the backend service. The
	// valid value is range from 1 to 600,000, default to 5,000.
	// +kubebuilder:validation:Optional
	Timeout *float64 `json:"timeout,omitempty" tf:"timeout,omitempty"`

	// Specifies the VPC channel ID. This parameter and backend_address are
	// alternative.
	// +kubebuilder:validation:Optional
	VPCChannelID *string `json:"vpcChannelId,omitempty" tf:"vpc_channel_id,omitempty"`
}

// APISpec defines the desired state of API
type APISpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     APIParameters `json:"forProvider"`
}

// APIStatus defines the observed state of API.
type APIStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        APIObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// API is the Schema for the APIs API. ""page_title: "flexibleengine_apig_api"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,flexibleengine}
type API struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              APISpec   `json:"spec"`
	Status            APIStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// APIList contains a list of APIs
type APIList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []API `json:"items"`
}

// Repository type metadata.
var (
	API_Kind             = "API"
	API_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: API_Kind}.String()
	API_KindAPIVersion   = API_Kind + "." + CRDGroupVersion.String()
	API_GroupVersionKind = CRDGroupVersion.WithKind(API_Kind)
)

func init() {
	SchemeBuilder.Register(&API{}, &APIList{})
}
