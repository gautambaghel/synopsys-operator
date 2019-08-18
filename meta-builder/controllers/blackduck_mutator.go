/*
 * Copyright (C) $year Synopsys, Inc.
 *
 *  Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 *  with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 *  under the License.
 */

/*

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

package controllers

import (
	"fmt"
	synopsysv1 "github.com/blackducksoftware/synopsys-operator/meta-builder/api/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/runtime"
	"strconv"
	"strings"
)

func patchBlackduck(blackduck *synopsysv1.Blackduck, objects map[string]runtime.Object) map[string]runtime.Object {
	patcher := BlackduckPatcher{
		blackduck: blackduck,
		objects:   objects,
	}
	return patcher.patch()
}

type BlackduckPatcher struct {
	blackduck *synopsysv1.Blackduck
	objects   map[string]runtime.Object
}

func (p *BlackduckPatcher) patch() map[string]runtime.Object {
	// TODO JD: Patching this way is costly. Consider iterating over the objects only once
	// and apply the necessary changes
	p.patchNamespace()
	p.patchStorage()
	p.patchLiveness()
	p.patchEnvirons()
	p.patchWebserverCertificates()
	p.patchPostgresConfig()
	p.patchImages()
	p.patchReplicas()
	p.patchAuthCert()
	p.patchProxyCert()
	// TODO - Patch ImageRegistries | RegistryConfiguration
	// DesiredState | ExposeService
	return p.objects
}

func (p *BlackduckPatcher) patchAuthCert() error {
	if len(p.blackduck.Spec.AuthCustomCA) == 0 {
		for _, v := range p.objects {
			switch v.(type) {
			case *v1.ReplicationController:
				removeVolumeAndVolumeMountFromRC(v.(*v1.ReplicationController), "blackduck-auth-custom-ca")
			}
		}
	} else {
		secret, ok := p.objects["Secret.blackduck-auth-custom-ca"]
		if !ok {
			return nil
		}

		if secret.(*v1.Secret).Data == nil {
			secret.(*v1.Secret).Data = make(map[string][]byte)
		}

		secret.(*v1.Secret).Data["AUTH_CUSTOM_CA"] = []byte(p.blackduck.Spec.AuthCustomCA)
	}
	return nil
}

func (p *BlackduckPatcher) patchProxyCert() error {
	if len(p.blackduck.Spec.ProxyCertificate) == 0 {
		for _, v := range p.objects {
			switch v.(type) {
			case *v1.ReplicationController:
				removeVolumeAndVolumeMountFromRC(v.(*v1.ReplicationController), "blackduck-proxy-certificate")
			}
		}
	} else {
		secret, ok := p.objects["Secret.blackduck-proxy-certificate"]
		if !ok {
			return nil
		}

		if secret.(*v1.Secret).Data == nil {
			secret.(*v1.Secret).Data = make(map[string][]byte)
		}

		secret.(*v1.Secret).Data["HUB_PROXY_CERT_FILE"] = []byte(p.blackduck.Spec.ProxyCertificate)
	}
	return nil
}

func (p *BlackduckPatcher) patchReplicas() error {
	for _, v := range p.objects {
		switch v.(type) {
		case *v1.ReplicationController:
			switch p.blackduck.Spec.DesiredState {
			case "STOP":
				v.(*v1.ReplicationController).Spec.Replicas = func(i int32) *int32 { return &i }(0)
			case "DBMIGRATE":
			// TODO
			default:
				// TODO apply replica from flavor configuration
			}
		}
	}
	return nil
}

func (p *BlackduckPatcher) patchImages() error {
	if len(p.blackduck.Spec.RegistryConfiguration.Registry) > 0 || len(p.blackduck.Spec.ImageRegistries) > 0 {
		for _, v := range p.objects {
			switch v.(type) {
			case *v1.ReplicationController:
				//for i := range v.(*v1.ReplicationController).Spec.Template.Spec.Containers {
				// TODO
				//}

			}

		}
	}
	return nil
}

func (p *BlackduckPatcher) patchPostgresConfig() error {
	cmConf, ok := p.objects["ConfigMap.blackduck-db-config"]
	if !ok {
		return nil
	}

	secretConf, ok := p.objects["Secret.blackduck-db-creds"]
	if !ok {
		return nil
	}

	if cmConf.(*v1.ConfigMap).Data == nil {
		cmConf.(*v1.ConfigMap).Data = make(map[string]string)
	}

	if secretConf.(*v1.Secret).Data == nil {
		secretConf.(*v1.Secret).Data = make(map[string][]byte)
	}

	if p.blackduck.Spec.ExternalPostgres != nil {
		cmConf.(*v1.ConfigMap).Data["HUB_POSTGRES_ADMIN"] = p.blackduck.Spec.ExternalPostgres.PostgresAdmin
		cmConf.(*v1.ConfigMap).Data["HUB_POSTGRES_ENABLE_SSL"] = strconv.FormatBool(p.blackduck.Spec.ExternalPostgres.PostgresSsl)
		cmConf.(*v1.ConfigMap).Data["HUB_POSTGRES_HOST"] = p.blackduck.Spec.ExternalPostgres.PostgresHost
		cmConf.(*v1.ConfigMap).Data["HUB_POSTGRES_PORT"] = strconv.Itoa(p.blackduck.Spec.ExternalPostgres.PostgresPort)
		cmConf.(*v1.ConfigMap).Data["HUB_POSTGRES_USER"] = p.blackduck.Spec.ExternalPostgres.PostgresUser

		secretConf.(*v1.Secret).Data["HUB_POSTGRES_ADMIN_PASSWORD_FILE"] = []byte(p.blackduck.Spec.ExternalPostgres.PostgresAdminPassword)
		secretConf.(*v1.Secret).Data["HUB_POSTGRES_USER_PASSWORD_FILE"] = []byte(p.blackduck.Spec.ExternalPostgres.PostgresUserPassword)

		// Delete the component required when deploying internal postgres
		delete(p.objects, "PersistentVolumeClaim.blackduck-postgres")
		delete(p.objects, "Job.blackduck-init-postgres")
		delete(p.objects, "ConfigMap.postgres-init-config")
		delete(p.objects, "Service.blackduck-postgres")
		delete(p.objects, "ReplicationController.blackduck-postgres")
	} else {
		cmConf.(*v1.ConfigMap).Data["HUB_POSTGRES_ADMIN"] = "blackduck"
		cmConf.(*v1.ConfigMap).Data["HUB_POSTGRES_ENABLE_SSL"] = "false"
		cmConf.(*v1.ConfigMap).Data["HUB_POSTGRES_HOST"] = "blackduck-postgres"
		cmConf.(*v1.ConfigMap).Data["HUB_POSTGRES_PORT"] = "5432"
		cmConf.(*v1.ConfigMap).Data["HUB_POSTGRES_USER"] = "blackduck_user"

		secretConf.(*v1.Secret).Data["HUB_POSTGRES_ADMIN_PASSWORD_FILE"] = []byte(p.blackduck.Spec.AdminPassword)
		secretConf.(*v1.Secret).Data["HUB_POSTGRES_USER_PASSWORD_FILE"] = []byte(p.blackduck.Spec.UserPassword)
		secretConf.(*v1.Secret).Data["HUB_POSTGRES_POSTGRES_PASSWORD_FILE"] = []byte(p.blackduck.Spec.PostgresPassword)
	}

	return nil
}

func (p *BlackduckPatcher) patchWebserverCertificates() error {

	if len(p.blackduck.Spec.Certificate) > 0 && len(p.blackduck.Spec.CertificateKey) > 0 {
		id := "Secret.blackduck-webserver-certificate"
		runtimeObject, ok := p.objects[id]
		if !ok {
			return nil
		}
		runtimeObject.(*v1.Secret).Data["WEBSERVER_CUSTOM_CERT_FILE"] = []byte(p.blackduck.Spec.Certificate)
		runtimeObject.(*v1.Secret).Data["WEBSERVER_CUSTOM_KEY_FILE"] = []byte(p.blackduck.Spec.CertificateKey)

	}

	return nil
}

func (p *BlackduckPatcher) patchEnvirons() error {
	ConfigMapUniqueID := "ConfigMap.blackduck-config"
	configMapRuntimeObject, ok := p.objects[ConfigMapUniqueID]
	if !ok {
		return nil
	}
	configMap := configMapRuntimeObject.(*v1.ConfigMap)
	for _, e := range p.blackduck.Spec.Environs {
		vals := strings.Split(e, ":") // TODO - doesn't handle multiple colons
		if len(vals) != 2 {
			fmt.Printf("Could not split environ '%s' on ':'\n", e) // TODO change to log
			continue
		}
		environKey := strings.TrimSpace(vals[0])
		environVal := strings.TrimSpace(vals[1])
		configMap.Data[environKey] = environVal
	}
	return nil
}
func (p *BlackduckPatcher) patchNamespace() error {
	accessor := meta.NewAccessor()
	for _, runtimeObject := range p.objects {
		accessor.SetNamespace(runtimeObject, p.blackduck.Spec.Namespace)
	}
	return nil
}

func (p *BlackduckPatcher) patchLiveness() error {
	// Removes liveness probes if Spec.LivenessProbes is set to false
	for _, v := range p.objects {
		switch v.(type) {
		case *v1.ReplicationController:
			if !p.blackduck.Spec.LivenessProbes {
				for i := range v.(*v1.ReplicationController).Spec.Template.Spec.Containers {
					v.(*v1.ReplicationController).Spec.Template.Spec.Containers[i].LivenessProbe = nil
				}
			}
		}
	}
	return nil
}

func (p *BlackduckPatcher) patchStorage() error {
	for k, v := range p.objects {
		switch v.(type) {
		case *v1.PersistentVolumeClaim:
			if !p.blackduck.Spec.PersistentStorage {
				delete(p.objects, k)
			} else {
				if len(p.blackduck.Spec.PVCStorageClass) > 0 {
					v.(*v1.PersistentVolumeClaim).Spec.StorageClassName = &p.blackduck.Spec.PVCStorageClass
				}
				for _, pvc := range p.blackduck.Spec.PVC {
					if strings.EqualFold(pvc.Name, v.(*v1.PersistentVolumeClaim).Name) {
						v.(*v1.PersistentVolumeClaim).Spec.VolumeName = pvc.VolumeName
						v.(*v1.PersistentVolumeClaim).Spec.StorageClassName = &pvc.StorageClass
						if quantity, err := resource.ParseQuantity(pvc.Size); err == nil {
							v.(*v1.PersistentVolumeClaim).Spec.Resources.Requests[v1.ResourceStorage] = quantity
						}
					}
				}
			}
		case *v1.ReplicationController:
			if !p.blackduck.Spec.PersistentStorage {
				for i := range v.(*v1.ReplicationController).Spec.Template.Spec.Volumes {
					// If PersistentVolumeClaim then we change it to emptyDir
					if v.(*v1.ReplicationController).Spec.Template.Spec.Volumes[i].VolumeSource.PersistentVolumeClaim != nil {
						v.(*v1.ReplicationController).Spec.Template.Spec.Volumes[i].VolumeSource = v1.VolumeSource{
							EmptyDir: &v1.EmptyDirVolumeSource{
								Medium:    v1.StorageMediumDefault,
								SizeLimit: nil,
							},
						}
					}
				}
			}
		}

	}
	return nil
}

func removeVolumeAndVolumeMountFromRC(rc *v1.ReplicationController, volumeName string) *v1.ReplicationController {
	for volumeNb, volume := range rc.Spec.Template.Spec.Volumes {
		if volume.Secret != nil && strings.Compare(volume.Secret.SecretName, volumeName) == 0 {
			rc.Spec.Template.Spec.Volumes = append(rc.Spec.Template.Spec.Volumes[:volumeNb], rc.Spec.Template.Spec.Volumes[volumeNb+1:]...)
			for containerNb, container := range rc.Spec.Template.Spec.Containers {
				for volumeMountNb, volumeMount := range container.VolumeMounts {
					if strings.Compare(volumeMount.Name, volume.Name) == 0 {
						rc.Spec.Template.Spec.Containers[containerNb].VolumeMounts = append(rc.Spec.Template.Spec.Containers[containerNb].VolumeMounts[:volumeMountNb], rc.Spec.Template.Spec.Containers[containerNb].VolumeMounts[volumeMountNb+1:]...)
					}
				}
			}
		}
	}
	return rc
}