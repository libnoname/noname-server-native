package api

// A generic struct to serialize data in web responses.
type ReturnData[T any] struct {
	// Success indicate whether the operation was successful.
	Success  bool   `json:"success"`
	Code     int    `json:"code"`
	Data     T      `json:"data,omitempty"`
	ErrorMsg string `json:"errorMsg,omitempty"`
}

// Return data when calling getFileList.
type FileListData struct {
	Files   []string `json:"files"`
	Folders []string `json:"folders"`
}

func SuccessResult[T any](data T) ReturnData[T] {
	return ReturnData[T]{
		Success:  true,
		Code:     200,
		Data:     data,
		ErrorMsg: "",
	}
}

func FailedResult(code int, err string) ReturnData[any] {
	return ReturnData[any]{
		Success:  false,
		Code:     code,
		Data:     nil,
		ErrorMsg: err,
	}
}
