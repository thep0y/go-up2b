/*
 * @Author: thepoy
 * @Email: thepoy@163.com
 * @File Name: response.go (c) 2021
 * @Created: 2021-06-24 11:51:36
 * @Modified: 2021-06-24 11:51:52
 */

package request

import "net/http"

type Response struct {
	StatusCode int
	Body       []byte
	Request    *http.Request
	Header     http.Header
}

func (r *Response) String() string {
	return string(r.Body)
}
