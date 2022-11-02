package containers

import (
	corev1 "k8s.io/api/core/v1"
)

const (
	graceTime       = 30
	TCPLivenessPort = 8161
)

func MakeContainer(hostingPodSpec *corev1.PodSpec, customResourceName string, imageName string, envVarArray []corev1.EnvVar) *corev1.Container {

	name := customResourceName + "-container"

	var container *corev1.Container

	if hostingPodSpec != nil {
		for _, existingContainer := range hostingPodSpec.Containers {
			if existingContainer.Name == name {
				container = &existingContainer
				break
			}
		}
	}
	if container == nil {
		container = &corev1.Container{
			Name:    name,
			Command: []string{"/bin/sh", "-c", "export STATEFUL_SET_ORDINAL=${HOSTNAME##*-};exec /opt/amq/bin/launch.sh", "start"},
		}
	}

	container.Image = imageName
	container.Env = envVarArray

	return container
}

func MakeInitContainer(hostingPodSpec *corev1.PodSpec, customResourceName string, imageName string, envVarArray []corev1.EnvVar) *corev1.Container {

	name := customResourceName + "-container-init"

	var container *corev1.Container

	if hostingPodSpec != nil {
		for _, existingContainer := range hostingPodSpec.InitContainers {
			if existingContainer.Name == name {
				container = &existingContainer
				break
			}
		}
	}
	if container == nil {
		container = &corev1.Container{
			Name:    name,
			Command: []string{"/bin/bash"},
		}
	}

	container.Image = imageName
	container.Env = envVarArray

	return container
}
