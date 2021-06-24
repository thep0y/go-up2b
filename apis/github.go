/*
 * @Author: thepoy
 * @Email: thepoy@163.com
 * @File Name: github.go
 * @Created: 2021-06-21 09:52:54
 * @Modified: 2021-06-24 11:58:09
 */

package apis

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/thep0y/go-up2b/models"
	"github.com/thep0y/go-up2b/request"
	"github.com/tidwall/gjson"
)

type Github struct {
	Config  *models.LoginInfo
	headers http.Header
	request *request.Request
	maxSize int
}

func (g *Github) NewRequest() {
	g.request = request.NewRequest(g.headers)
}

func (g *Github) MakeHeaders(headers map[string]string) error {
	g.headers = make(http.Header)
	for k, v := range headers {
		g.headers.Add(k, v)
	}
	return nil
}

func (g Github) GetHeaders() http.Header {
	return g.headers
}

func (g Github) String() string {
	return "github.com"
}

func (g Github) Login(config models.LoginInfo) (*models.LoginInfo, error) { return nil, nil }

func (g Github) checkSize(filepath string) error {
	f, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	if f.Size() > int64(g.maxSize)*int64(math.Pow(10, 6)) {
		return fmt.Errorf("the upper limit of the single upload image size of [ %s ] is %dMB, and the current file size is %dMB:\n%s", GITHUB, g.maxSize, f.Size()/int64(math.Pow(10, 6)), filepath)
	}
	return nil
}

func (g Github) UploadImage(imagePath string) (string, error) {
	err := g.checkSize(imagePath)
	if err != nil {
		return "", err
	}

	file, err := os.Open(imagePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	suffix := path.Ext(imagePath)

	filename := fmt.Sprintf(
		"%d%s",
		time.Now().UnixNano()/1000,
		suffix,
	)

	href := g.BaseURL() + filename

	reader := bufio.NewReader(file)
	fileBytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}

	_, localFilename := filepath.Split(imagePath)
	data := map[string]string{
		"content": base64.StdEncoding.EncodeToString(fileBytes),
		"message": "typora - " + localFilename,
	}

	resp, err := g.request.Put(href, data)
	if err != nil {
		return "", err
	}

	result := gjson.ParseBytes(resp.Body)
	if resp.StatusCode == http.StatusCreated {
		return g.cdnURL(result.Get("content.download_url").String()), nil
	} else {
		return "", fmt.Errorf(
			"the server returns an incorrect response [ %d ], message: %s",
			resp.StatusCode,
			result.Get("message").String(),
		)
	}
}

func (g Github) UploadImages(imagesPath []string) ([]string, error) {

	var wg sync.WaitGroup

	result := make(map[string]string, len(imagesPath))

	for _, path := range imagesPath {
		wg.Add(1)

		go func(p string) {
			defer wg.Done()
			u, err := g.UploadImage(p)
			if err == nil {
				result[p] = u
			}
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

func (g Github) BaseURL() string {
	return fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/contents/%s/",
		g.Config.Username,
		g.Config.Repo,
		g.Config.Folder,
	)
}

func (g Github) cdnURL(src string) string {
	p := strings.Split(src, "/main/")
	return fmt.Sprintf(
		"https://cdn.jsdelivr.net/gh/%s/%s/%s",
		g.Config.Username,
		g.Config.Repo,
		p[len(p)-1],
	)
}
