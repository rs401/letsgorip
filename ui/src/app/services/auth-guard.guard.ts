import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot, UrlTree } from '@angular/router';
import { Observable } from 'rxjs';
import { User } from '../models/user';
import { AuthService } from './auth.service';

@Injectable({
  providedIn: 'root'
})
export class AuthGuardGuard implements CanActivate {

  constructor(private auth: AuthService, private router: Router) {

  }

  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): Observable<boolean | UrlTree> | Promise<boolean | UrlTree> | boolean | UrlTree {
    console.log('AuthGuard isLoggedIn:' + this.auth.isLoggedIn());
    if (this.auth.isLoggedIn()) {
      let valid: boolean;
      this.auth.checkToken().subscribe({
        next: (val: boolean) => {
          if(val) {
            console.log('AuthGuard checkToken: ' + val);
            return true;
          } else {
            console.log('AuthGuard checkToken: ' + val);
            return false;
          }
        }
      });
      // return true;
    }
    this.router.navigate(['sign-in'], { queryParams: { returnUrl: state.url } });
    return false;
  }

}
