import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';
import { AppConstants } from 'src/app/Constants';
import { IBookmark } from '../bookmark-service/bookmark.service';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class CategoryService {
  _baseURL: string
  constructor(private _http: HttpClient, private _route: Router) {
    this._baseURL = `${AppConstants.baseURL}/${sessionStorage.getItem("userid")}/category`
  }
  getAllCategories(): Observable<ICategory[]> {
    return new Observable<ICategory[]>((observer) => {
      this._http.get(this._baseURL, {
        headers: this.setTokenToHeader()
      }).subscribe((data: ICategory[]) => {
        observer.next(data)
      }, (error) => {
        observer.error(error.error)
      })
    })
  }

  setTokenToHeader(): HttpHeaders {
    return new HttpHeaders().set('token', sessionStorage.getByID('token'));
  }
}




export interface ICategory {
  id: string,
  name: string,
  bookmarks: IBookmark[],
  userID: string
}