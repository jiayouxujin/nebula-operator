/*
Copyright 2021 Vesoft Inc.

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

package v1alpha1

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/pointer"
)

func (nc *NebulaCluster) GraphdComponent() NebulaClusterComponentter {
	return newGraphdComponent(nc)
}

func (nc *NebulaCluster) MetadComponent() NebulaClusterComponentter {
	return newMetadComponent(nc)
}

func (nc *NebulaCluster) StoragedComponent() NebulaClusterComponentter {
	return newStoragedComponent(nc)
}

func (nc *NebulaCluster) ComponentByType(typ ComponentType) (NebulaClusterComponentter, error) {
	switch typ {
	case GraphdComponentType:
		return nc.GraphdComponent(), nil
	case MetadComponentType:
		return nc.MetadComponent(), nil
	case StoragedComponentType:
		return nc.StoragedComponent(), nil
	}

	return nil, fmt.Errorf("unsupport component %s", typ)
}

func (nc *NebulaCluster) GetMetadThriftConnAddress() string {
	return nc.MetadComponent().GetConnAddress(MetadPortNameThrift)
}

func (nc *NebulaCluster) GetMetadEndpoints(portName string) []string {
	return nc.MetadComponent().GetEndpoints(portName)
}

func (nc *NebulaCluster) GetStoragedEndpoints(portName string) []string {
	return nc.StoragedComponent().GetEndpoints(portName)
}

func (nc *NebulaCluster) GetGraphdEndpoints(portName string) []string {
	return nc.GraphdComponent().GetEndpoints(portName)
}

func (nc *NebulaCluster) GetClusterName() string {
	return nc.Name
}

func (nc *NebulaCluster) GenerateOwnerReferences() []metav1.OwnerReference {
	return []metav1.OwnerReference{
		{
			APIVersion:         nc.APIVersion,
			Kind:               nc.Kind,
			Name:               nc.GetName(),
			UID:                nc.GetUID(),
			Controller:         pointer.BoolPtr(true),
			BlockOwnerDeletion: pointer.BoolPtr(true),
		},
	}
}

func (nc *NebulaCluster) IsPVReclaimEnabled() bool {
	enabled := nc.Spec.EnablePVReclaim
	if enabled == nil {
		return false
	}
	return *enabled
}

func (nc *NebulaCluster) IsAutoBalanceEnabled() bool {
	enabled := nc.Spec.Storaged.EnableAutoBalance
	if enabled == nil {
		return false
	}
	return *enabled
}

func (nc *NebulaCluster) IsForceUpdateEnabled() bool {
	enabled := nc.Spec.Storaged.EnableForceUpdate
	if enabled == nil {
		return false
	}
	return *enabled
}

func (nc *NebulaCluster) IsBREnabled() bool {
	enabled := nc.Spec.EnableBR
	if enabled == nil {
		return false
	}
	return *enabled
}

func (nc *NebulaCluster) IsLogRotateEnabled() bool {
	return nc.Spec.LogRotate != nil
}

func (nc *NebulaCluster) InsecureSkipVerify() bool {
	skip := nc.Spec.SSLCerts.InsecureSkipVerify
	if skip == nil {
		return false
	}
	return *skip
}

func (nc *NebulaCluster) IsGraphdSSLEnabled() bool {
	return nc.Spec.Graphd.Config["enable_graph_ssl"] == "true"
}

func (nc *NebulaCluster) IsMetadSSLEnabled() bool {
	return nc.Spec.Graphd.Config["enable_meta_ssl"] == "true" &&
		nc.Spec.Metad.Config["enable_meta_ssl"] == "true" &&
		nc.Spec.Storaged.Config["enable_meta_ssl"] == "true"
}

func (nc *NebulaCluster) IsClusterEnabled() bool {
	return nc.Spec.Graphd.Config["enable_ssl"] == "true" &&
		nc.Spec.Metad.Config["enable_ssl"] == "true" &&
		nc.Spec.Storaged.Config["enable_ssl"] == "true"
}
