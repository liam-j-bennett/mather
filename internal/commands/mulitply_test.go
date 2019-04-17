package commands

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestMultiplyCommand(t *testing.T) {
	tests := []struct {
		name         string
		args         []string
		expectArgs   []int
		expectReturn int
		expectOutput string
		expectError  error
	}{
		{
			"Invalid left-hand number",
			[]string{"a", "2"},
			[]int{0, 0},
			0,
			"",
			errors.New("invalid first number a"),
		},
		{
			"Invalid right-hand number",
			[]string{"2", "a"},
			[]int{0, 0},
			0,
			"",
			errors.New("invalid second number a"),
		},
		{
			"Valid number multiply",
			[]string{"2", "2"},
			[]int{2, 2},
			4,
			"Numbers multiply to 4\n",
			nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockClient := NewMockMatherClient(ctrl)
			if test.expectError == nil {
				mockClient.EXPECT().Multiply(test.expectArgs[0], test.expectArgs[1]).Return(test.expectReturn, nil)
			}
			output := &strings.Builder{}
			cmd := newMultiplyCommand(mockClient)
			cmd.SetArgs(test.args)
			cmd.SetOutput(output)
			_, err := cmd.ExecuteC()
			if test.expectError != nil {
				require.Error(t, err)
				assert.Equal(t, test.expectError.Error(), err.Error())
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.expectOutput, output.String())
			}
		})
	}
}
