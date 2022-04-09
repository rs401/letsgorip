import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { User } from '../models/user';
import { BehaviorSubject, catchError, map, Observable, tap } from 'rxjs';
import { environment } from 'src/environments/environment';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private userSubject: BehaviorSubject<User>;
  public user: Observable<User>;
  public token: string;
  readonly ROOT_URL = environment.root_url;
  cu?: User;
  validToken: boolean = false;

  constructor(private http: HttpClient) {
    this.userSubject = new BehaviorSubject<User>(
      JSON.parse(localStorage.getItem('currentUser') || '{}')
    );
    this.user = this.userSubject.asObservable();
    this.token = localStorage.getItem('lgrToken') || '';
  }

  signin(email: string, password: string) {
    return this.http.post(
      `${this.ROOT_URL}/signin/`,
      { email: email, password: password },
      { observe: 'response', responseType: 'json', withCredentials: true }
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
    return this.http.post(
      `${this.ROOT_URL}/signup/`,
      { name: name, email: email, password: password },
      { observe: "response", withCredentials: true }
    ).pipe(
      tap((data) => {
        let user: User = data.body as User;
        this.token = String(data.headers.get('Authorization'));
        localStorage.setItem('lgrToken', data.headers.get('Authorization') || '');
        localStorage.setItem('currentUser', JSON.stringify(user));
        this.userSubject.next(user);
      })
    );
  }

  signout() {
    localStorage.removeItem('currentUser');
    localStorage.removeItem('lgrToken');
    this.token = '';
    this.userSubject.next(new User);
  }

  isLoggedIn(): boolean {
    this.user.subscribe(user => this.cu = user);
    if (this.cu?.id === undefined) {
      return false;
    }
    return true;
  }

  checkToken(): Observable<boolean> {
    return this.http.get(
      `${this.ROOT_URL}/user/checktoken/`,
      {
        headers: new HttpHeaders({
          Authorization: this.token,
        }),
        observe: 'response',
        responseType: 'json',
        withCredentials: true,
      },
    ).pipe(
      map(response => {
        console.log('AuthService >> checkToken response: ' + response.ok);
        if(response.ok) {
          return true;
        } else {
          this.signout();
          return false;
        }
      }),
      catchError((err) => {
        console.log('auth-service catching err, err.status: ' + err.status);
        this.signout();
        throw err;
      }),
    );
  }

  getUser(id: number): Observable<User> {
    // let tmpUser: User = new User;
    return this.http.get<User>(
      `${this.ROOT_URL}/user/${id}/`,
      { observe: "body", withCredentials: true }
    );
  }

  getUserByUid(uid: string): Observable<User> {
    // let tmpUser: User = new User;
    return this.http.get<User>(
      `${this.ROOT_URL}/user/uid/${uid}/`,
      { observe: "body", withCredentials: true }
    );
  }

  verifyToken(token: string): Observable<Object> {
    return this.http.get(
      `https://oauth2.googleapis.com/tokeninfo?id_token=${token}`,
      { observe: "body"}
      ).pipe();
  }
  
  gAuthSignIn(response: any) {
    // Check that the aud claim contains app client ID
    if(response.aud !== environment.gauth_client_id) {
      // Token is valid from Google but not from our app.
      console.log('Error. Valid token from wrong app: ' + JSON.stringify(response));
      throw new Error("invalid token audience");
    }
    // Post User signin
    return this.http.post<User>(
      `${this.ROOT_URL}/signin/`,
      {
        uid: response.sub,
        name: response.name,
        email: response.email,
        email_verified: Boolean(response.email_verified),
        picture: response.picture,
      },
      { observe: 'response', responseType: 'json', withCredentials: true }
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

}
