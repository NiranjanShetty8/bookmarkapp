import { Component, OnInit } from '@angular/core';
import { UserService, IUser } from 'src/app/services/user-service/user.service';
import { FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'bookmarkapp-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})

export class RegisterComponent implements OnInit {
  registerForm: FormGroup
  loading: boolean
  user: IUser
  successMessage: string
  errorMessage: string
  password: string
  confirmPassword: string
  errorOccurred: boolean
  match: boolean
  errorString: string

  constructor(private _service: UserService) { }

  registerUser() {
    if (!this.passwordMatch()) {
      this.errorOccurred = true
      this.errorMessage = "Passwords do not match."
      return
    }

    this.loading = true
    this.user = this.registerForm.value
    this._service.register(this.user).subscribe((data: string) => {
      this.loading = false
      this.errorOccurred = false
      this.successMessage = `Hello ${this.user.username}, Your account has been created with ID:${data}
      Please Login to proceed.`
      this.registerForm.reset()

    }, (error) => {
      this.loading = false
      this.errorOccurred = true
      this.errorString = error.toString()
      if (this.errorString.startsWith("Error 1062", 0)) {
        this.errorMessage = "Username Already Exists. Try a different name."
        return
      }
      this.errorMessage = "Check Internet connection"

    })
  }

  passwordMatch(): boolean {
    this.password = this.registerForm.controls['password'].value
    this.confirmPassword = this.registerForm.controls['confirmPassword'].value
    if (this.password === this.confirmPassword) {
      this.match = true
      return true
    }
    this.match = false
    return false
  }

  ngOnInit() {
    this.loading = false
    this.initForm()
    this.passwordMatch()
  }

  private initForm() {
    this.registerForm = new FormGroup({
      username: new FormControl(null,
        [Validators.required, Validators.minLength(5), Validators.maxLength(30)]),
      password: new FormControl(null, [Validators.required, Validators.minLength(8)]),
      confirmPassword: new FormControl(null, [Validators.required])
    });
  }

}
