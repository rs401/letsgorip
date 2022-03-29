import { Component, Input, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { Router } from '@angular/router';
import { User } from 'src/app/models/user';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {

  @Input() user?: User;
  public query: FormControl = new FormControl('');

  constructor(private auth: AuthService, private router: Router) { }
  
  ngOnInit(): void {
  }
  
  signOut() {
    this.auth.signout();
    window.location.reload();
  }

  search() {
    this.router.navigateByUrl(`/search/${this.query.value}`)
  }

}
