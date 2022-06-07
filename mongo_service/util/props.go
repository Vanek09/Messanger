package util

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type PROPERTIES struct {
	props map[string]string
}

func ReadProperties(filename string) (PROPERTIES) {
	props_file, err := os.Open(filename)
	if err != nil {
		log.Println(err.Error())
		return PROPERTIES{}
	}
	defer props_file.Close()

	props := PROPERTIES{
		props: make(map[string]string),
	}

	reader := bufio.NewReader(props_file)
	for {
		line, err := reader.ReadString('\n')
		if sep := strings.Index(line, "="); sep > 0 {
			if key := strings.TrimSpace(line[:sep]); len(key) > 0 {
				var val string
				if len(line) > sep {
					val = strings.TrimSpace(line[sep + 1:])
				}
				props.props[key] = val
			}
		}
		if err == io.EOF {
			break
		} else if err != nil {
			return PROPERTIES{}
		}	
	}
	return props
}

func (m PROPERTIES) GetBoolean(key string, defaultValue bool) bool {
	prop := m.props[key]
	if prop != "" {
		val, err := strconv.ParseBool(prop)
		if err != nil {
			log.Println(err.Error())
			return defaultValue
		}
		return val
	}
	return defaultValue
}

func (m PROPERTIES) GetString(key string, defaultValue string) string {
	prop := m.props[key]
	if prop != "" {
		return prop
	}
	return defaultValue
}

func (m PROPERTIES) GetInteger(key string, defaultValue int64) int64 {
	prop := m.props[key]
	if prop != "" {
		val, err := strconv.ParseInt(prop, 0, 64)
		if err != nil {
			log.Println(err.Error())
			return defaultValue
		}
		return val
	}
	return defaultValue
}