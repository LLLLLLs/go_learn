// Time        : 2019/10/22
// Description :

package leaderboards

import "testing"

func TestLeaderBoard(t *testing.T) {
	event := DefaultEvent{
		Type:     1,
		Identify: "role",
		Value:    100,
		Extend:   nil,
	}
	subscribe(event)

	ce := CustomerEvent{
		Type:     2,
		Identify: "alliance",
		Group:    "role",
		Group2:   0,
		Value:    123,
		Extend: map[string]interface{}{
			"allianceName": "arthur",
		},
	}
	subscribe(ce)
}
