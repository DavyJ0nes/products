# Products API
[![Go Report Card](https://goreportcard.com/badge/github.com/DavyJ0nes/products)](https://goreportcard.com/report/github.com/DavyJ0nes/products)

## Table of Contents

- [Products API](#products-api)
    - [Table of Contents](#table-of-contents)
    - [Description](#description)
    - [Basic Full Transaction Example](#basic-full-transaction-example)
    - [Usage](#usage)
    - [TODO](#todo)
    - [License](#license)

## Description

Example of how to create a simple API that could be used as part of a commerce system or POS.

You can find information about the data models used [here](./api/models)

## Basic Full Transaction Example 

1. Run the docker container with: `make run-docker` 
2. Send a POST request to the container (this is done with curl): `make transaction-test`

You should see output similar to the following:
This output can be used by a receipt printer

```json
{
    "order_id": "b63ad3fc-31a8-48e9-9ecb-53779890e8e3",
    "formatted_products": [
        {
            "product_quantity": 1,
            "product_name": "Coffee Mug",
            "price": 5.99
        },
        {
            "product_quantity": 1,
            "product_name": "Coaster",
            "price": 2.5
        },
        {
            "product_quantity": 1,
            "product_name": "Glass Tumbler",
            "price": 12.99
        }
    ],
    "formatted_date_time": "06-07-2018 20:10:57",
    "subtotal": 19.43,
    "taxtotal": 3.89,
    "currency_conversion": 0.885697,
    "total": 23.32
}
```

## Usage

How to build and run the API as a service. More information on these commands can be found in the [Makefile](./Makefile)

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

# Make POST request to test creating a new transaction
make test-transaction
```

## TODO

- [ ] mock dependecy on currency convertor. Causes flaky tests due to changing conversion rates
- [ ] Abstract transaction totals to seperate package and look at cleaner implementation
- [ ] Implement Data store
- [ ] Add authentication to API
- [ ] Improve transactions to better handle multiple quantities of the same product
- [ ] Tidy up test suites for improved readability
- [ ] Improve the healthz endpoint
- [ ] Add readyz endpoint

## License

[Apache 2.0](./LICENSE)
