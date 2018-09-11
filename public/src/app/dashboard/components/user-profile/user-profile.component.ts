import { DropdownItem } from './../../../core/components/form-dropdown/form-dropdown.component';
import { Component, OnInit } from '@angular/core';

@Component({
    selector: 'app-user-profile',
    templateUrl: './user-profile.component.html'
})
export class UserProfileComponent implements OnInit {
    items: DropdownItem[];
    constructor() { }

    ngOnInit(): void {
        this.items = [
            {
                title: 'Change Password',
                func: null,
            },
            {
                title: 'Log Out',
                func: null,
            },
        ];
    }
}
