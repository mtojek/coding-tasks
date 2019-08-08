package coding_tasks

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestParseFile(t *testing.T) {
	f, err := ioutil.TempFile(".", "statsd-log-file-")
	if err != nil {
		log.Fatal(err)
	}

	f.WriteString(`test.gauge:1.000000|g|@0.999999
test.gauge:1.000000|g|#tagA:valueA,tagB:valueB
test.gauge:1.000000|g|@0.999999|#tagA:valueA
users.online:1|c|@0.5|#country:china`)
	f.Sync()
	ParseFile(f.Name())

	os.Remove(f.Name())
}