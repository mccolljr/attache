// Code generated by go-bindata. DO NOT EDIT.
// sources:
// templates/model.tpl
// templates/routes.json.tpl
// templates/routes.tpl
// templates/view_create.tpl
// templates/view_list.tpl
// templates/view_update.tpl

package cmd_gen

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
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
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesModelTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xc4\x94\x41\x8b\xdb\x3c\x10\x86\xcf\xd6\xaf\x98\xcf\xf0\x15\xbb\x64\x95\x7b\x4b\x0e\xa5\x4b\x21\xa4\xe4\xd0\x6c\x7b\x29\x3d\x28\xf2\x38\xab\x56\x96\x1d\x69\xdc\xc5\x08\xfd\xf7\x62\xd9\xf5\x3a\xcb\x66\xd9\x86\xa5\xbd\x85\xc9\xcc\xbc\x0f\xcf\xd8\x6e\x84\xfc\x21\x0e\x08\x55\x5d\xa0\x76\x8c\xa9\xaa\xa9\x2d\x41\xc6\x92\xb4\x10\x24\xf6\xc2\xe1\xd2\x1d\x75\xca\x92\xf4\xa0\xe8\xb6\xdd\x73\x59\x57\xcb\x4a\xca\x5a\xeb\xef\x76\x29\x88\x84\xbc\xc5\x94\xe5\x8c\x51\xd7\x20\x78\xcf\xb7\xa2\xc2\x10\xc0\x91\x6d\x25\x81\x07\xef\xad\x30\x07\x04\xfe\x41\xa1\x2e\x5c\x08\xde\xf3\x5d\xfc\x33\x16\x42\xe8\x87\x6e\xba\x06\x43\x78\xeb\x3d\x9a\xbe\x12\x18\x2b\x5b\x23\x61\x8b\x77\xd3\xc6\x2c\x87\x31\x8e\xef\xa8\xb6\x62\xaf\x11\x3c\x58\xa4\xd6\x1a\x30\x78\x97\x4d\x9d\xf9\x34\x9f\x55\xf0\x7a\x56\xbe\xe9\x87\xb2\xbc\x67\x53\xe6\x70\x3f\xed\x7d\x63\x95\xa1\x12\xd2\xff\x8f\x29\xf0\xd8\x36\xa3\x78\xb0\x65\x6d\x1c\x5a\xca\x72\xc8\x64\xad\xdb\xca\x38\xf8\xfa\x6d\xd8\xb8\x80\x9f\x42\xb7\xd8\x17\x94\x21\xb4\xa5\x90\xe8\x43\x0e\x9e\x25\xbf\x5b\x57\x53\xb3\x67\x49\xe2\xfd\x15\x9c\xd8\x81\xab\x10\x62\x1d\x54\x09\xa6\x26\xe0\xdb\x7a\xc8\x83\x5e\x1c\x9c\x70\xbe\x8f\x3b\x21\x84\xc5\xe8\x6d\xdc\x88\xa6\x18\xf6\x04\x96\x8c\x40\xab\x53\xa4\x3f\xcf\xae\xf8\xc3\xb3\x3d\x11\x3a\x68\x65\x81\x31\xef\x55\x09\xfc\x1a\x4b\xd1\x6a\xda\x60\x17\x5b\x1e\x93\xfa\xae\x24\xb4\xa3\x59\x8b\xae\xd5\x04\xee\xa8\xf9\xa7\xf8\x33\x1a\x54\xc5\x02\xd0\x5a\x78\xb3\x82\xa1\x81\x7f\x14\x8e\x86\x91\x75\x91\xe5\x2c\x51\x65\x6c\xf8\x6f\x05\x46\xe9\x7e\x24\x69\x84\x51\x32\x43\x6b\xf3\xc8\x55\xf1\xf5\x35\xac\x40\x15\x2c\xb0\x91\x39\x9c\x39\xf2\xe7\xa6\x10\x84\x7f\xef\xc8\x43\xde\xbf\x39\xf2\x94\x7d\xe1\x91\x1f\xf3\xb7\x43\x8d\xf2\xcc\x4b\xa2\x0c\xd5\x2f\x6b\x6f\x48\xbb\xd0\x5e\xc4\xb9\xd4\xdd\x94\xfc\xea\x05\xe5\x6d\xb0\x1b\xc0\x5d\x96\x4f\x32\xee\x3f\x57\x33\x3d\xe7\x19\xfb\xf7\x6e\x83\xdd\xb3\x94\xcc\xe9\xfa\xcf\xde\x19\xa6\x2f\xf1\x31\x8b\x48\x33\x55\x73\xae\x53\x83\xcf\x80\xeb\xa5\xc1\xdc\xda\x53\x5c\xbf\x02\x00\x00\xff\xff\x22\x18\x74\xfb\xad\x06\x00\x00")

func templatesModelTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesModelTpl,
		"templates/model.tpl",
	)
}

func templatesModelTpl() (*asset, error) {
	bytes, err := templatesModelTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/model.tpl", size: 1709, mode: os.FileMode(420), modTime: time.Unix(1535980617, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf8, 0x6d, 0x99, 0x11, 0xc1, 0xfd, 0x44, 0x68, 0xab, 0x92, 0x7f, 0xdf, 0xa8, 0xb6, 0xc, 0xf8, 0xfe, 0xee, 0x3b, 0xe, 0xe, 0xad, 0x6e, 0x7b, 0x3b, 0xb6, 0x19, 0xc9, 0xf8, 0x12, 0x7a, 0x7c}}
	return a, nil
}

var _templatesRoutesJsonTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x95\xcf\x6f\xd3\x30\x14\xc7\xcf\xf6\x5f\x61\x72\x98\x6c\x34\x39\x13\xda\x09\xb4\x03\x5b\x53\x7e\x68\x74\xa8\xed\xe0\x88\xdc\xf8\xb5\xf5\x70\xec\xf4\xc5\x21\x54\x51\xfe\x77\x94\xa6\x2d\x3f\x5a\x51\x36\x0a\x07\xb8\xf5\xf5\xc9\x7e\x4f\x9f\xef\xc7\x4a\xae\xd2\x8f\x6a\x06\x2c\x53\xc6\x51\x6a\xb2\xdc\x63\x60\x9c\x92\x68\x9a\x85\x88\x92\xc8\x41\x88\xe7\x21\xe4\xed\x6f\xe3\x63\xe3\xcb\x60\x6c\x5b\x68\x15\xd4\x44\x15\x10\x17\x8b\x55\x0d\x2e\xf5\xda\xb8\x59\x7c\x57\x78\x17\x51\x4a\xa2\x99\x09\xf3\x72\x22\x53\x9f\xc5\x59\x9a\x7a\x6b\xef\x30\x56\x21\xa8\x74\x0e\x11\x15\x94\x4e\x4b\x97\x32\x9e\xb2\xc7\x75\x2d\xaf\xbc\x0b\xf0\x39\x8c\x97\x39\x34\x8d\x60\x2f\x92\xf1\x87\xba\x96\xa3\xd4\xe7\x70\xa5\x32\xb0\x4d\x53\xd7\xf2\x8d\xd7\x60\xe5\x40\x65\xd0\x34\xd7\xa6\x08\x5c\xb0\x9a\x92\x8a\x3d\xbd\x60\xa9\x1c\x42\x91\x7b\x57\xc0\x7b\x34\x01\x90\x0b\x4a\x94\xb5\xa7\x0c\x10\xbb\x7e\xef\x92\x0b\xf9\xdc\x5a\xde\x8e\xe5\x82\xad\x37\x91\xa3\xe0\x51\x4d\x2c\xd4\x0c\x21\x94\xe8\x98\x83\x8a\x67\xed\xa4\x42\xfe\x30\x53\xb0\x46\x50\x62\xa6\xab\x4b\x1f\x5d\x30\x67\x2c\x3b\x39\xd9\x54\xc5\xc2\xca\x04\x71\xe0\x87\xbe\x2a\xda\xc5\xc8\x66\x44\x82\xe8\xb1\xaf\x82\xb2\x1c\x10\x05\x25\x0d\xdd\xf6\x86\xe0\x34\xe0\xeb\xd1\xcd\x80\x57\xa7\x4c\x59\x2b\x68\xf3\x9b\x68\x0e\x60\xc1\x4d\x63\x51\x42\xcb\x90\x12\xa3\xdb\xbf\x50\xf6\x3d\x66\xef\x94\x2d\x81\x47\x46\x47\x82\x92\x4f\x0a\x59\x50\x38\x83\xc0\xf6\x13\xd9\xd2\xf8\x8a\xb8\x6f\x9c\xe6\x27\xdd\xa9\x53\x66\xb4\x78\xf6\x2d\xae\x16\xcb\xfa\xc8\xc5\x3e\x64\xdf\x33\xe3\xe7\x67\xe7\x82\x92\x96\xd7\x03\x60\x76\x3b\x1c\xe2\xf9\xf6\x66\x74\x08\xe8\x00\xaa\x7b\x33\x9d\x78\xbd\xdc\xda\xd7\xbd\x1a\x39\x04\xa5\x5b\x03\x51\x5e\x7a\xbd\xdc\x51\xe9\x90\x32\xf7\x4a\xa3\x7d\x85\xf2\xd6\x65\x0a\x8b\xb9\xb2\xbc\x5b\x67\x1d\xcb\x6e\x24\x3f\x1b\xbb\x93\xf0\x2b\x57\x00\x06\xfe\xa0\xcb\x2a\xb9\x82\xf6\x12\x94\x06\xe4\x4f\xce\xce\x8e\x10\xcf\xff\xe7\xfb\x1f\x90\xeb\xf8\xe6\x90\x3d\xb0\x6e\x73\xad\x02\xfc\x25\x75\x7a\xc9\x75\x32\x4e\xfe\x65\x79\xf6\x20\x21\x84\x74\xdf\xb1\x5f\xd0\x68\x67\x9d\x1e\x58\x38\x66\x3c\x5f\x02\x00\x00\xff\xff\xca\x78\x46\x77\x5d\x08\x00\x00")

func templatesRoutesJsonTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesRoutesJsonTpl,
		"templates/routes.json.tpl",
	)
}

func templatesRoutesJsonTpl() (*asset, error) {
	bytes, err := templatesRoutesJsonTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/routes.json.tpl", size: 2141, mode: os.FileMode(420), modTime: time.Unix(1536005076, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x18, 0x9, 0xcf, 0xd7, 0x0, 0x63, 0x1a, 0x27, 0xea, 0x3b, 0x87, 0x96, 0xc9, 0x8e, 0x81, 0xee, 0x26, 0xc8, 0xb8, 0x37, 0x3d, 0x53, 0x1f, 0xd2, 0x8e, 0x12, 0xd5, 0xd9, 0xf9, 0xe5, 0x3f, 0x97}}
	return a, nil
}

var _templatesRoutesTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x96\x5f\x6f\xd3\x30\x14\xc5\x9f\xe3\x4f\x61\x22\x31\x39\xa8\xb8\x13\xda\x13\xa8\x42\xb0\xb6\x0c\xb1\x8d\x6a\x2d\xe3\x11\x79\xf6\x6d\x6b\xe6\x38\xd9\xcd\xed\xca\x14\xf9\xbb\x23\xf7\xef\xd8\xca\xa0\x63\x42\x4c\xbc\x25\x96\xef\xb9\xc7\xf7\x77\x12\xb9\x54\xfa\x5c\x8d\x80\xe7\xca\x7a\xc6\x6c\x5e\x16\x48\x5c\xb0\x24\x1d\xe6\x94\xb2\x24\xf5\x40\xcd\x31\x51\x19\x9f\x8d\x22\x75\xa6\x2a\x68\x56\x17\x2e\x65\x2c\x49\x47\x96\xc6\x93\x33\xa9\x8b\xbc\x99\x6b\x5d\x38\xf7\x15\x9b\x8a\x48\xe9\x31\xa4\x2c\x63\x6c\x38\xf1\x9a\x0b\xcd\x9f\xd5\xb5\xdc\x2f\x3c\xc1\x37\x1a\x5c\x95\x10\x42\xc6\xdf\x75\x06\x5f\xea\x5a\xf6\x75\x51\xc2\xbe\xca\xc1\x85\x50\xd7\xf2\xa8\x30\xe0\xe4\xb1\xca\x21\x84\x63\x98\x8a\x8c\xd7\x2c\xa9\x6b\x6e\x87\x7c\xbe\xb5\xef\xd5\x39\xf0\xe7\x21\xb0\x64\xd1\x48\x9e\x80\x37\x80\x07\x83\xa3\x43\xa1\x1b\x3c\x5d\x8a\xce\x76\x86\x20\x57\xaa\x03\x75\xe6\xe2\x82\x46\x50\x04\x69\x36\x13\x06\x57\xfd\x52\xee\xae\x72\x6f\x78\x08\x2c\xfc\xe1\x51\x0f\x6d\x45\xf3\xb3\x2a\xe7\x1a\x1c\x10\xf9\xcb\x16\xd7\xb2\xfd\x56\x64\xf2\x8d\x73\x22\xaa\x8b\x8c\x2f\x3d\xf6\xa9\xc0\x68\xa7\xe6\x08\x34\x41\xcf\x3d\x4c\x45\x1e\x05\x2b\x79\x43\x3a\xe3\x21\x63\x89\x1d\xce\x44\x9f\xb4\xb8\xb7\x8e\xef\xec\x2c\xdf\xaa\x0b\x27\x3b\x88\xc7\xc5\x49\x31\xad\x62\xff\xd5\x18\x3a\x88\x05\x76\x15\x29\x27\x00\x31\x63\x49\x60\x89\x96\x7d\xa0\x53\x0b\xd3\xb6\x22\x25\x94\x73\xd9\x83\xc3\x71\xb6\xa2\x7b\xa3\xb9\x5e\xfc\x30\x60\xe6\x50\x16\x34\x4e\xe0\x62\x02\x11\x14\x4b\xac\x89\x4b\x28\xbb\x05\xe6\xa7\xca\x4d\x40\xa4\xd6\xc4\xd6\x97\x0a\x39\x29\x1c\x01\xf1\xcd\x3c\x56\x2c\xd6\x80\xbb\xd6\x1b\xb1\x33\xaf\x6a\x70\x6b\xb2\x57\xd7\x61\x45\x28\x8b\x92\xd6\x26\x60\x3f\x12\x13\x7b\xbb\x7b\x19\x4b\x22\xad\x2d\x50\xce\x7b\x3f\x72\x9a\xbd\x8f\xfd\x35\xce\x85\xa5\x9f\xfd\x52\x36\x10\x5d\x51\x41\xd9\x53\x58\x41\x44\x2b\x6e\xa3\xb8\x6b\xa8\x5b\xc1\x5f\x0a\xc5\x3e\x6d\xd0\x85\x81\x75\x06\xe6\xc1\xda\xae\xf9\xad\x58\xbd\xf7\x15\x20\x2d\x45\xb7\x13\x5b\x33\x32\x16\x41\x53\x4f\x8d\x40\x0c\x73\x92\xfd\x12\xad\xa7\xa1\x58\xb1\xef\x29\x1a\x87\xd0\xbc\x49\xef\xb5\x35\xad\xa7\x97\x69\x63\x31\x8f\xf5\x20\x3e\xc0\x55\x9f\x70\xa2\xa9\x6b\xc1\x99\x10\xb2\x6c\x1b\xac\xff\xd9\x57\xfa\x17\xd3\xf2\xa9\x34\x8a\xe0\xd1\xa7\xa5\xdd\x39\xec\x0c\x3a\xbf\x95\x97\xe9\x32\x2f\x55\x59\xf8\x0a\x3e\xa3\x25\xc0\x18\x9b\x7f\x3b\x48\x53\x39\x73\x7a\x00\xca\x00\x8a\x17\xbb\xbb\x31\x48\xc9\xfc\x1e\xb0\x55\xa4\x16\x76\xda\xe0\xe0\xbe\xe0\x37\x78\x09\xdf\x03\x00\x00\xff\xff\xd5\xe2\xab\x35\x51\x0a\x00\x00")

func templatesRoutesTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesRoutesTpl,
		"templates/routes.tpl",
	)
}

func templatesRoutesTpl() (*asset, error) {
	bytes, err := templatesRoutesTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/routes.tpl", size: 2641, mode: os.FileMode(420), modTime: time.Unix(1536005156, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x70, 0x81, 0xf8, 0xec, 0xc0, 0x32, 0x45, 0xc3, 0x66, 0x7d, 0x44, 0xa, 0xbb, 0x61, 0x55, 0x1c, 0xdf, 0xb5, 0x3c, 0xa, 0x57, 0x9b, 0xf6, 0x80, 0xbc, 0x4b, 0x6f, 0x8b, 0x2c, 0xe3, 0x8a, 0xe8}}
	return a, nil
}

var _templatesView_createTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x52\x4d\xcb\xdb\x30\x0c\x3e\xb7\xbf\x42\x88\x5d\x93\xdc\x47\xd2\xcb\x60\xb0\xc3\xca\xe0\xdd\xcd\x98\xe1\xc4\x4a\x6b\x70\xe4\xe0\x28\xed\x4a\xc9\x7f\x1f\x49\x9d\xb6\xd9\x7b\x32\x7e\x24\x3d\x1f\x96\x95\xca\xe0\xcb\xd0\x84\x9e\x7e\x19\x39\xc3\xd7\x0a\xf2\x8f\xe7\x2d\xd3\x7a\x3f\x37\x5c\x9d\x9c\x21\xff\x19\x2c\xf9\x05\xbb\xdf\x2d\xb5\x8e\x09\x50\x9c\x78\xc2\x69\x3a\xd2\x15\x94\xca\x8f\xa6\x23\xad\xef\x77\x62\x3b\x4d\x6f\x6d\x75\xb0\x37\x9c\xa6\x7d\x69\xdd\x05\x1a\x6f\x86\xa1\xc2\x26\xb0\x18\xc7\x14\xf1\xb0\xdf\x6d\x0a\x26\xda\x19\xfb\x04\x66\x0b\xcd\x5c\xd9\x95\x6d\x88\x1d\xb0\xe9\xa8\x42\xa6\xeb\x1f\xa5\xf2\xdf\xa6\xf6\xa4\x35\x42\x47\x72\x0e\xb6\xc2\x3e\x0c\x82\x60\x1a\x71\x81\x2b\x54\xea\x15\x53\xeb\xe2\x35\x50\x30\x5d\x1f\xa4\xbb\xb2\x75\xe4\xed\x40\xf2\xb8\xee\x4a\x4f\x27\x62\x7b\xd8\xa4\x2b\x8b\x84\x3e\x7a\x94\x8a\x86\x4f\x04\xf9\xf7\x65\x56\xeb\x15\xce\xc0\xb5\xc0\x41\x20\x3f\x86\x1f\x3c\x50\x94\xb5\xf6\x9e\x6b\xce\x91\x9d\x62\x18\xfb\xe4\x61\x56\x35\x35\x79\x68\x43\x9c\x5d\xe7\x1f\x12\xc7\x46\x16\x76\xad\xf1\xf0\x3f\x52\x16\x4b\xfb\x73\xd8\x71\x3f\x0a\xc8\xad\xa7\x0a\x85\xfe\x0a\xa6\x57\xfa\xc4\xb4\x71\x30\x2f\x23\x06\x8f\xc5\x9a\xbc\xb0\xee\x72\x78\x45\x21\xb6\xcb\xe6\x13\x40\x6c\x9f\x61\xde\x05\x87\xb1\xee\x9c\x20\x5c\x8c\x1f\xa9\xc2\x6f\x91\x8c\xd0\x53\xa9\x16\x86\x5a\x38\xeb\xa3\xeb\x4c\xbc\xad\x62\x65\xb1\x79\xf6\xb2\x98\x1d\x2d\xfb\x4f\x26\xd2\x99\x8e\xf5\x77\x25\x17\xff\x02\x00\x00\xff\xff\xf5\x89\x17\x15\xc1\x02\x00\x00")

func templatesView_createTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesView_createTpl,
		"templates/view_create.tpl",
	)
}

func templatesView_createTpl() (*asset, error) {
	bytes, err := templatesView_createTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/view_create.tpl", size: 705, mode: os.FileMode(420), modTime: time.Unix(1536004481, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf7, 0x9e, 0x1d, 0x29, 0x8f, 0xee, 0xb9, 0x16, 0xfb, 0x5c, 0x7c, 0x7b, 0xa4, 0xca, 0x95, 0x45, 0xc0, 0x10, 0xd8, 0xb6, 0x4e, 0xe0, 0xd3, 0x17, 0x59, 0x47, 0x66, 0x10, 0xd8, 0x1f, 0xfc, 0x11}}
	return a, nil
}

var _templatesView_listTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x52\x31\x6b\xf3\x30\x10\x9d\x93\x5f\x21\x4c\x56\x5b\x43\xb6\x0f\xd9\xdf\x52\xba\xb4\x0d\x85\x94\x2e\x42\x83\x62\x5d\x2a\x81\x6b\x17\xf9\x9a\x10\x84\xfe\x7b\x91\x64\x3b\x4e\x93\x42\xe9\x62\x9d\xee\xdd\x9d\xde\x7b\x3e\xce\x73\xb2\xea\xeb\xee\x03\x9e\x25\x6a\xf2\xaf\x24\xc5\x76\xba\xe5\x42\x2c\x43\xc1\xd1\xa0\x26\xc5\x53\xa7\xa0\x89\x39\xe7\x14\xec\x4d\x0b\x24\x43\x83\x0d\x64\xde\x73\x5e\x6c\xe4\x3b\x08\x41\x1e\x4d\x8f\xce\x41\xab\xbc\x9f\xd5\xed\x3a\x75\xca\x62\x26\xcd\x7a\x35\x70\xbc\x93\x28\xbd\x5f\x32\x65\x0e\xa4\x6e\x64\xdf\x97\x59\xdd\xb5\x28\x4d\x0b\x36\xab\x96\x8b\x0b\x40\x5a\x15\x72\x57\xc9\x3c\x4e\x0e\xc8\x82\xe9\x75\x75\xc9\x83\x51\xbd\x4e\x10\xca\x5d\x03\x63\x5f\xba\xc4\x6f\xbe\xeb\xac\x02\x0b\x69\x78\x28\xd4\x20\x55\x8a\x17\x0c\xed\x10\x05\x13\xac\x6c\xdf\x80\x14\xf7\x06\x1a\xd5\x0b\x91\x80\x88\x98\x3d\x69\x3b\x24\xc5\xa6\xdb\x42\x03\x35\x4e\x20\x43\x1d\x18\x6d\xd1\x7e\xd6\x18\x1b\x85\x60\x14\x75\x35\x6b\x86\x56\x45\x4f\xe7\x89\xf1\xca\xe8\xc8\x20\x74\x4d\xc4\x18\x06\xd1\x29\x76\x6e\xe0\xe5\xfd\x2d\xd2\xab\xa4\x35\xfc\xd6\x97\x10\xcd\x1f\xfa\x93\xa0\xd1\x9b\xa9\xb0\x78\x80\x53\x10\xc0\x24\xd1\x16\xf6\x65\xc6\xf9\x79\x9f\x84\xa0\x9c\x27\x0e\x42\xfc\x37\xaa\x74\xae\xf8\x6e\x88\xf7\x59\x75\x65\x44\x10\x76\xa3\xf2\xa7\xb7\xa9\xbc\x31\x82\xd1\x33\xdb\xdf\x3b\x3d\xee\xee\x90\x9d\x9c\x66\x34\xca\x88\x3b\x48\x95\x39\x84\x05\x4d\xe7\x70\x9c\x97\x3e\x9d\x9c\xc7\xf9\x5f\x01\x00\x00\xff\xff\x84\xc7\xb0\xd9\x61\x03\x00\x00")

func templatesView_listTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesView_listTpl,
		"templates/view_list.tpl",
	)
}

func templatesView_listTpl() (*asset, error) {
	bytes, err := templatesView_listTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/view_list.tpl", size: 865, mode: os.FileMode(420), modTime: time.Unix(1536004472, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x46, 0xdb, 0xc4, 0x3a, 0xf7, 0x26, 0xed, 0xa7, 0x96, 0x9d, 0x9, 0x38, 0xcb, 0x8f, 0x47, 0x89, 0x40, 0xdc, 0xfe, 0xb0, 0x79, 0xb3, 0x3e, 0x2, 0x52, 0xf3, 0x9, 0x9f, 0x2, 0xce, 0x3d, 0xfc}}
	return a, nil
}

var _templatesView_updateTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x52\xcd\x8e\x9c\x3c\x10\x3c\x33\x4f\xd1\x6a\x7d\xd7\x81\xfb\x27\x98\x5c\x92\x5c\xa2\xac\x22\x6d\x92\x8b\x65\x45\x06\x37\x33\x96\x8c\x8d\x4c\x33\x1b\x84\x78\xf7\xc8\xe0\xf9\x61\xf7\xd4\x72\xb9\xbb\xba\xaa\x6c\x21\x8e\xf0\xdf\xd0\xf8\x9e\x7e\x28\xbe\xc0\xff\x15\xe4\xaf\xf7\xd3\x51\xca\x43\x6c\x78\x33\x7c\x81\xfc\xbb\xd7\x64\x57\x6c\x9e\x35\xb5\xc6\x11\x20\x1b\xb6\x84\xcb\xf2\x45\x1b\x06\x21\xf2\x17\xd5\x91\x94\xf3\x4c\x4e\x2f\xcb\x53\x5f\xed\xf5\x84\x2b\xb2\x71\xfd\x36\xf4\xf6\x59\xb1\x5a\x96\x43\xa9\xcd\x15\x1a\xab\x86\xa1\xc2\xc6\x3b\x56\xc6\x51\xc0\xd3\x21\xdb\x5d\xa8\xa0\x23\xf6\x01\x3c\xae\xcc\xf1\x26\x2b\x5b\x1f\x3a\x70\xaa\xa3\x0a\x49\x1b\xfe\x23\x44\xfe\x53\xd5\x96\xa4\x44\xe8\x88\x2f\x5e\x57\xd8\xfb\x81\x11\x54\xc3\xc6\xbb\x0a\x85\x78\x78\x97\xb2\x78\x0c\x7c\x32\xba\x9a\xe7\x5c\x88\xfc\x1b\x4d\xaf\x1c\xc6\x86\xbf\x1a\xb2\x5a\xca\x65\xd9\xb6\x65\x65\x1b\x81\x81\x78\x3b\x66\xa5\xa5\x33\x39\x7d\xda\x47\x51\x16\x09\xde\x9a\x84\x08\xca\x9d\x09\xf2\x95\x6d\x90\x32\xcd\x3e\xb9\x8a\x2e\x8e\xe7\xe0\xc7\x3e\x2d\x8a\xd4\xaa\x26\x0b\xad\x0f\x51\x72\xbe\xd3\x83\xa7\xf7\x48\x59\xac\xed\xf7\x61\xe3\xfa\x91\x81\xa7\x9e\x2a\x64\xfa\xcb\x98\x32\xfa\xc0\x94\x06\xb2\xec\xaa\xec\x48\x15\x6e\x01\xbc\x77\x0f\xf7\x36\x21\x4c\x0b\x3e\x40\xfe\xe2\x7f\xf5\x5a\x31\x41\x4c\x4b\xca\x40\x4a\x7b\x67\xa7\x0a\x39\x8c\x84\x42\x90\xd3\x37\xa7\x59\x96\x3d\xfb\x8c\x0f\x1e\xbc\xc5\xe2\x16\x62\xa1\xcd\xf5\x9e\xd5\xd3\xdc\xce\xc5\x30\xd6\x9d\x61\x84\xa4\x73\x5b\x8e\xb7\x00\x6b\x76\x50\xb3\x3b\xf6\xc1\x74\x2a\x4c\x37\xee\xb2\xd8\x3d\x58\x59\x44\x01\xeb\x97\x4a\x3b\x53\x4d\xe5\xf1\x87\xb7\x9a\xd4\xfc\x0b\x00\x00\xff\xff\x42\xe1\x97\xd7\x30\x03\x00\x00")

func templatesView_updateTplBytes() ([]byte, error) {
	return bindataRead(
		_templatesView_updateTpl,
		"templates/view_update.tpl",
	)
}

func templatesView_updateTpl() (*asset, error) {
	bytes, err := templatesView_updateTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/view_update.tpl", size: 816, mode: os.FileMode(420), modTime: time.Unix(1536004464, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe9, 0xd7, 0xe5, 0x4, 0xf1, 0x53, 0xfb, 0x4d, 0x51, 0x65, 0x5d, 0x1b, 0x9d, 0xc, 0xbc, 0x65, 0x61, 0xee, 0x13, 0x95, 0x20, 0x76, 0xb8, 0x3f, 0xbf, 0x5e, 0x5b, 0xc1, 0xde, 0xa1, 0xff, 0x2f}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/model.tpl": templatesModelTpl,

	"templates/routes.json.tpl": templatesRoutesJsonTpl,

	"templates/routes.tpl": templatesRoutesTpl,

	"templates/view_create.tpl": templatesView_createTpl,

	"templates/view_list.tpl": templatesView_listTpl,

	"templates/view_update.tpl": templatesView_updateTpl,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
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

var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"model.tpl":       &bintree{templatesModelTpl, map[string]*bintree{}},
		"routes.json.tpl": &bintree{templatesRoutesJsonTpl, map[string]*bintree{}},
		"routes.tpl":      &bintree{templatesRoutesTpl, map[string]*bintree{}},
		"view_create.tpl": &bintree{templatesView_createTpl, map[string]*bintree{}},
		"view_list.tpl":   &bintree{templatesView_listTpl, map[string]*bintree{}},
		"view_update.tpl": &bintree{templatesView_updateTpl, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory.
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

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
