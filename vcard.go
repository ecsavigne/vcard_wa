package vcard

import (
	"fmt"
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
	fmt.Println("ParseVCard")
	value, _ := io.ReadAll(val)
	cad := clean(string(value))
	fmt.Println(":\n", cad, "\n:")
	re := regexp.MustCompile(`(?m)^(` + strings.Join(tags, "|") + `)[^:\n]*:.*`)
	matches := re.FindAllString(cad, -1)
	vc := &VCard{value: cad, partRecord: matches}
	vc.createRecord()
	// return vc
	return nil
}

func clean(value string) string {
	return strings.NewReplacer(replacer...).Replace(value)
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
