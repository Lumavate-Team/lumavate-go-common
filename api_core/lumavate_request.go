package api_core

import (
  "fmt"
  "os"
  "net/http"
  "io/ioutil"
  "strconv"
  "strings"
	"bytes"
  "encoding/json"
  "github.com/Lumavate-Team/lumavate-go-common/models"
)

type LumavateRequest struct {
  Authorization string
  BaseUrl       string
}

func (this *LumavateRequest) GetAuth() string {
  return strings.TrimPrefix(this.Authorization, "Bearer " + this.Authorization)
}

func (this *LumavateRequest) Get(url string, use_single_token ...bool) ([]byte, string) {
  use_token := this.ExtractSingleTokenFlag(use_single_token)
	return this.Request("GET", url, []byte{}, use_token)
}

func (this *LumavateRequest) Post(url string, payload []byte, use_single_token ...bool) ([]byte, string) {
  use_token := this.ExtractSingleTokenFlag(use_single_token)
	return this.Request("POST", url, payload, use_token)
}

func (this *LumavateRequest) Put(url string, payload []byte, use_single_token ...bool) ([]byte, string) {
  use_token := this.ExtractSingleTokenFlag(use_single_token)
	return this.Request("PUT", url, payload, use_token)
}

func (this *LumavateRequest) Patch(url string, payload []byte, use_single_token ...bool) ([]byte, string) {
  use_token := this.ExtractSingleTokenFlag(use_single_token)
	return this.Request("PATCH", url, payload, use_token)
}

func (this *LumavateRequest) Delete(url string, payload []byte, use_single_token ...bool) ([]byte, string){
  use_token := this.ExtractSingleTokenFlag(use_single_token)
  return this.Request("DELETE", url, payload, use_token)
}

func (this *LumavateRequest) Request(method string, url string, payload []byte, use_single_token bool) ([]byte, string) {
  s := Signer{}
  baseUrl := this.BaseUrl
  if baseUrl == "" {
    baseUrl = os.Getenv("BASE_URL")
  } else {
    baseUrl = "https://" + this.BaseUrl
  }
  signed_widget_data_url := fmt.Sprintf("%s%s",
    baseUrl,
    s.GetSignature(strings.ToLower(method), url, payload))

	req, err := http.NewRequest(method, signed_widget_data_url, bytes.NewReader(payload))
  if err != nil {
    fmt.Println(err)
    return []byte{}, "500"
  }
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Authorization", "Bearer " + this.GetAuth())

  if use_single_token {
    token_obj, code := this.GetSingleUseToken()
    if code == 200 {
      req.Header.Add("Experience-Token", token_obj.Payload.Data.Token)
    } else {
      return []byte{}, strconv.Itoa(code)
    }
  }

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
    return body, strconv.Itoa(res.StatusCode)
  }
}

func (this *LumavateRequest) ExtractSingleTokenFlag(single_token []bool) bool{
  if len(single_token) == 1 {
    return single_token[0]
  }
  return false
}

func (this *LumavateRequest) GetSingleUseToken() (models.SingleUseToken, int) {

  t, status := this.Post("/pwa/v1/single-use-token", []byte{})
  if code, _ := strconv.Atoi(status); code != 200 {
    return models.SingleUseToken{}, code
  }

  var token models.SingleUseToken
  if err  := json.Unmarshal([]byte(t), &token); err != nil{
    fmt.Println(err)
    return models.SingleUseToken{}, 500
  }

  return token, 200
}
