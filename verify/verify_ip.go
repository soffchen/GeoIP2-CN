package main

import (
	"fmt"
	"log"
	"net"

	"github.com/oschwald/geoip2-golang"
)

const DATA_FILE = "Country.mmdb"

func main() {
	db, err := geoip2.Open(DATA_FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test IP list contains both Chinese and international IPs for verification
	// IPs from GFW that are classified as HK in ipip but China in maxmind:
	// 103.200.30.143, 103.228.130.61, 216.58.200.238, 103.200.30.245,
	// 118.184.26.113, 103.200.31.172, 69.171.235.101
	// Example IPs from China: 123.126.55.41, 117.23.61.238
	var list = []string{"103.200.30.143", "103.228.130.61", "216.58.200.238",
		"103.200.30.245", "118.184.26.113", "103.200.31.172", "69.171.235.101",
		"123.126.55.41", "117.23.61.238"}

	for _, ipTxt := range list {
		ip := net.ParseIP(ipTxt)
		record, err := db.Country(ip)
		if err != nil || record == nil {
			log.Fatal(err)
		}

		fmt.Printf("IP:%s - Locale:%s\n", ipTxt, record.Country.IsoCode)
	}
}
