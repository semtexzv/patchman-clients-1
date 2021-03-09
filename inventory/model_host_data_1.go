/*
 * Insights Host Inventory REST Interface
 *
 * REST interface for the Insights Platform Host Inventory application.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package inventory
// HostData1 struct for HostData1
type HostData1 struct {
	// The number of items on the current page
	Count int32 `json:"count"`
	// The page number
	Page int32 `json:"page"`
	// The number of items to return per page
	PerPage int32 `json:"per_page"`
	Results []ActiveTag `json:"results"`
	// Total number of items
	Total int32 `json:"total"`
}
