# Google Fonts

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Used By][used-by-svg]][used-by-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

This package provides a very limited number of [Google Fonts](https://fonts.google.com/), licensed under the Apache 2.0 license, that can be embedded in Go programs.

Currenty, the following font families are supported:

1. Roboto ([on `fonts.google.com`](https://fonts.google.com/specimen/Roboto))

The `main.go` file in the root folder auto-generates Go files.

## Adding fonts

1. Download the zip file from fonts.google.com
2. Unzip into new subfolder
3. Lower-case the subfolder name
4. Run `go run main.go <new_folder_name>`

 [build-status-svg]: https://github.com/grokify/google-fonts/workflows/test/badge.svg
 [build-status-url]: https://github.com/grokify/google-fonts/actions/workflows/test.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/google-fonts
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/google-fonts
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/google-fonts
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/google-fonts/v2
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/google-fonts/blob/main/LICENSE.txt
 [used-by-svg]: https://sourcegraph.com/github.com/grokify/google-fonts/-/badge.svg
 [used-by-url]: https://sourcegraph.com/github.com/grokify/google-fonts?badge
 [loc-svg]: https://tokei.rs/b1/github/grokify/google-fonts
 [repo-url]: https://github.com/grokify/google-fonts