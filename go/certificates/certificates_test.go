package certificates

import (
	"crypto/x509"
	"testing"

	"github.com/candiddev/shared/go/assert"
)

func TestGetSelfSigned(t *testing.T) {
	c, err := GetSelfSigned("testing")
	assert.HasErr(t, err, nil)

	cr, err := x509.ParseCertificate(c.Certificate[0])
	assert.HasErr(t, err, nil)

	assert.Equal(t, cr.Subject.CommonName, "testing")
}
