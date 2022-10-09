// Copyright 2022 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package onewire

import "errors"

var (
	ErrNoResponse = errors.New("no response")
	ErrBusFault   = errors.New("bus fault")
	ErrCRC        = errors.New("bad CRC")
	ErrDevType    = errors.New("bad device type")
)

// DCI means Data and Control Interface. It contains set of methods that
// should provide access to 1-Wire low-level signaling. It is used by Master
// to implement higher layer protocols.
type DCI interface {
	// Reset sends reset pulse. If there is no presence pulse received it
	// returns ErrNoResponse error.
	Reset() error

	// ReadBit generates read time slot on the bus. It returns received bit
	// value (0 or 1) or error.
	ReadBit() (bit int, err error)

	// WriteBit generates write slot on the bus. It sends 0 if bit&1 == 0 or 1
	// otherwise.
	WriteBit(bit int) error

	// ReadByte receives a byte by generating 8 read slots on the bus. It
	// returns read byte or error.
	ReadByte() (byte, error)

	// WriteByte sends b by generating 8 write slots on the bus.
	WriteByte(b byte) error
}
