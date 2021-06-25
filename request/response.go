/*
 * @Author: thepoy
 * @Email: thepoy@163.com
 * @File Name: response.go (c) 2021
 * @Created: 2021-06-24 11:51:36
 * @Modified: 2021-06-25 08:31:51
 */

package request

import "net/http"

// Response is a network response that contains
// useful information
type Response struct {
	// Response status code
	StatusCode int

	// Response body
	Body []byte

	// the request to get this Response
	Request *http.Request

	// Response header
	Header http.Header
}

// String returns the string of the response body
func (r *Response) String() string {
	return string(r.Body)
}
