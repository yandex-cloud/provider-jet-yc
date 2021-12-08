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

package iam

import (
	"bb.yandex-team.ru/crossplane/provider-jet-yc/config/resourcemanager"
	"fmt"
	"github.com/crossplane-contrib/terrajet/pkg/config"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "bb.yandex-team.ru/crossplane/provider-jet-yc/apis/iam/v1alpha1"
)

// Configure adds configurations for iam group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("yandex_iam_service_account", func(r *config.Resource) {
		r.References["folder_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", resourcemanager.ApisPackagePath, "Folder"),
		}
	})
	p.AddResourceConfigurator("yandex_iam_service_account_key", func(r *config.Resource) {
		r.References["service_account_id"] = config.Reference{
			Type: "ServiceAccount",
		}
	})
	p.AddResourceConfigurator("yandex_iam_service_account_iam_binding", func(r *config.Resource) {
		r.References["service_account_id"] = config.Reference{
			Type: "ServiceAccount",
		}
	})
}