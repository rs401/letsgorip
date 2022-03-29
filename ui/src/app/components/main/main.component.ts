import { Component, ElementRef, OnInit, Renderer2, ViewChild } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
import { Router, ActivatedRoute, ParamMap } from '@angular/router';
import { timer } from 'rxjs';
import { Forum } from 'src/app/models/forum';
import { ForumService } from 'src/app/services/forum.service';

@Component({
  selector: 'app-main',
  templateUrl: './main.component.html',
  styleUrls: ['./main.component.css']
})
export class MainComponent implements OnInit {
  
  public forums: Forum[] = [];
  public stateForm!: FormGroup;
  // private source = timer(1000, 3000);
  // private trigger = this.source.subscribe(() => this.scrollText())
  public scrollMsg: string = "";
  public lines: string[] = [
    "Find places to go off-road for all different vehicles.",
    "Find people to go off-road with.",
    "Find off-road parks.",
    "Find off-road events.",
    "Find a mud hole to lose your super-swampers in.",
    "Find a dune for your buggy.",
    "Find a parking lot to do donuts in",
  ];

  constructor(
    private forumService: ForumService,
    private fb: FormBuilder,
    private router: Router) {

  }

  ngOnInit(): void {
    this.forumService.getForums().subscribe({
      next: (res) => {
        res.forEach((forum) => {
          this.forums.push(forum);
        });
      },
    });
    this.stateForm = this.fb.group({
      state: [0]
    });
  }

  update() {
    this.router.navigateByUrl('/forum/' + this.stateForm.value.state);
  }


}
