package main

import (
	"encoding/json"
	"github.com/appscode/jsonpatch"
	"io/ioutil"
	"github.com/pkg/errors"
	"sigs.k8s.io/yaml"
)

func GenerateJsonPatchFromFile(fromFile, toFile string) (string, error) {
	fromData, err := ioutil.ReadFile(fromFile)
	if err != nil{
		return "", errors.Errorf("failed to read file %s. error: %v", fromFile, err)
	}

	toData, err := ioutil.ReadFile(toFile)
	if err != nil{
		return "", errors.Errorf("failed to read file %s. error: %v", toFile, err)
	}

	return GenerateJsonPatchFromBytes(fromData, toData)
}

func GenerateJsonPatchFromBytes(fromData, toData []byte) (string, error){
	fromJson, err := yaml.YAMLToJSON(fromData)
	if err != nil {
		return "", err
	}

	toJson, err := yaml.YAMLToJSON(toData)
	if err != nil {
		return "", err
	}

	jsonPatch, err := jsonpatch.CreatePatch(fromJson, toJson)
	if err != nil {
		return "", err
	}

	var patch []byte
	if outputType == fmtJSON {
		patch, err = json.MarshalIndent(jsonPatch, "", "  ")
	}else {
		patch, err = yaml.Marshal(jsonPatch)
	}
	return string(patch), err
}