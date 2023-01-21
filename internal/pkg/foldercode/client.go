package foldercode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bilgehannal/foldercode-cli/internal/env"
	"github.com/bilgehannal/foldercode-cli/internal/pkg/errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type FoldercodeClient struct{}

type Session struct {
	Id         string `json:"id"`
	FolderCode string `json:"folderCode"`
}

type FileResponse struct {
	Files []File `json:"files"`
}

type File struct {
	FileName string `json:"filename"`
	FileUrl  string `json:"file_url"`
}

func (fc FoldercodeClient) GenerateSession() (Session, error) {
	url := fmt.Sprintf("%s/api/v1/uploading_sessions", env.ApiUrlBase)
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return Session{}, err
	}
	res, err := client.Do(req)
	if err != nil {
		return Session{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Session{}, err
	}
	var session Session
	err = json.Unmarshal(body, &session)
	if err != nil {
		return Session{}, err
	}
	return session, nil
}

func (fc FoldercodeClient) UploadFile(s Session, fileName string) error {
	url := fmt.Sprintf("%s/upload?browserSession=%s", env.ApiUrlBase, s.Id)
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(fileName)
	defer file.Close()
	part1,
		errFile1 := writer.CreateFormFile("file", filepath.Base(fileName))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		return errFile1
	}
	err := writer.Close()
	if err != nil {
		return err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}

func (fc FoldercodeClient) GetFiles(code string) ([]File, error) {
	url := fmt.Sprintf("%s/api/v1/folder_files/%s", env.ApiUrlBase, code)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return []File{{}}, err
	}
	res, err := client.Do(req)
	if err != nil {
		return []File{{}}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []File{{}}, err
	}
	var fileResponse FileResponse
	err = json.Unmarshal(body, &fileResponse)
	if err != nil {
		return []File{{}}, err
	}
	return fileResponse.Files, nil
}

func (fc FoldercodeClient) DownloadFile(file File) error {
	out, err := os.Create(file.FileName)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(file.FileUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return errors.DownloadError{}
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
