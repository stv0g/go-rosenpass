module github.com/stv0g/go-rosenpass

go 1.20

require (
	github.com/cloudflare/circl v0.0.0-00010101000000-000000000000
	github.com/open-quantum-safe/liboqs-go v0.0.0-20230607110523-d1fd56d355f7
	github.com/pelletier/go-toml/v2 v2.0.8
	github.com/spf13/cobra v1.7.0
	golang.org/x/crypto v0.9.0
	golang.org/x/exp v0.0.0-20230522175609-2e198f4a06a1
)

require github.com/stretchr/testify v1.8.4 // testing

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/sys v0.8.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// For Classic McEliece support
// https://github.com/cloudflare/circl/pull/378
replace github.com/cloudflare/circl => github.com/pufferffish/circl v0.0.0-20221209095330-9dbf4684a331
