# Products API

## Table of Contents

- [Products API](#products-api)
    - [Table of Contents](#table-of-contents)
    - [Description](#description)
    - [Usage](#usage)
    - [TODO](#todo)
    - [License](#license)

## Description

Example of how to create a simple API that could be used as part of a commerce system or POS.

You can find information about the data models used [here](./api/models)

## Usage

How to build and run the API

```shell
# Basic run while testing
make run

# Run test suite over all packages
make test

# Build Docker Image
make build

# Run Docker Image
make run-docker

# Deploy to Kubernetes Cluster
make deploy
```

## TODO

- [ ] Implement Data store
- [ ] Add authentication to API
- [ ] Improve the healthz endpoint
- [ ] Add readyz endpoint

## License

[Apache 2.0](./LICENSE)
