import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { AppConstants } from 'src/app/Constants';
import { Observable } from 'rxjs';
import { IUser } from '../user-service/user.service';

@Injectable({
  providedIn: 'root'
})
export class AdminService {
  _baseURL: string
  loginAttempts: number = 3

  constructor(private _http: HttpClient, private _router: Router) {
    this._baseURL = AppConstants.baseURL
  }

  getAllusers(): Observable<IUser[]> {
    return new Observable<IUser[]>((observer) => {
      this._http.get(`${this._baseURL}/all`, {
        headers: this.setTokenToHeader()
      }).subscribe((data: IUser[]) => {
        observer.next(data)
      }, (error) => {
        observer.error(error.error)
      })
    })
  }

  deleteUser(userID: string): Observable<string> {
    return new Observable<string>((observer) => {
      this._http.delete(`${this._baseURL}/${userID}`, {
        headers: this.setTokenToHeader()
      }).subscribe((data: string) => {
        observer.next(data)
      }, (error) => {
        observer.error(error.error)
      })
    })
  }

  unlockUser(user: IUser): Observable<string> {
    user.loginAttempts = this.loginAttempts
    return new Observable<string>((observer) => {
      this._http.put(`${this._baseURL}/${user.id}`, {
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
