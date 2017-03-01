package config

import (
	"os"
	"strconv"
)

type Config struct {
	Target                   string
	PollIntervalInMs         int
	TotalNumberOfRequests    int
	ReportIntervalInRequests int
}

func Load() Config {
	var config Config
	config = Config{
		Target:                   "example.com",
		PollIntervalInMs:         1,
		TotalNumberOfRequests:    10,
		ReportIntervalInRequests: 5,
	}

	if targetString, ok := os.LookupEnv("TARGET"); ok {
		config.Target = targetString
	}

	if value, ok := os.LookupEnv("POLL_INTERVAL_IN_MS"); ok {
		if numValue, err := strconv.Atoi(value); err != nil {
			panic(err)
		} else {
			config.PollIntervalInMs = numValue
		}
	}

	//	if value, ok := os.LookupEnv("REPORT_INTERVAL_IN_SECONDS"); ok {
	//		if numValue, err := strconv.Atoi(value); err != nil {
	//			panic(err)
	//		} else {
	//			config.Report_interval_in_requests = numValue
	//		}
	//	}

	if value, ok := os.LookupEnv("TOTAL_NUMBER_OF_REQUESTS"); ok {
		if numValue, err := strconv.Atoi(value); err != nil {
			panic(err)
		} else {
			config.TotalNumberOfRequests = numValue
		}
	}

	if value, ok := os.LookupEnv("REPORT_INTERVAL_IN_REQUESTS"); ok {
		if numValue, err := strconv.Atoi(value); err != nil {
			panic(err)
		} else {
			config.ReportIntervalInRequests = numValue
		}
	}

	return config
}
