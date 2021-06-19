package main

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/yaml"
)

func GenerateStrategicMergePatchFromFile(fromFile, toFile string) (string, error){
	fromData, err := ioutil.ReadFile(fromFile)
	if err != nil{
		return "", errors.Errorf("failed to read file %s. error: %v", fromFile, err)
	}

	toData, err := ioutil.ReadFile(toFile)
	if err != nil{
		return "", errors.Errorf("failed to read file %s. error: %v", toFile, err)
	}

	return GenerateStrategicMergePatchFromByte(fromData, toData)
}

func GenerateStrategicMergePatchFromByte(fromData, toData []byte) (string, error){
	fromJson, err := yaml.YAMLToJSON(fromData)
	if err != nil {
		return "", err
	}

	toJson, err := yaml.YAMLToJSON(toData)
	if err != nil {
		return "", err
	}

	var u unstructured.Unstructured

	_, gvk, err := unstructured.UnstructuredJSONScheme.Decode(fromJson, nil, &u)
	if err != nil {
		return "", err
	}

	obj, err := scheme.Scheme.New(*gvk)
	if err != nil{
		return "", err
	}

	jsonPatch, err := strategicpatch.CreateTwoWayMergePatch(fromJson, toJson, obj)
	if err != nil{
		return "", err
	}
	var overlay map[string]interface{}
	err = json.Unmarshal(jsonPatch, &overlay)
	if err != nil {
		return "", err
	}
	//overlay["apiVersion"] = u.GetAPIVersion()
	//overlay["kind"] = u.GetKind()
	//err = unstructured.SetNestedField(overlay, u.GetName(), "metadata", "name")
	//if err != nil {
	//	return "", err
	//}

	var patch []byte
	if outputType == fmtJSON {
		patch, err = json.MarshalIndent(overlay, "", "  ")
	} else {
		patch, err = yaml.Marshal(overlay)
	}
	return string(patch), err
}