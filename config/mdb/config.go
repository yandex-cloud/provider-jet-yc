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

package mdb

import (
	"fmt"
	"strings"

	"github.com/upbound/upjet/pkg/config"

	"github.com/yandex-cloud/provider-jet-yc/config/iam"
	"github.com/yandex-cloud/provider-jet-yc/config/vpc"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/yandex-cloud/provider-jet-yc/apis/mdb/v1alpha1"
)

func usernames(attr map[string]interface{}) []string {
	usersInterface, ok := attr["user"]
	if !ok {
		return nil
	}
	users, ok := usersInterface.([]interface{})
	if !ok {
		return nil
	}
	result := make([]string, len(users))
	for i, userInterface := range users {
		user, ok := userInterface.(map[string]interface{})
		if !ok {
			continue
		}
		if username, ok := user["name"].(string); ok {
			result[i] = username
		}
	}
	return result
}

func passwords(attr map[string]interface{}) []string {
	usersInterface, ok := attr["user"]
	if !ok {
		return nil
	}
	users, ok := usersInterface.([]interface{})
	if !ok {
		return nil
	}
	result := make([]string, len(users))
	for i, userInterface := range users {
		user, ok := userInterface.(map[string]interface{})
		if !ok {
			continue
		}
		if password, ok := user["password"].(string); ok {
			result[i] = password
		}
	}
	return result
}

func databases(attr map[string]interface{}) []string {
	databasesInterface, ok := attr["database"]
	if !ok {
		return nil
	}
	databases, ok := databasesInterface.([]interface{})
	if !ok {
		return nil
	}
	result := make([]string, len(databases))
	for i, databaseInterface := range databases {
		database, ok := databaseInterface.(map[string]interface{})
		if !ok {
			continue
		}
		if databaseName, ok := database["name"].(string); ok {
			result[i] = databaseName
		}
	}
	return result
}

func fqdns(attr map[string]interface{}) []string {
	hostsInterface, ok := attr["host"]
	if !ok {
		return nil
	}
	hosts, ok := hostsInterface.([]interface{})
	if !ok {
		return nil
	}
	result := make([]string, len(hosts))
	for i, hostInterface := range hosts {
		host, ok := hostInterface.(map[string]interface{})
		if !ok {
			continue
		}
		if fqdn, ok := host["fqdn"].(string); ok {
			result[i] = fqdn
		}
	}
	return result
}

func hostnames(attr map[string]interface{}) []string {
	hostsInterface, ok := attr["host"]
	if !ok {
		return nil
	}
	hosts, ok := hostsInterface.([]interface{})
	if !ok {
		return nil
	}
	result := make([]string, len(hosts))
	for i, hostInterface := range hosts {
		host, ok := hostInterface.(map[string]interface{})
		if !ok {
			continue
		}
		if fqdn, ok := host["name"].(string); ok {
			result[i] = fqdn
		}
	}
	return result
}

func postgresqlConnectionStrings(attr map[string]interface{}) map[string]string {
	connstrings := make(map[string]string)
	hosts := fqdns(attr)
	for _, db := range databases(attr) {
		ps := passwords(attr)
		for i, u := range usernames(attr) {
			password := ps[i]
			connstrings[fmt.Sprintf("connection-string.%s.%s", u, db)] =
				fmt.Sprintf(
					"host=%s port=6432 sslmode=verify-full dbname=%s user=%s target_session_attrs=read-write password=%s",
					strings.Join(hosts, ","), db, u, password)
		}
	}
	return connstrings
}

func mysqlConnectionStrings(attr map[string]interface{}) map[string]string {
	connstrings := make(map[string]string)
	for _, db := range databases(attr) {
		ps := passwords(attr)
		for i, u := range usernames(attr) {
			password := ps[i]
			for _, host := range fqdns(attr) {
				connstrings[fmt.Sprintf("connection-string.%s.%s.%s", u, db, host)] =
					fmt.Sprintf("mysql://%s:%s@%s/%s", u, password, host, db)
			}
		}
	}
	return connstrings
}

func mongodbConnectionStrings(attr map[string]interface{}) map[string]string {
	connstrings := make(map[string]string)
	hosts := hostnames(attr)
	for _, db := range databases(attr) {
		ps := passwords(attr)
		for i, u := range usernames(attr) {
			password := ps[i]
			connstrings[fmt.Sprintf("connection-string.%s.%s", u, db)] =
				fmt.Sprintf(
					"mongodb://%s:%s@%s/%s",
					u, password, strings.Join(hosts, ","), db)
		}
	}
	return connstrings
}

func elasticsearchConnectionStrings(attr map[string]interface{}) map[string]string {
	connstrings := make(map[string]string)
	for _, db := range databases(attr) {
		ps := passwords(attr)
		for i, u := range usernames(attr) {
			password := ps[i]
			for _, host := range fqdns(attr) {
				connstrings[fmt.Sprintf("connection-string.%s.%s.%s", u, db, host)] =
					fmt.Sprintf("https://%s:%s@%s:9200/%s",
						u, password, host, db)
			}

		}
	}
	return connstrings
}

func postgresqlConnDetails(attr map[string]interface{}) map[string][]byte {
	conn := make(map[string][]byte)
	for i, v := range fqdns(attr) {
		conn[fmt.Sprintf("attribute.host.%d.fqdn", i)] = []byte(v)
	}
	for i, v := range usernames(attr) {
		conn[fmt.Sprintf("attribute.user.%d.name", i)] = []byte(v)
	}
	for i, v := range databases(attr) {
		conn[fmt.Sprintf("attribute.database.%d.name", i)] = []byte(v)
	}
	for k, v := range postgresqlConnectionStrings(attr) {
		conn[k] = []byte(v)
	}

	return conn
}

func mysqlConnDetails(attr map[string]interface{}) map[string][]byte {
	conn := make(map[string][]byte)
	for i, v := range fqdns(attr) {
		conn[fmt.Sprintf("attribute.host.%d.fqdn", i)] = []byte(v)
	}
	for i, v := range usernames(attr) {
		conn[fmt.Sprintf("attribute.user.%d.name", i)] = []byte(v)
	}
	for i, v := range databases(attr) {
		conn[fmt.Sprintf("attribute.database.%d.name", i)] = []byte(v)
	}
	for k, v := range mysqlConnectionStrings(attr) {
		conn[k] = []byte(v)
	}

	return conn
}

func mongodbConnDetails(attr map[string]interface{}) map[string][]byte {
	conn := make(map[string][]byte)
	for i, v := range hostnames(attr) {
		conn[fmt.Sprintf("attribute.host.%d.name", i)] = []byte(v)
	}
	for i, v := range usernames(attr) {
		conn[fmt.Sprintf("attribute.user.%d.name", i)] = []byte(v)
	}
	for i, v := range databases(attr) {
		conn[fmt.Sprintf("attribute.database.%d.name", i)] = []byte(v)
	}
	for k, v := range mongodbConnectionStrings(attr) {
		conn[k] = []byte(v)
	}

	return conn
}

func elasticsearchConnDetails(attr map[string]interface{}) map[string][]byte {
	conn := make(map[string][]byte)
	for i, v := range hostnames(attr) {
		conn[fmt.Sprintf("attribute.host.%d.name", i)] = []byte(v)
	}
	for i, v := range usernames(attr) {
		conn[fmt.Sprintf("attribute.user.%d.name", i)] = []byte(v)
	}
	for i, v := range databases(attr) {
		conn[fmt.Sprintf("attribute.database.%d.name", i)] = []byte(v)
	}
	for i, v := range fqdns(attr) {
		conn[fmt.Sprintf("attribute.fqdn.%d.name", i)] = []byte(v)
	}
	for k, v := range elasticsearchConnectionStrings(attr) {
		conn[k] = []byte(v)
	}
	if clusterID, ok := attr["id"]; ok {
		if clusterIDString, ok := clusterID.(string); ok {
			conn["attribute.id"] = []byte(clusterIDString)
		}
	}
	return conn
}

func redisConnDetails(attr map[string]interface{}) map[string][]byte {
	conn := make(map[string][]byte)
	for i, v := range fqdns(attr) {
		conn[fmt.Sprintf("attribute.host.%d.fqdn", i)] = []byte(v)
	}
	for i, v := range databases(attr) {
		conn[fmt.Sprintf("attribute.database.%d.name", i)] = []byte(v)
	}

	return conn
}

// Configure adds configurations for mdb group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("yandex_mdb_postgresql_cluster", func(r *config.Resource) {
		r.References["network_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["host.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["security_group_ids"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.UseAsync = true
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			return postgresqlConnDetails(attr), nil
		}
	})
	p.AddResourceConfigurator("yandex_mdb_postgresql_database", func(r *config.Resource) {
		r.References["cluster_id"] = config.Reference{
			Type: "PostgresqlCluster",
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("yandex_mdb_postgresql_user", func(r *config.Resource) {
		r.References["cluster_id"] = config.Reference{
			Type: "PostgresqlCluster",
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("yandex_mdb_mysql_cluster", func(r *config.Resource) {
		r.References["network_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["host.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["security_group_ids"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.UseAsync = true
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			return mysqlConnDetails(attr), nil
		}
	})
	p.AddResourceConfigurator("yandex_mdb_mysql_database", func(r *config.Resource) {
		r.References["cluster_id"] = config.Reference{
			Type: "MySQLCluster",
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("yandex_mdb_mysql_user", func(r *config.Resource) {
		r.References["cluster_id"] = config.Reference{
			Type: "MySQLCluster",
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("yandex_mdb_redis_cluster", func(r *config.Resource) {
		r.References["network_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["host.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["security_group_ids"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.UseAsync = true
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			return redisConnDetails(attr), nil
		}
	})
	p.AddResourceConfigurator("yandex_mdb_mongodb_cluster", func(r *config.Resource) {
		r.References["network_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["host.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["security_group_ids"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.UseAsync = true
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			return mongodbConnDetails(attr), nil
		}
	})
	p.AddResourceConfigurator("yandex_mdb_elasticsearch_cluster", func(r *config.Resource) {
		r.References["network_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["host.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["security_group_ids"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.References["service_account_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", iam.ApisPackagePath, "ServiceAccount"),
		}
		r.UseAsync = true
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			return elasticsearchConnDetails(attr), nil
		}
	})
}
