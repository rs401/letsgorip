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
  private threadSubject: BehaviorSubject<Thread>;
  public thread: Observable<Thread>

  constructor(private http: HttpClient) {
    this.forumSubject = new BehaviorSubject<Forum>(
      JSON.parse(localStorage.getItem('currentForum') || '{}')
    );
    this.forum = this.forumSubject.asObservable();
    this.threadSubject = new BehaviorSubject<Thread>(
      JSON.parse(localStorage.getItem('currentThread') || '{}')
    );
    this.thread = this.threadSubject.asObservable();
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
        localStorage.setItem('currentForum', JSON.stringify(f));
        this.forumSubject.next(f);
      })
    );
  }

  getThread(fid: number, tid: number) {
    return this.http.get<Thread>(
      `${this.ROOT_URL}/api/forum/${fid}/thread/${tid}/`,
      {observe: "response"}
    ).pipe(
      map((data) => {
        let t: Thread = data.body as Thread;
        localStorage.setItem('currentThread', JSON.stringify(t));
        this.threadSubject.next(t);
      })
    );
  }

  createThread(thread: Thread) {
    return this.http.post<Thread>(
      `${this.ROOT_URL}/api/forum/${thread.forum_id}/thread/`,
      JSON.stringify(thread),
    );
  }

}
