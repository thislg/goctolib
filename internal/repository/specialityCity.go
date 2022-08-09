package repository

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SearchFiltersContext struct {
	Label     string `json:"label"`
	ParamName string `json:"param_name"`
	Items     []struct {
		Name string      `json:"name"`
		ID   interface{} `json:"id"`
	} `json:"items"`
	Value interface{} `json:"value"`
	Icon  string      `json:"icon"`
	Order int         `json:"order"`
}

type SpecialityCityResponse struct {
	Data struct {
		Speciality struct {
			ID                            int    `json:"id"`
			Name                          string `json:"name"`
			Slug                          string `json:"slug"`
			SearchResultsPriority         bool   `json:"search_results_priority"`
			TelehealthCore                bool   `json:"telehealth_core"`
			Vaccination                   bool   `json:"vaccination"`
			CovidVaccinationHcpSpeciality bool   `json:"covid_vaccination_hcp_speciality"`
		} `json:"speciality"`
		Doctors []struct {
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
			ExactMatch         bool        `json:"exact_match"`
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
			PlaceID         interface{} `json:"place_id"`
			Telehealth      bool        `json:"telehealth"`
			TopSpecialities []string    `json:"top_specialities,omitempty"`
			Distance        string      `json:"distance,omitempty"`
		} `json:"doctors"`
		DirectoryDoctors []struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Address   string `json:"address"`
			City      string `json:"city"`
			Link      string `json:"link"`
			Position  struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"position"`
			NameWithTitle      string      `json:"name_with_title"`
			Speciality         string      `json:"speciality"`
			IsDirectory        bool        `json:"is_directory"`
			Kind               string      `json:"kind"`
			OrganizationStatus interface{} `json:"organization_status"`
			Zipcode            string      `json:"zipcode"`
			CloudinaryPublicID interface{} `json:"cloudinary_public_id"`
		} `json:"directory_doctors"`
		SearchFiltersContext []SearchFiltersContext `json:"search_filters_context"`
		Place                struct {
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Slug     string `json:"slug"`
			PlaceID  string `json:"place_id"`
			Position struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"position"`
		} `json:"place"`
		DirectoryTitle       string      `json:"directory_title"`
		DoctorTitle          string      `json:"doctor_title"`
		SecondaryDoctorTitle string      `json:"secondary_doctor_title"`
		TitlesByPriority     interface{} `json:"titles_by_priority"`
		HasMore              bool        `json:"has_more"`
		SearchScope          string      `json:"search_scope"`
		IsHub                bool        `json:"is_hub"`
		TotalSearchResults   int         `json:"total_search_results"`
		Type                 string      `json:"type"`
	} `json:"data"`
	Seo struct {
		AboutDoctolib struct {
			Title string `json:"title"`
			Links []struct {
				Name    string `json:"name"`
				URL     string `json:"url"`
				Encoded bool   `json:"encoded"`
				Target  string `json:"target"`
				Key     string `json:"key"`
			} `json:"links"`
		} `json:"about_doctolib"`
		AboutDoctolibPro struct {
			Title string `json:"title"`
			Links []struct {
				Name    string `json:"name"`
				URL     string `json:"url"`
				Target  string `json:"target"`
				Encoded bool   `json:"encoded"`
				Key     string `json:"key"`
			} `json:"links"`
		} `json:"about_doctolib_pro"`
		PopularLinks struct {
			Title string `json:"title"`
			Links []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"links"`
		} `json:"popular_links"`
		NeighborhoodLinks struct {
			Title string        `json:"title"`
			Links []interface{} `json:"links"`
		} `json:"neighborhood_links"`
		SpecialityDirectoryLinks struct {
			Title    interface{} `json:"title"`
			MainLink interface{} `json:"main_link"`
			Links    interface{} `json:"links"`
		} `json:"speciality_directory_links"`
		SublocalityDirectoryLinks struct {
			Title    interface{} `json:"title"`
			MainLink interface{} `json:"main_link"`
			Links    interface{} `json:"links"`
		} `json:"sublocality_directory_links"`
		SimilarLinks struct {
			Title string `json:"title"`
			Links []struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"links"`
		} `json:"similar_links"`
		PublicContent struct {
			Title string      `json:"title"`
			Text  interface{} `json:"text"`
		} `json:"public_content"`
	} `json:"seo"`
	Meta struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		PageType    string `json:"page_type"`
	} `json:"meta"`
	Jsonld struct {
		PlaceholderImage struct {
			Context            string `json:"@context"`
			Type               string `json:"@type"`
			Image              string `json:"image"`
			PrimaryImageOfPage struct {
				Context     string `json:"@context"`
				Type        string `json:"@type"`
				ContentURL  string `json:"contentUrl"`
				Description string `json:"description"`
			} `json:"primaryImageOfPage"`
		} `json:"placeholder_image"`
		Doctors []struct {
			Context          string      `json:"@context"`
			Type             string      `json:"@type"`
			Name             string      `json:"name"`
			MedicalSpecialty string      `json:"medicalSpecialty"`
			LegalName        interface{} `json:"legalName"`
			URL              string      `json:"url"`
			Address          struct {
				Type            string `json:"@type"`
				Name            string `json:"name"`
				StreetAddress   string `json:"streetAddress"`
				PostalCode      string `json:"postalCode"`
				AddressLocality string `json:"addressLocality"`
			} `json:"address"`
			PaymentAccepted string      `json:"paymentAccepted"`
			Description     interface{} `json:"description,omitempty"`
		} `json:"doctors"`
		Breadcrumbs struct {
			Context         string `json:"@context"`
			Type            string `json:"@type"`
			ItemListElement []struct {
				Type     string `json:"@type"`
				Position int    `json:"position"`
				Item     struct {
					Type string      `json:"@type"`
					Name string      `json:"name"`
					ID   string      `json:"id"`
					URL  interface{} `json:"url"`
				} `json:"item"`
			} `json:"itemListElement"`
		} `json:"breadcrumbs"`
	} `json:"jsonld"`
}

func GetSpecialityCityResponse(baseUrl string, verbose bool, specialty string, city string, motive string) SpecialityCityResponse {
	var url string = baseUrl + "/" + specialty + "/" + city + ".json"

	if motive != "" {
		url += "?ref_visit_motive_id=" + motive
	}

	var resp *http.Response = GetResponse(url, "GET", verbose)
	defer resp.Body.Close()
	var cityResponse SpecialityCityResponse
	var err = json.NewDecoder(resp.Body).Decode(&cityResponse)

	if err != nil {
		panic(err)
	}

	if verbose {
		fmt.Println("Got speciality: " + cityResponse.Data.Speciality.Name)
	}

	return cityResponse
}
