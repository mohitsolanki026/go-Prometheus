package metrics

import "github.com/mohitsolanki026/monitoring-Prometheus/pkg/collectors"

func UpdateMetrics() {
    CPUTemperature.Set(collectors.GetCPUTemperature())
    InternetSpeed.Set(collectors.GetInternetSpeed())
    CameraStatus.Set(collectors.GetCameraStatus())
}
