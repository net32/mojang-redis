package server

import (
	"encoding/json"
	"log"

	"github.com/net32/mojang-redis/model"
)

func FetchNames(uuid string) []model.NameHistoryEntry {
	data := UuidToNameHistory(uuid)
	var nameHistory []model.NameHistoryEntry
	err := json.Unmarshal([]byte(data.Json), &nameHistory)
	if err != nil {
		log.Println(err, "uuid:", uuid)
	}
	return nameHistory
}

func FetchProfile(uuid string) model.Profile {
	data := UuidToProfile(uuid, "false")
	var profile model.Profile
	err := json.Unmarshal([]byte(data.Json), &profile)
	if err != nil {
		log.Println(err, "uuid:", uuid)
	}
	return profile
}

func FetchProfileByName(userName string) (model.Profile, MojangResponse) {
	data := UsernameToUUID(userName)
	var profile model.Profile
	err := json.Unmarshal([]byte(data.Json), &profile)
	if err != nil {
		log.Println(err, "userName:", userName)
	}
	return profile, data
}

func HasPaid(userName string) (string, MojangResponse) {
	result := "failed"
	data := UsernameToUUID(userName)
	if data.Code == 200 {
		result = "true"
	}
	if data.Code == 404 {
		result = "false"
	}
	return result, data
}
