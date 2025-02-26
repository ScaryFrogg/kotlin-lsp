package rpc_test

import (
	"testing"

	"github.comgithub.com/ScaryFrogg/kotlin-lsp/rpc"
)

type TestStruct struct {
	Testing bool
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 16\r\n\r\n{\"Testing\":true}"
	actual := rpc.Encode(TestStruct{Testing: true})
	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}
