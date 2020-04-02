import { Component, OnInit } from '@angular/core';
import { BookmarkService, IBookmark } from 'src/app/services/bookmark-service/bookmark.service';
import { CategoryService, ICategory } from 'src/app/services/category-service/category.service';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { Location } from '@angular/common';


@Component({
  selector: 'bookmarkapp-bookmark',
  templateUrl: './bookmark.component.html',
  styleUrls: ['./bookmark.component.css']
})
export class BookmarkComponent implements OnInit {
  bookmarks: IBookmark[]
  bookmark: IBookmark
  bookmarkForm: FormGroup
  formName: string
  formOperationName: string
  formMessage: string
  allCategoryNames: string[]
  allCategories: ICategory[]
  userHomeLink: string
  loading: boolean

  constructor(private _service: BookmarkService,
    private categoryService: CategoryService, private modalService: NgbModal, private location: Location) {
    location.onUrlChange((urlChanged) => {
      var categoryID = (urlChanged.split("category/", 2)[1])
      if (location.path() != this.userHomeLink) {
        this.getAllBookmarks(categoryID)
      }
    })
  }



  getAllBookmarksOfUser() {
    this.loading = true
    this.location.replaceState(this.userHomeLink)
    this._service.getAllBookmarksOfUser().subscribe((data: IBookmark[]) => {
      this.bookmarks = data
      this.assignDisplayAndCategoryName()
      this.loading = false

    }, (error) => {
      this.loading = false
      alert(error)
    });
  }

  getAllBookmarks(categoryID: string) {
    this.loading = true
    this._service.getAllBookmarks(categoryID).subscribe((data: IBookmark[]) => {
      this.bookmarks = data
      this.assignDisplayAndCategoryName()
      this.loading = false
    }, (error) => {
      this.loading = false
      alert(error)
    });
  }

  getBookmarkByName(categoryID, bookmarkName: string) {
    this.loading = true
    this._service.getBookmarkByName(categoryID, bookmarkName).subscribe((data: IBookmark) => {
      this.bookmark = data
      this.loading = false
    }, (error) => {
      this.loading = false
      alert(error)
    });
  }

  getBookmarkByID(categoryID, bookmarkID: string) {
    this.loading = true
    this._service.getBookmarkByID(categoryID, bookmarkID).subscribe((data: IBookmark) => {
      this.bookmark = data
      this.loading = false
    }, (error) => {
      this.loading = false
      alert(error)

    });
  }

  addBookmark() {
    this.loading = true
    this.bookmark = this.bookmarkForm.value
    for (let category of this.allCategories) {
      if (this.bookmark.categoryName == category.name) {
        this.bookmark.categoryID = category.id
      }
    }
    this._service.addBookmark(this.bookmark).subscribe((data: string) => {
      this.loading = false
      alert(`Bookmark Added with ID ${data}`)
      location.reload()

    }, (error) => {
      this.loading = false
      alert(error)
    })
  }

  deleteBookmark(categoryID, bookmarkID: string) {
    if (!confirm("Are you sure you want to delete the bookmark?")) {
      return
    }
    this.loading = true
    this._service.deleteBookmark(categoryID, bookmarkID).subscribe((data: string) => {
      this.loading = false
      alert(data)
      location.reload()
    }, (error) => {
      this.loading = false
      alert(error)
    });
  }

  updateBookmark(bookmark: IBookmark) {
    this.bookmark = this.bookmarkForm.value
    if (this.bookmark.name === bookmark.name && this.bookmark.url === bookmark.url &&
      this.bookmark.categoryName === bookmark.categoryName) {
      alert("No changes made to book mark.");
      return
    }
    this.loading = true
    if (this.bookmark.categoryName != bookmark.categoryName) {
      for (let category of this.allCategories) {
        if (this.bookmark.categoryName == category.name) {
          this.bookmark.categoryID = category.id
        }
      }
    }
    this._service.updateBookmark(this.bookmark).subscribe((data: string) => {
      this.loading = false
      alert(data)
      location.reload()
    }, (error) => {
      this.loading = false
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
    this.loading = true
    for (let bookmark of this.bookmarks) {
      bookmark.display = true
      this.categoryService.getCategoryName(bookmark.categoryID).subscribe((data: string) => {
        bookmark.categoryName = data
        this.loading = false
      }, (error) => {
        this.loading = false
        alert(error)
      })
    }
  }

  openModal(content: any, catgeoryID?: string, bookmark?: IBookmark) {
    this.setAllCategoriesAndNames()
    if (bookmark) {
      this.formName = "Edit Bookmark";
      this.formOperationName = "Save Changes"
      this.formMessage = "you make some valid changes to fields."
      this.bookmarkForm.setValue(bookmark)
      this.modalService.open(content, { ariaLabelledBy: 'bookmark-modal' }).result.then(() => {
        console.log("calling update")
        this.updateBookmark(bookmark)
      }, () => {
        this.bookmarkForm.reset()
      });
      return
    }
    this.formName = "Add Bookmark";
    this.formOperationName = "Add Bookmark"
    this.formMessage = "you fill all fields with valid data."
    this.bookmarkForm.reset()
    this.modalService.open(content, { ariaLabelledBy: 'bookmark-modal' }).result.then((result) => {
      this.addBookmark()
    }, () => void 0);
  }

  setAllCategoriesAndNames() {
    this.loading = true
    this.allCategoryNames = [];
    this.categoryService.getAllCategories().subscribe((data: ICategory[]) => {
      this.allCategories = data
      for (let category of data) {
        this.allCategoryNames.push(category.name);
      }
      this.loading = false
    }, (error) => {
      this.loading = false
      alert(error)
    })
  }

  ngOnInit() {
    this.userHomeLink = `/${sessionStorage.getItem("username")}/home`
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
      categoryName: new FormControl(null, Validators.required)

    })
  }
}
