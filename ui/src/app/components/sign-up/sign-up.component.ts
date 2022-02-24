import { Component, OnInit } from '@angular/core';
import { FormControl, ReactiveFormsModule } from '@angular/forms';
import { Router } from '@angular/router';
import { first } from 'rxjs';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-sign-up',
  templateUrl: './sign-up.component.html',
  styleUrls: ['./sign-up.component.css']
})
export class SignUpComponent implements OnInit {

  nameControl: FormControl = new FormControl('');
  emailControl: FormControl = new FormControl('');
  passwordControl: FormControl = new FormControl('');
  password2Control: FormControl = new FormControl('');
  signUpMessage: string = '';

  constructor(
    private auth: AuthService,
    private router: Router,
    ) { }

  ngOnInit(): void {
  }

  signup(){
    const name = this.nameControl.value;
    const email = this.emailControl.value;
    const password = this.passwordControl.value;
    const password2 = this.password2Control.value;
    if(password != password2) {
      this.passwordControl.setValue('');
      this.password2Control.setValue('');
      this.showFlashMessage('passwords do not match');
      return;
    }
    if(name && email && password) {
      this.auth.signup(name, email, password).subscribe({
        next: (res) => {
          console.log('response: ' + res);
          this.router.navigateByUrl('/');
        },
        error: (err) => {this.showFlashMessage(err.error.error)},
        complete: () => console.info('complete')
      });
    } else if(!name){
      this.showFlashMessage("Name cannot be empty.");
    } else if(!email){
      this.showFlashMessage("Email cannot be empty.");
    } else if(!password){
      this.showFlashMessage("Password cannot be empty.");
    }
  }

  showFlashMessage(msg: string) {
    this.signUpMessage = msg;
  }

}
