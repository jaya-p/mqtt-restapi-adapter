package main

import (
	"testing"
)
 func TestuseGETMethod(t *testing.T) {
 	result = useGETMethod()
 	if result != "off" or result != "on"{
 		t.Error("Expected off or on")
 	}
 }


func TestusePOSTMethod(t *testing.T) {
	result = usePOSTMethod()
	if result != '{"message":"Response Success","status_code":200}'{
		t.Error("Expected 200 success")
	}
}