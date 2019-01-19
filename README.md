# go-ghost

Golang API client for Ghost publishing platform

This package uses both v0.1 and v2 of Ghost API because v2 doesn't support resource creation or update.

## Useful links

[Ghost API v0.1](https://api.ghost.org/v0.1/docs)

[Managing your posts, tags, and users with the Ghost Blog API](https://grantwinney.com/what-is-the-ghost-api/)

[Ghost API v0.1](https://docs.ghost.org/api/content/)

## Obtaining API key

- create account on your Ghost installation
- select Integrations in the left menu
- press "Add custom integration" button
- there you can find your API key

## Obtaining client ID and client secret

- you can find this inside table clients of your Ghost database
- try with client ID ghost-frontend or ghost-admin and correlated client secret