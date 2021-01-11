// Copyright 2021 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package component

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// FileSystem struct
type FileSystem struct {
	ctx context.Context
}

// NewFileSystem creates a new instance
func NewFileSystem(ctx context.Context) *FileSystem {
	fs := &FileSystem{
		ctx: ctx,
	}

	return fs
}

// WithCorrelation adds correlation id to context
func (fs *FileSystem) WithCorrelation(correlation string) *FileSystem {
	fs.ctx = context.WithValue(
		fs.ctx,
		CorralationID,
		correlation,
	)

	return fs
}

// ReadFile get the file content
func (fs *FileSystem) ReadFile(path string) (string, error) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

// EnsureTrailingSlash ensure there is a trailing slash
func (fs *FileSystem) EnsureTrailingSlash(dir string) string {
	return fmt.Sprintf(
		"%s%s",
		strings.TrimRight(dir, string(os.PathSeparator)),
		string(os.PathSeparator),
	)
}

// RemoveTrailingSlash removes any trailing slash
func (fs *FileSystem) RemoveTrailingSlash(dir string) string {
	return strings.TrimRight(dir, string(os.PathSeparator))
}

// RemoveStartingSlash removes any starting slash
func (fs *FileSystem) RemoveStartingSlash(dir string) string {
	return strings.TrimLeft(dir, string(os.PathSeparator))
}

// ClearDir removes all files and sub dirs
func (fs *FileSystem) ClearDir(dir string) error {
	files, err := filepath.Glob(filepath.Join(dir, "*"))

	if err != nil {
		return err
	}

	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return err
		}
	}

	return nil
}

// StoreFile stores a file content
func (fs *FileSystem) StoreFile(path, content string) error {
	dir := filepath.Dir(path)

	err := os.MkdirAll(dir, 0775)

	if err != nil {
		return err
	}

	f, err := os.Create(path)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(content)

	return err
}

// PathExists reports whether the path exists
func (fs *FileSystem) PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

// FileExists reports whether the named file exists
func (fs *FileSystem) FileExists(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		if fi.Mode().IsRegular() {
			return true
		}
	}

	return false
}

// DirExists reports whether the dir exists
func (fs *FileSystem) DirExists(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		if fi.Mode().IsDir() {
			return true
		}
	}

	return false
}

// EnsureDir ensures that directory exists
func (fs *FileSystem) EnsureDir(dirName string, mode int) error {
	err := os.MkdirAll(dirName, os.FileMode(mode))

	if err == nil || os.IsExist(err) {
		return nil
	}

	return err
}

// DeleteFile deletes a file
func (fs *FileSystem) DeleteFile(path string) error {
	return os.Remove(path)
}

// GetHostname gets the hostname
func (fs *FileSystem) GetHostname() (string, error) {
	hostname, err := os.Hostname()

	if err != nil {
		return "", err
	}

	return strings.ToLower(hostname), nil
}

// DeleteDir deletes a dir
func (fs *FileSystem) DeleteDir(dir string) error {
	err := os.RemoveAll(dir)

	if err != nil {
		return err
	}

	return nil
}
