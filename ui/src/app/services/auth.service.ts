import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { User } from '../models/user';
import { BehaviorSubject, catchError, Observable,  tap } from 'rxjs';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private userSubject: BehaviorSubject<User>;
  public user: Observable<User>;
  public token: string;
  readonly ROOT_URL = environment.root_url;

  constructor(private http: HttpClient) {
    this.userSubject = new BehaviorSubject<User>(
      JSON.parse(localStorage.getItem('currentUser') || '{}')
    );
    this.user = this.userSubject.asObservable();
    this.token = localStorage.getItem('lgrToken') || '';
  }

  signin(email: string, password: string) {
    return this.http.post(
      `${this.ROOT_URL}/api/signin/`,
      { email: email, password: password },
      {observe: 'response', responseType: 'json'}
      )
      .pipe(
        tap((data) => {
          let user: User = data.body as User;
          this.token = String(data.headers.get('Authorization'));
          localStorage.setItem('lgrToken', data.headers.get('Authorization') || '');
          localStorage.setItem('currentUser', JSON.stringify(user));
          this.userSubject.next(user);
        })
      );
  }

  signup(name: string, email: string, password: string) {
    console.log('Calling endpoint');
    return this.http.post(
      `${this.ROOT_URL}/api/signup/`,
      { name:name, email:email, password:password },
      {observe: "response"}
    )
      .pipe(
        tap((data) => {
          let user: User = data.body as User;
          this.token = String(data.headers.get('Authorization'));
          localStorage.setItem('lgrToken', data.headers.get('Authorization') || '');
          localStorage.setItem('currentUser', JSON.stringify(user));
          this.userSubject.next(user);
        }),
        catchError((err) => {return err;})
      );
  }

  signout() {
    localStorage.removeItem('currentUser');
    localStorage.removeItem('lgrToken');
    this.token = '';
    this.userSubject.next(new User);
  }

}
