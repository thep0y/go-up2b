/*
 * @Author: thepoy
 * @Email: email@example.com
 * @File Name: ipfs_test.go
 * @Created: 2021-07-27 15:02:53
 * @Modified: 2021-07-27 15:15:02
 */

package apis

import "testing"

func TestIpfs(t *testing.T) {
	i := new(Ipfs)
	i.MakeHeaders(nil)
	i.NewRequest()

	u, err := i.UploadImage("/mnt/c/Users/thepoy/Pictures/截图/微信截图_20210704171840.png")
	if err != nil {
		t.Error(err)
	}
	t.Log(u)

	us, err := i.UploadImages([]string{
		"/mnt/c/Users/thepoy/Pictures/截图/微信截图_20210704171840.png",
		"/mnt/c/Users/thepoy/Pictures/截图/微信截图_20210704171636.png",
		"/mnt/c/Users/thepoy/Pictures/截图/微信截图_20210704160614.png",
		"/mnt/c/Users/thepoy/Pictures/截图/微信截图_20210704151359.png",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(us)
}
