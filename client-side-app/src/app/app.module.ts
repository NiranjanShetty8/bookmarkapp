import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';

import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { UserHomeComponent } from './user-home/user-home.component';
import { CategoryComponent } from './category/category.component';
import { BookmarkComponent } from './bookmark/bookmark.component';
import { AddCategoryComponent } from './add-category/add-category.component';
import { EditCategoryComponent } from './edit-category/edit-category.component';
import { AddBookmarkComponent } from './add-bookmark/add-bookmark.component';
import { EditBookmarkComponent } from './edit-bookmark/edit-bookmark.component';
import { RouterModule } from '@angular/router';
import { routeArray } from './RouteConfig';

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
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
