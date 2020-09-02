package validator

import (
	"log"

	"github.com/AchoArnold/ov-chipkaart-dashboard/backend/shared/ov-chipkaart"
)

// Helpers are custom validations that are package agnostic
type Helpers struct {
	ovChipkaartAPIClient ovchipkaart.APIClient
}

// NewHelpers creates new helpers
func NewHelpers(ovChipkaartAPIClient ovchipkaart.APIClient) Helpers {
	return Helpers{ovChipkaartAPIClient}
}

// ValidateOvChipkaartCredentials checks that the username and password for an ov chipkaart are valid
func (h Helpers) ValidateOvChipkaartCredentials(username string, password string) (err error) {
	token, err := h.ovChipkaartAPIClient.GetAuthorisationToken(username, password)
	log.Println(token)
	return err
}
