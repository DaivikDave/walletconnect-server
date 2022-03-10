package relay

import (
	"crypto/sha256"
	"fmt"
)

func Sha256(data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	return fmt.Printf("%x", h.Sum(nil))
}
