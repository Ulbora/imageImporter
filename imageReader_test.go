package imageimporter

import "testing"

func TestRemoteImage_ReadImage(t *testing.T) {
	type fields struct {
		RemoteURL string
		Name      string
	}
	type args struct {
		name      string
		remoteURL string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name:   "test 1",
			fields: fields{
				// RemoteURL: "CMM55BA48A",
				// Name:      "https://3375977.app.netsuite.com/core/media/media.nl?id=100013&c=3375977&h=nzzZKBKt5M8cx9crVBWOOngkQkiJtkNp1GY9gH6LH6ap65It",
			},
			args: args{
				name:      "CMM55BA48A",
				remoteURL: "https://3375977.app.netsuite.com/core/media/media.nl?id=100013&c=3375977&h=nzzZKBKt5M8cx9crVBWOOngkQkiJtkNp1GY9gH6LH6ap65It",
			},
			// want: &ResponseID{
			// 	ID:      1,
			// 	Success: true,
			// },
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &RemoteImage{
				// RemoteURL: tt.fields.RemoteURL,
				// Name:      tt.fields.Name,
			}
			m.ReadImage(tt.args.name, tt.args.remoteURL)
		})
	}
}
