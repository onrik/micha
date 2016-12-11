package micha

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct{}

func (s testStruct) MarshalJSON() ([]byte, error) {
	return nil, errors.New("Some error")
}

func TestStructToValues(t *testing.T) {
	_, err := structToValues(testStruct{})
	assert.NotNil(t, err)

	_, err = structToValues([]string{"1"})
	assert.NotNil(t, err)
}
