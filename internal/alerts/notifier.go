package alert

import "fmt"

func TriggerAlert(currentTemp, threshold float64) {
    fmt.Printf("Alert: Temperature %.2f°C exceeds threshold of %.2f°C\n", currentTemp, threshold)
    // Further implementations: send email or push notification
}
