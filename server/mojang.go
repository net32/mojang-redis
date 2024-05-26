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

// This endpoint has been deprecated by Mojang and was removed on 13 September 2022 at 9:25 AM CET to "improve player safety and data privacy"
func UuidToNameHistory(uuid string) MojangResponse {
	return UuidToName(uuid, "names")
}

func UuidToProfile(uuid string, unsigned string) MojangResponse {
	URL := SESSION_URL + "session/minecraft/profile/" + uuid + "?unsigned=" + unsigned
	return mojangGet(URL)
}

func HasJoined(userName string, serverId string) MojangResponse {
	URL := SESSION_URL + fmt.Sprintf("session/minecraft/hasJoined?username=%s&serverId=%s", userName, serverId)
	return mojangGet(URL)
}

func BlockedServers() MojangResponse {
	URL := SESSION_URL + "blockedservers"
	return mojangGet(URL)
}

type MojangResponse struct {
	Code int    `json:"code"`
	Json string `json:"json"`
}

func mojangGet(URL string) MojangResponse {
	key := URL
	cache := HasCache(key)
	if cache.HasCache {
		return cache.Response
	}
	resp, err := http.Get(URL)
	return SaveCache(key, mojangResponse(resp, err)).Response
}

func mojangPost(URL string, jsonData []byte) MojangResponse {
	hashData := hex.EncodeToString(md5.New().Sum(jsonData))
	key := URL + hashData
	cache := HasCache(key)
	if cache.HasCache {
		return cache.Response
	}
	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(jsonData))
	return SaveCache(key, mojangResponse(resp, err)).Response
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
