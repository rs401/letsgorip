import { Component, OnInit, Input } from '@angular/core';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { Router } from '@angular/router';
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

  constructor(
    private auth: AuthService,
    private router: Router,
    ) { }

  ngOnInit(): void {
  }

  signin(){
    const email = this.emailControl.value;
    const password = this.passwordControl.value;
    if(email && password) {
      this.auth.signin(email, password).subscribe({
        next: (res) => {
          console.log('response: ' + res);
          this.router.navigateByUrl('/');
        },
        error: (err) => {this.showFlashMessage(err)},
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
