module github.com/shipmight/helm-playground

// Based on:
//   https://github.com/helm/helm/blob/v3.14.0/go.mod
// Commands:
//   go get github.com/BurntSushi/toml@v1.3.2
//   go get github.com/Masterminds/sprig/v3@v3.2.3
//   go get sigs.k8s.io/yaml@v1.3.0
// Update these commands with the versions and then run them!
// Also keep go version in sync!

go 1.21

require (
	github.com/BurntSushi/toml v1.3.2
	github.com/Masterminds/sprig/v3 v3.2.3
	sigs.k8s.io/yaml v1.3.0
)

require (
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver/v3 v3.2.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/huandu/xstrings v1.3.3 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	golang.org/x/crypto v0.3.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
