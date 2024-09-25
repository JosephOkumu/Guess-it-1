package guessit

import (
    "testing"
)

func TestAverage(t *testing.T) {
    type args struct {
        numslice []float64
    }
    tests := []struct {
        name string
        args args
        want float64
    }{
        {
            name: "A number",
            args: args{numslice: []float64{10.0}},
            want: 10,
        },
        {
            name: "Multiple numbers",
            args: args{numslice: []float64{1.0, 2.0, 3.0, 4.0, 5.0}},
            want: 3.0,
        },
        {
            name: "positive and negative",
            args: args{numslice: []float64{-1.0, -2.0, 3.0, 4.0}},
            want: 1.0,
        },
        {
            name: "Negative numbers",
            args: args{numslice: []float64{-1.0, -2.0, -3.0, -4.0, -5.0}},
            want: -3.0,
        },
        {
            name: "Similar numbers",
            args: args{numslice: []float64{2.0, 2.0, 2.0, 2.0, 2.0}},
            want: 2.0,
        },
        {
            name: "Floats",
            args: args{numslice: []float64{1.5, 2.5, 3.5, 4.5}},
            want: 3.0,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Average(tt.args.numslice); got != tt.want {
                t.Errorf("Mean() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestPredictRange(t *testing.T) {
    type args struct {
        numslice []float64
    }
    tests := []struct {
        name  string
        args  args
        want1 int
        want2 int
    }{
        {
            name:  "Single float",
            args:  args{numslice: []float64{10.0}},
            want1: 10,
            want2: 10,
        },
        {
            name:  "Multiple floats with normal distribution",
            args:  args{numslice: []float64{1.0, 2.0, 3.0, 4.0, 5.0}},
            want1: 1,
            want2: 6,
        },
        {
            name:  "positive and negative numbers",
            args:  args{numslice: []float64{-1.0, -2.0, 3.0, 4.0}},
            want1: -4,
            want2: 6,
        },
        {
            name:  "Negative numbers",
            args:  args{numslice: []float64{-1.0, -2.0, -3.0, -4.0, -5.0}},
            want1: -6,
            want2: -1,
        },
        {
            name:  "Similar numbers",
            args:  args{numslice: []float64{2.0, 2.0, 2.0, 2.0, 2.0}},
            want1: 2,
            want2: 2,
        },
        {
            name:  "Floats",
            args:  args{numslice: []float64{1.5, 2.5, 3.5, 4.5}},
            want1: 1,
            want2: 5,
        },
        {
            name:  "Large range of values",
            args:  args{numslice: []float64{-100.0, 0.0, 50.0, 100.0}},
            want1: -121,
            want2: 146,
        },
        {
            name:  "Small range of values",
            args:  args{numslice: []float64{1.0, 2.0, 2.5, 3.0}},
            want1: 1,
            want2: 3,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, got1 := PredictRange(tt.args.numslice)
            if got != tt.want1 {
                t.Errorf("Range() got = %v, want %v", got, tt.want1)
            }
            if got1 != tt.want2 {
                t.Errorf("Range() got1 = %v, want %v", got1, tt.want2)
            }
        })
    }
}

func TestVariance(t *testing.T) {
    type args struct {
        nums []float64
    }
    tests := []struct {
        name string
        args args
        want float64
    }{
        {
            name: "Single element",
            args: args{nums: []float64{10.0}},
            want: 0,
        },
        {
            name: "Multiple elements",
            args: args{nums: []float64{1.0, 2.0, 3.0, 4.0, 5.0}},
            want: 2.0,
        },
        {
            name: "All same numbers",
            args: args{nums: []float64{2.0, 2.0, 2.0, 2.0}},
            want: 0,
        },
        {
            name: "Mixed values",
            args: args{nums: []float64{-1.0, 0.0, 1.0}},
            want: 0.6666666666666666,
        },
        {
            name: "Large range of values",
            args: args{nums: []float64{-100.0, 0.0, 50.0, 100.0}},
            want: 5468.75,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Variance(tt.args.nums); got != tt.want {
                t.Errorf("Variance() = %v, want %v", got, tt.want)
            }
        })
    }
}

func TestStddev(t *testing.T) {
    type args struct {
        variance float64
    }
    tests := []struct {
        name string
        args args
        want float64
    }{
        {
            name: "Zero variance",
            args: args{variance: 0},
            want: 0,
        },
        {
            name: "Positive variance",
            args: args{variance: 4.0},
            want: 2.0,
        },
        {
            name: "Large variance",
            args: args{variance: 24850.0},
            want: 157.63882770434446,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := StdDev(tt.args.variance); got != tt.want {
                t.Errorf("Stddev() = %v, want %v", got, tt.want)
            }
        })
    }
}