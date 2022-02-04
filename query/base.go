package query

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

type base[T any] struct {
	from   string
	values url.Values
}

func newBase[T any]() *base[T] {
	return &base[T]{
		values: make(url.Values),
	}
}

// From sets the resource the query will be executed against. It must be the URL
// to a valid SOAPI endpoint. The format extension is ignored and it is expected
// that JSON result is supported by the endpoint. In case of multiple From calls,
// the last one is what will be used.
func (q *base[T]) From(from string) *base[T] {
	// Force json output.
	from = strings.TrimSuffix(from, filepath.Ext(from))
	q.from = from + ".json"

	return q
}

// Execute sends the query to the SOAPI endpoint and returns the result. Returns
// a nil error on success and a non-nil error on failure.
func (q *base[T]) Execute() (T, error) {
	var result T
	u, err := url.Parse(q.from)
	if err != nil {
		return result, err
	}

	u.RawQuery = q.values.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

