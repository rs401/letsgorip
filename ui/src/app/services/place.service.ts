import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { Place } from '../models/place';

@Injectable({
  providedIn: 'root'
})
export class PlaceService {

  readonly ROOT_URL = environment.root_url;

  constructor(private http: HttpClient) { }

  createPlace(token: string, place: Place) {
    return this.http.post<Place>(
      `${this.ROOT_URL}/place/`,
      JSON.stringify(place),
      { 
        headers: new HttpHeaders({
          'Content-Type':  'application/json',
          Authorization: token,
        }),
        observe: 'response', responseType: 'json', withCredentials: true }
    )
  }

  getPlaces() {
    return this.http.get<Place[]>(`${this.ROOT_URL}/place/`);
  }
}
