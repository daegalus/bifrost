package registry

import (
	"bifrost/providers/plugins"
	"bifrost/providers/types"
)
// ProviderRegistry the provider string -> provider following the Provider interface
var ProviderRegistry = map[string]types.Provider{}

// LoadProviders loads the provider plugins into the registry for use.
func LoadProviders() {
	plugins.GCP{}.Register(&ProviderRegistry)
}