package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type MapRoom struct {
	Id   string
	Name string
}

func main() {
	flag.Parse()
	host := flag.Arg(0)

	res, err := http.Get(fmt.Sprint("http://", host, "/api/v2/robot/capabilities/MapSegmentationCapability"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error making http request: %s\nMaking no changes.\n", err)
		os.Exit(0)
	}

	var body []byte
	body, err = io.ReadAll(res.Body)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading HTTP body: %s\n", err)
		os.Exit(1)
	}

	var rooms []MapRoom
	err = json.Unmarshal(body, &rooms)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing JSON: %s\n", err)
		os.Exit(1)
	}

	var sb strings.Builder
	var input_id string

	for i := 0; i < len(rooms); i++ {
		fmt.Fprintf(os.Stderr, "Found: %s\n", rooms[i].Name)
		input_id = strings.ReplaceAll(strings.ToLower(rooms[i].Name), " ", "")
		sb.WriteString(fmt.Sprintf("input_boolean.vacuum_%s:\n  room_id: %s \n", input_id, rooms[i].Id))
	}

	fmt.Print(sb.String())
}
