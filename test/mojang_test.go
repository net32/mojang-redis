package test

import (
	"testing"

	"github.com/net32/mojang-redis/server"
)

func TestUuidToNameHistory(t *testing.T) {
	t.Log("Testing UuidToNameHistory expected 200")
	response := server.UuidToNameHistory("c5870df7-44e9-495f-928a-0e3e8703a03e")
	if response.Code != 200 {
		t.Error("Expected 200 return code value is", response.Code)
	}
	t.Log("Response:", response)
}

func TestBlockedServers(t *testing.T) {
	t.Log("Testing BlockedServers expected 200")
	response := server.BlockedServers()
	if response.Code != 200 {
		t.Error("Expected 200 return code value is", response.Code)
	}
	t.Log("Response:", response)
}
