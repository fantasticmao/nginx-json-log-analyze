package handler

import "github.com/fantasticmao/nginx-json-log-analyzer/ioutil"

const (
	AnalyzeTypePvUv = iota
	AnalyzeTypeFieldIp
	AnalyzeTypeFieldUri
	AnalyzeTypeFieldUserAgent
	AnalyzeTypeFieldUserCountry
	AnalyzeTypeFieldUserCity
	AnalyzeTypeResponseStatus
	AnalyzeTypeTimeMeanCostUris
	AnalyzeTypeTimePercentCostUris
)

type Handler interface {
	Input(info *ioutil.LogInfo)

	Output(limit int)
}

func NewHandler(analyzeType int, percentile float64) Handler {
	switch analyzeType {
	case AnalyzeTypePvUv:
		return NewPvAndUvHandler()
	case AnalyzeTypeFieldIp:
		return NewMostMatchFieldHandler(AnalyzeTypeFieldIp)
	case AnalyzeTypeFieldUri:
		return NewMostMatchFieldHandler(AnalyzeTypeFieldUri)
	case AnalyzeTypeFieldUserAgent:
		return NewMostMatchFieldHandler(AnalyzeTypeFieldUserAgent)
	case AnalyzeTypeResponseStatus:
		return NewMostFrequentStatusHandler()
	case AnalyzeTypeTimeMeanCostUris:
		return NewTopTimeMeanCostUrisHandler()
	case AnalyzeTypeTimePercentCostUris:
		return NewTopTimePercentCostUrisHandler(percentile)
	default:
		ioutil.Fatal("unsupported analyze type: %v\n", analyzeType)
		return nil
	}
}