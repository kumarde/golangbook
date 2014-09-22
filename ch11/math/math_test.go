package math

import "testing"

type testpair struct{
    values[] float64
    average float64
    sum float64
}

var tests = []testpair{
    { []float64{1, 2}, 1.5, 3},
    { []float64{1,2,3,4,5,6}, 21.0/6.0, 21},
    { []float64{0}, 0, 0},
    { []float64{1,1,1,1}, 1, 4},
}

func TestAverage(t *testing.T){
    for _, pair := range tests{
        actual := Average(pair.values)
        expected := pair.average
        if expected != actual {
            t.Error(
                "For", pair.values,
                "expected", expected,
                "got", actual,
            )
        }
    }
}

func TestSum(t *testing.T){
    for _, pair := range tests{
        actual := Sum(pair.values)
        expected := pair.sum
        if expected != actual {
            t.Error(
                "For", pair.values,
                "expected", expected,
                "got", actual,
            )
        }
    }
}
