package query

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

type base struct {
	from   string
	values url.Values
}

func newBase() *base {
	return &base{
		values: make(url.Values),
	}
}

// From sets the resource the query will be executed against. It must be the URL
// to a valid SOAPI endpoint. The format extension is ignored and it is expected
// that JSON result is supported by the endpoint. In case of multiple From calls,
// the last one is what will be used.
func (q *base) From(from string) *base {
	// Force json output.
	from = strings.TrimSuffix(from, filepath.Ext(from))
	q.from = from + ".json"

	return q
}

// Execute sends the query to the SOAPI endpoint and returns the result. Returns
// a nil error on success and a non-nil error on failure.
func (q *base) Execute() ([]Result, error) {
	u, err := url.Parse(q.from)
	if err != nil {
		return nil, err
	}

	u.RawQuery = q.values.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	// Massage response a bit.
	var results []Result
	switch result.(type) {
	case []interface{}:
		for _, r := range result.([]interface{}) {
			rr := r.(map[string]interface{})
			results = append(results, Result(rr))
		}
	case map[string]interface{}:
		results = []Result{Result(result.(map[string]interface{}))}
	default:
		return nil, fmt.Errorf("unexpected result type: %T", result)
	}

	return results, nil
}
