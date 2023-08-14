package main

import (
	"fmt"
	"net/http"
)

var capitals = map[string]struct {
	Lat         float64
	Lng         float64
	Description string
}{
	"prague": {Lat: 50.0755, Lng: 14.4378, Description: "Prague is the capital and largest city of the Czech Republic."},
	"paris":  {Lat: 48.8566, Lng: 2.3522, Description: "Paris is the capital and most populous city of France."},
	"berlin": {Lat: 52.5200, Lng: 13.4050, Description: "Berlin is the capital and largest city of Germany."},
	"rome":   {Lat: 41.9028, Lng: 12.4964, Description: "Rome is the capital city of Italy and a special comune."},
	// Add more capital cities here with their coordinates
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		html := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>City Map</title>
			<script src="https://cdn.jsdelivr.net/npm/leaflet@1.7.1/dist/leaflet.js"></script>
			<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/leaflet@1.7.1/dist/leaflet.css" />
			<style>
				#map {
					height: 500px;
				}
			</style>
		</head>
		<body>
			<div>
				<label for="citySelect">Select a capital city:</label>
				<select id="citySelect">
					<option value="prague">Prague</option>
					<option value="paris">Paris</option>
					<option value="berlin">Berlin</option>
					<option value="rome">Rome</option>
					<!-- Add more capital cities here -->
				</select>
			</div>
			<div id="map"></div>
			<div id="cityInfo">
				<h2>City Information</h2>
				<p id="population"></p>
				<p id="area"></p>
				<p id="description"></p>
			</div>
			<script>
				var map = L.map('map').setView([50.0755, 14.4378], 13);
				L.tileLayer('https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png', {
					attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
				}).addTo(map);

				var marker = L.marker([50.0755, 14.4378]).addTo(map);
				marker.bindPopup("Prague").openPopup();

				var citySelect = document.getElementById('citySelect');
				citySelect.addEventListener('change', function() {
					var selectedCity = citySelect.value;
					var city = capitals[selectedCity];
					map.setView([city.Lat, city.Lng], 13);
					marker.setLatLng([city.Lat, city.Lng]);
					marker.bindPopup(selectedCity.charAt(0).toUpperCase() + selectedCity.slice(1)).openPopup();
					updateCityInfo(city.Description);
				});

				function updateCityInfo(description) {
					document.getElementById("description").innerHTML = description;
				}
			</script>
		</body>
		</html>
		`
		fmt.Fprintf(w, html)
	})

	fmt.Println("Server is running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
