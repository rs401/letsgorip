import { Component, Input, OnInit } from '@angular/core';
import { User } from 'src/app/models/user';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {

  @Input() user?: User;

  constructor(private auth: AuthService) { 
    // this.auth.user.subscribe( user => this.user = user);
  }
  
  ngOnInit(): void {
  }
  
  // isEmptyObject(obj: Object) {
  //   return (obj && (Object.keys(obj).length === 0));
  // }

}
