# gosaas

SaaS starter for Go and VueJS with gRPC

## What's included

This little starter stack is powered by VueJS and Go + gRPC.

### Why?

VueJS is a powerful and complete framework for front-end development. It's easy to get started and is perfect for Single Page Apps.

Go is fast evolving into a powerful language for microservices. Coupled with gRPC, it provides a world of strong-typed API contracts that are easy to work with, thanks to the flexibility of protobuf.

### Who should use this?

If you're familiar with Go and you're already convinced that gRPC is the way to go, then this stack is for you. 

Go is easy to pick up, but can be tricky to master. gRPC is quickly becoming the norm in web APIs, but it's quite a far cry from a REST API, and it's likely to take some getting used to.

In short, here's why you should consider go + gRPC:
 - a microservice approach is scalable and extensible
 - API-first is a solid way forward; you can add more clients (mobile apps, chat bots, integrations), and even expose the API to customers
 - microservices in `go` are nimble and performant
 - `proto` files are both code and documentation
 - with gRPC, you call methods with strong-typed arguments and responses / not hand-craft JSON
 - [protobufs are backwards and forwards compatible](https://developers.google.com/protocol-buffers/docs/gotutorial#extending-a-protocol-buffer)

### When to skip this

This approach doesn't work for everything. In particular, you should look elsewhere if:

 - the project is small enough: it doesn't need scalability, multiple components, or a mobile app
 - you don't need a SPA: you may not need a fancy UI for your project
 - you're not familiar with go, gRPC and Vue: it's easy to learn one, but not all 3 at the same time

## Components

Besides the Go and Vue codebase, these components are already setup and ready to be configured.

### Identity management

### Auth0

Auth0 is an identity and access management provider with deep integrations, concise and complete documentation, and SDKs that are easy to use. 

The Vue app uses the Oauth flow (Authorization code flow with PKCE)[https://auth0.com/docs/flows/authorization-code-flow-with-proof-key-for-code-exchange-pkce]. With this flow: 
 - the user logs in with Auth0
 - Auth0 provides an Access Token
 - Vue passes the access token to the Go backend 
 - and the Go backend (uses the access token as a JWT)[https://auth0.com/blog/authentication-in-golang/#Authorization-with-Golang] to get the user's profile

Links:

- https://auth0.com/docs/authorization/which-oauth-2-0-flow-should-i-use
- https://auth0.com/docs/flows/authorization-code-flow-with-proof-key-for-code-exchange-pkce
- https://auth0.com/docs/libraries/auth0-single-page-app-sdk
- https://blog.risingstack.com/auth0-vue-typescript-quickstart-docs/


### 3rd Party Providers

To use 3rd party logins such as Login with Google, you will need to setup an app in the (Google Developer Console)[https://console.developers.google.com/], and setup the Connection in Auth0. Otherwise, certain features will not work (such as the Silent auth that happens on a page refresh).

#### SuperTokens

Go to [https://supertokens.io/] and set things up there. 
