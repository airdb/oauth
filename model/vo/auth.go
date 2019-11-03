package vo

import (
	"github.com/airdb/passport/model/po"
)

type ProviderSecret struct {
	Provider string `json:"provider"`
	ClientID string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURI string `json:"redirect_uri"`
}

func FromPoProviderSecret(poSecret *po.Secret) *ProviderSecret {
	return &ProviderSecret{
		Provider:poSecret.Provider,
		ClientID:poSecret.ClientID,
		ClientSecret:poSecret.ClientSecret,
		RedirectURI:poSecret.RedirectURI,
	}
}

func FromPoProviderSecrets(poSecrets []*po.Secret) (secrets []*ProviderSecret) {
	for _, secret := range poSecrets {
		secrets = append(secrets, FromPoProviderSecret(secret))
	}
	return
}

func ListProvider() []*ProviderSecret{
	ProviderSecrets := FromPoProviderSecrets(po.ListProvider())

	return ProviderSecrets
}

func QueryProvider() *ProviderSecret {
	return FromPoProviderSecret(po.QueryProvider())
}