import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';
import { AppConstants } from 'src/app/Constants';
import { IBookmark } from '../bookmark-service/bookmark.service';
import { Observable, observable } from 'rxjs';

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

  getCategoryByName(categoryName: string): Observable<ICategory> {
    return new Observable<ICategory>((observer) => {
      this._http.get(`${this._baseURL}/name/${categoryName}`, {
        headers: this.setTokenToHeader()
      }).subscribe((data: ICategory) => {
        observer.next(data)
      }, (error) => {
        observer.error(error.error);
      });
    })
  }

  getCategoryByID(categoryID: string): Observable<ICategory> {
    return new Observable<ICategory>((observer) => {
      this._http.get(`${this._baseURL}/${categoryID}`, {
        headers: this.setTokenToHeader()
      }).subscribe((data: ICategory) => {
        observer.next(data)
      }, (error) => {
        observer.error(error.error)
      })
    })
  }

  getCategoryName(categoryID: string): Observable<string> {
    return new Observable<string>((observer) => {
      this._http.get(`${this._baseURL}/${categoryID}`, {
        headers: this.setTokenToHeader()
      }).subscribe((data: ICategory) => {
        observer.next(data.name)
      }, (error) => {
        observer.error(error.error)
      })
    })
  }

  addCategory(category: ICategory): Observable<string> {
    console.log(this.setTokenToHeader())
    return new Observable<string>((observer) => {
      this._http.post(this._baseURL, category, {
        headers: this.setTokenToHeader()
      }).subscribe((data: string) => {
        console.log(data)
        observer.next(data)
      }, (error) => {
        observer.error(error.error)
      });
    })
  }

  deleteCategory(categoryID: string): Observable<string> {
    return new Observable<string>((observer) => {
      this._http.delete(`${this._baseURL}/${categoryID}`, {
        headers: this.setTokenToHeader()
      }).subscribe((data: string) => {
        observer.next(data)
      }, (error) => {
        observer.error(error.error)
      })
    })
  }

  updateCategory(category: ICategory): Observable<string> {
    return new Observable<string>((observer) => {
      this._http.put(`${this._baseURL}/${category.id}`, category, {
        headers: this.setTokenToHeader()
      }).subscribe((data: string) => {
        observer.next(data)
      }, (error) => {
        observer.error(error.error)
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
  userID: string,
  display: boolean
}