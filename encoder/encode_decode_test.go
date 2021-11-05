package encoder

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestEncodingAndDecoding(t *testing.T) {
	rand.Seed(0)
	id := rand.Int63()

	encodedUrl := Encode(id)
	decodedId := Decode(encodedUrl)

	fmt.Println(encodedUrl)
	assert.Equal(t, id, decodedId)
}

