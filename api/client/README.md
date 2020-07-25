# Client Service

The Client api manages all inbound api requests for m3o client libraries.

## Overview

The micro api provides a public entrypoint mapping http/json requests to backend rpc service calls. 
This is great because /foo/bar translates to the foo service with endpoint Foo.Bar but building 
clients against this in the long term means we have to bake in a lot of things into the API. 
Instead writing a client api service lets us manage all api access for clients in one place.

## Usage

Clients are the m3o-{node, angular, java, ...} client libraries written to execute against 
the endpoint api.m3o.com/client or localhost:8080/api/client, whichever is preferrable.
