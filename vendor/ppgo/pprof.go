//	ppgo-pprof
//	快捷的pprof性能分析
//  依赖情况:
//          "net/http/pprof"

package ppgo

import (
	_ "net/http/pprof"
	"log"
	"net/http"
)

// 开启PProf性能分析
func OpenPProf(port string) {
	go func() {
		log.Println(http.ListenAndServe("localhost:" + port, nil))
	}()
}