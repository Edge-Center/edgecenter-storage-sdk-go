# EdgeCenter Storage API Golang SDK

[![Lint](https://github.com/Edge-Center/edgecenter-storage-sdk-go/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/Edge-Center/edgecenter-storage-sdk-go/actions/workflows/golangci-lint.yml)
[![Build and test](https://github.com/Edge-Center/edgecenter-storage-sdk-go/actions/workflows/go.yml/badge.svg)](https://github.com/Edge-Center/edgecenter-storage-sdk-go/actions/workflows/go.yml)

The purpose of this project is to cover EdgeCenter Storage API methods related with EdgeCenter Terraform plugin (https://github.com/Edge-Center/terraform-provider-edgecenter).

## Internal design of SDK

Since Storage service provides Swagger open docs so we generate an actual version of the client with https://goswagger.io/generate/client.html.
And we extend our wrapper of this client after. 
Read Makefile to get more technical details.

## License

This project is licensed under the Mozilla Public License 2.0, see [LICENSE](LICENSE).
Copyright (c) 2022 LLC "EDGECENTER"

This project includes code originally developed by G-Core Labs S.A. (Copyright G-Core Labs S.A.), licensed under the same license as this repository. Subsequent development by LLC "EDGECENTER".
