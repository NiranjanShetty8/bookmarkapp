import { Component, OnInit } from '@angular/core';
import { BookmarkService, IBookmark } from 'src/app/services/bookmark-service/bookmark.service';

@Component({
  selector: 'bookmarkapp-bookmark',
  templateUrl: './bookmark.component.html',
  styleUrls: ['./bookmark.component.css']
})
export class BookmarkComponent implements OnInit {
  bookmarks: IBookmark[]
  errorOccured: boolean
  errorMessage: string
  successMessage: string
  bookmark: IBookmark

  constructor(private _service: BookmarkService) { }

  getAllBookmarksOfUser() {
    this._service.getAllBookmarksOfUser().subscribe((data: IBookmark[]) => {
      this.errorOccured = false
      console.log("data", data)
      this.bookmarks = data
    }, (error) => {
      this.errorOccured = true
      this.errorMessage = error
    });
  }

  getAllBookmarks() {
    this._service.getAllBookmarks().subscribe((data: IBookmark[]) => {
      this.errorOccured = false
      console.log("data", data)
      this.bookmarks = data
    }, (error) => {
      this.errorOccured = true
      this.errorMessage = error
    });
  }

  getBookmarkByName(bookmarkName: string) {
    this._service.getBookmarkByName(bookmarkName).subscribe((data: IBookmark) => {
      this.errorOccured = false
      this.bookmark = data
    }, (error) => {
      this.errorOccured = true
      this.errorMessage = error
    });
  }

  getBookmarkByID(bookmarkID: string) {
    this._service.getBookmarkByID(bookmarkID).subscribe((data: IBookmark) => {
      this.errorOccured = false
      this.bookmark = data
    }, (error) => {
      this.errorOccured = true
      this.errorMessage = error
    });
  }

  addBookmark(bookmark: IBookmark) {
    this._service.addBookmark(bookmark).subscribe((data: string) => {
      this.errorOccured = false
      console.log(data)
      this.successMessage = data

    }, (error) => {
      this.errorOccured = true
      this.errorMessage = error
      console.log(error)
    })
  }

  deleteBookmark(bookmarkID: string) {
    this._service.deleteBookmark(bookmarkID).subscribe((data: string) => {
      this.errorOccured = false
      this.successMessage = data
    }, (error) => {
      this.errorOccured = true
      this.errorMessage = error
    });
  }

  updateBookmark(bookmark: IBookmark) {
    this._service.updateBookmark(bookmark).subscribe((data: string) => {
      this.errorOccured = false
      this.successMessage = data
    }, (error) => {
      this.errorOccured = true
      this.errorMessage = error
    });
  }
  ngOnInit() {
    this.getAllBookmarksOfUser()
  }

}
