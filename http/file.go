// http/file.go kee > 2020/12/10

package http

import (
	//"fmt"
	"errors"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"okauth/utils"
	"os"
	"strings"
)

type Files struct {
	Names  map[string][]FileStream
	Values utils.Values
}

type File interface {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Closer
}

type FileStream struct {
	Filename   string
	MimeHeader textproto.MIMEHeader
	Size       int64
	File       File
}

func setFile(header *multipart.FileHeader, file multipart.File) FileStream {
	return FileStream{
		Filename:   header.Filename,
		MimeHeader: header.Header,
		Size:       header.Size,
		File:       file,
	}
}

func ParseFormFile(r *http.Request) Files {
	formFiles := utils.Values{}
	names := make(map[string][]FileStream)
	for n, fl := range r.MultipartForm.File {

		for _, fh := range fl {
			x := n
			if -1 < strings.Index(x, "[") {
				x = strings.Replace(x, "[", ".", -1)
				x = strings.Replace(x, "]", "", -1)
			}
			if file, _, e := r.FormFile(n); e == nil {
				f := setFile(fh, file)
				formFiles.Set(x, f)
				names[n] = append(names[n], f)
			}
		}
	}
	return Files{
		Names:  names,
		Values: formFiles,
	}
}

func (f Files) Get(name string) (FileStream, error) {
	if -1 < strings.Index(name, "[") {
		if file, ok := f.Names[name]; ok {
			return file[0], nil
		} else {
			return FileStream{}, errors.New("invalid file range of " + name)
		}
	} else {
		file := f.Values.Get(name)
		switch file.(type) {
		case FileStream:
			{
				return file.(FileStream), nil
			}
		default:
			return FileStream{}, errors.New("invalid file range of " + name)
		}
	}
}

func (f Files) Has(name string) (ok bool) {
	if -1 < strings.Index(name, "[") {
		_, ok = f.Names[name]
		return
	}
	file := f.Values.Get(name)
	switch file.(type) {
	case FileStream:
		ok = true
	}
	return
}

func (f FileStream) GuessExtension() string {
	mimeTypes := NewMimeTypes()
	return mimeTypes.GuessExtension(f.MimeType())
}

func (f FileStream) Extension() string {
	ext := f.GuessExtension()
	if ext == "undefined" {
		fn := strings.Split(f.Filename, ".")
		ext = fn[len(fn)-1]
	}
	return ext
}

func (f FileStream) MimeType() string {
	return f.MimeHeader["Content-Type"][0]
}

func (f FileStream) Stream() File {
	return f.File
}

func (f FileStream) CopyTo(path string) {
	f.CopyAs(path, f.HashName())
}

func (f FileStream) CopyAs(path string, filename string) {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		panic(err)
	}

	if "/" != string(path[len(path)-1]) {
		path += "/"
	}
	path += filename
	out, e := os.Create(path)
	if e != nil {
		panic(e)
	}
	defer out.Close()
	_, e = io.Copy(out, f.File)
	if e != nil {
		panic(e)
	}
}

func (f FileStream) HashName() string {
	const letterBytes = "qwertyuiopasdfghjklzxcvbnm-QWERTYUIOPASDFGHJKLZXCVBNM_1234567890"
	b := make([]byte, 40)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b) + "." + f.GuessExtension()
}
