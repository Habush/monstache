package main

import (
	"github.com/rwynn/monstache/monstachemap"
)

func Map(input *monstachemap.MapperPluginInput) (output *monstachemap.MapperPluginOutput, err error) {
	doc := input.Document
	switch input.Collection {

	case "clinvar":
		fields := []string{"chrom", "pos", "id", "ref", "alt", "hgvs"}
		for key, _ := range doc {
			if !contains(fields, key) {
				delete(doc, key)
			}
		}
		break
	case "genes":
		fields := []string{"chrom", "id", "start", "end", "symbol", "entrezID"}

		for key, _ := range doc {
			if !contains(fields, key) {
				delete(doc, key)
			}
		}
		break;

	case "exac":
		fields := []string{"chrom", "pos", "ref", "alt", "hgvs"}
		for key, _ := range doc {
			if !contains(fields, key) {
				delete(doc, key)
			}
		}

		break;

	case "effect":
		fields := []string{"chrom", "pos", "ref", "alt", "hgvs", "rsId"}
		for key, _ := range doc {
			if !contains(fields, key) {
				delete(doc, key)
			}
		}

		break;
	case "g1k":
		fields := []string{"chrom", "pos", "ref", "alt", "hgvs", "rsId"}
		for key, _ := range doc {
			if !contains(fields, key) {
				delete(doc, key)
			}
		}

		break;
	case "dbnsfp":
		fields := []string{"chrom", "pos", "ref", "alt", "hgvs"}
		for key, _ := range doc {
			if !contains(fields, key) {
				delete(doc, key)
			}
		}
		break;
	case "acmg":
		fields := []string{"chrom", "pos", "ref", "alt", "hgvs"}
		for key, _ := range doc {
			if !contains(fields, key) {
				delete(doc, key)
			}
		}
		break;
	}

	output = &monstachemap.MapperPluginOutput{Document: doc}
	return
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
