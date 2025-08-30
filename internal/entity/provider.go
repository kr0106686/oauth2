package entity

type Provider struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Scopes       []string
	Endpoint     Endpoint
}

type Endpoint struct {
	AuthURL  string
	TokenURL string
	InfoURL  string
}
