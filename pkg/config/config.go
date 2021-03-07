package config

type Config struct {
	Docstores []Docstore `yaml:"docstores,omitempty"`
}

type Docstore struct {
	Name             string `yaml:"name"`
	URL              string `yaml:"url"`
	IDField          string `yaml:"idField,omitempty"`
	ConnectionString string `yaml:"connectionString,omitempty"`
}
