package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/KuaishouContainerService/quota-order-webhook/pkg/channel"
	"github.com/KuaishouContainerService/quota-order-webhook/pkg/routes"
	"github.com/comail/colog"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// Call Parse() to avoid noisy logs
	flag.CommandLine.Parse([]string{})

	colog.SetDefaultLevel(colog.LInfo)
	colog.SetMinLevel(colog.LInfo)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()
	level := StringToLevel(os.Getenv("LOG_LEVEL"))
	log.Print("Log level was set to ", strings.ToUpper(level.String()))
	colog.SetMinLevel(level)

	port := os.Getenv("PORT")
	if _, err := strconv.Atoi(port); err != nil {
		port = "39999"
	}

	channel.ApplyQuotaURL = os.Getenv("APPLY_QUOTA_URL")
	if channel.ApplyQuotaURL == "" {
		log.Fatal("APPLY_QUOTA_URL must be provided.")
	}
	log.Printf("debug: APPLY_QUOTA_URL is :%s", channel.ApplyQuotaURL)

	channel.ApplyQuotaToken = os.Getenv("APPLY_QUOTA_TOKEN")
	if channel.ApplyQuotaToken == "" {
		log.Fatal("APPLY_QUOTA_TOKEN must be provided.")
	}
	log.Printf("debug: APPLY_QUOTA_TOKEN is :%s", channel.ApplyQuotaToken)

	channel.TransferQuotaURL = os.Getenv("TRANSFER_QUOTA_URL")
	if channel.TransferQuotaURL == "" {
		log.Fatal("TRANSFER_QUOTA_URL must be provided.")
	}
	log.Printf("debug: TRANSFER_QUOTA_URL is :%s", channel.TransferQuotaURL)

	channel.TransferQuotaToken = os.Getenv("TRANSFER_QUOTA_TOKEN")
	if channel.TransferQuotaToken == "" {
		log.Fatal("TRANSFER_QUOTA_TOKEN must be provided.")
	}
	log.Printf("debug: TRANSFER_QUOTA_TOKEN is :%s", channel.TransferQuotaToken)

	onSubmit := channel.NewOnSubmit()
	onGetNextActs := channel.NewOnGetNextActs()

	router := httprouter.New()

	routes.AddPProf(router)
	routes.AddVersion(router)
	routes.AddOnSubmit(router, onSubmit)
	routes.AddOnGetNextActs(router, onGetNextActs)

	log.Printf("info: server starting on the port :%s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}

func StringToLevel(levelStr string) colog.Level {
	switch level := strings.ToUpper(levelStr); level {
	case "TRACE":
		return colog.LTrace
	case "DEBUG":
		return colog.LDebug
	case "INFO":
		return colog.LInfo
	case "WARNING":
		return colog.LWarning
	case "ERROR":
		return colog.LError
	case "ALERT":
		return colog.LAlert
	default:
		log.Printf("warning: LOG_LEVEL=\"%s\" is empty or invalid, fallling back to \"INFO\".\n", level)
		return colog.LInfo
	}
}
