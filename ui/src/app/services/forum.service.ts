import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { BehaviorSubject, map, Observable } from 'rxjs';
import { environment } from 'src/environments/environment';
import { Forum } from '../models/forum';
import { Thread } from '../models/thread';

@Injectable({
  providedIn: 'root'
})
export class ForumService {

  readonly ROOT_URL = environment.root_url;
  private forumSubject: BehaviorSubject<Forum>;
  public forum: Observable<Forum>

  constructor(private http: HttpClient) {
    this.forumSubject = new BehaviorSubject<Forum>(
      JSON.parse(localStorage.getItem('currentForum') || '{}')
    );
    this.forum = this.forumSubject.asObservable();
  }

  getForums() {
    return this.http.get<Forum[]>(`${this.ROOT_URL}/api/forum/`);
  }

  getThreads(id: number) {
    return this.http.get<Thread[]>(`${this.ROOT_URL}/api/forum/${id}/thread/`);
  }

  getForum(id: number) {
    return this.http.get<Forum>(
      `${this.ROOT_URL}/api/forum/${id}/`,
      {observe: "response"}
    ).pipe(
      map((data) => {
        let f: Forum = data.body as Forum;
        console.log('data:' + data.body)
        localStorage.setItem('currentForum', JSON.stringify(f));
        this.forumSubject.next(f);
      })
    );
  }

}
