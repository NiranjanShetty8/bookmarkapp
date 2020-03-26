import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { AppConstants } from '../../Constants'
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UserService {

  _baseURL: string
  constructor(private _http: HttpClient) {
    this._baseURL = AppConstants.baseURL
  }

  Register(user: IUser): Observable<any> {
    return new Observable<any>((observer) => {
      this._http.post(this._baseURL + "/register", user)
        .subscribe((data: any) => {
          observer.next(data)
        }, (error) => {
          observer.error(error.message)
        })

    });
  }

}



export interface IUser {
  id?: number,
  username: string,
  password: string,
  token?: string
}