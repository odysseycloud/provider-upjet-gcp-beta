// SPDX-FileCopyrightText: 2025 Upbound Inc. <https://upbound.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	"dario.cat/mergo"
	"github.com/pkg/errors"

	"github.com/crossplane/upjet/pkg/resource"
	"github.com/crossplane/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this Cluster
func (mg *Cluster) GetTerraformResourceType() string {
	return "google_container_cluster"
}

// GetConnectionDetailsMapping for this Cluster
func (tr *Cluster) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"master_auth[*].client_key": "status.atProvider.masterAuth[*].clientKey"}
}

// GetObservation of this Cluster
func (tr *Cluster) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Cluster
func (tr *Cluster) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Cluster
func (tr *Cluster) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Cluster
func (tr *Cluster) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Cluster
func (tr *Cluster) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this Cluster
func (tr *Cluster) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// GetInitParameters of this Cluster
func (tr *Cluster) GetMergedParameters(shouldMergeInitProvider bool) (map[string]any, error) {
	params, err := tr.GetParameters()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get parameters for resource '%q'", tr.GetName())
	}
	if !shouldMergeInitProvider {
		return params, nil
	}

	initParams, err := tr.GetInitParameters()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get init parameters for resource '%q'", tr.GetName())
	}

	// Note(lsviben): mergo.WithSliceDeepCopy is needed to merge the
	// slices from the initProvider to forProvider. As it also sets
	// overwrite to true, we need to set it back to false, we don't
	// want to overwrite the forProvider fields with the initProvider
	// fields.
	err = mergo.Merge(&params, initParams, mergo.WithSliceDeepCopy, func(c *mergo.Config) {
		c.Overwrite = false
	})
	if err != nil {
		return nil, errors.Wrapf(err, "cannot merge spec.initProvider and spec.forProvider parameters for resource '%q'", tr.GetName())
	}

	return params, nil
}

// LateInitialize this Cluster using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Cluster) LateInitialize(attrs []byte) (bool, error) {
	params := &ClusterParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}
	opts = append(opts, resource.WithNameFilter("AddonsConfig.DNSCacheConfig"))
	opts = append(opts, resource.WithNameFilter("AddonsConfig.GCPFilestoreCsiDriverConfig"))
	opts = append(opts, resource.WithNameFilter("AddonsConfig.GcsFuseCsiDriverConfig"))
	opts = append(opts, resource.WithNameFilter("AddonsConfig.NetworkPolicyConfig"))
	opts = append(opts, resource.WithNameFilter("ClusterAutoscaling.Enabled"))
	opts = append(opts, resource.WithNameFilter("ClusterAutoscaling.ResourceLimits"))
	opts = append(opts, resource.WithNameFilter("ClusterIPv4Cidr"))
	opts = append(opts, resource.WithNameFilter("DefaultMaxPodsPerNode"))
	opts = append(opts, resource.WithNameFilter("EnableAutopilot"))
	opts = append(opts, resource.WithNameFilter("EnableIntranodeVisibility"))
	opts = append(opts, resource.WithNameFilter("EnableShieldedNodes"))
	opts = append(opts, resource.WithNameFilter("IPAllocationPolicy"))
	opts = append(opts, resource.WithNameFilter("NetworkPolicy"))
	opts = append(opts, resource.WithNameFilter("NodeConfig"))
	opts = append(opts, resource.WithNameFilter("NodeVersion"))
	opts = append(opts, resource.WithNameFilter("WorkloadIdentityConfig"))

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Cluster) GetTerraformSchemaVersion() int {
	return 2
}
