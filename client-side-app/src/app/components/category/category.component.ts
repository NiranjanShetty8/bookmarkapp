import { Component, OnInit } from '@angular/core';
import { CategoryService, ICategory } from 'src/app/services/category-service/category.service';

@Component({
  selector: 'bookmarkapp-category',
  templateUrl: './category.component.html',
  styleUrls: ['./category.component.css']
})
export class CategoryComponent implements OnInit {
  categories: ICategory[]
  errorOccured: boolean
  errorMessage: string
  category: ICategory

  constructor(private _service: CategoryService, ) { }

  getAllCategories() {
    console.log("called")
    this._service.getAllCategories().subscribe((data: ICategory[]) => {
      this.errorOccured = false
      console.log("data", data)
      this.categories = data
    }, (error) => {
      console.log("in 1st error:", error)
      this.errorOccured = true
      this.errorMessage = error

    });
  }


  addCategory() {
    this.category = { name: "angular", userID: sessionStorage.getItem('userid') }
    console.log(this.category)
    this._service.addCategory(this.category).subscribe((data: string) => {
      console.log(data)

    }, (error) => {
      console.log(error)
    })
  }

  ngOnInit() {
    // this.addCategory()
    this.getAllCategories()
  }

}
