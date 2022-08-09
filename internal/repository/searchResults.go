package repository

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type SearchResultResponse struct {
	Availabilities []struct {
		Date  string        `json:"date"`
		Slots []interface{} `json:"slots"`
	} `json:"availabilities"`
	Total        int       `json:"total"`
	Reason       string    `json:"reason"`
	Message      string    `json:"message"`
	NextSlot     time.Time `json:"next_slot"`
	SearchResult struct {
		ID                 int         `json:"id"`
		IsDirectory        bool        `json:"is_directory"`
		Kind               string      `json:"kind"`
		Address            string      `json:"address"`
		City               string      `json:"city"`
		Zipcode            string      `json:"zipcode"`
		PracticeIds        []int       `json:"practice_ids"`
		Link               string      `json:"link"`
		CloudinaryPublicID string      `json:"cloudinary_public_id"`
		ProfileID          int         `json:"profile_id"`
		ExactMatch         interface{} `json:"exact_match"`
		PrioritySpeciality bool        `json:"priority_speciality"`
		FirstName          string      `json:"first_name"`
		LastName           string      `json:"last_name"`
		NameWithTitle      string      `json:"name_with_title"`
		Speciality         string      `json:"speciality"`
		OrganizationStatus interface{} `json:"organization_status"`
		RegulationSector   string      `json:"regulation_sector"`
		Position           struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"position"`
		PlaceID                    interface{} `json:"place_id"`
		Telehealth                 bool        `json:"telehealth"`
		VisitMotiveType            interface{} `json:"visit_motive_type"`
		AgendaIds                  []int       `json:"agenda_ids"`
		LandlineNumber             string      `json:"landline_number"`
		BookingTemporaryDisabled   bool        `json:"booking_temporary_disabled"`
		ResetVisitMotive           interface{} `json:"resetVisitMotive"`
		ToFinalizeStep             interface{} `json:"toFinalizeStep"`
		ToFinalizeStepWithoutState interface{} `json:"toFinalizeStepWithoutState"`
		URL                        string      `json:"url"`
	} `json:"search_result"`
}

func GetSearchResultResponse(baseUrl string, verbose bool, doctorID int, refVisitMotiveID string) SearchResultResponse {

	var limit string = "1"
	var url string = baseUrl + "/search_results/" + strconv.Itoa(doctorID) + ".json?limit=" + limit

	if refVisitMotiveID != "" {
		url += "&ref_visit_motive_id=" + refVisitMotiveID
	}

	var resp *http.Response = GetResponse(url, "GET", verbose)
	defer resp.Body.Close()
	var responseObject SearchResultResponse
	var err = json.NewDecoder(resp.Body).Decode(&responseObject)

	if err != nil {
		panic(err)
	}

	return responseObject
}

func GetSearchResultByCityResponse(baseUrl string, verbose bool, cityResponse SpecialityCityResponse, refVisitMotiveID string) []SearchResultResponse {
	var responses []SearchResultResponse

	for _, doctor := range cityResponse.Data.Doctors {
		responses = append(responses, GetSearchResultResponse(baseUrl, verbose, doctor.ID, refVisitMotiveID))
	}

	return responses
}
