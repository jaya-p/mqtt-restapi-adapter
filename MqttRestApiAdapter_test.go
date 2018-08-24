package main

import (
	"fmt"
	//import the Paho Go MQTT library
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)