import { Component, OnInit } from '@angular/core';
import { UserService, IUser } from 'src/app/services/user-service/user.service';
import { FormGroup, FormControl, Validators } from '@angular/forms';

@Component({
  selector: 'bookmarkapp-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  loginForm: FormGroup
  loading: boolean
  user: IUser
  errorOccurred: boolean
  errorMessage: string
  errorString: string

  constructor(private _service: UserService) { }

  userLogin() {
    this.loading = true
    this.user = this.loginForm.value
    this._service.login(this.user).subscribe((data: IUser) => {
      this.errorOccurred = false
      this.loading = false
      sessionStorage.setItem('userid', data.id)
      sessionStorage.setItem('token', data.token)
      this.loginForm.reset()

    }, (error) => {
      this.loading = false
      this.errorOccurred = true
      this.errorString = error.toString()

      if (this.errorString.startsWith("mismatch: Incorrect Username", 0)) {
        this.errorMessage = "Incorrect Username. New User? Register first!"
        return
      }
      this.errorMessage = error

    })
  }

  ngOnInit() {
    this.loading = false
    this.initForm()
  }

  private initForm() {
    this.loginForm = new FormGroup({
      username: new FormControl(null,
        [Validators.required, Validators.minLength(5), Validators.maxLength(25)]),
      password: new FormControl(null, [Validators.required, Validators.minLength(8)])
    });
  }

}
