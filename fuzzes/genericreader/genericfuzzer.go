package genericreader

import (
	"bytes"

	avro "github.com/kjuulh/go-avro"
	"github.com/kjuulh/go-avro/fuzzes"
)

var buf bytes.Buffer
var reader = avro.NewDatumReader(fuzzes.ComplexSchema)

func Fuzz(input []byte) int {
	var dest *avro.GenericRecord
	err := reader.Read(&dest, avro.NewBinaryDecoder(input))
	if err != nil {
		return 0
	}
	return 1
}
