package api

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	signaturePrefix = "sha1="
	signatureLength = 45 // len(SignaturePrefix) + len(hex(sha1))
)

func signBody(secret, body []byte) []byte {
	computed := hmac.New(sha1.New, secret)
	computed.Write(body)
	return []byte(computed.Sum(nil))
}

//https://gist.github.com/rjz/b51dc03061dbcff1c521
func verifySignature(secret []byte, signature string, body []byte) bool {
	if len(signature) != signatureLength || !strings.HasPrefix(signature, signaturePrefix) {
		return false
	}

	actual := make([]byte, 20)
	hex.Decode(actual, []byte(signature[5:]))

	return hmac.Equal(signBody(secret, body), actual)
}

type hookContext struct {
	Signature string
	Event     string
	ID        string
	Payload   []byte
}

func parseHook(secret []byte, req *http.Request) (*hookContext, bool) {
	var hc hookContext

	if hc.Signature = req.Header.Get("x-hub-signature"); len(hc.Signature) == 0 {
		return nil, false
	}

	if hc.Event = req.Header.Get("x-github-event"); len(hc.Event) == 0 {
		return nil, false
	}

	if hc.ID = req.Header.Get("x-github-delivery"); len(hc.ID) == 0 {
		return nil, false
	}

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		return nil, false
	}

	if !verifySignature(secret, hc.Signature, body) {
		return nil, false
	}

	hc.Payload = body

	return &hc, true
}

func (a *Api) postWebhook(w http.ResponseWriter, r *http.Request) {
	b := getBot(r)
	con, ok := parseHook([]byte(b.WebhookSecret), r)
	if !ok {
		a.httpError(w, 400, nil)
		return
	}

	switch con.Event {
	case "ping":
		w.WriteHeader(201)
	case "push":
		payload := make(map[string]interface{})
		err := json.Unmarshal(con.Payload, &payload)
		if err != nil {
			a.httpErrorWithMsg(w, 400, "content type is not application/json", nil)
			break
		}

		ref_, ok := payload["ref"]
		if !ok {
			a.httpError(w, 400, nil)
			break
		}

		ref, ok := ref_.(string)
		if !ok {
			a.httpError(w, 400, nil)
			break
		}

		if !strings.Contains(ref, "master") {
			a.httpErrorWithMsg(w, 400, "branch other than master is not available", nil)
			break
		}

		a.requestBuild(w, b)
	default:
		a.httpError(w, 400, nil)
	}
}
