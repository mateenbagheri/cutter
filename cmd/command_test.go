package cmd

import (
	"reflect"
	"testing"
)

func TestCommandParseFlags(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantFields string
		wantDelim  string
		wantFiles  []string
		wantErr    bool
	}{
		{
			name:       "basic with -f and file",
			args:       []string{"-f", "1,2", "data.txt"},
			wantFields: "1,2",
			wantDelim:  " ",
			wantFiles:  []string{"data.txt"},
			wantErr:    false,
		},
		{
			name:       "custom delimiter",
			args:       []string{"-d", ",", "-f", "1", "a.csv", "b.csv"},
			wantFields: "1",
			wantDelim:  ",",
			wantFiles:  []string{"a.csv", "b.csv"},
			wantErr:    false,
		},
		{
			name:       "no flags, just file",
			args:       []string{"file.txt"},
			wantFields: "",
			wantDelim:  " ",
			wantFiles:  []string{"file.txt"},
			wantErr:    false,
		},
		{
			name:    "invalid flag",
			args:    []string{"-x", "value", "file.txt"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var cmd Command
			err := cmd.ParseFlags(tt.args)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseFlags() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if cmd.Fields != tt.wantFields {
				t.Errorf("Fields = %q, want %q", cmd.Fields, tt.wantFields)
			}
			if cmd.Delimiter != tt.wantDelim {
				t.Errorf("Delimiter = %q, want %q", cmd.Delimiter, tt.wantDelim)
			}
			if !reflect.DeepEqual(cmd.Files, tt.wantFiles) {
				t.Errorf("Files = %v, want %v", cmd.Files, tt.wantFiles)
			}
		})
	}
}

func TestCommandValidate(t *testing.T) {
	tests := []struct {
		name    string
		files   []string
		wantErr bool
	}{
		{"has files", []string{"a.txt"}, false},
		{"multiple files", []string{"a.txt", "b.txt"}, false},
		{"no files", []string{}, true},
		{"nil slice", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := Command{Files: tt.files}
			err := cmd.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
