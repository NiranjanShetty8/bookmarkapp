import { Component, OnInit } from '@angular/core';
import { IUser, UserService } from 'src/app/services/user-service/user.service';
import { AdminService } from 'src/app/services/admin-service/admin.service';

@Component({
  selector: 'app-admin-home',
  templateUrl: './admin-home.component.html',
  styleUrls: ['./admin-home.component.css']
})
export class AdminHomeComponent implements OnInit {
  allUsers: IUser[]
  loading: boolean


  constructor(private adminService: AdminService, private userService: UserService) { }

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
    this.loading = true
    this.adminService.getAllusers().subscribe((data) => {
      this.allUsers = data
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
    })
  }

  ngOnInit(): void {
    this.getAllUsers()
  }

}
