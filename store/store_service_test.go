package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.psqlConnectionPool != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	initialUrl := "https://www.eddywm.com/lets-build-a-url-shortener-in-go/"
	id := int64(5577006791947779410)

	SaveUrlMapping(id, initialUrl)
	retrievedUrl := RetrieveOriginalUrl(id)

	assert.Equal(t, initialUrl, retrievedUrl)
}