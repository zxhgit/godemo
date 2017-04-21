//	ppgo-Curl
//	调用HTTP请求,依赖beego-curl
//  依赖情况:
//			"github.com/astaxie/beego"

package ppgo

import (
//"github.com/astaxie/beego/httplib"
	"ppgo/httplib"
)
type Curl struct {
}

// Get请求
func (this *Curl)CurlGet(url string) (string, error) {

	curl := httplib.Get(url)
	str, err := curl.String()
	if err != nil {
		return "", err
	}
	return str, nil
}



