import { Component, OnInit } from '@angular/core';
import { BookmarkService, IBookmark } from 'src/app/services/bookmark-service/bookmark.service';
import { CategoryService } from 'src/app/services/category-service/category.service';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';

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
  bookmarkForm: FormGroup
  formName: string
  formOperationName: string


  constructor(private _service: BookmarkService,
    private categoryService: CategoryService, private modalService: NgbModal) { }

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

  getAllBookmarks(categoryID: string) {
    this._service.getAllBookmarks(categoryID).subscribe((data: IBookmark[]) => {
      this.errorOccured = false
      console.log("data", data)
      this.bookmarks = data
    }, (error) => {
      this.errorOccured = true
      this.errorMessage = error
    });
  }

  getBookmarkByName(categoryID, bookmarkName: string) {
    this._service.getBookmarkByName(categoryID, bookmarkName).subscribe((data: IBookmark) => {
      this.errorOccured = false
      this.bookmark = data
    }, (error) => {
      this.errorOccured = true
      this.errorMessage = error
    });
  }

  getBookmarkByID(categoryID, bookmarkID: string) {
    this._service.getBookmarkByID(categoryID, bookmarkID).subscribe((data: IBookmark) => {
      this.errorOccured = false
      this.bookmark = data
    }, (error) => {
      this.errorOccured = true
      this.errorMessage = error
    });
  }

  addBookmark(categoryID: string) {
    this.bookmark = this.bookmarkForm.value
    this._service.addBookmark(categoryID, this.bookmark).subscribe((data: string) => {
      alert(`Bookmark Added with ID ${data}`)
    }, (error) => {
      alert(error)
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
    this.bookmark = this.bookmarkForm.value
    if (this.bookmark.name === bookmark.name && this.bookmark.url === bookmark.url &&
      this.bookmark.categoryID === bookmark.categoryID) {
      alert("No changes made to book mark.");
      return
    }
    this._service.updateBookmark(this.bookmark).subscribe((data: string) => {
      alert(data)
      location.reload()
    }, (error) => {
      alert(error)
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

  openModal(content: any, catgeoryID?: string, bookmark?: IBookmark) {
    if (bookmark) {
      this.formName = "Edit Bookmark";
      this.formOperationName = "Save Changes"
      this.bookmarkForm.setValue(bookmark)
      this.modalService.open(content, { ariaLabelledBy: 'bookmark-modal' }).result.then((result) => {
        console.log("calling update")
        this.updateBookmark(bookmark)
      }, () => {
        this.bookmarkForm.reset()
      });
      return
    }
    this.formName = "Add Bookmark";
    this.formOperationName = "Add Bookmark"
    this.bookmarkForm.reset()
    this.modalService.open(content, { ariaLabelledBy: 'bookmark-modal' }).result.then((result) => {
      this.addBookmark(catgeoryID)
    }, () => void 0);
  }

  ngOnInit() {
    this.getAllBookmarksOfUser()
    this.initAddForm()
  }

  private initAddForm() {
    this.bookmarkForm = new FormGroup({
      id: new FormControl(null),
      name: new FormControl(null, [Validators.required, Validators.minLength(3), Validators.maxLength(30)]),
      url: new FormControl(null, [Validators.required, Validators.minLength(8)]),
      categoryID: new FormControl(null),
      display: new FormControl(null),
      categoryName: new FormControl(null)

    })
  }
}
