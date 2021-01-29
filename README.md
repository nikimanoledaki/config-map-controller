# Kubernetes ConfigMap Controller

This is a project to build a Kubernetes Controller from scratch using the Watcher method. It is a Controller that watches for the creation of new ConfigMaps.

## How to run it
This project has the following dependencies. Make sure to install these locally before proceeding.
- [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)
- [direnv](https://direnv.net/#basic-installation) is used (it's the `.envrc` file) to export the KUBECONFIG environment variable in your shell. 

Follow the following steps to spin up this application.

1. Do `make up` to spin up a cluster called `playground` using kind.

2. Build the binary with `make build`.

3. Run the application with `make run`.

4. In a separate terminal, create a ConfigMap out of the manifest in `./manifests/configmap.yaml` and deploy it to your local kind cluster:

```
kubectl create configmap joke --from-file=./manifests/configmap.yaml
```

## Resources
Here are some of the resources that helped me along the way that I would like to credit:
- https://github.com/aclevername/config-map-controller
- https://github.com/kubernetes/client-go/tree/master/examples#advanced-concepts
- https://engineering.bitnami.com/articles/a-deep-dive-into-kubernetes-controllers.html
- https://github.com/kubernetes/community/blob/8decfe4/contributors/devel/controllers.md
- https://github.com/kubernetes/sample-controller
