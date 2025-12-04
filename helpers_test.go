package micha

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

type testStruct struct {
	Name string `json:"name"`
}

func TestStructToValues(t *testing.T) {
	_, err := structToValues([]string{"1"})
	require.NotNil(t, err)

	s := testStruct{
		Name: `"Rock & Roll"`,
	}
	values, err := structToValues(s)
	require.Nil(t, err)
	require.Equal(t, url.Values{
		"name": {`"Rock & Roll"`},
	}, values)
}
