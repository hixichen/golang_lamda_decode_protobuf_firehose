package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	pb "eventlog"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golang/protobuf/proto"
	"reflect"
)

type EventData struct {
	EventType uint32 `json:"eventType"`
}

func decodeProtobuf(in string) (string, error) {

	decodeBytes, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		return "", err
	}

	eventLog := &pb.EventLog{}
	err = proto.Unmarshal(decodeBytes, eventLog)
	if err != nil {
		return "", err
	}

	event := &EventData{}
	err = json.Unmarshal([]byte(eventLog.Data), &event)
	if err != nil {
		return "", err
	}

	retData, err := json.Marshal(event)
	if err != nil {
		return "", err
	}
	retData = append(retData, '\n')

	return base64.StdEncoding.EncodeToString(retData), nil
}

func handler(ctx context.Context, in map[string]interface{}) (interface{}, error) {

	t := in["records"]
	ret := make(map[string][]interface{})
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)
		for i := 0; i < s.Len(); i++ {

			value, _ := s.Index(i).Interface().(map[string]interface{})
			data := value["data"].(string)
			decodedData, err := decodeProtobuf(data)
			record := make(map[string]string)
			record["recordId"] = value["recordId"].(string)

			if err != nil {
				record["data"] = data
				record["result"] = "ProcessingFailed"
			} else {
				record["data"] = decodedData
				record["result"] = "Ok"
			}

			ret["records"] = append(ret["records"], record)
		}
	}
	return ret, nil
}

func main() {
	lambda.Start(handler)
}
