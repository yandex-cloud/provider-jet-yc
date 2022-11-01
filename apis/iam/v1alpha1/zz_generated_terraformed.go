/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this ServiceAccount
func (mg *ServiceAccount) GetTerraformResourceType() string {
	return "yandex_iam_service_account"
}

// GetConnectionDetailsMapping for this ServiceAccount
func (tr *ServiceAccount) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ServiceAccount
func (tr *ServiceAccount) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ServiceAccount
func (tr *ServiceAccount) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ServiceAccount
func (tr *ServiceAccount) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ServiceAccount
func (tr *ServiceAccount) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ServiceAccount
func (tr *ServiceAccount) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ServiceAccount using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ServiceAccount) LateInitialize(attrs []byte) (bool, error) {
	params := &ServiceAccountParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ServiceAccount) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ServiceAccountIAMMember
func (mg *ServiceAccountIAMMember) GetTerraformResourceType() string {
	return "yandex_iam_service_account_iam_member"
}

// GetConnectionDetailsMapping for this ServiceAccountIAMMember
func (tr *ServiceAccountIAMMember) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ServiceAccountIAMMember
func (tr *ServiceAccountIAMMember) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ServiceAccountIAMMember
func (tr *ServiceAccountIAMMember) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ServiceAccountIAMMember
func (tr *ServiceAccountIAMMember) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ServiceAccountIAMMember
func (tr *ServiceAccountIAMMember) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ServiceAccountIAMMember
func (tr *ServiceAccountIAMMember) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ServiceAccountIAMMember using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ServiceAccountIAMMember) LateInitialize(attrs []byte) (bool, error) {
	params := &ServiceAccountIAMMemberParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ServiceAccountIAMMember) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ServiceAccountKey
func (mg *ServiceAccountKey) GetTerraformResourceType() string {
	return "yandex_iam_service_account_key"
}

// GetConnectionDetailsMapping for this ServiceAccountKey
func (tr *ServiceAccountKey) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"private_key": "status.atProvider.privateKey"}
}

// GetObservation of this ServiceAccountKey
func (tr *ServiceAccountKey) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ServiceAccountKey
func (tr *ServiceAccountKey) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ServiceAccountKey
func (tr *ServiceAccountKey) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ServiceAccountKey
func (tr *ServiceAccountKey) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ServiceAccountKey
func (tr *ServiceAccountKey) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ServiceAccountKey using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ServiceAccountKey) LateInitialize(attrs []byte) (bool, error) {
	params := &ServiceAccountKeyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ServiceAccountKey) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ServiceAccountStaticAccessKey
func (mg *ServiceAccountStaticAccessKey) GetTerraformResourceType() string {
	return "yandex_iam_service_account_static_access_key"
}

// GetConnectionDetailsMapping for this ServiceAccountStaticAccessKey
func (tr *ServiceAccountStaticAccessKey) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"secret_key": "status.atProvider.secretKey"}
}

// GetObservation of this ServiceAccountStaticAccessKey
func (tr *ServiceAccountStaticAccessKey) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ServiceAccountStaticAccessKey
func (tr *ServiceAccountStaticAccessKey) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ServiceAccountStaticAccessKey
func (tr *ServiceAccountStaticAccessKey) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ServiceAccountStaticAccessKey
func (tr *ServiceAccountStaticAccessKey) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ServiceAccountStaticAccessKey
func (tr *ServiceAccountStaticAccessKey) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ServiceAccountStaticAccessKey using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ServiceAccountStaticAccessKey) LateInitialize(attrs []byte) (bool, error) {
	params := &ServiceAccountStaticAccessKeyParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ServiceAccountStaticAccessKey) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this FolderIAMBinding
func (mg *FolderIAMBinding) GetTerraformResourceType() string {
	return "yandex_resourcemanager_folder_iam_binding"
}

// GetConnectionDetailsMapping for this FolderIAMBinding
func (tr *FolderIAMBinding) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this FolderIAMBinding
func (tr *FolderIAMBinding) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this FolderIAMBinding
func (tr *FolderIAMBinding) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this FolderIAMBinding
func (tr *FolderIAMBinding) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this FolderIAMBinding
func (tr *FolderIAMBinding) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this FolderIAMBinding
func (tr *FolderIAMBinding) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this FolderIAMBinding using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *FolderIAMBinding) LateInitialize(attrs []byte) (bool, error) {
	params := &FolderIAMBindingParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *FolderIAMBinding) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this FolderIAMMember
func (mg *FolderIAMMember) GetTerraformResourceType() string {
	return "yandex_resourcemanager_folder_iam_member"
}

// GetConnectionDetailsMapping for this FolderIAMMember
func (tr *FolderIAMMember) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this FolderIAMMember
func (tr *FolderIAMMember) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this FolderIAMMember
func (tr *FolderIAMMember) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this FolderIAMMember
func (tr *FolderIAMMember) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this FolderIAMMember
func (tr *FolderIAMMember) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this FolderIAMMember
func (tr *FolderIAMMember) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this FolderIAMMember using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *FolderIAMMember) LateInitialize(attrs []byte) (bool, error) {
	params := &FolderIAMMemberParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *FolderIAMMember) GetTerraformSchemaVersion() int {
	return 0
}
