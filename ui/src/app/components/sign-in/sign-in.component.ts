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

  constructor(
    private auth: AuthService,
    private route: ActivatedRoute,
    private router: Router,
    ) { }

  ngOnInit(): void {
    this.returnURL = this.route.snapshot.queryParams['returnUrl'] || '/';
  }

  signin(){
    const email = this.emailControl.value;
    const password = this.passwordControl.value;
    if(email && password) {
      this.auth.signin(email, password).subscribe({
        next: (res) => {
          this.router.navigateByUrl(this.returnURL);
        },
        error: (err) => {this.showFlashMessage(err.error.error)},
        complete: () => console.info('complete')
      });
    } else {
      this.showFlashMessage("Email and Password cannot be empty.");
    }
  }

  showFlashMessage(msg: string) {
    this.signInMessage = msg;
  }

}
