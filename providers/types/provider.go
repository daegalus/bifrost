package types

// Provider interface for all providers to follow.
type Provider interface {
	Encrypt(string, []byte) ([]byte, error)
	Decrypt(string, []byte) ([]byte,error)
	Download(string, string) ([]byte, error)
	Upload(string, string, []byte) (bool, error)
	Register(*map[string]Provider)
	ProviderID() string
}

// ProviderConst string type for Provider constants
type ProviderConst string