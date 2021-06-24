/*
 * @Author: thepoy
 * @Email: thepoy@163.com
 * @File Name: smms.go (c) 2021
 * @Created:  2021-06-24 09:19:17
 * @Modified: 2021-06-24 10:43:17
 */

package apis

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/thep0y/go-logger/log"
	"github.com/thep0y/go-up2b/models"
	"github.com/thep0y/go-up2b/request"
	"github.com/tidwall/gjson"
)

type SmMs struct {
	Config     *models.LoginInfo
	headers    http.Header
	request    *request.Request
	ConfigFile string
	maxSize    int
}

func (s *SmMs) MakeHeaders(headers map[string]string) error {
	if len(headers) != 0 {
		return errors.New("[ sm.ms ] cannot use custom headers")
	}
	headers = map[string]string{
		"Authorization": s.Config.Token,
	}
	s.headers = make(http.Header)
	for k, v := range headers {
		s.headers.Add(k, v)
	}
	return nil
}

func (s *SmMs) NewRequest() {
	s.request = request.NewRequest(s.headers)
}

func (s *SmMs) GetHeaders() http.Header {
	return s.headers
}

func (s SmMs) String() string {
	return "sm.ms"
}

func (s SmMs) Login(config models.LoginInfo) (*models.LoginInfo, error) { return nil, nil }

func (s SmMs) checkSize(filepath string) error {
	f, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	if f.Size() > int64(s.maxSize)*int64(math.Pow(10, 6)) {
		return fmt.Errorf("the upper limit of the single upload image size of [ %s ] is %dMB, and the current file size is %dMB:\n%s", SMMS, s.maxSize, f.Size()/int64(math.Pow(10, 6)), filepath)
	}
	return nil
}

func (s SmMs) createFormFile(p string) (string, io.Reader, error) {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	_, filename := filepath.Split(p)
	fw, err := w.CreateFormFile("smfile", filename)
	if err != nil {
		return "", nil, err
	}

	f, err := os.Open(p)
	if err != nil {
		return "", nil, err
	}
	defer f.Close()

	_, err = io.Copy(fw, f)
	if err != nil {
		return "", nil, err
	}

	w.Close()
	return w.FormDataContentType(), buf, nil
}

func (s SmMs) UploadImage(imagePath string) (string, error) {
	href := s.url("upload")

	contentType, reader, err := s.createFormFile(imagePath)
	if err != nil {
		return "", err
	}

	resp, err := s.request.PostMultipartForm(href, contentType, reader)
	if err != nil {
		return "", err
	}

	result := gjson.ParseBytes(resp.Body)
	if result.Get("success").Bool() {
		return result.Get("data.url").String(), nil
	}

	return "", errors.New(result.Get("message").String())
}

func (s SmMs) UploadImages(imagesPath []string) ([]string, error) {
	var wg sync.WaitGroup

	result := make(map[string]string, len(imagesPath))

	for _, path := range imagesPath {
		wg.Add(1)

		go func(p string) {
			defer wg.Done()
			u, err := s.UploadImage(p)
			if err != nil {
				log.Error(err)
			}
			result[p] = u
		}(path)
	}

	wg.Wait()

	downloadURL := make([]string, 0)
	for _, p := range imagesPath {
		if u, ok := result[p]; ok {
			downloadURL = append(downloadURL, u)
		}
	}

	return downloadURL, nil
}

func (s SmMs) url(path string) string {
	return s.BaseURL() + path
}

func (s SmMs) BaseURL() string {
	return "https://sm.ms/api/v2/"
}
