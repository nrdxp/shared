package notify

import (
	"testing"

	"github.com/candiddev/shared/go/assert"
	"github.com/candiddev/shared/go/logger"
)

func TestValidDomain(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"no match": {
			input: "support@homechart.app",
			want:  true,
		},
		"match demo example": {
			input: "email@demo.example.com",
			want:  false,
		},
		"match": {
			input: "email@sOMETHING.com",
			want:  false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			s := SMTP{
				FromAddress: "no@example.com",
				NoEmailDomains: []string{
					"something.com",
				},
			}
			s.Setup(ctx, "", "", "", "")

			assert.Equal(t, s.ValidDomain(tc.input), tc.want)

			s.NoEmailDomains = []string{}
			s.Setup(ctx, "", "", "", "")

			assert.Equal(t, s.ValidDomain(tc.input), true)
		})
	}
}

// TODO make this work using a mock SMTP server.
func TestSMTPSend(t *testing.T) {
	logger.UseTestLogger(t)

	if c.SMTP.Hostname == "" || c.SMTP.Username == "" {
		t.Skip("No SMTP hostname/username defined")

		return
	}

	c.SMTP.FromAddress = c.SMTP.Username

	tests := map[string]struct {
		errSetup    bool
		errSend     bool
		inputConfig *notifyConfig
		inputTo     string
	}{
		"no hostname/template": {
			errSend:     true,
			errSetup:    true,
			inputConfig: &notifyConfig{},
		},
		"no recipients": {
			errSend:     true,
			inputConfig: c,
		},
		"good": {
			inputConfig: c,
			inputTo:     c.SMTP.Username,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.inputConfig.SMTP.Setup(ctx, "homechart", "https://homechart.app", "https://homechart.app/homechart.png", "/settings/notifications") != nil, tc.errSetup)

			assert.Equal(t, tc.inputConfig.SMTP.Send(ctx, SMTPMessage{
				Body:    "This is a test email.",
				To:      tc.inputTo,
				Subject: "Hello!",
			}) != nil, tc.errSend)
		})
	}
}

func TestSMTPGenerateBody(t *testing.T) {
	logger.UseTestLogger(t)

	c.SMTP.FromAddress = "from@example.com"

	err := c.SMTP.Setup(ctx, "appName", "https://example.com", "https://example.com/logo.png", "/unsubscribe")
	assert.Equal(t, err, nil)

	body, errr := c.SMTP.generateBody(SMTPMessage{
		Body:         "Hello World!",
		FooterFrom:   "This is a generated email from appName",
		FooterUpdate: "Update your settings",
		Subject:      "Hello!",
		To:           "person@example.com",
	})

	b := string(body)

	assert.Equal(t, errr, nil)
	assert.Contains(t, b, "To: person@example.com")
	assert.Contains(t, b, "Reply-To: "+c.SMTP.ReplyTo)
	assert.Contains(t, b, "Subject: Hello!")
	assert.Contains(t, b, "List-Unsubscribe: <https://example.com/unsubscribe")
	assert.Contains(t, b, "Hello World!")
	assert.Contains(t, b, "This is a generated email from appName (https://example.com)")
	assert.Contains(t, b, "Update your settings: https://example.com/unsubscribe")
	assert.Contains(t, b, `<a href="https://example.com"`)
	assert.Contains(t, b, `alt="appName Logo"`)
	assert.Contains(t, b, `src="https://example.com/logo.png"`)
	assert.Contains(t, b, `appName</span>`)
	assert.Contains(t, b, `This is a generated email from appName (<a href="https://example.com" style="color: #616161" target="_blank">https://example.com`)
	assert.Contains(t, b, `<a href="https://example.com/unsubscribe"`)
	assert.Contains(t, b, `>Update your settings</p>`)
}
