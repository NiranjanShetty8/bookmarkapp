import { Component, OnInit } from '@angular/core';
import { CategoryService, ICategory } from 'src/app/services/category-service/category.service';
import { NgbModal, ModalDismissReasons } from '@ng-bootstrap/ng-bootstrap';
import { FormGroup, FormControl, Validators } from '@angular/forms';

@Component({
  selector: 'bookmarkapp-category',
  templateUrl: './category.component.html',
  styleUrls: ['./category.component.css']
})
export class CategoryComponent implements OnInit {
  categories: ICategory[]
  errorOccured: boolean
  errorMessage: string
  successMessage: string
  category: ICategory
  categoryForm: FormGroup
  formName: string
  formOperationName: string

  constructor(private _service: CategoryService, private modalService: NgbModal) { }
  closeResult = '';


  getAllCategories() {
    this._service.getAllCategories().subscribe((data: ICategory[]) => {
      this.errorOccured = false
      this.categories = data
      this.assignDisplayValue()
    }, (error) => {
      this.errorOccured = true
      this.errorMessage = error
    });
  }

  assignDisplayValue() {
    for (let category of this.categories) {
      category.display = true
    }
  }

  getCategoryByName(categoryName: string) {
    this._service.getCategoryByName(categoryName).subscribe((data: ICategory) => {
      this.errorOccured = false
      this.category = data
    }, (error) => {
      this.errorOccured = true
      this.errorMessage = error
    });
  }

  getCategoryByID(categoryID: string) {
    this._service.getCategoryByID(categoryID).subscribe((data: ICategory) => {
      this.errorOccured = false
      this.category = data
    }, (error) => {
      this.errorOccured = true
      this.errorMessage = error
    });
  }

  addCategory() {
    this.category = this.categoryForm.value
    this._service.addCategory(this.category).subscribe((data: string) => {
      alert(`Category added with ID: ${data}`)
      location.reload()
    }, (error) => {
      alert(error)
    })
  }

  deleteCategory(categoryID: string) {
    if (!confirm("Warning! Deleting a Category will delete all the bookmarks in it.\
    \nAre you sure you want to delete it?")) {
      return
    }
    this._service.deleteCategory(categoryID).subscribe((data: string) => {
      this.errorOccured = false
      this.successMessage = data
      alert(data)
      location.reload()
    }, (error) => {
      this.errorOccured = true
      this.errorMessage = error
      alert(error)
    });
  }

  updateCategory(category: ICategory) {
    console.log(category)
    this.category = this.categoryForm.value
    if (this.category.name === category.name) {
      alert("Name has not been changed.")
      return
    }
    this._service.updateCategory(this.category).subscribe((data: string) => {
      alert("Category Updated")
      location.reload()
    }, (error) => {
      // if (error.toString() == "" )
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


  openModal(content: any, category?: ICategory) {
    if (category) {
      this.formName = "Edit Category";
      this.formOperationName = "Save Changes"
      this.categoryForm.setValue(category)
      this.modalService.open(content, { ariaLabelledBy: 'category-modal' }).result.then((result) => {
        this.updateCategory(category)
      }, () => void 0);
      return
    }
    this.formName = "Add Category";
    this.formOperationName = "Add Category"
    this.categoryForm.reset()
    this.modalService.open(content, { ariaLabelledBy: 'category-modal' }).result.then((result) => {
      this.addCategory()
    }, () => void 0);
  }

  ngOnInit() {
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
