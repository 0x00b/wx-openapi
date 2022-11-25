# Go API client for third_part_auth

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import third_part_auth "github.com/0x00b/wx-openapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), third_part_auth.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), third_part_auth.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), third_part_auth.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), third_part_auth.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.weixin.qq.com/cgi-bin*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthApi* | [**ComponentApiAuthorizerTokenPost**](docs/AuthApi.md#componentapiauthorizertokenpost) | **Post** /component/api_authorizer_token | 获取/刷新接口调用令牌
*AuthApi* | [**ComponentApiComponentTokenPost**](docs/AuthApi.md#componentapicomponenttokenpost) | **Post** /component/api_component_token | 令牌
*AuthApi* | [**ComponentApiCreatePreauthcodePost**](docs/AuthApi.md#componentapicreatepreauthcodepost) | **Post** /component/api_create_preauthcode | 预授权码
*AuthApi* | [**ComponentApiGetAuthorizerInfoPost**](docs/AuthApi.md#componentapigetauthorizerinfopost) | **Post** /component/api_get_authorizer_info | 获取授权方的帐号基本信息
*AuthApi* | [**ComponentApiQueryAuthPost**](docs/AuthApi.md#componentapiqueryauthpost) | **Post** /component/api_query_auth | 使用授权码获取授权信息
*AuthApi* | [**ComponentApiStartPushTicketPost**](docs/AuthApi.md#componentapistartpushticketpost) | **Post** /component/api_start_push_ticket | 启动ticket推送服务


## Documentation For Models

 - [AuthorizerTokenRequest](docs/AuthorizerTokenRequest.md)
 - [AuthorizerTokenResponse](docs/AuthorizerTokenResponse.md)
 - [ComponentTokenRequest](docs/ComponentTokenRequest.md)
 - [ComponentTokenResponse](docs/ComponentTokenResponse.md)
 - [CreatePreauthcodeRequest](docs/CreatePreauthcodeRequest.md)
 - [CreatePreauthcodeResponse](docs/CreatePreauthcodeResponse.md)
 - [GetAuthorizerInfoRequest](docs/GetAuthorizerInfoRequest.md)
 - [GetAuthorizerInfoResponse](docs/GetAuthorizerInfoResponse.md)
 - [GetAuthorizerInfoResponseAuthorizationInfo](docs/GetAuthorizerInfoResponseAuthorizationInfo.md)
 - [GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner](docs/GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInner.md)
 - [GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInnerFuncscopeCategory](docs/GetAuthorizerInfoResponseAuthorizationInfoFuncInfoInnerFuncscopeCategory.md)
 - [GetAuthorizerInfoResponseAuthorizerInfo](docs/GetAuthorizerInfoResponseAuthorizerInfo.md)
 - [GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo](docs/GetAuthorizerInfoResponseAuthorizerInfoBusinessInfo.md)
 - [GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo](docs/GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfo.md)
 - [GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfoCategoriesInner](docs/GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfoCategoriesInner.md)
 - [GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfoNetwork](docs/GetAuthorizerInfoResponseAuthorizerInfoMiniProgramInfoNetwork.md)
 - [GetAuthorizerInfoResponseAuthorizerInfoServiceTypeInfo](docs/GetAuthorizerInfoResponseAuthorizerInfoServiceTypeInfo.md)
 - [GetAuthorizerInfoResponseAuthorizerInfoVerifyTypeInfo](docs/GetAuthorizerInfoResponseAuthorizerInfoVerifyTypeInfo.md)
 - [QueryAuthRequest](docs/QueryAuthRequest.md)
 - [QueryAuthResponse](docs/QueryAuthResponse.md)
 - [QueryAuthResponseAuthorizationInfo](docs/QueryAuthResponseAuthorizationInfo.md)
 - [StartPushTicketRequest](docs/StartPushTicketRequest.md)
 - [StartPushTicketResponse](docs/StartPushTicketResponse.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


