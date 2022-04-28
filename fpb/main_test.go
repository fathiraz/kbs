package main

import (
	"fpb/model"
	"testing"
)

func TestMain_Main(t *testing.T) {
	type fields struct {
		Cakes  uint64
		Apples uint64
	}

	tests := []struct {
		name          string
		fields        fields
		wantBoxCakes  uint64
		wantBoxApples uint64
		wantBox       uint64
	}{
		{
			name: "Success Case" +
				" 4 cakes, 5 apples, 5 box",
			fields: fields{
				Cakes:  20,
				Apples: 25,
			},
			wantBox:       5,
			wantBoxCakes:  4,
			wantBoxApples: 5,
		},
		{
			name: "Success Case" +
				" 42 cakes, 56 apples, 14 box",
			fields: fields{
				Cakes:  42,
				Apples: 56,
			},
			wantBox:       14,
			wantBoxCakes:  3,
			wantBoxApples: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := model.Box{
				Cakes:  tt.fields.Cakes,
				Apples: tt.fields.Apples,
			}
			if got := b.GetBoxCakes(); got != tt.wantBoxCakes {
				t.Errorf("GetBoxCakes() = %v, want %v", got, tt.wantBoxCakes)
			}
			if got := b.GetBoxApples(); got != tt.wantBoxApples {
				t.Errorf("GetBoxApples() = %v, want %v", got, tt.wantBoxApples)
			}
			if got := b.GetBoxEvenly(); got != tt.wantBox {
				t.Errorf("GetBoxEvenly() = %v, want %v", got, tt.wantBox)
			}
		})
	}
}
