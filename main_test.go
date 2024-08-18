package main

import (
	"fmt"
	"slices"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		users                  []user
		mv                     move
		expectedFightLocations []piece
	}
	tests := []testCase{
		{
			users: []user{
				{
					name: "Toussaint",
					pieces: []piece{
						{
							location: "San Domingo",
							name:     "Cavalry",
						},
						{
							location: "San Domingo",
							name:     "Infantry",
						},
					},
				},
				{
					name: "Napoleon",
					pieces: []piece{
						{
							location: "France",
							name:     "Infantry",
						},
						{
							location: "Russia",
							name:     "Infantry",
						},
					},
				},
				{
					name: "Washington",
					pieces: []piece{
						{
							location: "United States",
							name:     "Artillery",
						},
					},
				},
			},
			mv: move{
				userName: "Toussaint",
				piece: piece{
					location: "United States",
					name:     "Cavalry",
				},
			},
			expectedFightLocations: []piece{
				{
					location: "United States",
					name:     "Artillery",
				},
			},
		},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{
				users: []user{
					{
						name: "Toussaint",
						pieces: []piece{
							{
								location: "San Domingo",
								name:     "Cavalry",
							},
							{
								location: "San Domingo",
								name:     "Infantry",
							},
						},
					},
					{
						name: "Napoleon",
						pieces: []piece{
							{
								location: "France",
								name:     "Infantry",
							},
							{
								location: "Russia",
								name:     "Infantry",
							},
							{
								location: "United States",
								name:     "Cavalry",
							},
						},
					},
					{
						name: "Washington",
						pieces: []piece{
							{
								location: "United States",
								name:     "Artillery",
							},
						},
					},
				},
				mv: move{
					userName: "Toussaint",
					piece: piece{
						location: "United States",
						name:     "Cavalry",
					},
				},
				expectedFightLocations: []piece{
					{
						location: "United States",
						name:     "Cavalry",
					},
					{
						location: "United States",
						name:     "Artillery",
					},
				},
			},
		}...)
	}

	for _, test := range tests {
		bufferedCh := make(chan move, 100)
		mover := user{}
		for _, u := range test.users {
			if u.name == test.mv.userName {
				mover = u
			}
		}
		if mover.name == "" {
			t.Errorf("Test Failed: user with name %v not found", test.mv.userName)
		}
		mover.march(test.mv.piece, bufferedCh)
		close(bufferedCh)
		output := doBattles(bufferedCh, test.users)
		if !slices.Equal(output, test.expectedFightLocations) {
			t.Errorf(`
Test Failed:
  users:
%v
  move: %v
  =>
  expected battle pieces:
%v
  actual battle pieces:
%v
			`,
				formatSlice(test.users),
				test.mv,
				formatSlice(test.expectedFightLocations),
				formatSlice(output),
			)
		} else {
			fmt.Printf(`
Test Passed:
  users:
%v
  move: %v
  =>
  expected battle pieces:
%v
  actual battle pieces:
%v
			`,
				formatSlice(test.users),
				test.mv,
				formatSlice(test.expectedFightLocations),
				formatSlice(output),
			)
		}
	}
}

func formatSlice[T any](slice []T) string {
	if slice == nil {
		return "nil"
	}
	output := ""
	for i, v := range slice {
		output += fmt.Sprintf("* %v", v)
		if i < len(slice)-1 {
			output += "\n"
		}
	}
	return output
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
