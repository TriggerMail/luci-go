// Copyright 2015 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

// AUTOGENERATED. DO NOT EDIT.

// Package main is generated by go.chromium.org/luci/tools/cmd/assets.
//
// It contains all [*.css *.html *.js *.tmpl] files found in the package as byte arrays.
package main

// GetAsset returns an asset by its name. Returns nil if no such asset exists.
func GetAsset(name string) []byte {
	return []byte(files[name])
}

// GetAssetString is version of GetAsset that returns string instead of byte
// slice. Returns empty string if no such asset exists.
func GetAssetString(name string) string {
	return files[name]
}

// GetAssetSHA256 returns the asset checksum. Returns nil if no such asset
// exists.
func GetAssetSHA256(name string) []byte {
	data := fileSha256s[name]
	if data == nil {
		return nil
	}
	return append([]byte(nil), data...)
}

// Assets returns a map of all assets.
func Assets() map[string]string {
	cpy := make(map[string]string, len(files))
	for k, v := range files {
		cpy[k] = v
	}
	return cpy
}

var files = map[string]string{
	"machine-provider-agent.bat.tmpl": string([]byte{58, 58,
		32, 109, 97, 99, 104, 105, 110, 101, 45, 112, 114, 111, 118, 105,
		100, 101, 114, 45, 97, 103, 101, 110, 116, 10, 58, 58, 10, 58,
		58, 32, 82, 117, 110, 115, 32, 116, 104, 101, 32, 77, 97, 99,
		104, 105, 110, 101, 32, 80, 114, 111, 118, 105, 100, 101, 114, 32,
		97, 103, 101, 110, 116, 32, 112, 114, 111, 99, 101, 115, 115, 46,
		10, 10, 64, 123, 123, 46, 65, 103, 101, 110, 116, 125, 125, 32,
		45, 103, 99, 101, 45, 115, 101, 114, 118, 105, 99, 101, 45, 97,
		99, 99, 111, 117, 110, 116, 32, 123, 123, 46, 83, 101, 114, 118,
		105, 99, 101, 65, 99, 99, 111, 117, 110, 116, 125, 125, 32, 45,
		115, 101, 114, 118, 101, 114, 32, 123, 123, 46, 83, 101, 114, 118,
		101, 114, 125, 125, 32, 45, 117, 115, 101, 114, 32, 123, 123, 46,
		85, 115, 101, 114, 125, 125, 10}),
	"machine-provider-agent.conf.tmpl": string([]byte{35, 32,
		109, 97, 99, 104, 105, 110, 101, 45, 112, 114, 111, 118, 105, 100,
		101, 114, 45, 97, 103, 101, 110, 116, 10, 35, 10, 35, 32, 82,
		117, 110, 115, 32, 116, 104, 101, 32, 77, 97, 99, 104, 105, 110,
		101, 32, 80, 114, 111, 118, 105, 100, 101, 114, 32, 97, 103, 101,
		110, 116, 32, 112, 114, 111, 99, 101, 115, 115, 46, 10, 10, 100,
		101, 115, 99, 114, 105, 112, 116, 105, 111, 110, 32, 34, 109, 97,
		99, 104, 105, 110, 101, 32, 112, 114, 111, 118, 105, 100, 101, 114,
		32, 97, 103, 101, 110, 116, 34, 10, 10, 115, 116, 97, 114, 116,
		32, 111, 110, 32, 40, 102, 105, 108, 101, 115, 121, 115, 116, 101,
		109, 32, 97, 110, 100, 32, 110, 101, 116, 45, 100, 101, 118, 105,
		99, 101, 45, 117, 112, 32, 73, 70, 65, 67, 69, 33, 61, 108,
		111, 41, 10, 115, 116, 111, 112, 32, 111, 110, 32, 115, 104, 117,
		116, 100, 111, 119, 110, 10, 10, 115, 99, 114, 105, 112, 116, 10,
		32, 32, 123, 123, 46, 65, 103, 101, 110, 116, 125, 125, 32, 45,
		103, 99, 101, 45, 115, 101, 114, 118, 105, 99, 101, 45, 97, 99,
		99, 111, 117, 110, 116, 32, 123, 123, 46, 83, 101, 114, 118, 105,
		99, 101, 65, 99, 99, 111, 117, 110, 116, 125, 125, 32, 45, 115,
		101, 114, 118, 101, 114, 32, 123, 123, 46, 83, 101, 114, 118, 101,
		114, 125, 125, 32, 45, 117, 115, 101, 114, 32, 123, 123, 46, 85,
		115, 101, 114, 125, 125, 10, 101, 110, 100, 32, 115, 99, 114, 105,
		112, 116, 10, 10, 114, 101, 115, 112, 97, 119, 110, 10, 114, 101,
		115, 112, 97, 119, 110, 32, 108, 105, 109, 105, 116, 32, 117, 110,
		108, 105, 109, 105, 116, 101, 100, 10, 10, 112, 111, 115, 116, 45,
		115, 116, 111, 112, 32, 101, 120, 101, 99, 32, 115, 108, 101, 101,
		112, 32, 53, 10}),
	"machine-provider-agent.service.tmpl": string([]byte{91, 85,
		110, 105, 116, 93, 10, 68, 101, 115, 99, 114, 105, 112, 116, 105,
		111, 110, 61, 77, 97, 99, 104, 105, 110, 101, 32, 80, 114, 111,
		118, 105, 100, 101, 114, 32, 97, 103, 101, 110, 116, 10, 65, 102,
		116, 101, 114, 61, 110, 101, 116, 119, 111, 114, 107, 46, 116, 97,
		114, 103, 101, 116, 10, 10, 91, 83, 101, 114, 118, 105, 99, 101,
		93, 10, 84, 121, 112, 101, 61, 115, 105, 109, 112, 108, 101, 10,
		82, 101, 115, 116, 97, 114, 116, 61, 97, 108, 119, 97, 121, 115,
		10, 82, 101, 115, 116, 97, 114, 116, 83, 101, 99, 61, 53, 10,
		69, 120, 101, 99, 83, 116, 97, 114, 116, 61, 123, 123, 46, 65,
		103, 101, 110, 116, 125, 125, 32, 45, 103, 99, 101, 45, 115, 101,
		114, 118, 105, 99, 101, 45, 97, 99, 99, 111, 117, 110, 116, 32,
		123, 123, 46, 83, 101, 114, 118, 105, 99, 101, 65, 99, 99, 111,
		117, 110, 116, 125, 125, 32, 45, 115, 101, 114, 118, 101, 114, 32,
		123, 123, 46, 83, 101, 114, 118, 101, 114, 125, 125, 32, 45, 117,
		115, 101, 114, 32, 123, 123, 46, 85, 115, 101, 114, 125, 125, 10,
		10, 91, 73, 110, 115, 116, 97, 108, 108, 93, 10, 87, 97, 110,
		116, 101, 100, 66, 121, 61, 109, 117, 108, 116, 105, 45, 117, 115,
		101, 114, 46, 116, 97, 114, 103, 101, 116, 10}),
	"swarming-start-bot.bat.tmpl": string([]byte{58, 58,
		32, 115, 119, 97, 114, 109, 105, 110, 103, 45, 115, 116, 97, 114,
		116, 45, 98, 111, 116, 32, 45, 32, 115, 119, 97, 114, 109, 105,
		110, 103, 32, 98, 111, 116, 32, 115, 116, 97, 114, 116, 117, 112,
		10, 58, 58, 10, 58, 58, 32, 85, 115, 101, 100, 32, 102, 111,
		114, 32, 115, 116, 97, 114, 116, 105, 110, 103, 32, 97, 32, 83,
		119, 97, 114, 109, 105, 110, 103, 32, 98, 111, 116, 32, 112, 114,
		111, 99, 101, 115, 115, 46, 10, 10, 58, 58, 32, 80, 114, 101,
		118, 101, 110, 116, 32, 83, 119, 97, 114, 109, 105, 110, 103, 32,
		102, 114, 111, 109, 32, 99, 111, 110, 102, 105, 103, 117, 114, 105,
		110, 103, 32, 105, 116, 115, 32, 111, 119, 110, 32, 97, 117, 116,
		111, 115, 116, 97, 114, 116, 46, 10, 83, 69, 84, 32, 83, 87,
		65, 82, 77, 73, 78, 71, 95, 69, 88, 84, 69, 82, 78, 65,
		76, 95, 66, 79, 84, 95, 83, 69, 84, 85, 80, 61, 49, 10,
		64, 67, 58, 92, 116, 111, 111, 108, 115, 92, 112, 121, 116, 104,
		111, 110, 92, 98, 105, 110, 92, 112, 121, 116, 104, 111, 110, 46,
		101, 120, 101, 32, 123, 123, 46, 80, 97, 116, 104, 125, 125, 32,
		115, 116, 97, 114, 116, 95, 98, 111, 116, 10}),
	"swarming-start-bot.conf.tmpl": string([]byte{35, 32,
		115, 119, 97, 114, 109, 105, 110, 103, 45, 115, 116, 97, 114, 116,
		45, 98, 111, 116, 32, 45, 32, 115, 119, 97, 114, 109, 105, 110,
		103, 32, 98, 111, 116, 32, 115, 116, 97, 114, 116, 117, 112, 10,
		35, 10, 35, 32, 85, 115, 101, 100, 32, 102, 111, 114, 32, 115,
		116, 97, 114, 116, 105, 110, 103, 32, 97, 32, 83, 119, 97, 114,
		109, 105, 110, 103, 32, 98, 111, 116, 32, 112, 114, 111, 99, 101,
		115, 115, 46, 10, 10, 100, 101, 115, 99, 114, 105, 112, 116, 105,
		111, 110, 32, 34, 115, 119, 97, 114, 109, 105, 110, 103, 32, 98,
		111, 116, 32, 115, 116, 97, 114, 116, 117, 112, 34, 10, 10, 115,
		116, 97, 114, 116, 32, 111, 110, 32, 40, 102, 105, 108, 101, 115,
		121, 115, 116, 101, 109, 32, 97, 110, 100, 32, 110, 101, 116, 45,
		100, 101, 118, 105, 99, 101, 45, 117, 112, 32, 73, 70, 65, 67,
		69, 33, 61, 108, 111, 41, 10, 115, 116, 111, 112, 32, 111, 110,
		32, 115, 104, 117, 116, 100, 111, 119, 110, 10, 10, 115, 99, 114,
		105, 112, 116, 10, 32, 32, 47, 117, 115, 114, 47, 98, 105, 110,
		47, 115, 117, 100, 111, 32, 45, 72, 32, 45, 117, 32, 123, 123,
		46, 85, 115, 101, 114, 125, 125, 32, 83, 87, 65, 82, 77, 73,
		78, 71, 95, 69, 88, 84, 69, 82, 78, 65, 76, 95, 66, 79,
		84, 95, 83, 69, 84, 85, 80, 61, 49, 32, 47, 117, 115, 114,
		47, 98, 105, 110, 47, 112, 121, 116, 104, 111, 110, 32, 123, 123,
		46, 80, 97, 116, 104, 125, 125, 32, 115, 116, 97, 114, 116, 95,
		98, 111, 116, 10, 101, 110, 100, 32, 115, 99, 114, 105, 112, 116,
		10}),
	"swarming-start-bot.service.tmpl": string([]byte{91, 85,
		110, 105, 116, 93, 10, 68, 101, 115, 99, 114, 105, 112, 116, 105,
		111, 110, 61, 83, 119, 97, 114, 109, 105, 110, 103, 32, 98, 111,
		116, 32, 115, 116, 97, 114, 116, 117, 112, 10, 65, 102, 116, 101,
		114, 61, 110, 101, 116, 119, 111, 114, 107, 46, 116, 97, 114, 103,
		101, 116, 10, 10, 91, 83, 101, 114, 118, 105, 99, 101, 93, 10,
		84, 121, 112, 101, 61, 115, 105, 109, 112, 108, 101, 10, 85, 115,
		101, 114, 61, 123, 123, 46, 85, 115, 101, 114, 125, 125, 10, 69,
		110, 118, 105, 114, 111, 110, 109, 101, 110, 116, 61, 83, 87, 65,
		82, 77, 73, 78, 71, 95, 69, 88, 84, 69, 82, 78, 65, 76,
		95, 66, 79, 84, 95, 83, 69, 84, 85, 80, 61, 49, 10, 69,
		120, 101, 99, 83, 116, 97, 114, 116, 61, 47, 117, 115, 114, 47,
		98, 105, 110, 47, 112, 121, 116, 104, 111, 110, 32, 123, 123, 46,
		80, 97, 116, 104, 125, 125, 32, 115, 116, 97, 114, 116, 95, 98,
		111, 116, 10, 10, 91, 73, 110, 115, 116, 97, 108, 108, 93, 10,
		87, 97, 110, 116, 101, 100, 66, 121, 61, 109, 117, 108, 116, 105,
		45, 117, 115, 101, 114, 46, 116, 97, 114, 103, 101, 116, 10}),
}

var fileSha256s = map[string][]byte{
	"machine-provider-agent.bat.tmpl": {125, 73,
		232, 119, 61, 107, 159, 152, 105, 88, 23, 239, 10, 24, 85, 68,
		59, 36, 105, 114, 195, 228, 186, 115, 85, 238, 180, 117, 224, 170,
		14, 199},
	"machine-provider-agent.conf.tmpl": {71, 237,
		73, 176, 85, 229, 154, 146, 140, 247, 128, 215, 158, 161, 19, 232,
		215, 97, 195, 187, 118, 146, 142, 22, 238, 206, 193, 180, 34, 13,
		214, 100},
	"machine-provider-agent.service.tmpl": {228, 23,
		38, 10, 145, 147, 50, 183, 167, 180, 244, 49, 230, 17, 98, 242,
		218, 68, 39, 78, 221, 199, 141, 24, 242, 67, 73, 92, 167, 222,
		8, 43},
	"swarming-start-bot.bat.tmpl": {223, 155,
		52, 23, 163, 162, 69, 31, 35, 140, 191, 28, 4, 51, 228, 141,
		7, 182, 247, 255, 182, 222, 218, 80, 171, 191, 144, 114, 152, 108,
		186, 166},
	"swarming-start-bot.conf.tmpl": {106, 84,
		13, 227, 223, 216, 36, 27, 173, 38, 118, 130, 200, 138, 232, 101,
		160, 62, 143, 0, 33, 151, 4, 194, 227, 218, 41, 167, 189, 63,
		113, 177},
	"swarming-start-bot.service.tmpl": {96, 182,
		122, 152, 147, 106, 192, 117, 43, 172, 201, 101, 137, 30, 94, 130,
		72, 57, 135, 160, 136, 124, 104, 104, 126, 204, 155, 163, 245, 128,
		140, 229},
}
