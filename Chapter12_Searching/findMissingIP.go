package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/damnever/bitarray"
)

func findMissingIP() {
	var ipLookup [65536]int
	var ipBitArray [65536]*bitarray.BitArray
	for i := range ipBitArray {
		ipBitArray[i] = bitarray.New(65536)
	}
	fs, err := os.Open("./geoip2-ipv4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fs.Close()
	fileScanner := bufio.NewScanner(fs)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		ipWords := strings.Split(line, ".")
		if len(ipWords) == 4 {
			a, err := strconv.Atoi(ipWords[0])
			if err != nil {
				log.Println(err)
				continue
			}
			b, err := strconv.Atoi(ipWords[1])
			if err != nil {
				log.Println(err)
				continue
			}
			ab := ((a << 8) & 0xff00) ^ (b & 0x00ff)
			// fmt.Println(ab)
			ipLookup[ab]++
		}
	}

	for ab, ipCount := range ipLookup {
		a := (ab >> 8) & 0x00ff
		b := (ab) & 0x00ff
		if ipCount == 0 {
			// ips with starting with these many 16 bit msb doesen't exist
			// fmt.Printf("%d.%d.c.d\n", a, b)
		} else {
			fs.Seek(0, 0)
			fileScanner = bufio.NewScanner(fs)

			for fileScanner.Scan() {
				line := fileScanner.Text()
				if err != nil {
					log.Println(err)
					continue
				}

				ip16bitMSB := fmt.Sprintf("%d.%d", a, b)
				if strings.Contains(line, ip16bitMSB) {

					ipWords := strings.Split(line, ".")
					if len(ipWords) == 4 {
						c, err := strconv.Atoi(ipWords[2])
						if err != nil {
							log.Println(err)
							continue
						}
						d, err := strconv.Atoi(ipWords[3])
						if err != nil {
							log.Println(err)
							continue
						}
						cd := ((c << 8) & 0xff00) ^ (d & 0x00ff)
						_, err = ipBitArray[ab].Put(cd, 1)
						if err != nil {
							fmt.Println(err)
						}
					}
				}

			}
		}
	}

	for ab, ipCount := range ipLookup {
		if ipCount > 0 {
			a := (ab >> 8) & 0x00ff
			b := (ab) & 0x00ff
			fmt.Printf("%d.%d ips : %d %d\n", a, b, ipCount, ipBitArray[ab].Count())
		}
	}
	fmt.Println("END")
}
