package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Biome struct {
	Id       int
	Name     string
	Passages []int
}

// list all biomes
var biomes = [...]Biome{
	{0, "Prisoners' Quarters", []int{1, 2, 3}},
	{1, "Dilapidated Arboretum", []int{4, 5, 6}},
	{2, "Promenade of the Condemned", []int{4, 5, 6, 7}},
	{3, "Toxic Sewers", []int{6, 7, 9}},
	{4, "Morass of the Banished", []int{10}},
	{5, "Prison Depths", []int{8}},
	{6, "The Ramparts", []int{11}},
	{7, "Corrupted Prison", []int{9}},
	{8, "Ossuary", []int{11}},
	{9, "Ancient Sewers", []int{12}},
	{10, "Nest", []int{13, 14, 16}},
	{11, "Black Bridge", []int{13, 14, 15}},
	{12, "Insufferable Crypt", []int{15, 16}},
	{13, "Fractured Shrines", []int{17, 18, 19}},
	{14, "Stilt Village", []int{18}},
	{15, "Slumbering Sanctuary", []int{19}},
	{16, "Graveyard", []int{19, 20}},
	{17, "Undying Shores", []int{}},
	{18, "Clock Tower", []int{21}},
	{19, "Forgotten Sepulcher", []int{21}},
	{20, "Cavern", []int{}},
	{21, "Clock Room", []int{22, 23}},
	{22, "Derelict Distillery", []int{}},
	{23, "High Peak Castle", []int{24}},
	{24, "Throne Room", []int{}},
}

var biomeIdToBiome = make(map[int]Biome)

func copySliceAndAddElement(biomes []Biome, biome Biome) []Biome {
	newVisitedBiomes := make([]Biome, len(biomes)+1)
	copy(newVisitedBiomes, biomes)
	newVisitedBiomes[len(newVisitedBiomes)-1] = biome
	return newVisitedBiomes
}

func findRoutesHelper(currentBiome Biome, targetBiome Biome, visitedBiomes []Biome, results *[][]Biome) {
	for _, passageBiomeId := range currentBiome.Passages {
		nextBiome := biomeIdToBiome[passageBiomeId]
		newVisitedBiomes := copySliceAndAddElement(visitedBiomes, currentBiome)
		findRoutesHelper(nextBiome, targetBiome, newVisitedBiomes, results)
	}
	if currentBiome.Id == targetBiome.Id {
		newVisitedBiomes := copySliceAndAddElement(visitedBiomes, currentBiome)
		*results = append(*results, newVisitedBiomes)
	}
}

func findRoutes(currentBiome Biome, targetBiome Biome) [][]Biome {
	var results [][]Biome
	findRoutesHelper(currentBiome, targetBiome, make([]Biome, 0), &results)
	return results
}

func printResults(results [][]Biome) {
	biomeIdToMaxIndex := make(map[int]int)
	maxPathLength := 0
	for _, path := range results {
		for i, biome := range path {
			if biomeIdToMaxIndex[biome.Id] < i {
				biomeIdToMaxIndex[biome.Id] = i
			}
		}
		if len(path) > maxPathLength {
			maxPathLength = len(path)
		}
	}
	//maxPathIndex := maxPathLength - 1

	indexToLength := make(map[int]int)

	for biomeId, index := range biomeIdToMaxIndex {
		biomeName := biomeIdToBiome[biomeId].Name
		biomeNameLength := len(biomeName)
		if indexToLength[index] < biomeNameLength {
			indexToLength[index] = biomeNameLength
		}
	}

	for _, path := range results {
		var currentIndex = 0
		fmt.Printf("(%2d)", len(path))
		for _, biome := range path {
			biomeIndex := biomeIdToMaxIndex[biome.Id]
			for biomeIndex > currentIndex {
				fmt.Printf("    %-"+strconv.Itoa(indexToLength[currentIndex])+"s", "")
				currentIndex++
			}
			fmt.Printf(" -> %-"+strconv.Itoa(indexToLength[currentIndex])+"s", biome.Name)
			currentIndex++
		}
		fmt.Println()
	}
}

func main() {
	var menuFlag = flag.Int("menu", 0, "only 1 is allowed")

	if *menuFlag < 0 || *menuFlag > 1 {
		fmt.Println("Invalid value of flag -menu")
		os.Exit(0)
	}

	// add biomes to a map by their ID
	for _, v := range biomes {
		biomeIdToBiome[v.Id] = v
	}

	// sort biomes alphabetically
	sort.Slice(biomes[:], func(i, j int) bool {
		return biomes[i].Name < biomes[j].Name
	})

	routes := findRoutes(biomeIdToBiome[0], biomeIdToBiome[24])
	printResults(routes)
}
