import { AuthService } from './../../../core/services/auth.service';
import { DropdownItem } from './../../../core/components/form-dropdown/form-dropdown.component';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { PopupService } from '../../../core/services/popup.service';

@Component({
    selector: 'app-user-profile',
    templateUrl: './user-profile.component.html'
})
export class UserProfileComponent implements OnInit {
    items: DropdownItem[];
    constructor(private authService: AuthService, private popupService: PopupService, private router: Router) { }

    ngOnInit(): void {
        this.items = [
            {
                title: 'Change Password',
                func: this.changePassword.bind(this),
            },
            {
                title: 'Log Out',
                func: this.logout.bind(this),
            },
        ];
    }

    changePassword(t: string): void {
      this.popupService.createMsg('unimplemented');
    }

    logout(t: string): void {
        this.authService.logout().subscribe(
            _ => {
                this.router.navigate(['/']);
            },
            error => {
                this.popupService.createMsg(`unknown error(${error})`);
            }
        );
    }
}
