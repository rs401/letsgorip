import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Post } from 'src/app/models/post';
import { User } from 'src/app/models/user';
import { AuthService } from 'src/app/services/auth.service';
import { ForumService } from 'src/app/services/forum.service';

@Component({
  selector: 'app-new-post',
  templateUrl: './new-post.component.html',
  styleUrls: ['./new-post.component.css']
})
export class NewPostComponent implements OnInit {

  public fid: number = 0;
  public tid: number = 0;
  public currentUser?: User;
  public token: string = '';
  public post: Post = new Post;
  titleControl: FormControl = new FormControl('');
  msgControl: FormControl = new FormControl('');
  newPostMessage: string = '';

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private forumService: ForumService,
    private auth: AuthService,
    ) {
    this.auth.user.subscribe( user => this.currentUser = user);
  }

  ngOnInit(): void {
    const routeParams = this.route.snapshot.paramMap;
    this.fid = Number(routeParams.get('id'));
    this.tid = Number(routeParams.get('tid'));
  }

  createPost() {
    const title = String(this.titleControl.value);
    const msg = String(this.msgControl.value);

    if(msg) {
      this.post.thread_id = this.tid;
      this.post.user_id = this.currentUser?.id;
      this.post.msg = msg;
      this.forumService.createPost(this.auth.token, this.fid, this.post).subscribe({
        next: () => { this.router.navigate([`/forum/${this.fid}/${this.tid}/`]) }
      });
    } else {
      this.showFlashMessage('Message cannot be empty.')
    }
  }

  showFlashMessage(msg: string) {
    this.newPostMessage = msg;
  }

}
