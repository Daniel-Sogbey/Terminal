package main

import (
	"crypto/sha1"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"time"
)

var urlMap = make(map[string]string)

const base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func main() {

	flag.String("url", "", "Long url")
	flag.Int("id", 0, "Customer Id")
	flag.Parse()

	url := os.Args[2]

	hash := generateShortUrl(url)

	urlMap[url] = hash

	fmt.Printf("UrlMap %s", urlMap)

}

func generateShortUrl(longUrl string) string {
	hasher := sha1.New()
	hasher.Write([]byte(longUrl))
	hashBytes := hasher.Sum(nil)

	shortUrl := base64.URLEncoding.EncodeToString(hashBytes)

	shortUrl = shortUrl[:5]
	return shortUrl
}

func hash(url string, time time.Time, customerId int) string {

	hashedData := fmt.Sprintf("%s%v%d", string(url[len(url)-(len(url)-12)]), time.Second(), (customerId + int(url[len(url)-(len(url)-3)])))

	fmt.Printf("ExtractedData : %v", hashedData)

	return hashedData
}
