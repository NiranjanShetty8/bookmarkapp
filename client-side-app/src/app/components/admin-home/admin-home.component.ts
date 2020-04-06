import { Component, OnInit } from '@angular/core';
import { IUser, UserService } from 'src/app/services/user-service/user.service';
import { AdminService } from 'src/app/services/admin-service/admin.service';
import { IBookmark } from 'src/app/services/bookmark-service/bookmark.service';

@Component({
  selector: 'app-admin-home',
  templateUrl: './admin-home.component.html',
  styleUrls: ['./admin-home.component.css']
})
export class AdminHomeComponent implements OnInit {
  allUsers: IUser[]
  allBookmarks: IBookmark[]
  loading: boolean
  displaySingleUser: boolean


  constructor(private adminService: AdminService) { }

  updateUser() {
    this.loading = true
    // this.userService.
  }

  unlockAccount(user: IUser) {
    this.loading = true
    this.adminService.unlockUser(user).subscribe((data) => {
      this.loading = false
      alert("Account Unlocked")
    }, (error) => {
      this.loading = false
      alert(error)
    })
  }

  getAllUsers() {
    this.displaySingleUser = false
    this.allBookmarks = []
    this.loading = true
    this.adminService.getAllusers().subscribe((data) => {
      this.allUsers = data
      for (let user of this.allUsers) {
        user.display = true
      }
      this.loading = false
    }, (error) => {
      this.loading = false
      alert(error)
    });
  }

  deleteUser(userID: string) {
    if (!confirm("Are you sure you want to delete the user?")) {
      return
    }
    this.loading = true
    this.adminService.deleteUser(userID).subscribe((data) => {
      alert(data)
      this.getAllUsers()
      this.loading = false
    }, (error) => {
      this.loading = false
      alert(error)
      this.getAllUsers()
    })
  }

  getAllBookmarksOfUser(user: IUser) {
    this.displaySingleUser = true
    for (let aUser of this.allUsers) {
      if (aUser.username === user.username) {
        console.log(aUser)
        for (let category of aUser.categories) {
          for (let bookmark of category.bookmarks) {
            console.log(bookmark)
            bookmark.display = true
            bookmark.categoryName = category.name
            this.allBookmarks.push(bookmark)
          }
        }
      }
    }
    console.log(this.allBookmarks)
  }

  getSpecificUser(event: any) {
    let name: string = event.target.value
    if (name == "") {
      for (let user of this.allUsers) {
        user.display = true
      }
      return
    }
    for (let user of this.allUsers) {
      let actualName = user.username.toLowerCase()
      if (actualName.indexOf(name) == -1) {
        user.display = false
      } else {
        user.display = true
      }
    }
  }

  ngOnInit(): void {
    this.getAllUsers()
  }

}
