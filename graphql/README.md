# GraphQL API

This is the graphql api served at /graphql

WIP

## Overview

The graphql api is a federated api gateway which fans out queries to different microservices 
on the backend. We're using gqlgen to generate the schema off the protos in this directory.

## Implementation

1. GitHub actions workflow to find per service protos (only backend services)
2. Generate the schemas needed
3. Code generated the resolvers
4. Register function/handler at known location
