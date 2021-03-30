package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/neverlock/openvidu-exporter/session"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
)

var (
	SESSION_API_ENDPOINT = "/openvidu/api/sessions"

	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
)

type ViduExporter struct {
	Host                         string
	User                         string
	Pass                         string
	ViduSessionAPIDesc           *prometheus.Desc //Session or Connection
	ViduConnectionPerSessionDesc *prometheus.Desc
}

func (v *ViduExporter) ReallyExpensiveAssessmentOfTheSystemState() (objectByHost map[string]int, sessionByHost map[string]int) {
	session, connection, connectionInsession, err := getSessionNumber(v.Host, v.User, v.Pass)
	if err != nil {
		return
	}
	objectByHost = map[string]int{
		"session":    session,
		"connection": connection,
	}
	sessionByHost = connectionInsession
	return
}

func getSessionNumber(host string, user string, pass string) (int, int, map[string]int, error) {
	req, err := http.NewRequest("GET", host+SESSION_API_ENDPOINT, nil)
	if err != nil {
		return 0, 0, nil, err
	}

	req.SetBasicAuth(user, pass)

	resp, err := client.Do(req)
	if err != nil {
		return 0, 0, nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return 0, 0, nil, err
	}
	sessionCount := session.GetSessionActive(body)

	//collect connection from each sessionID in to map[sessionID]connection
	connectionInsession := make(map[string]int)
	for i := 0; i < sessionCount; i++ {
		connectionInsession[session.GetSessionID(body, i)] = session.GetConnectionInSession(body, i)
	}
	return sessionCount, session.GetConnectionActive(body), connectionInsession, nil
}

func (v *ViduExporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- v.ViduSessionAPIDesc
	ch <- v.ViduConnectionPerSessionDesc
}

func (v *ViduExporter) Collect(ch chan<- prometheus.Metric) {
	objectByHost, sessionByHost := v.ReallyExpensiveAssessmentOfTheSystemState()
	for object, scount := range objectByHost {
		ch <- prometheus.MustNewConstMetric(
			v.ViduSessionAPIDesc,
			prometheus.GaugeValue,
			float64(scount),
			object,
		)
	}
	for sessionid, ccount := range sessionByHost {
		ch <- prometheus.MustNewConstMetric(
			v.ViduConnectionPerSessionDesc,
			prometheus.GaugeValue,
			float64(ccount),
			sessionid,
		)
	}
}

func NewExporter(viduHost string, viduUsername string, viduPassword string) *ViduExporter {
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
		ViduConnectionPerSessionDesc: prometheus.NewDesc(
			"vidu_connection_per_session",
			"Number of connection in each session in  OpenVidu server",
			[]string{"sessionID"},
			prometheus.Labels{"host": viduHost},
		),
	}

}

func readConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s \n", err))
	}
}

func main() {
	readConfig()

	exporterListen := viper.GetString("Vidu.ListenHost") + viper.GetString("Vidu.ListenPort")

	reg := prometheus.NewPedanticRegistry()

	for i := 1; i <= viper.GetInt("Vidu.MaxSub"); i++ {
		sub := fmt.Sprintf("Vidu.Vidu%d", i)
		Vidu := viper.Sub(sub)
		if Vidu == nil {
			panic("Vidu config not found")
		}
		exporter := NewExporter(Vidu.GetString("URL"), Vidu.GetString("User"), Vidu.GetString("Password"))
		reg.MustRegister(exporter)
	}

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
