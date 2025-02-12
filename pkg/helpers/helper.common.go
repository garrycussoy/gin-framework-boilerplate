package helpers

import (
	uuid "github.com/nu7hatch/gouuid"
)

// Function to check whether an array contains an element or not
func IsArrayContains(arr []string, str string) bool {
	for _, item := range arr {
		if item == str {
			return true
		}
	}
	return false
}

// func GenerateRandomString(n int) (string, error) {
// 	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
// 	ret := make([]byte, n)
// 	for i := 0; i < n; i++ {
// 		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
// 		if err != nil {
// 			return "", err
// 		}
// 		ret[i] = letters[num.Int64()]
// 	}

// 	return string(ret), nil
// }

// // This function is used to remove any field on a map which have empty string value
// func RemoveEmptyField(data map[string]string) map[string]string {
// 	for k, v := range data {
// 		if v == "" {
// 			delete(data, k)
// 		}
// 	}
// 	return data
// }

// Generate UUID
func GenerateUUID() (string, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

// // Extract values from nullable string (if null, we will set the value to an empty string)
// func ExtractNullableString(val *string) string {
// 	if val != nil {
// 		return *val
// 	}
// 	return ""
// }

// // Convert a string to a pointer string
// func CreatePointerString(str string) *string {
// 	tempStr := str
// 	return &tempStr
// }

// // Read a CSV file into list of list of strings
// func ReadCSVFile(file *multipart.FileHeader) ([][]string, error) {
// 	// Open the file
// 	openedFile, err := file.Open()
// 	if err != nil {
// 		return [][]string{}, err
// 	}

// 	// Define CSV reader
// 	reader := csv.NewReader(openedFile)

// 	// Read the CSV file and turn it into list of list of string
// 	records, err := reader.ReadAll()
// 	if err != nil {
// 		return [][]string{}, err
// 	}

// 	// Close the file
// 	openedFile.Close()

// 	return records, nil
// }

// // Open a file
// func MustOpen(filePath string) *os.File {
// 	r, _ := os.Open(filePath)
// 	return r
// }

// // Convert os.File to multipartForm
// // This function is used for testing-purpose only
// func ConvertFileToMultipart(filePath string) (*multipart.FileHeader, string, error) {
// 	// Open the file
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return nil, "", err
// 	}
// 	defer file.Close()

// 	// Create a buffer to hold the file in memory
// 	var buff bytes.Buffer
// 	buffWriter := io.Writer(&buff)

// 	// Create a new form and create a new file field
// 	formWriter := multipart.NewWriter(buffWriter)
// 	formPart, err := formWriter.CreateFormFile("file", filepath.Base(file.Name()))
// 	if err != nil {
// 		return nil, "", err
// 	}

// 	// Get content-type
// 	contentType := formWriter.FormDataContentType()

// 	// Copy the content of the file to the form's file field
// 	io.Copy(formPart, file)

// 	// Close the form writer after the copying process is finished
// 	// I don't use defer in here to avoid unexpected EOF error
// 	formWriter.Close()

// 	// Transform the bytes buffer into a form reader
// 	buffReader := bytes.NewReader(buff.Bytes())
// 	formReader := multipart.NewReader(buffReader, formWriter.Boundary())

// 	// Read the form components with max stored memory of 1MB
// 	multipartForm, err := formReader.ReadForm(1 << 20)
// 	if err != nil {
// 		return nil, "", err
// 	}

// 	// Return the multipart file header
// 	files := multipartForm.File["file"]
// 	return files[0], contentType, nil
// }
