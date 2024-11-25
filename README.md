# k8s examples

We will have 2 pods where the following scenarios will be evaluated:

1. Pod communication through services
2. Downward API
3. ENV variables

## k3s deployment

1. k3s ctr image import testapp.tar
2. k apply -f deployment.yaml

## Helpful tips

- alias k=kubectl
- k api-resources
