package main

import (
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/bouk/monkey"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Remove("test_download.csv")
	os.Remove("test_upload.csv")
	os.Remove("test_ping.csv")
	rand.Seed(time.Now().UTC().UnixNano())
	os.Exit(m.Run())
}

func TestSuccessfulDownload(t *testing.T) {
	patchDate := time.Now().AddDate(0, 0, -3)
	patch := monkey.Patch(time.Now, func() time.Time { return patchDate })
	test := &BenchMark{}

	test.hourly("test_download.csv", "test_upload.csv", "test_ping.csv")
	test.hourly("test_download.csv", "test_upload.csv", "test_ping.csv")

	assert.Len(t, test.downloads("test_download.csv", time.Now()), 2)
	patch.Unpatch()

	test.hourly("test_download.csv", "test_upload.csv", "test_ping.csv")
	test.hourly("test_download.csv", "test_upload.csv", "test_ping.csv")
	test.hourly("test_download.csv", "test_upload.csv", "test_ping.csv")

	assert.Len(t, test.downloads("test_download.csv", time.Now()), 3)
	assert.Len(t, test.downloads("test_download.csv", time.Time{}), 5)
	assert.NotZero(t, test.downloadAverage("test_download.csv", time.Now()))

	pastDate := time.Now().AddDate(0, 0, -1)
	assert.Len(t, test.downloads("test_download.csv", pastDate), 0)
	assert.Zero(t, test.downloadAverage("test_download.csv", pastDate))
}

func TestSuccessfulUpload(t *testing.T) {
	patchDate := time.Now().AddDate(0, 0, -3)
	patch := monkey.Patch(time.Now, func() time.Time { return patchDate })
	test := &BenchMark{}

	assert.Len(t, test.uploads("test_upload.csv", time.Now()), 2)
	patch.Unpatch()

	assert.Len(t, test.uploads("test_upload.csv", time.Now()), 3)
	assert.Len(t, test.uploads("test_upload.csv", time.Time{}), 5)
	assert.NotZero(t, test.uploadAverage("test_upload.csv", time.Now()))

	pastDate := time.Now().AddDate(0, 0, -1)
	assert.Len(t, test.uploads("test_upload.csv", pastDate), 0)
	assert.Zero(t, test.uploadAverage("test_upload.csv", pastDate))
}

func TestSuccessfulPing(t *testing.T) {
	patchDate := time.Now().AddDate(0, 0, -3)
	patch := monkey.Patch(time.Now, func() time.Time { return patchDate })
	test := &BenchMark{}

	assert.Len(t, test.pings("test_ping.csv", time.Now()), 2)
	patch.Unpatch()

	assert.Len(t, test.pings("test_ping.csv", time.Now()), 3)
	assert.Len(t, test.pings("test_ping.csv", time.Time{}), 5)
	assert.NotZero(t, test.pingAverage("test_ping.csv", time.Now()))

	pastDate := time.Now().AddDate(0, 0, -1)
	assert.Len(t, test.pings("test_ping.csv", pastDate), 0)
	assert.Zero(t, test.pingAverage("test_ping.csv", pastDate))
}
