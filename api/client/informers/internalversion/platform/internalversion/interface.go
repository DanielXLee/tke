/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by informer-gen. DO NOT EDIT.

package internalversion

import (
	internalinterfaces "tkestack.io/tke/api/client/informers/internalversion/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// CSIOperators returns a CSIOperatorInformer.
	CSIOperators() CSIOperatorInformer
	// Clusters returns a ClusterInformer.
	Clusters() ClusterInformer
	// ClusterCredentials returns a ClusterCredentialInformer.
	ClusterCredentials() ClusterCredentialInformer
	// ConfigMaps returns a ConfigMapInformer.
	ConfigMaps() ConfigMapInformer
	// CronHPAs returns a CronHPAInformer.
	CronHPAs() CronHPAInformer
	// Machines returns a MachineInformer.
	Machines() MachineInformer
	// PersistentEvents returns a PersistentEventInformer.
	PersistentEvents() PersistentEventInformer
	// Prometheuses returns a PrometheusInformer.
	Prometheuses() PrometheusInformer
	// Registries returns a RegistryInformer.
	Registries() RegistryInformer
	// TappControllers returns a TappControllerInformer.
	TappControllers() TappControllerInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// CSIOperators returns a CSIOperatorInformer.
func (v *version) CSIOperators() CSIOperatorInformer {
	return &cSIOperatorInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Clusters returns a ClusterInformer.
func (v *version) Clusters() ClusterInformer {
	return &clusterInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterCredentials returns a ClusterCredentialInformer.
func (v *version) ClusterCredentials() ClusterCredentialInformer {
	return &clusterCredentialInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ConfigMaps returns a ConfigMapInformer.
func (v *version) ConfigMaps() ConfigMapInformer {
	return &configMapInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// CronHPAs returns a CronHPAInformer.
func (v *version) CronHPAs() CronHPAInformer {
	return &cronHPAInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Machines returns a MachineInformer.
func (v *version) Machines() MachineInformer {
	return &machineInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// PersistentEvents returns a PersistentEventInformer.
func (v *version) PersistentEvents() PersistentEventInformer {
	return &persistentEventInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Prometheuses returns a PrometheusInformer.
func (v *version) Prometheuses() PrometheusInformer {
	return &prometheusInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Registries returns a RegistryInformer.
func (v *version) Registries() RegistryInformer {
	return &registryInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// TappControllers returns a TappControllerInformer.
func (v *version) TappControllers() TappControllerInformer {
	return &tappControllerInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
