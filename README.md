# Publisher
Service to publish posts from system.

# Getting Started

First step is to create a `.env` file in project root and change example values to your config. You can use `example.env` file from `_setup` folder as template.

And then, run `docker compose up -d` into project root to start project.

## Tests

The easy way to run tests is just run `docker compose up -d` command to start project with variables. So, enter in `synk_auth` with `docker exec` and run `go test ./tests -v`.

## Certificates

This app must run in HTTPS to authentication works properly. So, to install it, just setup `[mkcert](https://github.com/FiloSottile/mkcert)` into your machine and then run command below into root directory of this project.

```
mkcert -key-file ./.cert/key.pem -cert-file ./.cert/cert.pem localhost synk_publisher
```

## Network

You can use a custom network for this services, using then `synk_network` you must create before run it. So, to create on just run command below once during initial setup.

```
docker network create synk_network
```

# Routes

## Get info about app

> `GET` /about

### Response

```json
{
	"ok": true,
	"error": "",
	"info": {
		"server_port": "8080",
		"app_port": "8083",
		"db_working": true
	},
	"list": null
}
```

## Send message through Discord

> `GET` /discord/publish

### Request

> `webhook_url` can be get in the text chat settings on a server from Discord. So in that settings access Integrations > Webhooks > New Webhook.

```json
{
	"message": "showwwwwwwwwwwwww",
	"webhook_url": "https://discord.com/api/webhooks/123456789/asadsadfasdwefef323112ewefdwed"
}
```

### Response

```json
{
	"resource": {
		"ok": true,
		"error": ""
	},
	"post": {
		"id": "123456789",
		"channel_id": "123456789",
		"webhook_id": "123456789"
	},
	"raw": "{\"type\":0,\"content\":\"showwwwwwwwwwwwww\",\"mentions\":[],\"mention_roles\":[],\"attachments\":[],\"embeds\":[],\"timestamp\":\"2025-11-26T00:23:48.134000+00:00\",\"edited_timestamp\":null,\"flags\":0,\"components\":[],\"id\":\"123456789\",\"channel_id\":\"123456789\",\"author\":{\"id\":\"123456789\",\"username\":\"Captain Hook\",\"avatar\":null,\"discriminator\":\"0000\",\"public_flags\":0,\"flags\":0,\"bot\":true,\"global_name\":null,\"clan\":null,\"primary_guild\":null},\"pinned\":false,\"mention_everyone\":false,\"tts\":false,\"webhook_id\":\"123456789\"}\n"
}
```