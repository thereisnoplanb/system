package Collections

import "errors"

var ErrCollectionContainsNoElements = errors.New("collection contains no elements")
var ErrIndexIsLessThanZero = errors.New("index is less than zero")
var ErrIndexIsGreaterThanCount = errors.New("index is greater than count")
