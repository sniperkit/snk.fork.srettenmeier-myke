/*
Sniperkit-Bot
- Status: analyzed
*/

// Code generated by go-bindata.
// sources:
// tmp/LICENSES
// tmp/version
// DO NOT EDIT!

package core

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
	bytes []byte
	info  os.FileInfo
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

var _tmpLicenses = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x7d\x6d\x73\xdb\x38\xb2\xee\x77\xfe\x8a\xbe\xae\xba\x3b\xf6\x2d\x46\xce\x64\x5e\xee\xd9\x99\xad\xad\x51\x6c\x25\xd1\x8c\x23\xf9\x48\xf2\x64\x53\x5b\xfb\x01\x22\x21\x09\x63\x12\xe0\x02\xa4\x1c\x9d\x5f\x7f\xaa\x01\x90\x04\x25\xd9\xf1\xdb\xd6\x38\x71\xe7\xc3\x8c\x25\x91\x78\x69\x00\x0f\xfa\xe9\x06\xba\xc7\xe7\x83\x11\x4c\xc7\x17\x93\x93\x01\x9c\x0d\x4f\x06\xa3\xe9\x60\x1a\x45\x4b\x51\xae\xaa\x79\x2f\x51\xf9\xf1\x52\xf1\x4a\xab\xe3\x7c\x73\xc9\xa3\x28\x7a\x3f\x9c\xc1\x99\x48\xb8\x34\x3c\x8a\x4e\x54\xb1\xd1\x62\xb9\x2a\xe1\x30\x39\x82\x57\x2f\xbf\xfd\x11\xde\xaa\x41\xa5\x15\xcc\x34\x5b\xf3\x0c\xde\xe6\xf3\x77\x51\x74\xce\x75\x2e\x8c\x11\x4a\x82\x30\xb0\xe2\x9a\xcf\x37\xb0\xd4\x4c\x96\x3c\x8d\x61\xa1\x39\x07\xb5\x80\x64\xc5\xf4\x92\xc7\x50\x2a\x60\x72\x03\x05\xd7\x46\x49\x50\xf3\x92\x09\x29\xe4\x12\x18\x24\xaa\xd8\x44\x6a\x01\xe5\x4a\x18\x30\x6a\x51\x5e\x31\xcd\x81\xc9\x14\x98\x31\x2a\x11\xac\xe4\x29\xa4\x2a\xa9\x72\x2e\x4b\x56\x62\x7d\x0b\x91\x71\x03\x87\xe5\x8a\xc3\xc1\xd4\xbf\x71\x70\x64\x2b\x49\x39\xcb\x22\x21\x01\x7f\xab\x7f\x82\x2b\x51\xae\x54\x55\x82\xe6\xa6\xd4\x22\xc1\x32\x62\x10\x32\xc9\xaa\x14\xdb\x50\xff\x9c\x89\x5c\xf8\x1a\xf0\x75\x2b\x03\x13\x95\x0a\x2a\xc3\x63\xdb\xce\x18\x72\x95\x8a\x05\xfe\x9f\xdb\x6e\x15\xd5\x3c\x13\x66\x15\x43\x2a\xb0\xe8\x79\x55\xf2\x18\x0c\x7e\x69\x85\x19\x63\x3f\x8e\x95\x06\xc3\xb3\x2c\x4a\x54\x21\xb8\x01\xdb\xd7\xb6\x75\xf6\x19\x6c\x7a\x81\x02\x2d\xbd\x88\x0c\x7e\x73\xb5\x52\x79\xb7\x27\xc2\x44\x8b\x4a\x4b\x61\x56\xdc\xbe\x93\x2a\x30\xca\xd6\xf8\x07\x4f\x4a\xfc\x06\x1f\x5f\xa8\x2c\x53\x57\xd8\xb5\x44\xc9\x54\x60\x8f\xcc\x4f\x51\x34\x5b\x71\x60\x73\xb5\xe6\xb6\x2f\x6e\x88\xa5\x2a\x45\xe2\xc4\x6d\x07\xa0\x68\x47\xd5\xff\x64\x56\x2c\xcb\x60\xce\xbd\xc0\x78\x0a\x42\x02\x0b\xba\xa3\xb1\x7a\x53\x32\x59\x0a\x96\x41\xa1\xb4\xad\x6f\xbb\x9b\xbd\x28\x9a\xbd\x1b\xc0\x74\xfc\x66\xf6\xa1\x3f\x19\xc0\x70\x0a\xe7\x93\xf1\xef\xc3\xd3\xc1\x29\x1c\xf4\xa7\x30\x9c\x1e\xc4\xf0\x61\x38\x7b\x37\xbe\x98\xc1\x87\xfe\x64\xd2\x1f\xcd\x3e\xc2\xf8\x0d\xf4\x47\x1f\xe1\xb7\xe1\xe8\x34\x86\xc1\x3f\xce\x27\x83\xe9\x14\xc6\x93\x68\xf8\xfe\xfc\x6c\x38\x38\x8d\x61\x38\x3a\x39\xbb\x38\x1d\x8e\xde\xc2\xeb\x8b\x19\x8c\xc6\x33\x38\x1b\xbe\x1f\xce\x06\xa7\x30\x1b\x03\x56\xe8\x8b\x1a\x0e\xa6\x58\xd8\xfb\xc1\xe4\xe4\x5d\x7f\x34\xeb\xbf\x1e\x9e\x0d\x67\x1f\xe3\xe8\xcd\x70\x36\xc2\x32\xdf\x8c\x27\xd0\x87\xf3\xfe\x64\x36\x3c\xb9\x38\xeb\x4f\xe0\xfc\x62\x72\x3e\x9e\x0e\xa0\x3f\x3a\x85\xd1\x78\x34\x1c\xbd\x99\x0c\x47\x6f\x07\xef\x07\xa3\x59\x0f\x86\x23\x18\x8d\x61\xf0\xfb\x60\x34\x83\xe9\xbb\xfe\xd9\x19\x56\x15\xf5\x2f\x66\xef\xc6\x13\x6c\x1f\x9c\x8c\xcf\x3f\x4e\x86\x6f\xdf\xcd\xe0\xdd\xf8\xec\x74\x30\x99\xc2\x6b\x5c\x88\xfd\xd7\x67\x03\x57\xd5\xe8\x23\x9c\x9c\xf5\x87\xef\x63\x38\xed\xbf\xef\xbf\x1d\xd8\xb7\xc6\xb3\x77\x83\x49\x84\x8f\xb9\xd6\xc1\x87\x77\x03\xfc\x0a\xeb\xeb\x8f\xa0\x7f\x32\x1b\x8e\x47\xd8\x8d\x93\xf1\x68\x36\xe9\x9f\xcc\x62\x98\x8d\x27\xb3\xe6\xd5\x0f\xc3\xe9\x20\x86\xfe\x64\x38\x45\x81\xbc\x99\x8c\xdf\xc7\x11\x8a\x73\xfc\x06\x1f\x19\x8e\xf0\xbd\xd1\xc0\x95\x82\xa2\x86\xce\x88\x8c\x27\xf6\xf3\xc5\x74\xd0\x14\x08\xa7\x83\xfe\xd9\x70\xf4\x76\x8a\x2f\x63\x17\xeb\x87\x7b\x51\xd4\x81\x93\xf7\xcc\x94\x38\x73\x64\x6a\x8e\x4d\xa1\xc5\x32\x8a\xa2\xa9\xfd\x7f\x80\x26\x27\x16\x4d\xbe\x83\xe0\x61\x82\x91\xe7\x0c\x23\x11\x7e\x45\x30\xf2\xac\x60\xa4\x33\x7c\xdb\x30\xc2\xd4\xa5\xca\xc4\xf1\x52\x55\xa5\xc8\x0c\xfe\x08\x9f\xfb\xd7\x2f\x58\xb2\xe2\x8d\xf6\x72\xc3\x93\xbf\x73\x6d\x27\xe4\xab\xde\xcb\x18\x7e\x65\xb2\x62\x7a\x03\xaf\x5e\xbe\xfc\xfe\xda\x97\x56\x65\x59\xfc\x74\x7c\x7c\x75\x75\xd5\x63\xb6\x9a\x9e\xd2\xcb\x63\xbf\x28\xcd\xb1\x6d\xdd\x6c\x30\x79\x3f\xb5\xc3\x7b\x32\x1e\x9d\x0e\x51\x28\x6e\x1a\x5c\xa0\x10\x27\x83\xf3\xc9\xf8\xf4\xc2\xca\x2a\xb6\x4f\x9d\x0e\xa7\xb3\xc9\xf0\xf5\x05\x7e\x63\x0b\xf8\xb6\x07\xa7\x7c\x21\xa4\x5b\x5f\xbd\xba\xcb\x07\xbe\x47\x07\x7e\xe5\xe4\x9c\x39\x3c\x41\xe8\x34\x76\xa5\xb5\xab\x12\x16\x4a\x3b\x78\xd1\xbc\xd0\x2a\xad\x1c\x38\xf9\xa2\xf0\xd9\x06\x58\x50\x02\xcc\x40\x8a\x55\xf2\x14\xe6\x1b\x98\xf2\xc4\x15\xf2\x2d\x94\x2b\xad\xaa\xe5\x0a\xfe\x0a\x35\x92\xd6\x88\xb9\xdd\x2e\xa5\x77\x1a\xd6\xa2\x81\xba\x92\x5c\xe3\x8a\xe6\xb2\x14\xe5\x06\x58\x55\xae\x94\x16\xff\x63\xeb\xf3\xe5\xec\x7b\xa3\x5c\xb1\x12\x77\x01\x0b\xff\x88\x3b\x65\x3b\xb2\x41\x03\xf8\x92\x65\x30\xb0\x45\xef\x34\xa2\x92\xd8\x41\x0f\x1a\x2c\xb1\xa5\xd4\xad\xc0\xbd\x20\xcb\x7c\x31\xaa\x5c\x71\xdf\x40\x84\x1f\x5b\x75\xa2\x64\xa9\x55\x16\x03\x62\xa4\xff\x90\xd9\x46\xc7\xd8\x1b\xfc\xb6\x92\x29\xd7\x90\xa8\x3c\x57\xd2\x97\xe4\x1f\xb4\xf8\xef\xca\x71\x15\xf6\xe0\x8d\xd2\xb6\x1d\x45\xa5\x0b\x65\x6a\xcc\x16\x5e\xfa\x22\x1c\xa3\x03\x5f\xca\x81\xed\x8a\x81\x43\x71\xe4\x5e\x55\x57\x5c\xe3\xbe\xa0\x11\x98\x95\x06\x21\xdd\xdf\x76\x9b\x4a\x58\x65\x38\x3e\xe7\x4b\x71\x3f\x59\x09\x68\xc8\x99\x64\x4b\x8e\x83\x87\xf5\x9a\x2a\x59\xf9\x86\xc5\x70\xb5\xe2\xb6\xfb\xf3\x8d\x6b\x3d\xb3\x65\x87\x92\xb9\x12\x38\x9b\x94\x86\x43\x21\x8e\xdc\xf0\x98\x95\x28\xb0\xa4\x85\x58\x94\x76\x0b\x4e\xb0\xe8\xc3\x1f\x5e\xfe\xdf\x23\x5b\x9d\xd2\xdc\x0b\xbe\x2e\xa8\x2a\x11\xcf\xed\xe6\x68\x56\x4c\x73\x53\x97\x28\x8e\x60\xce\x25\x5f\x88\x04\xb1\xbe\x53\x7a\xd0\xce\x76\xc8\x3f\xaa\xea\x00\x0e\x95\xb6\x7f\xe9\x83\xa3\x70\xd4\x99\xb4\x32\x59\x8b\xb4\xc2\xb2\x34\x84\xf3\xc3\x17\xc0\x3f\x71\x9d\x08\x83\x0d\x69\x77\x26\x53\xab\x19\x28\x06\x3b\x2c\x3b\x53\x6d\xaa\x2a\x9d\xf0\x03\x5c\x5e\xf9\xf6\x4c\x2b\x34\x5f\x70\xad\x79\xea\x7e\x5d\x58\x89\x5f\x62\x15\x76\x73\x17\x89\x55\x01\x4c\x3d\xc0\xad\x9e\x30\xaf\xec\x7e\xe9\xf4\x04\xb7\xff\x36\xfa\x8a\xb1\x15\x42\xa2\x52\x1e\x77\xb5\x15\x5f\x8c\x7b\x20\xae\xd7\xff\x42\x2c\x2b\x1d\x68\x33\x6d\xd3\xc7\x76\x2b\xdf\x6d\x3a\xaa\x4f\xf6\x3b\xcd\x4d\x95\xd9\xf5\xb1\xd0\x2a\x87\x9c\x27\x2b\x26\x45\xc2\xea\x05\x52\x6a\x26\x0d\x3e\xc9\xea\x09\x65\xbf\xc9\xfc\xc7\x05\x30\x70\xe2\xb1\xc5\xc5\xdd\x0e\xfa\x32\xb6\xba\x99\xa8\xbc\x10\xb8\xa0\x94\xd3\x33\x5c\x37\x97\x5c\x72\xbd\xab\x9e\x85\xe8\x95\x28\xb9\x76\xe8\x6d\x15\x1a\xb7\x76\x73\x9e\x0a\x06\xe5\xa6\x08\xbb\xfd\x41\xe9\xcb\x1d\x50\xb8\x52\xfa\xd2\xb6\xd8\xe2\x10\xce\xb4\x76\x09\x08\x59\x77\xa3\x59\x00\x4e\x74\xbe\x5b\x39\x4b\x39\xb0\x35\x13\x19\x9b\x67\xf5\xfa\x0f\x70\x29\x46\x34\xc5\x09\x98\x30\x3f\x95\x58\x83\x0b\x5b\xda\x51\x0d\x6f\x21\x91\x42\x58\x29\x4b\xdc\x5b\xd2\x5a\xed\xc2\xd6\xfa\x22\x0e\x99\x04\xfe\x89\xe5\x45\x86\x3a\x1b\x14\x5a\xad\x85\x7f\x11\x9f\xec\x17\x05\x97\xa9\xf8\x04\x73\x9e\xa9\xab\xa3\x56\x0a\xa7\x5c\x8b\x35\x2b\xc5\x9a\x03\x0a\xc4\x1c\x6c\xcf\x00\xac\x63\xbf\x0c\x7c\xef\x7d\x49\x4e\x06\x75\xc3\xe7\xcc\xe0\xe0\x49\xbb\x14\x53\xac\x03\x67\xbf\x56\xb9\xc3\x2a\xac\xca\x0e\x17\xae\x85\xab\x95\x48\x56\x01\x18\xf0\x54\x94\x4a\xe3\x72\xd7\x7c\x2d\xec\x50\xe2\x2c\x96\xaa\xf4\xeb\x04\x78\xc6\xe6\x4a\xd7\x9f\x94\xae\x87\x39\x5c\x4d\xbe\x30\xdc\xe5\xb8\xe1\xb2\xb4\xd2\x67\xa8\xe1\x66\x76\x51\x80\xd2\x62\x29\x24\xcb\xf6\x8c\xf9\x2e\x1e\xd7\x38\xb5\xe8\x2c\xff\x18\xb6\xc5\xe7\xa5\x87\xb3\xd9\x8f\x9d\x2d\xde\xef\x1a\x9a\xe7\x4c\x34\xeb\x93\x17\x4c\xdb\x99\x82\x72\xb1\xdd\xc8\xb9\xe6\xd9\x06\x32\x21\x2f\xad\xe0\xe6\x42\xda\x79\x22\x59\xce\x8f\xea\x41\x17\xb2\xe4\x7a\xc1\x12\xbb\x49\xc4\xc1\x1e\xd9\x08\x75\xa7\x51\x28\x1d\xae\x16\xed\xa8\x9f\x20\x94\xfb\x3d\x7e\xef\x88\x6f\xaf\x81\x66\xc9\x06\xf5\x35\x02\xf4\x0b\xae\xde\x4b\x9b\x76\x60\x61\x9d\x31\xb1\x73\x38\xf5\x9a\x48\x5d\x92\x72\xb2\xb1\x6f\x29\x7d\x6d\xe3\xe3\x60\x51\x94\x88\xfa\x4a\xb2\x2c\xab\x61\xdb\x54\xf3\x5c\x94\x1e\x3c\x6a\xbd\xc3\xce\x2e\xdb\x72\x47\x1a\x65\xdb\x3c\x8b\xe3\x3b\x6a\x45\x3d\xca\x76\xbb\xbb\x71\xb7\x08\x15\x15\x44\x65\x5b\x3d\xce\xf7\x39\x5f\xb1\x6c\x01\x6a\x71\xbd\xf2\x72\xbb\xdd\x1e\x0e\x9a\x3e\x1d\xf8\xb2\xdc\x7e\xdf\xc0\xb2\x5a\x00\xcf\x78\x52\x6a\x25\x45\x12\xe3\x28\xcc\x59\x66\xe7\xd1\x95\xc6\xf7\xa4\x55\x3e\x2a\xe9\xa5\x0f\xb8\x0a\x42\xa1\xf3\x56\x50\x28\xa7\xd2\xb4\x8b\xc5\xca\xdf\xc4\x37\x6e\x45\x0d\x76\x85\x75\x28\x19\xb4\x09\x72\x26\x32\x7c\x39\x13\xa6\x34\x71\xb8\x65\x35\xaa\x90\xd9\x98\x92\xe7\x26\x84\x70\x61\x4c\xc5\x71\x0b\x49\xec\x1e\xe9\x9f\x70\xc3\x8f\x3b\x9f\xd3\x56\x1a\x5d\x2b\x14\x7a\x1c\xc0\x48\x67\x16\x04\xd2\x46\xb9\xa5\xc2\x24\x95\xb1\xbb\xbc\xad\x31\xb7\x78\xe9\xd5\xc8\x0f\x16\xf1\xda\xad\x89\x7f\xaa\x85\xd0\xed\x6b\x3d\x1f\x13\x25\x4d\x21\x92\x4a\x55\x26\xdb\x40\xce\xf4\x25\x42\x9f\x6e\xb5\xa3\x5a\xe5\xe2\x46\x2c\xa5\xc5\x7e\x21\xed\x18\x59\xc1\xee\x9d\x89\x08\x56\x07\x23\x55\x02\x83\x70\xad\xf6\x0e\x76\x97\xf0\x96\x7e\xdd\x74\xbb\x5e\x81\x9f\x55\x79\x42\x01\x3a\x0b\x40\xb7\x52\x58\x31\x03\x73\xce\x25\x68\x9e\x70\x8b\xe4\xf3\x4d\xa7\x9e\x76\x11\x1a\xfe\xef\x8a\xcb\x32\xc3\x6a\x13\xa5\x0b\xe5\xb6\x6b\x54\x78\x83\xe5\xe7\x80\xe8\x55\x0f\xde\xa2\x5a\x85\xd5\xb6\xb6\x9f\x5a\xb3\x82\x69\xd7\xc4\xb0\x97\xcc\x04\xcb\x2c\x44\x65\xce\x92\x15\x04\x02\xea\x18\x8b\xac\x5e\xf0\x51\x55\xc0\x50\xc3\x2b\x78\x59\xb1\xac\x9e\x7e\x57\x4a\x67\xe9\x95\x40\x5d\x43\x2a\xf9\xc2\x8e\xbc\x11\x6b\xfb\xf1\x45\x6d\x59\xd2\x6a\xc3\xb2\x72\xf3\x62\xa1\x39\x8f\x41\x68\xcd\xd7\x2a\x41\x20\xdf\xd9\xcd\x3d\xff\xc3\x0a\x6b\xb6\xc5\x63\x54\x07\x0b\x9c\xc7\x3b\x48\xd7\xc2\xb9\xb5\xf2\x24\xd9\x06\x27\x6a\x91\xb1\x4d\xdc\x7e\x53\x70\xed\xb6\xda\x2d\xa3\x4f\x60\x10\x0a\x16\x41\x83\xc5\x56\x59\xde\xa9\x71\xcf\x76\x6e\xb1\xc5\x0d\xd0\x77\xc1\x00\x9d\x33\x04\xdd\xaf\x60\x74\x0e\xf9\xa7\x84\x17\x25\x2e\x30\x53\xd6\x8b\xd1\x99\x02\x1d\x21\x3a\x82\xc2\xf5\x35\x18\xbd\x9c\x5d\xf2\x18\x56\x6c\xcd\xad\x96\x57\x37\xc8\xf2\x68\xb5\x58\xa0\x9e\xa7\xac\xc9\x2d\xf6\xff\x15\x79\xa1\x74\xe9\x06\xa6\xc1\x01\xaf\x28\x7b\xad\xd0\xc2\x4c\xdd\x33\x14\x81\x1b\xa3\xba\x56\x56\x14\x99\xb5\x76\xc9\x6c\xe3\xa4\x8c\xd8\xe5\x9b\x96\x64\x4c\xe4\xc6\x3f\x1b\x74\x6e\xbe\x71\x85\x84\xd2\x6d\x70\x53\xf2\x84\x1b\xc3\xb4\xb0\xab\x73\xa1\x85\x5c\xd6\x8c\x86\x8b\x7a\xef\x0b\x17\xfe\xa1\x39\x02\x96\x29\xc9\xfd\x8e\x98\xa8\x7c\x2e\x64\xa3\xd5\xdb\xd7\xb6\x5f\xa8\x3b\xe4\x18\xae\xdf\x6d\xad\x65\x11\x95\xbc\x6e\xe3\x7c\x15\x57\x38\x14\xf5\x5e\xd7\x83\xe1\x02\xc7\xbf\xe1\x42\xa6\x14\x25\xce\xe9\x66\x50\x4a\xb1\x74\x4d\x60\x4b\x86\x3f\x5b\x90\xf3\xc4\xfd\xb0\xdd\xb0\x1a\xdd\x5a\x2b\x63\x5e\x58\x81\x61\x37\x12\x55\xa1\xfe\xe4\x3e\x0b\x09\x0c\x32\x76\x65\x2a\x51\x62\x57\x33\xbe\x74\x9b\x00\x2b\x9b\xc6\xb7\x3a\xc1\x16\x2a\xde\x04\x70\x76\x4f\x70\x0d\x37\x9e\x6a\xb7\xe5\x24\xed\xe0\x6c\xea\x6e\xd5\xe3\x91\x5b\x4d\xb5\x5c\x71\xa7\x8a\x75\x67\x62\xad\x32\xd5\x64\xd4\xaf\x94\x9a\x68\xb4\x6b\xcc\x6f\x79\xb5\x56\xe5\x76\x07\x67\x48\x67\x65\x3d\x57\x58\x63\x31\x4d\x59\xd9\x4c\xbe\x46\xba\xc2\x58\x9e\x98\x3a\x28\xf8\xbe\x07\x13\x1e\x5a\x86\x7a\xb6\xea\x9c\x6d\x5a\x64\xdb\x46\xa1\x8e\xf5\x39\xc4\xa3\x1b\xb4\x3c\x3b\x24\xa8\x36\xf2\x54\x54\x79\xec\xe6\x11\x6a\x34\xce\x62\x5e\x2b\x42\x1d\xda\xec\xb6\xf0\x6b\x90\x2c\x6e\xa9\x90\x15\x48\x3b\xb5\x72\xce\xcb\x9b\x8c\xd7\x1e\x2e\xd8\x91\xeb\x69\x65\x4a\x58\x62\x7b\xb1\x79\x8e\x6f\x68\x9e\x88\x42\x70\x04\xad\x50\xf5\x6d\xd8\x21\xfe\xdb\xe9\xa8\xf3\x3f\x6c\x33\x89\x9f\xed\x36\x5a\xd7\x39\x0f\xea\x74\x86\x9b\x56\x95\x46\x1e\x65\xbd\x11\xd6\xa8\xa3\x71\x0a\x69\x95\x0b\x89\xf3\xc4\xb1\x47\x13\x54\x8f\x10\xd7\x4c\x69\x2c\x13\xa9\xfb\xd2\x0a\x83\xbb\x72\xba\x35\x27\x41\xcd\x9a\x97\x4c\x58\xbf\x85\xb7\xab\x37\x14\xde\xb2\x03\xb9\xd9\xe9\x5c\x50\x71\x53\x61\xe8\xa7\xf0\x46\x7c\xb7\x3b\xc6\x7e\x76\xc7\x08\x8b\x29\x47\xbd\x29\x0e\x94\x09\x3b\x45\xcb\x76\xb9\xf9\xbe\x39\x13\xc4\x9e\xf6\x6c\x43\x6a\x57\x73\x73\xe8\x59\x97\x61\x1b\x97\x2a\xab\xd0\x16\x5c\x63\x37\x1b\x7f\x11\xd3\x65\xbb\x71\x81\xd7\xe0\xb7\x3b\xda\x15\x5a\x7a\x84\xa0\xd5\x8c\xbf\x27\x7e\x38\xd4\x07\xa3\xf1\x6c\x78\x32\x38\x80\x92\x7f\x2a\xad\xbc\x71\xd9\xf9\x3a\x50\xe5\x0e\xea\x09\x57\x57\x00\x01\x7b\x56\xca\x8e\x64\xed\x78\x05\x45\xd5\xd4\x93\x81\xe6\x2c\xb5\x1c\xb3\x9d\x74\x7c\xaf\x58\x11\x94\x98\x90\x3c\x14\xbf\x07\x35\x8b\x0c\xae\x23\xb6\x0b\xf1\x6d\xe4\x1a\x14\xb3\x5f\xc2\x7b\xe5\x6a\x27\x1b\x2b\x21\xe3\xcc\x20\x9d\x0a\xad\xf4\xfe\x95\x76\xb5\x16\x19\x92\xe0\x9f\xea\x66\xb2\xba\x8d\xad\xac\x5b\x09\x75\x66\x95\xb9\xb1\x0d\x3f\x87\x60\xde\x99\x64\xe1\xba\xee\x1a\xa0\x40\x2c\x5a\x9c\xc1\x2d\x73\xd9\xee\x80\xbb\xe5\x2b\x1d\xef\x4a\x99\xd5\xba\x5e\x60\xe5\xf2\xdc\x60\x8f\x94\x16\x5b\x2b\xc5\x2a\x10\x6b\xae\xdd\x60\x95\x2b\xa1\xd3\x17\xd8\xc9\x4d\x33\x36\x52\xe9\x1c\x09\x33\x2a\x16\x9c\xe9\x1e\xcc\x56\x8e\x85\x21\x7e\xed\x8a\x39\x18\x6f\xab\x3c\x38\x2a\xdd\x18\xf9\x58\x16\x90\x57\xd4\x50\xba\xcd\xf1\x6b\xcb\xf9\x2e\x3b\xb6\xf9\x66\xdb\x60\x69\x8a\x7f\x6b\xe4\x3b\xe1\x8c\x0c\x4a\xa9\x9b\xee\x25\x74\x9b\x95\x10\x3b\xe9\x1b\x91\x76\xa6\x8e\xe5\x53\x4c\x62\xa5\x5c\xa6\x55\x5e\xab\xad\x9d\x19\x53\x03\x8b\xe3\x7f\xf5\x70\x6e\x63\x9a\x15\x70\x6d\xc4\x60\xd9\xfe\xc5\x64\xad\x55\x30\xe7\x4e\x0f\xd0\xd5\xf6\xfc\x73\x82\xb9\xce\x6f\xb1\x57\x44\x2d\xab\xb0\x6a\xab\x35\xd6\x3b\x05\x60\xcb\xf0\x15\x0c\x05\x16\xe2\xfb\x11\x36\x59\x69\x48\x05\x6a\xad\x1d\x2d\x77\x8f\x06\xdf\x9a\xf6\xf6\xb8\x8c\x5c\x31\x81\xaf\x48\x2d\xf6\xb4\x26\x6e\x97\xcd\xc2\x92\xc5\xcd\x35\x54\x24\xb4\xce\x35\x4b\xc9\x96\x87\x55\x07\xd6\xbc\xb6\x01\x3b\xde\xaa\xce\x2e\xdc\x68\xdd\x89\xca\x9d\x2a\x8d\xf3\xa8\x63\x96\x69\x98\xca\x16\x13\xe8\x0c\xc8\x0f\x96\xec\xd4\x3e\x6a\xcb\x55\x5b\x2d\xd0\xf4\xe0\x42\x66\xdc\x18\x3b\x68\xfc\x53\x91\x89\x44\x20\xfd\xb5\x25\x06\x0e\x92\xc6\xbe\xb1\xd9\xd6\x22\x03\x63\x56\x60\xc6\xba\xd6\x74\xd5\x6a\xfa\x58\xe3\xb6\x21\xa7\xf1\x9d\xb7\xd6\xe7\xbb\x50\xb3\xfa\x60\x02\x36\x33\x98\x30\xae\x08\xa7\xba\xa6\xb5\xf7\xd1\xbd\x3f\x52\x25\xbe\xd4\x78\x6f\xca\xda\xe5\x8f\xa4\x0c\x97\xed\xd2\xd2\x3b\xdc\x46\x6c\xd3\x4c\x55\x70\x6d\x78\xca\x9d\x23\x08\x97\x41\x30\x24\xbe\x22\xa7\x5d\x38\x03\x69\xc9\x5b\x4a\xb4\xd4\xdc\x4d\xfc\x8d\x5f\x21\x96\x91\xf1\x4f\x3c\x09\x20\xde\x02\x6f\x23\x10\xcd\x97\x4c\x3b\xbf\xd2\x36\xf7\xf0\xbe\x80\x1f\x7b\x30\xab\x15\x10\x83\xb0\x18\xe8\xd1\xa9\xb2\xc8\x59\x3a\x95\x3b\x3c\xab\xe0\x0e\x69\xb8\x46\xe3\xdb\xb5\x1b\x83\xe5\xdc\x04\x1a\x8d\x41\x42\xa8\xd7\x22\xe1\xe0\x3f\x2a\x0d\x7e\x0e\xbb\x87\xeb\x49\x5b\xb7\x38\x6e\xad\x4e\x9e\xa6\x6a\xfe\xef\x4a\x78\xef\x11\x6e\xe8\x46\x49\xbb\xa5\xdb\x21\xad\x4c\xa9\x72\xa6\x37\xb6\x35\x42\x42\xca\x4d\xa2\xc5\xdc\x0f\x45\x43\x3a\xc4\x52\xec\xda\x67\xeb\xd5\x54\x8f\x9b\xdf\x0d\xf6\x6c\x01\x4e\x52\xff\xbf\x07\xa7\xc2\x58\xea\xc4\x35\x3e\xf5\x81\x69\x94\xcb\xa6\x59\x04\x4d\x53\xe7\x1b\x47\x60\x2d\xf3\x46\x8a\xd5\xc2\x80\x1d\x45\x4b\x5e\x5a\x2b\x58\xdc\x0e\x98\x5f\xfb\xa6\x6d\xea\x21\xb6\x95\xb3\x64\xb5\x4d\x51\xc3\xa7\x45\x69\xba\x83\x7b\x04\xca\x7a\xfc\xfc\x51\x0f\x78\xdd\x9f\x0e\xa7\xb5\x70\xb7\x8e\x7d\x0c\x07\xfe\x0c\x45\xe3\x96\xef\x1c\x03\xe1\xc2\x79\x80\x3f\x15\x1a\x3b\xd9\xf4\x44\x58\x5c\x49\x03\x33\x69\xbc\xe7\x68\x4f\xec\x8c\xea\x4e\x54\xfe\xfc\xca\x0e\xc4\xaa\x05\xcc\x86\xb3\xb3\x41\x0c\xa3\xf1\xe8\x45\x78\xf6\x23\xde\x39\x42\x82\x05\x74\x4e\x91\xf8\x32\x76\xcf\x92\xb8\xdd\xd6\x79\x0b\x33\x9e\x21\x57\x33\x85\x92\x46\x58\xaf\x83\xf5\xcc\x38\x56\xd8\x9d\x2e\xac\x28\xb4\x2a\xb4\x40\xf5\xdc\x76\x78\x01\x95\xb5\x95\xda\xf9\xd7\x22\x6e\x60\x2f\xad\x8f\x4f\x55\xb9\xe5\x2a\x35\x5c\x0b\x63\x91\xbd\x39\x55\x65\xd7\xa6\x05\x75\xef\x67\xb5\xd6\xd8\xd0\xd1\xba\x4b\x66\xdd\xdc\xfb\xaf\x1e\x9c\xb5\xa7\xa5\xd4\x02\xce\x04\x9b\x8b\xcc\x3a\xcf\x87\xb8\xf3\x02\x5f\xe3\xdc\xc5\x76\xb8\x32\xa4\x82\xcc\x1a\x3b\xcb\x15\x57\x7a\x13\x98\x5a\x6a\x4f\x56\xa9\x74\x19\x9a\x0c\x24\x5f\x66\x62\xc9\x65\xc2\x8f\xe2\xc6\xdb\x1d\x77\x4c\xb9\x8d\xe5\xe7\xb3\xf3\xfd\xd0\x29\x0a\x06\x52\x9e\x89\xb9\x55\xe8\x6c\xe3\x96\x5a\x19\xd3\xf8\x2d\xea\x2a\x4b\x60\x49\x69\xac\x77\x7c\xff\xfa\x70\xe8\xd9\xd9\x3e\x94\x86\x79\x3d\x64\x99\xb0\x15\x7b\x8b\x80\x1d\x5a\x96\xb3\x65\xd7\x86\x8f\x6f\xd7\x47\x02\xda\xc3\x01\xa6\xe0\x89\x68\x8d\x6c\x42\x26\x22\x45\xc5\xd6\xb9\x12\x50\x81\x71\x36\x5d\xc1\xb2\xba\xd0\x1a\xa1\x93\x15\x43\x11\x71\x0d\x4c\x3b\x9f\x39\xee\xe2\xcd\x5e\x6d\xaa\xac\xdc\x26\xba\x56\x9a\x55\x83\x31\x95\xfb\x46\x48\x3f\x98\x01\xae\x86\x16\x83\xc3\x1b\x7d\xe2\x75\xab\xb0\xdb\x99\x72\x13\x76\xa9\x54\x7a\x25\xb2\xd0\x76\x78\x09\xa6\x54\x45\xc1\x96\xf6\x6c\x5d\x5e\x54\xd8\xf0\x05\x13\x59\xa5\xdd\x6e\xc4\xb2\x45\x25\x5b\xe5\xc6\x6e\x82\x7b\x4e\x82\x24\x2a\xcf\x71\xf2\x86\xf2\x70\x15\x73\x73\x14\xdb\x79\x88\x0a\xfa\xb6\x21\xce\x97\xd1\x18\xd3\x59\xba\x16\xd6\x49\xba\xf0\xc7\x37\x8c\x11\x5e\x08\xf5\xe1\x06\x5f\xbc\x5b\x01\x7f\xed\x41\x3f\xc1\x3d\x01\xa5\x50\x23\x2f\xd6\xdc\x6f\x37\xea\x60\x51\x7c\x58\xa1\xea\xde\x5d\xae\xdb\xce\xc2\x1b\xdd\x6d\xb5\x16\x9a\xac\x94\x72\x56\x50\x6b\xe9\xec\x38\xdb\xad\xcd\x15\x18\x2c\xb8\xc5\x93\x18\x98\x6d\x21\x93\x09\x77\x9d\x28\x9c\x19\xd4\xa3\xdf\xc6\xce\x3b\x9e\x4b\x51\x36\xeb\xb1\xf1\xde\x66\x75\xdb\x41\xcd\x33\x6f\x85\x32\xf5\x71\x46\x77\x34\xd2\xce\x46\x61\xec\x26\xe5\xf9\x95\x30\x1d\x77\x0f\xef\xc1\x3b\x75\x85\x4c\xc8\x51\xc9\x46\x60\x56\x9e\x41\xc1\x6d\xff\xec\x89\x16\x99\x05\xde\x90\x46\xe7\xf6\x6e\x11\x6b\xc4\xf5\x5f\x23\x90\xb6\x30\x6a\xdb\x6b\x35\x9d\xd6\x8b\xd2\x22\x7a\x6b\x29\x0a\xa6\x81\xb7\x09\x23\x67\x12\x0b\x87\xcf\xb8\xe0\xdd\x7a\xb7\xb2\x59\x34\xb2\x49\xf9\x82\xcb\xd4\xbd\xb1\x52\x59\xba\xc7\x74\xce\x74\x6e\x91\xa8\x56\xae\x1b\x29\xb6\xcb\xb9\xd2\xba\xf5\x96\x79\xcb\x31\x33\x86\x6b\x5c\x3e\xde\x88\x1a\xef\xda\x8d\xe7\x1b\xaf\x6c\xb4\x1d\xda\xa0\x04\x5a\x99\x36\xca\xfc\x55\x30\x1b\x03\xb5\xb1\x69\x8b\x9b\xc0\x83\xd1\x29\xee\xab\xfb\x8e\xc1\xd9\xdf\xfb\xe7\xe7\x83\xd1\xe9\xf0\x1f\x3f\xe1\x10\x5a\x6b\x41\x51\x64\x1b\x7f\x7c\x21\x3c\xba\x87\xbf\xd9\xa6\x5c\x35\xbe\x24\x00\x98\xdd\xf2\x85\xd8\x1f\xa3\xe8\x5a\x13\x6a\xb5\x5a\x89\x8c\xeb\x22\x43\xb4\x76\x6c\x2e\x6e\x99\xfc\x42\xf0\x2c\x35\xc0\x65\x92\x29\xe3\x40\x7f\xae\x59\x72\xc9\x4b\x03\x07\xff\xfc\xd7\x41\x4b\x52\x32\x96\xd4\xbb\xdd\xa6\x9e\x4c\x16\x55\x3d\xeb\x0b\x98\x74\x0f\x0e\x4f\x95\xfc\xa6\x39\x2f\x10\xac\xd1\xba\xf0\xff\x73\x04\x96\xad\x5b\x9a\x6a\x56\xaa\xca\x52\x54\xf1\x9b\x76\x78\x76\x10\x6c\xdb\x81\x6f\x16\xd7\x8a\xd9\xc8\x92\x7d\x6a\x1c\xa1\x96\xd4\xbb\x06\xf4\xe0\x03\x07\x96\x19\x05\x9a\xbb\xa7\xbd\x9d\xb4\x46\x71\xfb\xac\x9b\x37\xc6\x58\x8d\xd5\xd1\x2e\xab\x66\x16\xf5\x66\x5c\xbb\x56\xc3\x43\xbb\xee\x50\x73\xed\x1a\xc4\x17\x0f\x0a\x2d\xac\xe1\x1a\x31\xf8\x00\xf7\x8a\xae\xe7\xd3\x1f\x7e\xc1\x66\x72\x66\x44\xe3\x8f\xf7\x92\xab\xfd\xae\x8d\x79\xa6\x35\x72\x30\x9d\xac\xc4\xba\x46\xca\xd6\x99\xf8\xcf\xcd\x66\xb3\xf9\x17\xfc\xd3\xb6\x5b\x2d\xb6\xbd\xac\xff\xb2\x8f\xfb\x49\x92\x06\x9c\xa9\x3b\x7d\xe2\xf0\x40\xa8\x3f\x05\x5e\x9f\xb9\x3c\xfa\x19\x8b\xa8\xf9\x08\x02\x81\xdb\xbe\xbc\xf9\xbc\x56\xe3\x85\xf4\x34\xd4\x42\x63\x33\xa3\x1a\x15\x27\x60\xfd\xee\xa8\x7a\xc7\x4e\xdc\x4e\x64\x56\x36\x07\x5d\x3f\x73\xe4\xd4\xdf\xe5\x79\xf1\xaa\xf7\xd2\xbe\x72\x1b\x0d\xfd\x3a\xdd\xc3\x9f\x39\x8b\x42\x2b\x65\x47\x5e\x75\xf3\x84\xe9\x3c\x70\x9d\x06\xfe\x40\xf5\xbb\x56\xbc\xad\xd8\xa6\x9c\x77\x9a\x50\x4f\x72\xab\xd6\x2c\x44\x02\x19\x93\xcb\x8a\x2d\x39\x2c\xd5\x9a\x6b\xb9\x7d\xb2\xcf\x5b\x4b\x5a\x7d\xdd\xec\xf6\x6b\xe7\xf0\x71\xc1\x3f\x1d\x67\x6a\x19\x45\xd1\x21\x2e\xcc\xe0\x52\xd4\xd1\x9e\x5b\x51\x3f\xc0\xec\x57\x78\xa7\x32\x75\xc5\x36\xc9\xaa\xba\x84\xbf\x64\xe5\xcf\xe5\x1f\xbf\x94\x7f\xac\xda\x2f\xb1\x60\xf6\x97\x65\xf9\xf3\x23\x5d\x74\x88\xb6\x1c\x0d\x77\xbd\xe8\x10\x7d\x53\x1f\xa4\xff\xa6\xbd\xe8\x00\x77\xbb\xe8\x10\xdd\x78\xd1\x01\x6e\x75\xd1\x21\xba\xc5\x45\x07\xb8\xf9\xa2\x43\x74\xbb\x8b\x0e\x70\xf3\x45\x87\xe8\x3f\x76\xd1\x21\xda\xba\x2f\xf5\x88\x17\x1d\xbe\xb1\x6b\xef\x9b\xcf\x5c\x74\x88\xda\x8b\x0e\x70\xcf\x8b\x0e\xd1\x0e\x4b\xbd\xd7\x45\x87\x68\xef\x45\x07\xb8\xdb\x45\x87\xe8\x9a\x8b\x0e\x70\x97\x8b\x0e\xd1\xe7\x2e\x3a\xc0\xe7\x2e\x3a\x44\x77\xb9\x2f\x75\xed\x45\x87\xe5\x4a\xa5\xc6\x1c\x6f\x58\x9e\x45\x6e\xa2\x05\x68\x03\x87\xef\x87\xb3\x7d\x90\xf3\x3d\x4c\x59\x0e\x6f\xf1\x5d\xba\x38\xf5\x9c\x2f\x4e\xd1\xfd\xcb\xe7\x77\x71\x6a\xe7\xfe\xe5\x0e\x3c\xbc\xb2\x7c\xe2\xad\x82\xbe\x3b\xe3\xdb\x83\x7e\x96\xd5\xdb\xb2\xe6\x86\xeb\xb5\x3d\xf9\xd1\x3d\xf3\xe1\x8c\x6b\xce\xce\xec\xcf\x72\xe2\x37\x73\x21\x99\x76\x67\x52\xcd\xee\x89\x8d\xd0\x1b\xe4\xae\xcb\xb8\xd5\x63\x79\x40\xe7\x68\xc6\x75\xeb\x01\x5f\x8a\x72\x5e\xba\x33\x19\xff\x6f\xeb\x20\x8a\x9d\xbe\xe1\xc9\xd2\xe0\xec\x42\xeb\x8d\x68\xd7\x51\x54\xd3\x3a\xbb\x88\x32\x61\x4a\x47\x0c\xda\xda\x64\xba\xd5\x94\xb4\x31\x7a\xf7\xf6\xb7\x40\xc8\x50\x08\x75\x0b\xea\x63\x31\x4d\x23\xa2\xed\xc5\x7c\xbf\x46\xd4\xa0\xd9\x45\x5a\x0f\x5f\xfe\x24\x3c\x2b\xb9\x16\x2c\x0b\x2e\x02\xd4\x9c\x23\xea\x9c\xe2\x71\xfd\x19\x79\x15\x1b\x4b\xad\xa9\xd2\x5b\xa5\x96\x19\x87\xa1\x4c\x7a\x20\x55\xfb\x9b\xa9\x4f\x2e\x04\xc7\x98\x8c\x25\x2e\x73\x6b\xce\xb3\x60\xc7\x65\xaa\xb4\x33\xed\x15\x5a\xe5\xaa\xe4\xb5\xd3\xc3\x74\x6e\x03\x44\xdd\xad\xa4\x86\xf7\x46\x77\x2f\xb4\x08\x8e\x33\xb7\x88\x67\x21\x6a\x38\xdd\x8f\x51\xaf\x3f\xda\xd5\xb1\xbb\xb0\xbd\xcd\xc1\x5d\xaa\x1b\x4f\xa6\x51\xcd\x4c\xf0\x07\x5c\xe7\xbb\x6a\x50\x00\x4a\x01\x82\xc5\x35\x84\x45\x2d\x84\xc5\xb6\xd2\xdd\xd7\xf6\x60\x99\xad\x2f\x80\xb3\x68\x3f\x9c\x4d\x06\x70\x3a\x9c\x5a\xec\x19\x9c\x5e\x83\x64\x6d\x2f\xa3\xf1\x87\xd1\x60\xe2\x79\x54\xd3\xc5\x3d\x60\x76\x3a\x9c\x0c\x10\x8f\x86\xa3\xf6\xaf\x93\xe1\xe9\x60\x34\xeb\x9f\xc5\xd1\xf4\x7c\x70\x32\xec\x9f\x21\x84\x0f\xde\x9f\x9f\xf5\x27\x1f\x63\x5f\xe6\x74\xf0\xdf\x17\x83\xd1\x6c\xd8\x3f\x6b\x80\xf0\xf0\x33\x12\x39\x9f\x8c\x4f\x2e\x26\x16\x89\x51\x0c\xd3\x8b\xd7\xd3\xd9\x70\x76\x31\x1b\xc0\xdb\xf1\xf8\xd4\xca\x79\x3a\x98\xfc\x3e\x3c\x19\x4c\x7f\x86\xb3\xf1\xd4\x0a\xeb\x62\x3a\x88\xa3\xd3\xfe\xac\x6f\x2b\x3e\x9f\x8c\xdf\x0c\x67\xd3\x9f\xf1\xef\xd7\x17\xd3\xa1\x95\xd9\x70\x34\x1b\x4c\x26\x17\xe7\x08\x8f\x47\xf0\x6e\xfc\x61\xf0\xfb\x60\x02\x27\xfd\x8b\xe9\xe0\xd4\x0a\x77\x3c\xb2\x7a\xe0\xec\xdd\x60\x3c\xb1\x5b\xd3\x7e\xa4\x6e\xc1\x79\x3a\x9b\x0c\x4f\x66\xe1\x63\x88\xb1\xe3\xc9\x2c\x6a\xfb\x08\xa3\xc1\xdb\xb3\xe1\xdb\xc1\xe8\x64\xd0\xc1\xf1\xa3\x06\xc7\x2d\xf8\x7f\x84\x0f\xfd\x8f\xb5\x72\xe8\x61\x3a\xb2\x7f\x06\x13\x36\xb6\x03\x09\xc3\x37\xd0\x3f\xfd\x7d\x88\xcd\xf6\x0f\x9f\x8f\xa7\xd3\xa1\x9f\x26\x56\x64\x27\xef\xbc\xb8\xb7\x15\xc3\x3f\xb8\x31\x7c\x9d\x5e\x1e\x2f\xd5\x8b\x45\xc6\x96\x66\x3f\xd0\xff\x8a\x8f\xc1\x9a\x49\x48\xb9\x84\xdf\x04\x9f\x2b\x95\x5f\x83\xf7\x5f\x23\xdc\x3b\x37\xcb\x9f\x87\xf9\xde\x10\xf7\x28\xc0\xef\x2d\xba\x0f\x41\x7f\x67\x57\x7e\xa4\x2d\xc0\xf7\xed\x11\xf6\x01\x6f\xf3\xa4\xcd\x80\x36\x03\xda\x0c\xee\xb3\x19\xa8\x95\x3a\x5e\xaa\x54\x95\x5c\xae\xf7\x6d\x04\xdf\xc1\xaf\x6a\x25\xe1\x35\xd3\xa5\x92\x5b\x61\x9c\x9e\x86\xbd\x71\x9f\x7d\x80\xec\x8d\x5f\x98\xbd\xf1\x76\xf6\x81\x3f\xcd\xde\x18\xdd\xce\x3e\xf0\x39\x7b\x63\x74\x3b\xfb\xc0\xe7\xec\x8d\xd1\x9d\xec\x03\xd7\xd9\x1b\xa3\x7b\xd9\x07\xba\xf6\xc6\x10\x4a\x2e\x99\x4e\x05\x93\xca\x1c\x2b\xc3\x3f\x95\x64\x3e\x20\xf3\x01\x99\x0f\x48\x63\x24\x8d\xf1\xeb\xd2\x18\x73\x56\x96\xf2\x78\xa9\x5e\xe8\x4a\xf2\x2b\x91\x96\xab\x3b\xb8\x97\x7e\x84\x8f\xcc\x54\x2b\xa1\x15\xbc\x67\xa5\xa9\x72\x55\x2a\xf2\x33\x91\x9f\x89\xfc\x4c\xcf\xda\xcf\x14\xe2\x8b\xca\xf8\x65\x75\xa9\xe4\xa5\x3a\x2e\xb7\x4c\x93\x27\xde\x45\x3d\xdf\xc0\x18\x9f\x82\xdf\xf0\x31\xc2\x8f\xe7\x8c\x1f\x14\xe0\xf3\xf9\xe1\xc7\x4e\x80\xcf\x00\x3e\x8a\xcb\xe5\x31\xd7\x5a\xe9\xbd\x6e\x8d\x1f\x62\x38\x65\x6b\x0e\x27\x2b\x2e\xf9\x06\xfe\x96\xb2\x35\xff\x25\xb1\x1f\x7a\x92\x97\x7f\x8f\x9e\x38\x1f\x05\xc7\x47\x1f\x4a\x46\x3b\x3c\x30\x82\xbb\xd3\xd1\x3d\x2d\xb8\x25\x19\xdd\x6d\x44\x04\xf7\xa2\xa3\x7b\xfd\x19\x11\xdc\x9a\x90\x6e\x39\x33\x1e\x83\xd8\xd5\xe8\x11\xdd\x9b\xd8\xc1\x16\xb1\x8b\xee\x45\xec\xae\x81\x8f\xc9\x20\xba\x03\xb1\xf3\xbd\xbc\x9e\xd9\x45\xb7\x62\x76\x70\x1b\x66\x17\xdd\xc0\xec\xe0\x6e\xcc\x2e\xda\xcf\xec\xe0\x1e\xcc\x2e\xda\x61\x76\xf0\x00\x66\x17\x79\x66\x07\x4f\x96\xd9\xe9\xd4\x54\xf3\x15\x33\x48\xee\x78\xc6\x4c\x29\x92\xd2\xc6\x08\xfa\x7c\xf6\x86\xa9\x7d\x11\x4e\x99\x4c\x59\xc1\xa4\x20\x9d\xec\x39\xeb\x64\xc4\xe9\x9e\x9f\x4e\x76\x23\xa7\x33\xac\x54\x5a\x1c\x2f\x55\xaf\xaa\x44\xba\x8f\xd4\x7d\xf7\xc2\xe2\xc8\x7c\x03\xef\xd9\x27\x91\xc3\x6b\xbb\x72\xe0\x6f\xf3\x5f\x50\x9f\xc9\x95\xbc\xe4\x9b\x9e\xae\xfe\x4e\x4e\x47\x72\x3a\x92\xd3\xf1\x79\x3b\x1d\x03\x60\x29\x97\x55\xc6\x12\x23\x8e\xaf\x34\x2b\xf6\x11\xbe\xef\x63\x78\xad\x99\x4c\x95\x84\xd9\x4a\xe5\x46\xc9\xff\x18\xc5\xeb\xc4\xaf\x89\x23\xa2\x78\x0f\x6e\xc4\x4d\x47\xd6\x22\xf8\x53\x29\x9e\x5d\x9f\x8f\x42\xf1\xea\xd7\xa2\x27\x48\xf1\x9c\x97\xf0\x51\x28\x5e\x8d\x37\xd1\xc3\x29\x5e\xeb\xbc\x8b\x1e\x4a\xf1\xba\xce\xbb\xe8\x81\x14\xef\x71\x9d\x77\xe0\x29\x5e\xf4\x88\x14\xaf\x14\xe9\x15\xcb\xb2\xe3\xe5\x1f\x88\x84\x77\xf1\xdb\xfd\xaa\xcc\x0a\x5e\xb3\x4b\xae\x1f\x97\xdb\x81\x5a\x44\x8f\xc7\xed\x40\xc8\xe8\xd1\xb8\x1d\xaa\x35\x8f\xc3\xed\x5a\x15\x2c\x7a\x00\xb7\xdb\xa7\x82\x45\xc4\xed\xee\xc9\xed\x6a\x34\x8d\x1e\x81\xdb\x05\x2a\x58\xf4\x00\x6e\xb7\x07\x71\xa2\xfb\x73\xbb\x8e\x0a\x16\x3d\x94\xdb\x5d\xaf\x82\x79\x48\xc9\x59\x99\xdc\xed\x28\x00\x41\x0a\x41\x0a\x41\x0a\x41\x8a\x87\x14\x55\x5c\x2e\x7b\x42\xda\x1b\xeb\xbd\xf5\xab\x0e\x9b\x7b\xf5\xf2\xdb\x6f\x9d\x85\xe8\x84\x49\x65\x73\x48\xc1\x59\x89\xbc\xed\xa1\x41\x58\x1e\x1c\x81\xe5\x4e\xe1\x57\xee\x12\x7b\xe5\x31\x02\xaf\x3c\x3c\xea\xca\xa3\x85\x5c\x79\x8c\x78\x2b\x9f\x09\xb6\x32\xeb\xa0\x95\xc3\xf5\x2b\x8e\x3c\x5c\x69\x1f\xf2\xed\xad\x72\xf1\x7a\x4f\xfc\xcf\x6a\x01\x99\x98\xe3\x94\xf3\xc8\xb9\xaa\x8c\xa5\xee\xa6\x14\x16\x64\xd7\x5c\x07\xc9\x0e\xda\xdc\x45\x2d\xe8\xe1\x6b\x7e\x04\x7d\xf4\x77\x56\x88\xa4\xb7\x74\x89\x6e\xb8\x35\x00\xe8\xe6\x73\xc1\xb4\x09\x3e\x6a\xce\xd2\xe0\xa3\x49\x98\x94\xc1\x67\x1c\xd1\xe0\x23\xb6\x73\x15\x7e\x28\x6c\x10\x36\x6e\xbf\xdb\xd9\x62\x5f\xfe\x08\xbf\x09\x8d\xdd\x98\x8a\x5c\x49\xb5\xa6\x6d\xf6\xbe\xdb\x6c\xf4\x30\xaf\xcc\xf6\x36\x1b\x91\x57\x86\xbc\x32\x7f\xba\x57\x86\x32\x6a\x93\x73\x97\x0e\xdc\x11\x8c\x50\x46\x6d\xca\xa8\x4d\x19\xb5\x29\xa3\xb6\x2b\x88\x32\x6a\x53\x46\x6d\xca\xa8\x4d\x19\xb5\x81\x32\x6a\x53\x46\x6d\xca\xa8\x4d\x19\xb5\x29\xa3\x36\x65\xd4\xfe\x42\x73\x36\x53\x46\xed\xa7\x3c\x3a\x94\x51\x9b\x32\x6a\x53\x46\x6d\xca\xa8\x4d\x19\xb5\x29\xa3\x36\x65\xd4\xa6\x8c\xda\x94\x51\x9b\x32\x6a\x53\x46\x6d\xca\xa8\x4d\x19\xb5\x29\xa3\x36\x65\xd4\xa6\x8c\xda\x94\x51\x9b\x32\x6a\x53\x46\x6d\xca\xa8\x4d\x19\xb5\x29\xa3\x36\x65\xd4\xa6\x8c\xda\x94\x51\x9b\x32\x6a\x53\x46\x6d\xca\xa8\x4d\x19\xb5\x29\xd8\xd4\x97\x17\x6c\x8a\x32\x6a\x53\x46\x6d\xba\x38\xf5\x8c\x2f\x4e\xd1\xfd\xcb\xe7\x77\x71\x8a\x32\x6a\x53\x4a\x2c\x4a\x89\x45\x29\xb1\x28\x25\xd6\x53\x89\xaa\x47\x19\xb5\x29\xa3\x36\x65\xd4\xa6\x8c\xda\xb4\x19\xd0\x66\x40\x19\xb5\x29\xb8\xfd\x57\x6b\x6f\xa4\xe0\xf6\x5f\x64\x70\x7b\xca\xa8\x4d\xe6\x03\x32\x1f\x90\xc6\x48\x1a\xe3\xb3\xd1\x18\x29\xa3\x36\xf9\x99\xc8\xcf\x44\x7e\x26\xca\xa8\x4d\xf8\xf1\x04\xf0\x83\x02\x7c\x3e\x3f\xfc\xa0\x8c\xda\x5f\x53\xba\x35\xca\xa8\x4d\x19\xb5\x29\xa3\xf6\x13\x60\x76\x94\x51\x9b\x74\x32\xe2\x74\xa4\x93\x51\x46\x6d\x72\x3a\x92\xd3\x91\x9c\x8e\x4f\xd5\xe9\x48\x19\xb5\xbf\x34\x8a\x47\x19\xb5\x29\xa3\x36\x65\xd4\xa6\x8c\xda\xcf\x3c\x2f\x1f\xa5\xbf\x7d\x42\xdc\x8e\xd2\xdf\x52\x46\x6d\x82\x14\x82\x14\x82\x94\x27\x02\x29\x94\x51\x9b\x32\x6a\x53\x46\x6d\xca\xa8\xfd\x14\xb7\x59\xca\xa8\xfd\x84\xb6\x59\xf2\xca\x50\x46\x6d\x72\xee\xd2\x81\x3b\x82\x11\xca\xa8\xbd\xfb\x8f\x32\x6a\x53\x46\x6d\xca\xa8\x4d\x19\xb5\x29\xa3\x36\x65\xd4\xa6\x8c\xda\x94\x51\x9b\x32\x6a\x53\x46\xed\xeb\x94\x17\xca\xa8\x4d\x19\xb5\x29\xa3\x36\x65\xd4\xa6\x8c\xda\x94\x51\xbb\x4d\xf3\x44\x19\xb5\x29\xa3\x36\x65\xd4\xa6\x8c\xda\x94\x51\xbb\x2d\x84\x32\x6a\x53\x46\xed\xe6\x1f\x65\xd4\xa6\x8c\xda\x94\x51\x9b\x32\x6a\x53\x46\x6d\xa0\x8c\xda\x94\x51\x9b\x32\x6a\x53\x46\x6d\xca\xa8\x4d\x19\xb5\x5d\x39\x94\x51\x9b\x32\x6a\x53\x46\x6d\xca\xa8\x4d\x19\xb5\x29\xa3\x36\x65\xd4\xa6\x8c\xda\x94\x51\x9b\x32\x6a\x53\x46\x6d\x0a\x36\x45\x19\xb5\x29\xa3\x36\x5d\x9c\x7a\xc6\x17\xa7\xe8\xfe\xe5\xf3\xbb\x38\x45\x19\xb5\x29\x25\x16\xa5\xc4\xa2\x94\x58\x94\x12\xeb\xa9\x44\xd5\xa3\x24\xaa\x14\xcf\xf8\x6b\xa5\x98\x14\xcf\xf8\x8b\x8c\x67\x4c\x49\x54\x49\x63\x24\x8d\x91\x34\x46\xd2\x18\x9f\x8d\xc6\x48\x09\x31\x48\x81\x24\x05\x92\x14\xc8\xc7\x4f\x88\x41\x01\xde\x29\x1a\x33\xb9\x29\x28\x1a\x33\x05\x78\x27\x48\x21\x48\x21\x48\x79\x42\x90\x42\x01\xde\x29\xc0\x3b\x05\x78\xa7\x00\xef\x4f\x71\x9b\xa5\x00\xef\x4f\x68\x9b\xa5\x03\x46\x14\xe0\x9d\xce\x29\x52\x80\x77\x82\x11\x0a\xf0\xbe\xfb\x8f\x02\xbc\x53\x80\x77\x0a\xf0\x4e\x01\xde\x29\xc0\x3b\x05\x78\xa7\x00\xef\x14\xe0\x9d\x02\xbc\x53\x80\xf7\xeb\x94\x17\x0a\xf0\x4e\x01\xde\x29\xc0\x3b\x05\x78\xa7\x00\xef\x14\xe0\xbd\x8d\x3a\x46\x01\xde\x29\xc0\x3b\x05\x78\xa7\x00\xef\x14\xe0\xbd\x2d\x84\x02\xbc\x53\x80\xf7\xe6\x1f\x05\x78\xa7\x00\xef\x14\xe0\x9d\x02\xbc\x53\x80\x77\xa0\x00\xef\x14\xe0\x9d\x02\xbc\x53\x80\x77\x0a\xf0\x4e\x01\xde\x5d\x39\x14\xe0\x9d\x02\xbc\x53\x80\x77\x0a\xf0\x4e\x01\xde\x29\xc0\x3b\x05\x78\xa7\x00\xef\x14\xe0\x9d\x02\xbc\x53\x80\x77\x0a\x9e\x42\x01\xde\x29\xc0\x3b\x5d\x9c\x7a\xc6\x17\xa7\xe8\xfe\xe5\xf3\xbb\x38\x45\x01\xde\x29\x5c\x27\x85\xeb\xa4\x70\x9d\x14\xae\xf3\x6b\x0e\xd7\xf9\x07\x37\x86\xaf\xd3\xcb\xe3\xa5\x7a\xb1\xc8\xd8\xd2\xec\x07\xfa\x5f\xf1\x31\x58\x33\x09\x29\x97\xf0\x9b\xe0\x73\xa5\xf2\x6b\xf0\xfe\x6b\x84\x7b\xe7\x66\xf9\xf3\x30\xdf\x1b\xe2\x1e\x05\xf8\xbd\x45\xf7\x21\xe8\xef\xec\xca\x8f\xb4\x05\xf8\xbe\x3d\xc2\x3e\xe0\x6d\x9e\xb4\x19\xd0\x66\x40\x9b\x01\x65\xfb\x20\x7b\xe3\x17\x6a\x6f\xa4\x60\xcd\x5f\x64\xb0\x66\xca\xf6\x41\xe6\x03\x32\x1f\x90\xc6\x48\x1a\xe3\xb3\xd1\x18\x73\x56\x96\xf2\x78\xa9\x5e\xe8\x4a\xf2\x2b\x91\x96\x77\x0b\xa3\xfd\x91\x99\x6a\x25\xb4\x82\xf7\xac\x34\x55\xae\x4a\x45\x7e\x26\xf2\x33\x91\x9f\xe9\x59\xfb\x99\x42\x7c\x51\x19\xbf\xac\x2e\x95\xbc\x54\xc7\xa5\xd9\x97\x4c\xe8\x7b\x98\x6f\x60\x8c\x4f\xc1\x6f\xf8\x18\xe1\xc7\x73\xc6\x0f\x0a\xf0\xf9\xfc\xf0\x63\x27\xc0\x67\x00\x1f\xc5\xe5\xf2\x98\x6b\xad\xf4\x5e\xb7\xc6\x0f\x31\x9c\xb2\x35\x87\x93\x15\x97\x7c\x03\x7f\x4b\xd9\x9a\xff\x92\xd8\x0f\x3d\xc9\xcb\xbf\x47\x4f\x9c\x8f\x82\xe3\xa3\x0f\x25\xa3\x1d\x1e\x18\xc1\xdd\xe9\xe8\x9e\x16\xdc\x92\x8c\xee\x36\x22\x82\x7b\xd1\xd1\xbd\xfe\x8c\x08\x6e\x4d\x48\xb7\x9c\x19\x8f\x41\xec\x6a\xf4\x88\xee\x4d\xec\x60\x8b\xd8\x45\xf7\x22\x76\xd7\xc0\xc7\x64\x10\xdd\x81\xd8\xf9\x5e\x5e\xcf\xec\xa2\x5b\x31\x3b\xb8\x0d\xb3\x8b\x6e\x60\x76\x70\x37\x66\x17\xed\x67\x76\x70\x0f\x66\x17\xed\x30\x3b\x78\x00\xb3\x8b\x3c\xb3\x83\x27\xcb\xec\x74\x6a\xaa\xf9\x8a\x19\x24\x77\x3c\x63\xa6\x14\x49\x69\x63\x04\x6d\x99\xfd\xf7\xb0\xba\xa9\x7d\x11\x4e\x99\x4c\x59\xc1\xa4\x20\x9d\xec\x39\xeb\x64\xc4\xe9\x9e\x9f\x4e\x76\x23\xa7\xa3\x0c\xb1\xe4\x74\x24\xa7\x23\x39\x1d\x1f\x3f\x43\xec\xb2\xca\x58\x62\xc4\xf1\x95\x66\xc5\x3e\xc2\xf7\x7d\x0c\xaf\x35\x93\xa9\x92\x30\x5b\xa9\xdc\x28\xf9\x1f\xa3\x78\x9d\xf8\x35\x71\x44\x14\xef\xc1\x8d\xb8\xe9\xc8\x5a\x04\x7f\x2a\xc5\xb3\xeb\xf3\x51\x28\x5e\xfd\x5a\xf4\x04\x29\x9e\xf3\x12\x3e\x0a\xc5\xab\xf1\x26\x7a\x38\xc5\x6b\x9d\x77\xd1\x43\x29\x5e\xd7\x79\x17\x3d\x90\xe2\x3d\xae\xf3\x0e\x3c\xc5\x8b\x1e\x91\xe2\x51\x46\x6d\x4a\x7f\x4b\xdc\x8e\xd2\xdf\x52\x46\x6d\x82\x14\x82\x14\x82\x94\xa7\x04\x29\x94\x51\x9b\x32\x6a\x53\x46\x6d\xca\xa8\xfd\x14\xb7\x59\xca\xa8\xfd\x84\xb6\x59\xf2\xca\x3c\xc0\x2b\xf3\xbf\x01\x00\x00\xff\xff\x48\xde\x69\xc1\x3c\xd4\x01\x00")

func tmpLicensesBytes() ([]byte, error) {
	return bindataRead(
		_tmpLicenses,
		"tmp/LICENSES",
	)
}

func tmpLicenses() (*asset, error) {
	bytes, err := tmpLicensesBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmp/LICENSES", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _tmpVersion = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xca\x49\x2c\x49\x2d\x2e\xe1\x02\x04\x00\x00\xff\xff\xcb\xe7\x49\xef\x07\x00\x00\x00")

func tmpVersionBytes() ([]byte, error) {
	return bindataRead(
		_tmpVersion,
		"tmp/version",
	)
}

func tmpVersion() (*asset, error) {
	bytes, err := tmpVersionBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "tmp/version", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
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
	"tmp/LICENSES": tmpLicenses,
	"tmp/version":  tmpVersion,
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
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
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
	"tmp": {nil, map[string]*bintree{
		"LICENSES": {tmpLicenses, map[string]*bintree{}},
		"version":  {tmpVersion, map[string]*bintree{}},
	}},
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
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
