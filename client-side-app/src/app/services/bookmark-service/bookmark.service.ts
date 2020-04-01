import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';
import { AppConstants } from 'src/app/Constants';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class BookmarkService {
  _baseURL: string
  _allBookmarkURL: string

  constructor(private _http: HttpClient, private _router: Router) {
    this._baseURL = `${AppConstants.baseURL}/${sessionStorage.getItem("userid")}/category`
    this._allBookmarkURL = `${AppConstants.baseURL}/${sessionStorage.getItem("userid")}/bookmark/all`
  }


  getAllBookmarksOfUser(): Observable<IBookmark[]> {
    return new Observable<IBookmark[]>((observer) => {
      this._http.get(this._allBookmarkURL, {
        headers: this.setTokenToHeader()
      }).subscribe((data: IBookmark[]) => {
        observer.next(data)
      }, (error) => {
        observer.error(error.error)
      })
    })
  }

  getAllBookmarks(categoryID: string): Observable<IBookmark[]> {
    return new Observable<IBookmark[]>((observer) => {
      this._http.get(`${this._baseURL}/${categoryID}`, {
        headers: this.setTokenToHeader()
      }).subscribe((data: IBookmark[]) => {
        observer.next(data)
      }, (error) => {
        observer.error(error.error)
      })
    })
  }


  getBookmarkByName(categoryID, bookmarkName: string): Observable<IBookmark> {
    return new Observable<IBookmark>((observer) => {
      this._http.get(`${this._baseURL}/${categoryID}/bookmark/name/${bookmarkName}`, {
        headers: this.setTokenToHeader()
      }).subscribe((data: IBookmark) => {
        observer.next(data)
      }, (error) => {
        observer.error(error.error);
      });
    })
  }

  getBookmarkByID(categoryID, bookmarkID: string): Observable<IBookmark> {
    return new Observable<IBookmark>((observer) => {
      this._http.get(`${this._baseURL}/${categoryID}/bookmark/${bookmarkID}`, {
        headers: this.setTokenToHeader()
      }).subscribe((data: IBookmark) => {
        observer.next(data)
      }, (error) => {
        observer.error(error.error)
      })
    })
  }

  addBookmark(categoryID, bookmark: IBookmark): Observable<string> {
    return new Observable<string>((observer) => {
      this._http.post(`${this._baseURL}/${categoryID}`, bookmark, {
        headers: this.setTokenToHeader()
      }).subscribe((data: string) => {
        console.log(data)
        observer.next(data)
      }, (error) => {
        observer.error(error.error)
      });
    })
  }

  deleteBookmark(categoryID, bookmarkID: string): Observable<string> {
    return new Observable<string>((observer) => {
      this._http.delete(`${this._baseURL}/${categoryID}/bookmark/${bookmarkID}`, {
        headers: this.setTokenToHeader()
      }).subscribe((data: string) => {
        observer.next(data)
      }, (error) => {
        observer.error(error.error)
      })
    })
  }

  updateBookmark(bookmark: IBookmark): Observable<string> {
    return new Observable<string>((observer) => {
      this._http.put(`${this._baseURL}/${bookmark.categoryID}/bookmark/${bookmark.id}`, bookmark, {
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

export interface IBookmark {
  id: string
  name: string
  url: string
  categoryID: string
  display?: boolean
  categoryName?: string
}