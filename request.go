/**
 * @Author:      leafney
 * @Date:        2023-03-04 17:19
 * @Project:     rose
 * @HomePage:    https://github.com/leafney
 * @Description:
 */

package rose

// ReqUserAgentPC PC端 User-Agent
func ReqUserAgentPC() string {
	UserAgentList := []string{
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.54",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534.57.2 (KHTML, like Gecko) Version/5.1.7 Safari/534.57.2",
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/534.16 (KHTML, like Gecko) Chrome/10.0.648.133 Safari/534.16",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_0) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.1 (KHTML, like Gecko) Chrome/21.0.1180.71 Safari/537.1 LBBROWSER",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:90.0) Gecko/20100101 Firefox/90.0",
		"Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_8; en-us) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36",
	}

	return SliceRandomItemStr(UserAgentList)
}

// ReqUserAgentPhone 移动端 User-Agent
func ReqUserAgentPhone() string {
	UserAgentList := []string{
		"Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_3_3 like Mac OS X; en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
		"Mozilla/5.0 (iPad; U; CPU OS 4_3_3 like Mac OS X; en-us) AppleWebKit/533.17.9 (KHTML, like Gecko) Version/5.0.2 Mobile/8J2 Safari/6533.18.5",
		"Mozilla/5.0 (Linux; U; Android 2.3.7; en-us; Nexus One Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
		"Mozilla/5.0 (Linux; U; Android 11; zh-cn; Redmi K30 Pro Build/RKQ1.200826.002) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/79.0.3945.147 Mobile Safari/537.36 XiaoMi/MiuiBrowser/14.7.10",
		"Mozilla/5.0 (Linux; Android 10; EVR-AL00 Build/HUAWEIEVR-AL00; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/74.0.3729.186 Mobile Safari/537.36 baiduboxapp/11.0.5.12 (Baidu; P1 10)",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 16_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 16_6_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/20G81",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 14_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148",
	}
	return SliceRandomItemStr(UserAgentList)
}

const (
	ReqHeaderReferer        = "Referer"
	ReqHeaderUserAgent      = "User-Agent"
	ReqHeaderAuthorization  = "Authorization"
	ReqHeaderContentType    = "Content-Type"
	ReqHeaderAccept         = "Accept"
	ReqHeaderAcceptEncoding = "Accept-Encoding"
	ReqHeaderAcceptLanguage = "Accept-Language"
	ReqHeaderCacheControl   = "Cache-Control"
	ReqHeaderCookie         = "Cookie"
)

//// TODO 请求头参数
//HEADER_ACCEPT_ALL            = "*/*"
//HEADER_ACCEPT_JSON_PLAIN_ALL = "application/json, text/plain, */*"
//HEADER_ACCEPT_LANGUAGE       = "zh-CN,zh-Hans;q=0.9"
//HEADER_ACCEPT_ENCODING       = "gzip"
//HEADER_CONTENT_TYPE_JSON     = "application/json; charset=UTF-8"
//HEADER_CONTENT_TYPE_FORM     = "application/x-www-form-urlencoded"
//
