package main

import (
	"bufio"
	"strconv"
	"strings"
	"testing"
)

func TestCanSortContainers(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name: "Example 1",
			input: `2
					1 2
					2 1`,
			want: "yes",
		},
		{
			name: "Example 2",
			input: `3
					10 20 30
					1 1 1
					0 0 1`,
			want: "no",
		},
		{
			name: "test 1",
			input: `3
					1 1 1
					2 2 2
					3 3 3`,
			want: "no",
		},
		{
			name: "test 2",
			input: `2
					1 1
					1 1`,
			want: "yes",
		},
		{
			name: "test 3",
			input: `3
					3 1 2
					1 2 1
					2 1 2`,
			want: "yes",
		},
		{
			name: "test 4",
			input: `3
					0 0 0
					0 0 0
					0 0 0`,
			want: "yes",
		},
		{
			name: "test 5",
			input: `2
					2 1
					3 1`,
			want: "no",
		},
		{
			name: "test 6",
			input: `2
					1 999999999
					999999999 1`,
			want: "yes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.input))
			scanner.Scan()
			n, _ := strconv.Atoi(scanner.Text())

			containers := make([][]int, n)
			for i := 0; i < n; i++ {
				containers[i] = make([]int, n)
				scanner.Scan()
				values := strings.Fields(scanner.Text())
				for j := 0; j < n; j++ {
					containers[i][j], _ = strconv.Atoi(values[j])
				}
			}

			got := canSortContainers(n, containers)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
