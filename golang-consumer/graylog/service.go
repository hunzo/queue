package graylog

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang-consumer/config"
	"golang-consumer/models"
	"io"
	"log"
	"net/http"
	"os"
	_ "time/tzdata"
)

func LogToGraylog(payload models.GraylogPayload) error {

	GRAYLOG_SERVER := config.New().GraylogSrv

	reqBody, err := json.Marshal(models.GraylogPayload{
		Version:      payload.Version,
		Host:         os.Getenv("HOSTNAME"),
		ShortMessage: payload.ShortMessage,
		StartTime:    payload.StartTime,
		Count:        payload.Count,
		Check:        1,
	})

	fmt.Printf("\ncount message: %d\n", payload.Count)
	fmt.Printf("\nhostname: %s\n", os.Getenv("HOSTNAME"))

	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Printf("%s", GRAYLOG_SERVER)

	ret, err := http.Post(GRAYLOG_SERVER, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println(err)
		return err
	}

	// body, err := io.ReadAll(ret.Body)
	_, e := io.ReadAll(ret.Body)
	if e != nil {
		log.Println(e)
		return e
	}

	// log.Println(string(body))

	return nil
}
