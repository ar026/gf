// Copyright 2019 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

// Package gres provides resource management and packing/unpacking feature between files and bytes.
package gres

const (
	gPACKAGE_TEMPLATE = `package %s

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add(%s); err != nil {
		panic(err)
	}
}
`
)

var (
	// Default resource object.
	defaultResource = New()
)

// Default returns the default resource object.
func Default() *Resource {
	return defaultResource
}

// Add unpacks and adds the <content> into the default resource object.
// The unnecessary parameter <prefix> indicates the prefix
// for each file storing into current resource object.
func Add(content []byte, prefix ...string) error {
	return defaultResource.Add(content, prefix...)
}

// Load loads, unpacks and adds the data from <path> into the default resource object.
// The unnecessary parameter <prefix> indicates the prefix
// for each file storing into current resource object.
func Load(path string, prefix ...string) error {
	return defaultResource.Load(path, prefix...)
}

// Get returns the file with given path.
func Get(path string) *File {
	return defaultResource.Get(path)
}

// Contains checks whether the <path> exists in the default resource object.
func Contains(path string) bool {
	return defaultResource.Contains(path)
}

// Scan returns the files under the given path, the parameter <path> should be a folder type.
//
// The pattern parameter <pattern> supports multiple file name patterns,
// using the ',' symbol to separate multiple patterns.
//
// It scans directory recursively if given parameter <recursive> is true.
func Scan(path string, pattern string, recursive ...bool) []*File {
	return defaultResource.Scan(path, pattern, recursive...)
}

// Dump prints the files of the default resource object.
func Dump() {
	defaultResource.Dump()
}