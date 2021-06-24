/*
 * @Author: thepoy
 * @Email: thepoy@163.com
 * @File Name: models.go
 * @Created: 2021-06-21 09:52:54
 * @Modified: 2021-06-24 11:55:13
 */

//go:generate easyjson -all -snake_case $GOFILE

package models

type ImageBedCode int

var (
	imageBeds = []string{
		"sm.ms",
		"imgtu.com",
		"gitee.com",
		"github.com",
	}
)

// String returns the domain name of the
// image bed corresponding to the code
func (ibc ImageBedCode) String() string {
	return imageBeds[ibc]
}

//easyjson:json
type FileData map[string]string

// Config stores configuration information of
// all picture beds
type Config struct {
	// image bed code in use
	ImageBed ImageBedCode

	// An array with the same length as the
	// number of image beds
	AuthData [4]*LoginInfo
}

// LoginInfo stores the authentication information of
// each image bed
type LoginInfo struct {
	Token    string `json:"token,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Repo     string `json:"repo,omitempty"`
	Folder   string `json:"folder,omitempty"`
	Cookie   string `json:"cookie,omitempty"`
}

type GithubOK struct {
	Content struct {
		Sha         string
		DownloadURL string `json:"download_url"`
	}
}
