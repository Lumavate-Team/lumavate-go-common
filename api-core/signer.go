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

  key := fmt.Sprintf("%v\n%v\n%v\n%v",
                     method,
                     strings.ToLower(p),
                     qs.Encode(),
                     qs.Get("s-nonce"))

  h := hmac.New(sha256.New, privateKey)
  h.Write([]byte(key))
  signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
  qs.Add("s-signature", signature)
	qs.Del("s-hash")

  result := fmt.Sprintf("%v?%v", root, qs.Encode())
  return result
}

func (this *Signer) GetMD5Hash(body []byte) string {
    hasher := md5.New()
    hasher.Write(body)
    return hex.EncodeToString(hasher.Sum(nil))
}
