package common

import (
	"fmt"
	"testing"

	"github.com/Noofbiz/tmx"
)

// DecoderTest holds test data that encodes to the same expected result
// for each encoding
type DecoderTest struct {
	data     []tmx.Data
	expected []uint32
}

// decoderTests defines Input and expected output for testing the decoder.
var decoderTests = map[string]DecoderTest{
	"NoTiles": DecoderTest{
		expected: []uint32{0, 0, 0, 0},
		data: []tmx.Data{ // generated via Tiled 0.10.1 from 2x2 empty map
			tmx.Data{
				Encoding:    "base64",
				Compression: "zlib",
				Inner:       "eJxjYEAFAAAQAAE=",
			},
			tmx.Data{
				Encoding:    "base64",
				Compression: "gzip",
				Inner:       "H4sIAAAAAAAAA2NgQAUAVUu77BAAAAA=",
			},
			tmx.Data{
				Encoding:    "base64",
				Compression: "",
				Inner:       "AAAAAAAAAAAAAAAAAAAAAA==",
			},
			tmx.Data{
				Encoding:    "csv",
				Compression: "",
				Inner:       "0,0\n0,0",
			},
			tmx.Data{
				Encoding:    "",
				Compression: "",
				Inner:       "",
				Tiles: []tmx.TileData{
					tmx.TileData{},
					tmx.TileData{},
					tmx.TileData{},
					tmx.TileData{},
				},
			},
		},
	},
	"Adventure": DecoderTest{
		expected: []uint32{1, 2, 1, 1, 1, 2, 1, 2, 1, 1, 2, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 13, 14, 13, 1, 13, 14, 13, 14, 13, 1, 2, 14, 13, 1, 1, 2, 13, 14, 13, 14, 13, 14, 13, 14, 13, 1, 2, 1, 1, 2, 2, 1, 2, 1, 1, 2, 2, 1, 13, 13, 14, 2, 56, 1, 2, 1, 2, 1, 2, 1, 1, 2, 13, 13, 14, 14, 1, 14, 13, 1, 2, 14, 2, 1, 1, 1, 2, 1, 2, 2, 13, 14, 13, 14, 13, 1, 1, 13, 1, 1, 1, 1, 14, 1, 2, 14, 14, 14, 13, 13, 13, 14, 13, 14, 56, 1, 2, 1, 2, 1, 13, 13, 14, 1, 1, 1, 1, 14, 1, 2, 13, 14, 13, 13, 1, 2, 1, 1, 1, 1, 13, 14, 13, 14, 13, 1, 13, 14, 1, 1, 1, 1, 2, 1, 1, 2, 14, 1, 1, 13, 1, 55, 56, 1, 1, 1, 2, 1, 2, 1, 13, 14, 1, 1, 1, 1, 1, 14, 13, 1, 1, 1, 1, 2, 1, 2, 1, 2, 55, 56, 13, 14, 13, 14, 13, 1, 1, 2, 1, 1, 1, 1, 1, 1, 13, 13, 13, 13, 1, 1, 1, 1, 1, 2, 2, 13, 2, 1, 2, 1, 13, 13, 14, 13, 13, 13, 13, 13, 13, 1, 13, 14, 2, 1, 2, 2, 1, 2, 14, 14, 13, 14, 13, 14, 13, 1, 2, 14, 2, 1, 2, 1, 2, 1, 1, 2, 14, 1, 13, 14, 2, 2, 14, 55, 56, 1, 2, 1, 2, 1, 1, 2, 1, 2, 2, 2, 13, 14, 13, 1, 2, 1, 13, 13, 14, 14, 14, 2, 1, 2, 13, 14, 13, 14, 13, 1, 2, 1, 2, 14, 14, 13, 14, 13, 1, 2, 13, 13, 14, 13, 14, 14, 13, 14, 13, 14, 2, 1, 2, 1, 1, 1, 1, 2, 2, 2, 2, 14, 13, 1, 2, 13, 14, 14, 13, 1, 2, 1, 2, 1, 13, 14, 13, 14, 13, 1, 1, 13, 14, 14, 14, 14, 14, 1, 1, 1, 13, 14, 2, 1, 2, 55, 56, 14, 13, 14, 2, 1, 2, 1, 13, 1, 1, 13, 14, 13, 14, 2, 1, 2, 13, 14, 13, 14, 13, 14, 67, 1, 2, 1, 13, 14, 13, 14, 13, 1, 13, 13, 13, 14, 13, 14, 14, 1, 1, 13, 14, 1, 2, 1, 2, 55, 56, 14, 13, 14, 2, 1, 2, 1, 13, 13, 13, 13, 13, 13, 13, 14, 1, 13, 13, 14, 13, 14, 13, 14, 67, 1, 2, 1, 13, 14, 13, 14, 13, 1, 2, 1, 2, 1, 2, 1, 2, 13, 13, 14, 2, 1, 2, 1, 2, 1, 2, 67, 1, 2, 2, 1, 2, 1, 13, 14, 13, 14, 13, 14, 13, 14, 13, 14, 13, 14, 13, 14, 13, 14, 13, 14, 13, 13, 14, 14, 13, 14, 13, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 13, 14, 13, 14, 13, 14, 13, 14, 13, 14, 13, 14, 13, 14, 13, 14, 13, 14, 13, 14, 13, 14, 13, 14, 13, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 13, 14, 13, 14, 13, 14, 13, 14, 13, 14, 13, 14, 13, 14, 13, 14, 13, 14, 13, 14, 13, 14, 13, 14, 13, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1},
		data: []tmx.Data{
			tmx.Data{
				Encoding:    "base64",
				Compression: "zlib",
				Inner:       " eJztVcsOwzAIS3OH39nu+/9/Wic1kuXYJNp63AHRIoIxj+RorfVTDpBOenyzfVfHKXlp/kd7p/9hU/7ufJUv2wNifGzPBQ/MZ5zLy654VHV1vIZv0NkDsDpgj3Mh4jk+6FthBMXnWmCuiofD4Hpmm3k/IH83ly4+96OK0QGr6ofijTmjKB/sedUPjoUxcVbVjOM8VPu1mu8UeMOOfXHnMTeeI8c725zbas8dZ7WnYXxVPVTfsAaMkcLu7j+15ylE7Zea2YqHwlG+7l59bfJw9w/vdba5Prs8nCT5/MKjesf4rXB+iLUzB99qN88rHnfoO3nsvOd/HrV+A9pgE7E=",
			},
		},
	},
}

// TestLayerDecode tests the Later Data Decoder by feeding it encoded data
// and checking for expected output.
func TestLayerDecode(t *testing.T) {
	for name, dt := range decoderTests {
		for _, data := range dt.data {
			// Give each test a name
			var encoding string
			var compression string

			if len(data.Encoding) < 1 {
				encoding = "none"
			} else {
				encoding = data.Encoding
			}

			if len(data.Compression) < 1 {
				compression = "none"
			} else {
				compression = data.Compression
			}
			tiles := len(data.Tiles)
			var tname string
			tname = fmt.Sprintf("%v/%v-%v-%v", name, encoding, compression, tiles)
			t.Run(tname, func(t *testing.T) {

			})
		}
	}
}
