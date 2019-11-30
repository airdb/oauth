package vo

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

// Generate a random token
func randToken() string {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		log.Println(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}
