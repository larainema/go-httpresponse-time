// package httpresponse to implement a application 
// that will get HTTP response times over 5 minutes 
// from your location to https://gitlab.com.

package httpresponse

import (
	"log"
	"net"
	"net/http"
	"time"
)

// get http response times
func GetTime() {
	tp := myTransport()
	client := &http.Client{Transport: tp}

	resp, err := client.Get("https://gitlab.com")
	if err != nil {
		log.Fatalf("get error: %s: %s", err)
	}
	defer resp.Body.Close()

	log.Println("Response Time:", tp.Duration())
	log.Println("Connection Time:", tp.ConnDuration())

}

// run GetTime over 5 minis, and run 10 times 
func CronJob() {
	for i := 0; i < 10; i++ {
		GetTime()
		time.Sleep(5 * 60 * 1000 * time.Millisecond)
    	}

}
// run GetTime over 10 seconds, and run 10 times 
func CronJobShort() {
	for i := 0; i < 10; i++ {
		GetTime()
		time.Sleep(10 * 1000 * time.Millisecond)
    	}

}

// custom the Transport
type customTransport struct {
	rtp       http.RoundTripper
	dialer    *net.Dialer
	connStart time.Time
	connEnd   time.Time
	reqStart  time.Time
	reqEnd    time.Time
}

func myTransport() *customTransport {

	tr := &customTransport{
		dialer: &net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		},
	}
	tr.rtp = &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		Dial:                tr.dial,
		TLSHandshakeTimeout: 10 * time.Second,
	}
	return tr
}

// RoundTrip to get the start/end time of request
func (tr *customTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	tr.reqStart = time.Now()
	resp, err := tr.rtp.RoundTrip(r)
	tr.reqEnd = time.Now()
	return resp, err
}

// dial to get the start/end time of connection
func (tr *customTransport) dial(network, addr string) (net.Conn, error) {
	tr.connStart = time.Now()
	cn, err := tr.dialer.Dial(network, addr)
	tr.connEnd = time.Now()
	return cn, err
}

// count the connect time
func (tr *customTransport) ConnDuration() time.Duration {
	return tr.connEnd.Sub(tr.connStart)
}

// count the response time
func (tr *customTransport) Duration() time.Duration {
	return tr.reqEnd.Sub(tr.reqStart)
}
