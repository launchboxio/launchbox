package api

const (
	HTTPAddrEnvName      = "LAUNCHBOX_HTTP_ADDR"
	HTTPTokenEnvName     = "LAUNCHBOX_HTTP_TOKEN"
	HTTPTokenFileEnvName = "LAUNCHBOX_HTTP_TOKEN_FILE"
	HTTPAuthEnvName      = "LAUNCHBOX_HTTP_AUTH"
	HTTPSSLEnvName       = "LAUNCHBOX_HTTP_SSL"
	HTTPCAFile           = "LAUNCHBOX_CACERT"
	HTTPCAPath           = "LAUNCHBOX_CAPATH"
	HTTPClientCert       = "LAUNCHBOX_CLIENT_CERT"
	HTTPClientKey        = "LAUNCHBOX_CLIENT_KEY"
	HTTPSSLVerifyEnvName = "LAUNCHBOX_HTTP_SSL_VERIFY"
)

type Client struct{}
