/*
 * @Author: thepoy
 * @Email: thepoy@163.com
 * @File Name: interface.go
 * @Created: 2021-06-21 09:52:54
 * @Modified: 2021-06-24 11:49:06
 */

package apis

import (
	"errors"
	"net/http"

	"github.com/thep0y/go-up2b/models"
)

type Client interface {
	// Only the image bed that does not provide api
	// needs to implement the login method, such as
	// "imgtu.com"
	Login(models.LoginInfo) (*models.LoginInfo, error)

	// UploadImage uploads a image to the image bed
	UploadImage(imagePath string) (string, error)

	// UploadImages uploads multiple images to the image bed
	UploadImages(imagesPath []string) ([]string, error)

	// String returns the domain name of the image bed
	String() string

	// BaseURL returns the basic url of the image bed
	BaseURL() string

	// MakeHeaders creates the header required by the image
	// bed when uploading images
	MakeHeaders(map[string]string) error

	// GetHeaders returns the header used by the image bed
	// when uploading images
	GetHeaders() http.Header

	// NewRequest creates the request needed for the image bed
	NewRequest()

	// checkSize checks if there are any images that exceed
	// the limit size of the image bed
	checkSize(filepath string) error
}

// NewImageBedClient creates a image bed client
// according to the configuration file
func NewImageBedClient(config models.Config, configGile string) (Client, error) {
	switch config.ImageBed {
	case SMMS:
		smms := &SmMs{
			Config:  config.AuthData[SMMS],
			maxSize: 5,
		}
		smms.MakeHeaders(nil)
		smms.NewRequest()
		return smms, nil
	case IMGTU:
		imgtu := &Imgtu{
			Config:     config.AuthData[IMGTU],
			ConfigFile: configGile,
			maxSize:    10,
		}
		imgtu.MakeHeaders(nil)
		imgtu.NewRequest()

		return imgtu, nil
	case GITEE:
		gitee := &Gitee{
			Config:  config.AuthData[GITEE],
			maxSize: 1,
		}
		gitee.MakeHeaders(nil)
		gitee.NewRequest()

		return gitee, nil
	case GITHUB:
		github := &Github{
			Config:  config.AuthData[GITHUB],
			maxSize: 20,
		}

		github.MakeHeaders(map[string]string{
			"Accept":        "application/vnd.github.v3+json",
			"Authorization": "token " + github.Config.Token,
		})

		github.NewRequest()

		return github, nil
	default:
		return nil, errors.New("unknown image bed code")
	}
}
