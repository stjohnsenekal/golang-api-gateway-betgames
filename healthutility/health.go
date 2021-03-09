package healthutility

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

type Health struct {
	ServiceName         string
	Status              int16
	MessageCount        uint64
	ErrorCount          uint64
	QueryCount          uint64
	LastQueryRun        int64
	LastMessageReceived int64
	LastMessageProduced int64
	ServiceStartTime    int64
}

const (
	SERVICE_STATUS_UP                   = 1
	SERVICE_STATUS_DOWN                 = 0
	HEALTH_WEB_SERVICE_STARTED          = "HealthMonitorStarted: (%d)"
	HEALTH_WEB_SERVICE_REQUEST_RECEIVED = "HealthMonitorRequest: (%s)"
)

type Monitor struct {
	Health *Health
}

func (monitor *Monitor) get(w http.ResponseWriter, r *http.Request) {
	log.Printf(HEALTH_WEB_SERVICE_REQUEST_RECEIVED, "Health endpoint hit")
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(monitor.Health); err != nil {
		log.Fatal("HealthMonitorFailed:", err)
	}
}

func (monitor *Monitor) Listen(port int) {
	http.HandleFunc("/health/", monitor.get)
	log.Printf(HEALTH_WEB_SERVICE_STARTED, port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatal("HealthMonitorFailed:", err)
	}
}

func (health *Health) SetServiceName(name string) {
	health.ServiceName = name
}

func (health *Health) IncMessageCount() {
	atomic.AddUint64(&health.MessageCount, 1)
}

func (health *Health) IncErrorCount() {
	atomic.AddUint64(&health.ErrorCount, 1)
}

func (health *Health) IncQueryCount() {
	atomic.AddUint64(&health.QueryCount, 1)
}

func (health *Health) MessageReceived() {
	health.IncMessageCount()
	health.LastMessageReceived = time.Now().Unix()
}

func (health *Health) QueryRan() {
	health.IncQueryCount()
	health.LastQueryRun = time.Now().Unix()
}

func (health *Health) MessageProduced() {
	health.LastMessageProduced = time.Now().Unix()
}

func (health *Health) ServiceStarted() {
	health.ServiceStartTime = time.Now().Unix()
}
