package query

// Interface is the common interface for the SOSQL API query implementations.
type Interface interface {
	// Execute executes the query and returns the results.
	Execute() (Result, error)
}
