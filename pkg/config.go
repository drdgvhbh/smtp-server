package pkg

type Config struct {
	Port uint `toml:"PORT"`
	DomainName string `toml:"DOMAIN_NAME"`
}
