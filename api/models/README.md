# Product Models

## Description

Data models to be used by the Products API

## List of Models

### Product

Describes a Product within the system that has the following properties

- Name: Name of the Product
- Description: Short Description
- Colour: Colout of the Prouct if applicable
- SKU: Stock Keeping Unit. Used for lookups
- 

### Location

Describes a geographical location

## Assumptions

- One product can be linked with many locations.
- A location describes a geographical area with seperate tax laws.
- Location names must be unique.
- Taxes are tightly coupled to a location.
- a location may not have a tax associated with it.
- Each products price has to have a base location associated with it. If a product is being purchased from another location then the currency conversion is done at the time of checkout.
