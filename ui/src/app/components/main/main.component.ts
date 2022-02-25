import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup } from '@angular/forms';
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
  stateForm!: FormGroup;

  constructor(private forumService: ForumService, private fb:FormBuilder, private router: Router) {
    
  }

  ngOnInit(): void {
    this.forumService.getForums().subscribe({
      next: (res) => {
        res.forEach((forum) => {
          this.forums.push(forum);
        });
      },
      error: (err) => {console.log('Error: ' + err);},
      complete: () => {console.log('MainComponent: Completed GET Forums');},
    });
    this.stateForm = this.fb.group({
      state: [0]
    });
  }

  update() {
    console.log(this.stateForm.value)
    this.router.navigateByUrl('/forum/' + this.stateForm.value.state)
  }

}
