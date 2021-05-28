package main

import "fmt"

// struct program
type program struct {
	programID string
	seasons   map[string]season
}

type season struct {
	seasonID string
	episodes map[string]string
}

func queryMap(programName, episodeNum, seasonNum string, programs map[string]program) {
	program := programs[programName]
	season := program.seasons[seasonNum]
	episodeID := season.episodes[episodeNum]
	fmt.Printf("ProgramID %s, SeasonID %s, EpisodeID %s", program.programID, season.seasonID, episodeID)
}

func DemoMaps() {
	programs := make(map[string]program)
	programs = map[string]program{
		"P1": program{
			"758E7893-5906-4208-AFCF-57C69DCAE37C",
			map[string]season{
				"S1": season{
					"CA3901EB-2122-4A6E-BFCE-30EC1D93F3B9",
					map[string]string{
						"E1": "7B280231-ACB8-4503-BA09-6404D798593F",
						"E2": "7B280231-ACB8-4503-BA09-6404D798593G",
						"E3": "7B280231-ACB8-4503-BA09-6404D798593H",
						"E4": "7B280231-ACB8-4503-BA09-6404D798593I",
						"E5": "7B280231-ACB8-4503-BA09-6404D798593J",
					},
				},
				"S2": season{
					"CA3901EB-2122-4A6E-BFCE-30EC1D93F3C0",
					map[string]string{
						"E1": "7B280231-ACB8-4503-BA09-6404D798594G",
					},
				},
			},
		},
	}

	queryMap("P1", "S1", "E2", programs)

	// Retrieve, Delete, Insert, Test if a key is present etc..
	// https://tour.golang.org/moretypes/22
}
