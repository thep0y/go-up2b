// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonD2b7633eDecodeGithubComThep0yGoUp2bModels(in *jlexer.Lexer, out *LoginInfo) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "token":
			out.Token = string(in.String())
		case "username":
			out.Username = string(in.String())
		case "password":
			out.Password = string(in.String())
		case "repo":
			out.Repo = string(in.String())
		case "folder":
			out.Folder = string(in.String())
		case "cookie":
			out.Cookie = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComThep0yGoUp2bModels(out *jwriter.Writer, in LoginInfo) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Token != "" {
		const prefix string = ",\"token\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Token))
	}
	if in.Username != "" {
		const prefix string = ",\"username\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Username))
	}
	if in.Password != "" {
		const prefix string = ",\"password\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Password))
	}
	if in.Repo != "" {
		const prefix string = ",\"repo\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Repo))
	}
	if in.Folder != "" {
		const prefix string = ",\"folder\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Folder))
	}
	if in.Cookie != "" {
		const prefix string = ",\"cookie\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Cookie))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v LoginInfo) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComThep0yGoUp2bModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v LoginInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComThep0yGoUp2bModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *LoginInfo) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComThep0yGoUp2bModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *LoginInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComThep0yGoUp2bModels(l, v)
}
func easyjsonD2b7633eDecodeGithubComThep0yGoUp2bModels1(in *jlexer.Lexer, out *GithubOK) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "content":
			easyjsonD2b7633eDecode(in, &out.Content)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComThep0yGoUp2bModels1(out *jwriter.Writer, in GithubOK) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"content\":"
		out.RawString(prefix[1:])
		easyjsonD2b7633eEncode(out, in.Content)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GithubOK) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComThep0yGoUp2bModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GithubOK) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComThep0yGoUp2bModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GithubOK) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComThep0yGoUp2bModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GithubOK) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComThep0yGoUp2bModels1(l, v)
}
func easyjsonD2b7633eDecode(in *jlexer.Lexer, out *struct {
	Sha         string
	DownloadURL string `json:"download_url"`
}) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "sha":
			out.Sha = string(in.String())
		case "download_url":
			out.DownloadURL = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncode(out *jwriter.Writer, in struct {
	Sha         string
	DownloadURL string `json:"download_url"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"sha\":"
		out.RawString(prefix[1:])
		out.String(string(in.Sha))
	}
	{
		const prefix string = ",\"download_url\":"
		out.RawString(prefix)
		out.String(string(in.DownloadURL))
	}
	out.RawByte('}')
}
func easyjsonD2b7633eDecodeGithubComThep0yGoUp2bModels2(in *jlexer.Lexer, out *FileData) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
	} else {
		in.Delim('{')
		*out = make(FileData)
		for !in.IsDelim('}') {
			key := string(in.String())
			in.WantColon()
			var v1 string
			v1 = string(in.String())
			(*out)[key] = v1
			in.WantComma()
		}
		in.Delim('}')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComThep0yGoUp2bModels2(out *jwriter.Writer, in FileData) {
	if in == nil && (out.Flags&jwriter.NilMapAsEmpty) == 0 {
		out.RawString(`null`)
	} else {
		out.RawByte('{')
		v2First := true
		for v2Name, v2Value := range in {
			if v2First {
				v2First = false
			} else {
				out.RawByte(',')
			}
			out.String(string(v2Name))
			out.RawByte(':')
			out.String(string(v2Value))
		}
		out.RawByte('}')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v FileData) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComThep0yGoUp2bModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FileData) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComThep0yGoUp2bModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FileData) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComThep0yGoUp2bModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FileData) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComThep0yGoUp2bModels2(l, v)
}
func easyjsonD2b7633eDecodeGithubComThep0yGoUp2bModels3(in *jlexer.Lexer, out *Config) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "image_bed":
			out.ImageBed = ImageBedCode(in.Int())
		case "auth_data":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('[')
				v3 := 0
				for !in.IsDelim(']') {
					if v3 < 4 {
						if in.IsNull() {
							in.Skip()
							(out.AuthData)[v3] = nil
						} else {
							if (out.AuthData)[v3] == nil {
								(out.AuthData)[v3] = new(LoginInfo)
							}
							(*(out.AuthData)[v3]).UnmarshalEasyJSON(in)
						}
						v3++
					} else {
						in.SkipRecursive()
					}
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonD2b7633eEncodeGithubComThep0yGoUp2bModels3(out *jwriter.Writer, in Config) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"image_bed\":"
		out.RawString(prefix[1:])
		out.Int(int(in.ImageBed))
	}
	{
		const prefix string = ",\"auth_data\":"
		out.RawString(prefix)
		out.RawByte('[')
		for v4 := range in.AuthData {
			if v4 > 0 {
				out.RawByte(',')
			}
			if (in.AuthData)[v4] == nil {
				out.RawString("null")
			} else {
				(*(in.AuthData)[v4]).MarshalEasyJSON(out)
			}
		}
		out.RawByte(']')
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Config) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonD2b7633eEncodeGithubComThep0yGoUp2bModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Config) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonD2b7633eEncodeGithubComThep0yGoUp2bModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Config) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonD2b7633eDecodeGithubComThep0yGoUp2bModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Config) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonD2b7633eDecodeGithubComThep0yGoUp2bModels3(l, v)
}
