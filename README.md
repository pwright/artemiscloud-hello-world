# Artemiscloud Hello World

#### A minimal messaging solution deployed in a Kubernetes cluster

This example shows how you can use [Artemis][website] for messaging.

[website]: https://artemiscloud.io/

#### Contents

* [Overview](#overview)
* [Prerequisites](#prerequisites)
* [Step 1: Create a namespace](#step-1-create-a-namespace)
* [Step 2: Deploy operator](#step-2-deploy-operator)
* [Step 3: Check operator](#step-3-check-operator)
* [Step 4: Create a Broker instance](#step-4-create-a-broker-instance)
* [Step 5: Produce and check messages](#step-5-produce-and-check-messages)
* [Step 6: Cleaning up](#step-6-cleaning-up)
* [Summary](#summary)
* [About this example](#about-this-example)

## Overview

This example is a very simple multi-protocol messaging solution.
Before starting, make sure you have golang installed.

## Prerequisites

* The `kubectl` command-line tool, version 1.15 or later
  ([installation guide][install-kubectl])

* Access to at least one Kubernetes cluster, from [any provider you
  choose][kube-providers]

[install-kubectl]: https://kubernetes.io/docs/tasks/tools/install-kubectl/
[kube-providers]: https://skupper.io/start/kubernetes.html

## Step 1: Create a namespace

Use `kubectl create namespace` to create a `hello-world` namespace.

_**Console for hello-world:**_

~~~ shell
kubectl create namespace hello-world
kubectl config set-context --current --namespace hello-world
~~~

_Sample output:_

~~~ console
$ kubectl create namespace hello-world
namespace/hello-world created
~~~

## Step 2: Deploy operator

Deploy the Artemis operator to the current namespace.

_**Console for hello-world:**_

~~~ shell
cd subrepos/operator/deploy; bash ./install_opr.sh
~~~

## Step 3: Check operator

_**Console for hello-world:**_

~~~ shell
kubectl get pods
~~~

_Sample output:_

~~~ console
$ kubectl get pods
activemq-artemis-controller-manager-f8fb97ddd-bjrtv   1/1     Running   0          1m
~~~

## Step 4: Create a Broker instance

_**Console for hello-world:**_

~~~ shell
kubectl create -f subrepos/operator/examples/artemis-basic-deployment.yaml
~~~

_Sample output:_

~~~ console
$ kubectl create -f subrepos/operator/examples/artemis-basic-deployment.yaml
activemqartemis.broker.amq.io/ex-aao created
~~~

## Step 5: Produce and check messages

_**Console for hello-world:**_

~~~ shell
kubectl exec ex-aao-ss-0 -- amq-broker/bin/artemis producer --user x --password y --url tcp://ex-aao-ss-0:61616 --message-count=100
kubectl exec ex-aao-ss-0 -- amq-broker/bin/artemis queue stat --user x --password y --url tcp://ex-aao-ss-0:61616
~~~

_Sample output:_

~~~ console
$ kubectl exec ex-aao-ss-0 -- amq-broker/bin/artemis producer --user x --password y --url tcp://ex-aao-ss-0:61616 --message-count=100
Defaulted container ex-aao-container out of ex-aao-container, ex-aao-container-init (init)

$ kubectl exec ex-aao-ss-0 -- amq-broker/bin/artemis queue stat --user x --password y --url tcp://ex-aao-ss-0:61616
Defaulted container ex-aao-container out of ex-aao-container, ex-aao-container-init (init)
~~~

## Step 6: Cleaning up

_**Console for hello-world:**_

~~~ shell
kubectl delete namespaces hello-world
~~~

## Summary

This example shows how to set up the Artemis operator.

## Next steps

Check out the other [examples][examples] on the Skupper website.

## About this example

This example was produced using [Skewer][skewer], a library for
documenting and testing Skupper examples.

[skewer]: https://github.com/skupperproject/skewer

Skewer provides utility functions for generating the README and
running the example steps.  Use the `./plano` command in the project
root to see what is available.

To quickly stand up the example using Minikube, try the `./plano demo`
command.
