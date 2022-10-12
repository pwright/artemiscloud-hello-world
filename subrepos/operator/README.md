# Artemiscloud Hello World

#### A minimal messaging solution deployed in a Kubernetes cluster

This example is part of a [suite of examples][examples] showing the
different ways you can use [Skupper][website] to connect services
across cloud providers, data centers, and edge sites.

[website]: https://skupper.io/
[examples]: https://skupper.io/examples/index.html

#### Contents

* [Overview](#overview)
* [Prerequisites](#prerequisites)
* [Step 1: Create a namespace](#step-1-create-a-namespace)
* [Step 2: Cleaning up](#step-2-cleaning-up)
* [Summary](#summary)
* [About this example](#about-this-example)

## Overview

This example is a very simple multi-protocol messaging solution.

## Prerequisites

* The `kubectl` command-line tool, version 1.15 or later
  ([installation guide][install-kubectl])

* Access to at least one Kubernetes cluster, from [any provider you
  choose][kube-providers]

[install-kubectl]: https://kubernetes.io/docs/tasks/tools/install-kubectl/
[kube-providers]: https://skupper.io/start/kubernetes.html

## Step 1: Create a namespace

Use `kubectl create ns` to create a namespace.

_**Console for hello-world:**_

~~~ shell
kubectl create namespace hello-world
~~~

## Step 2: Cleaning up

_**Console for hello-world:**_

~~~ shell
namespace/hello-world created
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
