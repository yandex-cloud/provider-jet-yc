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

package dns

import (
	"fmt"

	"github.com/upbound/upjet/pkg/config"

	"github.com/yandex-cloud/provider-jet-yc/config/vpc"
)

// Configure adds configurations for dns group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("yandex_dns_zone", func(r *config.Resource) {
		r.References["private_networks"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
	})
	p.AddResourceConfigurator("yandex_dns_recordset", func(r *config.Resource) {
		r.References["zone_id"] = config.Reference{
			Type: "Zone",
		}
	})
}
