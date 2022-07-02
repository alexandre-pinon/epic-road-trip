package utils

import (
	"fmt"

	"github.com/alexandre-pinon/epic-road-trip/model"
)

func ConstraintStringify(constraint model.Constraints) string {
	var parameters string

	if constraint.MaxPrice != 0 {
		maxprice := fmt.Sprintf("&maxprice=%d", constraint.MaxPrice)
		parameters = parameters + maxprice
	}
	if constraint.MinPrice != 0 {
		minprice := fmt.Sprintf("&minprice=%d", constraint.MinPrice)
		parameters = parameters + minprice
	}
	if constraint.Radius != 0 {
		radius := fmt.Sprintf("&radius=%d", constraint.Radius)
		parameters = parameters + radius
	}
	if constraint.OpenNow {
		open := "&opennow"
		parameters = parameters + open
	}

	if parameters == "" {
		parameters = ""
		return parameters
	}

	return parameters
}