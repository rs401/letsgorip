import { Component, OnInit, Input } from '@angular/core';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-sign-in',
  templateUrl: './sign-in.component.html',
  styleUrls: ['./sign-in.component.css']
})
export class SignInComponent implements OnInit {

  emailControl: FormControl = new FormControl('');
  passwordControl: FormControl = new FormControl('');
  signInMessage: string = '';
  returnURL!: string;
  callbackCreds: string | null = null;

  constructor(
    private auth: AuthService,
    private route: ActivatedRoute,
    private router: Router,
  ) {
    this.returnURL = this.route.snapshot.queryParams['returnUrl'] || '/';
    const routeParams = this.route.snapshot.paramMap;
    this.callbackCreds = routeParams.get('creds');
    if (this.callbackCreds != null && this.callbackCreds != undefined) {
      // unwrap the jwt, verify it, send it to back-end.
      this.auth.verifyToken(this.callbackCreds).subscribe({
        next: (response) => {
          this.gAuthSignIn(response);
        }
      });
    }
  }

  gAuthSignIn(response: Object) {
    try {
      this.auth.gAuthSignIn(response).subscribe({
        next:  (res) => {
          this.router.navigateByUrl(this.returnURL);
        },
        error: (err) => { this.showFlashMessage(err.error.error) },
        complete: () => console.info('gAuthSignIn complete')
      });
    } catch (error: any) {
      this.showFlashMessage(error);
    }
  }

  ngOnInit(): void { }

  signin() {
    const email = this.emailControl.value;
    const password = this.passwordControl.value;
    if (email && password) {
      this.auth.signin(email, password).subscribe({
        next: (res) => {
          this.router.navigateByUrl(this.returnURL);
        },
        error: (err) => { this.showFlashMessage(err.error.error) },
        complete: () => console.info('signin complete')
      });
    } else {
      this.showFlashMessage("Email and Password cannot be empty.");
    }
  }

  showFlashMessage(msg: string) {
    this.signInMessage = msg;
  }

}
