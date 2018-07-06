# Product Models

## Description

Data models to be used by the Products API

- [Product Models](#product-models)
    - [Description](#description)
    - [List of Models](#list-of-models)
        - [Product](#product)
        - [Location](#location)
        - [Transactions](#transactions)
    - [Assumptions](#assumptions)

## List of Models

### Product

Describes a Product within the system that has the following attributes:

- **Name**: Name of the Product
- **Description**: Short Description
- **Colour**: Colour of the Product if applicable
- **SKU**: Stock Keeping Unit. Used for lookup
- **BasePrice**: Price of the product in the currency that it was created with
- **BaseCurrency**: The currency that the price is in. Used for conversion

### Location

Describes a geographical location. Includes currency and taxes for a location. Used to link a currency and where a transaction has occurred for currency conversion.

A location can be a country, county, city. The model is intentionally kept abstract to help better definition in the future if required.

Locations have the following attributes:

- **Name**: The Name of the Location
- **Currency**: The currency for the given location
    - **Name**: The name of the currency. Should adhere to [ISO4217](https://en.wikipedia.org/wiki/ISO_4217)
    - **Country**: The Name of the Country
    - **Symbol**: The Currency symbol.
- **Taxes**: Array of the taxes associated with the location.
    - **Name**: name of the tax (e.g. VAT).
    - **Amount**: decimal percentage of the tax.

### Transactions

Describes a transaction. A transaction is defined as the exchange of monies for a set of products.

A transaction's total is calculated as:

- **Subtotal**: sum of all prices in the transaction's location currency
- **Tax total**: sum of (subtotal * each tax for the location)
- **Total**: sum of Subtotal and Tax total


## Assumptions

- One product can be linked with many locations.
- A location describes a geographical area with separate tax laws.
- Location names must be unique.
- Taxes are tightly coupled to a location.
- a location may not have a tax associated with it.
- Each products price has to have a base location associated with it. If a product is being purchased from another location then the currency conversion is done at the time of checkout.
