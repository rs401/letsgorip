import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { Title } from '@angular/platform-browser';
import { ActivatedRoute, Router } from '@angular/router';
import { Thread } from 'src/app/models/thread';
import { User } from 'src/app/models/user';
import { AuthService } from 'src/app/services/auth.service';
import { ForumService } from 'src/app/services/forum.service';

@Component({
  selector: 'app-new-thread',
  templateUrl: './new-thread.component.html',
  styleUrls: ['./new-thread.component.css']
})
export class NewThreadComponent implements OnInit {

  public fid: number = 0;
  public currentUser?: User;
  public thread: Thread = new Thread;
  titleControl: FormControl = new FormControl('');
  msgControl: FormControl = new FormControl('');
  newThreadMessage: string = '';

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private forumService: ForumService,
    private auth: AuthService,
    private title: Title,
    ) {
    this.auth.user.subscribe( user => this.currentUser = user);
  }

  ngOnInit(): void {
    const routeParams = this.route.snapshot.paramMap;
    this.fid = Number(routeParams.get('id'));
    this.title.setTitle('LGR: Create new discussion thread');
  }

  createThread() {
    const title = String(this.titleControl.value);
    const msg = String(this.msgControl.value);

    if(title && msg) {
      this.thread.forum_id = this.fid;
      this.thread.user_id = this.currentUser?.id;
      this.thread.title = title;
      this.thread.msg = msg;
      this.forumService.createThread(this.auth.token, this.thread).subscribe({
        next: () => { this.router.navigate([`/forum/${this.fid}/`]) }
      });
    } else {
      this.showFlashMessage('Fields cannot be empty.')
    }
  }

  showFlashMessage(msg: string) {
    this.newThreadMessage = msg;
  }

}
