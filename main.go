package main

import (
	"log"
	"naoborot/capsula"
	"os"
	"strings"
)

func defaultHandler(resp *capsula.Response, req *capsula.Request) *capsula.Response {
	t := strings.ToLower(req.OriginalUtterance())
	rev := []rune(t)
	for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}
	reversed := string(rev)
	resp.Text(reversed)
	resp.TTS(reversed)
	return resp
}

func helloHandler(resp *capsula.Response, req *capsula.Request) *capsula.Response {
	message := "Я буду говорить задом наоборот все что Вы мне скажете. Скажите \"стоп\" \"хватит\" или \"пока\" чтобы закончить"
	resp.Text(message)
	resp.TTS(message)
	return resp
}

func endHandler(resp *capsula.Response, req *capsula.Request) *capsula.Response {
	message := "Пока пока..."
	resp.Text(message)
	resp.TTS(message)
	resp.EndSession()
	return resp
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	dr := capsula.NewDiaogRouter(true)
	dr.RegisterDefault(defaultHandler)
	dr.Register("привет", helloHandler)
	dr.Register("пока", endHandler)

	config := capsula.NewConfig(
		true,
		getEnv("SSL_CERT_FILE", "cert.cer"),
		getEnv("SSL_KEY_FILE", "key.key"),
		getEnv("LISTEN_ADDRESS", ":9000"),
		getEnv("WEBHOOK_URL", "/webhook"),
	)
	skill := capsula.NewSkill(config, dr)
	if err := skill.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
