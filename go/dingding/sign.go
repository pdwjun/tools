package main
// 签名计算代码示例（go）
import (
	"fmt"
	"crypto/hmac"
	"crypto/sha256"
    "encoding/base64"
	"net/url"
	"time"
	"strconv"
)

// https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq 
// 参考文档 https://golang.org/pkg/crypto/hmac/
func main() {
	//timestamp := "1575442836258"
	timestamp := strconv.FormatInt(time.Now().UnixNano() / 1e6, 10)
  // 加签生成的密钥
	secret  := "从钉钉机器人设置页面复制"
	var secret_enc []byte = []byte(secret)
	string_to_sign := timestamp + "\n" + secret
	string_to_sign_enc := []byte(string_to_sign)
	mac := hmac.New(sha256.New, secret_enc)
	mac.Write(string_to_sign_enc)
	hmac_code := mac.Sum(nil)

	sign := base64.StdEncoding.EncodeToString([]byte(hmac_code))

	signUrl := url.QueryEscape(sign)
	fmt.Println(timestamp )
	fmt.Println(signUrl )
}
//curl 'https://oapi.dingtalk.com/robot/send?access_token=xxxxxx&timestamp=1575442836258&sign=ttttttt' -H 'Content-Type: application/json' -d '{"msgtype": "text",   "text": {   "content": "通知通知"  }}'
