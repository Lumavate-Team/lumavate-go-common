package api_core

import (
    "crypto/md5"
    "crypto/hmac"
    "encoding/hex"
    "crypto/sha256"
    "encoding/base64"
    "net/url"
    "strings"
    "fmt"
    "os"
    "time"
    "math/rand"
		"sort"
)

type Signer struct {
}

func (this *Signer) GetSignature(method string, url_to_sign string, body []byte) string {
  publicKey := os.Getenv("PUBLIC_KEY")
  privateKey := []byte(os.Getenv("PRIVATE_KEY"))
  rand.Seed(time.Now().UTC().UnixNano())

  parts := strings.Split(url_to_sign, "?")
  root := parts[0]
  query := ""
  if len(parts) > 1 {
    query = parts[1]
  }

  rm,  _ := url.Parse(root)
  root = rm.RequestURI()

  qs, _ := url.ParseQuery(query)

  qs.Add("s-key", publicKey)
  qs.Add("s-time", fmt.Sprintf("%d", time.Now().Unix()))
  qs.Add("s-hash", this.GetMD5Hash(body))
  qs.Add("s-nonce", fmt.Sprintf("%d", rand.Intn(1000000000)))
	p, _ := url.PathUnescape(root)

	qs_string :=  orderQueryString(qs)

  key := fmt.Sprintf("%v\n%v\n%v\n%v",
                     method,
                     strings.ToLower(p),
                     qs_string,
                     qs.Get("s-nonce"))

  h := hmac.New(sha256.New, privateKey)
  h.Write([]byte(key))
  signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
  qs.Add("s-signature", url.QueryEscape(signature))
	qs.Del("s-hash")

	qs_string_final := orderQueryString(qs)

	result := fmt.Sprintf("%v?%v", root, qs_string_final)
  return result
}

func orderQueryString(qs map[string][]string) string{
	qs_keys := make([]string, 0)
	for k, _ := range(qs) {
		qs_keys = append(qs_keys, k)
	}
	sort.Strings(qs_keys)
	qs_full := make([]string, 0)
	for _, k := range(qs_keys) {
		for _, v := range(qs[k]){
			qs_full = append(qs_full, fmt.Sprintf("%s=%s", k, v))
		}
	}
	qs_string := strings.Join(qs_full, "&")
	return qs_string
}

func (this *Signer) GetMD5Hash(body []byte) string {
    hasher := md5.New()
    hasher.Write(body)
    return hex.EncodeToString(hasher.Sum(nil))
}
