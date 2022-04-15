package test

import (
	"testing"

	"github.com/net32/mojang-redis/server"
)

const HTTP_200_OK = 200
const HTTP_204_No_Content = 204
const HTTP_404_Not_Found = 404
const HTTP_405_Method_Not_Allowed = 405

func TestFetchProfileByName(t *testing.T) {
	userName := "net32"
	uuid := "c5870df744e9495f928a0e3e8703a03e"
	t.Log("Testing FetchProfileByName expected", HTTP_200_OK)
	profile, response := server.FetchProfileByName(userName)
	if response.Code != HTTP_200_OK {
		t.Errorf("Expected %d return %d", HTTP_200_OK, response.Code)
	}
	t.Log("Testing Profile Name expected", userName)
	if profile.Name != userName {
		t.Errorf("Expected %s return %s", userName, profile.Name)
	}
	t.Log("Testing UUID expected", uuid)
	if profile.UUID != uuid {
		t.Errorf("Expected %s return %s", uuid, profile.UUID)
	}
	t.Log("Response:", profile, response)
}

func TestUuidToNameHistory(t *testing.T) {
	t.Log("Testing UuidToNameHistory expected", HTTP_200_OK)
	response := server.UuidToNameHistory("c5870df7-44e9-495f-928a-0e3e8703a03e")
	if response.Code != HTTP_200_OK {
		t.Errorf("Expected %d return %d", HTTP_200_OK, response.Code)
	}
	t.Log("Response:", response)
}

func TestBlockedServers(t *testing.T) {
	t.Log("Testing BlockedServers expected", HTTP_200_OK)
	response := server.BlockedServers()
	if response.Code != HTTP_200_OK {
		t.Errorf("Expected %d return %d", HTTP_200_OK, response.Code)
	}
	t.Log("Response:", response.Code, "total", len(response.Json))
}
