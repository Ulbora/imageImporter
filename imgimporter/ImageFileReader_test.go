package imgimporter

import (
	"fmt"
	"io/ioutil"
	"testing"

	fr "github.com/Ulbora/FileReader"
)

func TestCvsFile_PrepRecords(t *testing.T) {
	var cr fr.CsvFileReader
	//sourceFile, err := ioutil.ReadFile("../full_file.csv")
	sourceFile, err := ioutil.ReadFile("fullsize_images_test.csv")
	fmt.Println("readFile err: ", err)
	rd := cr.GetNew()
	rec := rd.ReadCsvFile(sourceFile)
	fmt.Println("csv err: ", rec.CsvReadErr)
	fmt.Println("csv len: ", len(rec.CsvFileList))
	type args struct {
		files *fr.CsvFiles
	}
	tests := []struct {
		name  string
		p     *CvsFile
		args  args
		want  *[]ImageRecord
		want2 int
	}{
		// TODO: Add test cases.
		{
			name: "test 1",
			//p:    rec,
			// fields: fields{
			// 	// RemoteURL: "CMM55BA48A",
			// 	// Name:      "https://3375977.app.netsuite.com/core/media/media.nl?id=100013&c=3375977&h=nzzZKBKt5M8cx9crVBWOOngkQkiJtkNp1GY9gH6LH6ap65It",
			// },
			args: args{
				files: rec,
			},
			want2: 4,
			// want: &ResponseID{
			// 	ID:      1,
			// 	Success: true,
			// },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &CvsFile{}
			// if got := p.ReadImageURLs(tt.args.files); !reflect.DeepEqual(got, tt.want) {
			if got := p.ReadImageURLs(tt.args.files); len(*got) != tt.want2 {
				t.Errorf("CvsFile.PrepRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}
