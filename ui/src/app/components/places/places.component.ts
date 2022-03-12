import { HttpClient } from '@angular/common/http';
import { Component, OnInit, ViewChild } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { GoogleMap, MapInfoWindow, MapMarker } from '@angular/google-maps';
import { catchError, map, Observable, of } from 'rxjs';
import { Place } from 'src/app/models/place';
import { User } from 'src/app/models/user';
import { AuthService } from 'src/app/services/auth.service';
import { environment } from 'src/environments/environment';

@Component({
  selector: 'app-places',
  templateUrl: './places.component.html',
  styleUrls: ['./places.component.css']
})
export class PlacesComponent implements OnInit {
  readonly MAP_API_KEY = environment.map_api_key;

  currentUser?: User;
  states: string[] = ['Alabama', 'Alaska', 'Arizona', 'Arkansas', 'California', 'Colorado', 'Connecticut', 'Delaware', 'Florida', 'Georgia', 'Hawaii', 'Idaho', 'Illinois', 'Indiana', 'Iowa', 'Kansas', 'Kentucky', 'Louisiana', 'Maine', 'Maryland', 'Massachusetts', 'Michigan', 'Minnesota', 'Mississippi', 'Missouri', 'Montana', 'Nebraska', 'Nevada', 'New Hampshire', 'New Jersey', 'New Mexico', 'New York', 'North Carolina', 'North Dakota', 'Ohio', 'Oklahoma', 'Oregon', 'Pennsylvania', 'Rhode Island', 'South Carolina', 'South Dakota', 'Tennessee', 'Texas', 'Utah', 'Vermont', 'Virginia', 'Washington', 'West Virginia', 'Wisconsin', 'Wyoming'];
  stateForm: FormGroup;
  newPlaceForm: FormGroup = this.fb.group({});
  name: FormControl = new FormControl('', Validators.required);
  description: FormControl = new FormControl('', Validators.required);
  lat: FormControl = new FormControl('');
  lng: FormControl = new FormControl('');

  stateCenters: google.maps.LatLngLiteral[] = [{lat:32.318230,lng:-86.902298},
    {lat:66.160507,lng:-153.369141},
    {lat:34.048927,lng:-111.093735},
    {lat:34.799999,lng:-92.199997},
    {lat:36.778259,lng:-119.417931},
    {lat:39.113014,lng:-105.358887},
    {lat:41.599998,lng:-72.699997},
    {lat:39.000000,lng:-75.500000},
    {lat:27.994402,lng:-81.760254},
    {lat:33.247875,lng:-83.441162},
    {lat:19.741755,lng:-155.844437},
    {lat:44.068203,lng:-114.742043},
    {lat:40.000000,lng:-89.000000},
    {lat:40.273502,lng:-86.126976},
    {lat:42.032974,lng:-93.581543},
    {lat:38.500000,lng:-98.000000},
    {lat:37.839333,lng:-84.270020},
    {lat:30.391830,lng:-92.329102},
    {lat:45.367584,lng:-68.972168},
    {lat:39.045753,lng:-76.641273},
    {lat:42.407211,lng:-71.382439},
    {lat:44.182205,lng:-84.506836},
    {lat:46.392410,lng:-94.636230},
    {lat:33.000000,lng:-90.000000},
    {lat:38.573936,lng:-92.603760},
    {lat:46.965260,lng:-109.533691},
    {lat:41.500000,lng:-100.000000},
    {lat:39.876019,lng:-117.224121},
    {lat:44.000000,lng:-71.500000},
    {lat:39.833851,lng:-74.871826},
    {lat:34.307144,lng:-106.018066},
    {lat:43.000000,lng:-75.000000},
    {lat:35.782169,lng:-80.793457},
    {lat:47.650589,lng:-100.437012},
    {lat:40.367474,lng:-82.996216},
    {lat:36.084621,lng:-96.921387},
    {lat:44.000000,lng:-120.500000},
    {lat:41.203323,lng:-77.194527},
    {lat:41.742325,lng:-71.742332},
    {lat:33.836082,lng:-81.163727},
    {lat:44.500000,lng:-100.000000},
    {lat:35.860119,lng:-86.660156},
    {lat:31.000000,lng:-100.000000},
    {lat:39.419220,lng:-111.950684},
    {lat:44.000000,lng:-72.699997},
    {lat:37.926868,lng:-78.024902},
    {lat:47.751076,lng:-120.740135},
    {lat:39.000000,lng:-80.500000},
    {lat:44.500000,lng:-89.500000},
    {lat:43.075970,lng:-107.290283},
    ];

  mapLoading: boolean = false;
  mapLoaded: boolean = false;
  addingPlace: boolean = false;
  authenticated: boolean = this.auth.isLoggedIn();
  markers: Marker[] = [];
  places: Place[] = [];

  zoom: number = 4;
  center: google.maps.LatLngLiteral;
  options: google.maps.MapOptions = {
    mapTypeId: 'roadmap',
    zoomControl: true,
    scrollwheel: true,
    disableDoubleClickZoom: true,
    maxZoom: 15,
    minZoom: 2,
  };

  iwContent: string = '';

  @ViewChild(GoogleMap) map!: GoogleMap;
  @ViewChild(MapInfoWindow) infoWindow!: MapInfoWindow;

  constructor(private fb: FormBuilder, private auth: AuthService) {
    this.auth.user.subscribe( user => this.currentUser = user);
    this.center = {
      lat: 39.8097343,
      lng: -98.5556199,
    };
    this.stateForm = this.fb.group({
      state: [0]
    });
  }

  ngOnInit(): void {
    this.loadMap();
  }

  loadMap() {
    if (this.mapLoaded || this.mapLoading) {
      return;
    }
    
    // One way of doing this: dynamically load a script tag.
    this.mapLoading = true;    
    const mapsScript = document.createElement('script')
    mapsScript.setAttribute('async', '');
    mapsScript.src = `https://maps.googleapis.com/maps/api/js?key=${this.MAP_API_KEY}`;
    mapsScript.addEventListener('load', () => {
      this.mapLoaded = true;
      this.mapLoading = false;
    })
    document.head.appendChild(mapsScript); 
  }

  update(){
    this.map.panTo(this.stateCenters[this.stateForm.value.state])
    // this.center = this.stateCenters[this.stateForm.value.state];
    this.zoom = 8;
  }

  addMarker(event: google.maps.MapMouseEvent) {
    this.addingPlace = true;
    this.lat.setValue(event.latLng?.lat());
    this.lng.setValue(event.latLng?.lng());
  }

  createPlace() {
    // Need to create placeService to create and get places
    let place: Place = {
      user_id: this.currentUser?.id,
      name: this.name.value,
      description: this.description.value,
      latitude: this.lat.value,
      longitude: this.lng.value
    };
    this.places.push(place);
    // For now just going to add places to local array as markers.
    let marker: Marker = {
      options: {
        title: place.name
      },
      position: {
        lat: place.latitude!,
        lng: place.longitude!
      },
      name: place.name!,
      description: place.description!
    };

    this.markers.push(marker);
    this.cancelCreatePlace();
  }

  cancelCreatePlace() {
    this.addingPlace = false;
    this.name.setValue('');
    this.description.setValue('');
  }

  showInfo(marker: Marker, mapmarker: MapMarker) {
    this.iwContent = `
    <h3>${marker.name}</h3>
    <hr />
    <p>${marker.description}</p>
    `;
    this.infoWindow.open(mapmarker);
  }

}// END PlacesComponent class

interface Marker {
  options: google.maps.MarkerOptions,
  position: google.maps.LatLngLiteral,
  name: string,
  description: string
}
