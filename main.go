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

func readConfig() {
	viper.SetConfigName("config") // ชื่อ config file
	viper.AddConfigPath(".")      // ระบุ path ของ config file
	viper.AutomaticEnv()          // อ่าน value จาก ENV variable
	// แปลง _ underscore ใน env เป็น . dot notation ใน viper
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// อ่าน config
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
