/*
 * @Author: thepoy
 * @Email: thepoy@163.com
 * @File Name: imgtu.go (c) 2021
 * @Created:  2021-06-22 17:51:13
 * @Modified: 2021-07-03 20:44:15
 */

package apis

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/mailru/easyjson"
	"github.com/thep0y/go-logger/log"
	"github.com/thep0y/go-up2b/models"
	"github.com/thep0y/go-up2b/request"
	"github.com/tidwall/gjson"
)

type Imgtu struct {
	Config     *models.LoginInfo
	headers    http.Header
	request    *request.Request
	ConfigFile string
	maxSize    int
}

func (i *Imgtu) MakeHeaders(headers map[string]string) error {
	if headers == nil {
		headers = map[string]string{
			"Accept":     "application/json",
			"User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:85.0) Gecko/20100101 Firefox/85.0",
		}
	}
	i.headers = make(http.Header)
	for k, v := range headers {
		i.headers.Add(k, v)
	}
	return nil
}

func (i *Imgtu) NewRequest() {
	i.request = request.NewRequest(i.headers)
}

func (i Imgtu) GetHeaders() http.Header {
	return i.headers
}

func (i Imgtu) String() string {
	return "imgtu.com"
}

// Login gets cookies by username and password
func (i Imgtu) Login(config models.LoginInfo) (*models.LoginInfo, error) {
	href := i.url("login")
	token, cookie, err := i.parseAuthToken(href)
	if err != nil {
		return nil, err
	}

	headers := i.headers.Clone()
	headers.Add("Cookie", cookie)
	headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	headers.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := request.PostWithHeadersWithNoRedirect(
		href,
		headers,
		models.FileData{
			"login-subject": config.Username,
			"password":      config.Password,
			"auth_token":    token,
		},
	)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusMovedPermanently {
		finalCookie := cookie + "; " + strings.Split(resp.Header.Get("Set-Cookie"), "; ")[0]
		finalConfig := &models.LoginInfo{
			Token:    token,
			Cookie:   finalCookie,
			Username: config.Username,
			Password: config.Password,
		}
		return finalConfig, nil
	} else {
		return nil, fmt.Errorf("invalid response [ %d ], you need check your account and password", resp.StatusCode)
	}
}

func (i Imgtu) parseAuthToken(u string) (string, string, error) {
	resp, err := i.request.Get(u)
	if err != nil {
		return "", "", err
	}

	if resp.StatusCode == http.StatusOK {
		re, err := regexp.Compile(`PF.obj.config.auth_token = "([a-f0-9]{40})"`)
		if err != nil {
			return "", "", err
		}

		result := re.FindAllStringSubmatch(resp.String(), 1)
		if len(result) >= 1 {
			setCookie := strings.Split(resp.Header.Get("Set-Cookie"), "; ")[0]
			return result[0][1], setCookie, nil
		} else {
			return "", "", errors.New("`auth_token` was not found")
		}
	} else {
		return "", "", fmt.Errorf("response error [ %d ]", resp.StatusCode)
	}
}

func (i Imgtu) checkSize(filepath string) error {
	f, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	if f.Size() > int64(i.maxSize)*int64(math.Pow(10, 6)) {
		return fmt.Errorf("the upper limit of the single upload image size of [ %s ] is %dMB, and the current file size is %dMB:\n%s", IMGTU, i.maxSize, f.Size()/int64(math.Pow(10, 6)), filepath)
	}
	return nil
}

func (i Imgtu) createMuiltipartForm(p string) (string, io.Reader, error) {
	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)

	f, err := os.Open(p)
	if err != nil {
		return "", nil, err
	}
	defer f.Close()

	type_, err := w.CreateFormField("type")
	if err != nil {
		return "", nil, err
	}
	type_.Write([]byte("file"))

	action, err := w.CreateFormField("action")
	if err != nil {
		return "", nil, err
	}
	action.Write([]byte("upload"))

	timestamp, err := w.CreateFormField("timestamp")
	if err != nil {
		return "", nil, err
	}
	timestamp.Write([]byte(strconv.Itoa(int(time.Now().Local().UnixNano() / 1000 / 1000))))

	authToken, err := w.CreateFormField("auth_token")
	if err != nil {
		return "", nil, err
	}
	authToken.Write([]byte(i.Config.Token))

	nsfw, err := w.CreateFormField("nsfw")
	if err != nil {
		return "", nil, err
	}
	nsfw.Write([]byte("0"))

	_, filename := filepath.Split(p)
	ff, err := w.CreateFormFile("source", filename)
	if err != nil {
		return "", nil, err
	}
	io.Copy(ff, f)

	w.Close()
	return w.FormDataContentType(), buf, nil
}

func (i *Imgtu) updateAuthToken() error {
	resp, err := i.request.Get(i.BaseURL())
	if err != nil {
		return err
	}

	re, err := regexp.Compile(`PF.obj.config.auth_token = "([a-f0-9]{40})"`)
	if err != nil {
		return err
	}

	result := re.FindAllStringSubmatch(resp.String(), 1)
	if len(result) < 1 {
		return errors.New("`auth_token` was not found")
	}

	i.Config.Token = result[0][1]

	go func() {
		configData, err := os.ReadFile(i.ConfigFile)
		if err != nil {
			log.Error(err)
			return
		}

		config := new(models.Config)
		err = easyjson.Unmarshal(configData, config)
		if err != nil {
			log.Error(err)
			return
		}
		config.AuthData[IMGTU].Token = result[0][1]

		newConfigData, err := easyjson.Marshal(config)
		if err != nil {
			log.Error(err)
			return
		}

		err = ioutil.WriteFile(i.ConfigFile, newConfigData, 0o644)
		if err != nil {
			log.Error(err)
			return
		}
	}()

	return nil
}

func (i Imgtu) UploadImage(imagePath string) (string, error) {
	err := i.checkSize(imagePath)
	if err != nil {
		return "", err
	}

	file, err := os.Open(imagePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	i.headers.Add("Cookie", i.Config.Cookie)
	i.request = request.NewRequest(i.headers)

	contentType, reader, err := i.createMuiltipartForm(imagePath)
	if err != nil {
		return "", err
	}

	href := i.url("json")

	resp, err := i.request.PostMultipartForm(href, contentType, reader)
	if err != nil {
		return "", err
	}

	result := gjson.ParseBytes(resp.Body)
	if resp.StatusCode != http.StatusCreated {
		msg := result.Get("error.message").String()
		if msg == "请求被拒绝 (auth_token)" {
			log.Warn("the token has expired and will be automatically updated")
			err = i.updateAuthToken()
			if err != nil {
				return "", err
			}
			return i.UploadImage(imagePath)
		}
	}

	return result.Get("image.image.url").String(), nil
}

func (i Imgtu) UploadImages(imagesPath []string) ([]string, error) {
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

func (i Imgtu) url(path string) string {
	return i.BaseURL() + path
}

func (i Imgtu) BaseURL() string {
	return "https://imgtu.com/"
}
