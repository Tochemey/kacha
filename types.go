/*
 * MIT License
 *
 * Copyright (c) 2025 Arsene Tochemey Gandote
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package distcache

import "time"

// KV represents a key-value pair.
//
// The Key field holds the identifier for the value, while Value stores the data
// associated with that key.
type KV struct {
	// Key is the identifier for the value.
	Key string

	// Value contains the data associated with the key.
	Value []byte
}

// Entry represents a key-value pair with an optional expiration time.
//
// The Key field holds the identifier for the value, while Value stores the data
// associated with that key. Expiry is a Unix timestamp (in seconds) indicating
// when the key-value pair should expire. An Expiry value of 0 typically means
// that the pair does not expire.
type Entry struct {
	KV
	// Expiry is the expiration time of the key-value pair, represented as a Unix timestamp.
	// If Expiry is set to the zero value (time.Time{}), the entry does not expire.
	Expiry time.Time
}
