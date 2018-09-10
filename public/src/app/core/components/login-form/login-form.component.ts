import { Component, OnInit, EventEmitter, Output } from '@angular/core';
import { NgForm, FormGroup, FormBuilder, Validators } from '@angular/forms';

import {
  AuthService,
  WRONG_CRED,
  NOT_FOUND,
} from '../../services/auth.service';

@Component({
  selector: 'app-login-form',
  templateUrl: './login-form.component.html',
})
export class LoginFormComponent implements OnInit {
  @Output()
  OnSuccess = new EventEmitter<void>();
  formGroup: FormGroup;

  constructor(
    private formBuilder: FormBuilder,
    private authSerivce: AuthService
  ) {}

  ngOnInit(): void {
    this.formGroup = this.formBuilder.group({
      username: ['', Validators.required],
      password: ['', Validators.required],
    });
  }

  onSubmit(f: NgForm): void {
    if (f.valid) {
      this.authSerivce.login(f.value.username, f.value.password).subscribe(
        _ => {
          this.OnSuccess.emit();
        },
        error => {
          if (error === NOT_FOUND) {
            // this.wrongUsername = true;
          } else if (error === WRONG_CRED) {
            // this.wrongPassword = true;
          } else {
            // alert
          }
        }
      );
    }
  }
}
