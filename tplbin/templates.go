// Code generated by go-bindata. DO NOT EDIT.
// sources:
// templates/postgres.enum.go.tpl
// templates/postgres.foreignkey.go.tpl
// templates/postgres.index.go.tpl
// templates/postgres.proc.go.tpl
// templates/postgres.query.go.tpl
// templates/postgres.querytype.go.tpl
// templates/postgres.tuple.go.tpl
// templates/postgres.type.go.tpl
// templates/xo_db.go.tpl
// templates/xo_package.go.tpl

package tplbin


import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}


type asset struct {
	bytes []byte
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataPostgresenumgotpl = "" +
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x55\xc1\x6a\xdc\x30\x10\x3d\x5b\x5f\x31\x15\x85\x58\x21\xf1\xd2\x4b" +
	"\x0f\x81\x3d\x95\x1e\x9b\x43\xd3\xe6\x52\x7a\xd0\x7a\x47\x59\x11\x5b\x6e\x24\x79\x93\x60\xf4\xef\x65\x24\x99\x95" +
	"\xd3\x4d\x4b\x0a\x7b\x59\xc4\x68\xfc\xe6\xcd\x7b\x4f\xec\x34\x5d\xc2\x7b\xff\xfc\x0b\xe1\x6a\x0d\xcd\xb5\xec\x11" +
	"\x2e\x43\x60\xb1\xec\x76\x83\xf5\x54\xaf\xe3\xc9\xd0\x65\xea\xe5\x68\xc6\xfe\x56\x76\x1c\xb8\xc7\x27\xcf\x81\x6f" +
	"\x46\xc5\x81\x0f\xf7\x1c\xb8\xb3\x2d\x17\x07\x14\x8b\x7b\xb4\x0e\x09\xda\xc5\x21\x5f\x53\xe1\xd3\x60\x9c\x4f\x55" +
	"\xea\x5d\xad\x60\x9a\x32\x7c\x08\xa0\x1d\xf8\x1d\xc2\xd9\x34\x41\xf3\xd9\x8c\x7d\xfc\x89\xf4\x42\x38\x03\x1a\x0f" +
	"\xb1\x55\xd9\xa1\x07\xd7\xee\xb0\x97\xa9\xf9\x26\x9d\xa9\xad\x61\xb1\xa5\x84\x1d\xb5\xf1\x1f\x3e\x32\xd6\xd2\x70" +
	"\xa8\x23\x43\x2b\xcd\x1d\x42\x73\x2b\xbb\x11\x1d\x84\xc0\xaa\xc4\x45\xab\x17\xe4\x43\xa0\x01\x99\x44\x81\x3a\x4d" +
	"\x80\x9d\xfb\xb3\x58\xb4\xa2\xd9\xbe\xdc\xea\x56\x76\x71\xa9\x38\x37\x6e\x55\x7c\xdd\xb0\xea\x34\x0c\xd6\xe5\x94" +
	"\x7a\xe6\x11\xbd\x98\x89\x08\x96\xdb\xc9\x16\xc1\xc8\x99\x1b\x6f\xb5\xb9\x03\x8b\x7e\xb4\x26\xed\xe0\x52\x69\x1f" +
	"\x3f\x1a\x54\xac\x2d\x16\x50\xa3\x69\x81\x26\xe4\x1c\x85\x50\xde\x8b\x8c\x59\x8b\x19\x69\x62\xd5\x5e\x5a\xc8\xc9" +
	"\xca\x55\xc6\x2a\xf7\xa8\x7d\xbb\x83\x25\xd0\x2b\xc6\xb5\xd2\xe1\x69\xac\xbb\x62\x55\x35\x53\x5b\x03\x3f\x66\x20" +
	"\x2f\x75\xab\x02\x63\x55\xd2\x6b\x5e\x89\x85\xa8\xe5\x17\x69\xdd\x4e\x76\xdf\xf0\xc9\x43\x9f\xce\x6e\x19\x7d\xe3" +
	"\x07\xa0\x67\xf5\x6f\x0d\x0b\xac\x5a\x40\xfd\xe3\xe7\xe6\xd9\xe3\x05\xa0\xb5\x83\x15\xa4\x68\x66\x90\x2e\x16\x40" +
	"\xcd\xac\xbf\xb8\x00\xa3\x67\x72\xdf\x4d\x5f\xd0\x1b\xcd\x51\x82\xf1\xcd\xbd\x4a\xf0\x7c\xc1\x70\x01\x58\xd3\x47" +
	"\x99\x8c\x48\x2c\x89\x64\x76\x38\x39\x1e\x7b\x44\xf5\x57\x87\x8f\xcb\x4f\x16\x9d\x2f\xa8\xac\x4f\x93\x05\x76\x38" +
	"\xb1\x6a\x8b\x4a\x8e\x9d\xa7\xe1\xb3\xdd\xb4\x97\x6b\xae\xf1\xb1\xe6\xda\xec\x65\xa7\xb7\xa5\x7c\x5c\x2c\xc2\x71" +
	"\xd0\x3e\x2d\xe2\xa4\xd7\x4e\x69\xcc\xaf\xec\xa1\x5b\x6d\xad\xde\xa3\x4d\x22\x58\x4a\x07\x5a\x25\x5b\x04\x45\xea" +
	"\xbd\xe5\xc5\x45\x04\xca\x49\x89\x78\x24\x2d\x47\x63\x52\xa6\xe4\xa6\x95\xe6\x05\xd1\xad\xf4\x72\x23\x1d\xae\xdc" +
	"\x43\xd7\xd0\xbd\x79\x33\xd7\x65\x70\x08\xa3\x76\xb6\x3d\x80\x4c\xa1\xc8\xcc\x66\x54\x17\x30\xdc\xd3\x1f\x8a\xb3" +
	"\x6d\x93\xa3\x2f\x58\xa5\x15\xbc\x1b\xee\xa9\x05\x00\xfe\xcb\x91\xc5\xfa\xcb\xfc\x6e\x46\x25\x48\x83\xdf\x01\x00" +
	"\x00\xff\xff\x33\x36\x2f\x65\x36\x07\x00\x00"

func bindataPostgresenumgotplBytes() ([]byte, error) {
	return bindataRead(
		_bindataPostgresenumgotpl,
		"postgres.enum.go.tpl",
	)
}



func bindataPostgresenumgotpl() (*asset, error) {
	bytes, err := bindataPostgresenumgotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "postgres.enum.go.tpl",
		size: 0,
		md5checksum: "",
		mode: os.FileMode(0),
		modTime: time.Unix(0, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataPostgresforeignkeygotpl = "" +
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x41\x4b\xc4\x30\x10\x85\xcf\xe6\x57\xbc\x83\xb0\xc9\xb2\xdb\xde" +
	"\x05\x2f\xab\xe8\x41\x50\x10\x0f\x5e\xbb\xed\xd4\x16\xb7\x89\x4c\x52\xb5\x84\xfc\x77\x49\xda\xed\xd6\x65\x0f\x81" +
	"\xe1\x9b\x37\x99\xf7\xc6\xfb\x2d\xae\x6d\x63\xd8\xe1\xe6\x16\x32\x55\xba\xe8\x08\xd9\xdb\xf0\x45\xd9\x73\xd1\x91" +
	"\xc2\x36\x04\x91\xe7\xf0\x1e\x09\x20\x04\x30\xb9\x9e\xb5\x85\x6b\x28\xf1\x57\xaa\xe7\x81\xd8\x2f\xac\x35\x65\x5b" +
	"\x38\xaa\xf0\xd3\xba\x66\xd6\x2d\x45\x2b\x9b\xd0\x43\x4b\x87\x6a\x1e\x94\x27\x74\x67\x0e\xf1\xf5\x9d\x9e\x9a\x2a" +
	"\x13\x79\x1e\x9d\x3c\x92\x26\x4e\x9f\xd7\x6c\x3a\xd4\x86\xa9\xfd\xd0\xf8\xa4\x01\xab\x34\x3f\x82\x27\x1a\x16\xe5" +
	"\x71\x6b\x26\xea\x5e\x97\x69\xd1\x94\x3c\x04\xac\xcf\xcd\xa9\x65\x5c\x59\xed\xf1\xfe\x72\xbf\x53\x90\xeb\x0b\x69" +
	"\x37\x20\x66\xc3\x0a\x5e\x5c\x8d\x87\xb9\x74\x93\xdd\x30\xc1\x7f\x81\x65\xb5\xdf\x44\x75\x69\xf4\x37\xfd\xba\xa3" +
	"\xa5\xf1\x04\x27\x79\x74\x24\x82\x10\x7f\x01\x00\x00\xff\xff\xea\x89\x96\x81\xb0\x01\x00\x00"

func bindataPostgresforeignkeygotplBytes() ([]byte, error) {
	return bindataRead(
		_bindataPostgresforeignkeygotpl,
		"postgres.foreignkey.go.tpl",
	)
}



func bindataPostgresforeignkeygotpl() (*asset, error) {
	bytes, err := bindataPostgresforeignkeygotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "postgres.foreignkey.go.tpl",
		size: 0,
		md5checksum: "",
		mode: os.FileMode(0),
		modTime: time.Unix(0, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataPostgresindexgotpl = "" +
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\xdd\x4e\x1b\x3d\x10\xbd\xb6\x9f\x62\x3e\xeb\x13\x6c\xda\xb0\xb9" +
	"\x47\xca\x45\x0b\xa1\xad\x4a\xa1\x05\xaa\x22\x21\xd4\x38\xbb\xb3\xb0\xd2\xc6\xce\xda\x5e\x20\xb2\xfc\xee\xd5\xd8" +
	"\xbb\x94\x9f\x08\x41\x6f\x1c\xdb\xe3\x99\x73\xe6\x9c\xc9\x7a\xbf\x03\xff\xdb\x6b\x6d\x1c\xec\x4e\x21\x8b\x3b\x25" +
	"\x97\x08\xf9\xd9\x7a\x85\xf9\x11\x6d\x05\x1a\x23\x40\xd8\xb6\xb1\x8e\x36\xe5\x42\x80\x68\x05\x08\x83\x56\x80\x38" +
	"\x3f\x3e\xd4\x57\x02\xf2\x83\x1a\x9b\xd2\x8e\x60\x27\x04\x1e\xcb\x3a\xb9\x68\x30\x95\x2d\xae\x71\x29\x21\x3f\xed" +
	"\x7f\x63\xed\x33\x0a\xa7\x95\x60\x52\xe2\x64\x02\xde\x43\x7e\xd0\xa9\x22\x62\x87\x00\x06\x9d\xa9\xf1\x06\x2d\x48" +
	"\x30\xfa\x16\x2a\xa3\x97\xb0\xed\xfd\x00\x10\xc2\x36\x48\x0a\x52\xe2\x5f\xd6\x21\xe4\x7c\x32\xa1\x82\x9f\x50\xa1" +
	"\x91\x0e\xcb\x94\x5a\xab\x12\xef\x62\x81\xfc\x0b\x6d\xd3\xda\xe7\x6c\xe7\xbc\xea\x54\xf1\x94\x44\x56\x2e\xe0\xfc" +
	"\x78\xff\xa3\xf7\x70\xa5\x57\xd2\xc8\x65\x53\x5b\x37\xf4\x0c\xce\x74\x98\x96\x10\x46\x90\x79\x0f\x75\x05\x4a\xbb" +
	"\x7b\x04\xfb\x53\xd5\x6d\x0c\x5f\x5c\x7a\x0f\xa8\x4a\x08\xe1\xdd\x53\xc2\x63\x40\x63\xb4\x19\x81\xe7\xec\x46\x1a" +
	"\x3a\xa5\x1b\xce\xd9\x64\x02\xb6\x6d\xa0\xed\xd0\xac\x39\x2b\xb4\xb2\x0e\x92\x23\x30\x85\xf9\xe9\xec\x70\xb6\x77" +
	"\x06\x73\x78\xcf\x19\x9b\x7b\x0f\x85\x6e\xc8\x46\xdb\x03\xf4\x3c\x43\x18\x9e\x1c\x9c\x1c\x7f\x83\x87\x1a\x0e\x81" +
	"\x5f\x9f\x67\x27\x33\x78\x50\x21\x22\xde\x77\x2a\xe0\xc3\xd1\x3e\x08\x08\x61\x9e\x48\x99\x4e\x0d\xa4\xe2\x20\x64" +
	"\x89\xd4\x4b\x42\x55\xb2\xb1\x51\xa9\x38\x26\x75\xb5\x41\x25\xce\x88\x5b\x9a\xcb\x10\x68\x86\x9e\x6a\xe5\xe9\x49" +
	"\xca\x8e\xd7\xdf\x4d\xbd\x94\x66\xfd\x15\xd7\x31\x9d\xfd\xc6\xbb\xda\x3a\xbb\x1b\x21\xc7\xb1\x1e\xa9\x4e\x33\xc6" +
	"\x02\xe7\x8c\xb4\x9d\x42\xb9\xc8\x7f\x10\xf9\x13\x7d\xfb\x16\xe2\xf9\x69\x21\x15\xd9\x5c\x51\x74\x83\xd0\xd9\xca" +
	"\xd4\xca\x81\xd8\x12\x7d\x17\xa3\xd8\x2f\xab\xab\x68\xea\x7f\x53\x50\x75\x43\x36\x33\x83\xae\x33\x8a\x8e\xd1\xfd" +
	"\x44\xae\xbf\xdc\x7a\x28\xc2\x98\xde\x44\xc5\x30\xb1\xe0\xac\x8d\x29\xa4\xce\xd0\xc7\x9b\xd4\x7f\x1d\x1b\x56\x62" +
	"\x85\x06\xda\x7c\xaf\xd1\x16\xb3\x51\xb2\xbd\xd1\xb2\x04\x83\xb6\x6b\x9c\x25\xbe\x96\x58\x5c\x5c\x3e\x1b\x69\x1f" +
	"\x38\xab\x34\xa5\x1f\xe1\x9d\xcb\xe2\x68\xbf\xc6\xdb\x97\xcd\x7d\xe6\xee\x23\x7b\xa3\x84\xf1\x0f\x53\x48\xc5\x59" +
	"\x6f\x75\xfb\xcf\xa6\x6d\xd0\xe9\xb9\x50\x09\x94\x84\x98\x82\x5c\xad\x50\x95\x99\x41\x3b\x7e\xec\xe1\xe8\x91\xbd" +
	"\x31\x7e\x6f\x6a\xfc\x24\xf0\xc0\xf9\x9f\x00\x00\x00\xff\xff\xc3\x4f\x76\x7d\x93\x05\x00\x00"

func bindataPostgresindexgotplBytes() ([]byte, error) {
	return bindataRead(
		_bindataPostgresindexgotpl,
		"postgres.index.go.tpl",
	)
}



func bindataPostgresindexgotpl() (*asset, error) {
	bytes, err := bindataPostgresindexgotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "postgres.index.go.tpl",
		size: 0,
		md5checksum: "",
		mode: os.FileMode(0),
		modTime: time.Unix(0, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataPostgresprocgotpl = "" +
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\x4d\x6f\xda\x4c\x10\x3e\x7b\x7f\xc5\xbc\xe8\x55\x63\x47\x89\xdd" +
	"\x4b\x2e\x95\xb8\xb4\xa5\xa7\x28\x49\x43\x54\xe5\xd6\x38\xf6\x40\x56\x5a\x76\x61\x76\x4d\x13\xad\xe6\xbf\x57\xfb" +
	"\x01\x98\x96\x56\x2a\x12\x46\x8c\xf7\xf9\x1c\xdb\xfb\x4b\xf8\x5f\x1b\xf7\xcd\xc8\x1e\x3e\x4c\xa1\xd4\x08\xf5\x1d" +
	"\x99\xae\xbe\x47\x37\x90\x7e\x78\x5b\x23\x4c\xb6\x46\xf6\x93\x0a\x2e\x99\x45\x04\xb8\x61\xad\x30\x1d\x08\xa0\x45" +
	"\xab\x2c\xee\xef\xca\x05\x2c\x11\x4a\x85\xfa\x37\xa6\x0a\xae\x80\x59\x40\xfe\x78\x7f\xc4\x34\x85\x12\x37\x50\x5a" +
	"\x25\xbb\x13\x26\xde\xc3\x55\x05\x93\xf9\xec\xe1\xf6\xcb\xa4\xca\x52\xa8\xfb\x9d\xa5\x35\x99\x2e\x06\xb0\xdd\x0b" +
	"\xae\x5a\xa8\xe7\xf9\x37\xf2\x84\xcb\x4d\xbb\xc2\x6a\xec\xf2\x64\x54\x47\x72\xb9\x44\x9a\xc4\x83\x4d\x03\xde\x43" +
	"\x1d\x90\xc0\x0c\x5d\xab\x94\x05\xf7\x82\x60\x9d\x21\xec\x21\x88\x62\x3f\x10\xc2\x99\xf7\xd9\x03\x73\x19\x30\x81" +
	"\xf8\xae\xa5\x76\x65\x81\xb9\x82\xdd\x68\xac\xc5\x7c\x06\x46\x43\xff\x5c\x8b\xc5\xa0\xbb\xb1\x54\x99\x2d\x8e\xfb" +
	"\x61\xee\x9f\xe1\xdc\x6e\xd4\x6b\xfd\xf9\x63\x4c\xaf\x2c\xc6\xe1\xe3\x6d\x1e\xe8\x1e\x98\xbd\x87\xa5\x59\x07\x69" +
	"\x25\xad\x83\x3a\xbb\x70\x34\x60\xba\x04\x3f\xc1\x63\xe0\xdf\xad\x3e\xc2\x08\x5d\xf0\x95\x3d\xd6\xd9\xe4\x45\x30" +
	"\x96\xa8\x91\xc8\x50\x05\x5e\x14\xdb\x96\x00\x29\x7e\x0d\x09\x51\x34\x0d\xd8\x8d\x82\xcd\x80\xf4\x26\x8a\xce\x68" +
	"\xeb\xc2\xc0\x3a\x82\x29\x3c\xcd\x67\xd7\xb3\x4f\x0f\x90\x35\x8f\x32\x9d\xc3\x82\xcc\xea\x20\xf1\x4b\x8f\x9d\x51" +
	"\xdb\x56\xd9\x7d\x0a\xe6\xea\x29\xc9\xd1\xa0\xb3\xdc\x89\xae\xf6\x4f\x59\xd3\xc4\x69\x88\x36\x90\x16\x87\x96\xf6" +
	"\xa8\x43\x03\x29\x15\xa1\x83\x3f\x76\x21\x8a\xc7\xdb\x6b\xb3\x2c\x53\xb4\xbf\x35\x9d\xde\x09\xe6\x4a\x14\x27\x77" +
	"\x29\x8a\x50\xdf\x34\x6c\x7f\x8e\x0a\x3b\x57\xbe\x23\x74\x17\xb9\xb4\x8c\xc2\xc4\x31\x3a\xfb\x35\x24\xbe\x37\x3f" +
	"\xfe\xc5\x41\x3d\xef\x5a\x1d\xe9\x77\xb4\xa9\x81\x42\x2e\xe2\x0e\xff\x9b\x82\x96\x2a\x6c\xb5\x48\x35\xa5\xfc\x5a" +
	"\xaa\xa3\x0a\x6e\xa4\xda\x3f\x11\x48\x24\x0a\x16\x62\x07\x88\xd6\xb5\x54\xe2\xc8\xf5\xb8\xab\x4a\x14\xdf\x23\x2e" +
	"\xe5\x98\xbd\x62\x77\xb8\x93\x59\x02\xeb\xc8\x1f\x8f\xff\x88\x9f\x01\x00\x00\xff\xff\x06\x1d\xa2\x88\xaa\x04\x00" +
	"\x00"

func bindataPostgresprocgotplBytes() ([]byte, error) {
	return bindataRead(
		_bindataPostgresprocgotpl,
		"postgres.proc.go.tpl",
	)
}



func bindataPostgresprocgotpl() (*asset, error) {
	bytes, err := bindataPostgresprocgotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "postgres.proc.go.tpl",
		size: 0,
		md5checksum: "",
		mode: os.FileMode(0),
		modTime: time.Unix(0, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataPostgresquerygotpl = "" +
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x54\x5f\x6b\xdb\x3e\x14\x7d\xb6\x3e\xc5\xfd\x89\x50\xec\xdf\x5c\xe7" +
	"\xbd\xe0\x97\x75\x0c\x06\xa3\xd9\xbf\x87\x42\x29\x4c\x8d\xe5\x4c\xa0\x48\xb1\x24\x77\x0d\x42\xdf\x7d\x5c\xc9\x76" +
	"\xec\xa6\x65\xa3\x6f\xca\xf1\xfd\x73\xce\x3d\xf7\xc6\xfb\x4b\x58\xd9\x5f\xda\x38\xb8\xaa\x21\x8f\x2f\xc5\xf6\x1c" +
	"\xaa\x1f\xc7\x03\xaf\x6e\xf0\x49\xb9\x31\x14\xa8\xed\xa4\x75\xf8\x68\x1e\x28\xd0\x8e\x02\x35\xdc\x52\xa0\xb7\x9b" +
	"\xcf\x7a\x47\xa1\xfa\xda\x73\x73\xfc\xc2\x0c\xdb\xdb\x02\x2e\x43\x20\xb1\x76\x87\xe8\xb5\xde\xef\xb9\x72\x16\x7b" +
	"\xa4\xb8\x09\x19\x03\x45\x0b\xd5\x00\x46\x6c\xbd\x06\xef\x4f\xd0\x10\xc5\xa5\xe5\xf3\xcf\x91\x5f\x08\x60\x7a\x65" +
	"\x81\xc1\xb6\xb7\x4e\xef\x21\xf6\x2c\xc1\x70\xd7\x1b\x25\xd4\x0e\x0c\xb7\xbd\x74\x16\x98\x8d\x59\x27\x69\x21\x54" +
	"\xa9\xae\x6a\xb0\x45\xdb\xab\xed\xa2\x6e\xde\x3c\xc0\xed\xe6\xc3\x7b\xef\xc1\x30\xb5\xe3\x0b\x95\x10\x42\xb9\x88" +
	"\x1e\x6b\x43\x08\xde\x0f\x35\x0b\xc8\xbd\x47\x75\x4a\x3b\xa8\x36\x4a\x1e\x37\x0a\x03\xee\xee\xa7\x90\xff\x9f\x73" +
	"\x2a\x81\x1b\xa3\x4d\x01\x9e\x64\x8f\xcc\xe0\xaf\x84\x10\x92\xad\xd7\x60\x3b\x99\x24\x92\x2c\x95\xae\x3e\x29\xc7" +
	"\xcd\x41\x4b\xe6\x30\xfd\x91\x19\xac\x8d\xa3\x0a\x61\xab\x95\x75\x53\x2b\x48\x26\x42\x0d\x93\xa2\x95\x28\x61\x25" +
	"\x4f\xce\x24\xf2\xa2\x85\x95\xc0\x84\x77\x53\x6e\x42\x73\xa1\x1a\xfe\xf4\xdc\xd7\x95\x28\x30\x38\xb9\xf2\x4a\xc4" +
	"\x7c\x2a\xb3\x0e\x28\x02\xc1\xcb\x10\x7e\x7a\x8f\x54\xd2\x63\xb0\x24\x2a\x36\xbd\x1a\x15\xc7\x6d\xcb\x93\x8c\xd7" +
	"\x5c\x99\x0d\x7c\x39\x99\x85\x5d\x73\x32\x83\x57\xd3\x26\x9e\x7c\x4a\x0e\x20\xb1\x74\x25\x33\x9b\xc7\x42\x24\x43" +
	"\x83\x6a\x68\x1e\x12\x8f\x6f\xfa\xf7\x5f\x08\xbe\xcc\xa3\xa8\xbe\x6f\x99\xc2\x75\x69\x05\x97\x0d\x9e\xa1\x1d\x3a" +
	"\x7d\x44\xc0\x42\x7e\x30\x42\x39\xa0\x17\x74\xa0\x53\x44\xd6\x99\x68\xe3\x8e\xfc\x57\x83\x12\x12\xb7\x26\x4b\xbb" +
	"\x8f\x3f\xe3\x32\x91\x0c\x27\x39\x80\x17\x73\x35\x25\xc6\x9c\x6e\x0b\xd5\x74\x31\x05\x37\x62\x54\xf4\x36\x39\xff" +
	"\xc8\x2b\x6b\x78\xcb\x0d\x74\xd5\xb5\xd4\x96\xe7\x45\xb2\x5c\x6a\xd6\x8c\x77\x8b\xcc\xe3\x7f\xc7\xdd\xfd\xd9\xad" +
	"\xf8\x40\xb2\x56\x63\xfa\x0d\x7f\x72\x79\xbc\x99\x6c\x61\xd7\x55\x7d\xe6\x98\xc7\x69\xc4\x53\xda\x32\x45\xb2\xc1" +
	"\xbf\xee\xcd\xf3\x7f\x41\xe8\xb9\xd2\x68\x41\x54\x52\x03\x3b\x1c\xb8\x6a\x72\xc3\x6d\xb9\xb4\xa3\x58\x38\x15\xbf" +
	"\x4f\xfe\xa4\x83\x08\x84\xfc\x09\x00\x00\xff\xff\xb6\xe2\x6b\x23\xb5\x05\x00\x00"

func bindataPostgresquerygotplBytes() ([]byte, error) {
	return bindataRead(
		_bindataPostgresquerygotpl,
		"postgres.query.go.tpl",
	)
}



func bindataPostgresquerygotpl() (*asset, error) {
	bytes, err := bindataPostgresquerygotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "postgres.query.go.tpl",
		size: 0,
		md5checksum: "",
		mode: os.FileMode(0),
		modTime: time.Unix(0, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataPostgresquerytypegotpl = "" +
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8e\xb1\x0e\x82\x30\x10\x86\x67\x79\x8a\x7f\x30\x41\x07\xca\x6e\xe2" +
	"\x64\xe2\xe8\x22\x2f\x50\xe1\x50\x92\xb6\x90\xb6\xc4\x98\xcb\xbd\xbb\x29\xa0\xe2\xd0\x6b\x72\xff\xd7\xaf\x3f\x73" +
	"\x81\x6d\xd4\x37\x43\x38\x1c\xb1\x0b\xf5\x83\xac\x86\xba\x2e\x77\x95\x92\x79\x5e\xb4\xa5\x3d\x0a\x91\x2c\xbd\xe9" +
	"\x5a\xa8\x53\x6f\x2d\xb9\x38\xed\xca\x12\xcc\xbf\xd5\x42\x91\x09\xb4\x8e\x93\x03\x22\xf0\x34\x78\x0a\xe4\x62\x80" +
	"\x86\xef\x9f\x68\x7d\x6f\x91\x33\x7f\xba\x88\xe4\x6a\x36\xb8\x26\xc9\xe2\x6b\xa0\x3f\x43\x88\x7e\xac\x23\x78\x82" +
	"\xbc\x76\x77\x82\x3a\x77\x64\x9a\x90\xf0\xcd\x1a\x65\x86\xa7\x49\xa0\xaa\x34\x45\xf0\x6d\x6b\xd2\x19\xad\x5b\xd8" +
	"\xf5\x97\x92\x65\xef\x00\x00\x00\xff\xff\x5b\x7f\x83\xf0\x1d\x01\x00\x00"

func bindataPostgresquerytypegotplBytes() ([]byte, error) {
	return bindataRead(
		_bindataPostgresquerytypegotpl,
		"postgres.querytype.go.tpl",
	)
}



func bindataPostgresquerytypegotpl() (*asset, error) {
	bytes, err := bindataPostgresquerytypegotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "postgres.querytype.go.tpl",
		size: 0,
		md5checksum: "",
		mode: os.FileMode(0),
		modTime: time.Unix(0, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataPostgrestuplegotpl = "" +
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\xa8\xae\x56\xd0\xf3\x4b\xcc\x4d\x55\xa8\xad\x55\x28" +
	"\x2e\x29\x2a\x4d\x2e\x51\xa8\xe6\xaa\xae\xd6\x55\x28\x4a\xcc\x4b\x4f\x55\xd0\x73\xcb\x4c\xcd\x49\x29\x56\xa8\xad" +
	"\xe5\xe2\x44\x56\x5a\x5d\xad\x50\x94\x0a\x36\x40\x2f\x04\x44\xd6\xd6\x2a\x24\xa4\x24\x59\x29\x81\xd4\x80\xb5\x40" +
	"\x15\x2a\x25\x28\xe8\xeb\x2b\xa0\x8b\x82\x2d\x48\xcd\x4b\x01\x31\x6b\xb9\x00\x01\x00\x00\xff\xff\x12\x3f\xd6\x00" +
	"\x88\x00\x00\x00"

func bindataPostgrestuplegotplBytes() ([]byte, error) {
	return bindataRead(
		_bindataPostgrestuplegotpl,
		"postgres.tuple.go.tpl",
	)
}



func bindataPostgrestuplegotpl() (*asset, error) {
	bytes, err := bindataPostgrestuplegotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "postgres.tuple.go.tpl",
		size: 0,
		md5checksum: "",
		mode: os.FileMode(0),
		modTime: time.Unix(0, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataPostgrestypegotpl = "" +
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x58\x4d\x73\xda\x48\x13\x3e\x4b\xbf\xa2\xa3\x4a\xbd\x81\x37\x8e\xa8" +
	"\x1c\xf6\xb0\xae\xe2\x90\x35\xca\xae\x6b\x1d\x70\x0c\xec\xe6\x16\x06\xd4\x18\xad\xa5\x19\x3c\x33\xb2\x4d\xa9\xf4" +
	"\xdf\xb7\x66\x46\x12\xfa\x02\x84\x93\xda\x83\xc1\x68\xba\x7b\x7a\xfa\xe3\xe9\x67\x94\x24\x1f\xe0\xad\xd8\x30\x2e" +
	"\xe1\x72\x08\x3d\xfd\x1f\x25\x11\x82\x3b\x56\x9f\x0e\x72\xee\x80\xc3\x51\x38\xe0\x88\xc7\x50\x48\xf5\xd3\x5f\x3a" +
	"\xe0\x7c\x9b\xdc\xb0\x7b\xa7\x0f\x1f\xd2\xd4\xd6\x56\x24\x59\x86\x68\xac\xac\x36\x18\x11\x70\xa7\xd9\xf7\x4c\xad" +
	"\x98\x4f\x65\x75\xaf\x13\xac\xc1\xbd\x62\x51\x84\x54\xea\x67\x83\x01\x24\xc9\xfe\x51\x26\x85\xa1\xc0\xf2\xb2\xf6" +
	"\x2c\x4d\x81\xe3\x96\xa3\x40\x2a\x05\x10\xe0\xec\x19\xd6\x9c\x45\xf0\x2e\x49\x72\x5f\xd2\xf4\x9d\x6b\x2c\x50\x5f" +
	"\x19\x93\xbb\x2d\x56\x2c\x08\xc9\xe3\x95\x84\x44\x0b\x71\x42\xef\x11\xdc\xcf\x01\x86\xbe\x50\xe2\x56\x59\x34\x49" +
	"\x80\xa3\x36\xe0\xce\xd4\x67\x9a\xc2\xe2\x1f\xc1\xe8\xa5\x63\x3c\x0e\xd5\x5f\x1c\xd1\x4c\xde\x59\x40\x71\x98\xda" +
	"\x52\xd9\xa3\x3c\x08\xb7\x3c\x88\x08\xdf\xfd\x89\x3b\xf5\xd4\xb6\x06\x03\x78\x61\xb0\xd6\xae\xd8\xd6\x77\x7c\x09" +
	"\x84\x14\x17\xf0\xdd\xc7\x10\x25\xfa\xb0\x64\x2c\xb4\x93\x24\x37\x93\xda\xea\x47\xd3\xd0\x60\x00\x9e\x56\x05\x1f" +
	"\x25\xf2\x28\xa0\x28\x94\x98\xdc\x54\xe3\x60\xec\x43\x40\xf5\x8a\x4f\x24\x59\x12\x81\xae\xbd\x8e\xe9\x0a\x7a\x2a" +
	"\xa0\xa6\x44\xd2\x14\xfe\x5f\xd2\xeb\x67\xd6\x7b\x7d\xed\x10\x24\xb6\xc5\x51\xc6\x9c\x42\x59\xc5\xcd\xdc\x57\x5e" +
	"\x0e\x06\x30\xca\x8e\xb0\xe5\xec\x29\xf0\x95\x3f\x74\xcd\x78\x44\x64\xc0\x68\x9b\x6f\x1b\x22\x60\x89\x48\x21\x3f" +
	"\xbb\xce\xf2\x99\x7e\x66\x9b\x9e\x72\x34\xdb\x22\xf3\xf4\x9a\x0a\xe4\x12\x02\xfd\x25\x1a\x8e\x49\x76\xae\x17\xc6" +
	"\x60\xcf\x5f\xc2\xb7\xc9\xe8\xb7\x3e\x20\xe7\x8c\x2b\x67\x9e\x08\x57\x3f\xcc\x03\x93\xfe\x60\x0d\x24\xe4\x48\xfc" +
	"\x9d\xc9\xce\x05\x2c\x49\x10\xda\x56\xb0\x6e\x0d\xae\xb2\x92\x9f\x49\x5b\x11\xee\x18\x9f\x7b\x8e\x71\x1e\xd6\x24" +
	"\x08\xd1\xbf\xac\x9a\x14\x4e\xdf\xb6\xf6\xa5\x63\xba\xf4\x0b\xa1\x31\x09\x6f\x1f\x74\x03\x0c\x06\x20\x1e\xc3\x2c" +
	"\x02\xf0\x18\x23\xdf\x5d\xc0\xd6\x94\x18\x3c\xe0\x0e\xa2\x58\x48\x58\x62\x9e\x4c\xdf\xb6\x56\x8c\x0a\x09\x06\x2a" +
	"\x60\x08\x8b\xeb\xf1\xd4\xbb\x9b\xc1\xf5\x78\x36\x81\x72\x67\x42\x6f\x01\xef\x6d\xcb\x5a\x24\x09\xac\x58\xa8\x30" +
	"\x47\x94\x9a\x2f\x5b\xec\xc3\x5f\x9f\x6e\xe6\xde\xb4\x26\xfd\x44\xc2\x36\xe1\x85\x09\x1d\x8f\xa9\xf1\xd5\xb6\x34" +
	"\x48\xf5\x8c\x37\x17\x6a\x7f\xdd\x52\xd5\xcd\x8a\x58\xf6\x6d\x4b\x25\x61\x08\xfe\xd2\xfd\xaa\xf4\xef\xd8\x73\x67" +
	"\x5d\x77\xba\x22\xb4\xf7\xbf\x4a\x6e\x54\xf2\xf7\x0d\x59\xd4\x81\x4e\xa2\xda\xe9\xcd\x10\x68\x10\xd6\x52\xa7\x52" +
	"\xa2\x3a\x5b\x81\x5e\xa7\x1c\xe4\xb1\x87\xe5\x0e\x04\x3e\xc6\x48\x57\xf8\x93\xf2\xd0\xe2\xfd\x19\x89\x39\xa6\x7d" +
	"\xe7\xcd\xe6\x77\xe3\xeb\xf1\xef\xb0\xdf\xb7\xa2\x70\xc5\x42\x25\xff\x23\x19\x6d\x8f\xfd\x2b\x53\xdc\x66\xec\xa7" +
	"\xe7\xdc\xa0\xb9\xc9\x39\x4a\xd3\xa5\x26\x9d\xad\x3d\x3f\x04\xc9\x63\xb4\x0b\x30\xa3\x41\xb8\x1f\x05\x14\xa1\xb7" +
	"\x3f\x4e\x14\x87\x32\x38\x72\x26\xb3\xd0\x07\xc7\xc9\x8b\x6e\xbe\xf5\x89\x44\x88\xf5\x57\x13\xfd\x1a\xb3\xc2\x3a" +
	"\x09\x7f\xc6\x62\x0b\xfc\x35\xf0\x2f\x03\x40\x9f\xa1\xa0\xef\x64\x15\x00\x55\x20\xdf\x1c\x84\xc0\x36\x0c\x34\x47" +
	"\x28\x30\x50\x59\x05\xca\x32\xb3\x0a\x03\x15\x08\x16\x7b\x9a\x11\x50\xde\xad\x75\x46\x74\xdd\x2d\x22\xfc\x41\x0d" +
	"\x2d\xc6\x8d\xe9\x80\xd1\xfd\x96\x26\x53\xf7\x12\x7a\x10\x22\x6d\xe6\x03\xfa\xf0\x51\xe7\xc3\xca\x61\x40\x77\x01" +
	"\x3c\x07\x72\x03\x2b\x16\x6d\x99\x08\x24\x96\xd1\x40\x89\xd6\x5b\x7f\x7e\x3b\xfa\x34\xf3\xaa\x5d\x3f\xf5\x66\x79" +
	"\xeb\x56\x7b\xbf\x5a\x28\x4d\x8f\xf2\x1e\xd6\x5d\x3c\x84\x1e\xd4\x8c\x28\x04\x38\xcb\xc6\xdf\x7f\x78\x77\x5e\x09" +
	"\x05\x84\x3e\x62\x66\xa2\xa1\xea\xc0\xa7\xf1\x08\x1c\xe8\xdd\xa3\x14\x92\x70\xb9\x62\x31\x95\x87\xf7\xea\xeb\x6e" +
	"\x33\x40\x62\xd5\xa0\xc4\x3a\x06\x26\xdd\x1a\x06\xd2\xb4\x81\x1b\x0d\x19\xa3\xac\x41\xc0\xb2\xbe\x5f\x40\x81\x40" +
	"\xde\x0b\xae\xfe\xd3\xdd\xcb\x88\x63\x15\xd4\xba\x59\x60\x3f\x5c\x45\xa7\x67\xc0\xb1\xfa\xe9\xa8\x5d\xaf\x9c\x96" +
	"\xf9\x01\x43\x78\x6b\x04\x0e\xd6\x49\x61\xf8\xdc\x0a\xe9\x30\x21\x2e\xa0\xdb\x6c\xe8\x5a\x16\x3f\x75\xcb\x66\x31" +
	"\x98\xf1\x63\x65\x13\x68\x4a\x9e\x10\x04\x79\x6a\x01\xff\x06\xf5\x3d\x0d\xfe\xca\x5a\x1b\xf4\xd7\xf1\xb5\xb8\x51" +
	"\x94\xf1\xb5\x22\x51\x8c\x91\x02\x46\xdb\xa4\x0a\xae\xdd\x2f\x0e\x34\xdf\x6a\x0a\xb5\x45\xae\x2e\x1c\x02\x08\x85" +
	"\xd8\x3c\x52\xe8\x5c\xf2\xd6\x55\xe2\x5a\x65\x3c\x99\x79\x97\x70\xcb\x84\xbc\xe7\x38\xfd\x7a\x03\xbf\xba\xbf\xbc" +
	"\x07\x46\xc3\x5d\xa7\x79\x77\x80\xee\x1f\x9a\x77\xad\x84\xff\x28\xe3\x7f\x15\xe5\x2f\xc6\x5d\xb9\xdf\xcf\xe5\x8b" +
	"\xc7\x89\x7b\x93\x20\x1e\xa5\xee\x4a\x7c\x32\x86\xab\xc9\xf8\xf3\xcd\xf5\xd5\x4c\x47\x75\x6f\xbb\x05\xf6\xfa\x30" +
	"\x9a\x40\x86\x48\x65\x10\x3a\xe9\xd4\xb0\x2e\xba\xe5\xb8\x0e\x5e\xaa\x0a\x8e\xf7\xed\xea\x66\x3e\xf2\x46\x4e\x59" +
	"\x77\x91\x45\xad\x8c\x0f\xe7\xde\x2f\x5e\xd9\xe7\x46\xb7\xc9\x20\xab\x3d\x5c\xa4\xb5\x4a\x1d\x4f\x70\xc7\x2a\x79" +
	"\xac\xdd\x3c\x32\x12\x28\x24\x91\x18\xe9\x37\x2d\x2c\x0a\xa4\xa2\x3f\x7e\x8c\x0a\x06\x42\xb2\x7a\x00\xb6\xce\x5e" +
	"\x55\x00\x93\x1b\xe4\x20\x37\x84\x56\x48\x49\x89\xdb\x16\x6f\x00\x32\xa6\xd5\x04\x97\xd7\xdf\xef\x3b\xdf\xac\x5b" +
	"\x89\xe5\x51\x5e\x59\x0a\x52\x0e\x27\x4d\xb2\x78\x94\x2b\xd6\x2d\x74\xe7\x7e\xdd\xa9\x5f\xbd\x89\x47\xde\x8d\x37" +
	"\xf3\xe0\xf3\xdd\xe4\x4b\xb5\x89\x0f\xb0\xae\x23\x84\x2b\x1b\x90\xe7\xd4\xff\x01\x2a\x72\x66\x27\x1c\xb7\x72\xba" +
	"\x27\x6a\x3c\xe7\x04\xec\x1d\x8c\x58\x47\xb2\xf1\xb1\x53\x94\xba\xcc\xe7\x63\xf1\xe9\xa2\xdf\x35\x32\xb5\x3b\x67" +
	"\xfe\x0a\xcc\x6a\xaf\xe4\xf6\x2b\x67\xd9\xd0\xbf\x01\x00\x00\xff\xff\xc2\xe8\x45\xb3\x52\x16\x00\x00"

func bindataPostgrestypegotplBytes() ([]byte, error) {
	return bindataRead(
		_bindataPostgrestypegotpl,
		"postgres.type.go.tpl",
	)
}



func bindataPostgrestypegotpl() (*asset, error) {
	bytes, err := bindataPostgrestypegotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "postgres.type.go.tpl",
		size: 0,
		md5checksum: "",
		mode: os.FileMode(0),
		modTime: time.Unix(0, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataXodbgotpl = "" +
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x94\xdf\x6f\x9b\x48\x10\xc7\x9f\xcd\x5f\x31\x41\x77\x2a\xa4\x14\x2e" +
	"\xf7\x18\xc9\x0f\xd7\x6b\x5f\x4e\xf7\xd3\xa9\x4e\x27\x19\x9f\xbc\x86\xc1\x5e\x05\x76\xc9\xce\x82\x63\x59\xfe\xdf" +
	"\xab\xd9\x05\x17\xd2\xa6\x2f\x89\x59\xe6\x3b\x3f\x3e\xdf\x61\xb3\x0c\xfe\xfb\xeb\xc3\x7b\x90\x04\xf6\x80\x50\xe8" +
	"\xa6\xd1\x0a\xa4\xb2\x68\x2a\x51\x20\x54\xda\x40\x29\xac\xd8\x09\x42\xd0\x2d\x1a\x61\xa5\x56\x1c\x2c\x2c\x14\x42" +
	"\xc1\x0e\xa1\x23\x2c\xe1\x28\xed\x21\xc8\x32\xb0\xa7\x16\x09\x2a\xa3\x1b\xa0\xe2\x80\x8d\x80\x37\xe7\xf3\xf8\x33" +
	"\x7d\xf0\xff\x2f\x97\x37\x69\x90\x65\x1c\xff\xe9\x20\x09\xe8\xa0\xbb\xba\x84\xa3\x36\x8f\x2e\xd1\xb5\x64\x46\x4f" +
	"\x75\xfa\xe1\x3d\x08\x55\xce\xcf\x3e\x3d\xa7\x01\x97\x1a\xba\xbf\xf6\x7b\x0e\x16\x1f\x9f\xb1\x88\xc8\x1a\xa9\xf6" +
	"\x09\xa4\x69\x7a\x7d\x79\xbe\xc4\x10\xb1\x78\x85\xd4\xd5\x36\x01\x34\x46\x9b\x38\x58\xfc\xd3\xa1\x39\xbd\x2e\xb9" +
	"\x75\x1a\x7d\xa4\x17\x8a\x95\x3e\xbe\x2a\x1a\x35\xc1\x25\x08\x1c\xe3\xdf\xf5\x1e\x5a\xa3\x7b\x59\xa2\x47\x5d\xeb" +
	"\x3d\x54\x9d\x2a\x3c\xbe\xdd\x09\xf6\xa8\x18\x2f\x96\xf0\xd4\xa1\x91\x48\x69\xd0\x0b\x33\x48\x97\x2e\xf6\xd5\x72" +
	"\x67\xf0\x75\x1e\x0a\xa1\x14\x9a\x7f\x45\xdd\xa1\xf9\xae\xa9\xde\x27\x67\xa3\x6c\xda\x1a\x1b\x54\x16\x76\xda\x1e" +
	"\x58\xc2\xa9\x66\xb8\x87\xbc\xce\x07\x7a\xaa\xb3\xd2\xc8\x1e\x4d\x3a\xd6\x19\x33\xd3\x60\xca\x8b\x36\xa6\xee\x4c" +
	"\xb2\x05\x8b\x59\x9a\x01\xd5\x83\x1b\xf1\xa1\x96\x05\xf2\x00\x02\xc8\xfd\xd4\x15\xf8\xe1\xaf\x35\x26\x71\xeb\x8d" +
	"\x7f\xe7\x12\x3c\x75\xda\xe2\x47\x2a\x44\x8b\x2b\xdc\xe3\xf3\x88\xc1\xb8\x07\xab\xa1\x11\xb6\x38\x00\xba\x88\x12" +
	"\x8a\x83\x30\xa2\xb0\x68\x08\xa4\xe2\x72\x2e\x93\x67\xff\x55\xaa\xa5\xcf\xd2\xa6\x7f\x74\x64\x7f\xd5\x4d\x2b\x6b" +
	"\x8c\xb6\xd1\xfa\xff\x3c\xdf\x44\xeb\x3c\xdf\x9c\x7f\xbe\xc4\xb7\x71\x9e\x87\xdb\xf8\x6a\x08\x90\xb0\x92\x2a\x39" +
	"\x18\x3f\xe5\x39\xf7\x64\x32\x52\x1a\xb8\xdd\x88\x88\xe0\x76\x72\x1c\xbb\x84\x11\x99\x02\x66\xfe\xbb\xbd\x64\xbc" +
	"\xbb\xae\x4a\x40\x3f\xc2\xfd\x12\xc8\x14\x69\xb4\xde\xec\x4e\x16\xe3\x60\x21\x2b\xb8\xd1\x8f\x1c\xb2\x30\x68\x3b" +
	"\xa3\xbc\x86\xd2\x3f\xf1\x18\x85\x52\xf5\xa2\x96\xe5\xb4\x83\x30\x0e\x16\x97\x20\x58\x64\x19\x23\x52\x7b\xf4\x34" +
	"\x06\x6e\xe4\x1a\x2e\xa8\x87\x56\x18\x62\x2f\xc9\x1a\xae\xfa\x12\x59\xba\xc2\xb6\x16\x05\xfe\x52\xd7\x3e\xf9\xb0" +
	"\xc3\xd1\xae\xab\xe2\x04\xb6\x3f\xdc\x85\xcc\xca\xc9\x97\x57\x8b\x07\x11\xc7\x26\xb0\xcd\xf3\x2d\xff\xdd\x26\xf0" +
	"\xee\x2e\xf6\x2d\x19\x6c\x74\x8f\xb0\x33\xbc\x75\x13\xf5\xfa\xee\xbe\x46\xc5\xba\xf8\xdd\xdd\xc6\xc7\xee\x84\xac" +
	"\x41\x56\xa0\x55\x7d\x02\xad\xd0\xc1\x18\xa3\x60\xb9\x84\x9f\x1c\x96\x5b\x22\x58\x4e\x09\x44\xe3\x5a\x9d\x2f\xf1" +
	"\x17\x6c\x4a\xd6\x57\x30\x6e\x76\x7f\x63\x31\x0a\x83\xa2\x64\x14\x85\x23\x51\x50\xcf\x70\x57\xee\x30\x1a\x27\x9b" +
	"\x9d\xc4\x3c\x38\x97\x72\x37\x8b\x13\x99\x94\x5f\x47\xde\x31\x3e\xbc\x59\x72\x49\xd7\x61\xd5\xd8\xf4\x6f\x23\x95" +
	"\xad\xa2\x10\x9f\xa5\x95\x6a\x7f\x73\x0f\x3f\xf6\xb9\x0a\x5d\x82\x78\x66\xae\xef\xf2\xeb\xa9\x5c\x41\xc6\x38\x19" +
	"\xc8\x7f\x7a\xee\x3b\x7c\xb1\xad\xaf\x7c\xe9\xdf\xd9\xd7\xd9\xba\x3a\x5d\x14\x43\x34\xcd\x33\xde\xa3\x3c\x54\xcf" +
	"\x53\x37\xe2\xf1\x0b\xed\xc4\x7b\x43\x0c\x87\xab\xc8\x04\x88\x83\x8c\x5b\x42\x22\x87\xa2\x5f\xcb\x0d\x2c\x61\x1b" +
	"\x6e\xe1\xed\xb7\xb6\x66\xfe\x3c\x6c\xcf\x36\xcf\x87\x25\x4a\x58\xc9\x07\xa1\x7f\x86\xb7\x7c\xc0\xc4\x46\x2a\xe1" +
	"\x39\x9c\x64\xfe\x4d\x4b\x15\xf5\x09\x84\x49\xc8\xb1\xe1\x25\x4c\x26\xdc\xbe\x75\x59\xcd\xae\xc0\xeb\x9d\x35\xdc" +
	"\x56\xb3\x97\x41\xf0\x39\x00\x00\xff\xff\x7a\x63\xe4\xe0\x85\x07\x00\x00"

func bindataXodbgotplBytes() ([]byte, error) {
	return bindataRead(
		_bindataXodbgotpl,
		"xo_db.go.tpl",
	)
}



func bindataXodbgotpl() (*asset, error) {
	bytes, err := bindataXodbgotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "xo_db.go.tpl",
		size: 0,
		md5checksum: "",
		mode: os.FileMode(0),
		modTime: time.Unix(0, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataXopackagegotpl = "" +
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8c\x31\x6e\xf3\x30\x0c\x46\xe7\x9f\xa7\x20\xbc\xe4\xef\x22\x1e\xa2" +
	"\xe9\xd0\xa5\x29\xd0\x5c\x40\x96\x18\x59\x68\x24\xba\x14\x1d\x24\x30\x7c\xf7\x42\x29\xb2\x14\x9d\xde\x7b\x24\xf0" +
	"\x11\xe1\xbb\x0f\x9f\x3e\x31\xae\x2b\xba\x87\x6f\x1b\x06\xa9\xe6\x73\x6d\x68\x13\xa3\xdd\x66\x6e\x78\x12\xc5\x16" +
	"\x26\x2e\x1e\x77\xeb\xfa\x50\xf7\xf1\xc3\x6d\xdb\x39\x98\xff\x1c\x03\x20\xc2\x67\x89\x8c\x89\x2b\xab\x37\x8e\x38" +
	"\xde\xf0\x2a\x0e\xf7\x07\x7c\x3b\x1c\xf1\x65\xff\x7a\x74\x00\xb9\xcc\xa2\x86\xff\xe1\xdf\x10\xbd\xf9\xd1\x37\xa6" +
	"\xf6\x75\x1e\x7e\x35\x45\xcd\x17\xd6\x7e\xe6\x1a\x24\xe6\x9a\x28\xb4\xcb\xbd\x55\x45\x5b\xb7\x53\xb1\x0e\xe5\xc4" +
	"\xd7\xb9\x5b\x33\xcd\x35\xdd\x7f\x96\x0b\x77\xa6\x6c\xd3\x32\xba\x20\x85\x92\x48\x3a\x33\x2d\x4b\x8e\x03\x3c\x01" +
	"\x7c\x07\x00\x00\xff\xff\x45\x78\x64\x2c\x1a\x01\x00\x00"

func bindataXopackagegotplBytes() ([]byte, error) {
	return bindataRead(
		_bindataXopackagegotpl,
		"xo_package.go.tpl",
	)
}



func bindataXopackagegotpl() (*asset, error) {
	bytes, err := bindataXopackagegotplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "xo_package.go.tpl",
		size: 0,
		md5checksum: "",
		mode: os.FileMode(0),
		modTime: time.Unix(0, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"postgres.enum.go.tpl":       bindataPostgresenumgotpl,
	"postgres.foreignkey.go.tpl": bindataPostgresforeignkeygotpl,
	"postgres.index.go.tpl":      bindataPostgresindexgotpl,
	"postgres.proc.go.tpl":       bindataPostgresprocgotpl,
	"postgres.query.go.tpl":      bindataPostgresquerygotpl,
	"postgres.querytype.go.tpl":  bindataPostgresquerytypegotpl,
	"postgres.tuple.go.tpl":      bindataPostgrestuplegotpl,
	"postgres.type.go.tpl":       bindataPostgrestypegotpl,
	"xo_db.go.tpl":               bindataXodbgotpl,
	"xo_package.go.tpl":          bindataXopackagegotpl,
}

//
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}


type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"postgres.enum.go.tpl": {Func: bindataPostgresenumgotpl, Children: map[string]*bintree{}},
	"postgres.foreignkey.go.tpl": {Func: bindataPostgresforeignkeygotpl, Children: map[string]*bintree{}},
	"postgres.index.go.tpl": {Func: bindataPostgresindexgotpl, Children: map[string]*bintree{}},
	"postgres.proc.go.tpl": {Func: bindataPostgresprocgotpl, Children: map[string]*bintree{}},
	"postgres.query.go.tpl": {Func: bindataPostgresquerygotpl, Children: map[string]*bintree{}},
	"postgres.querytype.go.tpl": {Func: bindataPostgresquerytypegotpl, Children: map[string]*bintree{}},
	"postgres.tuple.go.tpl": {Func: bindataPostgrestuplegotpl, Children: map[string]*bintree{}},
	"postgres.type.go.tpl": {Func: bindataPostgrestypegotpl, Children: map[string]*bintree{}},
	"xo_db.go.tpl": {Func: bindataXodbgotpl, Children: map[string]*bintree{}},
	"xo_package.go.tpl": {Func: bindataXopackagegotpl, Children: map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
