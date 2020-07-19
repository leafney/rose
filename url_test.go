package rose

import "testing"

func TestUrlJoin(t *testing.T) {
	t.Log(UrlJoin("http://www.baidu.com","/abc/","/def"))
	// http://www.baidu.com/abc/def
}

func TestUrlJoinWithQuery(t *testing.T) {
	u,_:= UrlJoinWithQuery("http://www.baidu.com","abc/","/def/","search?q=golang")
	t.Log(u.String())
//	http://www.baidu.com/abc/def/search?q=golang
}