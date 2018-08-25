import { Component, OnInit, EventEmitter, Output } from '@angular/core';
import { NgForm } from '@angular/forms';

import { AuthService, WRONG_CRED, NOT_FOUND } from '../../services/auth.service';

@Component({
  selector: 'app-login-form',
  templateUrl: './login-form.component.html',
  styleUrls: ['./login-form.component.scss']
})
export class LoginFormComponent implements OnInit {
  @Output() OnSuccess = new EventEmitter<void>();

  wrongUsername: boolean;
  wrongPassword: boolean;

  constructor(private authSerivce: AuthService) {
  }

  ngOnInit() {
  }

  onSubmit(f: NgForm) {
    if (f.valid) {
      this.wrongUsername = false;
      this.wrongPassword = false;
      this.authSerivce.login(f.value.username, f.value.password).subscribe(
        data => {
          this.OnSuccess.emit();
        },
        error => {
          if (error === NOT_FOUND) {
            this.wrongUsername = true;
          } else if (error === WRONG_CRED) {
            this.wrongPassword = true;
          } else {
            // alert
          }
        }
      );
    }
  }
}
