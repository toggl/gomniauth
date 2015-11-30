package oauth2

import (
	"github.com/toggl/gomniauth/common"
)

// AuthorizationHeader returns the key, value pair to insert into an authorized request.
func AuthorizationHeader(creds *common.Credentials) (key, value string) {
	return "Authorization", "Bearer " + creds.Get(OAuth2KeyAccessToken).Str("Invalid")
}
