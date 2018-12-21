# Sample Envoy Sidecar for TLS Termination
This project is an example of using envoy (https://www.envoyproxy.io) as a proxy for service within a container. It runs as a sidecar much like Nginx with TLS terminating at envoy and then envoy communicating with the underlying service (via http) directly.

The code here is just a simple Go service that exposes an API over http. It utilizes our make system and all of those bits. To run this you will need first build the project:


```
	make docker
```

And then you can run the "system" which will spin up an instance of Vault to serve as our CA, and an Envoy sidecar for TLS termination fronting this service.


```
	docker-compose up
```

We run 2 instances of this service with 2 sidecars. To access server one you can go to the following URL:

```
	https://localhost:500/info
```

or the second instance at:
```
	https://localhost:550/info
```

The docker-compose.yml file is done such that direct ports exposed by this service are not visible from outside the docker network. Therefore you can not reach the services in an unsecure manner.