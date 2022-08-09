package main

import (
	"flag"
	"fmt"
	"golang.org/x/exp/slices"
	"internal/repository"
	"os/exec"
)

func main() {
	var baseUrl string = "https://www.doctolib.fr"
	cityPtr := flag.String("city", "", "Required city slug to search in. Example: paris")
	specialityPtr := flag.String("speciality", "", "Required specialty slug to search in. Example: medecin-generaliste")
	refVisitMotiveIdPtr := flag.String("motive", "", "Optional motive id to search in. Use verbose -v to get a list of motives. Example: 2167")
	verbosePtr := flag.Bool("v", false, "Verbose. Output more information.")
	notifyPtr := flag.Bool("notify", false, "Use notify-send to notify if there are results.")
	flag.Parse()

	if *cityPtr == "" || *specialityPtr == "" {
		flag.Usage()
		return
	}

	var specialityResponse repository.SpecialityCityResponse = repository.GetSpecialityCityResponse(baseUrl, *verbosePtr, *specialityPtr, *cityPtr, *refVisitMotiveIdPtr)

	if *refVisitMotiveIdPtr == "" && *verbosePtr {
		fmt.Println("No motive given. Get motive list...")

		idx := slices.IndexFunc(specialityResponse.Data.SearchFiltersContext, func(c repository.SearchFiltersContext) bool { return c.ParamName == "ref_visit_motive_id" })
		context := specialityResponse.Data.SearchFiltersContext[idx]

		fmt.Println("Valid motives ids are: ")

		for _, item := range context.Items {
			fmt.Println(item.ID, " - ", item.Name)
		}

		fmt.Println("Looking for availability for any motive...")
	}

	var searchResults []repository.SearchResultResponse = repository.GetSearchResultByCityResponse(baseUrl, *verbosePtr, specialityResponse, *refVisitMotiveIdPtr)
	var hasResults bool = false

	for _, searchResult := range searchResults {
		if searchResult.Total > 0 {
			fmt.Println("Available: " + searchResult.SearchResult.NameWithTitle + " " + baseUrl + searchResult.SearchResult.URL)
			hasResults = true
		}

		if !searchResult.NextSlot.IsZero() {
			fmt.Println("Next slot: " + searchResult.SearchResult.NameWithTitle + " " + searchResult.NextSlot.String() + " " + baseUrl + searchResult.SearchResult.URL)
			hasResults = true
		}
	}

	if *notifyPtr && hasResults {
		cmd := exec.Command("notify-send", "-a", "Goctolib", "--expire-time", "100000", "There are results to your appointment ("+*specialityPtr+" in "+*cityPtr+") search!")

		_, err := cmd.Output()

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
