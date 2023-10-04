package utils

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var myClient *http.Client

func Execute(url, method string, pData []byte, headers map[string]string) (result interface{}, err error) {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(pData))
	if err != nil {
		return "NewRequest err>>", err
	}
	if len(headers) > 0 {
		for key, value := range headers {
			request.Header.Add(key, value)
		}
	}
	if myClient == nil {
		myClient = &http.Client{}
	}
	response, err := myClient.Do(request)
	if err != nil {
		return "client.Do err>>", err
	}
	defer response.Body.Close()

	fmt.Println("response code>>", response.StatusCode)
	buff := &bytes.Buffer{}
	_, err = buff.ReadFrom(response.Body)
	if err != nil {
		return "readData err>>", err
	}
	return buff.Bytes(), nil

}

// 获取当前执行文件所在目录
func GetExeDir() (path string) {
	fPath, _ := os.Executable()
	return filepath.Dir(fPath)
}

// 输出文件到根目录 json
func writeJsonFile(pInfos []ProvinceInfo) {

	//写json
	buff, _ := json.Marshal(pInfos)
	fPath := fmt.Sprintf("%s\\json_%s.json", GetExeDir(), time.Now().Format("20060102150405"))
	os.WriteFile(fPath, buff, os.ModePerm)

}

// 输出文件到根目录   csv
func writeCsvFile(pInfos []ProvinceInfo) {

	// Create a new CSV file
	fPath := fmt.Sprintf("%s\\csv_%s.csv", GetExeDir(), time.Now().Format("20060102150405"))
	file, err := os.Create(fPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the CSV header
	header := []string{"type", "Province Name", "Province Code", "City Name", "City Code", "Area Name", "Area Code"}
	if err := writer.Write(header); err != nil {
		log.Fatal(err)
	}

	// Write data to the CSV file
	for _, pInfo := range pInfos {
		provinceName := pInfo.Name
		provinceCode := pInfo.Code
		//province
		record := []string{"省份", provinceName, provinceCode}
		writer.Write(record)
		for _, cInfo := range pInfo.CityInfo {
			cityName := cInfo.Name
			cityCode := cInfo.Code

			record = []string{"城市", provinceName, provinceCode, cityName, cityCode}
			writer.Write(record)
			for _, aInfo := range cInfo.AreaInfo {
				areaName := aInfo.Name
				areaCode := aInfo.Code

				record = []string{"区县", provinceName, provinceCode, cityName, cityCode, areaName, areaCode}
				if err := writer.Write(record); err != nil {
					log.Fatal(err)
				}
			}
		}
	}

	log.Printf("CSV file saved to: %s\n", fPath)
}
