package metrics

import (
    "github.com/prometheus/client_golang/prometheus"
)

var (
    CPUTemperature = prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "raspberry_pi_cpu_temperature_celsius",
        Help: "Current CPU temperature of Raspberry Pi.",
    })
    
    InternetSpeed = prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "raspberry_pi_internet_speed_mbps",
        Help: "Current internet speed in Mbps.",
    })

    CameraStatus = prometheus.NewGaugeVec(prometheus.GaugeOpts{
        Name: "raspberry_pi_camera_status",
        Help: "Camera status: 1 for active, 0 for inactive.",
    }, []string{"camera_id"})
)

func init() {
    // Register metrics
    prometheus.MustRegister(CPUTemperature)
    prometheus.MustRegister(InternetSpeed)
    prometheus.MustRegister(CameraStatus)
}
