package fuzzes

import (
	"bytes"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"testing"

	avro "github.com/kjuulh/go-avro"
)

func TestGenerateSpecificComplexFuzz(t *testing.T) {
	const folder = "specificreadercomplex/corpus/"
	w := avro.NewDatumWriter(ComplexSchema)

	var buf bytes.Buffer
	var fixed16 = []byte("0123456789abcdef")

	writeOut := func(name string, v *Complex) {
		if v.FixedField == nil {
			v.FixedField = fixed16
		}
		if v.EnumField == nil {
			v.EnumField = NewComplexEnumField()
			v.EnumField.SetIndex(3)
		}
		buf.Reset()
		err := w.Write(v, avro.NewBinaryEncoder(&buf))
		if err != nil {
			log.Fatal(err)
		}
		ioutil.WriteFile(folder+name, buf.Bytes(), 0644)
	}

	writeOut("strings-only.bin", &Complex{
		StringArray: []string{"abc", "def", "ghi", "jkl"},
		FixedField:  fixed16,
	})
	writeOut("longs-only.bin", &Complex{LongArray: []int64{978, -1, math.MaxInt64, math.MinInt64}})
	writeOut("map-ints.bin", &Complex{
		MapOfInts: map[string]int32{
			"aaa": 485,
			"bbb": math.MaxInt32,
			"ccc": math.MinInt32,
		},
	})
	writeOut("union-string.bin", &Complex{
		UnionField: "AAAAAAAAAABCDEF",
	})
	writeOut("union-bool.bin", &Complex{
		UnionField: true,
	})
}

var fixed16 = []byte("0123456789abcdef")

func fixComplex(v *Complex) {
	if v.FixedField == nil {
		v.FixedField = fixed16
	}
	if v.EnumField == nil {
		v.EnumField = NewComplexEnumField()
		v.EnumField.SetIndex(int32(rand.Intn(4)))
	}
}

func TestGenerateGenericFuzz(t *testing.T) {
	const folder = "genericreader/corpus/"
	w := avro.NewDatumWriter(CombinedSchema)

	var buf bytes.Buffer

	writeOut := func(name string, v *Combined) {
		if v.Complex != nil {
			fixComplex(v.Complex)
		}
		buf.Reset()
		err := w.Write(v, avro.NewBinaryEncoder(&buf))
		if err != nil {
			log.Fatal(err)
		}
		ioutil.WriteFile(folder+name, buf.Bytes(), 0644)
	}

	writeOut("strings-only.bin", &Combined{
		Complex: &Complex{
			StringArray: []string{"abc", "def", "ghi", "jkl"},
		},
	})
	writeOut("longs-only.bin", &Combined{
		Complex: &Complex{LongArray: []int64{978, -1, math.MaxInt64, math.MinInt64}},
	})
	writeOut("map-ints.bin", &Combined{
		Complex: &Complex{
			MapOfInts: map[string]int32{
				"aaa": 485,
				"bbb": math.MaxInt32,
				"ccc": math.MinInt32,
			},
		},
	})
	writeOut("primitives.bin", &Combined{
		Primitive: &Primitive{
			BooleanField: true,
			DoubleField:  4.8,
			FloatField:   9.2,
			StringField:  "abcdefg",
		},
	})
}
