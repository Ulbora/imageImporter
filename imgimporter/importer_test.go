package imgimporter

import "testing"

func TestImageImporter_ProcessImages(t *testing.T) {
	type args struct {
		csvFile string
	}
	tests := []struct {
		name string
		p    *ImageImporter
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test_1",
			args: args{
				csvFile: "fullsize_images_test.csv",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &ImageImporter{}
			p.ProcessImages(tt.args.csvFile)
		})
	}
}
