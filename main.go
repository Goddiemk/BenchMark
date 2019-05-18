package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

// Speedtest returns speed in random float values
type Speedtest struct{}

func (s *Speedtest) downloadSpeed() float64 {
	return rand.Float64()
}

func (s *Speedtest) uploadSpeed() float64 {
	return rand.Float64()
}

func (s *Speedtest) ping() float64 {
	return rand.Float64()
}

// SpeedData is the field format for csv
type SpeedData struct {
	Speed float64   `csv:"speed"`
	Time  time.Time `csv:"time"`
}

var downloadData []SpeedData
var uploadData []SpeedData
var pingData []SpeedData

// BenchMark holds a number of benchmarking methods
type BenchMark struct{}

func (b *BenchMark) hourly(download, upload, ping string) {
	s := &Speedtest{}
	downloadData = nil
	uploadData = nil
	pingData = nil

	downloadData = append(downloadData, SpeedData{s.downloadSpeed(), time.Now()})
	uploadData = append(uploadData, SpeedData{s.uploadSpeed(), time.Now()})
	pingData = append(pingData, SpeedData{s.ping(), time.Now()})

	csvOperation(download, downloadData)
	csvOperation(upload, uploadData)
	csvOperation(ping, pingData)
}

func (b *BenchMark) downloadAverage(csv string, t time.Time) float64 {
	var total float64
	var download []SpeedData

	downloadFile, _ := os.Open(csv)
	gocsv.UnmarshalFile(downloadFile, &download)

	for _, value := range download {
		if value.Time.Format("2006-01-02") == t.Format("2006-01-02") {
			total += value.Speed
		}
	}
	return (total / float64(len(download)))
}

func (b *BenchMark) uploadAverage(csv string, t time.Time) float64 {
	var total float64
	var upload []SpeedData

	uploadFile, _ := os.Open(csv)
	gocsv.UnmarshalFile(uploadFile, &upload)

	for _, value := range upload {
		if value.Time.Format("2006-01-02") == t.Format("2006-01-02") {
			total += value.Speed
		}
	}
	return (total / float64(len(upload)))
}

func (b *BenchMark) pingAverage(csv string, t time.Time) float64 {
	var total float64
	var ping []SpeedData

	pingFile, _ := os.Open(csv)
	gocsv.UnmarshalFile(pingFile, &ping)

	for _, value := range ping {
		if value.Time.Format("2006-01-02") == t.Format("2006-01-02") {
			total += value.Speed
		}
	}
	return (total / float64(len(ping)))
}

func (b *BenchMark) downloads(csv string, t time.Time) []SpeedData {
	var download []SpeedData
	var result []SpeedData

	downloadFile, _ := os.Open(csv)
	gocsv.UnmarshalFile(downloadFile, &download)

	if t.IsZero() {
		return download
	}

	for _, value := range download {
		if value.Time.Format("2006-01-02") == t.Format("2006-01-02") {
			result = append(result, SpeedData{value.Speed, value.Time})
		}
	}
	return result
}

func (b *BenchMark) uploads(csv string, t time.Time) []SpeedData {
	var upload []SpeedData
	var result []SpeedData

	uploadFile, _ := os.Open(csv)
	gocsv.UnmarshalFile(uploadFile, &upload)

	if t.IsZero() {
		return upload
	}

	for _, value := range upload {
		if value.Time.Format("2006-01-02") == t.Format("2006-01-02") {
			result = append(result, SpeedData{value.Speed, value.Time})
		}
	}
	return result
}

func (b *BenchMark) pings(csv string, t time.Time) []SpeedData {
	var ping []SpeedData
	var result []SpeedData

	pingFile, _ := os.Open(csv)
	gocsv.UnmarshalFile(pingFile, &ping)

	if t.IsZero() {
		return ping
	}

	for _, value := range ping {
		if value.Time.Format("2006-01-02") == t.Format("2006-01-02") {
			result = append(result, SpeedData{value.Speed, value.Time})
		}
	}
	return result
}

// csvOperation handles csv file write operations
func csvOperation(csv string, data []SpeedData) {
	var err error
	var file *os.File
	var createFlags = os.O_CREATE | os.O_APPEND | os.O_WRONLY
	var appendFlags = os.O_APPEND | os.O_WRONLY

	file, err = os.OpenFile(csv, appendFlags, 0600)
	switch err {
	case nil:
		gocsv.MarshalWithoutHeaders(&data, file)
	default:
		if file, err = os.OpenFile(csv, createFlags, 0600); err == nil {
			gocsv.MarshalFile(&data, file)
		}
	}
}

func main() {
}
