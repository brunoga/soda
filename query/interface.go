package query

// Interface is the common interface for the SOSQL API query implementations.
type Interface[T any] interface {
	// Execute executes the query and returns the results.
	Execute() (T, error)
}

