package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/neverlock/openvidu-exporter/session"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

var (
	SESSION_API_ENDPOINT = "/openvidu/api/sessions"

	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
)

type ViduExporter struct {
	Host               string
	User               string
	Pass               string
	ViduSessionAPIDesc *prometheus.Desc //Session or Connection
}

func (v *ViduExporter) ReallyExpensiveAssessmentOfTheSystemState() (objectByHost map[string]int) {
	session, connection, err := getSessionNumber(v.Host, v.User, v.Pass)
	if err != nil {
		return
	}
	objectByHost = map[string]int{
		"session":    session,
		"connection": connection,
	}
	return
}

func getSessionNumber(host string, user string, pass string) (int, int, error) {
	req, err := http.NewRequest("GET", host+SESSION_API_ENDPOINT, nil)
	if err != nil {
		return 0, 0, err
	}

	req.SetBasicAuth(user, pass)

	resp, err := client.Do(req)
	if err != nil {
		return 0, 0, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return 0, 0, err
	}
	return session.GetSessionActive(body), session.GetConnectionActive(body), nil
}

func (v *ViduExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- v.ViduSessionAPIDesc
}

func (v *ViduExporter) Collect(ch chan<- prometheus.Metric) {
	objectByHost := v.ReallyExpensiveAssessmentOfTheSystemState()
	for object, scount := range objectByHost {
		ch <- prometheus.MustNewConstMetric(
			v.ViduSessionAPIDesc,
			prometheus.GaugeValue,
			float64(scount),
			object,
		)
	}
}

func NewExporter(viduHost string, viduUsername string, viduPassword string) *ViduExporter {
	//test
	fmt.Println(viduUsername, viduPassword)
	//endtest
	return &ViduExporter{
		Host: viduHost,
		User: viduUsername,
		Pass: viduPassword,
		ViduSessionAPIDesc: prometheus.NewDesc(
			"vidu_session_usage",
			"Number of session use in OpenVidu server",
			[]string{"object"},
			prometheus.Labels{"host": viduHost},
		),
	}

}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Errorf("Error loading .env file, assume env variables are set.")
	}

	flag.Parse()

	viduHost := os.Getenv("VIDU_HOST")
	viduUsername := os.Getenv("VIDU_USERNAME")
	viduPassword := os.Getenv("VIDU_PASSWORD")
	exporterListen := os.Getenv("LISTEN_HOST") + os.Getenv("LISTEN_PORT")

	exporter := NewExporter(viduHost, viduUsername, viduPassword)
	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(exporter)

	gatherers := prometheus.Gatherers{
		prometheus.DefaultGatherer,
		reg,
	}
	h := promhttp.HandlerFor(gatherers,
		promhttp.HandlerOpts{
			ErrorLog:      log.NewErrorLogger(),
			ErrorHandling: promhttp.ContinueOnError,
		})
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})
	log.Infoln("Start server at ", exporterListen)
	if err := http.ListenAndServe(exporterListen, nil); err != nil {
		log.Errorf("Error occur when start server %v", err)
		os.Exit(1)
	}

}
