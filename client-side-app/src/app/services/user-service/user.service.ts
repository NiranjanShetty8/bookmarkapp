import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { AppConstants } from '../../Constants'
import { Observable } from 'rxjs';
import { Router } from '@angular/router';
import { ICategory } from '../category-service/category.service';

@Injectable({
  providedIn: 'root'
})

export class UserService {

  _baseURL: string
  constructor(private _http: HttpClient, private _router: Router) {
    this._baseURL = AppConstants.baseURL
  }

  register(user: IUser): Observable<string> {
    return new Observable<string>((observer) => {
      this._http.post(this._baseURL + "/register", user)
        .subscribe((data: string) => {
          observer.next(data)
        }, (error) => {
          observer.error(error.error)
        })

    });
  }

  login(user: IUser): Observable<IUser> {
    return new Observable<IUser>((observer) => {
      this._http.post(this._baseURL + "/login", user)
        .subscribe((data: IUser) => {
          sessionStorage.setItem('token', data.token)
          sessionStorage.setItem('userid', data.id)
          observer.next(data)
          this._router.navigate([data.username + '/home'])
        }, (error) => {
          observer.error(error.error)
        })
    });
  }

}


export interface IUser {
  id?: string,
  username: string,
  password: string,
  email?: string,
  profilePhoto?: any,
  superUser: boolean,
  categories?: ICategory,
  token?: string
}