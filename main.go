package main

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

func guid() string {
	var id [18]byte
	now := time.Now().UnixMilli()
	id[0] = byte(now >> 40)
	id[1] = byte(now >> 32)
	id[2] = byte(now >> 24)
	id[3] = byte(now >> 16)
	id[4] = byte(now >> 8)
	id[5] = byte(now)
	if _, err := rand.Read(id[6:]); err != nil {
		return ""
	}
	return hex.EncodeToString(id[:])
}
func main() { guid() }
