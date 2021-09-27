package wx

import "crypto/tls"

type ClientOption func(client *wxclient)

// WithInsecureSkipVerify specifies the `InsecureSkipVerify` to http client tls config.
func WithInsecureSkipVerify() ClientOption {
	return func(client *wxclient) {
		client.insecure = true
	}
}

// WithTLSCertificates specifies the certificate to http client tls config.
func WithTLSCertificates(certs ...tls.Certificate) ClientOption {
	return func(client *wxclient) {
		client.certs = certs
	}
}

func WithLogger() ClientOption {
	return func(client *wxclient) {

	}
}

// ActionOption configures how we set up the action
type ActionOption func(a *action)

// WithQuery specifies the `query` to Action.
func WithQuery(key, value string) ActionOption {
	return func(a *action) {
		a.query.Set(key, value)
	}
}

// WithBody specifies the `body` to Action.
func WithBody(f func() ([]byte, error)) ActionOption {
	return func(a *action) {
		a.body = f
	}
}

// WithWXML specifies the `wxml` to Action.
func WithWXML(f func(appid, mchid, nonce string) (WXML, error)) ActionOption {
	return func(a *action) {
		a.wxml = f
	}
}

// WithUploadField specifies the `upload field` to Action.
func WithUploadField(field *UploadField) ActionOption {
	return func(a *action) {
		a.uploadfield = field
	}
}

// WithDecode specifies the `decode` to Action.
func WithDecode(f func(resp []byte) error) ActionOption {
	return func(a *action) {
		a.decode = f
	}
}

// WithTLS specifies the `tls` to Action.
func WithTLS() ActionOption {
	return func(a *action) {
		a.tls = true
	}
}