// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cloudhsmv2_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func TestAccCloudHSMV2_serial(t *testing.T) {
	t.Parallel()

	testCases := map[string]map[string]func(t *testing.T){
		"Cluster": {
			acctest.CtBasic: testAccCluster_basic,
			"disappears":    testAccCluster_disappears,
			"tags":          testAccCluster_tags,
		},
		"Hsm": {
			"availabilityZone": testAccHSM_AvailabilityZone,
			acctest.CtBasic:    testAccHSM_basic,
			"disappears":       testAccHSM_disappears,
			"ipAddress":        testAccHSM_IPAddress,
		},
		"DataSource": {
			acctest.CtBasic: testAccDataSourceCluster_basic,
		},
	}

	acctest.RunSerialTests2Levels(t, testCases, 0)
}
