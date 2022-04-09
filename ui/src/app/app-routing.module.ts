import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AppComponent } from './app.component';
import { ForumComponent } from './components/forum/forum.component';
import { MainComponent } from './components/main/main.component';
import { NewPostComponent } from './components/new-post/new-post.component';
import { NewThreadComponent } from './components/new-thread/new-thread.component';
import { PlacesComponent } from './components/places/places.component';
import { PrivacyComponent } from './components/privacy/privacy.component';
import { RulesComponent } from './components/rules/rules.component';
import { SearchForumsComponent } from './components/search-forums/search-forums.component';
import { SignInComponent } from './components/sign-in/sign-in.component';
import { SignUpComponent } from './components/sign-up/sign-up.component';
import { ThreadComponent } from './components/thread/thread.component';
import { TosComponent } from './components/tos/tos.component';
import { AuthGuardGuard } from './services/auth-guard.guard';

const routes: Routes = [
  { path: '', component: MainComponent },
  { path: 'sign-in/:creds', component: SignInComponent },
  { path: 'sign-in', component: SignInComponent },
  { path: 'sign-up', component: SignUpComponent },
  { path: 'tos', component: TosComponent },
  { path: 'privacy', component: PrivacyComponent },
  { path: 'rules', component: RulesComponent },
  { path: 'places', component: PlacesComponent },
  { path: 'search/:query', component: SearchForumsComponent },
  { 
    path: 'forum',
    children: [
      { path: ':id', component: ForumComponent },
      { path: ':id/new', component: NewThreadComponent, canActivate: [AuthGuardGuard] },
      { path: ':id/:tid', component: ThreadComponent },
      { path: ':id/:tid/new', component: NewPostComponent, canActivate: [AuthGuardGuard] }
    ],
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes, {useHash: true})],
  exports: [RouterModule]
})
export class AppRoutingModule { }
