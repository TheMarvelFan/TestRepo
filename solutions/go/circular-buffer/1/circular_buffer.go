package circular

import "errors"

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Define the Buffer type here.
type Buffer struct {
    buf []byte
    latest int
    stale int
}

func NewBuffer(size int) *Buffer {
	arr := make([]byte, size)
    return &Buffer{
        stale: -1,
        latest: -1,
        buf: arr,
    }
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.stale == -1 {
        return byte(0), errors.New("Buffer is empty")
    }

    ret := b.buf[b.stale]

    if b.stale == b.latest {
        b.stale = -1
        b.latest = -1
    } else {
    	b.stale = (b.stale + 1) % len(b.buf)
    }

    return ret, nil
}

func (b *Buffer) WriteByte(c byte) error {
	if (b.latest + 1) % len(b.buf) == b.stale {
        return errors.New("Buffer is full")
    }

    if b.stale == -1 {
        b.stale ++
    }

    b.latest = (b.latest + 1) % len(b.buf)
    b.buf[b.latest] = c

    return nil
}

func (b *Buffer) Overwrite(c byte) {
	if (b.latest + 1) % len(b.buf) != b.stale {
        err := b.WriteByte(c)

        if err != nil {
            panic(err)
        }

        return
    }

    b.stale = (b.stale + 1) % len(b.buf)
    b.latest = (b.latest + 1) % len(b.buf)
    b.buf[b.latest] = c
}

func (b *Buffer) Reset() {
	b.latest = -1
    b.stale = -1
}
