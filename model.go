package adh

// Config holds settings that are persisted to disk under the user's $HOME
// directory, which describe an Azure DevOps organization/project that will be
// dealt with via API.
type Config struct {
	Organization string `json:"organization,omitempty"`
	Project      string `json:"project,omitempty"`
	Token        string `json:"token,omitempty"`
}

