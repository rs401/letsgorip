import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { map, Observable } from 'rxjs';
import { Post } from 'src/app/models/post';
import { Thread } from 'src/app/models/thread';
import { ForumService } from 'src/app/services/forum.service';

@Component({
  selector: 'app-thread',
  templateUrl: './thread.component.html',
  styleUrls: ['./thread.component.css']
})
export class ThreadComponent implements OnInit {
  public fid: number = 0;
  public tid: number = 0;
  public thread?: Thread;
  public posts: Post[] = [];

  constructor(private route: ActivatedRoute, private forumService: ForumService) {
    this.forumService.thread.subscribe(t => this.thread = t);
  }

  ngOnInit(): void {
    const routeParams = this.route.snapshot.paramMap;
    this.fid = Number(routeParams.get('id'));
    this.tid = Number(routeParams.get('tid'));
    this.forumService.getThread(this.fid, this.tid).subscribe({
      next: (res) => {console.log('threadComponent: response: ' + res);},
      error: (e) => {console.log('threadComponent: error: ' + JSON.stringify(e));},
    });
  }

}
