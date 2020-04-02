import { Component, OnInit } from '@angular/core';
import { CategoryService, ICategory } from 'src/app/services/category-service/category.service';
import { NgbModal, ModalDismissReasons } from '@ng-bootstrap/ng-bootstrap';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { Router, NavigationEnd, NavigationStart } from '@angular/router';
import { Location } from '@angular/common';

@Component({
  selector: 'bookmarkapp-category',
  templateUrl: './category.component.html',
  styleUrls: ['./category.component.css']
})
export class CategoryComponent implements OnInit {
  categories: ICategory[]
  category: ICategory
  categoryForm: FormGroup
  formName: string
  formOperationName: string
  formMessage: string
  userName: string
  userHomeLink: string
  loading: boolean

  constructor(private _service: CategoryService, private modalService: NgbModal,
    private location: Location) {
  }


  getAllCategories() {
    this.loading = true
    this._service.getAllCategories().subscribe((data: ICategory[]) => {
      this.categories = data
      this.assignDisplayValue()
      this.loading = false
    }, (error) => {
      this.loading = false
      alert(error)
    });
  }

  assignDisplayValue() {
    for (let category of this.categories) {
      category.display = true
    }
  }

  getCategoryByName(categoryName: string) {
    this.loading = true
    this._service.getCategoryByName(categoryName).subscribe((data: ICategory) => {
      this.category = data
      this.loading = false
    }, (error) => {
      this.loading = false
      alert(error)
    });
  }

  getCategoryByID(categoryID: string) {
    this.loading = true
    this._service.getCategoryByID(categoryID).subscribe((data: ICategory) => {
      this.category = data
      this.loading = false
    }, (error) => {
      this.loading = false
      alert(error)
    });
  }

  addCategory() {
    this.loading = true
    this.category = this.categoryForm.value
    this._service.addCategory(this.category).subscribe((data: string) => {
      this.loading = false
      alert(`Category added with ID: ${data}`)
      location.reload()
    }, (error) => {
      this.loading = false
      alert(error)
    })
  }

  deleteCategory(categoryID: string) {
    if (!confirm("Warning! Deleting a Category will delete all the bookmarks in it.\
    \nAre you sure you want to delete it?")) {
      return
    }
    this.loading = true
    this._service.deleteCategory(categoryID).subscribe((data: string) => {
      this.loading = false
      alert(data)
      location.reload()
    }, (error) => {
      this.loading = false
      alert(error)
    });
  }

  updateCategory(category: ICategory) {
    this.category = this.categoryForm.value
    if (this.category.name === category.name) {
      alert("Name has not been changed.")
      return
    }
    this.loading = true
    this._service.updateCategory(this.category).subscribe((data: string) => {
      this.loading = false
      alert("Category Updated")
      location.reload()
    }, (error) => {
      this.loading = false
      alert(error)
    });
  }

  getSpecificCategories(event: any) {
    let name = event.target.value
    if (name == "") {
      for (let category of this.categories) {
        category.display = true
      }
      return
    }
    for (let category of this.categories) {
      if (category.name.indexOf(name) == -1) {
        category.display = false
      } else {
        category.display = true
      }
    }
  }

  showBookmarksOfSpecificCAtegory(categoryID: string) {
    this.location.replaceState(`${this.userHomeLink}/category/${categoryID}`)
  }

  openModal(content: any, category?: ICategory) {
    if (category) {
      this.formName = "Edit Category";
      this.formOperationName = "Save Changes"
      this.formMessage = "you make some valid changes to fields."
      this.categoryForm.setValue(category)
      this.modalService.open(content, { ariaLabelledBy: 'category-modal' }).result.then((result) => {
        this.updateCategory(category)
      }, () => {
        this.categoryForm.reset()
      });
      return
    }
    this.formName = "Add Category";
    this.formOperationName = "Add Category"
    this.formMessage = "you fill all fields with valid data."
    this.categoryForm.reset()
    this.modalService.open(content, { ariaLabelledBy: 'category-modal' }).result.then((result) => {
      this.addCategory()
    }, () => void 0);
  }

  ngOnInit() {
    this.userHomeLink = `${sessionStorage.getItem("username")}/home`
    this.getAllCategories()
    this.initAddForm()

  }

  private initAddForm() {
    this.categoryForm = new FormGroup({
      id: new FormControl(null),
      name: new FormControl(null, [Validators.required, Validators.minLength(3), Validators.maxLength(30)]),
      bookmarks: new FormControl(null),
      userID: new FormControl(null),
      display: new FormControl(null)

    })
  }
}
