package main

import (
	"flag"
	"fmt"
	"log"
	rand "math/rand"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	count := flag.Int("count", 1, "number of rows to generate")
	flag.Parse()
	gofakeit.Seed(0)

	// Add fake lookup functions
	gofakeit.AddFuncLookup("stringdate", gofakeit.Info{
		Category:    "custom",
		Description: "YYYY-MM-DD",
		Example:     "1999-01-08",
		Output:      "string",
		Generate: func(f *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (any, error) {
			return gofakeit.Date().Format("2006-01-02"), nil
		},
	})

	gofakeit.AddFuncLookup("stringdatetime", gofakeit.Info{
		Category:    "custom",
		Description: "YYYY-MM-DDTHH:MM:SS.UUZ",
		Example:     "1985-04-12T23:20:50.52Z",
		Output:      "string",
		Generate: func(f *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (any, error) {
			return gofakeit.Date().Format(time.RFC3339), nil
		},
	})

	gofakeit.AddFuncLookup("stringtime", gofakeit.Info{
		Category:    "custom",
		Description: "[hour]:[minute]:[second].[subsecond]Z",
		Example:     "23:20:50.52Z",
		Output:      "string",
		Generate: func(f *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (any, error) {
			return gofakeit.Date().Format("15:04:05") + "." + fmt.Sprintf("%d", gofakeit.Number(10, 50)) + "Z", nil
		},
	})

	gofakeit.AddFuncLookup("slicestring", gofakeit.Info{
		Category:    "custom",
		Description: "Slice containing strings",
		Example:     "[\"hi\", \"world\"]",
		Output:      "slice", // not sure here...
		Generate: func(f *rand.Rand, m *gofakeit.MapParams, info *gofakeit.Info) (any, error) {
			var S []string
			gofakeit.Slice(&S)
			return S, nil
		},
	})

	// Generate CSV
	value, err := gofakeit.CSV(&gofakeit.CSVOptions{
		RowCount: *count,
		Fields: []gofakeit.Field{
			{Name: "IntField", Function: "number", Params: gofakeit.MapParams{"min": {"-1000"}, "max": {"1000"}}},
			{Name: "NumField", Function: "float64range", Params: gofakeit.MapParams{"min": {"-1000"}, "max": {"1000"}}},
			{Name: "StringField", Function: "noun"},
			{Name: "BoolField", Function: "bool"},
			{Name: "StringDateField", Function: "stringdate"},
			{Name: "StringDateTimeField", Function: "stringdatetime"},
			{Name: "StringTimeField", Function: "stringtime"},
			{Name: "ArrayField", Function: "slicestring"},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(value))
}
