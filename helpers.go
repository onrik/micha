package micha

import (
	"encoding/json"
	"net/url"
)

// Convert struct to url values map
// TODO: temp implementation
func structToValues(obj interface{}) (url.Values, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	rawMap := map[string]json.RawMessage{}
	if err := json.Unmarshal(data, &rawMap); err != nil {
		return nil, err
	}

	values := url.Values{}
	for key := range rawMap {
		values.Set(key, string(rawMap[key]))
	}

	return values, nil
}
