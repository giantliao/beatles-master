// Code generated by go-bindata. DO NOT EDIT.
// sources:
// contract/ERC20.sol (6.827kB)

package abires

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
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
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

var _contractErc20Sol = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x59\x5b\x6f\xdb\xca\x11\x7e\xd7\xaf\x98\xba\x0f\x95\x02\x1f\xc9\x55\x73\x4e\xd2\x38\x0a\xea\xdc\x0a\x03\x0d\x5c\x34\x6e\x1f\x1a\x04\xc6\x92\x3b\x14\x17\x21\x77\x99\xdd\xa1\x54\xb5\xc8\x7f\x2f\xf6\xc6\x9b\x28\xf9\xd6\xa6\x01\xfa\x26\x93\x33\xb3\xb3\xf3\x7d\xf3\xcd\x2e\x5d\x69\xb6\x2e\x19\x18\x55\x08\x2e\x68\x07\xaf\x56\x67\xf3\xa7\xf3\xe5\xd3\xf3\xc9\x44\x94\x95\xd2\x04\x27\xf3\xc5\xe5\xbb\xbf\xbc\x59\x9e\xcd\x8d\x2a\x4e\xce\x3b\x8f\x0d\xcb\xb0\x64\x94\x87\x17\x93\x54\x49\xd2\x2c\x25\x70\xe6\x20\x0c\x78\x47\xf8\xd7\x04\xa0\x36\x42\xae\xe1\x23\xcb\xf0\x03\xa3\x1c\x32\xa5\xa1\x16\x92\x96\x3f\xff\x72\x3e\x99\x00\x94\xac\xaa\xac\xc1\x94\x71\xae\xd1\x18\x58\xbd\x8a\xef\x67\x50\x69\xb1\x61\x84\x70\x93\xb0\x82\xc9\x14\xcd\x41\x97\x63\x61\x3a\x71\x58\x51\xa8\x2d\x72\x17\x26\xbc\x6e\x5f\x92\x22\x56\x7c\xac\xab\xaa\xd8\x9d\x4f\x00\x0c\x69\x1b\xb1\xaa\x93\x42\xa4\x90\x2a\x69\x88\x49\x02\xc9\x4a\x84\x15\x9c\xbc\x46\xa4\x02\xe1\x8d\x12\xf2\xe4\x88\xb9\xd9\x95\x89\x2a\x9c\xc3\xf5\x9f\xde\x38\x4b\xbb\xf0\xf3\x3d\x43\x8e\xa9\x28\x59\x61\x60\x05\xbf\x7d\x7e\x3e\xb1\x19\xba\x77\xba\x4e\x49\xe9\x69\xcc\xd6\x25\x39\x8b\xde\xb6\xc0\xd0\xcb\x1c\x56\xde\xe4\xdc\xbf\x89\x85\xfb\x54\x9a\xf5\xdc\xa0\xe4\xa8\x3f\xc3\x6a\x6f\xaf\xdf\xec\x72\x8b\x27\x4f\x26\x00\x4f\xe0\x0f\x1c\x37\x70\x6d\x0d\x40\xd6\x65\x82\x1a\x54\x06\xa4\xbe\xa0\x34\x20\x24\xe0\x3f\x84\x21\x94\x29\x5a\xe3\xc5\x04\x20\xab\x65\x4a\x42\x49\xe8\x04\x9d\x36\x29\x6e\x04\x6e\x41\x23\xd5\x5a\x1a\x98\x36\xd0\xfa\xcc\xfd\xf3\x3b\xa4\xf3\x47\x24\x03\x94\x23\x84\x0d\xb9\x94\x72\x04\x53\x61\x2a\x32\x81\x1c\x02\xf0\x73\xef\x53\x31\xcd\x4a\x50\x5b\x89\x1a\xae\x73\x8c\x6f\x81\x14\x7c\xad\x51\xef\x06\xb1\x82\x57\x48\xe7\x42\x36\xe4\xd0\x58\x69\x34\x28\xc9\x62\x6b\x7d\x58\xa9\x6a\x49\x2e\x32\x87\xc4\xc7\xa9\x98\x31\x83\x0c\x7a\x75\x09\xeb\x5c\x65\x0d\x3b\x5d\x62\xf7\x2b\x51\x83\xa4\xf3\xfd\x3c\x28\x53\xac\xd3\xfb\x16\x0b\x48\x73\x4c\xbf\xf4\x92\x6e\x60\xa4\x9c\x11\x30\x19\x0a\x14\xba\xc2\xfa\x30\x5b\x51\xcb\x92\x79\x08\xda\x2d\x64\xcc\xbe\x5b\xd0\x6d\x2e\xd2\xdc\xbe\xf7\xf0\x64\xb5\xe4\xa6\xef\x1b\x02\x1e\xf1\xde\x8a\xa2\xf0\x66\xfb\x31\x22\x24\x0d\x22\x1e\xf1\xdd\x10\x8f\x66\x6b\x86\x6c\x34\xb6\x61\xa2\x60\x49\x81\x4e\x70\x02\x53\xda\x7d\xf5\xe0\x71\xdb\xb7\xb5\x9d\xba\x82\xf7\x30\x3a\xed\x3d\x0a\x31\xec\xb3\x99\x7b\xe1\x01\x74\x3f\x2d\x88\x1d\xc0\x3a\x40\x4e\x60\x00\x65\xa8\x77\x40\xf2\x53\x88\x3a\x84\x34\x36\xa2\x66\xd2\x64\xa8\xfd\xfe\xdc\x76\xd8\x3e\xed\xbb\xac\x27\x35\xa4\x3c\xb5\x31\x7a\xfd\xb1\x61\x45\x8d\xde\xd8\x97\x91\x14\x24\xd8\x98\x6b\xe4\xfb\x64\x8e\x2f\xa7\x6d\xfc\xd3\x06\x1c\x17\xb0\xe1\x75\x53\x89\x44\xa9\xa2\xe5\xf3\xd7\x5a\x68\x9c\xfa\xb5\x5f\xae\xc6\x35\x6a\x76\xde\x33\x26\x05\xbf\x5a\xc5\x1d\x4d\xcf\x66\x33\xa7\xe0\x47\xf4\x6d\xec\xf9\xdc\xd4\x89\x5f\x76\x36\x54\x47\x52\x7d\x2f\x52\x9f\xe7\x8c\xf3\x9e\x35\x96\x82\x1a\x34\xa6\x6d\xdc\x53\x57\x81\xae\x65\xc0\x99\x74\x8d\x07\xda\xf4\xa2\xaa\xb4\xda\xe0\x88\x7a\x58\x08\xda\x56\xe8\xe0\x3c\xe4\xb9\xd5\x15\xcc\x59\x91\xd9\x67\x6d\x36\xa1\x6f\x5e\xe3\x96\x69\xf4\x7d\x9e\xe6\x4c\xae\x6d\xbf\xb0\x0e\xd7\x61\x2b\x28\x07\xca\x85\x81\x12\x29\x57\x1c\x12\x3b\xbf\x7c\x17\x6b\x61\xbe\x78\x5f\xa3\x4a\x54\x12\xa1\x64\x3b\xa8\x0d\x42\xa2\x9c\x17\x82\x2a\xb8\x5f\x89\x85\x54\x25\x6e\x3b\xd1\x93\x1d\xd4\x32\x53\x9a\x6a\x69\x67\xab\x23\x0d\xf3\x04\x52\x9a\xa3\x5d\x6a\x0e\x57\x12\xa1\x52\xc6\x08\xdb\xa7\x46\x15\x75\x54\xae\x52\x90\x58\x3b\xbf\x5c\x18\xbf\x8c\x66\x29\xda\x81\xc8\x85\x33\x12\xae\x50\x99\xd0\x86\x40\x23\xaf\x53\xec\x36\xf9\x6f\x4c\x27\x15\x52\x70\xe6\xb2\x34\x48\xce\x88\xa3\x11\x1a\x79\x60\x3f\xcb\x08\xf5\x96\x69\x6e\x5e\xf8\x95\x72\xa2\xca\xbc\x58\x2c\xd6\x82\xf2\x3a\x99\xa7\xaa\x5c\x20\xe5\xa8\xb1\x2e\x17\xef\x2e\xff\x6c\x16\xc2\x98\x1a\xcd\x62\x79\xf6\x6b\xf7\x2b\x55\x65\x89\x92\x7e\x5a\xfe\xf2\xbb\x9f\x97\x4f\x9f\x2d\x7f\x3f\xaa\x7e\xf7\x53\xbd\x03\xed\xd9\x11\x70\xd7\xa8\xd6\x9b\xc6\x14\xcd\xd3\x6b\x3a\xd0\xad\x07\xf5\x69\xdc\xc0\x78\xff\x45\x29\xeb\xb4\x59\xa3\x67\xb0\xf2\xeb\x74\xda\xc7\xf3\x9e\x15\xbd\xf6\x69\xb2\xbb\x5f\x0f\xf5\x85\xd1\x40\xa6\x55\x09\x96\xab\x9d\x4e\x62\x52\x59\xe8\x7a\x55\x75\x76\x87\x87\xd1\x4e\xd5\xb0\x65\x5e\x0b\x8d\x03\xa7\x0d\xdf\x8b\x63\xc3\xdf\x29\x4a\x47\x7d\x47\xe0\x6d\x0e\x76\x87\x61\xee\xe8\xf1\x3e\xd6\xf1\xe5\x7b\xad\xca\xfe\x00\xb3\x19\xf7\xe7\x17\x29\xff\x77\x8f\x06\x93\x91\x61\xd6\x67\x44\x67\x80\x1d\xd4\x6e\xbb\xd8\x50\xb5\x5b\xb3\x48\x13\x67\xf5\x28\x9d\x77\x11\x60\x6f\xe5\x47\x6a\xfb\xe1\x04\xe1\x58\xf6\x7b\xab\xf6\x67\x84\x03\xe0\x01\xd3\xe1\x52\xa6\x1a\x99\xc1\x47\x9e\xdd\x82\x0a\x80\xc9\x55\x5d\x70\x4b\xa4\x94\x15\x05\x72\xd8\xe6\x28\xa3\xe3\xcd\xa7\x9b\xb6\x5f\x57\x70\x36\x87\x6b\x05\xc2\x26\x60\x45\x2d\x04\x0a\x6b\x78\x3c\x85\x81\x04\x89\x1c\x9f\xdd\x4c\x70\x43\xa4\x73\x09\x00\xb6\x51\x82\xc3\xd2\x2d\x67\x60\x6a\xb5\x77\xcb\x04\x41\x2d\x49\x14\x3e\xa4\x53\x3c\xa7\xde\xdd\xd1\x60\x87\x91\x90\xc8\x67\xde\xc8\x52\x1a\x3e\x28\xa9\x0a\x41\xf9\xdb\x8b\x2b\xb8\xb6\x15\xb0\xb7\xce\xff\x98\xc4\x32\xce\x91\xff\xed\xa8\xce\x8a\x1e\x1a\x9d\x09\x37\x22\xbc\xd1\xf6\x62\xfc\x48\x19\x95\xae\xd7\x84\x6d\x0a\x0f\xe8\xc4\x47\xa9\xb3\x4f\xee\x16\x3b\xd7\x2b\x6d\x8e\xb3\xd9\x9d\xe5\xfc\x68\xd8\x3b\x37\xc3\x5b\xfc\x1f\x37\x03\xc7\xff\x97\x66\x30\x75\xe2\xbe\xe3\xdc\xd6\x11\x1c\xef\xde\x11\xd1\xf6\x3e\x1d\x31\xc8\xe3\xc7\x6c\x0b\x2b\xfd\x83\x44\xbf\x7b\x6f\x5c\x4a\x42\x2d\x59\xd1\x21\x9c\xed\x87\x52\x48\x32\xee\xb4\xdf\x82\x97\x63\xb8\x40\x5a\x02\x32\x63\xc4\x5a\x1a\x10\xd4\x9c\x49\xac\x75\x9a\x5a\xf3\x39\x5c\x5b\x0e\xa3\x4c\x59\x65\xea\x82\x11\xfa\x4b\x41\xa9\xb8\xc8\x44\xca\xfc\x21\x3e\x8b\x5f\x35\x0c\x98\x3a\xcd\xfd\xc2\x94\xa3\x8f\x56\x69\x55\xa1\x06\xdc\xa0\xcb\x44\xa3\xab\x09\xf9\xfb\x64\x47\x7d\xfd\x8a\x9e\x68\xe1\xb7\x0b\xe4\xb8\xaa\x31\x45\x11\x2e\x49\x96\x44\x84\xf1\x30\x36\x88\x52\xb6\x41\xca\x41\x8c\xa4\x71\x1d\xe1\xe6\x8d\x2d\x54\x73\x48\x0e\x09\xb4\x87\x64\x1f\x6c\x06\x22\x56\xb9\xcf\xb1\x98\xf0\x90\x63\x23\x5f\xe4\xba\x7f\x7a\x3d\xf5\xa1\x87\x47\x95\x10\xb2\x7f\x5e\x89\x0f\xf7\x1d\xfb\xa7\x8d\x36\x89\xd3\x76\x2f\x1d\xfb\xbb\xf3\x27\x71\xad\x75\x80\x3f\x2a\x03\x06\x6b\xb1\x41\x19\x88\x13\x58\x73\x2b\xb0\xdb\x5c\x19\x8c\x2a\x12\xc1\xb1\x4b\x0d\x7d\x8f\xc3\xd9\xf1\xe8\x83\x69\x5f\xfc\x57\xc0\x6c\x6c\x7c\x36\x2f\xc7\xc0\x69\xb4\xe5\x08\xf0\x56\x31\x1e\x04\xfc\x9e\xe3\x00\xf8\x06\xec\x2e\x03\xbe\x17\xf0\xa7\xc0\xed\xfd\xdb\x7d\x20\x75\x57\x2b\x77\x0f\xdf\xbf\x86\x67\x4a\x83\x61\x82\xb7\x32\xf3\x57\xe3\x95\xc5\x87\x6b\x80\xb1\x49\x34\x79\xfd\x18\xb4\x72\x77\xab\x07\x53\xab\xa5\x4d\x14\xfd\x88\xec\xe0\x26\xe4\xdc\x16\x0b\xf8\xe8\xcf\x28\x23\x1f\x22\xae\x2a\x94\x7f\xc7\xaa\xc2\x42\xc8\xc5\x3f\xc3\x8f\x9f\xe2\x3f\x71\xe2\x97\x89\x67\x67\xcf\xec\x96\x58\x9a\x62\x45\xc8\x4f\x63\xdc\xfe\xd9\x44\x22\x72\x37\xce\x1d\x9b\x98\xf4\x6a\x1d\xbf\x0b\x21\xd4\x15\x77\x92\xcb\xc2\x20\x9b\xf7\x87\xe7\xd8\x1e\xe0\x96\x3d\x3a\x26\x87\xe9\x3a\xe8\x04\xd7\xbc\xa3\xaa\xf5\xed\xdf\x01\x00\x00\xff\xff\xb8\xc6\xd0\xc2\xab\x1a\x00\x00")

func contractErc20SolBytes() ([]byte, error) {
	return bindataRead(
		_contractErc20Sol,
		"contract/ERC20.sol",
	)
}

func contractErc20Sol() (*asset, error) {
	bytes, err := contractErc20SolBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "contract/ERC20.sol", size: 6827, mode: os.FileMode(0644), modTime: time.Unix(1610105769, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xc9, 0x53, 0xfc, 0x31, 0xb4, 0x69, 0xcd, 0xc, 0xf2, 0x8d, 0x8b, 0x71, 0xfe, 0xab, 0xeb, 0x31, 0xb9, 0xb5, 0x2a, 0xd4, 0xb0, 0x53, 0x8f, 0xe0, 0x5f, 0x8a, 0x88, 0xd6, 0xf2, 0x10, 0xa8, 0x3b}}
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
	"contract/ERC20.sol": contractErc20Sol,
}

// AssetDebug is true if the assets were built with the debug flag enabled.
const AssetDebug = false

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
	"contract": &bintree{nil, map[string]*bintree{
		"ERC20.sol": &bintree{contractErc20Sol, map[string]*bintree{}},
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
