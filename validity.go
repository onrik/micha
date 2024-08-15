package micha

import (
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

var (
	ErrInvalidHash = errors.New("invalid hash")
)

// ValidateAuthCallback - https://core.telegram.org/widgets/login#checking-authorization
func ValidateAuthCallback(values url.Values, botToken string) error {
	secret := sha256.Sum256([]byte(botToken))

	return validateHash(values, secret[:])
}

// ValidateWabAppData - https://core.telegram.org/bots/webapps#validating-data-received-via-the-mini-app
func ValidateWabAppData(values url.Values, botToken string) error {
	hm := hmac.New(sha256.New, []byte("WebAppData"))
	_, err := hm.Write([]byte(botToken))
	if err != nil {
		return err
	}

	secret := hm.Sum(nil)

	return validateHash(values, secret)
}

func validateHash(values url.Values, secret []byte) error {
	hm := hmac.New(sha256.New, secret)
	_, err := hm.Write([]byte(buildCheckString(values)))
	if err != nil {
		return err
	}

	if fmt.Sprintf("%x", hm.Sum(nil)) != values.Get("hash") {
		return ErrInvalidHash
	}

	return nil
}

func buildCheckString(values url.Values) string {
	keys := []string{}
	for key := range values {
		if key == "hash" {
			continue
		}
		keys = append(keys, key)
	}

	sort.Strings(keys)
	parts := []string{}
	for _, key := range keys {
		parts = append(parts, fmt.Sprintf("%s=%s", key, values.Get(key)))
	}

	return strings.Join(parts, "\n")
}
