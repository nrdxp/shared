package notify

import (
	"testing"

	"github.com/candiddev/shared/go/assert"
)

var testClient = WebPushClient{
	Auth:     "lDVz2GD0oypIBSDIEnIe6A",
	Endpoint: "https://fcm.googleapis.com/fcm/send/dRRK3dOj3KM:APA91bHSisASrz9pTX30JmHlNkR2taunDWVcfCbQsY0c84439Jl9_sAs29t_hCkjk0S9wmb5Rka_cyWGkWnK9lwFFHp3MTtemUft4oxMiYkx1OZ-fuvJV9_MbwiF2sMIlqTlogisdT73",
	P256:     "BIG5gktqUlXxvHNnGnz8aGxcl1I_PBLVuStThK4APa_JNamU7yChBUkgA1OVrt8lcniCrgiQ2ddQ_gZUlYkpzcw",
}

func TestClientDe(t *testing.T) {
	a, p, err := testClient.decode()
	assert.HasErr(t, err, nil)
	assert.Equal(t, len(a), 16)
	assert.Equal(t, len(p), 65)
}
