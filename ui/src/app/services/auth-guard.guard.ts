import { Injectable } from '@angular/core';
import { ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot, UrlTree } from '@angular/router';
import { map, Observable } from 'rxjs';
import { User } from '../models/user';
import { AuthService } from './auth.service';

@Injectable({
  providedIn: 'root'
})
export class AuthGuardGuard implements CanActivate {

  constructor(private auth: AuthService, private router: Router) { }

  canActivate(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot): boolean | Observable<boolean> {
    console.log('AuthGuard isLoggedIn:' + this.auth.isLoggedIn());
    if (this.auth.isLoggedIn()) {
      try {
        return this.auth.checkToken().pipe(
          map((valid) => {
            if(valid) {
              return true;
            } else {
              // token expired
              this.router.navigate(['sign-in'], { queryParams: { returnUrl: state.url } });
              return false;
            }
          })
        );
      } catch (err) {
        this.router.navigate(['sign-in'], { queryParams: { returnUrl: state.url } });
        return false;
      }
      // return true;
    }
    this.router.navigate(['sign-in'], { queryParams: { returnUrl: state.url } });
    return false;
  }

}
