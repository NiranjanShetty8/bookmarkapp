import { Component, OnInit } from '@angular/core';
import { BookmarkService, IBookmark } from 'src/app/services/bookmark-service/bookmark.service';
import { CategoryService } from 'src/app/services/category-service/category.service';
import { NgbAlert } from '@ng-bootstrap/ng-bootstrap';

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

  constructor(private _service: BookmarkService, private categoryService: CategoryService) { }

  getAllBookmarksOfUser() {
    this._service.getAllBookmarksOfUser().subscribe((data: IBookmark[]) => {
      this.errorOccured = false
      this.bookmarks = data
      this.assignDisplayAndCategoryName()
    }, (error) => {
      this.errorOccured = true
      this.errorMessage = error
      alert(error)
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

  deleteBookmark(categoryID, bookmarkID: string) {
    if (!confirm("Are you sure you want to delete the bookmark?")) {
      return
    }
    this._service.deleteBookmark(categoryID, bookmarkID).subscribe((data: string) => {
      alert(data)
      location.reload()
    }, (error) => {
      alert(error)
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

  getSpecificBookmark(event: any) {
    let name = event.target.value

    if (name == "") {
      for (let bookmark of this.bookmarks) {
        bookmark.display = true
      }
      return
    }
    for (let bookmark of this.bookmarks) {
      if (bookmark.name.indexOf(name) == -1) {
        bookmark.display = false
      } else {
        bookmark.display = true
      }
    }
  }

  assignDisplayAndCategoryName() {
    for (let bookmark of this.bookmarks) {
      bookmark.display = true
      this.categoryService.getCategoryName(bookmark.categoryID).subscribe((data: string) => {
        this.errorOccured = false
        bookmark.categoryName = data
      }, (error) => {
        this.errorOccured = true
        this.errorMessage = error
        alert(error)
      })
    }
  }

  ngOnInit() {
    this.getAllBookmarksOfUser()
  }

}
