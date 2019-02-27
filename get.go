package adh

import (
	"fmt"
	"os"
	"regexp"
)

var configPath = os.ExpandEnv("${HOME}/.config/azure-devops-helper/adh.json")

// validatePAT makes sure that the given token matches the format of Azure DevOps
func validatePAT(token string) (string, error) {
	rx := regexp.MustCompile("^[a-z0-9]{52}$")
	if !rx.Match([]byte(token)) {
		return "", fmt.Errorf("the value for the 'token' key in the config file (%s) is not correctly formed", configPath)
	}

	return token, nil
}
