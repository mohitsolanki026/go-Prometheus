package metrics

import "github.com/mohitsolanki026/monitoring-Prometheus/pkg/collectors"

func UpdateMetrics() {
    CPUTemperature.Set(collectors.GetCPUTemperature())
    InternetSpeed.Set(collectors.GetInternetSpeed())
    CameraStatus.WithLabelValues("camera_1").Set(collectors.GetCameraStatus("camera_1"))
}
