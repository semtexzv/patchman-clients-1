/*
 * Insights Host Inventory REST Interface
 *
 * REST interface for the Insights Platform Host Inventory application.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package inventory
// CanonicalFactsInAllOf struct for CanonicalFactsInAllOf
type CanonicalFactsInAllOf struct {
	// A host’s Fully Qualified Domain Name.  This field is considered to be a canonical fact.
	Fqdn *string `json:"fqdn,omitempty"`
	// A UUID of the host machine BIOS.  This field is considered to be a canonical fact.
	BiosUuid *string `json:"bios_uuid,omitempty"`
	// Host’s reference in the external source e.g. AWS EC2, Azure, OpenStack, etc. This field is considered to be a canonical fact.
	ExternalId *string `json:"external_id,omitempty"`
	// An ID defined in /etc/insights-client/machine-id. This field is considered a canonical fact.
	InsightsId *string `json:"insights_id,omitempty"`
	// Host’s network IP addresses.  This field is considered to be a canonical fact.
	IpAddresses *[]string `json:"ip_addresses,omitempty"`
	// Host’s network interfaces MAC addresses.  This field is considered to be a canonical fact.
	MacAddresses *[]string `json:"mac_addresses,omitempty"`
	// A Machine ID of a RHEL host.  This field is considered to be a canonical fact.
	RhelMachineId *string `json:"rhel_machine_id,omitempty"`
	// A Red Hat Satellite ID of a RHEL host.  This field is considered to be a canonical fact.
	SatelliteId *string `json:"satellite_id,omitempty"`
	// A Red Hat Subcription Manager ID of a RHEL host.  This field is considered to be a canonical fact.
	SubscriptionManagerId *string `json:"subscription_manager_id,omitempty"`
}
