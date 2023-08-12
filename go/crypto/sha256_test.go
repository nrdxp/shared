package crypto

import (
	"os"
	"testing"

	"github.com/candiddev/shared/go/assert"
)

func TestSHA256File(t *testing.T) {
	f1, _ := os.OpenFile("testsha1", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	f1.WriteString("Hello")
	f1.Close()
	f1, _ = os.Open("testsha1")

	f2, _ := os.OpenFile("testsha2", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	f2.WriteString("Hello")
	f2.Close()
	f2, _ = os.Open("testsha2")

	f3, _ := os.OpenFile("testsha3", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	f3.WriteString("Hello!")
	f3.Close()
	f3, _ = os.Open("testsha3")

	f4, _ := os.Open("testsha4")

	s1, _ := SHA256File(f1)
	s2, _ := SHA256File(f2)
	s3, _ := SHA256File(f3)
	_, err := SHA256File(f4)

	assert.Equal(t, err != nil, true)
	assert.Equal(t, len(s1), 64)
	assert.Equal(t, s1, s2)
	assert.Equal(t, s1 != s3, true)

	os.Remove("testsha1")
	os.Remove("testsha2")
	os.Remove("testsha3")
}

func TestSHA256String(t *testing.T) {
	s1 := SHA256String("hello")
	s2 := SHA256String("hello")
	s3 := SHA256String("hello1")

	assert.Equal(t, s1, s2)
	assert.Equal(t, s1 != s3, true)
	assert.Equal(t, len(s1), 64)
}
