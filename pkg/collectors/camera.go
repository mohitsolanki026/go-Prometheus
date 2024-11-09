package collectors

func GetCameraStatus(cameraID string) float64 {
    if cameraID == "camera_1" {
        return 1 // Active
    }
    return 0 // Inactive
}
