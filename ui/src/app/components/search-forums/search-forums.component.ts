import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Forum } from 'src/app/models/forum';
import { Thread } from 'src/app/models/thread';
import { ForumService } from 'src/app/services/forum.service';

@Component({
  selector: 'app-search-forums',
  templateUrl: './search-forums.component.html',
  styleUrls: ['./search-forums.component.css']
})
export class SearchForumsComponent implements OnInit {
  public query: string = '';
  public threads: Thread[] = [];
  public forums: Forum[] = [];
  public stateAbrs: string[] = ['AL','AK','AZ','AR','CA','CO','CT','DE','DC','FL','GA','HI','ID','IL','IN','IA','KS','KY','LA','ME','MD','MA','MI','MN','MS','MO','MT','NE','NV','NH','NJ','NM','NY','NC','ND','OH','OK','OR','PA','PR','RI','SC','SD','TN','TX','UT','VT','VA','WA','WV','WI','WY'];

  constructor(private route: ActivatedRoute, private router: Router, private forumService: ForumService) { }

  ngOnInit(): void {
    const routeParams = this.route.snapshot.paramMap;
    this.query = routeParams.get('query')!;
    this.forumService.searchThreads(this.query).subscribe({
      next: (res) => {
        res.forEach((t) => {
          this.threads.push(t);
        });
      }
    });
    this.forumService.getForums().subscribe({
      next: (res) => {
        res.forEach((forum) => {
          this.forums.push(forum);
        });
      },
    });
  }

  viewThread(thread: Thread) {
    this.router.navigateByUrl(`/forum/${thread.forum_id}/${thread.id}`);
  }

}
