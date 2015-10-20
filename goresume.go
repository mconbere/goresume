package goresume

import (
	"io"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"io/ioutil"
	"net/url"

	"github.com/golang/protobuf/proto"
	md "github.com/russross/blackfriday"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"

	resumepb "proto"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	url, data, text, output, err := parseURL(r.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if url != nil {
		client := urlfetch.Client(ctx)
		resp, err := client.Get(url.String())
		defer resp.Body.Close()

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	resume, err := func() (*resumepb.Resume, error) {
		r := &resumepb.Resume{}
		if text == true {
			err := proto.UnmarshalText(string(data), r)
			if err != nil {
				return nil, fmt.Errorf("Could not unmarshal proto text: %v", err)
			}
			return r, nil
		}
		err := proto.Unmarshal(data, r)
		if err != nil {
			return nil, fmt.Errorf("Could not unmarshal proto data: %v", err)
		}
		return r, nil
	}()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := output(w, resume); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func outputBinary(w http.ResponseWriter, r *resumepb.Resume) error {
	w.Header().Add("Content-Type", "application/octet-stream")
	b, err := proto.Marshal(r)
	if err != nil {
		return err
	}
	_, err = w.Write(b)
	return err
}

func outputText(w http.ResponseWriter, r *resumepb.Resume) error {
	w.Header().Add("Content-Type", "text/plain")
	return proto.MarshalText(w, r)
}

func outputMarkdown(w http.ResponseWriter, r *resumepb.Resume) error {
	s := resumeToMarkdown(r)
	_, err := io.WriteString(w, s)
	return err
}

func outputHTML(w http.ResponseWriter, r *resumepb.Resume) error {
	s := resumeToMarkdown(r)
	b := md.MarkdownCommon([]byte(s))
	_, err := w.Write(b)
	return err
}

const (
	queryURL    string = "url"
	queryData   string = "data"
	queryText   string = "text"
	queryOutput string = "output"
)

type outputFunc func(http.ResponseWriter, *resumepb.Resume) error

var outputMap = map[string]outputFunc{
	"binary": outputBinary,
	"text": outputText,
	"markdown": outputMarkdown,
	"html": outputHTML,
}
var outputDefault = outputText

const textDefault bool = false

func parseURL(in *url.URL) (u *url.URL, data []byte, text bool, output outputFunc, err error) {
	values := in.Query()

	u, err = func() (*url.URL, error) {
		v := values.Get(queryURL)
		if v == "" {
			return nil, nil
		}
		return url.Parse(v)
	}()
	if err != nil {
		return
	}

	data, err = func() ([]byte, error) {
		v := values.Get(queryData)
		if v == "" {
			return nil, nil
		}
		return base64.StdEncoding.DecodeString(v)
	}()
	if err != nil {
		return
	}

	text, err = func() (bool, error) {
		v := values.Get(queryText)
		if v == "" {
			return textDefault, nil
		}
		return strconv.ParseBool(v)
	}()
	if err != nil {
		return
	}

	output, err = func() (outputFunc, error) {
		v := values.Get(queryOutput)
		if v == "" {
			return outputDefault, nil
		}
		o, ok := outputMap[v]
		if !ok {
			return nil, fmt.Errorf("unknown output %q", v)
		}
		return o, nil
	}()
	if err != nil {
		return
	}

	if u == nil && data == nil {
		err = fmt.Errorf("Query must contain %q or %q value", queryURL, queryData)
		return
	}

	if u != nil && data != nil {
		err = fmt.Errorf("Query must not contain both %q and %q value", queryURL, queryData)
		return
	}
	return
} 
