/*
Copyright 2019 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package resources

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"knative.dev/pkg/kmeta"
	"knative.dev/pkg/ptr"

	"github.com/mattmoor/vmware-sources/pkg/apis/sources/v1alpha1"
	"github.com/mattmoor/vmware-sources/pkg/reconciler/vsphere/resources/names"
	"github.com/mattmoor/vmware-sources/pkg/vsphere"
)

func MakeDeployment(ctx context.Context, vms *v1alpha1.VSphereSource, adapterImage string) *appsv1.Deployment {
	labels := map[string]string{
		"vspheresources.sources.knative.dev/name": vms.Name,
	}

	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:            names.DeploymentName(vms),
			Namespace:       vms.Namespace,
			OwnerReferences: []metav1.OwnerReference{*kmeta.NewControllerRef(vms)},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: ptr.Int32(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{{
						Name:  "adapter",
						Image: adapterImage,
						Env: []corev1.EnvVar{{
							Name: "NAMESPACE",
							ValueFrom: &corev1.EnvVarSource{
								FieldRef: &corev1.ObjectFieldSelector{
									FieldPath: "metadata.namespace",
								},
							},
						}, {
							Name:  "K_METRICS_CONFIG",
							Value: "{}",
						}, {
							Name:  "K_LOGGING_CONFIG",
							Value: "{}",
						}, {
							Name:  "GOVMOMI_ADDRESS",
							Value: vms.Spec.Address.String(),
						}, {
							Name:  "GOVMOMI_INSECURE",
							Value: fmt.Sprintf("%v", vms.Spec.SkipTLSVerify),
						}},
						VolumeMounts: []corev1.VolumeMount{{
							Name:      vsphere.VolumeName,
							ReadOnly:  true,
							MountPath: vsphere.MountPath,
						}},
					}},
					Volumes: []corev1.Volume{{
						Name: vsphere.VolumeName,
						VolumeSource: corev1.VolumeSource{
							Secret: &corev1.SecretVolumeSource{
								SecretName: vms.Spec.SecretRef.Name,
							},
						},
					}},
				},
			},
		},
	}
}
