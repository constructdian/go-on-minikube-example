# Go Hello World on Minikube

This project contains scripts and configuration to run a simple Go Hello World application on [minikube](https://github.com/kubernetes/minikube).

The Minikube environment conists of a Kubernetes __deployment__ containing two __pods__, and a Kubernetes __service__ that exposes the webapp acts as a load balancer to these two pods. See [kubernetes concepts](https://kubernetes.io/docs/concepts/) for more information.

## Prerequisites

This project assumes basic familiarity with the Linux shell and a Linux host with the following components correctly set up and installed:

* [minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/) with [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/#install-kubectl-on-linux)
* [docker](https://docs.docker.com/)

This project has been tested on Minikube v1.2.0 and Docker v18.09.7

If you wish to test the webapp locally in a browser, you will also need [xdg-utils](https://www.freedesktop.org/wiki/Software/xdg-utils/) and/or a browser of your choice such as [Lynx](https://lynx.browser.org) installed.

## Usage

1. Clone the git repository, or copy the project directory to a directory of your choice. cd to that directory.
2. To build the Docker Image and start the Kubernetes environment, Run __./build-and-start.sh__
3. To test that that the webapp is running, run one of the following (you may have to wait for a few seconds for the services to start)
    * To open automatically in a browser, run: __minikube service helloworld__
    * To get a URL that you can test locally yourself, run: __minikube service helloworld --url__
4. To stop the webapp and delete the Minikube environment, run __./stop.sh__

You can also run __./build-docker-image.sh__ to build the Docker Image without starting the Kubernetes environment, and __./start.sh__ if you wish to start the Kubernetes environment manually.
