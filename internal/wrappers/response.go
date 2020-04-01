package wrappers

import (
	"encoding/json"
	"net/http"

	projectsApi "github.com/checkmarxDev/scans/api/v1/rest/projects"
	scansApi "github.com/checkmarxDev/scans/api/v1/rest/scans"
	"github.com/pkg/errors"
)

const (
	failedToParseErr = "Failed to parse error response"
)

func handleScanResponseWithNoBody(resp *http.Response, err error,
	successStatusCode int) (*scansApi.ErrorModel, error) {
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(resp.Body)

	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusBadRequest, http.StatusInternalServerError:
		errorModel := scansApi.ErrorModel{}
		err = decoder.Decode(&errorModel)
		if err != nil {
			return nil, errors.Wrapf(err, failedToParseErr)
		}
		return &errorModel, nil
	case successStatusCode:
		return nil, nil

	default:
		return nil, errors.Errorf("Unknown response status code %d", resp.StatusCode)
	}
}

func handleScanResponseWithBody(resp *http.Response, err error,
	successStatusCode int) (*scansApi.ScanResponseModel, *scansApi.ErrorModel, error) {
	if err != nil {
		return nil, nil, err
	}
	decoder := json.NewDecoder(resp.Body)

	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusBadRequest, http.StatusInternalServerError:
		errorModel := scansApi.ErrorModel{}
		err = decoder.Decode(&errorModel)
		if err != nil {
			return responseScanParsingFailed(err)
		}
		return nil, &errorModel, nil
	case successStatusCode:
		model := scansApi.ScanResponseModel{}
		err = decoder.Decode(&model)
		if err != nil {
			return responseScanParsingFailed(err)
		}
		return &model, nil, nil

	default:
		return nil, nil, errors.Errorf("Unknown response status code %d", resp.StatusCode)
	}
}

func handleProjectResponseWithNoBody(resp *http.Response, err error,
	successStatusCode int) (*projectsApi.ErrorModel, error) {
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(resp.Body)

	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusBadRequest, http.StatusInternalServerError:
		errorModel := projectsApi.ErrorModel{}
		err = decoder.Decode(&errorModel)
		if err != nil {
			return nil, errors.Wrapf(err, failedToParseErr)
		}
		return &errorModel, nil
	case successStatusCode:
		return nil, nil

	default:
		return nil, errors.Errorf("Unknown response status code %d", resp.StatusCode)
	}
}

func handleProjectResponseWithBody(resp *http.Response, err error,
	successStatusCode int) (*projectsApi.ProjectResponseModel, *projectsApi.ErrorModel, error) {
	if err != nil {
		return nil, nil, err
	}
	decoder := json.NewDecoder(resp.Body)

	defer resp.Body.Close()
	switch resp.StatusCode {
	case http.StatusBadRequest, http.StatusInternalServerError:
		errorModel := projectsApi.ErrorModel{}
		err = decoder.Decode(&errorModel)
		if err != nil {
			return responseProjectParsingFailed(err)
		}
		return nil, &errorModel, nil
	case successStatusCode:
		model := projectsApi.ProjectResponseModel{}
		err = decoder.Decode(&model)
		if err != nil {
			return responseProjectParsingFailed(err)
		}
		return &model, nil, nil

	default:
		return nil, nil, errors.Errorf("Unknown response status code %d", resp.StatusCode)
	}
}

func responseScanParsingFailed(err error) (*scansApi.ScanResponseModel, *scansApi.ErrorModel, error) {
	msg := "Failed to parse scan response"
	return nil, nil, errors.Wrapf(err, msg)
}
func responseProjectParsingFailed(err error) (*projectsApi.ProjectResponseModel, *projectsApi.ErrorModel, error) {
	msg := "Failed to parse project response"
	return nil, nil, errors.Wrapf(err, msg)
}