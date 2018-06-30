package v13

/*

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

   http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

import (
	"github.com/apache/trafficcontrol/lib/go-log"
	"github.com/apache/trafficcontrol/lib/go-tc"
	"testing"
)

func TestDeliveryServiceMatches(t *testing.T) {
	CreateTestCDNs(t)
	CreateTestTypes(t)
	CreateTestProfiles(t)
	CreateTestStatuses(t)
	CreateTestDivisions(t)
	CreateTestRegions(t)
	CreateTestPhysLocations(t)
	CreateTestCacheGroups(t)
	CreateTestServers(t)
	CreateTestDeliveryServices(t)

	GetTestDeliveryServiceMatches(t)

	DeleteTestDeliveryServices(t)
	DeleteTestServers(t)
	DeleteTestCacheGroups(t)
	DeleteTestPhysLocations(t)
	DeleteTestRegions(t)
	DeleteTestDivisions(t)
	DeleteTestStatuses(t)
	DeleteTestProfiles(t)
	DeleteTestTypes(t)
	DeleteTestCDNs(t)
}

func GetTestDeliveryServiceMatches(t *testing.T) {
	log.Debugln("GetTestDeliveryServiceMatches")
	dsMatches, _, err := TOSession.GetDeliveryServiceMatches()
	if err != nil {
		t.Fatalf("cannot GET DeliveryService matches: %v\n", err)
	}

	dsMatchMap := map[tc.DeliveryServiceName][]string{}
	for _, ds := range dsMatches {
		dsMatchMap[ds.DSName] = ds.Patterns
	}

	for _, ds := range testData.DeliveryServices {
		if _, ok := dsMatchMap[tc.DeliveryServiceName(ds.XMLID)]; !ok {
			t.Fatalf("GET DeliveryService matches missing: %v\n", ds.XMLID)
		}
	}
}
