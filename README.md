## Setup for local testing

Server setup

Compose: `docker-compose up`

Shell:
- Build Docker: `docker build -f Dockerfile -t mather .`
- Run Docker: `docker run -d -p 5000:5000 --env PORT=5000 --env API_KEY=some-api-key mather`

CLI setup
- Build app: `go build -o mather ./cmd/mather`
- Run help command: `./mather --help`
- Run any of the commands against the local server: `./mather <add, subtract, multiply, divide> x y`

### Additional thinking
Below is information regarding the project

#### 12Factors Alignment
 - Codebase: This codebase defines the app and cli and can be deployed any amount of times.
 - Dependencies: At a code level, dependencies are abstracted and isolated via common
 interfaces and use of primitives for communication between service layers.
 - Config: Configuration can be done via. yaml files, or environment variables,
 with environment variables prioritised/overwrite config file variables.
 - Dependent Services: Currently, the app uses no other resources (e.g. database, another microservice). However,
 if it did, they would run as seperate containers within a k8s pod and would communicate with a client.
 For example, if the current matherapi called another service to produce the math results,
 it would use another gRPC client (abstracted by the interface defined by matherapi) to get it results.
 - CI/CD/Executing: At a low level, the building and running of the app (the Dockerfile) are in seperate
 stages. At a higher level, the CI would be run by CircleCI, CD run by ArgoCD (or similar) and running
 handled by k8s.
 - Stateless Process: The matherapi is a stateless process, any configuration is handled one off time
 at start up.
 - Services/Load balancing: In the `docs/k8s`, you will see the a service is defined that forwards
 onto a specific port binding of the app. The port binding of the app is controlled by a config variable.
 - Concurrency: At a code level, gRPC handles concurrency of incoming requests, spawning a go routine
 for each request (a similar thing happens in the http package of go). When deployed, if more concurrency was required,
 we can scale the app pod horizontally behind the k8s service.
 - Robust Up/Down: Matherapi is quite simple and starts fast. The down is handled using stop signals
 to stop handling new requests. Once deployed, it has both readiness and liveness checks.
 - Environment Parity: Environment parity is handled by kustomize. The large difference is have included
 is the use of a different k8s Service between the environments, to avoid dev being publicly available.
 - Logs: Logs are handled in the background and have no effect on business execution.
 - Admin: Currently, there are no specific admin tasks. However, if any were to occur
 (such as more secure secret retrieval), they would occur during the startup configuration.
 
#### Cloud Native
Matherapi uses not specific cloud technology and could run on any cloud provider. The hosted version
detailed in the External Client section is hosted in GKE on GCP, but uses nothing specific to GCP itself.
It could be argued that CircleCI isn't a cloud native solution to CI, and using something like Argo Workflows would achieve this.
However, for this task, I decided to use CircleCI for a quicker solution.

#### Event Store
Conceptual, all the functions of mather (+, -, *, /) could be applied to a stream of events of two numbers, where the
two numbers are then identified by an ID of some kind.
Therefore, the app would simply instantiate the required operation on the event store object. Although our operations
and events are simple, this could be used for further expansion of operators using the event data (e.g. x^y), or even
potentially act like a cache for past results on new operation instantiation.

#### External Client
If another microservice wanted to use matherapi, it could generate a gRPC client from the proto file in the docs. To share
the proto file it could hosted in a shared repo, or on another API/schema registry for access on startup.

The current CLI could access an external server by setting the `HOST` and `API_KEY` values in `cfg.yaml`, or in the env to
valid remote variables.
