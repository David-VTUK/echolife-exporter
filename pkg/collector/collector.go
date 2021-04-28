package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Stats get reset after disconnect, hence the use of Guage Type
type myMetrics struct {
	UpstreamCurrRate      prometheus.Gauge
	DownstreamCurrRate    prometheus.Gauge
	UpstreamCurrRate2     prometheus.Gauge
	DownstreamCurrRate2   prometheus.Gauge
	UpstreamMaxRate       prometheus.Gauge
	DownstreamMaxRate     prometheus.Gauge
	UpstreamNoiseMargin   prometheus.Gauge
	DownstreamNoiseMargin prometheus.Gauge
	UpstreamAttenuation   prometheus.Gauge
	DownstreamAttenuation prometheus.Gauge
	UpstreamPower         prometheus.Gauge
	DownstreamPower       prometheus.Gauge
	DownstreamHECErrors   prometheus.Gauge
	UpstreamHECErrors     prometheus.Gauge
	DownstreamCRCErrors   prometheus.Gauge
	UpstreamCRCErrors     prometheus.Gauge
	DownstreamFECErrors   prometheus.Gauge
	UpstreamFECErrors     prometheus.Gauge
	DownstreamHECErrors2  prometheus.Gauge
	UpstreamHECErrors2    prometheus.Gauge
	DownstreamCRCErrors2  prometheus.Gauge
	UpstreamCRCErrors2    prometheus.Gauge
	DownstreamFECErrors2  prometheus.Gauge
	UpstreamFECErrors2    prometheus.Gauge
}

func new() myMetrics {
	// Init and return a instance of MyMetrics
	mm := myMetrics{
		UpstreamCurrRate: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "upstream_current_rate",
				Help: "Current upstream rate",
			},
		),
		DownstreamCurrRate: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "downstream_current_rate",
				Help: "Current downstream rate",
			},
		),
		UpstreamCurrRate2: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "upstream2_current_rate",
				Help: "Current upstream2 rate",
			},
		),
		DownstreamCurrRate2: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "downstream2_current_rate",
				Help: "Current downstream2 rate",
			},
		),
		UpstreamMaxRate: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "upstream_max_rate",
				Help: "Current maximum (attainable) upstream rate",
			},
		),
		DownstreamMaxRate: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "downstream_max_rate",
				Help: "Current maximum (attainable) downstream rate",
			},
		),
		UpstreamNoiseMargin: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "upstream_noise_margin",
				Help: "Current upstream noise margin",
			},
		),
		DownstreamNoiseMargin: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "downstream_noise_margin",
				Help: "Current downstream noise margin",
			},
		),
		UpstreamAttenuation: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "upstream_attenuation",
				Help: "Current upstream attenuation",
			},
		),
		DownstreamAttenuation: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "downstream_attenuation",
				Help: "Current downstream attenuation",
			},
		),
		UpstreamPower: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "upstream_power",
				Help: "Current upstream power",
			},
		),
		DownstreamPower: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "downstream_power",
				Help: "Current downstream power",
			},
		),
		DownstreamHECErrors: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "downstream_hec_errors",
				Help: "Number of HEC downstream errors since last showtime",
			},
		),
		UpstreamHECErrors: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "upstream_hec_errors",
				Help: "Number of HEC upstream errors since last showtime",
			},
		),
		DownstreamCRCErrors: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "downstream_crc_errors",
				Help: "Number of CRC downstream errors since last showtime",
			},
		),
		UpstreamCRCErrors: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "upstream_crc_errors",
				Help: "Number of CRC upstream errors since last showtime",
			},
		),
		DownstreamFECErrors: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "downstream_fec_errors",
				Help: "Number of FEC downstream errors since last showtime",
			},
		),
		UpstreamFECErrors: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "upstream_fec_errors",
				Help: "Number of FEC upstream errors since last showtime",
			},
		),
		DownstreamHECErrors2: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "downstream_hec_errors2",
				Help: "Number of HEC downstream errors since last showtime",
			},
		),
		UpstreamHECErrors2: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "upstream_hec_errors2",
				Help: "Number of HEC upstream errors since last showtime",
			},
		),
		DownstreamCRCErrors2: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "downstream_crc_errors2",
				Help: "Number of CRC downstream errors since last showtime",
			},
		),
		UpstreamCRCErrors2: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "upstream_crc_errors2",
				Help: "Number of CRC upstream errors since last showtime",
			},
		),
		DownstreamFECErrors2: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "downstream_fec_errors2",
				Help: "Number of FEC downstream errors since last showtime",
			},
		),
		UpstreamFECErrors2: prometheus.NewGauge(
			prometheus.GaugeOpts{
				Name: "upstream_fec_errors2",
				Help: "Number of FEC upstream errors since last showtime",
			},
		),
	}

	//Register metrics with prometheus
	prometheus.MustRegister(mm.UpstreamCurrRate)
	prometheus.MustRegister(mm.DownstreamCurrRate)
	prometheus.MustRegister(mm.UpstreamCurrRate2)
	prometheus.MustRegister(mm.DownstreamCurrRate2)
	prometheus.MustRegister(mm.UpstreamMaxRate)
	prometheus.MustRegister(mm.DownstreamMaxRate)
	prometheus.MustRegister(mm.UpstreamNoiseMargin)
	prometheus.MustRegister(mm.DownstreamNoiseMargin)
	prometheus.MustRegister(mm.UpstreamAttenuation)
	prometheus.MustRegister(mm.DownstreamAttenuation)
	prometheus.MustRegister(mm.UpstreamPower)
	prometheus.MustRegister(mm.DownstreamPower)
	prometheus.MustRegister(mm.DownstreamHECErrors)
	prometheus.MustRegister(mm.UpstreamHECErrors)
	prometheus.MustRegister(mm.DownstreamCRCErrors)
	prometheus.MustRegister(mm.UpstreamCRCErrors)
	prometheus.MustRegister(mm.DownstreamFECErrors)
	prometheus.MustRegister(mm.UpstreamFECErrors)
	prometheus.MustRegister(mm.DownstreamHECErrors2)
	prometheus.MustRegister(mm.UpstreamHECErrors2)
	prometheus.MustRegister(mm.DownstreamCRCErrors2)
	prometheus.MustRegister(mm.UpstreamCRCErrors2)
	prometheus.MustRegister(mm.DownstreamFECErrors2)
	return mm
}

func Collect() {
	url := "http://192.168.1.3/html/status/xdslStatus.asp"
	modemMetrics := new()

	// Update values every 30s (simply increment them by 2)
	ticker := time.NewTicker(30 * time.Second)
	for range ticker.C {
		//log.Info("Updating metric values")

		var client http.Client
		resp, err := client.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			//Form Regex to extract all quoted strings
			re := regexp.MustCompile("\"(.*?)\"")
			extractedValues := re.FindAll(bodyBytes, -1)

			modemMetrics.UpstreamCurrRate.Set(convertToFloat(extractedValues[4]))
			modemMetrics.DownstreamCurrRate.Set(convertToFloat(extractedValues[5]))
			modemMetrics.UpstreamCurrRate2.Set(convertToFloat(extractedValues[6]))
			modemMetrics.DownstreamCurrRate2.Set(convertToFloat(extractedValues[7]))
			modemMetrics.UpstreamMaxRate.Set(convertToFloat(extractedValues[8]))
			modemMetrics.DownstreamMaxRate.Set(convertToFloat(extractedValues[9]))
			modemMetrics.UpstreamNoiseMargin.Set(convertToFloat(extractedValues[10]))
			modemMetrics.DownstreamNoiseMargin.Set(convertToFloat(extractedValues[11]))
			modemMetrics.UpstreamAttenuation.Set(convertToFloat(extractedValues[12]))
			modemMetrics.DownstreamAttenuation.Set(convertToFloat(extractedValues[13]))
			modemMetrics.UpstreamPower.Set(convertToFloat(extractedValues[14]))
			modemMetrics.DownstreamPower.Set(convertToFloat(extractedValues[15]))
			modemMetrics.DownstreamHECErrors.Set(convertToFloat(extractedValues[18]))
			modemMetrics.UpstreamHECErrors.Set(convertToFloat(extractedValues[19]))
			modemMetrics.DownstreamCRCErrors.Set(convertToFloat(extractedValues[20]))
			modemMetrics.UpstreamCRCErrors.Set(convertToFloat(extractedValues[21]))
			modemMetrics.DownstreamFECErrors.Set(convertToFloat(extractedValues[22]))
			modemMetrics.UpstreamFECErrors.Set(convertToFloat(extractedValues[23]))
			modemMetrics.DownstreamHECErrors2.Set(convertToFloat(extractedValues[24]))
			modemMetrics.UpstreamHECErrors2.Set(convertToFloat(extractedValues[25]))
			modemMetrics.DownstreamCRCErrors2.Set(convertToFloat(extractedValues[26]))
			modemMetrics.UpstreamCRCErrors2.Set(convertToFloat(extractedValues[27]))
			modemMetrics.DownstreamFECErrors2.Set(convertToFloat(extractedValues[28]))
			modemMetrics.UpstreamFECErrors2.Set(convertToFloat(extractedValues[29]))

		}
		resp.Body.Close()
	}
}

func convertToFloat(input []byte) float64 {
	floatvalue, err := strconv.ParseFloat(strings.Trim(string(input), `"`), 64)

	if err != nil {
		log.Fatal(err)
	}
	return floatvalue
}
