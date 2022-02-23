import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute, ParamMap } from '@angular/router';
import { Forum } from 'src/app/models/forum';
import { ForumService } from 'src/app/services/forum.service';

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css']
})
export class MainComponent implements OnInit {

  public forums: Forum[] = [];

  constructor(private forumService: ForumService) {

  }

  ngOnInit(): void {
    this.forumService.getForums().subscribe({
      next: (res) => {
        res.forEach((forum) => {
          this.forums.push(forum);
        });
      },
      error: (err) => {console.log('Error: ' + err);},
      complete: () => {console.log('Completed GET Forums');},
    });
  }

}
