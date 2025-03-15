// patchr:replace module github.com/wreulicke/{{ input "repo" }}
module github.com/wreulicke/go-cli-template

go 1.22

require github.com/spf13/cobra v1.8.1

require github.com/pmezard/go-difflib v1.0.0 // indirect

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/wreulicke/snap v0.0.0-20240620141153-a95fa41ce87c
)
