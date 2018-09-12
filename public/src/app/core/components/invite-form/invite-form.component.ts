import { PopupService } from './../../services/popup.service';
import { Component, OnInit, Input, EventEmitter, Output } from '@angular/core';
import { NgForm, FormGroup, FormBuilder, Validators } from '@angular/forms';

import { AuthService, DUPLICATE } from '../../services/auth.service';
import { STRINGS } from '../../../../locale/strings';

@Component({
  selector: 'app-invite-form',
  templateUrl: './invite-form.component.html',
})
export class InviteFormComponent implements OnInit {
  @Input()
  key: string;
  @Input()
  username: string;
  @Output()
  OnSuccess = new EventEmitter<void>();

  formGroup: FormGroup;

  constructor(
    private formBuilder: FormBuilder,
    private authSerivce: AuthService,
    private popupService: PopupService
  ) {}

  ngOnInit(): void {
    this.formGroup = this.formBuilder.group({
      nickname: ['', Validators.required],
      password: ['', Validators.required],
    });
  }

  onSubmit(f: NgForm): void {
    if (f.valid) {
      this.authSerivce
        .register(this.key, this.username, f.value.nickname, f.value.password)
        .subscribe(
          _ => {
            this.OnSuccess.emit();
          },
          error => {
            if (error === DUPLICATE) {
              this.popupService.createMsg(STRINGS.NICKNAME_IN_USE);
            } else {
              this.popupService.createMsg(`${STRINGS.UNKNOWN_ERROR} (${error})`);
            }
          }
        );
    }
  }
}
