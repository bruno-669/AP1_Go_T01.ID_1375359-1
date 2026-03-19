package main

import (
	"strings"
	"testing"
)

func TestHandlerWord(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "нормальный случай",
			input:   "foo bar foo\n2\n",
			want:    "foo bar",
			wantErr: false,
		},
		{
			name:    "обычное поведение с K меньшим, чем количество уникальных слов",
			input:   "faaa aaa aaa aaa aaa aaa aaa aaa aaa aaa bb cc cc cc cc cc cc cc cc ee ee ff ff ff ff ff ff dd dd dd dd gg gg gg\n2\n",
			want:    "aaa cc",
			wantErr: false,
		}, {
			name:    "тестовый кейс из примера задания",
			input:   "aa bb cc aa cc cc cc aa ab ac bb\n3\n",
			want:    "cc aa bb",
			wantErr: false,
		},
		{
			name:    "передача пустого списка слов",
			input:   "\n5\n",
			want:    "",
			wantErr: false,
		}, {
			name:    "передача списка слов, где K больше, чем число уникальных слов",
			input:   "faaa aaa aaa aaa aaa aaa aaa aaa aaa aaa bb cc cc cc cc cc cc cc cc ee ee ff ff ff ff ff ff dd dd dd dd gg gg gg\n56\n",
			want:    "aaa cc ff dd gg ee bb faaa",
			wantErr: false,
		},
		{
			name:    "нет второй строки",
			input:   "foo bar",
			want:    "",
			wantErr: true,
			errMsg:  "invalid input",
		},
		{
			name:    "число не число",
			input:   "foo bar\nabc\n",
			want:    "",
			wantErr: true,
			errMsg:  "invalid input",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			got, err := HandlerWord(r)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandlerWord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
				t.Errorf("HandlerWord() error = %q, want containing %q", err.Error(), tt.errMsg)
			}
			if got != tt.want {
				t.Errorf("HandlerWord() = %q, want %q", got, tt.want)
			}
		})
	}
}
