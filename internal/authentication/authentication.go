package authentication

import (
	"github.com/launchboxio/launchbox/internal/config"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/github"
	"github.com/markbates/goth/providers/google"
	"github.com/markbates/goth/providers/openidConnect"
	"os"
)

func New(conf *config.Config) {
	providers := []goth.Provider{}
	for auth, config := range conf.Auth {
		redirectUri := config.RedirectUri
		switch auth {
		case "google":
			providers = append(providers, google.New(
				os.Getenv("GOOGLE_KEY"),
				os.Getenv("GOOGLE_SECRET"),
				redirectUri,
			))
			break
		case "github":
			providers = append(providers, github.New(
				os.Getenv("GITHUB_KEY"),
				os.Getenv("GITHUB_SECRET"),
				redirectUri,
			))
			break
		case "bitbucket":
			providers = append(providers, google.New(
				os.Getenv("BITBUCKET_KET"),
				os.Getenv("BITBUCKET_SECRET"),
				redirectUri,
			))
			break
		case "gitlab":
			providers = append(providers, google.New(
				os.Getenv("GITLAB_KEY"),
				os.Getenv("GITLAB_SECRET"),
				redirectUri,
			))
			break
		case "auth0":
			providers = append(providers, google.New(
				os.Getenv("AUTH0_KEY"),
				os.Getenv("AUTH0_SECRET"),
				redirectUri,
				os.Getenv("AUTH0_DOMAIN"),
			))
			break
		case "gitea":
			providers = append(providers, google.New(
				os.Getenv("GITEA_KEY"),
				os.Getenv("GITEA_SECRET"),
				redirectUri,
			))
			break
		case "okta":
			providers = append(providers, google.New(
				os.Getenv("OKTA_KEY"),
				os.Getenv("OKTA_SECRET"),
				redirectUri,
				"openid", "profile", "email",
			))
		case "oidc":
			openidConnect, _ := openidConnect.New(os.Getenv("OPENID_CONNECT_KEY"), os.Getenv("OPENID_CONNECT_SECRET"), config.RedirectUri, os.Getenv("OPENID_CONNECT_DISCOVERY_URL"))
			if openidConnect != nil {
				providers = append(providers, openidConnect)
			}
			break
		}
	}
	goth.UseProviders(providers...)
}
