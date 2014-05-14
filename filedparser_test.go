package fieldparser

import "testing"
import "math"

func Test_ParseFields_1(t *testing.T) {
	var testfield = "1-"
	output, err := ParseFields(testfield)
	if err != nil {
		t.Error(err)
	}
	if output[0].StartPos == 1 && output[0].EndPos == math.MaxInt32 {
	} else {
		t.Logf("$#v\n",output)
		t.Error("parsing failed")
	}
}
func Test_ParseFields_2(t *testing.T) {
	var testfield = "10-20,11-50,-3"
	output, err := ParseFields(testfield)
	if err != nil {
		t.Error(err)
	}
	for i,j := range output {
		switch {
		case i==0 && j.StartPos == 1 && j.EndPos == 3:
		case i==1 && j.StartPos == 10 && j.EndPos == 50:
		default:
			t.Logf("%#v\n",j)
			t.Error("parsing failed")
		}
	}
}
func Test_ParseFields_3(t *testing.T) {
	var testfield = "-5,6-13,12-16,14-20,22-29,31-35,32-"
	output, err := ParseFields(testfield)
	if err != nil {
		t.Error(err)
	}
	for i,j := range output {
		switch {
		case i==0 && j.StartPos == 1 && j.EndPos == 5:
		case i==1 && j.StartPos == 6 && j.EndPos == 20:
		case i==2 && j.StartPos == 22 && j.EndPos == 29:
		case i==3 && j.StartPos == 31 && j.EndPos == math.MaxInt32:
		default:
			t.Logf("%#v\n",j)
			t.Error("parsing failed")
		}
	}
}
func Test_ParseFields_4(t *testing.T) {
	var testfield = "4,6,7,3"
	output, err := ParseFields(testfield)
	if err != nil {
		t.Error(err)
	}
	for i,j := range output {
		switch {
		case i==0 && j.StartPos == 3 && j.EndPos == -1:
		case i==1 && j.StartPos == 4 && j.EndPos == -1:
		case i==2 && j.StartPos == 6 && j.EndPos == -1:
		case i==3 && j.StartPos == 7 && j.EndPos == -1:
		default:
			t.Logf("%#v\n",j)
			t.Error("parsing failed")
		}
	}
}
func Test_ParseFields_5(t *testing.T) {
	var testfield = "-10,5,12,15-20,17-25,56,60-"
	output, err := ParseFields(testfield)
	if err != nil {
		t.Error(err)
	}
	for i,j := range output {
		switch {
		case i==0 && j.StartPos == 1 && j.EndPos == 10:
		case i==1 && j.StartPos == 12 && j.EndPos == -1:
		case i==2 && j.StartPos == 15 && j.EndPos == 25:
		case i==3 && j.StartPos == 56 && j.EndPos == -1:
		case i==4 && j.StartPos == 60 && j.EndPos == math.MaxInt32:
		default:
			t.Logf("%#v\n",j)
			t.Error("parsing failed")
		}
	}
}
