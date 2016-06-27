package healthz

import (
	"errors"
	"fmt"
	"net/http"
)

type VaultChecker struct {
	address string
}

func NewVaultChecker(address string) *VaultChecker {
	return &VaultChecker{address}
}

func (vc *VaultChecker) Ping() error {
	url := fmt.Sprintf("http://%s/sys/health", vc.address)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case 200:
		return nil
	case 429:
		errors.New("unsealed and in standby.")
	case 500:
		errors.New("Sealed or not initialized.")
	}
	return nil
}
