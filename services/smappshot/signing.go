package smappshot

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// signURL appends expires, computes HMAC-SHA256 sig, returns full signed URL.
func signURL(baseURL, path, secret string, expiry time.Duration, params url.Values) (string, error) {
	p := make(url.Values, len(params)+1)
	for k, v := range params {
		p[k] = v
	}
	p.Set("expires", strconv.FormatInt(time.Now().UTC().Add(expiry).Unix(), 10))

	sig, err := computeSignature(secret, path, p)
	if err != nil {
		return "", fmt.Errorf("smapp smappshot: signing failed: %w", err)
	}
	return baseURL + path + "?" + encodeParamsSorted(p) + "&sig=" + sig, nil
}

// computeSignature returns hex(HMAC-SHA256(secret, "GET\n{path}\n{sorted-params}")).
func computeSignature(secret, path string, params url.Values) (string, error) {
	message := "GET\n" + path + "\n" + encodeParamsSorted(params)
	mac := hmac.New(sha256.New, []byte(secret))
	if _, err := mac.Write([]byte(message)); err != nil {
		return "", err
	}
	return hex.EncodeToString(mac.Sum(nil)), nil
}

// encodeParamsSorted sorts keys case-insensitively and percent-encodes key=value pairs.
// sig must be excluded before calling.
func encodeParamsSorted(params url.Values) string {
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return strings.ToLower(keys[i]) < strings.ToLower(keys[j])
	})
	parts := make([]string, 0, len(keys))
	for _, k := range keys {
		for _, v := range params[k] {
			parts = append(parts, url.QueryEscape(k)+"="+url.QueryEscape(v))
		}
	}
	return strings.Join(parts, "&")
}

// formatLocation serializes a Location as "{lon},{lat}".
func formatLocation(loc Location) string {
	return strconv.FormatFloat(loc.Lon, 'f', -1, 64) + "," + strconv.FormatFloat(loc.Lat, 'f', -1, 64)
}

func validateLanguage(lang Language) error {
	switch lang {
	case LanguageFarsi, LanguageEnglish, LanguageArabic, LanguageKurdish:
		return nil
	}
	return fmt.Errorf("smapp smappshot: language must be one of fa, en, ar, ku, got %q", string(lang))
}

func defaultInt(val, fallback int) int {
	if val == 0 {
		return fallback
	}
	return val
}

func defaultLanguage(lang Language) Language {
	if lang == "" {
		return LanguageFarsi
	}
	return lang
}
