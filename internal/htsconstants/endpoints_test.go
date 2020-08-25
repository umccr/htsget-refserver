// Package htsconstants contains program constants
//
// Module endpoints_test tests module endpoints
package htsconstants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var endpointsTC = []struct {
	e   ServerEndpoint
	exp string
}{
	{ReadsTicket, "/reads/{id}"},
	{ReadsData, "/reads/data/{id}"},
	{VariantsServiceInfo, "/variants/service-info"},
	{FileBytes, "/file-bytes"},
}

func TestEndpoints(t *testing.T) {
	for _, tc := range endpointsTC {
		assert.Equal(t, tc.exp, tc.e.String())
	}
}
