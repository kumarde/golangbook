package math

//Generates an average from a slice of float64
func Average(xs []float64) float64{
    total := float64(0)
    for _, value := range xs {
        total += value
    }
    return total / float64(len(xs))
}

func Sum(xs []float64) float64{
    total := float64(0)
    for _, value := range xs {
        total += value
    }
    return total
}
