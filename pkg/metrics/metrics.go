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

    InternetStatus = prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "raspberry_pi_internet_status",
        Help: "Internet status: 1 for active, 0 for inactive.",
    })

    CameraStatus = prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "raspberry_pi_camera_status",
        Help: "Camera status: 1 for active, 0 for inactive.",
    })
    
    CameraStream = prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "raspberry_pi_camera_stream",
        Help: "Camera stream: 1 for active, 0 for inactive.",
    })

    DongalStatus = prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "raspberry_pi_dongal_status",
        Help: "Dongal status: 1 for active, 0 for inactive.",
    })

    CpuVoltage = prometheus.NewGauge(prometheus.GaugeOpts{
        Name: "raspberry_pi_cpu_voltage",
        Help: "Current CPU voltage of Raspberry Pi.",
    })
)

func init() {
    // Register metrics
    prometheus.MustRegister(CPUTemperature)
    prometheus.MustRegister(InternetSpeed)
    prometheus.MustRegister(CameraStatus)
    prometheus.MustRegister(CameraStream)
    prometheus.MustRegister(DongalStatus)
    prometheus.MustRegister(CpuVoltage)
    prometheus.MustRegister(InternetStatus)
}
