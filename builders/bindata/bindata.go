// Code generated by go-bindata. DO NOT EDIT.
// sources:
// builders/bindata/pipdeptree.py
package bindata

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

var _buildersBindataPipdeptreePy = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xdc\x7c\xdd\x93\xdb\x36\x92\xf8\xbb\xfe\x0a\xc4\x79\x20\x15\x73\xf8\x9b\x71\xfd\xaa\x6e\x4b\x75\x4c\xce\x17\x6f\xb6\x52\x57\x67\xa7\x12\x67\xf3\x30\x3b\x45\x43\x22\x24\xc1\x43\x11\x34\x00\x6a\xac\x4c\xcd\xff\x7e\xd5\x8d\x0f\x02\x24\xf5\x61\x3b\xbb\xb9\x3a\x3d\xd8\x23\x12\x68\x34\xfa\x1b\xdd\x0d\xad\xa5\xd8\x91\xb2\x5c\x77\xba\x93\xac\x2c\x09\xdf\xb5\x42\x6a\xd2\x4a\xde\xe8\x72\xdd\x35\x2b\xcd\x45\x33\xb3\x4f\x85\x72\x7f\xa9\x83\x9a\xe1\x54\xae\x99\xd4\x42\xd4\xca\xcd\x5c\x6d\x29\x6f\xcc\xbb\x95\xa8\x6b\x86\x00\xfc\xdb\x8a\xad\x69\x57\xeb\x8a\xaf\xb4\x03\x45\xe5\xa6\xa5\x52\x31\x33\x47\xb4\x4c\x52\x2d\xa4\x9b\x40\xb5\x96\x1b\xa6\x35\x93\x6e\xfc\x7b\x25\x2c\x7c\xf3\xa0\xe6\x4b\x37\xd8\xfc\x57\xee\x44\xd5\xd5\x6c\x36\xd3\xf2\xb0\x98\x11\x42\xc8\x31\x6c\xde\xc8\x8a\x49\x56\xbd\x02\x6c\xd8\xc7\x15\x6b\x35\xf9\x11\xdf\xfc\x55\x4a\x21\x83\xb9\xc2\x0c\x04\xb4\xa7\xe6\x0e\x56\x6a\x79\x9b\x97\xbc\xd1\x4c\x36\xb4\x76\xe3\x37\x4c\x97\xbc\x51\x9a\xd6\x35\xab\xca\x8a\x2b\x2d\xf9\xb2\x43\x74\x8e\xcc\xcc\x0d\x2d\x60\x44\xbe\x96\x8c\xfd\xce\x1c\xac\x1f\xa4\xf8\x9d\x35\x3f\xb3\x0f\x1d\x97\x6c\xc7\x9a\x33\xd8\xb7\xbc\xbd\x00\x8b\x6c\x02\xac\xa3\x79\x7b\xbf\x29\x25\x53\xa2\x93\x2b\xa6\x66\x5f\x13\xde\xd4\xbc\x61\x8b\xd9\xd7\x66\x81\x8d\xa4\xed\x76\xcf\x7f\x77\xab\x2c\xe9\xea\x9e\x35\x55\x46\x5e\x71\x7c\x35\x9b\xcd\xca\x72\xcf\xa4\xe2\xa2\x29\x4b\x52\x90\xe4\x3a\xbf\x79\x91\xdf\x24\xb3\xd9\x6c\x5d\x53\xad\x59\x43\x0a\x23\x39\x39\x00\x2c\x41\xaa\xe8\x12\x78\x38\xab\xd8\x9a\x2c\x3b\x5e\x1b\x6c\x4b\xde\x54\xec\x63\xda\xde\x6f\xd4\xdc\xec\xf0\xd9\xb3\x67\xff\x09\xaf\x09\x6d\x08\xbe\x04\x64\x15\x59\x1e\x88\xde\x32\x2e\xc9\x3d\x3b\x10\xaa\x08\x25\xc0\xbc\x7c\x86\x73\x16\x2d\x95\x74\x47\x6a\xae\x70\x6b\x6a\x61\xfe\x14\xeb\x78\xa3\xf9\xab\x80\x3e\x04\xc9\xd6\xc0\xfe\x11\x84\x64\xba\x93\x8d\x5a\xd8\x45\xc5\x1a\xd6\x0b\xd7\x86\xbf\x61\x75\x3b\x5c\x1f\x5a\xb6\x40\x24\x66\x0e\x6f\xfc\xdf\xc0\xc1\x17\x69\xda\xe6\xf7\xec\x00\x64\x53\xfa\x27\xba\xba\xa7\x1b\x96\xb6\xf3\x39\x59\x0b\x49\x5a\xc2\x1b\x04\x3f\xb7\x44\x59\x89\x46\x69\xd9\xad\x74\xa9\x25\x63\x29\xa2\xd1\xd3\xe4\x7b\xf7\x96\xc0\x5b\x22\x59\x2b\x99\x62\x8d\x46\x79\x8a\xb0\x45\x0e\xc2\x37\x84\x60\x29\xf4\x76\xcb\x00\x77\xe5\x46\xa2\xe4\x7b\x20\xbc\xd9\xe0\x53\x04\xfd\xc0\xeb\x9a\x2c\x19\x11\xcb\xf7\x6c\xa5\xcd\x8c\x43\xcb\x10\x4c\xb0\x11\x42\x9b\x0a\x27\xed\x69\xdd\x31\xe5\xa7\x39\xca\xff\xcc\x3e\xb8\x91\x16\x52\xcc\x2c\xa3\x7b\x80\x22\x50\x51\xd9\xbf\x09\x67\x79\xcf\x81\x11\xe7\x87\xbc\x42\x84\xdd\x40\x8b\x10\x97\xa4\x62\x2d\x6b\x2a\xd6\xac\xb8\xe7\xee\x65\xec\xca\xc8\x6d\x8f\x77\x2a\x33\x4b\xc3\x0d\xd3\xa9\x04\x56\xce\xe7\x38\x69\xf4\x01\x86\x4a\x64\x68\x2e\x8d\xba\xa9\x74\x7e\x37\x1e\xec\x19\x6f\xe0\x1a\xd2\xa5\x73\x27\x02\x4a\x48\xcd\x2a\xc3\x7f\xf8\xa7\x67\xff\x2f\x42\x6a\x35\xc1\xb9\x88\xfd\x30\xa5\x67\xb7\x14\x42\x93\xd6\x6c\x45\x81\xca\x3c\xb0\xba\x86\xff\x8d\x6c\x68\x26\x77\xac\xe2\x54\xb3\x60\x90\x64\x16\x07\x84\xc2\x1b\x1c\x4b\xeb\x76\x4b\x97\x4c\xf3\x15\xad\x8d\xe5\xf4\xf2\x66\x19\xdc\xd0\x1d\x9b\x62\x2f\x20\xb4\xf0\xba\xe3\xb9\x72\xb0\x7c\x5b\x6a\xca\x1b\x56\x01\x87\x57\xb4\xae\x79\xb3\x99\xa6\xee\xbb\x58\x35\xde\x11\xef\xc7\x62\x69\x30\xa8\x1b\x32\x84\x5c\x0f\x3c\x45\x1e\x99\xf9\x09\x41\x08\xde\xa7\x06\x5e\x7a\x9b\xde\x67\x16\x76\xba\xcf\x40\x0c\x8b\xde\x8b\xa5\xc9\x3d\x3b\x24\xf3\x63\x82\x11\xb3\xfe\x3e\x23\x7b\xa4\xaa\x64\x2c\xe7\x9a\xed\x40\x4a\xb2\x33\x33\x61\xbd\x9a\xee\x96\x15\x25\xf7\xfb\x05\xb9\xdf\xdf\x5e\xdf\x59\x59\x34\x52\xb3\xe6\x8d\x91\x99\x12\x58\x8e\x82\x83\x58\xf6\xd2\xf3\x03\x6f\x2a\x42\x8d\x44\xf0\x86\x50\x43\xff\xe5\x81\x70\x9d\x28\xb4\x69\x7f\x12\xeb\xcc\x82\x4a\xa3\x72\x2f\xd0\xb6\x5b\xd1\x42\x5c\x1b\x51\x31\xa2\x05\xee\x70\xc0\x6b\x1a\x8c\xe0\x6b\xb2\x16\x5d\x53\x11\x56\x2b\x46\x5e\x8b\x26\x66\xff\x8e\x7f\x64\xd5\x90\xd9\xaa\xab\x35\x29\xc8\x6d\xdb\xeb\x24\x72\x05\x8c\x4c\x3a\x07\x90\x68\xbb\x49\x51\x00\x56\x77\x38\x8b\x2a\xc5\xa4\x26\x35\x6b\x52\x03\x60\x0e\xd3\x6e\xaf\x33\x72\x73\x17\xca\x10\x60\x00\x10\xc2\x81\x45\x41\xae\x0d\x7e\xe6\xc9\xed\xf5\x9d\x65\x9f\x64\xe0\x49\xd9\x94\xd6\xff\x6c\x5e\x19\xbd\x8f\x59\x60\xb5\x0d\xcc\xa5\x1e\xd8\x76\xb3\x00\xd8\x74\xe4\x24\x28\xf5\x94\x2d\x0f\x0c\xf4\xc0\x94\xc3\x0c\x30\xe3\x38\x3e\x34\xf9\xc7\x0d\xf9\x1f\x24\x2e\x17\xab\xba\x25\xda\x84\xb2\x4f\x98\x78\xc4\xa1\x08\x63\xd5\x14\xb6\x67\x54\x76\xb5\x85\x58\x04\xe9\x57\x10\xc5\x74\xba\x42\xbe\x83\x54\xac\x80\xbd\x36\x9e\x41\xce\xf4\x06\xdb\xcc\x75\x3a\xad\x86\x4a\xbd\x98\x85\x6a\x8f\x3a\xbf\x57\x8b\x68\xcf\x28\xb9\xc5\x50\x79\xa5\xd1\xde\x3d\x2a\x38\x81\xb9\xd1\x24\x7c\x7d\x0b\x53\xef\x72\xda\x02\x91\xd3\xfb\x9c\xaa\xd2\x3a\x9d\xaa\x5c\x1e\xd2\x7d\x60\x8b\xf8\x9a\xdc\xe3\x76\x1a\xa3\xf9\xfd\x66\x17\x13\x70\x43\x50\x10\x2e\xa6\xf3\x3b\x50\x91\x48\xb6\xa5\x71\x32\x28\xb9\x9b\x8e\x29\xe5\xe2\x40\x88\xe1\x4a\x8c\x76\x2c\x9d\x8b\xe4\xbb\xa4\x17\xe5\xbf\xc1\x58\x23\x64\x66\x3c\xc8\x16\x45\x69\x79\xd8\xb2\x06\x63\xda\x4a\x30\xd5\x24\x70\x58\x11\x7b\x0e\x7a\xad\x67\x43\x33\x61\x17\x89\x4c\x85\xf5\x42\xc3\xa1\x16\x8d\x85\xfb\xc3\x2f\xac\x85\xdb\x0c\x5f\x93\xae\x81\xb8\xf4\x88\x91\xb1\x33\x22\xf9\x82\x08\xb2\xd9\xc4\x12\xe6\x8f\x0b\xf0\xd9\x91\x22\x3e\xb8\x38\xca\x18\xb6\x1c\x8b\xec\x03\x1a\x5b\x84\xcd\xf0\x5a\xb1\xd1\x80\x0d\xd3\xe0\x81\xd2\x5d\x46\x92\x20\x10\x4f\x3c\xed\xc1\x35\xac\x6a\xaa\x14\x71\xb1\x8c\x51\xdd\x9e\x1f\x2f\x97\x4a\x4b\xba\xd2\xc4\x0c\x03\x31\x7d\x90\x20\x52\x12\xf4\x1f\xad\xa9\x33\x1a\x7a\x4b\x35\xf2\xc7\xd2\xc5\x87\x94\x5c\xd9\xd9\x0d\x63\x95\x02\x1a\x2e\x19\x51\xdd\x12\x1f\xb2\x8a\x3c\x70\xbd\x05\x52\xd4\x28\x4d\xe6\xe4\x03\x2b\xe1\xf4\x77\x12\x4c\x84\x2c\x41\xe4\x84\xd0\xef\xd0\x08\x05\x0f\x97\x92\x36\xab\xed\x3b\xb2\x63\x7a\x2b\x2a\xb7\x28\x50\x1c\xff\x00\xf9\x2b\x4b\xde\x70\x5d\x96\xa9\x62\xf5\x3a\x03\x7c\x03\xc5\x83\x67\x79\x29\x96\xef\x49\x01\x6f\xe2\xe7\xad\x14\xb0\xb7\x12\x02\x17\xf3\x3e\x7a\x14\x0f\x46\x37\x80\x63\xbc\xa7\x34\x76\x3b\xc4\xdf\xe2\xb0\xc6\x23\xd7\x7c\xc4\xb1\xd7\x42\xff\xe8\x08\xc1\x2a\x64\xfb\x14\x24\xb3\xe9\x2f\x87\x65\x21\xb4\x54\xb2\x46\x17\xe0\x8f\x1c\xb8\xe2\x07\x5a\x2b\x16\x00\xe5\x6b\x34\x0e\x66\xe8\xc0\x2a\x98\xf5\x90\x08\x83\xdd\x5a\xdc\xfc\xf0\x58\x4c\x8f\x4f\xb6\x1b\x74\xd3\x71\xc6\x7f\x28\x90\x8d\x95\x61\xb4\xdf\x88\x19\x52\x42\xa4\x9b\xc6\x9c\x5d\x4b\x52\x8c\xcf\xb6\xe6\xa0\x09\x27\x09\x18\x9e\x91\xdb\x20\xfc\x76\xb8\x68\x99\xae\xe5\x3c\x07\x15\x6e\xd3\x79\x28\x48\x56\xa3\xbc\x2c\xf5\xe1\xd3\x84\xda\x79\xd1\x32\xe3\x42\x38\x80\xae\x05\x32\x9e\x9f\xfc\xfb\xe3\xf5\x53\xfa\xec\xf1\xe6\xe9\xd9\xfc\xdb\x24\x5f\x0b\xb9\xa3\xda\x42\x2b\x51\x69\xca\x32\x2f\x51\x02\xcb\x32\xf3\xc2\xd7\x2b\x73\x78\x8e\xb4\xff\xf7\x1a\xfd\x9b\x51\xdf\x40\xa1\x2f\x3a\xfc\x5a\x24\xad\xe1\x14\xcb\xf7\x8b\x53\xf3\xb4\x40\x33\x41\xc4\x9e\xc9\x78\xa6\x64\x1f\x16\x44\xb4\x30\x8a\xd6\xe3\xd3\x1f\xcc\xa4\x4a\x89\x15\x9e\x38\xf4\x96\xab\x89\x08\x20\x8c\x35\xc0\x78\xe4\xc6\xc8\x70\x45\x3a\xc5\xd6\x5d\x8d\xbb\xaa\xb8\x6a\x6b\x7a\xb0\x67\xd6\x09\x28\xe8\xee\x79\xe3\x62\x84\x4b\xec\x46\x06\xe8\xa3\x9e\x84\x16\xa4\x6b\x99\x4c\x03\x9c\x0c\x4b\xe6\xb9\x9f\x0f\x72\x19\xdb\x0a\x67\x8b\x55\xcb\x56\xa4\xe8\x63\x51\x3f\x40\xb2\x0f\xa4\x80\xd5\x3e\xd5\x90\x58\x3d\x35\xcf\x27\x55\x2d\x79\xbc\x7e\x2a\x8a\xc7\x9b\xa7\x58\xb4\x42\xc3\x96\x45\x58\x5e\xaa\xbe\xbd\x70\x86\x4a\xe9\xb5\x60\xfe\xe9\x96\xcc\x46\xd2\x9e\x22\x5c\xe1\xe6\x22\x72\x9d\xd8\xb0\x31\x56\xe0\xf8\x1c\xa1\x1d\xa4\x88\x01\x53\x73\x20\x32\x08\x86\x4f\x5a\x7d\xbb\xfa\x60\x95\xc5\x48\xd4\x02\x90\xcf\x8b\xe1\xf0\x29\x6a\xa6\x23\x10\x9e\x67\xe4\xd6\x25\x0e\x16\xe4\xf1\xc5\xd3\x5d\x12\x0d\x9d\x5f\xc8\xd1\x2c\xc0\xe9\xd3\x8d\x73\x64\xd9\x3d\x4f\x07\x31\xe1\xc0\xb4\xe1\x29\x05\x81\xd1\x50\xe9\xa7\x12\x14\x3c\x32\x60\x2e\x76\x0a\xd0\x09\x52\x2f\x5e\xb8\x46\x21\x69\x86\xd9\xa2\x02\xd1\x98\x42\x12\x63\x60\x23\x76\x92\x7d\x38\x82\x69\x68\x68\x9c\x29\xec\xcd\x53\x85\xc6\x8a\xc8\x30\x71\xea\xa0\xa0\x41\x72\x23\xd1\x92\x42\xf8\xb3\x62\x4a\x51\x79\x98\xb0\x4f\xce\x1a\x05\x7b\x45\xab\x94\xf7\x20\xad\x01\x8d\xa8\xf7\x61\x61\x8f\x72\x1e\x85\xd8\x80\x82\x75\xec\x01\xf8\x80\x75\x6a\x5f\xb3\x90\x02\x23\x97\x18\xe9\x77\xe8\xd9\x80\x78\x11\x7d\xf1\xdc\x34\xed\xd8\x1e\x31\x09\xb2\xf0\x1e\x6b\x9c\xd2\x48\x6c\x90\x8e\x32\xeb\x46\x46\x72\x3c\x9e\xd2\xe7\xb5\xad\x74\xbb\x79\xf6\xeb\x93\xf7\x8b\x81\xe0\x5c\xe0\x16\x83\xa0\x41\x8d\xa8\x14\xfa\xc1\xb7\x5b\x46\xde\x85\xa3\xdf\xf5\xc2\x72\xdc\x11\x82\x78\x06\x9e\xf0\xdd\x71\x6f\x1a\x80\x73\x51\xf1\xe0\x83\x4a\x13\xca\x21\x09\x7d\xd9\xaf\xaf\xff\xeb\xf5\x9b\xdf\x5e\x97\x7f\xff\xeb\xcf\xbf\xfc\xf8\xe6\x35\x29\x48\xf2\x5d\x72\xda\xcd\xa1\xee\x4c\xfa\xb9\x9e\x86\x17\xb8\x39\x4c\xd8\x16\x08\xcd\xc6\x70\xad\x14\x2d\x93\xfa\xe0\x57\x0f\x0d\xf1\x50\x6c\xe0\x99\x72\x66\x18\xb5\x1c\x9f\x8c\xe2\xa5\x2c\xc9\xdf\x0b\xde\xa4\xb7\x89\xfd\x43\xb5\x26\x79\xae\x30\x61\x83\x93\xee\x30\x59\x63\x20\xf6\xf9\x9f\x23\x48\x8d\x44\x6a\x88\x99\x75\x3a\x7e\x93\x93\x66\x33\x3e\xf7\x7a\xa1\x37\xb3\x06\x4c\x19\x07\xa1\x0e\xb4\x13\xe3\x9e\x61\x5c\x95\x2b\xd1\xac\x6b\xbe\xd2\xbc\xd9\x4c\x58\xda\x1f\x83\x1d\xf8\xe3\xac\x9b\xa2\xcc\x79\xcb\x99\x41\xf7\x3e\x54\xfa\xaf\x49\xd7\xdc\x37\xe2\xa1\x99\x00\x03\x66\xad\x56\x02\xcb\x10\x1c\x73\xa0\x24\x40\x26\x24\x10\xee\x60\x44\x49\x52\x14\x93\x04\x98\xa4\xe0\x5b\xd9\xf5\xee\x36\x70\xe3\xe9\x38\x88\x72\x0b\x46\x0f\x91\xd1\x49\x12\x12\xf7\x83\x3f\x07\x1b\x17\x0f\x9e\xf5\x4c\x2c\xe4\x16\x8e\xc1\x98\x23\x63\xac\xb5\xe1\x39\x03\x8b\x9b\xe9\x60\xc1\x69\x36\x8f\x89\x64\xf3\x30\x76\x9d\x3f\x33\x06\x1c\xe1\x16\xc6\x0b\x8e\xe8\x47\x55\xe0\x5c\x5c\x88\x02\x5e\x46\x86\xe3\x74\x14\x12\x05\x62\x9f\x1c\x4f\x9e\x24\x0c\x32\xca\x59\x9b\x4b\x65\xeb\x65\x73\x48\x2e\x0f\xe1\x7c\xf4\x56\x2d\xc8\xe3\xcd\x53\xd6\xeb\xd7\x54\x34\x47\xce\x44\x74\x16\xe5\x4b\x18\xf5\x85\x81\xdd\xff\x1e\x9f\x3e\x7a\x31\x31\xd9\x47\x78\xd3\xf1\x00\xb2\xef\xc9\x67\xd5\x71\xf3\x3e\xa9\x9e\x61\x46\xbb\xa4\x75\x5d\x80\xf1\xc9\x88\xda\x8a\x87\x52\x34\xf5\xe1\x78\x7a\xc4\x94\x5d\xf7\x70\x52\xd1\xc2\x04\x72\x5a\xd8\x04\xe0\x20\xc2\x3d\x99\x11\xb7\x01\x59\x9f\xaa\x36\xc3\x96\x42\xd4\x1e\xab\x05\x79\xd8\x32\xbd\x65\x12\x96\xc0\x1a\x2a\xad\x6b\x33\x7d\x73\xaf\x08\xd5\xbe\x2a\x32\x15\x27\xd8\x4f\xcd\xf6\xac\x26\x42\x12\xd8\x17\xd1\x5b\x81\xb5\x03\x6a\x6a\x00\xd3\x87\x65\xf7\x51\xdd\xf2\x6a\x5c\x36\xb5\xd9\x54\xa6\x7b\x72\x2d\xf0\xab\x40\xcd\x81\x73\xbd\x2f\x1e\xda\xf4\xdf\xd6\xb8\x97\x33\xcb\x89\x4e\xb7\x9d\xee\xcf\xf7\x3e\x5c\xa2\x72\x93\xf5\xa9\x5b\x60\x4d\x3e\xa2\x99\x55\x73\x4f\x31\x21\x8d\xd3\xde\x8a\x07\xa4\x12\x96\x23\xa3\x9a\x78\x10\x7f\x0f\x3f\x30\xc6\x60\x83\xa4\x4a\x14\x59\xd3\xbd\xe8\xa4\x4b\x0b\xb7\xbc\x25\x57\x57\xa6\x5f\xc3\x60\xd2\xd7\x1b\xa7\x24\x21\x2a\xc6\x92\x38\x71\x3c\xcc\x1a\x63\x59\x62\x54\xf4\xc5\x97\xc6\xde\x85\x75\x09\xe9\xeb\x12\xf2\x5c\x5d\xa2\x11\x15\x83\x69\x41\x45\x0b\x9f\x77\x8a\x95\xcb\xae\xae\x99\x86\xb7\xbd\xc5\x34\x78\xdd\xb3\x43\xe9\x4a\x25\x58\x0e\xbf\x37\x61\xcd\x7e\x7e\xb4\x78\x69\xc0\x6e\x98\x2e\xb1\xb2\x20\xb1\xef\xc3\x16\x2b\x9b\x85\x87\x88\x05\xf4\xc6\x40\xbb\xbd\xb3\x26\x08\x8c\xaf\x17\x2a\xcf\x1e\x87\x7a\x58\x98\xc3\x67\x63\x06\xfa\x1a\x1d\x84\x82\x0e\x12\x08\x43\x1b\xe7\x77\xc3\xd7\xa6\x96\x81\x0e\x0e\xb6\xef\xb5\xef\xfc\xfa\xfd\x72\xd6\x87\x07\x1c\xba\x0b\x8c\x6a\xf7\x31\x85\xf1\x83\xe4\x2b\x07\xbd\xd2\xc5\x75\x66\xba\x62\x86\x21\x38\x5f\x9b\xe7\xa0\x0a\xf0\x2a\x36\xea\xe6\x55\x41\xb0\xf0\x13\xed\xed\x2e\xc2\xdb\x46\x3e\x38\xca\xe6\x81\x0d\x12\xde\x63\x86\x0b\x4e\x65\x7c\x5b\xc9\xd6\xfc\x23\x44\x4f\x24\xf9\xc6\xe0\x4c\x9e\x93\x34\xb9\x22\x09\x16\x4b\x02\xf9\x19\x85\x60\x03\x2c\x2c\xa8\xe7\xfe\x59\xe0\x60\x5c\xf1\xd5\xbd\xea\xb7\x11\x88\xd1\x2d\x90\x72\x95\x11\x43\x4d\x4b\x40\xf3\xdf\xf3\x17\xc7\x8b\xe6\x86\xbe\xf8\xef\xf3\xdb\x55\x4c\xad\xe9\x22\xbd\x2f\xf5\x85\x72\x8c\x4c\x9c\x1e\x0f\xcc\x8a\x25\xcc\x57\xd7\x28\x6f\xee\x86\xfb\x7c\x5e\xa0\x98\xa5\x4e\x61\xdd\x0a\xf3\x51\xcc\x68\x26\x18\x59\xaa\x79\x83\x72\xe8\x66\x21\x35\xda\xf9\x40\x28\xed\x8e\x5c\xf0\xf7\x8f\xc6\x9e\x93\x70\xf6\x3c\x76\x85\xef\x95\x68\xac\x2b\x34\x54\x1c\x79\x39\x15\x26\x2c\x30\x05\x02\xab\x63\x83\xde\xc0\xc8\x05\x5d\x45\xfe\xad\xef\x00\xa2\xbe\x07\x68\x4b\xd5\x96\xa9\x8c\x30\xba\xda\xe2\x17\xb2\xa5\x7b\x30\x99\x2f\xc8\x9a\xb3\xba\xf2\x55\xc8\xab\xa8\x88\x07\xdf\x43\x5f\xd4\xb7\x73\x45\x1e\xea\x88\xcb\x1d\x14\x9f\xc3\x51\xbc\xd1\x76\xef\x0b\xd2\x88\x1c\xbd\x58\x4b\x57\xc6\x79\x59\x71\xc7\x76\x44\x12\xd9\xf9\x09\x02\x5c\x6e\xe5\x2d\x6f\x00\x46\x5e\x75\xbb\x56\xa5\xb7\x8f\x2e\x72\x4a\x16\x04\x6b\xae\x68\x6e\xe7\xc7\x65\x3a\x09\xf7\x9d\x2c\xc8\xed\xbe\x9f\x15\x56\x98\xef\x9e\x8e\x81\x38\x5a\xa9\x3e\xda\x7e\x12\x29\xdc\x84\x28\x85\xa1\xd5\xc5\xf2\xd4\x30\xa5\x59\xf5\xcf\x90\x28\x58\x69\x2d\xea\x5a\x3c\xc0\xb7\x23\xd2\x15\xa6\x79\xaf\x7c\x2f\x1f\xfc\x3d\x0c\x2d\xfd\x8b\x51\x5c\xfa\x7f\x4f\x42\xff\xf9\x71\xc8\xd9\x2e\x9b\x29\x97\x4a\xfe\xf5\x21\xc9\x09\x07\xfe\x07\xbb\x6d\x3f\xb0\x72\x0e\xdb\x2b\xf4\x39\x1f\x5d\xdd\x8e\x0f\x42\x77\x0e\xca\xf0\x78\x3b\x7e\x38\x71\xbc\x1d\x1f\x23\x8f\xad\x51\xdd\x4e\x9c\xe0\xc2\xdd\xdc\xc6\xa6\x0a\xfb\x47\x22\xc8\xd6\xa9\x5b\xda\x1a\x42\x5f\xee\xb2\x2f\x75\xd5\xa7\x5d\xb4\x1f\x1a\x60\xee\x5a\x2f\x66\xc7\x8c\xf6\xb4\xff\xcd\xa6\xed\x24\xcc\x29\x5d\x17\xb5\x35\x92\xe6\x90\x51\x9a\x63\x7f\x91\x54\x42\x07\xed\x31\x6f\xcc\x09\x24\xb0\x0c\x38\x9b\x50\x45\x44\xc3\x9c\x2e\xab\xae\x6d\x4d\x8f\xe3\xdf\xe0\xed\xdf\xf9\xef\xee\xe8\x62\xa0\x1e\xef\xcc\x1a\x02\x8e\x8e\x77\xe6\x18\x13\xe1\xb7\x88\x01\x0f\xcc\xcd\x84\xa5\xb1\x35\x57\xc4\xb2\x65\x2b\xbe\xe6\xac\x9a\x84\xe1\x8d\x10\x44\xea\x4b\xde\x50\x79\x18\xc2\x33\xc8\x22\x4e\x4d\x78\x3e\xb3\x70\x06\xc6\xeb\x10\x76\x06\x5c\xd2\xbc\x4e\xce\x74\xe1\xe0\x5d\x89\x34\xe9\xe1\x98\xe2\x24\xdd\x53\x5e\xc3\xa9\x30\x23\xcb\x4e\x0f\x8a\x3e\x01\x96\xc3\x5c\x4f\x62\xce\xb6\x39\xf9\xa9\x66\x54\x31\xe7\x51\x08\xd7\x79\x92\x91\x35\xaf\x59\xa1\x0e\x2a\x57\xba\x62\x32\xc8\x23\xc2\x33\xf6\x91\xeb\xf4\xa6\x3f\x2f\x45\x2c\xf2\x46\xd3\xec\x2f\xff\xe1\xcd\xcf\xff\xfd\xf2\xed\x2f\xa3\x7d\x3c\x5e\x3f\xf9\x2d\x04\x12\x14\xd1\x34\xf7\xe9\xc2\x68\x89\x51\x48\x72\x14\x5b\xbb\xd6\x2f\x1e\xbc\x15\x48\x42\x25\x5b\x90\xc7\xeb\x3e\x1f\x19\x41\x4c\x32\xe2\x32\xfb\xa6\xbf\x76\xb0\x99\xf9\x7c\xfe\x29\x24\x32\x3a\x53\x38\x4e\xa7\x56\xd7\xe2\x3d\xcd\x9c\x21\x69\x5d\xb9\xa3\x62\xed\x89\x0e\xbe\x41\xaf\x90\x9d\x35\x5d\x39\xae\xe9\x92\xd5\x36\xf9\xfc\x8f\x26\x4c\xc3\xc6\x19\x3e\x07\x64\x94\xd1\x43\xbc\x73\xb0\x2d\x83\x19\x08\xb9\xc0\x7f\xfb\xd1\x58\x6e\x64\x68\x8e\x60\x0f\xb1\x0d\x77\xb8\x54\xac\x3d\x5e\x15\x77\xe7\x60\x18\x3b\x2e\x70\xfb\xed\xd0\x61\x36\xd4\xe0\xc9\xaa\xcd\x10\x4f\x58\xed\x04\xe6\x08\xe4\x6b\xf2\x12\x02\x35\x27\x82\x10\x35\x09\x27\x89\x19\x61\x7b\x86\x9d\x79\xde\xc4\x71\x6c\x0a\xf4\x7e\x27\x9f\xd6\x86\xa2\x20\x68\x53\xc7\x5d\x3b\x88\xaa\xc9\xe4\xfb\xf5\x95\x69\x7d\xf6\xe5\x10\x6b\x5f\x31\xa5\x35\x54\x0e\xd0\x1e\x6b\xa8\x4c\xa6\x29\x03\xab\x43\xb4\xb0\xc0\x2a\xb6\xc2\x2e\x64\x0d\xd6\xfa\xd7\xb7\x3f\xfc\x05\x5b\xd9\x30\x19\x81\x6a\x01\x6f\x44\x67\x74\xd5\x80\xc1\xe5\x30\xd7\x84\xad\x57\x42\x29\xbe\xac\x6d\xa2\x2b\x32\x68\xd1\x16\x5a\xde\xb2\x74\x9e\x9b\xe5\xd2\xa4\xd3\xeb\xab\xbf\x24\x51\x53\xe1\xaf\x0d\x87\x77\xaf\x70\xc4\x74\x6f\x61\x08\xca\x3a\x2b\x73\x37\xcc\x7b\x2b\xf4\x5d\x86\x00\xbd\x7f\x7a\xd5\xed\x5a\xd3\x86\x4c\x35\x25\x1b\xd6\x30\x89\x85\xf2\xe5\xa1\xe7\x13\xe6\x46\x2b\xd1\x0d\x6e\xe5\x04\xf0\x4c\x49\xd5\x11\x17\x8c\x75\xe4\x28\x23\xc3\xce\x31\xd8\xc7\x7e\xab\x00\x42\x46\x12\xd6\xc0\xee\x92\xf9\xd0\xd0\x85\x78\xdb\x2c\x4f\x18\xd6\x60\x71\x4c\xa8\x7c\x5d\x89\x96\x35\xa9\xb5\x27\x80\x2d\xd8\x97\x46\xa4\xf3\x8c\x24\x0f\xcb\x64\x0e\x5c\x5c\x1e\x34\x53\x5a\x32\xba\x8b\x95\xa2\x7f\x9e\x3f\x48\xae\x59\xbc\xa8\xbf\xc7\xe3\xaa\x66\x25\x68\xe5\xa8\xa7\x1b\x9d\x68\x74\x50\x20\x0f\x5b\xbe\xda\x62\x96\x16\x05\xc2\x38\x43\x90\x36\x07\xcc\xa0\xef\x72\xaa\x32\xac\x60\x8b\x35\x11\x98\x06\x75\x99\x58\x4b\x7e\x96\x6f\x72\x73\x82\x7a\xa0\xa6\xd7\xb5\xbd\xdf\xdc\xb8\xb9\x0a\xbe\xbd\x28\x8a\x17\xf9\x35\x8a\xab\xf9\x76\x93\x5f\x83\xb4\x7b\x5d\x8b\x18\xd9\x67\xb4\xa3\xf5\xd1\xf5\xa7\x10\x6b\xcc\x07\x61\x02\xc6\x1f\x83\x16\xf2\xab\x6f\xfd\x51\xa9\x6b\x14\xd5\x5c\x41\xa4\xf0\xff\x9c\x22\xf6\x55\xe8\x28\x5c\x18\x37\x75\x07\x54\x3e\xda\xda\x8d\x46\x3e\x23\xf2\x4c\x7b\x36\x36\x20\x35\x44\x0e\xec\x27\x87\xe3\xee\x87\x7c\x50\x94\x9d\x8f\xad\x64\xf0\xfa\xb6\xf5\x7d\xd9\xd8\x3f\x11\x28\x5e\x58\x4c\xb5\x82\x72\x58\xd5\x7c\x75\x5c\x46\xec\x80\x58\x52\xa8\xf2\xf4\xd3\x5d\x5b\x0f\x0e\x99\xc1\x0d\xb8\xcb\x6f\xbe\x05\xc1\x22\x4c\xb2\x77\xb2\x76\xb4\x6d\x31\x04\xb3\xb7\xab\x9c\x00\x72\x20\x18\xf8\x11\xcd\xf7\x6c\xea\x4e\x56\x8c\x5e\x7c\xc9\x6c\x62\x47\x11\x9f\xad\x61\x11\x83\x13\xea\xbf\xee\x18\x88\xfc\x34\x48\xfa\xce\xf7\x2f\x16\x23\x9f\xa5\xde\xd1\x36\x1d\x5d\x1d\xca\xe2\xc3\x0c\xc8\xcd\x94\x90\x21\x52\x4e\xb8\xd2\x16\xab\x84\x19\x69\xe7\xb1\x90\xe1\x28\xd7\x9e\xcf\x74\x89\xc5\x6a\xe9\x50\x35\xdf\x48\xe1\xef\xe8\xe6\x2f\xe5\xa6\x03\x35\xfe\xc9\x8c\xab\x98\x5a\x49\x8e\xf1\x6a\xd1\x47\x6a\xc9\xab\xe1\x7d\x8e\xb5\xbd\x44\xe6\x9a\x08\xda\x83\xde\x8a\xc6\x5b\x20\x13\x2b\x58\xd4\xcc\xa2\x39\xad\xaa\x92\xda\xd5\xd2\xe4\x6a\x0f\xa1\xdf\xd5\x95\x3b\x43\x66\x84\xe2\x35\x8f\x22\xf1\x4f\x8e\x66\xc3\xec\x88\x22\x09\xc3\xca\xa0\x0d\xfe\xe4\xba\x6b\xb3\xae\xa9\xe7\x04\xcb\x2a\x2d\x24\x2b\xb5\xec\xd8\x89\x95\xb7\xac\x6e\x8b\xe4\x27\x74\xea\xa6\xc8\xa4\x04\xde\xa8\x13\x04\xdd\x01\xb1\xd7\x7a\xc1\xa3\xa8\xe4\x14\x1a\xd4\xa0\x41\xeb\xfa\xf3\x70\xf0\x15\x42\x8c\x5d\xa9\x26\x5a\xb4\xa6\xfc\x77\x72\xd9\xda\x2c\x5b\x8b\x15\xad\xaf\x20\x42\x39\xb1\xd0\x14\x56\x66\xf5\x71\x0d\x3c\xfc\x24\xd8\xa5\x42\x28\xd9\x73\xa9\x3b\x5a\xb3\x66\x6f\x8a\x90\x5b\xaa\xc8\xa6\x16\x4b\x5a\x13\xba\x82\xc3\xd3\xe8\xa8\x34\x00\x54\x89\xbe\xae\x67\x26\xd6\x87\x50\xe8\x22\x69\x9b\xfa\x9c\x94\x84\xce\xd0\xa2\x53\x4c\x5a\x52\x7c\x06\x23\xce\x90\xe2\x0d\x04\x81\x88\xbf\x45\xdb\x5e\x82\xb0\x27\x65\x58\x9b\x28\x10\x9c\x8a\xcb\xcf\xdc\xc6\x83\xd9\x06\xf8\xf8\xe1\x0e\xf0\x3e\x88\xd2\x45\x62\x5e\x1e\x85\xdf\x50\xb9\x51\x45\xf2\x5d\x12\xdc\xdd\x81\x83\xa2\x64\x4a\x9d\x98\xb6\xda\x0a\xbe\x62\xaa\x48\x13\x05\x11\xd4\x0a\xd6\x0b\xe6\x91\x64\x4d\x79\x9d\x9c\x48\x6b\x5f\x42\xc1\xdf\xa8\x6c\xd0\x73\x88\x46\x4b\x51\xe7\xe4\x99\x5b\xe1\x99\x09\x6e\x90\xba\x0f\x66\xd4\x59\x89\x82\x83\xbb\x35\x95\xd7\xc3\xf2\xb1\xde\xb2\x03\x86\x60\xd6\x5f\xe5\xe7\x80\x3d\xb3\xdb\xb6\x88\x78\x51\xf5\xc8\x50\xa3\xa2\x10\x5e\x9d\x01\x45\xeb\x07\x7a\x50\x1e\xb5\x9c\x3c\x03\xda\x4d\xee\xf0\x02\x68\x16\xcc\x0d\xf8\x1e\xda\x0c\x36\xf5\x16\xef\x12\x9a\x1b\x51\xfc\x2c\xc1\x7a\x6a\xe7\x9f\x29\x9f\xd2\xc8\xa7\x6d\x7b\xfd\x54\x25\x73\xf2\x88\x6d\x19\x97\x99\x9f\x5f\xb6\xe2\x41\x4d\x5d\x99\x74\x5a\x67\x51\x21\x6b\xaa\xb6\x10\x0e\x9d\x21\x81\xbb\x62\x39\xec\x90\xf0\xd7\x24\xdd\x4d\x27\x18\x74\x06\x96\xff\x59\x02\xdf\x34\x01\x86\xb1\x61\x0c\xef\x5f\xee\x48\xd7\x54\x0c\x93\x49\xbb\xcf\x25\x77\x6b\xc8\xed\x6d\xe3\x97\x69\xdf\xf7\x62\xb7\xa3\x44\x31\x08\x11\x61\x9f\x0e\xff\x89\xde\x0f\x94\xd2\x73\xa4\x0c\xd3\x79\x39\xf9\x11\x00\xe9\x8c\xa0\x1f\xf4\xc5\x1e\xbe\x69\x84\x64\xd5\xe7\x12\xe0\xbd\x21\xc0\x7b\x25\x46\xf6\xd0\x39\xb1\x58\xa8\xbe\x88\x40\xaf\x4c\xb3\xf7\x48\xd4\xa8\x32\xc9\x63\xd3\xe0\x82\x5b\x3b\x70\x56\x9f\x55\xdd\x67\x92\x3e\x3c\x0b\xfb\x51\xc8\x8e\x1e\x80\x2a\x9d\x32\xa7\x6d\xf6\xd1\xfe\x0c\x09\xfe\x48\xcb\x59\x1b\x85\xcb\x9b\x04\x24\xb6\x2a\x4b\x5e\x81\xe4\xd6\xb5\x3d\x33\x9a\x57\xea\x73\x89\x8d\x64\xbe\x82\x1d\xff\xf9\xb4\xb6\x87\x68\xec\xc7\x47\xa5\x3c\x43\x1a\xd4\x69\xba\x63\xe4\x81\x1e\xdc\xcf\x22\xb4\x35\x05\x19\x65\x1f\xb5\xe3\x01\xe6\x16\x0c\xe9\xed\x56\xfe\x28\x9a\x93\xd4\x66\x6d\x0c\x15\xe7\x9f\xcd\x03\x4c\x9f\x5c\x19\x7c\xbd\xcb\x8f\x92\x63\x5f\x68\x05\x4c\xc0\x4b\xc7\xd5\x89\x63\x09\xff\x33\x04\xb2\x39\x67\xf2\xd2\x65\xd4\xd1\x94\x02\x81\x5c\xd6\xb8\x4f\x53\x2f\x0f\xe7\xa0\xb9\xfc\x53\x86\x19\x8f\x05\xa9\x84\xce\xc8\xfb\x96\x6d\x32\xd2\x56\xeb\x8c\xb4\xcd\x26\x23\x6a\xbf\x39\x4b\x5e\xeb\x36\x0d\x95\xed\x19\x6a\x47\x79\x33\x3a\x3d\x85\x07\x2b\x7c\x03\xf1\x13\xa6\x86\x91\x3f\xf8\x1f\x70\x48\xb9\x3b\x86\x78\x80\x2e\x4e\xfd\x58\x4f\x8a\x21\xb9\xe9\x47\x84\x99\x79\xff\xfd\xdc\xef\x43\x44\x1f\x88\x28\x03\x28\xfe\xab\x2b\x70\xfa\x9f\xdc\x21\xc5\x91\x5f\xe1\xb1\x49\x48\x3c\x6c\x0f\x7e\x92\xa6\x1f\xda\x97\x24\x70\x19\x90\xdf\x61\x3a\xee\x58\xd7\x49\xf1\xff\xc7\x8d\x2f\xd7\x36\x5f\x17\xc2\xc3\x25\x4f\x00\x1d\xf7\x1f\x5c\x02\x39\x2e\x72\xf9\xd1\x56\x6e\x8b\x0b\x0a\x77\x63\x30\x83\x2a\x48\x3f\x3d\x4c\x46\x46\x38\x05\xe2\x56\xae\xcc\xef\x00\x5c\xbb\xd4\xf4\x20\xda\x5b\x8a\xce\x67\x88\x0f\x51\xd6\xcb\x54\x2d\xdc\x4f\x5f\x50\x7b\x79\xfd\xeb\x60\xae\x64\x84\xe1\xf5\xf6\x2a\xe2\x16\xa6\x03\xbf\x2a\x88\x0f\xdb\x7b\x3a\xc4\x59\xb5\xe9\x4c\xa6\x1f\xcc\xa3\x5c\xe7\xb0\x91\x0c\xeb\x41\x36\x7a\xff\xea\xab\xaf\xc8\x4f\x47\xf6\xd0\xc7\x53\xb8\x91\xc5\xa4\xad\x3a\x5a\x02\x22\x41\x7e\x86\x7d\xc0\x83\x55\x00\x7f\x9c\xa8\xf1\xf8\xdd\xc3\x06\xdb\x61\x8f\xb4\xe9\xfe\x1d\x8f\x36\xbb\xf9\x86\x3c\x06\x05\x9d\xfb\xcd\xa9\xda\x54\x88\x9c\xcb\x0c\xb1\x0f\x6a\x8c\x09\xb1\x3d\xea\xa6\x6b\x4e\xb2\x0f\xe3\xdb\xd2\x47\xb0\x0a\x30\x23\x57\x21\x6a\x16\xdc\x39\xf4\xec\xdc\xab\xe4\x9b\x7f\x7b\x31\x31\xb4\x17\x0a\x97\x0f\x1b\x25\x2c\x23\x51\xc0\x97\xa7\xa5\x80\x7c\x3f\x91\xd1\x74\x6c\x3f\xcf\x64\x9a\x91\x65\x66\xea\xff\x53\xab\xc5\x9c\xba\x7e\x22\xc5\xb7\xe4\xf1\xc6\xfc\xf7\xa2\xa7\x0e\x3d\xd3\xa3\x7e\xc1\x67\xf9\xe5\x20\xe2\xf6\x84\xa3\x87\xe4\x3f\x84\x81\x91\xe2\x17\x85\x3d\x98\xe3\x49\x32\x0d\xb5\x51\x48\x4b\xd7\xf9\xd4\xa5\x02\x67\xa9\x6e\x0c\xe0\xbe\xe3\xd6\xb4\x05\xe1\x0a\xbe\x04\xa1\xda\x9a\xeb\x34\xc9\x92\xf9\xdc\x2f\xef\x4f\x0b\x83\x6b\x5a\xd6\xd9\x0c\x3b\xf6\x5d\x6d\x12\xe7\xba\x33\x9b\xfd\x05\x9d\xe1\x2f\xe6\x1c\x65\x81\x6f\xfa\x47\x28\xb4\xae\xc3\xc6\x7f\xff\xd7\xd1\xe9\xf6\x52\x00\x4e\x36\xb9\x3d\x1b\x89\x21\xe5\x7b\x1d\xf0\x0d\x9c\x9e\x50\xb3\xd9\x8c\xaf\x89\xbb\xcc\x8f\x54\x2f\x4b\x88\x26\xca\xd2\xda\x5b\x5f\xbd\x36\x31\xc6\xfc\x7f\x02\x00\x00\xff\xff\x4a\x10\x27\xf4\xba\x51\x00\x00")

func buildersBindataPipdeptreePyBytes() ([]byte, error) {
	return bindataRead(
		_buildersBindataPipdeptreePy,
		"builders/bindata/pipdeptree.py",
	)
}

func buildersBindataPipdeptreePy() (*asset, error) {
	bytes, err := buildersBindataPipdeptreePyBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "builders/bindata/pipdeptree.py", size: 20922, mode: os.FileMode(420), modTime: time.Unix(1525410168, 0)}
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
	"builders/bindata/pipdeptree.py": buildersBindataPipdeptreePy,
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
	"builders": &bintree{nil, map[string]*bintree{
		"bindata": &bintree{nil, map[string]*bintree{
			"pipdeptree.py": &bintree{buildersBindataPipdeptreePy, map[string]*bintree{}},
		}},
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

