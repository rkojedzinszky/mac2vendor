package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/rkojedzinszky/mac2vendor/radix"
)

// OUIDescr describes an OUI entry
type OUIDescr struct {
	Prefix   string `json:"prefix"`
	Vendor   string `json:"vendor"`
	Comments string `json:"comments"`
}

// read a gzip compressed file
func readOuiFile(ouiPath string) (radix.Readonly, error) {
	fh, err := os.Open(ouiPath)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	gzreader, err := gzip.NewReader(fh)
	if err != nil {
		return nil, err
	}
	defer gzreader.Close()

	tree := radix.New()

	commentre := regexp.MustCompile("\\s*#.*")
	re := regexp.MustCompile("^([0-9a-fA-F:/-]+)\\s+(\\S+)\\s+(.*?)\\s*$")
	prefixmatch := regexp.MustCompile("^([0-9a-fA-F:-]+)/(\\d+)$")
	stripre := regexp.MustCompile("[:-]")

	scanner := bufio.NewScanner(gzreader)
	for scanner.Scan() {
		line := commentre.ReplaceAllString(scanner.Text(), "")
		subs := re.FindStringSubmatch(line)
		if subs == nil {
			continue
		}

		oui := &OUIDescr{
			Prefix:   subs[1],
			Vendor:   subs[2],
			Comments: subs[3],
		}

		prefix := oui.Prefix
		plen := 24

		pmatch := prefixmatch.FindStringSubmatch(prefix)
		if pmatch != nil {
			prefix = pmatch[1]
			if len, err := strconv.Atoi(pmatch[2]); err == nil {
				plen = len
			} else {
				return nil, err
			}
		}

		prefix = strings.ToLower(stripre.ReplaceAllString(prefix, ""))
		if plen%4 != 0 {
			return nil, fmt.Errorf("Prefix length not multiple of 4")
		}

		prefix = prefix[:plen/4]

		tree.Add(prefix, oui)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return tree.Readonly(), nil
}
