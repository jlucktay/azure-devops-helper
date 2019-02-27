package adh

type config struct {
	ApiVersion   string `json:"api-version,omitempty"`
	Organization string `json:"organization,omitempty"`
	Project      string `json:"project,omitempty"`
	Token        string `json:"token,omitempty"`
}

/*
https://
dev.azure.com/
DanClzAutomation/
DanMigrationFactory/
_apis/
build/
builds
?
api-version=5.0
*/
