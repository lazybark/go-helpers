package hasher

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/lazybark/go-helpers/fsw"
)

// HashFilePath opens specified file and calls to HashFile
//
// bl is the number of bytes for HashFile to create reading buffer. If provided 0, it will be set for 1024 (1KB)
func HashFilePath(path string, ht HashType, bl int) (hashed string, err error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("[HashFilePath] can not open file -> %w", err)
	}
	defer file.Close()

	hashed, err = HashFile(file, ht, bl)
	if err != nil {
		return "", fmt.Errorf("[HashFilePath] can not make hash -> %w", err)
	}

	return hashed, nil
}

// HashFile uses buffer with specified length to read and hash file as []byte
//
// bl is the number of bytes for HashFile to create reading buffer. If provided 0, it will be set for 1024 (1KB)
func HashFile(file fsw.IFileReader, ht HashType, bl int) (hashed string, err error) {
	r := bufio.NewReader(file)

	var h hash.Hash

	if ht == MD5 {
		h = md5.New()
	} else if ht == SHA1 {
		h = sha1.New()
	} else if ht == SHA256 {
		h = sha256.New()
	} else if ht == SHA512 {
		h = sha512.New()
	}

	// Just use default
	if bl == 0 {
		bl = 1024
	}

	buf := make([]byte, bl)
	n := 0

	for {
		n, err = r.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("[HashFile] can not Read to buffer -> %w", err)
		}
		_, err = h.Write(buf[:n])
		if err != nil {
			return "", fmt.Errorf("[HashFile] can not Write to Hash -> %w", err)
		}
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
