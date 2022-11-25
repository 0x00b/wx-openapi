package store

import "context"

// ThirdPartyStorage to store third platform ticket, token etc
type WXTPCacheStorage interface {
	GetVerifyTicket(c context.Context, componentAppId string) (string, error)
	SetVerifyTicket(c context.Context, componentAppId, ticket string) error
	GetComponentAccessToken(c context.Context, componentAppId string) (string, error)
	SetComponentAccessToken(c context.Context, componentAppId, token string, expire int32) error
	GetPreAuthCode(c context.Context, componentAppId string) (string, error)
	SetPreAuthCode(c context.Context, componentAppId, code string, expire int32) error
	SetAuthorizerAccessToken(c context.Context, authorizeAppId, token string, expire int32) error
	GetAuthorizerAccessToken(c context.Context, authorizeAppId string) (string, error)
	DelAuthorizerAccessToken(c context.Context, authorizeAppId string) error
}
