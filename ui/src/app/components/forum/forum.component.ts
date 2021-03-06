import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { map, Observable } from 'rxjs';
import { Forum } from 'src/app/models/forum';
import { Thread } from 'src/app/models/thread';
import { ForumService } from 'src/app/services/forum.service';

@Component({
  selector: 'app-forum',
  templateUrl: './forum.component.html',
  styleUrls: ['./forum.component.css']
})
export class ForumComponent implements OnInit {

  public fid: number = 0;
  public forum?: Forum;
  public threads: Thread[] = [];

  constructor(private route: ActivatedRoute, private router: Router, private forumService: ForumService) {
    this.forumService.forum.subscribe(f => this.forum = f);
  }

  ngOnInit(): void {
    const routeParams = this.route.snapshot.paramMap;
    this.fid = Number(routeParams.get('id'));
    this.forumService.getForum(this.fid).subscribe();
    
    this.forumService.getThreads(this.fid).subscribe({
      next: (res) => {
        res.forEach((thread) => {
          this.threads.push(thread);
        });
      },
      error: (e) => {console.log('Error: ' + e);},
    });
  }

  viewThread(thread: Thread) {
    this.router.navigateByUrl(`/forum/${thread.forum_id}/${thread.id}`);
    // this.router.navigateByUrl(`/forum/${thread.forum_id}/thread/${thread.id}`);
  }

}
