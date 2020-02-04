package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/user"
	"path"
	"time"
)

var REPO = ""

func setRepo() {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	REPO = path.Join(usr.HomeDir, "github/testgen")
}

// Create integration tests, i.e. feature files
func MakeIntegration() error {

	var n = 4 // num verts = num of components

	err, w := MakeWalk(n)
	if err != nil {
		fmt.Errorf("Failed to make the graph walk: %v", err)
		os.Exit(1)
	}

	fmt.Printf("MakeWalk: done.\n")
	fmt.Printf("%v\n", w)

	err = MakeFeature(n, w)
	if err != nil {
		fmt.Errorf("Failed to make a feature file from the graph walk")
		os.Exit(1)
	}

	fmt.Printf("MakeFeature: done.\n")

	return nil
}

// Append scenario from the library to the feature file to form a user
// story
// featureName:   name of the feature file being created
// scenarioName:  name of the scenario to be appended
func AddScenarioToFeature(featureFname string, scenarioName string) error {
	scenarioName = fmt.Sprintf("%s.feature", scenarioName)
	scenarioFname := path.Join(REPO, "lib", scenarioName)
	scenarioBytes, err := ioutil.ReadFile(scenarioFname)
	if err != nil {
		fmt.Errorf("Failed reading data from file: %s", err)
		return err
	}

	featureFile, err := os.OpenFile(featureFname, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Errorf("failed opening file: %s", err)
		return err
	}
	defer featureFile.Close()

	_, err = featureFile.Write(scenarioBytes)
	if err != nil {
		fmt.Errorf("failed writing to file: %s", err)
		return err
	}

	return nil
}

// create a feature.feature file from the graph walk
// w: graph walk (user story)
func MakeFeature(n int, walk []int) error {

	featureFullName := "feature.feature"
	featureFname := path.Join(REPO, "feats", featureFullName)
	m := make(map[int]string)

	m[1] = "setup"
	m[2] = "start"
	m[3] = "stop"
	m[4] = "delete"

	f, err := os.Create(featureFname)
	if err != nil {
		fmt.Println(err)
		fmt.Errorf("Failed to create a feature file")
		os.Exit(1)
	}
	defer f.Close()

	_, err = f.WriteString("Feature: Auto-generated feature\n\t Testing basic functionality of CRC\n\n")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	f.Sync()

	// for each step in the walk, append the respective scenario to
	// the story
	for i := 0; i < len(walk); i++ {
		if walk[i] != 0 {
			err = AddScenarioToFeature(featureFname, m[walk[i]])
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	}

	fmt.Printf("Written to %s\n", featureFname)
	return nil
}

// Create a test story as a valid walk on components
// n: number of components
func MakeWalk(num_verts int) (error, []int) {

	// number of vertices = number of components +2 (dummy start, dummy end)
	n := num_verts + 2
	// graph G
	G := make([][]int, n)
	// initialize non-adjacencies to -1
	for i := 0; i < n; i++ {
		G[i] = make([]int, n)
		for j := 0; j < n; j++ {
			G[i][j] = -1
		}
	}

	// adjacencies
	G[0][0] = 1 // dummy start only to vtx 1
	G[1][0] = 2
	G[2][0] = 3
	G[3][0] = 2
	G[3][1] = 4
	G[4][0] = 2
	G[4][1] = 5 // dummy end is a sink

	fmt.Printf("%v\n", G)

	var walk = []int{}

	var vtx = 0
	for {
		fmt.Println("----------")
		num_adj_vtx := 0
		for k := 0; k < n; k++ {
			if G[vtx][k] > -1 {
				num_adj_vtx = num_adj_vtx + 1
			} else {
				break // after first encountered -1, all are -1's
			}
		}

		fmt.Printf("vtx = %d\nnum_adj_vtx = %v\n", vtx, num_adj_vtx)
		// if no adjacencies, we're done (nowhere to go)
		var next int
		switch num_adj_vtx {
		case 0:
			fmt.Printf("Walk: %v\n", walk)
			fmt.Println("Done.\n")
			return nil, walk
		case 1:
			next = 0
		default:
			rand.Seed(time.Now().UnixNano()) // want to seed with diff value each time
			next = rand.Intn(num_adj_vtx)    // make each adjacent vertex equally likely
		}
		walk = append(walk, vtx) // extend walk with vertex
		fmt.Printf("next = %d\t", next)
		vtx = G[vtx][next]
		fmt.Printf("going to %d next\n", vtx)

		fmt.Printf("Length: %d\nWalk: %v\n", len(walk), walk)

	}

	return nil, walk
}

// create Feature file
func main() {

	err := MakeIntegration()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
