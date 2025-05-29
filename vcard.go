package main

import (
	"io"
	"maps"
	"regexp"
	"slices"
	"strings"
)

type VCard struct {
	value      string
	record     Record
	keys       []string
	partRecord []string
}

func ParseVCard(val io.Reader) *VCard {
	value, _ := io.ReadAll(val)
	cad := string(value)
	re := regexp.MustCompile(`(?m)^(` + strings.Join(tags, "|") + `)[^:\n]*:.*`)
	matches := re.FindAllString(cad, -1)
	vc := &VCard{value: cad, partRecord: matches}
	vc.createRecord()
	return vc
}

func (vc *VCard) createRecord() {
	vc.record = make(Record)
	for _, m := range vc.partRecord {
		tag := strings.Split(m, ":")[0]
		value := strings.Split(m, ":")[1]
		vc.record[tag] = value
	}
}

func (vc VCard) GetRecords() Record {
	return vc.record
}

func (vc VCard) GetTags() []string {
	if vc.keys == nil {
		vc.keys = slices.Collect(maps.Keys(vc.record))
	}
	return vc.keys
}

func (vc VCard) GetTagValue(tag string) string {
	return vc.record[tag]
}

// @tagContent: Substring part of a tag
func (vc VCard) GetTagContentValues(tagContent string) string {
	for _, k := range vc.GetTags() {
		if strings.Contains(k, tagContent) {
			return k
		}
	}
	return ""
}

func (vc VCard) GetPartRecord() []string {
	return vc.partRecord
}
