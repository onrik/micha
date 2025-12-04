package micha

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

type testStruct struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func TestStructToValues(t *testing.T) {
	_, err := structToValues([]string{"1"})
	require.NotNil(t, err)

	s := testStruct{
		ID:   3904834,
		Name: `"Rock & Roll"`,
	}
	values, err := structToValues(s)
	require.Nil(t, err)
	require.Equal(t, url.Values{
		"id":   {"3904834"},
		"name": {`"Rock & Roll"`},
	}, values)
}
