package api_core

import (
  "fmt"
  "os"
  "net/http"
  "io/ioutil"
  "strconv"
  "strings"
	"bytes"
)

type LumavateRequest struct {
  Authorization string
}

func (this *LumavateRequest) GetAuth() string {
  return strings.TrimPrefix(this.Authorization, "Bearer " + this.Authorization)
}

func (this *LumavateRequest) Get(url string) ([]byte, string) {
	return this.Request("GET", url, []byte{})
}

func (this *LumavateRequest) Post(url string, payload []byte) ([]byte, string) {
	return this.Request("POST", url, payload)
}

func (this *LumavateRequest) Put(url string, payload []byte) ([]byte, string) {
	return this.Request("PUT", url, payload)
}

func (this *LumavateRequest) Request(method string, url string, payload []byte) ([]byte, string) {
  s := Signer{}
  signed_widget_data_url := fmt.Sprintf("%s%s",
    os.Getenv("BASE_URL"),
    s.GetSignature(strings.ToLower(method), url, payload))

	req, err := http.NewRequest(method, signed_widget_data_url, bytes.NewReader(payload))
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Authorization", "Bearer " + this.GetAuth())
  if err != nil {
		fmt.Println(err)
    return []byte{}, "500"
  }

  res, err := http.DefaultClient.Do(req)
  if err != nil {
		fmt.Println(err)
    return []byte{}, "500"
  }

  defer res.Body.Close()
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
		fmt.Println(err)
    return []byte{}, "500"
  }

  if res.StatusCode == 200 {
    return body, "200"
  } else {
    return []byte{}, strconv.Itoa(res.StatusCode)
  }
}
