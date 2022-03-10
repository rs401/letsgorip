import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HeaderComponent } from './components/header/header.component';
import { MainComponent } from './components/main/main.component';
import { SignInComponent } from './components/sign-in/sign-in.component';
import { SignUpComponent } from './components/sign-up/sign-up.component';
import { FlashMessageComponent } from './components/flash-message/flash-message.component';
import { ForumComponent } from './components/forum/forum.component';
import { ThreadComponent } from './components/thread/thread.component';
import { NewThreadComponent } from './components/new-thread/new-thread.component';
import { NewPostComponent } from './components/new-post/new-post.component';
import { TosComponent } from './components/tos/tos.component';
import { PrivacyComponent } from './components/privacy/privacy.component';
import { RulesComponent } from './components/rules/rules.component';
import { ScrollTextComponent } from './components/scroll-text/scroll-text.component';
import { UserCardComponent } from './components/user-card/user-card.component';
import { PlacesComponent } from './components/places/places.component';
import { HashLocationStrategy, LocationStrategy } from '@angular/common';

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    MainComponent,
    SignInComponent,
    SignUpComponent,
    FlashMessageComponent,
    ForumComponent,
    ThreadComponent,
    NewThreadComponent,
    NewPostComponent,
    TosComponent,
    PrivacyComponent,
    RulesComponent,
    ScrollTextComponent,
    UserCardComponent,
    PlacesComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    ReactiveFormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
  ],
  providers: [{ provide: LocationStrategy, useClass: HashLocationStrategy}],
  bootstrap: [AppComponent]
})
export class AppModule { }
