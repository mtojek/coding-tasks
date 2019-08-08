package coding_tasks

import "strings"

/**
 * Given a string containing different sentences with keys (strings) and a list of keys (strings),
 * find all the sentences which contains the input keys (all of them) and return a list
 * with the other keys in this selected sentences without any duplicate.  
 */

func FindOtherKeys(sentences []string, keys []string) (unknown []string) {
	keysMap := map[string]bool{}
	for _, k := range keys {
		keysMap[k] = true
	}

	sentencesWithKeys := filterSentencesWithKeys(sentences, keysMap)

	unkMap := map[string]bool{}

	for _, sentence := range sentencesWithKeys {
		words := strings.Split(sentence, " ")
		for _, word := range words {
			if _, ok := keysMap[word]; !ok {
				unkMap[word] = true
			}
		}
	}

	for u := range unkMap {
		unknown = append(unknown, u)
	}
	return
}

func filterSentencesWithKeys(sentences []string, keys map[string]bool) (filtered []string) {
	for _, sentence := range sentences {
		words := strings.Split(sentence, " ")
		for _, word := range words {
			if _, ok := keys[word]; ok {
				filtered = append(filtered, sentence)
				break
			}
		}
	}
	return
}