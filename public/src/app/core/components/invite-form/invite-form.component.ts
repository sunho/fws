import { Component, OnInit, Input, EventEmitter, Output } from '@angular/core';
import { NgForm } from '@angular/forms';

import { AuthService, DUPLICATE } from '../../services/auth.service';

@Component({
  selector: 'app-invite-form',
  templateUrl: './invite-form.component.html',
  styleUrls: ['./invite-form.component.scss']
})
export class InviteFormComponent implements OnInit {
  @Input() key: string;
  @Input() username: string;
  @Output() OnSuccess = new EventEmitter<void>();

  wrongNickname: boolean;

  constructor(private authSerivce: AuthService) {
  }

  ngOnInit() {
  }

  onSubmit(f: NgForm) {
    if (f.valid) {
      this.wrongNickname = false;
      this.authSerivce.register(this.key, this.username, f.value.nickname, f.value.password).subscribe(
        data => {
          this.OnSuccess.emit();
        },
        error => {
          if (error === DUPLICATE) {
            this.wrongNickname = true;
          } else {
            // alert
          }
        }
      );
    }
  }
}
