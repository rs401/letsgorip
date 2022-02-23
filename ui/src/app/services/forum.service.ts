import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Forum } from '../models/forum';

@Injectable({
  providedIn: 'root'
})
export class ForumService {

  readonly ROOT_URL;

  constructor(private http: HttpClient) {
    this.ROOT_URL = 'http://192.168.49.2:32410';
  }

  getForums() {
    return this.http.get<Forum[]>(`${this.ROOT_URL}/api/forum/`);
  }

}
