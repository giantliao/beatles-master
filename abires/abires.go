// Code generated by go-bindata. DO NOT EDIT.
// sources:
// contract/ERC20.abi (3.468kB)

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

var _contractErc20Abi = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x96\x31\x6f\xdb\x30\x10\x85\xff\xcb\xcd\x9a\x0a\xb4\x08\xb4\x65\xe9\x56\x74\x68\xd0\xa5\xc8\x70\x12\x4f\x05\x01\xea\x48\x90\x47\xb9\x42\x90\xff\x5e\x88\x4d\xcc\xd8\x90\x2c\xcb\x36\x64\x75\xb2\xa1\x93\x8f\x8f\x1f\xdf\xa3\xef\xd7\x0b\x68\x76\x51\x02\x94\xe9\xab\x90\x67\x34\x4f\xbd\x23\x28\x21\x6a\x96\x4f\x9f\xbf\x40\x01\x8c\xed\xf0\x40\xac\xa0\x81\x02\xe4\xb0\xfe\xfa\x5c\x80\xc3\x1e\x2b\x43\x50\x36\x68\x02\x15\x10\x04\x85\xbe\x45\xc1\x4a\x1b\x2d\x3d\x94\xc0\x96\xdf\x5f\xda\x77\xa8\x2d\x07\xf1\xb1\x16\xeb\xe1\xb5\x78\x01\x64\xcb\x7d\x6b\x63\xd8\xf7\x39\x50\xa7\xe8\x0f\x29\x28\xc5\xc7\x54\x39\x10\x8b\x4a\x79\x0a\x21\x8b\xb5\x3b\x26\x9f\x97\x7a\xaf\x0f\xcb\x2c\xec\x14\x1c\xb1\x9a\xed\xb5\x17\x7c\x9a\x61\x87\x26\xd2\x38\xc3\xb7\x37\x1e\x9d\xf3\xb6\xfb\x08\x9a\x3a\x62\xb9\x39\xa0\xc6\xdb\xf6\x26\x7c\xc4\xae\x85\xe6\xc9\x23\x87\xe6\xe3\x51\x64\x34\xc9\x4b\xc8\x92\x55\x4f\xf9\x7a\x99\x55\x2e\xf5\x46\x56\x8d\xc6\xd8\x1d\x72\x3d\xec\xcd\x46\x39\x3b\x6e\x97\x26\xad\xd3\xb4\xcb\xbf\x6d\x22\xd7\xa2\x2d\x1f\x41\x1a\xb3\xcf\x55\x31\xb8\xf6\x6c\x31\xd9\x7e\x86\x51\x65\xad\x19\x03\x94\x9e\x5f\x7e\x0f\x4d\x30\xba\x99\x91\xf2\x2e\x2b\x34\x83\x13\xbe\x37\x5b\xf2\xc2\xd1\x3e\xb3\x5a\x45\xb5\x6e\xd1\x84\x79\xb1\x0f\x53\x52\x1f\xfe\x63\xd3\x86\x58\x89\xc7\x5a\x48\xfd\x9c\xb5\xaf\xa2\xda\x13\x06\x7a\x3c\x2f\xec\xeb\x1a\x79\x65\x6e\xa8\xd4\x39\xc8\x34\x6f\x18\xd9\x64\x26\xd2\xc7\x49\xa1\x41\xbc\xe6\xdf\x63\x52\xdf\x2a\x6b\x44\x37\xf4\x6d\x95\x78\x6d\x5d\x68\x1a\x2a\x7f\x44\xe7\x4c\xbf\xa5\x3b\x71\x79\x64\xa6\xc7\xa0\x6b\xff\x1a\x25\x8f\x3d\x9b\xc9\xc7\x72\x3e\xa7\xe6\xcd\x7b\x93\xfd\xfa\x4f\xdb\x3d\xe9\x3e\xff\x0d\x00\x00\xff\xff\x81\xbb\xb5\x41\x8c\x0d\x00\x00")

func contractErc20AbiBytes() ([]byte, error) {
	return bindataRead(
		_contractErc20Abi,
		"contract/ERC20.abi",
	)
}

func contractErc20Abi() (*asset, error) {
	bytes, err := contractErc20AbiBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "contract/ERC20.abi", size: 3468, mode: os.FileMode(0644), modTime: time.Unix(1610105940, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x23, 0x35, 0x29, 0x7e, 0xea, 0xd9, 0xbd, 0x8c, 0xc3, 0xb2, 0xeb, 0xe9, 0xd9, 0xf3, 0xb, 0x8a, 0xdf, 0x36, 0xbd, 0x51, 0x54, 0x58, 0x5d, 0xce, 0x9c, 0xdf, 0x1d, 0xe7, 0x83, 0x94, 0xb5, 0xee}}
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
	"contract/ERC20.abi": contractErc20Abi,
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
		"ERC20.abi": &bintree{contractErc20Abi, map[string]*bintree{}},
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
