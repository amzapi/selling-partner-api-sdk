package decryption

import (
	"io/ioutil"
	"reflect"
	"testing"
)

const ecryptedFileName1 = "./encrypted_report_example1.bin"
const decryptedFileName1 = "./decrypted_report_example1.bin"

func Test_decrypt(t *testing.T) {
	type args struct {
		encryptionKey string
		initVector    string
		dataToDecrypt []byte
	}

	encryptedData1, err := ioutil.ReadFile(ecryptedFileName1)
	if err != nil {
		t.Fatal("Failed to read the file " +ecryptedFileName1+ ": ", err )
	}
	expectedDecryptedData1, err := ioutil.ReadFile(decryptedFileName1)
	if err != nil {
		t.Fatal("Failed to read the file " +decryptedFileName1+ ": ", err )
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Successful decryption of a report",
			args : args{
				encryptionKey:"5XgmfCcQSWn5q4u69kM5O0XTsgP2Fzbq+GcHpWJOGR0=",
				initVector:"bXCrfJN8QrsWVRZ7J9aPIg==",
				dataToDecrypt: encryptedData1,
			},
			want : expectedDecryptedData1,
			wantErr: false,
		},
		{
			name: "Decryption of a report failed: invalid base64 key encoding",
			args : args{
				encryptionKey:"5XgmfCcQSWn5q4u69kM5O0XTsgP2Fzbq+GcHpWJOGR0=%",
				initVector:"bXCrfJN8QrsWVRZ7J9aPIg==",
				dataToDecrypt: encryptedData1,
			},
			want : nil,
			wantErr: true,
		},
		{
			name: "Decryption of a report failed: invalid base64 init vector encoding",
			args : args{
				encryptionKey:"5XgmfCcQSWn5q4u69kM5O0XTsgP2Fzbq+GcHpWJOGR0=",
				initVector:"bXCrfJN8QrsWVRZ7J9aPIg==%",
				dataToDecrypt: encryptedData1,
			},
			want : nil,
			wantErr: true,
		},
		{
			name: "Decryption of a report failed: invalid key length",
			args : args{
				encryptionKey:"dGVzdHdyb25na2V5",
				initVector:"bXCrfJN8QrsWVRZ7J9aPIg==",
				dataToDecrypt: encryptedData1,
			},
			want : nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Decrypt(tt.args.encryptionKey, tt.args.initVector, tt.args.dataToDecrypt)
			if (err != nil) != tt.wantErr {
				t.Errorf("decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}