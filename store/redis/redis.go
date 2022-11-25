package redis

import (
	"context"
	"time"

	"github.com/0x00b/wx-openapi/store"
	rediscli "github.com/go-redis/redis/v8"
)

// Storage implement store
type Storage struct {
	cli *rediscli.Client
}

var _ store.WXTPCacheStorage = &Storage{}

// New init
func New(host, password string, db int) *Storage {
	return &Storage{
		cli: rediscli.NewClient(&rediscli.Options{
			Addr:     host,
			Password: password,
			DB:       db,
		}),
	}
}

// NewStore
func NewStore(client *rediscli.Client) *Storage {
	return &Storage{
		cli: client,
	}
}

func (s *Storage) getVerifyTicketKey(c context.Context, thirdPartyAppId string) string {
	return "wx_verifyTicket_" + thirdPartyAppId
}

// GetVerifyTicket implement Storager
func (s *Storage) GetVerifyTicket(c context.Context, thirdPartyAppId string) (string, error) {
	cmd := s.cli.Get(c, s.getVerifyTicketKey(c, thirdPartyAppId))
	if isRedisNil(cmd.Err()) {
		return "", nil
	}
	if cmd.Err() != nil {
		return "", cmd.Err()
	}

	return cmd.Result()
}

// SetVerifyTicket implement Storager
func (s *Storage) SetVerifyTicket(c context.Context, thirdPartyAppId, ticket string) error {
	cmd := s.cli.Set(c, s.getVerifyTicketKey(c, thirdPartyAppId), ticket, time.Duration(0))
	return cmd.Err()
}

func (s *Storage) getComponentAccessTokenKey(c context.Context, thirdPartyAppId string) string {
	return "wx_componentAccessToken_" + thirdPartyAppId
}

// GetComponentAccessToken implement Storager
func (s *Storage) GetComponentAccessToken(c context.Context, thirdPartyAppId string) (string, error) {
	cmd := s.cli.Get(c, s.getComponentAccessTokenKey(c, thirdPartyAppId))
	if isRedisNil(cmd.Err()) {
		return "", nil
	}
	if cmd.Err() != nil {
		return "", cmd.Err()
	}

	return cmd.Result()
}

// SetComponentAccessToken implement Storager
func (s *Storage) SetComponentAccessToken(c context.Context, thirdPartyAppId, token string, expire int32) error {
	cmd := s.cli.Set(c, s.getComponentAccessTokenKey(c, thirdPartyAppId), token, time.Second*time.Duration(expire))
	return cmd.Err()
}

func (s *Storage) getPreAuthCodeKey(c context.Context, thirdPartyAppId string) string {
	return "wx_preAuthCode_" + thirdPartyAppId
}

// SetPreAuthCode set code
func (s *Storage) SetPreAuthCode(c context.Context, thirdPartyAppId, code string, expire int32) error {
	cmd := s.cli.Set(c, s.getPreAuthCodeKey(c, thirdPartyAppId), code, time.Second*time.Duration(expire))
	return cmd.Err()
}

// GetPreAuthCode get code
func (s *Storage) GetPreAuthCode(c context.Context, thirdPartyAppId string) (string, error) {
	cmd := s.cli.Get(c, s.getPreAuthCodeKey(c, thirdPartyAppId))
	if isRedisNil(cmd.Err()) {
		return "", nil
	}
	if cmd.Err() != nil {
		return "", cmd.Err()
	}

	return cmd.Result()
}

func (s *Storage) getAuthorizerAccessTokenKey(c context.Context, authorizeAppId string) string {
	return "wx_authorizerAccessToken_" + authorizeAppId
}

// SetAuthorizerAccessToken set Authorizer access token
func (s *Storage) SetAuthorizerAccessToken(c context.Context, authorizeAppId, token string, expire int32) error {
	cmd := s.cli.Set(c, s.getAuthorizerAccessTokenKey(c, authorizeAppId), token, time.Second*time.Duration(expire))
	return cmd.Err()
}

// GetAuthorizerAccessToken Get Authorizer access token
func (s *Storage) GetAuthorizerAccessToken(c context.Context, authorizeAppId string) (string, error) {
	cmd := s.cli.Get(c, s.getAuthorizerAccessTokenKey(c, authorizeAppId))
	if isRedisNil(cmd.Err()) {
		return "", nil
	}
	if cmd.Err() != nil {
		return "", cmd.Err()
	}

	return cmd.Result()
}

// DelAuthorizerAccessToken delete Authorizer access token
func (s *Storage) DelAuthorizerAccessToken(c context.Context, authorizeAppId string) error {
	cmd := s.cli.Del(c, s.getAuthorizerAccessTokenKey(c, authorizeAppId))
	return cmd.Err()
}

func isRedisNil(err error) bool {
	if rediscli.Nil == err {
		return true
	}
	return false
}
