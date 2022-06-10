package testutils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

// mock server for testing requests
func HttpMock(pattern string, statusCode int, respBody any) *httptest.Server {
	handler := http.NewServeMux()
	mockHandler := createMockHandler(statusCode, respBody)
	handler.HandleFunc(pattern, mockHandler)
	return httptest.NewServer(handler)
}

func createMockHandler(statusCode int, respBody any) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, req *http.Request) {
		data, err := json.Marshal(respBody)
		if err != nil {
			data = []byte("{}")
		}
		writer.WriteHeader(statusCode)
		_, err = writer.Write(data)
		if err != nil {
			fmt.Println("error writing data", err)
			writer.Write([]byte("{}"))
		}
	}
}

// #############
// ## example ##
// #############

// func TestGetExampleEmployees(t *testing.T) {
// 	resp := &EmployeesResponse{
// 		Status:  "success",
// 		Message: "abc",
// 		Data: []EmployeeData{
// 			{
// 				ID:             1,
// 				EmployeeName:   "felix",
// 				EmployeeSalary: 500,
// 				EmployeeAge:    31,
// 				ProfileImage:   "",
// 			},
// 		},
// 	}
// 	pattern := "/api/employees"
// 	srv := HttpMock(pattern, http.StatusOK, resp)
// 	defer srv.Close()
// 	url := srv.URL + pattern

// 	patternErr := "/api/employees/err"
// 	srvErr := HttpMock(patternErr, http.StatusBadRequest, resp)
// 	defer srvErr.Close()
// 	urlErr := srvErr.URL + patternErr

// 	type args struct {
// 		url string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    *EmployeesResponse
// 		wantErr bool
// 	}{
// 		{
// 			"get 200",
// 			args{url: url},
// 			resp,
// 			false,
// 		},
// 		{
// 			"get 400",
// 			args{url: urlErr},
// 			nil,
// 			true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := GetExampleEmployees(tt.args.url)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetExampleEmployees() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("GetExampleEmployees() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
