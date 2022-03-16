package unit

import (
	"fmt"
	"testing"

	"gitlab.ozon.dev/qa/teachers/qa-route-256/go-unit-test/mock"
)

func TestIsWin(t *testing.T) {
	tests := []struct {
		bet  int
		dice Dice
		want bool
	}{
		{
			bet:  1,
			dice: mock.NewDiceMock(t).ThrowMock.Return(1),
			want: true,
		},
		{
			bet:  2,
			dice: mock.NewDiceMock(t).ThrowMock.Return(1),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("bet = %d", tt.bet), func(t *testing.T) {
			got := IsWin(tt.bet, tt.dice)

			if got != tt.want {
				t.Errorf("IsWin() = %v, bet = %v, want %v", got, tt.bet, tt.want)
			}
		})
	}
}
