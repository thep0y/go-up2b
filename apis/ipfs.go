/*
 * @Author: thepoy
 * @Email: email@example.com
 * @File Name: ipfs.go
 * @Created: 2021-07-27 14:09:31
 * @Modified: 2021-07-27 15:11:18
 */

package apis

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"sync"

	"github.com/thep0y/go-logger/log"
	"github.com/thep0y/go-up2b/models"
	"github.com/thep0y/go-up2b/request"
	"github.com/tidwall/gjson"
)

type Ipfs struct {
	headers    http.Header
	request    *request.Request
	ConfigFile string
	maxSize    int
}

func (i Ipfs) Login(models.LoginInfo) (*models.LoginInfo, error) { return nil, nil }

func (i Ipfs) String() string {
	return "ipfs"
}

func (i Ipfs) UploadImage(imagePath string) (string, error) {
	resp, err := i.postMultipart(imagePath)
	if err != nil {
		return "", err
	}
	body := gjson.ParseBytes(resp.Body)
	if resp.StatusCode == 200 {
		hash := body.Get("Hash").String()
		return i.url(hash), nil
	}
	return "", fmt.Errorf("wrong response: %d", resp.StatusCode)
}

func (i Ipfs) UploadImages(imagesPath []string) ([]string, error) {
	var ch = make(chan uploadResult, len(imagesPath))

	var wg sync.WaitGroup

	for index, path := range imagesPath {
		wg.Add(1)

		go func(index int, p string) {
			defer wg.Done()
			u, err := i.UploadImage(p)
			if err == nil {
				ch <- uploadResult{index, u}
			} else {
				log.Error(err)
			}
		}(index, path)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var downloadURL = make([]string, len(imagesPath))

	for r := range ch {
		downloadURL[r.index] = r.url
	}
	return downloadURL, nil
}

func (i Ipfs) BaseURL() string {
	return "https://ipfsapi.glitch.me//api/v0/add?pin=true"
}

func (i *Ipfs) MakeHeaders(headers map[string]string) error {
	if headers == nil {
		headers = map[string]string{
			"Accept":     "application/json, text/javascript, */*; q=0.01",
			"User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:85.0) Gecko/20100101 Firefox/85.0",
		}
	}
	i.headers = make(http.Header)
	for k, v := range headers {
		i.headers.Add(k, v)
	}
	return nil
}

func (i Ipfs) postMultipart(imgPath string) (*request.Response, error) {
	// boundary := i.boundary()
	// i.headers.Add("Content-Type", "boundary=----WebKitFormBoundary"+boundary)
	contentType, reader, err := i.createBody("", imgPath)
	if err != nil {
		return nil, err
	}
	resp, err := i.request.PostMultipartForm(i.BaseURL(), contentType, reader)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (i Ipfs) createBody(boundary, filepath string) (string, io.Reader, error) {
	basename := path.Base(filepath)
	// dashBoundary := "------" + boundary

	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)

	f, err := os.Open(filepath)
	if err != nil {
		return "", nil, err
	}
	defer f.Close()

	fw, _ := w.CreateFormFile("file", basename)
	io.Copy(fw, f)

	w.Close()

	// buffer.WriteString(dashBoundary + "\r\n")
	// buffer.WriteString("Content-Disposition: form-data; name=file; filename=" + basename + "\r\n")
	// buffer.WriteString("Content-Type: image/png\r\n")
	// buffer.Write(content)
	// buffer.WriteString("\r\n")

	// buffer.WriteString(dashBoundary + "--\r\n")
	return w.FormDataContentType(), buf, nil
}

func (i Ipfs) boundary() string {
	const letterBytes = "123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, 16)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (i Ipfs) GetHeaders() http.Header { return i.headers }

func (i *Ipfs) NewRequest() {
	i.request = request.NewRequest(i.headers)
}

func (i Ipfs) checkSize(filepath string) error { return nil }

func (i Ipfs) url(hash string) string {
	return "https://cf-ipfs.com/ipfs/" + hash
}
