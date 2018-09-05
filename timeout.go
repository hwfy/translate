package translate

import (
	"net"
	"net/http"
	"time"
)

// getHttpClient 获取超时设置的HTTP服务,默认连接20s,读写25s
func getHttpClient(rw_time, connect_time time.Duration) http.Client {
	if rw_time == 0 {
		rw_time = 20 * time.Second
	}
	if connect_time == 0 {
		connect_time = 10 * time.Second
	}
	return http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(rw_time)                 //读写超时
				c, err := net.DialTimeout(netw, addr, connect_time) //连接超时
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}
}
