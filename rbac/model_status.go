/*
 * Role Based Access Control
 *
 * The API for Role Based Access Control.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package rbac
// Status struct for Status
type Status struct {
	ApiVersion int64 `json:"api_version"`
	Commit string `json:"commit,omitempty"`
	ServerAddress string `json:"server_address,omitempty"`
	PlatformInfo map[string]interface{} `json:"platform_info,omitempty"`
	PythonVersion string `json:"python_version,omitempty"`
	Modules map[string]interface{} `json:"modules,omitempty"`
}
