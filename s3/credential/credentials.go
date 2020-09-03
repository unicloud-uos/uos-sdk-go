package credential

const (
	StaticCredentialsProviderName = "StaticCredentialsProvider"
)

var DefaultCredentials = &Credentials{
	AccessKeyID:     "DefaultAccessKeyID",
	SecretAccessKey: "DefaultSecretAccessKey",
	ProviderName:    "DefaultCredentialsProvider",
}

// Credentials is the credentials value for individual credential fields.
type Credentials struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
	ProviderName    string // credential provider
}

// StaticCredentialsEmptyError is emitted when static credentials are empty.
type StaticCredentialsEmptyError struct{}

func (*StaticCredentialsEmptyError) Error() string {
	return "default credentials are empty"
}

// NewStaticCredentialsProvider return a StaticCredentialsProvider initialized with the credentials passed in.
func NewStaticCredentials(key, secret, session string) *Credentials {
	return &Credentials{
		AccessKeyID:     key,
		SecretAccessKey: secret,
		SessionToken:    session,
	}
}

// Retrieve returns the credentials or error if the credentials are invalid.
func (c Credentials) Retrieve() (Credentials, error) {
	if c.AccessKeyID == "" || c.SecretAccessKey == "" {
		return Credentials{ProviderName: StaticCredentialsProviderName}, &StaticCredentialsEmptyError{}
	}

	if len(c.ProviderName) == 0 {
		c.ProviderName = StaticCredentialsProviderName
	}

	return c, nil
}
