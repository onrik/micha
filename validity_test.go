package micha

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuildCheckString(t *testing.T) {
	values := url.Values{
		"id":         {"12807202"},
		"first_name": {"John"},
		"username":   {"doe"},
		"photo_url":  {"https://t.me/i/userpic/320/a4f15041-e6a3-4cf4-9363-b0cba2f66720.jpg"},
		"auth_date":  {"1722489598"},
		"hash":       {"a19016198a35deb3469a2af3c5aa1f40aa71941cf14cc20e7be8c6507b147061"},
	}

	require.Equal(t,
		"auth_date=1722489598\nfirst_name=John\nid=12807202\nphoto_url=https://t.me/i/userpic/320/a4f15041-e6a3-4cf4-9363-b0cba2f66720.jpg\nusername=doe",
		buildCheckString(values),
	)
}

func TestValidateHash(t *testing.T) {
	values := url.Values{
		"id":         {"12807202"},
		"first_name": {"John"},
		"username":   {"doe"},
		"photo_url":  {"https://t.me/i/userpic/320/a4f15041-e6a3-4cf4-9363-b0cba2f66720.jpg"},
		"auth_date":  {"1722489598"},
		"hash":       {"aba23cb08c508bd952abb2a9b3c2e9b3d6a2c8726145ccea53bd660d7995b586"},
	}

	// Test valid
	err := validateHash(values, []byte("111"))
	require.Nil(t, err)

	// Test invalid
	err = validateHash(values, []byte("222"))
	require.NotNil(t, err)
}
