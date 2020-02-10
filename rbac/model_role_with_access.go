/*
 * Role Based Access Control
 *
 * The API for Role Based Access Control.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package rbac
// RoleWithAccess struct for RoleWithAccess
type RoleWithAccess struct {
	Name string `json:"name"`
	Description string `json:"description,omitempty"`
	Uuid string `json:"uuid"`
	Created string `json:"created"`
	Modified string `json:"modified"`
	PolicyCount int32 `json:"policyCount,omitempty"`
	AccessCount int32 `json:"accessCount,omitempty"`
	Applications []string `json:"applications,omitempty"`
	System bool `json:"system,omitempty"`
	PlatformDefault bool `json:"platform_default,omitempty"`
	Access []Access `json:"access"`
}
