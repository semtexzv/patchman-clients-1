/*
 * VMaaS Webapp
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.4.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package vmaas
// CvesResponseCveList struct for CvesResponseCveList
type CvesResponseCveList struct {
	Impact string `json:"impact,omitempty"`
	PublicDate string `json:"public_date,omitempty"`
	Synopsis string `json:"synopsis,omitempty"`
	Description string `json:"description,omitempty"`
	ModifiedDate string `json:"modified_date,omitempty"`
	RedhatUrl string `json:"redhat_url,omitempty"`
	SecondaryUrl string `json:"secondary_url,omitempty"`
	Cvss2Score string `json:"cvss2_score,omitempty"`
	Cvss2Metrics string `json:"cvss2_metrics,omitempty"`
	Cvss3Score string `json:"cvss3_score,omitempty"`
	Cvss3Metrics string `json:"cvss3_metrics,omitempty"`
	CweList []string `json:"cwe_list,omitempty"`
	ErrataList []string `json:"errata_list,omitempty"`
	PackageList []string `json:"package_list,omitempty"`
	SourcePackageList []string `json:"source_package_list,omitempty"`
}
