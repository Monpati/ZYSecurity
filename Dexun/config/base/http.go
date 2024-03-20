package base

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

var (
	avgHttpProtocol = flag.String("http_protocol", "", "http_protocol")
	avgHttpHost     = flag.String("http_host", "", "http_host")
	avgHttpPort     = flag.Uint("http_port", 8080, "http_port")
	avgHttpUpload   = flag.String("http_upload", ".", "http_upload")
)

type HttpConf struct {
	Protocol string
	Host     string
	Port     int
	Upload   string
}

func NewHttpConf() *HttpConf {
	return &HttpConf{
		Protocol: *avgHttpProtocol,
		Host:     *avgHttpHost,
		Port:     int(*avgHttpPort),
		Upload:   *avgHttpUpload,
	}
}

func (p *HttpConf) LoadFromFile(name string) error {
	file, _ := os.Open(name)
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(p)
}

func (p *HttpConf) LoadFromBytes(data []byte) error {
	return json.Unmarshal(data, p)
}

func (p *HttpConf) SaveToFile(name string) error {
	flag := os.O_RDWR | os.O_TRUNC
	file, _ := os.OpenFile(name, flag, 0644)
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(p)
}

func (p *HttpConf) GetPrefix() string {
	return fmt.Sprintf("%s://%s", p.Protocol, p.Host)
}
