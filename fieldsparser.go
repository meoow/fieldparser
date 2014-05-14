package fieldparser

import "math"
import "strings"
import "errors"
import "strconv"
import "log"
import "os"
import "sort"
import . "github.com/meoow/mergeoverlap"

//import "fmt"

const (
	BOL       int = 1
	EOL       int = math.MaxInt32
	NOT_RANGE int = -1
)

type (
	FieldRange struct {
		StartPos int
		EndPos   int
	}
	FieldsList []*FieldRange
)

var logger = log.New(os.Stderr, "", 0)

func ParseFields(fstr string) (FieldsList, error) {
	ranges := strings.Split(fstr, ",")
	if len(ranges) == 0 {
		return nil, errors.New("can not parse fields list")
	}
	tempRanges := make([][2]int, 0, len(ranges))
	tempPoints := make([]int, 0, len(ranges))
	var s, e int
	for _, rng := range ranges {
		ste := strings.SplitN(rng, "-", 2)
		switch len(ste) {
		case 0:
			continue
		case 1:
			//if ste[0]=="" {continue}
			s = atoi_positive(ste[0])
			e = NOT_RANGE
		case 2:
			switch ste[0] {
			case "":
				s = BOL
			default:
				s = atoi_positive(ste[0])
			}
			switch ste[1] {
			case "":
				e = EOL
			default:
				e = atoi_positive(ste[1])
			}
			if s >= e {
				e = NOT_RANGE
			}
		}
		if e == NOT_RANGE {
			tempPoints = append(tempPoints, s)
		} else {
			tempRanges = append(tempRanges, [2]int{s, e})
		}
	}
	if len(tempRanges) > 0 {
		tempRanges = MergeOverlap(tempRanges)
	}
	//fmt.Println(tempRanges)
	tempPoints2 := make([]int, 0, len(ranges))

	var found bool
	for _, j := range tempPoints {
		found = false
		for _, k := range tempRanges {
			if k[0] <= j && k[1] >= j {
				found = true
				break
			}
		}
		if !found {
			tempPoints2 = append(tempPoints2, j)
		}
	}
	outfields := make(FieldsList, 0, len(tempRanges)+len(tempPoints2))
	for _, j := range tempRanges {
		outfields = append(outfields, &FieldRange{j[0], j[1]})
	}
	for _, j := range tempPoints2 {
		outfields = append(outfields, &FieldRange{j, NOT_RANGE})
	}
	sort.Sort(outfields)
	return outfields, nil
}

func (fl FieldsList) Len() int {
	return len(fl)
}

func (fl FieldsList) Swap(i, j int) {
	fl[i], fl[j] = fl[j], fl[i]
}

func (fl FieldsList) Less(i, j int) bool {
	return fl[i].StartPos < fl[j].StartPos
}

func atoi_positive(n string) int {
	np, err := strconv.ParseUint(n, 10, 32)
	if err != nil {
		logger.Fatal(err)
	}
	if np > math.MaxInt32 {
		logger.Fatal(errors.New("field number is too big"))
	}
	return int(np)
}
