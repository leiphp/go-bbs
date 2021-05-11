package libs

import (
	"bbs/configs"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	mathrand "math/rand"
	"net"
	"net/http"
	"strings"
	"time"
)


//获得本机ip
func ReturnCurrentIp() string {
	var ips []string
	addrs, _ := net.InterfaceAddrs()
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String())
		}
	}
	return ips[0]
}


/**
 * ReturnJson function
 * 通用返回信息
 * @return void
 * @author LeixiaoTian
 */
func ReturnJson(code int, msg string, result interface{}) interface{} {
	content := make(map[string]interface{})
	content["code"] = code

	if msg != "" {
		content["msg"] = msg
	} else {
		content["msg"] = configs.MsgCode[code]
	}

	//如果返回结果集为空
	if result == nil {
		result = []int{}
	}
	content["data"] = result
	content["timemap"] = time.Now().Unix()
	return content
}

//返回订单号
func ReturnOrderSn(prefix string) string {
	rand := fmt.Sprintf("%06v", mathrand.New(mathrand.NewSource(time.Now().UnixNano())).Int31n(100000000))
	return prefix + time.Now().Format("20060102150405") + rand
}

//http函数
func HttpRequest(method, url, parrams string) ([]byte, error) {

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println("[验证银行卡失败]" + err.Error())
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "APPCODE 5ce2686147d944e9b404010c7e409f40")

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("[验证银行卡失败]" + err.Error())
	}
	return body, err
}

//http请求函数封装
func HttpDo(method, url, parrams, contentType string) ([]byte, error) {

	client := &http.Client{}

	var (
		req *http.Request
		err error
	)
	//默认json数据发送
	if contentType == "" {
		req, err = http.NewRequest(method, url, bytes.NewBuffer([]byte(parrams)))
	} else {
		req, err = http.NewRequest(method, url, strings.NewReader(parrams))
	}
	if err != nil {
		log.Println(err)
	}

	//设置请求内容类型 默认json
	if contentType == "" {
		req.Header.Set("Content-Type", "application/json")
	} else {
		req.Header.Set("Content-Type", contentType)
	}

	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	return body, nil
}

//struct转json字符串
func StructToJson(data interface{}) string {

	byteStr, _ := json.Marshal(data)
	return string(byteStr)

}