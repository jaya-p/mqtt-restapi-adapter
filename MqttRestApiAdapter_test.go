package main

import (
	"testing"
	"strings"
)


 func TestUseGETMethod(t *testing.T) {
 	result := useGETMethod()
 	if (result == "off"){
 		
 	}else if (result == "on"){

 		}else{
 			t.Error("expect on or off")
 		}
 }


func TestUsePOSTMethod(t *testing.T) {
	result := usePOSTMethod()
	value := strings.Compare("{'message':'Response Success','status_code':200}", result)
	if value == 0{
		t.Error(value)

	}
}



func TestSubscribeMessage(t *testing.T){
	subscribeMessage("jinawi-dev-1","1")
	if(state != "on"){
		t.Error("Expected on")
	}
	subscribeMessage("jinawi-dev-1","0")
	if(state != "off"){
		t.Error("Expected off")
	}


}



func TestCheckMessage(t *testing.T){
	checkMessage("on", mqttClient)//dengan asumsi state awal mati
	if(state != "on"){
		t.Error("Expected on")
	}
	checkMessage("off", mqttClient)
	if(state != "off"){
		t.Error("Expected off")
	}
}