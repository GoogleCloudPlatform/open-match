package testing

import "open-match.dev/open-match/internal/future/pb"

// NewPbFilter is a wrapper to create protobuf filter objects
func NewPbFilter(attributeName string, min, max float64) *pb.Filter {
	return &pb.Filter{Attribute: attributeName, Min: min, Max: max}
}

// NewPbPool is a wrapper to create protobuf pool objects
func NewPbPool(poolName string, filters ...*pb.Filter) *pb.Pool {
	return &pb.Pool{
		Name:   poolName,
		Filter: filters,
	}
}
