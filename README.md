# AWS IAM User

![Go Version](https://img.shields.io/github/go-mod/go-version/conijnio/aws-iam-user)
[![License](https://img.shields.io/badge/License-Apache2-green.svg)](./LICENSE)
[![Maintenance](https://img.shields.io/badge/Maintained-yes-green.svg)](https://github.com/conijnio/aws-iam-user/graphs/commit-activity)
[![Workflow: ci](https://github.com/conijnio/aws-iam-user/actions/workflows/ci.yml/badge.svg)](https://github.com/conijnio/aws-iam-user/actions/workflows/go.yml)
[![Workflow: release](https://github.com/conijnio/aws-iam-user/actions/workflows/release.yml/badge.svg)](https://github.com/conijnio/aws-iam-user/actions/workflows/goreleaser.yml)
![Release](https://img.shields.io/github/v/release/conijnio/aws-iam-user)
[![Go Report Card](https://goreportcard.com/badge/github.com/conijnio/aws-iam-user)](https://goreportcard.com/report/github.com/conijnio/aws-iam-user)
[![Coverage Status](https://coveralls.io/repos/github/conijnio/aws-iam-user/badge.svg?branch=main)](https://coveralls.io/github/conijnio/aws-iam-user?branch=main)

When using IAM users it could be cumbersome to rotate your `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`.
This could lead into resources being in a non-compliant state. The `aws-iam-user` tool will address exactly that!

More information can be found on the [documentation pages](https://conijnio.github.io/aws-iam-user/).

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

## License

This project is free and open source software licensed under the [Apache 2.0 License](./LICENSE).
