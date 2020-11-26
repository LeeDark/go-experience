package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/mrjones/oauth"
)


type OAuth1 struct {
	ConsumerKey string
	ConsumerSecret string
	AccessToken string
	AccessSecret string
}

func escape(s string) string {
	t := make([]byte, 0, 3*len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if isEscapable(c) {
			t = append(t, '%')
			t = append(t, "0123456789ABCDEF"[c>>4])
			t = append(t, "0123456789ABCDEF"[c&15])
		} else {
			t = append(t, s[i])
		}
	}
	return string(t)
}

func isEscapable(b byte) bool {
	return !('A' <= b && b <= 'Z' || 'a' <= b && b <= 'z' || '0' <= b && b <= '9' || b == '-' || b == '.' || b == '_' || b == '~')

}

// Params being any key-value url query parameter pairs
func (auth OAuth1) BuildOAuth1Header(method, path string, params map[string]string) string {
	vals := url.Values{}
	vals.Add("oauth_nonce", generateNonce())
	vals.Add("oauth_consumer_key", auth.ConsumerKey)
	vals.Add("oauth_signature_method", "HMAC-SHA1")
	vals.Add("oauth_timestamp", strconv.Itoa(int(time.Now().Unix())))
	//vals.Add("oauth_token", auth.AccessToken)
	vals.Add("oauth_version", "1.0")

	for k, v := range params {
		vals.Add(k, v)
	}
	// net/url package QueryEscape escapes " " into "+", this replaces it with the percentage encoding of " "
	parameterString := strings.Replace(vals.Encode(), "+", "%20", -1)
	fmt.Println("parameterString: ", parameterString)

	// Calculating Signature Base String and Signing Key
	signatureBase := strings.ToUpper(method) + "&" + url.QueryEscape(strings.Split(path, "?")[0]) + "&" + url.QueryEscape(parameterString)
	fmt.Println("signatureBase: ", signatureBase)

	signingKey := url.QueryEscape(auth.ConsumerSecret) + "&" + url.QueryEscape(auth.AccessSecret)
	signature := calculateSignature(signatureBase, signingKey)

	return "OAuth oauth_consumer_key=\"" + url.QueryEscape(vals.Get("oauth_consumer_key")) +
		"\",oauth_nonce=\"" + url.QueryEscape(vals.Get("oauth_nonce")) +
		"\",oauth_signature=\"" + url.QueryEscape(signature) + "\", oauth_signature_method=\"" + url.QueryEscape(vals.Get("oauth_signature_method")) +
		"\",oauth_timestamp=\"" + url.QueryEscape(vals.Get("oauth_timestamp")) +
		"\",oauth_version=\"" + url.QueryEscape(vals.Get("oauth_version")) + "\""
}

func calculateSignature(base, key string) string {
	hash := hmac.New(sha1.New, []byte(key))
	hash.Write([]byte(base))
	signature := hash.Sum(nil)
	return base64.StdEncoding.EncodeToString(signature)
}

func generateNonce() string {
	const allowed= "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 12)
	for i := range b {
		b[i] = allowed[rand.Intn(len(allowed))]
	}
	return string(b)
}

func main() {
	//method := http.MethodPost
	//url := "http://hlrapi.mediafoncs.lt:3333/api/Number/lookup-msisdn-info"
	//payload := strings.NewReader("{\"RequestId\":\"1sdg45df-sdfh56sfg-sdfh545sd-s5s5s4\",\"Msisdns\":[\"37062603925\",\"221770524279\"]}")
	//
	//auth := OAuth1{
	//	ConsumerKey: "WNkKD7c2ke3PFE1PWUDMTDVPJs1t9U43",
	//	ConsumerSecret: "1sEHkM6vgAzCvbaPtHM194EmRGKWTs3eJU4phNsMgJEPC2Jgsyd9b5ZPZhasmPa",
	//	AccessToken: "",
	//	AccessSecret: "",
	//}
	//
	//authHeader := auth.BuildOAuth1Header(method, url, map[string]string {})
	//fmt.Println("Authorization: ", authHeader)
	//
	//req, err := http.NewRequest(method, url, payload)
	//if err != nil {
	//	fmt.Println("Error: ", err)
	//}
	//
	//req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Authorization", authHeader)
	//
	////req.Header.Add("Authorization", "OAuth oauth_consumer_key=\"WNkKD7c2ke3PFE1PWUDMTDVPJs1t9U43\", oauth_nonce=\"BpLnfgDsc2WD\", oauth_signature_method=\"HMAC-SHA1\", oauth_timestamp=\"1604500842\", oauth_version=\"1.0\", oauth_signature=\"ZpsQk%2FxRCThTF1xxIaQuSLzDAlg%3D\"")
	//
	////req.Header.Add("Authorization", "OAuth oauth_consumer_key=\"WNkKD7c2ke3PFE1PWUDMTDVPJs1t9U43\",oauth_signature_method=\"HMAC-SHA1\",oauth_timestamp=\"1604500700\",oauth_nonce=\"K90ns9BtiVI\",oauth_version=\"1.0\",oauth_signature=\"DoKcB5J1aPJ1FeAw1N%2FLrroI7t4%3D\"")
	//
	//if res, err := http.DefaultClient.Do(req); err == nil {
	//	fmt.Println(res.StatusCode)
	//
	//	body, err := ioutil.ReadAll(res.Body)
	//	if err != nil {
	//		fmt.Println("Error: ", err)
	//	}
	//	fmt.Println(string(body))
	//}

	//config := oauth1.NewConfig("WNkKD7c2ke3PFE1PWUDMTDVPJs1t9U43", "1sEHkM6vgAzCvbaPtHM194EmRGKWTs3eJU4phNsMgJEPC2Jgsyd9b5ZPZhasmPa")
	//token := oauth1.NewToken("", "")
	//
	//// httpClient will automatically authorize http.Request's
	//httpClient := config.Client(oauth1.NoContext, token)
	//
	//// example Twitter API request
	//path := "http://hlrapi.mediafoncs.lt:3333/api/Number/lookup-msisdn-info"
	//payload := strings.NewReader("{\"RequestId\":\"1sdg45df-sdfh56sfg-sdfh545sd-s5s5s4\",\"Msisdns\":[\"37062603925\",\"221770524279\"]}")
	//resp, _ := httpClient.Post(path, "application/json", payload)
	//defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Printf("Raw Response Body:\n%v\n", string(body))

	consumerKey := "WNkKD7c2ke3PFE1PWUDMTDVPJs1t9U43"
	consumerSecret := "1sEHkM6vgAzCvbaPtHM194EmRGKWTs3eJU4phNsMgJEPC2Jgsyd9b5ZPZhasmPa"
	c := oauth.NewConsumer(consumerKey, consumerSecret, oauth.ServiceProvider{})
	c.Debug(true)

	token := &oauth.AccessToken{"", "", nil}
	rt, err := c.MakeRoundTripper(token)
	if err != nil {
		fmt.Println("MakeRoundTripper Error: ", err)
	}

	method := http.MethodPost
	url := "http://hlrapi.mediafoncs.lt:3333/api/Number/lookup-msisdn-info"
	payload := strings.NewReader("{\"RequestId\":\"1sdg45df-sdfh56sfg-sdfh545sd-s5s5s4\",\"Msisdns\":[\"37062603925\",\"221770524279\"]}")
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println("NewRequest Error: ", err)
	} else {
		req.Header.Add("Content-Type", "application/json")

		resp, err := rt.RoundTrip(req)
		if err != nil {
			fmt.Println("RoundTrip Error: ", err)
		}

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("Raw Response Body:\n%v\n", string(body))
	}
}
