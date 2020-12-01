package config

import (
	"math/rand"
	"regexp"
	"time"
)

const (
	Aid       = 1233
	UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36"
)

var (
	SecUIDReg  = regexp.MustCompile(`(?m)secUid":"(.*?)"`)
	SeededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
)
