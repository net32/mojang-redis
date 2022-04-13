package server

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
)

const API_URL = "https://api.mojang.com/"
const SESSION_URL = "https://sessionserver.mojang.com/"

func UsernameToUUID(userName string) MojangResponse {
	URL := API_URL + "users/profiles/minecraft/" + userName
	return mojangGet(URL)
}

func UsernamesToUUIDs(jsonData []byte) MojangResponse {
	URL := API_URL + "profiles/minecraft"
	return mojangPost(URL, jsonData)
}

func UuidToName(uuid string, action string) MojangResponse {
	URL := API_URL + "user/profiles/" + uuid + "/" + action
	return mojangGet(URL)
}

func UuidToNameHistory(uuid string) MojangResponse {
	return UuidToName(uuid, "names")
}

func UuidToProfile(uuid string) MojangResponse {
	URL := SESSION_URL + "session/minecraft/profile/" + uuid
	return mojangGet(URL)
}

func BlockedServers() MojangResponse {
	URL := SESSION_URL + "blockedservers"
	return mojangGet(URL)
}

type MojangResponse struct {
	Code int
	Json string
}

func mojangGet(URL string) MojangResponse {
	key := URL
	cache := HasCache(key)
	if cache.hasCache {
		return cache.response
	}
	resp, err := http.Get(URL)
	return SaveCache(key, mojangResponse(resp, err)).response
}

func mojangPost(URL string, jsonData []byte) MojangResponse {
	hashData := hex.EncodeToString(md5.New().Sum(jsonData))
	key := URL + hashData
	cache := HasCache(key)
	if cache.hasCache {
		return cache.response
	}
	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(jsonData))
	return SaveCache(key, mojangResponse(resp, err)).response
}

func mojangResponse(resp *http.Response, err error) MojangResponse {
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return MojangResponse{resp.StatusCode, string(b)}
}
