package collectors

import (
    "math/rand"
)

func GetInternetSpeed() float64 {
    return 20.0 + rand.Float64()*5 // Random value for simulation
}
