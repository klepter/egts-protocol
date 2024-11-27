package egts

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	testEgtsSrServiceInfoBytes = []byte{0x02, 0x00, 0x02}
	testEgtsSrServiceInfo = SrServiceInfo{
		ServiceType:            2,
		ServiceStatement:       0,
		ServiceAttribute:       "0",
		ServiceRoutingPriority: "10",
	}
)

func TestEgtsSrServiceInfo_Encode(t *testing.T) {
	posDataBytes, err := testEgtsSrServiceInfo.Encode()

	if assert.NoError(t, err) {
		assert.Equal(t, posDataBytes, testEgtsSrServiceInfoBytes)
	}
}

func TestEgtsSrServiceInfo_Decode(t *testing.T) {
	serviceInfo := SrServiceInfo{}

	if assert.NoError(t, serviceInfo.Decode(testEgtsSrServiceInfoBytes)) {
		assert.Equal(t, serviceInfo, testEgtsSrServiceInfo)
	}
}
