package system

// Represents the method that defines a set of criteria and determines whether the specified object meets those criteria.
//
// Parameters:
//	obj T - The object to compare against the criteria defined within the method represented by this delegate.
//
// Returs:
//	bool - true if obj meets the criteria defined within the method represented by this delegate, false otherwise.
type Predicate[T any] func(item T) bool
