//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package exported

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestNopCloser(t *testing.T) {
	nc := NopCloser(strings.NewReader("foo"))
	if err := nc.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestHasStatusCode(t *testing.T) {
	if HasStatusCode(nil, http.StatusAccepted) {
		t.Fatal("unexpected success")
	}
	if HasStatusCode(&http.Response{}) {
		t.Fatal("unexpected success")
	}
	if HasStatusCode(&http.Response{StatusCode: http.StatusBadGateway}, http.StatusBadRequest) {
		t.Fatal("unexpected success")
	}
	if !HasStatusCode(&http.Response{StatusCode: http.StatusOK}, http.StatusAccepted, http.StatusOK, http.StatusNoContent) {
		t.Fatal("unexpected failure")
	}
}

func TestPayload(t *testing.T) {
	const val = "payload"
	resp := &http.Response{
		Body: io.NopCloser(strings.NewReader(val)),
	}
	b, err := Payload(resp)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != val {
		t.Fatalf("got %s, want %s", string(b), val)
	}
	b, err = Payload(resp)
	if err != nil {
		t.Fatal(err)
	}
	if string(b) != val {
		t.Fatalf("got %s, want %s", string(b), val)
	}
}
