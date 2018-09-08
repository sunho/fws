import { Component, OnInit, Input, EventEmitter, Output } from '@angular/core';
import { NgForm, FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AuthService, DUPLICATE } from '../../services/auth.service';

@Component({
  selector: 'app-invite-form',
  templateUrl: './invite-form.component.html'
})
export class InviteFormComponent implements OnInit {
  @Input() key: string;
  @Input() username: string;
  @Output() OnSuccess = new EventEmitter<void>();

  formGroup: FormGroup;

  constructor(private formBuilder: FormBuilder, private authSerivce: AuthService) {}

  ngOnInit() {
    this.formGroup = this.formBuilder.group({
      nickname: ['', Validators.required],
      password: ['', Validators.required]
    });
  }

  onSubmit(f: NgForm) {
    if (f.valid) {
      this.authSerivce.register(this.key, this.username, f.value.nickname, f.value.password).subscribe(
        data => {
          this.OnSuccess.emit();
        },
        error => {
          if (error === DUPLICATE) {
            // this.wrongNickname = true;
          } else {
            // alert
          }
        }
      );
    }
  }
}
