// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package onewire

type Master struct {
	DCI
}

func (m *Master) Write(data []byte) (int, error) {
	for n, b := range data {
		if err := m.WriteByte(b); err != nil {
			return n, err
		}
	}
	return len(data), nil
}

// Read behaves link io.ReadFull (reads len(data) bytes or returns error).
func (m *Master) Read(data []byte) (int, error) {
	for n := range data {
		b, err := m.ReadByte()
		if err != nil {
			return n, err
		}
		data[n] = b
	}
	return len(data), nil

}
