package braingames

import (
	cryptorand "crypto/rand"
	"encoding/binary"
	rand "math/rand"
)

func SetSeed() {
	var b [8]byte
	_, err := cryptorand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package")
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}
