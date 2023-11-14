package main

// swagger:meta
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

// CountriesListHandler retrieves a list of all countries.
// CountriesListHandler godoc
// @Summary Retrieve a list of all countries
// @Description Retrieve a list of all countries
// @Tags countries
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "JWT token"
// @Success 200 {object} map[string]interface{} "country data"
// @Router /countries [get]
func CountriesListHandler(w http.ResponseWriter, r *http.Request) {
	apiURL := "https://restcountries.com/v3.1/all"

	resp, err := http.Get(apiURL)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching country data: %s", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching country data. Status code: %d", resp.StatusCode))
		return
	}

	var countryData []map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&countryData)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("Error decoding country data: %s", err))

		// Print the raw response body for debugging
		body, _ := io.ReadAll(resp.Body)
		fmt.Fprintf(w, "Raw response body: %s", body)

		return
	}

	response := map[string]interface{}{"country": countryData}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// CountryDetailsHandler fetches detailed information about a specific country
// CountryDetailsHandler godoc
// @Summary Get detailed information about a specific country
// @Description Get detailed information about a specific country
// @Tags countries
// @Accept  json
// @Produce  json
// @Param name query string true "The name of the country to fetch."
// @Security ApiKeyAuth
// @Param Authorization header string true "JWT token"
// @Success 200 {object} map[string]interface{} "country data"
// @Failure 400 {object} ErrorResponse "Error: Please provide a valid country name in the query parameters"
// @Failure 500 {object} ErrorResponse "Error fetching country data or decoding country data"
// @Router /country [get]
func CountryDetailsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract country name from the query parameters
	countryName := r.URL.Query().Get("name")

	if countryName == "" {
		writeJSONError(w, http.StatusBadRequest, "Error: Please provide a country name in the query parameters (e.g., /country?name=India)")
		return
	}

	apiURL := fmt.Sprintf("https://restcountries.com/v3.1/name/%s", countryName)

	resp, err := http.Get(apiURL)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching country data: %s", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching country data. Status code: %d", resp.StatusCode))
		return
	}

	var countryData []map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&countryData)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("Error decoding country data: %s", err))

		// Print the raw response body for debugging
		body, _ := io.ReadAll(resp.Body)
		fmt.Fprintf(w, "Raw response body: %s", body)

		return
	}

	// Create a map with the desired structure
	responseMap := map[string]interface{}{"country_data": countryData}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseMap)
}

// CountriesFilterListHandler retrieves a filtered and sorted list of countries based on specified parameters.
// CountriesFilterListHandler godoc
// @Summary Get a filtered and sorted list of countries
// @Description Get a filtered and sorted list of countries based on specified parameters
// @Tags countries
// @Accept  json
// @Produce  json
// @Param population query integer false "Filter countries by population."
// @Param area query integer false "Filter countries by area."
// @Param language query string false "Filter countries by language."
// @Param sort query string false "Sort order (asc or desc)."
// @Param page query integer false "Page number for pagination."
// @Security ApiKeyAuth
// @Param Authorization header string true "JWT token"
// @Success 200 {object} map[string][]string "paginated list of countries"
// @Failure 400 {object} ErrorResponse "Invalid filter or page parameters"
// @Failure 500 {object} ErrorResponse "Error fetching or decoding countries data"
// @Router /countries/filter [get]
func CountriesFilterListHandler(w http.ResponseWriter, r *http.Request) {
	// Extract filter parameters from query
	populationFilterStr := r.URL.Query().Get("population")
	areaFilterStr := r.URL.Query().Get("area")
	languageFilter := r.URL.Query().Get("language")
	sortOrder := r.URL.Query().Get("sort") // "asc" or "desc"
	pageStr := r.URL.Query().Get("page")

	// Parse filter parameters
	var populationFilter, areaFilter, page int
	var err error

	if populationFilterStr != "" {
		populationFilter, err = strconv.Atoi(populationFilterStr)
		if err != nil {
			// w.WriteHeader(http.StatusBadRequest)
			// fmt.Fprint(w, "Error: Invalid population filter value")
			// return
			writeJSONError(w, http.StatusBadRequest, "Invalid population filter value")
			return
		}
	}

	if areaFilterStr != "" {
		areaFilter, err = strconv.Atoi(areaFilterStr)
		if err != nil {
			// w.WriteHeader(http.StatusBadRequest)
			// fmt.Fprint(w, "Error: Invalid area filter value")
			// return
			writeJSONError(w, http.StatusBadRequest, "Invalid area filter value")
			return
		}
	}

	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil || page <= 0 {
			// w.WriteHeader(http.StatusBadRequest)
			// fmt.Fprint(w, "Error: Invalid page number")
			// return
			writeJSONError(w, http.StatusBadRequest, "Invalid page number")
			return
		}
	} else {
		// Set default page value if not present
		page = 1
	}

	// Fetch all countries from the REST Countries API
	apiURL := "https://restcountries.com/v3.1/all"
	resp, err := http.Get(apiURL)
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		// fmt.Fprintf(w, "Error fetching countries data: %s", err)
		// return
		writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching countries data: %s", err))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// w.WriteHeader(http.StatusInternalServerError)
		// fmt.Fprintf(w, "Error fetching countries data. Status code: %d", resp.StatusCode)
		// return
		writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("Error fetching countries data. Status code: %d", resp.StatusCode))
		return
	}

	var countriesData []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&countriesData)
	if err != nil {
		// w.WriteHeader(http.StatusInternalServerError)
		// fmt.Fprintf(w, "Error decoding countries data: %s\n", err)
		writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("Error decoding countries data: %s", err))

		// Print the raw response body for debugging
		body, _ := io.ReadAll(resp.Body)
		// fmt.Fprintf(w, "Raw response body: %s", body)
		writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("Raw response body: %s", body))

		return
	}

	// Apply filters and sorting
	filteredCountries := filterAndSortCountries(countriesData, populationFilter, areaFilter, languageFilter, sortOrder)

	// Paginate the results
	startIndex := (page - 1) * 20
	// Ensure startIndex is within bounds
	if startIndex >= len(filteredCountries) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string][]string{"country": {}}) // Empty response for out-of-bounds page
		return
	}

	endIndex := startIndex + 20

	// Ensure endIndex is within bounds
	if endIndex > len(filteredCountries) {
		endIndex = len(filteredCountries)
	}

	// Extract country names
	var countryNames []string
	for _, country := range filteredCountries[startIndex:endIndex] {
		name, ok := country["name"].(map[string]interface{})["common"].(string)
		if ok {
			countryNames = append(countryNames, name)
		}
	}

	response := map[string][]string{"country": countryNames}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func writeJSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	errorResponse := ErrorResponse{
		Error: message,
	}

	json.NewEncoder(w).Encode(errorResponse)
}

// filterAndSortCountries filters and sorts the countries based on the specified parameters
func filterAndSortCountries(countries []map[string]interface{}, populationFilter, areaFilter int, languageFilter, sortOrder string) []map[string]interface{} {
	var filteredCountries []map[string]interface{}

	for _, country := range countries {
		population, popOk := country["population"].(float64)
		area, areaOk := country["area"].(float64)
		languages, langOk := country["languages"].(map[string]interface{})

		// Apply filters
		populationCondition := !popOk || (populationFilter <= 0 || int(population) > populationFilter)
		areaCondition := !areaOk || (areaFilter <= 0 || int(area) == areaFilter)
		languageCondition := !langOk || (languageFilter == "" || languages[languageFilter] != nil)

		// Include the country in filteredCountries only if all conditions are met
		if populationCondition && areaCondition && languageCondition {
			filteredCountries = append(filteredCountries, country)
		}
	}

	// Sort the filtered countries based on population in ascending or descending order
	if sortOrder == "desc" {
		sortCountriesDesc(filteredCountries)
	} else {
		sortCountriesAsc(filteredCountries)
	}

	return filteredCountries
}

// sortCountriesAsc sorts countries based on population in ascending order
// func sortCountriesAsc(countries []map[string]interface{}) {
// 	for i := 0; i < len(countries)-1; i++ {
// 		for j := i + 1; j < len(countries); j++ {
// 			population1, _ := countries[i]["population"].(float64)
// 			population2, _ := countries[j]["population"].(float64)

// 			if population1 > population2 {
// 				// Swap countries
// 				countries[i], countries[j] = countries[j], countries[i]
// 			}
// 		}
// 	}
// }

// // sortCountriesDesc sorts countries based on population in descending order
// func sortCountriesDesc(countries []map[string]interface{}) {
// 	for i := 0; i < len(countries)-1; i++ {
// 		for j := i + 1; j < len(countries); j++ {
// 			population1, _ := countries[i]["population"].(float64)
// 			population2, _ := countries[j]["population"].(float64)

// 			if population1 < population2 {
// 				// Swap countries
// 				countries[i], countries[j] = countries[j], countries[i]
// 			}
// 		}
// 	}
// }

// sortCountriesAsc sorts countries based on Name in ascending order
func sortCountriesAsc(countries []map[string]interface{}) {
	sort.Slice(countries, func(i, j int) bool {
		name1, _ := countries[i]["name"].(map[string]interface{})["common"].(string)
		name2, _ := countries[j]["name"].(map[string]interface{})["common"].(string)
		return strings.ToLower(name1) < strings.ToLower(name2)
	})
}

// sortCountriesDesc sorts countries based on Name in descending order
func sortCountriesDesc(countries []map[string]interface{}) {
	sort.Slice(countries, func(i, j int) bool {
		name1, _ := countries[i]["name"].(map[string]interface{})["common"].(string)
		name2, _ := countries[j]["name"].(map[string]interface{})["common"].(string)
		return strings.ToLower(name1) > strings.ToLower(name2)
	})
}
