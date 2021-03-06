package tosca_v1_3

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/tliron/puccini/tosca"
)

//
// ScalarUnit
//

type ScalarUnit struct {
	CanonicalString string      `json:"$string" yaml:"$string"`
	CanonicalNumber interface{} `json:"$number" yaml:"$number"` // float64 or uint64
	OriginalString  string      `json:"$originalString" yaml:"$originalString"`

	Scalar float64 `json:"scalar" yaml:"scalar"`
	Unit   string  `json:"unit" yaml:"unit"`

	count                 bool // if true, CanonicalNumber is uint64
	canonicalUnitSingular string
	canonicalUnitPlural   string
}

func ReadScalarUnit(context *tosca.Context, name string, canonicalUnit string, canonicalUnitSingular string, canonicalUnitPlural string, re *regexp.Regexp, measures ScalarUnitMeasures, count bool, caseSensitive bool) *ScalarUnit {
	self := ScalarUnit{
		count:                 count,
		canonicalUnitSingular: canonicalUnitSingular,
		canonicalUnitPlural:   canonicalUnitPlural,
	}

	if !context.ValidateType("string") {
		return &self
	}

	// Original
	self.OriginalString = *context.ReadString()

	// Regular expression
	matches := re.FindStringSubmatch(self.OriginalString)
	if len(matches) != 3 {
		context.ReportValueMalformed(name, "")
		return &self
	}

	// Scalar
	var err error
	if self.Scalar, err = strconv.ParseFloat(matches[1], 64); err != nil {
		context.ReportValueMalformed(name, fmt.Sprintf("%s", err))
		return &self
	}

	// Unit
	var measure float64
	self.Unit, measure = measures.Get(matches[2], caseSensitive)

	// Canonical
	if count {
		self.CanonicalNumber = uint64(math.Round(self.Scalar * measure))
		self.CanonicalString = fmt.Sprintf("%d %s", self.CanonicalNumber, canonicalUnit)
	} else {
		self.CanonicalNumber = self.Scalar * measure
		self.CanonicalString = fmt.Sprintf("%g %s", self.CanonicalNumber, canonicalUnit)
	}

	return &self
}

// fmt.Stringify interface
func (self *ScalarUnit) String() string {
	var singular bool

	if self.canonicalUnitSingular == self.canonicalUnitPlural {
		singular = false
	} else if self.count {
		singular = self.CanonicalNumber.(uint64) == 1
	} else {
		singular = self.CanonicalNumber.(float64) == 1.0
	}

	if singular {
		return fmt.Sprintf("1 %s", self.canonicalUnitSingular)
	} else {
		if self.count {
			return fmt.Sprintf("%d %s", self.CanonicalNumber.(uint64), self.canonicalUnitPlural)
		} else {
			return fmt.Sprintf("%g %s", self.CanonicalNumber.(float64), self.canonicalUnitPlural)
		}
	}
}

func (self *ScalarUnit) Compare(data interface{}) (int, error) {
	if scalarUnit, ok := data.(*ScalarUnit); ok {
		if self.count {
			return CompareUint64(self.CanonicalNumber.(uint64), scalarUnit.CanonicalNumber.(uint64)), nil
		} else {
			return CompareFloat64(self.CanonicalNumber.(float64), scalarUnit.CanonicalNumber.(float64)), nil
		}
	} else {
		return 0, errors.New("incompatible comparison")
	}
}

//
// ScalarUnitMeasures
//

type ScalarUnitMeasures map[string]float64

func (self ScalarUnitMeasures) Get(unit string, caseSensitive bool) (string, float64) {
	if caseSensitive {
		if measure, ok := self[unit]; ok {
			return unit, measure
		}
	} else {
		for canonicalUnit, measure := range self {
			if strings.EqualFold(canonicalUnit, unit) {
				return canonicalUnit, measure
			}
		}
	}

	panic("as long as the regexp does its job we should never get here")
}

//
// ScalarUnitSize
//
// [TOSCA-Simple-Profile-YAML-v1.3] @ 3.3.6.4
// [TOSCA-Simple-Profile-YAML-v1.2] @ 3.3.6.4
// [TOSCA-Simple-Profile-YAML-v1.1] @ 3.2.6.4
//

var ScalarUnitSizeRE = regexp.MustCompile(
	`^(?P<scalar>[0-9]*\.?[0-9]+(?:e[-+]?[0-9]+)?)\s*` +
		`(?i)(?P<unit>B|kB|KiB|MB|MiB|GB|GiB|TB|TiB)$`)

var ScalarUnitSizeMeasures = ScalarUnitMeasures{
	"B":   1,
	"kB":  1000,
	"KiB": 1024,
	"MB":  1000000,
	"MiB": 1048576,
	"GB":  1000000000,
	"GiB": 1073741824,
	"TB":  1000000000000,
	"TiB": 1099511627776,
}

// tosca.Reader signature
func ReadScalarUnitSize(context *tosca.Context) interface{} {
	return ReadScalarUnit(context, "scalar-unit.size", "B", "byte", "bytes", ScalarUnitSizeRE, ScalarUnitSizeMeasures, true, false)
}

//
// ScalarUnitTime
//
// [TOSCA-Simple-Profile-YAML-v1.3] @ 3.3.6.5
// [TOSCA-Simple-Profile-YAML-v1.2] @ 3.3.6.5
// [TOSCA-Simple-Profile-YAML-v1.1] @ 3.2.6.5
//

var ScalarUnitTimeRE = regexp.MustCompile(
	`^(?P<scalar>[0-9]*\.?[0-9]+(?:e[-+]?[0-9]+)?)\s*` +
		`(?i)(?P<unit>ns|us|ms|s|m|h|d)$`)

var ScalarUnitTimeMeasures = ScalarUnitMeasures{
	"ns": 0.000000001,
	"us": 0.000001,
	"ms": 0.001,
	"s":  1,
	"m":  60,
	"h":  3600,
	"d":  86400,
}

// tosca.Reader signature
func ReadScalarUnitTime(context *tosca.Context) interface{} {
	return ReadScalarUnit(context, "scalar-unit.time", "s", "second", "seconds", ScalarUnitTimeRE, ScalarUnitTimeMeasures, false, false)
}

//
// ScalarUnitFrequency
//
// [TOSCA-Simple-Profile-YAML-v1.3] @ 3.3.6.6
// [TOSCA-Simple-Profile-YAML-v1.2] @ 3.3.6.6
// [TOSCA-Simple-Profile-YAML-v1.1] @ 3.2.6.6
//

var ScalarUnitFrequencyRE = regexp.MustCompile(
	`^(?P<scalar>[0-9]*\.?[0-9]+(?:e[-+]?[0-9]+)?)\s*` +
		`(?i)(?P<unit>Hz|kHz|MHz|GHz)$`)

var ScalarUnitFrequencyMeasures = ScalarUnitMeasures{
	"Hz":  1,
	"kHz": 1000,
	"MHz": 1000000,
	"GHz": 1000000000,
}

// tosca.Reader signature
func ReadScalarUnitFrequency(context *tosca.Context) interface{} {
	return ReadScalarUnit(context, "scalar-unit.frequency", "Hz", "Hz", "Hz", ScalarUnitFrequencyRE, ScalarUnitFrequencyMeasures, false, false)
}

//
// ScalarUnitBitrate
//
// [TOSCA-Simple-Profile-YAML-v1.3] @ 3.3.6.7
//

var ScalarUnitBitrateRE = regexp.MustCompile(
	`^(?P<scalar>[0-9]*\.?[0-9]+(?:e[-+]?[0-9]+)?)\s*` +
		`(?P<unit>bps|Kbps|Kibps|Mbps|Mibps|Gbps|Gibps|Tbps|Tibps|Bps|KBps|KiBps|MBps|MiBps|GBps|GiBps|TBps|TiBps)$`)

// Case-sensitive!
var ScalarUnitBitrateMeasures = ScalarUnitMeasures{
	"bps":   1,
	"Kbps":  1000,
	"Kibps": 1024,
	"Mbps":  1000000,
	"Mibps": 1048576,
	"Gbps":  1000000000,
	"Gibps": 1073741824,
	"Tbps":  1000000000000,
	"Tibps": 1099511627776,
	"Bps":   8,
	"KBps":  8000,
	"KiBps": 8192,
	"MBps":  8000000,
	"MiBps": 8388608,
	"GBps":  8000000000,
	"GiBps": 8589934592,
	"TBps":  8000000000000,
	"TiBps": 8796093022208,
}

// tosca.Reader signature
func ReadScalarUnitBitrate(context *tosca.Context) interface{} {
	return ReadScalarUnit(context, "scalar-unit.bitrate", "bps", "bps", "bps", ScalarUnitBitrateRE, ScalarUnitBitrateMeasures, false, true)
}
