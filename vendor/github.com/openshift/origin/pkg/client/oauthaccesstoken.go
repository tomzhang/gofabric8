package client

import (
	oauthapi "github.com/openshift/origin/pkg/oauth/api"
)

// OAuthAccessTokensInterface has methods to work with OAuthAccessTokens resources in a namespace
type OAuthAccessTokensInterface interface {
	OAuthAccessTokens() OAuthAccessTokenInterface
}

// OAuthAccessTokenInterface exposes methods on OAuthAccessTokens resources.
type OAuthAccessTokenInterface interface {
	Create(token *oauthapi.OAuthAccessToken) (*oauthapi.OAuthAccessToken, error)
	Delete(name string) error
}

type oauthAccessTokenInterface struct {
	r *Client
}

func newOAuthAccessTokens(c *Client) *oauthAccessTokenInterface {
	return &oauthAccessTokenInterface{
		r: c,
	}
}

// Delete removes the OAuthAccessToken on server
func (c *oauthAccessTokenInterface) Delete(name string) (err error) {
	err = c.r.Delete().Resource("oAuthAccessTokens").Name(name).Do().Error()
	return
}

func (c *oauthAccessTokenInterface) Create(token *oauthapi.OAuthAccessToken) (result *oauthapi.OAuthAccessToken, err error) {
	result = &oauthapi.OAuthAccessToken{}
	err = c.r.Post().Resource("oAuthAccessTokens").Body(token).Do().Into(result)
	return
}
