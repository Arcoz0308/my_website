package error

import (
	"encoding/json"
)

type ApiError struct {
	Message        string
	Status         int
	Code           int
	Err            error
	AdditionalInfo map[string]interface{}
	Req            *RequestInfos
}
type RequestInfos struct {
	Path    string
	Body    string
	Headers map[string]string
}

func (a *ApiError) PrintErrorWithDebug() {
	logger.Logger.Errorf("error : %s", a.Err.Error())
	for k, v := range a.AdditionalInfo {
		logger.Logger.Debugf("%s = %#v", k, v)
	}
	if a.Req != nil {
		logger.Logger.Debugf("path : %s", a.Req.Path)
		if a.Req.Body != "" {
			logger.Logger.Debugf("body : %s", a.Req.Body)
		}
		logger.Logger.Debugf("headers : %v", a.Req.Headers)
	}
}

func (a *ApiError) Error() string {
	return a.Err.Error()
}
func (a *ApiError) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Status  int
		Message string
		Code    int
	}{
		a.Status,
		a.Message,
		a.Code,
	})
}
