import { Component, OnInit } from '@angular/core';
import { CategoryService, ICategory } from 'src/app/services/category-service/category.service';
import { of } from 'rxjs';
import { TestBed } from '@angular/core/testing';

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

  constructor(private _service: CategoryService) { }

  getAllCategories() {
    console.log("called")
    this._service.getAllCategories().subscribe((data: ICategory[]) => {
      this.errorOccured = false
      console.log("data", data)
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

  addCategory(category: ICategory) {
    this._service.addCategory(category).subscribe((data: string) => {
      this.errorOccured = false
      console.log(data)
      this.successMessage = data

    }, (error) => {
      this.errorOccured = true
      this.errorMessage = error
      console.log(error)
    })
  }

  deleteCategory(categoryID: string) {
    this._service.deleteCategory(categoryID).subscribe((data: string) => {
      this.errorOccured = false
      this.successMessage = data
    }, (error) => {
      this.errorOccured = true
      this.errorMessage = error
    });
  }

  updateCategory(category: ICategory) {
    this._service.updateCategory(category).subscribe((data: string) => {
      this.errorOccured = false
      this.successMessage = data
    }, (error) => {
      this.errorOccured = true
      this.errorMessage = error
    });
  }

  getSpecificCategories(event: any) {
    let name = event.target.value
    if (name = "") {
      for (let category of this.categories) {
        category.display = true
      }
      return
    }
    for (let category of this.categories) {
      if (category.name.indexOf(event.target.value) == -1) {
        category.display = false
      } else {
        category.display = true
      }
    }
  }

  ngOnInit() {
    this.getAllCategories()
  }

}
