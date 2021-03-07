# Universal Interface

Runtime agnostic, generic interface for all for cloud services, middleware or any other piece of infrastructure your apps may depend on.

This project aims to provide a generic and universal interface for services such as: document based databases, caches, blob storages, messaging brokers,... so you can write your app once and configure it to work with different databases or make use of different cloud services without having to change your code.

## Main Features

### Multi Runtime
Everything is made through an http API(the plan is to provide an OpenAPI schema definition), clients for different programming languagues will be provided.

### Platform agnostic
The project provides a generic interface for using different kinds of services. For now, the only implemented interface is for NoSQL Document databases, such as MongoDB or DynamoDB. 
For app developers the developer experience will be the same regardless of the database used. You may use an in-memory database during development but MongoDB in production and your code will be the same the only change will be in the configuration of the Universal Interface.

### Cloud Native
A common expected usecase is to deploy the interface as close to your app as possible, so the project has been designed to be lightweight, an example of deploying the interface as a sidecar container in kubernetes can be found [here](./examples/k8s). It's also easy to configure, see this [example config file](./examples/config.yaml).

## Try it out

The project is a work in progress, but if you find this thing useful contributions are welcome. :)

You can take a look at the [examples](./examples) folder, it may help you understand what this thing does and how does it work.

And if you are in pro mode and want to try this thing. There is a container image available at `quay.io/famargon/uninterface:latest`
