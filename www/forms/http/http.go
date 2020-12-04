package http

import (
	"github.com/golangee/forms-example/www/forms/dom"
	"io/ioutil"
	"net/http"
)

type Result struct {
	*http.Response
	Error error
}

func Get(url string, f func(res *http.Response, err error)) {
	go func() {
		defer dom.GlobalPanicHandler()
		//time.Sleep(2 * time.Second)
		res, err := http.Get(url)
		if err == nil {
			defer res.Body.Close()
		}

		f(res, err)
	}()
}

func GetText(url string, f func(res string, err error)) {
	Get(url, func(res *http.Response, err error) {
		if err != nil {
			f("", err)
			return
		}

		buf, err := ioutil.ReadAll(res.Body)
		if err != nil {
			f("", err)
			return
		}

		f(string(buf), nil)

	})
}
