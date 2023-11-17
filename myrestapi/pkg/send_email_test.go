package pkg

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestSendEmail(t *testing.T) {
	err := godotenv.Load("../.env")

	if err != nil {
		t.Error("Test failed")
	}

	ec := NewEmailClient()

	
	statusCode, err := ec.SendEmail("Hello World From Test", "mathematics06physics@gmail.com")

	if statusCode > 300 {
		t.Errorf("Expected a status code within the 200 range but got %d", statusCode)
	}

	if err != nil {
		t.Errorf("Test failed with err %v", err)
	}

}
