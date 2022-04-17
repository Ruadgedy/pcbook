package serializer

import (
	"github.com/Ruadgedy/pcbook/pb"
	"github.com/Ruadgedy/pcbook/sample"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProtobufToFile(t *testing.T)  {
	laptop := sample.NewLaptop()
	err := WriteProtobufToBinaryFile(laptop, "laptop.bin")
	assert.NoError(t,err)
}

func TestFileToProtobuf(t *testing.T) {
	laptop := &pb.Laptop{}
	err := ReadBinaryFileToProtobuf("laptop.bin",laptop)
	assert.NoError(t,err)
	assert.NotNil(t, laptop)

	t.Log(laptop)
}

func TestProtobufToJSON(t *testing.T)  {
	laptop := sample.NewLaptop()
	err := WriteProtobufToJSONFile(laptop, "laptop.json")
	assert.NoError(t,err)
}

func TestReadJsonFileToProtobuf(t *testing.T)  {
	laptop := &pb.Laptop{}
	err := ReadProtobufFromJSONFile(laptop, "laptop.json")
	assert.NoError(t,err)
	t.Log(laptop)
}