package coding_tasks

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/**
 * Sample log lines:
 *
 * test.gauge:1.000000|g|@0.999999
 * test.gauge:1.000000|g|#tagA:valueA,tagB:valueB
 * test.gauge:1.000000|g|@0.999999|#tagA:valueA
 * users.online:1|c|@0.5|#country:china
 */

type Tag struct {
	Name string
	Value string
}

func ParseFile(logFile string) {
	f, err := os.Open(logFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		tags := ParseLine(line)

		fmt.Println(tags)
	}
}

func ParseLine(line string) (tags []Tag) {
	tagsIndex := strings.Index(line, "#")
	if tagsIndex < 0 {
		return
	}

	tagNameValues := strings.Split(line[(tagsIndex+1):], ",")
	for _, tagNameValue := range tagNameValues {
		vals := strings.Split(tagNameValue, ":")
		tags = append(tags, Tag{
			Name: vals[0],
			Value: vals[1],
		})
	}
	return
}