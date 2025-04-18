// Copyright (c) 2021-2025 Doc.ai and/or its affiliates.
//
// Copyright (c) 2022-2025 Cisco and/or its affiliates.
//
// Copyright (c) 2024-2025 Pragmagic Inc. and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package interdomain

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bszirtes/integration-tests-1/extensions/parallel"
	"github.com/bszirtes/integration-tests-1/suites/interdomain/suites/basic"
	"github.com/bszirtes/integration-tests-1/suites/interdomain/suites/heal"
	"github.com/bszirtes/integration-tests-1/suites/interdomain/suites/ipsec"
	"github.com/bszirtes/integration-tests-1/suites/interdomain/suites/multiservicemesh"
)

func TestRunBasicInterdomainSuite(t *testing.T) {
	basicSuite := new(basic.Suite)
	parallel.Run(t, basicSuite,
		parallel.WithRunningTestsSynchronously(
			basicSuite.TestFloating_vl3_basic,
			basicSuite.TestFloating_vl3_scale_from_zero,
			basicSuite.TestFloating_vl3_dns,
			basicSuite.TestFloating_nse_composition))
}

func TestRunInterdomainIPSecSuite(t *testing.T) {
	parallel.Run(t, new(ipsec.Suite))
}

func TestRunInterdomainHealSuite(t *testing.T) {
	suite.Run(t, new(heal.Suite))
}

type msmSuite struct {
	multiservicemesh.Suite
}

func (s *msmSuite) BeforeTest(suiteName, testName string) {
	if testName == "TestNsm_istio" {
		s.T().Skip()
	}
}

func TestRunMultiServiceMeshSuite(t *testing.T) {
	suite.Run(t, new(msmSuite))
}
