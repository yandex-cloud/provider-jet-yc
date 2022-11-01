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

package compute

import (
	"fmt"

	"github.com/upbound/upjet/pkg/config"

	"github.com/yandex-cloud/provider-jet-yc/config/iam"
	"github.com/yandex-cloud/provider-jet-yc/config/vpc"
)

// Configure adds configurations for compute group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("yandex_compute_instance", func(r *config.Resource) {
		r.References["network_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["service_account_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", iam.ApisPackagePath, "ServiceAccount"),
		}
		r.References["network_interface.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["network_interface.security_group_ids"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.UseAsync = true
	})
}
