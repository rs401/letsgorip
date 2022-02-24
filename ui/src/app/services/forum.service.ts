import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { Forum } from '../models/forum';

@Injectable({
  providedIn: 'root'
})
export class ForumService {

  readonly ROOT_URL = environment.root_url;

  constructor(private http: HttpClient) {
  }

  getForums() {
    return this.http.get<Forum[]>(`${this.ROOT_URL}/api/forum/`);
  }

}
