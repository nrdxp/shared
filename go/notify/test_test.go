package notify

import (
	"testing"

	"github.com/candiddev/shared/go/assert"
	"github.com/candiddev/shared/go/logger"
)

func TestTestNotifier(t *testing.T) {
	logger.UseTestLogger(t)

	msg1 := SMTPMessage{
		To: "test1",
	}
	msg2 := SMTPMessage{
		To: "test2",
	}
	msg3 := SMTPMessage{
		To: "test3",
	}

	assert.Equal(t, Test.SendSMTP(ctx, msg1), nil)
	assert.Equal(t, Test.SendSMTP(ctx, msg2), nil)
	assert.Equal(t, Test.SMTPMessages(), []SMTPMessage{
		msg1,
		msg2,
	})
	assert.Equal(t, Test.SendSMTP(ctx, msg3), nil)
	assert.Equal(t, Test.SMTPMessages(), []SMTPMessage{
		msg3,
	})

	msg4 := WebPushMessage{
		Client: &WebPushClient{
			Endpoint: "hello",
		},
	}

	assert.Equal(t, Test.SendWebPush(ctx, msg4), nil)
	assert.Equal(t, Test.WebPushMessages(), []WebPushMessage{
		msg4,
	})
}
