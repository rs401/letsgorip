import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AppComponent } from './app.component';
import { ForumComponent } from './components/forum/forum.component';
import { MainComponent } from './components/main/main.component';
import { NewPostComponent } from './components/new-post/new-post.component';
import { NewThreadComponent } from './components/new-thread/new-thread.component';
import { SignInComponent } from './components/sign-in/sign-in.component';
import { SignUpComponent } from './components/sign-up/sign-up.component';
import { ThreadComponent } from './components/thread/thread.component';

const routes: Routes = [
  { path: '', component: MainComponent },
  { path: 'sign-in', component: SignInComponent },
  { path: 'sign-up', component: SignUpComponent },
  { 
    path: 'forum',
    children: [
      { path: ':id', component: ForumComponent },
      { path: ':id/new', component: NewThreadComponent },
      { path: ':id/:tid', component: ThreadComponent },
      { path: ':id/:tid/new', component: NewPostComponent }
    ],
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
