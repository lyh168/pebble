// Copyright 2018 The LevelDB-Go and Pebble Authors. All rights reserved. Use
// of this source code is governed by a BSD-style license that can be found in
// the LICENSE file.

package pebble

import (
	"github.com/petermattis/pebble/db"
	"github.com/petermattis/pebble/internal/rangedel"
)

type rangeDelLevel struct {
	iter internalIterator
	val  rangedel.Tombstone
}

func (l *rangeDelLevel) init(iter internalIterator) {
	l.iter = iter
}

// rangeDelMap provides a merged view of the range tombstones from
// multiple levels.
type rangeDelMap struct {
	// The sequence number at which reads are being performed. Tombstones that
	// are newer than this sequence number are ignored.
	seqNum uint64
	levels []rangeDelLevel
}

// ClearCache clears any cached tombstone information so that the next call to
// Deleted or Get will have to repopulate the cache. This is useful when the
// next call to Deleted or Get is expected to not hit the cache due to a Seek
// being performed on the iterator using this rangeDelMap.
func (m *rangeDelMap) ClearCache() {
}

// Deleted returns true if the specified key is covered by a newer range
// tombstone.
func (m *rangeDelMap) Deleted(key db.InternalKey) bool {
	return m.Get(key).Start.SeqNum() >= key.SeqNum()
}

// Get the range tombstone at the specified key. A tombstone is always
// returned, though it may cover an empty range of keys or the sequence number
// may be 0 to indicate that no tombstone covers the specified key.
func (m *rangeDelMap) Get(key db.InternalKey) rangedel.Tombstone {
	return rangedel.Tombstone{}
}