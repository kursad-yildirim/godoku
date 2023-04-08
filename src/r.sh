#! /bin/bash
OLD_VERSION=$(cat k8s/deployment.yaml | grep image: | awk -F ":" '{print $4}')
VERSION=0.18.0
sed -i "s/$OLD_VERSION/$VERSION/g" ./k8s/deployment.yaml
podman rmi default-route-openshift-image-registry.apps.tuff.tripko.local/tuff-apps/godoku:$OLD_VERSION
podman build -t default-route-openshift-image-registry.apps.tuff.tripko.local/tuff-apps/godoku:$VERSION ./
podman push default-route-openshift-image-registry.apps.tuff.tripko.local/tuff-apps/godoku:$VERSION --tls-verify=false

oc delete -f ./k8s/configmap.yaml
oc delete -f ./k8s/service.yaml
oc delete -f ./k8s/route.yaml
oc delete -f ./k8s/deployment.yaml

oc apply -f ./k8s/configmap.yaml
oc apply -f ./k8s/service.yaml
oc apply -f ./k8s/route.yaml
oc apply -f ./k8s/deployment.yaml
