# Golang Template

![Go Version](https://img.shields.io/github/go-mod/go-version/conijnio/golang-template)
[![License](https://img.shields.io/badge/License-Apache2-green.svg)](./LICENSE)
[![Maintenance](https://img.shields.io/badge/Maintained-yes-green.svg)](https://github.com/conijnio/golang-template/graphs/commit-activity)
[![Workflow: ci](https://github.com/conijnio/golang-template/actions/workflows/ci.yml/badge.svg)](https://github.com/conijnio/golang-template/actions/workflows/go.yml)
[![Workflow: release](https://github.com/conijnio/golang-template/actions/workflows/release.yml/badge.svg)](https://github.com/conijnio/golang-template/actions/workflows/goreleaser.yml)
![Release](https://img.shields.io/github/v/release/conijnio/golang-template)
[![Go Report Card](https://goreportcard.com/badge/github.com/conijnio/golang-template)](https://goreportcard.com/report/github.com/conijnio/golang-template)
[![Coverage Status](https://coveralls.io/repos/github/conijnio/golang-template/badge.svg?branch=main)](https://coveralls.io/github/conijnio/golang-template?branch=main)

Template repository for Golang projects.

## First steps!

Since you are starting from a template there are some steps that you need to take!

### Replace golang-template with your repository name

We need to replace all the `golang-template` occurrences with the new project name:

```shell
find . \( -iname '*.yaml' -o -iname '*.yml' -o -iname '*.md' -o -iname '*.go' -o -iname '*.mod' -o -iname '*.toml' \) -exec sed -i '' -e "s/golang-template/golang-project/g" {} \;
```

### Setup documentation pages

Because we use the hugo theme as a submodule we need to initialize it:

```shell
git submodule init
git submodule update
```

In GitHub go to the project **Settings**, **Pages** and select `GitHub Actions` as the **Source** under **Build and Deploy**.

Now you only need to write the documentation ;-)

### Create first release

Time to create an initial release!

```shell
git tag latest
git tag v0.1.0
git push --tags
```

## Prerequisites

You will need to install the following tools to successfully run the make targets:

```shell
go install github.com/fzipp/gocyclo/cmd/gocyclo@latest
go install github.com/uudashr/gocognit/cmd/gocognit@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
go install github.com/go-critic/go-critic/cmd/gocritic@latest
```

To make use of the pre-commit hooks you need to install [pre-commit](https://pre-commit.com) and execute the following command:

```shell
pre-commit install
```

## Commands

- `make build`, builds the project.
- `make complexity`, perform complexity scans on the codebase.
- `make coverage`, create and displays the code coverage report in HTML.
- `make help`, displays all the available options.
- `make lint`, performs linting actions on the codebase.
- `make test`, runs all the unit tests.

### Run goreleaser locally

Because we enabled signing you need to supply a `GPG_FINGERPRINT` environment variable. You will be prompted for a passphrase.

```shell
GPG_FINGERPRINT=C490C64E6938FD0C goreleaser release --snapshot --clean
```

Afterward, you can validate the signature using the following command:

```shell
gpg --verify [Signature] [File]
```

## License

This project is free and open source software licensed under the [Apache 2.0 License](./LICENSE).
