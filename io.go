/*	Copyright (c) 2022-present, Serhat Şevki Dinçer.
	This Source Code Form is subject to the terms of the Mozilla Public
	License, v. 2.0. If a copy of the MPL was not distributed with this
	file, You can obtain one at http://mozilla.org/MPL/2.0/.
*/

package rng

import (
	"unsafe"

	"github.com/jfcg/sixb"
)

// Fill buf with random bytes
func Fill(buf []byte) {
	// fill 8 bytes at a time
	ls := sixb.BtoU8(buf)
	for i := 0; i < len(ls); i++ {
		ls[i] = Get()
	}

	if r := len(ls) << 3; r < len(buf) {
		u := Get()
		src := (*[8]byte)(unsafe.Pointer(&u))[:]
		// fill 1 to 7 remaining bytes
		copy(buf[r:], src)
	}
}
