module github.com/shipmight/helm-playground

// Based on:
//   https://github.com/helm/helm/blob/v4.1.2/go.mod
// Commands:
//   go get github.com/BurntSushi/toml@v1.6.0
//   go get github.com/Masterminds/sprig/v3@v3.3.0
//   go get sigs.k8s.io/yaml@v1.6.0
// Update these commands with the versions and then run them!
// Also keep go version in sync!

go 1.25.0

require (
	github.com/BurntSushi/toml v1.6.0
	github.com/Masterminds/sprig/v3 v3.3.0
	sigs.k8s.io/yaml v1.6.0
)

require (
	dario.cat/mergo v1.0.1 // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver/v3 v3.3.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/huandu/xstrings v1.5.0 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	github.com/spf13/cast v1.7.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	go.yaml.in/yaml/v2 v2.4.2 // indirect
	golang.org/x/crypto v0.26.0 // indirect
)
