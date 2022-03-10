import { Component, Input, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { User } from 'src/app/models/user';
import { AuthService } from 'src/app/services/auth.service';

@Component({
  selector: 'app-user-card',
  templateUrl: './user-card.component.html',
  styleUrls: ['./user-card.component.css']
})
export class UserCardComponent implements OnInit {
  
  @Input() uid?: number;
  public user?: Observable<User>;

  constructor(private auth: AuthService) { }

  ngOnInit(): void {
    this.user = this.auth.getUser(this.uid!);
  }

}
