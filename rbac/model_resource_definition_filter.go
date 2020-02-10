/*
 * Role Based Access Control
 *
 * The API for Role Based Access Control.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package rbac
// ResourceDefinitionFilter struct for ResourceDefinitionFilter
type ResourceDefinitionFilter struct {
	Key string `json:"key"`
	Operation string `json:"operation"`
	Value string `json:"value"`
}
