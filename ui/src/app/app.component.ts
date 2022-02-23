import { Component } from '@angular/core';
import { User } from './models/user';
import { AuthService } from './services/auth.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  
  public currentUser?: User;

  constructor(private auth: AuthService) {
    this.auth.user.subscribe( user => this.currentUser = user);
  }
}
