// Package notify contains helper functions for sending notifications.
//
//nolint:staticcheck
package notify

import (
	"bytes"
	"context"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/candiddev/shared/go/assert"
	"github.com/candiddev/shared/go/logger"
	"github.com/golang-jwt/jwt/v4"
)

func TestWebPushSend(t *testing.T) {
	logger.UseTestLogger(t)

	srvPrv, srvPub, _ := NewWebPushVAPID()
	c := WebPush{
		BaseURL:         "hello",
		VAPIDPrivateKey: srvPrv,
	}
	curve := elliptic.P256()
	cliPrv, x, y, _ := elliptic.GenerateKey(curve, rand.Reader)
	cliPub := elliptic.Marshal(curve, x, y)

	msg := WebPushMessage{
		Actions: WebPushActions{
			Default:    "c",
			Target:     "a",
			TargetType: "b",
			Types: []WebPushActionType{
				1,
			},
		},
		Body:    "Hello world!",
		Subject: "Greetings",
		Topic:   "topic1",
		TTL:     100,
		Urgency: WebPushUrgencyHigh,
	}

	var gotBody []byte

	var gotHeader http.Header

	var resStatus int

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Del("User-Agent")
		r.Header.Del("Accept-Encoding")

		gotHeader = r.Header
		gotBody, _ = io.ReadAll(r.Body)

		w.WriteHeader(resStatus)
	}))

	cli := &WebPushClient{
		Auth:     "lDVz2GD0oypIBSDIEnIe6A",
		Endpoint: srv.URL,
		P256:     base64.RawURLEncoding.EncodeToString(cliPub),
	}

	tests := map[string]struct {
		client    *WebPushClient
		resStatus int
		srvPub    string
		wantErr   error
	}{
		"no server": {
			wantErr: ErrCancelled,
		},
		"no client": {
			srvPub:  srvPub,
			wantErr: ErrCancelled,
		},
		"bad client": {
			client: &WebPushClient{
				Auth:     "#*&&%*@(%*@)!(%_)",
				Endpoint: srv.URL,
				P256:     string(cliPub),
			},
			srvPub:  srvPub,
			wantErr: ErrSend,
		},
		"bad request": {
			client:    cli,
			resStatus: http.StatusBadRequest,
			srvPub:    srvPub,
			wantErr:   ErrSend,
		},
		"forbidden": {
			client:    cli,
			resStatus: http.StatusForbidden,
			srvPub:    srvPub,
			wantErr:   ErrSend,
		},
		"not found": {
			client:    cli,
			resStatus: http.StatusNotFound,
			srvPub:    srvPub,
			wantErr:   ErrMissing,
		},
		"gone": {
			client:    cli,
			resStatus: http.StatusGone,
			srvPub:    srvPub,
			wantErr:   ErrMissing,
		},
		"too large": {
			client:    cli,
			resStatus: http.StatusRequestEntityTooLarge,
			srvPub:    srvPub,
			wantErr:   ErrSend,
		},
		"too many requests": {
			client:    cli,
			resStatus: http.StatusTooManyRequests,
			srvPub:    srvPub,
			wantErr:   ErrSend,
		},
		"good": {
			client: &WebPushClient{
				Auth:     "lDVz2GD0oypIBSDIEnIe6A",
				Endpoint: srv.URL,
				P256:     base64.RawURLEncoding.EncodeToString(cliPub),
			},
			resStatus: http.StatusCreated,
			srvPub:    srvPub,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			c.VAPIDPublicKey = tc.srvPub
			m := msg
			m.Client = tc.client
			resStatus = tc.resStatus

			assert.HasErr(t, c.Send(ctx, &m), tc.wantErr)

			if tc.wantErr == nil {
				assert.Equal(t, gotHeader.Get("Content-Encoding"), "aes128gcm")
				assert.Equal(t, gotHeader.Get("Content-Type"), "application/octet-stream")
				assert.Equal(t, gotHeader.Get("Crypto-Key"), fmt.Sprintf("p256ecdsa=%s", srvPub))
				assert.Equal(t, gotHeader.Get("TTL"), "100")
				assert.Equal(t, gotHeader.Get("Topic"), "topic1")
				assert.Equal(t, gotHeader.Get("Urgency"), "high")

				c := jwt.MapClaims{
					"aud": "",
					"exp": jwt.NewNumericDate(time.Now()),
					"sub": "",
				}

				token, err := jwt.ParseWithClaims(strings.Split(gotHeader.Get("Authorization"), " ")[1], &c, func(t *jwt.Token) (any, error) {
					return &server.key.PublicKey, nil
				})

				assert.HasErr(t, err, nil)
				assert.Equal(t, token.Valid, true)
				assert.Equal(t, len(gotBody), 4096)
				assert.Equal(t, c["aud"].(string), strings.Join(strings.Split(srv.URL, ":")[:2], ":"))
				assert.Equal(t, time.Unix(int64(c["exp"].(float64)), 0).After(time.Now().Add(4*time.Hour)), true)
				assert.Equal(t, c["sub"], "hello")

				salt := gotBody[:16] // Salt is 16 characters
				l := int(gotBody[20])
				srvPub := gotBody[21 : 21+l]
				auth, _ := base64.RawURLEncoding.DecodeString(cli.Auth)

				gcm, nonce, _ := getWebPushCipherNonce(srvPub, cliPrv, cliPub, auth, salt, true)
				ct, err := gcm.Open([]byte{}, nonce, gotBody[21+l:], nil)
				assert.HasErr(t, err, nil)

				j := map[string]any{}

				assert.HasErr(t, json.Unmarshal(bytes.Split(ct, []byte("\x02"))[0], &j), nil)
				assert.Equal(t, j["actions"].(map[string]any)["default"], "c")
				assert.Equal(t, j["body"].(string), "Hello world!")
				assert.Equal(t, j["subject"].(string), "Greetings")
			}
		})
	}
}
