# Product

## Description

Package to interact with data about products within an e-commerce system

```
go get -u davyj0nes/products-api/product
```

## Assumptions

- One product can be linked with many regions.
- A region describes a geographical area with seperate tax laws.
- Region names must be unique.
- Taxes are tightly coupled to a region.
- a Region may not have a tax associated with it.
