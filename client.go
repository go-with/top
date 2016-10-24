package top

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/bitly/go-simplejson"
	"github.com/google/go-querystring/query"
)

const (
	apiVersion = "2.0"                 // API协议版本
	signMethod = "md5"                 // 签名的摘要算法
	bodyFormat = "json"                // 响应格式
	timeLayout = "2006-01-02 15:04:05" // 时间戳格式

	productionURL = "https://eco.taobao.com/router/rest"       // 正式环境请求地址
	sandboxURL    = "https://gw.api.tbsandbox.com/router/rest" // 沙箱环境请求地址
)

type Request interface {
	Method() string
}

type Client struct {
	AppKey       string // TOP分配给应用的AppKey
	AppSecret    string // TOP分配给应用的AppSecret
	Session      string // TOP颁发给应用的授权信息
	PartnerID    string // 合作伙伴身份标识
	TargetAppKey string // 被调用的目标AppKey
	Simplify     bool   // 是否采用精简JSON返回格式
	Sandbox      bool   // 是否采用沙箱环境
}

func NewClient(appKey, appSecret string) *Client {
	return &Client{
		AppKey:    appKey,
		AppSecret: appSecret,
	}
}

func (c *Client) Do(req Request) (js *simplejson.Json, err error) {
	// 请求参数
	params, err := query.Values(req)
	if err != nil {
		return
	}
	// 公共参数
	params.Set("method", req.Method())
	params.Set("app_key", c.AppKey)
	params.Set("timestamp", time.Now().Format(timeLayout))
	params.Set("format", bodyFormat)
	params.Set("v", apiVersion)
	if len(c.Session) > 0 {
		params.Set("session", c.Session)
	}
	if len(c.PartnerID) > 0 {
		params.Set("partner_id", c.PartnerID)
	}
	if len(c.TargetAppKey) > 0 {
		params.Set("target_app_key", c.TargetAppKey)
	}
	if c.Simplify != false {
		params.Set("simplify", strconv.FormatBool(c.Simplify))
	}
	params.Set("sign_method", signMethod)
	params.Set("sign", c.sign(params))

	var addr string
	if c.Sandbox {
		addr = sandboxURL
	} else {
		addr = productionURL
	}
	resp, err := http.PostForm(addr, params)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	js, err = simplejson.NewJson(body)
	return
}

// 参数签名
func (c *Client) sign(params url.Values) string {
	keys := make([]string, 0, len(params))
	for k, _ := range params {
		if k != "sign" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	var buf bytes.Buffer
	buf.WriteString(c.AppSecret)
	for _, k := range keys {
		buf.WriteString(k)
		buf.WriteString(params.Get(k))
	}
	buf.WriteString(c.AppSecret)

	sum := md5.Sum(buf.Bytes())
	str := hex.EncodeToString(sum[:])

	return strings.ToUpper(str)
}
