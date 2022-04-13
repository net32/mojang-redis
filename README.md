# :incoming_envelope: mojang-redis :incoming_envelope:
A Go MicroService for the Mojang API, to serve responses faster and create REDIS caches for each request.

## Quick start with Docker :whale:

```sh
git clone https://github.com/net32/mojang-redis.git
docker-compose up -d
```
## Examples Request :memo:
Request are the same of [Mojang API](https://wiki.vg/Mojang_API)

### [Username to UUID](https://wiki.vg/Mojang_API#Username_to_UUID)
```http
GET http://127.0.0.1:8080/users/profiles/minecraft/notch
```
### Response
```json
{
	"name": "Notch",
	"id": "069a79f444e94726a5befca90e38aaf5"
}
```
### [Usernames to UUIDs](https://wiki.vg/Mojang_API#Usernames_to_UUIDs)
```http
POST http://127.0.0.1:8080/profiles/minecraft
```
#### Body payload
```json
[
    "NeT32",
    "notch",
    "nonExistingPlayer"
]
```
### Response
```json
[
	{
		"id": "c5870df744e9495f928a0e3e8703a03e",
		"name": "net32"
	},
	{
		"id": "069a79f444e94726a5befca90e38aaf5",
		"name": "Notch"
	}
]
```
### [UUID to Profile and Skin/Cape](https://wiki.vg/Mojang_API#UUID_to_Profile_and_Skin.2FCape)
```http
GET http://127.0.0.1:8080/session/minecraft/profile/069a79f444e94726a5befca90e38aaf5
```
### Response
```json
{
	"id": "069a79f444e94726a5befca90e38aaf5",
	"name": "Notch",
	"properties": [
		{
			"name": "textures",
			"value": "ewogICJ0aW1lc3RhbXAiIDogMTY0OTgxMjMzNzk2NywKICAicHJvZmlsZUlkIiA6ICIwNjlhNzlmNDQ0ZTk0NzI2YTViZWZjYTkwZTM4YWFmNSIsCiAgInByb2ZpbGVOYW1lIiA6ICJOb3RjaCIsCiAgInRleHR1cmVzIiA6IHsKICAgICJTS0lOIiA6IHsKICAgICAgInVybCIgOiAiaHR0cDovL3RleHR1cmVzLm1pbmVjcmFmdC5uZXQvdGV4dHVyZS8yOTIwMDlhNDkyNWI1OGYwMmM3N2RhZGMzZWNlZjA3ZWE0Yzc0NzJmNjRlMGZkYzMyY2U1NTIyNDg5MzYyNjgwIgogICAgfQogIH0KfQ=="
		}
	]
}
```
