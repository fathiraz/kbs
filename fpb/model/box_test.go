package model

import "testing"

func TestBox_GetBoxCakes(t *testing.T) {
	type fields struct {
		Cakes  uint64
		Apples uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		{
			name: "Success Case",
			fields: fields{
				Cakes:  1,
				Apples: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Box{
				Cakes:  tt.fields.Cakes,
				Apples: tt.fields.Apples,
			}
			if got := b.GetBoxCakes(); got != tt.want {
				t.Errorf("GetBoxCakes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBox_GetBoxApples(t *testing.T) {
	type fields struct {
		Cakes  uint64
		Apples uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		{
			name: "Success Case",
			fields: fields{
				Cakes:  1,
				Apples: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := Box{
				Cakes:  tt.fields.Cakes,
				Apples: tt.fields.Apples,
			}
			if got := b.GetBoxApples(); got != tt.want {
				t.Errorf("GetBoxApples() = %v, want %v", got, tt.want)
			}
		})
	}
}
