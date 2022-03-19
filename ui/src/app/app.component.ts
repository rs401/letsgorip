import { Component } from '@angular/core';
import { Title } from '@angular/platform-browser';
import { User } from './models/user';
import { AuthService } from './services/auth.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  
  public currentUser?: User;
  public title: string = "Lets Go Rip!"

  constructor(private auth: AuthService, private titleService: Title) {
    this.auth.user.subscribe( user => this.currentUser = user);
  }

  ngOnInit(){
    this.titleService.setTitle(this.title);
  }
}
