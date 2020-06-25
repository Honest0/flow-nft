// Code generated by go-bindata. DO NOT EDIT.
// sources:
// ../../../transactions/mint_nft.cdc (636B)
// ../../../transactions/read_nft_data.cdc (563B)
// ../../../transactions/setup_account.cdc (610B)
// ../../../transactions/transfer_nft.cdc (553B)

package assets

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)

	if err != nil {
		return nil, fmt.Errorf("read %q: %w", name, err)
	}

	clErr := gz.Close()
	if clErr != nil {
		return nil, clErr
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

var _mint_nftCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x91\xc1\x8e\xa2\x40\x10\x86\xef\x3c\x45\xad\x87\x0d\x5c\x70\xcf\x66\xd5\x10\x94\xdb\xb2\x13\xe5\x05\x9a\xb6\xc4\xce\x34\xdd\x9d\xa2\x98\x71\x62\x7c\xf7\x49\xb7\x08\x8e\xb1\x0f\x90\xc0\xd7\x5f\xfe\xbf\x4a\xb5\xce\x12\x43\x69\x4d\xd1\x9b\x46\xd5\x1a\x2b\xfb\x8e\x06\x8e\x64\x5b\xf8\x73\x2e\x8b\x2a\xdb\x6c\x76\xdb\xfd\x3e\x1a\xc8\xed\x59\xb4\x4e\x63\x59\x54\x0f\x4c\xfe\xbf\xac\x76\x59\x3e\xb2\x51\xc4\x24\x4c\x27\x24\x2b\x6b\x62\x42\xa9\x9c\x42\xc3\x0b\xc8\x0e\x07\xc2\xae\x4b\xe0\x12\x01\x00\x84\x87\x46\x86\x56\x19\x46\x5a\xc0\xef\x49\x9f\x96\x45\xf5\x2f\x7c\x8e\x02\xe6\x08\x9d\x20\x8c\x3b\xd5\x18\x8f\x66\x3d\x9f\x32\x29\x6d\x6f\xd8\xeb\x02\xe3\x4f\x87\xfa\x98\xde\x7c\xb0\x84\x1b\x9d\xd6\x96\xc8\x7e\xfe\x7d\xa9\x5f\xc5\xbe\xc8\x02\xe6\x1d\x5b\x12\x0d\xce\xc7\x3f\xc9\xaf\x60\xbd\xde\xe4\x78\x46\xd9\x33\x0e\xd1\xef\xc9\xc7\x72\xb0\x84\x06\x79\x48\x34\x75\x4e\xa2\x67\x1c\xd5\x47\x88\x36\x22\x23\xe0\x4f\xda\x20\xe7\xc2\x89\x5a\x69\xc5\x5f\xf1\xdc\xf5\xb5\x56\xd2\x67\xca\xad\xd6\x18\x26\x3a\xe4\x1a\xaf\xdc\xdb\x5d\x9e\xd7\x98\x4e\x77\xde\x82\xe7\xba\x8a\x93\x1f\x77\xd7\x6b\x70\xc2\x28\x19\xcf\x72\xdb\xeb\x03\x18\xcb\xbe\xc6\x94\x93\xf0\x88\x84\x46\x22\xb0\x05\x3e\x21\xf8\xcd\x4f\xda\x59\xf2\x72\xf2\xe1\x55\x16\xd5\xe3\xee\xef\xca\x64\x18\xea\xf5\x3b\x00\x00\xff\xff\x86\xbd\x57\x55\x7c\x02\x00\x00"

func mint_nftCdcBytes() ([]byte, error) {
	return bindataRead(
		_mint_nftCdc,
		"mint_nft.cdc",
	)
}

func mint_nftCdc() (*asset, error) {
	bytes, err := mint_nftCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mint_nft.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x65, 0xec, 0xdd, 0x1a, 0x3a, 0xef, 0x7, 0x4d, 0x38, 0x64, 0x96, 0x68, 0x68, 0xbc, 0xad, 0xa3, 0x9, 0xc5, 0x46, 0xa7, 0x58, 0xd2, 0x4c, 0xa5, 0x5a, 0x4e, 0x91, 0xe7, 0x8b, 0x7f, 0xf0, 0x54}}
	return a, nil
}

var _read_nft_dataCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x91\xc1\x6e\xa3\x30\x10\x86\xef\x3c\xc5\x9f\xcb\x2e\x5c\xc2\xae\xb4\xda\x43\xb4\x59\x89\x26\xa1\xea\x85\x56\x09\x7d\x00\x63\x86\xc4\x2d\xd8\xc8\x1e\x2b\xa9\xa2\xbc\x7b\x05\x01\x12\xb5\xdc\xd0\x7c\xff\xcc\x7c\x63\xd5\xb4\xc6\x32\x32\xa3\x53\xaf\xf7\xaa\xa8\x29\x37\xef\xa4\x51\x59\xd3\xe0\xd7\x29\x4b\xf3\x64\xbd\xde\x6e\x76\xbb\x60\x20\x37\x27\xd1\xb4\x35\x65\x69\x7e\xc7\xac\x9e\xb3\x7c\x9b\xac\x26\x36\x88\x63\xe4\x07\xe5\xe0\xa4\x55\x2d\xc3\x92\x28\x1d\x1a\x62\x51\x0a\x16\x10\x85\xf1\x0c\xa1\xd1\x75\x51\x1a\x02\xde\x91\xfd\xe9\x20\x4d\x5d\x93\x64\x65\x74\xd0\xfa\x02\x95\xd7\x68\x84\xd2\xa1\x90\xd2\x78\xcd\x0b\x24\x65\x69\xc9\xb9\x68\x81\xd7\x27\xcd\x7f\xff\xe0\x1c\x04\x00\x10\xc7\x78\x24\x06\x1f\x08\xad\x2f\x6a\x25\x31\x44\x60\x8a\x37\x92\x0c\x53\xf5\x45\x73\xd4\x64\xc7\x1f\xee\x4c\xfb\x78\x4d\x3c\x94\x96\xd8\x13\x27\xd7\xec\x38\x36\x0a\x26\xe8\xb6\xe0\x83\xb1\xd6\x1c\xb1\xbc\xe6\x7a\xa0\xfb\xe6\x7b\xe2\x95\x68\x45\xa1\x6a\xc5\x1f\x61\x7c\xdd\x26\xee\x6e\x34\x45\xa3\xd9\x0d\x2f\xfa\x36\xff\x7e\x9c\xbf\xbe\xc0\xfc\xc6\xbf\xf4\x3d\x2e\xff\xc3\x68\x36\xd9\x0e\xe3\x05\x2c\x55\x64\x49\xcb\x4e\x07\x02\xae\x25\xa9\x2a\x25\xc7\xd3\x76\x9e\x77\x57\x1d\x3d\x74\xc5\x58\x7e\xb3\x19\xb6\xc9\xd2\x3c\x54\xe5\x02\xbf\x07\x71\x4b\xec\xad\xee\x32\x73\x55\x06\x97\xcf\x00\x00\x00\xff\xff\x27\x05\x12\x73\x33\x02\x00\x00"

func read_nft_dataCdcBytes() ([]byte, error) {
	return bindataRead(
		_read_nft_dataCdc,
		"read_nft_data.cdc",
	)
}

func read_nft_dataCdc() (*asset, error) {
	bytes, err := read_nft_dataCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "read_nft_data.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x19, 0x91, 0x2e, 0xdb, 0x50, 0x62, 0xf7, 0xef, 0xd9, 0x4, 0x67, 0x3c, 0xa4, 0xf9, 0xda, 0x7c, 0x2f, 0xb, 0xf9, 0x8b, 0xae, 0xaa, 0xe6, 0x65, 0x47, 0x9, 0xf0, 0x1b, 0x7c, 0xc8, 0x22, 0x94}}
	return a, nil
}

var _setup_accountCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x91\xcf\x6e\xf2\x30\x10\xc4\xef\x7e\x8a\xf9\x2e\x28\x48\x40\xbe\x33\x0a\xa8\x11\x90\x63\x5a\x41\x5e\xc0\xb8\x0b\x58\x38\x76\x64\x6f\x80\x0a\xf1\xee\x55\x42\x0b\xa1\x2d\x7b\xb0\xfc\x67\xfc\x9b\x1d\x5b\x97\x95\xf3\x8c\xdc\xd9\xac\xb6\x5b\xbd\x36\x54\xb8\x3d\x59\x6c\xbc\x2b\xf1\xff\x94\x67\x45\x3a\x9f\x2f\x17\xab\x95\xf8\x52\x2e\x4e\xb2\xac\x0c\xe5\x59\xd1\xd1\xcc\x5e\xf3\x62\x99\xce\x6e\xda\x38\x46\xb1\xd3\x01\xec\xa5\x0d\x52\xb1\x76\x16\x3a\xe0\xb8\x93\x0c\x69\x21\x95\x72\xb5\x65\x1c\x5d\x6d\xde\xe1\x6b\xdb\x5c\x60\x87\x40\x0c\xcd\x81\xcc\x06\x75\xd5\x6c\x78\x52\xa4\x0f\x84\x3c\x2b\x82\x10\x5d\xda\x59\x00\x40\xe5\xa9\x92\x9e\x22\xa9\x14\x8f\x91\xd6\xbc\x4b\xaf\xe8\x3e\xce\xa2\x55\x34\xa5\x37\x8d\x23\x8f\xd6\xce\x7b\x77\x4c\x7a\xf7\x08\xa3\x99\x33\x86\x5a\xe2\x34\x6a\xe2\x8c\x11\x07\x76\x5e\x6e\x29\x6e\x62\xdd\x4e\xfb\x98\x4c\x60\xb5\xe9\x62\x9b\x32\xc4\x50\x37\x15\x92\x61\xe7\x7d\x46\xca\x93\x64\x5a\x94\x15\x7f\xdc\x49\x51\x1f\x32\xfc\xc3\xcb\x9f\x4d\x3c\xb0\x1f\x16\x6d\x80\x20\x0f\x14\x25\xc3\xbb\xe1\x00\xec\x9e\xb6\x2c\x7e\x03\x8c\xb6\xfb\xa4\x77\xfe\xf9\xdb\x9d\x0e\xde\xea\xb5\xd1\xea\x32\x8d\xe2\xaa\x9d\x3d\x32\x07\x60\xe9\xb7\xc4\xcf\x3d\xbf\xed\x2e\xe2\x3a\x5e\xc4\x67\x00\x00\x00\xff\xff\xf9\x8c\x1f\x1f\x62\x02\x00\x00"

func setup_accountCdcBytes() ([]byte, error) {
	return bindataRead(
		_setup_accountCdc,
		"setup_account.cdc",
	)
}

func setup_accountCdc() (*asset, error) {
	bytes, err := setup_accountCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "setup_account.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x63, 0xd4, 0x4, 0x5f, 0x75, 0x3c, 0xf0, 0x3d, 0x1b, 0xcb, 0x9a, 0x32, 0x26, 0xc7, 0xe8, 0xeb, 0x7e, 0xab, 0xa7, 0xa6, 0xeb, 0x34, 0x1e, 0x16, 0x42, 0xd6, 0x38, 0x10, 0x77, 0x35, 0xe4, 0x85}}
	return a, nil
}

var _transfer_nftCdc = "\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x91\x41\x6f\x82\x40\x10\x85\xef\xfc\x8a\xf1\xd2\x2c\x49\xc5\x1e\x9a\x1e\x88\x9a\x10\xd4\xc4\x0b\x6d\x94\xfe\x80\x65\x19\x70\xd3\x75\x77\xb3\x0c\xd1\xc6\xf8\xdf\x1b\xac\x2e\x20\xa7\x49\xf8\xe6\xcd\x7b\x6f\xe5\xd1\x1a\x47\x90\x19\xbd\x69\x75\x2d\x0b\x85\xb9\xf9\x41\x0d\x95\x33\x47\x78\x3b\x67\x9b\x3c\x59\xad\x76\xeb\xfd\x3e\xb8\x93\xeb\x33\x3f\x5a\x85\xd9\x26\x1f\x30\xe9\x67\x96\xef\x92\xd4\xb3\x01\x39\xae\x1b\x2e\x48\x1a\xcd\x1c\x0a\x69\x25\x6a\x8a\x21\x29\x4b\x87\x4d\xf3\x0a\x27\x49\x87\xd2\xf1\xd3\x76\x15\xc3\xf7\x56\xd3\xc7\x7b\x08\x97\x00\x00\xc0\x3a\xb4\xdc\x21\xe3\x42\x74\x0b\x2d\x1d\x12\x21\x4c\xab\xe9\x01\x74\x9f\x42\x02\x2f\x0b\x0b\xa8\x91\xee\x54\x7f\x2d\x0c\x46\xb8\x30\x4a\xe1\xcd\xd0\x0e\x2b\x58\x40\xa7\x1f\x15\xc6\x39\x73\x9a\xbf\xf4\xa1\xa2\xd4\x73\x4b\xd6\x05\x8c\x61\xd6\x90\x71\xbc\xc6\x59\x17\xd4\xff\x0d\x27\x23\xf9\x12\xad\x69\x24\xfd\x6b\x7b\x0f\x51\x8d\x94\x72\xcb\x0b\xa9\x24\xfd\xb2\x99\x6d\x0b\x25\xc5\xb3\x90\xb7\x71\x79\x7e\x86\x81\x9b\xaf\xdb\xea\x75\xc9\xc2\xc9\x38\x98\xae\x08\xe6\xd3\x71\xbe\xe8\xd1\x2f\x1b\x16\xdd\xcf\x83\x6e\x7a\xe3\xd1\x7d\x64\xd4\x9d\x8e\x61\x3e\xd5\x15\x85\x37\xf0\x1a\x5c\xff\x02\x00\x00\xff\xff\x93\x21\xc1\xd5\x29\x02\x00\x00"

func transfer_nftCdcBytes() ([]byte, error) {
	return bindataRead(
		_transfer_nftCdc,
		"transfer_nft.cdc",
	)
}

func transfer_nftCdc() (*asset, error) {
	bytes, err := transfer_nftCdcBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "transfer_nft.cdc", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xe2, 0x50, 0x26, 0x21, 0x3e, 0x36, 0x40, 0xfa, 0x93, 0x33, 0xd5, 0xc3, 0xc5, 0xab, 0x7d, 0x66, 0x74, 0x89, 0x15, 0xa7, 0xf3, 0xbd, 0xbb, 0x82, 0xb2, 0xc5, 0xad, 0xe8, 0xb7, 0xdb, 0x8c, 0x52}}
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
	"mint_nft.cdc":      mint_nftCdc,
	"read_nft_data.cdc": read_nft_dataCdc,
	"setup_account.cdc": setup_accountCdc,
	"transfer_nft.cdc":  transfer_nftCdc,
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
	"mint_nft.cdc":      &bintree{mint_nftCdc, map[string]*bintree{}},
	"read_nft_data.cdc": &bintree{read_nft_dataCdc, map[string]*bintree{}},
	"setup_account.cdc": &bintree{setup_accountCdc, map[string]*bintree{}},
	"transfer_nft.cdc":  &bintree{transfer_nftCdc, map[string]*bintree{}},
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
