package slack

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

//nolint:lll
func TestVerify(t *testing.T) {
	t.Skip("DELETE THIS")
	slackSigningToken := []byte("8f742231b10e8888abcd99yyyzzz85a5")
	body := []byte("token=xyzz0WbapA4vBCDEFasx0q6G&team_id=T1DC2JH3J&team_domain=testteamnow&channel_id=G8PSS9T3V&channel_name=foobar&user_id=U2CERLKJA&user_name=roadrunner&command=%2Fwebhook-collect&text=&response_url=https%3A%2F%2Fhooks.slack.com%2Fcommands%2FT1DC2JH3J%2F397700885554%2F96rGlfmibIGlgcZRskXaIFfN&trigger_id=398738663015.47445629121.803a0bc887a14d10d2c447fce8b6703c")
	var timestamp int64 = 1531420618

	expectedHex := []byte("v0=a2114d57b48eac39b9ad189dd8316235a7b4a8d21a10bd27519666489c69b503")

	if ok := Verify(slackSigningToken, body, timestamp, expectedHex); !ok {
		t.Fatal("Failed to verify")
	}
}

//nolint:lll
func TestVerifyRequest(t *testing.T) {
	t.Skip("DELETE THIS")

	bodyRaw := "token=xyzz0WbapA4vBCDEFasx0q6G&team_id=T1DC2JH3J&team_domain=testteamnow&channel_id=G8PSS9T3V&channel_name=foobar&user_id=U2CERLKJA&user_name=roadrunner&command=%2Fwebhook-collect&text=&response_url=https%3A%2F%2Fhooks.slack.com%2Fcommands%2FT1DC2JH3J%2F397700885554%2F96rGlfmibIGlgcZRskXaIFfN&trigger_id=398738663015.47445629121.803a0bc887a14d10d2c447fce8b6703c"
	body := bytes.NewBufferString(bodyRaw)
	req, _ := http.NewRequest("POST", "/api/doesn'tmatter", body)
	req.Header.Add("X-Slack-Request-Timestamp", "1531420618")
	req.Header.Add("X-Slack-Signature", "v0=a2114d57b48eac39b9ad189dd8316235a7b4a8d21a10bd27519666489c69b503")
	slackSigningToken := []byte("8f742231b10e8888abcd99yyyzzz85a5")

	if !VerifyRequest(req, slackSigningToken) {
		t.Fatal("Failed to verify with the request object")
	}

	if b, err := ioutil.ReadAll(req.Body); err != nil || string(b) != bodyRaw {
		t.Fatal("It should not consume body")
	}
}
