package golearn

import (
	"testing"
)

func TestDefaultEngine_MaxSpeed(t *testing.T) {
	tests := []struct {
		name string
		e    DefaultEngine
		want int
	}{
		{
			name: "default engine should have 50 max speed",
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.MaxSpeed(); got != tt.want {
				t.Errorf("DefaultEngine.MaxSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTurboEngine_MaxSpeed(t *testing.T) {
	tests := []struct {
		name string
		e    TurboEngine
		want int
	}{
		{
			name: "default engine should have 100 max speed",
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.MaxSpeed(); got != tt.want {
				t.Errorf("TurboEngine.MaxSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCar_Speed(t *testing.T) {
	type fields struct {
		Speeder Speeder
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "speed should be 50 when use default engine",
			fields: fields{Speeder: &DefaultEngine{}},
			want:   50,
		},
		{
			name:   "speed should be 80 when use default engine",
			fields: fields{Speeder: &TurboEngine{}},
			want:   80,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Car{
				Speeder: tt.fields.Speeder,
			}
			if got := c.Speed(); got != tt.want {
				t.Errorf("Car.Speed() = %v, want %v", got, tt.want)
			}
		})
	}
}
