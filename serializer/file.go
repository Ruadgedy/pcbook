package serializer

import (
	"bytes"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
)

// WriteProtobufToBinaryFile writes protobuf message to binary file
func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	bytes, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	if err = ioutil.WriteFile(filename, bytes, 0644); err != nil {
		return err
	}
	return nil
}

func ReadBinaryFileToProtobuf(filename string, message proto.Message) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	if err = proto.Unmarshal(bytes, message); err != nil {
		return err
	}
	return nil
}

// WriteJsonToBinaryFile writes JSON message to binary file
func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	json, err := ProtobufToJSON(message)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, []byte(json), 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadProtobufFromJSONFile(message proto.Message, filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	unmarshaler := jsonpb.Unmarshaler{
		AllowUnknownFields: true,
	}
	buffer := bytes.NewBuffer(data)
	if err = unmarshaler.Unmarshal(buffer, message); err != nil {
		return err
	}
	return nil
}
