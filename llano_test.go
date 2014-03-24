package llano

import (
	"bufio"
	"bytes"
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"testing"
)

var (
	errNoRedirect       = errors.New("Not redirecting this client")
	errNoRedirectActual = errors.New("Get /200: Not redirecting this client")
)

type response struct {
	code int
	err  error
}

var responses map[string]response = map[string]response{
	"/200": response{200, nil},
	"/301": response{301, errNoRedirectActual},
	"/302": response{302, errNoRedirectActual},
	"/404": response{404, nil},
	"/500": response{500, nil},
}

func Test_BasicResponse(t *testing.T) {
	testAddress := "127.0.0.1:2020"
	http.DefaultClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {

		return errNoRedirect
	}
	go Standalone(testAddress, "OK")
	Convey("Basic responses should be in order", t, func() {
		for k, v := range responses {
			resp, err := http.Get("http://" + testAddress + k)
			if v.err == nil {
				So(err, ShouldBeNil)
			}
			if v.err != nil {
				So(err.Error(), ShouldEqual, v.err.Error())
			}
			So(resp.StatusCode, ShouldEqual, v.code)
		}
	})
	Convey("Echo response should... echo", t, func() {
		messageList := [...]string{"A", "B", "C"}
		for _, m := range messageList {
			r := bytes.NewReader([]byte(m))
			resp, err := http.Post("http://"+testAddress+"/echo", "text", r)
			So(err, ShouldBeNil)
			responseString, _ := bufio.NewReader(resp.Body).ReadString('\n')
			So(responseString, ShouldEqual, m)
		}
	})
}
