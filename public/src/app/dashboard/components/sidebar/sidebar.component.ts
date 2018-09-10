import { routeName } from './../../dashboard-routing.module';
import { SideBarItem } from './../../models/sidebar';
import { Component, OnInit, Input } from '@angular/core';

@Component({
    selector: 'app-sidebar',
    templateUrl: './sidebar.component.html'
})
export class SideBarComponent implements OnInit {
    items: SideBarItem[] = [
        { title: 'home', icon: 'home', href: `/${routeName}`},
        { title: 'build', icon: 'hammer', href: `/${routeName}/build`},
        { title: 'volumes', icon: 'hdd', href: `/${routeName}/volumes`},
        { title: 'configs', icon: 'document', href: `/${routeName}/configs`},
    ];
    constructor() { }

    ngOnInit(): void { }
}
