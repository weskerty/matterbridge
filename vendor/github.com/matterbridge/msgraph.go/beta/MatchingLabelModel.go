// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MatchingLabel undocumented
type MatchingLabel struct {
	// Object is the base model of MatchingLabel
	Object
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// Name undocumented
	Name *string `json:"name,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// ToolTip undocumented
	ToolTip *string `json:"toolTip,omitempty"`
	// PolicyTip undocumented
	PolicyTip *string `json:"policyTip,omitempty"`
	// IsEndpointProtectionEnabled undocumented
	IsEndpointProtectionEnabled *bool `json:"isEndpointProtectionEnabled,omitempty"`
	// ApplicationMode undocumented
	ApplicationMode *ApplicationMode `json:"applicationMode,omitempty"`
	// LabelActions undocumented
	LabelActions []LabelActionBase `json:"labelActions,omitempty"`
	// Priority undocumented
	Priority *int `json:"priority,omitempty"`
}