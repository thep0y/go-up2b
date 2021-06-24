/*
 * @Author: thepoy
 * @Email: thepoy@163.com
 * @File Name: imgtu_test.go (c) 2021
 * @Created:  2021-06-22 22:33:56
 * @Modified: 2021-06-22 23:05:32
 */

package apis

import (
	"fmt"
	"testing"

	"github.com/thep0y/go-up2b/models"
)

func TestLogin(t *testing.T) {
	config := &models.LoginInfo{
		Username: "timg_test",
		Password: "timg_test",
	}

	imgtu := new(Imgtu)
	imgtu.Config = config

	imgtu.MakeHeaders(map[string]string{
		"Accept":     "application/json",
		"User-Agent": "Mozilla/5.0 (X11; Linux x86_64; rv:85.0) Gecko/20100101 Firefox/85.0",
	})

	imgtu.NewRequest()

	info, err := imgtu.Login(*config)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(info)
}
