package collectors

import (
    "math/rand"
)

func GetCPUTemperature() float64 {
    return 50.0 + rand.Float64()*10 // Random value for simulation
}
