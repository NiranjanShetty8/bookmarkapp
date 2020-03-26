import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';

import { AppComponent } from './app.component';
import { HomeComponent } from './components/home/home.component';
import { LoginComponent } from './components/login/login.component';
import { RegisterComponent } from './components/register/register.component';
import { UserHomeComponent } from './components/user-home/user-home.component';
import { CategoryComponent } from './components/category/category.component';
import { BookmarkComponent } from './components/bookmark/bookmark.component';
import { AddCategoryComponent } from './components/add-category/add-category.component';
import { EditCategoryComponent } from './components/edit-category/edit-category.component';
import { AddBookmarkComponent } from './components/add-bookmark/add-bookmark.component';
import { EditBookmarkComponent } from './components/edit-bookmark/edit-bookmark.component';
import { RouterModule } from '@angular/router';
import { routeArray } from './RouteConfig';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule, ReactiveFormsModule } from '@angular/forms'


@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    LoginComponent,
    RegisterComponent,
    UserHomeComponent,
    CategoryComponent,
    BookmarkComponent,
    AddCategoryComponent,
    EditCategoryComponent,
    AddBookmarkComponent,
    EditBookmarkComponent
  ],
  imports: [
    BrowserModule,
    NgbModule,
    RouterModule.forRoot(routeArray),
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
