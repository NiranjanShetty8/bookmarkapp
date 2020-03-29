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
  private categories: any
  constructor(private _http: HttpClient, private _route: Router) {
    this._baseURL = `${AppConstants.baseURL}/${sessionStorage.getItem("userid")}/category`
    this.categories = []
  }
  getAllCategories(): Observable<ICategory[]> {
    return new Observable<ICategory[]>((observer) => {
      this._http.get(this._baseURL, {
        headers: new HttpHeaders().set("token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3N1ZWRBdCI6MTU4NTQ4MzMzNCwidXNlcklEIjoiZDU1Zjg3M2EtZjFiOC00NmRmLThhNWItMWU4ZGEwMmJiMjk3IiwidXNlcm5hbWUiOiJOaXJhbmphbiIsInZhbGlkVGlsbCI6MTU4NTQ4NTEzNH0.RqeoYxxc2sNjPl_Uui2ORerP6k3fI8VRCDKi77bDk1w")
      }).subscribe((data: ICategory[]) => {
        observer.next(data)
        console.log("data:", data)
      }, (error) => {
        console.log("in error")
        observer.error(error.error)
      })
    })
  }


  addCategory(category: ICategory): Observable<string> {
    console.log(this.setTokenToHeader())
    return new Observable<string>((observer) => {
      this._http.post(this._baseURL, category,
        {
          headers: this.setTokenToHeader()
        }
      ).subscribe((data: string) => {
        console.log(data)
        observer.next(data)
      }, (error) => {
        observer.next(error.error)
      })
    })
  }

  setTokenToHeader(): HttpHeaders {
    return new HttpHeaders().set("token", sessionStorage.getItem("token"))
  }
}




export interface ICategory {
  id?: string,
  name: string,
  bookmarks?: IBookmark[],
  userID: string
}