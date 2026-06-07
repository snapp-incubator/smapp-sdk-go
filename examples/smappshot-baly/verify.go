package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

// serverComputeSignature mirrors signed_url.go computeSignature exactly
func serverComputeSignature(secret, path string, q url.Values) string {
	keys := make([]string, 0, len(q))
	for k := range q {
		if strings.ToLower(k) == "sig" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return strings.ToLower(keys[i]) < strings.ToLower(keys[j])
	})

	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		parts = append(parts, url.QueryEscape(k)+"="+url.QueryEscape(q.Get(k)))
	}
	serialized := strings.Join(parts, "&")
	message := "GET\n" + path + "\n" + serialized
	fmt.Println("Server message:", message)

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}

func init() {
	// Parse the generated URL and verify server-side
	rawURL := "http://localhost:8080/api/v1/photo/ride?destinations=44.4%2C33.36&expires=1780756467&height=285&language=en&marker_type=0&origin=44.3661%2C33.3152&width=512&sig=e196ef3f0dbbecfde02aa17e740839c8201e0d4bbf52c9c4a80f622c1e1d2fbc"
	secret := "5fe9e0d860b7aa3409f41cca3c5f41359124988fb28cbcf58005eb0fc0f21a25"

	u, _ := url.Parse(rawURL)
	q := u.Query()

	serverSig := serverComputeSignature(secret, u.Path, q)
	sdkSig := q.Get("sig")

	fmt.Println("SDK sig:   ", sdkSig)
	fmt.Println("Server sig:", serverSig)
	fmt.Println("Match:", serverSig == sdkSig)
}
