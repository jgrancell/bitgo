# Bitgo

This project provides an API interface that calls the Coinbase API to check the spot price for
a fiat currency against Bitcoin.

## Architecture and Explanation

### Application
The application, `bitgo`, is written in Golang using three non-native Go modules:
- My own custom logger to support both klog and CLI format logs
- [httprouter](https://github.com/julienschmidt/httprouter) for easier routing and DRY code
- [promhttp](https://github.com/prometheus/client_golang) for prometheus metrics integration

The application runs at [https://bitgo.joshgrancell.com](https://bitgo.joshgrancell.com), where you can
see a full list of API endpoints.

### CI Pipeline
Github Actions is currently configured so that on any tagged release matching the regex `v*` it will:
- Build and publish a Docker image to Docker hub at `jgrancell/bitgo`.
- Spin up a bitgo container and test several endpoints to validate its functionality.

**NOTE**: The CI pipeline is expected to fail. There is a final Github Action that calls an invalid currency path
which triggers a failure. This is expected, and shows that we can catch errors when we can't decode JSON.

There is also an example of how to automate the Terraform deployment of the Helm Chart by tying it into ArgoCD. Again, this is only triggered on versioned releases, so the Helm Chart is only ever deploying tags that match `v*`


### GitOps
This repository, specifically the `deploy/charts` Helm Chart directory, is watched by ArgoCD. On updates, ArgoCD will reconcile any changes in the live cluster.

The Helm Chart includes the option to deploy in a Highly Available configuration. When this is enabled, the pods
will have an anti-affinity rule ensuring that they are all running on different nodes. Therefore, if a single Kubernetes node goes down it will not bring down the entire application.
