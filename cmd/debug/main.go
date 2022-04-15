package main

import (
	"fmt"

	"github.com/net32/mojang-redis/server"
)

func main() {
	fmt.Println("debug cmd")

	// NeT32, Bar
	profile, _ := server.FetchProfileByName("Bar")
	fmt.Println(profile.Name, profile.UUID, profile.Properties)

	profile = server.FetchProfile(profile.UUID)
	fmt.Println(profile.Name, profile.UUID, profile.Properties[0].Name)

	nameHistory := server.FetchNames(profile.UUID)
	fmt.Println(len(nameHistory))
	for i, entry := range nameHistory {
		fmt.Println(i, entry.Name, entry.ChangedToAt)
	}

}
