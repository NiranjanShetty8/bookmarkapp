import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { AppConstants } from 'src/app/Constants';

@Injectable({
  providedIn: 'root'
})
export class BookmarkService {
  baseURL: string
  constructor(private _http: HttpClient, private _router: Router) {
    this.baseURL = `${AppConstants.baseURL}/${sessionStorage.getItem("userid")}/category/\
    ${sessionStorage.getItem("categoryid")}`
    console.log(this.baseURL)
  }
}

export interface IBookmark {
  id: string
  name: string
  url: string
  categoryID: string
}