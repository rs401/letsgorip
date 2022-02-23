import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { SignIn, User } from '../models/user';
import { BehaviorSubject, catchError, map, Observable, pipe, tap } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  private userSubject: BehaviorSubject<User>;
  public user: Observable<User>;
  public token: string;
  readonly ROOT_URL;

  constructor(private http: HttpClient) {
    this.userSubject = new BehaviorSubject<User>(new User);
    this.user = this.userSubject.asObservable();
    this.token = JSON.parse(localStorage.getItem('lgr-token') || '{}');
    this.ROOT_URL = 'http://192.168.49.2:32410';
  }

  // public getUser(): User {
  //   // return this.currentUser;
  //   // return this.userSubject.value;
  // }

  /* Need to set token after signin or signup in localstorage and make a way to pull it on header component init. 
   * Need to revisit this after I gett some other endpoints going.
   */

  signin(email: string, password: string) {
    return this.http.post<User>(`${this.ROOT_URL}/api/signin/`, { email: email, password: password })
      .pipe(
        map((data) => {
          let user: User = data;//{
          //   id: data.id,
          //   name: data.name,
          //   email: data.email,
          // };
          this.userSubject.next(user);
          console.log(user);
        })
      );

    // .subscribe(
    //   (response) => {
    //     console.log('SignIn response: ' + response);
    //   }
    // );
  }

  signup(name: string, email: string, password: string) {
    console.log('Calling endpoint');
    // this.http.post<User>(`${this.ROOT_URL}/api/signup/`, { name:name, email:email, password:password })
    return this.http.post<User>(this.ROOT_URL + '/api/signup/', { name:name, email:email, password:password })
      .pipe(
        map((data) => {
          let user: User = data;
          // localStorage.setItem(user)
          this.userSubject.next(user);
          console.log(user);
          // return user;
        }),
        catchError((err) => {return err;})
      );

    // .subscribe(
    //   (response) => {
    //     console.log('SignUp response: ' + response);
    //   }
    // );
  }

}
